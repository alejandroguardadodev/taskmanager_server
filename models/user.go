package models

import (
	"time"

	"taskmanagerserver.com/api/validation"
)

type User struct {
	ID         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Username   string `jsong:"username" gorm:"unique" validate:"required"`
	FirstName  string `json:"firstname" gorm:"not null;type:varchar(100)"`
	MiddleName string `json:"middlename" gorm:"type:varchar(100)"`
	LastName   string `json:"lastname" gorm:"type:varchar(200)"`
	Token      string `json:"token" gorm:"type:text"`
	Active     bool   `json:"active" gorm:"default:true"`
	Password   string `json:"password" gorm:"not null;type:text" validate:"required"`
	CreatedAt  time.Time
	UpdateAt   time.Time
}

func (u User) GetDictionary() *Dictionary {
	dic := Dictionary{
		"id":         u.ID,
		"username":   u.Username,
		"firstname":  u.FirstName,
		"middlename": u.MiddleName,
		"lastname":   u.LastName,
		"created_at": u.CreatedAt,
	}

	return &dic
}

func (u User) Validate() (*[]validation.ErrField, error) {

	if err := validation.Validate.StructExcept(u, "User"); err != nil {
		errFields := validation.GetValidateInformation(err, u)

		return errFields, err
	}

	return nil, nil
}
