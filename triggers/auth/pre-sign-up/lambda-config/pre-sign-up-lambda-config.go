package lambdaconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type PreSignUpLambda struct {
	awslambda.Function
}

func NewPreSignUpLambda(scope constructs.Construct, id string) *PreSignUpLambda {
	lambda := awslambda.NewFunction(scope, jsii.String(id), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Architecture: awslambda.Architecture_ARM_64(),
		Code:         awslambda.AssetCode_FromAsset(jsii.String("./trigers/auth/pre-sign-up/lambda"), nil),
		Handler:      jsii.String("bootstrap"),
	})

	return &PreSignUpLambda{
		Function: lambda,
	}
}
