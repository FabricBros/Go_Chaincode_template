package main

import (
	"testing"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func TestLedgerQueries(t *testing.T) {
	var cc = &SimpleChaincode{}
	var a = shim.NewMockStub("mock1",cc)
	print(a.Name)
}