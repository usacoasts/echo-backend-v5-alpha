package usecase

import (
	"app/domain/model"
	"app/domain/repository"
	"time"

	"github.com/golang-jwt/jwt"
)

// ApiUserTokenUseCase interfase
type ApiUserTokenUseCase interface {
	CreateToken(id int) (string, error)
	InsertToken(user *model.ApiUserToken) (*model.ApiUserToken, error)
}

type apiUserTokenUseCase struct {
	repository.ApiUserTokenRepository
}

// NewApiUserTokenUseCase ApiUserTokenUseCaseを取得します.
func NewApiUserTokenUseCase(r repository.ApiUserTokenRepository) ApiUserTokenUseCase {
	return &apiUserTokenUseCase{r}
}

func (u *apiUserTokenUseCase) CreateToken(id int) (string, error) {
	// Set custom claims
	claims := &model.JwtCustomClaims{
		id,
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	returnToken, err := token.SignedString([]byte("secret"))

	return returnToken, err
}

func (u *apiUserTokenUseCase) InsertToken(apiUserToken *model.ApiUserToken) (*model.ApiUserToken, error) {
	return u.ApiUserTokenRepository.Create(apiUserToken)
}
