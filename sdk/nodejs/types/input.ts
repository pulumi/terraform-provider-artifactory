// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export interface AccessTokenAdminToken {
    instanceId: pulumi.Input<string>;
}

export interface ArtifactPropertyWebhookCriteria {
    /**
     * Trigger on any local repo
     */
    anyLocal: pulumi.Input<boolean>;
    /**
     * Trigger on any remote repo
     */
    anyRemote: pulumi.Input<boolean>;
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    excludePatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    includePatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Trigger on this list of repo keys
     */
    repoKeys: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ArtifactWebhookCriteria {
    /**
     * Trigger on any local repo
     */
    anyLocal: pulumi.Input<boolean>;
    /**
     * Trigger on any remote repo
     */
    anyRemote: pulumi.Input<boolean>;
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    excludePatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    includePatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Trigger on this list of repo keys
     */
    repoKeys: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ArtifactoryReleaseBundleWebhookCriteria {
    /**
     * Trigger on any release bundle
     */
    anyReleaseBundle: pulumi.Input<boolean>;
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    excludePatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    includePatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Trigger on this list of release bundle names
     */
    registeredReleaseBundleNames: pulumi.Input<pulumi.Input<string>[]>;
}

export interface BuildWebhookCriteria {
    /**
     * Trigger on any build
     */
    anyBuild: pulumi.Input<boolean>;
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    excludePatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    includePatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Trigger on this list of build names
     */
    selectedBuilds: pulumi.Input<pulumi.Input<string>[]>;
}

export interface DistributionWebhookCriteria {
    /**
     * Trigger on any release bundle
     */
    anyReleaseBundle: pulumi.Input<boolean>;
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    excludePatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    includePatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Trigger on this list of release bundle names
     */
    registeredReleaseBundleNames: pulumi.Input<pulumi.Input<string>[]>;
}

export interface DockerWebhookCriteria {
    /**
     * Trigger on any local repo
     */
    anyLocal: pulumi.Input<boolean>;
    /**
     * Trigger on any remote repo
     */
    anyRemote: pulumi.Input<boolean>;
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    excludePatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    includePatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Trigger on this list of repo keys
     */
    repoKeys: pulumi.Input<pulumi.Input<string>[]>;
}

export interface FederatedAlpineRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface FederatedBowerRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface FederatedCargoRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface FederatedChefRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface FederatedCocoapodsRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface FederatedComposerRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface FederatedConanRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface FederatedCondaRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface FederatedCranRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface FederatedDebianRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface FederatedDockerRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface FederatedGemsRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface FederatedGenericRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface FederatedGitltfsRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface FederatedGoRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface FederatedGradleRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface FederatedHelmRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface FederatedIvyRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface FederatedMavenRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface FederatedNpmRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface FederatedNugetRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface FederatedOpkgRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface FederatedPuppetRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface FederatedPypiRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface FederatedRpmRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface FederatedSbtRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface FederatedVagrantRepositoryMember {
    /**
     * Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Full URL to ending with the repository name
     */
    url: pulumi.Input<string>;
}

export interface OauthSettingsOauthProvider {
    /**
     * OAuth user info endpoint for the IdP.
     */
    apiUrl: pulumi.Input<string>;
    /**
     * OAuth authorization endpoint for the IdP.
     */
    authUrl: pulumi.Input<string>;
    /**
     * OAuth client ID configured on the IdP.
     */
    clientId: pulumi.Input<string>;
    /**
     * OAuth client secret configured on the IdP.
     */
    clientSecret: pulumi.Input<string>;
    /**
     * Enable the Artifactory OAuth provider.  Default value is `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Name of the Artifactory OAuth provider.
     */
    name: pulumi.Input<string>;
    /**
     * OAuth token endpoint for the IdP.
     */
    tokenUrl: pulumi.Input<string>;
    /**
     * Type of OAuth provider. (e.g., `github`, `google`, `cloudfoundry`, or `openId`)
     */
    type: pulumi.Input<string>;
}

export interface PermissionTargetBuild {
    /**
     * -
     */
    actions?: pulumi.Input<inputs.PermissionTargetBuildActions>;
    /**
     * Pattern of artifacts to exclude
     */
    excludesPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Pattern of artifacts to include
     */
    includesPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of repositories this permission target is applicable for
     */
    repositories: pulumi.Input<pulumi.Input<string>[]>;
}

export interface PermissionTargetBuildActions {
    /**
     * Groups this permission applies for.
     */
    groups?: pulumi.Input<pulumi.Input<inputs.PermissionTargetBuildActionsGroup>[]>;
    /**
     * Users this permission target applies for.
     */
    users?: pulumi.Input<pulumi.Input<inputs.PermissionTargetBuildActionsUser>[]>;
}

export interface PermissionTargetBuildActionsGroup {
    /**
     * Name of permission
     */
    name: pulumi.Input<string>;
    permissions: pulumi.Input<pulumi.Input<string>[]>;
}

export interface PermissionTargetBuildActionsUser {
    /**
     * Name of permission
     */
    name: pulumi.Input<string>;
    permissions: pulumi.Input<pulumi.Input<string>[]>;
}

export interface PermissionTargetReleaseBundle {
    /**
     * -
     */
    actions?: pulumi.Input<inputs.PermissionTargetReleaseBundleActions>;
    /**
     * Pattern of artifacts to exclude
     */
    excludesPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Pattern of artifacts to include
     */
    includesPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of repositories this permission target is applicable for
     */
    repositories: pulumi.Input<pulumi.Input<string>[]>;
}

export interface PermissionTargetReleaseBundleActions {
    /**
     * Groups this permission applies for.
     */
    groups?: pulumi.Input<pulumi.Input<inputs.PermissionTargetReleaseBundleActionsGroup>[]>;
    /**
     * Users this permission target applies for.
     */
    users?: pulumi.Input<pulumi.Input<inputs.PermissionTargetReleaseBundleActionsUser>[]>;
}

export interface PermissionTargetReleaseBundleActionsGroup {
    /**
     * Name of permission
     */
    name: pulumi.Input<string>;
    permissions: pulumi.Input<pulumi.Input<string>[]>;
}

export interface PermissionTargetReleaseBundleActionsUser {
    /**
     * Name of permission
     */
    name: pulumi.Input<string>;
    permissions: pulumi.Input<pulumi.Input<string>[]>;
}

export interface PermissionTargetRepo {
    /**
     * -
     */
    actions?: pulumi.Input<inputs.PermissionTargetRepoActions>;
    /**
     * Pattern of artifacts to exclude
     */
    excludesPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Pattern of artifacts to include
     */
    includesPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of repositories this permission target is applicable for
     */
    repositories: pulumi.Input<pulumi.Input<string>[]>;
}

export interface PermissionTargetRepoActions {
    /**
     * Groups this permission applies for.
     */
    groups?: pulumi.Input<pulumi.Input<inputs.PermissionTargetRepoActionsGroup>[]>;
    /**
     * Users this permission target applies for.
     */
    users?: pulumi.Input<pulumi.Input<inputs.PermissionTargetRepoActionsUser>[]>;
}

export interface PermissionTargetRepoActionsGroup {
    /**
     * Name of permission
     */
    name: pulumi.Input<string>;
    permissions: pulumi.Input<pulumi.Input<string>[]>;
}

export interface PermissionTargetRepoActionsUser {
    /**
     * Name of permission
     */
    name: pulumi.Input<string>;
    permissions: pulumi.Input<pulumi.Input<string>[]>;
}

export interface PermissionTargetsBuild {
    actions?: pulumi.Input<inputs.PermissionTargetsBuildActions>;
    excludesPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    includesPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    repositories: pulumi.Input<pulumi.Input<string>[]>;
}

export interface PermissionTargetsBuildActions {
    groups?: pulumi.Input<pulumi.Input<inputs.PermissionTargetsBuildActionsGroup>[]>;
    users?: pulumi.Input<pulumi.Input<inputs.PermissionTargetsBuildActionsUser>[]>;
}

export interface PermissionTargetsBuildActionsGroup {
    name: pulumi.Input<string>;
    permissions: pulumi.Input<pulumi.Input<string>[]>;
}

export interface PermissionTargetsBuildActionsUser {
    name: pulumi.Input<string>;
    permissions: pulumi.Input<pulumi.Input<string>[]>;
}

export interface PermissionTargetsReleaseBundle {
    actions?: pulumi.Input<inputs.PermissionTargetsReleaseBundleActions>;
    excludesPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    includesPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    repositories: pulumi.Input<pulumi.Input<string>[]>;
}

export interface PermissionTargetsReleaseBundleActions {
    groups?: pulumi.Input<pulumi.Input<inputs.PermissionTargetsReleaseBundleActionsGroup>[]>;
    users?: pulumi.Input<pulumi.Input<inputs.PermissionTargetsReleaseBundleActionsUser>[]>;
}

export interface PermissionTargetsReleaseBundleActionsGroup {
    name: pulumi.Input<string>;
    permissions: pulumi.Input<pulumi.Input<string>[]>;
}

export interface PermissionTargetsReleaseBundleActionsUser {
    name: pulumi.Input<string>;
    permissions: pulumi.Input<pulumi.Input<string>[]>;
}

export interface PermissionTargetsRepo {
    actions?: pulumi.Input<inputs.PermissionTargetsRepoActions>;
    excludesPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    includesPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    repositories: pulumi.Input<pulumi.Input<string>[]>;
}

export interface PermissionTargetsRepoActions {
    groups?: pulumi.Input<pulumi.Input<inputs.PermissionTargetsRepoActionsGroup>[]>;
    users?: pulumi.Input<pulumi.Input<inputs.PermissionTargetsRepoActionsUser>[]>;
}

export interface PermissionTargetsRepoActionsGroup {
    name: pulumi.Input<string>;
    permissions: pulumi.Input<pulumi.Input<string>[]>;
}

export interface PermissionTargetsRepoActionsUser {
    name: pulumi.Input<string>;
    permissions: pulumi.Input<pulumi.Input<string>[]>;
}

export interface PushReplicationReplication {
    enabled?: pulumi.Input<boolean>;
    /**
     * Requires password encryption to be turned off `POST /api/system/decrypt`
     */
    password?: pulumi.Input<string>;
    pathPrefix?: pulumi.Input<string>;
    socketTimeoutMillis?: pulumi.Input<number>;
    syncDeletes?: pulumi.Input<boolean>;
    syncProperties?: pulumi.Input<boolean>;
    syncStatistics?: pulumi.Input<boolean>;
    url?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
}

export interface ReleaseBundleWebhookCriteria {
    /**
     * Trigger on any release bundle
     */
    anyReleaseBundle: pulumi.Input<boolean>;
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    excludePatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
     */
    includePatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Trigger on this list of release bundle names
     */
    registeredReleaseBundleNames: pulumi.Input<pulumi.Input<string>[]>;
}

export interface RemoteCargoRepositoryContentSynchronisation {
    /**
     * If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is 'false'.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is 'false'.
     */
    propertiesEnabled?: pulumi.Input<boolean>;
    /**
     * If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is 'false'
     */
    sourceOriginAbsenceDetection?: pulumi.Input<boolean>;
    /**
     * If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is 'false'.
     */
    statisticsEnabled?: pulumi.Input<boolean>;
}

export interface RemoteDockerRepositoryContentSynchronisation {
    /**
     * If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is 'false'.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is 'false'.
     */
    propertiesEnabled?: pulumi.Input<boolean>;
    /**
     * If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is 'false'
     */
    sourceOriginAbsenceDetection?: pulumi.Input<boolean>;
    /**
     * If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is 'false'.
     */
    statisticsEnabled?: pulumi.Input<boolean>;
}

export interface RemoteGradleRepositoryContentSynchronisation {
    /**
     * If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is 'false'.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is 'false'.
     */
    propertiesEnabled?: pulumi.Input<boolean>;
    /**
     * If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is 'false'
     */
    sourceOriginAbsenceDetection?: pulumi.Input<boolean>;
    /**
     * If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is 'false'.
     */
    statisticsEnabled?: pulumi.Input<boolean>;
}

export interface RemoteHelmRepositoryContentSynchronisation {
    /**
     * If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is 'false'.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is 'false'.
     */
    propertiesEnabled?: pulumi.Input<boolean>;
    /**
     * If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is 'false'
     */
    sourceOriginAbsenceDetection?: pulumi.Input<boolean>;
    /**
     * If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is 'false'.
     */
    statisticsEnabled?: pulumi.Input<boolean>;
}

export interface RemoteMavenRepositoryContentSynchronisation {
    /**
     * If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is 'false'.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is 'false'.
     */
    propertiesEnabled?: pulumi.Input<boolean>;
    /**
     * If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is 'false'
     */
    sourceOriginAbsenceDetection?: pulumi.Input<boolean>;
    /**
     * If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is 'false'.
     */
    statisticsEnabled?: pulumi.Input<boolean>;
}

export interface RemoteNpmRepositoryContentSynchronisation {
    /**
     * If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is 'false'.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is 'false'.
     */
    propertiesEnabled?: pulumi.Input<boolean>;
    /**
     * If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is 'false'
     */
    sourceOriginAbsenceDetection?: pulumi.Input<boolean>;
    /**
     * If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is 'false'.
     */
    statisticsEnabled?: pulumi.Input<boolean>;
}

export interface RemotePypiRepositoryContentSynchronisation {
    /**
     * If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is 'false'.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is 'false'.
     */
    propertiesEnabled?: pulumi.Input<boolean>;
    /**
     * If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is 'false'
     */
    sourceOriginAbsenceDetection?: pulumi.Input<boolean>;
    /**
     * If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is 'false'.
     */
    statisticsEnabled?: pulumi.Input<boolean>;
}

export interface RemoteRepositoryContentSynchronisation {
    /**
     * If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is 'false'.
     */
    enabled?: pulumi.Input<boolean>;
}

export interface ReplicationConfigReplication {
    enabled?: pulumi.Input<boolean>;
    /**
     * Requires password encryption to be turned off `POST /api/system/decrypt`
     */
    password?: pulumi.Input<string>;
    pathPrefix?: pulumi.Input<string>;
    socketTimeoutMillis?: pulumi.Input<number>;
    syncDeletes?: pulumi.Input<boolean>;
    syncProperties?: pulumi.Input<boolean>;
    syncStatistics?: pulumi.Input<boolean>;
    url?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
}

export interface XrayPolicyRule {
    /**
     * (Required) Nested block describing the actions to be applied by the policy. Described below.
     */
    actions?: pulumi.Input<inputs.XrayPolicyRuleActions>;
    /**
     * (Required) Nested block describing the criteria for the policy. Described below.
     */
    criteria: pulumi.Input<inputs.XrayPolicyRuleCriteria>;
    /**
     * (Required) Name of the rule
     */
    name: pulumi.Input<string>;
    /**
     * (Required) Integer describing the rule priority
     */
    priority: pulumi.Input<number>;
}

export interface XrayPolicyRuleActions {
    /**
     * (Optional) Nested block describing artifacts that should be blocked for download if a violation is triggered. Described below.
     */
    blockDownload: pulumi.Input<inputs.XrayPolicyRuleActionsBlockDownload>;
    /**
     * (Optional) The severity of violation to be triggered if the `criteria` are met.
     */
    customSeverity?: pulumi.Input<string>;
    /**
     * (Optional) Whether or not the related CI build should be marked as failed if a violation is triggered. This option is only available when the policy is applied to an `xrayWatch` resource with a `type` of `builds`.
     */
    failBuild?: pulumi.Input<boolean>;
    /**
     * (Optional) A list of email addressed that will get emailed when a violation is triggered.
     */
    mails?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional) A list of Xray-configured webhook URLs to be invoked if a violation is triggered.
     */
    webhooks?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface XrayPolicyRuleActionsBlockDownload {
    /**
     * Whether or not to block download of artifacts that meet the artifact and severity `filters` for the associated `xrayWatch` resource.
     */
    active: pulumi.Input<boolean>;
    /**
     * Whether or not to block download of artifacts that meet the artifact `filters` for the associated `xrayWatch` resource but have not been scanned yet.
     */
    unscanned: pulumi.Input<boolean>;
}

export interface XrayPolicyRuleCriteria {
    /**
     * (Optional) Whether or not to allow components whose license cannot be determined (`true` or `false`).
     */
    allowUnknown?: pulumi.Input<boolean>;
    /**
     * (Optional) A list of OSS license names that may be attached to a component.
     */
    allowedLicenses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional) A list of OSS license names that may not be attached to a component.
     */
    bannedLicenses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional) Nested block describing a CVS score range to be impacted. Defined below.
     */
    cvssRange?: pulumi.Input<inputs.XrayPolicyRuleCriteriaCvssRange>;
    /**
     * (Optional) The minimum security vulnerability severity that will be impacted by the policy.
     */
    minSeverity?: pulumi.Input<string>;
}

export interface XrayPolicyRuleCriteriaCvssRange {
    /**
     * (Required) The beginning of the range of CVS scores (from 1-10) to flag.
     */
    from: pulumi.Input<number>;
    /**
     * (Required) The end of the range of CVS scores (from 1-10) to flag.
     */
    to: pulumi.Input<number>;
}

export interface XrayWatchAssignedPolicy {
    /**
     * The name of the policy that will be applied
     */
    name: pulumi.Input<string>;
    /**
     * The type of the policy
     */
    type: pulumi.Input<string>;
}

export interface XrayWatchResource {
    /**
     * The ID number of a binary manager resource
     */
    binMgrId?: pulumi.Input<string>;
    /**
     * Nested argument describing filters to be applied. Defined below.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.XrayWatchResourceFilter>[]>;
    /**
     * A name describing the resource
     */
    name: pulumi.Input<string>;
    /**
     * Type of repository (e.g. local or remote)
     */
    repoType?: pulumi.Input<string>;
    /**
     * Type of resource to be watched
     */
    type: pulumi.Input<string>;
}

export interface XrayWatchResourceFilter {
    /**
     * The type of filter, such as `regex` or `package-type`
     */
    type: pulumi.Input<string>;
    /**
     * The value of the filter, such as the text of the regex or name of the package type
     */
    value: pulumi.Input<string>;
}

