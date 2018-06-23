package main

import (
	. "github.com/onsi/ginkgo"
	//. "github.com/Go_Chaincode_template"
	. "github.com/onsi/gomega"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"reflect"
)

var _ = Describe("Marble operations", func() {
	var (
		marble *Marble
	)

	BeforeEach(func() {
		marble = &Marble{
			ObjectType: "Marble",
			Name:       "marble1",
			Color:      "red",
			Size:       100,
			Owner:      "john",
		}
	})

	Describe("Basic Marble operations", func() {
		Context("given a Document", func() {
			docu := []*Document{NewDocument("index", "document123")}
			scc := new(SimpleChaincode)
			stub := shim.NewMockStub("ex02", scc)

			It("query should return an identical document", func() {
				command := []byte("AddDocuments")
				arg1 := []byte("[{\"uuid\":\"index\",\"data\":\"document123\"}]")
				args := [][]byte{command, arg1}

				checkInvoke(stub, args)

				var m = queryDocument(stub, docu[0].Uuid)
				Expect(reflect.DeepEqual(m, docu[0])).To(BeTrue(), "DeepEqual should return true")
			})
			It("then update should modify the document", func() {
				command := []byte("UpdateDocument")
				arg1 := []byte("[{\"uuid\":\"index\",\"data\":\"document1234\"}]")
				args := [][]byte{command, arg1}

				checkInvoke(stub, args)

				var m = queryDocument(stub, docu[0].Uuid)
				Expect(m.Data).Should(Equal("document1234"), "should equal to the new value")
			})
		})

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
