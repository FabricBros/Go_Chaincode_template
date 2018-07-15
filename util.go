package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
)

var ignore = false

func getCN(stub shim.ChaincodeStubInterface)(string, error){
	cert, err := cid.GetX509Certificate(stub)
	if err != nil && ignore{
		logger.Error(err.Error())
		return "org", err
	}
	if cert == nil { //Will be nill for testing.
		return "org", nil
	}
	return cert.Subject.CommonName, nil
}

func buildPK(stub shim.ChaincodeStubInterface ,objectType string, args []string)(string, error){

	return stub.CreateCompositeKey(objectType, args)
}

