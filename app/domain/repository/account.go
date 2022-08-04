package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Account interface {
	// Fetch account which has specified username
	FindByUsername(ctx context.Context, username string) (*object.Account, error)
    CreateAccount(ctx context.Context, username string, passwordHash string) (int64, error)	
	FindById(ctx context.Context, id int64) (*object.Account, error)
}
