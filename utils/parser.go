package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func parseQueryToInt(c *fiber.Ctx, key string, defaultValue int) (int, error) {
	queryValue := c.Query(key, strconv.Itoa(defaultValue))
	return strconv.Atoi(queryValue)
}
