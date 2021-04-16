// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "aws:lambda/alias:Alias":
		r = &Alias{}
	case "aws:lambda/codeSigningConfig:CodeSigningConfig":
		r = &CodeSigningConfig{}
	case "aws:lambda/eventSourceMapping:EventSourceMapping":
		r = &EventSourceMapping{}
	case "aws:lambda/function:Function":
		r = &Function{}
	case "aws:lambda/functionEventInvokeConfig:FunctionEventInvokeConfig":
		r = &FunctionEventInvokeConfig{}
	case "aws:lambda/layerVersion:LayerVersion":
		r = &LayerVersion{}
	case "aws:lambda/permission:Permission":
		r = &Permission{}
	case "aws:lambda/provisionedConcurrencyConfig:ProvisionedConcurrencyConfig":
		r = &ProvisionedConcurrencyConfig{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := aws.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"aws",
		"lambda/alias",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"lambda/codeSigningConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"lambda/eventSourceMapping",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"lambda/function",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"lambda/functionEventInvokeConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"lambda/layerVersion",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"lambda/permission",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"lambda/provisionedConcurrencyConfig",
		&module{version},
	)
}