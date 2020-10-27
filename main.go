package main

import (
	"net/http"

	h "github.com/mythio/go-rest-starter/cmd/http"
	mysqlRepo "github.com/mythio/go-rest-starter/repository/persistent"
	"github.com/mythio/go-rest-starter/user"
	l "github.com/mythio/go-rest-starter/util/logger"
	"gopkg.in/go-playground/validator.v9"
)

var v *validator.Validate

func main() {
	l := l.NewZapLogger()
	repo, err := mysqlRepo.NewMongoRepository("localhost", 3306, "test", "root", "password", 30, l)

	if err != nil {
		l.Debug("err")
	}
	v = validator.New()
	service := user.NewUserService(repo, l, v)
	handler := h.NewHandler(service)

	http.HandleFunc("/", handler.Signup)

	http.ListenAndServe(":8080", nil)
}
