package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

func main(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("GoLang Lambda")

	return events.APIGatewayProxyResponse{
		Body:       "GoLang Lambda",
		StatusCode: 200,
	}, nil
}