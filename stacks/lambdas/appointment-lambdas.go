package lambdas

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"

	getUserAppointments "monospec-api/api/appointment/get-user-appointments/lambda-config"
)

type AppointmentLambdaNestedStack struct {
	awscdk.NestedStack
}

type AppointmentLambdasNestedStackProps struct {
	awscdk.NestedStackProps

	HttpApiId  *string
	HttpApiUrl *string
}

func NewAppointmentLambdaNestedStack(scope constructs.Construct, id string, props *AppointmentLambdasNestedStackProps) *AppointmentLambdaNestedStack {
	nestedStack := awscdk.NewNestedStack(scope, jsii.String(id), nil)

	httpApi := awsapigatewayv2.HttpApi_FromHttpApiAttributes(nestedStack, jsii.String("HttpApi"), &awsapigatewayv2.HttpApiAttributes{
		HttpApiId:   props.HttpApiId,
		ApiEndpoint: props.HttpApiUrl,
	})

	getUserAppointments.NewGetUserAppointmentsLambda(nestedStack, "GetUserAppointments", &getUserAppointments.GetUserAppointmentsLambdaProps{
		HttpApi: httpApi,
	})

	return &AppointmentLambdaNestedStack{
		NestedStack: nestedStack,
	}
}
