package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
)

func getCN(stub shim.ChaincodeStubInterface)(string, error){
	cert, err := cid.GetX509Certificate(stub)
	if err != nil{
		logger.Error(err.Error())
		return "", err
	}
	if cert == nil { //Will be nill for testing.
		return "", nil
	}
	return cert.Subject.CommonName, nil
}

func buildPK(stub shim.ChaincodeStubInterface ,objectType string, args []string)(string, error){

	return stub.CreateCompositeKey(objectType, args)
}

