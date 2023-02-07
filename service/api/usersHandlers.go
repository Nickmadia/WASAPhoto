package api

import (
	"WASAPhoto/service/api/reqcontext"
	"WASAPhoto/service/database"
	"WASAPhoto/service/objects"
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

const USERNAMEREX = "^[a-zA-Z1-9]*$"
const MAXUSERNAMELENGTH = 16

type UserInfo struct {
	Followers []objects.Profile       `json:"followers"`
	Following []objects.Profile       `json:"following"`
	Posts     []objects.PhotoMetadata `json:"posts"`
	IsBanned  bool                    `json:"is_banned"`
}

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	defer r.Body.Close()

	id, err := strconv.ParseUint(ps.ByName("user_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// auth checks
	// TODO remember to encapsulate
	auth, err := strconv.ParseUint(r.Header.Get("Authorization"), 10, 64)
	if err != nil {
		// must be authenticated
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// The db will check in the query if the user is banned
	profile, err := rt.db.GetUserProfile(id, auth)

	if errors.Is(err, database.ErrProfileDoesNotExist) {
		// profile not found return status 404
		w.WriteHeader(http.StatusNotFound)
		return
	} else if errors.Is(err, database.ErrUserIsBanned) {
		w.WriteHeader(http.StatusForbidden)
		return
	} else if err != nil {
		// internal errors raise 500
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	}

	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(profile.FromDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	}

}

func (rt *_router) updateUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

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

	var updatedUsername objects.Username

	err = json.NewDecoder(r.Body).Decode(&updatedUsername)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if !isValidUsername(updatedUsername.Text) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.UpdateUsername(id, updatedUsername.Text)

	if errors.Is(err, database.ErrProfileDoesNotExist) {
		// profile not found return status 404
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// if internal errors raise 500
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (rt *_router) fetchUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	defer r.Body.Close()
	if !r.URL.Query().Has("username") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fecthedUsername := r.URL.Query().Get("username")

	if !isValidUsername(fecthedUsername) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	auth, err := strconv.ParseUint(r.Header.Get("Authorization"), 10, 64)

	if err != nil {
		// must be authenticated
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// db returns max 20 profiles, will not include users who have banned the requesters
	profiles, err := rt.db.FetchUsername(fecthedUsername, auth)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	}

	apiProfiles := make([]objects.Profile, len(profiles))
	for i, element := range profiles {
		apiProfiles[i] = element.FromDatabase()
	}
	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(apiProfiles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	}
}

func (rt *_router) getUserInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	defer r.Body.Close()
	id, err := strconv.ParseUint(ps.ByName("user_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// auth checks
	auth, err := strconv.ParseUint(r.Header.Get("Authorization"), 10, 64)
	if err != nil {
		// must be authenticated
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// The db will check in the query if the user is banned
	followers, following, err := rt.db.GetUserInfo(id, auth)

	if errors.Is(err, database.ErrProfileDoesNotExist) {
		// profile not found return status 404
		w.WriteHeader(http.StatusNotFound)
		return
	} else if errors.Is(err, database.ErrUserIsBanned) {
		w.WriteHeader(http.StatusForbidden)
		return
	} else if err != nil {
		// internal errors raise 500
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	}
	posts, err := rt.db.GetUserPosts(id, auth)
	if errors.Is(err, database.ErrProfileDoesNotExist) {
		// profile not found return status 404
		w.WriteHeader(http.StatusNotFound)
		return
	} else if errors.Is(err, database.ErrUserIsBanned) {
		w.WriteHeader(http.StatusForbidden)
		return
	} else if err != nil {
		// internal errors raise 500
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	}
	isBanned, err := rt.db.IsBanned(auth, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	}
	response := new(UserInfo)
	response.IsBanned = isBanned
	response.Posts = posts
	for _, element := range followers {
		response.Followers = append(response.Followers, element.FromDatabase())
	}
	for _, element := range following {
		response.Following = append(response.Following, element.FromDatabase())
	}
	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Error(err)
		return
	}

}

// helpers
func isValidUsername(username string) bool {
	res, _ := regexp.MatchString(USERNAMEREX, username)
	return res && 0 < len(username) && len(username) <= MAXUSERNAMELENGTH
}

func hasPermission(idReq uint64, idOwner uint64) bool {
	return idReq == idOwner
}
