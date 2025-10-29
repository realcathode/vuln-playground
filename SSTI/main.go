package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

func sstiHandler(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		name = "Guest"
	}

	templateString := "Hello, " + name

	tmpl, err := template.New("ssti").Parse(templateString)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error parsing template: "+err.Error())
	}

	// the entire 'echo.Context' is passed
	// as data to the Execute function. This gives the template
	// access to all exported methods on the context object, such as .File(), .Redirect(), etc.
	err = tmpl.Execute(c.Response().Writer, c)
	if err != nil {
		log.Printf("Error executing template: %s", err)
	}
	return nil
}

func main() {
	e := echo.New()

	e.GET("/hello", sstiHandler)

	port := "8080"
	log.Printf("Starting vulnerable SSTI server (Echo) on :%s", port)
	log.Println("Test with: http://localhost:8080/hello?name=World")
	e.Logger.Fatal(e.Start(":" + port))
}
