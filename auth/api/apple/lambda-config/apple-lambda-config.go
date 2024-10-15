package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type AppleLambda struct {
	awslambda.Function
}

type AppleLambdaProps struct {
	HttpApi awsapigatewayv2.IHttpApi

	UserPool         awscognito.UserPool
	UserPoolClientId string
}

func NewAppleLambda(scope constructs.Construct, id string, props *AppleLambdaProps) *AppleLambda {
	lambda := awslambda.NewFunction(scope, jsii.String(id), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Architecture: awslambda.Architecture_ARM_64(),
		Code:         awslambda.AssetCode_FromAsset(jsii.String("./api/auth/apple/lambda"), nil),
		Handler:      jsii.String("bootstrap"),
		Environment:  &map[string]*string{},
	})

	awsapigatewayv2.NewHttpRoute(scope, jsii.String(id+"Route"), &awsapigatewayv2.HttpRouteProps{
		HttpApi:     props.HttpApi,
		RouteKey:    awsapigatewayv2.HttpRouteKey_With(jsii.String("/auth/apple"), awsapigatewayv2.HttpMethod_POST),
		Integration: awsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String(id+"Integration"), lambda, nil),
	})

	props.UserPool.Grant(lambda, jsii.String("cognito-idp:ListUsers"))

	return &AppleLambda{
		Function: lambda,
	}
}
