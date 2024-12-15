package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/learn-go-projects/gorm_fiber/database"
)

func main() {
	db := database.New()

	app := fiber.New()

	app.Get("/book", func(c *fiber.Ctx) error {
		books := db.GetManyBook()
		return c.JSON(books)

	})

	app.Get("/book/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		book, err := db.GetBook(uint(id))

		return c.JSON(book)
	})

	app.Post("/book", func(c *fiber.Ctx) error {
		var book database.Book
		if err := c.BodyParser(&book); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		db.CreateBook(&book)
		return c.JSON(fiber.Map{
			"msg": "create success",
		})
	})

	app.Put("/book/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		var book database.Book
		if err := c.BodyParser(&book); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		book.ID = uint(id)
		db.Updatebook(&book)

		return c.JSON(fiber.Map{
			"msg": "update success",
		})
	})

	app.Delete("/book/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		db.DeleteBook(uint(id))
		return c.JSON(fiber.Map{
			"msg": "delete success",
		})

	})

	// User api

	app.Post("/register", func(c *fiber.Ctx) error {
		user := new(database.User)
		if err := c.BodyParser(&user); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		err := db.CreateUser(user)

		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.JSON(fiber.Map{
			"msg": "register success",
		})

	})

	app.Post("/login", func(c *fiber.Ctx) error {
		user := new(database.User)
		if err := c.BodyParser(&user); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		token, err := db.LoginUser(user)
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 72),
			HTTPOnly: true,
		})

		return c.JSON(fiber.Map{
			"token": token,
		})

	})
	app.Listen(":3000")
}
