package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"

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

type SimpleChaincode struct {}


// Init initializes chaincode
// ===========================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Init")

	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
// ========================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	//fmt.Println("invoke is running " + function)

	//// Handle different functions
	//if function == "initMarble" { //create a new Marble
	//	return t.initMarble(stub, args)
	//} else if function == "transferMarble" { //change owner of a specific Marble
	//	return t.transferMarble(stub, args)
	//} else if function == "transferMarblesBasedOnColor" { //transfer all Marbles of a certain color
	//	return t.transferMarblesBasedOnColor(stub, args)
	//} else if function == "delete" { //delete a Marble
	//	return t.delete(stub, args)
	//} else if function == "readMarble" { //read a Marble
	//	return t.readMarble(stub, args)
	//} else if function == "queryMarblesByOwner" { //find Marbles for owner X using rich query
	//	return t.queryMarblesByOwner(stub, args)
	//} else if function == "queryMarbles" { //find Marbles based on an ad hoc rich query
	//	return t.queryMarbles(stub, args)
	//} else if function == "getHistoryForMarble" { //get history of values for a Marble
	//	return t.getHistoryForMarble(stub, args)
	//} else if function == "getMarblesByRange" { //get Marbles based on range query
	//	return t.getMarblesByRange(stub, args)
	//}

	// Check for health
	if function == "Ping" {
		return t.Ping(stub)
	}

	// PO operations
	if function == "AddPO" {
		return t.initPurchaseOrders(stub, args)
	}else if function == "RetrievePO" {
		return t.readPurchaseOrder(stub, args)
	}else if function == "UpdatePO" {
		return t.updatePurchaseOrders(stub, args)
	}

	// Entity operations
	if function == "AddEntityMaster" {
		return t.addEntity(stub, args)
	}else if function == "RetrieveEntityMaster" {
		return t.getEntity(stub, args)
	}else if function == "UpdateEntityMaster" {
		return t.updateEntities(stub, args)
	}


	// Document operations
	if function == "AddDocuments" {
		return t.initDocuments(stub, args)
	}else if function == "RetrieveDocument" {
		return t.readDocument(stub, args)
	}else if function == "UpdateDocument"{
		return t.updateDocuments(stub, args)
	}

	// Invoice operations
	if function == "AddInvoices" {
		return t.addInvoices(stub, args)
	} else if function == "RetrieveInvoice" {
		return t.getInvoice(stub, args)
	} else if function == "UpdateInvoices" {
		return t.updateInvoices(stub, args)
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}


func (this *SimpleChaincode) Ping(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("Ping: enter")
	defer logger.Info("Ping: exit")

	return shim.Success([]byte("Ok"))
}

