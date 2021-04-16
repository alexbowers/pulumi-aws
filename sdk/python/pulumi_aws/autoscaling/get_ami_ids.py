# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetAmiIdsResult',
    'AwaitableGetAmiIdsResult',
    'get_ami_ids',
]

@pulumi.output_type
class GetAmiIdsResult:
    """
    A collection of values returned by getAmiIds.
    """
    def __init__(__self__, arns=None, filters=None, id=None, names=None):
        if arns and not isinstance(arns, list):
            raise TypeError("Expected argument 'arns' to be a list")
        pulumi.set(__self__, "arns", arns)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)

    @property
    @pulumi.getter
    def arns(self) -> Sequence[str]:
        """
        A list of the Autoscaling Groups Arns in the current region.
        """
        return pulumi.get(self, "arns")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetAmiIdsFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        A list of the Autoscaling Groups in the current region.
        """
        return pulumi.get(self, "names")


class AwaitableGetAmiIdsResult(GetAmiIdsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAmiIdsResult(
            arns=self.arns,
            filters=self.filters,
            id=self.id,
            names=self.names)


def get_ami_ids(filters: Optional[Sequence[pulumi.InputType['GetAmiIdsFilterArgs']]] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAmiIdsResult:
    """
    The Autoscaling Groups data source allows access to the list of AWS
    ASGs within a specific region. This will allow you to pass a list of AutoScaling Groups to other resources.


    :param Sequence[pulumi.InputType['GetAmiIdsFilterArgs']] filters: A filter used to scope the list e.g. by tags. See [related docs](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_Filter.html).
    """
    __args__ = dict()
    __args__['filters'] = filters
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:autoscaling/getAmiIds:getAmiIds', __args__, opts=opts, typ=GetAmiIdsResult).value

    return AwaitableGetAmiIdsResult(
        arns=__ret__.arns,
        filters=__ret__.filters,
        id=__ret__.id,
        names=__ret__.names)