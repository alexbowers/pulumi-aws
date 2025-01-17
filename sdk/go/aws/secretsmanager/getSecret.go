// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package secretsmanager

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve metadata information about a Secrets Manager secret. To retrieve a secret value, see the `secretsmanager.SecretVersion`.
//
// ## Example Usage
// ### ARN
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/secretsmanager"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456"
// 		_, err := secretsmanager.LookupSecret(ctx, &secretsmanager.LookupSecretArgs{
// 			Arn: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Name
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/secretsmanager"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "example"
// 		_, err := secretsmanager.LookupSecret(ctx, &secretsmanager.LookupSecretArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupSecret(ctx *pulumi.Context, args *LookupSecretArgs, opts ...pulumi.InvokeOption) (*LookupSecretResult, error) {
	var rv LookupSecretResult
	err := ctx.Invoke("aws:secretsmanager/getSecret:getSecret", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecret.
type LookupSecretArgs struct {
	// The Amazon Resource Name (ARN) of the secret to retrieve.
	Arn *string `pulumi:"arn"`
	// The name of the secret to retrieve.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getSecret.
type LookupSecretResult struct {
	// The Amazon Resource Name (ARN) of the secret.
	Arn string `pulumi:"arn"`
	// A description of the secret.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Key Management Service (KMS) Customer Master Key (CMK) associated with the secret.
	KmsKeyId string `pulumi:"kmsKeyId"`
	Name     string `pulumi:"name"`
	// The resource-based policy document that's attached to the secret.
	Policy string `pulumi:"policy"`
	// Whether rotation is enabled or not.
	//
	// Deprecated: Use the aws_secretsmanager_secret_rotation data source instead
	RotationEnabled bool `pulumi:"rotationEnabled"`
	// Rotation Lambda function Amazon Resource Name (ARN) if rotation is enabled.
	//
	// Deprecated: Use the aws_secretsmanager_secret_rotation data source instead
	RotationLambdaArn string `pulumi:"rotationLambdaArn"`
	// Rotation rules if rotation is enabled.
	//
	// Deprecated: Use the aws_secretsmanager_secret_rotation data source instead
	RotationRules []GetSecretRotationRule `pulumi:"rotationRules"`
	// Tags of the secret.
	Tags map[string]string `pulumi:"tags"`
}
