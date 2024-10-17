package stacks

import (
	sharedEnums "monospec-api/shared/enums"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type AppleLambda struct {
	awslambda.Function
}

type AppleLambdaProps struct {
	HttpApi awsapigatewayv2.IHttpApi
}

func NewAppleLambda(scope constructs.Construct, id string, props *AppleLambdaProps) *AppleLambda {
	lambda := awslambda.NewFunction(scope, jsii.String(id), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Architecture: awslambda.Architecture_ARM_64(),
		Code:         awslambda.AssetCode_FromAsset(jsii.String("./auth/api/apple/lambda"), nil),
		Handler:      jsii.String("bootstrap"),
		Environment: &map[string]*string{
			sharedEnums.PostgresDatabaseUrl: jsii.String("postgres://postgres:vkwWOJh9610DxDTWisD8K,6-e5BBLN@monospecapistack-rdsnestedstackrdsnest-rds34d05673-lwt37qlx9xe6.cfiwiiwq0xla.eu-central-1.rds.amazonaws.com:5432/monospec"),
			sharedEnums.JWTPrivateKey:       jsii.String("3b1a57d1a7486e4f5a5c257e1ec6d9f1b7c4e53ee44b7d799e9b4b4177f1e4b9"),
		},
	})

	awsapigatewayv2.NewHttpRoute(scope, jsii.String(id+"Route"), &awsapigatewayv2.HttpRouteProps{
		HttpApi:     props.HttpApi,
		RouteKey:    awsapigatewayv2.HttpRouteKey_With(jsii.String("/auth/apple"), awsapigatewayv2.HttpMethod_POST),
		Integration: awsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String(id+"Integration"), lambda, nil),
	})

	return &AppleLambda{
		Function: lambda,
	}
}
