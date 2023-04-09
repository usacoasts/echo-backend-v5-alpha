package datastore

import (
	"app/domain/model"
	"app/domain/repository"
	"context"
	"gorm.io/gorm"
)

type userRepository struct {
	Conn *gorm.DB
}

// NewUserRepository UserRepositoryを取得します.
func NewUserRepository(Conn *gorm.DB) repository.UserRepository {
	return &userRepository{Conn}
}

func (r *userRepository) Fetch(ctx context.Context) ([]*model.User, error) {
	var (
		users []*model.User
		err   error
	)
	err = r.Conn.Order("id desc").Find(&users).Error
	return users, err
}

func (r *userRepository) FetchByID(ctx context.Context, id int) (*model.User, error) {
	u := &model.User{ID: id}
	err := r.Conn.First(u).Error
	return u, err
}

func (r *userRepository) FetchByUserID(id int) (*model.User, error) {
	u := &model.User{ID: id}
	err := r.Conn.First(u).Error
	return u, err
}

func (r *userRepository) FistByEmail(email string) (*model.User, error) {
	u := &model.User{Email: email}
	err := r.Conn.Where("email = ?", email).First(&u).Error
	return u, err
}

func (r *userRepository) Create(u *model.User) (*model.User, error) {
	err := r.Conn.Create(u).Error
	return u, err
}

/**
func (r *userRepository) Update(ctx context.Context, u *model.User) (*model.User, error) {
	err := r.Conn.Model(u).Updates(u).Error
	return u, err
}*/

func (r *userRepository) Delete(ctx context.Context, id int) error {
	u := &model.User{ID: id}
	err := r.Conn.Delete(u).Error
	return err
}
