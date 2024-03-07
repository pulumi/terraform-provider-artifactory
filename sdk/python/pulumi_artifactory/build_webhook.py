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

__all__ = ['BuildWebhookArgs', 'BuildWebhook']

@pulumi.input_type
class BuildWebhookArgs:
    def __init__(__self__, *,
                 criteria: pulumi.Input['BuildWebhookCriteriaArgs'],
                 event_types: pulumi.Input[Sequence[pulumi.Input[str]]],
                 handlers: pulumi.Input[Sequence[pulumi.Input['BuildWebhookHandlerArgs']]],
                 key: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a BuildWebhook resource.
        :param pulumi.Input['BuildWebhookCriteriaArgs'] criteria: Specifies where the webhook will be applied on which repositories.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_types: List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `uploaded`, `deleted`, `promoted`.
        :param pulumi.Input[Sequence[pulumi.Input['BuildWebhookHandlerArgs']]] handlers: At least one is required.
        :param pulumi.Input[str] key: The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        :param pulumi.Input[str] description: Webhook description. Max length 1000 characters.
        :param pulumi.Input[bool] enabled: Status of webhook. Default to `true`.
        """
        pulumi.set(__self__, "criteria", criteria)
        pulumi.set(__self__, "event_types", event_types)
        pulumi.set(__self__, "handlers", handlers)
        pulumi.set(__self__, "key", key)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def criteria(self) -> pulumi.Input['BuildWebhookCriteriaArgs']:
        """
        Specifies where the webhook will be applied on which repositories.
        """
        return pulumi.get(self, "criteria")

    @criteria.setter
    def criteria(self, value: pulumi.Input['BuildWebhookCriteriaArgs']):
        pulumi.set(self, "criteria", value)

    @property
    @pulumi.getter(name="eventTypes")
    def event_types(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `uploaded`, `deleted`, `promoted`.
        """
        return pulumi.get(self, "event_types")

    @event_types.setter
    def event_types(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "event_types", value)

    @property
    @pulumi.getter
    def handlers(self) -> pulumi.Input[Sequence[pulumi.Input['BuildWebhookHandlerArgs']]]:
        """
        At least one is required.
        """
        return pulumi.get(self, "handlers")

    @handlers.setter
    def handlers(self, value: pulumi.Input[Sequence[pulumi.Input['BuildWebhookHandlerArgs']]]):
        pulumi.set(self, "handlers", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Webhook description. Max length 1000 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Status of webhook. Default to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)


@pulumi.input_type
class _BuildWebhookState:
    def __init__(__self__, *,
                 criteria: Optional[pulumi.Input['BuildWebhookCriteriaArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 handlers: Optional[pulumi.Input[Sequence[pulumi.Input['BuildWebhookHandlerArgs']]]] = None,
                 key: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BuildWebhook resources.
        :param pulumi.Input['BuildWebhookCriteriaArgs'] criteria: Specifies where the webhook will be applied on which repositories.
        :param pulumi.Input[str] description: Webhook description. Max length 1000 characters.
        :param pulumi.Input[bool] enabled: Status of webhook. Default to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_types: List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `uploaded`, `deleted`, `promoted`.
        :param pulumi.Input[Sequence[pulumi.Input['BuildWebhookHandlerArgs']]] handlers: At least one is required.
        :param pulumi.Input[str] key: The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        """
        if criteria is not None:
            pulumi.set(__self__, "criteria", criteria)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if event_types is not None:
            pulumi.set(__self__, "event_types", event_types)
        if handlers is not None:
            pulumi.set(__self__, "handlers", handlers)
        if key is not None:
            pulumi.set(__self__, "key", key)

    @property
    @pulumi.getter
    def criteria(self) -> Optional[pulumi.Input['BuildWebhookCriteriaArgs']]:
        """
        Specifies where the webhook will be applied on which repositories.
        """
        return pulumi.get(self, "criteria")

    @criteria.setter
    def criteria(self, value: Optional[pulumi.Input['BuildWebhookCriteriaArgs']]):
        pulumi.set(self, "criteria", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Webhook description. Max length 1000 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Status of webhook. Default to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="eventTypes")
    def event_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `uploaded`, `deleted`, `promoted`.
        """
        return pulumi.get(self, "event_types")

    @event_types.setter
    def event_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "event_types", value)

    @property
    @pulumi.getter
    def handlers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BuildWebhookHandlerArgs']]]]:
        """
        At least one is required.
        """
        return pulumi.get(self, "handlers")

    @handlers.setter
    def handlers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BuildWebhookHandlerArgs']]]]):
        pulumi.set(self, "handlers", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)


class BuildWebhook(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 criteria: Optional[pulumi.Input[pulumi.InputType['BuildWebhookCriteriaArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 handlers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BuildWebhookHandlerArgs']]]]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an Artifactory webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.

        ## Example Usage

        .
        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        build_webhook = artifactory.BuildWebhook("build-webhook",
            criteria=artifactory.BuildWebhookCriteriaArgs(
                any_build=True,
                exclude_patterns=["bar/**"],
                include_patterns=["foo/**"],
                selected_builds=["build-id"],
            ),
            event_types=[
                "uploaded",
                "deleted",
                "promoted",
            ],
            handlers=[artifactory.BuildWebhookHandlerArgs(
                custom_http_headers={
                    "header-1": "value-1",
                    "header-2": "value-2",
                },
                proxy="proxy-key",
                secret="some-secret",
                url="http://tempurl.org/webhook",
            )],
            key="build-webhook")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['BuildWebhookCriteriaArgs']] criteria: Specifies where the webhook will be applied on which repositories.
        :param pulumi.Input[str] description: Webhook description. Max length 1000 characters.
        :param pulumi.Input[bool] enabled: Status of webhook. Default to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_types: List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `uploaded`, `deleted`, `promoted`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BuildWebhookHandlerArgs']]]] handlers: At least one is required.
        :param pulumi.Input[str] key: The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BuildWebhookArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Artifactory webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.

        ## Example Usage

        .
        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        build_webhook = artifactory.BuildWebhook("build-webhook",
            criteria=artifactory.BuildWebhookCriteriaArgs(
                any_build=True,
                exclude_patterns=["bar/**"],
                include_patterns=["foo/**"],
                selected_builds=["build-id"],
            ),
            event_types=[
                "uploaded",
                "deleted",
                "promoted",
            ],
            handlers=[artifactory.BuildWebhookHandlerArgs(
                custom_http_headers={
                    "header-1": "value-1",
                    "header-2": "value-2",
                },
                proxy="proxy-key",
                secret="some-secret",
                url="http://tempurl.org/webhook",
            )],
            key="build-webhook")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param BuildWebhookArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BuildWebhookArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 criteria: Optional[pulumi.Input[pulumi.InputType['BuildWebhookCriteriaArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 handlers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BuildWebhookHandlerArgs']]]]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BuildWebhookArgs.__new__(BuildWebhookArgs)

            if criteria is None and not opts.urn:
                raise TypeError("Missing required property 'criteria'")
            __props__.__dict__["criteria"] = criteria
            __props__.__dict__["description"] = description
            __props__.__dict__["enabled"] = enabled
            if event_types is None and not opts.urn:
                raise TypeError("Missing required property 'event_types'")
            __props__.__dict__["event_types"] = event_types
            if handlers is None and not opts.urn:
                raise TypeError("Missing required property 'handlers'")
            __props__.__dict__["handlers"] = handlers
            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
        super(BuildWebhook, __self__).__init__(
            'artifactory:index/buildWebhook:BuildWebhook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            criteria: Optional[pulumi.Input[pulumi.InputType['BuildWebhookCriteriaArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            handlers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BuildWebhookHandlerArgs']]]]] = None,
            key: Optional[pulumi.Input[str]] = None) -> 'BuildWebhook':
        """
        Get an existing BuildWebhook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['BuildWebhookCriteriaArgs']] criteria: Specifies where the webhook will be applied on which repositories.
        :param pulumi.Input[str] description: Webhook description. Max length 1000 characters.
        :param pulumi.Input[bool] enabled: Status of webhook. Default to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_types: List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `uploaded`, `deleted`, `promoted`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BuildWebhookHandlerArgs']]]] handlers: At least one is required.
        :param pulumi.Input[str] key: The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BuildWebhookState.__new__(_BuildWebhookState)

        __props__.__dict__["criteria"] = criteria
        __props__.__dict__["description"] = description
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["event_types"] = event_types
        __props__.__dict__["handlers"] = handlers
        __props__.__dict__["key"] = key
        return BuildWebhook(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def criteria(self) -> pulumi.Output['outputs.BuildWebhookCriteria']:
        """
        Specifies where the webhook will be applied on which repositories.
        """
        return pulumi.get(self, "criteria")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Webhook description. Max length 1000 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Status of webhook. Default to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="eventTypes")
    def event_types(self) -> pulumi.Output[Sequence[str]]:
        """
        List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `uploaded`, `deleted`, `promoted`.
        """
        return pulumi.get(self, "event_types")

    @property
    @pulumi.getter
    def handlers(self) -> pulumi.Output[Sequence['outputs.BuildWebhookHandler']]:
        """
        At least one is required.
        """
        return pulumi.get(self, "handlers")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        """
        return pulumi.get(self, "key")

