package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"reflect"
	"encoding/json"
	"testing"
	"fmt"
)

var (
	accruals = []*Accrual {{Uuid: "invoice123",Data:"hello"}}
)
func init () {
	scc = new(SimpleChaincode)
	stub = shim.NewMockStub("ex02", scc)
}

func TestAddAccruals(t *testing.T) {
	command := []byte("AddAccruals")
	arg1,_ := json.Marshal(accruals)
	args := [][]byte{command, arg1}

	checkInvoke(stub, args)

	var m = queryAccrual(stub, accruals[0].Uuid)
	if ! reflect.DeepEqual(m, accruals[0]) {
		t.Fail()
	}
}

func TestAccrualUpdate(t *testing.T){
			command := []byte("UpdateAccrual")
			var updateValue = "1234"

			accruals[0].Data = updateValue

			arg1,_ := json.Marshal(accruals)
			args := [][]byte{command, arg1}

			checkInvoke(stub, args)

			var m = queryAccrual(stub, accruals[0].Uuid)
			if m.Data != updateValue {
					t.Fail() //("Value should reflect updated value.")
			}
}


func queryAccrual(stub *shim.MockStub, name string) *Accrual {

	res := stub.MockInvoke("1", [][]byte{[]byte("RetrieveAccrual"), []byte(name)})
	if res.Status != shim.OK {
		fmt.Printf("queryAccrual %s failed with %s" , name, string(res.Message))
		return nil
	}
	if res.Payload == nil {
		fmt.Printf("queryAccrual %s failed with %s ", name, string(res.Message))
		return nil
	}

	item := &Accrual{}
	_ = json.Unmarshal(res.Payload,item)
	return item
}
