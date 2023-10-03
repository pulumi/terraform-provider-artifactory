# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['GeneralSecurityArgs', 'GeneralSecurity']

@pulumi.input_type
class GeneralSecurityArgs:
    def __init__(__self__, *,
                 enable_anonymous_access: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a GeneralSecurity resource.
        """
        GeneralSecurityArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            enable_anonymous_access=enable_anonymous_access,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             enable_anonymous_access: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if enable_anonymous_access is not None:
            _setter("enable_anonymous_access", enable_anonymous_access)

    @property
    @pulumi.getter(name="enableAnonymousAccess")
    def enable_anonymous_access(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enable_anonymous_access")

    @enable_anonymous_access.setter
    def enable_anonymous_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_anonymous_access", value)


@pulumi.input_type
class _GeneralSecurityState:
    def __init__(__self__, *,
                 enable_anonymous_access: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering GeneralSecurity resources.
        """
        _GeneralSecurityState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            enable_anonymous_access=enable_anonymous_access,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             enable_anonymous_access: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if enable_anonymous_access is not None:
            _setter("enable_anonymous_access", enable_anonymous_access)

    @property
    @pulumi.getter(name="enableAnonymousAccess")
    def enable_anonymous_access(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enable_anonymous_access")

    @enable_anonymous_access.setter
    def enable_anonymous_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_anonymous_access", value)


class GeneralSecurity(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enable_anonymous_access: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        This resource can be used to manage Artifactory's general security settings.

        Only a single `GeneralSecurity` resource is meant to be defined.

        ~>The `GeneralSecurity` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        # Configure Artifactory general security settings
        security = artifactory.GeneralSecurity("security", enable_anonymous_access=True)
        ```

        ## Import

        Current general security settings can be imported using `security` as the `ID`, e.g.

        ```sh
         $ pulumi import artifactory:index/generalSecurity:GeneralSecurity security security
        ```
         ~>The `artifactory_general_security` resource uses endpoints that are undocumented and may not work with SaaS environments, or may change without notice.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[GeneralSecurityArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource can be used to manage Artifactory's general security settings.

        Only a single `GeneralSecurity` resource is meant to be defined.

        ~>The `GeneralSecurity` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        # Configure Artifactory general security settings
        security = artifactory.GeneralSecurity("security", enable_anonymous_access=True)
        ```

        ## Import

        Current general security settings can be imported using `security` as the `ID`, e.g.

        ```sh
         $ pulumi import artifactory:index/generalSecurity:GeneralSecurity security security
        ```
         ~>The `artifactory_general_security` resource uses endpoints that are undocumented and may not work with SaaS environments, or may change without notice.

        :param str resource_name: The name of the resource.
        :param GeneralSecurityArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GeneralSecurityArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            GeneralSecurityArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enable_anonymous_access: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GeneralSecurityArgs.__new__(GeneralSecurityArgs)

            __props__.__dict__["enable_anonymous_access"] = enable_anonymous_access
        super(GeneralSecurity, __self__).__init__(
            'artifactory:index/generalSecurity:GeneralSecurity',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enable_anonymous_access: Optional[pulumi.Input[bool]] = None) -> 'GeneralSecurity':
        """
        Get an existing GeneralSecurity resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GeneralSecurityState.__new__(_GeneralSecurityState)

        __props__.__dict__["enable_anonymous_access"] = enable_anonymous_access
        return GeneralSecurity(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="enableAnonymousAccess")
    def enable_anonymous_access(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "enable_anonymous_access")

