package models

import (
	"strings"
	"time"

	"taskmanagerserver.com/api/models/customtypes"
	"taskmanagerserver.com/api/validation"
)

type ContactMethod struct {
	ID        uint                                `json:"id" gorm:"primaryKey"`
	Contact   string                              `json:"contact" gorm:"not null;type:varchar(300);unique" validate:"required"`
	Type      customtypes.CustomContactMethodType `json:"type" gorm:"not null;type:custom_contact_method_type"`
	Primary   bool                                `json:"primary" gorm:"default:false"`
	Verified  bool                                `json:"verified" gorm:"default:false"`
	UserID    uint                                `json:"user_id"`
	User      User                                `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdateAt  time.Time
}

func (cm ContactMethod) GetDictionary() *Dictionary {
	dic := Dictionary{
		"id":       cm.ID,
		"contact":  cm.Contact,
		"type":     cm.Type.String(),
		"primary":  cm.Primary,
		"verified": cm.Verified,
	}

	return &dic
}

func (cm *ContactMethod) Fix() {
	if len(cm.Contact) > 0 {
		cm.Contact = strings.ToLower(cm.Contact)
	}

	if len(cm.Type) > 0 {
		cm.Type.Scan(strings.ToUpper(cm.Type.String()))
	}
}

func (cm ContactMethod) Validate() (Dictionary, error) {
	if err := validation.Validate.StructExcept(cm, "User"); err != nil {
		errors := validation.GetValidateInformation(err, cm)

		return DictionarySetup(errors), err
	}

	return nil, nil
}
