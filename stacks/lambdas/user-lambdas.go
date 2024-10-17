package lambdas

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2authorizers"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"

	getMe "monospec-api/api/user/get-me/lambda-config"
)

type UserLambdasNestedStack struct {
	awscdk.NestedStack
}

type UserLambdasNestedStackProps struct {
	awscdk.NestedStackProps

	HttpApiId               *string
	HttpApiUrl              *string
	HeadersLambdaAuthorizer *awsapigatewayv2authorizers.HttpLambdaAuthorizer
}

func NewUserLambdasNestedStack(scope constructs.Construct, id string, props *UserLambdasNestedStackProps) *UserLambdasNestedStack {
	nestedStack := awscdk.NewNestedStack(scope, jsii.String(id), &props.NestedStackProps)

	httpApi := awsapigatewayv2.HttpApi_FromHttpApiAttributes(nestedStack, jsii.String("HttpApi"), &awsapigatewayv2.HttpApiAttributes{
		HttpApiId:   props.HttpApiId,
		ApiEndpoint: props.HttpApiUrl,
	})

	getMe.NewGetMeLambda(nestedStack, "GetMe", &getMe.GetMeLambdaProps{
		HttpApi:                 httpApi,
		HeadersLambdaAuthorizer: props.HeadersLambdaAuthorizer,
	})

	return &UserLambdasNestedStack{
		NestedStack: nestedStack,
	}
}
