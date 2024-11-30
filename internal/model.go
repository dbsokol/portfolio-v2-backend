package internal

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UUID      string    `gorm:"type:varchar(36)" json:"uuid"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	IsDeleted bool      `gorm:"default:false" json:"isDeleted"`
}

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	if m.UUID == "" {
		m.UUID = uuid.New().String()
	}
	return nil
}
