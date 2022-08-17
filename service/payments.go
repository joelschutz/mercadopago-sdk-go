package service

import (
	"encoding/json"

	"github.com/joelschutz/mercadopago-sdk-go"
	model "github.com/joelschutz/mercadopago-sdk-go/models"
	"github.com/joelschutz/mercadopago-sdk-go/request"
)

// ConsultPayment é o método responsável consultar as informações atualizadas de um pagamento no MercadoPago, incluindo Status.
func ConsultPayment(paymentID string, mercadoPagoAccessToken ...string) (*model.PaymentResponse, *model.ErrorResponse, error) {

	params := request.Params{
		Method:     "GET",
		PathParams: request.PathParams{paymentID},
		Headers:    map[string]interface{}{"Authorization": "Bearer " + mercadopago.GetAccessToken(mercadoPagoAccessToken...)},
		URL:        mercadopago.BASEURL + "/v1/payments",
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := mercadopago.ParseError(response.RawBody)
		return nil, resp, err
	}

	var paymentResponse model.PaymentResponse
	err = json.Unmarshal(response.RawBody, &paymentResponse)
	return &paymentResponse, nil, err
}

func CreatePayment(paymentRequest model.PaymentRequest, mercadoPagoAccessToken ...string) (*model.PaymentResponse, *model.ErrorResponse, error) {

	params := request.Params{
		Method:  "POST",
		Body:    paymentRequest,
		Headers: map[string]interface{}{"Authorization": "Bearer " + mercadopago.GetAccessToken(mercadoPagoAccessToken...)},
		URL:     mercadopago.BASEURL + "/v1/payments",
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := mercadopago.ParseError(response.RawBody)
		return nil, resp, err
	}

	var paymentResponse model.PaymentResponse
	err = json.Unmarshal(response.RawBody, &paymentResponse)
	return &paymentResponse, nil, err
}
