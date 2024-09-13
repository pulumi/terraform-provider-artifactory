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

__all__ = ['PackageCleanupPolicyArgs', 'PackageCleanupPolicy']

@pulumi.input_type
class PackageCleanupPolicyArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 search_criteria: pulumi.Input['PackageCleanupPolicySearchCriteriaArgs'],
                 cron_expression: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 duration_in_minutes: Optional[pulumi.Input[int]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 project_key: Optional[pulumi.Input[str]] = None,
                 skip_trashcan: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a PackageCleanupPolicy resource.
        :param pulumi.Input[str] key: Policy key. It has to be unique. It should not be used for other policies and configuration entities like archive policies, key pairs, repo layouts, property sets, backups, proxies, reverse proxies etc. A minimum of three characters is required and can include letters, numbers, underscore and hyphen.
        :param pulumi.Input[str] cron_expression: The Cron expression that sets the schedule of policy execution. For example, `0 0 2 * * ?` executes the policy every day at 02:00 AM. The minimum recurrent time for policy execution is 6 hours.
        :param pulumi.Input[int] duration_in_minutes: Enable and select the maximum duration for policy execution. Note: using this setting can cause the policy to stop before completion.
        :param pulumi.Input[bool] enabled: Enables or disabled the package cleanup policy. This allows the user to run the policy manually. If a policy has a valid cron expression, then it will be scheduled for execution based on it. If a policy is disabled, its future executions will be unscheduled. Defaults to `true`
        :param pulumi.Input[str] project_key: This attribute is used only for project-level cleanup policies, it is not used for global-level policies.
        :param pulumi.Input[bool] skip_trashcan: Enabling this setting results in packages being permanently deleted from Artifactory after the cleanup policy is executed instead of going to the Trash Can repository. Defaults to `false`.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "search_criteria", search_criteria)
        if cron_expression is not None:
            pulumi.set(__self__, "cron_expression", cron_expression)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if duration_in_minutes is not None:
            pulumi.set(__self__, "duration_in_minutes", duration_in_minutes)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if project_key is not None:
            pulumi.set(__self__, "project_key", project_key)
        if skip_trashcan is not None:
            pulumi.set(__self__, "skip_trashcan", skip_trashcan)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        Policy key. It has to be unique. It should not be used for other policies and configuration entities like archive policies, key pairs, repo layouts, property sets, backups, proxies, reverse proxies etc. A minimum of three characters is required and can include letters, numbers, underscore and hyphen.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="searchCriteria")
    def search_criteria(self) -> pulumi.Input['PackageCleanupPolicySearchCriteriaArgs']:
        return pulumi.get(self, "search_criteria")

    @search_criteria.setter
    def search_criteria(self, value: pulumi.Input['PackageCleanupPolicySearchCriteriaArgs']):
        pulumi.set(self, "search_criteria", value)

    @property
    @pulumi.getter(name="cronExpression")
    def cron_expression(self) -> Optional[pulumi.Input[str]]:
        """
        The Cron expression that sets the schedule of policy execution. For example, `0 0 2 * * ?` executes the policy every day at 02:00 AM. The minimum recurrent time for policy execution is 6 hours.
        """
        return pulumi.get(self, "cron_expression")

    @cron_expression.setter
    def cron_expression(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cron_expression", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="durationInMinutes")
    def duration_in_minutes(self) -> Optional[pulumi.Input[int]]:
        """
        Enable and select the maximum duration for policy execution. Note: using this setting can cause the policy to stop before completion.
        """
        return pulumi.get(self, "duration_in_minutes")

    @duration_in_minutes.setter
    def duration_in_minutes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "duration_in_minutes", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables or disabled the package cleanup policy. This allows the user to run the policy manually. If a policy has a valid cron expression, then it will be scheduled for execution based on it. If a policy is disabled, its future executions will be unscheduled. Defaults to `true`
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> Optional[pulumi.Input[str]]:
        """
        This attribute is used only for project-level cleanup policies, it is not used for global-level policies.
        """
        return pulumi.get(self, "project_key")

    @project_key.setter
    def project_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_key", value)

    @property
    @pulumi.getter(name="skipTrashcan")
    def skip_trashcan(self) -> Optional[pulumi.Input[bool]]:
        """
        Enabling this setting results in packages being permanently deleted from Artifactory after the cleanup policy is executed instead of going to the Trash Can repository. Defaults to `false`.
        """
        return pulumi.get(self, "skip_trashcan")

    @skip_trashcan.setter
    def skip_trashcan(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "skip_trashcan", value)


@pulumi.input_type
class _PackageCleanupPolicyState:
    def __init__(__self__, *,
                 cron_expression: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 duration_in_minutes: Optional[pulumi.Input[int]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 project_key: Optional[pulumi.Input[str]] = None,
                 search_criteria: Optional[pulumi.Input['PackageCleanupPolicySearchCriteriaArgs']] = None,
                 skip_trashcan: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering PackageCleanupPolicy resources.
        :param pulumi.Input[str] cron_expression: The Cron expression that sets the schedule of policy execution. For example, `0 0 2 * * ?` executes the policy every day at 02:00 AM. The minimum recurrent time for policy execution is 6 hours.
        :param pulumi.Input[int] duration_in_minutes: Enable and select the maximum duration for policy execution. Note: using this setting can cause the policy to stop before completion.
        :param pulumi.Input[bool] enabled: Enables or disabled the package cleanup policy. This allows the user to run the policy manually. If a policy has a valid cron expression, then it will be scheduled for execution based on it. If a policy is disabled, its future executions will be unscheduled. Defaults to `true`
        :param pulumi.Input[str] key: Policy key. It has to be unique. It should not be used for other policies and configuration entities like archive policies, key pairs, repo layouts, property sets, backups, proxies, reverse proxies etc. A minimum of three characters is required and can include letters, numbers, underscore and hyphen.
        :param pulumi.Input[str] project_key: This attribute is used only for project-level cleanup policies, it is not used for global-level policies.
        :param pulumi.Input[bool] skip_trashcan: Enabling this setting results in packages being permanently deleted from Artifactory after the cleanup policy is executed instead of going to the Trash Can repository. Defaults to `false`.
        """
        if cron_expression is not None:
            pulumi.set(__self__, "cron_expression", cron_expression)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if duration_in_minutes is not None:
            pulumi.set(__self__, "duration_in_minutes", duration_in_minutes)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if project_key is not None:
            pulumi.set(__self__, "project_key", project_key)
        if search_criteria is not None:
            pulumi.set(__self__, "search_criteria", search_criteria)
        if skip_trashcan is not None:
            pulumi.set(__self__, "skip_trashcan", skip_trashcan)

    @property
    @pulumi.getter(name="cronExpression")
    def cron_expression(self) -> Optional[pulumi.Input[str]]:
        """
        The Cron expression that sets the schedule of policy execution. For example, `0 0 2 * * ?` executes the policy every day at 02:00 AM. The minimum recurrent time for policy execution is 6 hours.
        """
        return pulumi.get(self, "cron_expression")

    @cron_expression.setter
    def cron_expression(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cron_expression", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="durationInMinutes")
    def duration_in_minutes(self) -> Optional[pulumi.Input[int]]:
        """
        Enable and select the maximum duration for policy execution. Note: using this setting can cause the policy to stop before completion.
        """
        return pulumi.get(self, "duration_in_minutes")

    @duration_in_minutes.setter
    def duration_in_minutes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "duration_in_minutes", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables or disabled the package cleanup policy. This allows the user to run the policy manually. If a policy has a valid cron expression, then it will be scheduled for execution based on it. If a policy is disabled, its future executions will be unscheduled. Defaults to `true`
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        Policy key. It has to be unique. It should not be used for other policies and configuration entities like archive policies, key pairs, repo layouts, property sets, backups, proxies, reverse proxies etc. A minimum of three characters is required and can include letters, numbers, underscore and hyphen.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> Optional[pulumi.Input[str]]:
        """
        This attribute is used only for project-level cleanup policies, it is not used for global-level policies.
        """
        return pulumi.get(self, "project_key")

    @project_key.setter
    def project_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_key", value)

    @property
    @pulumi.getter(name="searchCriteria")
    def search_criteria(self) -> Optional[pulumi.Input['PackageCleanupPolicySearchCriteriaArgs']]:
        return pulumi.get(self, "search_criteria")

    @search_criteria.setter
    def search_criteria(self, value: Optional[pulumi.Input['PackageCleanupPolicySearchCriteriaArgs']]):
        pulumi.set(self, "search_criteria", value)

    @property
    @pulumi.getter(name="skipTrashcan")
    def skip_trashcan(self) -> Optional[pulumi.Input[bool]]:
        """
        Enabling this setting results in packages being permanently deleted from Artifactory after the cleanup policy is executed instead of going to the Trash Can repository. Defaults to `false`.
        """
        return pulumi.get(self, "skip_trashcan")

    @skip_trashcan.setter
    def skip_trashcan(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "skip_trashcan", value)


class PackageCleanupPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cron_expression: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 duration_in_minutes: Optional[pulumi.Input[int]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 project_key: Optional[pulumi.Input[str]] = None,
                 search_criteria: Optional[pulumi.Input[Union['PackageCleanupPolicySearchCriteriaArgs', 'PackageCleanupPolicySearchCriteriaArgsDict']]] = None,
                 skip_trashcan: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Provides an Artifactory Package Cleanup Policy resource. This resource enable system administrators to define and customize policies based on specific criteria for removing unused binaries from across their JFrog platform. See [Rentation Policies](https://jfrog.com/help/r/jfrog-platform-administration-documentation/retention-policies) for more details.

        ~>Currently in beta and not yet globally available. A full rollout is scheduled for early October 2024.

        ## Import

        ```sh
        $ pulumi import artifactory:index/packageCleanupPolicy:PackageCleanupPolicy my-cleanup-policy my-policy
        ```

        ```sh
        $ pulumi import artifactory:index/packageCleanupPolicy:PackageCleanupPolicy my-cleanup-policy my-policy:myproj
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cron_expression: The Cron expression that sets the schedule of policy execution. For example, `0 0 2 * * ?` executes the policy every day at 02:00 AM. The minimum recurrent time for policy execution is 6 hours.
        :param pulumi.Input[int] duration_in_minutes: Enable and select the maximum duration for policy execution. Note: using this setting can cause the policy to stop before completion.
        :param pulumi.Input[bool] enabled: Enables or disabled the package cleanup policy. This allows the user to run the policy manually. If a policy has a valid cron expression, then it will be scheduled for execution based on it. If a policy is disabled, its future executions will be unscheduled. Defaults to `true`
        :param pulumi.Input[str] key: Policy key. It has to be unique. It should not be used for other policies and configuration entities like archive policies, key pairs, repo layouts, property sets, backups, proxies, reverse proxies etc. A minimum of three characters is required and can include letters, numbers, underscore and hyphen.
        :param pulumi.Input[str] project_key: This attribute is used only for project-level cleanup policies, it is not used for global-level policies.
        :param pulumi.Input[bool] skip_trashcan: Enabling this setting results in packages being permanently deleted from Artifactory after the cleanup policy is executed instead of going to the Trash Can repository. Defaults to `false`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PackageCleanupPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Artifactory Package Cleanup Policy resource. This resource enable system administrators to define and customize policies based on specific criteria for removing unused binaries from across their JFrog platform. See [Rentation Policies](https://jfrog.com/help/r/jfrog-platform-administration-documentation/retention-policies) for more details.

        ~>Currently in beta and not yet globally available. A full rollout is scheduled for early October 2024.

        ## Import

        ```sh
        $ pulumi import artifactory:index/packageCleanupPolicy:PackageCleanupPolicy my-cleanup-policy my-policy
        ```

        ```sh
        $ pulumi import artifactory:index/packageCleanupPolicy:PackageCleanupPolicy my-cleanup-policy my-policy:myproj
        ```

        :param str resource_name: The name of the resource.
        :param PackageCleanupPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PackageCleanupPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cron_expression: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 duration_in_minutes: Optional[pulumi.Input[int]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 project_key: Optional[pulumi.Input[str]] = None,
                 search_criteria: Optional[pulumi.Input[Union['PackageCleanupPolicySearchCriteriaArgs', 'PackageCleanupPolicySearchCriteriaArgsDict']]] = None,
                 skip_trashcan: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PackageCleanupPolicyArgs.__new__(PackageCleanupPolicyArgs)

            __props__.__dict__["cron_expression"] = cron_expression
            __props__.__dict__["description"] = description
            __props__.__dict__["duration_in_minutes"] = duration_in_minutes
            __props__.__dict__["enabled"] = enabled
            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            __props__.__dict__["project_key"] = project_key
            if search_criteria is None and not opts.urn:
                raise TypeError("Missing required property 'search_criteria'")
            __props__.__dict__["search_criteria"] = search_criteria
            __props__.__dict__["skip_trashcan"] = skip_trashcan
        super(PackageCleanupPolicy, __self__).__init__(
            'artifactory:index/packageCleanupPolicy:PackageCleanupPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cron_expression: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            duration_in_minutes: Optional[pulumi.Input[int]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            key: Optional[pulumi.Input[str]] = None,
            project_key: Optional[pulumi.Input[str]] = None,
            search_criteria: Optional[pulumi.Input[Union['PackageCleanupPolicySearchCriteriaArgs', 'PackageCleanupPolicySearchCriteriaArgsDict']]] = None,
            skip_trashcan: Optional[pulumi.Input[bool]] = None) -> 'PackageCleanupPolicy':
        """
        Get an existing PackageCleanupPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cron_expression: The Cron expression that sets the schedule of policy execution. For example, `0 0 2 * * ?` executes the policy every day at 02:00 AM. The minimum recurrent time for policy execution is 6 hours.
        :param pulumi.Input[int] duration_in_minutes: Enable and select the maximum duration for policy execution. Note: using this setting can cause the policy to stop before completion.
        :param pulumi.Input[bool] enabled: Enables or disabled the package cleanup policy. This allows the user to run the policy manually. If a policy has a valid cron expression, then it will be scheduled for execution based on it. If a policy is disabled, its future executions will be unscheduled. Defaults to `true`
        :param pulumi.Input[str] key: Policy key. It has to be unique. It should not be used for other policies and configuration entities like archive policies, key pairs, repo layouts, property sets, backups, proxies, reverse proxies etc. A minimum of three characters is required and can include letters, numbers, underscore and hyphen.
        :param pulumi.Input[str] project_key: This attribute is used only for project-level cleanup policies, it is not used for global-level policies.
        :param pulumi.Input[bool] skip_trashcan: Enabling this setting results in packages being permanently deleted from Artifactory after the cleanup policy is executed instead of going to the Trash Can repository. Defaults to `false`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PackageCleanupPolicyState.__new__(_PackageCleanupPolicyState)

        __props__.__dict__["cron_expression"] = cron_expression
        __props__.__dict__["description"] = description
        __props__.__dict__["duration_in_minutes"] = duration_in_minutes
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["key"] = key
        __props__.__dict__["project_key"] = project_key
        __props__.__dict__["search_criteria"] = search_criteria
        __props__.__dict__["skip_trashcan"] = skip_trashcan
        return PackageCleanupPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cronExpression")
    def cron_expression(self) -> pulumi.Output[Optional[str]]:
        """
        The Cron expression that sets the schedule of policy execution. For example, `0 0 2 * * ?` executes the policy every day at 02:00 AM. The minimum recurrent time for policy execution is 6 hours.
        """
        return pulumi.get(self, "cron_expression")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="durationInMinutes")
    def duration_in_minutes(self) -> pulumi.Output[Optional[int]]:
        """
        Enable and select the maximum duration for policy execution. Note: using this setting can cause the policy to stop before completion.
        """
        return pulumi.get(self, "duration_in_minutes")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        """
        Enables or disabled the package cleanup policy. This allows the user to run the policy manually. If a policy has a valid cron expression, then it will be scheduled for execution based on it. If a policy is disabled, its future executions will be unscheduled. Defaults to `true`
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        Policy key. It has to be unique. It should not be used for other policies and configuration entities like archive policies, key pairs, repo layouts, property sets, backups, proxies, reverse proxies etc. A minimum of three characters is required and can include letters, numbers, underscore and hyphen.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> pulumi.Output[Optional[str]]:
        """
        This attribute is used only for project-level cleanup policies, it is not used for global-level policies.
        """
        return pulumi.get(self, "project_key")

    @property
    @pulumi.getter(name="searchCriteria")
    def search_criteria(self) -> pulumi.Output['outputs.PackageCleanupPolicySearchCriteria']:
        return pulumi.get(self, "search_criteria")

    @property
    @pulumi.getter(name="skipTrashcan")
    def skip_trashcan(self) -> pulumi.Output[bool]:
        """
        Enabling this setting results in packages being permanently deleted from Artifactory after the cleanup policy is executed instead of going to the Trash Can repository. Defaults to `false`.
        """
        return pulumi.get(self, "skip_trashcan")

