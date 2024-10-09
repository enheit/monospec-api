package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ApiGatewayNestedStack struct {
	awscdk.NestedStack
	HttpApi awsapigatewayv2.HttpApi
}

func NewApiGatewayNestedStack(scope constructs.Construct, id string) *ApiGatewayNestedStack {
	nestedStack := awscdk.NewNestedStack(scope, jsii.String(id), nil)

	httpApi := awsapigatewayv2.NewHttpApi(nestedStack, jsii.String("HttpApi"), &awsapigatewayv2.HttpApiProps{
		CorsPreflight: &awsapigatewayv2.CorsPreflightOptions{
			AllowMethods: &[]awsapigatewayv2.CorsHttpMethod{
				awsapigatewayv2.CorsHttpMethod_OPTIONS,
				awsapigatewayv2.CorsHttpMethod_GET,
				awsapigatewayv2.CorsHttpMethod_PUT,
				awsapigatewayv2.CorsHttpMethod_POST,
				awsapigatewayv2.CorsHttpMethod_DELETE,
			},
		},
	})

	return &ApiGatewayNestedStack{
		NestedStack: nestedStack,
		HttpApi:     httpApi,
	}
}
