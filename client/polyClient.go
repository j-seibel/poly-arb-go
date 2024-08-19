package client

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"math/big"
	"strconv"

	"fmt"
	"time"

	"poly/arb/eip712"
	"poly/arb/signer"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func CreateLevel2Headers(requestArgs RequestArgs) L2Headers {
	timestamp := time.Now().Unix()
	// timestamp := int64(1724036502)
	timestampString := fmt.Sprintf("%d", timestamp)
	// fmt.Printf("HMACSignature: %v\n", buildHMACSignature(API_CREDS.api_secret, timestampString, requestArgs.method, requestArgs.requestPath, requestArgs.body))
	return L2Headers{
		POLYMARKET_SIGNER.account.String(),
		buildHMACSignature(API_CREDS.api_secret, timestampString, requestArgs.method, requestArgs.requestPath, requestArgs.body),
		timestampString,
		API_CREDS.api_key,
		API_CREDS.api_passphrase,
	}
}

func CreateLevel1Headers(nonce int64) L1Headers {
	timestamp := big.NewInt(time.Now().Unix())
	// timestamp := big.NewInt(1724005368)
	timestampString := fmt.Sprintf("%d", timestamp)

	signature, err := signClobAuthMessage(POLYMARKET_SIGNER.private_key, *timestamp, *big.NewInt(nonce))

	if err != nil {
		fmt.Println("Error signing CLOBAuth message")
	}
	return L1Headers{
		POLYMARKET_SIGNER.account.String(),
		"0x" + hex.EncodeToString(signature),
		timestampString,
		strconv.Itoa(int(nonce)),
	}
}
func generateRandomNonce() (int64, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return 0, err
	}
	return n.Int64(), nil
}

func GenerateUniqueNonce() int64 {
	timestamp := time.Now().UnixNano()
	randValue, _ := generateRandomNonce()
	return timestamp + randValue
}

func buildHMACSignature(secret, timestamp, method, requestPath string, body *string) string {
	// Decode the base64 secret
	base64Secret, err := base64.URLEncoding.DecodeString(secret)
	fmt.Printf("Secret: %v\n", base64Secret)
	if err != nil {
		return ""
	}

	// Build the message to be signed
	message := timestamp + method + requestPath
	// message := "1724036207GET/notifications"
	fmt.Printf("Message: %v\n", message)
	if body != nil {

		message += *body
	}

	// Create HMAC using SHA256
	h := hmac.New(sha256.New, base64Secret)
	h.Write([]byte(message))

	// Encode the signature in base64
	signature := base64.URLEncoding.EncodeToString(h.Sum(nil))

	return signature
}

const (
	ClobDomainName = "ClobAuthDomain"
	ClobVersion    = "1"
	MsgToSign      = "This message attests that I control the given wallet"
)

func signClobAuthMessage(privKey string, timestamp, nonce big.Int) ([]byte, error) {
	name := crypto.Keccak256Hash([]byte(ClobDomainName))
	version := crypto.Keccak256Hash([]byte(ClobVersion))
	chainId := big.NewInt(137)
	address := common.HexToAddress(POLYMARKET_SIGNER.account.Hex())
	domainSeperator, err := eip712.BuildEIP712DomainSeparatorNoContract(name, version, chainId)

	fmt.Printf("Domain Seperator: %v\n", domainSeperator)
	if err != nil {
		return []byte{}, err
	}

	fmt.Printf("timestamp: %s\n", timestamp.String())
	values := []interface{}{
		CLOB_ENCODE_HASH,
		address,
		timestamp.String(),
		&nonce,
		MsgToSign,
	}

	fmt.Printf("Values: %v\n", values)

	hash, err := eip712.HashTypedDataV4(domainSeperator, CLOB_REQUEST_STRUCTURE, values)

	if err != nil {
		fmt.Println("Error hashing typed data", err)
		return []byte{}, err
	}

	keyBytes, err := hex.DecodeString(privKey)
	if err != nil {
		//empty byte array
		return []byte{}, err
	}

	// Create the private key from the byte slice
	PK, _ := crypto.ToECDSA(keyBytes)
	return signer.Sign(PK, hash)
}
