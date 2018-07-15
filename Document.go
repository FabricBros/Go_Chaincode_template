package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	)

// For storing arbitrary documents.
type Document struct {
	ObjectType string `json:"docType"`
	Uuid 	   string	`json:"uuid"`
	Data		string `json:"data"`
}


// ============================================================
// initDocument - creates a new document and stores it in the chaincode state
// ============================================================
func (t *SimpleChaincode) initDocuments(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("adding Document")
	defer logger.Debug("exit adding Document")

	logger.Debug("adding: "+args[0])
	logger.Debug("with pk :"+args[1])

	stub.PutState(args[1], []byte(args[0]))

	return shim.Success(nil)
}


func (t *SimpleChaincode) readDocument(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debugf("Enter readDocument %s", args)
	defer logger.Debug("Exited readDocument")

	v,err := stub.GetState(args[0])
	if err!=nil{
		logger.Debug(err.Error())
		return shim.Error(err.Error())
	}

	return shim.Success(v)
}


func (t *SimpleChaincode) updateDocuments(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("Enter updateDocuments")
	defer logger.Debug("Exited updateDocuments")
	logger.Debug("adding: "+args[0])
	logger.Debug("with pk :"+args[1])

	stub.PutState(args[1], []byte(args[0]))


	return shim.Success(nil)
}
