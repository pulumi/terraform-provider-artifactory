// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # Artifactory Local Opkg Repository Resource
 *
 * Creates a local opkg repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const terraform_local_test_opkg_repo = new artifactory.LocalOpkgRepository("terraform-local-test-opkg-repo", {
 *     key: "terraform-local-test-opkg-repo",
 * });
 * ```
 */
export class LocalOpkgRepository extends pulumi.CustomResource {
    /**
     * Get an existing LocalOpkgRepository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LocalOpkgRepositoryState, opts?: pulumi.CustomResourceOptions): LocalOpkgRepository {
        return new LocalOpkgRepository(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/localOpkgRepository:LocalOpkgRepository';

    /**
     * Returns true if the given object is an instance of LocalOpkgRepository.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LocalOpkgRepository {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LocalOpkgRepository.__pulumiType;
    }

    /**
     * When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
     * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
     * security (e.g., cross-site scripting attacks).
     */
    public readonly archiveBrowsingEnabled!: pulumi.Output<boolean | undefined>;
    public readonly blackedOut!: pulumi.Output<boolean | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly downloadDirect!: pulumi.Output<boolean | undefined>;
    public readonly excludesPattern!: pulumi.Output<string>;
    public readonly includesPattern!: pulumi.Output<string>;
    /**
     * - the identity key of the repo
     */
    public readonly key!: pulumi.Output<string>;
    public readonly notes!: pulumi.Output<string | undefined>;
    public /*out*/ readonly packageType!: pulumi.Output<string>;
    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     */
    public readonly priorityResolution!: pulumi.Output<boolean | undefined>;
    public readonly propertySets!: pulumi.Output<string[] | undefined>;
    public readonly repoLayoutRef!: pulumi.Output<string>;
    public readonly xrayIndex!: pulumi.Output<boolean>;

    /**
     * Create a LocalOpkgRepository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LocalOpkgRepositoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LocalOpkgRepositoryArgs | LocalOpkgRepositoryState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LocalOpkgRepositoryState | undefined;
            inputs["archiveBrowsingEnabled"] = state ? state.archiveBrowsingEnabled : undefined;
            inputs["blackedOut"] = state ? state.blackedOut : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["downloadDirect"] = state ? state.downloadDirect : undefined;
            inputs["excludesPattern"] = state ? state.excludesPattern : undefined;
            inputs["includesPattern"] = state ? state.includesPattern : undefined;
            inputs["key"] = state ? state.key : undefined;
            inputs["notes"] = state ? state.notes : undefined;
            inputs["packageType"] = state ? state.packageType : undefined;
            inputs["priorityResolution"] = state ? state.priorityResolution : undefined;
            inputs["propertySets"] = state ? state.propertySets : undefined;
            inputs["repoLayoutRef"] = state ? state.repoLayoutRef : undefined;
            inputs["xrayIndex"] = state ? state.xrayIndex : undefined;
        } else {
            const args = argsOrState as LocalOpkgRepositoryArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            inputs["archiveBrowsingEnabled"] = args ? args.archiveBrowsingEnabled : undefined;
            inputs["blackedOut"] = args ? args.blackedOut : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["downloadDirect"] = args ? args.downloadDirect : undefined;
            inputs["excludesPattern"] = args ? args.excludesPattern : undefined;
            inputs["includesPattern"] = args ? args.includesPattern : undefined;
            inputs["key"] = args ? args.key : undefined;
            inputs["notes"] = args ? args.notes : undefined;
            inputs["priorityResolution"] = args ? args.priorityResolution : undefined;
            inputs["propertySets"] = args ? args.propertySets : undefined;
            inputs["repoLayoutRef"] = args ? args.repoLayoutRef : undefined;
            inputs["xrayIndex"] = args ? args.xrayIndex : undefined;
            inputs["packageType"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(LocalOpkgRepository.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LocalOpkgRepository resources.
 */
export interface LocalOpkgRepositoryState {
    /**
     * When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
     * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
     * security (e.g., cross-site scripting attacks).
     */
    archiveBrowsingEnabled?: pulumi.Input<boolean>;
    blackedOut?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    downloadDirect?: pulumi.Input<boolean>;
    excludesPattern?: pulumi.Input<string>;
    includesPattern?: pulumi.Input<string>;
    /**
     * - the identity key of the repo
     */
    key?: pulumi.Input<string>;
    notes?: pulumi.Input<string>;
    packageType?: pulumi.Input<string>;
    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     */
    priorityResolution?: pulumi.Input<boolean>;
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    repoLayoutRef?: pulumi.Input<string>;
    xrayIndex?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a LocalOpkgRepository resource.
 */
export interface LocalOpkgRepositoryArgs {
    /**
     * When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
     * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
     * security (e.g., cross-site scripting attacks).
     */
    archiveBrowsingEnabled?: pulumi.Input<boolean>;
    blackedOut?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    downloadDirect?: pulumi.Input<boolean>;
    excludesPattern?: pulumi.Input<string>;
    includesPattern?: pulumi.Input<string>;
    /**
     * - the identity key of the repo
     */
    key: pulumi.Input<string>;
    notes?: pulumi.Input<string>;
    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     */
    priorityResolution?: pulumi.Input<boolean>;
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    repoLayoutRef?: pulumi.Input<string>;
    xrayIndex?: pulumi.Input<boolean>;
}
