package lambdaconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2authorizers"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type GetSpecialistProfileLambda struct {
	awslambda.Function
}

type GetSpecialistProfileLambdaProps struct {
	HttpApi                 awsapigatewayv2.IHttpApi
	HeadersLambdaAuthorizer *awsapigatewayv2authorizers.HttpLambdaAuthorizer
}

func NewGetSpecialistProfileLambda(scope constructs.Construct, id string, props *GetSpecialistProfileLambdaProps) *GetSpecialistProfileLambda {
	lambda := awslambda.NewFunction(scope, jsii.String(id), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Architecture: awslambda.Architecture_ARM_64(),
		Code:         awslambda.AssetCode_FromAsset(jsii.String("./api/specialist/get-specialist-profile/lambda"), nil),
		Handler:      jsii.String("bootstrap"),
	})

	awsapigatewayv2.NewHttpRoute(scope, jsii.String(id+"Route"), &awsapigatewayv2.HttpRouteProps{
		HttpApi:     props.HttpApi,
		RouteKey:    awsapigatewayv2.HttpRouteKey_With(jsii.String("/specialists/{specialistId}"), awsapigatewayv2.HttpMethod_GET),
		Integration: awsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String(id+"Integration"), lambda, nil),
		Authorizer:  *props.HeadersLambdaAuthorizer,
	})

	return &GetSpecialistProfileLambda{
		Function: lambda,
	}
}