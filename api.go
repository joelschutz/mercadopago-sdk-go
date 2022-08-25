package mercadopago

import (
	"encoding/json"
	"os"

	model "github.com/joelschutz/mercadopago-sdk-go/models"
	"github.com/joelschutz/mercadopago-sdk-go/request"
)

const BASEURL = "https://api.mercadopago.com"

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// GetIdentificationTypes é o método responsável retornar todos o tipos de documento de identificação do MercadoPago.
func GetIdentificationTypes(mercadoPagoAccessToken ...string) ([]model.IdentificationType, *model.ErrorResponse, error) {

	params := request.Params{
		Method:  "GET",
		Headers: map[string]interface{}{"Authorization": "Bearer " + GetAccessToken(mercadoPagoAccessToken...)},
		URL:     BASEURL + "/v1/identification_types",
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := ParseError(response.RawBody)
		return nil, resp, err
	}

	var identificationTypes []model.IdentificationType
	err = json.Unmarshal(response.RawBody, &identificationTypes)
	return identificationTypes, nil, err
}

// GetPaymentMethods é o método responsável retornar todos o tipos de documento de identificação do MercadoPago.
func GetPaymentMethods(mercadoPagoAccessToken ...string) ([]model.PaymentMethod, *model.ErrorResponse, error) {

	params := request.Params{
		Method:  "GET",
		Headers: map[string]interface{}{"Authorization": "Bearer " + GetAccessToken(mercadoPagoAccessToken...)},
		URL:     BASEURL + "/v1/payment_methods",
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := ParseError(response.RawBody)
		return nil, resp, err
	}

	var paymentMethods []model.PaymentMethod
	err = json.Unmarshal(response.RawBody, &paymentMethods)
	return paymentMethods, nil, err
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// ParseError é a função que pega os dados do erro do MercadoPago e retorna em formato de Struct.
func ParseError(body []byte) (*model.ErrorResponse, error) {
	var errResponse model.ErrorResponse
	if err := json.Unmarshal(body, &errResponse); err != nil {
		return nil, err
	}
	return &errResponse, nil
}

// GetAccessToken é a função responsável por retornar o AccessToken do mercado pago.
// Caso tenha sido passado um token por parametro pegamos o token passado por parametro, se não pegamos da variavel de ambiente MERCADO_PAGO_ACCESS_TOKEN.
func GetAccessToken(mercadoPagoAccessToken ...string) string {
	if len(mercadoPagoAccessToken) > 0 {
		return mercadoPagoAccessToken[0]
	} else {
		return os.Getenv("MERCADO_PAGO_ACCESS_TOKEN")
	}
}
