package main

import (
	"golb-api/handler"
)

func main() {
	srv := handler.CreateHandler()

	srv.CreateGinGroup()
}
