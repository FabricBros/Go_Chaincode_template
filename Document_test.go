package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
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
	documents = []*Document{}
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
	command := []byte(ADD_DOCUMENTS)
	args := [][]byte{command, []byte("document"), []byte("pk")}

	checkInvoke(stub, args)

	var m = queryDocument(stub, "pk")
	fmt.Printf("%-v",m)
	if string(m) != "document"  {
		t.Fail()
	}
}
func TestDocumentUpdate(t *testing.T){
			command := []byte(UPDATE_DOCUMENTS)
			var updateValue = []byte("1234")
			var pk = []byte("pk123")
			args := [][]byte{command, updateValue,pk }

			checkInvoke(stub, args)

			var m = queryDocument(stub, string(pk))
			if m==nil || string(m) != string(updateValue) {
					t.Fail() //("Value should reflect updated value.")
			}
}


func queryDocument(stub *shim.MockStub, name string) []byte {

	res := stub.MockInvoke("1", [][]byte{[]byte(GET_DOCUMENTS), []byte(name)})
	if res.Status != shim.OK {
		fmt.Printf("queryDocument %s failed with %s" , name, string(res.Message))
		return nil
	}
	if res.Payload == nil {
		fmt.Printf("queryDocument %s failed with %s ", name, string(res.Message))
		return nil
	}

	return res.Payload
}
