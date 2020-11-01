package main

import (
	"context"
	"fmt"

	"github.com/mythio/go-rest-starter/pkg/common/db/mysql"
)

func main() {
	db, err := mysql.NewConnection("root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("err", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	res, err := db.ExecContext(ctx, `
		CREATE TABLE users (
			id					int AUTO_INCREMENT NOT NULL PRIMARY KEY,
			first_name	text,
			last_name		text,
			username		text,
			email				text,
			password		text,
			created_at 	bigint,
			updated_at 	bigint,
			deleted_at 	bigint
		)
	`)

	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(res)
}
