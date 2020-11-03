package api

import (
	"crypto/sha1"
	"fmt"

	"github.com/mythio/go-rest-starter/pkg/api/middleware"
	"github.com/mythio/go-rest-starter/pkg/api/routes/auth"
	authTransport "github.com/mythio/go-rest-starter/pkg/api/routes/auth/transport"
	"github.com/mythio/go-rest-starter/pkg/common/db/mysql"
	"github.com/mythio/go-rest-starter/pkg/common/util/jwt"
	"github.com/mythio/go-rest-starter/pkg/common/util/logger"
	"github.com/mythio/go-rest-starter/pkg/common/util/secure"
	"github.com/mythio/go-rest-starter/pkg/common/util/server"
)

// Start starts the API service
func Start() error {
	db, err := mysql.NewConnection("root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		return err
	}

	sec := secure.New(sha1.New())
	log := logger.New()
	tk, err := jwt.New("HS256", "secret", 5)
	if err != nil {
		fmt.Println(err)
	}

	s := server.New()
	s.Use(middleware.CheckAuthToken(tk))
	authRouter := s.Group("/auth")
	authTransport.NewHTTP(auth.InitService(db, sec, log, tk), authRouter)

	server.Start(s)

	return nil
}
