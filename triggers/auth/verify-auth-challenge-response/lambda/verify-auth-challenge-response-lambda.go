package main

import (
	"monospec-api/triggers/auth/verify-auth-challenge-response/scraps"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func VerifyAuthChallengeResponseLambda(event events.CognitoEventUserPoolsVerifyAuthChallenge) *events.CognitoEventUserPoolsVerifyAuthChallengeResponse {
	applePublicKeys, err := scraps.FetchApplePublicKeys()

	if err != nil {
		return &events.CognitoEventUserPoolsVerifyAuthChallengeResponse{
			AnswerCorrect: false,
		}
	}

	appleIdToken := event.Request.ChallengeAnswer.(string)

	err = scraps.VerifyAppleIdToken(appleIdToken, applePublicKeys)

	if err != nil {
		return &events.CognitoEventUserPoolsVerifyAuthChallengeResponse{
			AnswerCorrect: false,
		}
	}

	return &events.CognitoEventUserPoolsVerifyAuthChallengeResponse{
		AnswerCorrect: true,
	}
}

func main() {
	lambda.Start(VerifyAuthChallengeResponseLambda)
}
