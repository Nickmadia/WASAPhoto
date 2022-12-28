package api

import (
	"WASAPhoto/service/api/reqcontext"
	"WASAPhoto/service/database"
	"WASAPhoto/service/objects"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) likeMedia(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	//check auth
	auth, err := strconv.ParseUint(r.Header.Get("Authorization"), 10, 64)

	if err != nil {
		//must be authenticated
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	postId, err := strconv.ParseUint(ps.ByName("post_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
	err = rt.db.LikeMedia(auth, postId)

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

func (rt *_router) unlikeMedia(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	//check auth
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
	postId, err := strconv.ParseUint(ps.ByName("post_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.UnlikeMedia(auth, postId)

	if errors.Is(err, database.ErrResourceDoesNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if errors.Is(err, database.ErrUserDoesNotExist) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else if errors.Is(err, database.ErrUserIsBanned) {
		w.WriteHeader(http.StatusForbidden)
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

func (rt *_router) addComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	//check auth
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
	postId, err := strconv.ParseUint(ps.ByName("post_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var comment objects.CommentPlain

	json.NewDecoder(r.Body).Decode(&comment)
	if !isCommentValid(comment.Text) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = rt.db.CommentMedia(auth, postId, comment.Text)

	if errors.Is(err, database.ErrResourceDoesNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if errors.Is(err, database.ErrUserDoesNotExist) {
		w.WriteHeader(http.StatusUnauthorized)
		return
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

func (rt *_router) removeComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	auth, err := strconv.ParseUint(r.Header.Get("Authorization"), 10, 64)
	if err != nil {
		//must be authenticated
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	postId, err := strconv.ParseUint(ps.ByName("post_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	commentId, err := strconv.ParseUint(ps.ByName("comment_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.UncommentMedia(auth, postId, commentId)

	if errors.Is(err, database.ErrResourceDoesNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if errors.Is(err, database.ErrUserDoesNotExist) {
		w.WriteHeader(http.StatusUnauthorized)
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

func isCommentValid(comment string) bool {
	return len(comment) > 0 && len(comment) < 200
}
