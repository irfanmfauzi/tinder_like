package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type swipeMemberRepo struct {
	baseRepo
}

func NewSwipeMember(db *sqlx.DB) swipeMemberRepo {
	return swipeMemberRepo{
		baseRepo: baseRepo{db: db},
	}
}
func (s *swipeMemberRepo) InsertSwipeMember(ctx context.Context, tx TxProvider, memberId, swipedMemberId int64, isLiked bool) error {
	query := "INSERT INTO swipe_members (member_id, swiped_member_id, is_liked) VALUES ($1,$2,$3)"
	_, err := s.DB(tx).Exec(query, memberId, swipedMemberId, isLiked)
	if err != nil {
		return err
	}
	return nil
}
