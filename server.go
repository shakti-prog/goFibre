package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"goFibre/dbConnection"
	"goFibre/functionsFolder"
)

func main() {
	// Create a new Fiber instance
	app, session := dbconnection.ConnectToDB()

	config := cors.Config{
		AllowOrigins: "http://localhost:8000", // Replace with the origin you want to allow
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}

	app.Use(cors.New(config))

	// Define routes and handlers
	app.Get("/readData", func(c *fiber.Ctx) error {
		data := functionsfolder.ReadData(session)
		return c.JSON(data)
	})

	app.Post("/insertData", func(c *fiber.Ctx) error {
		readMessage := functionsfolder.InsertData(session, c)
		return c.JSON(readMessage)
	})

	app.Post("/uploadCsvData", func(c *fiber.Ctx) error {
		return c.JSON("Successfully updated CSV data")
	})

	app.Patch("/updateData", func(c *fiber.Ctx) error {
		return c.JSON("This is patch request")
	})

	app.Delete("/deleteData", func(c *fiber.Ctx) error {
		return c.JSON("This is delete request")
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		return functionsfolder.Login(c, session)
	})

	app.Post("/signUp", func(c *fiber.Ctx) error {
		return functionsfolder.SignUp(c, session)
	})

	//Port number
	err := app.Listen(":9000")
	if err != nil {
		fmt.Println("Failed to start server", err)
	} else {
		fmt.Println("Server started successfully")
	}

	defer session.Close()
}
