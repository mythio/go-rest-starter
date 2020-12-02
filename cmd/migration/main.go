package main

import (
	"context"
	"fmt"

	"github.com/mythio/go-rest-starter/pkg/common/db/mysql"
)

func main() {
	db, err := mysql.NewConnection("root:password@tcp(mythio_go-rest-starter_mysql:3306)/test")
	if err != nil {
		fmt.Println("err", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if _, err := db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			id					int AUTO_INCREMENT NOT NULL PRIMARY KEY,
			first_name	text,
			last_name		text,
			username		text,
			email				text,
			password		text,
			created_at 	bigint,
			updated_at 	bigint,
			deleted_at 	bigint
		);
	`); err != nil {
		fmt.Println("err", err)
	}

	if _, err := db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS posts (
			id					int AUTO_INCREMENT NOT NULL PRIMARY KEY,
			author_id		int,
			title				text,
			body 				text,
			likes				int,
			created_at 	bigint,
			updated_at 	bigint,
			deleted_at 	bigint
		);
	`); err != nil {
		fmt.Println("err", err)
	}

	fmt.Println("OK")
}
