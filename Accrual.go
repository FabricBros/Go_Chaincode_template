package main

import (
		"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// For storing arbitrary Accruals.
type Accrual struct {
	ObjectType string `json:"docType"`
	Uuid 	   string	`json:"uuid"`
	Data		string `json:"data"`
}

func NewAccrual( uuid,data string ) *Accrual {
	return &Accrual{
		ObjectType: "Accrual",
		Uuid: uuid,
		Data: data,
	}
}


// ============================================================
// initAccrual - creates a new Accrual and stores it in the chaincode state
// ============================================================
func (t *SimpleChaincode) initAccruals(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("adding Accrual")
	defer logger.Debug("exit adding Accrual")

	var items []Accrual

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


func (t *SimpleChaincode) readAccrual(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var name, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the Accrual to query")
	}

	name = args[0]
	valAsbytes, err := stub.GetState(name) //get the Accrual from chaincode state
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
func (t *SimpleChaincode) updateAccruals(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("Enter updateAccruals")
	defer logger.Debug("Exited updateAccruals")
	var items []Accrual

	err := json.Unmarshal([]byte(args[0]), &items)
	if err != nil {
		logger.Error("Error unmarshing Accruals json:", err)
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

