package main

import (
		"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// For storing arbitrary POs.
type PurchaseOrder struct {
	ObjectType		string
	Uuid			string
	PONo            string `json:"PO_no"`
	FromCoCo        string `json:"From_CoCo"`
	ToCoCo          string `json:"To_CoCo"`
	PODate          string `json:"PO_Date"`
	POType          string `json:"PO_Type"`
	POValidTillDate string `json:"PO_Valid_till_date"`
	Currency        string `json:"Currency"`
	TotalPOAmount   string `json:"Total_PO_Amount"`
	LineLevel []struct {
		LineNo             string `json:"Line_No."`
		ProductCode        string `json:"Product_Code"`
		Description        string `json:"Description"`
		Quantity           string `json:"Quantity"`
		MeasurementUnit    string `json:"Measurement_Unit"`
		PricePerUnit       string `json:"Price_per_Unit"`
		TotalPrice         string `json:"Total_Price"`
		LineItemTax1       string `json:"Line_item_Tax_1"`
		LineItemTaxAmount1 string `json:"Line_item_Tax_Amount_1_"`
		LineItemTax2       string `json:"Line_item_Tax_2"`
		LineItemTaxAmount2 string `json:"Line_item_Tax_Amount_2"`
		LineItemTax3       string `json:"Line_item_Tax_3"`
		LineItemTaxAmount3 string `json:"Line_item_Tax_Amount_3"`
		LineItemTax4       string `json:"Line_item_Tax_4"`
		LineItemTaxAmount4 string `json:"Line_item_Tax_Amount_4"`
		LineItemTax5       string `json:"Line_item_Tax_5"`
		LineItemTaxAmount5 string `json:"Line_item_Tax_Amount_5"`
		LineItemTotalPrice string `json:"Line_item_total_Price"`
	} `json:"line_level"`
}

func NewPurchaseOrder( uuid string ) *PurchaseOrder {
	return &PurchaseOrder{
		ObjectType: "PurchaseOrder",
		Uuid: uuid,
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

	for _, v := range items{
		pk := v.PONo
		v.ObjectType="PurchaseOrder"
		vBytes, err := json.Marshal(v)

		if err!= nil{
			logger.Debug("error marshaling", err)
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

	for _, v := range items{
		pk := v.Uuid
		v.ObjectType="PurchaseOrder"
		vBytes, err := json.Marshal(v)

		if err!= nil{
			logger.Debug("error marshaling", err)
			return shim.Error(err.Error())
		}
		stub.PutState(pk, vBytes)
	}

	return shim.Success(nil)
}

