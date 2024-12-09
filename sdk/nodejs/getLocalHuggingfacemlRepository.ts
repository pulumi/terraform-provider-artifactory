// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a data source for a local huggingfaceml repository
 */
export function getLocalHuggingfacemlRepository(args: GetLocalHuggingfacemlRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetLocalHuggingfacemlRepositoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getLocalHuggingfacemlRepository:getLocalHuggingfacemlRepository", {
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
 * A collection of arguments for invoking getLocalHuggingfacemlRepository.
 */
export interface GetLocalHuggingfacemlRepositoryArgs {
    /**
     * When set, you may view content such as HTML or Javadoc files directly from Artifactory.
     * This may not be safe and therefore requires strict content moderation to prevent malicious users from uploading content that may compromise security (e.g., cross-site scripting attacks).
     */
    archiveBrowsingEnabled?: boolean;
    /**
     * When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
     */
    blackedOut?: boolean;
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from AWS CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
     */
    cdnRedirect?: boolean;
    /**
     * Public description.
     */
    description?: string;
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from the cloud storage provider. Available in Enterprise+ and Edge licenses only.
     */
    downloadDirect?: boolean;
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**&#47;z/*.By default no artifacts are excluded.
     */
    excludesPattern?: string;
    /**
     * List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**&#47;z/*. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**&#47;*).
     */
    includesPattern?: string;
    /**
     * A mandatory identifier for the repository that must be unique. Must be 3 - 10 lowercase alphanumeric and hyphen characters. It cannot begin with a number or contain spaces or special characters.
     */
    key: string;
    /**
     * Internal description.
     */
    notes?: string;
    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     */
    priorityResolution?: boolean;
    projectEnvironments?: string[];
    /**
     * Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     */
    projectKey?: string;
    /**
     * List of property set name
     */
    propertySets?: string[];
    /**
     * Repository layout key for the local repository
     */
    repoLayoutRef?: string;
    /**
     * Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via Xray settings.
     */
    xrayIndex?: boolean;
}

/**
 * A collection of values returned by getLocalHuggingfacemlRepository.
 */
export interface GetLocalHuggingfacemlRepositoryResult {
    /**
     * When set, you may view content such as HTML or Javadoc files directly from Artifactory.
     * This may not be safe and therefore requires strict content moderation to prevent malicious users from uploading content that may compromise security (e.g., cross-site scripting attacks).
     */
    readonly archiveBrowsingEnabled?: boolean;
    /**
     * When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
     */
    readonly blackedOut?: boolean;
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from AWS CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
     */
    readonly cdnRedirect?: boolean;
    /**
     * Public description.
     */
    readonly description?: string;
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from the cloud storage provider. Available in Enterprise+ and Edge licenses only.
     */
    readonly downloadDirect?: boolean;
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**&#47;z/*.By default no artifacts are excluded.
     */
    readonly excludesPattern?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**&#47;z/*. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**&#47;*).
     */
    readonly includesPattern?: string;
    /**
     * A mandatory identifier for the repository that must be unique. Must be 3 - 10 lowercase alphanumeric and hyphen characters. It cannot begin with a number or contain spaces or special characters.
     */
    readonly key: string;
    /**
     * Internal description.
     */
    readonly notes?: string;
    readonly packageType: string;
    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     */
    readonly priorityResolution?: boolean;
    readonly projectEnvironments: string[];
    /**
     * Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     */
    readonly projectKey?: string;
    /**
     * List of property set name
     */
    readonly propertySets?: string[];
    /**
     * Repository layout key for the local repository
     */
    readonly repoLayoutRef?: string;
    /**
     * Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via Xray settings.
     */
    readonly xrayIndex?: boolean;
}
/**
 * Provides a data source for a local huggingfaceml repository
 */
export function getLocalHuggingfacemlRepositoryOutput(args: GetLocalHuggingfacemlRepositoryOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetLocalHuggingfacemlRepositoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("artifactory:index/getLocalHuggingfacemlRepository:getLocalHuggingfacemlRepository", {
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
 * A collection of arguments for invoking getLocalHuggingfacemlRepository.
 */
export interface GetLocalHuggingfacemlRepositoryOutputArgs {
    /**
     * When set, you may view content such as HTML or Javadoc files directly from Artifactory.
     * This may not be safe and therefore requires strict content moderation to prevent malicious users from uploading content that may compromise security (e.g., cross-site scripting attacks).
     */
    archiveBrowsingEnabled?: pulumi.Input<boolean>;
    /**
     * When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
     */
    blackedOut?: pulumi.Input<boolean>;
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from AWS CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
     */
    cdnRedirect?: pulumi.Input<boolean>;
    /**
     * Public description.
     */
    description?: pulumi.Input<string>;
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from the cloud storage provider. Available in Enterprise+ and Edge licenses only.
     */
    downloadDirect?: pulumi.Input<boolean>;
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**&#47;z/*.By default no artifacts are excluded.
     */
    excludesPattern?: pulumi.Input<string>;
    /**
     * List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**&#47;z/*. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**&#47;*).
     */
    includesPattern?: pulumi.Input<string>;
    /**
     * A mandatory identifier for the repository that must be unique. Must be 3 - 10 lowercase alphanumeric and hyphen characters. It cannot begin with a number or contain spaces or special characters.
     */
    key: pulumi.Input<string>;
    /**
     * Internal description.
     */
    notes?: pulumi.Input<string>;
    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     */
    priorityResolution?: pulumi.Input<boolean>;
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     */
    projectKey?: pulumi.Input<string>;
    /**
     * List of property set name
     */
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Repository layout key for the local repository
     */
    repoLayoutRef?: pulumi.Input<string>;
    /**
     * Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via Xray settings.
     */
    xrayIndex?: pulumi.Input<boolean>;
}
