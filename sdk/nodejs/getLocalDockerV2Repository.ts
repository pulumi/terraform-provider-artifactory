// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Retrieves a local Docker (V2) repository resource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const artifactoryLocalTestDockerV2Repository = artifactory.getLocalDockerV2Repository({
 *     key: "artifactory_local_test_docker_v2_repository",
 * });
 * ```
 */
export function getLocalDockerV2Repository(args: GetLocalDockerV2RepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetLocalDockerV2RepositoryResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getLocalDockerV2Repository:getLocalDockerV2Repository", {
        "archiveBrowsingEnabled": args.archiveBrowsingEnabled,
        "blackedOut": args.blackedOut,
        "blockPushingSchema1": args.blockPushingSchema1,
        "cdnRedirect": args.cdnRedirect,
        "description": args.description,
        "downloadDirect": args.downloadDirect,
        "excludesPattern": args.excludesPattern,
        "includesPattern": args.includesPattern,
        "key": args.key,
        "maxUniqueTags": args.maxUniqueTags,
        "notes": args.notes,
        "priorityResolution": args.priorityResolution,
        "projectEnvironments": args.projectEnvironments,
        "projectKey": args.projectKey,
        "propertySets": args.propertySets,
        "repoLayoutRef": args.repoLayoutRef,
        "tagRetention": args.tagRetention,
        "xrayIndex": args.xrayIndex,
    }, opts);
}

/**
 * A collection of arguments for invoking getLocalDockerV2Repository.
 */
export interface GetLocalDockerV2RepositoryArgs {
    archiveBrowsingEnabled?: boolean;
    blackedOut?: boolean;
    /**
     * When set, Artifactory will block the pushing of Docker images with manifest v2
     * schema 1 to this repository.
     */
    blockPushingSchema1?: boolean;
    cdnRedirect?: boolean;
    description?: string;
    downloadDirect?: boolean;
    excludesPattern?: string;
    includesPattern?: string;
    /**
     * the identity key of the repo.
     */
    key: string;
    /**
     * The maximum number of unique tags of a single Docker image to store in this repository.
     * Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there
     * is no limit. This only applies to manifest v2.
     */
    maxUniqueTags?: number;
    notes?: string;
    priorityResolution?: boolean;
    projectEnvironments?: string[];
    projectKey?: string;
    propertySets?: string[];
    repoLayoutRef?: string;
    /**
     * If greater than 1, overwritten tags will be saved by their digest, up to the set up
     * number. This only applies to manifest V2.
     */
    tagRetention?: number;
    xrayIndex?: boolean;
}

/**
 * A collection of values returned by getLocalDockerV2Repository.
 */
export interface GetLocalDockerV2RepositoryResult {
    /**
     * "The Docker API version to use. This cannot be set"
     */
    readonly apiVersion: string;
    readonly archiveBrowsingEnabled?: boolean;
    readonly blackedOut?: boolean;
    /**
     * When set, Artifactory will block the pushing of Docker images with manifest v2
     * schema 1 to this repository.
     */
    readonly blockPushingSchema1: boolean;
    readonly cdnRedirect?: boolean;
    readonly description?: string;
    readonly downloadDirect?: boolean;
    readonly excludesPattern?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includesPattern?: string;
    readonly key: string;
    /**
     * The maximum number of unique tags of a single Docker image to store in this repository.
     * Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there
     * is no limit. This only applies to manifest v2.
     */
    readonly maxUniqueTags?: number;
    readonly notes?: string;
    readonly packageType: string;
    readonly priorityResolution?: boolean;
    readonly projectEnvironments: string[];
    readonly projectKey?: string;
    readonly propertySets?: string[];
    readonly repoLayoutRef?: string;
    /**
     * If greater than 1, overwritten tags will be saved by their digest, up to the set up
     * number. This only applies to manifest V2.
     */
    readonly tagRetention?: number;
    readonly xrayIndex?: boolean;
}
/**
 * Retrieves a local Docker (V2) repository resource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const artifactoryLocalTestDockerV2Repository = artifactory.getLocalDockerV2Repository({
 *     key: "artifactory_local_test_docker_v2_repository",
 * });
 * ```
 */
export function getLocalDockerV2RepositoryOutput(args: GetLocalDockerV2RepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLocalDockerV2RepositoryResult> {
    return pulumi.output(args).apply((a: any) => getLocalDockerV2Repository(a, opts))
}

/**
 * A collection of arguments for invoking getLocalDockerV2Repository.
 */
export interface GetLocalDockerV2RepositoryOutputArgs {
    archiveBrowsingEnabled?: pulumi.Input<boolean>;
    blackedOut?: pulumi.Input<boolean>;
    /**
     * When set, Artifactory will block the pushing of Docker images with manifest v2
     * schema 1 to this repository.
     */
    blockPushingSchema1?: pulumi.Input<boolean>;
    cdnRedirect?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    downloadDirect?: pulumi.Input<boolean>;
    excludesPattern?: pulumi.Input<string>;
    includesPattern?: pulumi.Input<string>;
    /**
     * the identity key of the repo.
     */
    key: pulumi.Input<string>;
    /**
     * The maximum number of unique tags of a single Docker image to store in this repository.
     * Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there
     * is no limit. This only applies to manifest v2.
     */
    maxUniqueTags?: pulumi.Input<number>;
    notes?: pulumi.Input<string>;
    priorityResolution?: pulumi.Input<boolean>;
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    projectKey?: pulumi.Input<string>;
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    repoLayoutRef?: pulumi.Input<string>;
    /**
     * If greater than 1, overwritten tags will be saved by their digest, up to the set up
     * number. This only applies to manifest V2.
     */
    tagRetention?: pulumi.Input<number>;
    xrayIndex?: pulumi.Input<boolean>;
}
