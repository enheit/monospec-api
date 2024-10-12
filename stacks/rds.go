package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type RdsNestedStack struct {
	awscdk.NestedStack
	Rds awsrds.DatabaseInstance
}

type RdsNestedStackProps struct {
	Vpc awsec2.IVpc
}

func NewRdsNestedStack(scope constructs.Construct, id string, props *RdsNestedStackProps) *RdsNestedStack {
	nestedStack := awscdk.NewNestedStack(scope, jsii.String(id), nil)

	securityGroup := awsec2.NewSecurityGroup(nestedStack, jsii.String("RdsSecurityGroup"), &awsec2.SecurityGroupProps{
		Vpc:         props.Vpc,
		Description: jsii.String("RDS security group"),
	})

	securityGroup.AddIngressRule(awsec2.Peer_AnyIpv4(), awsec2.Port_Tcp(jsii.Number(5432)), jsii.String("Allow access from the internet"), nil)

	rds := awsrds.NewDatabaseInstance(nestedStack, jsii.String("Rds"), &awsrds.DatabaseInstanceProps{
		Engine: awsrds.DatabaseInstanceEngine_Postgres(&awsrds.PostgresInstanceEngineProps{
			Version: awsrds.PostgresEngineVersion_VER_16_3(),
		}),
		InstanceType:       awsec2.InstanceType_Of(awsec2.InstanceClass_BURSTABLE4_GRAVITON, awsec2.InstanceSize_MICRO),
		AllocatedStorage:   jsii.Number(20),
		Vpc:                props.Vpc,
		StorageType:        awsrds.StorageType_GP2,
		BackupRetention:    awscdk.Duration_Days(jsii.Number(0)),
		MultiAz:            jsii.Bool(false),
		RemovalPolicy:      awscdk.RemovalPolicy_DESTROY,
		PubliclyAccessible: jsii.Bool(true),
		VpcSubnets: &awsec2.SubnetSelection{
			SubnetType: awsec2.SubnetType_PUBLIC,
		},
		SecurityGroups: &[]awsec2.ISecurityGroup{
			securityGroup,
		},
		DeletionProtection: jsii.Bool(false),
		Credentials: awsrds.Credentials_FromGeneratedSecret(jsii.String("postgres"), &awsrds.CredentialsBaseOptions{
			SecretName: jsii.String("ms/postgres-password"),
		}),
		DatabaseName: jsii.String("monospec"),
	})

	return &RdsNestedStack{
		NestedStack: nestedStack,
		Rds:         rds,
	}
}
