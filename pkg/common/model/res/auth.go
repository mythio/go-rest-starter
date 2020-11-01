package res

import "github.com/mythio/go-rest-starter/pkg/common/model"

// AuthUser holds user and access_token
type AuthUser struct {
	User        model.User `json:"user"`
	AccessToken string     `json:"access_token"`
}
