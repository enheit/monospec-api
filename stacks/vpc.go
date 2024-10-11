package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type VpcNestedStack struct {
	awscdk.NestedStack
	DefaultVpc awsec2.IVpc
	// CustomVpc  awsec2.IVpc
}

func NewVpcNestedStack(scope constructs.Construct, id string) *VpcNestedStack {
	nestedStack := awscdk.NewNestedStack(scope, jsii.String(id), nil)

	defaultVpc := awsec2.Vpc_FromLookup(nestedStack, jsii.String("DefaultVpc"), &awsec2.VpcLookupOptions{
		IsDefault: jsii.Bool(true),
	})

	// customVpc := awsec2.NewVpc(nestedStack, jsii.String("CustomVpc"), &awsec2.VpcProps{
	// 	IpAddresses: awsec2.IpAddresses_Cidr(jsii.String("10.0.0.0/16")),
	// 	SubnetConfiguration: &[]*awsec2.SubnetConfiguration{
	// 		{
	// 			Name:       jsii.String("PublicSubnet"),
	// 			SubnetType: awsec2.SubnetType_PUBLIC,
	// 			CidrMask:   jsii.Number(24), // 10.0.0.0/24
	// 		},
	//      {
	//        Name:       jsii.String("PrivateSubnet"),
	//        SubnetType: awsec2.SubnetType_PRIVATE_WITH_NAT,
	//        CidrMask:   jsii.Number(24), // 10.0.1.0/24
	//      },
	// 	},
	// })

	return &VpcNestedStack{
		NestedStack: nestedStack,
		DefaultVpc:  defaultVpc,
		// CustomVpc:   customVpc,
	}
}
