package mysql

import (
	"database/sql"
	"fmt"

	"github.com/mythio/go-rest-starter/pkg/common/model"
)

// User represents the client for user table
type User struct{}

// Create creates a new user on database
func (u User) Create(db *sql.DB, user model.User) (model.User, error) {
	statement, err := db.Prepare(`
		INSERT INTO users
		(first_name, last_name, username, email, password, created_at, updated_at, deleted_at)
		VALUES(?, ?, ?, ?, ?, ?, ?, ?)
		`,
	)
	if err != nil {
		return model.User{}, err
	}

	user.Base.BeforeCreate()

	result, err := statement.Exec(user.FirstName, user.LastName, user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt, user.DeletedAt)
	if err != nil {
		return model.User{}, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return model.User{}, err
	}

	return u.FindByID(db, insertedID)
}

// FindByID returns single user by ID
func (User) FindByID(db *sql.DB, id int64) (model.User, error) {
	var user model.User
	row := db.QueryRow(
		`SELECT * FROM users WHERE id = (?)`, id,
	)

	if err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	); err != nil {
		fmt.Println("errrrr", err)
		return model.User{}, nil
	}

	return user, nil
}

// FindByEmail returns single user by Email
func (User) FindByEmail(db *sql.DB, email string) (model.User, error) {
	var user model.User

	row := db.QueryRow(
		`SELECT * FROM users WHERE email LIKE (?)`, email,
	)

	if err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	); err != nil {
		fmt.Println("errrrr", err)
		return model.User{}, nil
	}

	fmt.Println(user)

	return user, nil
}
