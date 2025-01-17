// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a list the AWS accounts that are designated as delegated administrators in this organization
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "SERVICE PRINCIPAL"
// 		_, err := organizations.GetDelegatedAdministrators(ctx, &organizations.GetDelegatedAdministratorsArgs{
// 			ServicePrincipal: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetDelegatedAdministrators(ctx *pulumi.Context, args *GetDelegatedAdministratorsArgs, opts ...pulumi.InvokeOption) (*GetDelegatedAdministratorsResult, error) {
	var rv GetDelegatedAdministratorsResult
	err := ctx.Invoke("aws:organizations/getDelegatedAdministrators:getDelegatedAdministrators", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDelegatedAdministrators.
type GetDelegatedAdministratorsArgs struct {
	// Specifies a service principal name. If specified, then the operation lists the delegated administrators only for the specified service. If you don't specify a service principal, the operation lists all delegated administrators for all services in your organization.
	ServicePrincipal *string `pulumi:"servicePrincipal"`
}

// A collection of values returned by getDelegatedAdministrators.
type GetDelegatedAdministratorsResult struct {
	// The list of delegated administrators in your organization, which have the following attributes:
	DelegatedAdministrators []GetDelegatedAdministratorsDelegatedAdministrator `pulumi:"delegatedAdministrators"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ServicePrincipal *string `pulumi:"servicePrincipal"`
}
