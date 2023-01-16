package api

import (
	"WASAPhoto/service/api/reqcontext"
	"WASAPhoto/service/database"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	defer r.Body.Close()
	id, err := strconv.ParseUint(ps.ByName("user_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	auth, err := strconv.ParseUint(r.Header.Get("Authorization"), 10, 64)

	if err != nil {
		// must be authenticated
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !hasPermission(auth, id) {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	followId, err := strconv.ParseUint(ps.ByName("followed_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = rt.db.FollowUser(auth, followId)

	if errors.Is(err, database.ErrResourceDoesNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if errors.Is(err, database.ErrUserDoesNotExist) {
		w.WriteHeader(http.StatusUnauthorized)
	} else if errors.Is(err, database.ErrUserIsBanned) {
		w.WriteHeader(http.StatusForbidden)
		return
	} else if errors.Is(err, database.ErrUserIsNotAuthenticated) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	defer r.Body.Close()
	id, err := strconv.ParseUint(ps.ByName("user_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	auth, err := strconv.ParseUint(r.Header.Get("Authorization"), 10, 64)

	if err != nil {
		// must be authenticated
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !hasPermission(auth, id) {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	followId, err := strconv.ParseUint(ps.ByName("followed_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.UnfollowUser(auth, followId)
	if errors.Is(err, database.ErrResourceDoesNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if errors.Is(err, database.ErrUserDoesNotExist) {
		w.WriteHeader(http.StatusUnauthorized)
	} else if errors.Is(err, database.ErrUserIsBanned) {
		w.WriteHeader(http.StatusForbidden)
		return
	} else if errors.Is(err, database.ErrUserIsNotAuthenticated) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
