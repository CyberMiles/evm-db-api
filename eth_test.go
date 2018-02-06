package evmdbapi

import (
	"testing"
)

const (
	AccountAddress = "0x103AEB8261a57bbe22134ecBC4a107cd9eFe68AD"
	ContractAddress = "0xca5294b2064d92edb07ecbfc9cc77718ee5eace1"
	TransactionHash = "0x4b714230cb20d33db66c21e6b8cdf72d73531893dcbb7147fd2d8ac5760999f3"
	BlockHash = "0xde6584ba4f8d5be98bde3542f0d5dea2e701caf534914e01f0797cdf7af2bded"
)

func TestEthGetBalance(t *testing.T) {
	balance, err := EthGetBalance(AccountAddress, LATEST)
	if err != nil {
		t.Errorf("Error - %v", err)
	}
	t.Logf("Balanace - %v", balance)
}

func TestGetTokenBalance(t *testing.T) {
	token, err := EthGetTokenBalance(ContractAddress, AccountAddress, LATEST)
	if err != nil {
		t.Errorf("Error - %v", err)
	}
	t.Logf("Balanace - %v%v", token.Balance, token.Symbol)
	t.Logf("Decimals - %v", token.Decimals)
}
