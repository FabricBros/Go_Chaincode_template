package main

import (
		"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// For storing arbitrary documents.
type Document struct {
	ObjectType string `json:"docType"`
	Uuid 	   string	`json:"uuid"`
	Data		string `json:"data"`
}

func NewDocument( uuid,data string ) *Document {
	return &Document{
		ObjectType: "document",
		Uuid: uuid,
		Data: data,
	}
}

//var ObjectType="document"


// ============================================================
// initDocument - creates a new document and stores it in the chaincode state
// ============================================================
func (t *SimpleChaincode) initDocuments(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("adding Document")
	defer logger.Debug("exit adding Document")

	var items []Document

	err := json.Unmarshal([]byte(args[0]), &items)
	if err != nil {
		logger.Error("Error unmarshing invoice json:", err)
		return shim.Error(err.Error())
	}

	for _, v := range items{
		v.ObjectType="document"
		pk := v.Uuid
		vBytes, err := json.Marshal(v)

		if err!= nil{
			logger.Debug("error marshaling", err)
			return shim.Error(err.Error())
		}
		stub.PutState(pk, vBytes)
	}

	return shim.Success(nil)
}


func (t *SimpleChaincode) readDocument(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var name, jsonResp string
	var err error
	logger.Debug("Enter readDocument")
	defer logger.Debug("Exited readDocument")

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the Document to query")
	}

	name = args[0]
	valAsbytes, err := stub.GetState(name) //get the Document from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + name + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Marble does not exist: " + name + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}



// query callback representing the query of a chaincode
func (t *SimpleChaincode) updateDocuments(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("Enter updateDocuments")
	defer logger.Debug("Exited updateDocuments")
	var items []Document

	err := json.Unmarshal([]byte(args[0]), &items)
	if err != nil {
		logger.Error("Error unmarshing documents json:", err)
		return shim.Error(err.Error())
	}

	for _, v := range items{
		pk := v.Uuid
		vBytes, err := json.Marshal(v)

		if err!= nil{
			logger.Debug("error marshaling", err)
			return shim.Error(err.Error())
		}
		stub.PutState(pk, vBytes)
	}

	return shim.Success(nil)
}

