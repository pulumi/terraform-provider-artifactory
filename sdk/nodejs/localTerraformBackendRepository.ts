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
 * const terraform_local_test_terraformbackend_repo = new artifactory.LocalTerraformBackendRepository("terraform-local-test-terraformbackend-repo", {key: "terraform-local-test-terraformbackend-repo"});
 * ```
 *
 * ## Import
 *
 * Local repositories can be imported using their name, e.g.
 *
 * ```sh
 * $ pulumi import artifactory:index/localTerraformBackendRepository:LocalTerraformBackendRepository terraform-local-test-terraformbackend-repo terraform-local-test-terraformbackend-repo
 * ```
 */
export class LocalTerraformBackendRepository extends pulumi.CustomResource {
    /**
     * Get an existing LocalTerraformBackendRepository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LocalTerraformBackendRepositoryState, opts?: pulumi.CustomResourceOptions): LocalTerraformBackendRepository {
        return new LocalTerraformBackendRepository(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/localTerraformBackendRepository:LocalTerraformBackendRepository';

    /**
     * Returns true if the given object is an instance of LocalTerraformBackendRepository.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LocalTerraformBackendRepository {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LocalTerraformBackendRepository.__pulumiType;
    }

    /**
     * When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
     * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
     * security (e.g., cross-site scripting attacks).
     */
    public readonly archiveBrowsingEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
     */
    public readonly blackedOut!: pulumi.Output<boolean | undefined>;
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from AWS
     * CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
     */
    public readonly cdnRedirect!: pulumi.Output<boolean | undefined>;
    /**
     * Public description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
     * storage provider. Available in Enterprise+ and Edge licenses only.
     */
    public readonly downloadDirect!: pulumi.Output<boolean | undefined>;
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
     * the identity key of the repo.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * Internal description.
     */
    public readonly notes!: pulumi.Output<string | undefined>;
    public /*out*/ readonly packageType!: pulumi.Output<string>;
    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     */
    public readonly priorityResolution!: pulumi.Output<boolean | undefined>;
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
     * List of property set name
     */
    public readonly propertySets!: pulumi.Output<string[] | undefined>;
    /**
     * Repository layout key for the local repository
     */
    public readonly repoLayoutRef!: pulumi.Output<string | undefined>;
    /**
     * Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
     * Xray settings.
     */
    public readonly xrayIndex!: pulumi.Output<boolean | undefined>;

    /**
     * Create a LocalTerraformBackendRepository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LocalTerraformBackendRepositoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LocalTerraformBackendRepositoryArgs | LocalTerraformBackendRepositoryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LocalTerraformBackendRepositoryState | undefined;
            resourceInputs["archiveBrowsingEnabled"] = state ? state.archiveBrowsingEnabled : undefined;
            resourceInputs["blackedOut"] = state ? state.blackedOut : undefined;
            resourceInputs["cdnRedirect"] = state ? state.cdnRedirect : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["downloadDirect"] = state ? state.downloadDirect : undefined;
            resourceInputs["excludesPattern"] = state ? state.excludesPattern : undefined;
            resourceInputs["includesPattern"] = state ? state.includesPattern : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["notes"] = state ? state.notes : undefined;
            resourceInputs["packageType"] = state ? state.packageType : undefined;
            resourceInputs["priorityResolution"] = state ? state.priorityResolution : undefined;
            resourceInputs["projectEnvironments"] = state ? state.projectEnvironments : undefined;
            resourceInputs["projectKey"] = state ? state.projectKey : undefined;
            resourceInputs["propertySets"] = state ? state.propertySets : undefined;
            resourceInputs["repoLayoutRef"] = state ? state.repoLayoutRef : undefined;
            resourceInputs["xrayIndex"] = state ? state.xrayIndex : undefined;
        } else {
            const args = argsOrState as LocalTerraformBackendRepositoryArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            resourceInputs["archiveBrowsingEnabled"] = args ? args.archiveBrowsingEnabled : undefined;
            resourceInputs["blackedOut"] = args ? args.blackedOut : undefined;
            resourceInputs["cdnRedirect"] = args ? args.cdnRedirect : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["downloadDirect"] = args ? args.downloadDirect : undefined;
            resourceInputs["excludesPattern"] = args ? args.excludesPattern : undefined;
            resourceInputs["includesPattern"] = args ? args.includesPattern : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["priorityResolution"] = args ? args.priorityResolution : undefined;
            resourceInputs["projectEnvironments"] = args ? args.projectEnvironments : undefined;
            resourceInputs["projectKey"] = args ? args.projectKey : undefined;
            resourceInputs["propertySets"] = args ? args.propertySets : undefined;
            resourceInputs["repoLayoutRef"] = args ? args.repoLayoutRef : undefined;
            resourceInputs["xrayIndex"] = args ? args.xrayIndex : undefined;
            resourceInputs["packageType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LocalTerraformBackendRepository.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LocalTerraformBackendRepository resources.
 */
export interface LocalTerraformBackendRepositoryState {
    /**
     * When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
     * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
     * security (e.g., cross-site scripting attacks).
     */
    archiveBrowsingEnabled?: pulumi.Input<boolean>;
    /**
     * When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
     */
    blackedOut?: pulumi.Input<boolean>;
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from AWS
     * CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
     */
    cdnRedirect?: pulumi.Input<boolean>;
    /**
     * Public description.
     */
    description?: pulumi.Input<string>;
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
     * storage provider. Available in Enterprise+ and Edge licenses only.
     */
    downloadDirect?: pulumi.Input<boolean>;
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
     * the identity key of the repo.
     */
    key?: pulumi.Input<string>;
    /**
     * Internal description.
     */
    notes?: pulumi.Input<string>;
    packageType?: pulumi.Input<string>;
    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     */
    priorityResolution?: pulumi.Input<boolean>;
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
     * List of property set name
     */
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Repository layout key for the local repository
     */
    repoLayoutRef?: pulumi.Input<string>;
    /**
     * Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
     * Xray settings.
     */
    xrayIndex?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a LocalTerraformBackendRepository resource.
 */
export interface LocalTerraformBackendRepositoryArgs {
    /**
     * When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
     * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
     * security (e.g., cross-site scripting attacks).
     */
    archiveBrowsingEnabled?: pulumi.Input<boolean>;
    /**
     * When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
     */
    blackedOut?: pulumi.Input<boolean>;
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from AWS
     * CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
     */
    cdnRedirect?: pulumi.Input<boolean>;
    /**
     * Public description.
     */
    description?: pulumi.Input<string>;
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
     * storage provider. Available in Enterprise+ and Edge licenses only.
     */
    downloadDirect?: pulumi.Input<boolean>;
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
     * the identity key of the repo.
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
     * List of property set name
     */
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Repository layout key for the local repository
     */
    repoLayoutRef?: pulumi.Input<string>;
    /**
     * Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
     * Xray settings.
     */
    xrayIndex?: pulumi.Input<boolean>;
}
