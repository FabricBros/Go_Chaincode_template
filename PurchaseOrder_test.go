package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"reflect"
	"encoding/json"
	"testing"
	"fmt"
	"io/ioutil"
)

var (
	pos = []*PurchaseOrder{NewPurchaseOrder("C1_PO1")}
)

func init() {
	pos[0].PONo = "PO1"
}

func queryPurchaseOrder(stub *shim.MockStub, name string) *PurchaseOrder {

	//print("queryDocument")
	res := stub.MockInvoke("1", [][]byte{[]byte("RetrievePO"), []byte(name)})
	if res.Status != shim.OK {
		fmt.Printf("queryPO %s failed with %s", name, string(res.Message))
		return nil
	}
	if res.Payload == nil {
		fmt.Printf("queryPO %s failed with %s ", name, string(res.Message))
		return nil
	}

	item := &PurchaseOrder{}
	_ = json.Unmarshal(res.Payload, item)
	return item
}

func TestUnmarshalPurchaseOrder(t *testing.T) {
	b, err := ioutil.ReadFile("./fixtures/purchaseOrder_ex1.json")
	if err != nil {
		fmt.Printf("failed to load example file. ")
		t.Fail()
	}
	json_example := []PurchaseOrder{}
	err = json.Unmarshal(b, &json_example)
	if err != nil {
		fmt.Printf("failed to Unmarshal example file. ")
		t.Fail()
	}
	if json_example[0].Buyer != "A5" || json_example[0].Seller != "A3" || len(json_example) != 13 {
		fmt.Printf("failed to Unmarshal example file. ")
		t.Fail()
	}
}

func TestAddPurchaseOrders(t *testing.T) {
	command := []byte("AddPOs")
	arg1, _ := json.Marshal(pos)
	args := [][]byte{command, arg1}

	checkInvoke(stub, args)
}

func TestQueryPurchaseOrders(t *testing.T) {
	var m = queryPurchaseOrder(stub, pos[0].Uuid)
	if (! reflect.DeepEqual(m, pos[0])) {
		t.Fail()
	}
}

func TestUpdatePO(t *testing.T) {
	command := []byte("UpdatePOs")
	var updateValue = "1234"

	pos[0].Buyer = updateValue

	arg1, _ := json.Marshal(pos)
	args := [][]byte{command, arg1}

	checkInvoke(stub, args)

	var m = queryPurchaseOrder(stub, pos[0].Uuid)
	if m.Buyer != updateValue {
		t.Fail()
	}
}
