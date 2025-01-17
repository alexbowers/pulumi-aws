# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ServerEndpointDetails',
    'UserHomeDirectoryMapping',
    'UserPosixProfile',
]

@pulumi.output_type
class ServerEndpointDetails(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "addressAllocationIds":
            suggest = "address_allocation_ids"
        elif key == "subnetIds":
            suggest = "subnet_ids"
        elif key == "vpcEndpointId":
            suggest = "vpc_endpoint_id"
        elif key == "vpcId":
            suggest = "vpc_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ServerEndpointDetails. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ServerEndpointDetails.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ServerEndpointDetails.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 address_allocation_ids: Optional[Sequence[str]] = None,
                 subnet_ids: Optional[Sequence[str]] = None,
                 vpc_endpoint_id: Optional[str] = None,
                 vpc_id: Optional[str] = None):
        """
        :param Sequence[str] address_allocation_ids: A list of address allocation IDs that are required to attach an Elastic IP address to your SFTP server's endpoint. This property can only be used when `endpoint_type` is set to `VPC`.
        :param Sequence[str] subnet_ids: A list of subnet IDs that are required to host your SFTP server endpoint in your VPC. This property can only be used when `endpoint_type` is set to `VPC`.
        :param str vpc_endpoint_id: The ID of the VPC endpoint. This property can only be used when `endpoint_type` is set to `VPC_ENDPOINT`
        :param str vpc_id: The VPC ID of the virtual private cloud in which the SFTP server's endpoint will be hosted. This property can only be used when `endpoint_type` is set to `VPC`.
        """
        if address_allocation_ids is not None:
            pulumi.set(__self__, "address_allocation_ids", address_allocation_ids)
        if subnet_ids is not None:
            pulumi.set(__self__, "subnet_ids", subnet_ids)
        if vpc_endpoint_id is not None:
            pulumi.set(__self__, "vpc_endpoint_id", vpc_endpoint_id)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="addressAllocationIds")
    def address_allocation_ids(self) -> Optional[Sequence[str]]:
        """
        A list of address allocation IDs that are required to attach an Elastic IP address to your SFTP server's endpoint. This property can only be used when `endpoint_type` is set to `VPC`.
        """
        return pulumi.get(self, "address_allocation_ids")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Optional[Sequence[str]]:
        """
        A list of subnet IDs that are required to host your SFTP server endpoint in your VPC. This property can only be used when `endpoint_type` is set to `VPC`.
        """
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter(name="vpcEndpointId")
    def vpc_endpoint_id(self) -> Optional[str]:
        """
        The ID of the VPC endpoint. This property can only be used when `endpoint_type` is set to `VPC_ENDPOINT`
        """
        return pulumi.get(self, "vpc_endpoint_id")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        """
        The VPC ID of the virtual private cloud in which the SFTP server's endpoint will be hosted. This property can only be used when `endpoint_type` is set to `VPC`.
        """
        return pulumi.get(self, "vpc_id")


@pulumi.output_type
class UserHomeDirectoryMapping(dict):
    def __init__(__self__, *,
                 entry: str,
                 target: str):
        """
        :param str entry: Represents an entry and a target.
        :param str target: Represents the map target.
        """
        pulumi.set(__self__, "entry", entry)
        pulumi.set(__self__, "target", target)

    @property
    @pulumi.getter
    def entry(self) -> str:
        """
        Represents an entry and a target.
        """
        return pulumi.get(self, "entry")

    @property
    @pulumi.getter
    def target(self) -> str:
        """
        Represents the map target.
        """
        return pulumi.get(self, "target")


@pulumi.output_type
class UserPosixProfile(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "secondaryGids":
            suggest = "secondary_gids"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in UserPosixProfile. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        UserPosixProfile.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        UserPosixProfile.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 gid: int,
                 uid: int,
                 secondary_gids: Optional[Sequence[int]] = None):
        """
        :param int gid: The POSIX group ID used for all EFS operations by this user.
        :param int uid: The POSIX user ID used for all EFS operations by this user.
        :param Sequence[int] secondary_gids: The secondary POSIX group IDs used for all EFS operations by this user.
        """
        pulumi.set(__self__, "gid", gid)
        pulumi.set(__self__, "uid", uid)
        if secondary_gids is not None:
            pulumi.set(__self__, "secondary_gids", secondary_gids)

    @property
    @pulumi.getter
    def gid(self) -> int:
        """
        The POSIX group ID used for all EFS operations by this user.
        """
        return pulumi.get(self, "gid")

    @property
    @pulumi.getter
    def uid(self) -> int:
        """
        The POSIX user ID used for all EFS operations by this user.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="secondaryGids")
    def secondary_gids(self) -> Optional[Sequence[int]]:
        """
        The secondary POSIX group IDs used for all EFS operations by this user.
        """
        return pulumi.get(self, "secondary_gids")


