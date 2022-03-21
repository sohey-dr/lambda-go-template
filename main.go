package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

func main(request events.APIGatewayProxyRequest) (error) {
	fmt.Println("Hello World")
	return nil
}
