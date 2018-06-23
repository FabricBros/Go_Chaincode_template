package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hyperledger/fabric/core/chaincode/shim"
		)

var _ = Describe("User operations", func() {
	var (
		users []*User
	)

	BeforeEach(func() {
		users = []*User{ NewUser("User1","Sample Data")}
	})

	Describe("User operations", func() {
		Context("given a User registers with the system", func() {
			It("query should return the default values", func() {
				scc := new(SimpleChaincode)
				stub := shim.NewMockStub("ex02", scc)

				item := NewUser("org1", "userId1")
				command := []byte("initUser")
				arg1 := []byte(item.GroupId)
				arg2 := []byte(item.UserId)
				args := [][]byte{command, arg1, arg2}

				checkInvoke(stub, args)

				var m = queryUser(stub, item.UserId)
				Expect(m.GroupId).To(Equal(item.GroupId), "the GroupId doesn't match")
				Expect(m.UserId).To(Equal(item.UserId), "the UserId doesn't match")
			})
		})
	})
})
