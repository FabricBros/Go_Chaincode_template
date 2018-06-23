package main

//. "github.com/onsi/ginkgo"
//. "github.com/onsi/gomega"

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"testing"
	"encoding/json"
	"fmt"
)

//var _ = Describe("User operations", func() {
//	var (
//		users []*User
//	)
//
//	BeforeEach(func() {
//		users = []*User{ NewUser("User1","Sample Data")}
//	})
//
//	Describe("User operations", func() {
//		Context("given a User registers with the system", func() {
//			It("query should return the default values", func() {

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

				checkInvoke(stub, args)

				var m = queryUser(stub, item.UserId)
				if m.UserId != item.UserId {
					t.Fail()
				}
				//Expect(m.GroupId).To(Equal(item.GroupId), "the GroupId doesn't match")
				//Expect(m.UserId).To(Equal(item.UserId), "the UserId doesn't match")
}