package test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestFeed(t *testing.T) {
	path := "feed/2"
	fmt.Println("path :> ", path)

	req := GetBaseReq(path, "GET", "", "2")

	resp := ExecReq(req)
	CheckRes(t, http.StatusOK, resp)

}
