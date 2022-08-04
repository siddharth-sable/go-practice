package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/siddharth-sable/go-fiber-crm-basic/database"
	"github.com/siddharth-sable/go-fiber-crm-basic/lead"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("Database Connection Failed!")
	}
	fmt.Println("Connected to database!")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	initDatabase()
	app := fiber.New()
	setUpRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
}
