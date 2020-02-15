package main

import (
	"horgh-consumer/app"
)

func main() {
	application, err := app.New()
	if err != nil {
		panic(err)
	}

	application.Wait()
}
