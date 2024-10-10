package model

import (
	"gorm.io/gorm"
)

type ModelId uint

type Model struct {
	//ModelBasic
	ID            ModelId        `json:"id" gorm:"primary_key" sctable:"title:id;isDefaultDisplay"`
	CreatedAt     int64          `json:"created_at" gorm:"<-:create,autoCreateTime" order:"true"`
	UpdatedAt     int64          `json:"updated_at" gorm:"autoUpdateTime" order:"true"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
	CreatedById   uint           `json:"-" gorm:"<-:create"`
	CreatedBy     *AccUser       `json:"created_by,omitempty"`
	UpdatedById   uint           `json:"-"`
	UpdatedBy     *AccUser       `json:"updated_by,omitempty"`
}

type AccUser struct {
	ID         ModelId `json:"id" gorm:"primary_key" sctable:"title:id;isDefaultDisplay"`
	Username   string  `json:"username,omitempty"`
	Email      string  `json:"email,omitempty"`
	FirstName  string  `json:"first_name,omitempty"`
	LastName   string  `json:"last_name,omitempty"`
	IsSysadmin bool    `json:"-"`
	Language   string  `json:"language"`
}

type Product struct {
	Model
}

type Order struct {
	Model
}
