package main

import (
	"testing"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
	"encoding/json"
	"reflect"
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

func queryMarble(t *testing.T, stub *shim.MockStub, name string) *Marble {

	res := stub.MockInvoke("1", [][]byte{[]byte("readMarble"), []byte(name)})
	if res.Status != shim.OK {
		fmt.Println("Query Name", name, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("Query", name, "failed to get value")
		t.FailNow()
	}

	Marble := &Marble{}
	_ = json.Unmarshal(res.Payload,Marble)

	return Marble
}

func getField(v *Marble, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	if field == "Size"{
		return fmt.Sprintf("%d", f.Int())
	}
	return f.String()
}

func checkMarble(t *testing.T, m *Marble, name string, value string) {
	//fmt.Printf("%-v \n %s \n", m, getField(m,name))
	if getField(m,name) != value {
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

func TestExample02_Create_2_Marbles(t *testing.T) {
	scc := new(SimpleChaincode)
	stub := shim.NewMockStub("ex02", scc)

	// Init A=123 B=234
	//var m = &Marble{Size: 100, Owner: "john", Color:"Yellow",Name:"marbl1"}

	checkInit(t, stub, [][]byte{[]byte("init")})
	checkInvoke(t, stub,[][]byte{[]byte("initMarble"),[]byte("Marble1"),[]byte("blue"),[]byte("35"),[]byte("john")} )
	var m = queryMarble(t, stub, "Marble1")
	//assert(m.Owner,"john")
	print(m)
}

func TestExample02_Query(t *testing.T) {
	scc := new(SimpleChaincode)
	stub := shim.NewMockStub("ex02", scc)

	// Init A=345 B=456
	checkInit(t, stub, [][]byte{[]byte("init"), []byte("A"), []byte("345"), []byte("B"), []byte("456")})
	checkInvoke(t, stub, [][]byte{[]byte("initMarble"), []byte("Marble1"), []byte("blue"), []byte("35"), []byte("bob")})
	checkInvoke(t, stub, [][]byte{[]byte("initMarble"), []byte("Marble2"), []byte("blue"), []byte("40"), []byte("john")})

	// Query A
	Marble1 := queryMarble(t, stub, "Marble1")
	checkMarble(t,Marble1,"Owner","bob")
	checkMarble(t, Marble1,"Size","35")
	//if fmt.Sprintf("%d",Marble1.Size) != "35" {
	//	fmt.Println("Query value", "Size", "was not", "35", "as expected")
	//	t.FailNow()
	//}

	// Query B
	//checkQuery(t, stub, "Marble2", "40")
}

func TestExample02_Invoke(t *testing.T) {
	scc := new(SimpleChaincode)
	stub := shim.NewMockStub("ex02", scc)

	// Init A=567 B=678
	//"Marbleid", "blue", "35", "bob"
	checkInit(t, stub, [][]byte{[]byte("init")})

	// Invoke A->B for 123
	checkInvoke(t, stub, [][]byte{[]byte("initMarble"), []byte("Marble1"), []byte("blue"), []byte("35"), []byte("bob")})
	checkInvoke(t, stub, [][]byte{[]byte("initMarble"), []byte("Marble2"), []byte("blue"), []byte("40"), []byte("john")})
	//checkQuery(t, stub, "Marble1", "35")
	//checkQuery(t, stub, "Marble2", "40")

	// Invoke B->A for 234
	checkInvoke(t, stub, [][]byte{[]byte("transferMarble"), []byte("Marble1"), []byte("john")})
	//checkQuery(t, stub, "bob", "678")
	//checkQuery(t, stub, "john", "567")
	//checkQuery(t, stub, "A", "678")
	//checkQuery(t, stub, "B", "567")s
}
