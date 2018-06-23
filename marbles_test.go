package main
//
//import (
//	. "github.com/onsi/ginkgo"
//	. "github.com/onsi/gomega"
//	"github.com/hyperledger/fabric/core/chaincode/shim"
//	"fmt"
//	"reflect"
//	)
//
//var _ = Describe("Marbles", func() {
//	var (
//		marble *Marble
//	)
//
//	BeforeEach(func() {
//		marble = &Marble{
//			ObjectType: "Marble",
//			Name:  "marble1",
//			Color: "red",
//			Size: 100,
//			Owner:  "john",
//		}
//	})
//
//	Describe("Basic Marble operations", func() {
//		Context("Marble of default values", func() {
//			It("Should have a name, size, owner, and color", func() {
//				Expect(marble.Name).To(Equal("marble1"))
//				Expect(marble.Size).To(Equal(100))
//				Expect(marble.Owner).To(Equal("john"))
//				Expect(marble.Color).To(Equal("red"))
//			})
//		})
//
//		Context("given initMarble with default", func() {
//			It("query should return the default values", func() {
//				scc := new(SimpleChaincode)
//				stub := shim.NewMockStub("ex02", scc)
//				command := []byte("initMarble")
//				arg1 := []byte(marble.Name)
//				arg2 := []byte(marble.Color)
//				arg3 := []byte(fmt.Sprintf("%d", marble.Size))
//				arg4 := []byte(marble.Owner)
//				args := [][]byte{command,arg1,arg2, arg3, arg4}
//
//				checkInvoke(stub,args)
//
//				var m = queryMarble(stub, marble.Name)
//				//fmt.Println(marble)
//				//fmt.Println(m)
//				Expect(reflect.DeepEqual(m, marble)).To(BeTrue(), "DeepEqual should return true")
//			})
//		})
//})})
//func queryMarble(stub *shim.MockStub, name string) *Marble {
//
//	//print("queryMarble")
//	res := stub.MockInvoke("1", [][]byte{[]byte("readMarble"), []byte(name)})
//
//	//Expect(res.Status).To(BeEquivalentTo(shim.OK), fmt.Sprintf("Query %s failed with %s" , name, string(res.Message)))
//	//Expect(res.Payload).ToNot(BeNil(), fmt.Sprintf("Query %s failed with %s ", name, string(res.Message)))
//
//	marble := &Marble{}
//	_ = json.Unmarshal(res.Payload,marble)
//	return marble
//}