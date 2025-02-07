package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/deposinator/models"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&user)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, fmt.Sprintf("error decoding user: %s", err), http.StatusBadRequest)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {

}

func Logout(w http.ResponseWriter, r *http.Request) {

}
