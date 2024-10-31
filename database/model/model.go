package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Model contains common columns for all tables.
type Model struct {
	ID        string   `gorm:"primarykey"`
	CreatedAt DateTime `gorm:"index"`
	UpdatedAt DateTime
}

// BeforeCreate generates the primary key and audit columns.
func (b *Model) BeforeCreate(tx *gorm.DB) error {
	b.ID = uuid.New().String()
	b.CreatedAt = Now()
	b.UpdatedAt = Now()
	return nil
}
