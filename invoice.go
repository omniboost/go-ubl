package ubl

import (
	"encoding/xml"
)

type Invoice struct {
	XMLName xml.Name   `xml:"Invoice"`
	Xmlns   []xml.Attr `xml:"-"`

	UBLVersionID                string                 `xml:"cbc:UBLVersionID"`
	CustomizationID             string                 `xml:"cbc:CustomizationID"`
	ProfileID                   string                 `xml:"cbc:ProfileID"`
	ID                          string                 `xml:"cbc:ID"`
	IssueDate                   string                 `xml:"cbc:IssueDate"`
	DueDate                     string                 `xml:"cbc:DueDate,omitempty"`
	InvoiceTypeCode             int                    `xml:"cbc:InvoiceTypeCode"`
	DocumentCurrencyCode        string                 `xml:"cbc:DocumentCurrencyCode"`
	BuyerReference              string                 `xml:"cbc:BuyerReference,omitempty"`
	OrderReference              string                 `xml:"cac:OrderReference>cbc:ID"`
	AdditionalDocumentReference []xmlDocumentReference `xml:"cac:AdditionalDocumentReference"`
	AccountingSupplierParty     xmlSupplierParty       `xml:"cac:AccountingSupplierParty"`
	AccountingCustomerParty     xmlCustomerParty       `xml:"cac:AccountingCustomerParty"`
	PaymentMeans                *xmlPaymentMeans       `xml:"cac:PaymentMeans,omitempty"`
	PaymentTerms                *xmlPaymentTerms       `xml:"cac:PaymentTerms,omitempty"`
	TaxTotal                    xmlTaxTotal            `xml:"cac:TaxTotal"`
	LegalMonetaryTotal          xmlMonetaryTotal       `xml:"cac:LegalMonetaryTotal"`
	InvoiceLine                 []xmlInvoiceLine       `xml:"cac:InvoiceLine"`
}

func (i Invoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "Invoice"}

	for _, ns := range i.Xmlns {
		start.Attr = append(start.Attr, ns)
	}

	type alias Invoice
	a := alias(i)
	return e.EncodeElement(a, start)
}

type xmlDocumentReference struct {
	ID                  string          `xml:"cbc:ID"`
	DocumentDescription string          `xml:"cbc:DocumentDescription"`
	Attachment          []xmlAttachment `xml:"cac:Attachment"`
}

type xmlAttachment struct {
	EmbeddedDocumentBinaryObject xmlEmbeddedDocumentBinaryObject `xml:"cbc:EmbeddedDocumentBinaryObject"`
}

type xmlEmbeddedDocumentBinaryObject struct {
	Value    string `xml:",chardata"`
	MimeCode string `xml:"mimeCode,attr"`
	Filename string `xml:"filename,attr"`
}

type xmlSupplierParty struct {
	Party xmlParty `xml:"cac:Party"`
}

type xmlCustomerParty struct {
	SupplierAssignedAccountID string   `xml:"cbc:SupplierAssignedAccountID"`
	Party                     xmlParty `xml:"cac:Party"`
}

type xmlParty struct {
	PartyName struct {
		Name string `xml:"cbc:Name"`
	} `xml:"cac:PartyName"`
	PostalAddress  xmlPostalAddress   `xml:"cac:PostalAddress"`
	PartyTaxScheme *xmlPartyTaxScheme `xml:"cac:PartyTaxScheme,omitempty"`
	Contact        struct {
		ElectronicMail string `xml:"cbc:ElectronicMail"`
	} `xml:"cac:Contact,omitempty"`
}

type xmlPostalAddress struct {
	StreetName string     `xml:"cbc:StreetName"`
	CityName   string     `xml:"cbc:CityName"`
	PostalZone string     `xml:"cbc:PostalZone"`
	Country    xmlCountry `xml:"cac:Country"`
}

type xmlPartyTaxScheme struct {
	CompanyID string       `xml:"cbc:CompanyID"`
	TaxScheme xmlTaxScheme `xml:"cac:TaxScheme"`
}

type xmlCountry struct {
	IdentificationCode string `xml:"cbc:IdentificationCode"`
}

type xmlPaymentMeans struct {
	PaymentMeansCode   int    `xml:"cbc:PaymentMeansCode"`
	PaymentDueDate     string `xml:"cbc:PaymentDueDate"`
	PaymentChannelCode string `xml:"cbc:PaymentChannelCode"`
}

type xmlFinancialAccount struct {
	ID                         string                        `xml:"cbc:ID"`
	FinancialInstitutionBranch xmlFinancialInstitutionBranch `xml:"cac:FinancialInstitutionBranch"`
}

type xmlFinancialInstitutionBranch struct {
	ID string `xml:"cbc:ID"`
}

type xmlPaymentTerms struct {
	Note string `xml:"cbc:Note"`
}

type xmlTaxTotal struct {
	TaxAmount   xmlAmount        `xml:"cbc:TaxAmount"`
	TaxSubtotal []xmlTaxSubtotal `xml:"cac:TaxSubtotal"`
}

type xmlTaxSubtotal struct {
	TaxableAmount xmlAmount       `xml:"cbc:TaxableAmount"`
	TaxAmount     xmlAmount       `xml:"cbc:TaxAmount"`
	Percent       *float64         `xml:"cbc:Percent,omitempty"`
	TaxCategory   *xmlTaxCategory `xml:"cac:TaxCategory,omitempty"`
}

type xmlMonetaryTotal struct {
	LineExtensionAmount xmlAmount `xml:"cbc:LineExtensionAmount"`
	TaxExclusiveAmount  xmlAmount `xml:"cbc:TaxExclusiveAmount"`
	TaxInclusiveAmount  xmlAmount `xml:"cbc:TaxInclusiveAmount"`
	PayableAmount       xmlAmount `xml:"cbc:PayableAmount"`
}

type xmlAmount struct {
	Value      float64 `xml:",chardata"`
	CurrencyID string  `xml:"currencyID,attr"`
}

// Possible values for the unitcode:
// https://docs.peppol.eu/poacc/billing/3.0/codelist/UNECERec20/
type xmlQuantity struct {
	Value    float64 `xml:",chardata"`
	UnitCode string  `xml:"unitCode,attr"`
}

type xmlInvoiceLine struct {
	ID                  string      `xml:"cbc:ID"`
	InvoicedQuantity    xmlQuantity `xml:"cbc:InvoicedQuantity"`
	LineExtensionAmount xmlAmount   `xml:"cbc:LineExtensionAmount"`
	TaxTotal            xmlTaxTotal `xml:"cac:TaxTotal"`
	Item                xmlItem     `xml:"cac:Item"`
	Price               xmlPrice    `xml:"cac:Price"`
}

type xmlItem struct {
	Description           string         `xml:"cbc:Description"`
	Name                  string         `xml:"cbc:Name"`
	ClassifiedTaxCategory xmlTaxCategory `xml:"cac:ClassifiedTaxCategory"`
}

type xmlTaxCategory struct {
	ID        string       `xml:"cbc:ID,omitempty"`
	Name      string       `xml:"cbc:Name,omitempty"`
	Percent   *float64      `xml:"cbc:Percent,omitempty"`
	TaxScheme xmlTaxScheme `xml:"cac:TaxScheme,omitempty"`
}

type xmlTaxScheme struct {
	ID string `xml:"cbc:ID"`
}

type xmlPrice struct {
	PriceAmount xmlAmount `xml:"cbc:PriceAmount"`
}
