// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class LdapSetting extends pulumi.CustomResource {
    /**
     * Get an existing LdapSetting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LdapSettingState, opts?: pulumi.CustomResourceOptions): LdapSetting {
        return new LdapSetting(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/ldapSetting:LdapSetting';

    /**
     * Returns true if the given object is an instance of LdapSetting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LdapSetting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LdapSetting.__pulumiType;
    }

    /**
     * (Optional) Auto created users will have access to their profile page and will be able to perform actions such as
     * generating an API key. Default value is "false".
     */
    public readonly allowUserToAccessProfile!: pulumi.Output<boolean | undefined>;
    /**
     * (Optional) When set, users are automatically created when using LDAP. Otherwise, users are transient and associated with
     * auto-join groups defined in Artifactory. Default value is "true".
     */
    public readonly autoCreateUser!: pulumi.Output<boolean | undefined>;
    /**
     * (Optional) An attribute that can be used to map a user's email address to a user created automatically in Artifactory.
     * Default value is "mail".
     */
    public readonly emailAttribute!: pulumi.Output<string | undefined>;
    /**
     * (Optional) Flag to enable or disable the ldap setting. Default value is "true".
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * (Required) Ldap setting name.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * (Optional) Protects against LDAP poisoning by filtering out users exposed to vulnerabilities. Default value is "true".
     */
    public readonly ldapPoisoningProtection!: pulumi.Output<boolean | undefined>;
    /**
     * (Required) Location of the LDAP server in the following format: ldap://myldapserver/dc=sampledomain,dc=com
     */
    public readonly ldapUrl!: pulumi.Output<string>;
    /**
     * (Optional) The full DN of the user that binds to the LDAP server to perform user searches. Only used with "search"
     * authentication.
     */
    public readonly managerDn!: pulumi.Output<string | undefined>;
    /**
     * (Optional) The password of the user that binds to the LDAP server to perform the search. Only used with "search"
     * authentication.
     */
    public readonly managerPassword!: pulumi.Output<string>;
    /**
     * (Optional) When set, supports paging results for the LDAP server. This feature requires that the LDAP server supports a
     * PagedResultsControl configuration. Default value is "true".
     */
    public readonly pagingSupportEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * (Optional) A context name to search in relative to the base DN of the LDAP URL. For example, 'ou=users' With the LDAP
     * Group Add-on enabled, it is possible to enter multiple search base entries separated by a pipe ('|') character.
     */
    public readonly searchBase!: pulumi.Output<string | undefined>;
    /**
     * (Optional) A filter expression used to search for the user DN used in LDAP authentication. This is an LDAP search filter
     * (as defined in 'RFC 2254') with optional arguments. In this case, the username is the only argument, and is denoted by
     * '{0}'. Possible examples are: (uid={0}) - This searches for a username match on the attribute. Authentication to LDAP is
     * performed from the DN found if successful.
     */
    public readonly searchFilter!: pulumi.Output<string | undefined>;
    /**
     * (Optional) When set, enables deep search through the sub tree of the LDAP URL + search base. Default value is "true".
     */
    public readonly searchSubTree!: pulumi.Output<boolean | undefined>;
    /**
     * (Optional) A DN pattern that can be used to log users directly in to LDAP. This pattern is used to create a DN string
     * for 'direct' user authentication where the pattern is relative to the base DN in the LDAP URL. The pattern argument {0}
     * is replaced with the username. This only works if anonymous binding is allowed and a direct user DN can be used, which
     * is not the default case for Active Directory (use User DN search filter instead). Example: uid={0},ou=People. Default
     * value is blank/empty.
     */
    public readonly userDnPattern!: pulumi.Output<string | undefined>;

    /**
     * Create a LdapSetting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LdapSettingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LdapSettingArgs | LdapSettingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LdapSettingState | undefined;
            resourceInputs["allowUserToAccessProfile"] = state ? state.allowUserToAccessProfile : undefined;
            resourceInputs["autoCreateUser"] = state ? state.autoCreateUser : undefined;
            resourceInputs["emailAttribute"] = state ? state.emailAttribute : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["ldapPoisoningProtection"] = state ? state.ldapPoisoningProtection : undefined;
            resourceInputs["ldapUrl"] = state ? state.ldapUrl : undefined;
            resourceInputs["managerDn"] = state ? state.managerDn : undefined;
            resourceInputs["managerPassword"] = state ? state.managerPassword : undefined;
            resourceInputs["pagingSupportEnabled"] = state ? state.pagingSupportEnabled : undefined;
            resourceInputs["searchBase"] = state ? state.searchBase : undefined;
            resourceInputs["searchFilter"] = state ? state.searchFilter : undefined;
            resourceInputs["searchSubTree"] = state ? state.searchSubTree : undefined;
            resourceInputs["userDnPattern"] = state ? state.userDnPattern : undefined;
        } else {
            const args = argsOrState as LdapSettingArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.ldapUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ldapUrl'");
            }
            resourceInputs["allowUserToAccessProfile"] = args ? args.allowUserToAccessProfile : undefined;
            resourceInputs["autoCreateUser"] = args ? args.autoCreateUser : undefined;
            resourceInputs["emailAttribute"] = args ? args.emailAttribute : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["ldapPoisoningProtection"] = args ? args.ldapPoisoningProtection : undefined;
            resourceInputs["ldapUrl"] = args ? args.ldapUrl : undefined;
            resourceInputs["managerDn"] = args ? args.managerDn : undefined;
            resourceInputs["managerPassword"] = args ? args.managerPassword : undefined;
            resourceInputs["pagingSupportEnabled"] = args ? args.pagingSupportEnabled : undefined;
            resourceInputs["searchBase"] = args ? args.searchBase : undefined;
            resourceInputs["searchFilter"] = args ? args.searchFilter : undefined;
            resourceInputs["searchSubTree"] = args ? args.searchSubTree : undefined;
            resourceInputs["userDnPattern"] = args ? args.userDnPattern : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LdapSetting.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LdapSetting resources.
 */
export interface LdapSettingState {
    /**
     * (Optional) Auto created users will have access to their profile page and will be able to perform actions such as
     * generating an API key. Default value is "false".
     */
    allowUserToAccessProfile?: pulumi.Input<boolean>;
    /**
     * (Optional) When set, users are automatically created when using LDAP. Otherwise, users are transient and associated with
     * auto-join groups defined in Artifactory. Default value is "true".
     */
    autoCreateUser?: pulumi.Input<boolean>;
    /**
     * (Optional) An attribute that can be used to map a user's email address to a user created automatically in Artifactory.
     * Default value is "mail".
     */
    emailAttribute?: pulumi.Input<string>;
    /**
     * (Optional) Flag to enable or disable the ldap setting. Default value is "true".
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * (Required) Ldap setting name.
     */
    key?: pulumi.Input<string>;
    /**
     * (Optional) Protects against LDAP poisoning by filtering out users exposed to vulnerabilities. Default value is "true".
     */
    ldapPoisoningProtection?: pulumi.Input<boolean>;
    /**
     * (Required) Location of the LDAP server in the following format: ldap://myldapserver/dc=sampledomain,dc=com
     */
    ldapUrl?: pulumi.Input<string>;
    /**
     * (Optional) The full DN of the user that binds to the LDAP server to perform user searches. Only used with "search"
     * authentication.
     */
    managerDn?: pulumi.Input<string>;
    /**
     * (Optional) The password of the user that binds to the LDAP server to perform the search. Only used with "search"
     * authentication.
     */
    managerPassword?: pulumi.Input<string>;
    /**
     * (Optional) When set, supports paging results for the LDAP server. This feature requires that the LDAP server supports a
     * PagedResultsControl configuration. Default value is "true".
     */
    pagingSupportEnabled?: pulumi.Input<boolean>;
    /**
     * (Optional) A context name to search in relative to the base DN of the LDAP URL. For example, 'ou=users' With the LDAP
     * Group Add-on enabled, it is possible to enter multiple search base entries separated by a pipe ('|') character.
     */
    searchBase?: pulumi.Input<string>;
    /**
     * (Optional) A filter expression used to search for the user DN used in LDAP authentication. This is an LDAP search filter
     * (as defined in 'RFC 2254') with optional arguments. In this case, the username is the only argument, and is denoted by
     * '{0}'. Possible examples are: (uid={0}) - This searches for a username match on the attribute. Authentication to LDAP is
     * performed from the DN found if successful.
     */
    searchFilter?: pulumi.Input<string>;
    /**
     * (Optional) When set, enables deep search through the sub tree of the LDAP URL + search base. Default value is "true".
     */
    searchSubTree?: pulumi.Input<boolean>;
    /**
     * (Optional) A DN pattern that can be used to log users directly in to LDAP. This pattern is used to create a DN string
     * for 'direct' user authentication where the pattern is relative to the base DN in the LDAP URL. The pattern argument {0}
     * is replaced with the username. This only works if anonymous binding is allowed and a direct user DN can be used, which
     * is not the default case for Active Directory (use User DN search filter instead). Example: uid={0},ou=People. Default
     * value is blank/empty.
     */
    userDnPattern?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LdapSetting resource.
 */
export interface LdapSettingArgs {
    /**
     * (Optional) Auto created users will have access to their profile page and will be able to perform actions such as
     * generating an API key. Default value is "false".
     */
    allowUserToAccessProfile?: pulumi.Input<boolean>;
    /**
     * (Optional) When set, users are automatically created when using LDAP. Otherwise, users are transient and associated with
     * auto-join groups defined in Artifactory. Default value is "true".
     */
    autoCreateUser?: pulumi.Input<boolean>;
    /**
     * (Optional) An attribute that can be used to map a user's email address to a user created automatically in Artifactory.
     * Default value is "mail".
     */
    emailAttribute?: pulumi.Input<string>;
    /**
     * (Optional) Flag to enable or disable the ldap setting. Default value is "true".
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * (Required) Ldap setting name.
     */
    key: pulumi.Input<string>;
    /**
     * (Optional) Protects against LDAP poisoning by filtering out users exposed to vulnerabilities. Default value is "true".
     */
    ldapPoisoningProtection?: pulumi.Input<boolean>;
    /**
     * (Required) Location of the LDAP server in the following format: ldap://myldapserver/dc=sampledomain,dc=com
     */
    ldapUrl: pulumi.Input<string>;
    /**
     * (Optional) The full DN of the user that binds to the LDAP server to perform user searches. Only used with "search"
     * authentication.
     */
    managerDn?: pulumi.Input<string>;
    /**
     * (Optional) The password of the user that binds to the LDAP server to perform the search. Only used with "search"
     * authentication.
     */
    managerPassword?: pulumi.Input<string>;
    /**
     * (Optional) When set, supports paging results for the LDAP server. This feature requires that the LDAP server supports a
     * PagedResultsControl configuration. Default value is "true".
     */
    pagingSupportEnabled?: pulumi.Input<boolean>;
    /**
     * (Optional) A context name to search in relative to the base DN of the LDAP URL. For example, 'ou=users' With the LDAP
     * Group Add-on enabled, it is possible to enter multiple search base entries separated by a pipe ('|') character.
     */
    searchBase?: pulumi.Input<string>;
    /**
     * (Optional) A filter expression used to search for the user DN used in LDAP authentication. This is an LDAP search filter
     * (as defined in 'RFC 2254') with optional arguments. In this case, the username is the only argument, and is denoted by
     * '{0}'. Possible examples are: (uid={0}) - This searches for a username match on the attribute. Authentication to LDAP is
     * performed from the DN found if successful.
     */
    searchFilter?: pulumi.Input<string>;
    /**
     * (Optional) When set, enables deep search through the sub tree of the LDAP URL + search base. Default value is "true".
     */
    searchSubTree?: pulumi.Input<boolean>;
    /**
     * (Optional) A DN pattern that can be used to log users directly in to LDAP. This pattern is used to create a DN string
     * for 'direct' user authentication where the pattern is relative to the base DN in the LDAP URL. The pattern argument {0}
     * is replaced with the username. This only works if anonymous binding is allowed and a direct user DN can be used, which
     * is not the default case for Active Directory (use User DN search filter instead). Example: uid={0},ou=People. Default
     * value is blank/empty.
     */
    userDnPattern?: pulumi.Input<string>;
}
