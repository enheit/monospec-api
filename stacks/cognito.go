package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"

	lambdaconf2 "monospec-api/triggers/auth/create-auth-challenge/lambda-config"
	lambdaconf1 "monospec-api/triggers/auth/define-auth-challenge/lambda-config"
	lambdaconf4 "monospec-api/triggers/auth/pre-sign-up/lambda-config"
	lambdaconf3 "monospec-api/triggers/auth/verify-auth-challenge-response/lambda-config"
)

type CongitoNestedStack struct {
	awscdk.NestedStack

	UserPool       *awscognito.UserPool
	UserPoolClient *awscognito.UserPoolClient
}

func NewCognitoNestedStack(scope constructs.Construct, id string, props *awscdk.NestedStackProps) *CongitoNestedStack {
	nestedStack := awscdk.NewNestedStack(scope, jsii.String(id), props)

	userPool := awscognito.NewUserPool(nestedStack, jsii.String("UserPool"), &awscognito.UserPoolProps{
		SelfSignUpEnabled: jsii.Bool(true),
		RemovalPolicy:     awscdk.RemovalPolicy_DESTROY,
		LambdaTriggers: &awscognito.UserPoolTriggers{
			DefineAuthChallenge:         lambdaconf1.NewDefineAuthChallengeLambda(nestedStack, "DefineAuthChallengeLambda"),
			CreateAuthChallenge:         lambdaconf2.NewCreateAuthChallengeLambda(nestedStack, "CreateAuthChallengeLambda"),
			VerifyAuthChallengeResponse: lambdaconf3.NewVerifyAuthChallengeResponseLambda(nestedStack, "VerifyAuthChallengeResponseLambda"),
			PreSignUp:                   lambdaconf4.NewPreSignUpLambda(nestedStack, "PreSignUpLambda"),
		},
	})

	userPoolClient := awscognito.NewUserPoolClient(nestedStack, jsii.String("UserPoolClient"), &awscognito.UserPoolClientProps{
		UserPool: userPool,
	})

	return &CongitoNestedStack{
		NestedStack:    nestedStack,
		UserPool:       &userPool,
		UserPoolClient: &userPoolClient,
	}
}
