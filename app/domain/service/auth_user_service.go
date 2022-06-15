package service

import (
	"app/domain/model"
	"app/domain/repository"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v5"
)

// AuthUserService ドメインサービスとして利用し,複数のエンティティやレポジトリを扱う処理をここで実装する.
// AuthUserService interfase
type AuthUserService interface {
	AuthUser(c echo.Context) (*model.User, error)
}

type authUserService struct {
	repository.UserRepository
}

// NewAuthUserService AuthUserServiceを取得します.
func NewAuthUserService(r repository.UserRepository) AuthUserService {
	return &authUserService{r}
}

// token情報から認証中のログイン情報を取得する
func (u *authUserService) AuthUser(c echo.Context) (*model.User, error) {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*model.JwtCustomClaims)
	Id := claims.Id
	return u.UserRepository.FetchByUserID(Id)
}
