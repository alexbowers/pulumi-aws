// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides information about a Lambda Alias.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/lambda"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := lambda.LookupAlias(ctx, &lambda.LookupAliasArgs{
// 			FunctionName: "my-lambda-func",
// 			Name:         "production",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupAlias(ctx *pulumi.Context, args *LookupAliasArgs, opts ...pulumi.InvokeOption) (*LookupAliasResult, error) {
	var rv LookupAliasResult
	err := ctx.Invoke("aws:lambda/getAlias:getAlias", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAlias.
type LookupAliasArgs struct {
	// Name of the aliased Lambda function.
	FunctionName string `pulumi:"functionName"`
	// Name of the Lambda alias.
	Name string `pulumi:"name"`
}

// A collection of values returned by getAlias.
type LookupAliasResult struct {
	// The Amazon Resource Name (ARN) identifying the Lambda function alias.
	Arn string `pulumi:"arn"`
	// Description of alias.
	Description  string `pulumi:"description"`
	FunctionName string `pulumi:"functionName"`
	// Lambda function version which the alias uses.
	FunctionVersion string `pulumi:"functionVersion"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ARN to be used for invoking Lambda Function from API Gateway - to be used in aws_api_gateway_integration's `uri`.
	InvokeArn string `pulumi:"invokeArn"`
	Name      string `pulumi:"name"`
}
