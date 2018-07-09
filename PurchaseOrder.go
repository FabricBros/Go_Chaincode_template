package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"bytes"
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
	Amount   string `json:"amount"`
	Buyer    string `json:"buyer"`
	Currency string `json:"currency"`
	Doc      string `json:"doc"`
	Quantity string    `json:"quantity"`
	RefID    string `json:"refId"`
	Seller   string `json:"seller"`
	Sku      string    `json:"sku"`
	Type     string `json:"type"`
	UnitCost string    `json:"unitCost"`
}

//func NewPurchaseOrder(uuid string) *PurchaseOrder {
//	return &PurchaseOrder{
//		ObjectType: "PurchaseOrder",
//		Uuid:       uuid,
//	}
//}

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
		pk := v.RefID
		//v.ObjectType = "PurchaseOrder"
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

	logger.Debug("writing returned payload")
	var buffer bytes.Buffer
	buffer.WriteString("[")
	buffer.WriteString(string(valAsbytes))
	buffer.WriteString("]")

	return shim.Success([]byte(buffer.Bytes()))
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
		pk := v.RefID
		//v.ObjectType = "PurchaseOrder"
		vBytes, err := json.Marshal(v)

		if err != nil {
			logger.Debug("error marshaling", err)
			return shim.Error(err.Error())
		}
		stub.PutState(pk, vBytes)
	}

	return shim.Success(nil)
}
