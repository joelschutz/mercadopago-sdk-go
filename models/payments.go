package model

import (
	"time"
)

// type PaymentResponse struct {
// 	CollectorID   int      `json:"collector_id,omitempty"`   // Nosso ID do MercadoPago
// 	OperationType string   `json:"operation_type,omitempty"` // Tipo da operação (regular_payment, money_transfer)
// 	Items         []Item   `json:"items,omitempty"`          // Itens vendidos
// 	Payer         Payer    `json:"payer,omitempty"`          // Informações do pagador da cobrança
// 	BackUrls      BackUrls `json:"back_urls,omitempty"`      // URLS de redirecionamento
// 	// Indica se o comprador será redirecionado automaticamente para o back_urls após a compra
// 	// Use "approved" para redirecionar apenas no caso de sucesso
// 	// User "all" para todos os casos
// 	AutoReturn         string         `json:"auto_return,omitempty"`
// 	PaymentMethods     PaymentMethods `json:"payment_methods,omitempty"`      // Configurações das condições de pagamento do pagamento
// 	ClientID           string         `json:"client_id,omitempty"`            // ID do cliente do MercadoPago
// 	Marketplace        string         `json:"marketplace,omitempty"`          // Indica de qual marketplace foi feito pagamento (padrão NENHUM)
// 	ApplicationFee     float64        `json:"application_fee,omitempty"`     // Comissão de Mercado cobrada pelo proprietario do aplicativo
// 	Shipments          Shipments      `json:"shipments,omitempty"`            // Informações de envio dos itens
// 	NotificationURL    string         `json:"notification_url,omitempty"`     // URL do Webhook que é chamada quando o Status do pagamento é atualizado
// 	ExternalReference  string         `json:"external_reference,omitempty"`   // Nosso ID de controle interno
// 	AdditionalInfo     string         `json:"additional_info,omitempty"`      // Informações adicionais do pagamento
// 	Expires            bool           `json:"expires,omitempty"`              // Indica se o pagamento possui possui data de expiração
// 	DateOfExpiration   *time.Time     `json:"date_of_expiration,omitempty"`   // Data de expiração de meios de pagamento em dinheiro
// 	ExpirationDateFrom *time.Time     `json:"expiration_date_from,omitempty"` // A partir de qual data o pagamento estara ativo
// 	ExpirationDateTo   *time.Time     `json:"expiration_date_to,omitempty"`   // Até qual data o pagamento estara ativo
// 	DateCreated        time.Time      `json:"date_created,omitempty"`         // Data de criação do pagamento (gerado pelo MercadoPago)
// 	ID                 string         `json:"id,omitempty"`                   // ID do pagamento do MercadoPago (gerado pelo MercadoPago)
// 	InitPoint          string         `json:"init_point,omitempty"`           // Link de pagamento do pagamento
// 	SandboxInitPoint   string         `json:"sandbox_init_point,omitempty"`   // Link de pagamento de staging do pagamento
// 	SiteID             string         `json:"site_id,omitempty"`              // ID do site do pagamento
// }

// PaymentRequest é a struct que é usada para fazer a request de um novo pagamento para o MercadoPago
type PaymentRequest struct {
	AdditionalInfo        PaymentAdditionalInfo  `json:"additional_info,omitempty"` // Informações adicionais do pagamento
	ApplicationFee        float64                `json:"application_fee,omitempty"` // Comissão de Mercado cobrada pelo proprietario do aplicativo
	BinaryMode            bool                   `json:"binary_mode,omitempty"`     // Indica se é o modo binaria de pagamento ou não
	CallbackUrl           string                 `json:"callback_url,omitempty"`
	CampaignID            int                    `json:"campaign_id,omitempty"`             // Nosso ID do MercadoPago
	Capture               bool                   `json:"capture,omitempty"`                 // Indica se o pagamento foi capturado ou não ???
	CouponAmount          float64                `json:"coupon_amount,omitempty"`           // Valor pago
	CouponCode            string                 `json:"coupon_code,omitempty"`             // Valor pago
	DateOfExpiration      *time.Time             `json:"date_of_expiration,omitempty"`      // Data de expiração de meios de pagamento em dinheiro
	Description           string                 `json:"description,omitempty"`             // Descrição do pagamento
	DifferentialPricingId string                 `json:"differential_pricing_id,omitempty"` // Identificador único da configuração de preço diferenciado
	ExternalReference     string                 `json:"external_reference,omitempty"`      // Nosso ID de controle interno
	Installments          int                    `json:"installments,omitempty"`            // Número máximo de parcelas
	IssuerID              string                 `json:"issuer_id,omitempty"`               // Identificador universal da moeda que será usada no pagamento no formato ISO-4217
	Metadata              map[string]interface{} `json:"metadata,omitempty"`
	NotificationURL       string                 `json:"notification_url,omitempty"`     // URL do Webhook que é chamada quando o Status do pagamento é atualizado
	Payer                 Payer                  `json:"payer,omitempty"`                // Informações do pagador
	PaymentMethodID       string                 `json:"payment_method_id,omitempty"`    // ID do método de pagamento (exemplo: master)
	StatementDescriptor   string                 `json:"statement_descriptor,omitempty"` // Descrição do pagamento que ira aparecer no extrato do cartão
	CardToken             string                 `json:"token,omitempty"`
	TransactionAmount     float64                `json:"transaction_amount,omitempty"` // Valor pago
}

// PaymentResponse é a struct que contém as informações do retorno da Consulta de um pagamento.
type PaymentResponse struct {
	IssuerID                  string                 `json:"issuer_id,omitempty"`
	SponsorID                 string                 `json:"sponsor_id,omitempty"`
	TaxesAmount               int                    `json:"taxes_amount,omitempty"`
	ShippingAmount            int                    `json:"shipping_amount,omitempty"`
	Metadata                  map[string]interface{} `json:"metadata,omitempty"`
	CouponAmount              float64                `json:"coupon_amount,omitempty"`
	AdditionalInfo            PaymentAdditionalInfo  `json:"additional_info,omitempty"`    // Informações adicionais do pagamento
	AuthorizationCode         *string                `json:"authorization_code,omitempty"` // Código de autorização do pagamento
	BinaryMode                bool                   `json:"binary_mode,omitempty"`        // Indica se é o modo binaria de pagamento ou não
	Capture                   bool                   `json:"capture,omitempty"`            // Indica se o pagamento foi capturado ou não ???
	Card                      PaymentConsultCard     `json:"card,omitempty"`               // Informações do cartão de crédito do pagamento
	CollectorID               int                    `json:"collector_id,omitempty"`
	CurrencyID                string                 `json:"currency_id,omitempty"`                 // Identificador universal da moeda que será usada no pagamento no formato ISO-4217
	DateApproved              *time.Time             `json:"date_approved,omitempty"`               // Data da aprovação do pagamento
	DateCreated               time.Time              `json:"date_created,omitempty"`                // Data da criação do pagamento
	DateLastUpdated           *time.Time             `json:"date_last_updated,omitempty"`           // Data da ultima atualização do pagamento
	DateOfExpiration          *time.Time             `json:"date_of_expiration,omitempty"`          // Data de expiração de meios de pagamento em dinheiro
	Description               string                 `json:"description,omitempty"`                 // Descrição do pagamento
	DifferentialPricingId     string                 `json:"differential_pricing_id,omitempty"`     // Identificador único da configuração de preço diferenciado
	ExternalReference         string                 `json:"external_reference,omitempty"`          // Nosso ID de controle interno
	FeeDetails                []FeeDetails           `json:"fee_details,omitempty"`                 // Informações sobre as taxas que foram aplicadas sobre o pagamento
	ID                        int                    `json:"id,omitempty"`                          // Identificador único do pagamento
	Installments              int                    `json:"installments,omitempty"`                // Número máximo de parcelas
	LiveMode                  bool                   `json:"live_mode,omitempty"`                   // ???
	MoneyReleaseDate          *time.Time             `json:"money_release_date,omitempty"`          // Data da liberação do dinheiro na nossa conta do MercadoPago
	MoneyReleaseStatus        *string                `json:"money_release_status,omitempty"`        // Esquema da liberação do dinheiro
	MoneyReleaseSchema        *string                `json:"money_release_schema,omitempty"`        // Esquema da liberação do dinheiro
	NotificationURL           string                 `json:"notification_url,omitempty"`            // URL do Webhook que é chamada quando o Status do pagamento é atualizado
	OperationType             string                 `json:"operation_type,omitempty"`              // Tipo do pagamento (consulte a documentação para saber oque significa)
	Payer                     Payer                  `json:"payer,omitempty"`                       // Informações do pagador
	PaymentMethodID           string                 `json:"payment_method_id,omitempty"`           // ID do método de pagamento (exemplo: master)
	PaymentTypeID             string                 `json:"payment_type_id,omitempty"`             // Tipo do método de pagamento (exemplo: credit_card)
	ProcessingMode            string                 `json:"processing_mode,omitempty"`             // Modo de processamento
	StatementDescriptor       string                 `json:"statement_descriptor,omitempty"`        // Descrição do pagamento que ira aparecer no extrato do cartão
	TransactionAmount         float64                `json:"transaction_amount,omitempty"`          // Valor pago
	TransactionAmountRefunded float64                `json:"transaction_amount_refunded,omitempty"` // Valor do reembolso
	TransactionDetails        TransactionDetails     `json:"transaction_details,omitempty"`         // Detalhes da transação

	// Código de barras do boleto
	// Nos testes que eu (Eduardo Mior) fiz, essa Struct só é retornada quando o pagamento é feito em boleto ou PEC(lotéricas), porém isso não esta documentado em nenhum lugar na documentação do MercadoPago
	Barcode *Barcode `json:"barcode,omitempty"`

	// QRCode do Pix e Chave Copia-e-Cola do Pix
	// Nos testes que eu (Eduardo Mior) fiz, essa Struct contém a Chave Copia-e-Cola no caso e ser PIX e no caso e ser cartão de crédito ou outras formas de pagamento
	// contém outras informações inuteis, porém isso não esta documentado em nenhum lugar na documetação do MercadoPago
	PointOfInteraction *PointOfInteraction `json:"point_of_interaction,omitempty"`

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
	Status string `json:"status,omitempty"`

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
	StatusDetail string `json:"status_detail,omitempty"`
}

// PaymentConsultCard é a struct que contém as informações do cartão de crédito que efetuou o pagamento
type PaymentConsultCard struct {
	Cardholder      Cardholder `json:"cardholder,omitempty"`        // Informações do dono do cartão
	DateCreated     *time.Time `json:"date_created,omitempty"`      // Data de criação do cartão
	DateLastUpdated *time.Time `json:"date_last_updated,omitempty"` // Data da ultima atualização do cartão
	ExpirationMonth int        `json:"expiration_month,omitempty"`  // Mês de expiração do cartão
	ExpirationYear  int        `json:"expiration_year,omitempty"`   // Ano de expiração do cartão
	FirstSixDigits  string     `json:"first_six_digits,omitempty"`  // Seis primeiros digitos do cartão
	LastFourDigits  string     `json:"last_four_digits,omitempty"`  // Ultimos quatro digitos do cartão
}

// PaymentAdditionalInfo é a struct que contém informações adicionais sobre o pagamento
type PaymentAdditionalInfo struct {
	IPAddress string                         `json:"ip_address,omitempty"` // IP do usuário que pagou
	Items     []Item                         `json:"items,omitempty"`
	Payer     PaymentAdditionalInfoPayer     `json:"payer,omitempty"`     // Informações do pagador da cobrança
	Shipments PaymentAdditionalInfoShipments `json:"shipments,omitempty"` // Informações de envio dos itens

	// Items     []Item `json:"items,omitempty"`
	// Por algum motivo eles retoram a Quantity do item e o UnityPrice como String e isso faz com que aconteça erro no Unmarshal do JSON.
	// Por esse motivo foi comentado o campo Items das informações adicionais.
}

type PaymentAdditionalInfoPayer struct {
	Address   PayerAddress `json:"address,omitempty"`    // Endereço do pagador
	Phone     string       `json:"phone,omitempty"`      // Email do pagador
	FirstName string       `json:"first_name,omitempty"` // Nome do pagador (não colocar caracteres especiais)
	LastName  string       `json:"last_name,omitempty"`  // Sobrenome do pagador (não colocar caracteres especiais)
}

type PaymentAdditionalInfoShipments struct {
	ReceiverAddress *Address `json:"receiver_address,omitempty"` // Endereço de envio
}
