package models

import (
	"time"

	"taskmanagerserver.com/api/models/customtypes"
)

type UserLog struct {
	ID          uint                      `json:"id" gorm:"primaryKey"`
	Description string                    `json:"description" gorm:"not null;type:text"`
	LogType     customtypes.CustomLogType `json:"log_type" gorm:"not null;type:custom_log_type"`
	UserID      uint                      `json:"user_id"`
	User        User                      `gorm:"foreignKey:UserID"`
	CreatedAt   time.Time
}

func RegisterUserAccountLog(userID uint, description string) UserLog {

	log := UserLog{
		UserID:      userID,
		Description: description,
		LogType:     customtypes.LOG_TYPE_ACCOUNT,
	}

	return log
}
