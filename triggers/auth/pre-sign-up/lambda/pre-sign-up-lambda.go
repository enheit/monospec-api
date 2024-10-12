package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func PreSignUpLambda(event events.CognitoEventUserPoolsPreSignup) *events.CognitoEventUserPoolsPreSignupResponse {
	return &events.CognitoEventUserPoolsPreSignupResponse{
		AutoConfirmUser: true,
		AutoVerifyEmail: true,
	}
}

func main() {
	lambda.Start(PreSignUpLambda)
}
