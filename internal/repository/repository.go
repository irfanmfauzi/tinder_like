package repository

import (
	"context"
	"database/sql"
	"tinder_like/internal/model/entity"
	"tinder_like/internal/model/request"
)

type TxProvider interface {
	Commit() error
	Rollback() error
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type QueryProvider interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type UserRepo interface {
	InsertUser(ctx context.Context, tx TxProvider, registerRequest request.RequestRegister) (userID int64, err error)
	FindUserByEmail(ctx context.Context, email string) (user entity.User, err error)
}

type MemberRepo interface {
	InsertMember(ctx context.Context, tx TxProvider, userID int64, name, gender string) error
	FindMemberByUserID(ctx context.Context, userID int64) (user entity.Member, err error)
	FetchMemberByGenderExceptSelf(ctx context.Context, gender string, userID int64) ([]entity.Member, error)
}

type SwipeMemberRepo interface {
	InsertSwipeMember(ctx context.Context, tx TxProvider, memberId, swipedMemberId int64, isLiked bool) error
}
