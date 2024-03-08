// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    /// <summary>
    /// Provides an Artifactory LDAP Setting resource.
    /// 
    /// This resource can be used to manage Artifactory's LDAP settings for user authentication.
    /// 
    /// When specified LDAP setting is active, Artifactory first attempts to authenticate the user against the LDAP server.
    /// If LDAP authentication fails, it then tries to authenticate via its internal database.
    /// 
    /// [API documentation](https://jfrog.com/help/r/jfrog-rest-apis/ldap-setting), [general documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/ldap).
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Artifactory = Pulumi.Artifactory;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var ldapName = new Artifactory.LdapSettingV2("ldapName", new()
    ///     {
    ///         AllowUserToAccessProfile = false,
    ///         AutoCreateUser = true,
    ///         EmailAttribute = "mail",
    ///         Enabled = true,
    ///         Key = "ldap_name",
    ///         LdapPoisoningProtection = true,
    ///         LdapUrl = "ldap://ldap_server_url",
    ///         ManagerDn = "mgr_dn",
    ///         ManagerPassword = "mgr_passwd_random",
    ///         PagingSupportEnabled = false,
    ///         SearchBase = "ou=users",
    ///         SearchFilter = "(uid={0})",
    ///         SearchSubTree = true,
    ///         UserDnPattern = "uid={0},ou=People",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import artifactory:index/ldapSettingV2:LdapSettingV2 ldap ldap1
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/ldapSettingV2:LdapSettingV2")]
    public partial class LdapSettingV2 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Auto created users will have access to their profile page and will be able to perform actions such as generating an API key. Default value is `false`.
        /// </summary>
        [Output("allowUserToAccessProfile")]
        public Output<bool> AllowUserToAccessProfile { get; private set; } = null!;

        /// <summary>
        /// When set, users are automatically created when using LDAP. Otherwise, users are transient and associated with auto-join groups defined in Artifactory. Default value is `true`.
        /// </summary>
        [Output("autoCreateUser")]
        public Output<bool> AutoCreateUser { get; private set; } = null!;

        /// <summary>
        /// An attribute that can be used to map a user's email address to a user created automatically in Artifactory. Default value is`mail`.
        /// </summary>
        [Output("emailAttribute")]
        public Output<string> EmailAttribute { get; private set; } = null!;

        /// <summary>
        /// Flag to enable or disable the ldap setting. Default value is `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// Ldap setting name.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// When this is set to `true`, an empty or missing usernames array will detach all users from the group.
        /// </summary>
        [Output("ldapPoisoningProtection")]
        public Output<bool> LdapPoisoningProtection { get; private set; } = null!;

        /// <summary>
        /// Location of the LDAP server in the following format: `ldap://myldapserver/dc=sampledomain,dc=com`
        /// </summary>
        [Output("ldapUrl")]
        public Output<string> LdapUrl { get; private set; } = null!;

        /// <summary>
        /// The full DN of the user that binds to the LDAP server to perform user searches. Only used with `search` authentication.
        /// </summary>
        [Output("managerDn")]
        public Output<string> ManagerDn { get; private set; } = null!;

        /// <summary>
        /// The password of the user that binds to the LDAP server to perform the search. Only used with `search` authentication.
        /// </summary>
        [Output("managerPassword")]
        public Output<string> ManagerPassword { get; private set; } = null!;

        /// <summary>
        /// When set, supports paging results for the LDAP server. This feature requires that the LDAP server supports a PagedResultsControl configuration. Default value is `true`.
        /// </summary>
        [Output("pagingSupportEnabled")]
        public Output<bool> PagingSupportEnabled { get; private set; } = null!;

        /// <summary>
        /// A context name to search in relative to the base DN of the LDAP URL. For example, 'ou=users' With the LDAP Group Add-on enabled, it is possible to enter multiple search base entries separated by a pipe ('|') character.
        /// </summary>
        [Output("searchBase")]
        public Output<string> SearchBase { get; private set; } = null!;

        /// <summary>
        /// A filter expression used to search for the user DN used in LDAP authentication. This is an LDAP search filter (as defined in 'RFC 2254') with optional arguments. In this case, the username is the only argument, and is denoted by '{0}'. Possible examples are: (uid={0}) - This searches for a username match on the attribute. Authentication to LDAP is performed from the DN found if successful.
        /// </summary>
        [Output("searchFilter")]
        public Output<string> SearchFilter { get; private set; } = null!;

        /// <summary>
        /// When set, enables deep search through the sub tree of the LDAP URL + search base. Default value is `true`.
        /// </summary>
        [Output("searchSubTree")]
        public Output<bool> SearchSubTree { get; private set; } = null!;

        /// <summary>
        /// A DN pattern that can be used to log users directly in to LDAP. This pattern is used to create a DN string for 'direct' user authentication where the pattern is relative to the base DN in the LDAP URL. The pattern argument {0} is replaced with the username. This only works if anonymous binding is allowed and a direct user DN can be used, which is not the default case for Active Directory (use User DN search filter instead). Example: uid={0},ou=People. Default value is blank/empty.
        /// </summary>
        [Output("userDnPattern")]
        public Output<string> UserDnPattern { get; private set; } = null!;


        /// <summary>
        /// Create a LdapSettingV2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LdapSettingV2(string name, LdapSettingV2Args args, CustomResourceOptions? options = null)
            : base("artifactory:index/ldapSettingV2:LdapSettingV2", name, args ?? new LdapSettingV2Args(), MakeResourceOptions(options, ""))
        {
        }

        private LdapSettingV2(string name, Input<string> id, LdapSettingV2State? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/ldapSettingV2:LdapSettingV2", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "managerPassword",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LdapSettingV2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LdapSettingV2 Get(string name, Input<string> id, LdapSettingV2State? state = null, CustomResourceOptions? options = null)
        {
            return new LdapSettingV2(name, id, state, options);
        }
    }

    public sealed class LdapSettingV2Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto created users will have access to their profile page and will be able to perform actions such as generating an API key. Default value is `false`.
        /// </summary>
        [Input("allowUserToAccessProfile")]
        public Input<bool>? AllowUserToAccessProfile { get; set; }

        /// <summary>
        /// When set, users are automatically created when using LDAP. Otherwise, users are transient and associated with auto-join groups defined in Artifactory. Default value is `true`.
        /// </summary>
        [Input("autoCreateUser")]
        public Input<bool>? AutoCreateUser { get; set; }

        /// <summary>
        /// An attribute that can be used to map a user's email address to a user created automatically in Artifactory. Default value is`mail`.
        /// </summary>
        [Input("emailAttribute")]
        public Input<string>? EmailAttribute { get; set; }

        /// <summary>
        /// Flag to enable or disable the ldap setting. Default value is `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Ldap setting name.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// When this is set to `true`, an empty or missing usernames array will detach all users from the group.
        /// </summary>
        [Input("ldapPoisoningProtection")]
        public Input<bool>? LdapPoisoningProtection { get; set; }

        /// <summary>
        /// Location of the LDAP server in the following format: `ldap://myldapserver/dc=sampledomain,dc=com`
        /// </summary>
        [Input("ldapUrl", required: true)]
        public Input<string> LdapUrl { get; set; } = null!;

        /// <summary>
        /// The full DN of the user that binds to the LDAP server to perform user searches. Only used with `search` authentication.
        /// </summary>
        [Input("managerDn")]
        public Input<string>? ManagerDn { get; set; }

        [Input("managerPassword")]
        private Input<string>? _managerPassword;

        /// <summary>
        /// The password of the user that binds to the LDAP server to perform the search. Only used with `search` authentication.
        /// </summary>
        public Input<string>? ManagerPassword
        {
            get => _managerPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _managerPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// When set, supports paging results for the LDAP server. This feature requires that the LDAP server supports a PagedResultsControl configuration. Default value is `true`.
        /// </summary>
        [Input("pagingSupportEnabled")]
        public Input<bool>? PagingSupportEnabled { get; set; }

        /// <summary>
        /// A context name to search in relative to the base DN of the LDAP URL. For example, 'ou=users' With the LDAP Group Add-on enabled, it is possible to enter multiple search base entries separated by a pipe ('|') character.
        /// </summary>
        [Input("searchBase")]
        public Input<string>? SearchBase { get; set; }

        /// <summary>
        /// A filter expression used to search for the user DN used in LDAP authentication. This is an LDAP search filter (as defined in 'RFC 2254') with optional arguments. In this case, the username is the only argument, and is denoted by '{0}'. Possible examples are: (uid={0}) - This searches for a username match on the attribute. Authentication to LDAP is performed from the DN found if successful.
        /// </summary>
        [Input("searchFilter")]
        public Input<string>? SearchFilter { get; set; }

        /// <summary>
        /// When set, enables deep search through the sub tree of the LDAP URL + search base. Default value is `true`.
        /// </summary>
        [Input("searchSubTree")]
        public Input<bool>? SearchSubTree { get; set; }

        /// <summary>
        /// A DN pattern that can be used to log users directly in to LDAP. This pattern is used to create a DN string for 'direct' user authentication where the pattern is relative to the base DN in the LDAP URL. The pattern argument {0} is replaced with the username. This only works if anonymous binding is allowed and a direct user DN can be used, which is not the default case for Active Directory (use User DN search filter instead). Example: uid={0},ou=People. Default value is blank/empty.
        /// </summary>
        [Input("userDnPattern")]
        public Input<string>? UserDnPattern { get; set; }

        public LdapSettingV2Args()
        {
        }
        public static new LdapSettingV2Args Empty => new LdapSettingV2Args();
    }

    public sealed class LdapSettingV2State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto created users will have access to their profile page and will be able to perform actions such as generating an API key. Default value is `false`.
        /// </summary>
        [Input("allowUserToAccessProfile")]
        public Input<bool>? AllowUserToAccessProfile { get; set; }

        /// <summary>
        /// When set, users are automatically created when using LDAP. Otherwise, users are transient and associated with auto-join groups defined in Artifactory. Default value is `true`.
        /// </summary>
        [Input("autoCreateUser")]
        public Input<bool>? AutoCreateUser { get; set; }

        /// <summary>
        /// An attribute that can be used to map a user's email address to a user created automatically in Artifactory. Default value is`mail`.
        /// </summary>
        [Input("emailAttribute")]
        public Input<string>? EmailAttribute { get; set; }

        /// <summary>
        /// Flag to enable or disable the ldap setting. Default value is `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Ldap setting name.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// When this is set to `true`, an empty or missing usernames array will detach all users from the group.
        /// </summary>
        [Input("ldapPoisoningProtection")]
        public Input<bool>? LdapPoisoningProtection { get; set; }

        /// <summary>
        /// Location of the LDAP server in the following format: `ldap://myldapserver/dc=sampledomain,dc=com`
        /// </summary>
        [Input("ldapUrl")]
        public Input<string>? LdapUrl { get; set; }

        /// <summary>
        /// The full DN of the user that binds to the LDAP server to perform user searches. Only used with `search` authentication.
        /// </summary>
        [Input("managerDn")]
        public Input<string>? ManagerDn { get; set; }

        [Input("managerPassword")]
        private Input<string>? _managerPassword;

        /// <summary>
        /// The password of the user that binds to the LDAP server to perform the search. Only used with `search` authentication.
        /// </summary>
        public Input<string>? ManagerPassword
        {
            get => _managerPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _managerPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// When set, supports paging results for the LDAP server. This feature requires that the LDAP server supports a PagedResultsControl configuration. Default value is `true`.
        /// </summary>
        [Input("pagingSupportEnabled")]
        public Input<bool>? PagingSupportEnabled { get; set; }

        /// <summary>
        /// A context name to search in relative to the base DN of the LDAP URL. For example, 'ou=users' With the LDAP Group Add-on enabled, it is possible to enter multiple search base entries separated by a pipe ('|') character.
        /// </summary>
        [Input("searchBase")]
        public Input<string>? SearchBase { get; set; }

        /// <summary>
        /// A filter expression used to search for the user DN used in LDAP authentication. This is an LDAP search filter (as defined in 'RFC 2254') with optional arguments. In this case, the username is the only argument, and is denoted by '{0}'. Possible examples are: (uid={0}) - This searches for a username match on the attribute. Authentication to LDAP is performed from the DN found if successful.
        /// </summary>
        [Input("searchFilter")]
        public Input<string>? SearchFilter { get; set; }

        /// <summary>
        /// When set, enables deep search through the sub tree of the LDAP URL + search base. Default value is `true`.
        /// </summary>
        [Input("searchSubTree")]
        public Input<bool>? SearchSubTree { get; set; }

        /// <summary>
        /// A DN pattern that can be used to log users directly in to LDAP. This pattern is used to create a DN string for 'direct' user authentication where the pattern is relative to the base DN in the LDAP URL. The pattern argument {0} is replaced with the username. This only works if anonymous binding is allowed and a direct user DN can be used, which is not the default case for Active Directory (use User DN search filter instead). Example: uid={0},ou=People. Default value is blank/empty.
        /// </summary>
        [Input("userDnPattern")]
        public Input<string>? UserDnPattern { get; set; }

        public LdapSettingV2State()
        {
        }
        public static new LdapSettingV2State Empty => new LdapSettingV2State();
    }
}
