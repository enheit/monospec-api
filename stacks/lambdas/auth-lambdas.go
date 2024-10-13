package lambdas

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"

	enter "monospec-api/api/auth/enter/lambda-config"
	logout "monospec-api/api/auth/logout/lambda-config"
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
	nestedStack := awscdk.NewNestedStack(scope, jsii.String(id), nil)

	httpApi := awsapigatewayv2.HttpApi_FromHttpApiAttributes(nestedStack, jsii.String("HttpApi"), &awsapigatewayv2.HttpApiAttributes{
		HttpApiId:   props.HttpApiId,
		ApiEndpoint: props.HttpApiUrl,
	})

	enter.NewEnterLambda(nestedStack, "Enter", &enter.EnterLambdaProps{
		HttpApi: httpApi,
	})

	logout.NewLougoutLambda(nestedStack, "Logout", &logout.LogoutLambdaProps{
		HttpApi: httpApi,
	})

	return &AuthLambdasNestedStack{
		NestedStack: nestedStack,
	}
}
