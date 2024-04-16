// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides an Artifactory Scoped Token resource. This can be used to create and manage Artifactory Scoped Tokens.
 *
 * !>Scoped Tokens will be stored in the raw state as plain-text. Read more about sensitive data in
 * state.
 *
 * ~>Token would not be saved by Artifactory if `expiresIn` is less than the persistency threshold value (default to 10800 seconds) set in Access configuration. See [Persistency Threshold](https://jfrog.com/help/r/jfrog-platform-administration-documentation/using-the-revocable-and-persistency-thresholds) for details.
 *
 * ## Example Usage
 *
 * ### S
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * //## Create a new Artifactory scoped token for an existing user
 * const scopedToken = new artifactory.ScopedToken("scoped_token", {username: "existing-user"});
 * //## **Note:** This assumes that the user `existing-user` has already been created in Artifactory by different means, i.e. manually or in a separate pulumi up.
 * //## Create a new Artifactory user and scoped token
 * const newUser = new artifactory.User("new_user", {
 *     name: "new_user",
 *     email: "new_user@somewhere.com",
 *     groups: ["readers"],
 * });
 * const scopedTokenUser = new artifactory.ScopedToken("scoped_token_user", {username: newUser.name});
 * //## Creates a new token for groups
 * const scopedTokenGroup = new artifactory.ScopedToken("scoped_token_group", {scopes: ["applied-permissions/groups:readers"]});
 * //## Create token with expiry
 * const scopedTokenNoExpiry = new artifactory.ScopedToken("scoped_token_no_expiry", {
 *     username: "existing-user",
 *     expiresIn: 7200,
 * });
 * //## Creates a refreshable token
 * const scopedTokenRefreshable = new artifactory.ScopedToken("scoped_token_refreshable", {
 *     username: "existing-user",
 *     refreshable: true,
 * });
 * //## Creates an administrator token
 * const admin = new artifactory.ScopedToken("admin", {
 *     username: "admin-user",
 *     scopes: ["applied-permissions/admin"],
 * });
 * //## Creates a token with an audience
 * const audience = new artifactory.ScopedToken("audience", {
 *     username: "admin-user",
 *     scopes: ["applied-permissions/admin"],
 *     audiences: ["jfrt@*"],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## References
 *
 * - https://jfrog.com/help/r/jfrog-platform-administration-documentation/access-tokens
 * - https://jfrog.com/help/r/jfrog-rest-apis/access-tokens
 *
 * ## Import
 *
 * Artifactory **does not** retain scoped tokens, and they cannot be imported into state.
 */
export class ScopedToken extends pulumi.CustomResource {
    /**
     * Get an existing ScopedToken resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ScopedTokenState, opts?: pulumi.CustomResourceOptions): ScopedToken {
        return new ScopedToken(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/scopedToken:ScopedToken';

    /**
     * Returns true if the given object is an instance of ScopedToken.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScopedToken {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ScopedToken.__pulumiType;
    }

    /**
     * Returns the access token to authenticate to Artifactory.
     */
    public /*out*/ readonly accessToken!: pulumi.Output<string>;
    /**
     * A list of the other instances or services that should accept this token identified by their Service-IDs. Limited to
     * total 255 characters. Default to '*@*' if not set. Service ID must begin with valid JFrog service type. Options: jfrt,
     * jfxr, jfpip, jfds, jfmc, jfac, jfevt, jfmd, jfcon, or *. For instructions to retrieve the Artifactory Service ID see
     * this [documentation](https://jfrog.com/help/r/jfrog-rest-apis/get-service-id)
     */
    public readonly audiences!: pulumi.Output<string[] | undefined>;
    /**
     * Free text token description. Useful for filtering and managing tokens. Limited to 1024 characters.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The amount of time, in seconds, it would take for the token to expire. An admin shall be able to set whether expiry is
     * mandatory, what is the default expiry, and what is the maximum expiry allowed. Must be non-negative. Default value is
     * based on configuration in 'access.config.yaml'. See [API
     * documentation](https://jfrog.com/help/r/jfrog-rest-apis/revoke-token-by-id) for details. Access Token would not be saved
     * by Artifactory if this is less than the persistence threshold value (default to 10800 seconds) set in Access
     * configuration. See [official
     * documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/using-the-revocable-and-persistency-thresholds)
     * for details.
     */
    public readonly expiresIn!: pulumi.Output<number>;
    /**
     * Returns the token expiry.
     */
    public /*out*/ readonly expiry!: pulumi.Output<number>;
    /**
     * The grant type used to authenticate the request. In this case, the only value supported is `client_credentials` which is
     * also the default value if this parameter is not specified.
     */
    public readonly grantType!: pulumi.Output<string>;
    /**
     * Also create a reference token which can be used like an API key. Default is `false`.
     */
    public readonly includeReferenceToken!: pulumi.Output<boolean>;
    /**
     * Returns the token issued at date/time.
     */
    public /*out*/ readonly issuedAt!: pulumi.Output<number>;
    /**
     * Returns the token issuer.
     */
    public /*out*/ readonly issuer!: pulumi.Output<string>;
    /**
     * The project for which this token is created. Enter the project name on which you want to apply this token.
     */
    public readonly projectKey!: pulumi.Output<string | undefined>;
    /**
     * Reference Token (alias to Access Token).
     */
    public /*out*/ readonly referenceToken!: pulumi.Output<string>;
    /**
     * Refresh token.
     */
    public /*out*/ readonly refreshToken!: pulumi.Output<string>;
    /**
     * Is this token refreshable? Default is `false`.
     */
    public readonly refreshable!: pulumi.Output<boolean>;
    /**
     * The scope of access that the token provides. Access to the REST API is always provided by default. Administrators can
     * set any scope, while non-admin users can only set the scope to a subset of the groups to which they belong. The
     * supported scopes include: * `applied-permissions/user` - provides user access. If left at the default setting, the token
     * will be created with the user-identity scope, which allows users to identify themselves in the Platform but does not
     * grant any specific access permissions.* `applied-permissions/admin` - the scope assigned to admin users.*
     * `applied-permissions/groups` - this scope assigns permissions to groups using the following format:
     * applied-permissions/groups:<group-name>[,<group-name>...]* `system:metrics:r` - for getting the service metrics*
     * `system:livelogs:r` - for getting the service livelogsr. The scope to assign to the token should be provided as a list
     * of scope tokens, limited to 500 characters in total. Resource Permissions From Artifactory 7.38.x, resource permissions
     * scoped tokens are also supported in the REST API. A permission can be represented as a scope token string in the
     * following format: `<resource-type>:<target>[/<sub-resource>]:<actions>` Where: `<resource-type>` - one of the permission
     * resource types, from a predefined closed list. Currently, the only resource type that is supported is the artifact
     * resource type. `<target>` - the target resource, can be exact name or a pattern `<sub-resource>` - optional, the target
     * sub-resource, can be exact name or a pattern `<actions>` - comma-separated list of action acronyms.The actions allowed
     * are `r`, `w`, `d`, `a`, `m`, `x`, `s`, or any combination of these actions. To allow all actions - use `*` Examples:
     * `["applied-permissions/user", "artifact:generic-local:r"]` `["applied-permissions/group",
     * "artifact:generic-local/path:*"]` `["applied-permissions/admin", "system:metrics:r", "artifact:generic-local:*"]`
     */
    public readonly scopes!: pulumi.Output<string[]>;
    /**
     * Returns the token type.
     */
    public /*out*/ readonly subject!: pulumi.Output<string>;
    /**
     * Returns the token type.
     */
    public /*out*/ readonly tokenType!: pulumi.Output<string>;
    /**
     * The user name for which this token is created. The username is based on the authenticated user - either from the user of
     * the authenticated token or based on the username (if basic auth was used). The username is then used to set the subject
     * of the token: <service-id>/users/<username>. Limited to 255 characters.
     */
    public readonly username!: pulumi.Output<string | undefined>;

    /**
     * Create a ScopedToken resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ScopedTokenArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScopedTokenArgs | ScopedTokenState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ScopedTokenState | undefined;
            resourceInputs["accessToken"] = state ? state.accessToken : undefined;
            resourceInputs["audiences"] = state ? state.audiences : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["expiresIn"] = state ? state.expiresIn : undefined;
            resourceInputs["expiry"] = state ? state.expiry : undefined;
            resourceInputs["grantType"] = state ? state.grantType : undefined;
            resourceInputs["includeReferenceToken"] = state ? state.includeReferenceToken : undefined;
            resourceInputs["issuedAt"] = state ? state.issuedAt : undefined;
            resourceInputs["issuer"] = state ? state.issuer : undefined;
            resourceInputs["projectKey"] = state ? state.projectKey : undefined;
            resourceInputs["referenceToken"] = state ? state.referenceToken : undefined;
            resourceInputs["refreshToken"] = state ? state.refreshToken : undefined;
            resourceInputs["refreshable"] = state ? state.refreshable : undefined;
            resourceInputs["scopes"] = state ? state.scopes : undefined;
            resourceInputs["subject"] = state ? state.subject : undefined;
            resourceInputs["tokenType"] = state ? state.tokenType : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as ScopedTokenArgs | undefined;
            resourceInputs["audiences"] = args ? args.audiences : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["expiresIn"] = args ? args.expiresIn : undefined;
            resourceInputs["grantType"] = args ? args.grantType : undefined;
            resourceInputs["includeReferenceToken"] = args ? args.includeReferenceToken : undefined;
            resourceInputs["projectKey"] = args ? args.projectKey : undefined;
            resourceInputs["refreshable"] = args ? args.refreshable : undefined;
            resourceInputs["scopes"] = args ? args.scopes : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["accessToken"] = undefined /*out*/;
            resourceInputs["expiry"] = undefined /*out*/;
            resourceInputs["issuedAt"] = undefined /*out*/;
            resourceInputs["issuer"] = undefined /*out*/;
            resourceInputs["referenceToken"] = undefined /*out*/;
            resourceInputs["refreshToken"] = undefined /*out*/;
            resourceInputs["subject"] = undefined /*out*/;
            resourceInputs["tokenType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["accessToken", "referenceToken", "refreshToken"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ScopedToken.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ScopedToken resources.
 */
export interface ScopedTokenState {
    /**
     * Returns the access token to authenticate to Artifactory.
     */
    accessToken?: pulumi.Input<string>;
    /**
     * A list of the other instances or services that should accept this token identified by their Service-IDs. Limited to
     * total 255 characters. Default to '*@*' if not set. Service ID must begin with valid JFrog service type. Options: jfrt,
     * jfxr, jfpip, jfds, jfmc, jfac, jfevt, jfmd, jfcon, or *. For instructions to retrieve the Artifactory Service ID see
     * this [documentation](https://jfrog.com/help/r/jfrog-rest-apis/get-service-id)
     */
    audiences?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Free text token description. Useful for filtering and managing tokens. Limited to 1024 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * The amount of time, in seconds, it would take for the token to expire. An admin shall be able to set whether expiry is
     * mandatory, what is the default expiry, and what is the maximum expiry allowed. Must be non-negative. Default value is
     * based on configuration in 'access.config.yaml'. See [API
     * documentation](https://jfrog.com/help/r/jfrog-rest-apis/revoke-token-by-id) for details. Access Token would not be saved
     * by Artifactory if this is less than the persistence threshold value (default to 10800 seconds) set in Access
     * configuration. See [official
     * documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/using-the-revocable-and-persistency-thresholds)
     * for details.
     */
    expiresIn?: pulumi.Input<number>;
    /**
     * Returns the token expiry.
     */
    expiry?: pulumi.Input<number>;
    /**
     * The grant type used to authenticate the request. In this case, the only value supported is `client_credentials` which is
     * also the default value if this parameter is not specified.
     */
    grantType?: pulumi.Input<string>;
    /**
     * Also create a reference token which can be used like an API key. Default is `false`.
     */
    includeReferenceToken?: pulumi.Input<boolean>;
    /**
     * Returns the token issued at date/time.
     */
    issuedAt?: pulumi.Input<number>;
    /**
     * Returns the token issuer.
     */
    issuer?: pulumi.Input<string>;
    /**
     * The project for which this token is created. Enter the project name on which you want to apply this token.
     */
    projectKey?: pulumi.Input<string>;
    /**
     * Reference Token (alias to Access Token).
     */
    referenceToken?: pulumi.Input<string>;
    /**
     * Refresh token.
     */
    refreshToken?: pulumi.Input<string>;
    /**
     * Is this token refreshable? Default is `false`.
     */
    refreshable?: pulumi.Input<boolean>;
    /**
     * The scope of access that the token provides. Access to the REST API is always provided by default. Administrators can
     * set any scope, while non-admin users can only set the scope to a subset of the groups to which they belong. The
     * supported scopes include: * `applied-permissions/user` - provides user access. If left at the default setting, the token
     * will be created with the user-identity scope, which allows users to identify themselves in the Platform but does not
     * grant any specific access permissions.* `applied-permissions/admin` - the scope assigned to admin users.*
     * `applied-permissions/groups` - this scope assigns permissions to groups using the following format:
     * applied-permissions/groups:<group-name>[,<group-name>...]* `system:metrics:r` - for getting the service metrics*
     * `system:livelogs:r` - for getting the service livelogsr. The scope to assign to the token should be provided as a list
     * of scope tokens, limited to 500 characters in total. Resource Permissions From Artifactory 7.38.x, resource permissions
     * scoped tokens are also supported in the REST API. A permission can be represented as a scope token string in the
     * following format: `<resource-type>:<target>[/<sub-resource>]:<actions>` Where: `<resource-type>` - one of the permission
     * resource types, from a predefined closed list. Currently, the only resource type that is supported is the artifact
     * resource type. `<target>` - the target resource, can be exact name or a pattern `<sub-resource>` - optional, the target
     * sub-resource, can be exact name or a pattern `<actions>` - comma-separated list of action acronyms.The actions allowed
     * are `r`, `w`, `d`, `a`, `m`, `x`, `s`, or any combination of these actions. To allow all actions - use `*` Examples:
     * `["applied-permissions/user", "artifact:generic-local:r"]` `["applied-permissions/group",
     * "artifact:generic-local/path:*"]` `["applied-permissions/admin", "system:metrics:r", "artifact:generic-local:*"]`
     */
    scopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Returns the token type.
     */
    subject?: pulumi.Input<string>;
    /**
     * Returns the token type.
     */
    tokenType?: pulumi.Input<string>;
    /**
     * The user name for which this token is created. The username is based on the authenticated user - either from the user of
     * the authenticated token or based on the username (if basic auth was used). The username is then used to set the subject
     * of the token: <service-id>/users/<username>. Limited to 255 characters.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ScopedToken resource.
 */
export interface ScopedTokenArgs {
    /**
     * A list of the other instances or services that should accept this token identified by their Service-IDs. Limited to
     * total 255 characters. Default to '*@*' if not set. Service ID must begin with valid JFrog service type. Options: jfrt,
     * jfxr, jfpip, jfds, jfmc, jfac, jfevt, jfmd, jfcon, or *. For instructions to retrieve the Artifactory Service ID see
     * this [documentation](https://jfrog.com/help/r/jfrog-rest-apis/get-service-id)
     */
    audiences?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Free text token description. Useful for filtering and managing tokens. Limited to 1024 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * The amount of time, in seconds, it would take for the token to expire. An admin shall be able to set whether expiry is
     * mandatory, what is the default expiry, and what is the maximum expiry allowed. Must be non-negative. Default value is
     * based on configuration in 'access.config.yaml'. See [API
     * documentation](https://jfrog.com/help/r/jfrog-rest-apis/revoke-token-by-id) for details. Access Token would not be saved
     * by Artifactory if this is less than the persistence threshold value (default to 10800 seconds) set in Access
     * configuration. See [official
     * documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/using-the-revocable-and-persistency-thresholds)
     * for details.
     */
    expiresIn?: pulumi.Input<number>;
    /**
     * The grant type used to authenticate the request. In this case, the only value supported is `client_credentials` which is
     * also the default value if this parameter is not specified.
     */
    grantType?: pulumi.Input<string>;
    /**
     * Also create a reference token which can be used like an API key. Default is `false`.
     */
    includeReferenceToken?: pulumi.Input<boolean>;
    /**
     * The project for which this token is created. Enter the project name on which you want to apply this token.
     */
    projectKey?: pulumi.Input<string>;
    /**
     * Is this token refreshable? Default is `false`.
     */
    refreshable?: pulumi.Input<boolean>;
    /**
     * The scope of access that the token provides. Access to the REST API is always provided by default. Administrators can
     * set any scope, while non-admin users can only set the scope to a subset of the groups to which they belong. The
     * supported scopes include: * `applied-permissions/user` - provides user access. If left at the default setting, the token
     * will be created with the user-identity scope, which allows users to identify themselves in the Platform but does not
     * grant any specific access permissions.* `applied-permissions/admin` - the scope assigned to admin users.*
     * `applied-permissions/groups` - this scope assigns permissions to groups using the following format:
     * applied-permissions/groups:<group-name>[,<group-name>...]* `system:metrics:r` - for getting the service metrics*
     * `system:livelogs:r` - for getting the service livelogsr. The scope to assign to the token should be provided as a list
     * of scope tokens, limited to 500 characters in total. Resource Permissions From Artifactory 7.38.x, resource permissions
     * scoped tokens are also supported in the REST API. A permission can be represented as a scope token string in the
     * following format: `<resource-type>:<target>[/<sub-resource>]:<actions>` Where: `<resource-type>` - one of the permission
     * resource types, from a predefined closed list. Currently, the only resource type that is supported is the artifact
     * resource type. `<target>` - the target resource, can be exact name or a pattern `<sub-resource>` - optional, the target
     * sub-resource, can be exact name or a pattern `<actions>` - comma-separated list of action acronyms.The actions allowed
     * are `r`, `w`, `d`, `a`, `m`, `x`, `s`, or any combination of these actions. To allow all actions - use `*` Examples:
     * `["applied-permissions/user", "artifact:generic-local:r"]` `["applied-permissions/group",
     * "artifact:generic-local/path:*"]` `["applied-permissions/admin", "system:metrics:r", "artifact:generic-local:*"]`
     */
    scopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The user name for which this token is created. The username is based on the authenticated user - either from the user of
     * the authenticated token or based on the username (if basic auth was used). The username is then used to set the subject
     * of the token: <service-id>/users/<username>. Limited to 255 characters.
     */
    username?: pulumi.Input<string>;
}
