package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// For storing arbitrary SettlementOrders.
type SettlementOrder struct {
	ObjectType             string
	Uuid                   string
	SONo                   string `json:"SO_no"`
	FromCoCo               string `json:"From_CoCo"`
	ToCoCo                 string `json:"To_CoCo"`
	SODate                 string `json:"SO_Date"`
	SOValidTillDate        string `json:"SO_Valid_till_date"`
	Currency               string `json:"Currency"`
	TotalSOAmount          string `json:"Total_SO_Amount"`
	UniqueSONo             string `json:"Unique_SO_no."`
	BuyerAccountingDetails string `json:"Buyer_Accounting_details"`
}

func NewSettlementOrder(uuid string) *SettlementOrder {
	return &SettlementOrder{
		ObjectType: "SettlementOrder",
		Uuid:       uuid,

	}
}

// ============================================================
// initSettlementOrder - creates a new SettlementOrder and stores it in the chaincode state
// ============================================================
func (t *SimpleChaincode) initSettlementOrders(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("adding SettlementOrders")
	defer logger.Debug("exit adding SettlementOrders")

	var items []SettlementOrder

	err := json.Unmarshal([]byte(args[0]), &items)
	if err != nil {
		logger.Error("Error unmarshing invoice json:", err)
		return shim.Error(err.Error())
	}

	for _, v := range items {
		pk := v.Uuid
		v.ObjectType = "SettlementOrder"

		vBytes, err := json.Marshal(v)

		if err != nil {
			logger.Debug("error marshaling", err)
			return shim.Error(err.Error())
		}
		stub.PutState(pk, vBytes)
	}

	return shim.Success(nil)
}

// ===============================================
// readSettlementOrder - read a SettlementOrder from chaincode state
// ===============================================
func (t *SimpleChaincode) readSettlementOrders(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var name, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the SettlementOrder to query")
	}

	name = args[0]
	valAsbytes, err := stub.GetState(name)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + name + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"SettlementOrder does not exist: " + name + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

// query callback representing the query of a chaincode
func (t *SimpleChaincode) updateSettlementOrders(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("Enter updateSettlementOrders")
	defer logger.Debug("Exited updateSettlementOrders")
	var items []SettlementOrder

	err := json.Unmarshal([]byte(args[0]), &items)
	if err != nil {
		logger.Error("Error unmarshing SettlementOrders json:", err)
		return shim.Error(err.Error())
	}

	for _, v := range items {
		pk := v.Uuid
		v.ObjectType = "SettlemenetOrder"
		vBytes, err := json.Marshal(v)

		if err != nil {
			logger.Debug("error marshaling", err)
			return shim.Error(err.Error())
		}
		stub.PutState(pk, vBytes)
	}

	return shim.Success(nil)
}
