package main

type EntityBankAcctTable struct {
	CoCo                      string `json:"CoCo"`
	BankAccountCode           string `json:"Bank_Account_Code"`
	BankAccountCurrency       string `json:"Bank_Account_Currency"`
	BankAccountTitle          string `json:"Bank_Account_title"`
	BankAccountNo             string `json:"Bank_account_no"`
	BankName                  string `json:"Bank_Name"`
	BankAddress               string `json:"Bank_Address"`
	EntityAddress             string `json:"Entity_Address"`
	SwiftBIC                  string `json:"Swift1BIC"`
	BankAccountContactName    string `json:"Bank_Account_contact_name"`
	BankAccountContactEmailID string `json:"Bank_Account_contact_email_ID"`
}

type EntityTable struct {
	GroupID  string `json:"Group_ID"`
	EntityID string `json:"Entity_ID"`
	CoCo     string `json:"CoCo"`
	//9DigitAlphaNumericWithFirst6DigitsBeingTheGroupIDAndLast3FromEntityID
	EntityName                                            string `json:"Entity_Name"`
	EntityCountryName                                     string `json:"Entity_Country_Name"`
	EntityCountryISOCode                                  string `json:"Entity_Country_ISO_Code"`
	EntityAddress1                                        string `json:"Entity_Address_1"`
	EntityAddress2                                        string `json:"Entity_Address_2"`
	City                                                  string `json:"City"`
	EntityState                                           string `json:"Entity_State"`
	EntityFunctionalCurrencyCode                          string `json:"Entity_functional_currency_code"`
	MultilateralNettingAllowedWithinCountry               string `json:"Multilateral_Netting_allowed_within_Country"`
	MultilateralNettingAllowedOutsideCountry              string `json:"Multilateral_Netting_allowed_outside_Country"`
	BiLateralNettingAllowedWithinCountry                  string `json:"Bi_lateral_Netting_allowed_within_country"`
	BiLateralNettingAllowedOutsideCountry                 string `json:"Bi_Lateral_Netting_allowed_outside_country"`
	PayMasterAllowedToPayWithinCountry                    string `json:"Pay_Master_allowed_to_pay_within_country"`
	IfYesPaymasterEntityID                                string `json:"If_Yes_Paymaster_Entity_ID"`
	PaymasterAllowedToPayOutsideCountry                   string `json:"Paymaster_allowed_to_pay_outside_country"`
	AllowedToReceiveFromPaymasterWithinCountry            string `json:"Allowed_to_receive_from_paymaster_within_country"`
	AllowedToReceiveFromPaymasterOutsideCountry           string `json:"Allowed_to_receive_from_paymaster_outside_country"`
	AdditonalApprovalRequiredToPayWithinCountry           string `json:"Additonal_approval_required_to_pay_within_country"`
	AdditionalApprovalRequiredToPayOutsideCountry         string `json:"Additional_approval_required_to_pay_outside_country"`
	AdditionalApprovalRequiredToReceiveFromOutsideCountry string `json:"Additional_approval_required_to_receive_from_outside_country"`
	InwardFXToBeConvertedOnshore                          string `json:"Inward_FX_to_be_converted_onshore"`
	OutwardFXCanBeTradedOffshore                          string `json:"Outward_FX_can_be_traded_offshore?"`
}

type GRN struct {
	GRNNo       string `json:"GRN_no"`
	ReceiptDate string `json:"Receipt_Date"`
	FromCoCo    string `json:"From_CoCo"`
	ToCoCo      string `json:"To_CoCo"`
	InvoiceNo   string `json:"Invoice_no"`
	InvoiceDate string `json:"Invoice_Date"`
	PONo        string `json:"PO_no"`
	PODate      string `json:"PO_Date"`
	LineLevel []struct {
		LineNo          string `json:"Line_No."`
		ProductCode     string `json:"Product_Code"`
		Description     string `json:"Description"`
		Quantity        string `json:"Quantity"`
		MeasurementUnit string `json:"Measurement_Unit"`
	} `json:"line_level"`
}

type GroupMaster struct {
	GroupName                 string `json:"Group_Name"`
	GroupID                   string `json:"Group_ID"`
	GroupContactPersonName    string `json:"Group_Contact_person_Name"`
	GroupContactPersonEmailID string `json:"Group_Contact_person_email_ID"`
	ParentEntityName          string `json:"Parent_Entity_Name"`
	GroupContactPersonAddress string `json:"Group_Contact_person_Address"`
}

type Inspection struct {
	ShipFromEntityCode string `json:"Ship_From_Entity_Code"`
	ShipToEntityCode   string `json:"Ship_To_Entity_Code"`
	ContactName        string `json:"Contact_Name"`
	GoodsInformation   string `json:"Goods_Information"`
	PONumber           string `json:"PO_Number"`
	SONumber           string `json:"SO_Number"`
	InvoiceNumber      string `json:"Invoice_Number"`
	SentDate           string `json:"Sent_Date"`
	ReceiptDate        string `json:"Receipt_Date"`
	TotalReceivedQty   string `json:"Total_Received_Qty"`
	TotalAcceptableQty string `json:"Total_Acceptable_Qty"`
	TotalRejectedQty   string `json:"Total_Rejected_Qty"`
	ReturnQty          string `json:"Return_Qty"`
	UnderTestingQty    string `json:"Under_Testing_Qty"`
	Comments           string `json:"Comments"`
	LineLevel []struct {
		Lineno          string `json:"Lineno"`
		ProductCode     string `json:"Product_Code"`
		Description     string `json:"Description"`
		QtyOrdered      string `json:"Qty_Ordered"`
		QtyShipped      string `json:"Qty_Shipped"`
		QtyReceived     string `json:"Qty_Received"`
		QtyAcceptable   string `json:"Qty_Acceptable"`
		QtyRejected     string `json:"Qty_Rejected"`
		QtyReturned     string `json:"Qty_Returned"`
		QtyUnderTesting string `json:"Qty_Under_Testing"`
		Comments        string `json:"Comments"`
	} `json:"line_level"`
}

type Invoices struct {
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


type UserTable struct {
	GroupID         string `json:"Group_ID"`
	UserID          string `json:"User_ID"`
	UserFirstName   string `json:"User_First_Name"`
	UserLastName    string `json:"User_Last_Name"`
	UserAddress     string `json:"User_Address"`
	UserEmailID     string `json:"User_email_ID"`
	UserContactNo   string `json:"User_Contact_no."`
	UserDesgination string `json:"User_desgination"`
	UserActiveDate  string `json:"User_active_date"`
	UserExpiryDate  string `json:"User_expiry_date"`
	UserStatus      string `json:"User_Status"`
	//	_(Active_/_Inactive)
}
