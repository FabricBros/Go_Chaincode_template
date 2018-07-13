package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"fmt"
			"bytes"
	"strings"
	"encoding/json"
)

type ErrorTransactions struct{
	Invoices []Invoice
	POs		 []PurchaseOrder
}

func init() {
	//logger.SetLevel(shim.LogDebug)
}

func NewErrorTransactions() *ErrorTransactions{
	var ret = ErrorTransactions{}
	ret.Invoices= make([]Invoice,0)
	ret.POs= make([]PurchaseOrder,0)

	return &ret
}
func (t *SimpleChaincode)  match_invoice(stub shim.ChaincodeStubInterface, v *Invoice) {
	if v.PONum == "A9854" { // TC 1
		var pk = "CN_AtlasUSAA9854"
		var postr,err = stub.GetState(pk)
		if err != nil {
			logger.Errorf("Failed to find %s", pk)
		}
		po := &PurchaseOrder{}
		err = json.Unmarshal(postr, po)
		if err != nil {
			logger.Errorf("Failed to unmarshal PO %s", po)
		}

		if po.Qty > v.Qty {
			po.Qty= po.Qty + (-1 * v.Qty)
			v.State="Ok"
		}else{
			v.Qty = po.Qty
			po.Qty = 0
			v.Amount = v.Qty*v.UnitCost
			v.State=fmt.Sprintf("Ok Corrected Qty to %f",v.Qty)
		}

		po.Amount=po.UnitCost*po.Qty

		vBytes, _ := json.Marshal(po)
		//fmt.Printf("PurchaseOrder: %-v\n", po)
		err = stub.PutState(pk, vBytes)
		if err != nil {
			logger.Errorf("Failed to save %s", vBytes)
		}
	}else if v.Uuid=="CN_AtlasUSA1354651A6908"{
		 // TC 2
		 v.Buyer = "A4"
		 v.State="Ok Invalid PO # by CPTY"
	}else if v.Uuid=="CN_AtlasAmericas546568A6910"{
		 // TC 3
		var pk = "CN_AtlasGlobalA6910"
		var postr,err = stub.GetState(pk)
		if err != nil {
			logger.Errorf("Failed to find %s", pk)
		}
		po := &PurchaseOrder{}
		err = json.Unmarshal(postr, po)
		if err != nil {
			logger.Errorf("Failed to unmarshal PO %s", po)
		}

		po.UnitCost=400
		po.Amount=po.UnitCost*po.Qty
		po.State="PO Corrected with new price - 400"
		vBytes, _ := json.Marshal(po)
		err = stub.PutState(pk, vBytes)
		if err != nil {
			logger.Errorf("Failed to save %s", vBytes)
		}
		v.State="Ok"
	}else if v.Uuid=="CN_AtlasTrading56546A691000"{
		// TC 4
		v.PONum = "A6909"
		v.State="Ok PO# Corrected to A6909"
	}else if v.Uuid=="CN_AtlasUSA1354651A5686"{
		v.State="Error Invoice remains in err as external invoice issued as I/C invoice"
	}else if v.Uuid=="CN_AtlasAmericas4684A69879" {
		v.State="Error Invoice remain in err, reason under investigation"
	}else {
		v.State="Ok"
	}
}


func (t *SimpleChaincode) getUnmatchedKeys(stub shim.ChaincodeStubInterface, args []string) []string {
	logger.Debug("enter getUnmatchedKeys")
	defer logger.Debug("exited getUnmatchedKeys")

	getUnmatchedInv, err := stub.GetStateByPartialCompositeKey("unmatched~type~uuid", []string{"invoice"})

	if err != nil {
		return nil
	}
	defer getUnmatchedInv.Close()

	var keys = make([]string, 0)

	// Iterate through result set and for each Marble found, transfer to newOwner
	var i int
	for i = 0; getUnmatchedInv.HasNext(); i++ {
		// Note that we don't get the value (2nd return variable), we'll just get the Marble name from the composite key
		responseRange, err := getUnmatchedInv.Next()
		if err != nil {
			return nil
		}

		// get the color and name from color~name composite key
		_, compositeKeyParts, err := stub.SplitCompositeKey(responseRange.Key)
		if err != nil {
			return nil
		}

		returnedUuid := compositeKeyParts[1]
		keys = append(keys, returnedUuid)
	}

	logger.Debug("- found an unmatched indexes: %s\n", keys)
	return keys
}

// Finds unmatched queries
func (t *SimpleChaincode) getUnmatched(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("enter get unmatched")
	defer logger.Debug("exited get unmatched")

	logger.Debug("- start getUnmatched ")

	var keys = t.getUnmatchedKeys(stub,args)
	if keys == nil {
		return shim.Error("failed to get unMatched ")
	}

	//responsePayload := fmt.Sprintf("Found unmatched invoices: %d", len(keys))
	//var unmatched = NewErrorTransactions()
	var items = []string{}

	for idx:=0; idx < len(keys); idx++ {
		invoiceByte, _ := stub.GetState(keys[idx])
		fmt.Printf("%d - %s\n", idx, string(invoiceByte))
		items = append(items,string(invoiceByte))
	}

	var buffer bytes.Buffer
	buffer.WriteString("[")
	buffer.WriteString(strings.Join(items,","))
	buffer.WriteString("]")

	//var marsh , err = json.Marshal(items)
	//if err != nil {
	//	return shim.Error("getUnmatched: failed to unmarshal")
	//}

	logger.Debugf("- end getUnmatched: %s", string([]byte(buffer.Bytes())))
	return shim.Success([]byte(buffer.Bytes()))
}
