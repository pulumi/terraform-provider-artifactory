# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ReleaseBundleV2PromotionArgs', 'ReleaseBundleV2Promotion']

@pulumi.input_type
class ReleaseBundleV2PromotionArgs:
    def __init__(__self__, *,
                 environment: pulumi.Input[str],
                 keypair_name: pulumi.Input[str],
                 version: pulumi.Input[str],
                 excluded_repository_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 included_repository_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_key: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ReleaseBundleV2Promotion resource.
        :param pulumi.Input[str] environment: Target environment
        :param pulumi.Input[str] keypair_name: Key-pair name to use for signature creation
        :param pulumi.Input[str] version: Version to promote
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_repository_keys: Defines specific repositories to exclude from the promotion.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] included_repository_keys: Defines specific repositories to include in the promotion. If this property is left undefined, all repositories (except those specifically excluded) are included in the promotion. Important: If one or more repositories are specifically included, all other repositories are excluded (regardless of what is defined in `excluded_repository_keys`).
        :param pulumi.Input[str] name: Name of Release Bundle
        :param pulumi.Input[str] project_key: Project key the Release Bundle belongs to
        """
        pulumi.set(__self__, "environment", environment)
        pulumi.set(__self__, "keypair_name", keypair_name)
        pulumi.set(__self__, "version", version)
        if excluded_repository_keys is not None:
            pulumi.set(__self__, "excluded_repository_keys", excluded_repository_keys)
        if included_repository_keys is not None:
            pulumi.set(__self__, "included_repository_keys", included_repository_keys)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_key is not None:
            pulumi.set(__self__, "project_key", project_key)

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Input[str]:
        """
        Target environment
        """
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: pulumi.Input[str]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter(name="keypairName")
    def keypair_name(self) -> pulumi.Input[str]:
        """
        Key-pair name to use for signature creation
        """
        return pulumi.get(self, "keypair_name")

    @keypair_name.setter
    def keypair_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "keypair_name", value)

    @property
    @pulumi.getter
    def version(self) -> pulumi.Input[str]:
        """
        Version to promote
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input[str]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter(name="excludedRepositoryKeys")
    def excluded_repository_keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Defines specific repositories to exclude from the promotion.
        """
        return pulumi.get(self, "excluded_repository_keys")

    @excluded_repository_keys.setter
    def excluded_repository_keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "excluded_repository_keys", value)

    @property
    @pulumi.getter(name="includedRepositoryKeys")
    def included_repository_keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Defines specific repositories to include in the promotion. If this property is left undefined, all repositories (except those specifically excluded) are included in the promotion. Important: If one or more repositories are specifically included, all other repositories are excluded (regardless of what is defined in `excluded_repository_keys`).
        """
        return pulumi.get(self, "included_repository_keys")

    @included_repository_keys.setter
    def included_repository_keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "included_repository_keys", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of Release Bundle
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> Optional[pulumi.Input[str]]:
        """
        Project key the Release Bundle belongs to
        """
        return pulumi.get(self, "project_key")

    @project_key.setter
    def project_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_key", value)


@pulumi.input_type
class _ReleaseBundleV2PromotionState:
    def __init__(__self__, *,
                 created: Optional[pulumi.Input[str]] = None,
                 created_millis: Optional[pulumi.Input[int]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 excluded_repository_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 included_repository_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 keypair_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_key: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ReleaseBundleV2Promotion resources.
        :param pulumi.Input[str] created: Timestamp when the new version was created (ISO 8601 standard).
        :param pulumi.Input[int] created_millis: Timestamp when the new version was created (in milliseconds).
        :param pulumi.Input[str] environment: Target environment
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_repository_keys: Defines specific repositories to exclude from the promotion.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] included_repository_keys: Defines specific repositories to include in the promotion. If this property is left undefined, all repositories (except those specifically excluded) are included in the promotion. Important: If one or more repositories are specifically included, all other repositories are excluded (regardless of what is defined in `excluded_repository_keys`).
        :param pulumi.Input[str] keypair_name: Key-pair name to use for signature creation
        :param pulumi.Input[str] name: Name of Release Bundle
        :param pulumi.Input[str] project_key: Project key the Release Bundle belongs to
        :param pulumi.Input[str] version: Version to promote
        """
        if created is not None:
            pulumi.set(__self__, "created", created)
        if created_millis is not None:
            pulumi.set(__self__, "created_millis", created_millis)
        if environment is not None:
            pulumi.set(__self__, "environment", environment)
        if excluded_repository_keys is not None:
            pulumi.set(__self__, "excluded_repository_keys", excluded_repository_keys)
        if included_repository_keys is not None:
            pulumi.set(__self__, "included_repository_keys", included_repository_keys)
        if keypair_name is not None:
            pulumi.set(__self__, "keypair_name", keypair_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_key is not None:
            pulumi.set(__self__, "project_key", project_key)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def created(self) -> Optional[pulumi.Input[str]]:
        """
        Timestamp when the new version was created (ISO 8601 standard).
        """
        return pulumi.get(self, "created")

    @created.setter
    def created(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created", value)

    @property
    @pulumi.getter(name="createdMillis")
    def created_millis(self) -> Optional[pulumi.Input[int]]:
        """
        Timestamp when the new version was created (in milliseconds).
        """
        return pulumi.get(self, "created_millis")

    @created_millis.setter
    def created_millis(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "created_millis", value)

    @property
    @pulumi.getter
    def environment(self) -> Optional[pulumi.Input[str]]:
        """
        Target environment
        """
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter(name="excludedRepositoryKeys")
    def excluded_repository_keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Defines specific repositories to exclude from the promotion.
        """
        return pulumi.get(self, "excluded_repository_keys")

    @excluded_repository_keys.setter
    def excluded_repository_keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "excluded_repository_keys", value)

    @property
    @pulumi.getter(name="includedRepositoryKeys")
    def included_repository_keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Defines specific repositories to include in the promotion. If this property is left undefined, all repositories (except those specifically excluded) are included in the promotion. Important: If one or more repositories are specifically included, all other repositories are excluded (regardless of what is defined in `excluded_repository_keys`).
        """
        return pulumi.get(self, "included_repository_keys")

    @included_repository_keys.setter
    def included_repository_keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "included_repository_keys", value)

    @property
    @pulumi.getter(name="keypairName")
    def keypair_name(self) -> Optional[pulumi.Input[str]]:
        """
        Key-pair name to use for signature creation
        """
        return pulumi.get(self, "keypair_name")

    @keypair_name.setter
    def keypair_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "keypair_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of Release Bundle
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> Optional[pulumi.Input[str]]:
        """
        Project key the Release Bundle belongs to
        """
        return pulumi.get(self, "project_key")

    @project_key.setter
    def project_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_key", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        Version to promote
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


class ReleaseBundleV2Promotion(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 excluded_repository_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 included_repository_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 keypair_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_key: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource enables you to promote Release Bundle V2 version. For more information, see [JFrog documentation](https://jfrog.com/help/r/jfrog-artifactory-documentation/promote-a-release-bundle-v2-to-a-target-environment).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        my_release_bundle_v2_promotion = artifactory.ReleaseBundleV2Promotion("my-release-bundle-v2-promotion",
            name="my-release-bundle-v2-artifacts",
            version="1.0.0",
            keypair_name="my-keypair-name",
            environment="DEV",
            included_repository_keys=["commons-qa-maven-local"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] environment: Target environment
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_repository_keys: Defines specific repositories to exclude from the promotion.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] included_repository_keys: Defines specific repositories to include in the promotion. If this property is left undefined, all repositories (except those specifically excluded) are included in the promotion. Important: If one or more repositories are specifically included, all other repositories are excluded (regardless of what is defined in `excluded_repository_keys`).
        :param pulumi.Input[str] keypair_name: Key-pair name to use for signature creation
        :param pulumi.Input[str] name: Name of Release Bundle
        :param pulumi.Input[str] project_key: Project key the Release Bundle belongs to
        :param pulumi.Input[str] version: Version to promote
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReleaseBundleV2PromotionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource enables you to promote Release Bundle V2 version. For more information, see [JFrog documentation](https://jfrog.com/help/r/jfrog-artifactory-documentation/promote-a-release-bundle-v2-to-a-target-environment).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        my_release_bundle_v2_promotion = artifactory.ReleaseBundleV2Promotion("my-release-bundle-v2-promotion",
            name="my-release-bundle-v2-artifacts",
            version="1.0.0",
            keypair_name="my-keypair-name",
            environment="DEV",
            included_repository_keys=["commons-qa-maven-local"])
        ```

        :param str resource_name: The name of the resource.
        :param ReleaseBundleV2PromotionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReleaseBundleV2PromotionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 excluded_repository_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 included_repository_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 keypair_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_key: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ReleaseBundleV2PromotionArgs.__new__(ReleaseBundleV2PromotionArgs)

            if environment is None and not opts.urn:
                raise TypeError("Missing required property 'environment'")
            __props__.__dict__["environment"] = environment
            __props__.__dict__["excluded_repository_keys"] = excluded_repository_keys
            __props__.__dict__["included_repository_keys"] = included_repository_keys
            if keypair_name is None and not opts.urn:
                raise TypeError("Missing required property 'keypair_name'")
            __props__.__dict__["keypair_name"] = keypair_name
            __props__.__dict__["name"] = name
            __props__.__dict__["project_key"] = project_key
            if version is None and not opts.urn:
                raise TypeError("Missing required property 'version'")
            __props__.__dict__["version"] = version
            __props__.__dict__["created"] = None
            __props__.__dict__["created_millis"] = None
        super(ReleaseBundleV2Promotion, __self__).__init__(
            'artifactory:index/releaseBundleV2Promotion:ReleaseBundleV2Promotion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created: Optional[pulumi.Input[str]] = None,
            created_millis: Optional[pulumi.Input[int]] = None,
            environment: Optional[pulumi.Input[str]] = None,
            excluded_repository_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            included_repository_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            keypair_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_key: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'ReleaseBundleV2Promotion':
        """
        Get an existing ReleaseBundleV2Promotion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created: Timestamp when the new version was created (ISO 8601 standard).
        :param pulumi.Input[int] created_millis: Timestamp when the new version was created (in milliseconds).
        :param pulumi.Input[str] environment: Target environment
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_repository_keys: Defines specific repositories to exclude from the promotion.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] included_repository_keys: Defines specific repositories to include in the promotion. If this property is left undefined, all repositories (except those specifically excluded) are included in the promotion. Important: If one or more repositories are specifically included, all other repositories are excluded (regardless of what is defined in `excluded_repository_keys`).
        :param pulumi.Input[str] keypair_name: Key-pair name to use for signature creation
        :param pulumi.Input[str] name: Name of Release Bundle
        :param pulumi.Input[str] project_key: Project key the Release Bundle belongs to
        :param pulumi.Input[str] version: Version to promote
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ReleaseBundleV2PromotionState.__new__(_ReleaseBundleV2PromotionState)

        __props__.__dict__["created"] = created
        __props__.__dict__["created_millis"] = created_millis
        __props__.__dict__["environment"] = environment
        __props__.__dict__["excluded_repository_keys"] = excluded_repository_keys
        __props__.__dict__["included_repository_keys"] = included_repository_keys
        __props__.__dict__["keypair_name"] = keypair_name
        __props__.__dict__["name"] = name
        __props__.__dict__["project_key"] = project_key
        __props__.__dict__["version"] = version
        return ReleaseBundleV2Promotion(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def created(self) -> pulumi.Output[str]:
        """
        Timestamp when the new version was created (ISO 8601 standard).
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter(name="createdMillis")
    def created_millis(self) -> pulumi.Output[int]:
        """
        Timestamp when the new version was created (in milliseconds).
        """
        return pulumi.get(self, "created_millis")

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Output[str]:
        """
        Target environment
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter(name="excludedRepositoryKeys")
    def excluded_repository_keys(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Defines specific repositories to exclude from the promotion.
        """
        return pulumi.get(self, "excluded_repository_keys")

    @property
    @pulumi.getter(name="includedRepositoryKeys")
    def included_repository_keys(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Defines specific repositories to include in the promotion. If this property is left undefined, all repositories (except those specifically excluded) are included in the promotion. Important: If one or more repositories are specifically included, all other repositories are excluded (regardless of what is defined in `excluded_repository_keys`).
        """
        return pulumi.get(self, "included_repository_keys")

    @property
    @pulumi.getter(name="keypairName")
    def keypair_name(self) -> pulumi.Output[str]:
        """
        Key-pair name to use for signature creation
        """
        return pulumi.get(self, "keypair_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of Release Bundle
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> pulumi.Output[Optional[str]]:
        """
        Project key the Release Bundle belongs to
        """
        return pulumi.get(self, "project_key")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        Version to promote
        """
        return pulumi.get(self, "version")

