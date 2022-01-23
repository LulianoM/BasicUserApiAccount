package controllers

import (
	"basicuserapiaccount/src/data"
	"basicuserapiaccount/src/middlewares"
	"basicuserapiaccount/src/structs"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var resp map[string]string

	if err := c.BodyParser(&resp); err != nil {
		return err
	}

	if resp["password"] != resp["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	user := structs.User{
		FirstName: resp["first_name"],
		LastName:  resp["last_name"],
		Email:     resp["email"],
	}

	user.SetPassword(resp["password"])

	data.DB.Create(&user)

	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	var user structs.User

	data.DB.Where("id = ?", "1").First(&user)

	return c.JSON(user)
}

func UpdateSensitiveInfo(c *fiber.Ctx) error {
	var resp map[string]string

	if err := c.BodyParser(&resp); err != nil {
		return err
	}

	id, _ := middlewares.GetUserId(c)

	user := structs.SensitiveInformation{
		CPF:       resp["cpf"],
		BirthDate: resp["birth_date"],
		Gender:    resp["gender"],
	}

	user.Id = id

	data.DB.Model(&user).Updates(&user)

	return c.JSON(user)
}

func UpdatePassword(c *fiber.Ctx) error {
	var resp map[string]string

	if err := c.BodyParser(&resp); err != nil {
		return err
	}

	if resp["password"] != resp["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	id, _ := middlewares.GetUserId(c)

	user := structs.User{}
	user.Id = id

	user.SetPassword(resp["password"])

	data.DB.Model(&user).Updates(&user)

	return c.JSON(user)
}
