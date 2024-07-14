package controllers

import (
	"taskmanagerserver.com/api/database"
	"taskmanagerserver.com/api/models"
)

func ControllRegisterUserAccountLog(userId uint, description string) {
	log := models.RegisterUserAccountLog(userId, description)

	database.DB.Create(&log)
}
