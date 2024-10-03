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

__all__ = ['VirtualCondaRepositoryArgs', 'VirtualCondaRepository']

@pulumi.input_type
class VirtualCondaRepositoryArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 artifactory_requests_can_retrieve_remote_artifacts: Optional[pulumi.Input[bool]] = None,
                 default_deployment_repo: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 excludes_pattern: Optional[pulumi.Input[str]] = None,
                 includes_pattern: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 project_environments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project_key: Optional[pulumi.Input[str]] = None,
                 repo_layout_ref: Optional[pulumi.Input[str]] = None,
                 repositories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retrieval_cache_period_seconds: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a VirtualCondaRepository resource.
        :param pulumi.Input[str] key: A mandatory identifier for the repository that must be unique. It cannot begin with a number or
               contain spaces or special characters.
        :param pulumi.Input[bool] artifactory_requests_can_retrieve_remote_artifacts: Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
               another Artifactory instance.
        :param pulumi.Input[str] default_deployment_repo: Default repository to deploy artifacts.
        :param pulumi.Input[str] description: Public description.
        :param pulumi.Input[str] excludes_pattern: List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
               artifacts are excluded.
        :param pulumi.Input[str] includes_pattern: List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
               used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
        :param pulumi.Input[str] notes: Internal description.
        :param pulumi.Input[str] project_key: Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
               assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
        :param pulumi.Input[str] repo_layout_ref: Repository layout key for the virtual repository
        :param pulumi.Input[Sequence[pulumi.Input[str]]] repositories: The effective list of actual repositories included in this virtual repository.
        :param pulumi.Input[int] retrieval_cache_period_seconds: This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
        """
        pulumi.set(__self__, "key", key)
        if artifactory_requests_can_retrieve_remote_artifacts is not None:
            pulumi.set(__self__, "artifactory_requests_can_retrieve_remote_artifacts", artifactory_requests_can_retrieve_remote_artifacts)
        if default_deployment_repo is not None:
            pulumi.set(__self__, "default_deployment_repo", default_deployment_repo)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if excludes_pattern is not None:
            pulumi.set(__self__, "excludes_pattern", excludes_pattern)
        if includes_pattern is not None:
            pulumi.set(__self__, "includes_pattern", includes_pattern)
        if notes is not None:
            pulumi.set(__self__, "notes", notes)
        if project_environments is not None:
            pulumi.set(__self__, "project_environments", project_environments)
        if project_key is not None:
            pulumi.set(__self__, "project_key", project_key)
        if repo_layout_ref is not None:
            pulumi.set(__self__, "repo_layout_ref", repo_layout_ref)
        if repositories is not None:
            pulumi.set(__self__, "repositories", repositories)
        if retrieval_cache_period_seconds is not None:
            pulumi.set(__self__, "retrieval_cache_period_seconds", retrieval_cache_period_seconds)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        A mandatory identifier for the repository that must be unique. It cannot begin with a number or
        contain spaces or special characters.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="artifactoryRequestsCanRetrieveRemoteArtifacts")
    def artifactory_requests_can_retrieve_remote_artifacts(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
        another Artifactory instance.
        """
        return pulumi.get(self, "artifactory_requests_can_retrieve_remote_artifacts")

    @artifactory_requests_can_retrieve_remote_artifacts.setter
    def artifactory_requests_can_retrieve_remote_artifacts(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "artifactory_requests_can_retrieve_remote_artifacts", value)

    @property
    @pulumi.getter(name="defaultDeploymentRepo")
    def default_deployment_repo(self) -> Optional[pulumi.Input[str]]:
        """
        Default repository to deploy artifacts.
        """
        return pulumi.get(self, "default_deployment_repo")

    @default_deployment_repo.setter
    def default_deployment_repo(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_deployment_repo", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Public description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="excludesPattern")
    def excludes_pattern(self) -> Optional[pulumi.Input[str]]:
        """
        List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
        artifacts are excluded.
        """
        return pulumi.get(self, "excludes_pattern")

    @excludes_pattern.setter
    def excludes_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "excludes_pattern", value)

    @property
    @pulumi.getter(name="includesPattern")
    def includes_pattern(self) -> Optional[pulumi.Input[str]]:
        """
        List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
        used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
        """
        return pulumi.get(self, "includes_pattern")

    @includes_pattern.setter
    def includes_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "includes_pattern", value)

    @property
    @pulumi.getter
    def notes(self) -> Optional[pulumi.Input[str]]:
        """
        Internal description.
        """
        return pulumi.get(self, "notes")

    @notes.setter
    def notes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notes", value)

    @property
    @pulumi.getter(name="projectEnvironments")
    def project_environments(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "project_environments")

    @project_environments.setter
    def project_environments(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "project_environments", value)

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> Optional[pulumi.Input[str]]:
        """
        Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
        assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
        """
        return pulumi.get(self, "project_key")

    @project_key.setter
    def project_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_key", value)

    @property
    @pulumi.getter(name="repoLayoutRef")
    def repo_layout_ref(self) -> Optional[pulumi.Input[str]]:
        """
        Repository layout key for the virtual repository
        """
        return pulumi.get(self, "repo_layout_ref")

    @repo_layout_ref.setter
    def repo_layout_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repo_layout_ref", value)

    @property
    @pulumi.getter
    def repositories(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The effective list of actual repositories included in this virtual repository.
        """
        return pulumi.get(self, "repositories")

    @repositories.setter
    def repositories(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "repositories", value)

    @property
    @pulumi.getter(name="retrievalCachePeriodSeconds")
    def retrieval_cache_period_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
        """
        return pulumi.get(self, "retrieval_cache_period_seconds")

    @retrieval_cache_period_seconds.setter
    def retrieval_cache_period_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retrieval_cache_period_seconds", value)


@pulumi.input_type
class _VirtualCondaRepositoryState:
    def __init__(__self__, *,
                 artifactory_requests_can_retrieve_remote_artifacts: Optional[pulumi.Input[bool]] = None,
                 default_deployment_repo: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 excludes_pattern: Optional[pulumi.Input[str]] = None,
                 includes_pattern: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 package_type: Optional[pulumi.Input[str]] = None,
                 project_environments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project_key: Optional[pulumi.Input[str]] = None,
                 repo_layout_ref: Optional[pulumi.Input[str]] = None,
                 repositories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retrieval_cache_period_seconds: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering VirtualCondaRepository resources.
        :param pulumi.Input[bool] artifactory_requests_can_retrieve_remote_artifacts: Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
               another Artifactory instance.
        :param pulumi.Input[str] default_deployment_repo: Default repository to deploy artifacts.
        :param pulumi.Input[str] description: Public description.
        :param pulumi.Input[str] excludes_pattern: List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
               artifacts are excluded.
        :param pulumi.Input[str] includes_pattern: List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
               used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
        :param pulumi.Input[str] key: A mandatory identifier for the repository that must be unique. It cannot begin with a number or
               contain spaces or special characters.
        :param pulumi.Input[str] notes: Internal description.
        :param pulumi.Input[str] project_key: Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
               assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
        :param pulumi.Input[str] repo_layout_ref: Repository layout key for the virtual repository
        :param pulumi.Input[Sequence[pulumi.Input[str]]] repositories: The effective list of actual repositories included in this virtual repository.
        :param pulumi.Input[int] retrieval_cache_period_seconds: This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
        """
        if artifactory_requests_can_retrieve_remote_artifacts is not None:
            pulumi.set(__self__, "artifactory_requests_can_retrieve_remote_artifacts", artifactory_requests_can_retrieve_remote_artifacts)
        if default_deployment_repo is not None:
            pulumi.set(__self__, "default_deployment_repo", default_deployment_repo)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if excludes_pattern is not None:
            pulumi.set(__self__, "excludes_pattern", excludes_pattern)
        if includes_pattern is not None:
            pulumi.set(__self__, "includes_pattern", includes_pattern)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if notes is not None:
            pulumi.set(__self__, "notes", notes)
        if package_type is not None:
            pulumi.set(__self__, "package_type", package_type)
        if project_environments is not None:
            pulumi.set(__self__, "project_environments", project_environments)
        if project_key is not None:
            pulumi.set(__self__, "project_key", project_key)
        if repo_layout_ref is not None:
            pulumi.set(__self__, "repo_layout_ref", repo_layout_ref)
        if repositories is not None:
            pulumi.set(__self__, "repositories", repositories)
        if retrieval_cache_period_seconds is not None:
            pulumi.set(__self__, "retrieval_cache_period_seconds", retrieval_cache_period_seconds)

    @property
    @pulumi.getter(name="artifactoryRequestsCanRetrieveRemoteArtifacts")
    def artifactory_requests_can_retrieve_remote_artifacts(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
        another Artifactory instance.
        """
        return pulumi.get(self, "artifactory_requests_can_retrieve_remote_artifacts")

    @artifactory_requests_can_retrieve_remote_artifacts.setter
    def artifactory_requests_can_retrieve_remote_artifacts(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "artifactory_requests_can_retrieve_remote_artifacts", value)

    @property
    @pulumi.getter(name="defaultDeploymentRepo")
    def default_deployment_repo(self) -> Optional[pulumi.Input[str]]:
        """
        Default repository to deploy artifacts.
        """
        return pulumi.get(self, "default_deployment_repo")

    @default_deployment_repo.setter
    def default_deployment_repo(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_deployment_repo", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Public description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="excludesPattern")
    def excludes_pattern(self) -> Optional[pulumi.Input[str]]:
        """
        List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
        artifacts are excluded.
        """
        return pulumi.get(self, "excludes_pattern")

    @excludes_pattern.setter
    def excludes_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "excludes_pattern", value)

    @property
    @pulumi.getter(name="includesPattern")
    def includes_pattern(self) -> Optional[pulumi.Input[str]]:
        """
        List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
        used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
        """
        return pulumi.get(self, "includes_pattern")

    @includes_pattern.setter
    def includes_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "includes_pattern", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        A mandatory identifier for the repository that must be unique. It cannot begin with a number or
        contain spaces or special characters.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def notes(self) -> Optional[pulumi.Input[str]]:
        """
        Internal description.
        """
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
    @pulumi.getter(name="projectEnvironments")
    def project_environments(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "project_environments")

    @project_environments.setter
    def project_environments(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "project_environments", value)

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> Optional[pulumi.Input[str]]:
        """
        Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
        assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
        """
        return pulumi.get(self, "project_key")

    @project_key.setter
    def project_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_key", value)

    @property
    @pulumi.getter(name="repoLayoutRef")
    def repo_layout_ref(self) -> Optional[pulumi.Input[str]]:
        """
        Repository layout key for the virtual repository
        """
        return pulumi.get(self, "repo_layout_ref")

    @repo_layout_ref.setter
    def repo_layout_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repo_layout_ref", value)

    @property
    @pulumi.getter
    def repositories(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The effective list of actual repositories included in this virtual repository.
        """
        return pulumi.get(self, "repositories")

    @repositories.setter
    def repositories(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "repositories", value)

    @property
    @pulumi.getter(name="retrievalCachePeriodSeconds")
    def retrieval_cache_period_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
        """
        return pulumi.get(self, "retrieval_cache_period_seconds")

    @retrieval_cache_period_seconds.setter
    def retrieval_cache_period_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retrieval_cache_period_seconds", value)


class VirtualCondaRepository(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 artifactory_requests_can_retrieve_remote_artifacts: Optional[pulumi.Input[bool]] = None,
                 default_deployment_repo: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 excludes_pattern: Optional[pulumi.Input[str]] = None,
                 includes_pattern: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 project_environments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project_key: Optional[pulumi.Input[str]] = None,
                 repo_layout_ref: Optional[pulumi.Input[str]] = None,
                 repositories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retrieval_cache_period_seconds: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Creates a virtual Conda repository.
        Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Conda+Repositories#CondaRepositories-VirtualRepositories).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        foo_conda = artifactory.VirtualCondaRepository("foo-conda",
            key="foo-conda",
            repositories=[],
            description="A test virtual repo",
            notes="Internal description",
            includes_pattern="com/jfrog/**,cloud/jfrog/**",
            excludes_pattern="com/google/**")
        ```

        ## Import

        Virtual repositories can be imported using their name, e.g.

        ```sh
        $ pulumi import artifactory:index/virtualCondaRepository:VirtualCondaRepository foo-conda foo-conda
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] artifactory_requests_can_retrieve_remote_artifacts: Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
               another Artifactory instance.
        :param pulumi.Input[str] default_deployment_repo: Default repository to deploy artifacts.
        :param pulumi.Input[str] description: Public description.
        :param pulumi.Input[str] excludes_pattern: List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
               artifacts are excluded.
        :param pulumi.Input[str] includes_pattern: List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
               used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
        :param pulumi.Input[str] key: A mandatory identifier for the repository that must be unique. It cannot begin with a number or
               contain spaces or special characters.
        :param pulumi.Input[str] notes: Internal description.
        :param pulumi.Input[str] project_key: Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
               assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
        :param pulumi.Input[str] repo_layout_ref: Repository layout key for the virtual repository
        :param pulumi.Input[Sequence[pulumi.Input[str]]] repositories: The effective list of actual repositories included in this virtual repository.
        :param pulumi.Input[int] retrieval_cache_period_seconds: This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VirtualCondaRepositoryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a virtual Conda repository.
        Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Conda+Repositories#CondaRepositories-VirtualRepositories).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        foo_conda = artifactory.VirtualCondaRepository("foo-conda",
            key="foo-conda",
            repositories=[],
            description="A test virtual repo",
            notes="Internal description",
            includes_pattern="com/jfrog/**,cloud/jfrog/**",
            excludes_pattern="com/google/**")
        ```

        ## Import

        Virtual repositories can be imported using their name, e.g.

        ```sh
        $ pulumi import artifactory:index/virtualCondaRepository:VirtualCondaRepository foo-conda foo-conda
        ```

        :param str resource_name: The name of the resource.
        :param VirtualCondaRepositoryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VirtualCondaRepositoryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 artifactory_requests_can_retrieve_remote_artifacts: Optional[pulumi.Input[bool]] = None,
                 default_deployment_repo: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 excludes_pattern: Optional[pulumi.Input[str]] = None,
                 includes_pattern: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 project_environments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project_key: Optional[pulumi.Input[str]] = None,
                 repo_layout_ref: Optional[pulumi.Input[str]] = None,
                 repositories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retrieval_cache_period_seconds: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VirtualCondaRepositoryArgs.__new__(VirtualCondaRepositoryArgs)

            __props__.__dict__["artifactory_requests_can_retrieve_remote_artifacts"] = artifactory_requests_can_retrieve_remote_artifacts
            __props__.__dict__["default_deployment_repo"] = default_deployment_repo
            __props__.__dict__["description"] = description
            __props__.__dict__["excludes_pattern"] = excludes_pattern
            __props__.__dict__["includes_pattern"] = includes_pattern
            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            __props__.__dict__["notes"] = notes
            __props__.__dict__["project_environments"] = project_environments
            __props__.__dict__["project_key"] = project_key
            __props__.__dict__["repo_layout_ref"] = repo_layout_ref
            __props__.__dict__["repositories"] = repositories
            __props__.__dict__["retrieval_cache_period_seconds"] = retrieval_cache_period_seconds
            __props__.__dict__["package_type"] = None
        super(VirtualCondaRepository, __self__).__init__(
            'artifactory:index/virtualCondaRepository:VirtualCondaRepository',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            artifactory_requests_can_retrieve_remote_artifacts: Optional[pulumi.Input[bool]] = None,
            default_deployment_repo: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            excludes_pattern: Optional[pulumi.Input[str]] = None,
            includes_pattern: Optional[pulumi.Input[str]] = None,
            key: Optional[pulumi.Input[str]] = None,
            notes: Optional[pulumi.Input[str]] = None,
            package_type: Optional[pulumi.Input[str]] = None,
            project_environments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            project_key: Optional[pulumi.Input[str]] = None,
            repo_layout_ref: Optional[pulumi.Input[str]] = None,
            repositories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            retrieval_cache_period_seconds: Optional[pulumi.Input[int]] = None) -> 'VirtualCondaRepository':
        """
        Get an existing VirtualCondaRepository resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] artifactory_requests_can_retrieve_remote_artifacts: Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
               another Artifactory instance.
        :param pulumi.Input[str] default_deployment_repo: Default repository to deploy artifacts.
        :param pulumi.Input[str] description: Public description.
        :param pulumi.Input[str] excludes_pattern: List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
               artifacts are excluded.
        :param pulumi.Input[str] includes_pattern: List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
               used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
        :param pulumi.Input[str] key: A mandatory identifier for the repository that must be unique. It cannot begin with a number or
               contain spaces or special characters.
        :param pulumi.Input[str] notes: Internal description.
        :param pulumi.Input[str] project_key: Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
               assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
        :param pulumi.Input[str] repo_layout_ref: Repository layout key for the virtual repository
        :param pulumi.Input[Sequence[pulumi.Input[str]]] repositories: The effective list of actual repositories included in this virtual repository.
        :param pulumi.Input[int] retrieval_cache_period_seconds: This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VirtualCondaRepositoryState.__new__(_VirtualCondaRepositoryState)

        __props__.__dict__["artifactory_requests_can_retrieve_remote_artifacts"] = artifactory_requests_can_retrieve_remote_artifacts
        __props__.__dict__["default_deployment_repo"] = default_deployment_repo
        __props__.__dict__["description"] = description
        __props__.__dict__["excludes_pattern"] = excludes_pattern
        __props__.__dict__["includes_pattern"] = includes_pattern
        __props__.__dict__["key"] = key
        __props__.__dict__["notes"] = notes
        __props__.__dict__["package_type"] = package_type
        __props__.__dict__["project_environments"] = project_environments
        __props__.__dict__["project_key"] = project_key
        __props__.__dict__["repo_layout_ref"] = repo_layout_ref
        __props__.__dict__["repositories"] = repositories
        __props__.__dict__["retrieval_cache_period_seconds"] = retrieval_cache_period_seconds
        return VirtualCondaRepository(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="artifactoryRequestsCanRetrieveRemoteArtifacts")
    def artifactory_requests_can_retrieve_remote_artifacts(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
        another Artifactory instance.
        """
        return pulumi.get(self, "artifactory_requests_can_retrieve_remote_artifacts")

    @property
    @pulumi.getter(name="defaultDeploymentRepo")
    def default_deployment_repo(self) -> pulumi.Output[Optional[str]]:
        """
        Default repository to deploy artifacts.
        """
        return pulumi.get(self, "default_deployment_repo")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Public description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="excludesPattern")
    def excludes_pattern(self) -> pulumi.Output[Optional[str]]:
        """
        List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
        artifacts are excluded.
        """
        return pulumi.get(self, "excludes_pattern")

    @property
    @pulumi.getter(name="includesPattern")
    def includes_pattern(self) -> pulumi.Output[Optional[str]]:
        """
        List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
        used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
        """
        return pulumi.get(self, "includes_pattern")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        A mandatory identifier for the repository that must be unique. It cannot begin with a number or
        contain spaces or special characters.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def notes(self) -> pulumi.Output[Optional[str]]:
        """
        Internal description.
        """
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter(name="packageType")
    def package_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "package_type")

    @property
    @pulumi.getter(name="projectEnvironments")
    def project_environments(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "project_environments")

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> pulumi.Output[Optional[str]]:
        """
        Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
        assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
        """
        return pulumi.get(self, "project_key")

    @property
    @pulumi.getter(name="repoLayoutRef")
    def repo_layout_ref(self) -> pulumi.Output[Optional[str]]:
        """
        Repository layout key for the virtual repository
        """
        return pulumi.get(self, "repo_layout_ref")

    @property
    @pulumi.getter
    def repositories(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The effective list of actual repositories included in this virtual repository.
        """
        return pulumi.get(self, "repositories")

    @property
    @pulumi.getter(name="retrievalCachePeriodSeconds")
    def retrieval_cache_period_seconds(self) -> pulumi.Output[Optional[int]]:
        """
        This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
        """
        return pulumi.get(self, "retrieval_cache_period_seconds")

