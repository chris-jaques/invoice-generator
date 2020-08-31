package wave

import (
	"invgen/config"
	"invgen/invoicing"
	"invgen/graphql"
	"fmt"
)

type WaveAPI struct {
	serverUrl string "https://gql.waveapps.com/graphql/public"
	accessToken string
	recipientId string
}

func NewWaveAPI(conf config.Config) *WaveAPI {
	c := conf.Wave
	return &WaveAPI{
		accessToken: c.AccessToken, 
		recipientId: c.RecipientId,
	}
}

type Invoice struct {
	invoiceNumber string
}

func (w WaveAPI) GenerateInvoice(date string, totalHours float64) string {
	
	invoice := invoicing.NewInvoiceDefinition(date, w.recipientId, "Test Invoice")
	client := graphql.NewClient(w.accessToken, w.serverUrl)

	inv := Invoice{}
	response := client.Query(&inv, nil)
	fmt.Printf("%s :: %s --> %s\n",invoice, client, response)
	return "" // invoice ID
}

func (w WaveAPI) SendInvoice(invoiceNumber string) bool {

	return true // success
}