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

	// GET RESULT ----------------------------------------
	errUserBodyParse := c.BodyParser(&user)
	errContactBodyParse := c.BodyParser(&contact)

	if errUserBodyParse != nil || errContactBodyParse != nil {
		log.Println("User Register BodyParse ERR:", errUserBodyParse.Error(), "User Register Contact Method BodyParse ERR:", errContactBodyParse.Error())
		return c.Status(http.StatusBadRequest).SendString(types.ERR_MSG_BAR_BODY_PARSE)
	}

	// Fix Value ------------------------------------------
	user.Fix()
	contact.Fix()

	// FIRST VALIDATION ERR -------------------------------
	errFields := []validation.ErrField{}

	if errUserValidationErr, err := user.Validate(); err != nil {
		errFields = append(errFields, *errUserValidationErr...)
	} else {
		var userCounts int64

		database.DB.Model(&models.User{}).Where(models.User{Username: user.Username}).Count(&userCounts)

		if userCounts > 0 {
			errFields = append(errFields, validation.ErrField{
				FieldName:  "Username",
				ErrorTitle: "Username already exits",
				Value:      user.Username,
			})
		}
	}

	if errContactMethodValidationErr, err := contact.Validate(); err != nil {
		errFields = append(errFields, *errContactMethodValidationErr...)
	} else {
		var contactCounts int64

		database.DB.Model(&models.ContactMethod{}).Where(models.ContactMethod{Contact: contact.Contact}).Count(&contactCounts)

		if contactCounts > 0 {
			errFields = append(errFields, validation.ErrField{
				FieldName:  "Contact",
				ErrorTitle: "Contact already exits",
				Value:      contact.Contact,
			})
		}
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
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"err_type": types.ERR_TYPE_BY_MULTIPLE_FIELDS,
			"fields": map[string]string{
				"title": "This project already exists",
			},
		})
	}

	return c.Status(http.StatusOK).JSON(contact.GetDictionary())
}
