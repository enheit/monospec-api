package lambdaconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type GetMeLambda struct {
	awslambda.Function
}

type GetMeLambdaProps struct {
	HttpApi awsapigatewayv2.IHttpApi
}

func NewGetMeLambda(scope constructs.Construct, id string, props *GetMeLambdaProps) *GetMeLambda {
	lambda := awslambda.NewFunction(scope, jsii.String(id), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Architecture: awslambda.Architecture_ARM_64(),
		Code:         awslambda.AssetCode_FromAsset(jsii.String("./api/user/get-me/lambda"), nil),
		Handler:      jsii.String("bootstrap"),
	})

	awsapigatewayv2.NewHttpRoute(scope, jsii.String(id+"Route"), &awsapigatewayv2.HttpRouteProps{
		HttpApi:     props.HttpApi,
		RouteKey:    awsapigatewayv2.HttpRouteKey_With(jsii.String("/me"), awsapigatewayv2.HttpMethod_GET),
		Integration: awsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String(id+"Integration"), lambda, nil),
	})

	return &GetMeLambda{
		Function: lambda,
	}
}
