// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package basic

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GovernanceABI is the input ABI used to generate the binding from.
const GovernanceABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"AssetPausedUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"committerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"CommitterStatusUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"ExecutorStatusUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"MonitorStatusUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"assetId\",\"type\":\"uint16\"}],\"name\":\"NewAsset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAssetGovernance\",\"name\":\"newAssetGovernance\",\"type\":\"address\"}],\"name\":\"NewAssetGovernance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newGovernor\",\"type\":\"address\"}],\"name\":\"NewGovernor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"VerifierStatusUpdate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"assetAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetGovernance\",\"outputs\":[{\"internalType\":\"contractAssetGovernance\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"assetsList\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAssetGovernance\",\"name\":\"_newAssetGovernance\",\"type\":\"address\"}],\"name\":\"changeAssetGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGovernor\",\"type\":\"address\"}],\"name\":\"changeGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"committers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initializationParameters\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAddressExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"monitors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkGovernor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"pausedAssets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"requireActiveCommitter\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"requireActiveMonitor\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"requireActiveVerifier\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"requireGovernor\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_assetId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_assetAddress\",\"type\":\"address\"}],\"name\":\"setAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_assetAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_assetPaused\",\"type\":\"bool\"}],\"name\":\"setAssetPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"setCommitter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"setMonitor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"setVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"upgradeParameters\",\"type\":\"bytes\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_assetAddr\",\"type\":\"address\"}],\"name\":\"validateAssetAddress\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"verifiers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// GovernanceBin is the compiled bytecode used for deploying new contracts.
var GovernanceBin = "0x608060405234801561001057600080fd5b5061001961001e565b610081565b7f8e94fed44239eb2314ab7a406345e6c5a8f0ccedf3b600de3d004e672c33abf480546001909155801561007e576040805162461bcd60e51b815260206004820152600260248201526118a160f11b604482015290519081900360640190fd5b50565b611093806100906000396000f3fe608060405234801561001057600080fd5b50600436106101585760003560e01c80635f4fea26116100c3578063d87e37481161007c578063d87e3748146104d1578063dbfc2967146104f7578063e4c0aaf414610534578063f39349ef1461055a578063f5e7d6fd14610562578063f5f84ed41461056a57610158565b80635f4fea26146103e55780636c8244871461040b5780637a3459d4146104315780639bd7760914610457578063a28e00321461047d578063ac9b5671146104a357610158565b806326eb0f4a1161011557806326eb0f4a1461029657806328b58e1a146102c457806331d8687b146102fe578063321e182b1461031f578063439fab91146103455780635bab15ca146103b557610158565b806301e1d1141461015d5780631afcac961461017c5780631c021956146101ac5780631e763ee3146101d25780632520ce5a146101f85780632539464514610226575b600080fd5b610165610590565b6040805161ffff9092168252519081900360200190f35b6101aa6004803603604081101561019257600080fd5b506001600160a01b03813516906020013515156105a1565b005b6101aa600480360360208110156101c257600080fd5b50356001600160a01b0316610631565b610165600480360360208110156101e857600080fd5b50356001600160a01b0316610693565b6101aa6004803603604081101561020e57600080fd5b506001600160a01b03813516906020013515156106a9565b6101aa6004803603602081101561023c57600080fd5b81019060208101813564010000000081111561025757600080fd5b82018360208201111561026957600080fd5b8035906020019184600183028401116401000000008311171561028b57600080fd5b50909250905061062d565b6101aa600480360360408110156102ac57600080fd5b506001600160a01b03813516906020013515156107b8565b6102ea600480360360208110156102da57600080fd5b50356001600160a01b0316610847565b604080519115158252519081900360200190f35b6102ea6004803603602081101561031457600080fd5b503561ffff1661085c565b6102ea6004803603602081101561033557600080fd5b50356001600160a01b0316610871565b6101aa6004803603602081101561035b57600080fd5b81019060208101813564010000000081111561037657600080fd5b82018360208201111561038857600080fd5b803590602001918460018302840111640100000000831117156103aa57600080fd5b509092509050610886565b6101aa600480360360408110156103cb57600080fd5b50803561ffff1690602001356001600160a01b03166108de565b6102ea600480360360208110156103fb57600080fd5b50356001600160a01b0316610c01565b6102ea6004803603602081101561042157600080fd5b50356001600160a01b0316610c16565b6101aa6004803603602081101561044757600080fd5b50356001600160a01b0316610c2b565b6101656004803603602081101561046d57600080fd5b50356001600160a01b0316610c8b565b6101aa6004803603602081101561049357600080fd5b50356001600160a01b0316610da9565b6101aa600480360360408110156104b957600080fd5b506001600160a01b0381351690602001351515610e0a565b6101aa600480360360208110156104e757600080fd5b50356001600160a01b0316610e99565b6105186004803603602081101561050d57600080fd5b503561ffff16610f0b565b604080516001600160a01b039092168252519081900360200190f35b6101aa6004803603602081101561054a57600080fd5b50356001600160a01b0316610f26565b610518610f98565b610518610fa7565b6101aa6004803603602081101561058057600080fd5b50356001600160a01b0316610fb6565b600154600160a01b900461ffff1681565b6105aa33610fb6565b6001600160a01b03821660009081526002602052604090205460ff1615158115151461062d576001600160a01b038216600081815260026020908152604091829020805460ff1916851515908117909155825190815291517fd158a06928d89556421a4e2eabd98c7ea10e974c07e399e3afc449d19eb391b29281900390910190a25b5050565b6001600160a01b03811660009081526004602052604090205460ff16610690576040805162461bcd60e51b815260206004820152600f60248201526e34b73b30b634b21036b7b734ba37b960891b604482015290519081900360640190fd5b50565b60066020526000908152604090205461ffff1681565b6106b233610fb6565b6000306001600160a01b0316639bd77609846040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561070157600080fd5b505afa158015610715573d6000803e3d6000fd5b505050506040513d602081101561072b57600080fd5b505161ffff811660009081526005602052604090205490915060ff161515821515146107b35761ffff8116600090815260056020908152604091829020805460ff1916851515908117909155825190815291516001600160a01b038616927ff7ca5545623b85829d3abcb6729eb021070bbe93cba64f5c4de6b17b805b7d0292908290030190a25b505050565b6107c133610fb6565b6001600160a01b03821660009081526004602052604090205460ff1615158115151461062d576001600160a01b038216600081815260046020908152604091829020805460ff1916851515908117909155825190815291517f6fbe3fec04b6d2f79e569f1a46ec4d1621e751112483b5ce9d81832f19fee5c19281900390910190a25050565b60046020526000908152604090205460ff1681565b60056020526000908152604090205460ff1681565b60086020526000908152604090205460ff1681565b61088e610ffd565b600080838360408110156108a157600080fd5b506001805482356001600160a01b039081166001600160a01b03199283161790925560008054602090940135909216921691909117905550505050565b6009546001600160a01b03163314610922576040805162461bcd60e51b8152602060048201526002602482015261314560f01b604482015290519081900360640190fd5b60008054906101000a90046001600160a01b03166001600160a01b0316637af87d196040518163ffffffff1660e01b815260040160206040518083038186803b15801561096e57600080fd5b505afa158015610982573d6000803e3d6000fd5b505050506040513d602081101561099857600080fd5b505161ffff838116911614156109da576040805162461bcd60e51b8152602060048201526002602482015261316760f01b604482015290519081900360640190fd5b61ffff82166000908152600760205260409020546001600160a01b031615610a2e576040805162461bcd60e51b815260206004820152600260248201526118b360f11b604482015290519081900360640190fd5b6001600160a01b03811660009081526008602052604090205460ff1615610a81576040805162461bcd60e51b8152602060048201526002602482015261316560f01b604482015290519081900360640190fd5b60008054906101000a90046001600160a01b03166001600160a01b0316637af87d196040518163ffffffff1660e01b815260040160206040518083038186803b158015610acd57600080fd5b505afa158015610ae1573d6000803e3d6000fd5b505050506040513d6020811015610af757600080fd5b50516001600160a01b03821660009081526006602052604090205461ffff90811691161415610b52576040805162461bcd60e51b8152602060048201526002602482015261316760f01b604482015290519081900360640190fd5b6001805461ffff60a01b198116600160a01b9182900461ffff908116840181169092021782556001600160a01b0383166000818152600660209081526040808320805461ffff191695891695861790558483526007825280832080546001600160a01b031916851790558383526008909152808220805460ff19169095179094559251919290917f990c18cf226b253448a6a051b5cadf51a61f4ca56703374cdd8bf9df04b2f9019190a35050565b60026020526000908152604090205460ff1681565b60036020526000908152604090205460ff1681565b6001600160a01b03811660009081526003602052604090205460ff16610690576040805162461bcd60e51b815260206004820152601060248201526f34b73b30b634b2103b32b934b334b2b960811b604482015290519081900360640190fd5b6001600160a01b0380821660009081526006602090815260408083205483548251637af87d1960e01b81529251949561ffff90921694911692637af87d199260048082019391829003018186803b158015610ce557600080fd5b505afa158015610cf9573d6000803e3d6000fd5b505050506040513d6020811015610d0f57600080fd5b505161ffff82811691161415610d51576040805162461bcd60e51b8152602060048201526002602482015261316960f01b604482015290519081900360640190fd5b6001600160a01b03831660009081526008602052604090205460ff16610da3576040805162461bcd60e51b8152602060048201526002602482015261326960f01b604482015290519081900360640190fd5b92915050565b6001600160a01b03811660009081526002602052604090205460ff16610690576040805162461bcd60e51b815260206004820152601160248201527034b73b30b634b21031b7b6b6b4ba3a32b960791b604482015290519081900360640190fd5b610e1333610fb6565b6001600160a01b03821660009081526003602052604090205460ff1615158115151461062d576001600160a01b038216600081815260036020908152604091829020805460ff1916851515908117909155825190815291517ff09ef994a79ba4638c91b1d9595c89e208c3a9b9403769b68670c678782f82d09281900390910190a25050565b610ea233610fb6565b6009546001600160a01b0382811691161461069057600980546001600160a01b0383166001600160a01b0319909116811790915560408051918252517fc7708091594580aae77e08cb1e2e3b9d2bd0cab0d2d95797248dd95f02d7299e9181900360200190a150565b6007602052600090815260409020546001600160a01b031681565b610f2f33610fb6565b6001546001600160a01b0382811691161461069057600180546001600160a01b0383166001600160a01b0319909116811790915560408051918252517f5425363a03f182281120f5919107c49c7a1a623acc1cbc6df468b6f0c11fcf8c9181900360200190a150565b6001546001600160a01b031681565b6009546001600160a01b031681565b6001546001600160a01b03828116911614610690576040805162461bcd60e51b8152602060048201526002602482015261316760f01b604482015290519081900360640190fd5b7f8e94fed44239eb2314ab7a406345e6c5a8f0ccedf3b600de3d004e672c33abf4805460019091558015610690576040805162461bcd60e51b815260206004820152600260248201526118a160f11b604482015290519081900360640190fdfea26469706673582212208296ed07e151b8f63b8a3a49e798b48f7b2570b7fa1972307ce6369dbc8c7d5664736f6c63430007060033"

// DeployGovernance deploys a new Ethereum zkbnb, binding an instance of Governance to it.
func DeployGovernance(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Governance, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GovernanceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Governance{GovernanceCaller: GovernanceCaller{contract: contract}, GovernanceTransactor: GovernanceTransactor{contract: contract}, GovernanceFilterer: GovernanceFilterer{contract: contract}}, nil
}

// Governance is an auto generated Go binding around an Ethereum zkbnb.
type Governance struct {
	GovernanceCaller     // Read-only binding to the zkbnb
	GovernanceTransactor // Write-only binding to the zkbnb
	GovernanceFilterer   // Log filterer for zkbnb events
}

// GovernanceCaller is an auto generated read-only Go binding around an Ethereum zkbnb.
type GovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceTransactor is an auto generated write-only Go binding around an Ethereum zkbnb.
type GovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceFilterer is an auto generated log filtering Go binding around an Ethereum zkbnb events.
type GovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceSession is an auto generated Go binding around an Ethereum zkbnb,
// with pre-set call and transact options.
type GovernanceSession struct {
	Contract     *Governance       // Generic zkbnb binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernanceCallerSession is an auto generated read-only Go binding around an Ethereum zkbnb,
// with pre-set call options.
type GovernanceCallerSession struct {
	Contract *GovernanceCaller // Generic zkbnb caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum zkbnb,
// with pre-set transact options.
type GovernanceTransactorSession struct {
	Contract     *GovernanceTransactor // Generic zkbnb transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GovernanceRaw is an auto generated low-level Go binding around an Ethereum zkbnb.
type GovernanceRaw struct {
	Contract *Governance // Generic zkbnb binding to access the raw methods on
}

// GovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum zkbnb.
type GovernanceCallerRaw struct {
	Contract *GovernanceCaller // Generic read-only zkbnb binding to access the raw methods on
}

// GovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum zkbnb.
type GovernanceTransactorRaw struct {
	Contract *GovernanceTransactor // Generic write-only zkbnb binding to access the raw methods on
}

// NewGovernance creates a new instance of Governance, bound to a specific deployed zkbnb.
func NewGovernance(address common.Address, backend bind.ContractBackend) (*Governance, error) {
	contract, err := bindGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Governance{GovernanceCaller: GovernanceCaller{contract: contract}, GovernanceTransactor: GovernanceTransactor{contract: contract}, GovernanceFilterer: GovernanceFilterer{contract: contract}}, nil
}

// NewGovernanceCaller creates a new read-only instance of Governance, bound to a specific deployed zkbnb.
func NewGovernanceCaller(address common.Address, caller bind.ContractCaller) (*GovernanceCaller, error) {
	contract, err := bindGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceCaller{contract: contract}, nil
}

// NewGovernanceTransactor creates a new write-only instance of Governance, bound to a specific deployed zkbnb.
func NewGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernanceTransactor, error) {
	contract, err := bindGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceTransactor{contract: contract}, nil
}

// NewGovernanceFilterer creates a new log filterer instance of Governance, bound to a specific deployed zkbnb.
func NewGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernanceFilterer, error) {
	contract, err := bindGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernanceFilterer{contract: contract}, nil
}

// bindGovernance binds a generic wrapper to an already deployed zkbnb.
func bindGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) zkbnb method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governance *GovernanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Governance.Contract.GovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the zkbnb, calling
// its default method if one is available.
func (_Governance *GovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.Contract.GovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) zkbnb method with params as input values.
func (_Governance *GovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governance.Contract.GovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) zkbnb method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governance *GovernanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Governance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the zkbnb, calling
// its default method if one is available.
func (_Governance *GovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) zkbnb method with params as input values.
func (_Governance *GovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governance.Contract.contract.Transact(opts, method, params...)
}

// AssetAddresses is a free data retrieval call binding the zkbnb method 0xdbfc2967.
//
// Solidity: function assetAddresses(uint16 ) view returns(address)
func (_Governance *GovernanceCaller) AssetAddresses(opts *bind.CallOpts, arg0 uint16) (common.Address, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "assetAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetAddresses is a free data retrieval call binding the zkbnb method 0xdbfc2967.
//
// Solidity: function assetAddresses(uint16 ) view returns(address)
func (_Governance *GovernanceSession) AssetAddresses(arg0 uint16) (common.Address, error) {
	return _Governance.Contract.AssetAddresses(&_Governance.CallOpts, arg0)
}

// AssetAddresses is a free data retrieval call binding the zkbnb method 0xdbfc2967.
//
// Solidity: function assetAddresses(uint16 ) view returns(address)
func (_Governance *GovernanceCallerSession) AssetAddresses(arg0 uint16) (common.Address, error) {
	return _Governance.Contract.AssetAddresses(&_Governance.CallOpts, arg0)
}

// AssetGovernance is a free data retrieval call binding the zkbnb method 0xf5e7d6fd.
//
// Solidity: function assetGovernance() view returns(address)
func (_Governance *GovernanceCaller) AssetGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "assetGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetGovernance is a free data retrieval call binding the zkbnb method 0xf5e7d6fd.
//
// Solidity: function assetGovernance() view returns(address)
func (_Governance *GovernanceSession) AssetGovernance() (common.Address, error) {
	return _Governance.Contract.AssetGovernance(&_Governance.CallOpts)
}

// AssetGovernance is a free data retrieval call binding the zkbnb method 0xf5e7d6fd.
//
// Solidity: function assetGovernance() view returns(address)
func (_Governance *GovernanceCallerSession) AssetGovernance() (common.Address, error) {
	return _Governance.Contract.AssetGovernance(&_Governance.CallOpts)
}

// AssetsList is a free data retrieval call binding the zkbnb method 0x1e763ee3.
//
// Solidity: function assetsList(address ) view returns(uint16)
func (_Governance *GovernanceCaller) AssetsList(opts *bind.CallOpts, arg0 common.Address) (uint16, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "assetsList", arg0)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// AssetsList is a free data retrieval call binding the zkbnb method 0x1e763ee3.
//
// Solidity: function assetsList(address ) view returns(uint16)
func (_Governance *GovernanceSession) AssetsList(arg0 common.Address) (uint16, error) {
	return _Governance.Contract.AssetsList(&_Governance.CallOpts, arg0)
}

// AssetsList is a free data retrieval call binding the zkbnb method 0x1e763ee3.
//
// Solidity: function assetsList(address ) view returns(uint16)
func (_Governance *GovernanceCallerSession) AssetsList(arg0 common.Address) (uint16, error) {
	return _Governance.Contract.AssetsList(&_Governance.CallOpts, arg0)
}

// Committers is a free data retrieval call binding the zkbnb method 0x5f4fea26.
//
// Solidity: function committers(address ) view returns(bool)
func (_Governance *GovernanceCaller) Committers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "committers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Committers is a free data retrieval call binding the zkbnb method 0x5f4fea26.
//
// Solidity: function committers(address ) view returns(bool)
func (_Governance *GovernanceSession) Committers(arg0 common.Address) (bool, error) {
	return _Governance.Contract.Committers(&_Governance.CallOpts, arg0)
}

// Committers is a free data retrieval call binding the zkbnb method 0x5f4fea26.
//
// Solidity: function committers(address ) view returns(bool)
func (_Governance *GovernanceCallerSession) Committers(arg0 common.Address) (bool, error) {
	return _Governance.Contract.Committers(&_Governance.CallOpts, arg0)
}

// IsAddressExists is a free data retrieval call binding the zkbnb method 0x321e182b.
//
// Solidity: function isAddressExists(address ) view returns(bool)
func (_Governance *GovernanceCaller) IsAddressExists(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "isAddressExists", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAddressExists is a free data retrieval call binding the zkbnb method 0x321e182b.
//
// Solidity: function isAddressExists(address ) view returns(bool)
func (_Governance *GovernanceSession) IsAddressExists(arg0 common.Address) (bool, error) {
	return _Governance.Contract.IsAddressExists(&_Governance.CallOpts, arg0)
}

// IsAddressExists is a free data retrieval call binding the zkbnb method 0x321e182b.
//
// Solidity: function isAddressExists(address ) view returns(bool)
func (_Governance *GovernanceCallerSession) IsAddressExists(arg0 common.Address) (bool, error) {
	return _Governance.Contract.IsAddressExists(&_Governance.CallOpts, arg0)
}

// Monitors is a free data retrieval call binding the zkbnb method 0x28b58e1a.
//
// Solidity: function monitors(address ) view returns(bool)
func (_Governance *GovernanceCaller) Monitors(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "monitors", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Monitors is a free data retrieval call binding the zkbnb method 0x28b58e1a.
//
// Solidity: function monitors(address ) view returns(bool)
func (_Governance *GovernanceSession) Monitors(arg0 common.Address) (bool, error) {
	return _Governance.Contract.Monitors(&_Governance.CallOpts, arg0)
}

// Monitors is a free data retrieval call binding the zkbnb method 0x28b58e1a.
//
// Solidity: function monitors(address ) view returns(bool)
func (_Governance *GovernanceCallerSession) Monitors(arg0 common.Address) (bool, error) {
	return _Governance.Contract.Monitors(&_Governance.CallOpts, arg0)
}

// NetworkGovernor is a free data retrieval call binding the zkbnb method 0xf39349ef.
//
// Solidity: function networkGovernor() view returns(address)
func (_Governance *GovernanceCaller) NetworkGovernor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "networkGovernor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NetworkGovernor is a free data retrieval call binding the zkbnb method 0xf39349ef.
//
// Solidity: function networkGovernor() view returns(address)
func (_Governance *GovernanceSession) NetworkGovernor() (common.Address, error) {
	return _Governance.Contract.NetworkGovernor(&_Governance.CallOpts)
}

// NetworkGovernor is a free data retrieval call binding the zkbnb method 0xf39349ef.
//
// Solidity: function networkGovernor() view returns(address)
func (_Governance *GovernanceCallerSession) NetworkGovernor() (common.Address, error) {
	return _Governance.Contract.NetworkGovernor(&_Governance.CallOpts)
}

// PausedAssets is a free data retrieval call binding the zkbnb method 0x31d8687b.
//
// Solidity: function pausedAssets(uint16 ) view returns(bool)
func (_Governance *GovernanceCaller) PausedAssets(opts *bind.CallOpts, arg0 uint16) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "pausedAssets", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PausedAssets is a free data retrieval call binding the zkbnb method 0x31d8687b.
//
// Solidity: function pausedAssets(uint16 ) view returns(bool)
func (_Governance *GovernanceSession) PausedAssets(arg0 uint16) (bool, error) {
	return _Governance.Contract.PausedAssets(&_Governance.CallOpts, arg0)
}

// PausedAssets is a free data retrieval call binding the zkbnb method 0x31d8687b.
//
// Solidity: function pausedAssets(uint16 ) view returns(bool)
func (_Governance *GovernanceCallerSession) PausedAssets(arg0 uint16) (bool, error) {
	return _Governance.Contract.PausedAssets(&_Governance.CallOpts, arg0)
}

// RequireActiveCommitter is a free data retrieval call binding the zkbnb method 0xa28e0032.
//
// Solidity: function requireActiveCommitter(address _address) view returns()
func (_Governance *GovernanceCaller) RequireActiveCommitter(opts *bind.CallOpts, _address common.Address) error {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "requireActiveCommitter", _address)

	if err != nil {
		return err
	}

	return err

}

// RequireActiveCommitter is a free data retrieval call binding the zkbnb method 0xa28e0032.
//
// Solidity: function requireActiveCommitter(address _address) view returns()
func (_Governance *GovernanceSession) RequireActiveCommitter(_address common.Address) error {
	return _Governance.Contract.RequireActiveCommitter(&_Governance.CallOpts, _address)
}

// RequireActiveCommitter is a free data retrieval call binding the zkbnb method 0xa28e0032.
//
// Solidity: function requireActiveCommitter(address _address) view returns()
func (_Governance *GovernanceCallerSession) RequireActiveCommitter(_address common.Address) error {
	return _Governance.Contract.RequireActiveCommitter(&_Governance.CallOpts, _address)
}

// RequireActiveMonitor is a free data retrieval call binding the zkbnb method 0x1c021956.
//
// Solidity: function requireActiveMonitor(address _address) view returns()
func (_Governance *GovernanceCaller) RequireActiveMonitor(opts *bind.CallOpts, _address common.Address) error {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "requireActiveMonitor", _address)

	if err != nil {
		return err
	}

	return err

}

// RequireActiveMonitor is a free data retrieval call binding the zkbnb method 0x1c021956.
//
// Solidity: function requireActiveMonitor(address _address) view returns()
func (_Governance *GovernanceSession) RequireActiveMonitor(_address common.Address) error {
	return _Governance.Contract.RequireActiveMonitor(&_Governance.CallOpts, _address)
}

// RequireActiveMonitor is a free data retrieval call binding the zkbnb method 0x1c021956.
//
// Solidity: function requireActiveMonitor(address _address) view returns()
func (_Governance *GovernanceCallerSession) RequireActiveMonitor(_address common.Address) error {
	return _Governance.Contract.RequireActiveMonitor(&_Governance.CallOpts, _address)
}

// RequireActiveVerifier is a free data retrieval call binding the zkbnb method 0x7a3459d4.
//
// Solidity: function requireActiveVerifier(address _address) view returns()
func (_Governance *GovernanceCaller) RequireActiveVerifier(opts *bind.CallOpts, _address common.Address) error {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "requireActiveVerifier", _address)

	if err != nil {
		return err
	}

	return err

}

// RequireActiveVerifier is a free data retrieval call binding the zkbnb method 0x7a3459d4.
//
// Solidity: function requireActiveVerifier(address _address) view returns()
func (_Governance *GovernanceSession) RequireActiveVerifier(_address common.Address) error {
	return _Governance.Contract.RequireActiveVerifier(&_Governance.CallOpts, _address)
}

// RequireActiveVerifier is a free data retrieval call binding the zkbnb method 0x7a3459d4.
//
// Solidity: function requireActiveVerifier(address _address) view returns()
func (_Governance *GovernanceCallerSession) RequireActiveVerifier(_address common.Address) error {
	return _Governance.Contract.RequireActiveVerifier(&_Governance.CallOpts, _address)
}

// RequireGovernor is a free data retrieval call binding the zkbnb method 0xf5f84ed4.
//
// Solidity: function requireGovernor(address _address) view returns()
func (_Governance *GovernanceCaller) RequireGovernor(opts *bind.CallOpts, _address common.Address) error {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "requireGovernor", _address)

	if err != nil {
		return err
	}

	return err

}

// RequireGovernor is a free data retrieval call binding the zkbnb method 0xf5f84ed4.
//
// Solidity: function requireGovernor(address _address) view returns()
func (_Governance *GovernanceSession) RequireGovernor(_address common.Address) error {
	return _Governance.Contract.RequireGovernor(&_Governance.CallOpts, _address)
}

// RequireGovernor is a free data retrieval call binding the zkbnb method 0xf5f84ed4.
//
// Solidity: function requireGovernor(address _address) view returns()
func (_Governance *GovernanceCallerSession) RequireGovernor(_address common.Address) error {
	return _Governance.Contract.RequireGovernor(&_Governance.CallOpts, _address)
}

// TotalAssets is a free data retrieval call binding the zkbnb method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint16)
func (_Governance *GovernanceCaller) TotalAssets(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the zkbnb method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint16)
func (_Governance *GovernanceSession) TotalAssets() (uint16, error) {
	return _Governance.Contract.TotalAssets(&_Governance.CallOpts)
}

// TotalAssets is a free data retrieval call binding the zkbnb method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint16)
func (_Governance *GovernanceCallerSession) TotalAssets() (uint16, error) {
	return _Governance.Contract.TotalAssets(&_Governance.CallOpts)
}

// ValidateAssetAddress is a free data retrieval call binding the zkbnb method 0x9bd77609.
//
// Solidity: function validateAssetAddress(address _assetAddr) view returns(uint16)
func (_Governance *GovernanceCaller) ValidateAssetAddress(opts *bind.CallOpts, _assetAddr common.Address) (uint16, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "validateAssetAddress", _assetAddr)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ValidateAssetAddress is a free data retrieval call binding the zkbnb method 0x9bd77609.
//
// Solidity: function validateAssetAddress(address _assetAddr) view returns(uint16)
func (_Governance *GovernanceSession) ValidateAssetAddress(_assetAddr common.Address) (uint16, error) {
	return _Governance.Contract.ValidateAssetAddress(&_Governance.CallOpts, _assetAddr)
}

// ValidateAssetAddress is a free data retrieval call binding the zkbnb method 0x9bd77609.
//
// Solidity: function validateAssetAddress(address _assetAddr) view returns(uint16)
func (_Governance *GovernanceCallerSession) ValidateAssetAddress(_assetAddr common.Address) (uint16, error) {
	return _Governance.Contract.ValidateAssetAddress(&_Governance.CallOpts, _assetAddr)
}

// Verifiers is a free data retrieval call binding the zkbnb method 0x6c824487.
//
// Solidity: function verifiers(address ) view returns(bool)
func (_Governance *GovernanceCaller) Verifiers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "verifiers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verifiers is a free data retrieval call binding the zkbnb method 0x6c824487.
//
// Solidity: function verifiers(address ) view returns(bool)
func (_Governance *GovernanceSession) Verifiers(arg0 common.Address) (bool, error) {
	return _Governance.Contract.Verifiers(&_Governance.CallOpts, arg0)
}

// Verifiers is a free data retrieval call binding the zkbnb method 0x6c824487.
//
// Solidity: function verifiers(address ) view returns(bool)
func (_Governance *GovernanceCallerSession) Verifiers(arg0 common.Address) (bool, error) {
	return _Governance.Contract.Verifiers(&_Governance.CallOpts, arg0)
}

// ChangeAssetGovernance is a paid mutator transaction binding the zkbnb method 0xd87e3748.
//
// Solidity: function changeAssetGovernance(address _newAssetGovernance) returns()
func (_Governance *GovernanceTransactor) ChangeAssetGovernance(opts *bind.TransactOpts, _newAssetGovernance common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "changeAssetGovernance", _newAssetGovernance)
}

// ChangeAssetGovernance is a paid mutator transaction binding the zkbnb method 0xd87e3748.
//
// Solidity: function changeAssetGovernance(address _newAssetGovernance) returns()
func (_Governance *GovernanceSession) ChangeAssetGovernance(_newAssetGovernance common.Address) (*types.Transaction, error) {
	return _Governance.Contract.ChangeAssetGovernance(&_Governance.TransactOpts, _newAssetGovernance)
}

// ChangeAssetGovernance is a paid mutator transaction binding the zkbnb method 0xd87e3748.
//
// Solidity: function changeAssetGovernance(address _newAssetGovernance) returns()
func (_Governance *GovernanceTransactorSession) ChangeAssetGovernance(_newAssetGovernance common.Address) (*types.Transaction, error) {
	return _Governance.Contract.ChangeAssetGovernance(&_Governance.TransactOpts, _newAssetGovernance)
}

// ChangeGovernor is a paid mutator transaction binding the zkbnb method 0xe4c0aaf4.
//
// Solidity: function changeGovernor(address _newGovernor) returns()
func (_Governance *GovernanceTransactor) ChangeGovernor(opts *bind.TransactOpts, _newGovernor common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "changeGovernor", _newGovernor)
}

// ChangeGovernor is a paid mutator transaction binding the zkbnb method 0xe4c0aaf4.
//
// Solidity: function changeGovernor(address _newGovernor) returns()
func (_Governance *GovernanceSession) ChangeGovernor(_newGovernor common.Address) (*types.Transaction, error) {
	return _Governance.Contract.ChangeGovernor(&_Governance.TransactOpts, _newGovernor)
}

// ChangeGovernor is a paid mutator transaction binding the zkbnb method 0xe4c0aaf4.
//
// Solidity: function changeGovernor(address _newGovernor) returns()
func (_Governance *GovernanceTransactorSession) ChangeGovernor(_newGovernor common.Address) (*types.Transaction, error) {
	return _Governance.Contract.ChangeGovernor(&_Governance.TransactOpts, _newGovernor)
}

// Initialize is a paid mutator transaction binding the zkbnb method 0x439fab91.
//
// Solidity: function initialize(bytes initializationParameters) returns()
func (_Governance *GovernanceTransactor) Initialize(opts *bind.TransactOpts, initializationParameters []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "initialize", initializationParameters)
}

// Initialize is a paid mutator transaction binding the zkbnb method 0x439fab91.
//
// Solidity: function initialize(bytes initializationParameters) returns()
func (_Governance *GovernanceSession) Initialize(initializationParameters []byte) (*types.Transaction, error) {
	return _Governance.Contract.Initialize(&_Governance.TransactOpts, initializationParameters)
}

// Initialize is a paid mutator transaction binding the zkbnb method 0x439fab91.
//
// Solidity: function initialize(bytes initializationParameters) returns()
func (_Governance *GovernanceTransactorSession) Initialize(initializationParameters []byte) (*types.Transaction, error) {
	return _Governance.Contract.Initialize(&_Governance.TransactOpts, initializationParameters)
}

// SetAsset is a paid mutator transaction binding the zkbnb method 0x5bab15ca.
//
// Solidity: function setAsset(uint16 _assetId, address _assetAddress) returns()
func (_Governance *GovernanceTransactor) SetAsset(opts *bind.TransactOpts, _assetId uint16, _assetAddress common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setAsset", _assetId, _assetAddress)
}

// SetAsset is a paid mutator transaction binding the zkbnb method 0x5bab15ca.
//
// Solidity: function setAsset(uint16 _assetId, address _assetAddress) returns()
func (_Governance *GovernanceSession) SetAsset(_assetId uint16, _assetAddress common.Address) (*types.Transaction, error) {
	return _Governance.Contract.SetAsset(&_Governance.TransactOpts, _assetId, _assetAddress)
}

// SetAsset is a paid mutator transaction binding the zkbnb method 0x5bab15ca.
//
// Solidity: function setAsset(uint16 _assetId, address _assetAddress) returns()
func (_Governance *GovernanceTransactorSession) SetAsset(_assetId uint16, _assetAddress common.Address) (*types.Transaction, error) {
	return _Governance.Contract.SetAsset(&_Governance.TransactOpts, _assetId, _assetAddress)
}

// SetAssetPaused is a paid mutator transaction binding the zkbnb method 0x2520ce5a.
//
// Solidity: function setAssetPaused(address _assetAddress, bool _assetPaused) returns()
func (_Governance *GovernanceTransactor) SetAssetPaused(opts *bind.TransactOpts, _assetAddress common.Address, _assetPaused bool) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setAssetPaused", _assetAddress, _assetPaused)
}

// SetAssetPaused is a paid mutator transaction binding the zkbnb method 0x2520ce5a.
//
// Solidity: function setAssetPaused(address _assetAddress, bool _assetPaused) returns()
func (_Governance *GovernanceSession) SetAssetPaused(_assetAddress common.Address, _assetPaused bool) (*types.Transaction, error) {
	return _Governance.Contract.SetAssetPaused(&_Governance.TransactOpts, _assetAddress, _assetPaused)
}

// SetAssetPaused is a paid mutator transaction binding the zkbnb method 0x2520ce5a.
//
// Solidity: function setAssetPaused(address _assetAddress, bool _assetPaused) returns()
func (_Governance *GovernanceTransactorSession) SetAssetPaused(_assetAddress common.Address, _assetPaused bool) (*types.Transaction, error) {
	return _Governance.Contract.SetAssetPaused(&_Governance.TransactOpts, _assetAddress, _assetPaused)
}

// SetCommitter is a paid mutator transaction binding the zkbnb method 0x1afcac96.
//
// Solidity: function setCommitter(address _validator, bool _active) returns()
func (_Governance *GovernanceTransactor) SetCommitter(opts *bind.TransactOpts, _validator common.Address, _active bool) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setCommitter", _validator, _active)
}

// SetCommitter is a paid mutator transaction binding the zkbnb method 0x1afcac96.
//
// Solidity: function setCommitter(address _validator, bool _active) returns()
func (_Governance *GovernanceSession) SetCommitter(_validator common.Address, _active bool) (*types.Transaction, error) {
	return _Governance.Contract.SetCommitter(&_Governance.TransactOpts, _validator, _active)
}

// SetCommitter is a paid mutator transaction binding the zkbnb method 0x1afcac96.
//
// Solidity: function setCommitter(address _validator, bool _active) returns()
func (_Governance *GovernanceTransactorSession) SetCommitter(_validator common.Address, _active bool) (*types.Transaction, error) {
	return _Governance.Contract.SetCommitter(&_Governance.TransactOpts, _validator, _active)
}

// SetMonitor is a paid mutator transaction binding the zkbnb method 0x26eb0f4a.
//
// Solidity: function setMonitor(address _validator, bool _active) returns()
func (_Governance *GovernanceTransactor) SetMonitor(opts *bind.TransactOpts, _validator common.Address, _active bool) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setMonitor", _validator, _active)
}

// SetMonitor is a paid mutator transaction binding the zkbnb method 0x26eb0f4a.
//
// Solidity: function setMonitor(address _validator, bool _active) returns()
func (_Governance *GovernanceSession) SetMonitor(_validator common.Address, _active bool) (*types.Transaction, error) {
	return _Governance.Contract.SetMonitor(&_Governance.TransactOpts, _validator, _active)
}

// SetMonitor is a paid mutator transaction binding the zkbnb method 0x26eb0f4a.
//
// Solidity: function setMonitor(address _validator, bool _active) returns()
func (_Governance *GovernanceTransactorSession) SetMonitor(_validator common.Address, _active bool) (*types.Transaction, error) {
	return _Governance.Contract.SetMonitor(&_Governance.TransactOpts, _validator, _active)
}

// SetVerifier is a paid mutator transaction binding the zkbnb method 0xac9b5671.
//
// Solidity: function setVerifier(address _validator, bool _active) returns()
func (_Governance *GovernanceTransactor) SetVerifier(opts *bind.TransactOpts, _validator common.Address, _active bool) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setVerifier", _validator, _active)
}

// SetVerifier is a paid mutator transaction binding the zkbnb method 0xac9b5671.
//
// Solidity: function setVerifier(address _validator, bool _active) returns()
func (_Governance *GovernanceSession) SetVerifier(_validator common.Address, _active bool) (*types.Transaction, error) {
	return _Governance.Contract.SetVerifier(&_Governance.TransactOpts, _validator, _active)
}

// SetVerifier is a paid mutator transaction binding the zkbnb method 0xac9b5671.
//
// Solidity: function setVerifier(address _validator, bool _active) returns()
func (_Governance *GovernanceTransactorSession) SetVerifier(_validator common.Address, _active bool) (*types.Transaction, error) {
	return _Governance.Contract.SetVerifier(&_Governance.TransactOpts, _validator, _active)
}

// Upgrade is a paid mutator transaction binding the zkbnb method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_Governance *GovernanceTransactor) Upgrade(opts *bind.TransactOpts, upgradeParameters []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "upgrade", upgradeParameters)
}

// Upgrade is a paid mutator transaction binding the zkbnb method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_Governance *GovernanceSession) Upgrade(upgradeParameters []byte) (*types.Transaction, error) {
	return _Governance.Contract.Upgrade(&_Governance.TransactOpts, upgradeParameters)
}

// Upgrade is a paid mutator transaction binding the zkbnb method 0x25394645.
//
// Solidity: function upgrade(bytes upgradeParameters) returns()
func (_Governance *GovernanceTransactorSession) Upgrade(upgradeParameters []byte) (*types.Transaction, error) {
	return _Governance.Contract.Upgrade(&_Governance.TransactOpts, upgradeParameters)
}

// GovernanceAssetPausedUpdateIterator is returned from FilterAssetPausedUpdate and is used to iterate over the raw logs and unpacked data for AssetPausedUpdate events raised by the Governance zkbnb.
type GovernanceAssetPausedUpdateIterator struct {
	Event *GovernanceAssetPausedUpdate // Event containing the zkbnb specifics and raw log

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
func (it *GovernanceAssetPausedUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceAssetPausedUpdate)
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
		it.Event = new(GovernanceAssetPausedUpdate)
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
func (it *GovernanceAssetPausedUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceAssetPausedUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceAssetPausedUpdate represents a AssetPausedUpdate event raised by the Governance zkbnb.
type GovernanceAssetPausedUpdate struct {
	Token  common.Address
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAssetPausedUpdate is a free log retrieval operation binding the zkbnb event 0xf7ca5545623b85829d3abcb6729eb021070bbe93cba64f5c4de6b17b805b7d02.
//
// Solidity: event AssetPausedUpdate(address indexed token, bool paused)
func (_Governance *GovernanceFilterer) FilterAssetPausedUpdate(opts *bind.FilterOpts, token []common.Address) (*GovernanceAssetPausedUpdateIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "AssetPausedUpdate", tokenRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceAssetPausedUpdateIterator{contract: _Governance.contract, event: "AssetPausedUpdate", logs: logs, sub: sub}, nil
}

// WatchAssetPausedUpdate is a free log subscription operation binding the zkbnb event 0xf7ca5545623b85829d3abcb6729eb021070bbe93cba64f5c4de6b17b805b7d02.
//
// Solidity: event AssetPausedUpdate(address indexed token, bool paused)
func (_Governance *GovernanceFilterer) WatchAssetPausedUpdate(opts *bind.WatchOpts, sink chan<- *GovernanceAssetPausedUpdate, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "AssetPausedUpdate", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceAssetPausedUpdate)
				if err := _Governance.contract.UnpackLog(event, "AssetPausedUpdate", log); err != nil {
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

// ParseAssetPausedUpdate is a log parse operation binding the zkbnb event 0xf7ca5545623b85829d3abcb6729eb021070bbe93cba64f5c4de6b17b805b7d02.
//
// Solidity: event AssetPausedUpdate(address indexed token, bool paused)
func (_Governance *GovernanceFilterer) ParseAssetPausedUpdate(log types.Log) (*GovernanceAssetPausedUpdate, error) {
	event := new(GovernanceAssetPausedUpdate)
	if err := _Governance.contract.UnpackLog(event, "AssetPausedUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceCommitterStatusUpdateIterator is returned from FilterCommitterStatusUpdate and is used to iterate over the raw logs and unpacked data for CommitterStatusUpdate events raised by the Governance zkbnb.
type GovernanceCommitterStatusUpdateIterator struct {
	Event *GovernanceCommitterStatusUpdate // Event containing the zkbnb specifics and raw log

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
func (it *GovernanceCommitterStatusUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceCommitterStatusUpdate)
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
		it.Event = new(GovernanceCommitterStatusUpdate)
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
func (it *GovernanceCommitterStatusUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceCommitterStatusUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceCommitterStatusUpdate represents a CommitterStatusUpdate event raised by the Governance zkbnb.
type GovernanceCommitterStatusUpdate struct {
	CommitterAddress common.Address
	IsActive         bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCommitterStatusUpdate is a free log retrieval operation binding the zkbnb event 0xd158a06928d89556421a4e2eabd98c7ea10e974c07e399e3afc449d19eb391b2.
//
// Solidity: event CommitterStatusUpdate(address indexed committerAddress, bool isActive)
func (_Governance *GovernanceFilterer) FilterCommitterStatusUpdate(opts *bind.FilterOpts, committerAddress []common.Address) (*GovernanceCommitterStatusUpdateIterator, error) {

	var committerAddressRule []interface{}
	for _, committerAddressItem := range committerAddress {
		committerAddressRule = append(committerAddressRule, committerAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "CommitterStatusUpdate", committerAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceCommitterStatusUpdateIterator{contract: _Governance.contract, event: "CommitterStatusUpdate", logs: logs, sub: sub}, nil
}

// WatchCommitterStatusUpdate is a free log subscription operation binding the zkbnb event 0xd158a06928d89556421a4e2eabd98c7ea10e974c07e399e3afc449d19eb391b2.
//
// Solidity: event CommitterStatusUpdate(address indexed committerAddress, bool isActive)
func (_Governance *GovernanceFilterer) WatchCommitterStatusUpdate(opts *bind.WatchOpts, sink chan<- *GovernanceCommitterStatusUpdate, committerAddress []common.Address) (event.Subscription, error) {

	var committerAddressRule []interface{}
	for _, committerAddressItem := range committerAddress {
		committerAddressRule = append(committerAddressRule, committerAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "CommitterStatusUpdate", committerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceCommitterStatusUpdate)
				if err := _Governance.contract.UnpackLog(event, "CommitterStatusUpdate", log); err != nil {
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

// ParseCommitterStatusUpdate is a log parse operation binding the zkbnb event 0xd158a06928d89556421a4e2eabd98c7ea10e974c07e399e3afc449d19eb391b2.
//
// Solidity: event CommitterStatusUpdate(address indexed committerAddress, bool isActive)
func (_Governance *GovernanceFilterer) ParseCommitterStatusUpdate(log types.Log) (*GovernanceCommitterStatusUpdate, error) {
	event := new(GovernanceCommitterStatusUpdate)
	if err := _Governance.contract.UnpackLog(event, "CommitterStatusUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceExecutorStatusUpdateIterator is returned from FilterExecutorStatusUpdate and is used to iterate over the raw logs and unpacked data for ExecutorStatusUpdate events raised by the Governance zkbnb.
type GovernanceExecutorStatusUpdateIterator struct {
	Event *GovernanceExecutorStatusUpdate // Event containing the zkbnb specifics and raw log

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
func (it *GovernanceExecutorStatusUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceExecutorStatusUpdate)
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
		it.Event = new(GovernanceExecutorStatusUpdate)
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
func (it *GovernanceExecutorStatusUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceExecutorStatusUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceExecutorStatusUpdate represents a ExecutorStatusUpdate event raised by the Governance zkbnb.
type GovernanceExecutorStatusUpdate struct {
	ExecutorAddress common.Address
	IsActive        bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterExecutorStatusUpdate is a free log retrieval operation binding the zkbnb event 0x9a01f9639f4fd4c2eea316eabcc61b3674b410002afdb96bf605cf288d37f9da.
//
// Solidity: event ExecutorStatusUpdate(address indexed executorAddress, bool isActive)
func (_Governance *GovernanceFilterer) FilterExecutorStatusUpdate(opts *bind.FilterOpts, executorAddress []common.Address) (*GovernanceExecutorStatusUpdateIterator, error) {

	var executorAddressRule []interface{}
	for _, executorAddressItem := range executorAddress {
		executorAddressRule = append(executorAddressRule, executorAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ExecutorStatusUpdate", executorAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceExecutorStatusUpdateIterator{contract: _Governance.contract, event: "ExecutorStatusUpdate", logs: logs, sub: sub}, nil
}

// WatchExecutorStatusUpdate is a free log subscription operation binding the zkbnb event 0x9a01f9639f4fd4c2eea316eabcc61b3674b410002afdb96bf605cf288d37f9da.
//
// Solidity: event ExecutorStatusUpdate(address indexed executorAddress, bool isActive)
func (_Governance *GovernanceFilterer) WatchExecutorStatusUpdate(opts *bind.WatchOpts, sink chan<- *GovernanceExecutorStatusUpdate, executorAddress []common.Address) (event.Subscription, error) {

	var executorAddressRule []interface{}
	for _, executorAddressItem := range executorAddress {
		executorAddressRule = append(executorAddressRule, executorAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ExecutorStatusUpdate", executorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceExecutorStatusUpdate)
				if err := _Governance.contract.UnpackLog(event, "ExecutorStatusUpdate", log); err != nil {
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

// ParseExecutorStatusUpdate is a log parse operation binding the zkbnb event 0x9a01f9639f4fd4c2eea316eabcc61b3674b410002afdb96bf605cf288d37f9da.
//
// Solidity: event ExecutorStatusUpdate(address indexed executorAddress, bool isActive)
func (_Governance *GovernanceFilterer) ParseExecutorStatusUpdate(log types.Log) (*GovernanceExecutorStatusUpdate, error) {
	event := new(GovernanceExecutorStatusUpdate)
	if err := _Governance.contract.UnpackLog(event, "ExecutorStatusUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceMonitorStatusUpdateIterator is returned from FilterMonitorStatusUpdate and is used to iterate over the raw logs and unpacked data for MonitorStatusUpdate events raised by the Governance zkbnb.
type GovernanceMonitorStatusUpdateIterator struct {
	Event *GovernanceMonitorStatusUpdate // Event containing the zkbnb specifics and raw log

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
func (it *GovernanceMonitorStatusUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceMonitorStatusUpdate)
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
		it.Event = new(GovernanceMonitorStatusUpdate)
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
func (it *GovernanceMonitorStatusUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceMonitorStatusUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceMonitorStatusUpdate represents a MonitorStatusUpdate event raised by the Governance zkbnb.
type GovernanceMonitorStatusUpdate struct {
	ExecutorAddress common.Address
	IsActive        bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMonitorStatusUpdate is a free log retrieval operation binding the zkbnb event 0x6fbe3fec04b6d2f79e569f1a46ec4d1621e751112483b5ce9d81832f19fee5c1.
//
// Solidity: event MonitorStatusUpdate(address indexed executorAddress, bool isActive)
func (_Governance *GovernanceFilterer) FilterMonitorStatusUpdate(opts *bind.FilterOpts, executorAddress []common.Address) (*GovernanceMonitorStatusUpdateIterator, error) {

	var executorAddressRule []interface{}
	for _, executorAddressItem := range executorAddress {
		executorAddressRule = append(executorAddressRule, executorAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "MonitorStatusUpdate", executorAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceMonitorStatusUpdateIterator{contract: _Governance.contract, event: "MonitorStatusUpdate", logs: logs, sub: sub}, nil
}

// WatchMonitorStatusUpdate is a free log subscription operation binding the zkbnb event 0x6fbe3fec04b6d2f79e569f1a46ec4d1621e751112483b5ce9d81832f19fee5c1.
//
// Solidity: event MonitorStatusUpdate(address indexed executorAddress, bool isActive)
func (_Governance *GovernanceFilterer) WatchMonitorStatusUpdate(opts *bind.WatchOpts, sink chan<- *GovernanceMonitorStatusUpdate, executorAddress []common.Address) (event.Subscription, error) {

	var executorAddressRule []interface{}
	for _, executorAddressItem := range executorAddress {
		executorAddressRule = append(executorAddressRule, executorAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "MonitorStatusUpdate", executorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceMonitorStatusUpdate)
				if err := _Governance.contract.UnpackLog(event, "MonitorStatusUpdate", log); err != nil {
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

// ParseMonitorStatusUpdate is a log parse operation binding the zkbnb event 0x6fbe3fec04b6d2f79e569f1a46ec4d1621e751112483b5ce9d81832f19fee5c1.
//
// Solidity: event MonitorStatusUpdate(address indexed executorAddress, bool isActive)
func (_Governance *GovernanceFilterer) ParseMonitorStatusUpdate(log types.Log) (*GovernanceMonitorStatusUpdate, error) {
	event := new(GovernanceMonitorStatusUpdate)
	if err := _Governance.contract.UnpackLog(event, "MonitorStatusUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceNewAssetIterator is returned from FilterNewAsset and is used to iterate over the raw logs and unpacked data for NewAsset events raised by the Governance zkbnb.
type GovernanceNewAssetIterator struct {
	Event *GovernanceNewAsset // Event containing the zkbnb specifics and raw log

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
func (it *GovernanceNewAssetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceNewAsset)
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
		it.Event = new(GovernanceNewAsset)
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
func (it *GovernanceNewAssetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceNewAssetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceNewAsset represents a NewAsset event raised by the Governance zkbnb.
type GovernanceNewAsset struct {
	AssetAddress common.Address
	AssetId      uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewAsset is a free log retrieval operation binding the zkbnb event 0x990c18cf226b253448a6a051b5cadf51a61f4ca56703374cdd8bf9df04b2f901.
//
// Solidity: event NewAsset(address indexed assetAddress, uint16 indexed assetId)
func (_Governance *GovernanceFilterer) FilterNewAsset(opts *bind.FilterOpts, assetAddress []common.Address, assetId []uint16) (*GovernanceNewAssetIterator, error) {

	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "NewAsset", assetAddressRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceNewAssetIterator{contract: _Governance.contract, event: "NewAsset", logs: logs, sub: sub}, nil
}

// WatchNewAsset is a free log subscription operation binding the zkbnb event 0x990c18cf226b253448a6a051b5cadf51a61f4ca56703374cdd8bf9df04b2f901.
//
// Solidity: event NewAsset(address indexed assetAddress, uint16 indexed assetId)
func (_Governance *GovernanceFilterer) WatchNewAsset(opts *bind.WatchOpts, sink chan<- *GovernanceNewAsset, assetAddress []common.Address, assetId []uint16) (event.Subscription, error) {

	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "NewAsset", assetAddressRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceNewAsset)
				if err := _Governance.contract.UnpackLog(event, "NewAsset", log); err != nil {
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

// ParseNewAsset is a log parse operation binding the zkbnb event 0x990c18cf226b253448a6a051b5cadf51a61f4ca56703374cdd8bf9df04b2f901.
//
// Solidity: event NewAsset(address indexed assetAddress, uint16 indexed assetId)
func (_Governance *GovernanceFilterer) ParseNewAsset(log types.Log) (*GovernanceNewAsset, error) {
	event := new(GovernanceNewAsset)
	if err := _Governance.contract.UnpackLog(event, "NewAsset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceNewAssetGovernanceIterator is returned from FilterNewAssetGovernance and is used to iterate over the raw logs and unpacked data for NewAssetGovernance events raised by the Governance zkbnb.
type GovernanceNewAssetGovernanceIterator struct {
	Event *GovernanceNewAssetGovernance // Event containing the zkbnb specifics and raw log

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
func (it *GovernanceNewAssetGovernanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceNewAssetGovernance)
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
		it.Event = new(GovernanceNewAssetGovernance)
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
func (it *GovernanceNewAssetGovernanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceNewAssetGovernanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceNewAssetGovernance represents a NewAssetGovernance event raised by the Governance zkbnb.
type GovernanceNewAssetGovernance struct {
	NewAssetGovernance common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewAssetGovernance is a free log retrieval operation binding the zkbnb event 0xc7708091594580aae77e08cb1e2e3b9d2bd0cab0d2d95797248dd95f02d7299e.
//
// Solidity: event NewAssetGovernance(address newAssetGovernance)
func (_Governance *GovernanceFilterer) FilterNewAssetGovernance(opts *bind.FilterOpts) (*GovernanceNewAssetGovernanceIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "NewAssetGovernance")
	if err != nil {
		return nil, err
	}
	return &GovernanceNewAssetGovernanceIterator{contract: _Governance.contract, event: "NewAssetGovernance", logs: logs, sub: sub}, nil
}

// WatchNewAssetGovernance is a free log subscription operation binding the zkbnb event 0xc7708091594580aae77e08cb1e2e3b9d2bd0cab0d2d95797248dd95f02d7299e.
//
// Solidity: event NewAssetGovernance(address newAssetGovernance)
func (_Governance *GovernanceFilterer) WatchNewAssetGovernance(opts *bind.WatchOpts, sink chan<- *GovernanceNewAssetGovernance) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "NewAssetGovernance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceNewAssetGovernance)
				if err := _Governance.contract.UnpackLog(event, "NewAssetGovernance", log); err != nil {
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

// ParseNewAssetGovernance is a log parse operation binding the zkbnb event 0xc7708091594580aae77e08cb1e2e3b9d2bd0cab0d2d95797248dd95f02d7299e.
//
// Solidity: event NewAssetGovernance(address newAssetGovernance)
func (_Governance *GovernanceFilterer) ParseNewAssetGovernance(log types.Log) (*GovernanceNewAssetGovernance, error) {
	event := new(GovernanceNewAssetGovernance)
	if err := _Governance.contract.UnpackLog(event, "NewAssetGovernance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceNewGovernorIterator is returned from FilterNewGovernor and is used to iterate over the raw logs and unpacked data for NewGovernor events raised by the Governance zkbnb.
type GovernanceNewGovernorIterator struct {
	Event *GovernanceNewGovernor // Event containing the zkbnb specifics and raw log

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
func (it *GovernanceNewGovernorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceNewGovernor)
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
		it.Event = new(GovernanceNewGovernor)
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
func (it *GovernanceNewGovernorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceNewGovernorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceNewGovernor represents a NewGovernor event raised by the Governance zkbnb.
type GovernanceNewGovernor struct {
	NewGovernor common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewGovernor is a free log retrieval operation binding the zkbnb event 0x5425363a03f182281120f5919107c49c7a1a623acc1cbc6df468b6f0c11fcf8c.
//
// Solidity: event NewGovernor(address newGovernor)
func (_Governance *GovernanceFilterer) FilterNewGovernor(opts *bind.FilterOpts) (*GovernanceNewGovernorIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "NewGovernor")
	if err != nil {
		return nil, err
	}
	return &GovernanceNewGovernorIterator{contract: _Governance.contract, event: "NewGovernor", logs: logs, sub: sub}, nil
}

// WatchNewGovernor is a free log subscription operation binding the zkbnb event 0x5425363a03f182281120f5919107c49c7a1a623acc1cbc6df468b6f0c11fcf8c.
//
// Solidity: event NewGovernor(address newGovernor)
func (_Governance *GovernanceFilterer) WatchNewGovernor(opts *bind.WatchOpts, sink chan<- *GovernanceNewGovernor) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "NewGovernor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceNewGovernor)
				if err := _Governance.contract.UnpackLog(event, "NewGovernor", log); err != nil {
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

// ParseNewGovernor is a log parse operation binding the zkbnb event 0x5425363a03f182281120f5919107c49c7a1a623acc1cbc6df468b6f0c11fcf8c.
//
// Solidity: event NewGovernor(address newGovernor)
func (_Governance *GovernanceFilterer) ParseNewGovernor(log types.Log) (*GovernanceNewGovernor, error) {
	event := new(GovernanceNewGovernor)
	if err := _Governance.contract.UnpackLog(event, "NewGovernor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceVerifierStatusUpdateIterator is returned from FilterVerifierStatusUpdate and is used to iterate over the raw logs and unpacked data for VerifierStatusUpdate events raised by the Governance zkbnb.
type GovernanceVerifierStatusUpdateIterator struct {
	Event *GovernanceVerifierStatusUpdate // Event containing the zkbnb specifics and raw log

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
func (it *GovernanceVerifierStatusUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceVerifierStatusUpdate)
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
		it.Event = new(GovernanceVerifierStatusUpdate)
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
func (it *GovernanceVerifierStatusUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceVerifierStatusUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceVerifierStatusUpdate represents a VerifierStatusUpdate event raised by the Governance zkbnb.
type GovernanceVerifierStatusUpdate struct {
	VerifierAddress common.Address
	IsActive        bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVerifierStatusUpdate is a free log retrieval operation binding the zkbnb event 0xf09ef994a79ba4638c91b1d9595c89e208c3a9b9403769b68670c678782f82d0.
//
// Solidity: event VerifierStatusUpdate(address indexed verifierAddress, bool isActive)
func (_Governance *GovernanceFilterer) FilterVerifierStatusUpdate(opts *bind.FilterOpts, verifierAddress []common.Address) (*GovernanceVerifierStatusUpdateIterator, error) {

	var verifierAddressRule []interface{}
	for _, verifierAddressItem := range verifierAddress {
		verifierAddressRule = append(verifierAddressRule, verifierAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "VerifierStatusUpdate", verifierAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceVerifierStatusUpdateIterator{contract: _Governance.contract, event: "VerifierStatusUpdate", logs: logs, sub: sub}, nil
}

// WatchVerifierStatusUpdate is a free log subscription operation binding the zkbnb event 0xf09ef994a79ba4638c91b1d9595c89e208c3a9b9403769b68670c678782f82d0.
//
// Solidity: event VerifierStatusUpdate(address indexed verifierAddress, bool isActive)
func (_Governance *GovernanceFilterer) WatchVerifierStatusUpdate(opts *bind.WatchOpts, sink chan<- *GovernanceVerifierStatusUpdate, verifierAddress []common.Address) (event.Subscription, error) {

	var verifierAddressRule []interface{}
	for _, verifierAddressItem := range verifierAddress {
		verifierAddressRule = append(verifierAddressRule, verifierAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "VerifierStatusUpdate", verifierAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceVerifierStatusUpdate)
				if err := _Governance.contract.UnpackLog(event, "VerifierStatusUpdate", log); err != nil {
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

// ParseVerifierStatusUpdate is a log parse operation binding the zkbnb event 0xf09ef994a79ba4638c91b1d9595c89e208c3a9b9403769b68670c678782f82d0.
//
// Solidity: event VerifierStatusUpdate(address indexed verifierAddress, bool isActive)
func (_Governance *GovernanceFilterer) ParseVerifierStatusUpdate(log types.Log) (*GovernanceVerifierStatusUpdate, error) {
	event := new(GovernanceVerifierStatusUpdate)
	if err := _Governance.contract.UnpackLog(event, "VerifierStatusUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
