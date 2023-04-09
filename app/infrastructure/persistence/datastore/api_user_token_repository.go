package datastore

import (
	"app/domain/model"
	"app/domain/repository"
	"gorm.io/gorm"
)

type apiUserTokenRepository struct {
	Conn *gorm.DB
}

// NewApiUserTokenRepository ApiUserTokenRepositoryを取得します.
func NewApiUserTokenRepository(Conn *gorm.DB) repository.ApiUserTokenRepository {
	return &apiUserTokenRepository{Conn}
}

func (r *apiUserTokenRepository) Create(apiUserToken *model.ApiUserToken) (*model.ApiUserToken, error) {
	err := r.Conn.Create(apiUserToken).Error
	return apiUserToken, err
}
