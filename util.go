package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
)

func getCN(stub shim.ChaincodeStubInterface)(string, error){
	cert, err := cid.GetX509Certificate(stub)
	if err != nil{
		logger.Debug(err.Error())
		return "", err
	}
	return cert.Subject.CommonName, nil
}

func buildPK(stub shim.ChaincodeStubInterface ,objectType string, args []string)(string, error){

	return stub.CreateCompositeKey(objectType, args)
}

