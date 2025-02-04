// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractMachServiceManager

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

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSSignatureCheckerNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// IMachServiceManagerAlertHeader is an auto generated low-level Go binding around an user-defined struct.
type IMachServiceManagerAlertHeader struct {
	MessageHash                [32]byte
	QuorumNumbers              []byte
	QuorumThresholdPercentages []byte
	ReferenceBlockNumber       uint32
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ContractMachServiceManagerMetaData contains all meta data concerning the ContractMachServiceManager contract.
var ContractMachServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"referenceBlockNumber\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint32[]\",\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"quorumApks\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"apkG2\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"sigma\",\"type\":\"tuple\"},{\"internalType\":\"uint32[]\",\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[][]\",\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\"}],\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\"}],\"name\":\"verifyRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"type\":\"constructor\",\"inputs\":[{\"name\":\"__avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"__registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"__stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"THRESHOLD_DENOMINATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addToAllowlist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"alertConfirmer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowlistEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"confirmAlert\",\"inputs\":[{\"name\":\"alertHeader\",\"type\":\"tuple\",\"internalType\":\"structIMachServiceManager.AlertHeader\",\"components\":[{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentages\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"contains\",\"inputs\":[{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disableAllowlist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enableAllowlist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_alertConfirmer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queryMessageHashes\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"querySize\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumThresholdPercentage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeAlert\",\"inputs\":[{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeFromAllowlist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalAlerts\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateQuorumThresholdPercentage\",\"inputs\":[{\"name\":\"thresholdPercentage\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AlertConfirmed\",\"inputs\":[{\"name\":\"alertHeaderHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AlertConfirmerChanged\",\"inputs\":[{\"name\":\"previousAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AlertRemoved\",\"inputs\":[{\"name\":\"messageHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowlistDisabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowlistEnabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAdded\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAllowed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDisallowed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemoved\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumThresholdPercentageChanged\",\"inputs\":[{\"name\":\"thresholdPercentages\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInAllowlist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientThresholdPercentages\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfirmer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidReferenceBlockNum\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidStartIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInAllowlist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6101606040526006805460ff191660011790553480156200001f57600080fd5b506040516200534d3803806200534d8339810160408190526200004291620002f1565b6001600160a01b0380841660c052808316608052811660a05281838183620000696200020c565b5050506001600160a01b03811660e081905260408051636830483560e01b815290516368304835916004808201926020929091908290030181865afa158015620000b7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000dd919062000345565b6001600160a01b0316610100816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000136573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200015c919062000345565b6001600160a01b0316610120816001600160a01b031681525050610100516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa158015620001b8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001de919062000345565b6001600160a01b03166101405250609d805460ff19166001179055620002036200020c565b5050506200036c565b6006546301000000900460ff16156200027b5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60065460ff6201000090910481161015620002d6576006805462ff0000191662ff000017905560405160ff81527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114620002ee57600080fd5b50565b6000806000606084860312156200030757600080fd5b83516200031481620002d8565b60208501519093506200032781620002d8565b60408501519092506200033a81620002d8565b809150509250925092565b6000602082840312156200035857600080fd5b81516200036581620002d8565b9392505050565b60805160a05160c05160e051610100516101205161014051614eec62000461600039600081816105550152611cba0152600081816103f70152611e9c01526000818161041e01528181612072015261223401526000818161046b015281816111b50152818161198501528181611b1d0152611d57015260008181610442015281816125c50152818161270e01526127c0015260008181610d8f01528181610eea01528181610f81015281816129b101528181612b340152612bd3015260008181610bba01528181610c4901528181610cc901528181612528015281816126a6015281816128ef0152612a8f0152614eec6000f3fe608060405234801561001057600080fd5b506004361061023c5760003560e01c80636b3aa72e1161013b578063b98d0908116100b8578063e481af9d1161007c578063e481af9d14610577578063ef0244581461057f578063f2fde38b14610587578063f8e86ece1461059a578063fabc1cbc146105ad57600080fd5b8063b98d090814610520578063c6a2aac81461052d578063cf8e629a14610535578063d8ff953d1461053d578063df5cf7231461055057600080fd5b80638da5cb5b116100ff5780638da5cb5b146104c957806394c8e4ff146104da5780639926ee7d146104e7578063a364f4da146104fa578063a98fb3551461050d57600080fd5b80636b3aa72e146104405780636d14a987146104665780636efb46361461048d578063715018a6146104ae578063886f1195146104b657600080fd5b8063429d5bf0116101c95780635ac86ab71161018d5780635ac86ab7146103b45780635c975abb146103d75780635da93d7e146103df5780635df45946146103f2578063683048351461041957600080fd5b8063429d5bf01461033f5780634deabc21146103525780634f2b85db1461037657806354f4a9171461038c578063595c6a67146103ac57600080fd5b80631d1a696d116102105780631d1a696d146102ab57806333cfb7b7146102ce578063358394d8146102ee57806339bc68e714610301578063416c7e5e1461032c57600080fd5b80628f38e71461024157806310d67a2f14610256578063136439dd14610269578063171f1d5b1461027c575b600080fd5b61025461024f3660046143b9565b6105c0565b005b610254610264366004614439565b610807565b610254610277366004614456565b6108ba565b61028f61028a36600461446f565b6109f9565b6040805192151583529015156020830152015b60405180910390f35b6102be6102b9366004614456565b610b83565b60405190151581526020016102a2565b6102e16102dc366004614439565b610b95565b6040516102a291906144c0565b6102546102fc36600461450d565b611064565b600454610314906001600160a01b031681565b6040516001600160a01b0390911681526020016102a2565b61025461033a36600461456e565b6111b3565b61025461034d36600461459a565b611329565b60065461036490610100900460ff1681565b60405160ff90911681526020016102a2565b61037e61137b565b6040519081526020016102a2565b61039f61039a3660046145b7565b61138c565b6040516102a291906145d9565b610254611479565b6102be6103c236600461459a565b60d054600160ff9092169190911b9081161490565b60d05461037e565b6102546103ed366004614439565b611540565b6103147f000000000000000000000000000000000000000000000000000000000000000081565b6103147f000000000000000000000000000000000000000000000000000000000000000081565b7f0000000000000000000000000000000000000000000000000000000000000000610314565b6103147f000000000000000000000000000000000000000000000000000000000000000081565b6104a061049b366004614611565b6115d2565b6040516102a2929190614704565b6102546124e9565b60cf54610314906001600160a01b031681565b6039546001600160a01b0316610314565b6006546102be9060ff1681565b6102546104f53660046147c4565b6124fd565b610254610508366004614439565b61267b565b61025461051b36600461486f565b6127a1565b609d546102be9060ff1681565b610254612823565b610254612863565b61025461054b366004614456565b6128a0565b6103147f000000000000000000000000000000000000000000000000000000000000000081565b6102e16128e9565b61037e606481565b610254610595366004614439565b612cb2565b6102546105a8366004614439565b612d28565b6102546105bb366004614456565b612de5565b60d054156105e95760405162461bcd60e51b81526004016105e0906148bf565b60405180910390fd5b6004546001600160a01b0316336001600160a01b03161461061d5760405163fc4a01bd60e01b815260040160405180910390fd5b32331461063d57604051636edaef2f60e11b815260040160405180910390fd5b4361064e60808401606085016148f6565b63ffffffff1611156106735760405163c15ef5b560e01b815260040160405180910390fd5b600061068661068184614911565b612f41565b90506000806106b28361069c60208801886149b1565b6106ac60808a0160608b016148f6565b886115d2565b9150915060005b6106c660408701876149b1565b90508110156107bc5760006106de60408801886149b1565b838181106106ee576106ee6149fe565b60065492013560f81c925050610100900460ff168110156107225760405163bbf727c160e01b815260040160405180910390fd5b8060ff168460200151838151811061073c5761073c6149fe565b602002602001015161074e9190614a2a565b6001600160601b031660648560000151848151811061076f5761076f6149fe565b60200260200101516001600160601b031661078a9190614a59565b10156107a957604051633916714960e21b815260040160405180910390fd5b50806107b481614a78565b9150506106b9565b506107c960008635612fbb565b506040518535815283907ffdda6f7d4825a4f1e4e97b50a26a69a8bcc3f4fcb1113cc14ce8e7098ca636659060200160405180910390a25050505050565b60cf60009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561085a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061087e9190614a93565b6001600160a01b0316336001600160a01b0316146108ae5760405162461bcd60e51b81526004016105e090614ab0565b6108b781612fce565b50565b60cf5460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610902573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109269190614afa565b6109425760405162461bcd60e51b81526004016105e090614b17565b60d054818116146109bb5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c697479000000000000000060648201526084016105e0565b60d081905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000187876000015188602001518860000151600060028110610a4157610a416149fe565b60200201518951600160200201518a60200151600060028110610a6657610a666149fe565b60200201518b60200151600160028110610a8257610a826149fe565b602090810291909101518c518d830151604051610adf9a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c610b029190614b5f565b9050610b75610b1b610b1488846130c5565b869061315c565b610b236131f0565b610b6b610b5c85610b56604080518082018252600080825260209182015281518083019092526001825260029082015290565b906130c5565b610b658c6132b0565b9061315c565b886201d4c0613340565b909890975095505050505050565b6000610b8f8183613564565b92915050565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa158015610c01573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c259190614b81565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa158015610c90573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cb49190614b9a565b90506001600160c01b0381161580610d4e57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d25573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d499190614bc3565b60ff16155b15610d6a57505060408051600081526020810190915292915050565b6000610d7e826001600160c01b031661357c565b90506000805b8251811015610e54577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f5848381518110610dce57610dce6149fe565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa158015610e12573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e369190614b81565b610e409083614be0565b915080610e4c81614a78565b915050610d84565b506000816001600160401b03811115610e6f57610e6f613fb2565b604051908082528060200260200182016040528015610e98578160200160208202803683370190505b5090506000805b8451811015611057576000858281518110610ebc57610ebc6149fe565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa158015610f31573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f559190614b81565b905060005b81811015611041576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610fcf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ff39190614c0f565b60000151868681518110611009576110096149fe565b6001600160a01b03909216602092830291909101909101528461102b81614a78565b955050808061103990614a78565b915050610f5a565b505050808061104f90614a78565b915050610e9f565b5090979650505050505050565b6006546301000000900460ff161580801561108b575060065460016201000090910460ff16105b806110ab5750303b1580156110ab575060065462010000900460ff166001145b61110e5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016105e0565b6006805462ff00001916620100001790558015611139576006805463ff000000191663010000001790555b611143858561363e565b61114c83613728565b6111558261377a565b6006805461ff00191661420017905580156111ac576006805463ff00000019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611211573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112359190614a93565b6001600160a01b0316336001600160a01b0316146112e15760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a4016105e0565b609d805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc906020015b60405180910390a150565b6113316137db565b6006805461ff00191661010060ff8416908102919091179091556040519081527fc3acdc4f4bc283baa27c4207eb2c32954fbb26960847c9e15c2f7c89701342449060200161131e565b60006113876000613835565b905090565b6060600061139a6000613835565b90508084106113bc576040516392c4425960e01b815260040160405180910390fd5b60006113c88486614be0565b9050818111156113d55750805b60006113e18683614c4e565b6001600160401b038111156113f8576113f8613fb2565b604051908082528060200260200182016040528015611421578160200160208202803683370190505b509050855b8281101561146f5761143960008261383f565b826114448984614c4e565b81518110611454576114546149fe565b602090810291909101015261146881614a78565b9050611426565b5095945050505050565b60cf5460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156114c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114e59190614afa565b6115015760405162461bcd60e51b81526004016105e090614b17565b60001960d081905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b6115486137db565b6001600160a01b03811660009081526005602052604090205460ff16611581576040516315ebf2b560e21b815260040160405180910390fd5b6001600160a01b038116600081815260056020908152604091829020805460ff1916905590519182527f8560daa191dd8e6fba276b053006b3990c46c94b842f85490f52c49b15cfe5cb910161131e565b60408051808201909152606080825260208201526000846116495760405162461bcd60e51b81526020600482015260376024820152600080516020614e9783398151915260448201527f7265733a20656d7074792071756f72756d20696e70757400000000000000000060648201526084016105e0565b60408301515185148015611661575060a08301515185145b8015611671575060c08301515185145b8015611681575060e08301515185145b6116eb5760405162461bcd60e51b81526020600482015260416024820152600080516020614e9783398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a4016105e0565b825151602084015151146117635760405162461bcd60e51b815260206004820152604460248201819052600080516020614e97833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a4016105e0565b4363ffffffff168463ffffffff16106117d25760405162461bcd60e51b815260206004820152603c6024820152600080516020614e9783398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b0000000060648201526084016105e0565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b0381111561181357611813613fb2565b60405190808252806020026020018201604052801561183c578160200160208202803683370190505b506020820152866001600160401b0381111561185a5761185a613fb2565b604051908082528060200260200182016040528015611883578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b038111156118b7576118b7613fb2565b6040519080825280602002602001820160405280156118e0578160200160208202803683370190505b5081526020860151516001600160401b0381111561190057611900613fb2565b604051908082528060200260200182016040528015611929578160200160208202803683370190505b50816020018190525060006119fb8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa1580156119d2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119f69190614bc3565b61384b565b905060005b876020015151811015611c9657611a4588602001518281518110611a2657611a266149fe565b6020026020010151805160009081526020918201519091526040902090565b83602001518281518110611a5b57611a5b6149fe565b60209081029190910101528015611b1b576020830151611a7c600183614c4e565b81518110611a8c57611a8c6149fe565b602002602001015160001c83602001518281518110611aad57611aad6149fe565b602002602001015160001c11611b1b576040805162461bcd60e51b8152602060048201526024810191909152600080516020614e9783398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f7274656460648201526084016105e0565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec635184602001518381518110611b6057611b606149fe565b60200260200101518b8b600001518581518110611b7f57611b7f6149fe565b60200260200101516040518463ffffffff1660e01b8152600401611bbc9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611bd9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bfd9190614b9a565b6001600160c01b031683600001518281518110611c1c57611c1c6149fe565b602002602001018181525050611c82610b14611c568486600001518581518110611c4857611c486149fe565b6020026020010151166138d5565b8a602001518481518110611c6c57611c6c6149fe565b602002602001015161390090919063ffffffff16565b945080611c8e81614a78565b915050611a00565b5050611ca1836139e4565b609d5490935060ff16600081611cb8576000611d3a565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015611d16573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d3a9190614b81565b905060005b8a8110156123b8578215611e9a578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f86818110611d9657611d966149fe565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015611dd6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dfa9190614b81565b611e049190614be0565b11611e9a5760405162461bcd60e51b81526020600482015260666024820152600080516020614e9783398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c4016105e0565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d84818110611edb57611edb6149fe565b9050013560f81c60f81b60f81c8c8c60a001518581518110611eff57611eff6149fe565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015611f5b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f7f9190614c65565b6001600160401b031916611fa28a604001518381518110611a2657611a266149fe565b67ffffffffffffffff19161461203e5760405162461bcd60e51b81526020600482015260616024820152600080516020614e9783398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c4016105e0565b61206e89604001518281518110612057576120576149fe565b60200260200101518761315c90919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d848181106120b1576120b16149fe565b9050013560f81c60f81b60f81c8c8c60c0015185815181106120d5576120d56149fe565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612131573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121559190614c90565b8560200151828151811061216b5761216b6149fe565b6001600160601b03909216602092830291909101820152850151805182908110612197576121976149fe565b6020026020010151856000015182815181106121b5576121b56149fe565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a60200151518110156123a35761222d866000015182815181106121ff576121ff6149fe565b60200260200101518f8f86818110612219576122196149fe565b600192013560f81c9290921c811614919050565b15612391577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f86818110612273576122736149fe565b9050013560f81c60f81b60f81c8e89602001518581518110612297576122976149fe565b60200260200101518f60e0015188815181106122b5576122b56149fe565b602002602001015187815181106122ce576122ce6149fe565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015612332573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123569190614c90565b875180518590811061236a5761236a6149fe565b6020026020010181815161237e9190614cab565b6001600160601b03169052506001909101905b8061239b81614a78565b9150506121d9565b505080806123b090614a78565b915050611d3f565b5050506000806123d28c868a606001518b608001516109f9565b91509150816124435760405162461bcd60e51b81526020600482015260436024820152600080516020614e9783398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a4016105e0565b806124a45760405162461bcd60e51b81526020600482015260396024820152600080516020614e9783398151915260448201527f7265733a207369676e617475726520697320696e76616c69640000000000000060648201526084016105e0565b505060008782602001516040516020016124bf929190614cd3565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b6124f16137db565b6124fb6000613728565b565b60d0541561251d5760405162461bcd60e51b81526004016105e0906148bf565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146125655760405162461bcd60e51b81526004016105e090614d1b565b60065460ff16801561259057506001600160a01b03821660009081526005602052604090205460ff16155b156125ae576040516315ebf2b560e21b815260040160405180910390fd5b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d906125fc9085908590600401614de0565b600060405180830381600087803b15801561261657600080fd5b505af115801561262a573d6000803e3d6000fd5b50505050612642826002613a7f90919063ffffffff16565b506040516001600160a01b038316907fac6fa858e9350a46cec16539926e0fde25b7629f84b5a72bffaae4df888ae86d90600090a25050565b60d0541561269b5760405162461bcd60e51b81526004016105e0906148bf565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146126e35760405162461bcd60e51b81526004016105e090614d1b565b6126ee600282613a94565b506040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da90602401600060405180830381600087803b15801561275257600080fd5b505af1158015612766573d6000803e3d6000fd5b50506040516001600160a01b03841692507f80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d9150600090a250565b6127a96137db565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb355906127f5908490600401614e2b565b600060405180830381600087803b15801561280f57600080fd5b505af11580156111ac573d6000803e3d6000fd5b61282b6137db565b6006805460ff191660011790556040517f8a943acd5f4e6d3df7565a4a08a93f6b04cc31bb6c01ca4aef7abd6baf455ec390600090a1565b61286b6137db565b6006805460ff191690556040517f2d35c8d348a345fd7b3b03b7cfcf7ad0b60c2d46742d5ca536342e4185becb0790600090a1565b6128a86137db565b6128b3600082613aa9565b50604080518281523360208201527f1bdeffc0337373bf2f6fd4219080133eeaaee0554206d9bb24a019d96973c1eb910161131e565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561294b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061296f9190614bc3565b60ff1690508061298d57505060408051600081526020810190915290565b6000805b82811015612a4257604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015612a00573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a249190614b81565b612a2e9083614be0565b915080612a3a81614a78565b915050612991565b506000816001600160401b03811115612a5d57612a5d613fb2565b604051908082528060200260200182016040528015612a86578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612aeb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b0f9190614bc3565b60ff16811015612ca857604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015612b83573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ba79190614b81565b905060005b81811015612c93576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015612c21573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c459190614c0f565b60000151858581518110612c5b57612c5b6149fe565b6001600160a01b039092166020928302919091019091015283612c7d81614a78565b9450508080612c8b90614a78565b915050612bac565b50508080612ca090614a78565b915050612a8d565b5090949350505050565b612cba6137db565b6001600160a01b038116612d1f5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016105e0565b6108b781613728565b612d306137db565b6001600160a01b038116612d575760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03811660009081526005602052604090205460ff1615612d91576040516302eeeca960e41b815260040160405180910390fd5b6001600160a01b038116600081815260056020908152604091829020805460ff1916600117905590519182527fdde65206cdee4ea27ef1b170724ba50b41ad09a3bf2dda12935fc40c4dbf6e75910161131e565b60cf60009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612e38573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e5c9190614a93565b6001600160a01b0316336001600160a01b031614612e8c5760405162461bcd60e51b81526004016105e090614ab0565b60d05419811960d054191614612f0a5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c697479000000000000000060648201526084016105e0565b60d081905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016109ee565b6000612f7e82604080518082019091526000808252602082015250604080518082019091528151815260609091015163ffffffff16602082015290565b6040805182516020808301919091529092015163ffffffff1690820152606001604051602081830303815290604052805190602001209050919050565b6000612fc78383613ab5565b9392505050565b6001600160a01b03811661305c5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a4016105e0565b60cf54604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a160cf80546001600160a01b0319166001600160a01b0392909216919091179055565b60408051808201909152600080825260208201526130e1613ed8565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa905080801561311457613116565bfe5b50806131545760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b60448201526064016105e0565b505092915050565b6040805180820190915260008082526020820152613178613ef6565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156131145750806131545760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b60448201526064016105e0565b6131f8613f14565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b6040805180820190915260008082526020820152600080806132e0600080516020614e7783398151915286614b5f565b90505b6132ec81613b04565b9093509150600080516020614e77833981519152828309831415613326576040805180820190915290815260208101919091529392505050565b600080516020614e778339815191526001820890506132e3565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190613372613f39565b60005b600281101561353757600061338b826006614a59565b905084826002811061339f5761339f6149fe565b602002015151836133b1836000614be0565b600c81106133c1576133c16149fe565b60200201528482600281106133d8576133d86149fe565b602002015160200151838260016133ef9190614be0565b600c81106133ff576133ff6149fe565b6020020152838260028110613416576134166149fe565b6020020151515183613429836002614be0565b600c8110613439576134396149fe565b6020020152838260028110613450576134506149fe565b6020020151516001602002015183613469836003614be0565b600c8110613479576134796149fe565b6020020152838260028110613490576134906149fe565b6020020151602001516000600281106134ab576134ab6149fe565b6020020151836134bc836004614be0565b600c81106134cc576134cc6149fe565b60200201528382600281106134e3576134e36149fe565b6020020151602001516001600281106134fe576134fe6149fe565b60200201518361350f836005614be0565b600c811061351f5761351f6149fe565b6020020152508061352f81614a78565b915050613375565b50613540613f58565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b60008181526001830160205260408120541515612fc7565b606060008061358a846138d5565b61ffff166001600160401b038111156135a5576135a5613fb2565b6040519080825280601f01601f1916602001820160405280156135cf576020820181803683370190505b5090506000805b8251821080156135e7575061010081105b15612ca8576001811b93508584161561362e578060f81b838381518110613610576136106149fe565b60200101906001600160f81b031916908160001a9053508160010191505b61363781614a78565b90506135d6565b60cf546001600160a01b031615801561365f57506001600160a01b03821615155b6136e15760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a4016105e0565b60d081905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a261372482612fce565b5050565b603980546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600480546001600160a01b038381166001600160a01b031983168117909355604080519190921680825260208201939093527f175f27847b3568e0da876ffca3dc0bb52db4e6556346aedb530c6fe86610da27910160405180910390a15050565b6039546001600160a01b031633146124fb5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105e0565b6000610b8f825490565b6000612fc78383613b86565b60008061385784613bb0565b9050808360ff166001901b11612fc75760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c75650060648201526084016105e0565b6000805b8215610b8f576138ea600184614c4e565b90921691806138f881614e3e565b9150506138d9565b60408051808201909152600080825260208201526102008261ffff161061395c5760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b60448201526064016105e0565b8161ffff1660011415613970575081610b8f565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff16106139d957600161ffff871660ff83161c811614156139bc576139b9848461315c565b93505b6139c6838461315c565b92506201fffe600192831b16910161398c565b509195945050505050565b60408051808201909152600080825260208201528151158015613a0957506020820151155b15613a27575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020614e778339815191528460200151613a5a9190614b5f565b613a7290600080516020614e77833981519152614c4e565b905292915050565b919050565b6000612fc7836001600160a01b038416613ab5565b6000612fc7836001600160a01b038416613d3d565b6000612fc78383613d3d565b6000818152600183016020526040812054613afc57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610b8f565b506000610b8f565b60008080600080516020614e778339815191526003600080516020614e7783398151915286600080516020614e77833981519152888909090890506000613b7a827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020614e77833981519152613e30565b91959194509092505050565b6000826000018281548110613b9d57613b9d6149fe565b9060005260206000200154905092915050565b600061010082511115613c395760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a4016105e0565b8151613c4757506000919050565b60008083600081518110613c5d57613c5d6149fe565b0160200151600160f89190911c81901b92505b8451811015613d3457848181518110613c8b57613c8b6149fe565b0160200151600160f89190911c1b9150828211613d205760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a4016105e0565b91811791613d2d81614a78565b9050613c70565b50909392505050565b60008181526001830160205260408120548015613e26576000613d61600183614c4e565b8554909150600090613d7590600190614c4e565b9050818114613dda576000866000018281548110613d9557613d956149fe565b9060005260206000200154905080876000018481548110613db857613db86149fe565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080613deb57613deb614e60565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610b8f565b6000915050610b8f565b600080613e3b613f58565b613e43613f76565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa9250828015613114575082613ecd5760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c75726500000000000060448201526064016105e0565b505195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518060400160405280613f27613f94565b8152602001613f34613f94565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715613fea57613fea613fb2565b60405290565b60405161010081016001600160401b0381118282101715613fea57613fea613fb2565b604051601f8201601f191681016001600160401b038111828210171561403b5761403b613fb2565b604052919050565b60006001600160401b0382111561405c5761405c613fb2565b5060051b60200190565b803563ffffffff81168114613a7a57600080fd5b600082601f83011261408b57600080fd5b813560206140a061409b83614043565b614013565b82815260059290921b840181019181810190868411156140bf57600080fd5b8286015b848110156140e1576140d481614066565b83529183019183016140c3565b509695505050505050565b6000604082840312156140fe57600080fd5b614106613fc8565b9050813581526020820135602082015292915050565b600082601f83011261412d57600080fd5b8135602061413d61409b83614043565b82815260069290921b8401810191818101908684111561415c57600080fd5b8286015b848110156140e15761417288826140ec565b835291830191604001614160565b600082601f83011261419157600080fd5b614199613fc8565b8060408401858111156141ab57600080fd5b845b818110156141c55780358452602093840193016141ad565b509095945050505050565b6000608082840312156141e257600080fd5b6141ea613fc8565b90506141f68383614180565b81526142058360408401614180565b602082015292915050565b600082601f83011261422157600080fd5b8135602061423161409b83614043565b82815260059290921b8401810191818101908684111561425057600080fd5b8286015b848110156140e15780356001600160401b038111156142735760008081fd5b6142818986838b010161407a565b845250918301918301614254565b600061018082840312156142a257600080fd5b6142aa613ff0565b905081356001600160401b03808211156142c357600080fd5b6142cf8583860161407a565b835260208401359150808211156142e557600080fd5b6142f18583860161411c565b6020840152604084013591508082111561430a57600080fd5b6143168583860161411c565b604084015261432885606086016141d0565b606084015261433a8560e086016140ec565b608084015261012084013591508082111561435457600080fd5b6143608583860161407a565b60a084015261014084013591508082111561437a57600080fd5b6143868583860161407a565b60c08401526101608401359150808211156143a057600080fd5b506143ad84828501614210565b60e08301525092915050565b600080604083850312156143cc57600080fd5b82356001600160401b03808211156143e357600080fd5b90840190608082870312156143f757600080fd5b9092506020840135908082111561440d57600080fd5b5061441a8582860161428f565b9150509250929050565b6001600160a01b03811681146108b757600080fd5b60006020828403121561444b57600080fd5b8135612fc781614424565b60006020828403121561446857600080fd5b5035919050565b600080600080610120858703121561448657600080fd5b8435935061449786602087016140ec565b92506144a686606087016141d0565b91506144b58660e087016140ec565b905092959194509250565b6020808252825182820181905260009190848201906040850190845b818110156145015783516001600160a01b0316835292840192918401916001016144dc565b50909695505050505050565b6000806000806080858703121561452357600080fd5b843561452e81614424565b935060208501359250604085013561454581614424565b9150606085013561455581614424565b939692955090935050565b80151581146108b757600080fd5b60006020828403121561458057600080fd5b8135612fc781614560565b60ff811681146108b757600080fd5b6000602082840312156145ac57600080fd5b8135612fc78161458b565b600080604083850312156145ca57600080fd5b50508035926020909101359150565b6020808252825182820181905260009190848201906040850190845b81811015614501578351835292840192918401916001016145f5565b60008060008060006080868803121561462957600080fd5b8535945060208601356001600160401b038082111561464757600080fd5b818801915088601f83011261465b57600080fd5b81358181111561466a57600080fd5b89602082850101111561467c57600080fd5b602083019650945061469060408901614066565b935060608801359150808211156146a657600080fd5b506146b38882890161428f565b9150509295509295909350565b600081518084526020808501945080840160005b838110156146f95781516001600160601b0316875295820195908201906001016146d4565b509495945050505050565b604081526000835160408084015261471f60808401826146c0565b90506020850151603f1984830301606085015261473c82826146c0565b925050508260208301529392505050565b60006001600160401b0383111561476657614766613fb2565b614779601f8401601f1916602001614013565b905082815283838301111561478d57600080fd5b828260208301376000602084830101529392505050565b600082601f8301126147b557600080fd5b612fc78383356020850161474d565b600080604083850312156147d757600080fd5b82356147e281614424565b915060208301356001600160401b03808211156147fe57600080fd5b908401906060828703121561481257600080fd5b60405160608101818110838211171561482d5761482d613fb2565b60405282358281111561483f57600080fd5b61484b888286016147a4565b82525060208301356020820152604083013560408201528093505050509250929050565b60006020828403121561488157600080fd5b81356001600160401b0381111561489757600080fd5b8201601f810184136148a857600080fd5b6148b78482356020840161474d565b949350505050565b6020808252601c908201527f5061757361626c653a20636f6e74726163742069732070617573656400000000604082015260600190565b60006020828403121561490857600080fd5b612fc782614066565b60006080823603121561492357600080fd5b604051608081016001600160401b03828210818311171561494657614946613fb2565b8160405284358352602085013591508082111561496257600080fd5b61496e368387016147a4565b6020840152604085013591508082111561498757600080fd5b50614994368286016147a4565b6040830152506149a660608401614066565b606082015292915050565b6000808335601e198436030181126149c857600080fd5b8301803591506001600160401b038211156149e257600080fd5b6020019150368190038213156149f757600080fd5b9250929050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006001600160601b0380831681851681830481118215151615614a5057614a50614a14565b02949350505050565b6000816000190483118215151615614a7357614a73614a14565b500290565b6000600019821415614a8c57614a8c614a14565b5060010190565b600060208284031215614aa557600080fd5b8151612fc781614424565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600060208284031215614b0c57600080fd5b8151612fc781614560565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b600082614b7c57634e487b7160e01b600052601260045260246000fd5b500690565b600060208284031215614b9357600080fd5b5051919050565b600060208284031215614bac57600080fd5b81516001600160c01b0381168114612fc757600080fd5b600060208284031215614bd557600080fd5b8151612fc78161458b565b60008219821115614bf357614bf3614a14565b500190565b80516001600160601b0381168114613a7a57600080fd5b600060408284031215614c2157600080fd5b614c29613fc8565b8251614c3481614424565b8152614c4260208401614bf8565b60208201529392505050565b600082821015614c6057614c60614a14565b500390565b600060208284031215614c7757600080fd5b815167ffffffffffffffff1981168114612fc757600080fd5b600060208284031215614ca257600080fd5b612fc782614bf8565b60006001600160601b0383811690831681811015614ccb57614ccb614a14565b039392505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b83811015614d0e57815185529382019390820190600101614cf2565b5092979650505050505050565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b6000815180845260005b81811015614db957602081850181015186830182015201614d9d565b81811115614dcb576000602083870101525b50601f01601f19169290920160200192915050565b60018060a01b0383168152604060208201526000825160606040840152614e0a60a0840182614d93565b90506020840151606084015260408401516080840152809150509392505050565b602081526000612fc76020830184614d93565b600061ffff80831681811415614e5657614e56614a14565b6001019392505050565b634e487b7160e01b600052603160045260246000fdfe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a264697066735822122086d4c584ff6ec3c54140c7f4d498c7da6bdaf74a928c5af8b10d778933c7239764736f6c634300080c0033",
}

// ContractMachServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMachServiceManagerMetaData.ABI instead.
var ContractMachServiceManagerABI = ContractMachServiceManagerMetaData.ABI

// ContractMachServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMachServiceManagerMetaData.Bin instead.
var ContractMachServiceManagerBin = ContractMachServiceManagerMetaData.Bin

// DeployContractMachServiceManager deploys a new Ethereum contract, binding an instance of ContractMachServiceManager to it.
func DeployContractMachServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, __avsDirectory common.Address, __registryCoordinator common.Address, __stakeRegistry common.Address) (common.Address, *types.Transaction, *ContractMachServiceManager, error) {
	parsed, err := ContractMachServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractMachServiceManagerBin), backend, __avsDirectory, __registryCoordinator, __stakeRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractMachServiceManager{ContractMachServiceManagerCaller: ContractMachServiceManagerCaller{contract: contract}, ContractMachServiceManagerTransactor: ContractMachServiceManagerTransactor{contract: contract}, ContractMachServiceManagerFilterer: ContractMachServiceManagerFilterer{contract: contract}}, nil
}

// ContractMachServiceManager is an auto generated Go binding around an Ethereum contract.
type ContractMachServiceManager struct {
	ContractMachServiceManagerCaller     // Read-only binding to the contract
	ContractMachServiceManagerTransactor // Write-only binding to the contract
	ContractMachServiceManagerFilterer   // Log filterer for contract events
}

// ContractMachServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractMachServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractMachServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractMachServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractMachServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractMachServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractMachServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractMachServiceManagerSession struct {
	Contract     *ContractMachServiceManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ContractMachServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractMachServiceManagerCallerSession struct {
	Contract *ContractMachServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ContractMachServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractMachServiceManagerTransactorSession struct {
	Contract     *ContractMachServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ContractMachServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractMachServiceManagerRaw struct {
	Contract *ContractMachServiceManager // Generic contract binding to access the raw methods on
}

// ContractMachServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractMachServiceManagerCallerRaw struct {
	Contract *ContractMachServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractMachServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractMachServiceManagerTransactorRaw struct {
	Contract *ContractMachServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractMachServiceManager creates a new instance of ContractMachServiceManager, bound to a specific deployed contract.
func NewContractMachServiceManager(address common.Address, backend bind.ContractBackend) (*ContractMachServiceManager, error) {
	contract, err := bindContractMachServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractMachServiceManager{ContractMachServiceManagerCaller: ContractMachServiceManagerCaller{contract: contract}, ContractMachServiceManagerTransactor: ContractMachServiceManagerTransactor{contract: contract}, ContractMachServiceManagerFilterer: ContractMachServiceManagerFilterer{contract: contract}}, nil
}

// NewContractMachServiceManagerCaller creates a new read-only instance of ContractMachServiceManager, bound to a specific deployed contract.
func NewContractMachServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractMachServiceManagerCaller, error) {
	contract, err := bindContractMachServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractMachServiceManagerCaller{contract: contract}, nil
}

// NewContractMachServiceManagerTransactor creates a new write-only instance of ContractMachServiceManager, bound to a specific deployed contract.
func NewContractMachServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractMachServiceManagerTransactor, error) {
	contract, err := bindContractMachServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractMachServiceManagerTransactor{contract: contract}, nil
}

// NewContractMachServiceManagerFilterer creates a new log filterer instance of ContractMachServiceManager, bound to a specific deployed contract.
func NewContractMachServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractMachServiceManagerFilterer, error) {
	contract, err := bindContractMachServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractMachServiceManagerFilterer{contract: contract}, nil
}

// bindContractMachServiceManager binds a generic wrapper to an already deployed contract.
func bindContractMachServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMachServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractMachServiceManager *ContractMachServiceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractMachServiceManager.Contract.ContractMachServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractMachServiceManager *ContractMachServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.ContractMachServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractMachServiceManager *ContractMachServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.ContractMachServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractMachServiceManager *ContractMachServiceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractMachServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractMachServiceManager *ContractMachServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractMachServiceManager *ContractMachServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.contract.Transact(opts, method, params...)
}

// THRESHOLDDENOMINATOR is a free data retrieval call binding the contract method 0xef024458.
//
// Solidity: function THRESHOLD_DENOMINATOR() view returns(uint256)
func (_ContractMachServiceManager *ContractMachServiceManagerCaller) THRESHOLDDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractMachServiceManager.contract.Call(opts, &out, "THRESHOLD_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// THRESHOLDDENOMINATOR is a free data retrieval call binding the contract method 0xef024458.
//
// Solidity: function THRESHOLD_DENOMINATOR() view returns(uint256)
func (_ContractMachServiceManager *ContractMachServiceManagerSession) THRESHOLDDENOMINATOR() (*big.Int, error) {
	return _ContractMachServiceManager.Contract.THRESHOLDDENOMINATOR(&_ContractMachServiceManager.CallOpts)
}

// THRESHOLDDENOMINATOR is a free data retrieval call binding the contract method 0xef024458.
//
// Solidity: function THRESHOLD_DENOMINATOR() view returns(uint256)
func (_ContractMachServiceManager *ContractMachServiceManagerCallerSession) THRESHOLDDENOMINATOR() (*big.Int, error) {
	return _ContractMachServiceManager.Contract.THRESHOLDDENOMINATOR(&_ContractMachServiceManager.CallOpts)
}

// AlertConfirmer is a free data retrieval call binding the contract method 0x39bc68e7.
//
// Solidity: function alertConfirmer() view returns(address)
func (_ContractMachServiceManager *ContractMachServiceManagerCaller) AlertConfirmer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractMachServiceManager.contract.Call(opts, &out, "alertConfirmer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AlertConfirmer is a free data retrieval call binding the contract method 0x39bc68e7.
//
// Solidity: function alertConfirmer() view returns(address)
func (_ContractMachServiceManager *ContractMachServiceManagerSession) AlertConfirmer() (common.Address, error) {
	return _ContractMachServiceManager.Contract.AlertConfirmer(&_ContractMachServiceManager.CallOpts)
}

// AlertConfirmer is a free data retrieval call binding the contract method 0x39bc68e7.
//
// Solidity: function alertConfirmer() view returns(address)
func (_ContractMachServiceManager *ContractMachServiceManagerCallerSession) AlertConfirmer() (common.Address, error) {
	return _ContractMachServiceManager.Contract.AlertConfirmer(&_ContractMachServiceManager.CallOpts)
}

// AllowlistEnabled is a free data retrieval call binding the contract method 0x94c8e4ff.
//
// Solidity: function allowlistEnabled() view returns(bool)
func (_ContractMachServiceManager *ContractMachServiceManagerCaller) AllowlistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractMachServiceManager.contract.Call(opts, &out, "allowlistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowlistEnabled is a free data retrieval call binding the contract method 0x94c8e4ff.
//
// Solidity: function allowlistEnabled() view returns(bool)
func (_ContractMachServiceManager *ContractMachServiceManagerSession) AllowlistEnabled() (bool, error) {
	return _ContractMachServiceManager.Contract.AllowlistEnabled(&_ContractMachServiceManager.CallOpts)
}

// AllowlistEnabled is a free data retrieval call binding the contract method 0x94c8e4ff.
//
// Solidity: function allowlistEnabled() view returns(bool)
func (_ContractMachServiceManager *ContractMachServiceManagerCallerSession) AllowlistEnabled() (bool, error) {
	return _ContractMachServiceManager.Contract.AllowlistEnabled(&_ContractMachServiceManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractMachServiceManager *ContractMachServiceManagerCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractMachServiceManager.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractMachServiceManager *ContractMachServiceManagerSession) AvsDirectory() (common.Address, error) {
	return _ContractMachServiceManager.Contract.AvsDirectory(&_ContractMachServiceManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractMachServiceManager *ContractMachServiceManagerCallerSession) AvsDirectory() (common.Address, error) {
	return _ContractMachServiceManager.Contract.AvsDirectory(&_ContractMachServiceManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractMachServiceManager *ContractMachServiceManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractMachServiceManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractMachServiceManager *ContractMachServiceManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractMachServiceManager.Contract.BlsApkRegistry(&_ContractMachServiceManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractMachServiceManager *ContractMachServiceManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractMachServiceManager.Contract.BlsApkRegistry(&_ContractMachServiceManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractMachServiceManager *ContractMachServiceManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractMachServiceManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

	if err != nil {
		return *new(IBLSSignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerQuorumStakeTotals)).(*IBLSSignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractMachServiceManager *ContractMachServiceManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractMachServiceManager.Contract.CheckSignatures(&_ContractMachServiceManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractMachServiceManager *ContractMachServiceManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractMachServiceManager.Contract.CheckSignatures(&_ContractMachServiceManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Contains is a free data retrieval call binding the contract method 0x1d1a696d.
//
// Solidity: function contains(bytes32 messageHash) view returns(bool)
func (_ContractMachServiceManager *ContractMachServiceManagerCaller) Contains(opts *bind.CallOpts, messageHash [32]byte) (bool, error) {
	var out []interface{}
	err := _ContractMachServiceManager.contract.Call(opts, &out, "contains", messageHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Contains is a free data retrieval call binding the contract method 0x1d1a696d.
//
// Solidity: function contains(bytes32 messageHash) view returns(bool)
func (_ContractMachServiceManager *ContractMachServiceManagerSession) Contains(messageHash [32]byte) (bool, error) {
	return _ContractMachServiceManager.Contract.Contains(&_ContractMachServiceManager.CallOpts, messageHash)
}

// Contains is a free data retrieval call binding the contract method 0x1d1a696d.
//
// Solidity: function contains(bytes32 messageHash) view returns(bool)
func (_ContractMachServiceManager *ContractMachServiceManagerCallerSession) Contains(messageHash [32]byte) (bool, error) {
	return _ContractMachServiceManager.Contract.Contains(&_ContractMachServiceManager.CallOpts, messageHash)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractMachServiceManager *ContractMachServiceManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractMachServiceManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractMachServiceManager *ContractMachServiceManagerSession) Delegation() (common.Address, error) {
	return _ContractMachServiceManager.Contract.Delegation(&_ContractMachServiceManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractMachServiceManager *ContractMachServiceManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractMachServiceManager.Contract.Delegation(&_ContractMachServiceManager.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractMachServiceManager *ContractMachServiceManagerCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractMachServiceManager.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractMachServiceManager *ContractMachServiceManagerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractMachServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractMachServiceManager.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractMachServiceManager *ContractMachServiceManagerCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractMachServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractMachServiceManager.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractMachServiceManager *ContractMachServiceManagerCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractMachServiceManager.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractMachServiceManager *ContractMachServiceManagerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractMachServiceManager.Contract.GetRestakeableStrategies(&_ContractMachServiceManager.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractMachServiceManager *ContractMachServiceManagerCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractMachServiceManager.Contract.GetRestakeableStrategies(&_ContractMachServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractMachServiceManager *ContractMachServiceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractMachServiceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractMachServiceManager *ContractMachServiceManagerSession) Owner() (common.Address, error) {
	return _ContractMachServiceManager.Contract.Owner(&_ContractMachServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractMachServiceManager *ContractMachServiceManagerCallerSession) Owner() (common.Address, error) {
	return _ContractMachServiceManager.Contract.Owner(&_ContractMachServiceManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractMachServiceManager *ContractMachServiceManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractMachServiceManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractMachServiceManager *ContractMachServiceManagerSession) Paused(index uint8) (bool, error) {
	return _ContractMachServiceManager.Contract.Paused(&_ContractMachServiceManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractMachServiceManager *ContractMachServiceManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractMachServiceManager.Contract.Paused(&_ContractMachServiceManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractMachServiceManager *ContractMachServiceManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractMachServiceManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractMachServiceManager *ContractMachServiceManagerSession) Paused0() (*big.Int, error) {
	return _ContractMachServiceManager.Contract.Paused0(&_ContractMachServiceManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractMachServiceManager *ContractMachServiceManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractMachServiceManager.Contract.Paused0(&_ContractMachServiceManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractMachServiceManager *ContractMachServiceManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractMachServiceManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractMachServiceManager *ContractMachServiceManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractMachServiceManager.Contract.PauserRegistry(&_ContractMachServiceManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractMachServiceManager *ContractMachServiceManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractMachServiceManager.Contract.PauserRegistry(&_ContractMachServiceManager.CallOpts)
}

// QueryMessageHashes is a free data retrieval call binding the contract method 0x54f4a917.
//
// Solidity: function queryMessageHashes(uint256 start, uint256 querySize) view returns(bytes32[])
func (_ContractMachServiceManager *ContractMachServiceManagerCaller) QueryMessageHashes(opts *bind.CallOpts, start *big.Int, querySize *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _ContractMachServiceManager.contract.Call(opts, &out, "queryMessageHashes", start, querySize)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// QueryMessageHashes is a free data retrieval call binding the contract method 0x54f4a917.
//
// Solidity: function queryMessageHashes(uint256 start, uint256 querySize) view returns(bytes32[])
func (_ContractMachServiceManager *ContractMachServiceManagerSession) QueryMessageHashes(start *big.Int, querySize *big.Int) ([][32]byte, error) {
	return _ContractMachServiceManager.Contract.QueryMessageHashes(&_ContractMachServiceManager.CallOpts, start, querySize)
}

// QueryMessageHashes is a free data retrieval call binding the contract method 0x54f4a917.
//
// Solidity: function queryMessageHashes(uint256 start, uint256 querySize) view returns(bytes32[])
func (_ContractMachServiceManager *ContractMachServiceManagerCallerSession) QueryMessageHashes(start *big.Int, querySize *big.Int) ([][32]byte, error) {
	return _ContractMachServiceManager.Contract.QueryMessageHashes(&_ContractMachServiceManager.CallOpts, start, querySize)
}

// QuorumThresholdPercentage is a free data retrieval call binding the contract method 0x4deabc21.
//
// Solidity: function quorumThresholdPercentage() view returns(uint8)
func (_ContractMachServiceManager *ContractMachServiceManagerCaller) QuorumThresholdPercentage(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ContractMachServiceManager.contract.Call(opts, &out, "quorumThresholdPercentage")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// QuorumThresholdPercentage is a free data retrieval call binding the contract method 0x4deabc21.
//
// Solidity: function quorumThresholdPercentage() view returns(uint8)
func (_ContractMachServiceManager *ContractMachServiceManagerSession) QuorumThresholdPercentage() (uint8, error) {
	return _ContractMachServiceManager.Contract.QuorumThresholdPercentage(&_ContractMachServiceManager.CallOpts)
}

// QuorumThresholdPercentage is a free data retrieval call binding the contract method 0x4deabc21.
//
// Solidity: function quorumThresholdPercentage() view returns(uint8)
func (_ContractMachServiceManager *ContractMachServiceManagerCallerSession) QuorumThresholdPercentage() (uint8, error) {
	return _ContractMachServiceManager.Contract.QuorumThresholdPercentage(&_ContractMachServiceManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractMachServiceManager *ContractMachServiceManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractMachServiceManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractMachServiceManager *ContractMachServiceManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractMachServiceManager.Contract.RegistryCoordinator(&_ContractMachServiceManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractMachServiceManager *ContractMachServiceManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractMachServiceManager.Contract.RegistryCoordinator(&_ContractMachServiceManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractMachServiceManager *ContractMachServiceManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractMachServiceManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractMachServiceManager *ContractMachServiceManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractMachServiceManager.Contract.StakeRegistry(&_ContractMachServiceManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractMachServiceManager *ContractMachServiceManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractMachServiceManager.Contract.StakeRegistry(&_ContractMachServiceManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractMachServiceManager *ContractMachServiceManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractMachServiceManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractMachServiceManager *ContractMachServiceManagerSession) StaleStakesForbidden() (bool, error) {
	return _ContractMachServiceManager.Contract.StaleStakesForbidden(&_ContractMachServiceManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractMachServiceManager *ContractMachServiceManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractMachServiceManager.Contract.StaleStakesForbidden(&_ContractMachServiceManager.CallOpts)
}

// TotalAlerts is a free data retrieval call binding the contract method 0x4f2b85db.
//
// Solidity: function totalAlerts() view returns(uint256)
func (_ContractMachServiceManager *ContractMachServiceManagerCaller) TotalAlerts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractMachServiceManager.contract.Call(opts, &out, "totalAlerts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAlerts is a free data retrieval call binding the contract method 0x4f2b85db.
//
// Solidity: function totalAlerts() view returns(uint256)
func (_ContractMachServiceManager *ContractMachServiceManagerSession) TotalAlerts() (*big.Int, error) {
	return _ContractMachServiceManager.Contract.TotalAlerts(&_ContractMachServiceManager.CallOpts)
}

// TotalAlerts is a free data retrieval call binding the contract method 0x4f2b85db.
//
// Solidity: function totalAlerts() view returns(uint256)
func (_ContractMachServiceManager *ContractMachServiceManagerCallerSession) TotalAlerts() (*big.Int, error) {
	return _ContractMachServiceManager.Contract.TotalAlerts(&_ContractMachServiceManager.CallOpts)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractMachServiceManager *ContractMachServiceManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractMachServiceManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractMachServiceManager *ContractMachServiceManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractMachServiceManager.Contract.TrySignatureAndApkVerification(&_ContractMachServiceManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractMachServiceManager *ContractMachServiceManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractMachServiceManager.Contract.TrySignatureAndApkVerification(&_ContractMachServiceManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// AddToAllowlist is a paid mutator transaction binding the contract method 0xf8e86ece.
//
// Solidity: function addToAllowlist(address operator) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactor) AddToAllowlist(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractMachServiceManager.contract.Transact(opts, "addToAllowlist", operator)
}

// AddToAllowlist is a paid mutator transaction binding the contract method 0xf8e86ece.
//
// Solidity: function addToAllowlist(address operator) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerSession) AddToAllowlist(operator common.Address) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.AddToAllowlist(&_ContractMachServiceManager.TransactOpts, operator)
}

// AddToAllowlist is a paid mutator transaction binding the contract method 0xf8e86ece.
//
// Solidity: function addToAllowlist(address operator) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactorSession) AddToAllowlist(operator common.Address) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.AddToAllowlist(&_ContractMachServiceManager.TransactOpts, operator)
}

// ConfirmAlert is a paid mutator transaction binding the contract method 0x008f38e7.
//
// Solidity: function confirmAlert((bytes32,bytes,bytes,uint32) alertHeader, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactor) ConfirmAlert(opts *bind.TransactOpts, alertHeader IMachServiceManagerAlertHeader, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractMachServiceManager.contract.Transact(opts, "confirmAlert", alertHeader, nonSignerStakesAndSignature)
}

func (_ContractMachServiceManager *ContractMachServiceManagerTransactor) VerifyRequest(opts *bind.TransactOpts, reqId [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractMachServiceManager.contract.Transact(opts, "verifyRequest", reqId, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature)
}

// ConfirmAlert is a paid mutator transaction binding the contract method 0x008f38e7.
//
// Solidity: function confirmAlert((bytes32,bytes,bytes,uint32) alertHeader, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerSession) ConfirmAlert(alertHeader IMachServiceManagerAlertHeader, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.ConfirmAlert(&_ContractMachServiceManager.TransactOpts, alertHeader, nonSignerStakesAndSignature)
}

// ConfirmAlert is a paid mutator transaction binding the contract method 0x008f38e7.
//
// Solidity: function confirmAlert((bytes32,bytes,bytes,uint32) alertHeader, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactorSession) ConfirmAlert(alertHeader IMachServiceManagerAlertHeader, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.ConfirmAlert(&_ContractMachServiceManager.TransactOpts, alertHeader, nonSignerStakesAndSignature)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractMachServiceManager.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractMachServiceManager.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractMachServiceManager.TransactOpts, operator)
}

// DisableAllowlist is a paid mutator transaction binding the contract method 0xcf8e629a.
//
// Solidity: function disableAllowlist() returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactor) DisableAllowlist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractMachServiceManager.contract.Transact(opts, "disableAllowlist")
}

// DisableAllowlist is a paid mutator transaction binding the contract method 0xcf8e629a.
//
// Solidity: function disableAllowlist() returns()
func (_ContractMachServiceManager *ContractMachServiceManagerSession) DisableAllowlist() (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.DisableAllowlist(&_ContractMachServiceManager.TransactOpts)
}

// DisableAllowlist is a paid mutator transaction binding the contract method 0xcf8e629a.
//
// Solidity: function disableAllowlist() returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactorSession) DisableAllowlist() (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.DisableAllowlist(&_ContractMachServiceManager.TransactOpts)
}

// EnableAllowlist is a paid mutator transaction binding the contract method 0xc6a2aac8.
//
// Solidity: function enableAllowlist() returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactor) EnableAllowlist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractMachServiceManager.contract.Transact(opts, "enableAllowlist")
}

// EnableAllowlist is a paid mutator transaction binding the contract method 0xc6a2aac8.
//
// Solidity: function enableAllowlist() returns()
func (_ContractMachServiceManager *ContractMachServiceManagerSession) EnableAllowlist() (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.EnableAllowlist(&_ContractMachServiceManager.TransactOpts)
}

// EnableAllowlist is a paid mutator transaction binding the contract method 0xc6a2aac8.
//
// Solidity: function enableAllowlist() returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactorSession) EnableAllowlist() (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.EnableAllowlist(&_ContractMachServiceManager.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x358394d8.
//
// Solidity: function initialize(address _pauserRegistry, uint256 _initialPausedStatus, address _initialOwner, address _alertConfirmer) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, _initialPausedStatus *big.Int, _initialOwner common.Address, _alertConfirmer common.Address) (*types.Transaction, error) {
	return _ContractMachServiceManager.contract.Transact(opts, "initialize", _pauserRegistry, _initialPausedStatus, _initialOwner, _alertConfirmer)
}

// Initialize is a paid mutator transaction binding the contract method 0x358394d8.
//
// Solidity: function initialize(address _pauserRegistry, uint256 _initialPausedStatus, address _initialOwner, address _alertConfirmer) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerSession) Initialize(_pauserRegistry common.Address, _initialPausedStatus *big.Int, _initialOwner common.Address, _alertConfirmer common.Address) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.Initialize(&_ContractMachServiceManager.TransactOpts, _pauserRegistry, _initialPausedStatus, _initialOwner, _alertConfirmer)
}

// Initialize is a paid mutator transaction binding the contract method 0x358394d8.
//
// Solidity: function initialize(address _pauserRegistry, uint256 _initialPausedStatus, address _initialOwner, address _alertConfirmer) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactorSession) Initialize(_pauserRegistry common.Address, _initialPausedStatus *big.Int, _initialOwner common.Address, _alertConfirmer common.Address) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.Initialize(&_ContractMachServiceManager.TransactOpts, _pauserRegistry, _initialPausedStatus, _initialOwner, _alertConfirmer)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractMachServiceManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.Pause(&_ContractMachServiceManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.Pause(&_ContractMachServiceManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractMachServiceManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractMachServiceManager *ContractMachServiceManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.PauseAll(&_ContractMachServiceManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.PauseAll(&_ContractMachServiceManager.TransactOpts)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractMachServiceManager.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.RegisterOperatorToAVS(&_ContractMachServiceManager.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.RegisterOperatorToAVS(&_ContractMachServiceManager.TransactOpts, operator, operatorSignature)
}

// RemoveAlert is a paid mutator transaction binding the contract method 0xd8ff953d.
//
// Solidity: function removeAlert(bytes32 messageHash) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactor) RemoveAlert(opts *bind.TransactOpts, messageHash [32]byte) (*types.Transaction, error) {
	return _ContractMachServiceManager.contract.Transact(opts, "removeAlert", messageHash)
}

// RemoveAlert is a paid mutator transaction binding the contract method 0xd8ff953d.
//
// Solidity: function removeAlert(bytes32 messageHash) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerSession) RemoveAlert(messageHash [32]byte) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.RemoveAlert(&_ContractMachServiceManager.TransactOpts, messageHash)
}

// RemoveAlert is a paid mutator transaction binding the contract method 0xd8ff953d.
//
// Solidity: function removeAlert(bytes32 messageHash) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactorSession) RemoveAlert(messageHash [32]byte) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.RemoveAlert(&_ContractMachServiceManager.TransactOpts, messageHash)
}

// RemoveFromAllowlist is a paid mutator transaction binding the contract method 0x5da93d7e.
//
// Solidity: function removeFromAllowlist(address operator) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactor) RemoveFromAllowlist(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractMachServiceManager.contract.Transact(opts, "removeFromAllowlist", operator)
}

// RemoveFromAllowlist is a paid mutator transaction binding the contract method 0x5da93d7e.
//
// Solidity: function removeFromAllowlist(address operator) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerSession) RemoveFromAllowlist(operator common.Address) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.RemoveFromAllowlist(&_ContractMachServiceManager.TransactOpts, operator)
}

// RemoveFromAllowlist is a paid mutator transaction binding the contract method 0x5da93d7e.
//
// Solidity: function removeFromAllowlist(address operator) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactorSession) RemoveFromAllowlist(operator common.Address) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.RemoveFromAllowlist(&_ContractMachServiceManager.TransactOpts, operator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractMachServiceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractMachServiceManager *ContractMachServiceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.RenounceOwnership(&_ContractMachServiceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.RenounceOwnership(&_ContractMachServiceManager.TransactOpts)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractMachServiceManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.SetPauserRegistry(&_ContractMachServiceManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.SetPauserRegistry(&_ContractMachServiceManager.TransactOpts, newPauserRegistry)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractMachServiceManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.SetStaleStakesForbidden(&_ContractMachServiceManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.SetStaleStakesForbidden(&_ContractMachServiceManager.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractMachServiceManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.TransferOwnership(&_ContractMachServiceManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.TransferOwnership(&_ContractMachServiceManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractMachServiceManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.Unpause(&_ContractMachServiceManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.Unpause(&_ContractMachServiceManager.TransactOpts, newPausedStatus)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ContractMachServiceManager.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.UpdateAVSMetadataURI(&_ContractMachServiceManager.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.UpdateAVSMetadataURI(&_ContractMachServiceManager.TransactOpts, _metadataURI)
}

// UpdateQuorumThresholdPercentage is a paid mutator transaction binding the contract method 0x429d5bf0.
//
// Solidity: function updateQuorumThresholdPercentage(uint8 thresholdPercentage) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactor) UpdateQuorumThresholdPercentage(opts *bind.TransactOpts, thresholdPercentage uint8) (*types.Transaction, error) {
	return _ContractMachServiceManager.contract.Transact(opts, "updateQuorumThresholdPercentage", thresholdPercentage)
}

// UpdateQuorumThresholdPercentage is a paid mutator transaction binding the contract method 0x429d5bf0.
//
// Solidity: function updateQuorumThresholdPercentage(uint8 thresholdPercentage) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerSession) UpdateQuorumThresholdPercentage(thresholdPercentage uint8) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.UpdateQuorumThresholdPercentage(&_ContractMachServiceManager.TransactOpts, thresholdPercentage)
}

// UpdateQuorumThresholdPercentage is a paid mutator transaction binding the contract method 0x429d5bf0.
//
// Solidity: function updateQuorumThresholdPercentage(uint8 thresholdPercentage) returns()
func (_ContractMachServiceManager *ContractMachServiceManagerTransactorSession) UpdateQuorumThresholdPercentage(thresholdPercentage uint8) (*types.Transaction, error) {
	return _ContractMachServiceManager.Contract.UpdateQuorumThresholdPercentage(&_ContractMachServiceManager.TransactOpts, thresholdPercentage)
}

// ContractMachServiceManagerAlertConfirmedIterator is returned from FilterAlertConfirmed and is used to iterate over the raw logs and unpacked data for AlertConfirmed events raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerAlertConfirmedIterator struct {
	Event *ContractMachServiceManagerAlertConfirmed // Event containing the contract specifics and raw log

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
func (it *ContractMachServiceManagerAlertConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMachServiceManagerAlertConfirmed)
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
		it.Event = new(ContractMachServiceManagerAlertConfirmed)
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
func (it *ContractMachServiceManagerAlertConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMachServiceManagerAlertConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMachServiceManagerAlertConfirmed represents a AlertConfirmed event raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerAlertConfirmed struct {
	AlertHeaderHash [32]byte
	MessageHash     [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAlertConfirmed is a free log retrieval operation binding the contract event 0xfdda6f7d4825a4f1e4e97b50a26a69a8bcc3f4fcb1113cc14ce8e7098ca63665.
//
// Solidity: event AlertConfirmed(bytes32 indexed alertHeaderHash, bytes32 messageHash)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) FilterAlertConfirmed(opts *bind.FilterOpts, alertHeaderHash [][32]byte) (*ContractMachServiceManagerAlertConfirmedIterator, error) {

	var alertHeaderHashRule []interface{}
	for _, alertHeaderHashItem := range alertHeaderHash {
		alertHeaderHashRule = append(alertHeaderHashRule, alertHeaderHashItem)
	}

	logs, sub, err := _ContractMachServiceManager.contract.FilterLogs(opts, "AlertConfirmed", alertHeaderHashRule)
	if err != nil {
		return nil, err
	}
	return &ContractMachServiceManagerAlertConfirmedIterator{contract: _ContractMachServiceManager.contract, event: "AlertConfirmed", logs: logs, sub: sub}, nil
}

// WatchAlertConfirmed is a free log subscription operation binding the contract event 0xfdda6f7d4825a4f1e4e97b50a26a69a8bcc3f4fcb1113cc14ce8e7098ca63665.
//
// Solidity: event AlertConfirmed(bytes32 indexed alertHeaderHash, bytes32 messageHash)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) WatchAlertConfirmed(opts *bind.WatchOpts, sink chan<- *ContractMachServiceManagerAlertConfirmed, alertHeaderHash [][32]byte) (event.Subscription, error) {

	var alertHeaderHashRule []interface{}
	for _, alertHeaderHashItem := range alertHeaderHash {
		alertHeaderHashRule = append(alertHeaderHashRule, alertHeaderHashItem)
	}

	logs, sub, err := _ContractMachServiceManager.contract.WatchLogs(opts, "AlertConfirmed", alertHeaderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMachServiceManagerAlertConfirmed)
				if err := _ContractMachServiceManager.contract.UnpackLog(event, "AlertConfirmed", log); err != nil {
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

// ParseAlertConfirmed is a log parse operation binding the contract event 0xfdda6f7d4825a4f1e4e97b50a26a69a8bcc3f4fcb1113cc14ce8e7098ca63665.
//
// Solidity: event AlertConfirmed(bytes32 indexed alertHeaderHash, bytes32 messageHash)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) ParseAlertConfirmed(log types.Log) (*ContractMachServiceManagerAlertConfirmed, error) {
	event := new(ContractMachServiceManagerAlertConfirmed)
	if err := _ContractMachServiceManager.contract.UnpackLog(event, "AlertConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMachServiceManagerAlertConfirmerChangedIterator is returned from FilterAlertConfirmerChanged and is used to iterate over the raw logs and unpacked data for AlertConfirmerChanged events raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerAlertConfirmerChangedIterator struct {
	Event *ContractMachServiceManagerAlertConfirmerChanged // Event containing the contract specifics and raw log

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
func (it *ContractMachServiceManagerAlertConfirmerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMachServiceManagerAlertConfirmerChanged)
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
		it.Event = new(ContractMachServiceManagerAlertConfirmerChanged)
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
func (it *ContractMachServiceManagerAlertConfirmerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMachServiceManagerAlertConfirmerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMachServiceManagerAlertConfirmerChanged represents a AlertConfirmerChanged event raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerAlertConfirmerChanged struct {
	PreviousAddress common.Address
	NewAddress      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAlertConfirmerChanged is a free log retrieval operation binding the contract event 0x175f27847b3568e0da876ffca3dc0bb52db4e6556346aedb530c6fe86610da27.
//
// Solidity: event AlertConfirmerChanged(address previousAddress, address newAddress)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) FilterAlertConfirmerChanged(opts *bind.FilterOpts) (*ContractMachServiceManagerAlertConfirmerChangedIterator, error) {

	logs, sub, err := _ContractMachServiceManager.contract.FilterLogs(opts, "AlertConfirmerChanged")
	if err != nil {
		return nil, err
	}
	return &ContractMachServiceManagerAlertConfirmerChangedIterator{contract: _ContractMachServiceManager.contract, event: "AlertConfirmerChanged", logs: logs, sub: sub}, nil
}

// WatchAlertConfirmerChanged is a free log subscription operation binding the contract event 0x175f27847b3568e0da876ffca3dc0bb52db4e6556346aedb530c6fe86610da27.
//
// Solidity: event AlertConfirmerChanged(address previousAddress, address newAddress)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) WatchAlertConfirmerChanged(opts *bind.WatchOpts, sink chan<- *ContractMachServiceManagerAlertConfirmerChanged) (event.Subscription, error) {

	logs, sub, err := _ContractMachServiceManager.contract.WatchLogs(opts, "AlertConfirmerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMachServiceManagerAlertConfirmerChanged)
				if err := _ContractMachServiceManager.contract.UnpackLog(event, "AlertConfirmerChanged", log); err != nil {
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

// ParseAlertConfirmerChanged is a log parse operation binding the contract event 0x175f27847b3568e0da876ffca3dc0bb52db4e6556346aedb530c6fe86610da27.
//
// Solidity: event AlertConfirmerChanged(address previousAddress, address newAddress)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) ParseAlertConfirmerChanged(log types.Log) (*ContractMachServiceManagerAlertConfirmerChanged, error) {
	event := new(ContractMachServiceManagerAlertConfirmerChanged)
	if err := _ContractMachServiceManager.contract.UnpackLog(event, "AlertConfirmerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMachServiceManagerAlertRemovedIterator is returned from FilterAlertRemoved and is used to iterate over the raw logs and unpacked data for AlertRemoved events raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerAlertRemovedIterator struct {
	Event *ContractMachServiceManagerAlertRemoved // Event containing the contract specifics and raw log

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
func (it *ContractMachServiceManagerAlertRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMachServiceManagerAlertRemoved)
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
		it.Event = new(ContractMachServiceManagerAlertRemoved)
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
func (it *ContractMachServiceManagerAlertRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMachServiceManagerAlertRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMachServiceManagerAlertRemoved represents a AlertRemoved event raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerAlertRemoved struct {
	MessageHash [32]byte
	Sender      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAlertRemoved is a free log retrieval operation binding the contract event 0x1bdeffc0337373bf2f6fd4219080133eeaaee0554206d9bb24a019d96973c1eb.
//
// Solidity: event AlertRemoved(bytes32 messageHash, address sender)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) FilterAlertRemoved(opts *bind.FilterOpts) (*ContractMachServiceManagerAlertRemovedIterator, error) {

	logs, sub, err := _ContractMachServiceManager.contract.FilterLogs(opts, "AlertRemoved")
	if err != nil {
		return nil, err
	}
	return &ContractMachServiceManagerAlertRemovedIterator{contract: _ContractMachServiceManager.contract, event: "AlertRemoved", logs: logs, sub: sub}, nil
}

// WatchAlertRemoved is a free log subscription operation binding the contract event 0x1bdeffc0337373bf2f6fd4219080133eeaaee0554206d9bb24a019d96973c1eb.
//
// Solidity: event AlertRemoved(bytes32 messageHash, address sender)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) WatchAlertRemoved(opts *bind.WatchOpts, sink chan<- *ContractMachServiceManagerAlertRemoved) (event.Subscription, error) {

	logs, sub, err := _ContractMachServiceManager.contract.WatchLogs(opts, "AlertRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMachServiceManagerAlertRemoved)
				if err := _ContractMachServiceManager.contract.UnpackLog(event, "AlertRemoved", log); err != nil {
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

// ParseAlertRemoved is a log parse operation binding the contract event 0x1bdeffc0337373bf2f6fd4219080133eeaaee0554206d9bb24a019d96973c1eb.
//
// Solidity: event AlertRemoved(bytes32 messageHash, address sender)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) ParseAlertRemoved(log types.Log) (*ContractMachServiceManagerAlertRemoved, error) {
	event := new(ContractMachServiceManagerAlertRemoved)
	if err := _ContractMachServiceManager.contract.UnpackLog(event, "AlertRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMachServiceManagerAllowlistDisabledIterator is returned from FilterAllowlistDisabled and is used to iterate over the raw logs and unpacked data for AllowlistDisabled events raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerAllowlistDisabledIterator struct {
	Event *ContractMachServiceManagerAllowlistDisabled // Event containing the contract specifics and raw log

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
func (it *ContractMachServiceManagerAllowlistDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMachServiceManagerAllowlistDisabled)
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
		it.Event = new(ContractMachServiceManagerAllowlistDisabled)
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
func (it *ContractMachServiceManagerAllowlistDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMachServiceManagerAllowlistDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMachServiceManagerAllowlistDisabled represents a AllowlistDisabled event raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerAllowlistDisabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAllowlistDisabled is a free log retrieval operation binding the contract event 0x2d35c8d348a345fd7b3b03b7cfcf7ad0b60c2d46742d5ca536342e4185becb07.
//
// Solidity: event AllowlistDisabled()
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) FilterAllowlistDisabled(opts *bind.FilterOpts) (*ContractMachServiceManagerAllowlistDisabledIterator, error) {

	logs, sub, err := _ContractMachServiceManager.contract.FilterLogs(opts, "AllowlistDisabled")
	if err != nil {
		return nil, err
	}
	return &ContractMachServiceManagerAllowlistDisabledIterator{contract: _ContractMachServiceManager.contract, event: "AllowlistDisabled", logs: logs, sub: sub}, nil
}

// WatchAllowlistDisabled is a free log subscription operation binding the contract event 0x2d35c8d348a345fd7b3b03b7cfcf7ad0b60c2d46742d5ca536342e4185becb07.
//
// Solidity: event AllowlistDisabled()
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) WatchAllowlistDisabled(opts *bind.WatchOpts, sink chan<- *ContractMachServiceManagerAllowlistDisabled) (event.Subscription, error) {

	logs, sub, err := _ContractMachServiceManager.contract.WatchLogs(opts, "AllowlistDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMachServiceManagerAllowlistDisabled)
				if err := _ContractMachServiceManager.contract.UnpackLog(event, "AllowlistDisabled", log); err != nil {
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

// ParseAllowlistDisabled is a log parse operation binding the contract event 0x2d35c8d348a345fd7b3b03b7cfcf7ad0b60c2d46742d5ca536342e4185becb07.
//
// Solidity: event AllowlistDisabled()
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) ParseAllowlistDisabled(log types.Log) (*ContractMachServiceManagerAllowlistDisabled, error) {
	event := new(ContractMachServiceManagerAllowlistDisabled)
	if err := _ContractMachServiceManager.contract.UnpackLog(event, "AllowlistDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMachServiceManagerAllowlistEnabledIterator is returned from FilterAllowlistEnabled and is used to iterate over the raw logs and unpacked data for AllowlistEnabled events raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerAllowlistEnabledIterator struct {
	Event *ContractMachServiceManagerAllowlistEnabled // Event containing the contract specifics and raw log

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
func (it *ContractMachServiceManagerAllowlistEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMachServiceManagerAllowlistEnabled)
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
		it.Event = new(ContractMachServiceManagerAllowlistEnabled)
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
func (it *ContractMachServiceManagerAllowlistEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMachServiceManagerAllowlistEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMachServiceManagerAllowlistEnabled represents a AllowlistEnabled event raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerAllowlistEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAllowlistEnabled is a free log retrieval operation binding the contract event 0x8a943acd5f4e6d3df7565a4a08a93f6b04cc31bb6c01ca4aef7abd6baf455ec3.
//
// Solidity: event AllowlistEnabled()
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) FilterAllowlistEnabled(opts *bind.FilterOpts) (*ContractMachServiceManagerAllowlistEnabledIterator, error) {

	logs, sub, err := _ContractMachServiceManager.contract.FilterLogs(opts, "AllowlistEnabled")
	if err != nil {
		return nil, err
	}
	return &ContractMachServiceManagerAllowlistEnabledIterator{contract: _ContractMachServiceManager.contract, event: "AllowlistEnabled", logs: logs, sub: sub}, nil
}

// WatchAllowlistEnabled is a free log subscription operation binding the contract event 0x8a943acd5f4e6d3df7565a4a08a93f6b04cc31bb6c01ca4aef7abd6baf455ec3.
//
// Solidity: event AllowlistEnabled()
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) WatchAllowlistEnabled(opts *bind.WatchOpts, sink chan<- *ContractMachServiceManagerAllowlistEnabled) (event.Subscription, error) {

	logs, sub, err := _ContractMachServiceManager.contract.WatchLogs(opts, "AllowlistEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMachServiceManagerAllowlistEnabled)
				if err := _ContractMachServiceManager.contract.UnpackLog(event, "AllowlistEnabled", log); err != nil {
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

// ParseAllowlistEnabled is a log parse operation binding the contract event 0x8a943acd5f4e6d3df7565a4a08a93f6b04cc31bb6c01ca4aef7abd6baf455ec3.
//
// Solidity: event AllowlistEnabled()
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) ParseAllowlistEnabled(log types.Log) (*ContractMachServiceManagerAllowlistEnabled, error) {
	event := new(ContractMachServiceManagerAllowlistEnabled)
	if err := _ContractMachServiceManager.contract.UnpackLog(event, "AllowlistEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMachServiceManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerInitializedIterator struct {
	Event *ContractMachServiceManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractMachServiceManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMachServiceManagerInitialized)
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
		it.Event = new(ContractMachServiceManagerInitialized)
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
func (it *ContractMachServiceManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMachServiceManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMachServiceManagerInitialized represents a Initialized event raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractMachServiceManagerInitializedIterator, error) {

	logs, sub, err := _ContractMachServiceManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractMachServiceManagerInitializedIterator{contract: _ContractMachServiceManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractMachServiceManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractMachServiceManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMachServiceManagerInitialized)
				if err := _ContractMachServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) ParseInitialized(log types.Log) (*ContractMachServiceManagerInitialized, error) {
	event := new(ContractMachServiceManagerInitialized)
	if err := _ContractMachServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMachServiceManagerOperatorAddedIterator is returned from FilterOperatorAdded and is used to iterate over the raw logs and unpacked data for OperatorAdded events raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerOperatorAddedIterator struct {
	Event *ContractMachServiceManagerOperatorAdded // Event containing the contract specifics and raw log

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
func (it *ContractMachServiceManagerOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMachServiceManagerOperatorAdded)
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
		it.Event = new(ContractMachServiceManagerOperatorAdded)
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
func (it *ContractMachServiceManagerOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMachServiceManagerOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMachServiceManagerOperatorAdded represents a OperatorAdded event raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerOperatorAdded struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdded is a free log retrieval operation binding the contract event 0xac6fa858e9350a46cec16539926e0fde25b7629f84b5a72bffaae4df888ae86d.
//
// Solidity: event OperatorAdded(address indexed operator)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) FilterOperatorAdded(opts *bind.FilterOpts, operator []common.Address) (*ContractMachServiceManagerOperatorAddedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractMachServiceManager.contract.FilterLogs(opts, "OperatorAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractMachServiceManagerOperatorAddedIterator{contract: _ContractMachServiceManager.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

// WatchOperatorAdded is a free log subscription operation binding the contract event 0xac6fa858e9350a46cec16539926e0fde25b7629f84b5a72bffaae4df888ae86d.
//
// Solidity: event OperatorAdded(address indexed operator)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *ContractMachServiceManagerOperatorAdded, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractMachServiceManager.contract.WatchLogs(opts, "OperatorAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMachServiceManagerOperatorAdded)
				if err := _ContractMachServiceManager.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
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

// ParseOperatorAdded is a log parse operation binding the contract event 0xac6fa858e9350a46cec16539926e0fde25b7629f84b5a72bffaae4df888ae86d.
//
// Solidity: event OperatorAdded(address indexed operator)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) ParseOperatorAdded(log types.Log) (*ContractMachServiceManagerOperatorAdded, error) {
	event := new(ContractMachServiceManagerOperatorAdded)
	if err := _ContractMachServiceManager.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMachServiceManagerOperatorAllowedIterator is returned from FilterOperatorAllowed and is used to iterate over the raw logs and unpacked data for OperatorAllowed events raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerOperatorAllowedIterator struct {
	Event *ContractMachServiceManagerOperatorAllowed // Event containing the contract specifics and raw log

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
func (it *ContractMachServiceManagerOperatorAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMachServiceManagerOperatorAllowed)
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
		it.Event = new(ContractMachServiceManagerOperatorAllowed)
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
func (it *ContractMachServiceManagerOperatorAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMachServiceManagerOperatorAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMachServiceManagerOperatorAllowed represents a OperatorAllowed event raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerOperatorAllowed struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorAllowed is a free log retrieval operation binding the contract event 0xdde65206cdee4ea27ef1b170724ba50b41ad09a3bf2dda12935fc40c4dbf6e75.
//
// Solidity: event OperatorAllowed(address operator)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) FilterOperatorAllowed(opts *bind.FilterOpts) (*ContractMachServiceManagerOperatorAllowedIterator, error) {

	logs, sub, err := _ContractMachServiceManager.contract.FilterLogs(opts, "OperatorAllowed")
	if err != nil {
		return nil, err
	}
	return &ContractMachServiceManagerOperatorAllowedIterator{contract: _ContractMachServiceManager.contract, event: "OperatorAllowed", logs: logs, sub: sub}, nil
}

// WatchOperatorAllowed is a free log subscription operation binding the contract event 0xdde65206cdee4ea27ef1b170724ba50b41ad09a3bf2dda12935fc40c4dbf6e75.
//
// Solidity: event OperatorAllowed(address operator)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) WatchOperatorAllowed(opts *bind.WatchOpts, sink chan<- *ContractMachServiceManagerOperatorAllowed) (event.Subscription, error) {

	logs, sub, err := _ContractMachServiceManager.contract.WatchLogs(opts, "OperatorAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMachServiceManagerOperatorAllowed)
				if err := _ContractMachServiceManager.contract.UnpackLog(event, "OperatorAllowed", log); err != nil {
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

// ParseOperatorAllowed is a log parse operation binding the contract event 0xdde65206cdee4ea27ef1b170724ba50b41ad09a3bf2dda12935fc40c4dbf6e75.
//
// Solidity: event OperatorAllowed(address operator)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) ParseOperatorAllowed(log types.Log) (*ContractMachServiceManagerOperatorAllowed, error) {
	event := new(ContractMachServiceManagerOperatorAllowed)
	if err := _ContractMachServiceManager.contract.UnpackLog(event, "OperatorAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMachServiceManagerOperatorDisallowedIterator is returned from FilterOperatorDisallowed and is used to iterate over the raw logs and unpacked data for OperatorDisallowed events raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerOperatorDisallowedIterator struct {
	Event *ContractMachServiceManagerOperatorDisallowed // Event containing the contract specifics and raw log

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
func (it *ContractMachServiceManagerOperatorDisallowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMachServiceManagerOperatorDisallowed)
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
		it.Event = new(ContractMachServiceManagerOperatorDisallowed)
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
func (it *ContractMachServiceManagerOperatorDisallowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMachServiceManagerOperatorDisallowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMachServiceManagerOperatorDisallowed represents a OperatorDisallowed event raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerOperatorDisallowed struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorDisallowed is a free log retrieval operation binding the contract event 0x8560daa191dd8e6fba276b053006b3990c46c94b842f85490f52c49b15cfe5cb.
//
// Solidity: event OperatorDisallowed(address operator)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) FilterOperatorDisallowed(opts *bind.FilterOpts) (*ContractMachServiceManagerOperatorDisallowedIterator, error) {

	logs, sub, err := _ContractMachServiceManager.contract.FilterLogs(opts, "OperatorDisallowed")
	if err != nil {
		return nil, err
	}
	return &ContractMachServiceManagerOperatorDisallowedIterator{contract: _ContractMachServiceManager.contract, event: "OperatorDisallowed", logs: logs, sub: sub}, nil
}

// WatchOperatorDisallowed is a free log subscription operation binding the contract event 0x8560daa191dd8e6fba276b053006b3990c46c94b842f85490f52c49b15cfe5cb.
//
// Solidity: event OperatorDisallowed(address operator)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) WatchOperatorDisallowed(opts *bind.WatchOpts, sink chan<- *ContractMachServiceManagerOperatorDisallowed) (event.Subscription, error) {

	logs, sub, err := _ContractMachServiceManager.contract.WatchLogs(opts, "OperatorDisallowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMachServiceManagerOperatorDisallowed)
				if err := _ContractMachServiceManager.contract.UnpackLog(event, "OperatorDisallowed", log); err != nil {
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

// ParseOperatorDisallowed is a log parse operation binding the contract event 0x8560daa191dd8e6fba276b053006b3990c46c94b842f85490f52c49b15cfe5cb.
//
// Solidity: event OperatorDisallowed(address operator)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) ParseOperatorDisallowed(log types.Log) (*ContractMachServiceManagerOperatorDisallowed, error) {
	event := new(ContractMachServiceManagerOperatorDisallowed)
	if err := _ContractMachServiceManager.contract.UnpackLog(event, "OperatorDisallowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMachServiceManagerOperatorRemovedIterator is returned from FilterOperatorRemoved and is used to iterate over the raw logs and unpacked data for OperatorRemoved events raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerOperatorRemovedIterator struct {
	Event *ContractMachServiceManagerOperatorRemoved // Event containing the contract specifics and raw log

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
func (it *ContractMachServiceManagerOperatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMachServiceManagerOperatorRemoved)
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
		it.Event = new(ContractMachServiceManagerOperatorRemoved)
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
func (it *ContractMachServiceManagerOperatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMachServiceManagerOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMachServiceManagerOperatorRemoved represents a OperatorRemoved event raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerOperatorRemoved struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemoved is a free log retrieval operation binding the contract event 0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d.
//
// Solidity: event OperatorRemoved(address indexed operator)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) FilterOperatorRemoved(opts *bind.FilterOpts, operator []common.Address) (*ContractMachServiceManagerOperatorRemovedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractMachServiceManager.contract.FilterLogs(opts, "OperatorRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractMachServiceManagerOperatorRemovedIterator{contract: _ContractMachServiceManager.contract, event: "OperatorRemoved", logs: logs, sub: sub}, nil
}

// WatchOperatorRemoved is a free log subscription operation binding the contract event 0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d.
//
// Solidity: event OperatorRemoved(address indexed operator)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) WatchOperatorRemoved(opts *bind.WatchOpts, sink chan<- *ContractMachServiceManagerOperatorRemoved, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractMachServiceManager.contract.WatchLogs(opts, "OperatorRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMachServiceManagerOperatorRemoved)
				if err := _ContractMachServiceManager.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
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

// ParseOperatorRemoved is a log parse operation binding the contract event 0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d.
//
// Solidity: event OperatorRemoved(address indexed operator)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) ParseOperatorRemoved(log types.Log) (*ContractMachServiceManagerOperatorRemoved, error) {
	event := new(ContractMachServiceManagerOperatorRemoved)
	if err := _ContractMachServiceManager.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMachServiceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerOwnershipTransferredIterator struct {
	Event *ContractMachServiceManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractMachServiceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMachServiceManagerOwnershipTransferred)
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
		it.Event = new(ContractMachServiceManagerOwnershipTransferred)
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
func (it *ContractMachServiceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMachServiceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMachServiceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractMachServiceManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractMachServiceManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractMachServiceManagerOwnershipTransferredIterator{contract: _ContractMachServiceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractMachServiceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractMachServiceManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMachServiceManagerOwnershipTransferred)
				if err := _ContractMachServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractMachServiceManagerOwnershipTransferred, error) {
	event := new(ContractMachServiceManagerOwnershipTransferred)
	if err := _ContractMachServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMachServiceManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerPausedIterator struct {
	Event *ContractMachServiceManagerPaused // Event containing the contract specifics and raw log

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
func (it *ContractMachServiceManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMachServiceManagerPaused)
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
		it.Event = new(ContractMachServiceManagerPaused)
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
func (it *ContractMachServiceManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMachServiceManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMachServiceManagerPaused represents a Paused event raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractMachServiceManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractMachServiceManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractMachServiceManagerPausedIterator{contract: _ContractMachServiceManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractMachServiceManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractMachServiceManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMachServiceManagerPaused)
				if err := _ContractMachServiceManager.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) ParsePaused(log types.Log) (*ContractMachServiceManagerPaused, error) {
	event := new(ContractMachServiceManagerPaused)
	if err := _ContractMachServiceManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMachServiceManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerPauserRegistrySetIterator struct {
	Event *ContractMachServiceManagerPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *ContractMachServiceManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMachServiceManagerPauserRegistrySet)
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
		it.Event = new(ContractMachServiceManagerPauserRegistrySet)
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
func (it *ContractMachServiceManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMachServiceManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMachServiceManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractMachServiceManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractMachServiceManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractMachServiceManagerPauserRegistrySetIterator{contract: _ContractMachServiceManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractMachServiceManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractMachServiceManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMachServiceManagerPauserRegistrySet)
				if err := _ContractMachServiceManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractMachServiceManagerPauserRegistrySet, error) {
	event := new(ContractMachServiceManagerPauserRegistrySet)
	if err := _ContractMachServiceManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMachServiceManagerQuorumThresholdPercentageChangedIterator is returned from FilterQuorumThresholdPercentageChanged and is used to iterate over the raw logs and unpacked data for QuorumThresholdPercentageChanged events raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerQuorumThresholdPercentageChangedIterator struct {
	Event *ContractMachServiceManagerQuorumThresholdPercentageChanged // Event containing the contract specifics and raw log

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
func (it *ContractMachServiceManagerQuorumThresholdPercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMachServiceManagerQuorumThresholdPercentageChanged)
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
		it.Event = new(ContractMachServiceManagerQuorumThresholdPercentageChanged)
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
func (it *ContractMachServiceManagerQuorumThresholdPercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMachServiceManagerQuorumThresholdPercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMachServiceManagerQuorumThresholdPercentageChanged represents a QuorumThresholdPercentageChanged event raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerQuorumThresholdPercentageChanged struct {
	ThresholdPercentages uint8
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterQuorumThresholdPercentageChanged is a free log retrieval operation binding the contract event 0xc3acdc4f4bc283baa27c4207eb2c32954fbb26960847c9e15c2f7c8970134244.
//
// Solidity: event QuorumThresholdPercentageChanged(uint8 thresholdPercentages)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) FilterQuorumThresholdPercentageChanged(opts *bind.FilterOpts) (*ContractMachServiceManagerQuorumThresholdPercentageChangedIterator, error) {

	logs, sub, err := _ContractMachServiceManager.contract.FilterLogs(opts, "QuorumThresholdPercentageChanged")
	if err != nil {
		return nil, err
	}
	return &ContractMachServiceManagerQuorumThresholdPercentageChangedIterator{contract: _ContractMachServiceManager.contract, event: "QuorumThresholdPercentageChanged", logs: logs, sub: sub}, nil
}

// WatchQuorumThresholdPercentageChanged is a free log subscription operation binding the contract event 0xc3acdc4f4bc283baa27c4207eb2c32954fbb26960847c9e15c2f7c8970134244.
//
// Solidity: event QuorumThresholdPercentageChanged(uint8 thresholdPercentages)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) WatchQuorumThresholdPercentageChanged(opts *bind.WatchOpts, sink chan<- *ContractMachServiceManagerQuorumThresholdPercentageChanged) (event.Subscription, error) {

	logs, sub, err := _ContractMachServiceManager.contract.WatchLogs(opts, "QuorumThresholdPercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMachServiceManagerQuorumThresholdPercentageChanged)
				if err := _ContractMachServiceManager.contract.UnpackLog(event, "QuorumThresholdPercentageChanged", log); err != nil {
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

// ParseQuorumThresholdPercentageChanged is a log parse operation binding the contract event 0xc3acdc4f4bc283baa27c4207eb2c32954fbb26960847c9e15c2f7c8970134244.
//
// Solidity: event QuorumThresholdPercentageChanged(uint8 thresholdPercentages)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) ParseQuorumThresholdPercentageChanged(log types.Log) (*ContractMachServiceManagerQuorumThresholdPercentageChanged, error) {
	event := new(ContractMachServiceManagerQuorumThresholdPercentageChanged)
	if err := _ContractMachServiceManager.contract.UnpackLog(event, "QuorumThresholdPercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMachServiceManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractMachServiceManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

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
func (it *ContractMachServiceManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMachServiceManagerStaleStakesForbiddenUpdate)
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
		it.Event = new(ContractMachServiceManagerStaleStakesForbiddenUpdate)
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
func (it *ContractMachServiceManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMachServiceManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMachServiceManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractMachServiceManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractMachServiceManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractMachServiceManagerStaleStakesForbiddenUpdateIterator{contract: _ContractMachServiceManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractMachServiceManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractMachServiceManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMachServiceManagerStaleStakesForbiddenUpdate)
				if err := _ContractMachServiceManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
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

// ParseStaleStakesForbiddenUpdate is a log parse operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractMachServiceManagerStaleStakesForbiddenUpdate, error) {
	event := new(ContractMachServiceManagerStaleStakesForbiddenUpdate)
	if err := _ContractMachServiceManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMachServiceManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerUnpausedIterator struct {
	Event *ContractMachServiceManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractMachServiceManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMachServiceManagerUnpaused)
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
		it.Event = new(ContractMachServiceManagerUnpaused)
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
func (it *ContractMachServiceManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMachServiceManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMachServiceManagerUnpaused represents a Unpaused event raised by the ContractMachServiceManager contract.
type ContractMachServiceManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractMachServiceManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractMachServiceManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractMachServiceManagerUnpausedIterator{contract: _ContractMachServiceManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractMachServiceManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractMachServiceManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMachServiceManagerUnpaused)
				if err := _ContractMachServiceManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractMachServiceManager *ContractMachServiceManagerFilterer) ParseUnpaused(log types.Log) (*ContractMachServiceManagerUnpaused, error) {
	event := new(ContractMachServiceManagerUnpaused)
	if err := _ContractMachServiceManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
