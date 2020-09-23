package main

import (
	"context"
	"encoding/json"
	"fmt"
	"kfc-microservice/event"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const (
	PRODUCT_PATH = "/products"
	STORE_PATH   = "/stores"

	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

// Handler is the Lambda function handler
func Handler(ctx context.Context, payload events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var resp []byte

	switch path := payload.Path; path {
	case PRODUCT_PATH:
		switch method := payload.HTTPMethod; method {
		case GET:
			fmt.Printf("GET method for products.\n")
		case POST:
			newStock, err := event.CreateStock(payload.Body)
			if err != nil {
				panic(err)
			}

			resp, _ = json.Marshal(newStock)
		case PUT:
			fmt.Printf("PUT method for products.\n")
		case DELETE:
			fmt.Printf("DELETE method for products.\n")
		}

	case STORE_PATH:
		switch method := payload.HTTPMethod; method {
		case GET:
			fmt.Printf("GET method for stocks.\n")
		case POST:
			newStockLoc, err := event.CreateStockLocation(payload.Body)
			if err != nil {
				panic(err)
			}

			resp, _ = json.Marshal(newStockLoc)
		case PUT:
			fmt.Printf("PUT method for stocks.\n")
		case DELETE:
			fmt.Printf("DELETE method for stocks.\n")
		}
	default:
		fmt.Printf("panik: %s.\n", path)
	}

	return events.APIGatewayProxyResponse{
		Body:       string(resp),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
