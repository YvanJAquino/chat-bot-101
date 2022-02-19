package handlers

import (
	"fmt"
	"net/http"

	gchat "github.com/yaq-cc/go-gchat/objects"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	var req gchat.ChatRequest
	req.FromHTTPRequest(r)
	fmt.Println(req.User.DisplayName)
	fmt.Fprint(w, "Hello World!")
}
