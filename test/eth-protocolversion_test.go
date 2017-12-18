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
 * @file eth-protocolversion_test.go
 * @authors:
 *   Reginaldo Costa <regcostajr@gmail.com>
 * @date 2017
 */

package test

import (
	"fmt"
	"testing"

	web3 "github.com/regcostajr/go-web3"
)

func Eth_ProtocolVersion(connection *web3.Web3) error {
	version, err := connection.Eth.GetProtocolVersion()

	if err != nil {
		return err
	}

	fmt.Println(version)

	return nil
}

func TestGetProtocolVersion_IPCConnection(t *testing.T) {
	err := Eth_ProtocolVersion(IPCConnection())

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestGetProtocolVersion_HTTPConnection(t *testing.T) {
	err := Eth_ProtocolVersion(HTTPConnection())

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
