package client

import (
	"poly/arb/eip712"

	"poly/arb/abi"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const HOST = "https://clob.polymarket.com"
const CHAIN_ID = 137
const PK = os.Env("PRIVATE_KEY")
const FUNDER = os.ENV("FUNDER")

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
	os.Env("API_KEY"),
	os.Env("API_SECRET"),
	os.ENV("FUNDER"),
}

var POLYMARKET_SIGNER = Signer{
	PK,
	common.HexToAddress("0xA3D381B8C135cEd27efbbd3f231a0E1B6B931ad0"),
	137,
}
