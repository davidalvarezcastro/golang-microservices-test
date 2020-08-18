package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/davidalvarezcastro/golang-microservices-test/src/api/models/repositories"
	"github.com/davidalvarezcastro/golang-microservices-test/src/api/services"
	"github.com/davidalvarezcastro/golang-microservices-test/src/api/utils/errors"
)

var (
	success map[string]string
	failed  map[string]errors.APIError
)

type createRepoResult struct {
	Request repositories.CreateRepoRequest
	Result  *repositories.CreateRepoResponse
	Error   errors.APIError
}

func getRequets() []repositories.CreateRepoRequest {
	result := make([]repositories.CreateRepoRequest, 0)

	file, err := os.Open("/home/david/Desktop/requests.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		request := repositories.CreateRepoRequest{
			Name:        line,
			Description: "created using api at repo golang-microservices-test",
		}

		result = append(result, request)
	}

	return result
}

func main() {
	requests := getRequets()
	input := make(chan createRepoResult)

	// this is our limiter of concurrency of our script, in this case is limited to execute 10 requests
	buffer := make(chan bool, 10)
	var wg sync.WaitGroup

	go handleResults(&wg, input)

	fmt.Println(fmt.Sprintf("about to process %d requests", len(requests)))
	for _, request := range requests {
		// 1st, 2nd, ... request until 10th; then it blocks execution
		buffer <- true

		go CreateRepo(buffer, request, input)
	}

	wg.Wait()

	close(input)
}

func handleResults(wg *sync.WaitGroup, input chan createRepoResult) {
	for result := range input {
		if result.Error != nil {
			failed[result.Request.Name] = result.Error
		} else {
			success[result.Request.Name] = result.Result.Name
		}

		wg.Done()
	}
}
func a(wg *sync.WaitGroup) {
	fmt.Println(&wg)
}

// CreateRepo executes the repositories.CreateRepo function using concurrency
func CreateRepo(buffer chan bool, request repositories.CreateRepoRequest, output chan createRepoResult) {
	result, err := services.RepositoryService.CreateRepo(request)

	output <- createRepoResult{
		Request: request,
		Result:  result,
		Error:   err,
	}

	// fake reader: consumes one element in the buffer to allow adding one more element and executing another request
	<-buffer
}
