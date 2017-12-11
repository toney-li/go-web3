package test

import (
	"strings"
	"testing"
	"time"
)

func TestGetBlockByNumber_IPCConnection(t *testing.T) {
	err := Eth_GetBalance(IPCConnection())

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestGetBlockByNumber_HTTPConnection(t *testing.T) {
	//previous block Dec-09-2017 10:28:29 AM +UTC
	//Dec-09-2017 10:28:31 AM +UTC should be block 4701754, hash 0xaec14e98d578351a23d5ab40f65ee431063f582b5d736bbc0705a2eef0fb8f9d
	//next block Dec-09-2017 10:28:55 AM +UTC
	expectedBlockHash := "0xaec14e98d578351a23d5ab40f65ee431063f582b5d736bbc0705a2eef0fb8f9d"
	var blockNumber int64 = 4701754

	block, err := HTTPConnection().Eth.GetBlockByNumber(blockNumber, false)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if block == nil {
		t.Error("Block returned is nil")
		t.Fail()
	}

	actualBlockDate := time.Unix(block.Timestamp.ToInt64(), 0)
	expectedBlockDate := time.Date(2017, 12, 9, 10, 28, 31, 0, time.UTC)

	if strings.Compare(block.Hash, expectedBlockHash) != 0 {
		t.Errorf("Expected block hash %v, got %v", expectedBlockHash, block.Hash)
	}
	if block.Number.ToInt64() != blockNumber {
		t.Errorf("Expected block number %v, got %v", blockNumber, block.Number)
	}
	if actualBlockDate.Equal(expectedBlockDate) {
		t.Errorf("Expected block date %v, got %v", expectedBlockDate, actualBlockDate)
	}
}
