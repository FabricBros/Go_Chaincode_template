package main

import (
		"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
		"fmt"
	)


var ErrorTransactions= make([]interface{},0)

func init(){
	//logger.SetLevel(shim.LogDebug)
}

// Finds unmatched queries
func (t *SimpleChaincode) getUnmatched(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("enter get unmatched")
	defer logger.Debug("exited get unmatched")

	logger.Debug("- start getUnmatched ")

	//
	getUnmatchedInv, err := stub.GetStateByPartialCompositeKey("unmatched~type~uuid", []string{"invoice"})
	//
	//fmt.Println("Running loop")
	//var i = 0
	//
	//for getUnmatchedInv.HasNext() {
	//	kv, err := getUnmatchedInv.Next()
	//	fmt.Println("Loop", i, "got", kv.Key, kv.Value, err)
	//	i+=1
	//}
	//

	// Query the color~name index by color
	// This will execute a key range query on all keys starting with 'color'
	//coloredMarbleResultsIterator, err := stub.GetStateByPartialCompositeKey("unmatched~uuid", []string{"unmatched",})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer getUnmatchedInv.Close()

	// Iterate through result set and for each Marble found, transfer to newOwner
	var i int
	for i = 0; getUnmatchedInv.HasNext(); i++ {
		// Note that we don't get the value (2nd return variable), we'll just get the Marble name from the composite key
		responseRange, err := getUnmatchedInv.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		// get the color and name from color~name composite key
		objectType, compositeKeyParts, err := stub.SplitCompositeKey(responseRange.Key)
		if err != nil {
			return shim.Error(err.Error())
		}
		returnedUuid := compositeKeyParts[1]
		fmt.Printf("- found an unmatched item index:%s name:%s\n", objectType, returnedUuid)
	}

	responsePayload := fmt.Sprintf("Found unmatched invoices: %d", i)
	fmt.Println("- end getUnmatched: " + responsePayload)
	return shim.Success([]byte(responsePayload))
}
