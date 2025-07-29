package api

import "net/http"

func (cfg *apiConfig) create_user(rw http.ResponseWriter, req *http.Request) {
	user, err := unmarshal_json[User](req)
	if err != nil {
		respond_with_error(rw, http.StatusBadRequest, "unable to parse json from request")
		return
	}
	db_user, err := cfg.db_queries.CreateUser(req.Context(), user.Email)

	if err != nil {
		respond_with_error(rw, http.StatusInternalServerError, "Something went wrong, user not created")
		return
	}

	respond_with_json(rw, http.StatusCreated, User(db_user))
}