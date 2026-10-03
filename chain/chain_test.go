package chain_test

import (
	"context"
	"encoding/hex"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/skybytescode/geth-solidity-contracts/chain"
	"github.com/skybytescode/geth-solidity-contracts/internal/testchain"
)

func TestParseKeyNeverEchoesTheKey(t *testing.T) {
	bad := "0x" + strings.Repeat("zz", 32)
	_, err := chain.ParseKey(bad)
	require.Error(t, err)
	require.NotContains(t, err.Error(), "zz")

	c := testchain.New(t)
	key, err := chain.ParseKey("0x" + hex.EncodeToString(crypto.FromECDSA(c.Owner.Key)))
	require.NoError(t, err)
	require.Equal(t, c.Owner.Addr, chain.Address(key))
}

func TestSendETH(t *testing.T) {
	c := testchain.New(t)
	ctx := context.Background()
	before, err := c.Client.BalanceAt(ctx, c.Alice.Addr, nil)
	require.NoError(t, err)

	tx, err := chain.SendETH(ctx, c.Client, c.Owner.Key, c.Alice.Addr, chain.Ether(5))
	c.Succeeded(t, tx, err)

	after, err := c.Client.BalanceAt(ctx, c.Alice.Addr, nil)
	require.NoError(t, err)
	require.Zero(t, after.Sub(after, before).Cmp(chain.Ether(5)))
	require.Equal(t, uint8(2), tx.Type(), "EIP-1559 transaction")
}

// The original scan returned every transaction on the chain, whatever the account.
func TestTransactionsOfReturnsOnlyTheAccountsTransactions(t *testing.T) {
	c := testchain.New(t)
	ctx := context.Background()

	toAlice, err := chain.SendETH(ctx, c.Client, c.Owner.Key, c.Alice.Addr, chain.Ether(1))
	c.Succeeded(t, toAlice, err)
	toBob, err := chain.SendETH(ctx, c.Client, c.Owner.Key, c.Bob.Addr, chain.Ether(1))
	c.Succeeded(t, toBob, err)
	fromAlice, err := chain.SendETH(ctx, c.Client, c.Alice.Key, c.Bob.Addr, chain.Ether(1))
	c.Succeeded(t, fromAlice, err)

	head, err := c.Client.BlockNumber(ctx)
	require.NoError(t, err)
	txs, err := chain.TransactionsOf(ctx, c.Client, c.Alice.Addr, 0, head)
	require.NoError(t, err)

	require.Len(t, txs, 2)
	require.Equal(t, toAlice.Hash(), txs[0].Hash)
	require.Equal(t, c.Owner.Addr, txs[0].From)
	require.Equal(t, fromAlice.Hash(), txs[1].Hash)
	require.Equal(t, c.Alice.Addr, txs[1].From)
}
