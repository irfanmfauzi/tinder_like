package handler

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"tinder_like/internal/model/request"
	"tinder_like/internal/model/response"
	"tinder_like/internal/repository"

	"github.com/jmoiron/sqlx"
)

type MemberHandlerConfig struct {
	Db              *sqlx.DB
	UserRepo        repository.UserRepo
	MemberRepo      repository.MemberRepo
	SwipeMemberRepo repository.SwipeMemberRepo
}

type memberHandler struct {
	db              *sqlx.DB
	userRepo        repository.UserRepo
	memberRepo      repository.MemberRepo
	swipeMemberRepo repository.SwipeMemberRepo
}

func NewMemberHandler(cfg MemberHandlerConfig) memberHandler {
	return memberHandler{
		db:              cfg.Db,
		userRepo:        cfg.UserRepo,
		memberRepo:      cfg.MemberRepo,
		swipeMemberRepo: cfg.SwipeMemberRepo,
	}
}

func (m *memberHandler) GetMembers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	member := r.Context().Value("member").(map[string]interface{})
	gender := member["Gender"].(string)
	sentGender := ""

	if gender == "Male" {
		sentGender = "Female"
	} else {
		sentGender = "Male"
	}

	userID := r.Context().Value("userId").(int64)

	members, err := m.memberRepo.FetchMemberByGenderExceptSelf(r.Context(), sentGender, userID)
	if err != nil {
		slog.Error("Failed to Fetch Member", "Error", err)
		resp, _ := json.Marshal(response.GenericResponse{Success: false, Message: "Something Wrong With System"})
		w.Write(resp)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, _ := json.Marshal(response.GetMemberResponse{
		GenericResponse: response.GenericResponse{
			Success: true,
			Message: "Success ",
		}, Data: members,
	})
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (m *memberHandler) SwipeMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		resp, _ := json.Marshal(response.GenericResponse{Success: false, Message: "Something Wrong With System"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(resp)
		slog.Error("Failed to Read Body", "Error", err)
		return
	}
	defer r.Body.Close()

	req := request.RequestSwipedMember{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		resp, _ := json.Marshal(response.GenericResponse{Success: false, Message: "Something Wrong With System"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(resp)
		slog.Error("Failed to Unmarshal", "Error", err)
		return
	}

	member := r.Context().Value("member").(map[string]interface{})
	memberId := int64(member["Id"].(float64))

	tx, err := m.db.BeginTxx(r.Context(), nil)
	if err != nil {
		resp, _ := json.Marshal(response.GenericResponse{Success: false, Message: "Something Wrong With System"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(resp)
		slog.Error("Failed to Begin Transaction", "Error", err)
		return
	}
	defer tx.Rollback()

	err = m.swipeMemberRepo.InsertSwipeMember(r.Context(), tx, memberId, req.SwipeMemberId, req.IsLiked)
	if err != nil {
		slog.Info("Debug", "SwipeMemberId", req.SwipeMemberId, "MemberId", memberId)
		resp, _ := json.Marshal(response.GenericResponse{Success: false, Message: "Something Wrong With System"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(resp)
		slog.Error("Failed to InsertSwipeMember", "Error", err)
		return
	}

	resp, _ := json.Marshal(response.GenericResponse{Success: true, Message: "Success Swiping"})
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

	tx.Commit()
}
