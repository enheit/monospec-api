package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func VerifyAuthChallengeResponseLambda(event events.CognitoEventUserPoolsVerifyAuthChallenge) (*events.CognitoEventUserPoolsVerifyAuthChallengeResponse, error) {
	fmt.Println(event)

	return &events.CognitoEventUserPoolsVerifyAuthChallengeResponse{}, nil
}

func main() {
	lambda.Start(VerifyAuthChallengeResponseLambda)
}
