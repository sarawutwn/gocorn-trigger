package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/robfig/cron"
)

func main() {
	app := fiber.New(fiber.Config{
		// Prefork: true,
	})
	c := cron.New()
	c.AddFunc("@every 1m", func() {
		fmt.Println(time.Now())
	})
	c.Start()
	app.Listen(":10000")
}
