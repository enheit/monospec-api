package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func CreateAuthChallengeLambda(event events.CognitoEventUserPoolsCreateAuthChallenge) (*events.CognitoEventUserPoolsCreateAuthChallengeResponse, error) {
	fmt.Println(event)

	return &events.CognitoEventUserPoolsCreateAuthChallengeResponse{}, nil
}

func main() {
	lambda.Start(CreateAuthChallengeLambda)
}
