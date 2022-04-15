// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export interface AccessTokenAdminToken {
    instanceId: string;
}

export interface ArtifactPropertyWebhookCriteria {
    /**
     * Trigger on any local repo
     */
    anyLocal: boolean;
    /**
     * Trigger on any remote repo
     */
    anyRemote: boolean;
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    excludePatterns?: string[];
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    includePatterns?: string[];
    /**
     * Trigger on this list of repo keys
     */
    repoKeys: string[];
}

export interface ArtifactWebhookCriteria {
    /**
     * Trigger on any local repo
     */
    anyLocal: boolean;
    /**
     * Trigger on any remote repo
     */
    anyRemote: boolean;
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    excludePatterns?: string[];
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    includePatterns?: string[];
    /**
     * Trigger on this list of repo keys
     */
    repoKeys: string[];
}

export interface ArtifactoryReleaseBundleWebhookCriteria {
    /**
     * Trigger on any release bundle
     */
    anyReleaseBundle: boolean;
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    excludePatterns?: string[];
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    includePatterns?: string[];
    /**
     * Trigger on this list of release bundle names
     */
    registeredReleaseBundleNames: string[];
}

export interface BuildWebhookCriteria {
    /**
     * Trigger on any build
     */
    anyBuild: boolean;
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    excludePatterns?: string[];
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    includePatterns?: string[];
    /**
     * Trigger on this list of build names
     */
    selectedBuilds: string[];
}

export interface DistributionWebhookCriteria {
    /**
     * Trigger on any release bundle
     */
    anyReleaseBundle: boolean;
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    excludePatterns?: string[];
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    includePatterns?: string[];
    /**
     * Trigger on this list of release bundle names
     */
    registeredReleaseBundleNames: string[];
}

export interface DockerWebhookCriteria {
    /**
     * Trigger on any local repo
     */
    anyLocal: boolean;
    /**
     * Trigger on any remote repo
     */
    anyRemote: boolean;
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    excludePatterns?: string[];
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    includePatterns?: string[];
    /**
     * Trigger on this list of repo keys
     */
    repoKeys: string[];
}

export interface FederatedAlpineRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface FederatedBowerRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface FederatedCargoRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface FederatedChefRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface FederatedCocoapodsRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface FederatedComposerRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface FederatedConanRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface FederatedCondaRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface FederatedCranRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface FederatedDebianRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface FederatedDockerRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface FederatedGemsRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface FederatedGenericRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface FederatedGitltfsRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface FederatedGoRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface FederatedGradleRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface FederatedHelmRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface FederatedIvyRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface FederatedMavenRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface FederatedNpmRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface FederatedNugetRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface FederatedOpkgRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface FederatedPuppetRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface FederatedPypiRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface FederatedRpmRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface FederatedSbtRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface FederatedVagrantRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: boolean;
    /**
     * Full URL to ending with the repository name
     */
    url: string;
}

export interface OauthSettingsOauthProvider {
    /**
     * OAuth user info endpoint for the IdP.
     */
    apiUrl: string;
    /**
     * OAuth authorization endpoint for the IdP.
     */
    authUrl: string;
    /**
     * OAuth client ID configured on the IdP.
     */
    clientId: string;
    /**
     * OAuth client secret configured on the IdP.
     */
    clientSecret: string;
    /**
     * Enable the Artifactory OAuth provider.  Default value is `true`.
     */
    enabled?: boolean;
    /**
     * Name of the Artifactory OAuth provider.
     */
    name: string;
    /**
     * OAuth token endpoint for the IdP.
     */
    tokenUrl: string;
    /**
     * Type of OAuth provider. (e.g., `github`, `google`, `cloudfoundry`, or `openId`)
     */
    type: string;
}

export interface PermissionTargetBuild {
    /**
     * -
     */
    actions?: outputs.PermissionTargetBuildActions;
    /**
     * Pattern of artifacts to exclude
     */
    excludesPatterns?: string[];
    /**
     * Pattern of artifacts to include
     */
    includesPatterns?: string[];
    /**
     * List of repositories this permission target is applicable for
     */
    repositories: string[];
}

export interface PermissionTargetBuildActions {
    /**
     * Groups this permission applies for.
     */
    groups?: outputs.PermissionTargetBuildActionsGroup[];
    /**
     * Users this permission target applies for.
     */
    users?: outputs.PermissionTargetBuildActionsUser[];
}

export interface PermissionTargetBuildActionsGroup {
    /**
     * Name of permission
     */
    name: string;
    permissions: string[];
}

export interface PermissionTargetBuildActionsUser {
    /**
     * Name of permission
     */
    name: string;
    permissions: string[];
}

export interface PermissionTargetReleaseBundle {
    /**
     * -
     */
    actions?: outputs.PermissionTargetReleaseBundleActions;
    /**
     * Pattern of artifacts to exclude
     */
    excludesPatterns?: string[];
    /**
     * Pattern of artifacts to include
     */
    includesPatterns?: string[];
    /**
     * List of repositories this permission target is applicable for
     */
    repositories: string[];
}

export interface PermissionTargetReleaseBundleActions {
    /**
     * Groups this permission applies for.
     */
    groups?: outputs.PermissionTargetReleaseBundleActionsGroup[];
    /**
     * Users this permission target applies for.
     */
    users?: outputs.PermissionTargetReleaseBundleActionsUser[];
}

export interface PermissionTargetReleaseBundleActionsGroup {
    /**
     * Name of permission
     */
    name: string;
    permissions: string[];
}

export interface PermissionTargetReleaseBundleActionsUser {
    /**
     * Name of permission
     */
    name: string;
    permissions: string[];
}

export interface PermissionTargetRepo {
    /**
     * -
     */
    actions?: outputs.PermissionTargetRepoActions;
    /**
     * Pattern of artifacts to exclude
     */
    excludesPatterns?: string[];
    /**
     * Pattern of artifacts to include
     */
    includesPatterns?: string[];
    /**
     * List of repositories this permission target is applicable for
     */
    repositories: string[];
}

export interface PermissionTargetRepoActions {
    /**
     * Groups this permission applies for.
     */
    groups?: outputs.PermissionTargetRepoActionsGroup[];
    /**
     * Users this permission target applies for.
     */
    users?: outputs.PermissionTargetRepoActionsUser[];
}

export interface PermissionTargetRepoActionsGroup {
    /**
     * Name of permission
     */
    name: string;
    permissions: string[];
}

export interface PermissionTargetRepoActionsUser {
    /**
     * Name of permission
     */
    name: string;
    permissions: string[];
}

export interface PermissionTargetsBuild {
    /**
     * -
     */
    actions?: outputs.PermissionTargetsBuildActions;
    /**
     * Pattern of artifacts to exclude
     */
    excludesPatterns?: string[];
    /**
     * Pattern of artifacts to include
     */
    includesPatterns?: string[];
    /**
     * List of repositories this permission target is applicable for
     */
    repositories: string[];
}

export interface PermissionTargetsBuildActions {
    /**
     * Groups this permission applies for.
     */
    groups?: outputs.PermissionTargetsBuildActionsGroup[];
    /**
     * Users this permission target applies for.
     */
    users?: outputs.PermissionTargetsBuildActionsUser[];
}

export interface PermissionTargetsBuildActionsGroup {
    /**
     * Name of permission
     */
    name: string;
    permissions: string[];
}

export interface PermissionTargetsBuildActionsUser {
    /**
     * Name of permission
     */
    name: string;
    permissions: string[];
}

export interface PermissionTargetsReleaseBundle {
    /**
     * -
     */
    actions?: outputs.PermissionTargetsReleaseBundleActions;
    /**
     * Pattern of artifacts to exclude
     */
    excludesPatterns?: string[];
    /**
     * Pattern of artifacts to include
     */
    includesPatterns?: string[];
    /**
     * List of repositories this permission target is applicable for
     */
    repositories: string[];
}

export interface PermissionTargetsReleaseBundleActions {
    /**
     * Groups this permission applies for.
     */
    groups?: outputs.PermissionTargetsReleaseBundleActionsGroup[];
    /**
     * Users this permission target applies for.
     */
    users?: outputs.PermissionTargetsReleaseBundleActionsUser[];
}

export interface PermissionTargetsReleaseBundleActionsGroup {
    /**
     * Name of permission
     */
    name: string;
    permissions: string[];
}

export interface PermissionTargetsReleaseBundleActionsUser {
    /**
     * Name of permission
     */
    name: string;
    permissions: string[];
}

export interface PermissionTargetsRepo {
    /**
     * -
     */
    actions?: outputs.PermissionTargetsRepoActions;
    /**
     * Pattern of artifacts to exclude
     */
    excludesPatterns?: string[];
    /**
     * Pattern of artifacts to include
     */
    includesPatterns?: string[];
    /**
     * List of repositories this permission target is applicable for
     */
    repositories: string[];
}

export interface PermissionTargetsRepoActions {
    /**
     * Groups this permission applies for.
     */
    groups?: outputs.PermissionTargetsRepoActionsGroup[];
    /**
     * Users this permission target applies for.
     */
    users?: outputs.PermissionTargetsRepoActionsUser[];
}

export interface PermissionTargetsRepoActionsGroup {
    /**
     * Name of permission
     */
    name: string;
    permissions: string[];
}

export interface PermissionTargetsRepoActionsUser {
    /**
     * Name of permission
     */
    name: string;
    permissions: string[];
}

export interface PushReplicationReplication {
    enabled: boolean;
    password: string;
    pathPrefix?: string;
    /**
     * Proxy key from Artifactory Proxies setting
     */
    proxy?: string;
    socketTimeoutMillis: number;
    syncDeletes: boolean;
    syncProperties: boolean;
    syncStatistics: boolean;
    url: string;
    username: string;
}

export interface ReleaseBundleWebhookCriteria {
    /**
     * Trigger on any release bundle
     */
    anyReleaseBundle: boolean;
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    excludePatterns?: string[];
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    includePatterns?: string[];
    /**
     * Trigger on this list of release bundle names
     */
    registeredReleaseBundleNames: string[];
}

export interface RemoteAlpineRepositoryContentSynchronisation {
    enabled?: boolean;
    propertiesEnabled?: boolean;
    sourceOriginAbsenceDetection?: boolean;
    statisticsEnabled?: boolean;
}

export interface RemoteBowerRepositoryContentSynchronisation {
    enabled?: boolean;
    propertiesEnabled?: boolean;
    sourceOriginAbsenceDetection?: boolean;
    statisticsEnabled?: boolean;
}

export interface RemoteCargoRepositoryContentSynchronisation {
    /**
     * If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is 'false'.
     */
    enabled?: boolean;
    /**
     * If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is 'false'.
     */
    propertiesEnabled?: boolean;
    /**
     * If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is 'false'
     */
    sourceOriginAbsenceDetection?: boolean;
    /**
     * If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is 'false'.
     */
    statisticsEnabled?: boolean;
}

export interface RemoteChefRepositoryContentSynchronisation {
    enabled?: boolean;
    propertiesEnabled?: boolean;
    sourceOriginAbsenceDetection?: boolean;
    statisticsEnabled?: boolean;
}

export interface RemoteCocoapodsRepositoryContentSynchronisation {
    enabled?: boolean;
    propertiesEnabled?: boolean;
    sourceOriginAbsenceDetection?: boolean;
    statisticsEnabled?: boolean;
}

export interface RemoteComposerRepositoryContentSynchronisation {
    enabled?: boolean;
    propertiesEnabled?: boolean;
    sourceOriginAbsenceDetection?: boolean;
    statisticsEnabled?: boolean;
}

export interface RemoteConanRepositoryContentSynchronisation {
    enabled?: boolean;
    propertiesEnabled?: boolean;
    sourceOriginAbsenceDetection?: boolean;
    statisticsEnabled?: boolean;
}

export interface RemoteCondaRepositoryContentSynchronisation {
    enabled?: boolean;
    propertiesEnabled?: boolean;
    sourceOriginAbsenceDetection?: boolean;
    statisticsEnabled?: boolean;
}

export interface RemoteCranRepositoryContentSynchronisation {
    enabled?: boolean;
    propertiesEnabled?: boolean;
    sourceOriginAbsenceDetection?: boolean;
    statisticsEnabled?: boolean;
}

export interface RemoteDebianRepositoryContentSynchronisation {
    enabled?: boolean;
    propertiesEnabled?: boolean;
    sourceOriginAbsenceDetection?: boolean;
    statisticsEnabled?: boolean;
}

export interface RemoteDockerRepositoryContentSynchronisation {
    /**
     * If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is 'false'.
     */
    enabled?: boolean;
    /**
     * If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is 'false'.
     */
    propertiesEnabled?: boolean;
    /**
     * If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is 'false'
     */
    sourceOriginAbsenceDetection?: boolean;
    /**
     * If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is 'false'.
     */
    statisticsEnabled?: boolean;
}

export interface RemoteGemsRepositoryContentSynchronisation {
    enabled?: boolean;
    propertiesEnabled?: boolean;
    sourceOriginAbsenceDetection?: boolean;
    statisticsEnabled?: boolean;
}

export interface RemoteGenericRepositoryContentSynchronisation {
    /**
     * If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is 'false'.
     */
    enabled?: boolean;
    /**
     * If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is 'false'.
     */
    propertiesEnabled?: boolean;
    /**
     * If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is 'false'
     */
    sourceOriginAbsenceDetection?: boolean;
    /**
     * If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is 'false'.
     */
    statisticsEnabled?: boolean;
}

export interface RemoteGitlfsRepositoryContentSynchronisation {
    enabled?: boolean;
    propertiesEnabled?: boolean;
    sourceOriginAbsenceDetection?: boolean;
    statisticsEnabled?: boolean;
}

export interface RemoteGoRepositoryContentSynchronisation {
    enabled?: boolean;
    propertiesEnabled?: boolean;
    sourceOriginAbsenceDetection?: boolean;
    statisticsEnabled?: boolean;
}

export interface RemoteGradleRepositoryContentSynchronisation {
    /**
     * If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is 'false'.
     */
    enabled?: boolean;
    /**
     * If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is 'false'.
     */
    propertiesEnabled?: boolean;
    /**
     * If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is 'false'
     */
    sourceOriginAbsenceDetection?: boolean;
    /**
     * If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is 'false'.
     */
    statisticsEnabled?: boolean;
}

export interface RemoteHelmRepositoryContentSynchronisation {
    /**
     * If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is 'false'.
     */
    enabled?: boolean;
    /**
     * If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is 'false'.
     */
    propertiesEnabled?: boolean;
    /**
     * If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is 'false'
     */
    sourceOriginAbsenceDetection?: boolean;
    /**
     * If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is 'false'.
     */
    statisticsEnabled?: boolean;
}

export interface RemoteIvyRepositoryContentSynchronisation {
    /**
     * If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is 'false'.
     */
    enabled?: boolean;
    /**
     * If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is 'false'.
     */
    propertiesEnabled?: boolean;
    /**
     * If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is 'false'
     */
    sourceOriginAbsenceDetection?: boolean;
    /**
     * If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is 'false'.
     */
    statisticsEnabled?: boolean;
}

export interface RemoteMavenRepositoryContentSynchronisation {
    /**
     * If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is 'false'.
     */
    enabled?: boolean;
    /**
     * If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is 'false'.
     */
    propertiesEnabled?: boolean;
    /**
     * If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is 'false'
     */
    sourceOriginAbsenceDetection?: boolean;
    /**
     * If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is 'false'.
     */
    statisticsEnabled?: boolean;
}

export interface RemoteNpmRepositoryContentSynchronisation {
    /**
     * If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is 'false'.
     */
    enabled?: boolean;
    /**
     * If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is 'false'.
     */
    propertiesEnabled?: boolean;
    /**
     * If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is 'false'
     */
    sourceOriginAbsenceDetection?: boolean;
    /**
     * If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is 'false'.
     */
    statisticsEnabled?: boolean;
}

export interface RemoteNugetRepositoryContentSynchronisation {
    enabled?: boolean;
    propertiesEnabled?: boolean;
    sourceOriginAbsenceDetection?: boolean;
    statisticsEnabled?: boolean;
}

export interface RemoteOpkgRepositoryContentSynchronisation {
    enabled?: boolean;
    propertiesEnabled?: boolean;
    sourceOriginAbsenceDetection?: boolean;
    statisticsEnabled?: boolean;
}

export interface RemoteP2RepositoryContentSynchronisation {
    enabled?: boolean;
    propertiesEnabled?: boolean;
    sourceOriginAbsenceDetection?: boolean;
    statisticsEnabled?: boolean;
}

export interface RemotePuppetRepositoryContentSynchronisation {
    enabled?: boolean;
    propertiesEnabled?: boolean;
    sourceOriginAbsenceDetection?: boolean;
    statisticsEnabled?: boolean;
}

export interface RemotePypiRepositoryContentSynchronisation {
    /**
     * If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is 'false'.
     */
    enabled?: boolean;
    /**
     * If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is 'false'.
     */
    propertiesEnabled?: boolean;
    /**
     * If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is 'false'
     */
    sourceOriginAbsenceDetection?: boolean;
    /**
     * If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is 'false'.
     */
    statisticsEnabled?: boolean;
}

export interface RemoteRpmRepositoryContentSynchronisation {
    enabled?: boolean;
    propertiesEnabled?: boolean;
    sourceOriginAbsenceDetection?: boolean;
    statisticsEnabled?: boolean;
}

export interface RemoteSbtRepositoryContentSynchronisation {
    /**
     * If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is 'false'.
     */
    enabled?: boolean;
    /**
     * If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is 'false'.
     */
    propertiesEnabled?: boolean;
    /**
     * If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is 'false'
     */
    sourceOriginAbsenceDetection?: boolean;
    /**
     * If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is 'false'.
     */
    statisticsEnabled?: boolean;
}

export interface RemoteVcsRepositoryContentSynchronisation {
    enabled?: boolean;
    propertiesEnabled?: boolean;
    sourceOriginAbsenceDetection?: boolean;
    statisticsEnabled?: boolean;
}

export interface ReplicationConfigReplication {
    enabled: boolean;
    /**
     * Requires password encryption to be turned off `POST /api/system/decrypt`
     */
    password: string;
    pathPrefix?: string;
    /**
     * Proxy key from Artifactory Proxies setting
     */
    proxy?: string;
    socketTimeoutMillis: number;
    syncDeletes: boolean;
    syncProperties: boolean;
    syncStatistics: boolean;
    url?: string;
    username?: string;
}

