package helper

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

var (
	Req req
	Res res
)

type req struct{}
type res struct{}

type Paginate struct {
	Page  int `json:"page" example:"0"`
	Limit int `json:"limit" example:"0"`
	Total int `json:"total" example:"0"`
}

type queryResult struct {
	Page     int
	Limit    int
	Keyword  string
	OrderBy  string
	SortedBy string
}

func (*req) Offset(page int, limit int) int {
	return (page - 1) * limit
}

func (*req) QueryStr(c *fiber.Ctx) queryResult {
	getPage := strings.TrimSpace(c.Query("page", "1"))
	getLimit := strings.TrimSpace(c.Query("limit", "10"))
	getKeyword := strings.TrimSpace(c.Query("keyword", ""))
	getOrderBy := strings.TrimSpace(c.Query("orderBy", "id"))
	getSortedBy := strings.TrimSpace(c.Query("sortedBy", "asc"))

	newPage := Str2Int(getPage)
	if newPage < 1 {
		newPage = 1
	}

	newLimit := Str2Int(getLimit)
	if newLimit < 1 {
		newLimit = 1
	}
	if newLimit > 100 {
		newLimit = 100
	}

	sortedByToLower := strings.ToLower(getSortedBy)
	if sortedByToLower != "asc" && sortedByToLower != "desc" {
		getSortedBy = "asc"
	}

	return queryResult{
		Page:     newPage,
		Limit:    newLimit,
		Keyword:  getKeyword,
		OrderBy:  getOrderBy,
		SortedBy: getSortedBy,
	}
}

func (*res) SendCustom(c *fiber.Ctx, data interface{}, statusCode ...int) error {
	code := fiber.StatusOK // Default status code
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	return c.Status(code).JSON(data)
}

func (*res) SendSuccess(c *fiber.Ctx, msg string) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": msg,
	})
}

func (*res) SendData(c *fiber.Ctx, msg string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    data,
		"message": msg,
	})
}

func (*res) SendDatas(c *fiber.Ctx, msg string, data interface{}, paginate Paginate) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":       data,
		"pagination": paginate,
		"message":    msg,
	})
}

func (*res) SendErrorMsg(c *fiber.Ctx, msg string, statusCode ...int) error {
	code := fiber.StatusBadRequest // Default status code
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	return c.Status(code).JSON(fiber.Map{
		"message": msg,
	})
}

func (*res) SendErrors(c *fiber.Ctx, msg string, err interface{}) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"errors":  err,
		"message": msg,
	})
}
