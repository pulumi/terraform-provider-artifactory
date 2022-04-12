# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ArtifactWebhookArgs', 'ArtifactWebhook']

@pulumi.input_type
class ArtifactWebhookArgs:
    def __init__(__self__, *,
                 criteria: pulumi.Input['ArtifactWebhookCriteriaArgs'],
                 event_types: pulumi.Input[Sequence[pulumi.Input[str]]],
                 key: pulumi.Input[str],
                 url: pulumi.Input[str],
                 custom_http_headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 proxy: Optional[pulumi.Input[str]] = None,
                 secret: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ArtifactWebhook resource.
        :param pulumi.Input['ArtifactWebhookCriteriaArgs'] criteria: Specifies where the webhook will be applied on which repositories.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_types: List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow
               values: deployed, deleted, moved, copied
        :param pulumi.Input[str] key: Key of webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        :param pulumi.Input[str] url: Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_http_headers: Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
        :param pulumi.Input[str] description: Description of webhook. Max length 1000 characters.
        :param pulumi.Input[bool] enabled: Status of webhook. Default to 'true'
        :param pulumi.Input[str] proxy: Proxy key from Artifactory Proxies setting
        :param pulumi.Input[str] secret: Secret authentication token that will be sent to the configured URL.
        """
        pulumi.set(__self__, "criteria", criteria)
        pulumi.set(__self__, "event_types", event_types)
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "url", url)
        if custom_http_headers is not None:
            pulumi.set(__self__, "custom_http_headers", custom_http_headers)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if proxy is not None:
            pulumi.set(__self__, "proxy", proxy)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)

    @property
    @pulumi.getter
    def criteria(self) -> pulumi.Input['ArtifactWebhookCriteriaArgs']:
        """
        Specifies where the webhook will be applied on which repositories.
        """
        return pulumi.get(self, "criteria")

    @criteria.setter
    def criteria(self, value: pulumi.Input['ArtifactWebhookCriteriaArgs']):
        pulumi.set(self, "criteria", value)

    @property
    @pulumi.getter(name="eventTypes")
    def event_types(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow
        values: deployed, deleted, moved, copied
        """
        return pulumi.get(self, "event_types")

    @event_types.setter
    def event_types(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "event_types", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        Key of webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def url(self) -> pulumi.Input[str]:
        """
        Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter(name="customHttpHeaders")
    def custom_http_headers(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
        """
        return pulumi.get(self, "custom_http_headers")

    @custom_http_headers.setter
    def custom_http_headers(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "custom_http_headers", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of webhook. Max length 1000 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Status of webhook. Default to 'true'
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def proxy(self) -> Optional[pulumi.Input[str]]:
        """
        Proxy key from Artifactory Proxies setting
        """
        return pulumi.get(self, "proxy")

    @proxy.setter
    def proxy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "proxy", value)

    @property
    @pulumi.getter
    def secret(self) -> Optional[pulumi.Input[str]]:
        """
        Secret authentication token that will be sent to the configured URL.
        """
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret", value)


@pulumi.input_type
class _ArtifactWebhookState:
    def __init__(__self__, *,
                 criteria: Optional[pulumi.Input['ArtifactWebhookCriteriaArgs']] = None,
                 custom_http_headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 proxy: Optional[pulumi.Input[str]] = None,
                 secret: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ArtifactWebhook resources.
        :param pulumi.Input['ArtifactWebhookCriteriaArgs'] criteria: Specifies where the webhook will be applied on which repositories.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_http_headers: Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
        :param pulumi.Input[str] description: Description of webhook. Max length 1000 characters.
        :param pulumi.Input[bool] enabled: Status of webhook. Default to 'true'
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_types: List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow
               values: deployed, deleted, moved, copied
        :param pulumi.Input[str] key: Key of webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        :param pulumi.Input[str] proxy: Proxy key from Artifactory Proxies setting
        :param pulumi.Input[str] secret: Secret authentication token that will be sent to the configured URL.
        :param pulumi.Input[str] url: Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
        """
        if criteria is not None:
            pulumi.set(__self__, "criteria", criteria)
        if custom_http_headers is not None:
            pulumi.set(__self__, "custom_http_headers", custom_http_headers)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if event_types is not None:
            pulumi.set(__self__, "event_types", event_types)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if proxy is not None:
            pulumi.set(__self__, "proxy", proxy)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def criteria(self) -> Optional[pulumi.Input['ArtifactWebhookCriteriaArgs']]:
        """
        Specifies where the webhook will be applied on which repositories.
        """
        return pulumi.get(self, "criteria")

    @criteria.setter
    def criteria(self, value: Optional[pulumi.Input['ArtifactWebhookCriteriaArgs']]):
        pulumi.set(self, "criteria", value)

    @property
    @pulumi.getter(name="customHttpHeaders")
    def custom_http_headers(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
        """
        return pulumi.get(self, "custom_http_headers")

    @custom_http_headers.setter
    def custom_http_headers(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "custom_http_headers", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of webhook. Max length 1000 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Status of webhook. Default to 'true'
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="eventTypes")
    def event_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow
        values: deployed, deleted, moved, copied
        """
        return pulumi.get(self, "event_types")

    @event_types.setter
    def event_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "event_types", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        Key of webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def proxy(self) -> Optional[pulumi.Input[str]]:
        """
        Proxy key from Artifactory Proxies setting
        """
        return pulumi.get(self, "proxy")

    @proxy.setter
    def proxy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "proxy", value)

    @property
    @pulumi.getter
    def secret(self) -> Optional[pulumi.Input[str]]:
        """
        Secret authentication token that will be sent to the configured URL.
        """
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class ArtifactWebhook(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 criteria: Optional[pulumi.Input[pulumi.InputType['ArtifactWebhookCriteriaArgs']]] = None,
                 custom_http_headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 proxy: Optional[pulumi.Input[str]] = None,
                 secret: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ArtifactWebhook resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ArtifactWebhookCriteriaArgs']] criteria: Specifies where the webhook will be applied on which repositories.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_http_headers: Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
        :param pulumi.Input[str] description: Description of webhook. Max length 1000 characters.
        :param pulumi.Input[bool] enabled: Status of webhook. Default to 'true'
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_types: List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow
               values: deployed, deleted, moved, copied
        :param pulumi.Input[str] key: Key of webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        :param pulumi.Input[str] proxy: Proxy key from Artifactory Proxies setting
        :param pulumi.Input[str] secret: Secret authentication token that will be sent to the configured URL.
        :param pulumi.Input[str] url: Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ArtifactWebhookArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ArtifactWebhook resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ArtifactWebhookArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ArtifactWebhookArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 criteria: Optional[pulumi.Input[pulumi.InputType['ArtifactWebhookCriteriaArgs']]] = None,
                 custom_http_headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 proxy: Optional[pulumi.Input[str]] = None,
                 secret: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
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
            __props__ = ArtifactWebhookArgs.__new__(ArtifactWebhookArgs)

            if criteria is None and not opts.urn:
                raise TypeError("Missing required property 'criteria'")
            __props__.__dict__["criteria"] = criteria
            __props__.__dict__["custom_http_headers"] = custom_http_headers
            __props__.__dict__["description"] = description
            __props__.__dict__["enabled"] = enabled
            if event_types is None and not opts.urn:
                raise TypeError("Missing required property 'event_types'")
            __props__.__dict__["event_types"] = event_types
            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            __props__.__dict__["proxy"] = proxy
            __props__.__dict__["secret"] = secret
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__.__dict__["url"] = url
        super(ArtifactWebhook, __self__).__init__(
            'artifactory:index/artifactWebhook:ArtifactWebhook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            criteria: Optional[pulumi.Input[pulumi.InputType['ArtifactWebhookCriteriaArgs']]] = None,
            custom_http_headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            key: Optional[pulumi.Input[str]] = None,
            proxy: Optional[pulumi.Input[str]] = None,
            secret: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'ArtifactWebhook':
        """
        Get an existing ArtifactWebhook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ArtifactWebhookCriteriaArgs']] criteria: Specifies where the webhook will be applied on which repositories.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_http_headers: Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
        :param pulumi.Input[str] description: Description of webhook. Max length 1000 characters.
        :param pulumi.Input[bool] enabled: Status of webhook. Default to 'true'
        :param pulumi.Input[Sequence[pulumi.Input[str]]] event_types: List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow
               values: deployed, deleted, moved, copied
        :param pulumi.Input[str] key: Key of webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        :param pulumi.Input[str] proxy: Proxy key from Artifactory Proxies setting
        :param pulumi.Input[str] secret: Secret authentication token that will be sent to the configured URL.
        :param pulumi.Input[str] url: Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ArtifactWebhookState.__new__(_ArtifactWebhookState)

        __props__.__dict__["criteria"] = criteria
        __props__.__dict__["custom_http_headers"] = custom_http_headers
        __props__.__dict__["description"] = description
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["event_types"] = event_types
        __props__.__dict__["key"] = key
        __props__.__dict__["proxy"] = proxy
        __props__.__dict__["secret"] = secret
        __props__.__dict__["url"] = url
        return ArtifactWebhook(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def criteria(self) -> pulumi.Output['outputs.ArtifactWebhookCriteria']:
        """
        Specifies where the webhook will be applied on which repositories.
        """
        return pulumi.get(self, "criteria")

    @property
    @pulumi.getter(name="customHttpHeaders")
    def custom_http_headers(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
        """
        return pulumi.get(self, "custom_http_headers")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of webhook. Max length 1000 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Status of webhook. Default to 'true'
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="eventTypes")
    def event_types(self) -> pulumi.Output[Sequence[str]]:
        """
        List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow
        values: deployed, deleted, moved, copied
        """
        return pulumi.get(self, "event_types")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        Key of webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def proxy(self) -> pulumi.Output[Optional[str]]:
        """
        Proxy key from Artifactory Proxies setting
        """
        return pulumi.get(self, "proxy")

    @property
    @pulumi.getter
    def secret(self) -> pulumi.Output[Optional[str]]:
        """
        Secret authentication token that will be sent to the configured URL.
        """
        return pulumi.get(self, "secret")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
        """
        return pulumi.get(self, "url")

