# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ReleaseBundleV2CustomWebhookArgs', 'ReleaseBundleV2CustomWebhook']

@pulumi.input_type
class ReleaseBundleV2CustomWebhookArgs:
    def __init__(__self__, *,
                 event_types: pulumi.Input[Sequence[pulumi.Input[str]]],
                 key: pulumi.Input[str],
                 criteria: Optional[pulumi.Input['ReleaseBundleV2CustomWebhookCriteriaArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 handlers: Optional[pulumi.Input[Sequence[pulumi.Input['ReleaseBundleV2CustomWebhookHandlerArgs']]]] = None):
        """
        The set of arguments for constructing a ReleaseBundleV2CustomWebhook resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_types: List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `release_bundle_v2_started`, `release_bundle_v2_failed`, `release_bundle_v2_completed`.
        :param pulumi.Input[str] key: The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        :param pulumi.Input['ReleaseBundleV2CustomWebhookCriteriaArgs'] criteria: Specifies where the webhook will be applied on which repositories.
        :param pulumi.Input[str] description: Webhook description. Max length 1000 characters.
        :param pulumi.Input[bool] enabled: Status of webhook. Default to `true`.
        :param pulumi.Input[Sequence[pulumi.Input['ReleaseBundleV2CustomWebhookHandlerArgs']]] handlers: At least one is required.
        """
        pulumi.set(__self__, "event_types", event_types)
        pulumi.set(__self__, "key", key)
        if criteria is not None:
            pulumi.set(__self__, "criteria", criteria)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if handlers is not None:
            pulumi.set(__self__, "handlers", handlers)

    @property
    @pulumi.getter(name="eventTypes")
    def event_types(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `release_bundle_v2_started`, `release_bundle_v2_failed`, `release_bundle_v2_completed`.
        """
        return pulumi.get(self, "event_types")

    @event_types.setter
    def event_types(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "event_types", value)

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
    def criteria(self) -> Optional[pulumi.Input['ReleaseBundleV2CustomWebhookCriteriaArgs']]:
        """
        Specifies where the webhook will be applied on which repositories.
        """
        return pulumi.get(self, "criteria")

    @criteria.setter
    def criteria(self, value: Optional[pulumi.Input['ReleaseBundleV2CustomWebhookCriteriaArgs']]):
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
    @pulumi.getter
    def handlers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ReleaseBundleV2CustomWebhookHandlerArgs']]]]:
        """
        At least one is required.
        """
        return pulumi.get(self, "handlers")

    @handlers.setter
    def handlers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ReleaseBundleV2CustomWebhookHandlerArgs']]]]):
        pulumi.set(self, "handlers", value)


@pulumi.input_type
class _ReleaseBundleV2CustomWebhookState:
    def __init__(__self__, *,
                 criteria: Optional[pulumi.Input['ReleaseBundleV2CustomWebhookCriteriaArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 handlers: Optional[pulumi.Input[Sequence[pulumi.Input['ReleaseBundleV2CustomWebhookHandlerArgs']]]] = None,
                 key: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ReleaseBundleV2CustomWebhook resources.
        :param pulumi.Input['ReleaseBundleV2CustomWebhookCriteriaArgs'] criteria: Specifies where the webhook will be applied on which repositories.
        :param pulumi.Input[str] description: Webhook description. Max length 1000 characters.
        :param pulumi.Input[bool] enabled: Status of webhook. Default to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_types: List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `release_bundle_v2_started`, `release_bundle_v2_failed`, `release_bundle_v2_completed`.
        :param pulumi.Input[Sequence[pulumi.Input['ReleaseBundleV2CustomWebhookHandlerArgs']]] handlers: At least one is required.
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
    def criteria(self) -> Optional[pulumi.Input['ReleaseBundleV2CustomWebhookCriteriaArgs']]:
        """
        Specifies where the webhook will be applied on which repositories.
        """
        return pulumi.get(self, "criteria")

    @criteria.setter
    def criteria(self, value: Optional[pulumi.Input['ReleaseBundleV2CustomWebhookCriteriaArgs']]):
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
        List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `release_bundle_v2_started`, `release_bundle_v2_failed`, `release_bundle_v2_completed`.
        """
        return pulumi.get(self, "event_types")

    @event_types.setter
    def event_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "event_types", value)

    @property
    @pulumi.getter
    def handlers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ReleaseBundleV2CustomWebhookHandlerArgs']]]]:
        """
        At least one is required.
        """
        return pulumi.get(self, "handlers")

    @handlers.setter
    def handlers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ReleaseBundleV2CustomWebhookHandlerArgs']]]]):
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


class ReleaseBundleV2CustomWebhook(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 criteria: Optional[pulumi.Input[Union['ReleaseBundleV2CustomWebhookCriteriaArgs', 'ReleaseBundleV2CustomWebhookCriteriaArgsDict']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 handlers: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ReleaseBundleV2CustomWebhookHandlerArgs', 'ReleaseBundleV2CustomWebhookHandlerArgsDict']]]]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an Artifactory custom webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        release_bundle_v2_custom_webhook = artifactory.ReleaseBundleV2CustomWebhook("release-bundle-v2-custom-webhook",
            key="release-bundle-v2-custom-webhook",
            event_types=[
                "release_bundle_v2_started",
                "release_bundle_v2_failed",
                "release_bundle_v2_completed",
            ],
            criteria={
                "any_release_bundle": False,
                "selected_release_bundles": ["bundle-name"],
                "include_patterns": ["foo/**"],
                "exclude_patterns": ["bar/**"],
            },
            handlers=[{
                "url": "https://tempurl.org",
                "secrets": {
                    "secretName1": "value1",
                    "secretName2": "value2",
                },
                "http_headers": {
                    "headerName1": "value1",
                    "headerName2": "value2",
                },
                "payload": "{ \\"ref\\": \\"main\\" , \\"inputs\\": { \\"artifact_path\\": \\"test-repo/repo-path\\" } }",
            }])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['ReleaseBundleV2CustomWebhookCriteriaArgs', 'ReleaseBundleV2CustomWebhookCriteriaArgsDict']] criteria: Specifies where the webhook will be applied on which repositories.
        :param pulumi.Input[str] description: Webhook description. Max length 1000 characters.
        :param pulumi.Input[bool] enabled: Status of webhook. Default to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_types: List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `release_bundle_v2_started`, `release_bundle_v2_failed`, `release_bundle_v2_completed`.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ReleaseBundleV2CustomWebhookHandlerArgs', 'ReleaseBundleV2CustomWebhookHandlerArgsDict']]]] handlers: At least one is required.
        :param pulumi.Input[str] key: The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReleaseBundleV2CustomWebhookArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Artifactory custom webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        release_bundle_v2_custom_webhook = artifactory.ReleaseBundleV2CustomWebhook("release-bundle-v2-custom-webhook",
            key="release-bundle-v2-custom-webhook",
            event_types=[
                "release_bundle_v2_started",
                "release_bundle_v2_failed",
                "release_bundle_v2_completed",
            ],
            criteria={
                "any_release_bundle": False,
                "selected_release_bundles": ["bundle-name"],
                "include_patterns": ["foo/**"],
                "exclude_patterns": ["bar/**"],
            },
            handlers=[{
                "url": "https://tempurl.org",
                "secrets": {
                    "secretName1": "value1",
                    "secretName2": "value2",
                },
                "http_headers": {
                    "headerName1": "value1",
                    "headerName2": "value2",
                },
                "payload": "{ \\"ref\\": \\"main\\" , \\"inputs\\": { \\"artifact_path\\": \\"test-repo/repo-path\\" } }",
            }])
        ```

        :param str resource_name: The name of the resource.
        :param ReleaseBundleV2CustomWebhookArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReleaseBundleV2CustomWebhookArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 criteria: Optional[pulumi.Input[Union['ReleaseBundleV2CustomWebhookCriteriaArgs', 'ReleaseBundleV2CustomWebhookCriteriaArgsDict']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 handlers: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ReleaseBundleV2CustomWebhookHandlerArgs', 'ReleaseBundleV2CustomWebhookHandlerArgsDict']]]]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ReleaseBundleV2CustomWebhookArgs.__new__(ReleaseBundleV2CustomWebhookArgs)

            __props__.__dict__["criteria"] = criteria
            __props__.__dict__["description"] = description
            __props__.__dict__["enabled"] = enabled
            if event_types is None and not opts.urn:
                raise TypeError("Missing required property 'event_types'")
            __props__.__dict__["event_types"] = event_types
            __props__.__dict__["handlers"] = handlers
            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
        super(ReleaseBundleV2CustomWebhook, __self__).__init__(
            'artifactory:index/releaseBundleV2CustomWebhook:ReleaseBundleV2CustomWebhook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            criteria: Optional[pulumi.Input[Union['ReleaseBundleV2CustomWebhookCriteriaArgs', 'ReleaseBundleV2CustomWebhookCriteriaArgsDict']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            handlers: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ReleaseBundleV2CustomWebhookHandlerArgs', 'ReleaseBundleV2CustomWebhookHandlerArgsDict']]]]] = None,
            key: Optional[pulumi.Input[str]] = None) -> 'ReleaseBundleV2CustomWebhook':
        """
        Get an existing ReleaseBundleV2CustomWebhook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['ReleaseBundleV2CustomWebhookCriteriaArgs', 'ReleaseBundleV2CustomWebhookCriteriaArgsDict']] criteria: Specifies where the webhook will be applied on which repositories.
        :param pulumi.Input[str] description: Webhook description. Max length 1000 characters.
        :param pulumi.Input[bool] enabled: Status of webhook. Default to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_types: List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `release_bundle_v2_started`, `release_bundle_v2_failed`, `release_bundle_v2_completed`.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ReleaseBundleV2CustomWebhookHandlerArgs', 'ReleaseBundleV2CustomWebhookHandlerArgsDict']]]] handlers: At least one is required.
        :param pulumi.Input[str] key: The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ReleaseBundleV2CustomWebhookState.__new__(_ReleaseBundleV2CustomWebhookState)

        __props__.__dict__["criteria"] = criteria
        __props__.__dict__["description"] = description
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["event_types"] = event_types
        __props__.__dict__["handlers"] = handlers
        __props__.__dict__["key"] = key
        return ReleaseBundleV2CustomWebhook(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def criteria(self) -> pulumi.Output[Optional['outputs.ReleaseBundleV2CustomWebhookCriteria']]:
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
    def enabled(self) -> pulumi.Output[bool]:
        """
        Status of webhook. Default to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="eventTypes")
    def event_types(self) -> pulumi.Output[Sequence[str]]:
        """
        List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `release_bundle_v2_started`, `release_bundle_v2_failed`, `release_bundle_v2_completed`.
        """
        return pulumi.get(self, "event_types")

    @property
    @pulumi.getter
    def handlers(self) -> pulumi.Output[Optional[Sequence['outputs.ReleaseBundleV2CustomWebhookHandler']]]:
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

