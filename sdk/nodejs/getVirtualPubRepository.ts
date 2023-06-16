// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getVirtualPubRepository(args: GetVirtualPubRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualPubRepositoryResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getVirtualPubRepository:getVirtualPubRepository", {
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
 * A collection of arguments for invoking getVirtualPubRepository.
 */
export interface GetVirtualPubRepositoryArgs {
    artifactoryRequestsCanRetrieveRemoteArtifacts?: boolean;
    defaultDeploymentRepo?: string;
    description?: string;
    excludesPattern?: string;
    includesPattern?: string;
    key: string;
    notes?: string;
    projectEnvironments?: string[];
    projectKey?: string;
    repoLayoutRef?: string;
    repositories?: string[];
}

/**
 * A collection of values returned by getVirtualPubRepository.
 */
export interface GetVirtualPubRepositoryResult {
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
export function getVirtualPubRepositoryOutput(args: GetVirtualPubRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVirtualPubRepositoryResult> {
    return pulumi.output(args).apply((a: any) => getVirtualPubRepository(a, opts))
}

/**
 * A collection of arguments for invoking getVirtualPubRepository.
 */
export interface GetVirtualPubRepositoryOutputArgs {
    artifactoryRequestsCanRetrieveRemoteArtifacts?: pulumi.Input<boolean>;
    defaultDeploymentRepo?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    excludesPattern?: pulumi.Input<string>;
    includesPattern?: pulumi.Input<string>;
    key: pulumi.Input<string>;
    notes?: pulumi.Input<string>;
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    projectKey?: pulumi.Input<string>;
    repoLayoutRef?: pulumi.Input<string>;
    repositories?: pulumi.Input<pulumi.Input<string>[]>;
}
