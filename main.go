package main

import "github.com/pradist/promotion/router"

func main() {
	app := router.New()
	app.Listen(":8080")
}
