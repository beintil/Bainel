package main

import (
	"Bainel/cmd/configs"
	"Bainel/internal/pkg/myapp"
)

func main() {
	myapp.Run()
	configs.GetMongoURL()
}
