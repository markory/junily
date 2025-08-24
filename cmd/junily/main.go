package main

import (
	_ "embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/markory/junily/internal/assets"
)


func main() {
	app := fiber.New()

	// We need to "strip" the folder so that /static maps correctly
    subFS, err := fs.Sub(assets.StaticFiles, "dashboard")
    if err != nil {
        log.Fatal(err)
    }

	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.FS(subFS),
		Browse: true,
	}))

	log.Fatal(app.Listen(":3000"))
}