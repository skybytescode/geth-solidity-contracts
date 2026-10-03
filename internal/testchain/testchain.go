// Package testchain starts an in-process go-ethereum chain with funded
// accounts for tests. Keys are generated per test and never leave memory.
package testchain

import (
	"context"
	"crypto/ecdsa"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"

	"github.com/skybytescode/geth-solidity-contracts/chain"
)

// Account is a funded test account.
type Account struct {
	Key  *ecdsa.PrivateKey
	Addr common.Address
	Opts *bind.TransactOpts
}

// Chain is a simulated chain; call Mine after sending transactions.
type Chain struct {
	Backend           *simulated.Backend
	Client            simulated.Client
	Owner, Alice, Bob Account
}

// New starts a chain where owner, alice and bob each hold 100 ETH.
func New(t *testing.T) *Chain {
	t.Helper()
	keys := make([]*ecdsa.PrivateKey, 3)
	alloc := types.GenesisAlloc{}
	for i := range keys {
		key, err := crypto.GenerateKey()
		if err != nil {
			t.Fatal(err)
		}
		keys[i] = key
		alloc[crypto.PubkeyToAddress(key.PublicKey)] = types.Account{Balance: chain.Ether(100)}
	}
	backend := simulated.NewBackend(alloc)
	t.Cleanup(func() { backend.Close() })

	c := &Chain{Backend: backend, Client: backend.Client()}
	accounts := []*Account{&c.Owner, &c.Alice, &c.Bob}
	for i, key := range keys {
		opts, err := chain.Transactor(context.Background(), c.Client, key)
		if err != nil {
			t.Fatal(err)
		}
		*accounts[i] = Account{Key: key, Addr: chain.Address(key), Opts: opts}
	}
	return c
}

// Mine seals the pending transactions into a block.
func (c *Chain) Mine() { c.Backend.Commit() }

// Succeeded mines tx and fails the test unless it succeeded.
func (c *Chain) Succeeded(t *testing.T, tx *types.Transaction, err error) *types.Receipt {
	t.Helper()
	if err != nil {
		t.Fatalf("sending transaction: %v", err)
	}
	c.Mine()
	receipt, err := c.Client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		t.Fatal(err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		t.Fatalf("transaction %s reverted", tx.Hash())
	}
	return receipt
}
