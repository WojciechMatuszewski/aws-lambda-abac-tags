package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

type Payload struct {
	Message string `json:"message"`
}

func handler(ctx context.Context) (Payload, error) {
	return Payload{Message: "ToBeInvokedPayload"}, nil
}
