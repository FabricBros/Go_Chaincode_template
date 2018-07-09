package access

import "github.com/hyperledger/fabric/core/chaincode/shim"

type AccessControl interface{
	GetCommonName(stub shim.ChaincodeStubInterface) (string, error)

}
