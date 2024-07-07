package controllers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"taskmanagerserver.com/api/database"
	"taskmanagerserver.com/api/models"
	"taskmanagerserver.com/api/types"
	"taskmanagerserver.com/api/validation"
)

func AuthRegister(c *fiber.Ctx) error {
	var user *models.User
	var contact *models.ContactMethod

	errUserBodyParse := c.BodyParser(&user)
	errContactBodyParse := c.BodyParser(&contact)

	if errUserBodyParse != nil || errContactBodyParse != nil {
		log.Println("User Register BodyParse ERR:", errUserBodyParse.Error(), "User Register Contact Method BodyParse ERR:", errContactBodyParse.Error())
		return c.Status(http.StatusBadRequest).SendString(types.ERR_MSG_BAR_BODY_PARSE)
	}

	errFields := []validation.ErrField{}

	if adultErrFields, err := user.Validate(); err != nil {
		errFields = append(errFields, *adultErrFields...)
	}

	if contactMethodErrFields, err := contact.Validate(); err != nil {
		errFields = append(errFields, *contactMethodErrFields...)
	}

	if len(errFields) > 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"err_fields": errFields,
			"err_type":   types.ERR_TYPE_BY_MULTIPLE_FIELDS,
			"user":       user,
		})
	}

	if err := database.DB.Create(&user).Error; err != nil {
		log.Println("User Register User ERR:", err)

	}

	return c.Status(http.StatusOK).JSON(contact.GetDictionary())
}
