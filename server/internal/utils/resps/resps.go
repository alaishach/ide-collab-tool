// Package resps
package resps

import (
	"encoding/json"
	"net/http"
	"server/internal/consts"
	"server/internal/err/panics"
)

const halfYear = 3600 * 24 * 180

func SetServerCookie(w http.ResponseWriter, name string, value string) {
	if consts.ENV == "dev" {
		http.SetCookie(w, &http.Cookie{Name: name, Value: value, Domain: consts.SERVER_DOMAIN, Path: "/api", MaxAge: halfYear, HttpOnly: true, Secure: false, SameSite: http.SameSiteDefaultMode})
	} else {
		http.SetCookie(w, &http.Cookie{Name: name, Value: value, Domain: consts.SERVER_DOMAIN, Path: "/api", MaxAge: halfYear, HttpOnly: true, Secure: true, SameSite: http.SameSiteDefaultMode})
	}
}

// -------------------------------- RESPONSES
// For consistency every request and response will be in JSON format so even RespMessage is json with just a message attribute

func RespJSON[T any](w http.ResponseWriter, status int, v map[string]T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		panics.PanicErr("writeJSON failed to json stringify: ", err)
	}
}

func RespMessage(w http.ResponseWriter, status int, msg string) {
	RespJSON(w, status, map[string]string{"message": msg})
}
