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

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

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

	app.Post("/updateSrStatus/:no/:status", func(c *fiber.Ctx) error {
	   return functionsfolder.UpdateSr(c,session)
 
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
