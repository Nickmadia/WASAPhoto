package api

import (
	"WASAPhoto/service/api/reqcontext"
	"WASAPhoto/service/database"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMedia(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	postId, err := strconv.ParseUint(ps.ByName("post_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//check auth
	auth, err := strconv.ParseUint(r.Header.Get("Authorization"), 10, 64)
	if err != nil {
		//must be authenticated
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//delegate check to db  wether the user is banned or not
	img, err := rt.db.GetMedia(auth, postId)

	if errors.Is(err, database.ErrProfileDoesNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if errors.Is(err, database.ErrUserIsBanned) {
		w.WriteHeader(http.StatusForbidden)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	}

	w.Header().Set("content-type", "image/png")
	_, err = w.Write([]byte(*img))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	}
}

func (rt *_router) getMediaMetadata(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	postId, err := strconv.ParseUint(ps.ByName("post_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//check auth
	auth, err := strconv.ParseUint(r.Header.Get("Authorization"), 10, 64)
	if err != nil {
		//must be authenticated
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//delegate check to db  wether the user is banned or not
	imgMetadata, err := rt.db.GetMediaMetadata(auth, postId)

	if errors.Is(err, database.ErrResourceDoesNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if errors.Is(err, database.ErrUserIsBanned) {
		w.WriteHeader(http.StatusForbidden)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	}

	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(*imgMetadata)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	}
}

func (rt *_router) postMedia(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	//check auth
	auth, err := strconv.ParseUint(r.Header.Get("Authorization"), 10, 64)

	if err != nil {
		//must be authenticated
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	img, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	imgStr := string(img)
	imgId, err := rt.db.UploadImage(auth, &imgStr)

	if errors.Is(err, database.ErrResourceDoesNotExist) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else if errors.Is(err, database.ErrUserIsNotAuthenticated) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(imgId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	}
}

func (rt *_router) deleteMedia(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	postId, err := strconv.ParseUint(ps.ByName("post_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//check auth
	auth, err := strconv.ParseUint(r.Header.Get("Authorization"), 10, 64)
	if err != nil {
		//must be authenticated
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = rt.db.DeleteMedia(auth, postId)

	if errors.Is(err, database.ErrResourceDoesNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if errors.Is(err, database.ErrUserDoesNotOwnTheResource) {
		w.WriteHeader(http.StatusForbidden)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
