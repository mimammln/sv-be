package repository

import (
	"article/driver"
)

type DatabaseResponse struct {
	Code    int    `json:"code" validate:"required"`
	Message string `json:"message" validate:"required"`
}

func GenerateDatabase() *DatabaseResponse {
	Res := &DatabaseResponse{500, "Internal Server Error"}

	driver.GenerateDatabase()

	Res = &DatabaseResponse{200, "Database Ready!"}
	return Res
}
