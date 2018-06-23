package main

import (
	//. "github.com/Go_Chaincode_template"
	. "github.com/onsi/gomega"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"reflect"
	"fmt"
	"testing"
	"encoding/json"
)

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func checkState(t *testing.T, stub *shim.MockStub, name string, value string) {
	bytes := stub.State[name]
	if bytes == nil {
		fmt.Println("State", name, "failed to get value")
		t.FailNow()
	}
	if string(bytes) != value {
		fmt.Println("State value", name, "was not", value, "as expected")
		t.FailNow()
	}
}

func queryMarble(stub *shim.MockStub, name string) *Marble {

	//print("queryMarble")
	res := stub.MockInvoke("1", [][]byte{[]byte("readMarble"), []byte(name)})

	Expect(res.Status).To(BeEquivalentTo(shim.OK), fmt.Sprintf("Query %s failed with %s" , name, string(res.Message)))
	Expect(res.Payload).ToNot(BeNil(), fmt.Sprintf("Query %s failed with %s ", name, string(res.Message)))

	marble := &Marble{}
	_ = json.Unmarshal(res.Payload,marble)
	return marble
}
func queryDocument(stub *shim.MockStub, name string) *Document {

	//print("queryDocument")
	res := stub.MockInvoke("1", [][]byte{[]byte("RetrieveDocument"), []byte(name)})

	Expect(res.Status).To(BeEquivalentTo(shim.OK), fmt.Sprintf("queryDocument %s failed with %s" , name, string(res.Message)))
	Expect(res.Payload).ToNot(BeNil(), fmt.Sprintf("queryDocument %s failed with %s ", name, string(res.Message)))

	item := &Document{}
	//fmt.Printf("%s", res.Payload)
	_ = json.Unmarshal(res.Payload,item)
	return item
}

func queryUser(stub *shim.MockStub, name string) *User {

	//print("queryUser")
	res := stub.MockInvoke("1", [][]byte{[]byte("readUser"), []byte(name)})

	Expect(res.Status).To(BeEquivalentTo(shim.OK), fmt.Sprintf("queryUser %s failed with %s" , name, string(res.Message)))
	Expect(res.Payload).ToNot(BeNil(), fmt.Sprintf("queryUser %s failed with %s ", name, string(res.Message)))

	item := &User{}
	_ = json.Unmarshal(res.Payload,item)
	return item
}

func getField(v *Marble, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	if field == "Size"{
		return fmt.Sprintf("%d", f.Int())
	}
	return f.String()
}

func checkInvoke(stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	Expect(res.Status).To(BeEquivalentTo(shim.OK), fmt.Sprintf("checkInvoke %s failed with: %s", args, string(res.Message)))
}
