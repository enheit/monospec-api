package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func DefineAuthChallengeLambda(event events.CognitoEventUserPoolsDefineAuthChallenge) *events.CognitoEventUserPoolsDefineAuthChallengeResponse {
	// Check if this is the first challenge
	if len(event.Request.Session) == 0 {
		// This is the first challenge, so we'll use CUSTOM_CHALLENGE
		return &events.CognitoEventUserPoolsDefineAuthChallengeResponse{
			ChallengeName:      "CUSTOM_CHALLENGE",
			FailAuthentication: false,
			IssueTokens:        false,
		}
	}

	// Check the last challenge
	lastChallenge := event.Request.Session[len(event.Request.Session)-1]

	if lastChallenge.ChallengeName == "CUSTOM_CHALLENGE" {
		if lastChallenge.ChallengeResult {
			// The custom challenge was passed, so we can issue tokens
			return &events.CognitoEventUserPoolsDefineAuthChallengeResponse{
				ChallengeName:      "CUSTOM_CHALLENGE",
				FailAuthentication: false,
				IssueTokens:        true,
			}
		} else {
			// The custom challenge failed, so we fail the authentication
			return &events.CognitoEventUserPoolsDefineAuthChallengeResponse{
				ChallengeName:      "CUSTOM_CHALLENGE",
				FailAuthentication: true,
				IssueTokens:        false,
			}
		}
	} else {
		// Unexpected challenge name, fail the authentication
		return &events.CognitoEventUserPoolsDefineAuthChallengeResponse{
			ChallengeName:      "CUSTOM_CHALLENGE",
			FailAuthentication: true,
			IssueTokens:        false,
		}
	}
}

func main() {
	lambda.Start(DefineAuthChallengeLambda)
}
