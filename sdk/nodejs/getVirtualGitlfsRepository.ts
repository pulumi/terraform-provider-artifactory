// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Retrieves a virtual Git LFS repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const virtual_gitlfs = artifactory.getVirtualGitlfsRepository({
 *     key: "virtual-gitlfs",
 * });
 * ```
 */
export function getVirtualGitlfsRepository(args: GetVirtualGitlfsRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualGitlfsRepositoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getVirtualGitlfsRepository:getVirtualGitlfsRepository", {
        "artifactoryRequestsCanRetrieveRemoteArtifacts": args.artifactoryRequestsCanRetrieveRemoteArtifacts,
        "defaultDeploymentRepo": args.defaultDeploymentRepo,
        "description": args.description,
        "excludesPattern": args.excludesPattern,
        "includesPattern": args.includesPattern,
        "key": args.key,
        "notes": args.notes,
        "projectEnvironments": args.projectEnvironments,
        "projectKey": args.projectKey,
        "repoLayoutRef": args.repoLayoutRef,
        "repositories": args.repositories,
    }, opts);
}

/**
 * A collection of arguments for invoking getVirtualGitlfsRepository.
 */
export interface GetVirtualGitlfsRepositoryArgs {
    artifactoryRequestsCanRetrieveRemoteArtifacts?: boolean;
    defaultDeploymentRepo?: string;
    description?: string;
    excludesPattern?: string;
    includesPattern?: string;
    /**
     * the identity key of the repo.
     */
    key: string;
    notes?: string;
    projectEnvironments?: string[];
    projectKey?: string;
    repoLayoutRef?: string;
    repositories?: string[];
}

/**
 * A collection of values returned by getVirtualGitlfsRepository.
 */
export interface GetVirtualGitlfsRepositoryResult {
    readonly artifactoryRequestsCanRetrieveRemoteArtifacts?: boolean;
    readonly defaultDeploymentRepo?: string;
    readonly description?: string;
    readonly excludesPattern?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includesPattern?: string;
    readonly key: string;
    readonly notes?: string;
    readonly packageType: string;
    readonly projectEnvironments: string[];
    readonly projectKey?: string;
    readonly repoLayoutRef?: string;
    readonly repositories?: string[];
}
/**
 * Retrieves a virtual Git LFS repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const virtual_gitlfs = artifactory.getVirtualGitlfsRepository({
 *     key: "virtual-gitlfs",
 * });
 * ```
 */
export function getVirtualGitlfsRepositoryOutput(args: GetVirtualGitlfsRepositoryOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetVirtualGitlfsRepositoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("artifactory:index/getVirtualGitlfsRepository:getVirtualGitlfsRepository", {
        "artifactoryRequestsCanRetrieveRemoteArtifacts": args.artifactoryRequestsCanRetrieveRemoteArtifacts,
        "defaultDeploymentRepo": args.defaultDeploymentRepo,
        "description": args.description,
        "excludesPattern": args.excludesPattern,
        "includesPattern": args.includesPattern,
        "key": args.key,
        "notes": args.notes,
        "projectEnvironments": args.projectEnvironments,
        "projectKey": args.projectKey,
        "repoLayoutRef": args.repoLayoutRef,
        "repositories": args.repositories,
    }, opts);
}

/**
 * A collection of arguments for invoking getVirtualGitlfsRepository.
 */
export interface GetVirtualGitlfsRepositoryOutputArgs {
    artifactoryRequestsCanRetrieveRemoteArtifacts?: pulumi.Input<boolean>;
    defaultDeploymentRepo?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    excludesPattern?: pulumi.Input<string>;
    includesPattern?: pulumi.Input<string>;
    /**
     * the identity key of the repo.
     */
    key: pulumi.Input<string>;
    notes?: pulumi.Input<string>;
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    projectKey?: pulumi.Input<string>;
    repoLayoutRef?: pulumi.Input<string>;
    repositories?: pulumi.Input<pulumi.Input<string>[]>;
}
