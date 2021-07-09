package main

import (
	"os"
	"strconv"

	"project.com/mecabapi/api"
)

func GetPort() int {
	port := os.Getenv("PORT")
	i, err := strconv.Atoi(port)
	if err != nil {
		return 3000
	}
	return i
}

func main() {
	api.Bootstrap(GetPort())
}
