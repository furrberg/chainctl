// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package qweezxc

import (
	"errors"
	"math/big"
	"strings"

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
)

// QweezxcMetaData contains all meta data concerning the Qweezxc contract.
var QweezxcMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b506040518060400160405280600781526020017f717765657a7863000000000000000000000000000000000000000000000000008152506040518060400160405280600481526020017f515a584300000000000000000000000000000000000000000000000000000000815250816003908161008b919061048f565b50806004908161009b919061048f565b5050506100d8336100b06100dd60201b60201c565b60ff16600a6100bf91906106ba565b620f42406100cd9190610704565b6100e560201b60201c565b610818565b5f6012905090565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610153576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161014a9061079f565b60405180910390fd5b6101645f838361024860201b60201c565b8060025f82825461017591906107bd565b92505081905550805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546101c791906107bd565b925050819055508173ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161022b91906107ff565b60405180910390a36102445f838361024d60201b60201c565b5050565b505050565b505050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806102cd57607f821691505b6020821081036102e0576102df610289565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026103427fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610307565b61034c8683610307565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61039061038b61038684610364565b61036d565b610364565b9050919050565b5f819050919050565b6103a983610376565b6103bd6103b582610397565b848454610313565b825550505050565b5f5f905090565b6103d46103c5565b6103df8184846103a0565b505050565b5b81811015610402576103f75f826103cc565b6001810190506103e5565b5050565b601f82111561044757610418816102e6565b610421846102f8565b81016020851015610430578190505b61044461043c856102f8565b8301826103e4565b50505b505050565b5f82821c905092915050565b5f6104675f198460080261044c565b1980831691505092915050565b5f61047f8383610458565b9150826002028217905092915050565b61049882610252565b67ffffffffffffffff8111156104b1576104b061025c565b5b6104bb82546102b6565b6104c6828285610406565b5f60209050601f8311600181146104f7575f84156104e5578287015190505b6104ef8582610474565b865550610556565b601f198416610505866102e6565b5f5b8281101561052c57848901518255600182019150602085019450602081019050610507565b868310156105495784890151610545601f891682610458565b8355505b6001600288020188555050505b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f8160011c9050919050565b5f5f8291508390505b60018511156105e0578086048111156105bc576105bb61055e565b5b60018516156105cb5780820291505b80810290506105d98561058b565b94506105a0565b94509492505050565b5f826105f857600190506106b3565b81610605575f90506106b3565b816001811461061b576002811461062557610654565b60019150506106b3565b60ff8411156106375761063661055e565b5b8360020a91508482111561064e5761064d61055e565b5b506106b3565b5060208310610133831016604e8410600b84101617156106895782820a9050838111156106845761068361055e565b5b6106b3565b6106968484846001610597565b925090508184048111156106ad576106ac61055e565b5b81810290505b9392505050565b5f6106c482610364565b91506106cf83610364565b92506106fc7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846105e9565b905092915050565b5f61070e82610364565b915061071983610364565b925082820261072781610364565b9150828204841483151761073e5761073d61055e565b5b5092915050565b5f82825260208201905092915050565b7f45524332303a206d696e7420746f20746865207a65726f2061646472657373005f82015250565b5f610789601f83610745565b915061079482610755565b602082019050919050565b5f6020820190508181035f8301526107b68161077d565b9050919050565b5f6107c782610364565b91506107d283610364565b92508282019050808211156107ea576107e961055e565b5b92915050565b6107f981610364565b82525050565b5f6020820190506108125f8301846107f0565b92915050565b6111b2806108255f395ff3fe608060405234801561000f575f5ffd5b50600436106100a7575f3560e01c8063395093511161006f578063395093511461016557806370a082311461019557806395d89b41146101c5578063a457c2d7146101e3578063a9059cbb14610213578063dd62ed3e14610243576100a7565b806306fdde03146100ab578063095ea7b3146100c957806318160ddd146100f957806323b872dd14610117578063313ce56714610147575b5f5ffd5b6100b3610273565b6040516100c09190610acc565b60405180910390f35b6100e360048036038101906100de9190610b7d565b610303565b6040516100f09190610bd5565b60405180910390f35b610101610325565b60405161010e9190610bfd565b60405180910390f35b610131600480360381019061012c9190610c16565b61032e565b60405161013e9190610bd5565b60405180910390f35b61014f61035c565b60405161015c9190610c81565b60405180910390f35b61017f600480360381019061017a9190610b7d565b610364565b60405161018c9190610bd5565b60405180910390f35b6101af60048036038101906101aa9190610c9a565b61039a565b6040516101bc9190610bfd565b60405180910390f35b6101cd6103df565b6040516101da9190610acc565b60405180910390f35b6101fd60048036038101906101f89190610b7d565b61046f565b60405161020a9190610bd5565b60405180910390f35b61022d60048036038101906102289190610b7d565b6104e4565b60405161023a9190610bd5565b60405180910390f35b61025d60048036038101906102589190610cc5565b610506565b60405161026a9190610bfd565b60405180910390f35b60606003805461028290610d30565b80601f01602080910402602001604051908101604052809291908181526020018280546102ae90610d30565b80156102f95780601f106102d0576101008083540402835291602001916102f9565b820191905f5260205f20905b8154815290600101906020018083116102dc57829003601f168201915b5050505050905090565b5f5f61030d610588565b905061031a81858561058f565b600191505092915050565b5f600254905090565b5f5f610338610588565b9050610345858285610752565b6103508585856107dd565b60019150509392505050565b5f6012905090565b5f5f61036e610588565b905061038f8185856103808589610506565b61038a9190610d8d565b61058f565b600191505092915050565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b6060600480546103ee90610d30565b80601f016020809104026020016040519081016040528092919081815260200182805461041a90610d30565b80156104655780601f1061043c57610100808354040283529160200191610465565b820191905f5260205f20905b81548152906001019060200180831161044857829003601f168201915b5050505050905090565b5f5f610479610588565b90505f6104868286610506565b9050838110156104cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104c290610e30565b60405180910390fd5b6104d8828686840361058f565b60019250505092915050565b5f5f6104ee610588565b90506104fb8185856107dd565b600191505092915050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b5f33905090565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036105fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105f490610ebe565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361066b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066290610f4c565b60405180910390fd5b8060015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516107459190610bfd565b60405180910390a3505050565b5f61075d8484610506565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146107d757818110156107c9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107c090610fb4565b60405180910390fd5b6107d6848484840361058f565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361084b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161084290611042565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036108b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108b0906110d0565b60405180910390fd5b6108c4838383610a52565b5f5f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015610947576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093e9061115e565b60405180910390fd5b8181035f5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550815f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546109d59190610d8d565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610a399190610bfd565b60405180910390a3610a4c848484610a57565b50505050565b505050565b505050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610a9e82610a5c565b610aa88185610a66565b9350610ab8818560208601610a76565b610ac181610a84565b840191505092915050565b5f6020820190508181035f830152610ae48184610a94565b905092915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610b1982610af0565b9050919050565b610b2981610b0f565b8114610b33575f5ffd5b50565b5f81359050610b4481610b20565b92915050565b5f819050919050565b610b5c81610b4a565b8114610b66575f5ffd5b50565b5f81359050610b7781610b53565b92915050565b5f5f60408385031215610b9357610b92610aec565b5b5f610ba085828601610b36565b9250506020610bb185828601610b69565b9150509250929050565b5f8115159050919050565b610bcf81610bbb565b82525050565b5f602082019050610be85f830184610bc6565b92915050565b610bf781610b4a565b82525050565b5f602082019050610c105f830184610bee565b92915050565b5f5f5f60608486031215610c2d57610c2c610aec565b5b5f610c3a86828701610b36565b9350506020610c4b86828701610b36565b9250506040610c5c86828701610b69565b9150509250925092565b5f60ff82169050919050565b610c7b81610c66565b82525050565b5f602082019050610c945f830184610c72565b92915050565b5f60208284031215610caf57610cae610aec565b5b5f610cbc84828501610b36565b91505092915050565b5f5f60408385031215610cdb57610cda610aec565b5b5f610ce885828601610b36565b9250506020610cf985828601610b36565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610d4757607f821691505b602082108103610d5a57610d59610d03565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610d9782610b4a565b9150610da283610b4a565b9250828201905080821115610dba57610db9610d60565b5b92915050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f775f8201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b5f610e1a602583610a66565b9150610e2582610dc0565b604082019050919050565b5f6020820190508181035f830152610e4781610e0e565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f206164645f8201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b5f610ea8602483610a66565b9150610eb382610e4e565b604082019050919050565b5f6020820190508181035f830152610ed581610e9c565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f2061646472655f8201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b5f610f36602283610a66565b9150610f4182610edc565b604082019050919050565b5f6020820190508181035f830152610f6381610f2a565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000005f82015250565b5f610f9e601d83610a66565b9150610fa982610f6a565b602082019050919050565b5f6020820190508181035f830152610fcb81610f92565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f2061645f8201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b5f61102c602583610a66565b915061103782610fd2565b604082019050919050565b5f6020820190508181035f83015261105981611020565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f20616464725f8201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b5f6110ba602383610a66565b91506110c582611060565b604082019050919050565b5f6020820190508181035f8301526110e7816110ae565b9050919050565b7f45524332303a207472616e7366657220616d6f756e74206578636565647320625f8201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b5f611148602683610a66565b9150611153826110ee565b604082019050919050565b5f6020820190508181035f8301526111758161113c565b905091905056fea2646970667358221220d34d3bb73ba190670eef2e0059321fdaa64ce41fe50467d081052e7db24d0d6864736f6c634300081e0033",
}

// QweezxcABI is the input ABI used to generate the binding from.
// Deprecated: Use QweezxcMetaData.ABI instead.
var QweezxcABI = QweezxcMetaData.ABI

// QweezxcBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use QweezxcMetaData.Bin instead.
var QweezxcBin = QweezxcMetaData.Bin

// DeployQweezxc deploys a new Ethereum contract, binding an instance of Qweezxc to it.
func DeployQweezxc(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Qweezxc, error) {
	parsed, err := QweezxcMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(QweezxcBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Qweezxc{QweezxcCaller: QweezxcCaller{contract: contract}, QweezxcTransactor: QweezxcTransactor{contract: contract}, QweezxcFilterer: QweezxcFilterer{contract: contract}}, nil
}

// Qweezxc is an auto generated Go binding around an Ethereum contract.
type Qweezxc struct {
	QweezxcCaller     // Read-only binding to the contract
	QweezxcTransactor // Write-only binding to the contract
	QweezxcFilterer   // Log filterer for contract events
}

// QweezxcCaller is an auto generated read-only Go binding around an Ethereum contract.
type QweezxcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QweezxcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QweezxcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QweezxcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QweezxcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QweezxcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QweezxcSession struct {
	Contract     *Qweezxc          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QweezxcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QweezxcCallerSession struct {
	Contract *QweezxcCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// QweezxcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QweezxcTransactorSession struct {
	Contract     *QweezxcTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// QweezxcRaw is an auto generated low-level Go binding around an Ethereum contract.
type QweezxcRaw struct {
	Contract *Qweezxc // Generic contract binding to access the raw methods on
}

// QweezxcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QweezxcCallerRaw struct {
	Contract *QweezxcCaller // Generic read-only contract binding to access the raw methods on
}

// QweezxcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QweezxcTransactorRaw struct {
	Contract *QweezxcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQweezxc creates a new instance of Qweezxc, bound to a specific deployed contract.
func NewQweezxc(address common.Address, backend bind.ContractBackend) (*Qweezxc, error) {
	contract, err := bindQweezxc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Qweezxc{QweezxcCaller: QweezxcCaller{contract: contract}, QweezxcTransactor: QweezxcTransactor{contract: contract}, QweezxcFilterer: QweezxcFilterer{contract: contract}}, nil
}

// NewQweezxcCaller creates a new read-only instance of Qweezxc, bound to a specific deployed contract.
func NewQweezxcCaller(address common.Address, caller bind.ContractCaller) (*QweezxcCaller, error) {
	contract, err := bindQweezxc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QweezxcCaller{contract: contract}, nil
}

// NewQweezxcTransactor creates a new write-only instance of Qweezxc, bound to a specific deployed contract.
func NewQweezxcTransactor(address common.Address, transactor bind.ContractTransactor) (*QweezxcTransactor, error) {
	contract, err := bindQweezxc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QweezxcTransactor{contract: contract}, nil
}

// NewQweezxcFilterer creates a new log filterer instance of Qweezxc, bound to a specific deployed contract.
func NewQweezxcFilterer(address common.Address, filterer bind.ContractFilterer) (*QweezxcFilterer, error) {
	contract, err := bindQweezxc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QweezxcFilterer{contract: contract}, nil
}

// bindQweezxc binds a generic wrapper to an already deployed contract.
func bindQweezxc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := QweezxcMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Qweezxc *QweezxcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Qweezxc.Contract.QweezxcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Qweezxc *QweezxcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Qweezxc.Contract.QweezxcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Qweezxc *QweezxcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Qweezxc.Contract.QweezxcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Qweezxc *QweezxcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Qweezxc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Qweezxc *QweezxcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Qweezxc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Qweezxc *QweezxcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Qweezxc.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Qweezxc *QweezxcCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Qweezxc.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Qweezxc *QweezxcSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Qweezxc.Contract.Allowance(&_Qweezxc.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Qweezxc *QweezxcCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Qweezxc.Contract.Allowance(&_Qweezxc.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Qweezxc *QweezxcCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Qweezxc.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Qweezxc *QweezxcSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Qweezxc.Contract.BalanceOf(&_Qweezxc.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Qweezxc *QweezxcCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Qweezxc.Contract.BalanceOf(&_Qweezxc.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Qweezxc *QweezxcCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Qweezxc.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Qweezxc *QweezxcSession) Decimals() (uint8, error) {
	return _Qweezxc.Contract.Decimals(&_Qweezxc.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Qweezxc *QweezxcCallerSession) Decimals() (uint8, error) {
	return _Qweezxc.Contract.Decimals(&_Qweezxc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Qweezxc *QweezxcCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Qweezxc.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Qweezxc *QweezxcSession) Name() (string, error) {
	return _Qweezxc.Contract.Name(&_Qweezxc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Qweezxc *QweezxcCallerSession) Name() (string, error) {
	return _Qweezxc.Contract.Name(&_Qweezxc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Qweezxc *QweezxcCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Qweezxc.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Qweezxc *QweezxcSession) Symbol() (string, error) {
	return _Qweezxc.Contract.Symbol(&_Qweezxc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Qweezxc *QweezxcCallerSession) Symbol() (string, error) {
	return _Qweezxc.Contract.Symbol(&_Qweezxc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Qweezxc *QweezxcCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Qweezxc.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Qweezxc *QweezxcSession) TotalSupply() (*big.Int, error) {
	return _Qweezxc.Contract.TotalSupply(&_Qweezxc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Qweezxc *QweezxcCallerSession) TotalSupply() (*big.Int, error) {
	return _Qweezxc.Contract.TotalSupply(&_Qweezxc.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Qweezxc *QweezxcTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Qweezxc.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Qweezxc *QweezxcSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Qweezxc.Contract.Approve(&_Qweezxc.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Qweezxc *QweezxcTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Qweezxc.Contract.Approve(&_Qweezxc.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Qweezxc *QweezxcTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Qweezxc.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Qweezxc *QweezxcSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Qweezxc.Contract.DecreaseAllowance(&_Qweezxc.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Qweezxc *QweezxcTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Qweezxc.Contract.DecreaseAllowance(&_Qweezxc.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Qweezxc *QweezxcTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Qweezxc.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Qweezxc *QweezxcSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Qweezxc.Contract.IncreaseAllowance(&_Qweezxc.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Qweezxc *QweezxcTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Qweezxc.Contract.IncreaseAllowance(&_Qweezxc.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Qweezxc *QweezxcTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Qweezxc.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Qweezxc *QweezxcSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Qweezxc.Contract.Transfer(&_Qweezxc.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Qweezxc *QweezxcTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Qweezxc.Contract.Transfer(&_Qweezxc.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Qweezxc *QweezxcTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Qweezxc.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Qweezxc *QweezxcSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Qweezxc.Contract.TransferFrom(&_Qweezxc.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Qweezxc *QweezxcTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Qweezxc.Contract.TransferFrom(&_Qweezxc.TransactOpts, from, to, amount)
}

// QweezxcApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Qweezxc contract.
type QweezxcApprovalIterator struct {
	Event *QweezxcApproval // Event containing the contract specifics and raw log

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
func (it *QweezxcApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QweezxcApproval)
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
		it.Event = new(QweezxcApproval)
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
func (it *QweezxcApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QweezxcApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QweezxcApproval represents a Approval event raised by the Qweezxc contract.
type QweezxcApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Qweezxc *QweezxcFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*QweezxcApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Qweezxc.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &QweezxcApprovalIterator{contract: _Qweezxc.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Qweezxc *QweezxcFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *QweezxcApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Qweezxc.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QweezxcApproval)
				if err := _Qweezxc.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Qweezxc *QweezxcFilterer) ParseApproval(log types.Log) (*QweezxcApproval, error) {
	event := new(QweezxcApproval)
	if err := _Qweezxc.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QweezxcTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Qweezxc contract.
type QweezxcTransferIterator struct {
	Event *QweezxcTransfer // Event containing the contract specifics and raw log

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
func (it *QweezxcTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QweezxcTransfer)
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
		it.Event = new(QweezxcTransfer)
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
func (it *QweezxcTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QweezxcTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QweezxcTransfer represents a Transfer event raised by the Qweezxc contract.
type QweezxcTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Qweezxc *QweezxcFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*QweezxcTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Qweezxc.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &QweezxcTransferIterator{contract: _Qweezxc.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Qweezxc *QweezxcFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *QweezxcTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Qweezxc.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QweezxcTransfer)
				if err := _Qweezxc.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Qweezxc *QweezxcFilterer) ParseTransfer(log types.Log) (*QweezxcTransfer, error) {
	event := new(QweezxcTransfer)
	if err := _Qweezxc.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
