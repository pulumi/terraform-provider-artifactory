// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * # Artifactory File Data Source
 *
 * Provides an Artifactory file datasource. This can be used to download a file from a given Artifactory repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const my-file = artifactory.getFile({
 *     repository: "repo-key",
 *     path: "/path/to/the/artifact.zip",
 *     outputPath: "tmp/artifact.zip",
 * });
 * ```
 */
export function getFile(args: GetFileArgs, opts?: pulumi.InvokeOptions): Promise<GetFileResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getFile:getFile", {
        "forceOverwrite": args.forceOverwrite,
        "outputPath": args.outputPath,
        "path": args.path,
        "pathIsAliased": args.pathIsAliased,
        "repository": args.repository,
    }, opts);
}

/**
 * A collection of arguments for invoking getFile.
 */
export interface GetFileArgs {
    /**
     * If set to true, an existing file in the outputPath will be overwritten. Default: `false`
     */
    forceOverwrite?: boolean;
    /**
     * The local path the file should be downloaded to.
     */
    outputPath: string;
    /**
     * The path to the file within the repository.
     */
    path: string;
    /**
     * If set to `true`, the provider will get the artifact directly from Artifactory without attempting to resolve it or verify it and will delegate this to artifactory if the file exists. When using a smart remote repository, it is recommended to set this attribute to `true`. This is necessary to ensure that the provider fetches the artifact directly from Artifactory. If this attribute is not set or is set to `false`, there is a risk of fetching the `-cache` directory in Artifactory, potentially resulting in resource expiration and a 404 error.
     */
    pathIsAliased?: boolean;
    /**
     * Name of the repository where the file is stored.
     */
    repository: string;
}

/**
 * A collection of values returned by getFile.
 */
export interface GetFileResult {
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
    readonly forceOverwrite?: boolean;
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
    readonly outputPath: string;
    readonly path: string;
    readonly pathIsAliased?: boolean;
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
/**
 * # Artifactory File Data Source
 *
 * Provides an Artifactory file datasource. This can be used to download a file from a given Artifactory repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const my-file = artifactory.getFile({
 *     repository: "repo-key",
 *     path: "/path/to/the/artifact.zip",
 *     outputPath: "tmp/artifact.zip",
 * });
 * ```
 */
export function getFileOutput(args: GetFileOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFileResult> {
    return pulumi.output(args).apply((a: any) => getFile(a, opts))
}

/**
 * A collection of arguments for invoking getFile.
 */
export interface GetFileOutputArgs {
    /**
     * If set to true, an existing file in the outputPath will be overwritten. Default: `false`
     */
    forceOverwrite?: pulumi.Input<boolean>;
    /**
     * The local path the file should be downloaded to.
     */
    outputPath: pulumi.Input<string>;
    /**
     * The path to the file within the repository.
     */
    path: pulumi.Input<string>;
    /**
     * If set to `true`, the provider will get the artifact directly from Artifactory without attempting to resolve it or verify it and will delegate this to artifactory if the file exists. When using a smart remote repository, it is recommended to set this attribute to `true`. This is necessary to ensure that the provider fetches the artifact directly from Artifactory. If this attribute is not set or is set to `false`, there is a risk of fetching the `-cache` directory in Artifactory, potentially resulting in resource expiration and a 404 error.
     */
    pathIsAliased?: pulumi.Input<boolean>;
    /**
     * Name of the repository where the file is stored.
     */
    repository: pulumi.Input<string>;
}
