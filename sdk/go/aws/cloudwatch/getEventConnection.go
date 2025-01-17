// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a EventBridge connection.
//
// > **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudwatch.LookupEventConnection(ctx, &cloudwatch.LookupEventConnectionArgs{
// 			Name: "test",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupEventConnection(ctx *pulumi.Context, args *LookupEventConnectionArgs, opts ...pulumi.InvokeOption) (*LookupEventConnectionResult, error) {
	var rv LookupEventConnectionResult
	err := ctx.Invoke("aws:cloudwatch/getEventConnection:getEventConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEventConnection.
type LookupEventConnectionArgs struct {
	// The name of the connection.
	Name string `pulumi:"name"`
}

// A collection of values returned by getEventConnection.
type LookupEventConnectionResult struct {
	// The ARN (Amazon Resource Name) for the connection.
	Arn string `pulumi:"arn"`
	// The type of authorization to use to connect. One of `API_KEY`,`BASIC`,`OAUTH_CLIENT_CREDENTIALS`.
	AuthorizationType string `pulumi:"authorizationType"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the connection.
	Name string `pulumi:"name"`
	// The ARN (Amazon Resource Name) for the secret created from the authorization parameters specified for the connection.
	SecretArn string `pulumi:"secretArn"`
}
