package client

import (
	"poly/arb/eip712"

	"poly/arb/abi"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const HOST = "https://clob.polymarket.com"
const CHAIN_ID = 137
const PK = "3273d17583f924c65451317e658532dc62a9cc52505183eb4b65f121660e8ed1"
const FUNDER = "0x6cd02aAfEEb049150014D3D9356613897Ce54e6C"

var NEG_RISK_CONTRACT_CONFIG = &ContractConfig{
	"0xd91E80cF2E7be2e162c6513ceD06f1dD0dA35296",
	"0x9c4e1703476e875070ee25b56a58b008cfb8fa78",
	"0x69308FB512518e39F9b16112fA8d994F4e2Bf8bB",
}

const MSG_TO_SIGN = "This message attests that I control the given wallet"

var CLOB_REQUEST_STRUCTURE = []abi.Type{
	eip712.Bytes32,
	eip712.Address,
	eip712.String,
	eip712.Uint256,
	eip712.String,
}

var CLOB_ENCODE_HASH = crypto.Keccak256Hash([]byte("ClobAuth(address address,string timestamp,uint256 nonce,string message)"))

var API_CREDS = APICreds{
	"2cade072-7d71-7480-0ed0-418f0de112fa",
	"Q4qyjHxWpCwNsPKjfnsQdUE0PFGRhlVcXH-4RE4CUYk=",
	"f64fe106f25ccb2ec95c46e4dbb0e1d7cdb442cb54f6df4a4ac6d542bda155ca",
}

var POLYMARKET_SIGNER = Signer{
	PK,
	common.HexToAddress("0xA3D381B8C135cEd27efbbd3f231a0E1B6B931ad0"),
	137,
}
