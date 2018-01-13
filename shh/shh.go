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
 * @file ssh.go
 * @authors:
 *   Reginaldo Costa <regcostajr@gmail.com>
 * @date 2017
 */

package shh

import (
	"github.com/alexlitz/go-web3/complex/types"
	"github.com/alexlitz/go-web3/dto"
	"github.com/alexlitz/go-web3/providers"
)

// SHH - The Net Module
type SHH struct {
	provider providers.ProviderInterface
}

// NewSHH - Net Module constructor to set the default provider
func NewSHH(provider providers.ProviderInterface) *SHH {
	shh := new(SHH)
	shh.provider = provider
	return shh
}

// GetVersion - Returns the current whisper protocol version.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#shh_version
// Parameters:
//    - none
// Returns:
//    - String - The current whisper protocol version
func (shh *SHH) GetVersion() (string, error) {

	pointer := &dto.RequestResult{}

	err := shh.provider.SendRequest(pointer, "shh_version", nil)

	if err != nil {
		return "", err
	}

	return pointer.ToString()

}

// Post - Sends a whisper message.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#shh_post
// Parameters:
//     1 .Object - The whisper post object:
//	  	- from: DATA, 60 Bytes - (optional) The identity of the sender.
//    	- to: DATA, 60 Bytes - (optional) The identity of the receiver. When present whisper will encrypt the message so that only the receiver can decrypt it.
//   	- topics: Array of DATA - Array of DATA topics, for the receiver to identify messages.
//    	- payload: DATA - The payload of the message.
//    	- priority: QUANTITY - The integer of the priority in a rang from ... (?).
//    	- ttl: QUANTITY - integer of the time to live in seconds.
// Returns:
// 	  - Boolean - returns true if the message was sent, otherwise false.
func (shh *SHH) Post(from string, to string, topics []string, payload string, priority types.ComplexIntParameter, ttl types.ComplexIntParameter) (bool, error) {

	params := make([]dto.WhisperMessage, 1)
	params[0].From = from
	params[0].To = to
	params[0].Topics = topics
	params[0].Payload = payload
	params[0].Priority = priority.ToHex()
	params[0].TTL = ttl.ToHex()

	pointer := &dto.RequestResult{}

	err := shh.provider.SendRequest(pointer, "shh_post", params)

	if err != nil {
		return false, err
	}

	return pointer.ToBoolean()

}

// NewIdentity - Creates a whisper identity in the client.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#shh_newidentity
// Parameters:
//      - None
// Returns:
// 	- DATA, 60 Bytes - the address of the new identiy.
func (ssh *SHH) NewIdentity() (string, error) {

        pointer := &dto.RequestResult{}

        err := shh.provider.SendRequest(pointer, "shh_newIdentity", nil)

        if err != nil {
		return "", err
	}

        return pointer.ToString()
}

// HasIdentity - Creates a whisper identity in the client.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#shh_hasidentity
// Parameters:
//       - None
// Returns:
// 	  - Boolean - returns true if the client holds the privatekey for that identity, otherwise false
func (ssh *SHH) NewIdentity() (bool, error) {

        pointer := &dto.RequestResult{}

        err := shh.provider.SendRequest(pointer, "shh_hasIdentity", nil)

        if err != nil {
                return false, err
        }

        return pointer.ToBoolean()
}

// NewGroup - Creates a group
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#shh_newgroup
// Parameters
//       - None
// Returns:
//       - DATA, 60 Bytes - the address of the new group.
func (ssh *SHH) NewGroup() (string, error) {

        pointer := &dto.RequestResult{}

        err := shh.provider.SendRequest(pointer, "shh_newGroup", nil)

        if err != nil {
		return "", err
	}

        return pointer.ToString()
}

// AddToGroup - 
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#shh_addtogroup
// Parameters
//       - Identity: DATA, 60 Bytes - The identity address to add to a group (?).
// Returns:
//       - Boolean - returns true if the identity was successfully added to the group, otherwise false
func (ssh *SHH) AddToGroup(id string) (bool, error) {

        params := [1]string{id}

        pointer := &dto.RequestResult{}

        err := shh.provider.SendRequest(pointer, "shh_addToGroup", params)

        if err != nil {
                return false, err
        }

        return pointer.ToBoolean()
}


// NewFilter - Creates filter to notify, when client receives whisper message matching the filter options. 
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#shh_newfilter
// Parameters
//       - Object - The filter options:
//           to: DATA, 60 Bytes - (optional) Identity of the receiver. When present it 
//              will try to decrypt any incoming message if the client holds the private key to this identity.
//           topics: Array of DATA - Array of DATA topics which the incoming message's topics should match. 
//              You can use the following combinations:
//              [A, B] = A && B
//              [A, [B, C]] = A && (B || C)
//              [null, A, B] = ANYTHING && A && B null works as a wildcard
// Returns:
//       - QUANTITY - The newly created filter.
func (ssh *SHH) NewFilter(to string, topics []string) {

        params := make(dto.SHHFilterParameters, 1)
        params[0].Topics = topics
        params[0].To = to

        pointer := &dto.RequestResult{}

        err := shh.provider.SendRequest(pointer, "shh_newFilter", params)

        if err != nil {
		return "", err
	}

        return pointer.ToString()
}

// UninstallFilter - Uninstalls a filter with given id. Should always be called when watch is no longer needed. 
//      Additonally Filters timeout when they aren't requested with shh_getFilterChanges for a period of time.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#shh_uninstallfilter
// Parameters
//       - QUANTITY - The filter id.
// Returns:
//       - Boolean - true if the filter was successfully uninstalled, otherwise false.
func (ssh *SHH) UninstallFilter(id string) (bool, filter) {

        params := [1]string{id}

        pointer := &dto.RequestResult{}

        err := shh.provider.SendRequest(pointer, "shh_uninstallFilter", params)

        if err != nil {
                return false, err
        }

        return pointer.ToBoolean()
}

// GetFilterChanges - Polling method for whisper filters. Returns new messages since the last call of this method.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#shh_getfilterchanges 
// Parameters
//       - QUANTITY - The filter id.
// Returns:
//       - Array - Array of messages received since last poll:
//              hash: DATA, 32 Bytes (?) - The hash of the message.
//              from: DATA, 60 Bytes - The sender of the message, if a sender was specified.
//              to: DATA, 60 Bytes - The receiver of the message, if a receiver was specified.
//              expiry: QUANTITY - Integer of the time in seconds when this message should expire (?).
//              ttl: QUANTITY - Integer of the time the message should float in the system in seconds (?).
//              sent: QUANTITY - Integer of the unix timestamp when the message was sent.
//              topics: Array of DATA - Array of DATA topics the message contained.
//              payload: DATA - The payload of the message.
//              workProved: QUANTITY - Integer of the work this message required before it was sent 
func (ssh *SHH) GetFilterChanges(id string) []dto.WhisperMessage {

        params := [1]string{id}

        pointer := &dto.RequestResult{}

        err := shh.provider.SendRequest(pointer, "shh_getFilterChanges", params)

        if err != nil {
                return nil, err
        }

        return pointer.ToMessageArray()
}

// GetMessages - Get all messages matching a filter. Unlike shh_getFilterChanges this returns all messages.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#shh_getmessages
// Parameters
//       - QUANTITY - The filter id
// Returns:
//       - Array - Array of messages received since last poll:
//              hash: DATA, 32 Bytes (?) - The hash of the message.
//              from: DATA, 60 Bytes - The sender of the message, if a sender was specified.
//              to: DATA, 60 Bytes - The receiver of the message, if a receiver was specified.
//              expiry: QUANTITY - Integer of the time in seconds when this message should expire (?).
//              ttl: QUANTITY - Integer of the time the message should float in the system in seconds (?).
//              sent: QUANTITY - Integer of the unix timestamp when the message was sent.
//              topics: Array of DATA - Array of DATA topics the message contained.
//              payload: DATA - The payload of the message.
//              workProved: QUANTITY - Integer of the work this message required before it was sent 
func (ssh *SHH) GetMessages(id string) []dto.WhisperMessage {

        params := [1]string{id}

        pointer := &dto.RequestResult{}

        err := shh.provider.SendRequest(pointer, "shh_getMessages", params)

        if err != nil {
                return nil, err
        }

        return pointer.ToMessageArray()
}
