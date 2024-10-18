package lambdas

import (
	getSpecialistProfile "monospec-api/api/specialist/get-specialist-profile/lambda-config"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2authorizers"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type SpecialistLambdasNestedStack struct {
	awscdk.NestedStack
}

type SpecialistLambdasNestedStackProps struct {
	awscdk.NestedStackProps

	HttpApiId               *string
	HttpApiUrl              *string
	HeadersLambdaAuthorizer *awsapigatewayv2authorizers.HttpLambdaAuthorizer
}

func NewSpecialistLambdasNestedStack(scope constructs.Construct, id string, props *SpecialistLambdasNestedStackProps) *SpecialistLambdasNestedStack {
	nestedStack := awscdk.NewNestedStack(scope, jsii.String(id), &props.NestedStackProps)

	httpApi := awsapigatewayv2.HttpApi_FromHttpApiAttributes(nestedStack, jsii.String("HttpApi"), &awsapigatewayv2.HttpApiAttributes{
		HttpApiId:   props.HttpApiId,
		ApiEndpoint: props.HttpApiUrl,
	})

	getSpecialistProfile.NewGetSpecialistProfileLambda(nestedStack, "GetSpecialistProfile", &getSpecialistProfile.GetSpecialistProfileLambdaProps{
		HttpApi:                 httpApi,
		HeadersLambdaAuthorizer: props.HeadersLambdaAuthorizer,
	})

	return &SpecialistLambdasNestedStack{
		NestedStack: nestedStack,
	}
}
