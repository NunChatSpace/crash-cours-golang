package main

import (
	"github.com/NunChatSpace/crash-cours-golang/http"
)

func main() {
	s := http.Build()

	s.Run(":8080")
}
