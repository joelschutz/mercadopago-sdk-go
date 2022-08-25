package service

import (
	"encoding/json"

	"github.com/joelschutz/mercadopago-sdk-go"
	model "github.com/joelschutz/mercadopago-sdk-go/models"
	"github.com/joelschutz/mercadopago-sdk-go/request"
)

// CreatePreference é o método responsável por criar um preference no MercadoPago.
func CreateAuthToken(authParams model.AuthParams, mercadoPagoAccessToken ...string) (*model.AuthResponse, *model.ErrorResponse, error) {

	params := request.Params{
		Method:      "POST",
		QueryParams: request.QueryParams(authParams),
		Headers:     map[string]interface{}{"Authorization": "Bearer " + mercadopago.GetAccessToken(mercadoPagoAccessToken...)},
		URL:         mercadopago.BASEURL + "/oauth/token",
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := mercadopago.ParseError(response.RawBody)
		return nil, resp, err
	}

	var preferenceResponse model.AuthResponse
	err = json.Unmarshal(response.RawBody, &preferenceResponse)
	return &preferenceResponse, nil, err
}
