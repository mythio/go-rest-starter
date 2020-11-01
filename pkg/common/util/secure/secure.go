package secure

import (
	"hash"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	h hash.Hash
}

// New initializes security service
func New(h hash.Hash) *Service {
	return &Service{h: h}
}

// Hash hashes the password using bcrypt
func (*Service) Hash(password string) string {
	hashedPw, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPw)
}

// Verify returns true if the hash matches the same of the password
func (*Service) Verify(hash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
