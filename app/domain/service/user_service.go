package service

import (
	"app/domain/repository"
	"context"
)

// UserService ドメインサービスとして利用し,複数のエンティティやレポジトリを扱う処理をここで実装する.
type UserService interface {
	DoSomething(ctx context.Context, foo int) error
}

type userService struct {
	repository.UserRepository
}

// NewUserService UserServiceを取得します.
func NewUserService(r repository.UserRepository) UserService {
	return &userService{r}
}

func (u *userService) DoSomething(ctx context.Context, foo int) error {
	// some code
	return nil
}
