package api

import (
	"WASAPhoto/service/api/reqcontext"
	"WASAPhoto/service/database"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userId, err := strconv.ParseUint(ps.ByName("user_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//TODO fix the check
	auth, err := strconv.ParseUint(r.Header.Get("Authorization"), 10, 64)
	if err != nil {
		//must be authenticated
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !hasPermission(auth, userId) {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	//delegate check to db  wether the user is banned or not
	postList, err := rt.db.GetStream(userId)

	if errors.Is(err, database.ErrUserIsNotAuthenticated) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(postList)
}
