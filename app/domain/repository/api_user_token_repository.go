package repository

import (
	"app/domain/model"
)

// ApiUserTokenRepository interface
type ApiUserTokenRepository interface {
	Create(apiUserToken *model.ApiUserToken) (*model.ApiUserToken, error)
}
