package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"testing"
	"encoding/json"
	"fmt"
)

func queryUser(stub *shim.MockStub, name string) *User {

	//print("queryUser")
	res := stub.MockInvoke("1", [][]byte{[]byte("readUser"), []byte(name)})

	if res.Status != shim.OK {
		fmt.Printf("queryUser %s failed with %s" , name, string(res.Message))
		return nil
	}
	if res.Payload == nil {
		fmt.Printf("queryUser %s failed with %s ", name, string(res.Message))
		return nil
	}

	item := &User{}
	_ = json.Unmarshal(res.Payload,item)
	return item
}

func TestUserInit(t *testing.T){
				scc := new(SimpleChaincode)
				stub := shim.NewMockStub("ex02", scc)

				item := NewUser("org1", "userId1")
				command := []byte("initUser")
				arg1 := []byte(item.GroupId)
				arg2 := []byte(item.UserId)
				args := [][]byte{command, arg1, arg2}

				err := checkInvoke(stub, args)
				if err != nil {
					fmt.Printf("Failed to create user: %s", err)
				}

				var m = queryUser(stub, item.UserId)
				if m.UserId != item.UserId {
					fmt.Printf("Failed to retrieve User %s ", item.UserId)
					t.Fail()
				}
}