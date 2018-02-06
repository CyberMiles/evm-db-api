package evmdbapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func SendRequest(method string, params interface{}) (*RPCResponse, error) {
	// body
	objReq := RPCRequest{JSONRPC: "2.0", ID: 1, Params: params, Method: method}
	data, err := json.Marshal(objReq)
	if err != nil {
		return nil, err
	}
	body := strings.NewReader(string(data))

	// url
	prefix := "http://"
	if MyProvider.secure {
		prefix = "https://"
	}
	url := prefix+MyProvider.address

	// send request
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	client := &http.Client{ Timeout: time.Second * time.Duration(MyProvider.timeout) }
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// parse response
	var bodyResp []byte
	var objResp RPCResponse
	if resp.StatusCode == http.StatusOK {
		bodyResp, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
	}
	err = json.Unmarshal(bodyResp, &objResp)
	if err != nil {
		return nil, err
	}
	return &objResp, nil
}
