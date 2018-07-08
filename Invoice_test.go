package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
		"encoding/json"
	"testing"
	"fmt"
	"io/ioutil"
	)

var (
	invoices = []Invoice {}
)
func init () {
	scc = new(SimpleChaincode)
	stub = shim.NewMockStub("ex02", scc)
	logger.SetLevel(shim.LogDebug)

	invoices = append(invoices, Invoice{Uuid: "invoice123", Ref:"123",Seller:"Foo", Buyer:"Bar"})

}

func TestUnmarshalInvoice(t *testing.T ){
	b, err := ioutil.ReadFile("./fixtures/invoice_ex1.json")
	if err != nil {
		fmt.Printf("failed to load example file. ")
		t.Fail()
	}
	json_example := []Invoice{}
	err = json.Unmarshal(b, &json_example)
	if err != nil {
		fmt.Printf("failed to Unmarshal example file. Error: %s ", err.Error())
		t.Fail()
	}
	if json_example[0].Seller!="A3" || json_example[0].Buyer!="A5" || len(json_example) != 21 {
		fmt.Printf("failed to Unmarshal data from example file. ")
		t.Fail()
	}
}

func addInvoice() error{
	command := []byte("AddInvoices")
	arg1,_ := json.Marshal(invoices)
	args := [][]byte{command, arg1}

	err := checkInvoke(stub, args)
	return err
}

func TestAddInvoices(t *testing.T) {
	var err = addInvoice()
	if err!=nil {
		fmt.Printf("Failed to AddInvoices: %s", err)
	}

	var m = queryInvoice(stub, invoices[0].Uuid)

	if m == nil ||  m[0].Ref != invoices[0].Ref || m[0].Seller != invoices[0].Seller {
		t.Fail()
	}
}

func TestInvoiceUpdate(t *testing.T){
	addInvoice()

	command := []byte("UpdateInvoices")
	var updateValue = "1234"

	invoices[0].Seller = updateValue

	arg1,_ := json.Marshal(invoices)
	args := [][]byte{command, arg1}

	checkInvoke(stub, args)

	var m = queryInvoice(stub, invoices[0].Uuid)
	if m == nil || m[0].Seller != updateValue {
			t.Fail() //("Value should reflect updated value.")
	}
}

func queryInvoice(stub *shim.MockStub, name string) []Invoice {
	logger.Debugf("queryInvoice: %s", name)
	defer logger.Debug("queryInvoice out")

	res := stub.MockInvoke("1", [][]byte{[]byte("RetrieveInvoice"), []byte(name)})
	if res.Status != shim.OK {
		fmt.Printf("queryInvoice %s failed with %s" , name, string(res.Message))
		return nil
	}

	if res.Payload == nil {
		fmt.Printf("queryInvoice %s failed with %s ", name, string(res.Message))
		return nil
	}

	logger.Debugf("Payload: %s", res.Payload)
	item := make([]Invoice, 0)
	//item := Invoice{}
	err := json.Unmarshal(res.Payload,&item)
	if err != nil {
		fmt.Printf("Failed to unmarshal: %s", err)
	}

	return item
}
