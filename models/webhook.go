package model

// WebhookPaymentID é a struct que contém o ID do pagamento que foi feito pelo pagador final, esse ID deve ser usado para consultar as informações do pagamento.
type WebhookPaymentID struct {
	ID string `json:"id"` // ID do pagamento que é usado para consultar o status
}

// Cardholder é a struct que contém as informações do dono do cartão
type Cardholder struct {
	Identification PayerIdentification `json:"identification"` // Informações de identificação do cartão de crédito
	Name           string              `json:"name"`           // Nome do dono do cartão de crédito
}
