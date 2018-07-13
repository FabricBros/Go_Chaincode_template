package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//FabricKey
// 	Buyer
// Doc
// Ref (PO#)
// 	Seller
// SKU
// Qty
// Curr
// Unit cost
// Amount
// Type

// For storing arbitrary POs.
type PurchaseOrder struct {
	ObjectType string
	Uuid       string  `json:"FabricKey"`
	Buyer      string  `json:"Buyer"`
	PONo       string  `json:"PO_no"`
	Doc        string  `json:"Doc"` // "PO"
	Ref        string  `json:"Ref"` // PO#
	Seller     string  `json:"Seller"`
	SKU        string  `json:"SKU"`
	Qty        float32 `json:"Qty,string"`
	Curr       string  `json:"Curr" ` //EUR USD
	UnitCost   float32 `json:"UnitCost,string"`
	Amount     float32 `json:"Amount,string"`
	Type       POType  `json:"Type"` // STD NTE
	State      string  `json:"State"`
}

//
type POType string

const (
	STDTYPE = "STD"
	NTETYPE = "NTETYPE"
)

func NewPurchaseOrder(uuid string) *PurchaseOrder {
	return &PurchaseOrder{
		ObjectType: "PurchaseOrder",
		Uuid:       uuid,
	}
}

// ============================================================
// initPO - creates a new PO and stores it in the chaincode state
// ============================================================
func (t *SimpleChaincode) initPurchaseOrders(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("adding PO")
	defer logger.Debug("exit adding PO")

	var items []PurchaseOrder

	err := json.Unmarshal([]byte(args[0]), &items)
	if err != nil {
		logger.Error("Error unmarshing invoice json:", err)
		return shim.Error(err.Error())
	}

	logger.Debugf("We have: %d items", len(items))
	for _, v := range items {
		pk := v.Uuid
		v.ObjectType = "PurchaseOrder"
		vBytes, err := json.Marshal(v)

		if err != nil {
			logger.Error("error marshaling", err)
			return shim.Error(err.Error())
		}
		stub.PutState(pk, vBytes)
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) readPurchaseOrder(stub shim.ChaincodeStubInterface, args []string) pb.Response {
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
		jsonResp = "{\"Error\":\"PurchaseOrder does not exist: " + name + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

// query callback representing the query of a chaincode
func (t *SimpleChaincode) updatePurchaseOrders(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("Enter updatePOs")
	defer logger.Debug("Exited updatePOs")
	var items []PurchaseOrder

	err := json.Unmarshal([]byte(args[0]), &items)
	if err != nil {
		logger.Error("Error unmarshing POs json:", err)
		return shim.Error(err.Error())
	}

	for _, v := range items {
		pk := v.Uuid
		v.ObjectType = "PurchaseOrder"
		vBytes, err := json.Marshal(v)

		if err != nil {
			logger.Debug("error marshaling", err)
			return shim.Error(err.Error())
		}
		stub.PutState(pk, vBytes)
	}

	return shim.Success(nil)
}
