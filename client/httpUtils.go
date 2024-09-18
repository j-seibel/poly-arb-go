package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func GetWithL0Headers(endpoint string, jsonData []byte) ([]byte, error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Set the headers
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header["User-Agent"] = []string{"py_clob_client"}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	return body, nil
}

func PostWithL0Headers(endpoint string, jsonData []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Set the headers
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header["User-Agent"] = []string{"py_clob_client"}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	return body, nil
}

func PostWithL1Headers(endpoint string, headers L1Headers, jsonData []byte) ([]byte, error) {
	// Create a new POST request
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Set the headers
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header["POLY_ADDRESS"] = []string{headers.POLY_ADDRESS}
	req.Header["POLY_SIGNATURE"] = []string{headers.POLY_SIGNATURE}
	req.Header["POLY_TIMESTAMP"] = []string{headers.POLY_TIMESTAMP}
	req.Header["POLY_NONCE"] = []string{headers.POLY_NONCE}
	// req.Header["Accept-Encoding"] = []string{"gzip, deflate"}
	req.Header["User-Agent"] = []string{"py_clob_client"}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	return body, nil
}

func GetWithL1Headers(endpoint string, headers L1Headers, jsonData []byte) ([]byte, error) {
	// Create a new POST request
	req, err := http.NewRequest("GET", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Set the headers
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header["POLY_ADDRESS"] = []string{headers.POLY_ADDRESS}
	req.Header["POLY_SIGNATURE"] = []string{headers.POLY_SIGNATURE}
	req.Header["POLY_TIMESTAMP"] = []string{headers.POLY_TIMESTAMP}
	req.Header["POLY_NONCE"] = []string{headers.POLY_NONCE}
	// req.Header["Accept-Encoding"] = []string{"gzip, deflate"}
	req.Header["User-Agent"] = []string{"py_clob_client"}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	return body, nil

}

func PostWithL2Headers(endpoint string, headers L2Headers, jsonData []byte) ([]byte, error) {
	// Create a new POST request
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Set the headers
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header["POLY_ADDRESS"] = []string{headers.POLY_ADDRESS}
	req.Header["POLY_SIGNATURE"] = []string{headers.POLY_SIGNATURE}
	req.Header["POLY_TIMESTAMP"] = []string{headers.POLY_TIMESTAMP}
	req.Header["POLY_API_KEY"] = []string{headers.POLY_API_KEY}
	req.Header["POLY_PASSPHRASE"] = []string{headers.POLY_PASSPHRASE}
	// req.Header["Accept-Encoding"] = []string{"gzip, deflate"}
	req.Header["User-Agent"] = []string{"py_clob_client"}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error sending request: ", err)
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	fmt.Println("response body: ", string(body))
	// print(string(body))
	return body, nil
}

func GetWithL2Headers(endpoint string, headers L2Headers, jsonData []byte) ([]byte, error) {
	req, err := http.NewRequest("GET", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Set the headers
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header["POLY_ADDRESS"] = []string{headers.POLY_ADDRESS}
	req.Header["POLY_SIGNATURE"] = []string{headers.POLY_SIGNATURE}
	req.Header["POLY_TIMESTAMP"] = []string{headers.POLY_TIMESTAMP}
	req.Header["POLY_API_KEY"] = []string{headers.POLY_API_KEY}
	req.Header["POLY_PASSPHRASE"] = []string{headers.POLY_PASSPHRASE}
	// req.Header["Accept-Encoding"] = []string{"gzip, deflate"}
	req.Header["User-Agent"] = []string{"py_clob_client"}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	return body, nil
}
