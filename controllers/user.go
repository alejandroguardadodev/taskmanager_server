package controllers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"taskmanagerserver.com/api/database"
	"taskmanagerserver.com/api/models"
	"taskmanagerserver.com/api/tools"
	"taskmanagerserver.com/api/types"
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
	errFields := models.Dictionary{}

	if errUserValidationErr, err := user.Validate(); err != nil {
		errFields = errFields.Add(errUserValidationErr)
	} else {
		var userCounts int64

		database.DB.Model(&models.User{}).Where(models.User{Username: user.Username}).Count(&userCounts)

		if userCounts > 0 {
			errFields["Username"] = "Username already exits"
		}
	}

	if errContactMethodValidationErr, err := contact.Validate(); err != nil {
		errFields = errFields.Add(errContactMethodValidationErr)
	} else {
		var contactCounts int64

		database.DB.Model(&models.ContactMethod{}).Where(models.ContactMethod{Contact: contact.Contact}).Count(&contactCounts)

		if contactCounts > 0 {
			errFields["Contact"] = "Contact already exits"
		}
	}

	// CHECK ERR LIST AND SHOW ERR MESSAGE
	if len(errFields) > 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"fields":   errFields,
			"err_type": types.ERR_TYPE_BY_MULTIPLE_FIELDS,
		})
	}

	// CREATE USER
	if err := database.DB.Create(&user).Error; err != nil {
		log.Println("User Register User ERR:", err)

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"err_type": types.ERR_TYPE_MESSAGE,
			"message":  "Unexpected error",
		})
	}

	contact.UserID = user.ID

	// CREATE CONTACT TABLE
	if err := database.DB.Create(&contact).Error; err != nil {
		log.Println("User Register Contact ERR:", err)

		if deleterr := database.DB.Where(&models.User{ID: user.ID}).Delete(&[]models.User{}).Error; deleterr != nil { // REMOVE USER IF THERE NO WAY TO SAVE CONTACT
			log.Println("User Register Deleting USER ERR:", deleterr)
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"err_type": types.ERR_TYPE_MESSAGE,
			"message":  "Unexpected error",
		})
	}

	// USER LOG
	ControllRegisterUserAccountLog(user.ID, types.LOG_MSG_ACCOUNT_CREATED)

	// CREATE USER TOKEN
	token, tokenErr := tools.GenerateJWTToken(user.ID, contact.Contact, contact.Type.String())

	if tokenErr != nil {
		log.Println("User Register Token ERR:", tokenErr)

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"err_type": types.ERR_TYPE_MESSAGE,
			"message":  "Unexpected error",
		})
	}

	// SEND TOKEN
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"user":  user.GetDictionaryWithPrimaryContact(*contact),
		"token": token,
	})
}
