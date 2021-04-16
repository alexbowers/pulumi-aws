# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ResolverDnsSecConfigArgs', 'ResolverDnsSecConfig']

@pulumi.input_type
class ResolverDnsSecConfigArgs:
    def __init__(__self__, *,
                 resource_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a ResolverDnsSecConfig resource.
        :param pulumi.Input[str] resource_id: The ID of the virtual private cloud (VPC) that you're updating the DNSSEC validation status for.
        """
        pulumi.set(__self__, "resource_id", resource_id)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Input[str]:
        """
        The ID of the virtual private cloud (VPC) that you're updating the DNSSEC validation status for.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_id", value)


@pulumi.input_type
class _ResolverDnsSecConfigState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 owner_id: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 validation_status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ResolverDnsSecConfig resources.
        :param pulumi.Input[str] arn: The ARN for a configuration for DNSSEC validation.
        :param pulumi.Input[str] owner_id: The owner account ID of the virtual private cloud (VPC) for a configuration for DNSSEC validation.
        :param pulumi.Input[str] resource_id: The ID of the virtual private cloud (VPC) that you're updating the DNSSEC validation status for.
        :param pulumi.Input[str] validation_status: The validation status for a DNSSEC configuration. The status can be one of the following: `ENABLING`, `ENABLED`, `DISABLING` and `DISABLED`.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if owner_id is not None:
            pulumi.set(__self__, "owner_id", owner_id)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if validation_status is not None:
            pulumi.set(__self__, "validation_status", validation_status)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN for a configuration for DNSSEC validation.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> Optional[pulumi.Input[str]]:
        """
        The owner account ID of the virtual private cloud (VPC) for a configuration for DNSSEC validation.
        """
        return pulumi.get(self, "owner_id")

    @owner_id.setter
    def owner_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner_id", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the virtual private cloud (VPC) that you're updating the DNSSEC validation status for.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="validationStatus")
    def validation_status(self) -> Optional[pulumi.Input[str]]:
        """
        The validation status for a DNSSEC configuration. The status can be one of the following: `ENABLING`, `ENABLED`, `DISABLING` and `DISABLED`.
        """
        return pulumi.get(self, "validation_status")

    @validation_status.setter
    def validation_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "validation_status", value)


class ResolverDnsSecConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Route 53 Resolver DNSSEC config resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_vpc = aws.ec2.Vpc("exampleVpc",
            cidr_block="10.0.0.0/16",
            enable_dns_support=True,
            enable_dns_hostnames=True)
        example_resolver_dns_sec_config = aws.route53.ResolverDnsSecConfig("exampleResolverDnsSecConfig", resource_id=example_vpc.id)
        ```

        ## Import

         Route 53 Resolver DNSSEC configs can be imported using the Route 53 Resolver DNSSEC config ID, e.g.

        ```sh
         $ pulumi import aws:route53/resolverDnsSecConfig:ResolverDnsSecConfig example rdsc-be1866ecc1683e95
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] resource_id: The ID of the virtual private cloud (VPC) that you're updating the DNSSEC validation status for.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ResolverDnsSecConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Route 53 Resolver DNSSEC config resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_vpc = aws.ec2.Vpc("exampleVpc",
            cidr_block="10.0.0.0/16",
            enable_dns_support=True,
            enable_dns_hostnames=True)
        example_resolver_dns_sec_config = aws.route53.ResolverDnsSecConfig("exampleResolverDnsSecConfig", resource_id=example_vpc.id)
        ```

        ## Import

         Route 53 Resolver DNSSEC configs can be imported using the Route 53 Resolver DNSSEC config ID, e.g.

        ```sh
         $ pulumi import aws:route53/resolverDnsSecConfig:ResolverDnsSecConfig example rdsc-be1866ecc1683e95
        ```

        :param str resource_name: The name of the resource.
        :param ResolverDnsSecConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResolverDnsSecConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ResolverDnsSecConfigArgs.__new__(ResolverDnsSecConfigArgs)

            if resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_id'")
            __props__.__dict__["resource_id"] = resource_id
            __props__.__dict__["arn"] = None
            __props__.__dict__["owner_id"] = None
            __props__.__dict__["validation_status"] = None
        super(ResolverDnsSecConfig, __self__).__init__(
            'aws:route53/resolverDnsSecConfig:ResolverDnsSecConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            owner_id: Optional[pulumi.Input[str]] = None,
            resource_id: Optional[pulumi.Input[str]] = None,
            validation_status: Optional[pulumi.Input[str]] = None) -> 'ResolverDnsSecConfig':
        """
        Get an existing ResolverDnsSecConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN for a configuration for DNSSEC validation.
        :param pulumi.Input[str] owner_id: The owner account ID of the virtual private cloud (VPC) for a configuration for DNSSEC validation.
        :param pulumi.Input[str] resource_id: The ID of the virtual private cloud (VPC) that you're updating the DNSSEC validation status for.
        :param pulumi.Input[str] validation_status: The validation status for a DNSSEC configuration. The status can be one of the following: `ENABLING`, `ENABLED`, `DISABLING` and `DISABLED`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ResolverDnsSecConfigState.__new__(_ResolverDnsSecConfigState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["owner_id"] = owner_id
        __props__.__dict__["resource_id"] = resource_id
        __props__.__dict__["validation_status"] = validation_status
        return ResolverDnsSecConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN for a configuration for DNSSEC validation.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Output[str]:
        """
        The owner account ID of the virtual private cloud (VPC) for a configuration for DNSSEC validation.
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[str]:
        """
        The ID of the virtual private cloud (VPC) that you're updating the DNSSEC validation status for.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="validationStatus")
    def validation_status(self) -> pulumi.Output[str]:
        """
        The validation status for a DNSSEC configuration. The status can be one of the following: `ENABLING`, `ENABLED`, `DISABLING` and `DISABLED`.
        """
        return pulumi.get(self, "validation_status")
