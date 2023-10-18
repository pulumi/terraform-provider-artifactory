# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetFederatedRpmRepositoryResult',
    'AwaitableGetFederatedRpmRepositoryResult',
    'get_federated_rpm_repository',
    'get_federated_rpm_repository_output',
]

@pulumi.output_type
class GetFederatedRpmRepositoryResult:
    """
    A collection of values returned by getFederatedRpmRepository.
    """
    def __init__(__self__, archive_browsing_enabled=None, blacked_out=None, calculate_yum_metadata=None, cdn_redirect=None, cleanup_on_delete=None, description=None, download_direct=None, enable_file_lists_indexing=None, excludes_pattern=None, id=None, includes_pattern=None, key=None, members=None, notes=None, package_type=None, primary_keypair_ref=None, priority_resolution=None, project_environments=None, project_key=None, property_sets=None, repo_layout_ref=None, secondary_keypair_ref=None, xray_index=None, yum_group_file_names=None, yum_root_depth=None):
        if archive_browsing_enabled and not isinstance(archive_browsing_enabled, bool):
            raise TypeError("Expected argument 'archive_browsing_enabled' to be a bool")
        pulumi.set(__self__, "archive_browsing_enabled", archive_browsing_enabled)
        if blacked_out and not isinstance(blacked_out, bool):
            raise TypeError("Expected argument 'blacked_out' to be a bool")
        pulumi.set(__self__, "blacked_out", blacked_out)
        if calculate_yum_metadata and not isinstance(calculate_yum_metadata, bool):
            raise TypeError("Expected argument 'calculate_yum_metadata' to be a bool")
        pulumi.set(__self__, "calculate_yum_metadata", calculate_yum_metadata)
        if cdn_redirect and not isinstance(cdn_redirect, bool):
            raise TypeError("Expected argument 'cdn_redirect' to be a bool")
        pulumi.set(__self__, "cdn_redirect", cdn_redirect)
        if cleanup_on_delete and not isinstance(cleanup_on_delete, bool):
            raise TypeError("Expected argument 'cleanup_on_delete' to be a bool")
        pulumi.set(__self__, "cleanup_on_delete", cleanup_on_delete)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if download_direct and not isinstance(download_direct, bool):
            raise TypeError("Expected argument 'download_direct' to be a bool")
        pulumi.set(__self__, "download_direct", download_direct)
        if enable_file_lists_indexing and not isinstance(enable_file_lists_indexing, bool):
            raise TypeError("Expected argument 'enable_file_lists_indexing' to be a bool")
        pulumi.set(__self__, "enable_file_lists_indexing", enable_file_lists_indexing)
        if excludes_pattern and not isinstance(excludes_pattern, str):
            raise TypeError("Expected argument 'excludes_pattern' to be a str")
        pulumi.set(__self__, "excludes_pattern", excludes_pattern)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if includes_pattern and not isinstance(includes_pattern, str):
            raise TypeError("Expected argument 'includes_pattern' to be a str")
        pulumi.set(__self__, "includes_pattern", includes_pattern)
        if key and not isinstance(key, str):
            raise TypeError("Expected argument 'key' to be a str")
        pulumi.set(__self__, "key", key)
        if members and not isinstance(members, list):
            raise TypeError("Expected argument 'members' to be a list")
        pulumi.set(__self__, "members", members)
        if notes and not isinstance(notes, str):
            raise TypeError("Expected argument 'notes' to be a str")
        pulumi.set(__self__, "notes", notes)
        if package_type and not isinstance(package_type, str):
            raise TypeError("Expected argument 'package_type' to be a str")
        pulumi.set(__self__, "package_type", package_type)
        if primary_keypair_ref and not isinstance(primary_keypair_ref, str):
            raise TypeError("Expected argument 'primary_keypair_ref' to be a str")
        pulumi.set(__self__, "primary_keypair_ref", primary_keypair_ref)
        if priority_resolution and not isinstance(priority_resolution, bool):
            raise TypeError("Expected argument 'priority_resolution' to be a bool")
        pulumi.set(__self__, "priority_resolution", priority_resolution)
        if project_environments and not isinstance(project_environments, list):
            raise TypeError("Expected argument 'project_environments' to be a list")
        pulumi.set(__self__, "project_environments", project_environments)
        if project_key and not isinstance(project_key, str):
            raise TypeError("Expected argument 'project_key' to be a str")
        pulumi.set(__self__, "project_key", project_key)
        if property_sets and not isinstance(property_sets, list):
            raise TypeError("Expected argument 'property_sets' to be a list")
        pulumi.set(__self__, "property_sets", property_sets)
        if repo_layout_ref and not isinstance(repo_layout_ref, str):
            raise TypeError("Expected argument 'repo_layout_ref' to be a str")
        pulumi.set(__self__, "repo_layout_ref", repo_layout_ref)
        if secondary_keypair_ref and not isinstance(secondary_keypair_ref, str):
            raise TypeError("Expected argument 'secondary_keypair_ref' to be a str")
        pulumi.set(__self__, "secondary_keypair_ref", secondary_keypair_ref)
        if xray_index and not isinstance(xray_index, bool):
            raise TypeError("Expected argument 'xray_index' to be a bool")
        pulumi.set(__self__, "xray_index", xray_index)
        if yum_group_file_names and not isinstance(yum_group_file_names, str):
            raise TypeError("Expected argument 'yum_group_file_names' to be a str")
        pulumi.set(__self__, "yum_group_file_names", yum_group_file_names)
        if yum_root_depth and not isinstance(yum_root_depth, int):
            raise TypeError("Expected argument 'yum_root_depth' to be a int")
        pulumi.set(__self__, "yum_root_depth", yum_root_depth)

    @property
    @pulumi.getter(name="archiveBrowsingEnabled")
    def archive_browsing_enabled(self) -> Optional[bool]:
        return pulumi.get(self, "archive_browsing_enabled")

    @property
    @pulumi.getter(name="blackedOut")
    def blacked_out(self) -> Optional[bool]:
        return pulumi.get(self, "blacked_out")

    @property
    @pulumi.getter(name="calculateYumMetadata")
    def calculate_yum_metadata(self) -> Optional[bool]:
        return pulumi.get(self, "calculate_yum_metadata")

    @property
    @pulumi.getter(name="cdnRedirect")
    def cdn_redirect(self) -> Optional[bool]:
        return pulumi.get(self, "cdn_redirect")

    @property
    @pulumi.getter(name="cleanupOnDelete")
    def cleanup_on_delete(self) -> Optional[bool]:
        return pulumi.get(self, "cleanup_on_delete")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="downloadDirect")
    def download_direct(self) -> Optional[bool]:
        return pulumi.get(self, "download_direct")

    @property
    @pulumi.getter(name="enableFileListsIndexing")
    def enable_file_lists_indexing(self) -> Optional[bool]:
        return pulumi.get(self, "enable_file_lists_indexing")

    @property
    @pulumi.getter(name="excludesPattern")
    def excludes_pattern(self) -> Optional[str]:
        return pulumi.get(self, "excludes_pattern")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="includesPattern")
    def includes_pattern(self) -> Optional[str]:
        return pulumi.get(self, "includes_pattern")

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def members(self) -> Optional[Sequence['outputs.GetFederatedRpmRepositoryMemberResult']]:
        """
        The list of Federated members and must contain this repository URL (configured base URL
        `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
        Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
        to set up Federated repositories correctly.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def notes(self) -> Optional[str]:
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter(name="packageType")
    def package_type(self) -> str:
        return pulumi.get(self, "package_type")

    @property
    @pulumi.getter(name="primaryKeypairRef")
    def primary_keypair_ref(self) -> Optional[str]:
        return pulumi.get(self, "primary_keypair_ref")

    @property
    @pulumi.getter(name="priorityResolution")
    def priority_resolution(self) -> Optional[bool]:
        return pulumi.get(self, "priority_resolution")

    @property
    @pulumi.getter(name="projectEnvironments")
    def project_environments(self) -> Sequence[str]:
        return pulumi.get(self, "project_environments")

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> Optional[str]:
        return pulumi.get(self, "project_key")

    @property
    @pulumi.getter(name="propertySets")
    def property_sets(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "property_sets")

    @property
    @pulumi.getter(name="repoLayoutRef")
    def repo_layout_ref(self) -> Optional[str]:
        return pulumi.get(self, "repo_layout_ref")

    @property
    @pulumi.getter(name="secondaryKeypairRef")
    def secondary_keypair_ref(self) -> Optional[str]:
        return pulumi.get(self, "secondary_keypair_ref")

    @property
    @pulumi.getter(name="xrayIndex")
    def xray_index(self) -> Optional[bool]:
        return pulumi.get(self, "xray_index")

    @property
    @pulumi.getter(name="yumGroupFileNames")
    def yum_group_file_names(self) -> Optional[str]:
        return pulumi.get(self, "yum_group_file_names")

    @property
    @pulumi.getter(name="yumRootDepth")
    def yum_root_depth(self) -> Optional[int]:
        return pulumi.get(self, "yum_root_depth")


class AwaitableGetFederatedRpmRepositoryResult(GetFederatedRpmRepositoryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFederatedRpmRepositoryResult(
            archive_browsing_enabled=self.archive_browsing_enabled,
            blacked_out=self.blacked_out,
            calculate_yum_metadata=self.calculate_yum_metadata,
            cdn_redirect=self.cdn_redirect,
            cleanup_on_delete=self.cleanup_on_delete,
            description=self.description,
            download_direct=self.download_direct,
            enable_file_lists_indexing=self.enable_file_lists_indexing,
            excludes_pattern=self.excludes_pattern,
            id=self.id,
            includes_pattern=self.includes_pattern,
            key=self.key,
            members=self.members,
            notes=self.notes,
            package_type=self.package_type,
            primary_keypair_ref=self.primary_keypair_ref,
            priority_resolution=self.priority_resolution,
            project_environments=self.project_environments,
            project_key=self.project_key,
            property_sets=self.property_sets,
            repo_layout_ref=self.repo_layout_ref,
            secondary_keypair_ref=self.secondary_keypair_ref,
            xray_index=self.xray_index,
            yum_group_file_names=self.yum_group_file_names,
            yum_root_depth=self.yum_root_depth)


def get_federated_rpm_repository(archive_browsing_enabled: Optional[bool] = None,
                                 blacked_out: Optional[bool] = None,
                                 calculate_yum_metadata: Optional[bool] = None,
                                 cdn_redirect: Optional[bool] = None,
                                 cleanup_on_delete: Optional[bool] = None,
                                 description: Optional[str] = None,
                                 download_direct: Optional[bool] = None,
                                 enable_file_lists_indexing: Optional[bool] = None,
                                 excludes_pattern: Optional[str] = None,
                                 includes_pattern: Optional[str] = None,
                                 key: Optional[str] = None,
                                 members: Optional[Sequence[pulumi.InputType['GetFederatedRpmRepositoryMemberArgs']]] = None,
                                 notes: Optional[str] = None,
                                 primary_keypair_ref: Optional[str] = None,
                                 priority_resolution: Optional[bool] = None,
                                 project_environments: Optional[Sequence[str]] = None,
                                 project_key: Optional[str] = None,
                                 property_sets: Optional[Sequence[str]] = None,
                                 repo_layout_ref: Optional[str] = None,
                                 secondary_keypair_ref: Optional[str] = None,
                                 xray_index: Optional[bool] = None,
                                 yum_group_file_names: Optional[str] = None,
                                 yum_root_depth: Optional[int] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFederatedRpmRepositoryResult:
    """
    Retrieves a federated Rpm repository.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_artifactory as artifactory

    federated_test_rpm_repo = artifactory.get_federated_rpm_repository(key="federated-test-rpm-repo")
    ```


    :param str key: the identity key of the repo.
    :param Sequence[pulumi.InputType['GetFederatedRpmRepositoryMemberArgs']] members: The list of Federated members and must contain this repository URL (configured base URL
           `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
           Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
           to set up Federated repositories correctly.
    """
    __args__ = dict()
    __args__['archiveBrowsingEnabled'] = archive_browsing_enabled
    __args__['blackedOut'] = blacked_out
    __args__['calculateYumMetadata'] = calculate_yum_metadata
    __args__['cdnRedirect'] = cdn_redirect
    __args__['cleanupOnDelete'] = cleanup_on_delete
    __args__['description'] = description
    __args__['downloadDirect'] = download_direct
    __args__['enableFileListsIndexing'] = enable_file_lists_indexing
    __args__['excludesPattern'] = excludes_pattern
    __args__['includesPattern'] = includes_pattern
    __args__['key'] = key
    __args__['members'] = members
    __args__['notes'] = notes
    __args__['primaryKeypairRef'] = primary_keypair_ref
    __args__['priorityResolution'] = priority_resolution
    __args__['projectEnvironments'] = project_environments
    __args__['projectKey'] = project_key
    __args__['propertySets'] = property_sets
    __args__['repoLayoutRef'] = repo_layout_ref
    __args__['secondaryKeypairRef'] = secondary_keypair_ref
    __args__['xrayIndex'] = xray_index
    __args__['yumGroupFileNames'] = yum_group_file_names
    __args__['yumRootDepth'] = yum_root_depth
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('artifactory:index/getFederatedRpmRepository:getFederatedRpmRepository', __args__, opts=opts, typ=GetFederatedRpmRepositoryResult).value

    return AwaitableGetFederatedRpmRepositoryResult(
        archive_browsing_enabled=pulumi.get(__ret__, 'archive_browsing_enabled'),
        blacked_out=pulumi.get(__ret__, 'blacked_out'),
        calculate_yum_metadata=pulumi.get(__ret__, 'calculate_yum_metadata'),
        cdn_redirect=pulumi.get(__ret__, 'cdn_redirect'),
        cleanup_on_delete=pulumi.get(__ret__, 'cleanup_on_delete'),
        description=pulumi.get(__ret__, 'description'),
        download_direct=pulumi.get(__ret__, 'download_direct'),
        enable_file_lists_indexing=pulumi.get(__ret__, 'enable_file_lists_indexing'),
        excludes_pattern=pulumi.get(__ret__, 'excludes_pattern'),
        id=pulumi.get(__ret__, 'id'),
        includes_pattern=pulumi.get(__ret__, 'includes_pattern'),
        key=pulumi.get(__ret__, 'key'),
        members=pulumi.get(__ret__, 'members'),
        notes=pulumi.get(__ret__, 'notes'),
        package_type=pulumi.get(__ret__, 'package_type'),
        primary_keypair_ref=pulumi.get(__ret__, 'primary_keypair_ref'),
        priority_resolution=pulumi.get(__ret__, 'priority_resolution'),
        project_environments=pulumi.get(__ret__, 'project_environments'),
        project_key=pulumi.get(__ret__, 'project_key'),
        property_sets=pulumi.get(__ret__, 'property_sets'),
        repo_layout_ref=pulumi.get(__ret__, 'repo_layout_ref'),
        secondary_keypair_ref=pulumi.get(__ret__, 'secondary_keypair_ref'),
        xray_index=pulumi.get(__ret__, 'xray_index'),
        yum_group_file_names=pulumi.get(__ret__, 'yum_group_file_names'),
        yum_root_depth=pulumi.get(__ret__, 'yum_root_depth'))


@_utilities.lift_output_func(get_federated_rpm_repository)
def get_federated_rpm_repository_output(archive_browsing_enabled: Optional[pulumi.Input[Optional[bool]]] = None,
                                        blacked_out: Optional[pulumi.Input[Optional[bool]]] = None,
                                        calculate_yum_metadata: Optional[pulumi.Input[Optional[bool]]] = None,
                                        cdn_redirect: Optional[pulumi.Input[Optional[bool]]] = None,
                                        cleanup_on_delete: Optional[pulumi.Input[Optional[bool]]] = None,
                                        description: Optional[pulumi.Input[Optional[str]]] = None,
                                        download_direct: Optional[pulumi.Input[Optional[bool]]] = None,
                                        enable_file_lists_indexing: Optional[pulumi.Input[Optional[bool]]] = None,
                                        excludes_pattern: Optional[pulumi.Input[Optional[str]]] = None,
                                        includes_pattern: Optional[pulumi.Input[Optional[str]]] = None,
                                        key: Optional[pulumi.Input[str]] = None,
                                        members: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetFederatedRpmRepositoryMemberArgs']]]]] = None,
                                        notes: Optional[pulumi.Input[Optional[str]]] = None,
                                        primary_keypair_ref: Optional[pulumi.Input[Optional[str]]] = None,
                                        priority_resolution: Optional[pulumi.Input[Optional[bool]]] = None,
                                        project_environments: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                        project_key: Optional[pulumi.Input[Optional[str]]] = None,
                                        property_sets: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                        repo_layout_ref: Optional[pulumi.Input[Optional[str]]] = None,
                                        secondary_keypair_ref: Optional[pulumi.Input[Optional[str]]] = None,
                                        xray_index: Optional[pulumi.Input[Optional[bool]]] = None,
                                        yum_group_file_names: Optional[pulumi.Input[Optional[str]]] = None,
                                        yum_root_depth: Optional[pulumi.Input[Optional[int]]] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFederatedRpmRepositoryResult]:
    """
    Retrieves a federated Rpm repository.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_artifactory as artifactory

    federated_test_rpm_repo = artifactory.get_federated_rpm_repository(key="federated-test-rpm-repo")
    ```


    :param str key: the identity key of the repo.
    :param Sequence[pulumi.InputType['GetFederatedRpmRepositoryMemberArgs']] members: The list of Federated members and must contain this repository URL (configured base URL
           `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
           Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
           to set up Federated repositories correctly.
    """
    ...
