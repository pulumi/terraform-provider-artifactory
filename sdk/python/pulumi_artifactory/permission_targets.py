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

__all__ = ['PermissionTargetsArgs', 'PermissionTargets']

@pulumi.input_type
class PermissionTargetsArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 build: Optional[pulumi.Input['PermissionTargetsBuildArgs']] = None,
                 release_bundle: Optional[pulumi.Input['PermissionTargetsReleaseBundleArgs']] = None,
                 repo: Optional[pulumi.Input['PermissionTargetsRepoArgs']] = None):
        """
        The set of arguments for constructing a PermissionTargets resource.
        :param pulumi.Input[str] name: Name of permission.
        :param pulumi.Input['PermissionTargetsBuildArgs'] build: As for repo but for artifactory-build-info permissions.
        :param pulumi.Input['PermissionTargetsReleaseBundleArgs'] release_bundle: As for repo for for release-bundles permissions.
        :param pulumi.Input['PermissionTargetsRepoArgs'] repo: Repository permission configuration.
        """
        pulumi.set(__self__, "name", name)
        if build is not None:
            pulumi.set(__self__, "build", build)
        if release_bundle is not None:
            pulumi.set(__self__, "release_bundle", release_bundle)
        if repo is not None:
            pulumi.set(__self__, "repo", repo)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name of permission.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def build(self) -> Optional[pulumi.Input['PermissionTargetsBuildArgs']]:
        """
        As for repo but for artifactory-build-info permissions.
        """
        return pulumi.get(self, "build")

    @build.setter
    def build(self, value: Optional[pulumi.Input['PermissionTargetsBuildArgs']]):
        pulumi.set(self, "build", value)

    @property
    @pulumi.getter(name="releaseBundle")
    def release_bundle(self) -> Optional[pulumi.Input['PermissionTargetsReleaseBundleArgs']]:
        """
        As for repo for for release-bundles permissions.
        """
        return pulumi.get(self, "release_bundle")

    @release_bundle.setter
    def release_bundle(self, value: Optional[pulumi.Input['PermissionTargetsReleaseBundleArgs']]):
        pulumi.set(self, "release_bundle", value)

    @property
    @pulumi.getter
    def repo(self) -> Optional[pulumi.Input['PermissionTargetsRepoArgs']]:
        """
        Repository permission configuration.
        """
        return pulumi.get(self, "repo")

    @repo.setter
    def repo(self, value: Optional[pulumi.Input['PermissionTargetsRepoArgs']]):
        pulumi.set(self, "repo", value)


@pulumi.input_type
class _PermissionTargetsState:
    def __init__(__self__, *,
                 build: Optional[pulumi.Input['PermissionTargetsBuildArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 release_bundle: Optional[pulumi.Input['PermissionTargetsReleaseBundleArgs']] = None,
                 repo: Optional[pulumi.Input['PermissionTargetsRepoArgs']] = None):
        """
        Input properties used for looking up and filtering PermissionTargets resources.
        :param pulumi.Input['PermissionTargetsBuildArgs'] build: As for repo but for artifactory-build-info permissions.
        :param pulumi.Input[str] name: Name of permission.
        :param pulumi.Input['PermissionTargetsReleaseBundleArgs'] release_bundle: As for repo for for release-bundles permissions.
        :param pulumi.Input['PermissionTargetsRepoArgs'] repo: Repository permission configuration.
        """
        if build is not None:
            pulumi.set(__self__, "build", build)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if release_bundle is not None:
            pulumi.set(__self__, "release_bundle", release_bundle)
        if repo is not None:
            pulumi.set(__self__, "repo", repo)

    @property
    @pulumi.getter
    def build(self) -> Optional[pulumi.Input['PermissionTargetsBuildArgs']]:
        """
        As for repo but for artifactory-build-info permissions.
        """
        return pulumi.get(self, "build")

    @build.setter
    def build(self, value: Optional[pulumi.Input['PermissionTargetsBuildArgs']]):
        pulumi.set(self, "build", value)

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
    @pulumi.getter(name="releaseBundle")
    def release_bundle(self) -> Optional[pulumi.Input['PermissionTargetsReleaseBundleArgs']]:
        """
        As for repo for for release-bundles permissions.
        """
        return pulumi.get(self, "release_bundle")

    @release_bundle.setter
    def release_bundle(self, value: Optional[pulumi.Input['PermissionTargetsReleaseBundleArgs']]):
        pulumi.set(self, "release_bundle", value)

    @property
    @pulumi.getter
    def repo(self) -> Optional[pulumi.Input['PermissionTargetsRepoArgs']]:
        """
        Repository permission configuration.
        """
        return pulumi.get(self, "repo")

    @repo.setter
    def repo(self, value: Optional[pulumi.Input['PermissionTargetsRepoArgs']]):
        pulumi.set(self, "repo", value)


class PermissionTargets(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 build: Optional[pulumi.Input[pulumi.InputType['PermissionTargetsBuildArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 release_bundle: Optional[pulumi.Input[pulumi.InputType['PermissionTargetsReleaseBundleArgs']]] = None,
                 repo: Optional[pulumi.Input[pulumi.InputType['PermissionTargetsRepoArgs']]] = None,
                 __props__=None):
        """
        Provides an Artifactory permission target resource. This can be used to create and manage Artifactory permission targets.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        # Create a new Artifactory permission target called testpermission
        test_perm = artifactory.PermissionTarget("test-perm",
            build=artifactory.PermissionTargetBuildArgs(
                actions=artifactory.PermissionTargetBuildActionsArgs(
                    users=[artifactory.PermissionTargetBuildActionsUserArgs(
                        name="anonymous",
                        permissions=[
                            "read",
                            "write",
                        ],
                    )],
                ),
                includes_patterns=["**"],
                repositories=["artifactory-build-info"],
            ),
            name="test-perm",
            release_bundle=artifactory.PermissionTargetReleaseBundleArgs(
                actions=artifactory.PermissionTargetReleaseBundleActionsArgs(
                    users=[artifactory.PermissionTargetReleaseBundleActionsUserArgs(
                        name="anonymous",
                        permissions=["read"],
                    )],
                ),
                includes_patterns=["**"],
                repositories=["release-bundles"],
            ),
            repo=artifactory.PermissionTargetRepoArgs(
                actions=artifactory.PermissionTargetRepoActionsArgs(
                    groups=[artifactory.PermissionTargetRepoActionsGroupArgs(
                        name="readers",
                        permissions=["read"],
                    )],
                    users=[artifactory.PermissionTargetRepoActionsUserArgs(
                        name="anonymous",
                        permissions=[
                            "read",
                            "write",
                        ],
                    )],
                ),
                excludes_patterns=["bar/**"],
                includes_patterns=["foo/**"],
                repositories=["example-repo-local"],
            ))
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
         $ pulumi import artifactory:index/permissionTargets:PermissionTargets terraform-test-permission mypermission
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PermissionTargetsBuildArgs']] build: As for repo but for artifactory-build-info permissions.
        :param pulumi.Input[str] name: Name of permission.
        :param pulumi.Input[pulumi.InputType['PermissionTargetsReleaseBundleArgs']] release_bundle: As for repo for for release-bundles permissions.
        :param pulumi.Input[pulumi.InputType['PermissionTargetsRepoArgs']] repo: Repository permission configuration.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PermissionTargetsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Artifactory permission target resource. This can be used to create and manage Artifactory permission targets.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        # Create a new Artifactory permission target called testpermission
        test_perm = artifactory.PermissionTarget("test-perm",
            build=artifactory.PermissionTargetBuildArgs(
                actions=artifactory.PermissionTargetBuildActionsArgs(
                    users=[artifactory.PermissionTargetBuildActionsUserArgs(
                        name="anonymous",
                        permissions=[
                            "read",
                            "write",
                        ],
                    )],
                ),
                includes_patterns=["**"],
                repositories=["artifactory-build-info"],
            ),
            name="test-perm",
            release_bundle=artifactory.PermissionTargetReleaseBundleArgs(
                actions=artifactory.PermissionTargetReleaseBundleActionsArgs(
                    users=[artifactory.PermissionTargetReleaseBundleActionsUserArgs(
                        name="anonymous",
                        permissions=["read"],
                    )],
                ),
                includes_patterns=["**"],
                repositories=["release-bundles"],
            ),
            repo=artifactory.PermissionTargetRepoArgs(
                actions=artifactory.PermissionTargetRepoActionsArgs(
                    groups=[artifactory.PermissionTargetRepoActionsGroupArgs(
                        name="readers",
                        permissions=["read"],
                    )],
                    users=[artifactory.PermissionTargetRepoActionsUserArgs(
                        name="anonymous",
                        permissions=[
                            "read",
                            "write",
                        ],
                    )],
                ),
                excludes_patterns=["bar/**"],
                includes_patterns=["foo/**"],
                repositories=["example-repo-local"],
            ))
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
         $ pulumi import artifactory:index/permissionTargets:PermissionTargets terraform-test-permission mypermission
        ```

        :param str resource_name: The name of the resource.
        :param PermissionTargetsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PermissionTargetsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 build: Optional[pulumi.Input[pulumi.InputType['PermissionTargetsBuildArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 release_bundle: Optional[pulumi.Input[pulumi.InputType['PermissionTargetsReleaseBundleArgs']]] = None,
                 repo: Optional[pulumi.Input[pulumi.InputType['PermissionTargetsRepoArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PermissionTargetsArgs.__new__(PermissionTargetsArgs)

            __props__.__dict__["build"] = build
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["release_bundle"] = release_bundle
            __props__.__dict__["repo"] = repo
        super(PermissionTargets, __self__).__init__(
            'artifactory:index/permissionTargets:PermissionTargets',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            build: Optional[pulumi.Input[pulumi.InputType['PermissionTargetsBuildArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            release_bundle: Optional[pulumi.Input[pulumi.InputType['PermissionTargetsReleaseBundleArgs']]] = None,
            repo: Optional[pulumi.Input[pulumi.InputType['PermissionTargetsRepoArgs']]] = None) -> 'PermissionTargets':
        """
        Get an existing PermissionTargets resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PermissionTargetsBuildArgs']] build: As for repo but for artifactory-build-info permissions.
        :param pulumi.Input[str] name: Name of permission.
        :param pulumi.Input[pulumi.InputType['PermissionTargetsReleaseBundleArgs']] release_bundle: As for repo for for release-bundles permissions.
        :param pulumi.Input[pulumi.InputType['PermissionTargetsRepoArgs']] repo: Repository permission configuration.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PermissionTargetsState.__new__(_PermissionTargetsState)

        __props__.__dict__["build"] = build
        __props__.__dict__["name"] = name
        __props__.__dict__["release_bundle"] = release_bundle
        __props__.__dict__["repo"] = repo
        return PermissionTargets(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def build(self) -> pulumi.Output[Optional['outputs.PermissionTargetsBuild']]:
        """
        As for repo but for artifactory-build-info permissions.
        """
        return pulumi.get(self, "build")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of permission.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="releaseBundle")
    def release_bundle(self) -> pulumi.Output[Optional['outputs.PermissionTargetsReleaseBundle']]:
        """
        As for repo for for release-bundles permissions.
        """
        return pulumi.get(self, "release_bundle")

    @property
    @pulumi.getter
    def repo(self) -> pulumi.Output[Optional['outputs.PermissionTargetsRepo']]:
        """
        Repository permission configuration.
        """
        return pulumi.get(self, "repo")

