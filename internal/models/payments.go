package model

import (
	"time"
)

type PaymentResponse struct {
	CollectorID   int      `json:"collector_id"`   // Nosso ID do MercadoPago
	OperationType string   `json:"operation_type"` // Tipo da operação (regular_payment, money_transfer)
	Items         []Item   `json:"items"`          // Itens vendidos
	Payer         Payer    `json:"payer"`          // Informações do pagador da cobrança
	BackUrls      BackUrls `json:"back_urls"`      // URLS de redirecionamento
	// Indica se o comprador será redirecionado automaticamente para o back_urls após a compra
	// Use "approved" para redirecionar apenas no caso de sucesso
	// User "all" para todos os casos
	AutoReturn         string         `json:"auto_return"`
	PaymentMethods     PaymentMethods `json:"payment_methods"`      // Configurações das condições de pagamento do pagamento
	ClientID           string         `json:"client_id"`            // ID do cliente do MercadoPago
	Marketplace        string         `json:"marketplace"`          // Indica de qual marketplace foi feito pagamento (padrão NENHUM)
	MarketplaceFee     float64        `json:"marketplace_fee"`      // Comissão de Mercado cobrada pelo proprietario do aplicativo
	Shipments          Shipments      `json:"shipments"`            // Informações de envio dos itens
	NotificationURL    string         `json:"notification_url"`     // URL do Webhook que é chamada quando o Status do pagamento é atualizado
	ExternalReference  string         `json:"external_reference"`   // Nosso ID de controle interno
	AdditionalInfo     string         `json:"additional_info"`      // Informações adicionais do pagamento
	Expires            bool           `json:"expires"`              // Indica se o pagamento possui possui data de expiração
	DateOfExpiration   *time.Time     `json:"date_of_expiration"`   // Data de expiração de meios de pagamento em dinheiro
	ExpirationDateFrom *time.Time     `json:"expiration_date_from"` // A partir de qual data o pagamento estara ativo
	ExpirationDateTo   *time.Time     `json:"expiration_date_to"`   // Até qual data o pagamento estara ativo
	DateCreated        time.Time      `json:"date_created"`         // Data de criação do pagamento (gerado pelo MercadoPago)
	ID                 string         `json:"id"`                   // ID do pagamento do MercadoPago (gerado pelo MercadoPago)
	InitPoint          string         `json:"init_point"`           // Link de pagamento do pagamento
	SandboxInitPoint   string         `json:"sandbox_init_point"`   // Link de pagamento de staging do pagamento
	SiteID             string         `json:"site_id"`              // ID do site do pagamento
}

// PaymentRequest é a struct que é usada para fazer a request de um novo pagamento para o MercadoPago
type PaymentRequest struct {
	ExternalReference string `json:"external_reference"` // Nosso ID de controle interno
	Items             []Item `json:"items"`              // Itens vendidos
	AdditionalInfo    string `json:"additional_info"`    // Informações adicionais do pagamento
	// Indica se o comprador será redirecionado automaticamente para o back_urls após a compra
	// Use "approved" para redirecionar apenas no caso de sucesso
	// User "all" para todos os casos
	AutoReturn          string               `json:"auto_return"`
	BackUrls            BackUrls             `json:"back_urls"`                      // URLS de redirecionamento
	DateOfExpiration    *time.Time           `json:"date_of_expiration"`             // Data de expiração de meios de pagamento em dinheiro
	ExpirationDateFrom  *time.Time           `json:"expiration_date_from"`           // A partir de qual data o pagamento estara ativo
	ExpirationDateTo    *time.Time           `json:"expiration_date_to"`             // Até qual data o pagamento estara ativo
	Expires             bool                 `json:"expires"`                        // Indica se o pagamento possui possui data de expiração
	DifferentialPricing *DifferentialPricing `json:"differential_pricing,omitempty"` // Configuração do preço diferenciado para este pagamento
	Marketplace         string               `json:"marketplace"`                    // Indica de qual marketplace foi feito pagamento (padrão NENHUM)
	MarketplaceFee      float64              `json:"marketplace_fee"`                // Comissão de Mercado cobrada pelo proprietario do aplicativo
	ApplicationFee      float64              `json:"application_fee"`                // Comissão de Mercado cobrada pelo proprietario do aplicativo
	NotificationURL     string               `json:"notification_url"`               // URL do Webhook que é chamada quando o Status do pagamento é atualizado
	Payer               Payer                `json:"payer"`                          // Informações do pagador da cobrança
	PaymentMethods      PaymentMethods       `json:"payment_methods"`                // Configurações das condições de pagamento do pagamento
	StatementDescriptor string               `json:"statement_descriptor,omitempty"` // Descrição do pagamento que ira aparecer no extrato do cartão
	Shipments           Shipments            `json:"shipments"`                      // Informações de envio dos itens
	Tracks              []Track              `json:"tracks,omitempty"`               // Lista trackeamentos que serão executados durante a interação do fluxo de pagamento
}

// PaymentConsultResponse é a struct que contém as informações do retorno da Consulta de um pagamento.
type PaymentConsultResponse struct {
	AdditionalInfo            PaymentConsultAdditionalInfo `json:"additional_info"`                // Informações adicionais do pagamento
	AuthorizationCode         *string                      `json:"authorization_code"`             // Código de autorização do pagamento
	BinaryMode                bool                         `json:"binary_mode"`                    // Indica se é o modo binaria de pagamento ou não
	Captured                  bool                         `json:"captured"`                       // Indica se o pagamento foi capturado ou não ???
	Card                      PaymentConsultCard           `json:"card"`                           // Informações do cartão de crédito do pagamento
	CollectorID               int                          `json:"collector_id"`                   // Nosso ID do MercadoPago
	CurrencyID                string                       `json:"currency_id"`                    // Identificador universal da moeda que será usada no pagamento no formato ISO-4217
	DateApproved              *time.Time                   `json:"date_approved"`                  // Data da aprovação do pagamento
	DateCreated               time.Time                    `json:"date_created"`                   // Data da criação do pagamento
	DateLastUpdated           *time.Time                   `json:"date_last_updated"`              // Data da ultima atualização do pagamento
	DateOfExpiration          *time.Time                   `json:"date_of_expiration"`             // Data de expiração de meios de pagamento em dinheiro
	Description               string                       `json:"description"`                    // Descrição do pagamento
	DifferentialPricingId     string                       `json:"differential_pricing_id"`        // Identificador único da configuração de preço diferenciado
	ExternalReference         string                       `json:"external_reference"`             // Nosso ID de controle interno
	FeeDetails                []FeeDetails                 `json:"fee_details"`                    // Informações sobre as taxas que foram aplicadas sobre o pagamento
	ID                        int                          `json:"id"`                             // Identificador único do pagamento
	Installments              int                          `json:"installments"`                   // Número máximo de parcelas
	LiveMode                  bool                         `json:"live_mode"`                      // ???
	MoneyReleaseDate          *time.Time                   `json:"money_release_date"`             // Data da liberação do dinheiro na nossa conta do MercadoPago
	MoneyReleaseSchema        *string                      `json:"money_release_schema"`           // Esquema da liberação do dinheiro
	NotificationURL           string                       `json:"notification_url"`               // URL do Webhook que é chamada quando o Status do pagamento é atualizado
	OperationType             string                       `json:"operation_type"`                 // Tipo do pagamento (consulte a documentação para saber oque significa)
	Payer                     Payer                        `json:"payer"`                          // Informações do pagador
	PaymentMethodID           string                       `json:"payment_method_id"`              // ID do método de pagamento (exemplo: master)
	PaymentTypeID             string                       `json:"payment_type_id"`                // Tipo do método de pagamento (exemplo: credit_card)
	ProcessingModes           string                       `json:"processing_modes"`               // Modo de processamento
	StatementDescriptor       string                       `json:"statement_descriptor,omitempty"` // Descrição do pagamento que ira aparecer no extrato do cartão
	TransactionAmount         float64                      `json:"transaction_amount"`             // Valor pago
	TransactionAmountRefunded float64                      `json:"transaction_amount_refunded"`    // Valor do reembolso
	TransactionDetails        TransactionDetails           `json:"transaction_details"`            // Detalhes da transação

	// Código de barras do boleto
	// Nos testes que eu (Eduardo Mior) fiz, essa Struct só é retornada quando o pagamento é feito em boleto ou PEC(lotéricas), porém isso não esta documentado em nenhum lugar na documentação do MercadoPago
	Barcode *Barcode `json:"barcode"`

	// QRCode do Pix e Chave Copia-e-Cola do Pix
	// Nos testes que eu (Eduardo Mior) fiz, essa Struct contém a Chave Copia-e-Cola no caso e ser PIX e no caso e ser cartão de crédito ou outras formas de pagamento
	// contém outras informações inuteis, porém isso não esta documentado em nenhum lugar na documetação do MercadoPago
	PointOfInteraction *PointOfInteraction `json:"point_of_interaction"`

	// Status do pagamento (segundo documentação oficial do mercado pago)
	// approved - Pagamento aprovado
	// pending - Pagamento pendente
	// authorized - Pagamento autorizado porém ainda não aprovado
	// in_process - O pagamento esta sendo revisado pelo MercadoPago
	// in_mediation - Foi aberta uma disputa no pagamento e ele esta em revisão
	// rejected - Pagamento rejeitado (cartão de crédito)
	// cancelled - O pagamento foi cancelado por uma das partes ou porque ele expirou
	// refunded - O pagamento foi reembolsado para o usuário
	// charged_back - Foi feito um estorno do pagamento no cartão de crédito do usuário
	Status string `json:"status"`

	// Detalhes sobre o status do pagamento. A lista completa pode ser consultada em https://www.mercadopago.com.br/developers/pt/guides/online-payments/checkout-api/handling-responses
	// cc_rejected_card_disabled (cartão de crédito bloqueado)
	// cc_rejected_blacklist (cartão de crédito recusado)
	// cc_rejected_high_risk (cartão de crédito recusado)
	// rejected_high_risk (paypal recusado)
	// pending_review_manual (pagamendo sendo revisado pelo mercadopago)
	// pending_waiting_transfer (aguardando pagamento do pix / aguardando transferência do dinheiro)
	// pending_waiting_payment (aguardando pagamento do boleto ou do PEC(lotérica))
	// accredited (aprovado)
	// https://prnt.sc/1ta9w4b
	StatusDetail string `json:"status_detail"`
}

// PaymentConsultCard é a struct que contém as informações do cartão de crédito que efetuou o pagamento
type PaymentConsultCard struct {
	Cardholder      Cardholder `json:"cardholder"`        // Informações do dono do cartão
	DateCreated     *time.Time `json:"date_created"`      // Data de criação do cartão
	DateLastUpdated *time.Time `json:"date_last_updated"` // Data da ultima atualização do cartão
	ExpirationMonth int        `json:"expiration_month"`  // Mês de expiração do cartão
	ExpirationYear  int        `json:"expiration_year"`   // Ano de expiração do cartão
	FirstSixDigits  string     `json:"first_six_digits"`  // Seis primeiros digitos do cartão
	LastFourDigits  string     `json:"last_four_digits"`  // Ultimos quatro digitos do cartão
}

// PaymentConsultAdditionalInfo é a struct que contém informações adicionais sobre o pagamento
type PaymentConsultAdditionalInfo struct {
	IPAddress string `json:"ip_address"` // IP do usuário que pagou

	// Items     []Item `json:"items"`
	// Por algum motivo eles retoram a Quantity do item e o UnityPrice como String e isso faz com que aconteça erro no Unmarshal do JSON.
	// Por esse motivo foi comentado o campo Items das informações adicionais.
}
