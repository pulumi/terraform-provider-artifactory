// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Retrieves a local Bower repository.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const local-test-bower-repo = artifactory.getLocalBowerRepository({
 *     key: "local-test-bower-repo",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getLocalBowerRepository(args: GetLocalBowerRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetLocalBowerRepositoryResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getLocalBowerRepository:getLocalBowerRepository", {
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
 * A collection of arguments for invoking getLocalBowerRepository.
 */
export interface GetLocalBowerRepositoryArgs {
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
     * (Required) the identity key of the repo.
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
 * A collection of values returned by getLocalBowerRepository.
 */
export interface GetLocalBowerRepositoryResult {
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
    /**
     * (Required) the identity key of the repo.
     */
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
 * Retrieves a local Bower repository.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const local-test-bower-repo = artifactory.getLocalBowerRepository({
 *     key: "local-test-bower-repo",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getLocalBowerRepositoryOutput(args: GetLocalBowerRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLocalBowerRepositoryResult> {
    return pulumi.output(args).apply((a: any) => getLocalBowerRepository(a, opts))
}

/**
 * A collection of arguments for invoking getLocalBowerRepository.
 */
export interface GetLocalBowerRepositoryOutputArgs {
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
     * (Required) the identity key of the repo.
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
