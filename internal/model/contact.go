package model

import (
	"time"
)

type Contact struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string         `gorm:"type:varchar(100);not null;check:length(name) >= 3" json:"name"`
	Email     string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Phone     string         `gorm:"type:varchar(20);uniqueIndex;not null;check:length(phone) >= 8" json:"phone"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}
