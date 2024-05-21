package repository

import (
	"context"
	"log/slog"
	"tinder_like/internal/model/entity"

	"github.com/jmoiron/sqlx"
)

type memberRepo struct {
	baseRepo
}

func NewMemberRepo(db *sqlx.DB) memberRepo {
	return memberRepo{
		baseRepo: baseRepo{db: db},
	}
}

func (u *memberRepo) InsertMember(ctx context.Context, tx TxProvider, user_id int64, name, gender string) error {
	query := "INSERT INTO members (user_id,name,gender) VALUES ($1,$2,$3)"

	_, err := u.DB(tx).ExecContext(ctx, query, user_id, name, gender)
	if err != nil {
		return err
	}

	return nil
}

func (u *memberRepo) FindMemberByUserID(ctx context.Context, userID int64) (user entity.Member, err error) {
	query := "SELECT id,user_id,name,gender FROM members where user_id = $1"

	err = u.db.GetContext(ctx, &user, query, userID)
	if err != nil {
		slog.Info("FindMemberByUserID", "user_id", userID)
		return user, err
	}

	return user, nil
}

func (u *memberRepo) FetchMemberByGenderExceptSelf(ctx context.Context, gender string, userID int64) ([]entity.Member, error) {
	query := "SELECT * FROM members where gender = $1 and user_id != $2 ORDER BY RANDOM() LIMIT 10"
	result := []entity.Member{}

	err := u.db.SelectContext(ctx, &result, query, gender, userID)
	if err != nil {
		return result, err
	}

	return result, nil

}
