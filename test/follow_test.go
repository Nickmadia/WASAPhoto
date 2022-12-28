package test

import (
	"fmt"
	"net/http"
	"testing"
)

// TODO CHECK IF A RESOURCE ALREADY EXIST
func TestFollow(t *testing.T) {
	path := "users/2/follow/1"
	fmt.Println("path :> ", path)

	req := GetBaseReq(path, "PUT", "", "2")

	resp := ExecReq(req)
	CheckRes(t, http.StatusNoContent, resp)
}

func TestUnfollow(t *testing.T) {
	path := "users/1/follow/2"
	fmt.Println("path :> ", path)

	req := GetBaseReq(path, "DELETE", "", "1")

	resp := ExecReq(req)
	CheckRes(t, http.StatusNoContent, resp)

}
