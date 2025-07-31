package api

import (
	"log"
	"net/http"

	"github.com/WillKopa/boot_dev_chirpy/internal/auth"
	"github.com/WillKopa/boot_dev_chirpy/internal/database"
)

func (cfg *apiConfig) update_user(rw http.ResponseWriter, req *http.Request) {
	type Update_user struct {
		Password		string		`json:"password"`
		Email			string		`json:"email"`
	}
	auth_token, err := auth.GetBearerToken(req.Header)
	if err != nil {
		respond_with_error(rw, http.StatusUnauthorized, "Missing JWT in header")
		return
	}

	token_user_id, err := auth.ValidateJWT(auth_token, cfg.secret)
	if err != nil {
		respond_with_error(rw, http.StatusUnauthorized, "YOU SHALL NOT UPDATE!")
		return
	}

	update_params, err := unmarshal_json[Update_user](req)
	if err != nil {
		respond_with_error(rw, http.StatusInternalServerError, "error parsing json from request")
		return
	}

	hashed_pw, err := auth.HashPassword(update_params.Password)
	if err != nil {
		respond_with_error(rw, http.StatusInternalServerError, "error hashing password")
		return
	}

	params := database.UpdateUserEmailPasswordParams{
		ID: token_user_id,
		Email: update_params.Email,
		HashedPassword: hashed_pw,
	}

	updated_user, err := cfg.db_queries.UpdateUserEmailPassword(req.Context(), params)

	if err != nil {
		log.Printf("error updating password and email: %s", err)
		respond_with_error(rw, http.StatusInternalServerError, "error updating email and password")
		return
	}

	respond_with_json(rw, http.StatusOK, User {
		ID: updated_user.ID,
		CreatedAt: updated_user.CreatedAt,
		UpdatedAt: updated_user.UpdatedAt,
		Email: update_params.Email,
	})
}