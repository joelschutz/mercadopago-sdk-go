package model

import (
	"time"

	"github.com/joelschutz/mercadopago-sdk-go/request"
)

// PaymentResponse é a struct que é usada para receber os dados do request de novo pagamento do MercadoPago.
type PreferenceResponse struct {
	CollectorID   int      `json:"collector_id,omitempty"`   // Nosso ID do MercadoPago
	OperationType string   `json:"operation_type,omitempty"` // Tipo da operação (regular_payment, money_transfer)
	Items         []Item   `json:"items,omitempty"`          // Itens vendidos
	Payer         Payer    `json:"payer,omitempty"`          // Informações do pagador da cobrança
	BackUrls      BackUrls `json:"back_urls,omitempty"`      // URLS de redirecionamento
	// Indica se o comprador será redirecionado automaticamente para o back_urls após a compra
	// Use "approved" para redirecionar apenas no caso de sucesso
	// User "all" para todos os casos
	AutoReturn         string         `json:"auto_return,omitempty"`
	PaymentMethods     PaymentMethods `json:"payment_methods,omitempty"`      // Configurações das condições de pagamento do pagamento
	ClientID           string         `json:"client_id,omitempty"`            // ID do cliente do MercadoPago
	Marketplace        string         `json:"marketplace,omitempty"`          // Indica de qual marketplace foi feito pagamento (padrão NENHUM)
	MarketplaceFee     float64        `json:"marketplace_fee,omitempty"`      // Comissão de Mercado cobrada pelo proprietario do aplicativo
	Shipments          Shipments      `json:"shipments,omitempty"`            // Informações de envio dos itens
	NotificationURL    string         `json:"notification_url,omitempty"`     // URL do Webhook que é chamada quando o Status do pagamento é atualizado
	ExternalReference  string         `json:"external_reference,omitempty"`   // Nosso ID de controle interno
	AdditionalInfo     string         `json:"additional_info,omitempty"`      // Informações adicionais do pagamento
	Expires            bool           `json:"expires,omitempty"`              // Indica se o pagamento possui possui data de expiração
	DateOfExpiration   *time.Time     `json:"date_of_expiration,omitempty"`   // Data de expiração de meios de pagamento em dinheiro
	ExpirationDateFrom *time.Time     `json:"expiration_date_from,omitempty"` // A partir de qual data o pagamento estara ativo
	ExpirationDateTo   *time.Time     `json:"expiration_date_to,omitempty"`   // Até qual data o pagamento estara ativo
	DateCreated        time.Time      `json:"date_created,omitempty"`         // Data de criação do pagamento (gerado pelo MercadoPago)
	ID                 string         `json:"id,omitempty"`                   // ID do pagamento do MercadoPago (gerado pelo MercadoPago)
	InitPoint          string         `json:"init_point,omitempty"`           // Link de pagamento do pagamento
	SandboxInitPoint   string         `json:"sandbox_init_point,omitempty"`   // Link de pagamento de staging do pagamento
	SiteID             string         `json:"site_id,omitempty"`              // ID do site do pagamento
}

// PaymentRequest é a struct que é usada para fazer a request de um novo pagamento para o MercadoPago
type PreferenceRequest struct {
	ExternalReference string `json:"external_reference,omitempty"` // Nosso ID de controle interno
	Items             []Item `json:"items,omitempty"`              // Itens vendidos
	AdditionalInfo    string `json:"additional_info,omitempty"`    // Informações adicionais do pagamento
	// Indica se o comprador será redirecionado automaticamente para o back_urls após a compra
	// Use "approved" para redirecionar apenas no caso de sucesso
	// User "all" para todos os casos
	AutoReturn          string               `json:"auto_return,omitempty"`
	BackUrls            BackUrls             `json:"back_urls,omitempty"`            // URLS de redirecionamento
	DateOfExpiration    *time.Time           `json:"date_of_expiration,omitempty"`   // Data de expiração de meios de pagamento em dinheiro
	ExpirationDateFrom  *time.Time           `json:"expiration_date_from,omitempty"` // A partir de qual data o pagamento estara ativo
	ExpirationDateTo    *time.Time           `json:"expiration_date_to,omitempty"`   // Até qual data o pagamento estara ativo
	Expires             bool                 `json:"expires,omitempty"`              // Indica se o pagamento possui possui data de expiração
	DifferentialPricing *DifferentialPricing `json:"differential_pricing,omitempty"` // Configuração do preço diferenciado para este pagamento
	Marketplace         string               `json:"marketplace,omitempty"`          // Indica de qual marketplace foi feito pagamento (padrão NENHUM)
	MarketplaceFee      float64              `json:"marketplace_fee,omitempty"`      // Comissão de Mercado cobrada pelo proprietario do aplicativo
	ApplicationFee      float64              `json:"application_fee,omitempty"`      // Comissão de Mercado cobrada pelo proprietario do aplicativo
	NotificationURL     string               `json:"notification_url,omitempty"`     // URL do Webhook que é chamada quando o Status do pagamento é atualizado
	Payer               Payer                `json:"payer,omitempty"`                // Informações do pagador da cobrança
	PaymentMethods      PaymentMethods       `json:"payment_methods,omitempty"`      // Configurações das condições de pagamento do pagamento
	StatementDescriptor string               `json:"statement_descriptor,omitempty"` // Descrição do pagamento que ira aparecer no extrato do cartão
	Shipments           Shipments            `json:"shipments,omitempty"`            // Informações de envio dos itens
	Tracks              []Track              `json:"tracks,omitempty"`               // Lista trackeamentos que serão executados durante a interação do fluxo de pagamento
}

// PaymentSearchParams é um map[string]interface{} contém os filtros usados no método de Search de pagamentos
type PreferenceSearchParams request.QueryParams

// PreferenceSearchResponse é a struct que contém todas as informações que são retornadas pela API no método de Search de pagamentos
type PreferenceSearchResponse struct {
	NextOffset int                               `json:"next_offset,omitempty"` // Número de ínicio da próxima busca
	Total      int                               `json:"total,omitempty"`       // Total de itens encontrados na busca
	Elements   []PreferenceSearchElementResponse `json:"elements,omitempty"`    // Pagamentos retornados da busca
}

// PreferenceSearchElementResponse é a struct que contém toda as informações do pagamentos que são retornados no método de Search de pagamentos
type PreferenceSearchElementResponse struct {
	ShippingMode       string     `json:"shipping_mode,omitempty"`
	ID                 string     `json:"id,omitempty"`
	CollectorID        int        `json:"collector_id,omitempty"`
	CorporationID      string     `json:"corporation_id,omitempty"`
	ExternalReference  string     `json:"external_reference,omitempty"`
	PayerID            *string    `json:"payer_id,omitempty"`
	PayerEmail         string     `json:"payer_email,omitempty"`
	ProcessingModes    []string   `json:"processing_modes,omitempty"`
	ProductID          string     `json:"product_id,omitempty"`
	DateCreated        time.Time  `json:"date_created,omitempty"`
	ExpirationDateFrom *time.Time `json:"expiration_date_from,omitempty"`
	ExpirationDateTo   *time.Time `json:"expiration_date_to,omitempty"`
	Marketplace        string     `json:"marketplace,omitempty"`
	ClientID           string     `json:"client_id,omitempty"`
	SiteID             string     `json:"site_id,omitempty"`
	Expires            bool       `json:"expires,omitempty"`
	Items              []string   `json:"items,omitempty"`
	OperationType      string     `json:"operation_type,omitempty"`
}
