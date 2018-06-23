package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"reflect"
	"fmt"
	"testing"
	)

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func checkState(t *testing.T, stub *shim.MockStub, name string, value string) {
	bytes := stub.State[name]
	if bytes == nil {
		fmt.Println("State", name, "failed to get value")
		t.FailNow()
	}
	if string(bytes) != value {
		fmt.Println("State value", name, "was not", value, "as expected")
		t.FailNow()
	}
}


func getField(v *Marble, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	if field == "Size"{
		return fmt.Sprintf("%d", f.Int())
	}
	return f.String()
}

func checkInvoke(stub *shim.MockStub, args [][]byte) error{
	resp := stub.MockInvoke("1", args)
	if resp.Status != shim.OK {
		//assert
		return fmt.Errorf("Expected 200 got %d", resp.Status)
	}
	return nil
	//Expect(resp.Status).To(BeEquivalentTo(shim.OK), fmt.Sprintf("checkInvoke %s failed with: %s", args, string(res.Message)))
}

