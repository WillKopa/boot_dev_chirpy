package api

import (
	"fmt"
	"net/http"

	"github.com/WillKopa/boot_dev_chirpy/internal/auth"
	"github.com/WillKopa/boot_dev_chirpy/internal/database"
)

func (cfg *apiConfig) create_user(rw http.ResponseWriter, req *http.Request) {
	type User_request struct {
		Password	string	`json:"password"`
		Email		string	`json:"email"`
	}
	
	user_request, err := unmarshal_json[User_request](req)
	if err != nil {
		respond_with_error(rw, http.StatusBadRequest, "unable to parse json from request")
		return
	}
	if len(user_request.Password) < 1 {
		respond_with_error(rw, http.StatusBadRequest, fmt.Sprintf("password is too short: must be at least %d characters long", 1))
	}

	hashed_password, err := auth.HashPassword(user_request.Password)
	if err != nil {
		respond_with_error(rw, http.StatusInternalServerError, "error saving password")
		return
	}

	db_params := database.CreateUserParams {
		Email: user_request.Email,
		HashedPassword: hashed_password,
	}

	db_user, err := cfg.db_queries.CreateUser(req.Context(), db_params)

	if err != nil {
		respond_with_error(rw, http.StatusInternalServerError, "Something went wrong, user not created")
		return
	}

	respond_with_json(rw, http.StatusCreated, User{
		ID: db_user.ID,
		CreatedAt: db_user.CreatedAt,
		UpdatedAt: db_user.UpdatedAt,
		Email: db_params.Email,
		IsChirpyRed: db_user.IsChirpyRed.Bool,
	})
}