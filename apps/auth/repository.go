package auth

import (
	"context"
	"database/sql"

	internalerror "github.com/ArdyJunata/grpc-go/internal/error"
	"github.com/jmoiron/sqlx"
)

func newRepository(db *sqlx.DB) (repo repository) {
	repo = repository{
		db: db,
	}

	return
}

type repository struct {
	db *sqlx.DB
}

// GetAuthByUsername implements Repository.
func (r repository) GetAuthByUsername(ctx context.Context, username string) (model Auth, err error) {
	query := `
		SELECT
			id, username
			,password
			,created_at, updated_at
		FROM auth
		WHERE username=$1
	`

	stmt, err := r.db.PreparexContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()
	err = stmt.GetContext(ctx, &model, username)
	if err != nil {
		if err == sql.ErrNoRows {
			err = internalerror.ErrNotFound
			return
		}
		return
	}

	return
}

// CreateAuth implements Repository.
func (r repository) CreateAuth(ctx context.Context, model Auth) (err error) {
	query := `
		INSERT INTO auth (
			username,
			password
		) VALUES (
			:username,
			:password
		)
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, model)

	if err != nil {
		return
	}

	return
}
