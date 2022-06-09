package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func processUrl(url string) string {
	r, _ := regexp.Compile("\\d{4}/problem/[A-E]")
	return strings.Replace(r.FindString(url), "/problem/", "", -1) + ".cpp"
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
