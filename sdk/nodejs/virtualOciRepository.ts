// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates a virtual OCI repository.
 *
 * Official documentation can be found [here](https://jfrog.com/help/r/jfrog-artifactory-documentation/set-up-virtual-oci-repositories).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const my_oci_virtual = new artifactory.VirtualOciRepository("my-oci-virtual", {
 *     key: "my-oci-virtual",
 *     repositories: [
 *         "my-oci-local",
 *         "my-oci-remote",
 *     ],
 *     description: "A test virtual OCI repo",
 *     notes: "Internal description",
 *     resolveOciTagsByTimestamp: true,
 * });
 * ```
 *
 * ## Import
 *
 * Virtual OCI repositories can be imported using their name, e.g.
 *
 * ```sh
 * $ pulumi import artifactory:index/virtualOciRepository:VirtualOciRepository my-oci-virtual my-oci-virtual
 * ```
 */
export class VirtualOciRepository extends pulumi.CustomResource {
    /**
     * Get an existing VirtualOciRepository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualOciRepositoryState, opts?: pulumi.CustomResourceOptions): VirtualOciRepository {
        return new VirtualOciRepository(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/virtualOciRepository:VirtualOciRepository';

    /**
     * Returns true if the given object is an instance of VirtualOciRepository.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualOciRepository {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualOciRepository.__pulumiType;
    }

    /**
     * Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
     * another Artifactory instance.
     */
    public readonly artifactoryRequestsCanRetrieveRemoteArtifacts!: pulumi.Output<boolean | undefined>;
    /**
     * Default repository to deploy artifacts.
     */
    public readonly defaultDeploymentRepo!: pulumi.Output<string | undefined>;
    /**
     * Public description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**&#47;z/*.By default no
     * artifacts are excluded.
     */
    public readonly excludesPattern!: pulumi.Output<string | undefined>;
    /**
     * List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**&#47;z/*. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**&#47;*).
     */
    public readonly includesPattern!: pulumi.Output<string | undefined>;
    /**
     * A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * Internal description.
     */
    public readonly notes!: pulumi.Output<string | undefined>;
    public /*out*/ readonly packageType!: pulumi.Output<string>;
    /**
     * Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
     * Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
     * attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
     * be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
     */
    public readonly projectEnvironments!: pulumi.Output<string[]>;
    /**
     * Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     */
    public readonly projectKey!: pulumi.Output<string | undefined>;
    /**
     * Repository layout key for the virtual repository
     */
    public readonly repoLayoutRef!: pulumi.Output<string | undefined>;
    /**
     * The effective list of actual repositories included in this virtual repository.
     */
    public readonly repositories!: pulumi.Output<string[] | undefined>;
    /**
     * When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
     */
    public readonly resolveOciTagsByTimestamp!: pulumi.Output<boolean | undefined>;

    /**
     * Create a VirtualOciRepository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualOciRepositoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualOciRepositoryArgs | VirtualOciRepositoryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VirtualOciRepositoryState | undefined;
            resourceInputs["artifactoryRequestsCanRetrieveRemoteArtifacts"] = state ? state.artifactoryRequestsCanRetrieveRemoteArtifacts : undefined;
            resourceInputs["defaultDeploymentRepo"] = state ? state.defaultDeploymentRepo : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["excludesPattern"] = state ? state.excludesPattern : undefined;
            resourceInputs["includesPattern"] = state ? state.includesPattern : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["notes"] = state ? state.notes : undefined;
            resourceInputs["packageType"] = state ? state.packageType : undefined;
            resourceInputs["projectEnvironments"] = state ? state.projectEnvironments : undefined;
            resourceInputs["projectKey"] = state ? state.projectKey : undefined;
            resourceInputs["repoLayoutRef"] = state ? state.repoLayoutRef : undefined;
            resourceInputs["repositories"] = state ? state.repositories : undefined;
            resourceInputs["resolveOciTagsByTimestamp"] = state ? state.resolveOciTagsByTimestamp : undefined;
        } else {
            const args = argsOrState as VirtualOciRepositoryArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            resourceInputs["artifactoryRequestsCanRetrieveRemoteArtifacts"] = args ? args.artifactoryRequestsCanRetrieveRemoteArtifacts : undefined;
            resourceInputs["defaultDeploymentRepo"] = args ? args.defaultDeploymentRepo : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["excludesPattern"] = args ? args.excludesPattern : undefined;
            resourceInputs["includesPattern"] = args ? args.includesPattern : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["projectEnvironments"] = args ? args.projectEnvironments : undefined;
            resourceInputs["projectKey"] = args ? args.projectKey : undefined;
            resourceInputs["repoLayoutRef"] = args ? args.repoLayoutRef : undefined;
            resourceInputs["repositories"] = args ? args.repositories : undefined;
            resourceInputs["resolveOciTagsByTimestamp"] = args ? args.resolveOciTagsByTimestamp : undefined;
            resourceInputs["packageType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VirtualOciRepository.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VirtualOciRepository resources.
 */
export interface VirtualOciRepositoryState {
    /**
     * Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
     * another Artifactory instance.
     */
    artifactoryRequestsCanRetrieveRemoteArtifacts?: pulumi.Input<boolean>;
    /**
     * Default repository to deploy artifacts.
     */
    defaultDeploymentRepo?: pulumi.Input<string>;
    /**
     * Public description.
     */
    description?: pulumi.Input<string>;
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**&#47;z/*.By default no
     * artifacts are excluded.
     */
    excludesPattern?: pulumi.Input<string>;
    /**
     * List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**&#47;z/*. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**&#47;*).
     */
    includesPattern?: pulumi.Input<string>;
    /**
     * A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.
     */
    key?: pulumi.Input<string>;
    /**
     * Internal description.
     */
    notes?: pulumi.Input<string>;
    packageType?: pulumi.Input<string>;
    /**
     * Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
     * Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
     * attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
     * be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
     */
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     */
    projectKey?: pulumi.Input<string>;
    /**
     * Repository layout key for the virtual repository
     */
    repoLayoutRef?: pulumi.Input<string>;
    /**
     * The effective list of actual repositories included in this virtual repository.
     */
    repositories?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
     */
    resolveOciTagsByTimestamp?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a VirtualOciRepository resource.
 */
export interface VirtualOciRepositoryArgs {
    /**
     * Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
     * another Artifactory instance.
     */
    artifactoryRequestsCanRetrieveRemoteArtifacts?: pulumi.Input<boolean>;
    /**
     * Default repository to deploy artifacts.
     */
    defaultDeploymentRepo?: pulumi.Input<string>;
    /**
     * Public description.
     */
    description?: pulumi.Input<string>;
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**&#47;z/*.By default no
     * artifacts are excluded.
     */
    excludesPattern?: pulumi.Input<string>;
    /**
     * List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**&#47;z/*. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**&#47;*).
     */
    includesPattern?: pulumi.Input<string>;
    /**
     * A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.
     */
    key: pulumi.Input<string>;
    /**
     * Internal description.
     */
    notes?: pulumi.Input<string>;
    /**
     * Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
     * Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
     * attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
     * be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
     */
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     */
    projectKey?: pulumi.Input<string>;
    /**
     * Repository layout key for the virtual repository
     */
    repoLayoutRef?: pulumi.Input<string>;
    /**
     * The effective list of actual repositories included in this virtual repository.
     */
    repositories?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
     */
    resolveOciTagsByTimestamp?: pulumi.Input<boolean>;
}
