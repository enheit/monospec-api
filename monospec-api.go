package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"monospec-api/stacks"
)

type MonospecApiStackProps struct {
	awscdk.StackProps
}

func NewMonospecApiStack(scope constructs.Construct, id string, props *MonospecApiStackProps) awscdk.Stack {
	var sprops awscdk.StackProps

	if props != nil {
		sprops = props.StackProps
	}

	stack := awscdk.NewStack(scope, &id, &sprops)

	vpc := stacks.NewVpcNestedStack(stack, "Vpc")

	apiGateWay := stacks.NewApiGatewayNestedStack(stack, "ApiGateway")

	stacks.NewRdsNestedStack(stack, "Rds", &stacks.RdsNestedStackProps{
		Vpc: vpc.DefaultVpc,
	})

	stacks.NewLambdasNestedStack(stack, "Lambdas", &stacks.LambdasNestedStackProps{
		HttpApiId:  apiGateWay.HttpApi.HttpApiId(),
		HttpApiUrl: apiGateWay.HttpApi.ApiEndpoint(),
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewMonospecApiStack(app, "MonospecApiStack", &MonospecApiStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String("536697255890"),
		Region:  jsii.String("eu-central-1"),
	}
}
