package main

import "github.com/pradist/promotion/router"

func main() {
	app := router.New()
	err := app.Run(":8080")
	if err != nil {
		panic(err)
	}
}
