// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # Artifactory File Info Data Source
 *
 * Provides an Artifactory fileinfo datasource. This can be used to read metadata of files stored in Artifactory repositories.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * // 
 * const my_file = pulumi.output(artifactory.getFileinfo({
 *     path: "/path/to/the/artifact.zip",
 *     repository: "repo-key",
 * }));
 * ```
 */
export function getFileinfo(args: GetFileinfoArgs, opts?: pulumi.InvokeOptions): Promise<GetFileinfoResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("artifactory:index/getFileinfo:getFileinfo", {
        "path": args.path,
        "repository": args.repository,
    }, opts);
}

/**
 * A collection of arguments for invoking getFileinfo.
 */
export interface GetFileinfoArgs {
    /**
     * The path to the file within the repository.
     */
    path: string;
    /**
     * Name of the repository where the file is stored.
     */
    repository: string;
}

/**
 * A collection of values returned by getFileinfo.
 */
export interface GetFileinfoResult {
    /**
     * The time & date when the file was created.
     */
    readonly created: string;
    /**
     * The user who created the file.
     */
    readonly createdBy: string;
    /**
     * The URI that can be used to download the file.
     */
    readonly downloadUri: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The time & date when the file was last modified.
     */
    readonly lastModified: string;
    /**
     * The time & date when the file was last updated.
     */
    readonly lastUpdated: string;
    /**
     * MD5 checksum of the file.
     */
    readonly md5: string;
    /**
     * The mimetype of the file.
     */
    readonly mimetype: string;
    /**
     * The user who last modified the file.
     */
    readonly modifiedBy: string;
    readonly path: string;
    readonly repository: string;
    /**
     * SHA1 checksum of the file.
     */
    readonly sha1: string;
    /**
     * SHA256 checksum of the file.
     */
    readonly sha256: string;
    /**
     * The size of the file.
     */
    readonly size: number;
}

export function getFileinfoOutput(args: GetFileinfoOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFileinfoResult> {
    return pulumi.output(args).apply(a => getFileinfo(a, opts))
}

/**
 * A collection of arguments for invoking getFileinfo.
 */
export interface GetFileinfoOutputArgs {
    /**
     * The path to the file within the repository.
     */
    path: pulumi.Input<string>;
    /**
     * Name of the repository where the file is stored.
     */
    repository: pulumi.Input<string>;
}
