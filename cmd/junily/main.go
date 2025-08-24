package main

import (
	_ "embed"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/markory/junily/internal/assets"
)


func main() {
	app := fiber.New()

	// Access file "image.png" under `static/` directory via URL: `http://<server>/static/image.png`.
	// Without `PathPrefix`, you have to access it via URL:
	// `http://<server>/static/static/image.png`.
	app.Use("/static", filesystem.New(filesystem.Config{
		Root: http.FS(assets.DashboardStaticFiles),
		PathPrefix: "static",
		Browse: true,
	}))

	log.Fatal(app.Listen(":3000"))
}