package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TokenRepository struct {
	DB *pgxpool.Pool
}

func NewTokenRepository(db *pgxpool.Pool) *TokenRepository {
	return &TokenRepository{
		DB: db,
	}
}

func (r *TokenRepository) TokenBlackList(token string) error {
	_, err := r.DB.Exec(
		context.Background(),
		`INSERT INTO token_blacklist(token,expired_at)
		VALUES($1,now()+interval '24 hour')`,
		token,
	)
	return err
}

func (r *TokenRepository) IsBlackListed(token string) bool {
	var count int

	err := r.DB.QueryRow(
		context.Background(),
		`SELECT COUNT(*)
			FROM token_blacklist
			WHERE token=$1`,
		token,
	).Scan(&count)

	if err != nil {
		println("BLACKLIST CHECK ERROR:", err.Error())
	}

	println("BLACKLIST COUNT", count)

	return count > 0
}
