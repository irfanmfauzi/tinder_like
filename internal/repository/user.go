package repository

import (
	"context"
	"tinder_like/internal/model/entity"
	"tinder_like/internal/model/request"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	baseRepo
}

func NewUserRepo(db *sqlx.DB) userRepo {
	return userRepo{
		baseRepo: baseRepo{db: db},
	}
}

func (u *userRepo) InsertUser(ctx context.Context, tx TxProvider, registerRequest request.RequestRegister) (user_id int64, err error) {
	query := "INSERT INTO users (email, password, is_premium, is_verified, is_infinite_quota) VALUES ($1,$2,$3,$4,$5) RETURNING id"
	isPremium := registerRequest.Premium != request.PremiumRequestRegister{}

	err = u.DB(tx).GetContext(ctx, &user_id, query, registerRequest.Email, registerRequest.Password, isPremium, registerRequest.Premium.IsVerified, registerRequest.Premium.IsInfiniteQuota)
	if err != nil {
		return user_id, err
	}

	return user_id, nil
}

func (u *userRepo) FindUserByEmail(ctx context.Context, email string) (user entity.User, err error) {
	query := "SELECT id,email,password,is_premium,is_verified,is_infinite_quota FROM users where email = $1"

	err = u.db.GetContext(ctx, &user, query, email)
	if err != nil {
		return user, err
	}

	return user, nil
}
