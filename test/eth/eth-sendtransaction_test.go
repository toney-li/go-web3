/********************************************************************************
   This file is part of go-web3.
   go-web3 is free software: you can redistribute it and/or modify
   it under the terms of the GNU Lesser General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   go-web3 is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Lesser General Public License for more details.
   You should have received a copy of the GNU Lesser General Public License
   along with go-web3.  If not, see <http://www.gnu.org/licenses/>.
*********************************************************************************/

/**
 * @file eth-sendtransaction_test.go
 * @authors:
 *   Reginaldo Costa <regcostajr@gmail.com>
 * @date 2017
 */
package test

import (
	"testing"

	"github.com/toney-li/go-web3"
	"github.com/toney-li/go-web3/complex/types"
	"github.com/toney-li/go-web3/dto"
	"github.com/toney-li/go-web3/providers"
	"math/big"
)

func TestEthSendTransaction(t *testing.T) {

	var connection = web3.NewWeb3(providers.NewHTTPProvider("127.0.0.1:8547", 10, false))

	_, err := connection.Eth.GetCoinbase()

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	transaction := new(dto.TransactionParameters)
	transaction.From = "0x5dd179db112df1b88b22bfeaade4818502066ee0"
	transaction.To = "0x8bab64f25be09973cd1bcdf90650eb49aa85e825"
	transaction.Value = big.NewInt(0).Mul(big.NewInt(1), big.NewInt(1))
	transaction.Gas = big.NewInt(40000)
	transaction.Data = types.ComplexString("p2p transaction")

	txID, err := connection.Eth.SendTransaction(transaction)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(txID)

}
