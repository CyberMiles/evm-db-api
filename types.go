package evmdbapi

import (
	"math/big"
)

type RPCRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Params  interface{} `json:"params,omitempty"`
	Method  string      `json:"method,omitempty"`
}

type RPCResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Error   *Error  		`json:"error,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type TransactionCall struct {
	From     string 		`json:"from,omitempty"`
	To       string 		`json:"to,omitempty"`
	Gas      string 		`json:"gas,omitempty"`
	GasPrice string 		`json:"gasPrice,omitempty"`
	Value    string 		`json:"value,omitempty"`
	Data     string 		`json:"data,omitempty"`
}

type TokenBalance struct {
	Balance 	*big.Int 	`json:"balance"`
	Decimals	int64 		`json:"decimals"`
	Symbol 		string		`json:"symbol"`
}
