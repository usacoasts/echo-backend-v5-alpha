package usecase

import (
	"app/domain/model"
	"app/domain/repository"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// UserUseCase interfase
type UserUseCase interface {
	GetUserId(id int) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	Login(ctx context.Context, user *model.User) (*model.User, error)
}

type userUseCase struct {
	repository.UserRepository
}

// NewUserUseCase UserUseCaseを取得します.
func NewUserUseCase(r repository.UserRepository) UserUseCase {
	return &userUseCase{r}
}

func (u *userUseCase) GetUserId(id int) (*model.User, error) {
	return u.UserRepository.FetchByUserID(id)
}

func (u *userUseCase) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	name := user.Name
	email := user.Email
	p := user.Password
	hashed, _ := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	password := string(hashed)

	return u.UserRepository.Create(&model.User{Name: name, Email: email, Password: password})
}

func (u *userUseCase) Login(ctx context.Context, user *model.User) (*model.User, error) {
	email := user.Email
	getUser := &model.User{}
	getUser, err := u.UserRepository.FistByEmail(email)

	if err := bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(user.Password)); err != nil {
		fmt.Println("error")
		fmt.Println("error")
		log.Fatal(err)
	}

	return getUser, err
}
