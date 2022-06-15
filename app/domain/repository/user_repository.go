package repository

import (
	"app/domain/model"
	"context"
)

// UserRepository interface
type UserRepository interface {
	Fetch(ctx context.Context) ([]*model.User, error)
	FetchByID(ctx context.Context, id int) (*model.User, error)
	FetchByUserID(id int) (*model.User, error)
	FistByEmail(email string) (*model.User, error)
	Create(user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	Delete(ctx context.Context, id int) error
}
