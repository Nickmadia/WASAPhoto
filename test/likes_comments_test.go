package test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestLike(t *testing.T) {
	path := "media/1/likes/1"
	fmt.Println("path :> ", path)

	req := GetBaseReq(path, "PUT", "", "1")

	resp := ExecReq(req)
	CheckRes(t, http.StatusNoContent, resp)
}
func TestUnLike(t *testing.T) {
	path := "media/1/likes/1"
	fmt.Println("path :> ", path)

	req := GetBaseReq(path, "DELETE", "", "1")

	resp := ExecReq(req)
	CheckRes(t, http.StatusNoContent, resp)

}

func TestComment(t *testing.T) {
	path := "media/1/comments/comment/3"
	fmt.Println("path :> ", path)
	body := `{"comment_text":"cool pic"}`
	req := GetBaseReq(path, "POST", body, "3")

	resp := ExecReq(req)
	CheckRes(t, http.StatusNoContent, resp)
}
func TestUnComment(t *testing.T) {
	path := "media/1/comments/delete/3"
	fmt.Println("path :> ", path)
	body := `{"comment_text":"cool pic"}`
	req := GetBaseReq(path, "DELETE", body, "3")

	resp := ExecReq(req)
	CheckRes(t, http.StatusNoContent, resp)

}
