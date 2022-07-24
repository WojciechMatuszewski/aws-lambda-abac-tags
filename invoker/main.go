package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	awslambda "github.com/aws/aws-sdk-go-v2/service/lambda"
	awslambdatypes "github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	ststypes "github.com/aws/aws-sdk-go-v2/service/sts/types"
)

func handler(ctx context.Context, _ events.APIGatewayProxyRequest) (
	events.APIGatewayProxyResponse, error) {

	cfg, err := config.LoadDefaultConfig(ctx) // config.WithCredentialsProvider(

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	stsClient := sts.NewFromConfig(cfg)
	lambdaCfg, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(
			stscreds.NewAssumeRoleProvider(
				stsClient,
				os.Getenv("ROLE_ARN"),
				func(o *stscreds.AssumeRoleOptions) {
					o.Tags = append(o.Tags, ststypes.Tag{Key: aws.String("Team"), Value: aws.String("Falcon")})
				},
			),
		),
	)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	client := awslambda.NewFromConfig(lambdaCfg)

	functionName := os.Getenv("FUNCTION_ARN")
	out, err := client.Invoke(ctx, &awslambda.InvokeInput{
		FunctionName:   aws.String(functionName),
		InvocationType: awslambdatypes.InvocationTypeRequestResponse,
		Payload:        []byte("{}"),
	})
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(out.Payload),
	}, nil
}

func main() {
	lambda.Start(handler)
}
