// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Retrieves a virtual Rpm repository.
 */
export function getVirtualRpmRepository(args: GetVirtualRpmRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualRpmRepositoryResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getVirtualRpmRepository:getVirtualRpmRepository", {
        "artifactoryRequestsCanRetrieveRemoteArtifacts": args.artifactoryRequestsCanRetrieveRemoteArtifacts,
        "defaultDeploymentRepo": args.defaultDeploymentRepo,
        "description": args.description,
        "excludesPattern": args.excludesPattern,
        "includesPattern": args.includesPattern,
        "key": args.key,
        "notes": args.notes,
        "primaryKeypairRef": args.primaryKeypairRef,
        "projectEnvironments": args.projectEnvironments,
        "projectKey": args.projectKey,
        "repoLayoutRef": args.repoLayoutRef,
        "repositories": args.repositories,
        "secondaryKeypairRef": args.secondaryKeypairRef,
    }, opts);
}

/**
 * A collection of arguments for invoking getVirtualRpmRepository.
 */
export interface GetVirtualRpmRepositoryArgs {
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
    /**
     * (Optional) The primary GPG key to be used to sign packages.
     */
    primaryKeypairRef?: string;
    projectEnvironments?: string[];
    projectKey?: string;
    repoLayoutRef?: string;
    repositories?: string[];
    /**
     * (Optional) The secondary GPG key to be used to sign packages.
     */
    secondaryKeypairRef?: string;
}

/**
 * A collection of values returned by getVirtualRpmRepository.
 */
export interface GetVirtualRpmRepositoryResult {
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
    /**
     * (Optional) The primary GPG key to be used to sign packages.
     */
    readonly primaryKeypairRef?: string;
    readonly projectEnvironments: string[];
    readonly projectKey?: string;
    readonly repoLayoutRef?: string;
    readonly repositories?: string[];
    /**
     * (Optional) The secondary GPG key to be used to sign packages.
     */
    readonly secondaryKeypairRef?: string;
}
/**
 * Retrieves a virtual Rpm repository.
 */
export function getVirtualRpmRepositoryOutput(args: GetVirtualRpmRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVirtualRpmRepositoryResult> {
    return pulumi.output(args).apply((a: any) => getVirtualRpmRepository(a, opts))
}

/**
 * A collection of arguments for invoking getVirtualRpmRepository.
 */
export interface GetVirtualRpmRepositoryOutputArgs {
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
    /**
     * (Optional) The primary GPG key to be used to sign packages.
     */
    primaryKeypairRef?: pulumi.Input<string>;
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    projectKey?: pulumi.Input<string>;
    repoLayoutRef?: pulumi.Input<string>;
    repositories?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional) The secondary GPG key to be used to sign packages.
     */
    secondaryKeypairRef?: pulumi.Input<string>;
}
