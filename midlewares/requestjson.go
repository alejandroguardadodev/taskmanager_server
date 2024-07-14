package midlewares

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

func RouteRequestToJSON(c *fiber.Ctx) error {
	var dat map[string]string

	if err := json.Unmarshal(c.Body(), &dat); err != nil {
		log.Println("Request Err: ", err.Error())
	}

	c.Locals("requestbody", dat)

	return c.Next()
}
