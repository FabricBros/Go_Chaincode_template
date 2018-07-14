package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
	"testing"
	"fmt"
	"io/ioutil"
)

var (
	invoices = []Invoice{}
)

func init() {
	scc = new(SimpleChaincode)
	stub = shim.NewMockStub("ex02", scc)
	//stub.
	logger.SetLevel(shim.LogDebug)

	invoices = append(invoices, Invoice{PoNumber: "po123", RefID: "123", Seller: "Foo", Buyer: "Bar"})

}

func TestUnmarshalInvoice(t *testing.T) {
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
	if json_example[0].Seller != "A3" || json_example[0].Buyer != "A5" || len(json_example) != 21 {
		fmt.Printf("failed to Unmarshal data from example file. ")
		t.Fail()
	}
}

func addInvoice() error {
	command := []byte("AddInvoices")
	arg1, _ := json.Marshal(invoices)
	args := [][]byte{command, arg1}

	err := checkInvoke(stub, args)
	return err
}

func TestAddInvoices(t *testing.T) {
	var err = addInvoice()
	if err != nil {
		fmt.Printf("Failed to AddInvoices: %s", err)
	}
	//var pk = buildPK(stub,"Invoice",{})
	//var m = queryInvoice(stub, invoices)

	//if m == nil || m[0].Ref != invoices[0].Ref || m[0].Seller != invoices[0].Seller {
	//	t.Fail()
	//}
}

func TestInvoiceUpdate(t *testing.T) {
	addInvoice()

	command := []byte("UpdateInvoices")
	var updateValue = "1234"

	invoices[0].Seller = updateValue

	arg1, _ := json.Marshal(invoices)
	args := [][]byte{command, arg1}

	checkInvoke(stub, args)

	//var m = queryInvoice(stub, invoices[0].Uuid)
	//if m == nil || m[0].Seller != updateValue {
	//	t.Fail() //("Value should reflect updated value.")
	//}
}

func queryUnmatched(stub *shim.MockStub) []map[string]string {
	logger.Debugf("queryUnmatched")
	defer logger.Debug("queryUnmatched out")

	res := stub.MockInvoke("1", [][]byte{[]byte(GET_UNMATCHED), []byte("")})
	if res.Status != shim.OK {
		fmt.Printf("queryUnmatched failed with %s", string(res.Message))
		return nil
	}

	if res.Payload == nil {
		fmt.Printf("queryInvoice failed with %s ", string(res.Message))
		return nil
	}

	//logger.Debugf("Payload: %s", string(res.Payload))
	item := make([]map[string]string, 0)
	err := json.Unmarshal(res.Payload, &item)
	if err != nil {
		logger.Errorf("Failed to unmarshal: %s", err)
	}

	return item
}

func queryInvoice(stub *shim.MockStub, ref,ponum string) []Invoice {
	logger.Debugf("queryInvoice: %s", ref+ponum)
	defer logger.Debug("queryInvoice out")

	res := stub.MockInvoke("1", [][]byte{[]byte("RetreiveInvoice"), []byte(ref),[]byte(ponum)})
	if res.Status != shim.OK {
		fmt.Printf("queryInvoice failed with %s",  string(res.Message))
		return nil
	}

	if res.Payload == nil {
		fmt.Printf("queryInvoice failed with %s ",  string(res.Message))
		return nil
	}

	//logger.Debugf("Payload: %s", res.Payload)
	item := make([]Invoice, 0)
	//item := Invoice{}
	err := json.Unmarshal(res.Payload, &item)
	if err != nil {
		fmt.Printf("Failed to unmarshal: %s", err)
	}

	return item
}
