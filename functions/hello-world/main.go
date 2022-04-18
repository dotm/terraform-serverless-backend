package main

import (
	"dotm/terraform-serverless-backend/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Greeting struct {
	Message string `json:"message"`
}

func greet(name string) Greeting {
	if name == "" {
		name = "World"
	}
	return Greeting{
		Message: utils.AddExclamation(fmt.Sprintf("Hello, %s!", name)),
	}
}

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

// Add a helper for handling errors. This logs any error to os.Stderr
// and returns a 500 Internal Server Error response that the AWS API
// Gateway understands.
func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := req.QueryStringParameters["Name"]

	resp := greet(name)

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(jsonResp),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
