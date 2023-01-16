package api

import (
	"WASAPhoto/service/api/reqcontext"
	"WASAPhoto/service/objects"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) signIn(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var username objects.Username

	err := json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	}
	if !isValidUsername(username.Text) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	// in case the username does not exist create a user profile with that username and return the ID
	id, err := rt.db.SignInOrLogin(username.Text)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	}

	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	}
}
