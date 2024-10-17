package repos

import (
	"context"
	"errors"
	"fmt"
	"monospec-api/auth/api/apple/types"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	context context.Context
	dbPool  *pgxpool.Pool
}

func NewUserRepo(dbPool *pgxpool.Pool, context context.Context) *UserRepo {
	return &UserRepo{
		dbPool:  dbPool,
		context: context,
	}
}

func (u *UserRepo) GetUserByAppleSub(sub string) (*types.User, error) {
	var user types.User

	query := "SELECT u.id, u.first_name, u.email FROM users u JOIN user_identities ui ON u.id = ui.user_id WHERE ui.provider_id = $1"

	err := u.dbPool.QueryRow(u.context, query, sub).Scan(&user.Id, &user.Name, &user.Email)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to get user by apple sub: %w", err)
	}

	return &user, nil
}

func (u *UserRepo) UpdateUserLoginAt(id int64) error {
	_, err := u.dbPool.Exec(u.context, "UPDATE users SET login_at = $1 WHERE id = $2", time.Now(), id)

	if err != nil {
		return fmt.Errorf("failed to update user login at: %w", err)
	}

	return nil
}

func (u *UserRepo) CreateUser(appleSub string, firstName string, email string, isEmailVerified bool) (*types.User, error) {
	now := time.Now()
	tx, err := u.dbPool.BeginTx(u.context, pgx.TxOptions{})

	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	queryString := "INSERT INTO users (first_name, email, email_verified, login_at) VALUES ($1, $2, $3, $4) RETURNING id"

	var userId int64
	err = tx.QueryRow(u.context, queryString, firstName, email, isEmailVerified, now).Scan(&userId)

	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	_, err = tx.Exec(u.context, "INSERT INTO user_identities (user_id, auth_provider, provider_id) VALUES ($1, $2, $3)", userId, "apple", appleSub)

	if err != nil {
		return nil, fmt.Errorf("failed to insert user identity: %w", err)
	}

	err = tx.Commit(u.context)

	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	user := &types.User{
		Id:    userId,
		Name:  firstName,
		Email: email,
	}

	return user, nil
}
