// Package chain holds the go-ethereum helpers the CLI and the tests share:
// loading a key, building transactors, sending ETH and finding an account's
// transactions.
package chain

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// Backend is what the helpers need from a node. Both *ethclient.Client and
// the simulated backend's client implement it.
type Backend interface {
	bind.ContractBackend
	bind.DeployBackend
	ChainID(ctx context.Context) (*big.Int, error)
	BalanceAt(ctx context.Context, account common.Address, block *big.Int) (*big.Int, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
}

// ParseKey reads a hex private key, with or without a 0x prefix. The error
// never contains the key.
func ParseKey(hexKey string) (*ecdsa.PrivateKey, error) {
	key, err := crypto.HexToECDSA(strings.TrimPrefix(strings.TrimSpace(hexKey), "0x"))
	if err != nil {
		return nil, errors.New("invalid private key: expected 64 hex characters")
	}
	return key, nil
}

// Address returns the account address of key.
func Address(key *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(key.PublicKey)
}

// Transactor builds transaction options that sign with key for the backend's
// chain ID.
func Transactor(ctx context.Context, b Backend, key *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	chainID, err := b.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("reading chain ID: %w", err)
	}
	opts, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		return nil, err
	}
	opts.Context = ctx
	return opts, nil
}

// SendETH sends wei from key's account to `to` as an EIP-1559 transaction and
// returns it once it is accepted by the node (not yet mined).
func SendETH(ctx context.Context, b Backend, key *ecdsa.PrivateKey, to common.Address, wei *big.Int) (*types.Transaction, error) {
	from := Address(key)
	chainID, err := b.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("reading chain ID: %w", err)
	}
	nonce, err := b.PendingNonceAt(ctx, from)
	if err != nil {
		return nil, fmt.Errorf("reading nonce: %w", err)
	}
	tip, err := b.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, fmt.Errorf("suggesting gas tip: %w", err)
	}
	head, err := b.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("reading latest header: %w", err)
	}
	baseFee := head.BaseFee
	if baseFee == nil {
		baseFee = new(big.Int)
	}
	feeCap := new(big.Int).Add(tip, new(big.Int).Mul(baseFee, big.NewInt(2)))
	gas, err := b.EstimateGas(ctx, ethereum.CallMsg{From: from, To: &to, Value: wei})
	if err != nil {
		return nil, fmt.Errorf("estimating gas: %w", err)
	}

	tx, err := types.SignNewTx(key, types.LatestSignerForChainID(chainID), &types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: tip,
		GasFeeCap: feeCap,
		Gas:       gas,
		To:        &to,
		Value:     wei,
	})
	if err != nil {
		return nil, err
	}
	if err := b.SendTransaction(ctx, tx); err != nil {
		return nil, fmt.Errorf("sending transaction: %w", err)
	}
	return tx, nil
}

// TxInfo is one transaction involving an account.
type TxInfo struct {
	Hash  common.Hash
	Block uint64
	From  common.Address
	To    *common.Address // nil for contract creation
	Value *big.Int
}

// TransactionsOf scans blocks first..last (inclusive) for transactions sent
// from or to account. It reads every block, so keep the range small on a real
// network; an indexer is the right tool for full history.
func TransactionsOf(ctx context.Context, b Backend, account common.Address, first, last uint64) ([]TxInfo, error) {
	chainID, err := b.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("reading chain ID: %w", err)
	}
	signer := types.LatestSignerForChainID(chainID)

	var found []TxInfo
	for n := first; n <= last; n++ {
		block, err := b.BlockByNumber(ctx, new(big.Int).SetUint64(n))
		if err != nil {
			return nil, fmt.Errorf("reading block %d: %w", n, err)
		}
		for _, tx := range block.Transactions() {
			from, err := types.Sender(signer, tx)
			if err != nil {
				return nil, fmt.Errorf("recovering sender of %s: %w", tx.Hash(), err)
			}
			if from != account && (tx.To() == nil || *tx.To() != account) {
				continue
			}
			found = append(found, TxInfo{Hash: tx.Hash(), Block: n, From: from, To: tx.To(), Value: tx.Value()})
		}
	}
	return found, nil
}

// Ether converts whole ETH to wei.
func Ether(eth int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(eth), big.NewInt(1e18))
}
