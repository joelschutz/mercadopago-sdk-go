package service

import (
	"encoding/json"

	"github.com/joelschutz/mercadopago-sdk-go"
	model "github.com/joelschutz/mercadopago-sdk-go/models"
	"github.com/joelschutz/mercadopago-sdk-go/request"
)

// CreatePreference é o método responsável por criar um preference no MercadoPago.
func CreatePreference(preferenceRequest model.PreferenceRequest, mercadoPagoAccessToken ...string) (*model.PreferenceResponse, *model.ErrorResponse, error) {

	params := request.Params{
		Method:  "POST",
		Body:    preferenceRequest,
		Headers: map[string]interface{}{"Authorization": "Bearer " + mercadopago.GetAccessToken(mercadoPagoAccessToken...)},
		URL:     mercadopago.BASEURL + "/checkout/preferences",
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := mercadopago.ParseError(response.RawBody)
		return nil, resp, err
	}

	var preferenceResponse model.PreferenceResponse
	err = json.Unmarshal(response.RawBody, &preferenceResponse)
	return &preferenceResponse, nil, err
}

// UpdatePreference é o método responsável por atualizar as informações de um pagamento no MercadoPago.
func UpdatePreference(preferenceID string, preferenceRequest model.PreferenceRequest, mercadoPagoAccessToken ...string) (*model.PreferenceResponse, *model.ErrorResponse, error) {

	params := request.Params{
		Method:     "PUT",
		PathParams: request.PathParams{preferenceID},
		Body:       preferenceRequest,
		Headers:    map[string]interface{}{"Authorization": "Bearer " + mercadopago.GetAccessToken(mercadoPagoAccessToken...)},
		URL:        mercadopago.BASEURL + "/checkout/preferences",
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := mercadopago.ParseError(response.RawBody)
		return nil, resp, err
	}

	var preferenceResponse model.PreferenceResponse
	err = json.Unmarshal(response.RawBody, &preferenceResponse)
	return &preferenceResponse, nil, err
}

// GetPreference é o método responsável buscar todas as informações de um pagamento no MercadoPago.
func GetPreference(preferenceID string, mercadoPagoAccessToken ...string) (*model.PreferenceResponse, *model.ErrorResponse, error) {

	params := request.Params{
		Method:     "GET",
		PathParams: request.PathParams{preferenceID},
		Headers:    map[string]interface{}{"Authorization": "Bearer " + mercadopago.GetAccessToken(mercadoPagoAccessToken...)},
		URL:        mercadopago.BASEURL + "/checkout/preferences",
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := mercadopago.ParseError(response.RawBody)
		return nil, resp, err
	}

	var preferenceResponse model.PreferenceResponse
	err = json.Unmarshal(response.RawBody, &preferenceResponse)
	return &preferenceResponse, nil, err
}

// SearchPreferences é o método responsável buscar todas as informações de um pagamento no MercadoPago.
// Como não existe nenhuma documentação completa sobre como esse EndPoint funciona então ele recebe por parametro qualquer filtro.
// Segundo oque consta nos SDKs oficiais e alguns não oficiais do MercadoPago, esse EndPoint é baseado em "Criteria Filters", ou seja,
// você pode filtrar por qualquer campo do pagamento usando qualquer operador, exemplo external_reference=525.
func SearchPreferences(searchParams model.PreferenceSearchParams, mercadoPagoAccessToken ...string) (*model.PreferenceSearchResponse, *model.ErrorResponse, error) {

	params := request.Params{
		Method:      "GET",
		QueryParams: request.QueryParams(searchParams),
		Headers:     map[string]interface{}{"Authorization": "Bearer " + mercadopago.GetAccessToken(mercadoPagoAccessToken...)},
		URL:         mercadopago.BASEURL + "/checkout/preferences/search",
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := mercadopago.ParseError(response.RawBody)
		return nil, resp, err
	}

	var preferenceSearchResponse model.PreferenceSearchResponse
	err = json.Unmarshal(response.RawBody, &preferenceSearchResponse)
	return &preferenceSearchResponse, nil, err
}
