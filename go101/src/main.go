package main

import (
	"fmt"
	"go-101/app"
)

func main() {
	fmt.Println("About to start app")
	app.Start(
		app.InitRoutes(),
	)
}
