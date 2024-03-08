# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetLocalSbtRepositoryResult',
    'AwaitableGetLocalSbtRepositoryResult',
    'get_local_sbt_repository',
    'get_local_sbt_repository_output',
]

@pulumi.output_type
class GetLocalSbtRepositoryResult:
    """
    A collection of values returned by getLocalSbtRepository.
    """
    def __init__(__self__, archive_browsing_enabled=None, blacked_out=None, cdn_redirect=None, checksum_policy_type=None, description=None, download_direct=None, excludes_pattern=None, handle_releases=None, handle_snapshots=None, id=None, includes_pattern=None, key=None, max_unique_snapshots=None, notes=None, package_type=None, priority_resolution=None, project_environments=None, project_key=None, property_sets=None, repo_layout_ref=None, snapshot_version_behavior=None, suppress_pom_consistency_checks=None, xray_index=None):
        if archive_browsing_enabled and not isinstance(archive_browsing_enabled, bool):
            raise TypeError("Expected argument 'archive_browsing_enabled' to be a bool")
        pulumi.set(__self__, "archive_browsing_enabled", archive_browsing_enabled)
        if blacked_out and not isinstance(blacked_out, bool):
            raise TypeError("Expected argument 'blacked_out' to be a bool")
        pulumi.set(__self__, "blacked_out", blacked_out)
        if cdn_redirect and not isinstance(cdn_redirect, bool):
            raise TypeError("Expected argument 'cdn_redirect' to be a bool")
        pulumi.set(__self__, "cdn_redirect", cdn_redirect)
        if checksum_policy_type and not isinstance(checksum_policy_type, str):
            raise TypeError("Expected argument 'checksum_policy_type' to be a str")
        pulumi.set(__self__, "checksum_policy_type", checksum_policy_type)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if download_direct and not isinstance(download_direct, bool):
            raise TypeError("Expected argument 'download_direct' to be a bool")
        pulumi.set(__self__, "download_direct", download_direct)
        if excludes_pattern and not isinstance(excludes_pattern, str):
            raise TypeError("Expected argument 'excludes_pattern' to be a str")
        pulumi.set(__self__, "excludes_pattern", excludes_pattern)
        if handle_releases and not isinstance(handle_releases, bool):
            raise TypeError("Expected argument 'handle_releases' to be a bool")
        pulumi.set(__self__, "handle_releases", handle_releases)
        if handle_snapshots and not isinstance(handle_snapshots, bool):
            raise TypeError("Expected argument 'handle_snapshots' to be a bool")
        pulumi.set(__self__, "handle_snapshots", handle_snapshots)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if includes_pattern and not isinstance(includes_pattern, str):
            raise TypeError("Expected argument 'includes_pattern' to be a str")
        pulumi.set(__self__, "includes_pattern", includes_pattern)
        if key and not isinstance(key, str):
            raise TypeError("Expected argument 'key' to be a str")
        pulumi.set(__self__, "key", key)
        if max_unique_snapshots and not isinstance(max_unique_snapshots, int):
            raise TypeError("Expected argument 'max_unique_snapshots' to be a int")
        pulumi.set(__self__, "max_unique_snapshots", max_unique_snapshots)
        if notes and not isinstance(notes, str):
            raise TypeError("Expected argument 'notes' to be a str")
        pulumi.set(__self__, "notes", notes)
        if package_type and not isinstance(package_type, str):
            raise TypeError("Expected argument 'package_type' to be a str")
        pulumi.set(__self__, "package_type", package_type)
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
        if snapshot_version_behavior and not isinstance(snapshot_version_behavior, str):
            raise TypeError("Expected argument 'snapshot_version_behavior' to be a str")
        pulumi.set(__self__, "snapshot_version_behavior", snapshot_version_behavior)
        if suppress_pom_consistency_checks and not isinstance(suppress_pom_consistency_checks, bool):
            raise TypeError("Expected argument 'suppress_pom_consistency_checks' to be a bool")
        pulumi.set(__self__, "suppress_pom_consistency_checks", suppress_pom_consistency_checks)
        if xray_index and not isinstance(xray_index, bool):
            raise TypeError("Expected argument 'xray_index' to be a bool")
        pulumi.set(__self__, "xray_index", xray_index)

    @property
    @pulumi.getter(name="archiveBrowsingEnabled")
    def archive_browsing_enabled(self) -> Optional[bool]:
        return pulumi.get(self, "archive_browsing_enabled")

    @property
    @pulumi.getter(name="blackedOut")
    def blacked_out(self) -> Optional[bool]:
        return pulumi.get(self, "blacked_out")

    @property
    @pulumi.getter(name="cdnRedirect")
    def cdn_redirect(self) -> Optional[bool]:
        return pulumi.get(self, "cdn_redirect")

    @property
    @pulumi.getter(name="checksumPolicyType")
    def checksum_policy_type(self) -> Optional[str]:
        """
        Checksum policy determines how Artifactory behaves when a client checksum for a deployed
        resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
        `client-checksums` and `generated-checksums`. For more details, please refer
        to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
        .
        """
        return pulumi.get(self, "checksum_policy_type")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="downloadDirect")
    def download_direct(self) -> Optional[bool]:
        return pulumi.get(self, "download_direct")

    @property
    @pulumi.getter(name="excludesPattern")
    def excludes_pattern(self) -> Optional[str]:
        return pulumi.get(self, "excludes_pattern")

    @property
    @pulumi.getter(name="handleReleases")
    def handle_releases(self) -> Optional[bool]:
        """
        If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`
        .
        """
        return pulumi.get(self, "handle_releases")

    @property
    @pulumi.getter(name="handleSnapshots")
    def handle_snapshots(self) -> Optional[bool]:
        """
        If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default
        is `true`.
        """
        return pulumi.get(self, "handle_snapshots")

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
    @pulumi.getter(name="maxUniqueSnapshots")
    def max_unique_snapshots(self) -> Optional[int]:
        """
        The maximum number of unique snapshots of a single artifact to store. Once the number of
        snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and
        unique snapshots are not cleaned up.
        """
        return pulumi.get(self, "max_unique_snapshots")

    @property
    @pulumi.getter
    def notes(self) -> Optional[str]:
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter(name="packageType")
    def package_type(self) -> str:
        return pulumi.get(self, "package_type")

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
    @pulumi.getter(name="snapshotVersionBehavior")
    def snapshot_version_behavior(self) -> Optional[str]:
        """
        Specifies the naming convention for Maven SNAPSHOT versions. The options are
        -
        """
        return pulumi.get(self, "snapshot_version_behavior")

    @property
    @pulumi.getter(name="suppressPomConsistencyChecks")
    def suppress_pom_consistency_checks(self) -> Optional[bool]:
        """
        By default, Artifactory keeps your repositories healthy by refusing POMs with
        incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed
        path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting the
        Suppress POM Consistency Checks checkbox. True by default for Gradle repository.
        """
        return pulumi.get(self, "suppress_pom_consistency_checks")

    @property
    @pulumi.getter(name="xrayIndex")
    def xray_index(self) -> Optional[bool]:
        return pulumi.get(self, "xray_index")


class AwaitableGetLocalSbtRepositoryResult(GetLocalSbtRepositoryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLocalSbtRepositoryResult(
            archive_browsing_enabled=self.archive_browsing_enabled,
            blacked_out=self.blacked_out,
            cdn_redirect=self.cdn_redirect,
            checksum_policy_type=self.checksum_policy_type,
            description=self.description,
            download_direct=self.download_direct,
            excludes_pattern=self.excludes_pattern,
            handle_releases=self.handle_releases,
            handle_snapshots=self.handle_snapshots,
            id=self.id,
            includes_pattern=self.includes_pattern,
            key=self.key,
            max_unique_snapshots=self.max_unique_snapshots,
            notes=self.notes,
            package_type=self.package_type,
            priority_resolution=self.priority_resolution,
            project_environments=self.project_environments,
            project_key=self.project_key,
            property_sets=self.property_sets,
            repo_layout_ref=self.repo_layout_ref,
            snapshot_version_behavior=self.snapshot_version_behavior,
            suppress_pom_consistency_checks=self.suppress_pom_consistency_checks,
            xray_index=self.xray_index)


def get_local_sbt_repository(archive_browsing_enabled: Optional[bool] = None,
                             blacked_out: Optional[bool] = None,
                             cdn_redirect: Optional[bool] = None,
                             checksum_policy_type: Optional[str] = None,
                             description: Optional[str] = None,
                             download_direct: Optional[bool] = None,
                             excludes_pattern: Optional[str] = None,
                             handle_releases: Optional[bool] = None,
                             handle_snapshots: Optional[bool] = None,
                             includes_pattern: Optional[str] = None,
                             key: Optional[str] = None,
                             max_unique_snapshots: Optional[int] = None,
                             notes: Optional[str] = None,
                             priority_resolution: Optional[bool] = None,
                             project_environments: Optional[Sequence[str]] = None,
                             project_key: Optional[str] = None,
                             property_sets: Optional[Sequence[str]] = None,
                             repo_layout_ref: Optional[str] = None,
                             snapshot_version_behavior: Optional[str] = None,
                             suppress_pom_consistency_checks: Optional[bool] = None,
                             xray_index: Optional[bool] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLocalSbtRepositoryResult:
    """
    Retrieves a local Sbt repository.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_artifactory as artifactory

    local_test_sbt_repo = artifactory.get_local_sbt_repository(key="local-test-sbt-repo")
    ```
    <!--End PulumiCodeChooser -->


    :param str checksum_policy_type: Checksum policy determines how Artifactory behaves when a client checksum for a deployed
           resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
           `client-checksums` and `generated-checksums`. For more details, please refer
           to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
           .
    :param bool handle_releases: If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`
           .
    :param bool handle_snapshots: If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default
           is `true`.
    :param str key: the identity key of the repo.
    :param int max_unique_snapshots: The maximum number of unique snapshots of a single artifact to store. Once the number of
           snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and
           unique snapshots are not cleaned up.
    :param str snapshot_version_behavior: Specifies the naming convention for Maven SNAPSHOT versions. The options are
           -
    :param bool suppress_pom_consistency_checks: By default, Artifactory keeps your repositories healthy by refusing POMs with
           incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed
           path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting the
           Suppress POM Consistency Checks checkbox. True by default for Gradle repository.
    """
    __args__ = dict()
    __args__['archiveBrowsingEnabled'] = archive_browsing_enabled
    __args__['blackedOut'] = blacked_out
    __args__['cdnRedirect'] = cdn_redirect
    __args__['checksumPolicyType'] = checksum_policy_type
    __args__['description'] = description
    __args__['downloadDirect'] = download_direct
    __args__['excludesPattern'] = excludes_pattern
    __args__['handleReleases'] = handle_releases
    __args__['handleSnapshots'] = handle_snapshots
    __args__['includesPattern'] = includes_pattern
    __args__['key'] = key
    __args__['maxUniqueSnapshots'] = max_unique_snapshots
    __args__['notes'] = notes
    __args__['priorityResolution'] = priority_resolution
    __args__['projectEnvironments'] = project_environments
    __args__['projectKey'] = project_key
    __args__['propertySets'] = property_sets
    __args__['repoLayoutRef'] = repo_layout_ref
    __args__['snapshotVersionBehavior'] = snapshot_version_behavior
    __args__['suppressPomConsistencyChecks'] = suppress_pom_consistency_checks
    __args__['xrayIndex'] = xray_index
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('artifactory:index/getLocalSbtRepository:getLocalSbtRepository', __args__, opts=opts, typ=GetLocalSbtRepositoryResult).value

    return AwaitableGetLocalSbtRepositoryResult(
        archive_browsing_enabled=pulumi.get(__ret__, 'archive_browsing_enabled'),
        blacked_out=pulumi.get(__ret__, 'blacked_out'),
        cdn_redirect=pulumi.get(__ret__, 'cdn_redirect'),
        checksum_policy_type=pulumi.get(__ret__, 'checksum_policy_type'),
        description=pulumi.get(__ret__, 'description'),
        download_direct=pulumi.get(__ret__, 'download_direct'),
        excludes_pattern=pulumi.get(__ret__, 'excludes_pattern'),
        handle_releases=pulumi.get(__ret__, 'handle_releases'),
        handle_snapshots=pulumi.get(__ret__, 'handle_snapshots'),
        id=pulumi.get(__ret__, 'id'),
        includes_pattern=pulumi.get(__ret__, 'includes_pattern'),
        key=pulumi.get(__ret__, 'key'),
        max_unique_snapshots=pulumi.get(__ret__, 'max_unique_snapshots'),
        notes=pulumi.get(__ret__, 'notes'),
        package_type=pulumi.get(__ret__, 'package_type'),
        priority_resolution=pulumi.get(__ret__, 'priority_resolution'),
        project_environments=pulumi.get(__ret__, 'project_environments'),
        project_key=pulumi.get(__ret__, 'project_key'),
        property_sets=pulumi.get(__ret__, 'property_sets'),
        repo_layout_ref=pulumi.get(__ret__, 'repo_layout_ref'),
        snapshot_version_behavior=pulumi.get(__ret__, 'snapshot_version_behavior'),
        suppress_pom_consistency_checks=pulumi.get(__ret__, 'suppress_pom_consistency_checks'),
        xray_index=pulumi.get(__ret__, 'xray_index'))


@_utilities.lift_output_func(get_local_sbt_repository)
def get_local_sbt_repository_output(archive_browsing_enabled: Optional[pulumi.Input[Optional[bool]]] = None,
                                    blacked_out: Optional[pulumi.Input[Optional[bool]]] = None,
                                    cdn_redirect: Optional[pulumi.Input[Optional[bool]]] = None,
                                    checksum_policy_type: Optional[pulumi.Input[Optional[str]]] = None,
                                    description: Optional[pulumi.Input[Optional[str]]] = None,
                                    download_direct: Optional[pulumi.Input[Optional[bool]]] = None,
                                    excludes_pattern: Optional[pulumi.Input[Optional[str]]] = None,
                                    handle_releases: Optional[pulumi.Input[Optional[bool]]] = None,
                                    handle_snapshots: Optional[pulumi.Input[Optional[bool]]] = None,
                                    includes_pattern: Optional[pulumi.Input[Optional[str]]] = None,
                                    key: Optional[pulumi.Input[str]] = None,
                                    max_unique_snapshots: Optional[pulumi.Input[Optional[int]]] = None,
                                    notes: Optional[pulumi.Input[Optional[str]]] = None,
                                    priority_resolution: Optional[pulumi.Input[Optional[bool]]] = None,
                                    project_environments: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                    project_key: Optional[pulumi.Input[Optional[str]]] = None,
                                    property_sets: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                    repo_layout_ref: Optional[pulumi.Input[Optional[str]]] = None,
                                    snapshot_version_behavior: Optional[pulumi.Input[Optional[str]]] = None,
                                    suppress_pom_consistency_checks: Optional[pulumi.Input[Optional[bool]]] = None,
                                    xray_index: Optional[pulumi.Input[Optional[bool]]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLocalSbtRepositoryResult]:
    """
    Retrieves a local Sbt repository.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_artifactory as artifactory

    local_test_sbt_repo = artifactory.get_local_sbt_repository(key="local-test-sbt-repo")
    ```
    <!--End PulumiCodeChooser -->


    :param str checksum_policy_type: Checksum policy determines how Artifactory behaves when a client checksum for a deployed
           resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
           `client-checksums` and `generated-checksums`. For more details, please refer
           to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
           .
    :param bool handle_releases: If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`
           .
    :param bool handle_snapshots: If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default
           is `true`.
    :param str key: the identity key of the repo.
    :param int max_unique_snapshots: The maximum number of unique snapshots of a single artifact to store. Once the number of
           snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and
           unique snapshots are not cleaned up.
    :param str snapshot_version_behavior: Specifies the naming convention for Maven SNAPSHOT versions. The options are
           -
    :param bool suppress_pom_consistency_checks: By default, Artifactory keeps your repositories healthy by refusing POMs with
           incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed
           path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting the
           Suppress POM Consistency Checks checkbox. True by default for Gradle repository.
    """
    ...
