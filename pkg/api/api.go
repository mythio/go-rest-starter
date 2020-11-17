package api

import (
	"crypto/sha1"
	"fmt"

	"github.com/mythio/go-rest-starter/pkg/api/middleware"
	"github.com/mythio/go-rest-starter/pkg/api/routes/auth"
	authTransport "github.com/mythio/go-rest-starter/pkg/api/routes/auth/transport"
	"github.com/mythio/go-rest-starter/pkg/api/routes/post"
	postTransport "github.com/mythio/go-rest-starter/pkg/api/routes/post/transport"
	"github.com/mythio/go-rest-starter/pkg/common/db/mysql"
	"github.com/mythio/go-rest-starter/pkg/common/db/redis"
	"github.com/mythio/go-rest-starter/pkg/common/util/jwt"
	"github.com/mythio/go-rest-starter/pkg/common/util/logger"
	"github.com/mythio/go-rest-starter/pkg/common/util/secure"
	"github.com/mythio/go-rest-starter/pkg/common/util/server"
)

// Start starts the API service
func Start() error {
	db, err := mysql.NewConnection("root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println(err)
		return err
	}

	redis.NewConnection("")

	sec := secure.New(sha1.New())
	log := logger.New()
	tk, err := jwt.New("HS256", "secret", 50)
	if err != nil {
		fmt.Println(err)
		return err
	}

	s := server.New()
	s.Use(middleware.CheckAuthToken(tk))
	authRouter := s.Group("/auth")
	authTransport.NewHTTP(auth.InitService(db, sec, log, tk), authRouter)

	postRouter := s.Group("/post")
	postTransport.NewHTTP(post.InitService(db), postRouter)

	server.Start(s)

	return nil
}
