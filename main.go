package main

import (
	"gogin/router"
)

func main() {
	router := router.InitRouter()
	router.Run(":8000")
}
