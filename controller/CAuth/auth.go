package CAuth

import (
	"encoding/json"
	"errors"
	"net/http"

	"lms/config/auth"
	"lms/models/MAuth"
	"lms/responses"
	"lms/responses/formaterror"
)

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	type token_refresh struct {
		Refresh_token string `json:"refresh_token"`
	}
	var token_old token_refresh
	var new_token map[string]string
	err := json.NewDecoder(r.Body).Decode(&token_old)
	if err != nil {
		panic(err.Error())
	}
	if auth.TokenCek(token_old.Refresh_token) != nil {
		responses.JSON(w, http.StatusInternalServerError, "Expired")
		return
	} else {
		id, err := auth.ExtractTokenID(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		new_token, err = auth.BuatToken(id)
		if err != nil {
			formattedError := formaterror.FormatError(err.Error())
			responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
			return
		}
		responses.JSON(w, http.StatusOK, new_token)
	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	var user MAuth.Peserta
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user.Persiapan("")
	err = user.Validasi()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, _ := user.ProsesLogin()
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	if token == nil {
		m := make(map[string]string)
		m["status"] = "failed"
		m["pesan"] = "username/password salah"
		responses.JSON(w, http.StatusOK, m)
	} else {
		responses.JSON(w, http.StatusOK, token)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:   "token",
		MaxAge: -1}
	http.SetCookie(w, &c)
	w.Write([]byte("Old cookie deleted. Logged out!\n"))
}
