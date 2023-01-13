package test

import (
	"encoding/base64"
	"fmt"
	"io"

	"net/http"
	"os"
	"testing"
)

func TestMedia(t *testing.T) {
	path := "media"
	fmt.Println("path :> ", path)

	img, err := os.ReadFile("img.jpeg")
	if err != nil {
		t.Error(err)
	}
	b64 := base64.StdEncoding.EncodeToString(img)
	req := GetBaseReq(path, "POST", b64, "1")

	resp := ExecReq(req)
	CheckRes(t, http.StatusCreated, resp)

}
func TestGetMedia(t *testing.T) {

	path := "media/1"
	fmt.Println("path :> ", path)

	req := GetBaseReq(path, "GET", "", "1")

	resp := ExecReq(req)
	// checkRes(t, http.StatusOK, resp)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	var img []byte
	img, err = base64.StdEncoding.DecodeString(string(body))
	if err != nil {
		t.Error(err)
	}
	err = os.WriteFile("imgdb.jpeg", img, 0644)
	if err != nil {
		t.Error(err)
	}

}

func TestMediaMetadata(t *testing.T) {

	path := "media/1/info"
	fmt.Println("path :> ", path)

	req := GetBaseReq(path, "GET", "", "1")

	resp := ExecReq(req)
	CheckRes(t, http.StatusOK, resp)
}

//ok
/* func DelMedia(t *testing.T, id string) {
	path := "media/" + id
	fmt.Println("path :> ", path)

	req := GetBaseReq(path, "DELETE", "", "1")

	resp := ExecReq(req)
	CheckRes(t, http.StatusNoContent, resp)

} */
