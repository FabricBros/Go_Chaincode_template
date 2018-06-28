package main

import (
	"bytes"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"encoding/json"
)

type Invoice struct {
	ObjectType			string
	Uuid				string `json: uuid`
	InvoiceNumber       string `json:"Invoice_Number"`
	FromCoCo            string `json:"From_CoCo"`
	ToCoCo              string `json:"To_CoCo"`
	FromCoCoName        string `json:"From_CoCo_Name"`
	ToCoCoName          string `json:"To_CoCo_Name"`
	FromCoCoFullAddress string `json:"From_CoCo_full_Address"`
	ToCoCoFullAddress   string `json:"To_CoCo_full_Address"`
	InvoiceDate         string `json:"Invoice_Date"`
	InvoiceLoadDate     string `json:"Invoice_Load_Date"`
	InvoiceSettledDate  string `json:"Invoice_settled_date"`
	LBInvoiceRef        string `json:"LB_Invoice_Ref"`
	///MMDDYYA1B2C3D4
	InvoiceHeaderLevelFreeFormText1   string `json:"Invoice_header_level_free_form_text_1"`
	InvoiceHeaderLevelFreeFormText2   string `json:"Invoice_header_level_free_form_text_2"`
	InvoiceCurrencyCode               string `json:"Invoice_Currency_Code"`
	TotalInvoiceAmount                string `json:"Total_Invoice_Amount"`
	InvoiceAmountInBillerCoCoCurrency string `json:"Invoice_amount_in_Biller_CoCo_Currency"`
	InvoiceAmountInBuyerCoCoCurrency  string `json:"Invoice_amount_in_Buyer_CoCo_currency"`
	BillerBankCurrencyCode            string `json:"Biller_Bank_Currency_code"`
	BuyerBankCurrencyCode             string `json:"Buyer_Bank_currency_Code"`
	InvoiceTypePOSO                   string `json:"Invoice_type_(PO/SO)"`
	PONo                              string `json:"PO_no"`
	SONo                              string `json:"SO_no"`
	NoOfLineItems                     string `json:"No_of_line_items"`
	TotalOfInvoiceLineAmount          string `json:"Total_of_invoice_line_amount"`
	TotalInvoiceLevelCharges          string `json:"Total_invoice_level_charges"`
	InvoiceLevelIndirectTaxCharges    string `json:"Invoice_level_indirect_tax_charges"`
	InvoiceLevelOtherTaxCharges       string `json:"Invoice_level_other_tax_charges"`
	InvoiceLevelFreightCharges        string `json:"Invoice_level_freight_charges"`
	InvoiceLevelShippingCharges       string `json:"Invoice_level_shipping_charges"`
	TaxType1                          string `json:"Tax_type_1"`
	TaxType1Charges                   string `json:"Tax_type_1_charges"`
	TaxType2                          string `json:"Tax_type_2"`
	TaxType2Charges                   string `json:"Tax_type_2_charges"`
	TaxType3                          string `json:"Tax_type_3"`
	TaxType3Charges                   string `json:"Tax_type_3_charges"`
	TaxType4                          string `json:"Tax_type_4"`
	TaxType4Charges                   string `json:"Tax_type_4_charges"`
	TaxType5                          string `json:"Tax_type_5"`
	TaxType5Charges                   string `json:"Tax_type_5_charges"`
	InvoiceBuyerActionRequired        string `json:"Invoice_buyer_action_required"`
	InvoiceBuyerActionUserID          string `json:"Invoice_buyer_action_user_ID"`
	InvoiceSellerActionRequired       string `json:"Invoice_seller_action_required"`
	InvoiceSellerActionID             string `json:"Invoice_seller_action_ID"`
	InvoiceWHTDeducted                string `json:"Invoice_WHT_deducted"`
	//Yes/No
	InvoiceWHTAmountInInvoiceCurrency    string `json:"Invoice_WHT_Amount_in_Invoice_currency"`
	InvoicePaymentAmount                 string `json:"Invoice_Payment_Amount"`
	InvoicePaymentAmountInBillerCurrency string `json:"Invoice_payment_amount_in_biller_currency"`
	InvoicePaymentAmountInBuyerCurrency  string `json:"Invoice_payment_amount_in_buyer_currency"`
	InvoiceNetted                        string `json:"Invoice_netted"`
	//Yes No
	InvoiceNettingID string `json:"Invoice_Netting_ID"`
	LBPaymentType    string `json:"LB_Payment_Type"`
	//ENUM [Wire/ FX/ NoWire]
	LBPayID string `json:"LB_Pay_ID"`
	LineLevel []struct {
		LineNo             string `json:"Line_No."`
		ProductCode        string `json:"Product_Code"`
		Description        string `json:"Description"`
		Quantity           string `json:"Quantity"`
		MeasurementUnit    string `json:"Measurement_Unit"`
		PricePerUnit       string `json:"Price_per_Unit"`
		TotalPrice         string `json:"Total_Price"`
		LineItemTax1       string `json:"Line_item_Tax_1"`
		LineItemTaxAmount1 string `json:"Line_item_Tax_Amount_1_"`
		LineItemTax2       string `json:"Line_item_Tax_2"`
		LineItemTaxAmount2 string `json:"Line_item_Tax_Amount_2"`
		LineItemTax3       string `json:"Line_item_Tax_3"`
		LineItemTaxAmount3 string `json:"Line_item_Tax_Amount_3"`
		LineItemTax4       string `json:"Line_item_Tax_4"`
		LineItemTaxAmount4 string `json:"Line_item_Tax_Amount_4"`
		LineItemTax5       string `json:"Line_item_Tax_5"`
		LineItemTaxAmount5 string `json:"Line_item_Tax_Amount_5"`
		LineItemTotalPrice string `json:"Line_item_total_Price"`
	} `json:"line_level"`
}



// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) addInvoices(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("adding invoice")
	defer logger.Debug("exit adding invoice")

	var invoices []Invoice

	err := json.Unmarshal([]byte(args[0]), &invoices)
	if err != nil {
		logger.Error("Error unmarshing invoice json:", err)
		return shim.Error(err.Error())
	}

	for _, v := range invoices{
		pk := v.InvoiceNumber
		vBytes, err := json.Marshal(v)

		if err!= nil{
			logger.Debug("error marshaling", err)
			return shim.Error(err.Error())
		}
		stub.PutState(pk, vBytes)
	}

	return shim.Success(nil)
}

// Deletes an entity from state
func (t *SimpleChaincode) getInvoice(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("enter get invoice")
	defer logger.Debug("exited get invoice")

	invoice_pk := args[0]
	var invoice Invoice

	invoiceByte, err := stub.GetState(invoice_pk)
	if err != nil{
		return shim.Error(err.Error())
	}
	err = json.Unmarshal(invoiceByte, &invoice)
	if err != nil{
		logger.Error(err)
		shim.Error(err.Error())
	}
	logger.Debug("getInvoice:")
	logger.Debug(invoice)
	var buffer bytes.Buffer
	buffer.WriteString("[")
	buffer.WriteString(string(invoiceByte))
	buffer.WriteString("]")

	return shim.Success([]byte(buffer.Bytes()))
}

// query callback representing the query of a chaincode
func (t *SimpleChaincode) updateInvoices(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("Enter updateInvoices")
	defer logger.Debug("Exited updateInvoices")
	var invoices []Invoice

	err := json.Unmarshal([]byte(args[0]), &invoices)
	if err != nil {
		logger.Error("Error unmarshing invoice json:", err)
		return shim.Error(err.Error())
	}

	for _, v := range invoices{
		pk := v.InvoiceNumber
		vBytes, err := json.Marshal(v)

		if err!= nil{
			logger.Debug("error marshaling", err)
			return shim.Error(err.Error())
		}
		stub.PutState(pk, vBytes)
	}

	return shim.Success(nil)
}
