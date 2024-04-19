# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ReplicationConfigArgs', 'ReplicationConfig']

@pulumi.input_type
class ReplicationConfigArgs:
    def __init__(__self__, *,
                 cron_exp: pulumi.Input[str],
                 repo_key: pulumi.Input[str],
                 enable_event_replication: Optional[pulumi.Input[bool]] = None,
                 replications: Optional[pulumi.Input[Sequence[pulumi.Input['ReplicationConfigReplicationArgs']]]] = None):
        """
        The set of arguments for constructing a ReplicationConfig resource.
        :param pulumi.Input[str] cron_exp: Cron expression to control the operation frequency.
        """
        pulumi.set(__self__, "cron_exp", cron_exp)
        pulumi.set(__self__, "repo_key", repo_key)
        if enable_event_replication is not None:
            pulumi.set(__self__, "enable_event_replication", enable_event_replication)
        if replications is not None:
            pulumi.set(__self__, "replications", replications)

    @property
    @pulumi.getter(name="cronExp")
    def cron_exp(self) -> pulumi.Input[str]:
        """
        Cron expression to control the operation frequency.
        """
        return pulumi.get(self, "cron_exp")

    @cron_exp.setter
    def cron_exp(self, value: pulumi.Input[str]):
        pulumi.set(self, "cron_exp", value)

    @property
    @pulumi.getter(name="repoKey")
    def repo_key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "repo_key")

    @repo_key.setter
    def repo_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "repo_key", value)

    @property
    @pulumi.getter(name="enableEventReplication")
    def enable_event_replication(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enable_event_replication")

    @enable_event_replication.setter
    def enable_event_replication(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_event_replication", value)

    @property
    @pulumi.getter
    def replications(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ReplicationConfigReplicationArgs']]]]:
        return pulumi.get(self, "replications")

    @replications.setter
    def replications(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ReplicationConfigReplicationArgs']]]]):
        pulumi.set(self, "replications", value)


@pulumi.input_type
class _ReplicationConfigState:
    def __init__(__self__, *,
                 cron_exp: Optional[pulumi.Input[str]] = None,
                 enable_event_replication: Optional[pulumi.Input[bool]] = None,
                 replications: Optional[pulumi.Input[Sequence[pulumi.Input['ReplicationConfigReplicationArgs']]]] = None,
                 repo_key: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ReplicationConfig resources.
        :param pulumi.Input[str] cron_exp: Cron expression to control the operation frequency.
        """
        if cron_exp is not None:
            pulumi.set(__self__, "cron_exp", cron_exp)
        if enable_event_replication is not None:
            pulumi.set(__self__, "enable_event_replication", enable_event_replication)
        if replications is not None:
            pulumi.set(__self__, "replications", replications)
        if repo_key is not None:
            pulumi.set(__self__, "repo_key", repo_key)

    @property
    @pulumi.getter(name="cronExp")
    def cron_exp(self) -> Optional[pulumi.Input[str]]:
        """
        Cron expression to control the operation frequency.
        """
        return pulumi.get(self, "cron_exp")

    @cron_exp.setter
    def cron_exp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cron_exp", value)

    @property
    @pulumi.getter(name="enableEventReplication")
    def enable_event_replication(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enable_event_replication")

    @enable_event_replication.setter
    def enable_event_replication(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_event_replication", value)

    @property
    @pulumi.getter
    def replications(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ReplicationConfigReplicationArgs']]]]:
        return pulumi.get(self, "replications")

    @replications.setter
    def replications(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ReplicationConfigReplicationArgs']]]]):
        pulumi.set(self, "replications", value)

    @property
    @pulumi.getter(name="repoKey")
    def repo_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "repo_key")

    @repo_key.setter
    def repo_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repo_key", value)


class ReplicationConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cron_exp: Optional[pulumi.Input[str]] = None,
                 enable_event_replication: Optional[pulumi.Input[bool]] = None,
                 replications: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReplicationConfigReplicationArgs']]]]] = None,
                 repo_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        > This resource is deprecated in favor of `PushReplication` resource.

        Provides an Artifactory replication config resource. This can be used to create and manage Artifactory replications.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        # Create a replication between two artifactory local repositories
        provider_test_source = artifactory.LocalMavenRepository("provider_test_source", key="provider_test_source")
        provider_test_dest = artifactory.LocalMavenRepository("provider_test_dest", key="provider_test_dest")
        foo_rep = artifactory.ReplicationConfig("foo-rep",
            repo_key=provider_test_source.key,
            cron_exp="0 0 * * * ?",
            enable_event_replication=True,
            replications=[artifactory.ReplicationConfigReplicationArgs(
                url="$var.artifactory_url",
                username="$var.artifactory_username",
                password="$var.artifactory_password",
            )])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Replication configs can be imported using their repo key, e.g.

        ```sh
        $ pulumi import artifactory:index/replicationConfig:ReplicationConfig foo-rep provider_test_source
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cron_exp: Cron expression to control the operation frequency.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReplicationConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        > This resource is deprecated in favor of `PushReplication` resource.

        Provides an Artifactory replication config resource. This can be used to create and manage Artifactory replications.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        # Create a replication between two artifactory local repositories
        provider_test_source = artifactory.LocalMavenRepository("provider_test_source", key="provider_test_source")
        provider_test_dest = artifactory.LocalMavenRepository("provider_test_dest", key="provider_test_dest")
        foo_rep = artifactory.ReplicationConfig("foo-rep",
            repo_key=provider_test_source.key,
            cron_exp="0 0 * * * ?",
            enable_event_replication=True,
            replications=[artifactory.ReplicationConfigReplicationArgs(
                url="$var.artifactory_url",
                username="$var.artifactory_username",
                password="$var.artifactory_password",
            )])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Replication configs can be imported using their repo key, e.g.

        ```sh
        $ pulumi import artifactory:index/replicationConfig:ReplicationConfig foo-rep provider_test_source
        ```

        :param str resource_name: The name of the resource.
        :param ReplicationConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReplicationConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cron_exp: Optional[pulumi.Input[str]] = None,
                 enable_event_replication: Optional[pulumi.Input[bool]] = None,
                 replications: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReplicationConfigReplicationArgs']]]]] = None,
                 repo_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ReplicationConfigArgs.__new__(ReplicationConfigArgs)

            if cron_exp is None and not opts.urn:
                raise TypeError("Missing required property 'cron_exp'")
            __props__.__dict__["cron_exp"] = cron_exp
            __props__.__dict__["enable_event_replication"] = enable_event_replication
            __props__.__dict__["replications"] = replications
            if repo_key is None and not opts.urn:
                raise TypeError("Missing required property 'repo_key'")
            __props__.__dict__["repo_key"] = repo_key
        super(ReplicationConfig, __self__).__init__(
            'artifactory:index/replicationConfig:ReplicationConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cron_exp: Optional[pulumi.Input[str]] = None,
            enable_event_replication: Optional[pulumi.Input[bool]] = None,
            replications: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReplicationConfigReplicationArgs']]]]] = None,
            repo_key: Optional[pulumi.Input[str]] = None) -> 'ReplicationConfig':
        """
        Get an existing ReplicationConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cron_exp: Cron expression to control the operation frequency.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ReplicationConfigState.__new__(_ReplicationConfigState)

        __props__.__dict__["cron_exp"] = cron_exp
        __props__.__dict__["enable_event_replication"] = enable_event_replication
        __props__.__dict__["replications"] = replications
        __props__.__dict__["repo_key"] = repo_key
        return ReplicationConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cronExp")
    def cron_exp(self) -> pulumi.Output[str]:
        """
        Cron expression to control the operation frequency.
        """
        return pulumi.get(self, "cron_exp")

    @property
    @pulumi.getter(name="enableEventReplication")
    def enable_event_replication(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "enable_event_replication")

    @property
    @pulumi.getter
    def replications(self) -> pulumi.Output[Optional[Sequence['outputs.ReplicationConfigReplication']]]:
        return pulumi.get(self, "replications")

    @property
    @pulumi.getter(name="repoKey")
    def repo_key(self) -> pulumi.Output[str]:
        return pulumi.get(self, "repo_key")

