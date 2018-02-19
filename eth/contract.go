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
 * @file contract.go
 * @authors:
 *   Reginaldo Costa <regcostajr@gmail.com>
 * @date 2018
 */

package eth

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/regcostajr/go-web3/complex/types"
	"github.com/regcostajr/go-web3/dto"

	"github.com/regcostajr/go-web3/utils"
)

// Contract ...
type Contract struct {
	super     *Eth
	abi       string
	functions map[string][]string
}

// NewContract - Contract abstraction
func (eth *Eth) NewContract(abi string) (*Contract, error) {

	contract := new(Contract)
	var mockInterface interface{}

	err := json.Unmarshal([]byte(abi), &mockInterface)

	if err != nil {
		return nil, err
	}

	jsonInterface := mockInterface.([]interface{})
	contract.functions = make(map[string][]string)
	for index := 0; index < len(jsonInterface); index++ {
		function := jsonInterface[index].(map[string]interface{})

		if function["name"] == nil {
			continue
		}

		functionName := function["name"].(string)
		contract.functions[functionName] = make([]string, 0)

		if function["inputs"] == nil {
			continue
		}

		inputs := function["inputs"].([]interface{})
		for paramIndex := 0; paramIndex < len(inputs); paramIndex++ {
			params := inputs[paramIndex].(map[string]interface{})
			contract.functions[functionName] = append(contract.functions[functionName], params["type"].(string))
		}
	}

	contract.abi = abi
	contract.super = eth

	return contract, nil
}

// prepareTransaction ...
func (contract *Contract) prepareTransaction(transaction *dto.TransactionParameters, functionName string, args ...interface{}) (*dto.TransactionParameters, error) {

	function, ok := contract.functions[functionName]
	if !ok {
		return nil, errors.New("Function not finded on passed abi")
	}

	fullFunction := functionName + "("

	comma := ""
	for arg := range function {
		fullFunction += comma + function[arg]
		comma = ","
	}

	fullFunction += ")"

	utils := utils.NewUtils(contract.super.provider)
	sha3Function, err := utils.Sha3(types.ComplexString(fullFunction))

	if err != nil {
		return nil, err
	}

	var data string

	for index := 0; index < len(function); index++ {
		switch argType := function[index]; argType {
		case "address":
			data += args[index].(string)
		default:
			data += fmt.Sprintf("0x%x", args[index].(string))
		}
	}

	transaction.Data = types.ComplexString(sha3Function[0:10] + data)

	return transaction, nil

}

func (contract *Contract) Call(transaction *dto.TransactionParameters, functionName string, args ...interface{}) (string, error) {

	transaction, err := contract.prepareTransaction(transaction, functionName, args)

	if err != nil {
		return "", err
	}

	return contract.super.Call(transaction)

}

func (contract *Contract) Send(transaction *dto.TransactionParameters, functionName string, args ...interface{}) (string, error) {

	transaction, err := contract.prepareTransaction(transaction, functionName, args)

	if err != nil {
		return "", err
	}

	return contract.super.SendTransaction(transaction)

}
