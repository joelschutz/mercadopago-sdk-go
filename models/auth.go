package model

import "github.com/joelschutz/mercadopago-sdk-go/request"

// Example Query
// map[string]interface{}{
// 	"client_id": "client_id"
// 	"client_secret": "client_secret"
// 	"grant_type": "grant_type"
// 	"code": "code"
// 	"redirect_uri": "redirect_uri"
// 	"refresh_token": "refresh_token"
// }
type AuthParams request.QueryParams

type AuthResponse struct {
	UserId       int64  `json:"user_id,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
	ExpiresIn    int64  `json:"expires_in,omitempty"`
	Scope        string `json:"scope,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	PublicKey    string `json:"public_key,omitempty"`
	LiveMode     bool   `json:"live_mode,omitempty"`
}
