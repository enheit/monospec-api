package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"

	"monospec-api/stacks/lambdas"
)

type LambdasNestedStack struct {
	awscdk.NestedStack
}

type LambdasNestedStackProps struct {
	HttpApiId  *string
	HttpApiUrl *string

	UserPool       *awscognito.UserPool
	UserPoolClient *awscognito.UserPoolClient
}

func NewLambdasNestedStack(scope constructs.Construct, id string, props *LambdasNestedStackProps) *LambdasNestedStack {
	nestedStack := awscdk.NewNestedStack(scope, jsii.String(id), nil)

	lambdas.NewUserLambdasNestedStack(nestedStack, "GetMeLambdas", &lambdas.UserLambdasNestedStackProps{
		HttpApiId:  props.HttpApiId,
		HttpApiUrl: props.HttpApiUrl,
	})

	lambdas.NewAppointmentLambdaNestedStack(nestedStack, "AppointmentLambdas", &lambdas.AppointmentLambdasNestedStackProps{
		HttpApiId:  props.HttpApiId,
		HttpApiUrl: props.HttpApiUrl,
	})

	lambdas.NewAuthLambdasNestedStack(nestedStack, "AuthLambdas", &lambdas.AuthLambdasNestedStackProps{
		HttpApiId:  props.HttpApiId,
		HttpApiUrl: props.HttpApiUrl,
	})

	return &LambdasNestedStack{
		NestedStack: nestedStack,
	}
}
