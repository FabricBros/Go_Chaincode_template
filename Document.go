package main

import "fmt"

// For storing arbitrary documents.
type Document struct {
	ObjectType string `json:"docType"`
	Uuid 	   string	`json:"uuid"`
	Data		string `json:"data"`
}

func NewDocument( uuid,data string ) *Document {
	return &Document{
		ObjectType: "document",
		Uuid: uuid,
		Data: data,
	}
}


// ============================================================
// initDocument - creates a new document and stores it in the chaincode state
// ============================================================
func (t *SimpleChaincode) initDocument(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	//   0			1
	// "uuid"	"arbitrary data"
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init Document")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}

	docu := NewDocument(args[0], args[1])

	// ==== Check if Marble already exists ====
	MarbleAsBytes, err := stub.GetState(docu.Uuid)
	if err != nil {
		return shim.Error("Failed to get Document: " + err.Error())
	} else if MarbleAsBytes != nil {
		fmt.Println("This Document already exists: " + docu.Uuid)
		return shim.Error("This Document already exists: " + docu.Uuid)
	}

	// ==== Create Marble object and marshal to JSON ====
	docuJSONasBytes, err := json.Marshal(docu)
	if err != nil {
		return shim.Error(err.Error())
	}
	//Alternatively, build the Marble json string manually if you don't want to use struct marshalling
	//MarbleJSONasString := `{"docType":"Marble",  "name": "` + MarbleName + `", "color": "` + color + `", "size": ` + strconv.Itoa(size) + `, "owner": "` + owner + `"}`
	//MarbleJSONasBytes := []byte(str)

	// === Save Marble to state ===
	err = stub.PutState(docu.Uuid, docuJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== Marble saved and indexed. Return success ====
	fmt.Println("- end init Document")
	return shim.Success(nil)
}


// ===============================================
// readMarble - read a Marble from chaincode state
// ===============================================
func (t *SimpleChaincode) readDocument(stub shim.ChaincodeStubInterface, args []string) pb.Response {
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
