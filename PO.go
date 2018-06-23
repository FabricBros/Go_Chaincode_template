package main

import (
		"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// For storing arbitrary POs.
type PO struct {
	ObjectType string `json:"docType"`
	Uuid 	   string	`json:"uuid"`
	Data		string `json:"data"`
}

func NewPO( uuid,data string ) *PO {
	return &PO{
		ObjectType: "PO",
		Uuid: uuid,
		Data: data,
	}
}


// ============================================================
// initPO - creates a new PO and stores it in the chaincode state
// ============================================================
func (t *SimpleChaincode) initPOs(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("adding PO")
	defer logger.Debug("exit adding PO")

	var items []PO

	err := json.Unmarshal([]byte(args[0]), &items)
	if err != nil {
		logger.Error("Error unmarshing invoice json:", err)
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


func (t *SimpleChaincode) readPO(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var name, jsonResp string
	var err error
	logger.Debug("Enter readPO")
	defer logger.Debug("Exited readPO")

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the PO to query")
	}

	name = args[0]
	valAsbytes, err := stub.GetState(name) //get the PO from chaincode state
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
func (t *SimpleChaincode) updatePOs(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("Enter updatePOs")
	defer logger.Debug("Exited updatePOs")
	var items []PO

	err := json.Unmarshal([]byte(args[0]), &items)
	if err != nil {
		logger.Error("Error unmarshing POs json:", err)
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

