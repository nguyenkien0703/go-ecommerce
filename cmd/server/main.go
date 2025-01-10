package main

import (
	"example.com/go-ecommerce-backend-api/internal/initialize"
)

func main() {
	// r := routers.NewRouter()
	// r.Run()// default 8080, if you want to change port, you can pass it as an argumnent to Run(":8081")
	initialize.Run()

}
