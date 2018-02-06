package evmdbapi

import (
	"fmt"
	"strconv"
	"math/big"
	"encoding/hex"
)

func EthGetBalance(address string, blockNumberOrTag string) (*big.Int, error) {
	resp, err := SendRequest("eth_getBalance", []string{address, blockNumberOrTag})
	if err != nil {
		return big.NewInt(0), err
	}
	if resp.Error != nil {
		return big.NewInt(0), fmt.Errorf(resp.Error.Message)
	}
	return ParseBigInt(resp.Result.(string)), nil
}

func EthCall(tx *TransactionCall, blockNumberOrTag string) (string, error) {
	resp, err := SendRequest("eth_call", []interface{}{tx, blockNumberOrTag})
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}

func EthGetTokenBalance(contract string, account string, blockNumberOrTag string) (*TokenBalance, error) {
	// balanceOf
	data := fmt.Sprintf("0x%s000000000000000000000000%s", 
		MethodHash("balanceOf(address)"), account[2:])
	tx := TransactionCall{To: contract, Data: data}

	resp, err := EthCall(&tx, blockNumberOrTag)
	if err != nil {
    return nil, err
	}
	balance := ParseBigInt(resp)

	// decimals
	data = fmt.Sprintf("0x%s", MethodHash("decimals()"))
	tx.Data = data

	resp, err = EthCall(&tx, blockNumberOrTag)
	if err != nil {
    return nil, err
	}
	decimals, err := strconv.ParseInt(resp, 0, 64)
	if err != nil {
    return nil, err
	}

	// symbol
	data = fmt.Sprintf("0x%s", MethodHash("symbol()"))
	tx.Data = data

	resp, err = EthCall(&tx, blockNumberOrTag)
	if err != nil {
    return nil, err
	}
  symbol, err := hex.DecodeString(resp[2:])
  if err != nil {
    return nil, err
  }

  ret := TokenBalance{Balance: balance, Decimals: decimals, Symbol: string(symbol)}
	return &ret, nil
}
