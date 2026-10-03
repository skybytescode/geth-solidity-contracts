// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
	_ = time.Tick
	_ = context.Background
)

// MyNFTMetaData contains all meta data concerning the MyNFT contract.
var MyNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526001600755348015610014575f5ffd5b5060405161146c38038061146c83398101604081905261003391610158565b8060405180604001604052806005815260200164135e53919560da1b815250604051806040016040528060048152602001631353919560e21b815250815f908161007d919061029d565b50600161008a828261029d565b5050506001600160a01b0381166100ba57604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b6100c3816100d8565b5060086100d0838261029d565b505050610357565b600680546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b634e487b7160e01b5f52604160045260245ffd5b80516001600160a01b0381168114610153575f5ffd5b919050565b5f5f60408385031215610169575f5ffd5b82516001600160401b0381111561017e575f5ffd5b8301601f8101851361018e575f5ffd5b80516001600160401b038111156101a7576101a7610129565b604051601f8201601f19908116603f011681016001600160401b03811182821017156101d5576101d5610129565b6040528181528282016020018710156101ec575f5ffd5b8160208401602083015e5f602083830101528094505050506102106020840161013d565b90509250929050565b600181811c9082168061022d57607f821691505b60208210810361024b57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561029857805f5260205f20601f840160051c810160208510156102765750805b601f840160051c820191505b81811015610295575f8155600101610282565b50505b505050565b81516001600160401b038111156102b6576102b6610129565b6102ca816102c48454610219565b84610251565b6020601f8211600181146102fc575f83156102e55750848201515b5f19600385901b1c1916600184901b178455610295565b5f84815260208120601f198516915b8281101561032b578785015182556020948501946001909201910161030b565b508482101561034857868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b611108806103645f395ff3fe608060405234801561000f575f5ffd5b5060043610610111575f3560e01c806370a082311161009e578063a22cb4651161006e578063a22cb46514610228578063b88d4fde1461023b578063c87b56dd1461024e578063e985e9c514610261578063f2fde38b14610274575f5ffd5b806370a08231146101f4578063715018a6146102075780638da5cb5b1461020f57806395d89b4114610220575f5ffd5b806318160ddd116100e457806318160ddd1461019257806323b872dd146101a857806342842e0e146101bb5780636352211e146101ce5780636a627842146101e1575f5ffd5b806301ffc9a71461011557806306fdde031461013d578063081812fc14610152578063095ea7b31461017d575b5f5ffd5b610128610123366004610d7b565b610287565b60405190151581526020015b60405180910390f35b6101456102d8565b6040516101349190610dc4565b610165610160366004610dd6565b610367565b6040516001600160a01b039091168152602001610134565b61019061018b366004610e03565b61038e565b005b61019a61039d565b604051908152602001610134565b6101906101b6366004610e2b565b6103b2565b6101906101c9366004610e2b565b610440565b6101656101dc366004610dd6565b61045f565b61019a6101ef366004610e65565b610469565b61019a610202366004610e65565b610496565b6101906104db565b6006546001600160a01b0316610165565b6101456104ee565b610190610236366004610e7e565b6104fd565b610190610249366004610ecb565b610508565b61014561025c366004610dd6565b610520565b61012861026f366004610fa8565b610585565b610190610282366004610e65565b6105b2565b5f6001600160e01b031982166380ac58cd60e01b14806102b757506001600160e01b03198216635b5e139f60e01b145b806102d257506301ffc9a760e01b6001600160e01b03198316145b92915050565b60605f80546102e690610fd9565b80601f016020809104026020016040519081016040528092919081815260200182805461031290610fd9565b801561035d5780601f106103345761010080835404028352916020019161035d565b820191905f5260205f20905b81548152906001019060200180831161034057829003601f168201915b5050505050905090565b5f610371826105ef565b505f828152600460205260409020546001600160a01b03166102d2565b610399828233610627565b5050565b5f60016007546103ad9190611025565b905090565b6001600160a01b0382166103e057604051633250574960e11b81525f60048201526024015b60405180910390fd5b5f6103ec838333610634565b9050836001600160a01b0316816001600160a01b03161461043a576040516364283d7b60e01b81526001600160a01b03808616600483015260248201849052821660448201526064016103d7565b50505050565b61045a83838360405180602001604052805f815250610508565b505050565b5f6102d2826105ef565b5f610472610726565b60078054905f61048183611038565b9190505590506104918282610753565b919050565b5f6001600160a01b0382166104c0576040516322718ad960e21b81525f60048201526024016103d7565b506001600160a01b03165f9081526003602052604090205490565b6104e3610726565b6104ec5f61076c565b565b6060600180546102e690610fd9565b6103993383836107bd565b6105138484846103b2565b61043a3385858585610884565b606061052b826105ef565b505f6105356109ac565b90505f8151116105535760405180602001604052805f81525061057e565b8061055d846109bb565b60405160200161056e929190611067565b6040516020818303038152906040525b9392505050565b6001600160a01b039182165f90815260056020908152604080832093909416825291909152205460ff1690565b6105ba610726565b6001600160a01b0381166105e357604051631e4fbdf760e01b81525f60048201526024016103d7565b6105ec8161076c565b50565b5f818152600260205260408120546001600160a01b0316806102d257604051637e27328960e01b8152600481018490526024016103d7565b61045a8383836001610a4b565b5f828152600260205260408120546001600160a01b039081169083161561066057610660818486610b4f565b6001600160a01b0381161561069a5761067b5f855f5f610a4b565b6001600160a01b0381165f90815260036020526040902080545f190190555b6001600160a01b038516156106c8576001600160a01b0385165f908152600360205260409020805460010190555b5f8481526002602052604080822080546001600160a01b0319166001600160a01b0389811691821790925591518793918516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4949350505050565b6006546001600160a01b031633146104ec5760405163118cdaa760e01b81523360048201526024016103d7565b610399828260405180602001604052805f815250610bb3565b600680546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b6001600160a01b0383166107e65760405163a9fbf51f60e01b81525f60048201526024016103d7565b6001600160a01b03821661081857604051630b61174360e31b81526001600160a01b03831660048201526024016103d7565b6001600160a01b038381165f81815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6001600160a01b0383163b156109a557604051630a85bd0160e11b81526001600160a01b0384169063150b7a02906108c690889088908790879060040161107b565b6020604051808303815f875af1925050508015610900575060408051601f3d908101601f191682019092526108fd918101906110b7565b60015b610967573d80801561092d576040519150601f19603f3d011682016040523d82523d5f602084013e610932565b606091505b5080515f0361095f57604051633250574960e11b81526001600160a01b03851660048201526024016103d7565b805160208201fd5b6001600160e01b03198116630a85bd0160e11b146109a357604051633250574960e11b81526001600160a01b03851660048201526024016103d7565b505b5050505050565b6060600880546102e690610fd9565b60605f6109c783610bca565b60010190505f8167ffffffffffffffff8111156109e6576109e6610eb7565b6040519080825280601f01601f191660200182016040528015610a10576020820181803683370190505b5090508181016020015b5f19016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a8504945084610a1a57509392505050565b8080610a5f57506001600160a01b03821615155b15610b20575f610a6e846105ef565b90506001600160a01b03831615801590610a9a5750826001600160a01b0316816001600160a01b031614155b8015610aad5750610aab8184610585565b155b15610ad65760405163a9fbf51f60e01b81526001600160a01b03841660048201526024016103d7565b8115610b1e5783856001600160a01b0316826001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b50505f90815260046020526040902080546001600160a01b0319166001600160a01b0392909216919091179055565b610b5a838383610ca1565b61045a576001600160a01b038316610b8857604051637e27328960e01b8152600481018290526024016103d7565b60405163177e802f60e01b81526001600160a01b0383166004820152602481018290526044016103d7565b610bbd8383610d05565b61045a335f858585610884565b5f8072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b8310610c085772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef81000000008310610c34576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc100008310610c5257662386f26fc10000830492506010015b6305f5e1008310610c6a576305f5e100830492506008015b6127108310610c7e57612710830492506004015b60648310610c90576064830492506002015b600a83106102d25760010192915050565b5f6001600160a01b03831615801590610cfd5750826001600160a01b0316846001600160a01b03161480610cda5750610cda8484610585565b80610cfd57505f828152600460205260409020546001600160a01b038481169116145b949350505050565b6001600160a01b038216610d2e57604051633250574960e11b81525f60048201526024016103d7565b5f610d3a83835f610634565b90506001600160a01b0381161561045a576040516339e3563760e11b81525f60048201526024016103d7565b6001600160e01b0319811681146105ec575f5ffd5b5f60208284031215610d8b575f5ffd5b813561057e81610d66565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f61057e6020830184610d96565b5f60208284031215610de6575f5ffd5b5035919050565b80356001600160a01b0381168114610491575f5ffd5b5f5f60408385031215610e14575f5ffd5b610e1d83610ded565b946020939093013593505050565b5f5f5f60608486031215610e3d575f5ffd5b610e4684610ded565b9250610e5460208501610ded565b929592945050506040919091013590565b5f60208284031215610e75575f5ffd5b61057e82610ded565b5f5f60408385031215610e8f575f5ffd5b610e9883610ded565b915060208301358015158114610eac575f5ffd5b809150509250929050565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f5f60808587031215610ede575f5ffd5b610ee785610ded565b9350610ef560208601610ded565b925060408501359150606085013567ffffffffffffffff811115610f17575f5ffd5b8501601f81018713610f27575f5ffd5b803567ffffffffffffffff811115610f4157610f41610eb7565b604051601f8201601f19908116603f0116810167ffffffffffffffff81118282101715610f7057610f70610eb7565b604052818152828201602001891015610f87575f5ffd5b816020840160208301375f6020838301015280935050505092959194509250565b5f5f60408385031215610fb9575f5ffd5b610fc283610ded565b9150610fd060208401610ded565b90509250929050565b600181811c90821680610fed57607f821691505b60208210810361100b57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b818103818111156102d2576102d2611011565b5f6001820161104957611049611011565b5060010190565b5f81518060208401855e5f93019283525090919050565b5f610cfd6110758386611050565b84611050565b6001600160a01b03858116825284166020820152604081018390526080606082018190525f906110ad90830184610d96565b9695505050505050565b5f602082840312156110c7575f5ffd5b815161057e81610d6656fea2646970667358221220d433f18c1ab4954693a4fb3421b7778c33037d85ccad0d56376fd921367e8f7964736f6c634300081e0033",
}

// MyNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use MyNFTMetaData.ABI instead.
var MyNFTABI = MyNFTMetaData.ABI

// MyNFTBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MyNFTMetaData.Bin instead.
var MyNFTBin = MyNFTMetaData.Bin

// DeployMyNFT deploys a new Ethereum contract, binding an instance of MyNFT to it.
func DeployMyNFT(auth *bind.TransactOpts, backend bind.ContractBackend, baseURI_ string, initialOwner common.Address) (common.Address, *types.Transaction, *MyNFT, error) {
	parsed, err := MyNFTMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MyNFTBin), backend, baseURI_, initialOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyNFT{MyNFTCaller: MyNFTCaller{contract: contract}, MyNFTTransactor: MyNFTTransactor{contract: contract}, MyNFTFilterer: MyNFTFilterer{contract: contract}}, nil
}

// MyNFT is an auto generated Go binding around an Ethereum contract.
type MyNFT struct {
	MyNFTCaller     // Read-only binding to the contract
	MyNFTTransactor // Write-only binding to the contract
	MyNFTFilterer   // Log filterer for contract events
}

// MyNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type MyNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MyNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyNFTSession struct {
	Contract     *MyNFT            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyNFTCallerSession struct {
	Contract *MyNFTCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MyNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyNFTTransactorSession struct {
	Contract     *MyNFTTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type MyNFTRaw struct {
	Contract *MyNFT // Generic contract binding to access the raw methods on
}

// MyNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyNFTCallerRaw struct {
	Contract *MyNFTCaller // Generic read-only contract binding to access the raw methods on
}

// MyNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyNFTTransactorRaw struct {
	Contract *MyNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyNFT creates a new instance of MyNFT, bound to a specific deployed contract.
func NewMyNFT(address common.Address, backend bind.ContractBackend) (*MyNFT, error) {
	contract, err := bindMyNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyNFT{MyNFTCaller: MyNFTCaller{contract: contract}, MyNFTTransactor: MyNFTTransactor{contract: contract}, MyNFTFilterer: MyNFTFilterer{contract: contract}}, nil
}

// NewMyNFTCaller creates a new read-only instance of MyNFT, bound to a specific deployed contract.
func NewMyNFTCaller(address common.Address, caller bind.ContractCaller) (*MyNFTCaller, error) {
	contract, err := bindMyNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyNFTCaller{contract: contract}, nil
}

// NewMyNFTTransactor creates a new write-only instance of MyNFT, bound to a specific deployed contract.
func NewMyNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*MyNFTTransactor, error) {
	contract, err := bindMyNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyNFTTransactor{contract: contract}, nil
}

// NewMyNFTFilterer creates a new log filterer instance of MyNFT, bound to a specific deployed contract.
func NewMyNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*MyNFTFilterer, error) {
	contract, err := bindMyNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyNFTFilterer{contract: contract}, nil
}

// bindMyNFT binds a generic wrapper to an already deployed contract.
func bindMyNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MyNFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyNFT *MyNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyNFT.Contract.MyNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyNFT *MyNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyNFT.Contract.MyNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyNFT *MyNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyNFT.Contract.MyNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyNFT *MyNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyNFT *MyNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyNFT *MyNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyNFT.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MyNFT *MyNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MyNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MyNFT *MyNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MyNFT.Contract.BalanceOf(&_MyNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MyNFT *MyNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MyNFT.Contract.BalanceOf(&_MyNFT.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MyNFT *MyNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MyNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MyNFT *MyNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _MyNFT.Contract.GetApproved(&_MyNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MyNFT *MyNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _MyNFT.Contract.GetApproved(&_MyNFT.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MyNFT *MyNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _MyNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MyNFT *MyNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MyNFT.Contract.IsApprovedForAll(&_MyNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MyNFT *MyNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MyNFT.Contract.IsApprovedForAll(&_MyNFT.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyNFT *MyNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MyNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyNFT *MyNFTSession) Name() (string, error) {
	return _MyNFT.Contract.Name(&_MyNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyNFT *MyNFTCallerSession) Name() (string, error) {
	return _MyNFT.Contract.Name(&_MyNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyNFT *MyNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MyNFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyNFT *MyNFTSession) Owner() (common.Address, error) {
	return _MyNFT.Contract.Owner(&_MyNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyNFT *MyNFTCallerSession) Owner() (common.Address, error) {
	return _MyNFT.Contract.Owner(&_MyNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MyNFT *MyNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MyNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MyNFT *MyNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _MyNFT.Contract.OwnerOf(&_MyNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MyNFT *MyNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _MyNFT.Contract.OwnerOf(&_MyNFT.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MyNFT *MyNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MyNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MyNFT *MyNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MyNFT.Contract.SupportsInterface(&_MyNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MyNFT *MyNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MyNFT.Contract.SupportsInterface(&_MyNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyNFT *MyNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MyNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyNFT *MyNFTSession) Symbol() (string, error) {
	return _MyNFT.Contract.Symbol(&_MyNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyNFT *MyNFTCallerSession) Symbol() (string, error) {
	return _MyNFT.Contract.Symbol(&_MyNFT.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MyNFT *MyNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _MyNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MyNFT *MyNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _MyNFT.Contract.TokenURI(&_MyNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MyNFT *MyNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _MyNFT.Contract.TokenURI(&_MyNFT.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MyNFT *MyNFTCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyNFT.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MyNFT *MyNFTSession) TotalSupply() (*big.Int, error) {
	return _MyNFT.Contract.TotalSupply(&_MyNFT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MyNFT *MyNFTCallerSession) TotalSupply() (*big.Int, error) {
	return _MyNFT.Contract.TotalSupply(&_MyNFT.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MyNFT *MyNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MyNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MyNFT *MyNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MyNFT.Contract.Approve(&_MyNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MyNFT *MyNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MyNFT.Contract.Approve(&_MyNFT.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 tokenId)
func (_MyNFT *MyNFTTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _MyNFT.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 tokenId)
func (_MyNFT *MyNFTSession) Mint(to common.Address) (*types.Transaction, error) {
	return _MyNFT.Contract.Mint(&_MyNFT.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 tokenId)
func (_MyNFT *MyNFTTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _MyNFT.Contract.Mint(&_MyNFT.TransactOpts, to)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MyNFT *MyNFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyNFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MyNFT *MyNFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _MyNFT.Contract.RenounceOwnership(&_MyNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MyNFT *MyNFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MyNFT.Contract.RenounceOwnership(&_MyNFT.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MyNFT *MyNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MyNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MyNFT *MyNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MyNFT.Contract.SafeTransferFrom(&_MyNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MyNFT *MyNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MyNFT.Contract.SafeTransferFrom(&_MyNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_MyNFT *MyNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _MyNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_MyNFT *MyNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _MyNFT.Contract.SafeTransferFrom0(&_MyNFT.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_MyNFT *MyNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _MyNFT.Contract.SafeTransferFrom0(&_MyNFT.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MyNFT *MyNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _MyNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MyNFT *MyNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MyNFT.Contract.SetApprovalForAll(&_MyNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MyNFT *MyNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MyNFT.Contract.SetApprovalForAll(&_MyNFT.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MyNFT *MyNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MyNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MyNFT *MyNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MyNFT.Contract.TransferFrom(&_MyNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MyNFT *MyNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MyNFT.Contract.TransferFrom(&_MyNFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MyNFT *MyNFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MyNFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MyNFT *MyNFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MyNFT.Contract.TransferOwnership(&_MyNFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MyNFT *MyNFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MyNFT.Contract.TransferOwnership(&_MyNFT.TransactOpts, newOwner)
}

// MyNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MyNFT contract.
type MyNFTApprovalIterator struct {
	Event *MyNFTApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyNFTApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyNFTApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyNFTApproval represents a Approval event raised by the MyNFT contract.
type MyNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MyNFT *MyNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*MyNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MyNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MyNFTApprovalIterator{contract: _MyNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MyNFT *MyNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MyNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MyNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyNFTApproval)
				if err := _MyNFT.contract.UnpackLog(event, "Approval", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MyNFT *MyNFTFilterer) ParseApproval(log types.Log) (*MyNFTApproval, error) {
	event := new(MyNFTApproval)
	if err := _MyNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the MyNFT contract.
type MyNFTApprovalForAllIterator struct {
	Event *MyNFTApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyNFTApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyNFTApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyNFTApprovalForAll represents a ApprovalForAll event raised by the MyNFT contract.
type MyNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MyNFT *MyNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*MyNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MyNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MyNFTApprovalForAllIterator{contract: _MyNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MyNFT *MyNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MyNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MyNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyNFTApprovalForAll)
				if err := _MyNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MyNFT *MyNFTFilterer) ParseApprovalForAll(log types.Log) (*MyNFTApprovalForAll, error) {
	event := new(MyNFTApprovalForAll)
	if err := _MyNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyNFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MyNFT contract.
type MyNFTOwnershipTransferredIterator struct {
	Event *MyNFTOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyNFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyNFTOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyNFTOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyNFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyNFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyNFTOwnershipTransferred represents a OwnershipTransferred event raised by the MyNFT contract.
type MyNFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MyNFT *MyNFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MyNFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MyNFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MyNFTOwnershipTransferredIterator{contract: _MyNFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MyNFT *MyNFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MyNFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MyNFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyNFTOwnershipTransferred)
				if err := _MyNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MyNFT *MyNFTFilterer) ParseOwnershipTransferred(log types.Log) (*MyNFTOwnershipTransferred, error) {
	event := new(MyNFTOwnershipTransferred)
	if err := _MyNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MyNFT contract.
type MyNFTTransferIterator struct {
	Event *MyNFTTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyNFTTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyNFTTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyNFTTransfer represents a Transfer event raised by the MyNFT contract.
type MyNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MyNFT *MyNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*MyNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MyNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MyNFTTransferIterator{contract: _MyNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MyNFT *MyNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MyNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MyNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyNFTTransfer)
				if err := _MyNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MyNFT *MyNFTFilterer) ParseTransfer(log types.Log) (*MyNFTTransfer, error) {
	event := new(MyNFTTransfer)
	if err := _MyNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
