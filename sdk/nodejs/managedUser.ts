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
 * const test_user = new artifactory.ManagedUser("test-user", {
 *     email: "test-user@artifactory-terraform.com",
 *     groups: [
 *         "readers",
 *         "logged-in-users",
 *     ],
 *     password: "my super secret password",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import artifactory:index/managedUser:ManagedUser test-user myusername
 * ```
 */
export class ManagedUser extends pulumi.CustomResource {
    /**
     * Get an existing ManagedUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ManagedUserState, opts?: pulumi.CustomResourceOptions): ManagedUser {
        return new ManagedUser(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/managedUser:ManagedUser';

    /**
     * Returns true if the given object is an instance of ManagedUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagedUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagedUser.__pulumiType;
    }

    /**
     * (Optional, Default: false) When enabled, this user is an administrator with all the ensuing privileges.
     */
    public readonly admin!: pulumi.Output<boolean>;
    /**
     * (Optional, Default: true) When enabled, this user can only access the system through the REST API. This option cannot be set if the user has Admin privileges.
     */
    public readonly disableUiAccess!: pulumi.Output<boolean>;
    /**
     * Email for user.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * List of groups this user is a part of. If no groups set, `readers` group will be added by default. If other groups are assigned, `readers` must be added to the list manually to avoid state drift.
     */
    public readonly groups!: pulumi.Output<string[]>;
    /**
     * (Optional, Default: false) When enabled, disables the fallback mechanism for using an internal password when external authentication (such as LDAP) is enabled.
     */
    public readonly internalPasswordDisabled!: pulumi.Output<boolean>;
    /**
     * Username for user. May contain lowercase letters, numbers and symbols: '.-_@'
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * (Optional, Sensitive) Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * (Optional, Default: true) When enabled, this user can update their profile details (except for the password. Only an administrator can update the password). There may be cases in which you want to leave this unset to prevent users from updating their profile. For example, a departmental user with a single password shared between all department members.
     */
    public readonly profileUpdatable!: pulumi.Output<boolean>;

    /**
     * Create a ManagedUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagedUserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ManagedUserArgs | ManagedUserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ManagedUserState | undefined;
            resourceInputs["admin"] = state ? state.admin : undefined;
            resourceInputs["disableUiAccess"] = state ? state.disableUiAccess : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["groups"] = state ? state.groups : undefined;
            resourceInputs["internalPasswordDisabled"] = state ? state.internalPasswordDisabled : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["profileUpdatable"] = state ? state.profileUpdatable : undefined;
        } else {
            const args = argsOrState as ManagedUserArgs | undefined;
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            resourceInputs["admin"] = args ? args.admin : undefined;
            resourceInputs["disableUiAccess"] = args ? args.disableUiAccess : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["groups"] = args ? args.groups : undefined;
            resourceInputs["internalPasswordDisabled"] = args ? args.internalPasswordDisabled : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["profileUpdatable"] = args ? args.profileUpdatable : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ManagedUser.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ManagedUser resources.
 */
export interface ManagedUserState {
    /**
     * (Optional, Default: false) When enabled, this user is an administrator with all the ensuing privileges.
     */
    admin?: pulumi.Input<boolean>;
    /**
     * (Optional, Default: true) When enabled, this user can only access the system through the REST API. This option cannot be set if the user has Admin privileges.
     */
    disableUiAccess?: pulumi.Input<boolean>;
    /**
     * Email for user.
     */
    email?: pulumi.Input<string>;
    /**
     * List of groups this user is a part of. If no groups set, `readers` group will be added by default. If other groups are assigned, `readers` must be added to the list manually to avoid state drift.
     */
    groups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional, Default: false) When enabled, disables the fallback mechanism for using an internal password when external authentication (such as LDAP) is enabled.
     */
    internalPasswordDisabled?: pulumi.Input<boolean>;
    /**
     * Username for user. May contain lowercase letters, numbers and symbols: '.-_@'
     */
    name?: pulumi.Input<string>;
    /**
     * (Optional, Sensitive) Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters
     */
    password?: pulumi.Input<string>;
    /**
     * (Optional, Default: true) When enabled, this user can update their profile details (except for the password. Only an administrator can update the password). There may be cases in which you want to leave this unset to prevent users from updating their profile. For example, a departmental user with a single password shared between all department members.
     */
    profileUpdatable?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ManagedUser resource.
 */
export interface ManagedUserArgs {
    /**
     * (Optional, Default: false) When enabled, this user is an administrator with all the ensuing privileges.
     */
    admin?: pulumi.Input<boolean>;
    /**
     * (Optional, Default: true) When enabled, this user can only access the system through the REST API. This option cannot be set if the user has Admin privileges.
     */
    disableUiAccess?: pulumi.Input<boolean>;
    /**
     * Email for user.
     */
    email: pulumi.Input<string>;
    /**
     * List of groups this user is a part of. If no groups set, `readers` group will be added by default. If other groups are assigned, `readers` must be added to the list manually to avoid state drift.
     */
    groups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional, Default: false) When enabled, disables the fallback mechanism for using an internal password when external authentication (such as LDAP) is enabled.
     */
    internalPasswordDisabled?: pulumi.Input<boolean>;
    /**
     * Username for user. May contain lowercase letters, numbers and symbols: '.-_@'
     */
    name?: pulumi.Input<string>;
    /**
     * (Optional, Sensitive) Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters
     */
    password?: pulumi.Input<string>;
    /**
     * (Optional, Default: true) When enabled, this user can update their profile details (except for the password. Only an administrator can update the password). There may be cases in which you want to leave this unset to prevent users from updating their profile. For example, a departmental user with a single password shared between all department members.
     */
    profileUpdatable?: pulumi.Input<boolean>;
}
