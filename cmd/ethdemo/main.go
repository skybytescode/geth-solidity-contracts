// Command ethdemo deploys the contracts in this repository to an Ethereum
// node and exercises them.
//
//	ETH_RPC_URL      node to talk to (default http://127.0.0.1:8545)
//	ETH_PRIVATE_KEY  hex key of a funded account, for commands that send transactions
//
// Usage:
//
//	ethdemo demo                       deploy all five contracts and walk through them
//	ethdemo balance <address>          show an account's ETH balance
//	ethdemo send <address> <eth>       send whole ETH from ETH_PRIVATE_KEY's account
//	ethdemo txs <address> [blocks]     list the account's transactions in the last blocks (default 100)
package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"

	"github.com/skybytescode/geth-solidity-contracts/bindings"
	"github.com/skybytescode/geth-solidity-contracts/chain"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		return errors.New("usage: ethdemo demo | balance <address> | send <address> <eth> | txs <address> [blocks]")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	url := os.Getenv("ETH_RPC_URL")
	if url == "" {
		url = "http://127.0.0.1:8545"
	}
	client, err := ethclient.DialContext(ctx, url)
	if err != nil {
		return fmt.Errorf("connecting to %s: %w", url, err)
	}
	defer client.Close()

	switch args[0] {
	case "balance":
		if len(args) != 2 {
			return errors.New("usage: ethdemo balance <address>")
		}
		addr, err := address(args[1])
		if err != nil {
			return err
		}
		bal, err := client.BalanceAt(ctx, addr, nil)
		if err != nil {
			return err
		}
		fmt.Printf("%s  %s ETH\n", addr.Hex(), eth(bal))
		return nil

	case "send":
		if len(args) != 3 {
			return errors.New("usage: ethdemo send <address> <eth>")
		}
		key, err := keyFromEnv()
		if err != nil {
			return err
		}
		to, err := address(args[1])
		if err != nil {
			return err
		}
		amount, err := strconv.ParseInt(args[2], 10, 64)
		if err != nil || amount <= 0 {
			return errors.New("the amount must be a positive whole number of ETH")
		}
		tx, err := chain.SendETH(ctx, client, key, to, chain.Ether(amount))
		if err != nil {
			return err
		}
		if _, err := mined(ctx, client, tx); err != nil {
			return err
		}
		fmt.Printf("sent %d ETH to %s in tx %s\n", amount, to.Hex(), tx.Hash().Hex())
		return nil

	case "txs":
		if len(args) < 2 || len(args) > 3 {
			return errors.New("usage: ethdemo txs <address> [blocks]")
		}
		addr, err := address(args[1])
		if err != nil {
			return err
		}
		window := uint64(100)
		if len(args) == 3 {
			if window, err = strconv.ParseUint(args[2], 10, 64); err != nil {
				return errors.New("blocks must be a positive number")
			}
		}
		head, err := client.BlockNumber(ctx)
		if err != nil {
			return err
		}
		first := uint64(0)
		if head > window {
			first = head - window
		}
		txs, err := chain.TransactionsOf(ctx, client, addr, first, head)
		if err != nil {
			return err
		}
		fmt.Printf("%d transactions of %s in blocks %d-%d\n", len(txs), addr.Hex(), first, head)
		for _, tx := range txs {
			to := "new contract"
			if tx.To != nil {
				to = short(tx.To.Hex())
			}
			fmt.Printf("  block %-4d %s  %s -> %-13s  %s ETH\n", tx.Block, short(tx.Hash.Hex()), short(tx.From.Hex()), to, eth(tx.Value))
		}
		return nil

	case "demo":
		key, err := keyFromEnv()
		if err != nil {
			return err
		}
		return demo(ctx, client, key)
	}
	return fmt.Errorf("unknown command %q", args[0])
}

// demo deploys every contract and walks through its main features, including
// the attacks the contracts must refuse. Alice is a throwaway account whose
// key only lives in memory.
func demo(ctx context.Context, client *ethclient.Client, ownerKey *ecdsa.PrivateKey) error {
	owner, err := chain.Transactor(ctx, client, ownerKey)
	if err != nil {
		return err
	}
	aliceKey, err := crypto.GenerateKey()
	if err != nil {
		return err
	}
	alice, err := chain.Transactor(ctx, client, aliceKey)
	if err != nil {
		return err
	}
	ownerAddr, aliceAddr := chain.Address(ownerKey), chain.Address(aliceKey)
	call := &bind.CallOpts{Context: ctx}
	step := func(format string, a ...any) { fmt.Printf("  "+format+"\n", a...) }
	refused := func(what string, err error) {
		if err == nil {
			step("!! %s was NOT refused", what)
			return
		}
		step("✓ %s refused", what)
	}

	fmt.Printf("owner %s, alice %s (throwaway)\n", short(ownerAddr.Hex()), short(aliceAddr.Hex()))
	tx, err := chain.SendETH(ctx, client, ownerKey, aliceAddr, chain.Ether(2))
	if err := check(ctx, client, tx, err); err != nil {
		return err
	}
	step("funded alice with 2 ETH for gas")

	fmt.Println("\nStorage")
	addr, tx, storage, err := bindings.DeployStorage(owner, client)
	if err := check(ctx, client, tx, err); err != nil {
		return err
	}
	step("deployed at %s", addr.Hex())
	tx, err = storage.Store(owner, big.NewInt(42))
	if err := check(ctx, client, tx, err); err != nil {
		return err
	}
	value, err := storage.Retrieve(call)
	if err != nil {
		return err
	}
	step("store(42) -> retrieve() = %s", value)

	fmt.Println("\nTokenERC20 (MTK)")
	addr, tx, token, err := bindings.DeployTokenERC20(owner, client, big.NewInt(1_000_000))
	if err := check(ctx, client, tx, err); err != nil {
		return err
	}
	step("deployed at %s with 1,000,000 MTK", addr.Hex())
	tx, err = token.Transfer(owner, aliceAddr, chain.Ether(250))
	if err := check(ctx, client, tx, err); err != nil {
		return err
	}
	bal, _ := token.BalanceOf(call, aliceAddr)
	step("transfer 250 MTK to alice -> alice holds %s MTK", eth(bal))
	_, err = token.TransferFrom(alice, ownerAddr, aliceAddr, chain.Ether(1))
	refused("alice spending the owner's tokens without an allowance", err)

	fmt.Println("\nMyNFT (ERC-721)")
	addr, tx, nft, err := bindings.DeployMyNFT(owner, client, "ipfs://collection/", ownerAddr)
	if err := check(ctx, client, tx, err); err != nil {
		return err
	}
	step("deployed at %s", addr.Hex())
	tx, err = nft.Mint(owner, ownerAddr)
	if err := check(ctx, client, tx, err); err != nil {
		return err
	}
	uri, _ := nft.TokenURI(call, big.NewInt(1))
	step("owner minted token #1 -> tokenURI %s", uri)
	_, err = nft.Mint(alice, aliceAddr)
	refused("alice minting", err)
	_, err = nft.TransferFrom(alice, ownerAddr, aliceAddr, big.NewInt(1))
	refused("alice taking token #1 without approval", err)
	tx, err = nft.Approve(owner, aliceAddr, big.NewInt(1))
	if err := check(ctx, client, tx, err); err != nil {
		return err
	}
	tx, err = nft.TransferFrom(alice, ownerAddr, aliceAddr, big.NewInt(1))
	if err := check(ctx, client, tx, err); err != nil {
		return err
	}
	holder, _ := nft.OwnerOf(call, big.NewInt(1))
	step("after the owner approves her, alice transfers it -> owner of #1 is alice: %v", holder == aliceAddr)

	fmt.Println("\nEthSender (faucet: 1 ETH per address per hour)")
	addr, tx, faucet, err := bindings.DeployEthSender(owner, client, chain.Ether(1), big.NewInt(3600))
	if err := check(ctx, client, tx, err); err != nil {
		return err
	}
	tx, err = chain.SendETH(ctx, client, ownerKey, addr, chain.Ether(5))
	if err := check(ctx, client, tx, err); err != nil {
		return err
	}
	step("deployed at %s and funded with 5 ETH", addr.Hex())
	tx, err = faucet.SendEth(alice)
	if err := check(ctx, client, tx, err); err != nil {
		return err
	}
	sent, _ := faucet.GetTotalSent(call, aliceAddr)
	step("alice claims -> received %s ETH in total", eth(sent))
	_, err = faucet.SendEth(alice)
	refused("alice claiming again within the hour", err)

	fmt.Println("\nMyERC1155Token (GOLD, SILVER, SWORD, SHIELD)")
	addr, tx, items, err := bindings.DeployMyERC1155Token(owner, client, ownerAddr)
	if err := check(ctx, client, tx, err); err != nil {
		return err
	}
	step("deployed at %s", addr.Hex())
	ids := []*big.Int{big.NewInt(0), big.NewInt(2)}
	tx, err = items.SafeBatchTransferFrom(owner, ownerAddr, aliceAddr, ids, []*big.Int{big.NewInt(100), big.NewInt(1)}, nil)
	if err := check(ctx, client, tx, err); err != nil {
		return err
	}
	gold, _ := items.BalanceOf(call, aliceAddr, big.NewInt(0))
	swords, _ := items.BalanceOf(call, aliceAddr, big.NewInt(2))
	step("batch transfer to alice -> %s GOLD, %s SWORD", gold, swords)
	_, err = items.Burn(alice, ownerAddr, big.NewInt(0), big.NewInt(1))
	refused("alice burning the owner's GOLD", err)
	tx, err = items.Burn(alice, aliceAddr, big.NewInt(0), big.NewInt(40))
	if err := check(ctx, client, tx, err); err != nil {
		return err
	}
	gold, _ = items.BalanceOf(call, aliceAddr, big.NewInt(0))
	step("alice burns 40 of her own GOLD -> %s left", gold)

	fmt.Println("\ndone")
	return nil
}

// check waits for a sent transaction and fails unless it succeeded.
func check(ctx context.Context, client *ethclient.Client, tx *types.Transaction, err error) error {
	if err != nil {
		return err
	}
	_, err = mined(ctx, client, tx)
	return err
}

func mined(ctx context.Context, client *ethclient.Client, tx *types.Transaction) (*types.Receipt, error) {
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		return nil, err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return nil, fmt.Errorf("transaction %s reverted", tx.Hash().Hex())
	}
	return receipt, nil
}

func keyFromEnv() (*ecdsa.PrivateKey, error) {
	v := os.Getenv("ETH_PRIVATE_KEY")
	if v == "" {
		return nil, errors.New("ETH_PRIVATE_KEY is not set")
	}
	return chain.ParseKey(v)
}

func address(s string) (common.Address, error) {
	if !common.IsHexAddress(s) {
		return common.Address{}, fmt.Errorf("%q is not an address", s)
	}
	return common.HexToAddress(s), nil
}

// eth formats wei as ETH (or as whole tokens with 18 decimals).
func eth(wei *big.Int) string {
	f := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether))
	return f.Text('f', 4)
}

func short(s string) string {
	if len(s) > 14 {
		return s[:8] + "…" + s[len(s)-4:]
	}
	return s
}
