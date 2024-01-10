package main

import "begingo/server"

func main() {
	r := server.NewRouter()

	r.Run(":8888")
}
