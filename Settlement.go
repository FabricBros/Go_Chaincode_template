package main

import (
		"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// For storing arbitrary Settlements.
type Settlement struct {
	ObjectType string `json:"docType"`
	Uuid 	   string	`json:"uuid"`
	Data		string `json:"data"`
}

func NewSettlement( uuid,data string ) *Settlement {
	return &Settlement{
		ObjectType: "settlement",
		Uuid: uuid,
		Data: data,
	}
}


// ============================================================
// initSettlement - creates a new settlement and stores it in the chaincode state
// ============================================================
func (t *SimpleChaincode) initSettlements(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("adding settlements")
	defer logger.Debug("exit adding settlements")

	var items []Settlement

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


// ===============================================
// readMarble - read a Marble from chaincode state
// ===============================================
func (t *SimpleChaincode) readSettlements(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var name, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the Settlement to query")
	}

	name = args[0]
	valAsbytes, err := stub.GetState(name) //get the Marble from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + name + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Settlement does not exist: " + name + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}


// query callback representing the query of a chaincode
func (t *SimpleChaincode) updateSettlements(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("Enter updateSettlements")
	defer logger.Debug("Exited updateSettlements")
	var items []Settlement

	err := json.Unmarshal([]byte(args[0]), &items)
	if err != nil {
		logger.Error("Error unmarshing Settlements json:", err)
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

