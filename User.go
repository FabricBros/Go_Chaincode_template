package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"encoding/json"
)
/// For storing arbitrary documents.
type User struct {
	ObjectType 		   string `json:"docType"`
	GroupId 	   	   string	`json:"GroupId"`
	UserId 	   		   string	`json:"UserId"`
	UserFirstName 	   string	`json:"UserFirstName"`
	UserAddress 	   string	`json:"UserAddress"`
	UserEmailId 	   string	`json:"UserEmailId"`
	UserContactNo 	   string	`json:"UserContactNo"`
	UserDesignation    string	`json:"UserDesignation"`
	UserExpiryDate	   string	`json:"UserExpiryDate"`
	UserStatus   	   string	`json:"UserStatus"`
	CorrelationId		string 	`json:"CorrelationId"`
}

func NewUser( GroupId, UserId string  ) *User {
	return &User{
		ObjectType: "user",
		GroupId: GroupId,
		UserId: UserId,
	}
}



// ============================================================
// initDocument - creates a new document and stores it in the chaincode state
// ============================================================
func InitUser (stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	//   0			1
	// "groupid"	"userid"
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init User")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}

	item := NewUser(args[0], args[1])

	// ==== Check if Marble already exists ====
	MarbleAsBytes, err := stub.GetState(item.UserId)
	if err != nil {
		return shim.Error("Failed to get Document: " + err.Error())
	} else if MarbleAsBytes != nil {
		fmt.Println("This User already exists: " + item.UserId)
		return shim.Error("This User already exists: " + item.UserId)
	}

	// ==== Create Marble object and marshal to JSON ====
	docuJSONasBytes, err := json.Marshal(item)
	if err != nil {
		return shim.Error(err.Error())
	}
	//Alternatively, build the Marble json string manually if you don't want to use struct marshalling
	//MarbleJSONasString := `{"docType":"Marble",  "name": "` + MarbleName + `", "color": "` + color + `", "size": ` + strconv.Itoa(size) + `, "owner": "` + owner + `"}`
	//MarbleJSONasBytes := []byte(str)

	// === Save Marble to state ===
	err = stub.PutState(item.UserId, docuJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== Marble saved and indexed. Return success ====
	fmt.Println("- end init User")
	return shim.Success(nil)
}

// ===============================================
// readMarble - read a Marble from chaincode state
// ===============================================
func ReadUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var name, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the Document to query")
	}

	name = args[0]
	valAsbytes, err := stub.GetState(name) //get the Marble from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + name + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Document does not exist: " + name + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}
