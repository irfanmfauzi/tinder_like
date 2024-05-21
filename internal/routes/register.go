package routes

import (
	"net/http"
	handleAuth "tinder_like/internal/handler/auth"
	handleMember "tinder_like/internal/handler/member"
	"tinder_like/internal/repository"
	"tinder_like/middleware"

	"github.com/jmoiron/sqlx"
)

func RegisterRoute(db *sqlx.DB) http.Handler {
	mux := http.NewServeMux()

	userRepo := repository.NewUserRepo(db)
	memberRepo := repository.NewMemberRepo(db)
	swipeMemberRepo := repository.NewSwipeMember(db)

	authHandler := handleAuth.NewAuthHandler(handleAuth.AuthHandlerConfig{Db: db, UserRepo: &userRepo, MemberRepo: &memberRepo})

	mux.HandleFunc("POST /api/v1/auth/register", authHandler.Register)
	mux.HandleFunc("POST /api/v1/auth/login", authHandler.Login)

	memberHandler := handleMember.NewMemberHandler(handleMember.MemberHandlerConfig{Db: db, UserRepo: &userRepo, MemberRepo: &memberRepo, SwipeMemberRepo: &swipeMemberRepo})

	mux.Handle("GET /api/v1/members", middleware.VerifyToken(http.HandlerFunc(memberHandler.GetMembers)))
	mux.Handle("POST /api/v1/members/swipe", middleware.VerifyToken(http.HandlerFunc(memberHandler.SwipeMember)))

	return mux
}
