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

__all__ = [
    'GetRemoteRpmRepositoryResult',
    'AwaitableGetRemoteRpmRepositoryResult',
    'get_remote_rpm_repository',
    'get_remote_rpm_repository_output',
]

@pulumi.output_type
class GetRemoteRpmRepositoryResult:
    """
    A collection of values returned by getRemoteRpmRepository.
    """
    def __init__(__self__, allow_any_host_auth=None, assumed_offline_period_secs=None, blacked_out=None, block_mismatching_mime_types=None, bypass_head_requests=None, cdn_redirect=None, client_tls_certificate=None, content_synchronisation=None, description=None, disable_proxy=None, download_direct=None, enable_cookie_management=None, excludes_pattern=None, hard_fail=None, id=None, includes_pattern=None, key=None, list_remote_folder_items=None, local_address=None, metadata_retrieval_timeout_secs=None, mismatching_mime_types_override_list=None, missed_cache_period_seconds=None, notes=None, offline=None, package_type=None, password=None, priority_resolution=None, project_environments=None, project_key=None, property_sets=None, proxy=None, query_params=None, remote_repo_layout_ref=None, repo_layout_ref=None, retrieval_cache_period_seconds=None, share_configuration=None, socket_timeout_millis=None, store_artifacts_locally=None, synchronize_properties=None, unused_artifacts_cleanup_period_hours=None, url=None, username=None, xray_index=None):
        if allow_any_host_auth and not isinstance(allow_any_host_auth, bool):
            raise TypeError("Expected argument 'allow_any_host_auth' to be a bool")
        pulumi.set(__self__, "allow_any_host_auth", allow_any_host_auth)
        if assumed_offline_period_secs and not isinstance(assumed_offline_period_secs, int):
            raise TypeError("Expected argument 'assumed_offline_period_secs' to be a int")
        pulumi.set(__self__, "assumed_offline_period_secs", assumed_offline_period_secs)
        if blacked_out and not isinstance(blacked_out, bool):
            raise TypeError("Expected argument 'blacked_out' to be a bool")
        pulumi.set(__self__, "blacked_out", blacked_out)
        if block_mismatching_mime_types and not isinstance(block_mismatching_mime_types, bool):
            raise TypeError("Expected argument 'block_mismatching_mime_types' to be a bool")
        pulumi.set(__self__, "block_mismatching_mime_types", block_mismatching_mime_types)
        if bypass_head_requests and not isinstance(bypass_head_requests, bool):
            raise TypeError("Expected argument 'bypass_head_requests' to be a bool")
        pulumi.set(__self__, "bypass_head_requests", bypass_head_requests)
        if cdn_redirect and not isinstance(cdn_redirect, bool):
            raise TypeError("Expected argument 'cdn_redirect' to be a bool")
        pulumi.set(__self__, "cdn_redirect", cdn_redirect)
        if client_tls_certificate and not isinstance(client_tls_certificate, str):
            raise TypeError("Expected argument 'client_tls_certificate' to be a str")
        pulumi.set(__self__, "client_tls_certificate", client_tls_certificate)
        if content_synchronisation and not isinstance(content_synchronisation, dict):
            raise TypeError("Expected argument 'content_synchronisation' to be a dict")
        pulumi.set(__self__, "content_synchronisation", content_synchronisation)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disable_proxy and not isinstance(disable_proxy, bool):
            raise TypeError("Expected argument 'disable_proxy' to be a bool")
        pulumi.set(__self__, "disable_proxy", disable_proxy)
        if download_direct and not isinstance(download_direct, bool):
            raise TypeError("Expected argument 'download_direct' to be a bool")
        pulumi.set(__self__, "download_direct", download_direct)
        if enable_cookie_management and not isinstance(enable_cookie_management, bool):
            raise TypeError("Expected argument 'enable_cookie_management' to be a bool")
        pulumi.set(__self__, "enable_cookie_management", enable_cookie_management)
        if excludes_pattern and not isinstance(excludes_pattern, str):
            raise TypeError("Expected argument 'excludes_pattern' to be a str")
        pulumi.set(__self__, "excludes_pattern", excludes_pattern)
        if hard_fail and not isinstance(hard_fail, bool):
            raise TypeError("Expected argument 'hard_fail' to be a bool")
        pulumi.set(__self__, "hard_fail", hard_fail)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if includes_pattern and not isinstance(includes_pattern, str):
            raise TypeError("Expected argument 'includes_pattern' to be a str")
        pulumi.set(__self__, "includes_pattern", includes_pattern)
        if key and not isinstance(key, str):
            raise TypeError("Expected argument 'key' to be a str")
        pulumi.set(__self__, "key", key)
        if list_remote_folder_items and not isinstance(list_remote_folder_items, bool):
            raise TypeError("Expected argument 'list_remote_folder_items' to be a bool")
        pulumi.set(__self__, "list_remote_folder_items", list_remote_folder_items)
        if local_address and not isinstance(local_address, str):
            raise TypeError("Expected argument 'local_address' to be a str")
        pulumi.set(__self__, "local_address", local_address)
        if metadata_retrieval_timeout_secs and not isinstance(metadata_retrieval_timeout_secs, int):
            raise TypeError("Expected argument 'metadata_retrieval_timeout_secs' to be a int")
        pulumi.set(__self__, "metadata_retrieval_timeout_secs", metadata_retrieval_timeout_secs)
        if mismatching_mime_types_override_list and not isinstance(mismatching_mime_types_override_list, str):
            raise TypeError("Expected argument 'mismatching_mime_types_override_list' to be a str")
        pulumi.set(__self__, "mismatching_mime_types_override_list", mismatching_mime_types_override_list)
        if missed_cache_period_seconds and not isinstance(missed_cache_period_seconds, int):
            raise TypeError("Expected argument 'missed_cache_period_seconds' to be a int")
        pulumi.set(__self__, "missed_cache_period_seconds", missed_cache_period_seconds)
        if notes and not isinstance(notes, str):
            raise TypeError("Expected argument 'notes' to be a str")
        pulumi.set(__self__, "notes", notes)
        if offline and not isinstance(offline, bool):
            raise TypeError("Expected argument 'offline' to be a bool")
        pulumi.set(__self__, "offline", offline)
        if package_type and not isinstance(package_type, str):
            raise TypeError("Expected argument 'package_type' to be a str")
        pulumi.set(__self__, "package_type", package_type)
        if password and not isinstance(password, str):
            raise TypeError("Expected argument 'password' to be a str")
        pulumi.set(__self__, "password", password)
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
        if proxy and not isinstance(proxy, str):
            raise TypeError("Expected argument 'proxy' to be a str")
        pulumi.set(__self__, "proxy", proxy)
        if query_params and not isinstance(query_params, str):
            raise TypeError("Expected argument 'query_params' to be a str")
        pulumi.set(__self__, "query_params", query_params)
        if remote_repo_layout_ref and not isinstance(remote_repo_layout_ref, str):
            raise TypeError("Expected argument 'remote_repo_layout_ref' to be a str")
        pulumi.set(__self__, "remote_repo_layout_ref", remote_repo_layout_ref)
        if repo_layout_ref and not isinstance(repo_layout_ref, str):
            raise TypeError("Expected argument 'repo_layout_ref' to be a str")
        pulumi.set(__self__, "repo_layout_ref", repo_layout_ref)
        if retrieval_cache_period_seconds and not isinstance(retrieval_cache_period_seconds, int):
            raise TypeError("Expected argument 'retrieval_cache_period_seconds' to be a int")
        pulumi.set(__self__, "retrieval_cache_period_seconds", retrieval_cache_period_seconds)
        if share_configuration and not isinstance(share_configuration, bool):
            raise TypeError("Expected argument 'share_configuration' to be a bool")
        pulumi.set(__self__, "share_configuration", share_configuration)
        if socket_timeout_millis and not isinstance(socket_timeout_millis, int):
            raise TypeError("Expected argument 'socket_timeout_millis' to be a int")
        pulumi.set(__self__, "socket_timeout_millis", socket_timeout_millis)
        if store_artifacts_locally and not isinstance(store_artifacts_locally, bool):
            raise TypeError("Expected argument 'store_artifacts_locally' to be a bool")
        pulumi.set(__self__, "store_artifacts_locally", store_artifacts_locally)
        if synchronize_properties and not isinstance(synchronize_properties, bool):
            raise TypeError("Expected argument 'synchronize_properties' to be a bool")
        pulumi.set(__self__, "synchronize_properties", synchronize_properties)
        if unused_artifacts_cleanup_period_hours and not isinstance(unused_artifacts_cleanup_period_hours, int):
            raise TypeError("Expected argument 'unused_artifacts_cleanup_period_hours' to be a int")
        pulumi.set(__self__, "unused_artifacts_cleanup_period_hours", unused_artifacts_cleanup_period_hours)
        if url and not isinstance(url, str):
            raise TypeError("Expected argument 'url' to be a str")
        pulumi.set(__self__, "url", url)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)
        if xray_index and not isinstance(xray_index, bool):
            raise TypeError("Expected argument 'xray_index' to be a bool")
        pulumi.set(__self__, "xray_index", xray_index)

    @property
    @pulumi.getter(name="allowAnyHostAuth")
    def allow_any_host_auth(self) -> Optional[bool]:
        return pulumi.get(self, "allow_any_host_auth")

    @property
    @pulumi.getter(name="assumedOfflinePeriodSecs")
    def assumed_offline_period_secs(self) -> Optional[int]:
        return pulumi.get(self, "assumed_offline_period_secs")

    @property
    @pulumi.getter(name="blackedOut")
    def blacked_out(self) -> Optional[bool]:
        return pulumi.get(self, "blacked_out")

    @property
    @pulumi.getter(name="blockMismatchingMimeTypes")
    def block_mismatching_mime_types(self) -> Optional[bool]:
        return pulumi.get(self, "block_mismatching_mime_types")

    @property
    @pulumi.getter(name="bypassHeadRequests")
    def bypass_head_requests(self) -> Optional[bool]:
        return pulumi.get(self, "bypass_head_requests")

    @property
    @pulumi.getter(name="cdnRedirect")
    def cdn_redirect(self) -> Optional[bool]:
        return pulumi.get(self, "cdn_redirect")

    @property
    @pulumi.getter(name="clientTlsCertificate")
    def client_tls_certificate(self) -> str:
        return pulumi.get(self, "client_tls_certificate")

    @property
    @pulumi.getter(name="contentSynchronisation")
    def content_synchronisation(self) -> 'outputs.GetRemoteRpmRepositoryContentSynchronisationResult':
        return pulumi.get(self, "content_synchronisation")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="disableProxy")
    def disable_proxy(self) -> Optional[bool]:
        return pulumi.get(self, "disable_proxy")

    @property
    @pulumi.getter(name="downloadDirect")
    def download_direct(self) -> Optional[bool]:
        return pulumi.get(self, "download_direct")

    @property
    @pulumi.getter(name="enableCookieManagement")
    def enable_cookie_management(self) -> Optional[bool]:
        return pulumi.get(self, "enable_cookie_management")

    @property
    @pulumi.getter(name="excludesPattern")
    def excludes_pattern(self) -> Optional[str]:
        return pulumi.get(self, "excludes_pattern")

    @property
    @pulumi.getter(name="hardFail")
    def hard_fail(self) -> Optional[bool]:
        return pulumi.get(self, "hard_fail")

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
    @pulumi.getter(name="listRemoteFolderItems")
    def list_remote_folder_items(self) -> Optional[bool]:
        return pulumi.get(self, "list_remote_folder_items")

    @property
    @pulumi.getter(name="localAddress")
    def local_address(self) -> Optional[str]:
        return pulumi.get(self, "local_address")

    @property
    @pulumi.getter(name="metadataRetrievalTimeoutSecs")
    def metadata_retrieval_timeout_secs(self) -> Optional[int]:
        return pulumi.get(self, "metadata_retrieval_timeout_secs")

    @property
    @pulumi.getter(name="mismatchingMimeTypesOverrideList")
    def mismatching_mime_types_override_list(self) -> Optional[str]:
        return pulumi.get(self, "mismatching_mime_types_override_list")

    @property
    @pulumi.getter(name="missedCachePeriodSeconds")
    def missed_cache_period_seconds(self) -> Optional[int]:
        return pulumi.get(self, "missed_cache_period_seconds")

    @property
    @pulumi.getter
    def notes(self) -> Optional[str]:
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter
    def offline(self) -> Optional[bool]:
        return pulumi.get(self, "offline")

    @property
    @pulumi.getter(name="packageType")
    def package_type(self) -> str:
        return pulumi.get(self, "package_type")

    @property
    @pulumi.getter
    def password(self) -> Optional[str]:
        return pulumi.get(self, "password")

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
    @pulumi.getter
    def proxy(self) -> Optional[str]:
        return pulumi.get(self, "proxy")

    @property
    @pulumi.getter(name="queryParams")
    def query_params(self) -> Optional[str]:
        return pulumi.get(self, "query_params")

    @property
    @pulumi.getter(name="remoteRepoLayoutRef")
    def remote_repo_layout_ref(self) -> Optional[str]:
        return pulumi.get(self, "remote_repo_layout_ref")

    @property
    @pulumi.getter(name="repoLayoutRef")
    def repo_layout_ref(self) -> Optional[str]:
        return pulumi.get(self, "repo_layout_ref")

    @property
    @pulumi.getter(name="retrievalCachePeriodSeconds")
    def retrieval_cache_period_seconds(self) -> Optional[int]:
        return pulumi.get(self, "retrieval_cache_period_seconds")

    @property
    @pulumi.getter(name="shareConfiguration")
    def share_configuration(self) -> bool:
        return pulumi.get(self, "share_configuration")

    @property
    @pulumi.getter(name="socketTimeoutMillis")
    def socket_timeout_millis(self) -> Optional[int]:
        return pulumi.get(self, "socket_timeout_millis")

    @property
    @pulumi.getter(name="storeArtifactsLocally")
    def store_artifacts_locally(self) -> Optional[bool]:
        return pulumi.get(self, "store_artifacts_locally")

    @property
    @pulumi.getter(name="synchronizeProperties")
    def synchronize_properties(self) -> Optional[bool]:
        return pulumi.get(self, "synchronize_properties")

    @property
    @pulumi.getter(name="unusedArtifactsCleanupPeriodHours")
    def unused_artifacts_cleanup_period_hours(self) -> Optional[int]:
        return pulumi.get(self, "unused_artifacts_cleanup_period_hours")

    @property
    @pulumi.getter
    def url(self) -> Optional[str]:
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def username(self) -> Optional[str]:
        return pulumi.get(self, "username")

    @property
    @pulumi.getter(name="xrayIndex")
    def xray_index(self) -> Optional[bool]:
        return pulumi.get(self, "xray_index")


class AwaitableGetRemoteRpmRepositoryResult(GetRemoteRpmRepositoryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRemoteRpmRepositoryResult(
            allow_any_host_auth=self.allow_any_host_auth,
            assumed_offline_period_secs=self.assumed_offline_period_secs,
            blacked_out=self.blacked_out,
            block_mismatching_mime_types=self.block_mismatching_mime_types,
            bypass_head_requests=self.bypass_head_requests,
            cdn_redirect=self.cdn_redirect,
            client_tls_certificate=self.client_tls_certificate,
            content_synchronisation=self.content_synchronisation,
            description=self.description,
            disable_proxy=self.disable_proxy,
            download_direct=self.download_direct,
            enable_cookie_management=self.enable_cookie_management,
            excludes_pattern=self.excludes_pattern,
            hard_fail=self.hard_fail,
            id=self.id,
            includes_pattern=self.includes_pattern,
            key=self.key,
            list_remote_folder_items=self.list_remote_folder_items,
            local_address=self.local_address,
            metadata_retrieval_timeout_secs=self.metadata_retrieval_timeout_secs,
            mismatching_mime_types_override_list=self.mismatching_mime_types_override_list,
            missed_cache_period_seconds=self.missed_cache_period_seconds,
            notes=self.notes,
            offline=self.offline,
            package_type=self.package_type,
            password=self.password,
            priority_resolution=self.priority_resolution,
            project_environments=self.project_environments,
            project_key=self.project_key,
            property_sets=self.property_sets,
            proxy=self.proxy,
            query_params=self.query_params,
            remote_repo_layout_ref=self.remote_repo_layout_ref,
            repo_layout_ref=self.repo_layout_ref,
            retrieval_cache_period_seconds=self.retrieval_cache_period_seconds,
            share_configuration=self.share_configuration,
            socket_timeout_millis=self.socket_timeout_millis,
            store_artifacts_locally=self.store_artifacts_locally,
            synchronize_properties=self.synchronize_properties,
            unused_artifacts_cleanup_period_hours=self.unused_artifacts_cleanup_period_hours,
            url=self.url,
            username=self.username,
            xray_index=self.xray_index)


def get_remote_rpm_repository(allow_any_host_auth: Optional[bool] = None,
                              assumed_offline_period_secs: Optional[int] = None,
                              blacked_out: Optional[bool] = None,
                              block_mismatching_mime_types: Optional[bool] = None,
                              bypass_head_requests: Optional[bool] = None,
                              cdn_redirect: Optional[bool] = None,
                              client_tls_certificate: Optional[str] = None,
                              content_synchronisation: Optional[pulumi.InputType['GetRemoteRpmRepositoryContentSynchronisationArgs']] = None,
                              description: Optional[str] = None,
                              disable_proxy: Optional[bool] = None,
                              download_direct: Optional[bool] = None,
                              enable_cookie_management: Optional[bool] = None,
                              excludes_pattern: Optional[str] = None,
                              hard_fail: Optional[bool] = None,
                              includes_pattern: Optional[str] = None,
                              key: Optional[str] = None,
                              list_remote_folder_items: Optional[bool] = None,
                              local_address: Optional[str] = None,
                              metadata_retrieval_timeout_secs: Optional[int] = None,
                              mismatching_mime_types_override_list: Optional[str] = None,
                              missed_cache_period_seconds: Optional[int] = None,
                              notes: Optional[str] = None,
                              offline: Optional[bool] = None,
                              password: Optional[str] = None,
                              priority_resolution: Optional[bool] = None,
                              project_environments: Optional[Sequence[str]] = None,
                              project_key: Optional[str] = None,
                              property_sets: Optional[Sequence[str]] = None,
                              proxy: Optional[str] = None,
                              query_params: Optional[str] = None,
                              remote_repo_layout_ref: Optional[str] = None,
                              repo_layout_ref: Optional[str] = None,
                              retrieval_cache_period_seconds: Optional[int] = None,
                              share_configuration: Optional[bool] = None,
                              socket_timeout_millis: Optional[int] = None,
                              store_artifacts_locally: Optional[bool] = None,
                              synchronize_properties: Optional[bool] = None,
                              unused_artifacts_cleanup_period_hours: Optional[int] = None,
                              url: Optional[str] = None,
                              username: Optional[str] = None,
                              xray_index: Optional[bool] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRemoteRpmRepositoryResult:
    """
    Retrieves a remote Rpm repository.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_artifactory as artifactory

    remote_rpm = artifactory.get_remote_rpm_repository(key="remote-rpm")
    ```


    :param str key: the identity key of the repo.
    """
    __args__ = dict()
    __args__['allowAnyHostAuth'] = allow_any_host_auth
    __args__['assumedOfflinePeriodSecs'] = assumed_offline_period_secs
    __args__['blackedOut'] = blacked_out
    __args__['blockMismatchingMimeTypes'] = block_mismatching_mime_types
    __args__['bypassHeadRequests'] = bypass_head_requests
    __args__['cdnRedirect'] = cdn_redirect
    __args__['clientTlsCertificate'] = client_tls_certificate
    __args__['contentSynchronisation'] = content_synchronisation
    __args__['description'] = description
    __args__['disableProxy'] = disable_proxy
    __args__['downloadDirect'] = download_direct
    __args__['enableCookieManagement'] = enable_cookie_management
    __args__['excludesPattern'] = excludes_pattern
    __args__['hardFail'] = hard_fail
    __args__['includesPattern'] = includes_pattern
    __args__['key'] = key
    __args__['listRemoteFolderItems'] = list_remote_folder_items
    __args__['localAddress'] = local_address
    __args__['metadataRetrievalTimeoutSecs'] = metadata_retrieval_timeout_secs
    __args__['mismatchingMimeTypesOverrideList'] = mismatching_mime_types_override_list
    __args__['missedCachePeriodSeconds'] = missed_cache_period_seconds
    __args__['notes'] = notes
    __args__['offline'] = offline
    __args__['password'] = password
    __args__['priorityResolution'] = priority_resolution
    __args__['projectEnvironments'] = project_environments
    __args__['projectKey'] = project_key
    __args__['propertySets'] = property_sets
    __args__['proxy'] = proxy
    __args__['queryParams'] = query_params
    __args__['remoteRepoLayoutRef'] = remote_repo_layout_ref
    __args__['repoLayoutRef'] = repo_layout_ref
    __args__['retrievalCachePeriodSeconds'] = retrieval_cache_period_seconds
    __args__['shareConfiguration'] = share_configuration
    __args__['socketTimeoutMillis'] = socket_timeout_millis
    __args__['storeArtifactsLocally'] = store_artifacts_locally
    __args__['synchronizeProperties'] = synchronize_properties
    __args__['unusedArtifactsCleanupPeriodHours'] = unused_artifacts_cleanup_period_hours
    __args__['url'] = url
    __args__['username'] = username
    __args__['xrayIndex'] = xray_index
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('artifactory:index/getRemoteRpmRepository:getRemoteRpmRepository', __args__, opts=opts, typ=GetRemoteRpmRepositoryResult).value

    return AwaitableGetRemoteRpmRepositoryResult(
        allow_any_host_auth=pulumi.get(__ret__, 'allow_any_host_auth'),
        assumed_offline_period_secs=pulumi.get(__ret__, 'assumed_offline_period_secs'),
        blacked_out=pulumi.get(__ret__, 'blacked_out'),
        block_mismatching_mime_types=pulumi.get(__ret__, 'block_mismatching_mime_types'),
        bypass_head_requests=pulumi.get(__ret__, 'bypass_head_requests'),
        cdn_redirect=pulumi.get(__ret__, 'cdn_redirect'),
        client_tls_certificate=pulumi.get(__ret__, 'client_tls_certificate'),
        content_synchronisation=pulumi.get(__ret__, 'content_synchronisation'),
        description=pulumi.get(__ret__, 'description'),
        disable_proxy=pulumi.get(__ret__, 'disable_proxy'),
        download_direct=pulumi.get(__ret__, 'download_direct'),
        enable_cookie_management=pulumi.get(__ret__, 'enable_cookie_management'),
        excludes_pattern=pulumi.get(__ret__, 'excludes_pattern'),
        hard_fail=pulumi.get(__ret__, 'hard_fail'),
        id=pulumi.get(__ret__, 'id'),
        includes_pattern=pulumi.get(__ret__, 'includes_pattern'),
        key=pulumi.get(__ret__, 'key'),
        list_remote_folder_items=pulumi.get(__ret__, 'list_remote_folder_items'),
        local_address=pulumi.get(__ret__, 'local_address'),
        metadata_retrieval_timeout_secs=pulumi.get(__ret__, 'metadata_retrieval_timeout_secs'),
        mismatching_mime_types_override_list=pulumi.get(__ret__, 'mismatching_mime_types_override_list'),
        missed_cache_period_seconds=pulumi.get(__ret__, 'missed_cache_period_seconds'),
        notes=pulumi.get(__ret__, 'notes'),
        offline=pulumi.get(__ret__, 'offline'),
        package_type=pulumi.get(__ret__, 'package_type'),
        password=pulumi.get(__ret__, 'password'),
        priority_resolution=pulumi.get(__ret__, 'priority_resolution'),
        project_environments=pulumi.get(__ret__, 'project_environments'),
        project_key=pulumi.get(__ret__, 'project_key'),
        property_sets=pulumi.get(__ret__, 'property_sets'),
        proxy=pulumi.get(__ret__, 'proxy'),
        query_params=pulumi.get(__ret__, 'query_params'),
        remote_repo_layout_ref=pulumi.get(__ret__, 'remote_repo_layout_ref'),
        repo_layout_ref=pulumi.get(__ret__, 'repo_layout_ref'),
        retrieval_cache_period_seconds=pulumi.get(__ret__, 'retrieval_cache_period_seconds'),
        share_configuration=pulumi.get(__ret__, 'share_configuration'),
        socket_timeout_millis=pulumi.get(__ret__, 'socket_timeout_millis'),
        store_artifacts_locally=pulumi.get(__ret__, 'store_artifacts_locally'),
        synchronize_properties=pulumi.get(__ret__, 'synchronize_properties'),
        unused_artifacts_cleanup_period_hours=pulumi.get(__ret__, 'unused_artifacts_cleanup_period_hours'),
        url=pulumi.get(__ret__, 'url'),
        username=pulumi.get(__ret__, 'username'),
        xray_index=pulumi.get(__ret__, 'xray_index'))


@_utilities.lift_output_func(get_remote_rpm_repository)
def get_remote_rpm_repository_output(allow_any_host_auth: Optional[pulumi.Input[Optional[bool]]] = None,
                                     assumed_offline_period_secs: Optional[pulumi.Input[Optional[int]]] = None,
                                     blacked_out: Optional[pulumi.Input[Optional[bool]]] = None,
                                     block_mismatching_mime_types: Optional[pulumi.Input[Optional[bool]]] = None,
                                     bypass_head_requests: Optional[pulumi.Input[Optional[bool]]] = None,
                                     cdn_redirect: Optional[pulumi.Input[Optional[bool]]] = None,
                                     client_tls_certificate: Optional[pulumi.Input[Optional[str]]] = None,
                                     content_synchronisation: Optional[pulumi.Input[Optional[pulumi.InputType['GetRemoteRpmRepositoryContentSynchronisationArgs']]]] = None,
                                     description: Optional[pulumi.Input[Optional[str]]] = None,
                                     disable_proxy: Optional[pulumi.Input[Optional[bool]]] = None,
                                     download_direct: Optional[pulumi.Input[Optional[bool]]] = None,
                                     enable_cookie_management: Optional[pulumi.Input[Optional[bool]]] = None,
                                     excludes_pattern: Optional[pulumi.Input[Optional[str]]] = None,
                                     hard_fail: Optional[pulumi.Input[Optional[bool]]] = None,
                                     includes_pattern: Optional[pulumi.Input[Optional[str]]] = None,
                                     key: Optional[pulumi.Input[str]] = None,
                                     list_remote_folder_items: Optional[pulumi.Input[Optional[bool]]] = None,
                                     local_address: Optional[pulumi.Input[Optional[str]]] = None,
                                     metadata_retrieval_timeout_secs: Optional[pulumi.Input[Optional[int]]] = None,
                                     mismatching_mime_types_override_list: Optional[pulumi.Input[Optional[str]]] = None,
                                     missed_cache_period_seconds: Optional[pulumi.Input[Optional[int]]] = None,
                                     notes: Optional[pulumi.Input[Optional[str]]] = None,
                                     offline: Optional[pulumi.Input[Optional[bool]]] = None,
                                     password: Optional[pulumi.Input[Optional[str]]] = None,
                                     priority_resolution: Optional[pulumi.Input[Optional[bool]]] = None,
                                     project_environments: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                     project_key: Optional[pulumi.Input[Optional[str]]] = None,
                                     property_sets: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                     proxy: Optional[pulumi.Input[Optional[str]]] = None,
                                     query_params: Optional[pulumi.Input[Optional[str]]] = None,
                                     remote_repo_layout_ref: Optional[pulumi.Input[Optional[str]]] = None,
                                     repo_layout_ref: Optional[pulumi.Input[Optional[str]]] = None,
                                     retrieval_cache_period_seconds: Optional[pulumi.Input[Optional[int]]] = None,
                                     share_configuration: Optional[pulumi.Input[Optional[bool]]] = None,
                                     socket_timeout_millis: Optional[pulumi.Input[Optional[int]]] = None,
                                     store_artifacts_locally: Optional[pulumi.Input[Optional[bool]]] = None,
                                     synchronize_properties: Optional[pulumi.Input[Optional[bool]]] = None,
                                     unused_artifacts_cleanup_period_hours: Optional[pulumi.Input[Optional[int]]] = None,
                                     url: Optional[pulumi.Input[Optional[str]]] = None,
                                     username: Optional[pulumi.Input[Optional[str]]] = None,
                                     xray_index: Optional[pulumi.Input[Optional[bool]]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRemoteRpmRepositoryResult]:
    """
    Retrieves a remote Rpm repository.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_artifactory as artifactory

    remote_rpm = artifactory.get_remote_rpm_repository(key="remote-rpm")
    ```


    :param str key: the identity key of the repo.
    """
    ...
