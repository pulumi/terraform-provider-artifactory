// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource can be used to manage Artifactory's Repository Layout settings. See [Repository Layouts](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts) in the Artifactory Wiki documentation for more details.
 *
 * ~>The `artifactory.RepositoryLayout` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const custom_layout = new artifactory.RepositoryLayout("custom-layout", {
 *     name: "custom-layout",
 *     artifactPathPattern: "[orgPath]/[module]/[baseRev](-[folderItegRev])/[module]-[baseRev](-[fileItegRev])(-[classifier]).[ext]",
 *     distinctiveDescriptorPathPattern: true,
 *     descriptorPathPattern: "[orgPath]/[module]/[baseRev](-[folderItegRev])/[module]-[baseRev](-[fileItegRev])(-[classifier]).pom",
 *     folderIntegrationRevisionRegexp: "Foo",
 *     fileIntegrationRevisionRegexp: "Foo|(?:(?:[0-9]{8}.[0-9]{6})-(?:[0-9]+))",
 * });
 * ```
 *
 * ## Import
 *
 * Repository layout can be imported using its name, e.g.
 *
 * ```sh
 * $ pulumi import artifactory:index/repositoryLayout:RepositoryLayout custom-layout custom-layout
 * ```
 */
export class RepositoryLayout extends pulumi.CustomResource {
    /**
     * Get an existing RepositoryLayout resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryLayoutState, opts?: pulumi.CustomResourceOptions): RepositoryLayout {
        return new RepositoryLayout(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/repositoryLayout:RepositoryLayout';

    /**
     * Returns true if the given object is an instance of RepositoryLayout.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RepositoryLayout {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RepositoryLayout.__pulumiType;
    }

    /**
     * Please refer to: [Path Patterns](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts#RepositoryLayouts-ModulesandPathPatternsusedbyRepositoryLayouts) in the Artifactory Wiki documentation.
     */
    public readonly artifactPathPattern!: pulumi.Output<string>;
    /**
     * Please refer to: [Descriptor Path Patterns](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts#RepositoryLayouts-DescriptorPathPatterns) in the Artifactory Wiki documentation.
     */
    public readonly descriptorPathPattern!: pulumi.Output<string | undefined>;
    /**
     * When set, `descriptorPathPattern` will be used. Default to `false`.
     */
    public readonly distinctiveDescriptorPathPattern!: pulumi.Output<boolean>;
    /**
     * A regular expression matching the integration revision string appearing in a file name as part of the artifact's path. For example, `SNAPSHOT|(?:(?:[0-9]{8}.[0-9]{6})-(?:[0-9]+))`, in Maven. Note! Take care not to introduce any regexp capturing groups within this expression. If not applicable use `.*`
     */
    public readonly fileIntegrationRevisionRegexp!: pulumi.Output<string>;
    /**
     * A regular expression matching the integration revision string appearing in a folder name as part of the artifact's path. For example, `SNAPSHOT`, in Maven. Note! Take care not to introduce any regexp capturing groups within this expression. If not applicable use `.*`
     */
    public readonly folderIntegrationRevisionRegexp!: pulumi.Output<string>;
    /**
     * Layout name
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a RepositoryLayout resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RepositoryLayoutArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepositoryLayoutArgs | RepositoryLayoutState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RepositoryLayoutState | undefined;
            resourceInputs["artifactPathPattern"] = state ? state.artifactPathPattern : undefined;
            resourceInputs["descriptorPathPattern"] = state ? state.descriptorPathPattern : undefined;
            resourceInputs["distinctiveDescriptorPathPattern"] = state ? state.distinctiveDescriptorPathPattern : undefined;
            resourceInputs["fileIntegrationRevisionRegexp"] = state ? state.fileIntegrationRevisionRegexp : undefined;
            resourceInputs["folderIntegrationRevisionRegexp"] = state ? state.folderIntegrationRevisionRegexp : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as RepositoryLayoutArgs | undefined;
            if ((!args || args.artifactPathPattern === undefined) && !opts.urn) {
                throw new Error("Missing required property 'artifactPathPattern'");
            }
            if ((!args || args.fileIntegrationRevisionRegexp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fileIntegrationRevisionRegexp'");
            }
            if ((!args || args.folderIntegrationRevisionRegexp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'folderIntegrationRevisionRegexp'");
            }
            resourceInputs["artifactPathPattern"] = args ? args.artifactPathPattern : undefined;
            resourceInputs["descriptorPathPattern"] = args ? args.descriptorPathPattern : undefined;
            resourceInputs["distinctiveDescriptorPathPattern"] = args ? args.distinctiveDescriptorPathPattern : undefined;
            resourceInputs["fileIntegrationRevisionRegexp"] = args ? args.fileIntegrationRevisionRegexp : undefined;
            resourceInputs["folderIntegrationRevisionRegexp"] = args ? args.folderIntegrationRevisionRegexp : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RepositoryLayout.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RepositoryLayout resources.
 */
export interface RepositoryLayoutState {
    /**
     * Please refer to: [Path Patterns](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts#RepositoryLayouts-ModulesandPathPatternsusedbyRepositoryLayouts) in the Artifactory Wiki documentation.
     */
    artifactPathPattern?: pulumi.Input<string>;
    /**
     * Please refer to: [Descriptor Path Patterns](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts#RepositoryLayouts-DescriptorPathPatterns) in the Artifactory Wiki documentation.
     */
    descriptorPathPattern?: pulumi.Input<string>;
    /**
     * When set, `descriptorPathPattern` will be used. Default to `false`.
     */
    distinctiveDescriptorPathPattern?: pulumi.Input<boolean>;
    /**
     * A regular expression matching the integration revision string appearing in a file name as part of the artifact's path. For example, `SNAPSHOT|(?:(?:[0-9]{8}.[0-9]{6})-(?:[0-9]+))`, in Maven. Note! Take care not to introduce any regexp capturing groups within this expression. If not applicable use `.*`
     */
    fileIntegrationRevisionRegexp?: pulumi.Input<string>;
    /**
     * A regular expression matching the integration revision string appearing in a folder name as part of the artifact's path. For example, `SNAPSHOT`, in Maven. Note! Take care not to introduce any regexp capturing groups within this expression. If not applicable use `.*`
     */
    folderIntegrationRevisionRegexp?: pulumi.Input<string>;
    /**
     * Layout name
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RepositoryLayout resource.
 */
export interface RepositoryLayoutArgs {
    /**
     * Please refer to: [Path Patterns](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts#RepositoryLayouts-ModulesandPathPatternsusedbyRepositoryLayouts) in the Artifactory Wiki documentation.
     */
    artifactPathPattern: pulumi.Input<string>;
    /**
     * Please refer to: [Descriptor Path Patterns](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts#RepositoryLayouts-DescriptorPathPatterns) in the Artifactory Wiki documentation.
     */
    descriptorPathPattern?: pulumi.Input<string>;
    /**
     * When set, `descriptorPathPattern` will be used. Default to `false`.
     */
    distinctiveDescriptorPathPattern?: pulumi.Input<boolean>;
    /**
     * A regular expression matching the integration revision string appearing in a file name as part of the artifact's path. For example, `SNAPSHOT|(?:(?:[0-9]{8}.[0-9]{6})-(?:[0-9]+))`, in Maven. Note! Take care not to introduce any regexp capturing groups within this expression. If not applicable use `.*`
     */
    fileIntegrationRevisionRegexp: pulumi.Input<string>;
    /**
     * A regular expression matching the integration revision string appearing in a folder name as part of the artifact's path. For example, `SNAPSHOT`, in Maven. Note! Take care not to introduce any regexp capturing groups within this expression. If not applicable use `.*`
     */
    folderIntegrationRevisionRegexp: pulumi.Input<string>;
    /**
     * Layout name
     */
    name?: pulumi.Input<string>;
}
