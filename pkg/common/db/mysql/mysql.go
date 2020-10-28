package mysql

import (
	"github.com/mythio/go-rest-starter/pkg/common/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnection(connectionURL string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(connectionURL), &gorm.Config{})

	db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
