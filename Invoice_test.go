package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
		"encoding/json"
	"testing"
	"fmt"
	)

var (
	invoices = []Invoice {{InvoiceNumber: "invoice123",FromCoCo:"hello"}}
)
func init () {
	scc = new(SimpleChaincode)
	stub = shim.NewMockStub("ex02", scc)
}

func TestAddInvoices(t *testing.T) {
	command := []byte("AddInvoices")
	arg1,_ := json.Marshal(invoices)
	args := [][]byte{command, arg1}

	err := checkInvoke(stub, args)
	if err!=nil {
		fmt.Printf("Failed to AddInvoices: %s", err)
	}

	var m = queryInvoice(stub, invoices[0].InvoiceNumber)

	if len(m)==0 ||  m[0].InvoiceNumber != invoices[0].InvoiceNumber || m[0].FromCoCo != invoices[0].FromCoCo {
		t.Fail()
	}
}

func TestInvoiceUpdate(t *testing.T){
			command := []byte("UpdateInvoices")
			var updateValue = "1234"

			invoices[0].FromCoCo = updateValue

			arg1,_ := json.Marshal(invoices)
			args := [][]byte{command, arg1}

			checkInvoke(stub, args)

			var m = queryInvoice(stub, invoices[0].InvoiceNumber)
			if len(m) == 0 || m[0].FromCoCo != updateValue {
					t.Fail() //("Value should reflect updated value.")
			}
}


func queryInvoice(stub *shim.MockStub, name string) []Invoice {

	res := stub.MockInvoke("1", [][]byte{[]byte("RetrieveInvoice"), []byte(name)})
	if res.Status != shim.OK {
		fmt.Printf("queryInvoice %s failed with %s" , name, string(res.Message))
		return nil
	}
	if res.Payload == nil {
		fmt.Printf("queryInvoice %s failed with %s ", name, string(res.Message))
		return nil
	}
	//fmt.Printf("Payload: %s", res.Payload)
	item := make([]Invoice, 0)
	err := json.Unmarshal(res.Payload,&item)
	if err != nil {
		fmt.Printf("Failed to unmarshal: %s", err)
	}
	return item
}
