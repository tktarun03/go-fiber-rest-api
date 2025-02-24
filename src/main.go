package main

import (
    "github.com/gofiber/fiber/v2"
)

type Item struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
}

var items []Item

func main() {
    app := fiber.New()

    app.Get("/api/items", func(c *fiber.Ctx) error {
        return c.JSON(items)
    })

    app.Post("/api/items", func(c *fiber.Ctx) error {
        var item Item
        if err := c.BodyParser(&item); err != nil {
            return err
        }
        items = append(items, item)
        return c.JSON(item)
    })

    app.Listen(":3000")
}
