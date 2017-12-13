package test

import (
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	web3 "github.com/regcostajr/go-web3"
	"github.com/regcostajr/go-web3/complex/types"
)

func Eth_GetBlockByNumber(connection *web3.Web3) error {
	//previous block Dec-09-2017 10:28:29 AM +UTC
	//Dec-09-2017 10:28:31 AM +UTC should be block 4701754, hash 0xaec14e98d578351a23d5ab40f65ee431063f582b5d736bbc0705a2eef0fb8f9d
	//next block Dec-09-2017 10:28:55 AM +UTC
	expectedBlockHash := "0xaec14e98d578351a23d5ab40f65ee431063f582b5d736bbc0705a2eef0fb8f9d"
	var blockNumber types.ComplexIntParameter = 436

	block, err := connection.Eth.GetBlockByNumber(blockNumber, false)

	if err != nil {
		return err
	}
	if block == nil {
		return errors.New("Block returned is nil")
	}

	actualBlockDate := time.Unix(block.Timestamp.ToInt64(), 0)
	expectedBlockDate := time.Date(2017, 12, 9, 10, 28, 31, 0, time.UTC)

	if strings.Compare(block.Hash, expectedBlockHash) != 0 {
		return fmt.Errorf("Expected block hash %v, got %v", expectedBlockHash, block.Hash)
	}
	if block.Number.ToInt64() != int64(blockNumber) {
		return fmt.Errorf("Expected block number %v, got %v", blockNumber, block.Number)
	}
	if actualBlockDate.Equal(expectedBlockDate) {
		return fmt.Errorf("Expected block date %v, got %v", expectedBlockDate, actualBlockDate)
	}

	return nil
}

func TestGetBlockByNumber_IPCConnection(t *testing.T) {
	err := Eth_GetBlockByNumber(IPCConnection())
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestGetBlockByNumber_HTTPConnection(t *testing.T) {
	err := Eth_GetBlockByNumber(HTTPConnection())
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
