package invoicing

type InvoiceLine struct {
	Item  string
	Quantity int
	Amount float64
	TaxID string
}

type InvoiceDefinition struct {
	InvoiceDate string
	RecipientID string
	Notes string
	Lines []InvoiceLine
}

func NewInvoiceDefinition(date string, recipientID string, notes string) *InvoiceDefinition {
	return &InvoiceDefinition{ date, recipientID, notes, []InvoiceLine{} }
}

func (inv InvoiceDefinition) AddLine(item string, quantity int, amount float64, taxId string) {
	inv.Lines = append(inv.Lines, InvoiceLine{item,quantity,amount,taxId})
}

type InvoicingAPI interface {
	generateInvoice(invoice InvoiceDefinition) string
	sendInvoice(invoiceNumber string) bool
}