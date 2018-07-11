package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
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


func checkInvoke(stub *shim.MockStub, args [][]byte) error{
	//logger.Debugf("checkInvoke: %s", args)
	resp := stub.MockInvoke("1", args)
	if resp.Status != shim.OK {
		logger.Errorf("Failed to Invoke %s", args)
		return fmt.Errorf("expected 200 got %d", resp.Status)
	}
	return nil
}

