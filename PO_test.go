package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"reflect"
	"encoding/json"
	"testing"
	"fmt"
)

func queryPO(stub *shim.MockStub, name string) *PO {

	//print("queryDocument")
	res := stub.MockInvoke("1", [][]byte{[]byte("RetrievePO"), []byte(name)})
	if res.Status != shim.OK {
		fmt.Printf("queryPO %s failed with %s" , name, string(res.Message))
		return nil
	}
	if res.Payload == nil {
		fmt.Printf("queryPO %s failed with %s ", name, string(res.Message))
		return nil
	}

	item := &PO{}
	_ = json.Unmarshal(res.Payload,item)
	return item
}

var (
	pos =[]*PO{ NewPO("PO1","Sample Data")}
)

func TestAddPOs(t *testing.T){
				command := []byte("AddPOs")
				arg1,_ := json.Marshal(pos)
				args := [][]byte{command, arg1}

				checkInvoke(stub, args)}

func TestQueryPOs(t *testing.T){
		var m = queryPO(stub, pos[0].Uuid)
		if (! reflect.DeepEqual(m, pos[0])) {
			t.Fail()
		}
	}

func TestUpdatePO(t *testing.T){
		command := []byte("UpdatePOs")
		var updateValue = "1234"

		pos[0].Data = updateValue

		arg1,_ := json.Marshal(pos)
		args := [][]byte{command, arg1}

		checkInvoke(stub, args)

		var m = queryPO(stub, pos[0].Uuid)
		if m.Data != updateValue {
			t.Fail()
		}
}
