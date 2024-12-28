package main

import (
	"example.com/go-ecommerce-backend-api/internal/routers"
)

func main() {
	r := routers.NewRouter()
	r.Run()
}