package mercadopago

import (
	"os"
	"testing"
)

// Definindo a variavel de ambiente MERCADO_PAGO_ACCESS_TOKEN que é usada pelo SDK
func init() {
	os.Setenv("MERCADO_PAGO_ACCESS_TOKEN", "TEST-1234936262199689-021433-a3e53345eecx2de47d25336123c21fc1-144567999")
}

// Testando geração de um pagamento
func TestSuccessOnCreatePreference(t *testing.T) {

	response, mercadopagoErr, err := CreatePreference(PreferenceRequest{
		ExternalReference: "test-00001",
		Items: []Item{
			{
				Title:     "Pagamendo mensalidade PagueTry",
				Quantity:  1,
				UnitPrice: 50,
			},
		},
		Payer: Payer{
			Identification: PayerIdentification{
				Type:   "CPF",
				Number: "12345678909",
			},
			Name:    "Rannielli Cruz",
			Surname: "Montagna",
			Email:   "raniellimontagna@hotmail.com",
		},
	}, "TEST-1234936262199689-021433-a3e53345eecx2de47d25336123c21fc1-144567999")

	if err != nil {
		t.Error("Erro inesperado!")
		t.Error(err.Error())

	} else if mercadopagoErr != nil {
		t.Error("Erro não tratado MercadoPago!")
		t.Error(mercadopagoErr.Message)
		t.Error(mercadopagoErr.Status)
		t.Error(mercadopagoErr.Error)

	} else {
		t.Log(response.InitPoint) // Sucesso!
	}

}

// Testando geração de um pagamento
func TestSuccessOnCreatePreferenceWithTokenParam(t *testing.T) {

	response, mercadopagoErr, err := CreatePreference(PreferenceRequest{
		ExternalReference: "test-00001",
		Items: []Item{
			{
				Title:     "Pagamendo mensalidade PagueTry",
				Quantity:  1,
				UnitPrice: 50,
			},
		},
		Payer: Payer{
			Identification: PayerIdentification{
				Type:   "CPF",
				Number: "12345678909",
			},
			Name:    "Rannielli Cruz",
			Surname: "Montagna",
			Email:   "raniellimontagna@hotmail.com",
		},
	})

	if err != nil {
		t.Error("Erro inesperado!")
		t.Error(err.Error())

	} else if mercadopagoErr != nil {
		t.Error("Erro não tratado MercadoPago!")
		t.Error(mercadopagoErr.Message)
		t.Error(mercadopagoErr.Status)
		t.Error(mercadopagoErr.Error)

	} else {
		t.Log(response.InitPoint) // Sucesso!
	}

}

// Testando tratamento de erro na geração de um pagamento (não informando campo obrigatório)
func TestFieldErrorOnCreatePreference(t *testing.T) {

	response, mercadopagoErr, err := CreatePreference(PreferenceRequest{
		ExternalReference: "test-00002",
		Items: []Item{
			{
				Title: "Pagamendo mensalidade PagueTry",
				// Não iremos informar o preço do item e a quantidade do item que são 2 campos obrigatórios
			},
		},
		Payer: Payer{
			Identification: PayerIdentification{
				Type:   "CPF",
				Number: "12345678909",
			},
			Name:    "Rannielli Cruz",
			Surname: "Montagna",
			Email:   "raniellimontagna@hotmail.com",
		},
	})

	if err != nil {
		t.Error("Erro inesperado!")
		t.Error(err.Error())

	} else if mercadopagoErr != nil {
		t.Log("Erro caputado com sucesso!") // Sucesso
		t.Log(mercadopagoErr.Message)
		t.Log(mercadopagoErr.Status)
		t.Log(mercadopagoErr.Error)

	} else {
		t.Error("Erro não capturado!")
		t.Error(response)
	}

}

// Testando atualização de um pagamento
func TestSuccessOnUpdatePreference(t *testing.T) {

	response, mercadopagoErr, err := UpdatePreference("825927174-5423394f-06f1-4d2b-8545-35ebecf70008", PreferenceRequest{
		ExternalReference: "test-00001",
		Items: []Item{
			{
				Title:     "Pagamendo semestralidade PagueTry",
				Quantity:  1,
				UnitPrice: 300,
			},
		},
		Payer: Payer{
			Identification: PayerIdentification{
				Type:   "CPF",
				Number: "12345678909",
			},
			Name:    "Matus",
			Surname: "Serafa",
			Email:   "mateus.silva@hotmail.com",
		},
	})

	if err != nil {
		t.Error("Erro inesperado!")
		t.Error(err.Error())

	} else if mercadopagoErr != nil {
		t.Error("Erro não tratado MercadoPago!")
		t.Error(mercadopagoErr.Message)
		t.Error(mercadopagoErr.Status)
		t.Error(mercadopagoErr.Error)

	} else {
		t.Log(response.InitPoint) // Sucesso!
	}

}

// Testando tratamento de erro na atualização de um pagamento (não informando campo obrigatório)
func TestFieldErrorOnUpdatePreference(t *testing.T) {

	response, mercadopagoErr, err := UpdatePreference("825927174-5423394f-06f1-4d2b-8545-35ebecf70008", PreferenceRequest{
		ExternalReference: "test-00001",
		Items: []Item{
			{
				Title: "Pagamendo mensalidade PagueTry",
				// Não iremos informar o preço do item e a quantidade do item que são 2 campos obrigatórios
			},
		},
		Payer: Payer{
			Identification: PayerIdentification{
				Type:   "CPF",
				Number: "12345678909",
			},
			Name:    "Rannielli Cruz",
			Surname: "Montagna",
			Email:   "raniellimontagna@hotmail.com",
		},
	})

	if err != nil {
		t.Error("Erro inesperado!")
		t.Error(err.Error())

	} else if mercadopagoErr != nil {
		t.Log("Erro caputado com sucesso!") // Sucesso
		t.Log(mercadopagoErr.Message)
		t.Log(mercadopagoErr.Status)
		t.Log(mercadopagoErr.Error)

	} else {
		t.Error("Erro não capturado!")
		t.Error(response)
	}

}

// Testando consulta de um pagamento
func TestSuccessOnGetPreference(t *testing.T) {

	response, mercadopagoErr, err := GetPreference("825927174-5423394f-06f1-4d2b-8545-35ebecf70008")

	if err != nil {
		t.Error("Erro inesperado!")
		t.Error(err.Error())

	} else if mercadopagoErr != nil {
		t.Error("Erro não tratado MercadoPago!")
		t.Error(mercadopagoErr.Message)
		t.Error(mercadopagoErr.Status)
		t.Error(mercadopagoErr.Error)

	} else {
		t.Log(response.InitPoint) // Sucesso!
	}

}

// Testando erro na consulta de um pagamento (pagamento inexistente)
func TestErrorOnGetPreference(t *testing.T) {

	response, mercadopagoErr, err := GetPreference("test-inexistente")

	if err != nil {
		t.Error("Erro inesperado!")
		t.Error(err.Error())

	} else if mercadopagoErr != nil {
		t.Log("Erro caputado com sucesso!") // Sucesso
		t.Log(mercadopagoErr.Message)
		t.Log(mercadopagoErr.Status)
		t.Log(mercadopagoErr.Error)

	} else {
		t.Error("Erro não capturado!")
		t.Error(response)
	}

}

// Testando busca de um pagamento atráves do filtro external_reference
func TestSuccessOnSearchPreferences(t *testing.T) {

	response, mercadopagoErr, err := SearchPreferences(PreferenceSearchParams{"external_reference": "test-00001"})

	if err != nil {
		t.Error("Erro inesperado!")
		t.Error(err.Error())

	} else if mercadopagoErr != nil {
		t.Error("Erro não tratado MercadoPago!")
		t.Error(mercadopagoErr.Message)
		t.Error(mercadopagoErr.Status)
		t.Error(mercadopagoErr.Error)

	} else {
		t.Log(response) // Sucesso!
	}

}

// Testando busca dos tipos de documento de identificação
func TestSuccessOnGetIdentificationTypes(t *testing.T) {

	identificationTypes, mercadopagoErr, err := GetIdentificationTypes()

	if err != nil {
		t.Error("Erro inesperado!")
		t.Error(err.Error())

	} else if mercadopagoErr != nil {
		t.Error("Erro não tratado MercadoPago!")
		t.Error(mercadopagoErr.Message)
		t.Error(mercadopagoErr.Status)
		t.Error(mercadopagoErr.Error)

	} else {
		t.Log(identificationTypes) // Sucesso!
	}

}

// Testando busca dos meios de pagamento
func TestSuccessOnGetPaymentMethods(t *testing.T) {

	paymentMethods, mercadopagoErr, err := GetPaymentMethods()

	if err != nil {
		t.Error("Erro inesperado!")
		t.Error(err.Error())

	} else if mercadopagoErr != nil {
		t.Error("Erro não tratado MercadoPago!")
		t.Error(mercadopagoErr.Message)
		t.Error(mercadopagoErr.Status)
		t.Error(mercadopagoErr.Error)

	} else {
		t.Log(paymentMethods) // Sucesso!
	}

}

// Testando consulta de situação de um pagamento
func TestErrorOnConsultStatusPayment(t *testing.T) {

	response, mercadopagoErr, err := ConsultPayment("1241420907")

	if err != nil {
		t.Error("Erro inesperado!")
		t.Error(err.Error())

	} else if mercadopagoErr != nil {
		t.Error("Erro não tratado MercadoPago!")
		t.Error(mercadopagoErr.Message)
		t.Error(mercadopagoErr.Status)
		t.Error(mercadopagoErr.Error)

	} else {
		t.Log(response)
	}

}
