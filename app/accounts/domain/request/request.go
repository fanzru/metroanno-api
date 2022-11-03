package request

import (
	"github.com/volatiletech/null/v9"
)

type UserRegisterReq struct {
	SubjectPreference null.String `json:"subject_preference" gorm:"subject_preference"`
	Username          string      `json:"username" gorm:"username" validate:"required,min=6"`
	Password          string      `json:"password" gorm:"password" validate:"required,min=6"`
	Contact           string      `json:"contact" gorm:"contact" validate:"required"`
	Age               int64       `json:"age" gorm:"age" validate:"required"`
}

type UserLoginReq struct {
	Username string `json:"username" gorm:"username" validate:"required,min=6"`
	Password string `json:"password" gorm:"password" validate:"required,min=6"`
}
