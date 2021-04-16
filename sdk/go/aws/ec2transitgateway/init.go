// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2transitgateway

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
	case "aws:ec2transitgateway/peeringAttachment:PeeringAttachment":
		r = &PeeringAttachment{}
	case "aws:ec2transitgateway/prefixListReference:PrefixListReference":
		r = &PrefixListReference{}
	case "aws:ec2transitgateway/route:Route":
		r = &Route{}
	case "aws:ec2transitgateway/routeTable:RouteTable":
		r = &RouteTable{}
	case "aws:ec2transitgateway/routeTableAssociation:RouteTableAssociation":
		r = &RouteTableAssociation{}
	case "aws:ec2transitgateway/routeTablePropagation:RouteTablePropagation":
		r = &RouteTablePropagation{}
	case "aws:ec2transitgateway/transitGateway:TransitGateway":
		r = &TransitGateway{}
	case "aws:ec2transitgateway/vpcAttachment:VpcAttachment":
		r = &VpcAttachment{}
	case "aws:ec2transitgateway/vpcAttachmentAccepter:VpcAttachmentAccepter":
		r = &VpcAttachmentAccepter{}
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
		"ec2transitgateway/peeringAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2transitgateway/prefixListReference",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2transitgateway/route",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2transitgateway/routeTable",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2transitgateway/routeTableAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2transitgateway/routeTablePropagation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2transitgateway/transitGateway",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2transitgateway/vpcAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ec2transitgateway/vpcAttachmentAccepter",
		&module{version},
	)
}