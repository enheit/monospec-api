package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type UserLambdaNestedStack struct {
	awscdk.NestedStack
}

type UserLambdaNestedStackProps struct {
	awscdk.NestedStackProps

	HttpApiId  *string
	HttpApiUrl *string
}

func NewUserLambdaNestedStack(scope constructs.Construct, id string, props *UserLambdaNestedStackProps) *UserLambdaNestedStack {
	nestedStack := awscdk.NewNestedStack(scope, jsii.String(id), nil)

	httpApi := awsapigatewayv2.HttpApi_FromHttpApiAttributes(nestedStack, jsii.String("HttpApi"), &awsapigatewayv2.HttpApiAttributes{
		HttpApiId:   props.HttpApiId,
		ApiEndpoint: props.HttpApiUrl,
	})

	NewGetMeLambda(nestedStack, "GetMe", &GetMeLambdaProps{
		httpApi: httpApi,
	})

	return &UserLambdaNestedStack{
		NestedStack: nestedStack,
	}
}
