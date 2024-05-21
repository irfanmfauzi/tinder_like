package response

import "tinder_like/internal/model/entity"

type GenericResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	GenericResponse
	Data TokenResponse `json:"data"`
}

type GetMemberResponse struct {
	GenericResponse
	Data []entity.Member `json:"data"`
}
