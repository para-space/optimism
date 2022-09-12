// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L1BlockStorageLayoutJSON = "{\"storage\":[{\"astId\":1898,\"contract\":\"contracts/L2/L1Block.sol:L1Block\",\"label\":\"number\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint64\"},{\"astId\":1901,\"contract\":\"contracts/L2/L1Block.sol:L1Block\",\"label\":\"timestamp\",\"offset\":8,\"slot\":\"0\",\"type\":\"t_uint64\"},{\"astId\":1904,\"contract\":\"contracts/L2/L1Block.sol:L1Block\",\"label\":\"basefee\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"},{\"astId\":1907,\"contract\":\"contracts/L2/L1Block.sol:L1Block\",\"label\":\"hash\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_bytes32\"},{\"astId\":1910,\"contract\":\"contracts/L2/L1Block.sol:L1Block\",\"label\":\"sequenceNumber\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_uint64\"}],\"types\":{\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"}}}"

var L1BlockStorageLayout = new(solc.StorageLayout)

var L1BlockDeployedBin = "0x608060405234801561001057600080fd5b50600436106100885760003560e01c806364ca23ef1161005b57806364ca23ef146100dc5780638381f58a14610109578063b80777ea1461011d578063e591b2821461013d57600080fd5b8063042c2f571461008d57806309bd5a60146100a257806354fd4d50146100be5780635cf24969146100d3575b600080fd5b6100a061009b3660046104a6565b61017d565b005b6100ab60025481565b6040519081526020015b60405180910390f35b6100c66102a9565b6040516100b5919061052d565b6100ab60015481565b6003546100f09067ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020016100b5565b6000546100f09067ffffffffffffffff1681565b6000546100f09068010000000000000000900467ffffffffffffffff1681565b61015873deaddeaddeaddeaddeaddeaddeaddeaddead000181565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100b5565b3373deaddeaddeaddeaddeaddeaddeaddeaddead000114610224576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603b60248201527f4c31426c6f636b3a206f6e6c7920746865206465706f7369746f72206163636f60448201527f756e742063616e20736574204c3120626c6f636b2076616c7565730000000000606482015260840160405180910390fd5b6000805467ffffffffffffffff9687167fffffffffffffffffffffffffffffffff0000000000000000000000000000000090911617680100000000000000009587169590950294909417909355600191909155600255600380547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001691909216179055565b60606102d47f000000000000000000000000000000000000000000000000000000000000000061034c565b6102fd7f000000000000000000000000000000000000000000000000000000000000000061034c565b6103267f000000000000000000000000000000000000000000000000000000000000000061034c565b6040516020016103389392919061057e565b604051602081830303815290604052905090565b60608160000361038f57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156103b957806103a381610623565b91506103b29050600a8361068a565b9150610393565b60008167ffffffffffffffff8111156103d4576103d461069e565b6040519080825280601f01601f1916602001820160405280156103fe576020820181803683370190505b5090505b8415610481576104136001836106cd565b9150610420600a866106e4565b61042b9060306106f8565b60f81b81838151811061044057610440610710565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535061047a600a8661068a565b9450610402565b949350505050565b803567ffffffffffffffff811681146104a157600080fd5b919050565b600080600080600060a086880312156104be57600080fd5b6104c786610489565b94506104d560208701610489565b935060408601359250606086013591506104f160808701610489565b90509295509295909350565b60005b83811015610518578181015183820152602001610500565b83811115610527576000848401525b50505050565b602081526000825180602084015261054c8160408501602087016104fd565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b600084516105908184602089016104fd565b80830190507f2e0000000000000000000000000000000000000000000000000000000000000080825285516105cc816001850160208a016104fd565b600192019182015283516105e78160028401602088016104fd565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610654576106546105f4565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000826106995761069961065b565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000828210156106df576106df6105f4565b500390565b6000826106f3576106f361065b565b500690565b6000821982111561070b5761070b6105f4565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L1BlockStorageLayoutJSON), L1BlockStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1Block"] = L1BlockStorageLayout
	deployedBytecodes["L1Block"] = L1BlockDeployedBin
}
