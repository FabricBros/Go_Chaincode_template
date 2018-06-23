package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var logger = shim.NewLogger("main logger")

// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}