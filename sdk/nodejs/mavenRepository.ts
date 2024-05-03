// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates a virtual Maven repository.
 * Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Maven+Repository).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const bar = new artifactory.LocalMavenRepository("bar", {
 *     key: "bar",
 *     repoLayoutRef: "maven-2-default",
 * });
 * const baz = new artifactory.RemoteMavenRepository("baz", {
 *     key: "baz",
 *     url: "https://search.maven.com/",
 *     repoLayoutRef: "maven-2-default",
 * });
 * const maven_virt_repo = new artifactory.MavenRepository("maven-virt-repo", {
 *     key: "maven-virt-repo",
 *     repoLayoutRef: "maven-2-default",
 *     repositories: [
 *         bar.key,
 *         baz.key,
 *     ],
 *     description: "A test virtual repo",
 *     notes: "Internal description",
 *     includesPattern: "com/jfrog/**,cloud/jfrog/**",
 *     excludesPattern: "com/google/**",
 *     forceMavenAuthentication: true,
 *     pomRepositoryReferencesCleanupPolicy: "discard_active_reference",
 * });
 * ```
 *
 * ## Import
 *
 * Virtual repositories can be imported using their name, e.g.
 *
 * ```sh
 * $ pulumi import artifactory:index/mavenRepository:MavenRepository maven-virt-repo maven-virt-repo
 * ```
 */
export class MavenRepository extends pulumi.CustomResource {
    /**
     * Get an existing MavenRepository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MavenRepositoryState, opts?: pulumi.CustomResourceOptions): MavenRepository {
        return new MavenRepository(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/mavenRepository:MavenRepository';

    /**
     * Returns true if the given object is an instance of MavenRepository.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MavenRepository {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MavenRepository.__pulumiType;
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
     * Forces authentication when fetching from remote repos.
     */
    public readonly forceMavenAuthentication!: pulumi.Output<boolean>;
    /**
     * List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**&#47;z/*. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**&#47;*).
     */
    public readonly includesPattern!: pulumi.Output<string | undefined>;
    /**
     * A mandatory identifier for the repository that must be unique. It cannot begin with a number or
     * contain spaces or special characters.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * The keypair used to sign artifacts
     */
    public readonly keyPair!: pulumi.Output<string | undefined>;
    /**
     * Internal description.
     */
    public readonly notes!: pulumi.Output<string | undefined>;
    public /*out*/ readonly packageType!: pulumi.Output<string>;
    /**
     * One of: `"discardActiveReference", "discardAnyReference", "nothing"`
     */
    public readonly pomRepositoryReferencesCleanupPolicy!: pulumi.Output<string>;
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
     * Create a MavenRepository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MavenRepositoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MavenRepositoryArgs | MavenRepositoryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MavenRepositoryState | undefined;
            resourceInputs["artifactoryRequestsCanRetrieveRemoteArtifacts"] = state ? state.artifactoryRequestsCanRetrieveRemoteArtifacts : undefined;
            resourceInputs["defaultDeploymentRepo"] = state ? state.defaultDeploymentRepo : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["excludesPattern"] = state ? state.excludesPattern : undefined;
            resourceInputs["forceMavenAuthentication"] = state ? state.forceMavenAuthentication : undefined;
            resourceInputs["includesPattern"] = state ? state.includesPattern : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["keyPair"] = state ? state.keyPair : undefined;
            resourceInputs["notes"] = state ? state.notes : undefined;
            resourceInputs["packageType"] = state ? state.packageType : undefined;
            resourceInputs["pomRepositoryReferencesCleanupPolicy"] = state ? state.pomRepositoryReferencesCleanupPolicy : undefined;
            resourceInputs["projectEnvironments"] = state ? state.projectEnvironments : undefined;
            resourceInputs["projectKey"] = state ? state.projectKey : undefined;
            resourceInputs["repoLayoutRef"] = state ? state.repoLayoutRef : undefined;
            resourceInputs["repositories"] = state ? state.repositories : undefined;
        } else {
            const args = argsOrState as MavenRepositoryArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            resourceInputs["artifactoryRequestsCanRetrieveRemoteArtifacts"] = args ? args.artifactoryRequestsCanRetrieveRemoteArtifacts : undefined;
            resourceInputs["defaultDeploymentRepo"] = args ? args.defaultDeploymentRepo : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["excludesPattern"] = args ? args.excludesPattern : undefined;
            resourceInputs["forceMavenAuthentication"] = args ? args.forceMavenAuthentication : undefined;
            resourceInputs["includesPattern"] = args ? args.includesPattern : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["keyPair"] = args ? args.keyPair : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["pomRepositoryReferencesCleanupPolicy"] = args ? args.pomRepositoryReferencesCleanupPolicy : undefined;
            resourceInputs["projectEnvironments"] = args ? args.projectEnvironments : undefined;
            resourceInputs["projectKey"] = args ? args.projectKey : undefined;
            resourceInputs["repoLayoutRef"] = args ? args.repoLayoutRef : undefined;
            resourceInputs["repositories"] = args ? args.repositories : undefined;
            resourceInputs["packageType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MavenRepository.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MavenRepository resources.
 */
export interface MavenRepositoryState {
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
     * Forces authentication when fetching from remote repos.
     */
    forceMavenAuthentication?: pulumi.Input<boolean>;
    /**
     * List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**&#47;z/*. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**&#47;*).
     */
    includesPattern?: pulumi.Input<string>;
    /**
     * A mandatory identifier for the repository that must be unique. It cannot begin with a number or
     * contain spaces or special characters.
     */
    key?: pulumi.Input<string>;
    /**
     * The keypair used to sign artifacts
     */
    keyPair?: pulumi.Input<string>;
    /**
     * Internal description.
     */
    notes?: pulumi.Input<string>;
    packageType?: pulumi.Input<string>;
    /**
     * One of: `"discardActiveReference", "discardAnyReference", "nothing"`
     */
    pomRepositoryReferencesCleanupPolicy?: pulumi.Input<string>;
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
}

/**
 * The set of arguments for constructing a MavenRepository resource.
 */
export interface MavenRepositoryArgs {
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
     * Forces authentication when fetching from remote repos.
     */
    forceMavenAuthentication?: pulumi.Input<boolean>;
    /**
     * List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**&#47;z/*. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**&#47;*).
     */
    includesPattern?: pulumi.Input<string>;
    /**
     * A mandatory identifier for the repository that must be unique. It cannot begin with a number or
     * contain spaces or special characters.
     */
    key: pulumi.Input<string>;
    /**
     * The keypair used to sign artifacts
     */
    keyPair?: pulumi.Input<string>;
    /**
     * Internal description.
     */
    notes?: pulumi.Input<string>;
    /**
     * One of: `"discardActiveReference", "discardAnyReference", "nothing"`
     */
    pomRepositoryReferencesCleanupPolicy?: pulumi.Input<string>;
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
}
