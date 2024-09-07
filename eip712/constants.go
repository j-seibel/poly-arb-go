package eip712

import (
	"poly/arb/abi"

	"github.com/ethereum/go-ethereum/crypto"
)

var (
	_EIP712_DOMAIN = []abi.Type{
		Bytes32, // typehash
		Bytes32, // name
		Bytes32, // version
		Uint256, // chainId
		Address, // verifyingContract
	}

	_EIP712_DOMAIN_HASH = crypto.Keccak256Hash(
		[]byte("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"),
	)

	_EIP712_DOMAIN_NO_VERIFYING_CONTRACT = []abi.Type{
		Bytes32, // typehash
		Bytes32, // name
		Bytes32, // version
		Uint256, // chainId
	}

	_EIP712_DOMAIN_HASH_NO_VERIFYING_CONTRACT = crypto.Keccak256Hash(
		[]byte("EIP712Domain(string name,string version,uint256 chainId)"),
	)
)

var (
	_ORDER_STRUCTURE = []abi.Type{
		Bytes32, // typehash
		Uint256, // salt
		Address, // maker
		Address, // signer
		Address, // taker
		Uint256, // tokenId
		Uint256, // makerAmount
		Uint256, // takerAmount
		Uint256, // expiration
		Uint256, // nonce
		Uint256, // feeRateBps
		Uint8,   // side
		Uint8,   // signatureType
	}
)

var (
	_ORDER_STRUCTURE_HASH = crypto.Keccak256Hash(
		[]byte("Order(uint256 salt,address maker,address signer,address taker,uint256 tokenId,uint256 makerAmount,uint256 takerAmount,uint256 expiration,uint256 nonce,uint256 feeRateBps,uint8 side,uint8 signatureType)"),
	)
)
