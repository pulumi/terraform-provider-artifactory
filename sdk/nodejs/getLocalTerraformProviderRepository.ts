// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const local-test-terraform-provider-repo = artifactory.getLocalTerraformProviderRepository({
 *     key: "local-test-terraform-provider-repo",
 * });
 * ```
 */
export function getLocalTerraformProviderRepository(args: GetLocalTerraformProviderRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetLocalTerraformProviderRepositoryResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getLocalTerraformProviderRepository:getLocalTerraformProviderRepository", {
        "archiveBrowsingEnabled": args.archiveBrowsingEnabled,
        "blackedOut": args.blackedOut,
        "cdnRedirect": args.cdnRedirect,
        "description": args.description,
        "downloadDirect": args.downloadDirect,
        "excludesPattern": args.excludesPattern,
        "includesPattern": args.includesPattern,
        "key": args.key,
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
 * A collection of arguments for invoking getLocalTerraformProviderRepository.
 */
export interface GetLocalTerraformProviderRepositoryArgs {
    archiveBrowsingEnabled?: boolean;
    blackedOut?: boolean;
    cdnRedirect?: boolean;
    /**
     * (Optional)
     */
    description?: string;
    downloadDirect?: boolean;
    excludesPattern?: string;
    includesPattern?: string;
    /**
     * the identity key of the repo.
     */
    key: string;
    /**
     * (Optional)
     */
    notes?: string;
    priorityResolution?: boolean;
    projectEnvironments?: string[];
    projectKey?: string;
    propertySets?: string[];
    repoLayoutRef?: string;
    xrayIndex?: boolean;
}

/**
 * A collection of values returned by getLocalTerraformProviderRepository.
 */
export interface GetLocalTerraformProviderRepositoryResult {
    readonly archiveBrowsingEnabled?: boolean;
    readonly blackedOut?: boolean;
    readonly cdnRedirect?: boolean;
    /**
     * (Optional)
     */
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
     * (Optional)
     */
    readonly notes?: string;
    readonly packageType: string;
    readonly priorityResolution?: boolean;
    readonly projectEnvironments: string[];
    readonly projectKey?: string;
    readonly propertySets?: string[];
    readonly repoLayoutRef?: string;
    readonly xrayIndex?: boolean;
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const local-test-terraform-provider-repo = artifactory.getLocalTerraformProviderRepository({
 *     key: "local-test-terraform-provider-repo",
 * });
 * ```
 */
export function getLocalTerraformProviderRepositoryOutput(args: GetLocalTerraformProviderRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLocalTerraformProviderRepositoryResult> {
    return pulumi.output(args).apply((a: any) => getLocalTerraformProviderRepository(a, opts))
}

/**
 * A collection of arguments for invoking getLocalTerraformProviderRepository.
 */
export interface GetLocalTerraformProviderRepositoryOutputArgs {
    archiveBrowsingEnabled?: pulumi.Input<boolean>;
    blackedOut?: pulumi.Input<boolean>;
    cdnRedirect?: pulumi.Input<boolean>;
    /**
     * (Optional)
     */
    description?: pulumi.Input<string>;
    downloadDirect?: pulumi.Input<boolean>;
    excludesPattern?: pulumi.Input<string>;
    includesPattern?: pulumi.Input<string>;
    /**
     * the identity key of the repo.
     */
    key: pulumi.Input<string>;
    /**
     * (Optional)
     */
    notes?: pulumi.Input<string>;
    priorityResolution?: pulumi.Input<boolean>;
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    projectKey?: pulumi.Input<string>;
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    repoLayoutRef?: pulumi.Input<string>;
    xrayIndex?: pulumi.Input<boolean>;
}
