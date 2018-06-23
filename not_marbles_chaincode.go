/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

// ====CHAINCODE EXECUTION SAMPLES (CLI) ==================

// ==== Invoke Marbles ====
// peer chaincode invoke -C myc1 -n Marbles -c '{"Args":["initMarble","Marble1","blue","35","tom"]}'
// peer chaincode invoke -C myc1 -n Marbles -c '{"Args":["initMarble","Marble2","red","50","tom"]}'
// peer chaincode invoke -C myc1 -n Marbles -c '{"Args":["initMarble","Marble3","blue","70","tom"]}'
// peer chaincode invoke -C myc1 -n Marbles -c '{"Args":["transferMarble","Marble2","jerry"]}'
// peer chaincode invoke -C myc1 -n Marbles -c '{"Args":["transferMarblesBasedOnColor","blue","jerry"]}'
// peer chaincode invoke -C myc1 -n Marbles -c '{"Args":["delete","Marble1"]}'

// ==== Query Marbles ====
// peer chaincode query -C myc1 -n Marbles -c '{"Args":["readMarble","Marble1"]}'
// peer chaincode query -C myc1 -n Marbles -c '{"Args":["getMarblesByRange","Marble1","Marble3"]}'
// peer chaincode query -C myc1 -n Marbles -c '{"Args":["getHistoryForMarble","Marble1"]}'

// Rich Query (Only supported if CouchDB is used as state database):
//   peer chaincode query -C myc1 -n Marbles -c '{"Args":["queryMarblesByOwner","tom"]}'
//   peer chaincode query -C myc1 -n Marbles -c '{"Args":["queryMarbles","{\"selector\":{\"owner\":\"tom\"}}"]}'

// INDEXES TO SUPPORT COUCHDB RICH QUERIES
//
// Indexes in CouchDB are required in order to make JSON queries efficient and are required for
// any JSON query with a sort. As of Hyperledger Fabric 1.1, indexes may be packaged alongside
// chaincode in a META-INF/statedb/couchdb/indexes directory. Each index must be defined in its own
// text file with extension *.json with the index definition formatted in JSON following the
// CouchDB index JSON syntax as documented at:
// http://docs.couchdb.org/en/2.1.1/api/database/find.html#db-index
//
// This Marbles02 example chaincode demonstrates a packaged
// index which you can find in META-INF/statedb/couchdb/indexes/indexOwner.json.
// For deployment of chaincode to production environments, it is recommended
// to define any indexes alongside chaincode so that the chaincode and supporting indexes
// are deployed automatically as a unit, once the chaincode has been installed on a peer and
// instantiated on a channel. See Hyperledger Fabric documentation for more details.
//
// If you have access to the your peer's CouchDB state database in a development environment,
// you may want to iteratively test various indexes in support of your chaincode queries.  You
// can use the CouchDB Fauxton interface or a command line curl utility to create and update
// indexes. Then once you finalize an index, include the index definition alongside your
// chaincode in the META-INF/statedb/couchdb/indexes directory, for packaging and deployment
// to managed environments.
//
// In the examples below you can find index definitions that support Marbles02
// chaincode queries, along with the syntax that you can use in development environments
// to create the indexes in the CouchDB Fauxton interface or a curl command line utility.
//

//Example hostname:port configurations to access CouchDB.
//
//To access CouchDB docker container from within another docker container or from vagrant environments:
// http://couchdb:5984/
//
//Inside couchdb docker container
// http://127.0.0.1:5984/

// Index for docType, owner.
// Note that docType and owner fields must be prefixed with the "data" wrapper
//
// Index definition for use with Fauxton interface
// {"index":{"fields":["data.docType","data.owner"]},"ddoc":"indexOwnerDoc", "name":"indexOwner","type":"json"}
//
// Example curl command line to define index in the CouchDB channel_chaincode database
// curl -i -X POST -H "Content-Type: application/json" -d "{\"index\":{\"fields\":[\"data.docType\",\"data.owner\"]},\"name\":\"indexOwner\",\"ddoc\":\"indexOwnerDoc\",\"type\":\"json\"}" http://hostname:port/myc1_Marbles/_index
//

// Index for docType, owner, size (descending order).
// Note that docType, owner and size fields must be prefixed with the "data" wrapper
//
// Index definition for use with Fauxton interface
// {"index":{"fields":[{"data.size":"desc"},{"data.docType":"desc"},{"data.owner":"desc"}]},"ddoc":"indexSizeSortDoc", "name":"indexSizeSortDesc","type":"json"}
//
// Example curl command line to define index in the CouchDB channel_chaincode database
// curl -i -X POST -H "Content-Type: application/json" -d "{\"index\":{\"fields\":[{\"data.size\":\"desc\"},{\"data.docType\":\"desc\"},{\"data.owner\":\"desc\"}]},\"ddoc\":\"indexSizeSortDoc\", \"name\":\"indexSizeSortDesc\",\"type\":\"json\"}" http://hostname:port/myc1_Marbles/_index

// Rich Query with index design doc and index name specified (Only supported if CouchDB is used as state database):
//   peer chaincode query -C myc1 -n Marbles -c '{"Args":["queryMarbles","{\"selector\":{\"docType\":\"Marble\",\"owner\":\"tom\"}, \"use_index\":[\"_design/indexOwnerDoc\", \"indexOwner\"]}"]}'

// Rich Query with index design doc specified only (Only supported if CouchDB is used as state database):
//   peer chaincode query -C myc1 -n Marbles -c '{"Args":["queryMarbles","{\"selector\":{\"docType\":{\"$eq\":\"Marble\"},\"owner\":{\"$eq\":\"tom\"},\"size\":{\"$gt\":0}},\"fields\":[\"docType\",\"owner\",\"size\"],\"sort\":[{\"size\":\"desc\"}],\"use_index\":\"_design/indexSizeSortDoc\"}"]}'

package main

import (
		"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

var logger = shim.NewLogger("main logger")


// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init initializes chaincode
// ===========================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Init")

	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
// ========================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	//fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "initMarble" { //create a new Marble
		return t.initMarble(stub, args)
	} else if function == "transferMarble" { //change owner of a specific Marble
		return t.transferMarble(stub, args)
	} else if function == "transferMarblesBasedOnColor" { //transfer all Marbles of a certain color
		return t.transferMarblesBasedOnColor(stub, args)
	} else if function == "delete" { //delete a Marble
		return t.delete(stub, args)
	} else if function == "readMarble" { //read a Marble
		return t.readMarble(stub, args)
	} else if function == "queryMarblesByOwner" { //find Marbles for owner X using rich query
		return t.queryMarblesByOwner(stub, args)
	} else if function == "queryMarbles" { //find Marbles based on an ad hoc rich query
		return t.queryMarbles(stub, args)
	} else if function == "getHistoryForMarble" { //get history of values for a Marble
		return t.getHistoryForMarble(stub, args)
	} else if function == "getMarblesByRange" { //get Marbles based on range query
		return t.getMarblesByRange(stub, args)
	}

	// PO operations
	if function == "AddPOs" {
		return t.initPOs(stub, args)
	}else if function == "RetrievePO" {
		return t.readPO(stub, args)
	}else if function == "UpdatePOs" {
		return t.updatePOs(stub, args)
	}

	// User operations
	if function == "initUser" {
		return t.initUser(stub, args)
	}else if function == "readUser" {
		return t.readUser(stub, args)
	}else if function == "updateUser" {
		return t.updateUser(stub, args)
	}

	// Settlement operations
	if function == "AddSetlements" {
		return t.initSettlements(stub, args)
	}else if function == "RetrieveSettlement" { //read a Marble
		return t.readSettlements(stub, args)
	}else if function == "UpdateSettlement"{
		return t.updateSettlements(stub, args)
	}
	// Accrual operations
	if function == "AddAccruals" {
		return t.initAccruals(stub, args)
	}else if function == "RetrieveAccrual" {
		return t.readAccrual(stub, args)
	}else if function == "UpdateAccrual"{
		return t.updateAccruals(stub, args)
	}

	// Document operations
	if function == "AddDocuments" {
		return t.initDocuments(stub, args)
	}else if function == "RetrieveDocument" { //read a Marble
		return t.readDocument(stub, args)
	}else if function == "UpdateDocument"{
		return t.updateDocuments(stub, args)
	}

	// Invoice operations
	if function == "AddInvoices" {
		return t.addInvoices(stub, args)
	} else if function == "RetrieveInvoice" {
		return t.getInvoice(stub, args)
	} else if function == "UpdateInvoices" {
		return t.updateInvoices(stub, args)
	}


	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

// ==================================================
// delete - remove a Marble key/value pair from state
// ==================================================
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var MarbleJSON Marble
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	MarbleName := args[0]

	// to maintain the color~name index, we need to read the Marble first and get its color
	valAsbytes, err := stub.GetState(MarbleName) //get the Marble from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + MarbleName + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Marble does not exist: " + MarbleName + "\"}"
		return shim.Error(jsonResp)
	}

	err = json.Unmarshal([]byte(valAsbytes), &MarbleJSON)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON of: " + MarbleName + "\"}"
		return shim.Error(jsonResp)
	}

	err = stub.DelState(MarbleName) //remove the Marble from chaincode state
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}

	// maintain the index
	indexName := "color~name"
	colorNameIndexKey, err := stub.CreateCompositeKey(indexName, []string{MarbleJSON.Color, MarbleJSON.Name})
	if err != nil {
		return shim.Error(err.Error())
	}

	//  Delete index entry to state.
	err = stub.DelState(colorNameIndexKey)
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}
	return shim.Success(nil)
}