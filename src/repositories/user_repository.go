package repositories

import (
	"barebone-go-crud/src/models/entity"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (int64, error)
}
