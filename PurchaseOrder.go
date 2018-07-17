package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"bytes"
)


// For storing arbitrary POs.
type PurchaseOrder struct {
	Amount   float32 `json:"amount,string"`
	Buyer    string `json:"buyer"`
	Currency string `json:"currency"`
	Doc      string `json:"doc"`
	Quantity float32    `json:"quantity,string"`
	RefID    string `json:"refId"`
	Seller   string `json:"seller"`
	Sku      string    `json:"sku"`
	Type     POType `json:"type"`
	UnitCost float32    `json:"unitCost,string"`
	//State 	string `json:"state"`
}

//
type POType string

const (
	STDTYPE = "STD"
	NTETYPE = "NTETYPE"
)

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

		cn, err := getCN(stub)

		if err != nil {
			logger.Errorf("err: %-v", err.Error())
			return shim.Error(err.Error())
		}
		var attr = []string{cn, v.RefID}
		//logger.Debugf("attr: %-v",attr )
		pk, err := buildPK(stub, "PurchaseOrder", attr)
		if err != nil{
			logger.Error(err.Error())
			return shim.Error(err.Error())
		}
		//logger.Debug("using pk "+pk)

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
	logger.Debug("Enter readPO: %s", args)
	defer logger.Debug("Exited readPO")

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the PO to query")
	}

	//name = args[0]
	cn, err := getCN(stub)
	if err!=nil{
		logger.Error(err)
		return shim.Error(err.Error())
	}
	var attr = []string{cn, args[0]}

	pk, err := buildPK(stub, "PurchaseOrder", attr)

	//logger.Errorf("find %s", pk)
	valAsbytes, err := stub.GetState(pk) //get the PO from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + name + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"PurchaseOrder does not exist: " + name + "\"}"
		return shim.Error(jsonResp)
	}

	logger.Debug("writing returned payload %s",string(valAsbytes))
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
		cn, err := getCN(stub)
		if err != nil{
			logger.Error(err.Error())
			return shim.Error(err.Error())
		}
		var attr = []string{cn, v.RefID}
		pk, err := buildPK(stub, "PurchaseOrder", attr)
		if err != nil{
			logger.Error(err.Error())
			return shim.Error(err.Error())
		}
		logger.Debug("using pk "+pk)

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
