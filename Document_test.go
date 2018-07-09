package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"reflect"
	"encoding/json"
	"testing"
	"os"
	"fmt"
)

var (
	documents []*Document
	stub *shim.MockStub
	scc *SimpleChaincode
)

func mySetupFunction() {
	//print("Setup")
	documents = []*Document{ NewDocument("Document1","Sample Data")}
	scc = new(SimpleChaincode)
	stub = shim.NewMockStub("ex02", scc)
}

func myTeardownFunction() {
	//print("Teardown")
}

func TestMain(m *testing.M) {
	mySetupFunction()
	retCode := m.Run()
	myTeardownFunction()
	os.Exit(retCode)
}


func TestInvoke(t *testing.T) {
	command := []byte("AddDocuments")
	arg1,_ := json.Marshal(documents)
	args := [][]byte{command, arg1}

	checkInvoke(stub, args)

	var m = queryDocument(stub, documents[0].Uuid)
	if ! reflect.DeepEqual(m, documents[0]) {
		t.Fail()
	}
}
func TestDocumentUpdate(t *testing.T){
			command := []byte("UpdateDocument")
			var updateValue = "1234"

			documents[0].Data = updateValue

			arg1,_ := json.Marshal(documents)
			args := [][]byte{command, arg1}

			checkInvoke(stub, args)

			var m = queryDocument(stub, documents[0].Uuid)
			if m==nil || m.Data != updateValue {
					t.Fail() //("Value should reflect updated value.")
			}
}


func queryDocument(stub *shim.MockStub, name string) *Document {

	res := stub.MockInvoke("1", [][]byte{[]byte("RetrieveDocument"), []byte(name)})
	if res.Status != shim.OK {
		fmt.Printf("queryDocument %s failed with %s" , name, string(res.Message))
		return nil
	}
	if res.Payload == nil {
		fmt.Printf("queryDocument %s failed with %s ", name, string(res.Message))
		return nil
	}

	item := &Document{}
	_ = json.Unmarshal(res.Payload,item)
	return item
}
