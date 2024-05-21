package handler

import (
	"database/sql"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"time"
	"tinder_like/internal/model"
	"tinder_like/internal/model/entity"
	"tinder_like/internal/model/request"
	"tinder_like/internal/model/response"
	"tinder_like/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
)

type AuthHandlerConfig struct {
	Db         *sqlx.DB
	UserRepo   repository.UserRepo
	MemberRepo repository.MemberRepo
}

type authHandler struct {
	db         *sqlx.DB
	userRepo   repository.UserRepo
	memberRepo repository.MemberRepo
}

func NewAuthHandler(cfg AuthHandlerConfig) authHandler {
	return authHandler{
		db:         cfg.Db,
		userRepo:   cfg.UserRepo,
		memberRepo: cfg.MemberRepo,
	}
}

func (a *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()
	request := request.RequestRegister{}

	err = json.Unmarshal(body, &request)
	if err != nil {
		resp, _ := json.Marshal(response.GenericResponse{Success: false, Message: err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(resp)
		return
	}

	ctx := r.Context()

	tx, err := a.db.BeginTxx(ctx, nil)
	if err != nil {
		slog.Error("Failed to Begin Transaction", "Error", err)
		resp, _ := json.Marshal(response.GenericResponse{Success: false, Message: "Something Wrong With System"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(resp)
		return
	}
	defer tx.Rollback()

	user_id, err := a.userRepo.InsertUser(ctx, tx, request)
	if err != nil {
		slog.Error("Failed to Insert User", "Error", err)
		resp, _ := json.Marshal(response.GenericResponse{Success: false, Message: "Something Wrong With System"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(resp)
		return
	}

	err = a.memberRepo.InsertMember(ctx, tx, user_id, request.Name, request.Gender)
	if err != nil {
		slog.Error("Failed to Insert Member", "Error", err)
		resp, _ := json.Marshal(response.GenericResponse{Success: false, Message: "Something Wrong With System"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(resp)
		return
	}

	err = tx.Commit()
	if err != nil {
		slog.Error("Failed to Commit", "Error", err)
		resp, _ := json.Marshal(response.GenericResponse{Success: false, Message: "Something Wrong With System"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(resp)
		return
	}

	resp, _ := json.Marshal(response.GenericResponse{Success: true, Message: "Success Register"})
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

func (a *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application-json")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()
	request := request.RequestLogin{}

	err = json.Unmarshal(body, &request)
	if err != nil {
		resp, _ := json.Marshal(response.GenericResponse{Success: false, Message: err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(resp)
		return
	}
	ctx := r.Context()

	user, err := a.userRepo.FindUserByEmail(ctx, request.Email)
	if err != nil {
		msg := ""
		code := http.StatusInternalServerError

		if err == sql.ErrNoRows {
			msg = "Email or Password is Wrong"
			code = http.StatusBadRequest
			slog.Info("Failed to Find User By Email", "Error", err)
		} else {
			msg = "Something Wrong With System"
			slog.Error("Failed to Find User By Email", "Error", err)
		}

		resp, _ := json.Marshal(response.GenericResponse{Success: false, Message: msg})
		w.WriteHeader(code)
		w.Write(resp)

		return
	}

	if user.Password != request.Password {
		resp, _ := json.Marshal(response.GenericResponse{Success: false, Message: "Email or Password is Wrong"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(resp)
		return
	}

	member, err := a.memberRepo.FindMemberByUserID(ctx, user.Id)
	if err != nil {
		slog.Error("Failed to Find Member By User ID", "Error", err)
		resp, _ := json.Marshal(response.GenericResponse{Success: false, Message: "Something Wrong With System"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(resp)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, model.CustomClaim{
		User: entity.User{
			Id:              user.Id,
			Email:           user.Email,
			IsPremium:       user.IsPremium,
			IsVerified:      user.IsVerified,
			IsInfiniteQuota: user.IsInfiniteQuota,
		},
		Member: member,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24).UTC()),
		},
	})

	tokenString, err := token.SignedString([]byte("REPLACE_THIS_WITH_ENV_VAR_SECRET"))
	if err != nil {
		slog.Error("Failed to Signed String", "Error", err)
		resp, _ := json.Marshal(response.GenericResponse{Success: false, Message: "Something Wrong With System"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(resp)

	}

	resp, _ := json.Marshal(
		response.LoginResponse{
			GenericResponse: response.GenericResponse{Success: true, Message: "Login Success"},
			Data:            response.TokenResponse{Token: tokenString},
		},
	)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}
