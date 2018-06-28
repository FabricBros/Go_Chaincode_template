package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
		"encoding/json"
	"testing"
	"fmt"
	"io/ioutil"
)

var (
	SettlementOrders =[]*SettlementOrder{ NewSettlementOrder("SettlementOrder1")}
)

func init(){
		SettlementOrders[0].SONo = "SONO1"
}

func querySettlementOrder(stub *shim.MockStub, name string) *SettlementOrder {

	//print("queryDocument")
	res := stub.MockInvoke("1", [][]byte{[]byte("RetrieveSettlement"), []byte(name)})
	if res.Status != shim.OK {
		fmt.Printf("querySettlementOrder %s failed with %s" , name, string(res.Message))
		return nil
	}
	if res.Payload == nil {
		fmt.Printf("querySettlementOrder %s failed with %s ", name, string(res.Message))
		return nil
	}

	item := NewSettlementOrder(name)
	_ = json.Unmarshal(res.Payload,item)
	return item
}



func TestUnmarshalSettlementOrder(t *testing.T ){
	b, err := ioutil.ReadFile("./fixtures/settlementOrder_ex1.json")
	if err != nil {
		fmt.Printf("failed to load example file. ")
		t.Fail()
	}
	json_example := []SettlementOrder{}
	err = json.Unmarshal(b, &json_example)
	if err != nil {
		fmt.Printf("failed to Unmarshal example file. Error: %s ", err.Error())
		t.Fail()
	}
	if len(json_example)!= 1 || json_example[0].FromCoCo!="CO2" || json_example[0].ToCoCo!="CO1" || json_example[0].SONo!="SO3201" {
		fmt.Printf("failed to Unmarshal data from example file. ")
		t.Fail()
	}
}

func TestAddSettlementOrders(t *testing.T){
				command := []byte("AddSettlements")
				arg1,_ := json.Marshal(SettlementOrders)
				args := [][]byte{command, arg1}

				checkInvoke(stub, args)}

func TestQuerySettlementOrders(t *testing.T){
		var m = querySettlementOrder(stub, SettlementOrders[0].Uuid)
		if m == nil || m.SONo != SettlementOrders[0].SONo {
			fmt.Printf("Query SettlementOrder failed with %-v", m)
			t.Fail()
		}
	}

func TestUpdateSettlementOrder(t *testing.T){
		command := []byte("UpdateSettlements")
		var updateValue = "1234"

		SettlementOrders[0].FromCoCo = updateValue

		arg1,_ := json.Marshal(SettlementOrders)
		args := [][]byte{command, arg1}

		checkInvoke(stub, args)

		var m = querySettlementOrder(stub, SettlementOrders[0].Uuid)
		if m.FromCoCo != updateValue {
			t.Fail()
		}
}
