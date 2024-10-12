package auth

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CreateAuthChallengeLambda struct {
	awslambda.Function
}

func NewCreateAuthChallengeLambda(scope constructs.Construct, id string) *CreateAuthChallengeLambda {
	lambda := awslambda.NewFunction(scope, jsii.String(id), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Architecture: awslambda.Architecture_ARM_64(),
		Code:         awslambda.AssetCode_FromAsset(jsii.String("./trigers/auth/create-auth-challenge/lambda"), nil),
		Handler:      jsii.String("bootstrap"),
	})

	return &CreateAuthChallengeLambda{
		Function: lambda,
	}
}
