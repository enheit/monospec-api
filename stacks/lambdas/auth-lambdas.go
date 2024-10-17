package lambdas

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"

	apple "monospec-api/auth/api/apple/lambda-config"
	logout "monospec-api/auth/api/logout/lambda-config"
)

type AuthLambdasNestedStack struct {
	awscdk.NestedStack
}

type AuthLambdasNestedStackProps struct {
	awscdk.NestedStackProps

	HttpApiId  *string
	HttpApiUrl *string

	UserPool       *awscognito.UserPool
	UserPoolClient *awscognito.UserPoolClient
}

func NewAuthLambdasNestedStack(scope constructs.Construct, id string, props *AuthLambdasNestedStackProps) *AuthLambdasNestedStack {
	nestedStack := awscdk.NewNestedStack(scope, jsii.String(id), &props.NestedStackProps)

	httpApi := awsapigatewayv2.HttpApi_FromHttpApiAttributes(nestedStack, jsii.String("HttpApi"), &awsapigatewayv2.HttpApiAttributes{
		HttpApiId:   props.HttpApiId,
		ApiEndpoint: props.HttpApiUrl,
	})

	apple.NewAppleLambda(nestedStack, "AppleEnter", &apple.AppleLambdaProps{
		HttpApi: httpApi,
	})

	logout.NewLougoutLambda(nestedStack, "Logout", &logout.LogoutLambdaProps{
		HttpApi: httpApi,
	})

	return &AuthLambdasNestedStack{
		NestedStack: nestedStack,
	}
}
