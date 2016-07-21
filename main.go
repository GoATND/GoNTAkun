package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/syo-sa1982/GoNTAkun/resources"
)

func rooter(e *echo.Echo) *echo.Echo {
	fmt.Println(e)

	resource := resources.Resource{}
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.Get("/api/event", resource.GetEvents())
	e.Post("/api/event", resource.CreateEvent())
	e.Get("/api/event/:id", resource.GetEvent())
	e.Put("/api/event/:id", resource.UpdateEvent())
	e.Delete("/api/event/:id", resource.DeleteEvent())

	e.Post("/api/join", resource.JoinEvent())
	e.Delete("/api/join/:join_id", resource.CancelEvent())

	e.Post("/api/comment", resource.CreateComment())
	e.Delete("/api/comment/:comment_id", resource.DeleteComment())

	e.Static("/", "public")

	return e
}

func main() {
	fmt.Println("main")
	e := rooter(echo.New())
	fmt.Println("Server running at http://localhost:60000")
	// Start server
	e.Run(standard.New(":60000"))
}

