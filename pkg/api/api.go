package api

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/mythio/go-rest-starter/pkg/api/middleware"
	"github.com/mythio/go-rest-starter/pkg/api/routes/auth"
	authTransport "github.com/mythio/go-rest-starter/pkg/api/routes/auth/transport"
	"github.com/mythio/go-rest-starter/pkg/api/routes/post"
	postTransport "github.com/mythio/go-rest-starter/pkg/api/routes/post/transport"
	"github.com/mythio/go-rest-starter/pkg/common/db/mysql"
	"github.com/mythio/go-rest-starter/pkg/common/util/jwt"
	"github.com/mythio/go-rest-starter/pkg/common/util/secure"
	"github.com/mythio/go-rest-starter/pkg/common/util/server"
)

// Start starts the API service
func Start() error {
	time.Sleep(10 * time.Second)
	db, err := mysql.NewConnection("root:password@tcp(localhost:3306)/test")
	if err != nil {
		fmt.Println(err)
		return err
	}

	sec := secure.New(sha1.New())
	tk, err := jwt.New("HS256", "secret", 50)
	if err != nil {
		fmt.Println(err)
		return err
	}

	s := server.New()
	s.Use(middleware.CheckAuthToken(tk))
	authRouter := s.Group("/auth")
	authTransport.NewHTTP(auth.InitService(db, sec, tk), authRouter)

	postRouter := s.Group("/post")
	postTransport.NewHTTP(post.InitService(db), postRouter)

	server.Start(s)

	return nil
}
