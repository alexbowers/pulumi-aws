// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssoadmin

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an IAM inline policy for a Single Sign-On (SSO) Permission Set resource
//
// > **NOTE:** AWS Single Sign-On (SSO) only supports one IAM inline policy per `ssoadmin.PermissionSet` resource.
// Creating or updating this resource will automatically [Provision the Permission Set](https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_ProvisionPermissionSet.html) to apply the corresponding updates to all assigned accounts.
//
// ## Import
//
// SSO Permission Set Inline Policies can be imported using the `permission_set_arn` and `instance_arn` separated by a comma (`,`) e.g.
//
// ```sh
//  $ pulumi import aws:ssoadmin/permissionSetInlinePolicy:PermissionSetInlinePolicy example arn:aws:sso:::permissionSet/ssoins-2938j0x8920sbj72/ps-80383020jr9302rk,arn:aws:sso:::instance/ssoins-2938j0x8920sbj72
// ```
type PermissionSetInlinePolicy struct {
	pulumi.CustomResourceState

	// The IAM inline policy to attach to a Permission Set.
	InlinePolicy pulumi.StringOutput `pulumi:"inlinePolicy"`
	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	InstanceArn pulumi.StringOutput `pulumi:"instanceArn"`
	// The Amazon Resource Name (ARN) of the Permission Set.
	PermissionSetArn pulumi.StringOutput `pulumi:"permissionSetArn"`
}

// NewPermissionSetInlinePolicy registers a new resource with the given unique name, arguments, and options.
func NewPermissionSetInlinePolicy(ctx *pulumi.Context,
	name string, args *PermissionSetInlinePolicyArgs, opts ...pulumi.ResourceOption) (*PermissionSetInlinePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InlinePolicy == nil {
		return nil, errors.New("invalid value for required argument 'InlinePolicy'")
	}
	if args.InstanceArn == nil {
		return nil, errors.New("invalid value for required argument 'InstanceArn'")
	}
	if args.PermissionSetArn == nil {
		return nil, errors.New("invalid value for required argument 'PermissionSetArn'")
	}
	var resource PermissionSetInlinePolicy
	err := ctx.RegisterResource("aws:ssoadmin/permissionSetInlinePolicy:PermissionSetInlinePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPermissionSetInlinePolicy gets an existing PermissionSetInlinePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPermissionSetInlinePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PermissionSetInlinePolicyState, opts ...pulumi.ResourceOption) (*PermissionSetInlinePolicy, error) {
	var resource PermissionSetInlinePolicy
	err := ctx.ReadResource("aws:ssoadmin/permissionSetInlinePolicy:PermissionSetInlinePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PermissionSetInlinePolicy resources.
type permissionSetInlinePolicyState struct {
	// The IAM inline policy to attach to a Permission Set.
	InlinePolicy *string `pulumi:"inlinePolicy"`
	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	InstanceArn *string `pulumi:"instanceArn"`
	// The Amazon Resource Name (ARN) of the Permission Set.
	PermissionSetArn *string `pulumi:"permissionSetArn"`
}

type PermissionSetInlinePolicyState struct {
	// The IAM inline policy to attach to a Permission Set.
	InlinePolicy pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	InstanceArn pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the Permission Set.
	PermissionSetArn pulumi.StringPtrInput
}

func (PermissionSetInlinePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionSetInlinePolicyState)(nil)).Elem()
}

type permissionSetInlinePolicyArgs struct {
	// The IAM inline policy to attach to a Permission Set.
	InlinePolicy string `pulumi:"inlinePolicy"`
	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	InstanceArn string `pulumi:"instanceArn"`
	// The Amazon Resource Name (ARN) of the Permission Set.
	PermissionSetArn string `pulumi:"permissionSetArn"`
}

// The set of arguments for constructing a PermissionSetInlinePolicy resource.
type PermissionSetInlinePolicyArgs struct {
	// The IAM inline policy to attach to a Permission Set.
	InlinePolicy pulumi.StringInput
	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	InstanceArn pulumi.StringInput
	// The Amazon Resource Name (ARN) of the Permission Set.
	PermissionSetArn pulumi.StringInput
}

func (PermissionSetInlinePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionSetInlinePolicyArgs)(nil)).Elem()
}

type PermissionSetInlinePolicyInput interface {
	pulumi.Input

	ToPermissionSetInlinePolicyOutput() PermissionSetInlinePolicyOutput
	ToPermissionSetInlinePolicyOutputWithContext(ctx context.Context) PermissionSetInlinePolicyOutput
}

func (*PermissionSetInlinePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionSetInlinePolicy)(nil))
}

func (i *PermissionSetInlinePolicy) ToPermissionSetInlinePolicyOutput() PermissionSetInlinePolicyOutput {
	return i.ToPermissionSetInlinePolicyOutputWithContext(context.Background())
}

func (i *PermissionSetInlinePolicy) ToPermissionSetInlinePolicyOutputWithContext(ctx context.Context) PermissionSetInlinePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionSetInlinePolicyOutput)
}

func (i *PermissionSetInlinePolicy) ToPermissionSetInlinePolicyPtrOutput() PermissionSetInlinePolicyPtrOutput {
	return i.ToPermissionSetInlinePolicyPtrOutputWithContext(context.Background())
}

func (i *PermissionSetInlinePolicy) ToPermissionSetInlinePolicyPtrOutputWithContext(ctx context.Context) PermissionSetInlinePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionSetInlinePolicyPtrOutput)
}

type PermissionSetInlinePolicyPtrInput interface {
	pulumi.Input

	ToPermissionSetInlinePolicyPtrOutput() PermissionSetInlinePolicyPtrOutput
	ToPermissionSetInlinePolicyPtrOutputWithContext(ctx context.Context) PermissionSetInlinePolicyPtrOutput
}

type permissionSetInlinePolicyPtrType PermissionSetInlinePolicyArgs

func (*permissionSetInlinePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PermissionSetInlinePolicy)(nil))
}

func (i *permissionSetInlinePolicyPtrType) ToPermissionSetInlinePolicyPtrOutput() PermissionSetInlinePolicyPtrOutput {
	return i.ToPermissionSetInlinePolicyPtrOutputWithContext(context.Background())
}

func (i *permissionSetInlinePolicyPtrType) ToPermissionSetInlinePolicyPtrOutputWithContext(ctx context.Context) PermissionSetInlinePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionSetInlinePolicyPtrOutput)
}

// PermissionSetInlinePolicyArrayInput is an input type that accepts PermissionSetInlinePolicyArray and PermissionSetInlinePolicyArrayOutput values.
// You can construct a concrete instance of `PermissionSetInlinePolicyArrayInput` via:
//
//          PermissionSetInlinePolicyArray{ PermissionSetInlinePolicyArgs{...} }
type PermissionSetInlinePolicyArrayInput interface {
	pulumi.Input

	ToPermissionSetInlinePolicyArrayOutput() PermissionSetInlinePolicyArrayOutput
	ToPermissionSetInlinePolicyArrayOutputWithContext(context.Context) PermissionSetInlinePolicyArrayOutput
}

type PermissionSetInlinePolicyArray []PermissionSetInlinePolicyInput

func (PermissionSetInlinePolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*PermissionSetInlinePolicy)(nil))
}

func (i PermissionSetInlinePolicyArray) ToPermissionSetInlinePolicyArrayOutput() PermissionSetInlinePolicyArrayOutput {
	return i.ToPermissionSetInlinePolicyArrayOutputWithContext(context.Background())
}

func (i PermissionSetInlinePolicyArray) ToPermissionSetInlinePolicyArrayOutputWithContext(ctx context.Context) PermissionSetInlinePolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionSetInlinePolicyArrayOutput)
}

// PermissionSetInlinePolicyMapInput is an input type that accepts PermissionSetInlinePolicyMap and PermissionSetInlinePolicyMapOutput values.
// You can construct a concrete instance of `PermissionSetInlinePolicyMapInput` via:
//
//          PermissionSetInlinePolicyMap{ "key": PermissionSetInlinePolicyArgs{...} }
type PermissionSetInlinePolicyMapInput interface {
	pulumi.Input

	ToPermissionSetInlinePolicyMapOutput() PermissionSetInlinePolicyMapOutput
	ToPermissionSetInlinePolicyMapOutputWithContext(context.Context) PermissionSetInlinePolicyMapOutput
}

type PermissionSetInlinePolicyMap map[string]PermissionSetInlinePolicyInput

func (PermissionSetInlinePolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*PermissionSetInlinePolicy)(nil))
}

func (i PermissionSetInlinePolicyMap) ToPermissionSetInlinePolicyMapOutput() PermissionSetInlinePolicyMapOutput {
	return i.ToPermissionSetInlinePolicyMapOutputWithContext(context.Background())
}

func (i PermissionSetInlinePolicyMap) ToPermissionSetInlinePolicyMapOutputWithContext(ctx context.Context) PermissionSetInlinePolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionSetInlinePolicyMapOutput)
}

type PermissionSetInlinePolicyOutput struct {
	*pulumi.OutputState
}

func (PermissionSetInlinePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionSetInlinePolicy)(nil))
}

func (o PermissionSetInlinePolicyOutput) ToPermissionSetInlinePolicyOutput() PermissionSetInlinePolicyOutput {
	return o
}

func (o PermissionSetInlinePolicyOutput) ToPermissionSetInlinePolicyOutputWithContext(ctx context.Context) PermissionSetInlinePolicyOutput {
	return o
}

func (o PermissionSetInlinePolicyOutput) ToPermissionSetInlinePolicyPtrOutput() PermissionSetInlinePolicyPtrOutput {
	return o.ToPermissionSetInlinePolicyPtrOutputWithContext(context.Background())
}

func (o PermissionSetInlinePolicyOutput) ToPermissionSetInlinePolicyPtrOutputWithContext(ctx context.Context) PermissionSetInlinePolicyPtrOutput {
	return o.ApplyT(func(v PermissionSetInlinePolicy) *PermissionSetInlinePolicy {
		return &v
	}).(PermissionSetInlinePolicyPtrOutput)
}

type PermissionSetInlinePolicyPtrOutput struct {
	*pulumi.OutputState
}

func (PermissionSetInlinePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PermissionSetInlinePolicy)(nil))
}

func (o PermissionSetInlinePolicyPtrOutput) ToPermissionSetInlinePolicyPtrOutput() PermissionSetInlinePolicyPtrOutput {
	return o
}

func (o PermissionSetInlinePolicyPtrOutput) ToPermissionSetInlinePolicyPtrOutputWithContext(ctx context.Context) PermissionSetInlinePolicyPtrOutput {
	return o
}

type PermissionSetInlinePolicyArrayOutput struct{ *pulumi.OutputState }

func (PermissionSetInlinePolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PermissionSetInlinePolicy)(nil))
}

func (o PermissionSetInlinePolicyArrayOutput) ToPermissionSetInlinePolicyArrayOutput() PermissionSetInlinePolicyArrayOutput {
	return o
}

func (o PermissionSetInlinePolicyArrayOutput) ToPermissionSetInlinePolicyArrayOutputWithContext(ctx context.Context) PermissionSetInlinePolicyArrayOutput {
	return o
}

func (o PermissionSetInlinePolicyArrayOutput) Index(i pulumi.IntInput) PermissionSetInlinePolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PermissionSetInlinePolicy {
		return vs[0].([]PermissionSetInlinePolicy)[vs[1].(int)]
	}).(PermissionSetInlinePolicyOutput)
}

type PermissionSetInlinePolicyMapOutput struct{ *pulumi.OutputState }

func (PermissionSetInlinePolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PermissionSetInlinePolicy)(nil))
}

func (o PermissionSetInlinePolicyMapOutput) ToPermissionSetInlinePolicyMapOutput() PermissionSetInlinePolicyMapOutput {
	return o
}

func (o PermissionSetInlinePolicyMapOutput) ToPermissionSetInlinePolicyMapOutputWithContext(ctx context.Context) PermissionSetInlinePolicyMapOutput {
	return o
}

func (o PermissionSetInlinePolicyMapOutput) MapIndex(k pulumi.StringInput) PermissionSetInlinePolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PermissionSetInlinePolicy {
		return vs[0].(map[string]PermissionSetInlinePolicy)[vs[1].(string)]
	}).(PermissionSetInlinePolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(PermissionSetInlinePolicyOutput{})
	pulumi.RegisterOutputType(PermissionSetInlinePolicyPtrOutput{})
	pulumi.RegisterOutputType(PermissionSetInlinePolicyArrayOutput{})
	pulumi.RegisterOutputType(PermissionSetInlinePolicyMapOutput{})
}