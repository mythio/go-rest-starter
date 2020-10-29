package model

import (
	"time"

	"gorm.io/gorm"
)

// Base contains common fields for all tables
type Base struct {
	ID        uint32 `json:"id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt int64  `json:"deleted_at,omitempty" pg:",soft_delete"`
}

// BeforeCreate hooks into insert operations, setting createdAt and updatedAt to current time
func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.CreatedAt = time.Now().UTC().Unix()
	return nil
}

// BeforeUpdate hooks into update operati	ons, setting updatedAt to current time
func (b *Base) BeforeUpdate(tx *gorm.DB) (err error) {
	b.UpdatedAt = time.Now().UTC().Unix()
	return nil
}
