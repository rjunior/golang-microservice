package main

import (
	"fmt"
	"os"

	"github.com/rjunior/golang-microservice/src/api/domain/github/repositories"
)

func getRequests() []repositories.CreateRepoRequest {
	result := make([]repositories.CreateRepoRequest, 0)

	_, err := os.Open("/requests.txt")
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	requests := getRequests()

	fmt.Println((fmt.Sprintf("about to process %d requests", len(requests))))
}
