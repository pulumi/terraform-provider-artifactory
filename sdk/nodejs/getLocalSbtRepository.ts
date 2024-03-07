// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Retrieves a local Sbt repository.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const local-test-sbt-repo = artifactory.getLocalSbtRepository({
 *     key: "local-test-sbt-repo",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getLocalSbtRepository(args: GetLocalSbtRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetLocalSbtRepositoryResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getLocalSbtRepository:getLocalSbtRepository", {
        "archiveBrowsingEnabled": args.archiveBrowsingEnabled,
        "blackedOut": args.blackedOut,
        "cdnRedirect": args.cdnRedirect,
        "checksumPolicyType": args.checksumPolicyType,
        "description": args.description,
        "downloadDirect": args.downloadDirect,
        "excludesPattern": args.excludesPattern,
        "handleReleases": args.handleReleases,
        "handleSnapshots": args.handleSnapshots,
        "includesPattern": args.includesPattern,
        "key": args.key,
        "maxUniqueSnapshots": args.maxUniqueSnapshots,
        "notes": args.notes,
        "priorityResolution": args.priorityResolution,
        "projectEnvironments": args.projectEnvironments,
        "projectKey": args.projectKey,
        "propertySets": args.propertySets,
        "repoLayoutRef": args.repoLayoutRef,
        "snapshotVersionBehavior": args.snapshotVersionBehavior,
        "suppressPomConsistencyChecks": args.suppressPomConsistencyChecks,
        "xrayIndex": args.xrayIndex,
    }, opts);
}

/**
 * A collection of arguments for invoking getLocalSbtRepository.
 */
export interface GetLocalSbtRepositoryArgs {
    archiveBrowsingEnabled?: boolean;
    blackedOut?: boolean;
    cdnRedirect?: boolean;
    /**
     * Checksum policy determines how Artifactory behaves when a client checksum for a deployed
     * resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
     * `client-checksums` and `generated-checksums`. For more details, please refer
     * to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
     * .
     */
    checksumPolicyType?: string;
    description?: string;
    downloadDirect?: boolean;
    excludesPattern?: string;
    /**
     * If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`
     * .
     */
    handleReleases?: boolean;
    /**
     * If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default
     * is `true`.
     */
    handleSnapshots?: boolean;
    includesPattern?: string;
    /**
     * the identity key of the repo.
     */
    key: string;
    /**
     * The maximum number of unique snapshots of a single artifact to store. Once the number of
     * snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and
     * unique snapshots are not cleaned up.
     */
    maxUniqueSnapshots?: number;
    notes?: string;
    priorityResolution?: boolean;
    projectEnvironments?: string[];
    projectKey?: string;
    propertySets?: string[];
    repoLayoutRef?: string;
    /**
     * Specifies the naming convention for Maven SNAPSHOT versions. The options are
     * -
     */
    snapshotVersionBehavior?: string;
    /**
     * By default, Artifactory keeps your repositories healthy by refusing POMs with
     * incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed
     * path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting the
     * Suppress POM Consistency Checks checkbox. True by default for Gradle repository.
     */
    suppressPomConsistencyChecks?: boolean;
    xrayIndex?: boolean;
}

/**
 * A collection of values returned by getLocalSbtRepository.
 */
export interface GetLocalSbtRepositoryResult {
    readonly archiveBrowsingEnabled?: boolean;
    readonly blackedOut?: boolean;
    readonly cdnRedirect?: boolean;
    /**
     * Checksum policy determines how Artifactory behaves when a client checksum for a deployed
     * resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
     * `client-checksums` and `generated-checksums`. For more details, please refer
     * to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
     * .
     */
    readonly checksumPolicyType?: string;
    readonly description?: string;
    readonly downloadDirect?: boolean;
    readonly excludesPattern?: string;
    /**
     * If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`
     * .
     */
    readonly handleReleases?: boolean;
    /**
     * If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default
     * is `true`.
     */
    readonly handleSnapshots?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includesPattern?: string;
    readonly key: string;
    /**
     * The maximum number of unique snapshots of a single artifact to store. Once the number of
     * snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and
     * unique snapshots are not cleaned up.
     */
    readonly maxUniqueSnapshots?: number;
    readonly notes?: string;
    readonly packageType: string;
    readonly priorityResolution?: boolean;
    readonly projectEnvironments: string[];
    readonly projectKey?: string;
    readonly propertySets?: string[];
    readonly repoLayoutRef?: string;
    /**
     * Specifies the naming convention for Maven SNAPSHOT versions. The options are
     * -
     */
    readonly snapshotVersionBehavior?: string;
    /**
     * By default, Artifactory keeps your repositories healthy by refusing POMs with
     * incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed
     * path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting the
     * Suppress POM Consistency Checks checkbox. True by default for Gradle repository.
     */
    readonly suppressPomConsistencyChecks?: boolean;
    readonly xrayIndex?: boolean;
}
/**
 * Retrieves a local Sbt repository.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const local-test-sbt-repo = artifactory.getLocalSbtRepository({
 *     key: "local-test-sbt-repo",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getLocalSbtRepositoryOutput(args: GetLocalSbtRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLocalSbtRepositoryResult> {
    return pulumi.output(args).apply((a: any) => getLocalSbtRepository(a, opts))
}

/**
 * A collection of arguments for invoking getLocalSbtRepository.
 */
export interface GetLocalSbtRepositoryOutputArgs {
    archiveBrowsingEnabled?: pulumi.Input<boolean>;
    blackedOut?: pulumi.Input<boolean>;
    cdnRedirect?: pulumi.Input<boolean>;
    /**
     * Checksum policy determines how Artifactory behaves when a client checksum for a deployed
     * resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
     * `client-checksums` and `generated-checksums`. For more details, please refer
     * to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
     * .
     */
    checksumPolicyType?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    downloadDirect?: pulumi.Input<boolean>;
    excludesPattern?: pulumi.Input<string>;
    /**
     * If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`
     * .
     */
    handleReleases?: pulumi.Input<boolean>;
    /**
     * If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default
     * is `true`.
     */
    handleSnapshots?: pulumi.Input<boolean>;
    includesPattern?: pulumi.Input<string>;
    /**
     * the identity key of the repo.
     */
    key: pulumi.Input<string>;
    /**
     * The maximum number of unique snapshots of a single artifact to store. Once the number of
     * snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and
     * unique snapshots are not cleaned up.
     */
    maxUniqueSnapshots?: pulumi.Input<number>;
    notes?: pulumi.Input<string>;
    priorityResolution?: pulumi.Input<boolean>;
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    projectKey?: pulumi.Input<string>;
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    repoLayoutRef?: pulumi.Input<string>;
    /**
     * Specifies the naming convention for Maven SNAPSHOT versions. The options are
     * -
     */
    snapshotVersionBehavior?: pulumi.Input<string>;
    /**
     * By default, Artifactory keeps your repositories healthy by refusing POMs with
     * incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed
     * path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting the
     * Suppress POM Consistency Checks checkbox. True by default for Gradle repository.
     */
    suppressPomConsistencyChecks?: pulumi.Input<boolean>;
    xrayIndex?: pulumi.Input<boolean>;
}
