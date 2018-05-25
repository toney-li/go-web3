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
 * @file shh-addToGroup_test.go
 * @authors:
 *   Alex Litzenberger
 * @date 2018
 */

package test

import (
	"testing"

	shh "github.com/alexlitz/go-web3/shh"
	"github.com/alexlitz/go-web3/providers"
)

func TestSHHAddToGroup(t *testing.T) {

	var instance = shh.NewSHH(providers.NewHTTPProvider("127.0.0.1:8545", 10, false))

	to := "0x04f96a5e25610293e42a73908e93ccc8c4d4dc0edcfa9fa872f50cb214e08ebf61a03e245533f97284d442460f2998cd41858798ddfd4d661997d3940272b717b1"

	_, err := instance.AddToGroup(to)

	if err != nil {
       	       t.Error(err)
	       t.FailNow()
	}
}