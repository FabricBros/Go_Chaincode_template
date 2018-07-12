package main

import (
	"bytes"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"encoding/json"
	)

//FabricKey	Seller	Date	Ref	Buyer	PO #	SKU	Qty	Curr	Unit cost	Amount
type Invoice struct {
	Amount   string `json:"amount"`
	Buyer    string `json:"buyer"`
	Currency string `json:"currency"`
	Date     string `json:"date"`
	PoNumber string `json:"poNumber"`
	Quantity string    `json:"quantity"`
	RefID    string    `json:"refId"`
	Seller   string `json:"seller"`
	Sku      string    `json:"sku"`
	UnitCost string    `json:"unitCost"`
}

func init(){
	//logger.SetLevel(shim.LogDebug)
}
// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) addInvoices(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("adding invoice")
	defer logger.Debug("exit adding invoice")

	var items []Invoice
	logger.Debug("addInvoice:"+args[0])

	err := json.Unmarshal([]byte(args[0]), &items)
	if err != nil {
		logger.Error("Error unmarshing invoice json:", err)
		return shim.Error(err.Error())
	}

	logger.Debugf("We have: %d items", len(items))

	//create PK CN+REF+PO
	for _, v := range items {
		//logger.Debugf("Adding: %-v", v)
		cn , err:= getCN(stub)
		if err != nil{
			logger.Debug(err.Error())
			shim.Error(err.Error())
		}
		logger.Debug("adding item:")
		logger.Debug(v)

		var attr = []string{cn, v.RefID, v.PoNumber}

		pk, err := buildPK(stub, "Invoice", attr)

		logger.Debug("Invoice has pk: "+pk)

		vBytes, err := json.Marshal(v)

		if err != nil {
			logger.Debug("error marshaling", err)
			return shim.Error(err.Error())
		}
		stub.PutState(pk, vBytes)
	}

	return shim.Success(nil)
}

// Deletes an entity from state
func (t *SimpleChaincode) getInvoice(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("enter get invoice")
	defer logger.Debug("exited get invoice")

	//key := args[0]
	cn, err := getCN(stub)

	var attr = []string{cn, args[0], args[1]}
	pk, err := buildPK(stub, "Invoice", attr)
	var invoice Invoice

	logger.Debug("getInvoice using pk:"+pk)


	invoiceByte, err := stub.GetState(pk)
	logger.Debugf("got back from state %s", invoiceByte)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = json.Unmarshal(invoiceByte, &invoice)
	if err != nil {
		logger.Error(err)
		return shim.Error(err.Error())
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
		pk := v.RefID
		vBytes, err := json.Marshal(v)

		if err != nil {
			logger.Debug("error marshaling", err)
			return shim.Error(err.Error())
		}
		stub.PutState(pk, vBytes)
	}

	return shim.Success(nil)
}
