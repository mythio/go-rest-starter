package model

import (
	"time"
)

// Base contains common fields for all tables
type Base struct {
	ID        int64
	CreatedAt int64
	UpdatedAt int64
	DeletedAt int64
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
