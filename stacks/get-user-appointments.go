package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type GetUserAppointmentsLambda struct {
	awslambda.Function
}

type GetUserAppointmentsLambdaProps struct {
	httpApi awsapigatewayv2.IHttpApi
}

func NewGetUserAppointmentsLambda(scope constructs.Construct, id string, props *GetUserAppointmentsLambdaProps) *GetUserAppointmentsLambda {
	lambda := awslambda.NewFunction(scope, jsii.String(id), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Architecture: awslambda.Architecture_ARM_64(),
		Code:         awslambda.AssetCode_FromAsset(jsii.String("./lambdas/get-user-appointments"), nil),
		Handler:      jsii.String("bootstrap"),
	})

	awsapigatewayv2.NewHttpRoute(scope, jsii.String(id+"Route"), &awsapigatewayv2.HttpRouteProps{
		HttpApi:     props.httpApi,
		RouteKey:    awsapigatewayv2.HttpRouteKey_With(jsii.String("/appointments"), awsapigatewayv2.HttpMethod_GET),
		Integration: awsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String(id+"Integration"), lambda, nil),
	})

	return &GetUserAppointmentsLambda{
		Function: lambda,
	}
}
