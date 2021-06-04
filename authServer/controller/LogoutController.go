package controller

import (
	"net/http"

	"github.com/partha-sen/ostd/authServer/token"
)

func HandleLogOut(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		jti := r.Header.Get("TOKEN_ID")
		token.GlobalTokenStore.Remove(jti)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
