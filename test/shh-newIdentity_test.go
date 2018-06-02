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
 * @file shh-newIdentity_test.go
 * @authors:
 *   Alex Litzenberger
 * @date 2018
 */

package test

import (
	"errors"
	"fmt"
	"testing"

	hex "encoding/hex"
	shh "github.com/alexlitz/go-web3/shh"
	"github.com/alexlitz/go-web3/providers"
)

func TestSHHNewIdentity(t *testing.T) {

	var instance = shh.NewSHH(providers.NewHTTPProvider("127.0.0.1:8545", 10, false))

	id, err := instance.NewIdentity()

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fmt.Println(id)
	
	if _, err := hex.DecodeString(id); err != nil {
		t.Error(errors.New("Invalid shh id"))
		t.Fail()
	}
}
