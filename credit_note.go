package ubl

import "encoding/xml"

type CreditNote struct {
	XMLName xml.Name   `xml:"CreditNote"`
	Xmlns   []xml.Attr `xml:"-"`

	UBLVersionID                string                 `xml:"cbc:UBLVersionID"`
	CustomizationID             string                 `xml:"cbc:CustomizationID"`
	ProfileID                   string                 `xml:"cbc:ProfileID"`
	ID                          string                 `xml:"cbc:ID"`
	IssueDate                   string                 `xml:"cbc:IssueDate"`
	// CreditNoteTypeCode          int                    `xml:"cbc:CreditNoteTypeCode"`
	DocumentCurrencyCode        string                 `xml:"cbc:DocumentCurrencyCode"`
	BuyerReference              string                 `xml:"cbc:BuyerReference,omitempty"`
	// OrderReference              string                 `xml:"cac:OrderReference>cbc:ID,omitempty"`
	AdditionalDocumentReference []xmlDocumentReference `xml:"cac:AdditionalDocumentReference,omitempty"`
	AccountingSupplierParty     xmlSupplierParty       `xml:"cac:AccountingSupplierParty"`
	AccountingCustomerParty     xmlCustomerParty       `xml:"cac:AccountingCustomerParty"`
	// PaymentMeans                *xmlPaymentMeans       `xml:"cac:PaymentMeans,omitempty"`
	// PaymentTerms                *xmlPaymentTerms       `xml:"cac:PaymentTerms,omitempty"`
	TaxTotal                    xmlTaxTotal            `xml:"cac:TaxTotal"`
	LegalMonetaryTotal          xmlMonetaryTotal       `xml:"cac:LegalMonetaryTotal"`
	CreditNoteLine              []xmlCreditNoteLine    `xml:"cac:CreditNoteLine"`
}

func (i CreditNote) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "CreditNote"}

	for _, ns := range i.Xmlns {
		start.Attr = append(start.Attr, ns)
	}

	type alias CreditNote
	a := alias(i)
	return e.EncodeElement(a, start)
}

type xmlCreditNoteLine struct {
	ID                  string      `xml:"cbc:ID"`
	CreditedQuantity    xmlQuantity `xml:"cbc:CreditedQuantity"`
	LineExtensionAmount xmlAmount   `xml:"cbc:LineExtensionAmount"`
	TaxTotal            xmlTaxTotal `xml:"cac:TaxTotal"`
	Item                xmlItem     `xml:"cac:Item"`
	Price               xmlPrice    `xml:"cac:Price"`
}
