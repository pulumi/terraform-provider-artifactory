// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class FederatedGenericRepository extends pulumi.CustomResource {
    /**
     * Get an existing FederatedGenericRepository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FederatedGenericRepositoryState, opts?: pulumi.CustomResourceOptions): FederatedGenericRepository {
        return new FederatedGenericRepository(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/federatedGenericRepository:FederatedGenericRepository';

    /**
     * Returns true if the given object is an instance of FederatedGenericRepository.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FederatedGenericRepository {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FederatedGenericRepository.__pulumiType;
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
    public readonly key!: pulumi.Output<string>;
    /**
     * The list of Federated members. If a Federated member receives a request that does not include the repository URL, it
     * will automatically be added with the combination of the configured base URL and `key` field value. Note that each of the
     * federated members will need to have a base URL set. Please follow the
     * [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
     * to set up Federated repositories correctly.
     */
    public readonly members!: pulumi.Output<outputs.FederatedGenericRepositoryMember[]>;
    public readonly notes!: pulumi.Output<string | undefined>;
    public /*out*/ readonly packageType!: pulumi.Output<string>;
    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     */
    public readonly priorityResolution!: pulumi.Output<boolean | undefined>;
    /**
     * Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
     */
    public readonly projectEnvironments!: pulumi.Output<string[] | undefined>;
    /**
     * Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
     * with project key, separated by a dash.
     */
    public readonly projectKey!: pulumi.Output<string | undefined>;
    public readonly propertySets!: pulumi.Output<string[] | undefined>;
    /**
     * Repository layout key for the local repository
     */
    public readonly repoLayoutRef!: pulumi.Output<string | undefined>;
    public readonly xrayIndex!: pulumi.Output<boolean>;

    /**
     * Create a FederatedGenericRepository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FederatedGenericRepositoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FederatedGenericRepositoryArgs | FederatedGenericRepositoryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FederatedGenericRepositoryState | undefined;
            resourceInputs["archiveBrowsingEnabled"] = state ? state.archiveBrowsingEnabled : undefined;
            resourceInputs["blackedOut"] = state ? state.blackedOut : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["downloadDirect"] = state ? state.downloadDirect : undefined;
            resourceInputs["excludesPattern"] = state ? state.excludesPattern : undefined;
            resourceInputs["includesPattern"] = state ? state.includesPattern : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["notes"] = state ? state.notes : undefined;
            resourceInputs["packageType"] = state ? state.packageType : undefined;
            resourceInputs["priorityResolution"] = state ? state.priorityResolution : undefined;
            resourceInputs["projectEnvironments"] = state ? state.projectEnvironments : undefined;
            resourceInputs["projectKey"] = state ? state.projectKey : undefined;
            resourceInputs["propertySets"] = state ? state.propertySets : undefined;
            resourceInputs["repoLayoutRef"] = state ? state.repoLayoutRef : undefined;
            resourceInputs["xrayIndex"] = state ? state.xrayIndex : undefined;
        } else {
            const args = argsOrState as FederatedGenericRepositoryArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            resourceInputs["archiveBrowsingEnabled"] = args ? args.archiveBrowsingEnabled : undefined;
            resourceInputs["blackedOut"] = args ? args.blackedOut : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["downloadDirect"] = args ? args.downloadDirect : undefined;
            resourceInputs["excludesPattern"] = args ? args.excludesPattern : undefined;
            resourceInputs["includesPattern"] = args ? args.includesPattern : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
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
        super(FederatedGenericRepository.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FederatedGenericRepository resources.
 */
export interface FederatedGenericRepositoryState {
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
    key?: pulumi.Input<string>;
    /**
     * The list of Federated members. If a Federated member receives a request that does not include the repository URL, it
     * will automatically be added with the combination of the configured base URL and `key` field value. Note that each of the
     * federated members will need to have a base URL set. Please follow the
     * [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
     * to set up Federated repositories correctly.
     */
    members?: pulumi.Input<pulumi.Input<inputs.FederatedGenericRepositoryMember>[]>;
    notes?: pulumi.Input<string>;
    packageType?: pulumi.Input<string>;
    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     */
    priorityResolution?: pulumi.Input<boolean>;
    /**
     * Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
     */
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
     * with project key, separated by a dash.
     */
    projectKey?: pulumi.Input<string>;
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Repository layout key for the local repository
     */
    repoLayoutRef?: pulumi.Input<string>;
    xrayIndex?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a FederatedGenericRepository resource.
 */
export interface FederatedGenericRepositoryArgs {
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
    key: pulumi.Input<string>;
    /**
     * The list of Federated members. If a Federated member receives a request that does not include the repository URL, it
     * will automatically be added with the combination of the configured base URL and `key` field value. Note that each of the
     * federated members will need to have a base URL set. Please follow the
     * [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
     * to set up Federated repositories correctly.
     */
    members: pulumi.Input<pulumi.Input<inputs.FederatedGenericRepositoryMember>[]>;
    notes?: pulumi.Input<string>;
    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     */
    priorityResolution?: pulumi.Input<boolean>;
    /**
     * Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
     */
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
     * with project key, separated by a dash.
     */
    projectKey?: pulumi.Input<string>;
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Repository layout key for the local repository
     */
    repoLayoutRef?: pulumi.Input<string>;
    xrayIndex?: pulumi.Input<boolean>;
}
