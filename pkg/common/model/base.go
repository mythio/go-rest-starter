package model

import (
	"time"
)

// Base contains common fields for all tables
type Base struct {
	ID        int64 `json:"id" db:"id"`
	CreatedAt int64 `json:"created_at" db:"created_at"`
	UpdatedAt int64 `json:"updated_at" db:"updated_at"`
	DeletedAt int64 `json:"deleted_at,omitempty" db:"deleted_at"`
}

// BeforeCreate hooks into insert operations, setting createdAt and updatedAt to current time
func (b *Base) BeforeCreate() {
	b.CreatedAt = time.Now().UTC().Unix()
	b.UpdatedAt = time.Now().UTC().Unix()
}

// BeforeUpdate hooks into update operati	ons, setting updatedAt to current time
func (b *Base) BeforeUpdate() {
	b.UpdatedAt = time.Now().UTC().Unix()
}
