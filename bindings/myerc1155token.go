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

// MyERC1155TokenMetaData contains all meta data concerning the MyERC1155Token contract.
var MyERC1155TokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC1155InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valuesLength\",\"type\":\"uint256\"}],\"name\":\"ERC1155InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC1155MissingApprovalForAll\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHIELD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SILVER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SWORD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"burnBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mintBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161214d38038061214d83398101604081905261002e9161064f565b806040518060600160405280602981526020016121246029913961005181610119565b506001600160a01b03811661008057604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b61008981610129565b506100ac335f6103e860405180602001604052805f81525061017a60201b60201c565b6100cf3360016103e860405180602001604052805f81525061017a60201b60201c565b6100f1336002600a60405180602001604052805f81525061017a60201b60201c565b610113336003600a60405180602001604052805f81525061017a60201b60201c565b5061094d565b60026101258282610711565b5050565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b6001600160a01b0384166101a357604051632bfa23e760e11b81525f6004820152602401610077565b604080516001808252602082018690528183019081526060820185905260808201909252906101d65f87848487846101de565b505050505050565b6101ea86868686610237565b6001600160a01b038516156101d6573381156102135761020e818888888888610447565b61022e565b6020858101519085015161022b838a8a85858a610568565b50505b50505050505050565b80518251146102665781518151604051635b05999160e01b815260048101929092526024820152604401610077565b335f5b8351811015610368576020818102858101820151908501909101516001600160a01b0388161561031a575f828152602081815260408083206001600160a01b038c168452909152902054818110156102f4576040516303dee4c560e01b81526001600160a01b038a166004820152602481018290526044810183905260648101849052608401610077565b5f838152602081815260408083206001600160a01b038d16845290915290209082900390555b6001600160a01b0387161561035e575f828152602081815260408083206001600160a01b038b168452909152812080548392906103589084906107cb565b90915550505b5050600101610269565b5082516001036103e85760208301515f906020840151909150856001600160a01b0316876001600160a01b0316846001600160a01b03167fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f6285856040516103d9929190918252602082015260400190565b60405180910390a45050610440565b836001600160a01b0316856001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb868660405161043792919061082a565b60405180910390a45b5050505050565b6001600160a01b0384163b156101d65760405163bc197c8160e01b81526001600160a01b0385169063bc197c819061048b9089908990889088908890600401610885565b6020604051808303815f875af19250505080156104c5575060408051601f3d908101601f191682019092526104c2918101906108e2565b60015b61052c573d8080156104f2576040519150601f19603f3d011682016040523d82523d5f602084013e6104f7565b606091505b5080515f0361052457604051632bfa23e760e11b81526001600160a01b0386166004820152602401610077565b805160208201fd5b6001600160e01b0319811663bc197c8160e01b1461022e57604051632bfa23e760e11b81526001600160a01b0386166004820152602401610077565b6001600160a01b0384163b156101d65760405163f23a6e6160e01b81526001600160a01b0385169063f23a6e61906105ac9089908990889088908890600401610909565b6020604051808303815f875af19250505080156105e6575060408051601f3d908101601f191682019092526105e3918101906108e2565b60015b610613573d8080156104f2576040519150601f19603f3d011682016040523d82523d5f602084013e6104f7565b6001600160e01b0319811663f23a6e6160e01b1461022e57604051632bfa23e760e11b81526001600160a01b0386166004820152602401610077565b5f6020828403121561065f575f5ffd5b81516001600160a01b0381168114610675575f5ffd5b9392505050565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806106a457607f821691505b6020821081036106c257634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561070c57805f5260205f20601f840160051c810160208510156106ed5750805b601f840160051c820191505b81811015610440575f81556001016106f9565b505050565b81516001600160401b0381111561072a5761072a61067c565b61073e816107388454610690565b846106c8565b6020601f821160018114610770575f83156107595750848201515b5f19600385901b1c1916600184901b178455610440565b5f84815260208120601f198516915b8281101561079f578785015182556020948501946001909201910161077f565b50848210156107bc57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b808201808211156107ea57634e487b7160e01b5f52601160045260245ffd5b92915050565b5f8151808452602084019350602083015f5b82811015610820578151865260209586019590910190600101610802565b5093949350505050565b604081525f61083c60408301856107f0565b828103602084015261084e81856107f0565b95945050505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b6001600160a01b0386811682528516602082015260a0604082018190525f906108b0908301866107f0565b82810360608401526108c281866107f0565b905082810360808401526108d68185610857565b98975050505050505050565b5f602082840312156108f2575f5ffd5b81516001600160e01b031981168114610675575f5ffd5b6001600160a01b03868116825285166020820152604081018490526060810183905260a0608082018190525f9061094290830184610857565b979650505050505050565b6117ca8061095a5f395ff3fe608060405234801561000f575f5ffd5b506004361061011b575f3560e01c80636b20c454116100a9578063e3e55f081161006e578063e3e55f0814610243578063e985e9c51461024b578063f242432a1461025e578063f2fde38b14610271578063f5298aca14610284575f5ffd5b80636b20c454146101e7578063715018a6146101fa578063731133e9146102025780638da5cb5b14610215578063a22cb46514610230575f5ffd5b80631f7fdffa116100ef5780631f7fdffa146101905780632eb2c2d6146101a55780633e4bee38146101b85780634e1273f4146101bf5780635b2725ed146101df575f5ffd5b8062fdd58e1461011f57806301ffc9a7146101455780630e89341c1461016857806313dc989f14610188575b5f5ffd5b61013261012d366004610ecd565b610297565b6040519081526020015b60405180910390f35b610158610153366004610f0a565b6102be565b604051901515815260200161013c565b61017b610176366004610f2c565b61030d565b60405161013c9190610f71565b610132600281565b6101a361019e3660046110be565b6103b6565b005b6101a36101b336600461115a565b6103d0565b6101325f81565b6101d26101cd366004611206565b6103ef565b60405161013c9190611301565b610132600381565b6101a36101f5366004611313565b6104be565b6101a36104d7565b6101a3610210366004611385565b6104ea565b6003546040516001600160a01b03909116815260200161013c565b6101a361023e3660046113c9565b6104fe565b610132600181565b610158610259366004611402565b61050d565b6101a361026c366004611433565b61053a565b6101a361027f366004611486565b610550565b6101a361029236600461149f565b61058d565b5f818152602081815260408083206001600160a01b03861684529091529020545b92915050565b5f6001600160e01b03198216636cdb3d1360e11b14806102ee57506001600160e01b031982166303a24d0760e21b145b806102b857506301ffc9a760e01b6001600160e01b03198316146102b8565b606081610333576040518060800160405280605b8152602001611680605b913992915050565b6001820361035a576040518060800160405280605d81526020016116db605d913992915050565b60028203610381576040518060800160405280605c8152602001611624605c913992915050565b600382036103a8576040518060800160405280605d8152602001611738605d913992915050565b6102b8826105a1565b919050565b6103be610633565b6103ca84848484610660565b50505050565b6103db335b86610698565b6103e885858585856106f2565b5050505050565b606081518351146104255781518351604051635b05999160e01b8152600481019290925260248201526044015b60405180910390fd5b5f83516001600160401b0381111561043f5761043f610f83565b604051908082528060200260200182016040528015610468578160200160208202803683370190505b5090505f5b84518110156104b65760208082028601015161049190602080840287010151610297565b8282815181106104a3576104a36114cf565b602090810291909101015260010161046d565b509392505050565b6104c783610752565b6104d28383836107a1565b505050565b6104df610633565b6104e85f6107e6565b565b6104f2610633565b6103ca84848484610837565b61050933838361089b565b5050565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205460ff1690565b610543336103d5565b6103e88585858585610958565b610558610633565b6001600160a01b03811661058157604051631e4fbdf760e01b81525f600482015260240161041c565b61058a816107e6565b50565b61059683610752565b6104d28383836109e5565b6060600280546105b0906114e3565b80601f01602080910402602001604051908101604052809291908181526020018280546105dc906114e3565b80156106275780601f106105fe57610100808354040283529160200191610627565b820191905f5260205f20905b81548152906001019060200180831161060a57829003601f168201915b50505050509050919050565b6003546001600160a01b031633146104e85760405163118cdaa760e01b815233600482015260240161041c565b6001600160a01b03841661068957604051632bfa23e760e11b81525f600482015260240161041c565b6103ca5f858585856001610a48565b816001600160a01b0316816001600160a01b0316141580156106c157506106bf818361050d565b155b156105095760405163711bec9160e11b81526001600160a01b0380841660048301528216602482015260440161041c565b6001600160a01b03841661071b57604051632bfa23e760e11b81525f600482015260240161041c565b6001600160a01b03851661074357604051626a0d4560e21b81525f600482015260240161041c565b6103e885858585856001610a48565b6001600160a01b03811633148015906107725750610770813361050d565b155b1561058a5760405163711bec9160e11b81523360048201526001600160a01b038216602482015260440161041c565b6001600160a01b0383166107c957604051626a0d4560e21b81525f600482015260240161041c565b6104d2835f848460405180602001604052805f8152506001610a48565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b6001600160a01b03841661086057604051632bfa23e760e11b81525f600482015260240161041c565b604080516001808252602082018690528183019081526060820185905260808201909252906108935f8784848784610a48565b505050505050565b6001600160a01b0383166108c457604051631f18c42760e11b81525f600482015260240161041c565b6001600160a01b0382166108ec5760405162ced3e160e81b81525f600482015260240161041c565b6001600160a01b038381165f81815260016020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6001600160a01b03841661098157604051632bfa23e760e11b81525f600482015260240161041c565b6001600160a01b0385166109a957604051626a0d4560e21b81525f600482015260240161041c565b604080516001808252602082018690528183019081526060820185905260808201909252906109dc87878484875f610a48565b50505050505050565b6001600160a01b038316610a0d57604051626a0d4560e21b81525f600482015260240161041c565b604080516001808252602082018590528183019081526060820184905260a082019092525f608082018181529192916103e891879185908590835b610a5486868686610aa0565b6001600160a01b0385161561089357338115610a7d57610a78818888888888610caf565b6109dc565b60208581015190850151610a95838a8a85858a610dd0565b505050505050505050565b8051825114610acf5781518151604051635b05999160e01b81526004810192909252602482015260440161041c565b335f5b8351811015610bd1576020818102858101820151908501909101516001600160a01b03881615610b83575f828152602081815260408083206001600160a01b038c16845290915290205481811015610b5d576040516303dee4c560e01b81526001600160a01b038a16600482015260248101829052604481018390526064810184905260840161041c565b5f838152602081815260408083206001600160a01b038d16845290915290209082900390555b6001600160a01b03871615610bc7575f828152602081815260408083206001600160a01b038b16845290915281208054839290610bc190849061151b565b90915550505b5050600101610ad2565b508251600103610c515760208301515f906020840151909150856001600160a01b0316876001600160a01b0316846001600160a01b03167fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f628585604051610c42929190918252602082015260400190565b60405180910390a450506103e8565b836001600160a01b0316856001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb8686604051610ca092919061153a565b60405180910390a45050505050565b6001600160a01b0384163b156108935760405163bc197c8160e01b81526001600160a01b0385169063bc197c8190610cf39089908990889088908890600401611567565b6020604051808303815f875af1925050508015610d2d575060408051601f3d908101601f19168201909252610d2a918101906115c4565b60015b610d94573d808015610d5a576040519150601f19603f3d011682016040523d82523d5f602084013e610d5f565b606091505b5080515f03610d8c57604051632bfa23e760e11b81526001600160a01b038616600482015260240161041c565b805160208201fd5b6001600160e01b0319811663bc197c8160e01b146109dc57604051632bfa23e760e11b81526001600160a01b038616600482015260240161041c565b6001600160a01b0384163b156108935760405163f23a6e6160e01b81526001600160a01b0385169063f23a6e6190610e1490899089908890889088906004016115df565b6020604051808303815f875af1925050508015610e4e575060408051601f3d908101601f19168201909252610e4b918101906115c4565b60015b610e7b573d808015610d5a576040519150601f19603f3d011682016040523d82523d5f602084013e610d5f565b6001600160e01b0319811663f23a6e6160e01b146109dc57604051632bfa23e760e11b81526001600160a01b038616600482015260240161041c565b80356001600160a01b03811681146103b1575f5ffd5b5f5f60408385031215610ede575f5ffd5b610ee783610eb7565b946020939093013593505050565b6001600160e01b03198116811461058a575f5ffd5b5f60208284031215610f1a575f5ffd5b8135610f2581610ef5565b9392505050565b5f60208284031215610f3c575f5ffd5b5035919050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f610f256020830184610f43565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b0381118282101715610fbf57610fbf610f83565b604052919050565b5f6001600160401b03821115610fdf57610fdf610f83565b5060051b60200190565b5f82601f830112610ff8575f5ffd5b813561100b61100682610fc7565b610f97565b8082825260208201915060208360051b86010192508583111561102c575f5ffd5b602085015b83811015611049578035835260209283019201611031565b5095945050505050565b5f82601f830112611062575f5ffd5b81356001600160401b0381111561107b5761107b610f83565b61108e601f8201601f1916602001610f97565b8181528460208386010111156110a2575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f5f608085870312156110d1575f5ffd5b6110da85610eb7565b935060208501356001600160401b038111156110f4575f5ffd5b61110087828801610fe9565b93505060408501356001600160401b0381111561111b575f5ffd5b61112787828801610fe9565b92505060608501356001600160401b03811115611142575f5ffd5b61114e87828801611053565b91505092959194509250565b5f5f5f5f5f60a0868803121561116e575f5ffd5b61117786610eb7565b945061118560208701610eb7565b935060408601356001600160401b0381111561119f575f5ffd5b6111ab88828901610fe9565b93505060608601356001600160401b038111156111c6575f5ffd5b6111d288828901610fe9565b92505060808601356001600160401b038111156111ed575f5ffd5b6111f988828901611053565b9150509295509295909350565b5f5f60408385031215611217575f5ffd5b82356001600160401b0381111561122c575f5ffd5b8301601f8101851361123c575f5ffd5b803561124a61100682610fc7565b8082825260208201915060208360051b85010192508783111561126b575f5ffd5b6020840193505b828410156112945761128384610eb7565b825260209384019390910190611272565b945050505060208301356001600160401b038111156112b1575f5ffd5b6112bd85828601610fe9565b9150509250929050565b5f8151808452602084019350602083015f5b828110156112f75781518652602095860195909101906001016112d9565b5093949350505050565b602081525f610f2560208301846112c7565b5f5f5f60608486031215611325575f5ffd5b61132e84610eb7565b925060208401356001600160401b03811115611348575f5ffd5b61135486828701610fe9565b92505060408401356001600160401b0381111561136f575f5ffd5b61137b86828701610fe9565b9150509250925092565b5f5f5f5f60808587031215611398575f5ffd5b6113a185610eb7565b9350602085013592506040850135915060608501356001600160401b03811115611142575f5ffd5b5f5f604083850312156113da575f5ffd5b6113e383610eb7565b9150602083013580151581146113f7575f5ffd5b809150509250929050565b5f5f60408385031215611413575f5ffd5b61141c83610eb7565b915061142a60208401610eb7565b90509250929050565b5f5f5f5f5f60a08688031215611447575f5ffd5b61145086610eb7565b945061145e60208701610eb7565b9350604086013592506060860135915060808601356001600160401b038111156111ed575f5ffd5b5f60208284031215611496575f5ffd5b610f2582610eb7565b5f5f5f606084860312156114b1575f5ffd5b6114ba84610eb7565b95602085013595506040909401359392505050565b634e487b7160e01b5f52603260045260245ffd5b600181811c908216806114f757607f821691505b60208210810361151557634e487b7160e01b5f52602260045260245ffd5b50919050565b808201808211156102b857634e487b7160e01b5f52601160045260245ffd5b604081525f61154c60408301856112c7565b828103602084015261155e81856112c7565b95945050505050565b6001600160a01b0386811682528516602082015260a0604082018190525f90611592908301866112c7565b82810360608401526115a481866112c7565b905082810360808401526115b88185610f43565b98975050505050505050565b5f602082840312156115d4575f5ffd5b8151610f2581610ef5565b6001600160a01b03868116825285166020820152604081018490526060810183905260a0608082018190525f9061161890830184610f43565b97965050505050505056fe687474703a2f2f3132372e302e302e313a383038302f697066732f516d643569467469396d695774734775696a7841686e583859723173724b6e4a585a72727275486d324e754434363f66696c656e616d653d73776f72642e706e67687474703a2f2f3132372e302e302e313a383038302f697066732f516d643767434d3248416d31425a57454b4c6e526b5362626b386654694c5a324c785454596e345a4339696b7a783f66696c656e616d653d676f6c642e706e67687474703a2f2f3132372e302e302e313a383038302f697066732f516d546744587556746e4c343958616b4365527a654e6d4562597762506e65335a464e6773447372417a754d316b3f66696c656e616d653d73696c7665722e706e67687474703a2f2f3132372e302e302e313a383038302f697066732f516d564638614568426e64537a6d6b4e4336334556557353635a68636752454e333843565936554632787146334a3f66696c656e616d653d736869656c642e706e67a2646970667358221220de9ac453ee5cf41d58d683c6599c32c87aaf8fc2beb7c0191eac4bb02c09683464736f6c634300081e0033687474703a2f2f3132372e302e302e313a383038302f697066732f7b69647d3f66696c656e616d653d",
}

// MyERC1155TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use MyERC1155TokenMetaData.ABI instead.
var MyERC1155TokenABI = MyERC1155TokenMetaData.ABI

// MyERC1155TokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MyERC1155TokenMetaData.Bin instead.
var MyERC1155TokenBin = MyERC1155TokenMetaData.Bin

// DeployMyERC1155Token deploys a new Ethereum contract, binding an instance of MyERC1155Token to it.
func DeployMyERC1155Token(auth *bind.TransactOpts, backend bind.ContractBackend, initialOwner common.Address) (common.Address, *types.Transaction, *MyERC1155Token, error) {
	parsed, err := MyERC1155TokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MyERC1155TokenBin), backend, initialOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyERC1155Token{MyERC1155TokenCaller: MyERC1155TokenCaller{contract: contract}, MyERC1155TokenTransactor: MyERC1155TokenTransactor{contract: contract}, MyERC1155TokenFilterer: MyERC1155TokenFilterer{contract: contract}}, nil
}

// MyERC1155Token is an auto generated Go binding around an Ethereum contract.
type MyERC1155Token struct {
	MyERC1155TokenCaller     // Read-only binding to the contract
	MyERC1155TokenTransactor // Write-only binding to the contract
	MyERC1155TokenFilterer   // Log filterer for contract events
}

// MyERC1155TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type MyERC1155TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyERC1155TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MyERC1155TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyERC1155TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyERC1155TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyERC1155TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyERC1155TokenSession struct {
	Contract     *MyERC1155Token   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyERC1155TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyERC1155TokenCallerSession struct {
	Contract *MyERC1155TokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MyERC1155TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyERC1155TokenTransactorSession struct {
	Contract     *MyERC1155TokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MyERC1155TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type MyERC1155TokenRaw struct {
	Contract *MyERC1155Token // Generic contract binding to access the raw methods on
}

// MyERC1155TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyERC1155TokenCallerRaw struct {
	Contract *MyERC1155TokenCaller // Generic read-only contract binding to access the raw methods on
}

// MyERC1155TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyERC1155TokenTransactorRaw struct {
	Contract *MyERC1155TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyERC1155Token creates a new instance of MyERC1155Token, bound to a specific deployed contract.
func NewMyERC1155Token(address common.Address, backend bind.ContractBackend) (*MyERC1155Token, error) {
	contract, err := bindMyERC1155Token(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyERC1155Token{MyERC1155TokenCaller: MyERC1155TokenCaller{contract: contract}, MyERC1155TokenTransactor: MyERC1155TokenTransactor{contract: contract}, MyERC1155TokenFilterer: MyERC1155TokenFilterer{contract: contract}}, nil
}

// NewMyERC1155TokenCaller creates a new read-only instance of MyERC1155Token, bound to a specific deployed contract.
func NewMyERC1155TokenCaller(address common.Address, caller bind.ContractCaller) (*MyERC1155TokenCaller, error) {
	contract, err := bindMyERC1155Token(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyERC1155TokenCaller{contract: contract}, nil
}

// NewMyERC1155TokenTransactor creates a new write-only instance of MyERC1155Token, bound to a specific deployed contract.
func NewMyERC1155TokenTransactor(address common.Address, transactor bind.ContractTransactor) (*MyERC1155TokenTransactor, error) {
	contract, err := bindMyERC1155Token(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyERC1155TokenTransactor{contract: contract}, nil
}

// NewMyERC1155TokenFilterer creates a new log filterer instance of MyERC1155Token, bound to a specific deployed contract.
func NewMyERC1155TokenFilterer(address common.Address, filterer bind.ContractFilterer) (*MyERC1155TokenFilterer, error) {
	contract, err := bindMyERC1155Token(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyERC1155TokenFilterer{contract: contract}, nil
}

// bindMyERC1155Token binds a generic wrapper to an already deployed contract.
func bindMyERC1155Token(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MyERC1155TokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyERC1155Token *MyERC1155TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyERC1155Token.Contract.MyERC1155TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyERC1155Token *MyERC1155TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyERC1155Token.Contract.MyERC1155TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyERC1155Token *MyERC1155TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyERC1155Token.Contract.MyERC1155TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyERC1155Token *MyERC1155TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyERC1155Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyERC1155Token *MyERC1155TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyERC1155Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyERC1155Token *MyERC1155TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyERC1155Token.Contract.contract.Transact(opts, method, params...)
}

// GOLD is a free data retrieval call binding the contract method 0x3e4bee38.
//
// Solidity: function GOLD() view returns(uint256)
func (_MyERC1155Token *MyERC1155TokenCaller) GOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyERC1155Token.contract.Call(opts, &out, "GOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GOLD is a free data retrieval call binding the contract method 0x3e4bee38.
//
// Solidity: function GOLD() view returns(uint256)
func (_MyERC1155Token *MyERC1155TokenSession) GOLD() (*big.Int, error) {
	return _MyERC1155Token.Contract.GOLD(&_MyERC1155Token.CallOpts)
}

// GOLD is a free data retrieval call binding the contract method 0x3e4bee38.
//
// Solidity: function GOLD() view returns(uint256)
func (_MyERC1155Token *MyERC1155TokenCallerSession) GOLD() (*big.Int, error) {
	return _MyERC1155Token.Contract.GOLD(&_MyERC1155Token.CallOpts)
}

// SHIELD is a free data retrieval call binding the contract method 0x5b2725ed.
//
// Solidity: function SHIELD() view returns(uint256)
func (_MyERC1155Token *MyERC1155TokenCaller) SHIELD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyERC1155Token.contract.Call(opts, &out, "SHIELD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SHIELD is a free data retrieval call binding the contract method 0x5b2725ed.
//
// Solidity: function SHIELD() view returns(uint256)
func (_MyERC1155Token *MyERC1155TokenSession) SHIELD() (*big.Int, error) {
	return _MyERC1155Token.Contract.SHIELD(&_MyERC1155Token.CallOpts)
}

// SHIELD is a free data retrieval call binding the contract method 0x5b2725ed.
//
// Solidity: function SHIELD() view returns(uint256)
func (_MyERC1155Token *MyERC1155TokenCallerSession) SHIELD() (*big.Int, error) {
	return _MyERC1155Token.Contract.SHIELD(&_MyERC1155Token.CallOpts)
}

// SILVER is a free data retrieval call binding the contract method 0xe3e55f08.
//
// Solidity: function SILVER() view returns(uint256)
func (_MyERC1155Token *MyERC1155TokenCaller) SILVER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyERC1155Token.contract.Call(opts, &out, "SILVER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SILVER is a free data retrieval call binding the contract method 0xe3e55f08.
//
// Solidity: function SILVER() view returns(uint256)
func (_MyERC1155Token *MyERC1155TokenSession) SILVER() (*big.Int, error) {
	return _MyERC1155Token.Contract.SILVER(&_MyERC1155Token.CallOpts)
}

// SILVER is a free data retrieval call binding the contract method 0xe3e55f08.
//
// Solidity: function SILVER() view returns(uint256)
func (_MyERC1155Token *MyERC1155TokenCallerSession) SILVER() (*big.Int, error) {
	return _MyERC1155Token.Contract.SILVER(&_MyERC1155Token.CallOpts)
}

// SWORD is a free data retrieval call binding the contract method 0x13dc989f.
//
// Solidity: function SWORD() view returns(uint256)
func (_MyERC1155Token *MyERC1155TokenCaller) SWORD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyERC1155Token.contract.Call(opts, &out, "SWORD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SWORD is a free data retrieval call binding the contract method 0x13dc989f.
//
// Solidity: function SWORD() view returns(uint256)
func (_MyERC1155Token *MyERC1155TokenSession) SWORD() (*big.Int, error) {
	return _MyERC1155Token.Contract.SWORD(&_MyERC1155Token.CallOpts)
}

// SWORD is a free data retrieval call binding the contract method 0x13dc989f.
//
// Solidity: function SWORD() view returns(uint256)
func (_MyERC1155Token *MyERC1155TokenCallerSession) SWORD() (*big.Int, error) {
	return _MyERC1155Token.Contract.SWORD(&_MyERC1155Token.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_MyERC1155Token *MyERC1155TokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MyERC1155Token.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_MyERC1155Token *MyERC1155TokenSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _MyERC1155Token.Contract.BalanceOf(&_MyERC1155Token.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_MyERC1155Token *MyERC1155TokenCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _MyERC1155Token.Contract.BalanceOf(&_MyERC1155Token.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_MyERC1155Token *MyERC1155TokenCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _MyERC1155Token.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_MyERC1155Token *MyERC1155TokenSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _MyERC1155Token.Contract.BalanceOfBatch(&_MyERC1155Token.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_MyERC1155Token *MyERC1155TokenCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _MyERC1155Token.Contract.BalanceOfBatch(&_MyERC1155Token.CallOpts, accounts, ids)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_MyERC1155Token *MyERC1155TokenCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _MyERC1155Token.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_MyERC1155Token *MyERC1155TokenSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _MyERC1155Token.Contract.IsApprovedForAll(&_MyERC1155Token.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_MyERC1155Token *MyERC1155TokenCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _MyERC1155Token.Contract.IsApprovedForAll(&_MyERC1155Token.CallOpts, account, operator)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyERC1155Token *MyERC1155TokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MyERC1155Token.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyERC1155Token *MyERC1155TokenSession) Owner() (common.Address, error) {
	return _MyERC1155Token.Contract.Owner(&_MyERC1155Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyERC1155Token *MyERC1155TokenCallerSession) Owner() (common.Address, error) {
	return _MyERC1155Token.Contract.Owner(&_MyERC1155Token.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MyERC1155Token *MyERC1155TokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MyERC1155Token.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MyERC1155Token *MyERC1155TokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MyERC1155Token.Contract.SupportsInterface(&_MyERC1155Token.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MyERC1155Token *MyERC1155TokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MyERC1155Token.Contract.SupportsInterface(&_MyERC1155Token.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_MyERC1155Token *MyERC1155TokenCaller) Uri(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _MyERC1155Token.contract.Call(opts, &out, "uri", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_MyERC1155Token *MyERC1155TokenSession) Uri(tokenId *big.Int) (string, error) {
	return _MyERC1155Token.Contract.Uri(&_MyERC1155Token.CallOpts, tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_MyERC1155Token *MyERC1155TokenCallerSession) Uri(tokenId *big.Int) (string, error) {
	return _MyERC1155Token.Contract.Uri(&_MyERC1155Token.CallOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address from, uint256 id, uint256 amount) returns()
func (_MyERC1155Token *MyERC1155TokenTransactor) Burn(opts *bind.TransactOpts, from common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _MyERC1155Token.contract.Transact(opts, "burn", from, id, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address from, uint256 id, uint256 amount) returns()
func (_MyERC1155Token *MyERC1155TokenSession) Burn(from common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _MyERC1155Token.Contract.Burn(&_MyERC1155Token.TransactOpts, from, id, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address from, uint256 id, uint256 amount) returns()
func (_MyERC1155Token *MyERC1155TokenTransactorSession) Burn(from common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _MyERC1155Token.Contract.Burn(&_MyERC1155Token.TransactOpts, from, id, amount)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address from, uint256[] ids, uint256[] amounts) returns()
func (_MyERC1155Token *MyERC1155TokenTransactor) BurnBatch(opts *bind.TransactOpts, from common.Address, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _MyERC1155Token.contract.Transact(opts, "burnBatch", from, ids, amounts)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address from, uint256[] ids, uint256[] amounts) returns()
func (_MyERC1155Token *MyERC1155TokenSession) BurnBatch(from common.Address, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _MyERC1155Token.Contract.BurnBatch(&_MyERC1155Token.TransactOpts, from, ids, amounts)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address from, uint256[] ids, uint256[] amounts) returns()
func (_MyERC1155Token *MyERC1155TokenTransactorSession) BurnBatch(from common.Address, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _MyERC1155Token.Contract.BurnBatch(&_MyERC1155Token.TransactOpts, from, ids, amounts)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address to, uint256 id, uint256 amount, bytes data) returns()
func (_MyERC1155Token *MyERC1155TokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MyERC1155Token.contract.Transact(opts, "mint", to, id, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address to, uint256 id, uint256 amount, bytes data) returns()
func (_MyERC1155Token *MyERC1155TokenSession) Mint(to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MyERC1155Token.Contract.Mint(&_MyERC1155Token.TransactOpts, to, id, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address to, uint256 id, uint256 amount, bytes data) returns()
func (_MyERC1155Token *MyERC1155TokenTransactorSession) Mint(to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MyERC1155Token.Contract.Mint(&_MyERC1155Token.TransactOpts, to, id, amount, data)
}

// MintBatch is a paid mutator transaction binding the contract method 0x1f7fdffa.
//
// Solidity: function mintBatch(address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_MyERC1155Token *MyERC1155TokenTransactor) MintBatch(opts *bind.TransactOpts, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _MyERC1155Token.contract.Transact(opts, "mintBatch", to, ids, amounts, data)
}

// MintBatch is a paid mutator transaction binding the contract method 0x1f7fdffa.
//
// Solidity: function mintBatch(address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_MyERC1155Token *MyERC1155TokenSession) MintBatch(to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _MyERC1155Token.Contract.MintBatch(&_MyERC1155Token.TransactOpts, to, ids, amounts, data)
}

// MintBatch is a paid mutator transaction binding the contract method 0x1f7fdffa.
//
// Solidity: function mintBatch(address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_MyERC1155Token *MyERC1155TokenTransactorSession) MintBatch(to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _MyERC1155Token.Contract.MintBatch(&_MyERC1155Token.TransactOpts, to, ids, amounts, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MyERC1155Token *MyERC1155TokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyERC1155Token.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MyERC1155Token *MyERC1155TokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _MyERC1155Token.Contract.RenounceOwnership(&_MyERC1155Token.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MyERC1155Token *MyERC1155TokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MyERC1155Token.Contract.RenounceOwnership(&_MyERC1155Token.TransactOpts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_MyERC1155Token *MyERC1155TokenTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _MyERC1155Token.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_MyERC1155Token *MyERC1155TokenSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _MyERC1155Token.Contract.SafeBatchTransferFrom(&_MyERC1155Token.TransactOpts, from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_MyERC1155Token *MyERC1155TokenTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _MyERC1155Token.Contract.SafeBatchTransferFrom(&_MyERC1155Token.TransactOpts, from, to, ids, values, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_MyERC1155Token *MyERC1155TokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MyERC1155Token.contract.Transact(opts, "safeTransferFrom", from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_MyERC1155Token *MyERC1155TokenSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MyERC1155Token.Contract.SafeTransferFrom(&_MyERC1155Token.TransactOpts, from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_MyERC1155Token *MyERC1155TokenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MyERC1155Token.Contract.SafeTransferFrom(&_MyERC1155Token.TransactOpts, from, to, id, value, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MyERC1155Token *MyERC1155TokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _MyERC1155Token.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MyERC1155Token *MyERC1155TokenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MyERC1155Token.Contract.SetApprovalForAll(&_MyERC1155Token.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MyERC1155Token *MyERC1155TokenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MyERC1155Token.Contract.SetApprovalForAll(&_MyERC1155Token.TransactOpts, operator, approved)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MyERC1155Token *MyERC1155TokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MyERC1155Token.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MyERC1155Token *MyERC1155TokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MyERC1155Token.Contract.TransferOwnership(&_MyERC1155Token.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MyERC1155Token *MyERC1155TokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MyERC1155Token.Contract.TransferOwnership(&_MyERC1155Token.TransactOpts, newOwner)
}

// MyERC1155TokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the MyERC1155Token contract.
type MyERC1155TokenApprovalForAllIterator struct {
	Event *MyERC1155TokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *MyERC1155TokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyERC1155TokenApprovalForAll)
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
		it.Event = new(MyERC1155TokenApprovalForAll)
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
func (it *MyERC1155TokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyERC1155TokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyERC1155TokenApprovalForAll represents a ApprovalForAll event raised by the MyERC1155Token contract.
type MyERC1155TokenApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_MyERC1155Token *MyERC1155TokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*MyERC1155TokenApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MyERC1155Token.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MyERC1155TokenApprovalForAllIterator{contract: _MyERC1155Token.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_MyERC1155Token *MyERC1155TokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MyERC1155TokenApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MyERC1155Token.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyERC1155TokenApprovalForAll)
				if err := _MyERC1155Token.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_MyERC1155Token *MyERC1155TokenFilterer) ParseApprovalForAll(log types.Log) (*MyERC1155TokenApprovalForAll, error) {
	event := new(MyERC1155TokenApprovalForAll)
	if err := _MyERC1155Token.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyERC1155TokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MyERC1155Token contract.
type MyERC1155TokenOwnershipTransferredIterator struct {
	Event *MyERC1155TokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MyERC1155TokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyERC1155TokenOwnershipTransferred)
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
		it.Event = new(MyERC1155TokenOwnershipTransferred)
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
func (it *MyERC1155TokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyERC1155TokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyERC1155TokenOwnershipTransferred represents a OwnershipTransferred event raised by the MyERC1155Token contract.
type MyERC1155TokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MyERC1155Token *MyERC1155TokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MyERC1155TokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MyERC1155Token.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MyERC1155TokenOwnershipTransferredIterator{contract: _MyERC1155Token.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MyERC1155Token *MyERC1155TokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MyERC1155TokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MyERC1155Token.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyERC1155TokenOwnershipTransferred)
				if err := _MyERC1155Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MyERC1155Token *MyERC1155TokenFilterer) ParseOwnershipTransferred(log types.Log) (*MyERC1155TokenOwnershipTransferred, error) {
	event := new(MyERC1155TokenOwnershipTransferred)
	if err := _MyERC1155Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyERC1155TokenTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the MyERC1155Token contract.
type MyERC1155TokenTransferBatchIterator struct {
	Event *MyERC1155TokenTransferBatch // Event containing the contract specifics and raw log

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
func (it *MyERC1155TokenTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyERC1155TokenTransferBatch)
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
		it.Event = new(MyERC1155TokenTransferBatch)
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
func (it *MyERC1155TokenTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyERC1155TokenTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyERC1155TokenTransferBatch represents a TransferBatch event raised by the MyERC1155Token contract.
type MyERC1155TokenTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_MyERC1155Token *MyERC1155TokenFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*MyERC1155TokenTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MyERC1155Token.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MyERC1155TokenTransferBatchIterator{contract: _MyERC1155Token.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_MyERC1155Token *MyERC1155TokenFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *MyERC1155TokenTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MyERC1155Token.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyERC1155TokenTransferBatch)
				if err := _MyERC1155Token.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_MyERC1155Token *MyERC1155TokenFilterer) ParseTransferBatch(log types.Log) (*MyERC1155TokenTransferBatch, error) {
	event := new(MyERC1155TokenTransferBatch)
	if err := _MyERC1155Token.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyERC1155TokenTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the MyERC1155Token contract.
type MyERC1155TokenTransferSingleIterator struct {
	Event *MyERC1155TokenTransferSingle // Event containing the contract specifics and raw log

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
func (it *MyERC1155TokenTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyERC1155TokenTransferSingle)
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
		it.Event = new(MyERC1155TokenTransferSingle)
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
func (it *MyERC1155TokenTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyERC1155TokenTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyERC1155TokenTransferSingle represents a TransferSingle event raised by the MyERC1155Token contract.
type MyERC1155TokenTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_MyERC1155Token *MyERC1155TokenFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*MyERC1155TokenTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MyERC1155Token.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MyERC1155TokenTransferSingleIterator{contract: _MyERC1155Token.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_MyERC1155Token *MyERC1155TokenFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *MyERC1155TokenTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MyERC1155Token.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyERC1155TokenTransferSingle)
				if err := _MyERC1155Token.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_MyERC1155Token *MyERC1155TokenFilterer) ParseTransferSingle(log types.Log) (*MyERC1155TokenTransferSingle, error) {
	event := new(MyERC1155TokenTransferSingle)
	if err := _MyERC1155Token.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyERC1155TokenURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the MyERC1155Token contract.
type MyERC1155TokenURIIterator struct {
	Event *MyERC1155TokenURI // Event containing the contract specifics and raw log

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
func (it *MyERC1155TokenURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyERC1155TokenURI)
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
		it.Event = new(MyERC1155TokenURI)
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
func (it *MyERC1155TokenURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyERC1155TokenURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyERC1155TokenURI represents a URI event raised by the MyERC1155Token contract.
type MyERC1155TokenURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_MyERC1155Token *MyERC1155TokenFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*MyERC1155TokenURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MyERC1155Token.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &MyERC1155TokenURIIterator{contract: _MyERC1155Token.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_MyERC1155Token *MyERC1155TokenFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *MyERC1155TokenURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MyERC1155Token.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyERC1155TokenURI)
				if err := _MyERC1155Token.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_MyERC1155Token *MyERC1155TokenFilterer) ParseURI(log types.Log) (*MyERC1155TokenURI, error) {
	event := new(MyERC1155TokenURI)
	if err := _MyERC1155Token.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
