package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type LambdaNestedStack struct {
	awscdk.NestedStack
}

type LambdaNestedStackProps struct {
	HttpApiId  *string
	HttpApiUrl *string
}

func NewLambdaNestedStack(scope constructs.Construct, id string, props *LambdaNestedStackProps) *LambdaNestedStack {
	nestedStack := awscdk.NewNestedStack(scope, jsii.String(id), nil)

	NewUserLambdaNestedStack(nestedStack, "GetMeLambda", &UserLambdaNestedStackProps{
		HttpApiId:  props.HttpApiId,
		HttpApiUrl: props.HttpApiUrl,
	})

	NewAppointmentLambdaNestedStack(nestedStack, "AppointmentLambdaNestedStack", &AppointmentLambdaNestedStackProps{
		HttpApiId:  props.HttpApiId,
		HttpApiUrl: props.HttpApiUrl,
	})

	return &LambdaNestedStack{
		NestedStack: nestedStack,
	}
}
