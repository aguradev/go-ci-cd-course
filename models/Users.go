package models

import (
	users "praktikum/models/Users"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id        uint           `gorm:"primaryKey" json:"id,omitempty"`
	Name      string         `gorm:"type:varchar(255)" json:"name,omitempty"`
	Email     string         `gorm:"type:varchar(50)" json:"email,omitempty"`
	Password  string         `gorm:"type:varchar(255)" json:"password,omitempty"`
	Blogs     []Blogs        `gorm:"foreignKey:UserID" json:"blogs,omitempty"`
	CreatedAt time.Time      `gorm:"autoCreateTime"  json:"created_at,omitempty"`
	UpdateAt  time.Time      `gorm:"autoUpdateTime:mili" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (user Users) CreateFormData(Request users.UserCreateRequest) Users {

	var getUser Users

	getUser.Name = Request.Name
	getUser.Email = Request.Email
	getUser.Password = Request.Password

	return getUser
}
