package access

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strings"
	"encoding/pem"
	"fmt"
	"crypto/x509"
)

func GetCommonName(stub shim.ChaincodeStubInterface) (string){



	creator, err := stub.GetCreator()
	var cn string

	if err != nil{
		//error
	}else{
		var certPem = string(creator[:])
		certPem = certPem[strings.Index(certPem, "-----"):]

		block, _ := pem.Decode([]byte(certPem))

		if block == nil{
			fmt.Println("block is null")
		}else{

			cert, certErr := x509.ParseCertificate(block.Bytes)

			if certErr != nil{
				fmt.Println("error parsing cert, "+certErr.Error())
			}else{
				cn =  cert.Subject.CommonName
			}

		}
	}

	return cn
}
