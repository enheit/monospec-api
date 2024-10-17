package lambdaconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type VerifyAccessTokenLambda struct {
	awslambda.Function
}

func NewVerifyAccessTokenLambda(scope constructs.Construct, id string) *VerifyAccessTokenLambda {
	lambda := awslambda.NewFunction(scope, jsii.String(id), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Architecture: awslambda.Architecture_ARM_64(),
		Code:         awslambda.AssetCode_FromAsset(jsii.String("./auth/middlewares/verify-access-token/lambda"), nil),
		Handler:      jsii.String("bootstrap"),
		// Environment:  &map[string]*string{},
	})

	return &VerifyAccessTokenLambda{
		Function: lambda,
	}
}
