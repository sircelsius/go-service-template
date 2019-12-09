package main

import (
	"fmt"

	"github.com/sircelsius/go-service-template/internal/http"
)

func main() {
	_ = http.NewClient("google")
	fmt.Printf("Hello World!")
}
