package main

import (
	"testing"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
	"encoding/json"
)

func TestLedgerQueries(t *testing.T) {
	var cc = &SimpleChaincode{}
	var a = shim.NewMockStub("mock1",cc)
	print(a.Name)
}


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

func checkQuery(t *testing.T, stub *shim.MockStub, name string, value string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("readMarble"), []byte(name)})
	if res.Status != shim.OK {
		fmt.Println("Query Name", name, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("Query", name, "failed to get value")
		t.FailNow()
	}
	//fmt.Printf("%s",res.Payload)
	marble := &marble{}
	_ = json.Unmarshal(res.Payload,marble)

	if fmt.Sprintf("%d",marble.Size) != value {
		fmt.Println("Query value", name, "was not", value, "as expected")
		t.FailNow()
	}
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
}

func TestExample02_Init(t *testing.T) {
	scc := new(SimpleChaincode)
	stub := shim.NewMockStub("ex02", scc)

	// Init A=123 B=234
	checkInit(t, stub, [][]byte{[]byte("init"), []byte("A"), []byte("123"), []byte("B"), []byte("234")})

	checkState(t, stub, "A", "123")
	checkState(t, stub, "B", "234")
}

func TestExample02_Query(t *testing.T) {
	scc := new(SimpleChaincode)
	stub := shim.NewMockStub("ex02", scc)

	// Init A=345 B=456
	checkInit(t, stub, [][]byte{[]byte("init"), []byte("A"), []byte("345"), []byte("B"), []byte("456")})

	// Query A
	checkQuery(t, stub, "A", "345")

	// Query B
	checkQuery(t, stub, "B", "456")
}

func TestExample02_Invoke(t *testing.T) {
	scc := new(SimpleChaincode)
	stub := shim.NewMockStub("ex02", scc)

	// Init A=567 B=678
	//"marbleid", "blue", "35", "bob"
	checkInit(t, stub, [][]byte{[]byte("init")})

	// Invoke A->B for 123
	checkInvoke(t, stub, [][]byte{[]byte("initMarble"), []byte("marble1"), []byte("blue"), []byte("35"), []byte("bob")})
	checkInvoke(t, stub, [][]byte{[]byte("initMarble"), []byte("marble2"), []byte("blue"), []byte("40"), []byte("john")})
	checkQuery(t, stub, "marble1", "35")
	checkQuery(t, stub, "marble2", "40")

	// Invoke B->A for 234
	checkInvoke(t, stub, [][]byte{[]byte("transferMarble"), []byte("marble1"), []byte("john")})
	checkQuery(t, stub, "bob", "678")
	checkQuery(t, stub, "john", "567")
	checkQuery(t, stub, "A", "678")
	checkQuery(t, stub, "B", "567")
}
