package test

import (
	"testing"

	web3 "github.com/regcostajr/go-web3"
	"github.com/regcostajr/go-web3/eth/block"
)

func Eth_GetBalance(connection *web3.Web3) error {
	accounts, err := ListAccounts(connection)

	if err != nil {
		return err
	}

	_, err = connection.Eth.GetBalance(accounts[0], block.LATEST)

	if err != nil {
		return err
	}

	return nil
}

func TestGetBalance_IPCConnection(t *testing.T) {
	err := Eth_GetBalance(IPCConnection())

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestGetBalance_HTTPConnection(t *testing.T) {
	err := Eth_GetBalance(HTTPConnection())

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
