// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Group extends pulumi.CustomResource {
    /**
     * Get an existing Group resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupState, opts?: pulumi.CustomResourceOptions): Group {
        return new Group(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/group:Group';

    /**
     * Returns true if the given object is an instance of Group.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Group {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Group.__pulumiType;
    }

    public readonly adminPrivileges!: pulumi.Output<boolean>;
    public readonly autoJoin!: pulumi.Output<boolean>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly detachAllUsers!: pulumi.Output<boolean | undefined>;
    public readonly name!: pulumi.Output<string>;
    /**
     * (Optional) When this override is set, User in the group can set Xray security and compliance policies. Default value is
     * 'false'.
     */
    public readonly policyManager!: pulumi.Output<boolean | undefined>;
    public readonly realm!: pulumi.Output<string>;
    public readonly realmAttributes!: pulumi.Output<string | undefined>;
    /**
     * (Optional) When this override is set, User in the group can manage Xray Reports. Default value is 'false'.
     */
    public readonly reportsManager!: pulumi.Output<boolean | undefined>;
    public readonly usersNames!: pulumi.Output<string[] | undefined>;
    /**
     * (Optional) When this override is set, User in the group can manage Xray Watches on any resource type. Default value is
     * 'false'.
     */
    public readonly watchManager!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Group resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: GroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupArgs | GroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupState | undefined;
            resourceInputs["adminPrivileges"] = state ? state.adminPrivileges : undefined;
            resourceInputs["autoJoin"] = state ? state.autoJoin : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["detachAllUsers"] = state ? state.detachAllUsers : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["policyManager"] = state ? state.policyManager : undefined;
            resourceInputs["realm"] = state ? state.realm : undefined;
            resourceInputs["realmAttributes"] = state ? state.realmAttributes : undefined;
            resourceInputs["reportsManager"] = state ? state.reportsManager : undefined;
            resourceInputs["usersNames"] = state ? state.usersNames : undefined;
            resourceInputs["watchManager"] = state ? state.watchManager : undefined;
        } else {
            const args = argsOrState as GroupArgs | undefined;
            resourceInputs["adminPrivileges"] = args ? args.adminPrivileges : undefined;
            resourceInputs["autoJoin"] = args ? args.autoJoin : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["detachAllUsers"] = args ? args.detachAllUsers : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["policyManager"] = args ? args.policyManager : undefined;
            resourceInputs["realm"] = args ? args.realm : undefined;
            resourceInputs["realmAttributes"] = args ? args.realmAttributes : undefined;
            resourceInputs["reportsManager"] = args ? args.reportsManager : undefined;
            resourceInputs["usersNames"] = args ? args.usersNames : undefined;
            resourceInputs["watchManager"] = args ? args.watchManager : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Group.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Group resources.
 */
export interface GroupState {
    adminPrivileges?: pulumi.Input<boolean>;
    autoJoin?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    detachAllUsers?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    /**
     * (Optional) When this override is set, User in the group can set Xray security and compliance policies. Default value is
     * 'false'.
     */
    policyManager?: pulumi.Input<boolean>;
    realm?: pulumi.Input<string>;
    realmAttributes?: pulumi.Input<string>;
    /**
     * (Optional) When this override is set, User in the group can manage Xray Reports. Default value is 'false'.
     */
    reportsManager?: pulumi.Input<boolean>;
    usersNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional) When this override is set, User in the group can manage Xray Watches on any resource type. Default value is
     * 'false'.
     */
    watchManager?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Group resource.
 */
export interface GroupArgs {
    adminPrivileges?: pulumi.Input<boolean>;
    autoJoin?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    detachAllUsers?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    /**
     * (Optional) When this override is set, User in the group can set Xray security and compliance policies. Default value is
     * 'false'.
     */
    policyManager?: pulumi.Input<boolean>;
    realm?: pulumi.Input<string>;
    realmAttributes?: pulumi.Input<string>;
    /**
     * (Optional) When this override is set, User in the group can manage Xray Reports. Default value is 'false'.
     */
    reportsManager?: pulumi.Input<boolean>;
    usersNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional) When this override is set, User in the group can manage Xray Watches on any resource type. Default value is
     * 'false'.
     */
    watchManager?: pulumi.Input<boolean>;
}
