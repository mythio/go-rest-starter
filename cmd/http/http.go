package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mythio/go-rest-starter/user"
)

type UserHandler interface {
	Signup(http.ResponseWriter, *http.Request)
}

type handler struct {
	userService user.UserService
}

func NewHandler(userService user.UserService) UserHandler {
	return &handler{
		userService,
	}
}

func (h *handler) Signup(w http.ResponseWriter, r *http.Request) {
	user := &user.User{}
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if err := json.Unmarshal(requestBody, &user); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	h.userService.Signup(user)
	w.WriteHeader(200)
}
