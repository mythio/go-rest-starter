package mysql

import (
	"fmt"
	"time"

	"github.com/mythio/go-rest-starter/user"
	"github.com/mythio/go-rest-starter/util/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mySQLRepository struct {
	client  *gorm.DB
	timeout time.Duration
}

func NewMongoRepository(host string, port uint16, database string, userName string, password string, timeout uint8, logger logger.Logger) (user.Repository, error) {
	repo := &mySQLRepository{
		timeout: time.Duration(timeout) * time.Second,
	}
	connectionURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", userName, password, host, port, database)

	db, err := gorm.Open(mysql.Open(connectionURL), &gorm.Config{})
	db.AutoMigrate(&user.User{})
	if err != nil {
		return nil, err
	}

	repo.client = db

	return repo, nil
}

func (r *mySQLRepository) Create(user *user.User) error {
	r.client.Create(&user)
	return nil
}

func (r *mySQLRepository) Find(user *user.User) (*user.User, error) {
	r.client.Where("email_id = ?", user.EmailID).First(&user)

	fmt.Println(user)
	return nil, nil
}

// 	ctx, cancel := context.WithTimeout(context.Background(), r.timeout*time.Second)
// 	defer cancel()

// 	user := &user.User{}
// 	row := r.client.QueryRowContext(ctx, "select * from user where user.email_id = $1", emailID)
// 	err := row.Scan(user)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }

// func (r *mySQLRepository) Create(user *user.User) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), r.timeout*time.Second)
// 	defer cancel()

// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}

// 	user.Password = string(hashedPassword)
// 	user.UUID = uuid.New().String()
// 	user.CreatedAt = time.Now().UTC().Unix()

// 	queryString, err := r.client.PrepareContext(ctx, "insert into product(first_name, last_name, email_id, password, uuid, created_at) values ($1, $2, $3, $4, $5, $6)")
// 	if err != nil {
// 		return err
// 	}
// 	result, err := queryString.ExecContext(ctx,
// 		user.FirstName,
// 		user.LastName,
// 		user.EmailID,
// 		user.Password,
// 		user.UUID,
// 		user.CreatedAt,
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println(fmt.Sprintf("%v", result))

// 	return nil
// }

// func (r *mySQLRepository) CloseConnection() {
// 	r.client.Close()
// }
