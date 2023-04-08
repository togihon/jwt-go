package main

import "jwt-go/app/handler"

func main() {
	var PORT = ":8080"

	handler.StartServer().Run(PORT)
}
