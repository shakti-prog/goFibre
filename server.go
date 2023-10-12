package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"goFibre/dbConnection"
	"goFibre/functionsFolder"
)

func main() {
	// Create a new Fiber instance
	app, session := dbconnection.ConnectToDB()

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*") // Allow requests from any origin
		c.Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		return c.Next()
	})

	// Define routes and handlers
	app.Get("/readData", func(c *fiber.Ctx) error {
		data := functionsfolder.ReadData(session)
		return c.JSON(data)
	})

	app.Post("/insertData", func(c *fiber.Ctx) error {
		readMessage := functionsfolder.InsertData(session, c)
		return c.JSON(readMessage)
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		return functionsfolder.Login(c, session)
	})

	app.Post("/signUp", func(c *fiber.Ctx) error {
		return functionsfolder.SignUp(c, session)
	})

	app.Get("/getSrData", func(c *fiber.Ctx) error {
		return functionsfolder.GetSrData(c, session)
	})

	app.Post("/createSr", func(c *fiber.Ctx) error {
		return functionsfolder.CreateNewSr(c, session)
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
