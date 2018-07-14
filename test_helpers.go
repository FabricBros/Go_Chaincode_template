package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
		)
func checkInvoke(stub *shim.MockStub, args [][]byte) error{
	resp := stub.MockInvoke("1", args)
	if resp.Status != shim.OK {
		logger.Errorf("Failed to Invoke %s", args)
		return fmt.Errorf("expected 200 got %d", resp.Status)
	}
	return nil
}

