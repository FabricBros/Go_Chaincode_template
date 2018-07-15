package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"testing"
	"io/ioutil"
	"strings"
	"fmt"
)

func init() {
	scc = new(SimpleChaincode)
	stub = shim.NewMockStub("ex02", scc)
}

func loadDataset(t *testing.T) error {

	b, err := ioutil.ReadFile("./fixtures/purchaseOrder_ex1.json")
	if err != nil {
		logger.Errorf("failed to load example file. ")
		t.Fail()
	}

	command := []byte(ADD_PO)
	args := [][]byte{command, b}
	err = checkInvoke(stub, args)
	if err != nil {
		logger.Errorf("Failed to PurchaseOrders: %s", err)
		t.Fail()

	}

	b, err = ioutil.ReadFile("./fixtures/invoice_ex1.json")
	if err != nil {
		logger.Error("failed to load example file. ")
		t.Fail()
	}

	command = []byte(ADD_INVOICES)
	args = [][]byte{command, b}
	err = checkInvoke(stub, args)
	if err != nil {
		logger.Errorf("Failed to AddInvoices: %s", err)
		t.Fail()
	}

	return err
}

func TestLoadDataset(t *testing.T) {
	loadDataset(t)
	po := queryPurchaseOrder(stub, "CN_AtlasIndiaA1235")
	inv := queryInvoice(stub,"68109","A1235")
	logger.Debugf("PO: %s", po)
	logger.Debugf("Inv: %s", inv)
}

//func TestMatching_InvoiceWithoutPO(t *testing.T) {
//	var invoices_str = []byte(`[{
//    "FabricKey": "CN_AtlasTrading80199A9854",
//    "Seller": "A6",
//    "Date": "2-Jan",
//    "Ref": "80199",
//    "Buyer": "A2",
//    "PONum": "A9854",
//    "SKU": "23598",
//    "Qty": "300",
//    "Curr": "USD",
//    "UnitCost": "100",
//    "Amount": "30000"
//  },{
//    "FabricKey": "CN_AtlasT1ading80199A9855",
//    "Seller": "A6",
//    "Date": "2-Jan",
//    "Ref": "80199",
//    "Buyer": "A2",
//    "PONum": "A9854",
//    "SKU": "23598",
//    "Qty": "300",
//    "Curr": "USD",
//    "UnitCost": "100",
//    "Amount": "30000"
//  }]`)
//
//	command2 := []byte(ADD_INVOICES)
//	args2 := [][]byte{command2, invoices_str}
//
//	err := checkInvoke(stub, args2)
//	if err != nil {
//		t.Errorf("failed to create Invoice")
//	}
//
//	getInv := queryInvoice(stub, "CN_AtlasTrading80199A9854")
//
//	if getInv == nil {
//		t.Errorf("failed to retrieve Invoice")
//	}
//
//	var items = queryUnmatched(stub)
//
//	if items == nil || len(items) != 2 {
//		t.Errorf("failed to retrieve unmatched invoices ~ len(%d)", len(items))
//	}
//
//	//var inv =
//
//	if items[0]["Seller"] != "A6" {
//		t.Errorf("unmatched invoice should have seller A6")
//	}
//}

//func TestMatching_Invoice2Po(t *testing.T) {
//	var pos_str = []byte(`[{"FabricKey": "CN_AtlasUSAA9854",
//"Buyer": "A2", "Doc": "PO", "Ref": "A9854",
//"Seller": "A6","SKU": "23598","Qty": "1000",
//"Curr": "USD","UnitCost": "100","Amount": "100000.00",
//"Type": "NTE"}]`)
//
//	command := []byte(ADD_PO)
//	args := [][]byte{command, pos_str}
//
//	var invoices_str = []byte(`[{
//    "FabricKey": "CN_AtlasTrading80199A9854",
//    "Seller": "A6",
//    "Date": "2-Jan",
//    "Ref": "80199",
//    "Buyer": "A2",
//    "PONum": "A9854",
//    "SKU": "23598",
//    "Qty": "300",
//    "Curr": "USD",
//    "UnitCost": "100",
//    "Amount": "30000"
//  }]`)
//
//	command2 := []byte(ADD_INVOICES)
//	args2 := [][]byte{command2, invoices_str}
//
//	err := checkInvoke(stub, args)
//	if err != nil {
//		t.Errorf("Failed to create PO")
//	}
//
//	err = checkInvoke(stub, args2)
//	if err != nil {
//		t.Errorf("Failed to create Invoice")
//	}
//
//	getPO := queryPurchaseOrder(stub, "CN_AtlasUSAA9854")
//	if getPO == nil {
//		t.Errorf("Failed to retrieve PO")
//	}
//
//	getInv := queryInvoice(stub, "CN_AtlasTrading80199A9854")
//
//	if getInv == nil {
//		t.Errorf("Failed to retrieve Invoice")
//	}
//
//	var items = queryUnmatched(stub)
//
//	if items == nil || len(items) != 0 {
//		t.Errorf("Failed to match invoice to PO ~ len(%d)", len(items))
//	}
//
//}

// Test case 1.
// FabricKey  CN_AtlasTrading80203A9854
// Seller	A6
// Buyer A2
// Ref 80203
// PO # A9854
// SKU 23598
// Qty 180
// Curr USD
// Unit 100
// Amount 18,000
// Inv Stts error
// If dispute reason 	```PO Exceed NTE quantity```
// Disp res date 25-Jan
// Dispute resolution steps taken ```Old PO amended for the quantity```
func TestMatching_POExceedsNTE(t *testing.T) {
	loadDataset(t)
	//PO:
	//"CN_AtlasUSAA9854	A2	PO	A9854	A6	23598	1,000	USD	100	 $100,000.00 	NTE"
	//Inv:
	//FabricKey				Seller	Date	Ref	Buyer	PO #	SKU	Qty	Curr	Unit cost	Amount
	//CN_AtlasTrading80199A9854	A6	2-Jan	80199	A2	A9854	23598	300	USD	100	30,000
	//CN_AtlasTrading80200A9854	A6	7-Jan	80200	A2	A9854	23598	200	USD	100	20,000
	//CN_AtlasTrading80201A9854	A6	10-Jan	80201	A2	A9854	23598	150	USD	100	15,000
	//CN_AtlasTrading80202A9854	A6	19-Jan	80202	A2	A9854	23598	200	USD	100	20,000
	//CN_AtlasTrading80203A9854	A6	20-Jan	80203	A2	A9854	23598	180	USD	100	18,000
	inv := queryInvoice(stub,"80203","A9854")
	if inv[0].Quantity != 150 {
		t.Errorf("Failed to match invoice, qty should be 150 is %f", inv[0].Quantity)
	}
}

// Test case 2
//CN_AtlasUSA1354651A6908
//A2
//A4
//1354651
//A6908
//12345
//100
//USD
//200
//20,000
//error
//Invalid PO # by CPTY
//5-Jan
//CPTY corrected

func TestMatching_CorrectedInvoiceBuyer(t *testing.T) {
	loadDataset(t)
	//PO:
	//"CN_AtlasUSAA9854	A2	PO	A9854	A6	23598	1,000	USD	100	 $100,000.00 	NTE"
	//Inv:
	//FabricKey					Seller	Date	Ref	Buyer	PO #	SKU	Qty	Curr	Unit cost	Amount
	//CN_AtlasTrading80199A9854	A6	2-Jan	80199	A2	A9854	23598	300	USD	100	30,000
	//CN_AtlasTrading80200A9854	A6	7-Jan	80200	A2	A9854	23598	200	USD	100	20,000
	//CN_AtlasTrading80201A9854	A6	10-Jan	80201	A2	A9854	23598	150	USD	100	15,000
	//CN_AtlasTrading80202A9854	A6	19-Jan	80202	A2	A9854	23598	200	USD	100	20,000
	//CN_AtlasTrading80203A9854	A6	20-Jan	80203	A2	A9854	23598	180	USD	100	18,000
	inv := queryInvoice(stub,"1354651","A6908")
	if inv[0].Buyer != "A4" {
		t.Errorf("Failed to match and update invoice-> Buyer should be A4 is %s", inv[0].Buyer)
	}

}

// Test case 3
//CN_AtlasAmericas546568A6910
// A1
// A4
// 546568
// A6910
// 25412
// 125
// USD
// 400
// 50,000
// error
// Price exceed PO price
// 22-Jan
// PO Corrected with new price
func TestMatching_InvPriceExceedsPOPrice(t *testing.T) {
	loadDataset(t)
	//PO:
	//"CN_AtlasUSAA9854	A2	PO	A9854	A6	23598	1,000	USD	100	 $100,000.00 	NTE"
	//Inv:
	//FabricKey	Seller	Date	Ref	Buyer	PO #	SKU	Qty	Curr	Unit cost	Amount
	//CN_AtlasTrading80199A9854	A6	2-Jan	80199	A2	A9854	23598	300	USD	100	30,000
	//CN_AtlasTrading80200A9854	A6	7-Jan	80200	A2	A9854	23598	200	USD	100	20,000
	//CN_AtlasTrading80201A9854	A6	10-Jan	80201	A2	A9854	23598	150	USD	100	15,000
	//CN_AtlasTrading80202A9854	A6	19-Jan	80202	A2	A9854	23598	200	USD	100	20,000
	//CN_AtlasTrading80203A9854	A6	20-Jan	80203	A2	A9854	23598	180	USD	100	18,000
	item := queryPurchaseOrder(stub, "A6910")
	fmt.Printf("\n%-v\n", item)
	if item[0].UnitCost != 400.0 {
		t.Errorf("Failed to update purchase order with new Unit cost should be 400.0 is %f", item[0].UnitCost)
	}
}

// Test case 4
//CN_AtlasTrading56546A691000
//A6
//A4
//56546
//A691000
//23598
//150
//USD
//100
//15,000
//error
//Invalid PO #
//18-Jan
//PO# Corrected
func TestMatching_InvMatchAndUpdatePONum(t *testing.T) {
	loadDataset(t)
	//PO:
	//CN_AtlasGlobalA6909	A4	PO	A6909	A6	23598	150	USD	100	15000.00	STD				//"CN_AtlasUSAA9854	A2	PO	A9854	A6	23598	1,000	USD	100	 $100,000.00 	NTE"
	//Inv:
	//FabricKey	Seller	Date	Ref	Buyer	PO #	SKU	Qty	Curr	Unit cost	Amount
	//CN_AtlasTrading56546A691000	A6	3-Jan	56546	A4	A691000	23598	150	USD	100	15,000
	item := queryInvoice(stub,"56546","A691000")
	if item[0].PoNumber != "A6909" {
		t.Errorf("Failed to update Invoice PONum. Should be A6909 is %s", item[0].PoNumber)
	}
}

// Test case 5
//CN_AtlasUSA1354651A5686
//A2
//A6
//1354651
//A5686
//654864
//100
//USD
//200
//20,000
//error
//Invalid CPT
// Invoice remains in err as external invoice issued as I/C invoice
func TestMatching_InvMatch(t *testing.T) {
	loadDataset(t)
	//PO:
	//CN_AtlasGlobalA6909	A4	PO	A6909	A6	23598	150	USD	100	15000.00	STD				//"CN_AtlasUSAA9854	A2	PO	A9854	A6	23598	1,000	USD	100	 $100,000.00 	NTE"
	//Inv:
	//FabricKey	Seller	Date	Ref	Buyer	PO #	SKU	Qty	Curr	Unit cost	Amount
	//CN_AtlasUSA1354651A5686	A2	3-Jan	1354651	A6	A5686	654864	100	USD	200	20,000
	item := queryInvoice(stub, "1354651","A5686")
	if strings.Contains(item[0].State, "Ok") {
		t.Errorf("Failed to Invoice should be in error unmatched state is %s", item[0].State)
	}
}

// Test case 6
//CN_AtlasAmericas4684A69879
// A1
// A5
// 4684
// A69879
// 65468746
// 100
// USD
// 180
// 18,000
// error
// Invalid PO #
// Invoice remain in err, reason under investigation

func TestMatching_InvMismatchExternal(t *testing.T) {
	loadDataset(t)
	//PO:
	//CN_AtlasGlobalA6909	A4	PO	A6909	A6	23598	150	USD	100	15000.00	STD				//"CN_AtlasUSAA9854	A2	PO	A9854	A6	23598	1,000	USD	100	 $100,000.00 	NTE"
	//Inv:
	//FabricKey	Seller	Date	Ref	Buyer	PO #	SKU	Qty	Curr	Unit cost	Amount
	//CN_AtlasUSA1354651A5686	A2	3-Jan	1354651	A6	A5686	654864	100	USD	200	20,000
	item := queryInvoice(stub, "4684","A69879")

	if strings.Contains(item[0].State, "Ok") {
		t.Errorf("Failed to Invoice should be in error state is %s", item[0].State)
	}
}
func TestMatching_List_Unmatched(t *testing.T) {
	loadDataset(t)
	//PO:
	//CN_AtlasGlobalA6909	A4	PO	A6909	A6	23598	150	USD	100	15000.00	STD				//"CN_AtlasUSAA9854	A2	PO	A9854	A6	23598	1,000	USD	100	 $100,000.00 	NTE"
	//Inv:
	//FabricKey	Seller	Date	Ref	Buyer	PO #	SKU	Qty	Curr	Unit cost	Amount
	//CN_AtlasUSA1354651A5686	A2	3-Jan	1354651	A6	A5686	654864	100	USD	200	20,000
	item := queryUnmatched(stub)
	if len(item) != 2 {
		t.Errorf("unmatched should return 2 items not %d", len(item))
	}
}
