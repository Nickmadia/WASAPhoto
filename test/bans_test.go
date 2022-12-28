package test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestBan(t *testing.T) {
	path := "users/1/banned/2"
	fmt.Println("path :> ", path)

	req := GetBaseReq(path, "PUT", "", "1")

	resp := ExecReq(req)
	CheckRes(t, http.StatusNoContent, resp)
}

func TestUnban(t *testing.T) {
	path := "users/1/banned/2"
	fmt.Println("path :> ", path)

	req := GetBaseReq(path, "DELETE", "", "1")

	resp := ExecReq(req)
	CheckRes(t, http.StatusNoContent, resp)
}
