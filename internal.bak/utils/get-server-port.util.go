package utils

import (
	"fmt"
	"os"
	"strconv"
)

func GetServerPort() int {
	portStr := os.Getenv("PORT")

	if portStr == "" {
		portStr = "3000"
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Println("error portStr formatting:", err)
	}

	return port
}
