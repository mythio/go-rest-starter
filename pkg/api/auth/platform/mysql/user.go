package mysql

import (
	"github.com/mythio/go-rest-starter/pkg/common/model"
	"gorm.io/gorm"
)

// User represents the client for user table
type User struct{}

// Create creates a new user on database
func (User) Create(db *gorm.DB, user model.User) (model.User, error) {
	result := db.Create(&user)

	return user, result.Error
}

// FindByID returns single user by ID
func (User) FindByID(db *gorm.DB, id uint32) (model.User, error) {
	var user model.User
	result := db.Find(&user, "id = ? and deleted_at = ?", id, 0)

	return user, result.Error
}

// FindByEmail returns single user by Email
func (User) FindByEmail(db *gorm.DB, email string) (model.User, error) {
	var user model.User
	result := db.Where("email = ? and deleted_at = ?", email, 0).First(&user)

	return user, result.Error
}
