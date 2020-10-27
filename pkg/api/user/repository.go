package user

type Repository interface {
	Create(user *User) error
	Find(user *User) (*User, error)
}
