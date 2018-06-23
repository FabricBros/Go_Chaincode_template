package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
		"github.com/hyperledger/fabric/core/chaincode/shim"
	"reflect"
	"encoding/json"
)

var _ = Describe("Document operations", func() {
	var (
		documents []*Document
	)

	BeforeEach(func() {
		documents = []*Document{ NewDocument("Document1","Sample Data")}
	})

	Describe("Basic Document operations", func() {
		Context("given a Document", func() {
			scc := new(SimpleChaincode)
			stub := shim.NewMockStub("ex02", scc)

			It("query should return an identical document", func() {
				command := []byte("AddDocuments")
				arg1,_ := json.Marshal(documents)
				args := [][]byte{command, arg1}

				checkInvoke(stub, args)

				var m = queryDocument(stub, documents[0].Uuid)
				Expect(reflect.DeepEqual(m, documents[0])).To(BeTrue(), "DeepEqual should return true")
			})
			It("then update should modify the document", func() {
				command := []byte("UpdateDocument")
				var updateValue = "1234"

				documents[0].Data = updateValue

				arg1,_ := json.Marshal(documents)
				args := [][]byte{command, arg1}

				checkInvoke(stub, args)

				var m = queryDocument(stub, documents[0].Uuid)
				Expect(m.Data).Should(Equal(updateValue), "should equal to the new value")
			})
		})
	})
})
