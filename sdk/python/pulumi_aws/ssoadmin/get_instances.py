# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetInstancesResult',
    'AwaitableGetInstancesResult',
    'get_instances',
]

@pulumi.output_type
class GetInstancesResult:
    """
    A collection of values returned by getInstances.
    """
    def __init__(__self__, arns=None, id=None, identity_store_ids=None):
        if arns and not isinstance(arns, list):
            raise TypeError("Expected argument 'arns' to be a list")
        pulumi.set(__self__, "arns", arns)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identity_store_ids and not isinstance(identity_store_ids, list):
            raise TypeError("Expected argument 'identity_store_ids' to be a list")
        pulumi.set(__self__, "identity_store_ids", identity_store_ids)

    @property
    @pulumi.getter
    def arns(self) -> Sequence[str]:
        """
        Set of Amazon Resource Names (ARNs) of the SSO Instances.
        """
        return pulumi.get(self, "arns")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="identityStoreIds")
    def identity_store_ids(self) -> Sequence[str]:
        """
        Set of identifiers of the identity stores connected to the SSO Instances.
        """
        return pulumi.get(self, "identity_store_ids")


class AwaitableGetInstancesResult(GetInstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstancesResult(
            arns=self.arns,
            id=self.id,
            identity_store_ids=self.identity_store_ids)


def get_instances(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstancesResult:
    """
    Use this data source to get ARNs and Identity Store IDs of Single Sign-On (SSO) Instances.
    """
    __args__ = dict()
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ssoadmin/getInstances:getInstances', __args__, opts=opts, typ=GetInstancesResult).value

    return AwaitableGetInstancesResult(
        arns=__ret__.arns,
        id=__ret__.id,
        identity_store_ids=__ret__.identity_store_ids)