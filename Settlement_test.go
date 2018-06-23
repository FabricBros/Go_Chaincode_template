package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"reflect"
	"encoding/json"
	"testing"
	"fmt"
)

func querySettlement(stub *shim.MockStub, name string) *Settlement {

	//print("queryDocument")
	res := stub.MockInvoke("1", [][]byte{[]byte("RetrieveSettlement"), []byte(name)})
	if res.Status != shim.OK {
		fmt.Printf("querySettlement %s failed with %s" , name, string(res.Message))
		return nil
	}
	if res.Payload == nil {
		fmt.Printf("querySettlement %s failed with %s ", name, string(res.Message))
		return nil
	}

	item := &Settlement{}
	_ = json.Unmarshal(res.Payload,item)
	return item
}

var (
	settlements =[]*Settlement{ NewSettlement("Settlement1","Sample Data")}
)

func TestAddSettlements(t *testing.T){
				command := []byte("AddSettlements")
				arg1,_ := json.Marshal(settlements)
				args := [][]byte{command, arg1}

				checkInvoke(stub, args)}

func TestQuerySettlements(t *testing.T){
		var m = querySettlement(stub, settlements[0].Uuid)
		if ! reflect.DeepEqual(m, settlements[0])  {
			fmt.Printf("Does not deep equal %-v", m)
			t.Fail()
		}
	}

func TestUpdateSettlement(t *testing.T){
		command := []byte("UpdateSettlements")
		var updateValue = "1234"

	settlements[0].Data = updateValue

		arg1,_ := json.Marshal(settlements)
		args := [][]byte{command, arg1}

		checkInvoke(stub, args)

		var m = querySettlement(stub, settlements[0].Uuid)
		if m.Data != updateValue {
			t.Fail()
		}
}
