package main

import (
	"log"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

func processUrl(url string) string {
	r, _ := regexp.Compile("(\\d{4})/\\w*/?([A-E])")
	return r.ReplaceAllString(r.FindString(url), `$1$2`) + ".cpp"
}

func postRoute(c *fiber.Ctx) error {
	d := new(Data)
	if err := c.BodyParser(d); err != nil {
		return err
	}
	log.Println(processUrl(d.Url))
	compile(processUrl(d.Url))
	runTests(d.Tests)
	return c.SendString("Post request")
}

func runServer(port string) {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	app.Post("/", postRoute)
	app.Listen(port)
}
