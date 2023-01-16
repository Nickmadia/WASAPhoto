package test

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
)

const BASEURL = "http://localhost:3000/"

func CheckRes(t *testing.T, expected int, resp *http.Response) string {
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != expected {
		t.Errorf("Expected %d got %d\n ", expected, resp.StatusCode)
	}
	if string(body) == "" {
		fmt.Println("EMPTY BODY")
	} else {
		fmt.Printf("Body : %s\n", body)
	}
	return string(body)
}

func ExecReq(req *http.Request) *http.Response {
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func GetBaseReq(path string, method string, body string, auth string) *http.Request {
	if body != "" {
		req, err := http.NewRequest(method, BASEURL+path, bytes.NewReader([]byte(body)))
		if err != nil {
			return nil
		}
		req.Header.Set("Authorization", auth)
		req.Header.Set("Content-Type", "application/json")
		return req
	} else {
		req, err := http.NewRequest(method, BASEURL+path, nil)
		if err != nil {
			return nil
		}
		req.Header.Set("Authorization", auth)
		return req
	}

}
