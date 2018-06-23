package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
		"github.com/hyperledger/fabric/core/chaincode/shim"
	"reflect"
	"encoding/json"
	"fmt"
)

func queryPO(stub *shim.MockStub, name string) *PO {

	//print("queryDocument")
	res := stub.MockInvoke("1", [][]byte{[]byte("RetrievePO"), []byte(name)})

	Expect(res.Status).To(BeEquivalentTo(shim.OK), fmt.Sprintf("queryPO %s failed with %s" , name, string(res.Message)))
	Expect(res.Payload).ToNot(BeNil(), fmt.Sprintf("queryPO %s failed with %s ", name, string(res.Message)))

	item := &PO{}
	//fmt.Printf("%s", res.Payload)
	_ = json.Unmarshal(res.Payload,item)
	return item
}


var _ = Describe("PO operations", func() {
	var (
		pos []*PO
	)

	BeforeEach(func() {
		pos = []*PO{ NewPO("PO1","Sample Data")}
	})

	Describe("Basic PO operations", func() {
		Context("given a PO", func() {
			scc := new(SimpleChaincode)
			stub := shim.NewMockStub("ex02", scc)

			It("add POs should succeed", func() {
				command := []byte("AddPOs")
				arg1,_ := json.Marshal(pos)
				args := [][]byte{command, arg1}

				checkInvoke(stub, args)
			})

			It("then query should return an identical po", func() {
				var m = queryPO(stub, pos[0].Uuid)
				Expect(reflect.DeepEqual(m, pos[0])).To(BeTrue(), "DeepEqual should return true")
			})

			It("then update should modify the po", func() {
				command := []byte("UpdatePOs")
				var updateValue = "1234"

				pos[0].Data = updateValue

				arg1,_ := json.Marshal(pos)
				args := [][]byte{command, arg1}

				checkInvoke(stub, args)

				var m = queryPO(stub, pos[0].Uuid)
				Expect(m.Data).Should(Equal(updateValue), "should equal to the new value")
			})
		})
	})
})
