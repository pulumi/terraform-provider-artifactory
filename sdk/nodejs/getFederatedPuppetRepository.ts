// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getFederatedPuppetRepository(args: GetFederatedPuppetRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetFederatedPuppetRepositoryResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getFederatedPuppetRepository:getFederatedPuppetRepository", {
        "archiveBrowsingEnabled": args.archiveBrowsingEnabled,
        "blackedOut": args.blackedOut,
        "cdnRedirect": args.cdnRedirect,
        "cleanupOnDelete": args.cleanupOnDelete,
        "description": args.description,
        "downloadDirect": args.downloadDirect,
        "excludesPattern": args.excludesPattern,
        "includesPattern": args.includesPattern,
        "key": args.key,
        "members": args.members,
        "notes": args.notes,
        "priorityResolution": args.priorityResolution,
        "projectEnvironments": args.projectEnvironments,
        "projectKey": args.projectKey,
        "propertySets": args.propertySets,
        "repoLayoutRef": args.repoLayoutRef,
        "xrayIndex": args.xrayIndex,
    }, opts);
}

/**
 * A collection of arguments for invoking getFederatedPuppetRepository.
 */
export interface GetFederatedPuppetRepositoryArgs {
    archiveBrowsingEnabled?: boolean;
    blackedOut?: boolean;
    cdnRedirect?: boolean;
    cleanupOnDelete?: boolean;
    description?: string;
    downloadDirect?: boolean;
    excludesPattern?: string;
    includesPattern?: string;
    key: string;
    members?: inputs.GetFederatedPuppetRepositoryMember[];
    notes?: string;
    priorityResolution?: boolean;
    projectEnvironments?: string[];
    projectKey?: string;
    propertySets?: string[];
    repoLayoutRef?: string;
    xrayIndex?: boolean;
}

/**
 * A collection of values returned by getFederatedPuppetRepository.
 */
export interface GetFederatedPuppetRepositoryResult {
    readonly archiveBrowsingEnabled?: boolean;
    readonly blackedOut?: boolean;
    readonly cdnRedirect?: boolean;
    readonly cleanupOnDelete?: boolean;
    readonly description?: string;
    readonly downloadDirect?: boolean;
    readonly excludesPattern: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includesPattern: string;
    readonly key: string;
    readonly members?: outputs.GetFederatedPuppetRepositoryMember[];
    readonly notes?: string;
    readonly packageType: string;
    readonly priorityResolution?: boolean;
    readonly projectEnvironments: string[];
    readonly projectKey?: string;
    readonly propertySets?: string[];
    readonly repoLayoutRef?: string;
    readonly xrayIndex?: boolean;
}
export function getFederatedPuppetRepositoryOutput(args: GetFederatedPuppetRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFederatedPuppetRepositoryResult> {
    return pulumi.output(args).apply((a: any) => getFederatedPuppetRepository(a, opts))
}

/**
 * A collection of arguments for invoking getFederatedPuppetRepository.
 */
export interface GetFederatedPuppetRepositoryOutputArgs {
    archiveBrowsingEnabled?: pulumi.Input<boolean>;
    blackedOut?: pulumi.Input<boolean>;
    cdnRedirect?: pulumi.Input<boolean>;
    cleanupOnDelete?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    downloadDirect?: pulumi.Input<boolean>;
    excludesPattern?: pulumi.Input<string>;
    includesPattern?: pulumi.Input<string>;
    key: pulumi.Input<string>;
    members?: pulumi.Input<pulumi.Input<inputs.GetFederatedPuppetRepositoryMemberArgs>[]>;
    notes?: pulumi.Input<string>;
    priorityResolution?: pulumi.Input<boolean>;
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    projectKey?: pulumi.Input<string>;
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    repoLayoutRef?: pulumi.Input<string>;
    xrayIndex?: pulumi.Input<boolean>;
}
