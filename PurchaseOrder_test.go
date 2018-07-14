package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
		"encoding/json"
	"testing"
	"fmt"
	"io/ioutil"
)

var (
	pos = []*PurchaseOrder{}
)

func init() {
	po1 := &PurchaseOrder{}
	po1.RefID = "PO1"
	po1.Buyer="A1"
	po1.Seller="B1"
	po1.Doc="Document"
	po1.Sku="1910"
	po1.Quantity=10.0
	po1.Currency="USD"
	po1.UnitCost=10.0
	po1.Amount=10.0
	po1.Type=STDTYPE
	pos = append(pos,po1)
}

func queryPurchaseOrder(stub *shim.MockStub, ref string) []PurchaseOrder {

	//print("queryDocument")
	res := stub.MockInvoke("1", [][]byte{[]byte(GET_PO), []byte(ref)})
	if res.Status != shim.OK {
		fmt.Printf("queryPO %s failed with %s", ref, string(res.Message))
		return nil
	}
	if res.Payload == nil {
		fmt.Printf("queryPO %s failed with %s ", ref, string(res.Message))
		return nil
	}

	item := []PurchaseOrder{}
	_ = json.Unmarshal(res.Payload, &item)
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

func AddPo(){
	command := []byte(ADD_PO)
	arg1, _ := json.Marshal(pos)
	args := [][]byte{command, arg1}

	checkInvoke(stub, args)
}

func TestAddPurchaseOrders(t *testing.T) {
		AddPo()
}

func TestQueryPurchaseOrders(t *testing.T) {
	AddPo()

	//var m = queryPurchaseOrder(stub, pos[0].Uuid)

	//if ! reflect.DeepEqual(m, pos[0]) {
	//	t.Fail()
	//}
}

func TestUpdatePO(t *testing.T) {
	AddPo()

	command := []byte(UPDATE_PO)
	var updateValue = "1234"

	pos[0].Buyer = updateValue

	arg1, _ := json.Marshal(pos)
	args := [][]byte{command, arg1}

	checkInvoke(stub, args)

	//var m = queryPurchaseOrder(stub, pos[0].Uuid)
	//if m.Buyer != updateValue {
	//	t.Fail()
	//}
}
