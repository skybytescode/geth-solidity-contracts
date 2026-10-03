package bindings_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/skybytescode/geth-solidity-contracts/bindings"
	"github.com/skybytescode/geth-solidity-contracts/chain"
	"github.com/skybytescode/geth-solidity-contracts/internal/testchain"
)

// Every test deploys its contracts on a fresh in-process chain, so no node is
// needed. Attacks must fail and leave the state unchanged.

var call = &bind.CallOpts{Context: context.Background()}

func TestStorage(t *testing.T) {
	c := testchain.New(t)
	_, tx, storage, err := bindings.DeployStorage(c.Owner.Opts, c.Client)
	c.Succeeded(t, tx, err)

	tx, err = storage.Store(c.Alice.Opts, big.NewInt(42))
	receipt := c.Succeeded(t, tx, err)

	value, err := storage.Retrieve(call)
	require.NoError(t, err)
	require.Equal(t, int64(42), value.Int64())

	event, err := storage.ParseValueChanged(*receipt.Logs[0])
	require.NoError(t, err)
	require.Equal(t, int64(42), event.NewValue.Int64())
}

func TestTokenERC20TransfersAndAllowances(t *testing.T) {
	c := testchain.New(t)
	_, tx, token, err := bindings.DeployTokenERC20(c.Owner.Opts, c.Client, big.NewInt(1000))
	c.Succeeded(t, tx, err)

	supply, err := token.TotalSupply(call)
	require.NoError(t, err)
	requireAmount(t, chain.Ether(1000), supply) // 1000 tokens with 18 decimals

	tx, err = token.Transfer(c.Owner.Opts, c.Alice.Addr, chain.Ether(100))
	c.Succeeded(t, tx, err)
	requireBalance(t, token, c.Alice.Addr, chain.Ether(100))

	// Bob cannot move Alice's tokens until she approves him, and only up to the allowance.
	_, err = token.TransferFrom(c.Bob.Opts, c.Alice.Addr, c.Bob.Addr, chain.Ether(1))
	require.Error(t, err, "no allowance yet")

	tx, err = token.Approve(c.Alice.Opts, c.Bob.Addr, chain.Ether(30))
	c.Succeeded(t, tx, err)
	tx, err = token.TransferFrom(c.Bob.Opts, c.Alice.Addr, c.Bob.Addr, chain.Ether(30))
	c.Succeeded(t, tx, err)
	_, err = token.TransferFrom(c.Bob.Opts, c.Alice.Addr, c.Bob.Addr, chain.Ether(1))
	require.Error(t, err, "the allowance is used up")

	requireBalance(t, token, c.Alice.Addr, chain.Ether(70))
	requireBalance(t, token, c.Bob.Addr, chain.Ether(30))

	_, err = token.Transfer(c.Alice.Opts, c.Bob.Addr, chain.Ether(71))
	require.Error(t, err, "more than her balance")
	_, err = token.Transfer(c.Alice.Opts, common.Address{}, chain.Ether(1))
	require.Error(t, err, "tokens sent to the zero address are lost")
}

func requireBalance(t *testing.T, token *bindings.TokenERC20, who common.Address, want *big.Int) {
	t.Helper()
	got, err := token.BalanceOf(call, who)
	require.NoError(t, err)
	requireAmount(t, want, got)
}

// The first NFT contract let anyone mint and transfer any token; transferFrom
// only checked that `from` owned it.
func TestMyNFTOwnershipIsEnforced(t *testing.T) {
	c := testchain.New(t)
	_, tx, nft, err := bindings.DeployMyNFT(c.Owner.Opts, c.Client, "ipfs://collection/", c.Owner.Addr)
	c.Succeeded(t, tx, err)

	_, err = nft.Mint(c.Bob.Opts, c.Bob.Addr)
	require.Error(t, err, "only the owner mints")

	tx, err = nft.Mint(c.Owner.Opts, c.Alice.Addr)
	c.Succeeded(t, tx, err)
	requireOwner(t, nft, 1, c.Alice.Addr)

	_, err = nft.TransferFrom(c.Bob.Opts, c.Alice.Addr, c.Bob.Addr, big.NewInt(1))
	require.Error(t, err, "Bob must not be able to take Alice's token")
	requireOwner(t, nft, 1, c.Alice.Addr)

	tx, err = nft.Approve(c.Alice.Opts, c.Bob.Addr, big.NewInt(1))
	c.Succeeded(t, tx, err)
	tx, err = nft.TransferFrom(c.Bob.Opts, c.Alice.Addr, c.Bob.Addr, big.NewInt(1))
	c.Succeeded(t, tx, err)
	requireOwner(t, nft, 1, c.Bob.Addr)

	_, err = nft.TransferFrom(c.Bob.Opts, c.Alice.Addr, c.Bob.Addr, big.NewInt(1))
	require.Error(t, err, "the approval is cleared once the token moves")

	uri, err := nft.TokenURI(call, big.NewInt(1))
	require.NoError(t, err)
	require.Equal(t, "ipfs://collection/1", uri)
	supply, err := nft.TotalSupply(call)
	require.NoError(t, err)
	require.Equal(t, int64(1), supply.Int64())
}

func requireOwner(t *testing.T, nft *bindings.MyNFT, id int64, want common.Address) {
	t.Helper()
	got, err := nft.OwnerOf(call, big.NewInt(id))
	require.NoError(t, err)
	require.Equal(t, want, got)
}

// The first faucet sent 42 ETH to any caller, again and again, until empty.
func TestEthSenderLimitsClaims(t *testing.T) {
	c := testchain.New(t)
	addr, tx, faucet, err := bindings.DeployEthSender(c.Owner.Opts, c.Client, chain.Ether(1), big.NewInt(3600))
	c.Succeeded(t, tx, err)
	tx, err = chain.SendETH(context.Background(), c.Client, c.Owner.Key, addr, chain.Ether(3))
	c.Succeeded(t, tx, err)

	tx, err = faucet.SendEth(c.Alice.Opts)
	c.Succeeded(t, tx, err)
	requireFaucetBalance(t, c, addr, chain.Ether(2))

	_, err = faucet.SendEth(c.Alice.Opts)
	require.Error(t, err, "a second claim inside the cooldown")
	requireFaucetBalance(t, c, addr, chain.Ether(2))

	require.NoError(t, c.Backend.AdjustTime(time.Hour+time.Second))
	c.Mine()
	tx, err = faucet.SendEth(c.Alice.Opts)
	c.Succeeded(t, tx, err)

	sent, err := faucet.GetTotalSent(call, c.Alice.Addr)
	require.NoError(t, err)
	requireAmount(t, chain.Ether(2), sent)

	_, err = faucet.Withdraw(c.Bob.Opts, chain.Ether(1))
	require.Error(t, err, "only the owner withdraws")
	tx, err = faucet.Withdraw(c.Owner.Opts, chain.Ether(1))
	c.Succeeded(t, tx, err)
	requireFaucetBalance(t, c, addr, big.NewInt(0))

	_, err = faucet.SendEth(c.Bob.Opts)
	require.Error(t, err, "an empty faucet refuses claims")
}

func requireFaucetBalance(t *testing.T, c *testchain.Chain, faucet common.Address, want *big.Int) {
	t.Helper()
	got, err := c.Client.BalanceAt(context.Background(), faucet, nil)
	require.NoError(t, err)
	requireAmount(t, want, got)
}

func TestMyERC1155MintTransferAndBurn(t *testing.T) {
	c := testchain.New(t)
	_, tx, items, err := bindings.DeployMyERC1155Token(c.Owner.Opts, c.Client, c.Owner.Addr)
	c.Succeeded(t, tx, err)
	const gold, silver, sword = 0, 1, 2

	requireItems(t, items, c.Owner.Addr, gold, 1000)
	requireItems(t, items, c.Owner.Addr, sword, 10)

	ids := []*big.Int{big.NewInt(gold), big.NewInt(silver)}
	tx, err = items.MintBatch(c.Owner.Opts, c.Alice.Addr, ids, []*big.Int{big.NewInt(50), big.NewInt(20)}, nil)
	c.Succeeded(t, tx, err)
	requireItems(t, items, c.Alice.Addr, gold, 50)
	requireItems(t, items, c.Alice.Addr, silver, 20)

	_, err = items.Mint(c.Bob.Opts, c.Bob.Addr, big.NewInt(gold), big.NewInt(1), nil)
	require.Error(t, err, "only the owner mints")

	tx, err = items.SafeTransferFrom(c.Alice.Opts, c.Alice.Addr, c.Bob.Addr, big.NewInt(gold), big.NewInt(10), nil)
	c.Succeeded(t, tx, err)
	requireItems(t, items, c.Bob.Addr, gold, 10)

	_, err = items.Burn(c.Bob.Opts, c.Alice.Addr, big.NewInt(gold), big.NewInt(1))
	require.Error(t, err, "Bob cannot burn Alice's gold")

	tx, err = items.SetApprovalForAll(c.Alice.Opts, c.Bob.Addr, true)
	c.Succeeded(t, tx, err)
	tx, err = items.BurnBatch(c.Bob.Opts, c.Alice.Addr, ids, []*big.Int{big.NewInt(40), big.NewInt(20)})
	c.Succeeded(t, tx, err)
	requireItems(t, items, c.Alice.Addr, gold, 0)
	requireItems(t, items, c.Alice.Addr, silver, 0)

	uri, err := items.Uri(call, big.NewInt(gold))
	require.NoError(t, err)
	require.Contains(t, uri, "gold.png")
}

func requireItems(t *testing.T, items *bindings.MyERC1155Token, who common.Address, id, want int64) {
	t.Helper()
	got, err := items.BalanceOf(call, who, big.NewInt(id))
	require.NoError(t, err)
	require.Equal(t, want, got.Int64())
}

// requireAmount compares by value; big.Int can represent zero in two ways.
func requireAmount(t *testing.T, want, got *big.Int) {
	t.Helper()
	require.Zero(t, want.Cmp(got), "want %s, got %s", want, got)
}
