package api

import (
	"crypto/sha1"

	"github.com/mythio/go-rest-starter/pkg/api/auth"
	ut "github.com/mythio/go-rest-starter/pkg/api/auth/transport"
	"github.com/mythio/go-rest-starter/pkg/common/db/mysql"
	"github.com/mythio/go-rest-starter/pkg/common/util"
	"github.com/mythio/go-rest-starter/pkg/common/util/logger"
	"github.com/mythio/go-rest-starter/pkg/common/util/server"
)

// Start starts the API service
func Start() error {
	db, err := mysql.NewConnection("root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		return err
	}

	sec := util.New(sha1.New())
	log := logger.New()

	e := server.New()
	ut.NewHTTP(auth.InitService(db, sec, log), e)

	server.Start(e)

	return nil
}
