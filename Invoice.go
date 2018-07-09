package main

import (
	"bytes"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"encoding/json"
	)

//FabricKey	Seller	Date	Ref	Buyer	PO #	SKU	Qty	Curr	Unit cost	Amount
type Invoice struct {
	ObjectType string
	Uuid       string `json:"FabricKey"`
	Seller     string `json:"Seller"`
	Date       string `json:"Date"`
	Ref        string `json:"Ref"`
	Buyer      string `json:"Buyer"`
	PONum      string `json:"PONum"`
	SKU        string `json:"SKU"`
	Qty        int `json:"Qty,string"`
	Curr       string `json:"Curr"`
	UnitCost   string `json:"UnitCost"`
	Amount     string `json:"Amount"`
	State		string `json:"State"`
}

func init(){
	//logger.SetLevel(shim.LogDebug)
}
// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) addInvoices(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("adding invoice")
	defer logger.Debug("exit adding invoice")

	var items []Invoice

	err := json.Unmarshal([]byte(args[0]), &items)
	if err != nil {
		logger.Error("Error unmarshing invoice json:", err)
		return shim.Error(err.Error())
	}

	logger.Debugf("We have: %d items", len(items))
	for _, v := range items {
		//logger.Debugf("Adding: %-v", v)
		pk := v.Uuid
		v.ObjectType="Invoice"
		vBytes, err := json.Marshal(v)

		if err != nil {
			logger.Debug("error marshaling", err)
			return shim.Error(err.Error())
		}

		stub.PutState(pk, vBytes)

		//  ==== Index the Invoice to enable unmatched-based range queries ====
		//  An 'index' is a normal key/value entry in state.
		//  The key is a composite key, with the elements that you want to range query on listed first.
		//  In our case, the composite key is based on indexName~color~name.
		//  This will enable very efficient state range queries based on composite keys matching indexName~color~*
		indexName := "unmatched~type~uuid"
		colorNameIndexKey, err := stub.CreateCompositeKey(indexName, []string{"invoice",v.Uuid})
		if err != nil {
			return shim.Error(err.Error())
		}
		//  Save index entry to state. Only the key name is needed, no need to store a duplicate copy of the Marble.
		//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
		value := []byte{0x00}
		stub.PutState(colorNameIndexKey, value)

	}

	return shim.Success(nil)
}

// Deletes an entity from state
func (t *SimpleChaincode) getInvoice(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("enter get invoice")
	defer logger.Debug("exited get invoice")

	pk := args[0]
	var invoice Invoice

	invoiceByte, err := stub.GetState(pk)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = json.Unmarshal(invoiceByte, &invoice)
	if err != nil {
		logger.Error(err)
		shim.Error(err.Error())
	}
	logger.Debug("getInvoice:")
	logger.Debug(invoice)
	var buffer bytes.Buffer
	buffer.WriteString("[")
	buffer.WriteString(string(invoiceByte))
	buffer.WriteString("]")

	return shim.Success([]byte(buffer.Bytes()))
}

// query callback representing the query of a chaincode
func (t *SimpleChaincode) updateInvoices(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("Enter updateInvoices")
	defer logger.Debug("Exited updateInvoices")
	var invoices []Invoice

	err := json.Unmarshal([]byte(args[0]), &invoices)
	if err != nil {
		logger.Error("Error unmarshing invoice json:", err)
		return shim.Error(err.Error())
	}

	for _, v := range invoices {
		pk := v.Uuid
		vBytes, err := json.Marshal(v)

		if err != nil {
			logger.Debug("error marshaling", err)
			return shim.Error(err.Error())
		}
		stub.PutState(pk, vBytes)
	}

	return shim.Success(nil)
}
