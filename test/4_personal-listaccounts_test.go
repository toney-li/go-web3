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
 * @file personal-listaccounts_test.go
 * @authors:
 *   Reginaldo Costa <regcostajr@gmail.com>
 * @date 2017
 */
package test

import (
	"testing"

	"github.com/regcostajr/go-web3"
)

func Personal_ListAccounts(connection *web3.Web3) error {
	_, err := connection.Personal.ListAccounts()

	if err != nil {
		return err
	}

	return nil
}

func TestPersonal_ListAccounts_IPCConnection(t *testing.T) {
	err := Personal_ListAccounts(IPCConnection())
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestPersonal_ListAccounts_HTTPConnection(t *testing.T) {
	err := Personal_ListAccounts(HTTPConnection())
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
