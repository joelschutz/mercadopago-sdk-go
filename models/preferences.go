package model

import (
	"time"

	"github.com/joelschutz/mercadopago-sdk-go/internal/request"
)

// PaymentResponse é a struct que é usada para receber os dados do request de novo pagamento do MercadoPago.
type PreferenceResponse struct {
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
type PreferenceRequest struct {
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

// PaymentSearchParams é um map[string]interface{} contém os filtros usados no método de Search de pagamentos
type PreferenceSearchParams request.QueryParams

// PreferenceSearchResponse é a struct que contém todas as informações que são retornadas pela API no método de Search de pagamentos
type PreferenceSearchResponse struct {
	NextOffset int                               `json:"next_offset"` // Número de ínicio da próxima busca
	Total      int                               `json:"total"`       // Total de itens encontrados na busca
	Elements   []PreferenceSearchElementResponse `json:"elements"`    // Pagamentos retornados da busca
}

// PreferenceSearchElementResponse é a struct que contém toda as informações do pagamentos que são retornados no método de Search de pagamentos
type PreferenceSearchElementResponse struct {
	ShippingMode       string     `json:"shipping_mode"`
	ID                 string     `json:"id"`
	CollectorID        int        `json:"collector_id"`
	CorporationID      string     `json:"corporation_id"`
	ExternalReference  string     `json:"external_reference"`
	PayerID            *string    `json:"payer_id"`
	PayerEmail         string     `json:"payer_email"`
	ProcessingModes    []string   `json:"processing_modes"`
	ProductID          string     `json:"product_id"`
	DateCreated        time.Time  `json:"date_created"`
	ExpirationDateFrom *time.Time `json:"expiration_date_from"`
	ExpirationDateTo   *time.Time `json:"expiration_date_to"`
	Marketplace        string     `json:"marketplace"`
	ClientID           string     `json:"client_id"`
	SiteID             string     `json:"site_id"`
	Expires            bool       `json:"expires"`
	Items              []string   `json:"items"`
	OperationType      string     `json:"operation_type"`
}
