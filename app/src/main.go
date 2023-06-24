package main

import "github.com/labstack/echo/v4"

func main() {
	router := echo.New()
	router.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	router.Logger.Fatal(router.Start(":8080"))
}
