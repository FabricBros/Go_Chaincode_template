package main

import (
	"bytes"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"encoding/json"
)

type Invoice struct {
	BillerBankCurrencyCode            string `json:"biller_Bank_Currency_code"`
	BuyerBankCurrencyCode             string `json:"buyer_Bank_currency_Code"`
	FromCoCo                          string `json:"from_CoCo"`
	FromCoCoName                      string `json:"from_CoCo_Name"`
	FromCoCoFullAddress               string `json:"from_CoCo_full_Address"`
	InvoiceCurrencyCode               string `json:"invoice_Currency_Code"`
	InvoiceDate                       string `json:"invoice_Date"`
	InvoiceLoadDate                   string `json:"invoice_Load_Date"`
	InvoiceNettingID                  string `json:"invoice_Netting_ID"`
	InvoiceNumber                     string `json:"invoice_Number"`
	InvoicePaymentAmount              string `json:"invoice_Payment_Amount"`
	InvoiceWHTAmountInInvoiceCurrency string `json:"invoice_WHT_Amount_in_Invoice_currency"`
	InvoiceWHTDeducted                string `json:"invoice_WHT_deducted"`
	InvoiceAmountInBillerCoCoCurrency string `json:"invoice_amount_in_Biller_CoCo_Currency"`
	InvoiceAmountInBuyerCoCoCurrency  string `json:"invoice_amount_in_Buyer_CoCo_currency"`
	InvoiceBuyerActionRequired        string `json:"invoice_buyer_action_required"`
	InvoiceBuyerActionUserID          string `json:"invoice_buyer_action_user_ID"`
	InvoiceHeaderLevelFreeFormText1   string `json:"invoice_header_level_free_form_text_1"`
	InvoiceHeaderLevelFreeFormText2   string `json:"invoice_header_level_free_form_text_2"`
	InvoiceLevelFreightCharges        string `json:"invoice_level_freight_charges"`
	InvoiceLevelIndirectTaxCharges    string `json:"invoice_level_indirect_tax_charges"`
	InvoiceLevelOtherTaxCharges       string `json:"invoice_level_other_tax_charges"`
	InvoiceLevelShippingCharges       string `json:"invoice_level_shipping_charges"`
	InvoiceLineLevel []struct {
		Description        string `json:"description"`
		LineNoInvoice      string `json:"line_No_invoice"`
		LineItemTax1       string `json:"line_item_Tax_1"`
		LineItemTax2       string `json:"line_item_Tax_2"`
		LineItemTax3       string `json:"line_item_Tax_3"`
		LineItemTax4       string `json:"line_item_Tax_4"`
		LineItemTax5       string `json:"line_item_Tax_5"`
		LineItemTaxAmount1 string `json:"line_item_Tax_Amount_1_"`
		LineItemTaxAmount2 string `json:"line_item_Tax_Amount_2"`
		LineItemTaxAmount3 string `json:"line_item_Tax_Amount_3"`
		LineItemTaxAmount4 string `json:"line_item_Tax_Amount_4"`
		LineItemTaxAmount5 string `json:"line_item_Tax_Amount_5"`
		LineItemTotalPrice string `json:"line_item_total_Price"`
		MeasurementUnit    string `json:"measurement_Unit"`
		PricePerUnit       string `json:"price_per_Unit"`
		ProductCode        string `json:"product_Code"`
		Quantity           string `json:"quantity"`
		TotalPrice         string `json:"total_Price"`
	} `json:"invoice_line_level"`
	InvoiceNetted                        string `json:"invoice_netted"`
	InvoicePaymentAmountInBillerCurrency string `json:"invoice_payment_amount_in_biller_currency"`
	InvoicePaymentAmountInBuyerCurrency  string `json:"invoice_payment_amount_in_buyer_currency"`
	InvoiceSellerActionID                string `json:"invoice_seller_action_ID"`
	InvoiceSellerActionRequired          string `json:"invoice_seller_action_required"`
	InvoiceSettledDate                   string `json:"invoice_settled_date"`
	InvoiceType                          string `json:"invoice_type"`
	LbInvoiceRef                         string `json:"lb_Invoice_Ref"`
	LbPayID                              string `json:"lb_Pay_ID"`
	LbPaymentType                        string `json:"lb_Payment_Type"`
	NoOfLineItems                        string `json:"no_of_line_items"`
	PoNo                                 string `json:"po_no"`
	SoNo                                 string `json:"so_no"`
	TaxType1                             string `json:"tax_type_1"`
	TaxType1Charges                      string `json:"tax_type_1_charges"`
	TaxType2                             string `json:"tax_type_2"`
	TaxType2Charges                      string `json:"tax_type_2_charges"`
	TaxType3                             string `json:"tax_type_3"`
	TaxType3Charges                      string `json:"tax_type_3_charges"`
	TaxType4                             string `json:"tax_type_4"`
	TaxType4Charges                      string `json:"tax_type_4_charges"`
	TaxType5                             string `json:"tax_type_5"`
	TaxType5Charges                      string `json:"tax_type_5_charges"`
	ToCoCo                               string `json:"to_CoCo"`
	ToCoCoName                           string `json:"to_CoCo_Name"`
	ToCoCoFullAddress                    string `json:"to_CoCo_full_Address"`
	TotalInvoiceAmount                   string `json:"total_Invoice_Amount"`
	TotalInvoiceLevelCharges             string `json:"total_invoice_level_charges"`
	TotalOfInvoiceLineAmount             string `json:"total_of_invoice_line_amount"`
}



// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) addInvoice(stub shim.ChaincodeStubInterface, args []string) pb.Response {
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
func (t *SimpleChaincode) updateInvoice(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	logger.Debug("Enter updateInvoice")
	defer logger.Debug("Exited updateInvoice")
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
