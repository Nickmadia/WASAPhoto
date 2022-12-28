package api

import (
	"WASAPhoto/service/api/reqcontext"
	"WASAPhoto/service/database"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	auth, err := strconv.ParseUint(r.Header.Get("Authorization"), 10, 64)

	if err != nil {
		//must be authenticated
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	id, err := strconv.ParseUint(ps.ByName("user_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	banId, err := strconv.ParseUint(ps.ByName("ban_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !hasPermission(auth, id) {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	err = rt.db.BanUser(auth, banId)
	//TODO consider modifiying userIsBannedError
	if errors.Is(err, database.ErrResourceDoesNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if errors.Is(err, database.ErrUserDoesNotExist) {
		w.WriteHeader(http.StatusUnauthorized)
	} else if errors.Is(err, database.ErrUserIsBanned) {
		w.WriteHeader(http.StatusForbidden)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	auth, err := strconv.ParseUint(r.Header.Get("Authorization"), 10, 64)

	if err != nil {
		//must be authenticated
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	id, err := strconv.ParseUint(ps.ByName("user_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !hasPermission(auth, id) {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	banId, err := strconv.ParseUint(ps.ByName("ban_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.UnbanUser(auth, banId)
	//TODO consider modifiying userIsBannedError
	if errors.Is(err, database.ErrResourceDoesNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if errors.Is(err, database.ErrUserDoesNotExist) {
		w.WriteHeader(http.StatusUnauthorized)
	} else if errors.Is(err, database.ErrUserIsBanned) {
		w.WriteHeader(http.StatusForbidden)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
