package model

// BackUrls é a struct que contém as URLs que são utilizadas para redicionar o usuário após a pagamentor ser realizado ou após acontecer algum erro
type BackUrls struct {
	Success string `json:"success"` // URL de redirecionamento para pagamentos aprovados (exemplo pix)
	Pending string `json:"pending"` // URL de redirecionamento para pagamentos que estão com o pagamento pendente (exemplo boleto)
	Failure string `json:"failure"` // URL de redirecionamento para pagamentos que falharão (exemplo cartão de crédito)
}

// DifferentialPricing é a struct que contém o identificador único da configuração de preço diferenciado que sera aplicado ao pagamento
type DifferentialPricing struct {
	ID string `json:"id"` // Identificador único da configuração de preço diferenciado
}

// Payer é a struct que é usada para identificar para quem é o pagamento.
type Payer struct {
	Identification PayerIdentification `json:"identification"` // Documento de identificação do pagador
	Address        PayerAddress        `json:"address"`        // Endereço do pagador
	Email          string              `json:"email"`          // Email do pagador
	FirstName      string              `json:"first_name"`     // Nome do pagador (não colocar caracteres especiais)
	LastName       string              `json:"last_name"`      // Sobrenome do pagador (não colocar caracteres especiais)
}

// PayerIdentification é a struct que contém as informações de identificação do pagador
type PayerIdentification struct {
	Type   string `json:"type"`   // Tipo do documento de identificação (use CPF ou CNPJ)
	Number string `json:"number"` // Número do documentação de identificação
}

// PayerAddress é a struct que contém as informações do endereço do pagador
type PayerAddress struct {
	ZipCode      string  `json:"zip_code"`      // CEP do endereço do pagador
	StreetName   string  `json:"street_name"`   // Nome da rua do endereço do pagador
	StreetNumber *string `json:"street_number"` // Número do endereço do pagador
}

// PaymentMethods é a struct que contém as informações das configurações do pagamento
type PaymentMethods struct {
	ExcludedPaymentMethods []PaymentMethodID `json:"excluded_payment_methods"`  // Lista dos meios de pagamentos bloqueados para esse pagamento
	ExcludedPaymentTypes   []PaymentMethodID `json:"excluded_payment_types"`    // Lista dos tipos de pagamento bloqueados para esse pagamento
	DefaultPaymentMethodID *string           `json:"default_payment_method_id"` // Método de pagamento padrão (preferido, preferencial) para esse pagamento
	Installments           *int              `json:"installments"`              // Número máximo de parcelas
	DefaultInstallments    *int              `json:"default_installments"`      // Número máximo de parcelas padrão
}

// PaymentMethodID é a struct que contém o ID do método de pagamento
type PaymentMethodID struct {
	ID string `json:"id"` // Identificador do método de pagamento do MercadoPago
}

// Item é a struct que contém as informações dos itens que serão cobrados no pagamento
type Item struct {
	ID          string  `json:"id"`                    // Identificador interno nosso de controle
	Title       string  `json:"title"`                 // Titulo do item que é exibido na hora do pagamento
	Description string  `json:"description"`           // Descrição do item (não é exibido na hora do pagamento)
	PictureURL  string  `json:"picture_url,omitempty"` // Imagem do item (imagem que fica nas OG MetaTags do HTML)
	CategoryID  string  `json:"category_id"`           // Identificador da categoria interno nosso de controle
	Quantity    float64 `json:"quantity"`              // Quantidade do item vendido (obrigatório)
	CurrencyID  string  `json:"currency_id"`           // Identificador universal da moeda que será usada no pagamento no formato ISO-4217
	UnitPrice   float64 `json:"unit_price"`            // Preço unitário do item vendido (obrigatório)
}

// Shipments é a struct que contém as informações de entrega/envio dos itens do pagamento
type Shipments struct {
	Mode                  string   `json:"mode"`                    // Modo de envio (custom, me2, not_specified)
	LocalPickup           bool     `json:"local_pickup"`            // Preferência de remoção dos pacotes naagência (somente me2)
	Dimensions            string   `json:"dimensions"`              // Tamanho do pacote em cm x cm x cm (somente me2)
	DefaultShippingMethod int      `json:"default_shipping_method"` // Método de envio padrão no _checkout_ (somente me2)
	FreeMethos            []int    `json:"free_methods"`            // IDs dos métodos de envio com frete grátis (somente me2)
	Cost                  float64  `json:"cost"`                    // Custo do frete (somente custom)
	FreeShipping          bool     `json:"free_shipping"`           // Preferência por frete grátis (somente custom)
	ReceiverAddress       *Address `json:"receiver_address"`        // Endereço de envio
}

// Address é a struct que contém as informações do endereço de envio dos itens de um pagamento
type Address struct {
	ZipCode      string `json:"zip_codigo"`   // CEP do endereço de envio
	StreetName   string `json:"street_name"`  // Nome da rua do endereço de envio
	CityName     string `json:"city_name"`    // Nome da cidade do endereço de envio
	StateName    string `json:"state_name"`   // Nome do estado do endereço do envio
	StreetNumber *int   `json:"stree_number"` // Número do endereço de envio
	Floor        string `json:"floor"`        // Número do andar
	Apartment    string `json:"apartment"`    // Número do apartamento
}

// Track é a struct que contém as informações dos trackeamentos que serão executados durante o fluxo de pagamento
type Track struct {
	Type   string `json:"type"`  // Tipo do trackamento (google_ad ou facebook_ad)
	Values string `json:"value"` // Valores de configuração de acordo com o tipo do track (para mais informações consulte a documentação oficial)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// IdentificationType é a struct que contém as informações de um tipo de documento de identificação
type IdentificationType struct {
	ID        string `json:"id"`         // ID único do tipo de documento de identificação
	Name      string `json:"name"`       // Nome do tipo de documento de identificação
	Type      string `json:"type"`       // Tipo do dado do documento de identificação
	MinLength int    `json:"min_length"` // Tamanho mínimo do documento de identificação
	MaxLength int    `json:"max_length"` // Tamanho máximo do documento de identificação
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// PaymentMethod é a struct que contém as informações de um método de pagamento
type PaymentMethod struct {
	ID                    string                               `json:"id"`                     // ID único do método de pagamento
	Name                  string                               `json:"name"`                   // Nome do método de pagamento
	PaymentTypeID         string                               `json:"payment_type_id"`        // Tipo do meio de pagamento (ticket, atm, credit_card, debit_card, prepaid_card)
	Status                string                               `json:"status"`                 // Status do meio de pagamento (active, deactive, temporally_deactive)
	SecureThumbnail       string                               `json:"secure_thumbnail"`       // Logo do método de pagamento que deve ser mostrada em sites seguros
	Thumbnail             string                               `json:"thumbnail"`              // Logo do método de pagamento
	DeferredCapture       string                               `json:"deferred_capture"`       // Indica se a captura pode ser lenta ou não
	Settings              []PaymentMethodSettings              `json:"settings"`               // Configurações do método de pagamento
	AdditionalInfoNeeded  []string                             `json:"additional_info_needed"` // Lista de informações que devem ser fornecidas pelo pagador
	MinAllowedAmount      float64                              `json:"min_allowed_amount"`     // Mínimo valor que pode ser processado com este meio de pagamento
	MaxAllowedAmount      int                                  `json:"max_allowed_amount"`     // Máxilo valor que pode ser processado com este meio de pagamento
	AccreditationTime     int                                  `json:"accreditation_time"`     // Tempo de processamento do pagamento
	FinancialInstitutions []PaymentMethodFinancialInstitutions `json:"financial_institutions"` // Instituições financeiras de processamento do meio de pagamento
	ProcessingModes       []string                             `json:"processing_modes"`       // Modos de processamento
}

// PaymentMethodSettings é a struct que contém as informações das configurações de um método de pagamento
type PaymentMethodSettings struct {
	Bin          PaymentMethodSettingsBin          `json:"bin"`
	CardNumber   PaymentMethodSettingsCardNumber   `json:"card_number"`   // Informações de configuração do cartão de crédito
	SecurityCode PaymentMethodSettingsSecurityCode `json:"security_code"` // Informações de configuração do código de segurança
}

type PaymentMethodSettingsBin struct {
	Pattern             string `json:"pattern"`              // Expressão regular, representando bines aceitos
	ExclusionPattern    string `json:"exclusion_pattern"`    // Expressão regular, representando bines excluídos
	InstallmentsPattern string `json:"installments_pattern"` // Expressão regular, representando bines aceitos para pagar com mais de uma cota
}

// PaymentMethodSettingsCardNumber é a struct que contém as informações das configurações do cartão de crédito
type PaymentMethodSettingsCardNumber struct {
	Length     int    `json:"length"`     // Comprimento do núemro do cartão de crédito
	Validation string `json:"validation"` // Validação do número do cartão de crédito (standart, none)
}

// PaymentMethodSettingsSecurityCode é a struct que contém as informações do código de seguraça
type PaymentMethodSettingsSecurityCode struct {
	Mode         string `json:"mode"`          // Indica se o código de segurança é obrigatório ou opcional (mandatory, optional)
	Length       int    `json:"length"`        // Comprimento do código de segurança
	CardLocation string `json:"card_location"` // Localização do código de segurança do cartão de crédito (back, front)
}

// PaymentMethodFinancialInstitutions é a struct que contém as informações das instituições financeiras que processamento o método de pagamento
type PaymentMethodFinancialInstitutions struct {
	ID          string `json:"id"`          // ID único da instituição financeira (exemplo atm)
	Description string `json:"description"` // Nome descritivo da instituição financeira
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// WebhookResponse é a struct que é usada para receber os dados do request que o MercadoPago faz para o nosso webhook.
type WebhookResponse struct {
	ID          int              `json:"id"`           // ID único da resposta do Webhook
	LiveMode    bool             `json:"live_mode"`    // ???
	Type        string           `json:"type"`         // Tipo do evento, exemplo: payment
	DateCreated string           `json:"date_created"` // Data de criação do Webhook
	UserID      string           `json:"user_id"`      // Nosso ID do MercadoPago
	APIVersion  string           `json:"api_version"`  // Versão da API, exemplo: v1
	Action      string           `json:"action"`       // Ação do evento, exemplo: payment.created
	Data        WebhookPaymentID `json:"data"`         // Struct com o ID do pagamento
}

// FeeDetails é a struct que contém as informações sobre a taxa que foi cobrada sobre o pagamento
type FeeDetails struct {
	Amount   float64 `json:"amount"`    // Porcentagem da taxa que foi paga
	FeePayer string  `json:"fee_payer"` // Indica que ira pagar a taxa
	Type     string  `json:"type"`      // Indica quem esta cobrando a taxa
}

// TransactionDetails é a struct que contém as informações sobre os detalhes da transação
type TransactionDetails struct {
	ExternalResourceURL      *string `json:"external_resource_url"`       // URL do Boleto ou do PEC caso a forma de pagamento for boleto ou lotéricas
	FinancialInstitution     *string `json:"financial_institution"`       // Instituição financeira responsavel pelo pagamento
	TotalPaidAmount          float64 `json:"total_paid_amount"`           // Valor total pago
	InstallmentAmount        float64 `json:"installment_amount"`          // Valor do pagamento / Valor da parcela
	NetReceivedAmount        float64 `json:"net_received_amount"`         // Valor liquido recebido com o valor descontado das taxas
	OverpaidAmount           float64 `json:"overpaid_amount"`             // Valor pago em excesso ???
	PayableDeferralPeriod    string  `json:"payable_deferral_period"`     // ????
	PaymentMethodReferenceID string  `json:"payment_method_reference_id"` // ID de referencia do método de pagamento
	TransactionID            *string `json:"transaction_id"`              // ID da transação
}

// Barcode é a struct que contém o código de barras do boleto
type Barcode struct {
	Content string `json:"content"` // Código de barras do boleto
}

// PointOfInteraction é a struct que contém as informações dos dados da transação no caso do pagamento.
// No caso de ser PIX possui a chave do PIX no TransactionData.
// No caso de ser cartão de crédito ou outras formas de pagamento possui algumas outras informações irrelevantes.
type PointOfInteraction struct {
	TransactionData *TransactionData `json:"transaction_data"` // Informações do QRCode
}

// TransactionData é a struct que contém as informações do Base64 do QRCode e a chave Pix Copia-e-Cola
type TransactionData struct {
	QrCode       string `json:"qr_code"`        // Chave Pix Copia-e-Cola
	QrCodeBase64 string `json:"qr_code_base64"` // Base64 Do QRCode do Pix
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// ErrorResponse é a struct que é usada para receber os retornos de erro do MercadoPago.
type ErrorResponse struct {
	Error   string `json:"error"`   // Slug do erro que retornou
	Message string `json:"message"` // Mensagem de erro relacinada ao campo
	Status  int    `json:"status"`  // Mensagem de erro relacinada ao campo
}
