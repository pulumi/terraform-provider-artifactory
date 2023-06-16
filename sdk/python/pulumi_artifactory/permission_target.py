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

__all__ = ['PermissionTargetArgs', 'PermissionTarget']

@pulumi.input_type
class PermissionTargetArgs:
    def __init__(__self__, *,
                 builds: Optional[pulumi.Input[Sequence[pulumi.Input['PermissionTargetBuildArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 release_bundles: Optional[pulumi.Input[Sequence[pulumi.Input['PermissionTargetReleaseBundleArgs']]]] = None,
                 repos: Optional[pulumi.Input[Sequence[pulumi.Input['PermissionTargetRepoArgs']]]] = None):
        """
        The set of arguments for constructing a PermissionTarget resource.
        :param pulumi.Input[Sequence[pulumi.Input['PermissionTargetBuildArgs']]] builds: As for repo but for artifactory-build-info permissions.
        :param pulumi.Input[str] name: Name of permission.
        :param pulumi.Input[Sequence[pulumi.Input['PermissionTargetReleaseBundleArgs']]] release_bundles: As for repo for for release-bundles permissions.
        :param pulumi.Input[Sequence[pulumi.Input['PermissionTargetRepoArgs']]] repos: Repository permission configuration.
        """
        if builds is not None:
            pulumi.set(__self__, "builds", builds)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if release_bundles is not None:
            pulumi.set(__self__, "release_bundles", release_bundles)
        if repos is not None:
            pulumi.set(__self__, "repos", repos)

    @property
    @pulumi.getter
    def builds(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PermissionTargetBuildArgs']]]]:
        """
        As for repo but for artifactory-build-info permissions.
        """
        return pulumi.get(self, "builds")

    @builds.setter
    def builds(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PermissionTargetBuildArgs']]]]):
        pulumi.set(self, "builds", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of permission.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="releaseBundles")
    def release_bundles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PermissionTargetReleaseBundleArgs']]]]:
        """
        As for repo for for release-bundles permissions.
        """
        return pulumi.get(self, "release_bundles")

    @release_bundles.setter
    def release_bundles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PermissionTargetReleaseBundleArgs']]]]):
        pulumi.set(self, "release_bundles", value)

    @property
    @pulumi.getter
    def repos(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PermissionTargetRepoArgs']]]]:
        """
        Repository permission configuration.
        """
        return pulumi.get(self, "repos")

    @repos.setter
    def repos(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PermissionTargetRepoArgs']]]]):
        pulumi.set(self, "repos", value)


@pulumi.input_type
class _PermissionTargetState:
    def __init__(__self__, *,
                 builds: Optional[pulumi.Input[Sequence[pulumi.Input['PermissionTargetBuildArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 release_bundles: Optional[pulumi.Input[Sequence[pulumi.Input['PermissionTargetReleaseBundleArgs']]]] = None,
                 repos: Optional[pulumi.Input[Sequence[pulumi.Input['PermissionTargetRepoArgs']]]] = None):
        """
        Input properties used for looking up and filtering PermissionTarget resources.
        :param pulumi.Input[Sequence[pulumi.Input['PermissionTargetBuildArgs']]] builds: As for repo but for artifactory-build-info permissions.
        :param pulumi.Input[str] name: Name of permission.
        :param pulumi.Input[Sequence[pulumi.Input['PermissionTargetReleaseBundleArgs']]] release_bundles: As for repo for for release-bundles permissions.
        :param pulumi.Input[Sequence[pulumi.Input['PermissionTargetRepoArgs']]] repos: Repository permission configuration.
        """
        if builds is not None:
            pulumi.set(__self__, "builds", builds)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if release_bundles is not None:
            pulumi.set(__self__, "release_bundles", release_bundles)
        if repos is not None:
            pulumi.set(__self__, "repos", repos)

    @property
    @pulumi.getter
    def builds(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PermissionTargetBuildArgs']]]]:
        """
        As for repo but for artifactory-build-info permissions.
        """
        return pulumi.get(self, "builds")

    @builds.setter
    def builds(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PermissionTargetBuildArgs']]]]):
        pulumi.set(self, "builds", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of permission.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="releaseBundles")
    def release_bundles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PermissionTargetReleaseBundleArgs']]]]:
        """
        As for repo for for release-bundles permissions.
        """
        return pulumi.get(self, "release_bundles")

    @release_bundles.setter
    def release_bundles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PermissionTargetReleaseBundleArgs']]]]):
        pulumi.set(self, "release_bundles", value)

    @property
    @pulumi.getter
    def repos(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PermissionTargetRepoArgs']]]]:
        """
        Repository permission configuration.
        """
        return pulumi.get(self, "repos")

    @repos.setter
    def repos(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PermissionTargetRepoArgs']]]]):
        pulumi.set(self, "repos", value)


class PermissionTarget(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 builds: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PermissionTargetBuildArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 release_bundles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PermissionTargetReleaseBundleArgs']]]]] = None,
                 repos: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PermissionTargetRepoArgs']]]]] = None,
                 __props__=None):
        """
        Provides an Artifactory permission target resource. This can be used to create and manage Artifactory permission targets.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        # Create a new Artifactory permission target called testpermission
        test_perm = artifactory.PermissionTarget("test-perm",
            builds=[artifactory.PermissionTargetBuildArgs(
                actions=[artifactory.PermissionTargetBuildActionArgs(
                    users=[artifactory.PermissionTargetBuildActionUserArgs(
                        name="anonymous",
                        permissions=[
                            "read",
                            "write",
                        ],
                    )],
                )],
                includes_patterns=["**"],
                repositories=["artifactory-build-info"],
            )],
            release_bundles=[artifactory.PermissionTargetReleaseBundleArgs(
                actions=[artifactory.PermissionTargetReleaseBundleActionArgs(
                    users=[artifactory.PermissionTargetReleaseBundleActionUserArgs(
                        name="anonymous",
                        permissions=["read"],
                    )],
                )],
                includes_patterns=["**"],
                repositories=["release-bundles"],
            )],
            repos=[artifactory.PermissionTargetRepoArgs(
                actions=[artifactory.PermissionTargetRepoActionArgs(
                    groups=[artifactory.PermissionTargetRepoActionGroupArgs(
                        name="readers",
                        permissions=["read"],
                    )],
                    users=[artifactory.PermissionTargetRepoActionUserArgs(
                        name="anonymous",
                        permissions=[
                            "read",
                            "write",
                        ],
                    )],
                )],
                excludes_patterns=["bar/**"],
                includes_patterns=["foo/**"],
                repositories=["example-repo-local"],
            )])
        ```
        ## Permissions

        The provider supports the following `permission` enums:

        * `read`
        * `write`
        * `annotate`
        * `delete`
        * `manage`
        * `managedXrayMeta`
        * `distribute`

        The values can be mapped to the permissions from the official [documentation](https://www.jfrog.com/confluence/display/JFROG/Permissions):

        * `read` - matches `Read` permissions.
        * `write` - matches `  Deploy / Cache / Create ` permissions.
        * `annotate` - matches `Annotate` permissions.
        * `delete` - matches `Delete / Overwrite` permissions.
        * `manage` - matches `Manage` permissions.
        * `managedXrayMeta` - matches `Manage Xray Metadata` permissions.
        * `distribute` - matches `Distribute` permissions.

        ## Import

        Permission targets can be imported using their name, e.g.

        ```sh
         $ pulumi import artifactory:index/permissionTarget:PermissionTarget terraform-test-permission mypermission
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PermissionTargetBuildArgs']]]] builds: As for repo but for artifactory-build-info permissions.
        :param pulumi.Input[str] name: Name of permission.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PermissionTargetReleaseBundleArgs']]]] release_bundles: As for repo for for release-bundles permissions.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PermissionTargetRepoArgs']]]] repos: Repository permission configuration.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PermissionTargetArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Artifactory permission target resource. This can be used to create and manage Artifactory permission targets.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        # Create a new Artifactory permission target called testpermission
        test_perm = artifactory.PermissionTarget("test-perm",
            builds=[artifactory.PermissionTargetBuildArgs(
                actions=[artifactory.PermissionTargetBuildActionArgs(
                    users=[artifactory.PermissionTargetBuildActionUserArgs(
                        name="anonymous",
                        permissions=[
                            "read",
                            "write",
                        ],
                    )],
                )],
                includes_patterns=["**"],
                repositories=["artifactory-build-info"],
            )],
            release_bundles=[artifactory.PermissionTargetReleaseBundleArgs(
                actions=[artifactory.PermissionTargetReleaseBundleActionArgs(
                    users=[artifactory.PermissionTargetReleaseBundleActionUserArgs(
                        name="anonymous",
                        permissions=["read"],
                    )],
                )],
                includes_patterns=["**"],
                repositories=["release-bundles"],
            )],
            repos=[artifactory.PermissionTargetRepoArgs(
                actions=[artifactory.PermissionTargetRepoActionArgs(
                    groups=[artifactory.PermissionTargetRepoActionGroupArgs(
                        name="readers",
                        permissions=["read"],
                    )],
                    users=[artifactory.PermissionTargetRepoActionUserArgs(
                        name="anonymous",
                        permissions=[
                            "read",
                            "write",
                        ],
                    )],
                )],
                excludes_patterns=["bar/**"],
                includes_patterns=["foo/**"],
                repositories=["example-repo-local"],
            )])
        ```
        ## Permissions

        The provider supports the following `permission` enums:

        * `read`
        * `write`
        * `annotate`
        * `delete`
        * `manage`
        * `managedXrayMeta`
        * `distribute`

        The values can be mapped to the permissions from the official [documentation](https://www.jfrog.com/confluence/display/JFROG/Permissions):

        * `read` - matches `Read` permissions.
        * `write` - matches `  Deploy / Cache / Create ` permissions.
        * `annotate` - matches `Annotate` permissions.
        * `delete` - matches `Delete / Overwrite` permissions.
        * `manage` - matches `Manage` permissions.
        * `managedXrayMeta` - matches `Manage Xray Metadata` permissions.
        * `distribute` - matches `Distribute` permissions.

        ## Import

        Permission targets can be imported using their name, e.g.

        ```sh
         $ pulumi import artifactory:index/permissionTarget:PermissionTarget terraform-test-permission mypermission
        ```

        :param str resource_name: The name of the resource.
        :param PermissionTargetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PermissionTargetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 builds: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PermissionTargetBuildArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 release_bundles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PermissionTargetReleaseBundleArgs']]]]] = None,
                 repos: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PermissionTargetRepoArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PermissionTargetArgs.__new__(PermissionTargetArgs)

            __props__.__dict__["builds"] = builds
            __props__.__dict__["name"] = name
            __props__.__dict__["release_bundles"] = release_bundles
            __props__.__dict__["repos"] = repos
        super(PermissionTarget, __self__).__init__(
            'artifactory:index/permissionTarget:PermissionTarget',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            builds: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PermissionTargetBuildArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            release_bundles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PermissionTargetReleaseBundleArgs']]]]] = None,
            repos: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PermissionTargetRepoArgs']]]]] = None) -> 'PermissionTarget':
        """
        Get an existing PermissionTarget resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PermissionTargetBuildArgs']]]] builds: As for repo but for artifactory-build-info permissions.
        :param pulumi.Input[str] name: Name of permission.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PermissionTargetReleaseBundleArgs']]]] release_bundles: As for repo for for release-bundles permissions.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PermissionTargetRepoArgs']]]] repos: Repository permission configuration.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PermissionTargetState.__new__(_PermissionTargetState)

        __props__.__dict__["builds"] = builds
        __props__.__dict__["name"] = name
        __props__.__dict__["release_bundles"] = release_bundles
        __props__.__dict__["repos"] = repos
        return PermissionTarget(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def builds(self) -> pulumi.Output[Optional[Sequence['outputs.PermissionTargetBuild']]]:
        """
        As for repo but for artifactory-build-info permissions.
        """
        return pulumi.get(self, "builds")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of permission.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="releaseBundles")
    def release_bundles(self) -> pulumi.Output[Optional[Sequence['outputs.PermissionTargetReleaseBundle']]]:
        """
        As for repo for for release-bundles permissions.
        """
        return pulumi.get(self, "release_bundles")

    @property
    @pulumi.getter
    def repos(self) -> pulumi.Output[Optional[Sequence['outputs.PermissionTargetRepo']]]:
        """
        Repository permission configuration.
        """
        return pulumi.get(self, "repos")

