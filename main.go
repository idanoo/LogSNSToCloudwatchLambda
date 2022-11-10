package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, snsEvent events.SNSEvent) {
	for _, record := range snsEvent.Records {
		fmt.Printf("Message: %s", record.SNS.Message)
	}
}

func main() {
	lambda.Start(handler)
}
