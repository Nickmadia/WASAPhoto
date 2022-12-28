package test

import (
	"fmt"
	"net/http"
	"testing"
)

/*
	 type TestPaths struct {
		baseUrl string `default:"http://localhost:3000/"`

}
*/

func TestUpdateUsername(t *testing.T) {
	path := "users/1/username"
	fmt.Println("path :> ", path)

	var query = `{"username":"changed"}`
	req := GetBaseReq(path, "PUT", query, "1")

	resp := ExecReq(req)
	CheckRes(t, http.StatusNoContent, resp)

}
func TestFetchUsername(t *testing.T) {
	path := `result?username=ch`
	fmt.Println("path :> ", path)

	var query = ``
	req := GetBaseReq(path, "GET", query, "1")

	resp := ExecReq(req)
	CheckRes(t, http.StatusOK, resp)
}
func TestGetUserProfile(t *testing.T) {
	path := "users/2"
	fmt.Println("path :> ", path)

	var query = ``
	req := GetBaseReq(path, "GET", query, "1")

	resp := ExecReq(req)
	CheckRes(t, http.StatusOK, resp)
}
