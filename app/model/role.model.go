package model

import "time"

type Role struct {
	Id        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar(255);unique"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp(0);not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp(0);not null;default:CURRENT_TIMESTAMP"`

	RoleHashPermission []RoleHasPermission `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
