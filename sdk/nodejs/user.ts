// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides an Artifactory user resource. This can be used to create and manage Artifactory users.
 * The password is a required field by the [Artifactory API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-CreateorReplaceUser), but we made it optional in this resource to accommodate the scenario where the password is not needed and will be reset by the actual user later.\
 * When the optional attribute `password` is omitted, a random password is generated according to current Artifactory password policy.
 *
 * > The generated password won't be stored in the TF state and can not be recovered. The user must reset the password to be able to log in. An admin can always generate the access key for the user as well. The password change won't trigger state drift. We don't recommend to use this resource unless there is a specific use case for it. Recommended resource is `artifactory.ManagedUser`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * // Create a new Artifactory user called terraform
 * const test_user = new artifactory.User("test-user", {
 *     email: "test-user@artifactory-terraform.com",
 *     groups: [
 *         "logged-in-users",
 *         "readers",
 *     ],
 *     name: "terraform",
 *     password: "my super secret password",
 * });
 * ```
 * ## Managing groups relationship
 *
 * See our recommendation on how to manage user-group relationship.
 *
 * ## Import
 *
 * Users can be imported using their name, e.g.
 *
 * ```sh
 *  $ pulumi import artifactory:index/user:User test-user myusername
 * ```
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
     */
    public readonly admin!: pulumi.Output<boolean | undefined>;
    /**
     * When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
     */
    public readonly disableUiAccess!: pulumi.Output<boolean | undefined>;
    /**
     * Email for user.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * List of groups this user is a part of. **Notes:** If this attribute is not specified then user's group membership set to empty. User will not be part of default "readers" group automatically.
     */
    public readonly groups!: pulumi.Output<string[] | undefined>;
    /**
     * When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
     */
    public readonly internalPasswordDisabled!: pulumi.Output<boolean | undefined>;
    /**
     * Username for user.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
     */
    public readonly profileUpdatable!: pulumi.Output<boolean | undefined>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserState | undefined;
            resourceInputs["admin"] = state ? state.admin : undefined;
            resourceInputs["disableUiAccess"] = state ? state.disableUiAccess : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["groups"] = state ? state.groups : undefined;
            resourceInputs["internalPasswordDisabled"] = state ? state.internalPasswordDisabled : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["profileUpdatable"] = state ? state.profileUpdatable : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
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
        super(User.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
     */
    admin?: pulumi.Input<boolean>;
    /**
     * When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
     */
    disableUiAccess?: pulumi.Input<boolean>;
    /**
     * Email for user.
     */
    email?: pulumi.Input<string>;
    /**
     * List of groups this user is a part of. **Notes:** If this attribute is not specified then user's group membership set to empty. User will not be part of default "readers" group automatically.
     */
    groups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
     */
    internalPasswordDisabled?: pulumi.Input<boolean>;
    /**
     * Username for user.
     */
    name?: pulumi.Input<string>;
    /**
     * Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters.
     */
    password?: pulumi.Input<string>;
    /**
     * When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
     */
    profileUpdatable?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
     */
    admin?: pulumi.Input<boolean>;
    /**
     * When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
     */
    disableUiAccess?: pulumi.Input<boolean>;
    /**
     * Email for user.
     */
    email: pulumi.Input<string>;
    /**
     * List of groups this user is a part of. **Notes:** If this attribute is not specified then user's group membership set to empty. User will not be part of default "readers" group automatically.
     */
    groups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
     */
    internalPasswordDisabled?: pulumi.Input<boolean>;
    /**
     * Username for user.
     */
    name: pulumi.Input<string>;
    /**
     * Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters.
     */
    password?: pulumi.Input<string>;
    /**
     * When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
     */
    profileUpdatable?: pulumi.Input<boolean>;
}
