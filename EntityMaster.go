package main

import (
	"bytes"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
	"fmt"
)

type EntityMaster struct {
	Account           string    `json:"account"`
	AdditionalReview  string `json:"additionalReview"`
	Bank              string `json:"bank"`
	Country           string `json:"country"`
	FnlCurr           string `json:"fnlCurr"`
	GlEntityCode      string `json:"glEntityCode"`
	Group             string `json:"group"`
	NettingSettRules  string `json:"nettingSettRules"`
	Paymaster         string `json:"paymaster"`
	PaymasterEligible string `json:"paymasterEligible"`
	SubName           string `json:"subName"`
	Wht               string `json:"wht"`
}

func init(){
	//logger.SetLevel(shim.LogDebug)
}
// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) addEntity(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("adding invoice")
	defer logger.Debug("exit adding invoice")

	var items []EntityMaster

	err := json.Unmarshal([]byte(args[0]), &items)
	if err != nil {
		logger.Error("Error unmarshing invoice json:", err)
		return shim.Error(err.Error())
	}

	logger.Debugf("We have: %d items", len(items))
	idTool, err := cid.New(stub)
	if err != nil{
		return shim.Error(err.Error())
	}

	cert, err:= idTool.GetX509Certificate()
	if err != nil{
		fmt.Println("error getting cert")
	}

	nameOfcaller :=cert.Subject
	cn := nameOfcaller.CommonName

	logger.Debug("name of caller is :"+nameOfcaller.CommonName )

	//create PK CN+REF+PO


	for _, v := range items {
		var attr []string = []string{cn}

		pks, err := buildPK(stub, "EntityMaster", attr)
		if err != nil{
			logger.Debug(err.Error())
			return shim.Error(err.Error())
		}
		//logger.Debugf("Adding: %-v", v)
		pk := pks
		vBytes, err := json.Marshal(v)

		if err != nil {
			logger.Debug("error marshaling", err)
			return shim.Error(err.Error())
		}
		stub.PutState(pk, vBytes)
	}

	return shim.Success(nil)
}

// Deletes an entity from state
func (t *SimpleChaincode) getEntity(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("enter get invoice")
	defer logger.Debug("exited get invoice")

	pks, err := buildPK(stub, "EntityMaster", args)
	if err != nil{
		logger.Debug(err.Error())
		return shim.Error(err.Error())
	}

	var entityMaster EntityMaster

	entityId, err := stub.GetState(pks)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = json.Unmarshal(entityId, &entityMaster)
	if err != nil {
		logger.Error(err)
		return shim.Error(err.Error())
	}
	logger.Debug("getEntity:")
	logger.Debug(entityMaster)
	var buffer bytes.Buffer
	buffer.WriteString("[")
	buffer.WriteString(string(entityId))
	buffer.WriteString("]")

	return shim.Success([]byte(buffer.Bytes()))
}

// query callback representing the query of a chaincode
func (t *SimpleChaincode) updateEntities(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("Enter updateInvoices")
	defer logger.Debug("Exited updateInvoices")
	var entity []EntityMaster

	err := json.Unmarshal([]byte(args[0]), &entity)
	if err != nil {
		logger.Error("Error unmarshing invoice json:", err)
		return shim.Error(err.Error())
	}

	for _, v := range entity {
		cn, err := getCN(stub)
		if err != nil{
			logger.Debug(err)
			return shim.Error(err.Error())
		}
		var attr []string = []string{cn}

		pks, err := buildPK(stub, "EntityMaster", attr)
		if err != nil{
			logger.Debug(err.Error())
			return shim.Error(err.Error())
		}
		//logger.Debugf("Adding: %-v", v)
		pk := pks

		vBytes, err := json.Marshal(v)

		if err != nil {
			logger.Debug("error marshaling", err)
			return shim.Error(err.Error())
		}
		stub.PutState(pk, vBytes)
	}

	return shim.Success(nil)
}
