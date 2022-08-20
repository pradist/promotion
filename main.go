package main

import "github.com/pradist/promotion/router"

func main() {
	app := router.New()
	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
