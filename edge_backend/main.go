package main

import (
	"app/database"
	"app/router"
	"app/scheduler"
	"log"

	"github.com/gofiber/fiber/v2"

	// "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	//runtime.GOMAXPROCS(4)
	//if !fiber.IsChild() {
	//	go handler.ScanJob()
	//}

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "App Name",
	})
	app.Use(recover.New())
	go scheduler.StartAutomationScheduler()

	database.ConnectDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))

}
