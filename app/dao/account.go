package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	account struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewAccount(db *sqlx.DB) repository.Account {
	return &account{db: db}
}

// FindByUsername : ユーザ名からユーザを取得
func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}

// ユーザー作成
func (r *account) CreateAccount(_ context.Context, username string, passwordHash string) (int64, error) {	
	ins, err := r.db.Preparex("insert INTO account (username, password_hash) VALUES (?, ?)")	
	if err != nil {		
		return 0, err	
	}	

	var id int64	
	result, err := ins.Exec(username, passwordHash)	
	if err != nil {
		log.Fatalf("Exec-miss")
	}
	id, err = result.LastInsertId()	
	if err != nil {
		log.Fatalf(("LastInsertId-miss"))
	}	
	return id, nil
}


