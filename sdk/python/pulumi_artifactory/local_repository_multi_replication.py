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

__all__ = ['LocalRepositoryMultiReplicationArgs', 'LocalRepositoryMultiReplication']

@pulumi.input_type
class LocalRepositoryMultiReplicationArgs:
    def __init__(__self__, *,
                 cron_exp: pulumi.Input[str],
                 repo_key: pulumi.Input[str],
                 enable_event_replication: Optional[pulumi.Input[bool]] = None,
                 replications: Optional[pulumi.Input[Sequence[pulumi.Input['LocalRepositoryMultiReplicationReplicationArgs']]]] = None):
        """
        The set of arguments for constructing a LocalRepositoryMultiReplication resource.
        :param pulumi.Input[str] cron_exp: A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
        :param pulumi.Input[str] repo_key: Repository name.
        :param pulumi.Input[bool] enable_event_replication: When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is `false`.
        :param pulumi.Input[Sequence[pulumi.Input['LocalRepositoryMultiReplicationReplicationArgs']]] replications: List of replications minimum 1 element.
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
        A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
        """
        return pulumi.get(self, "cron_exp")

    @cron_exp.setter
    def cron_exp(self, value: pulumi.Input[str]):
        pulumi.set(self, "cron_exp", value)

    @property
    @pulumi.getter(name="repoKey")
    def repo_key(self) -> pulumi.Input[str]:
        """
        Repository name.
        """
        return pulumi.get(self, "repo_key")

    @repo_key.setter
    def repo_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "repo_key", value)

    @property
    @pulumi.getter(name="enableEventReplication")
    def enable_event_replication(self) -> Optional[pulumi.Input[bool]]:
        """
        When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is `false`.
        """
        return pulumi.get(self, "enable_event_replication")

    @enable_event_replication.setter
    def enable_event_replication(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_event_replication", value)

    @property
    @pulumi.getter
    def replications(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LocalRepositoryMultiReplicationReplicationArgs']]]]:
        """
        List of replications minimum 1 element.
        """
        return pulumi.get(self, "replications")

    @replications.setter
    def replications(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LocalRepositoryMultiReplicationReplicationArgs']]]]):
        pulumi.set(self, "replications", value)


@pulumi.input_type
class _LocalRepositoryMultiReplicationState:
    def __init__(__self__, *,
                 cron_exp: Optional[pulumi.Input[str]] = None,
                 enable_event_replication: Optional[pulumi.Input[bool]] = None,
                 replications: Optional[pulumi.Input[Sequence[pulumi.Input['LocalRepositoryMultiReplicationReplicationArgs']]]] = None,
                 repo_key: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LocalRepositoryMultiReplication resources.
        :param pulumi.Input[str] cron_exp: A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
        :param pulumi.Input[bool] enable_event_replication: When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is `false`.
        :param pulumi.Input[Sequence[pulumi.Input['LocalRepositoryMultiReplicationReplicationArgs']]] replications: List of replications minimum 1 element.
        :param pulumi.Input[str] repo_key: Repository name.
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
        A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
        """
        return pulumi.get(self, "cron_exp")

    @cron_exp.setter
    def cron_exp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cron_exp", value)

    @property
    @pulumi.getter(name="enableEventReplication")
    def enable_event_replication(self) -> Optional[pulumi.Input[bool]]:
        """
        When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is `false`.
        """
        return pulumi.get(self, "enable_event_replication")

    @enable_event_replication.setter
    def enable_event_replication(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_event_replication", value)

    @property
    @pulumi.getter
    def replications(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LocalRepositoryMultiReplicationReplicationArgs']]]]:
        """
        List of replications minimum 1 element.
        """
        return pulumi.get(self, "replications")

    @replications.setter
    def replications(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LocalRepositoryMultiReplicationReplicationArgs']]]]):
        pulumi.set(self, "replications", value)

    @property
    @pulumi.getter(name="repoKey")
    def repo_key(self) -> Optional[pulumi.Input[str]]:
        """
        Repository name.
        """
        return pulumi.get(self, "repo_key")

    @repo_key.setter
    def repo_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repo_key", value)


class LocalRepositoryMultiReplication(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cron_exp: Optional[pulumi.Input[str]] = None,
                 enable_event_replication: Optional[pulumi.Input[bool]] = None,
                 replications: Optional[pulumi.Input[Sequence[pulumi.Input[Union['LocalRepositoryMultiReplicationReplicationArgs', 'LocalRepositoryMultiReplicationReplicationArgsDict']]]]] = None,
                 repo_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a local repository replication resource, also referred to as Artifactory push replication. This can be used to create and manage Artifactory local repository replications using [Multi-push Replication API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-CreateorReplaceLocalMulti-pushReplication).

        Push replication is used to synchronize Local Repositories, and is implemented by the Artifactory server on the near end invoking a synchronization of artifacts to the far end.

        See the [Official Documentation](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-PushReplication).

        This resource replaces `PushReplication` and used to create a replication of one local repository to multiple repositories on the remote server.

        > This resource requires Artifactory Enterprise license. Use `LocalRepositorySingleReplication` with other licenses.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        config = pulumi.Config()
        # The base URL of the Artifactory deployment
        artifactory_url = config.require("artifactoryUrl")
        # The username for the Artifactory
        artifactory_username = config.require("artifactoryUsername")
        # The password for the Artifactory
        artifactory_password = config.require("artifactoryPassword")
        # Create a replication between two artifactory local repositories
        provider_test_source = artifactory.LocalMavenRepository("provider_test_source", key="provider_test_source")
        provider_test_dest = artifactory.LocalMavenRepository("provider_test_dest", key="provider_test_dest")
        provider_test_dest1 = artifactory.LocalMavenRepository("provider_test_dest1", key="provider_test_dest1")
        foo_rep = artifactory.LocalRepositoryMultiReplication("foo-rep",
            repo_key=provider_test_source.key,
            cron_exp="0 0 * * * ?",
            enable_event_replication=True,
            replications=[
                {
                    "url": provider_test_dest.key.apply(lambda key: f"{artifactory_url}/artifactory/{key}"),
                    "username": "$var.artifactory_username",
                    "password": "$var.artifactory_password",
                    "enabled": True,
                },
                {
                    "url": provider_test_dest1.key.apply(lambda key: f"{artifactory_url}/artifactory/{key}"),
                    "username": "$var.artifactory_username",
                    "password": "$var.artifactory_password",
                    "enabled": True,
                },
            ])
        ```

        ## Import

        Push replication configs can be imported using their repo key, e.g.

        ```sh
        $ pulumi import artifactory:index/localRepositoryMultiReplication:LocalRepositoryMultiReplication foo-rep provider_test_source
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cron_exp: A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
        :param pulumi.Input[bool] enable_event_replication: When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is `false`.
        :param pulumi.Input[Sequence[pulumi.Input[Union['LocalRepositoryMultiReplicationReplicationArgs', 'LocalRepositoryMultiReplicationReplicationArgsDict']]]] replications: List of replications minimum 1 element.
        :param pulumi.Input[str] repo_key: Repository name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LocalRepositoryMultiReplicationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a local repository replication resource, also referred to as Artifactory push replication. This can be used to create and manage Artifactory local repository replications using [Multi-push Replication API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-CreateorReplaceLocalMulti-pushReplication).

        Push replication is used to synchronize Local Repositories, and is implemented by the Artifactory server on the near end invoking a synchronization of artifacts to the far end.

        See the [Official Documentation](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-PushReplication).

        This resource replaces `PushReplication` and used to create a replication of one local repository to multiple repositories on the remote server.

        > This resource requires Artifactory Enterprise license. Use `LocalRepositorySingleReplication` with other licenses.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        config = pulumi.Config()
        # The base URL of the Artifactory deployment
        artifactory_url = config.require("artifactoryUrl")
        # The username for the Artifactory
        artifactory_username = config.require("artifactoryUsername")
        # The password for the Artifactory
        artifactory_password = config.require("artifactoryPassword")
        # Create a replication between two artifactory local repositories
        provider_test_source = artifactory.LocalMavenRepository("provider_test_source", key="provider_test_source")
        provider_test_dest = artifactory.LocalMavenRepository("provider_test_dest", key="provider_test_dest")
        provider_test_dest1 = artifactory.LocalMavenRepository("provider_test_dest1", key="provider_test_dest1")
        foo_rep = artifactory.LocalRepositoryMultiReplication("foo-rep",
            repo_key=provider_test_source.key,
            cron_exp="0 0 * * * ?",
            enable_event_replication=True,
            replications=[
                {
                    "url": provider_test_dest.key.apply(lambda key: f"{artifactory_url}/artifactory/{key}"),
                    "username": "$var.artifactory_username",
                    "password": "$var.artifactory_password",
                    "enabled": True,
                },
                {
                    "url": provider_test_dest1.key.apply(lambda key: f"{artifactory_url}/artifactory/{key}"),
                    "username": "$var.artifactory_username",
                    "password": "$var.artifactory_password",
                    "enabled": True,
                },
            ])
        ```

        ## Import

        Push replication configs can be imported using their repo key, e.g.

        ```sh
        $ pulumi import artifactory:index/localRepositoryMultiReplication:LocalRepositoryMultiReplication foo-rep provider_test_source
        ```

        :param str resource_name: The name of the resource.
        :param LocalRepositoryMultiReplicationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LocalRepositoryMultiReplicationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cron_exp: Optional[pulumi.Input[str]] = None,
                 enable_event_replication: Optional[pulumi.Input[bool]] = None,
                 replications: Optional[pulumi.Input[Sequence[pulumi.Input[Union['LocalRepositoryMultiReplicationReplicationArgs', 'LocalRepositoryMultiReplicationReplicationArgsDict']]]]] = None,
                 repo_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LocalRepositoryMultiReplicationArgs.__new__(LocalRepositoryMultiReplicationArgs)

            if cron_exp is None and not opts.urn:
                raise TypeError("Missing required property 'cron_exp'")
            __props__.__dict__["cron_exp"] = cron_exp
            __props__.__dict__["enable_event_replication"] = enable_event_replication
            __props__.__dict__["replications"] = replications
            if repo_key is None and not opts.urn:
                raise TypeError("Missing required property 'repo_key'")
            __props__.__dict__["repo_key"] = repo_key
        super(LocalRepositoryMultiReplication, __self__).__init__(
            'artifactory:index/localRepositoryMultiReplication:LocalRepositoryMultiReplication',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cron_exp: Optional[pulumi.Input[str]] = None,
            enable_event_replication: Optional[pulumi.Input[bool]] = None,
            replications: Optional[pulumi.Input[Sequence[pulumi.Input[Union['LocalRepositoryMultiReplicationReplicationArgs', 'LocalRepositoryMultiReplicationReplicationArgsDict']]]]] = None,
            repo_key: Optional[pulumi.Input[str]] = None) -> 'LocalRepositoryMultiReplication':
        """
        Get an existing LocalRepositoryMultiReplication resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cron_exp: A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
        :param pulumi.Input[bool] enable_event_replication: When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is `false`.
        :param pulumi.Input[Sequence[pulumi.Input[Union['LocalRepositoryMultiReplicationReplicationArgs', 'LocalRepositoryMultiReplicationReplicationArgsDict']]]] replications: List of replications minimum 1 element.
        :param pulumi.Input[str] repo_key: Repository name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LocalRepositoryMultiReplicationState.__new__(_LocalRepositoryMultiReplicationState)

        __props__.__dict__["cron_exp"] = cron_exp
        __props__.__dict__["enable_event_replication"] = enable_event_replication
        __props__.__dict__["replications"] = replications
        __props__.__dict__["repo_key"] = repo_key
        return LocalRepositoryMultiReplication(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cronExp")
    def cron_exp(self) -> pulumi.Output[str]:
        """
        A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
        """
        return pulumi.get(self, "cron_exp")

    @property
    @pulumi.getter(name="enableEventReplication")
    def enable_event_replication(self) -> pulumi.Output[bool]:
        """
        When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is `false`.
        """
        return pulumi.get(self, "enable_event_replication")

    @property
    @pulumi.getter
    def replications(self) -> pulumi.Output[Optional[Sequence['outputs.LocalRepositoryMultiReplicationReplication']]]:
        """
        List of replications minimum 1 element.
        """
        return pulumi.get(self, "replications")

    @property
    @pulumi.getter(name="repoKey")
    def repo_key(self) -> pulumi.Output[str]:
        """
        Repository name.
        """
        return pulumi.get(self, "repo_key")

