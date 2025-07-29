package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/WillKopa/boot_dev_chirpy/internal/auth"
)

func (cfg *apiConfig) login (rw http.ResponseWriter, req *http.Request) {
	type Login_request struct {
		Password	string	`json:"password"`
		Email		string	`json:"email"`
	}

	login_request, err := unmarshal_json[Login_request](req)
	if err != nil {
		respond_with_error(rw, http.StatusInternalServerError, "error parsing json from request")
		return
	}

	user, err := cfg.db_queries.GetUserByEmail(req.Context(), login_request.Email)
	if err == sql.ErrNoRows {
		log.Printf("error getting user by email: %s", err)
		respond_with_error(rw, http.StatusNotFound, "No user matches the given email")
		return
	} else if err != nil {
		log.Printf("error getting user by email: %s", err)
		respond_with_error(rw, http.StatusInternalServerError, "error, please try again later")
		return
	}

	err = auth.CheckPasswordHash(login_request.Password, user.HashedPassword)
	if err != nil {
		respond_with_error(rw, http.StatusUnauthorized, "Incorrect email or password")
		return
	}
	respond_with_json(rw, http.StatusOK, User{
		ID: user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email: user.Email,
	})
}