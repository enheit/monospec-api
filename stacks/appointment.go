package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type AppointmentLambdaNestedStack struct {
	awscdk.NestedStack
}

type AppointmentLambdaNestedStackProps struct {
	awscdk.NestedStackProps

	HttpApiId  *string
	HttpApiUrl *string
}

func NewAppointmentLambdaNestedStack(scope constructs.Construct, id string, props *AppointmentLambdaNestedStackProps) *AppointmentLambdaNestedStack {
	nestedStack := awscdk.NewNestedStack(scope, jsii.String(id), nil)

	httpApi := awsapigatewayv2.HttpApi_FromHttpApiAttributes(nestedStack, jsii.String("HttpApi"), &awsapigatewayv2.HttpApiAttributes{
		HttpApiId:   props.HttpApiId,
		ApiEndpoint: props.HttpApiUrl,
	})

	NewGetUserAppointmentsLambda(nestedStack, "GetUserAppointments", &GetUserAppointmentsLambdaProps{
		httpApi: httpApi,
	})

	return &AppointmentLambdaNestedStack{
		NestedStack: nestedStack,
	}
}
