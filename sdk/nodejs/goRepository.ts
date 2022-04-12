// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class GoRepository extends pulumi.CustomResource {
    /**
     * Get an existing GoRepository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GoRepositoryState, opts?: pulumi.CustomResourceOptions): GoRepository {
        return new GoRepository(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/goRepository:GoRepository';

    /**
     * Returns true if the given object is an instance of GoRepository.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GoRepository {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GoRepository.__pulumiType;
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
     * A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
     * field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**&#47;z/*.By default no
     * artifacts are excluded.
     */
    public readonly excludesPattern!: pulumi.Output<string | undefined>;
    /**
     * When set (default), Artifactory will automatically follow remote VCS roots in 'go-import' meta tags to download remote
     * modules.
     */
    public readonly externalDependenciesEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * An allow list of Ant-style path patterns that determine which remote VCS roots Artifactory will follow to download
     * remote modules from, when presented with 'go-import' meta tags in the remote repository response. By default, this is
     * set to '**', which means that remote modules may be downloaded from any external VCS source.
     */
    public readonly externalDependenciesPatterns!: pulumi.Output<string[] | undefined>;
    /**
     * List of artifact patterns to include when evaluating artifact requests in the form of x/y/**&#47;z/*. When used, only
     * artifacts matching one of the include patterns are served. By default, all artifacts are included (**&#47;*).
     */
    public readonly includesPattern!: pulumi.Output<string | undefined>;
    /**
     * The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
     * contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
     * 'libs-release-local').
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * A free text field to add additional notes about the repository. These are only visible to the administrator.
     */
    public readonly notes!: pulumi.Output<string | undefined>;
    /**
     * The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
     */
    public /*out*/ readonly packageType!: pulumi.Output<string>;
    /**
     * Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
     */
    public readonly projectEnvironments!: pulumi.Output<string[] | undefined>;
    /**
     * Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
     * repository to a project, repository key must be prefixed with project key, separated by a dash.
     */
    public readonly projectKey!: pulumi.Output<string | undefined>;
    /**
     * Repository layout key for the local repository
     */
    public readonly repoLayoutRef!: pulumi.Output<string | undefined>;
    /**
     * The effective list of actual repositories included in this virtual repository.
     */
    public readonly repositories!: pulumi.Output<string[] | undefined>;
    /**
     * This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
     * repositories. A value of 0 indicates no caching.
     */
    public readonly retrievalCachePeriodSeconds!: pulumi.Output<number | undefined>;

    /**
     * Create a GoRepository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GoRepositoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GoRepositoryArgs | GoRepositoryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GoRepositoryState | undefined;
            resourceInputs["artifactoryRequestsCanRetrieveRemoteArtifacts"] = state ? state.artifactoryRequestsCanRetrieveRemoteArtifacts : undefined;
            resourceInputs["defaultDeploymentRepo"] = state ? state.defaultDeploymentRepo : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["excludesPattern"] = state ? state.excludesPattern : undefined;
            resourceInputs["externalDependenciesEnabled"] = state ? state.externalDependenciesEnabled : undefined;
            resourceInputs["externalDependenciesPatterns"] = state ? state.externalDependenciesPatterns : undefined;
            resourceInputs["includesPattern"] = state ? state.includesPattern : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["notes"] = state ? state.notes : undefined;
            resourceInputs["packageType"] = state ? state.packageType : undefined;
            resourceInputs["projectEnvironments"] = state ? state.projectEnvironments : undefined;
            resourceInputs["projectKey"] = state ? state.projectKey : undefined;
            resourceInputs["repoLayoutRef"] = state ? state.repoLayoutRef : undefined;
            resourceInputs["repositories"] = state ? state.repositories : undefined;
            resourceInputs["retrievalCachePeriodSeconds"] = state ? state.retrievalCachePeriodSeconds : undefined;
        } else {
            const args = argsOrState as GoRepositoryArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            resourceInputs["artifactoryRequestsCanRetrieveRemoteArtifacts"] = args ? args.artifactoryRequestsCanRetrieveRemoteArtifacts : undefined;
            resourceInputs["defaultDeploymentRepo"] = args ? args.defaultDeploymentRepo : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["excludesPattern"] = args ? args.excludesPattern : undefined;
            resourceInputs["externalDependenciesEnabled"] = args ? args.externalDependenciesEnabled : undefined;
            resourceInputs["externalDependenciesPatterns"] = args ? args.externalDependenciesPatterns : undefined;
            resourceInputs["includesPattern"] = args ? args.includesPattern : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["projectEnvironments"] = args ? args.projectEnvironments : undefined;
            resourceInputs["projectKey"] = args ? args.projectKey : undefined;
            resourceInputs["repoLayoutRef"] = args ? args.repoLayoutRef : undefined;
            resourceInputs["repositories"] = args ? args.repositories : undefined;
            resourceInputs["retrievalCachePeriodSeconds"] = args ? args.retrievalCachePeriodSeconds : undefined;
            resourceInputs["packageType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GoRepository.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GoRepository resources.
 */
export interface GoRepositoryState {
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
     * A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
     * field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
     */
    description?: pulumi.Input<string>;
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**&#47;z/*.By default no
     * artifacts are excluded.
     */
    excludesPattern?: pulumi.Input<string>;
    /**
     * When set (default), Artifactory will automatically follow remote VCS roots in 'go-import' meta tags to download remote
     * modules.
     */
    externalDependenciesEnabled?: pulumi.Input<boolean>;
    /**
     * An allow list of Ant-style path patterns that determine which remote VCS roots Artifactory will follow to download
     * remote modules from, when presented with 'go-import' meta tags in the remote repository response. By default, this is
     * set to '**', which means that remote modules may be downloaded from any external VCS source.
     */
    externalDependenciesPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of artifact patterns to include when evaluating artifact requests in the form of x/y/**&#47;z/*. When used, only
     * artifacts matching one of the include patterns are served. By default, all artifacts are included (**&#47;*).
     */
    includesPattern?: pulumi.Input<string>;
    /**
     * The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
     * contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
     * 'libs-release-local').
     */
    key?: pulumi.Input<string>;
    /**
     * A free text field to add additional notes about the repository. These are only visible to the administrator.
     */
    notes?: pulumi.Input<string>;
    /**
     * The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
     */
    packageType?: pulumi.Input<string>;
    /**
     * Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
     */
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
     * repository to a project, repository key must be prefixed with project key, separated by a dash.
     */
    projectKey?: pulumi.Input<string>;
    /**
     * Repository layout key for the local repository
     */
    repoLayoutRef?: pulumi.Input<string>;
    /**
     * The effective list of actual repositories included in this virtual repository.
     */
    repositories?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
     * repositories. A value of 0 indicates no caching.
     */
    retrievalCachePeriodSeconds?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a GoRepository resource.
 */
export interface GoRepositoryArgs {
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
     * A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
     * field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
     */
    description?: pulumi.Input<string>;
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**&#47;z/*.By default no
     * artifacts are excluded.
     */
    excludesPattern?: pulumi.Input<string>;
    /**
     * When set (default), Artifactory will automatically follow remote VCS roots in 'go-import' meta tags to download remote
     * modules.
     */
    externalDependenciesEnabled?: pulumi.Input<boolean>;
    /**
     * An allow list of Ant-style path patterns that determine which remote VCS roots Artifactory will follow to download
     * remote modules from, when presented with 'go-import' meta tags in the remote repository response. By default, this is
     * set to '**', which means that remote modules may be downloaded from any external VCS source.
     */
    externalDependenciesPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of artifact patterns to include when evaluating artifact requests in the form of x/y/**&#47;z/*. When used, only
     * artifacts matching one of the include patterns are served. By default, all artifacts are included (**&#47;*).
     */
    includesPattern?: pulumi.Input<string>;
    /**
     * The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
     * contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
     * 'libs-release-local').
     */
    key: pulumi.Input<string>;
    /**
     * A free text field to add additional notes about the repository. These are only visible to the administrator.
     */
    notes?: pulumi.Input<string>;
    /**
     * Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
     */
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
     * repository to a project, repository key must be prefixed with project key, separated by a dash.
     */
    projectKey?: pulumi.Input<string>;
    /**
     * Repository layout key for the local repository
     */
    repoLayoutRef?: pulumi.Input<string>;
    /**
     * The effective list of actual repositories included in this virtual repository.
     */
    repositories?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
     * repositories. A value of 0 indicates no caching.
     */
    retrievalCachePeriodSeconds?: pulumi.Input<number>;
}
