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
	UserId       int64  `json:"user_id"`
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
	PublicKey    string `json:"public_key"`
	LiveMode     bool   `json:"live_mode"`
}
