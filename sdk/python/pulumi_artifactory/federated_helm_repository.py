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

__all__ = ['FederatedHelmRepositoryArgs', 'FederatedHelmRepository']

@pulumi.input_type
class FederatedHelmRepositoryArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 members: pulumi.Input[Sequence[pulumi.Input['FederatedHelmRepositoryMemberArgs']]],
                 archive_browsing_enabled: Optional[pulumi.Input[bool]] = None,
                 blacked_out: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 download_direct: Optional[pulumi.Input[bool]] = None,
                 excludes_pattern: Optional[pulumi.Input[str]] = None,
                 includes_pattern: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 priority_resolution: Optional[pulumi.Input[bool]] = None,
                 project_environments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project_key: Optional[pulumi.Input[str]] = None,
                 property_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 repo_layout_ref: Optional[pulumi.Input[str]] = None,
                 xray_index: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a FederatedHelmRepository resource.
        :param pulumi.Input[str] key: - the identity key of the repo
        :param pulumi.Input[Sequence[pulumi.Input['FederatedHelmRepositoryMemberArgs']]] members: - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
        :param pulumi.Input[bool] archive_browsing_enabled: When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
               therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
               security (e.g., cross-site scripting attacks).
        :param pulumi.Input[bool] priority_resolution: Setting repositories with priority will cause metadata to be merged only from repositories set with this field
        :param pulumi.Input[Sequence[pulumi.Input[str]]] project_environments: Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
        :param pulumi.Input[str] project_key: Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
               with project key, separated by a dash.
        :param pulumi.Input[str] repo_layout_ref: Repository layout key for the local repository
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "members", members)
        if archive_browsing_enabled is not None:
            pulumi.set(__self__, "archive_browsing_enabled", archive_browsing_enabled)
        if blacked_out is not None:
            pulumi.set(__self__, "blacked_out", blacked_out)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if download_direct is not None:
            pulumi.set(__self__, "download_direct", download_direct)
        if excludes_pattern is not None:
            pulumi.set(__self__, "excludes_pattern", excludes_pattern)
        if includes_pattern is not None:
            pulumi.set(__self__, "includes_pattern", includes_pattern)
        if notes is not None:
            pulumi.set(__self__, "notes", notes)
        if priority_resolution is not None:
            pulumi.set(__self__, "priority_resolution", priority_resolution)
        if project_environments is not None:
            pulumi.set(__self__, "project_environments", project_environments)
        if project_key is not None:
            pulumi.set(__self__, "project_key", project_key)
        if property_sets is not None:
            pulumi.set(__self__, "property_sets", property_sets)
        if repo_layout_ref is not None:
            pulumi.set(__self__, "repo_layout_ref", repo_layout_ref)
        if xray_index is not None:
            pulumi.set(__self__, "xray_index", xray_index)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        - the identity key of the repo
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Input[Sequence[pulumi.Input['FederatedHelmRepositoryMemberArgs']]]:
        """
        - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: pulumi.Input[Sequence[pulumi.Input['FederatedHelmRepositoryMemberArgs']]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter(name="archiveBrowsingEnabled")
    def archive_browsing_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
        therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
        security (e.g., cross-site scripting attacks).
        """
        return pulumi.get(self, "archive_browsing_enabled")

    @archive_browsing_enabled.setter
    def archive_browsing_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "archive_browsing_enabled", value)

    @property
    @pulumi.getter(name="blackedOut")
    def blacked_out(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "blacked_out")

    @blacked_out.setter
    def blacked_out(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "blacked_out", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="downloadDirect")
    def download_direct(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "download_direct")

    @download_direct.setter
    def download_direct(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "download_direct", value)

    @property
    @pulumi.getter(name="excludesPattern")
    def excludes_pattern(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "excludes_pattern")

    @excludes_pattern.setter
    def excludes_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "excludes_pattern", value)

    @property
    @pulumi.getter(name="includesPattern")
    def includes_pattern(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "includes_pattern")

    @includes_pattern.setter
    def includes_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "includes_pattern", value)

    @property
    @pulumi.getter
    def notes(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "notes")

    @notes.setter
    def notes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notes", value)

    @property
    @pulumi.getter(name="priorityResolution")
    def priority_resolution(self) -> Optional[pulumi.Input[bool]]:
        """
        Setting repositories with priority will cause metadata to be merged only from repositories set with this field
        """
        return pulumi.get(self, "priority_resolution")

    @priority_resolution.setter
    def priority_resolution(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "priority_resolution", value)

    @property
    @pulumi.getter(name="projectEnvironments")
    def project_environments(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
        """
        return pulumi.get(self, "project_environments")

    @project_environments.setter
    def project_environments(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "project_environments", value)

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> Optional[pulumi.Input[str]]:
        """
        Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
        with project key, separated by a dash.
        """
        return pulumi.get(self, "project_key")

    @project_key.setter
    def project_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_key", value)

    @property
    @pulumi.getter(name="propertySets")
    def property_sets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "property_sets")

    @property_sets.setter
    def property_sets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "property_sets", value)

    @property
    @pulumi.getter(name="repoLayoutRef")
    def repo_layout_ref(self) -> Optional[pulumi.Input[str]]:
        """
        Repository layout key for the local repository
        """
        return pulumi.get(self, "repo_layout_ref")

    @repo_layout_ref.setter
    def repo_layout_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repo_layout_ref", value)

    @property
    @pulumi.getter(name="xrayIndex")
    def xray_index(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "xray_index")

    @xray_index.setter
    def xray_index(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "xray_index", value)


@pulumi.input_type
class _FederatedHelmRepositoryState:
    def __init__(__self__, *,
                 archive_browsing_enabled: Optional[pulumi.Input[bool]] = None,
                 blacked_out: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 download_direct: Optional[pulumi.Input[bool]] = None,
                 excludes_pattern: Optional[pulumi.Input[str]] = None,
                 includes_pattern: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input['FederatedHelmRepositoryMemberArgs']]]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 package_type: Optional[pulumi.Input[str]] = None,
                 priority_resolution: Optional[pulumi.Input[bool]] = None,
                 project_environments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project_key: Optional[pulumi.Input[str]] = None,
                 property_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 repo_layout_ref: Optional[pulumi.Input[str]] = None,
                 xray_index: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering FederatedHelmRepository resources.
        :param pulumi.Input[bool] archive_browsing_enabled: When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
               therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
               security (e.g., cross-site scripting attacks).
        :param pulumi.Input[str] key: - the identity key of the repo
        :param pulumi.Input[Sequence[pulumi.Input['FederatedHelmRepositoryMemberArgs']]] members: - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
        :param pulumi.Input[bool] priority_resolution: Setting repositories with priority will cause metadata to be merged only from repositories set with this field
        :param pulumi.Input[Sequence[pulumi.Input[str]]] project_environments: Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
        :param pulumi.Input[str] project_key: Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
               with project key, separated by a dash.
        :param pulumi.Input[str] repo_layout_ref: Repository layout key for the local repository
        """
        if archive_browsing_enabled is not None:
            pulumi.set(__self__, "archive_browsing_enabled", archive_browsing_enabled)
        if blacked_out is not None:
            pulumi.set(__self__, "blacked_out", blacked_out)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if download_direct is not None:
            pulumi.set(__self__, "download_direct", download_direct)
        if excludes_pattern is not None:
            pulumi.set(__self__, "excludes_pattern", excludes_pattern)
        if includes_pattern is not None:
            pulumi.set(__self__, "includes_pattern", includes_pattern)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if notes is not None:
            pulumi.set(__self__, "notes", notes)
        if package_type is not None:
            pulumi.set(__self__, "package_type", package_type)
        if priority_resolution is not None:
            pulumi.set(__self__, "priority_resolution", priority_resolution)
        if project_environments is not None:
            pulumi.set(__self__, "project_environments", project_environments)
        if project_key is not None:
            pulumi.set(__self__, "project_key", project_key)
        if property_sets is not None:
            pulumi.set(__self__, "property_sets", property_sets)
        if repo_layout_ref is not None:
            pulumi.set(__self__, "repo_layout_ref", repo_layout_ref)
        if xray_index is not None:
            pulumi.set(__self__, "xray_index", xray_index)

    @property
    @pulumi.getter(name="archiveBrowsingEnabled")
    def archive_browsing_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
        therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
        security (e.g., cross-site scripting attacks).
        """
        return pulumi.get(self, "archive_browsing_enabled")

    @archive_browsing_enabled.setter
    def archive_browsing_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "archive_browsing_enabled", value)

    @property
    @pulumi.getter(name="blackedOut")
    def blacked_out(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "blacked_out")

    @blacked_out.setter
    def blacked_out(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "blacked_out", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="downloadDirect")
    def download_direct(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "download_direct")

    @download_direct.setter
    def download_direct(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "download_direct", value)

    @property
    @pulumi.getter(name="excludesPattern")
    def excludes_pattern(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "excludes_pattern")

    @excludes_pattern.setter
    def excludes_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "excludes_pattern", value)

    @property
    @pulumi.getter(name="includesPattern")
    def includes_pattern(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "includes_pattern")

    @includes_pattern.setter
    def includes_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "includes_pattern", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        - the identity key of the repo
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FederatedHelmRepositoryMemberArgs']]]]:
        """
        - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FederatedHelmRepositoryMemberArgs']]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def notes(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "notes")

    @notes.setter
    def notes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notes", value)

    @property
    @pulumi.getter(name="packageType")
    def package_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "package_type")

    @package_type.setter
    def package_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "package_type", value)

    @property
    @pulumi.getter(name="priorityResolution")
    def priority_resolution(self) -> Optional[pulumi.Input[bool]]:
        """
        Setting repositories with priority will cause metadata to be merged only from repositories set with this field
        """
        return pulumi.get(self, "priority_resolution")

    @priority_resolution.setter
    def priority_resolution(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "priority_resolution", value)

    @property
    @pulumi.getter(name="projectEnvironments")
    def project_environments(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
        """
        return pulumi.get(self, "project_environments")

    @project_environments.setter
    def project_environments(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "project_environments", value)

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> Optional[pulumi.Input[str]]:
        """
        Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
        with project key, separated by a dash.
        """
        return pulumi.get(self, "project_key")

    @project_key.setter
    def project_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_key", value)

    @property
    @pulumi.getter(name="propertySets")
    def property_sets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "property_sets")

    @property_sets.setter
    def property_sets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "property_sets", value)

    @property
    @pulumi.getter(name="repoLayoutRef")
    def repo_layout_ref(self) -> Optional[pulumi.Input[str]]:
        """
        Repository layout key for the local repository
        """
        return pulumi.get(self, "repo_layout_ref")

    @repo_layout_ref.setter
    def repo_layout_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repo_layout_ref", value)

    @property
    @pulumi.getter(name="xrayIndex")
    def xray_index(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "xray_index")

    @xray_index.setter
    def xray_index(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "xray_index", value)


class FederatedHelmRepository(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 archive_browsing_enabled: Optional[pulumi.Input[bool]] = None,
                 blacked_out: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 download_direct: Optional[pulumi.Input[bool]] = None,
                 excludes_pattern: Optional[pulumi.Input[str]] = None,
                 includes_pattern: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FederatedHelmRepositoryMemberArgs']]]]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 priority_resolution: Optional[pulumi.Input[bool]] = None,
                 project_environments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project_key: Optional[pulumi.Input[str]] = None,
                 property_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 repo_layout_ref: Optional[pulumi.Input[str]] = None,
                 xray_index: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        ## # Artifactory Federated Helm Repository Resource

        Creates a federated Helm repository

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        terraform_federated_test_helm_repo = artifactory.FederatedHelmRepository("terraform-federated-test-helm-repo",
            key="terraform-federated-test-helm-repo",
            members=[
                artifactory.FederatedHelmRepositoryMemberArgs(
                    enabled=True,
                    url="http://tempurl.org/artifactory/terraform-federated-test-helm-repo",
                ),
                artifactory.FederatedHelmRepositoryMemberArgs(
                    enabled=True,
                    url="http://tempurl2.org/artifactory/terraform-federated-test-helm-repo-2",
                ),
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] archive_browsing_enabled: When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
               therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
               security (e.g., cross-site scripting attacks).
        :param pulumi.Input[str] key: - the identity key of the repo
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FederatedHelmRepositoryMemberArgs']]]] members: - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
        :param pulumi.Input[bool] priority_resolution: Setting repositories with priority will cause metadata to be merged only from repositories set with this field
        :param pulumi.Input[Sequence[pulumi.Input[str]]] project_environments: Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
        :param pulumi.Input[str] project_key: Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
               with project key, separated by a dash.
        :param pulumi.Input[str] repo_layout_ref: Repository layout key for the local repository
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FederatedHelmRepositoryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Artifactory Federated Helm Repository Resource

        Creates a federated Helm repository

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        terraform_federated_test_helm_repo = artifactory.FederatedHelmRepository("terraform-federated-test-helm-repo",
            key="terraform-federated-test-helm-repo",
            members=[
                artifactory.FederatedHelmRepositoryMemberArgs(
                    enabled=True,
                    url="http://tempurl.org/artifactory/terraform-federated-test-helm-repo",
                ),
                artifactory.FederatedHelmRepositoryMemberArgs(
                    enabled=True,
                    url="http://tempurl2.org/artifactory/terraform-federated-test-helm-repo-2",
                ),
            ])
        ```

        :param str resource_name: The name of the resource.
        :param FederatedHelmRepositoryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FederatedHelmRepositoryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 archive_browsing_enabled: Optional[pulumi.Input[bool]] = None,
                 blacked_out: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 download_direct: Optional[pulumi.Input[bool]] = None,
                 excludes_pattern: Optional[pulumi.Input[str]] = None,
                 includes_pattern: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FederatedHelmRepositoryMemberArgs']]]]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 priority_resolution: Optional[pulumi.Input[bool]] = None,
                 project_environments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project_key: Optional[pulumi.Input[str]] = None,
                 property_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 repo_layout_ref: Optional[pulumi.Input[str]] = None,
                 xray_index: Optional[pulumi.Input[bool]] = None,
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
            __props__ = FederatedHelmRepositoryArgs.__new__(FederatedHelmRepositoryArgs)

            __props__.__dict__["archive_browsing_enabled"] = archive_browsing_enabled
            __props__.__dict__["blacked_out"] = blacked_out
            __props__.__dict__["description"] = description
            __props__.__dict__["download_direct"] = download_direct
            __props__.__dict__["excludes_pattern"] = excludes_pattern
            __props__.__dict__["includes_pattern"] = includes_pattern
            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            if members is None and not opts.urn:
                raise TypeError("Missing required property 'members'")
            __props__.__dict__["members"] = members
            __props__.__dict__["notes"] = notes
            __props__.__dict__["priority_resolution"] = priority_resolution
            __props__.__dict__["project_environments"] = project_environments
            __props__.__dict__["project_key"] = project_key
            __props__.__dict__["property_sets"] = property_sets
            __props__.__dict__["repo_layout_ref"] = repo_layout_ref
            __props__.__dict__["xray_index"] = xray_index
            __props__.__dict__["package_type"] = None
        super(FederatedHelmRepository, __self__).__init__(
            'artifactory:index/federatedHelmRepository:FederatedHelmRepository',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            archive_browsing_enabled: Optional[pulumi.Input[bool]] = None,
            blacked_out: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            download_direct: Optional[pulumi.Input[bool]] = None,
            excludes_pattern: Optional[pulumi.Input[str]] = None,
            includes_pattern: Optional[pulumi.Input[str]] = None,
            key: Optional[pulumi.Input[str]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FederatedHelmRepositoryMemberArgs']]]]] = None,
            notes: Optional[pulumi.Input[str]] = None,
            package_type: Optional[pulumi.Input[str]] = None,
            priority_resolution: Optional[pulumi.Input[bool]] = None,
            project_environments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            project_key: Optional[pulumi.Input[str]] = None,
            property_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            repo_layout_ref: Optional[pulumi.Input[str]] = None,
            xray_index: Optional[pulumi.Input[bool]] = None) -> 'FederatedHelmRepository':
        """
        Get an existing FederatedHelmRepository resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] archive_browsing_enabled: When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
               therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
               security (e.g., cross-site scripting attacks).
        :param pulumi.Input[str] key: - the identity key of the repo
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FederatedHelmRepositoryMemberArgs']]]] members: - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
        :param pulumi.Input[bool] priority_resolution: Setting repositories with priority will cause metadata to be merged only from repositories set with this field
        :param pulumi.Input[Sequence[pulumi.Input[str]]] project_environments: Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
        :param pulumi.Input[str] project_key: Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
               with project key, separated by a dash.
        :param pulumi.Input[str] repo_layout_ref: Repository layout key for the local repository
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FederatedHelmRepositoryState.__new__(_FederatedHelmRepositoryState)

        __props__.__dict__["archive_browsing_enabled"] = archive_browsing_enabled
        __props__.__dict__["blacked_out"] = blacked_out
        __props__.__dict__["description"] = description
        __props__.__dict__["download_direct"] = download_direct
        __props__.__dict__["excludes_pattern"] = excludes_pattern
        __props__.__dict__["includes_pattern"] = includes_pattern
        __props__.__dict__["key"] = key
        __props__.__dict__["members"] = members
        __props__.__dict__["notes"] = notes
        __props__.__dict__["package_type"] = package_type
        __props__.__dict__["priority_resolution"] = priority_resolution
        __props__.__dict__["project_environments"] = project_environments
        __props__.__dict__["project_key"] = project_key
        __props__.__dict__["property_sets"] = property_sets
        __props__.__dict__["repo_layout_ref"] = repo_layout_ref
        __props__.__dict__["xray_index"] = xray_index
        return FederatedHelmRepository(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="archiveBrowsingEnabled")
    def archive_browsing_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
        therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
        security (e.g., cross-site scripting attacks).
        """
        return pulumi.get(self, "archive_browsing_enabled")

    @property
    @pulumi.getter(name="blackedOut")
    def blacked_out(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "blacked_out")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="downloadDirect")
    def download_direct(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "download_direct")

    @property
    @pulumi.getter(name="excludesPattern")
    def excludes_pattern(self) -> pulumi.Output[str]:
        return pulumi.get(self, "excludes_pattern")

    @property
    @pulumi.getter(name="includesPattern")
    def includes_pattern(self) -> pulumi.Output[str]:
        return pulumi.get(self, "includes_pattern")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        - the identity key of the repo
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence['outputs.FederatedHelmRepositoryMember']]:
        """
        - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def notes(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter(name="packageType")
    def package_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "package_type")

    @property
    @pulumi.getter(name="priorityResolution")
    def priority_resolution(self) -> pulumi.Output[Optional[bool]]:
        """
        Setting repositories with priority will cause metadata to be merged only from repositories set with this field
        """
        return pulumi.get(self, "priority_resolution")

    @property
    @pulumi.getter(name="projectEnvironments")
    def project_environments(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
        """
        return pulumi.get(self, "project_environments")

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> pulumi.Output[Optional[str]]:
        """
        Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
        with project key, separated by a dash.
        """
        return pulumi.get(self, "project_key")

    @property
    @pulumi.getter(name="propertySets")
    def property_sets(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "property_sets")

    @property
    @pulumi.getter(name="repoLayoutRef")
    def repo_layout_ref(self) -> pulumi.Output[Optional[str]]:
        """
        Repository layout key for the local repository
        """
        return pulumi.get(self, "repo_layout_ref")

    @property
    @pulumi.getter(name="xrayIndex")
    def xray_index(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "xray_index")

