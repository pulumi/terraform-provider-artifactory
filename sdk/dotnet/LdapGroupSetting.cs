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
    /// This resource can be used to manage Artifactory's LDAP Group settings for user authentication.
    /// 
    /// LDAP Groups Add-on allows you to synchronize your LDAP groups with the system and leverage your existing organizational
    /// structure for managing group-based permissions.
    /// 
    /// ~&gt;The `artifactory.LdapGroupSetting` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Artifactory = Pulumi.Artifactory;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Configure Artifactory LDAP setting
    ///     var ldapGroupName = new Artifactory.LdapGroupSetting("ldap_group_name", new()
    ///     {
    ///         Name = "ldap_group_name",
    ///         LdapSettingKey = "ldap_name",
    ///         GroupBaseDn = "",
    ///         GroupNameAttribute = "cn",
    ///         GroupMemberAttribute = "uniqueMember",
    ///         SubTree = true,
    ///         Filter = "(objectClass=groupOfNames)",
    ///         DescriptionAttribute = "description",
    ///         Strategy = "STATIC",
    ///     });
    /// 
    /// });
    /// ```
    /// Note: `Name` argument has to match to the resource name.Reference Link: [JFrog LDAP](https://www.jfrog.com/confluence/display/JFROG/LDAP)
    /// 
    /// ## Import
    /// 
    /// LDAP Group setting can be imported using the key, e.g.
    /// 
    /// ```sh
    /// $ pulumi import artifactory:index/ldapGroupSetting:LdapGroupSetting ldap_group_name ldap_group_name
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/ldapGroupSetting:LdapGroupSetting")]
    public partial class LdapGroupSetting : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An attribute on the group entry which denoting the group description. Used when importing groups.
        /// </summary>
        [Output("descriptionAttribute")]
        public Output<string> DescriptionAttribute { get; private set; } = null!;

        /// <summary>
        /// The LDAP filter used to search for group entries. Used for importing groups.
        /// </summary>
        [Output("filter")]
        public Output<string> Filter { get; private set; } = null!;

        /// <summary>
        /// A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
        /// </summary>
        [Output("groupBaseDn")]
        public Output<string?> GroupBaseDn { get; private set; } = null!;

        /// <summary>
        /// A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember,member).
        /// </summary>
        [Output("groupMemberAttribute")]
        public Output<string> GroupMemberAttribute { get; private set; } = null!;

        /// <summary>
        /// Attribute on the group entry denoting the group name. Used when importing groups.
        /// </summary>
        [Output("groupNameAttribute")]
        public Output<string> GroupNameAttribute { get; private set; } = null!;

        /// <summary>
        /// The LDAP setting key you want to use for group retrieval. The value for this field corresponds to 'enabledLdap' field of the ldap group setting XML block of system configuration.
        /// </summary>
        [Output("ldapSettingKey")]
        public Output<string> LdapSettingKey { get; private set; } = null!;

        /// <summary>
        /// Ldap group setting name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas:
        /// - STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN.
        /// - DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member.
        /// - HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org indicates that user1 belongs to two groups: uk and developers.
        /// </summary>
        [Output("strategy")]
        public Output<string> Strategy { get; private set; } = null!;

        /// <summary>
        /// When set, enables deep search through the sub-tree of the LDAP URL + Search Base. True by default.
        /// </summary>
        [Output("subTree")]
        public Output<bool?> SubTree { get; private set; } = null!;


        /// <summary>
        /// Create a LdapGroupSetting resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LdapGroupSetting(string name, LdapGroupSettingArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/ldapGroupSetting:LdapGroupSetting", name, args ?? new LdapGroupSettingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LdapGroupSetting(string name, Input<string> id, LdapGroupSettingState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/ldapGroupSetting:LdapGroupSetting", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LdapGroupSetting resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LdapGroupSetting Get(string name, Input<string> id, LdapGroupSettingState? state = null, CustomResourceOptions? options = null)
        {
            return new LdapGroupSetting(name, id, state, options);
        }
    }

    public sealed class LdapGroupSettingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An attribute on the group entry which denoting the group description. Used when importing groups.
        /// </summary>
        [Input("descriptionAttribute", required: true)]
        public Input<string> DescriptionAttribute { get; set; } = null!;

        /// <summary>
        /// The LDAP filter used to search for group entries. Used for importing groups.
        /// </summary>
        [Input("filter", required: true)]
        public Input<string> Filter { get; set; } = null!;

        /// <summary>
        /// A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
        /// </summary>
        [Input("groupBaseDn")]
        public Input<string>? GroupBaseDn { get; set; }

        /// <summary>
        /// A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember,member).
        /// </summary>
        [Input("groupMemberAttribute", required: true)]
        public Input<string> GroupMemberAttribute { get; set; } = null!;

        /// <summary>
        /// Attribute on the group entry denoting the group name. Used when importing groups.
        /// </summary>
        [Input("groupNameAttribute", required: true)]
        public Input<string> GroupNameAttribute { get; set; } = null!;

        /// <summary>
        /// The LDAP setting key you want to use for group retrieval. The value for this field corresponds to 'enabledLdap' field of the ldap group setting XML block of system configuration.
        /// </summary>
        [Input("ldapSettingKey", required: true)]
        public Input<string> LdapSettingKey { get; set; } = null!;

        /// <summary>
        /// Ldap group setting name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas:
        /// - STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN.
        /// - DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member.
        /// - HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org indicates that user1 belongs to two groups: uk and developers.
        /// </summary>
        [Input("strategy", required: true)]
        public Input<string> Strategy { get; set; } = null!;

        /// <summary>
        /// When set, enables deep search through the sub-tree of the LDAP URL + Search Base. True by default.
        /// </summary>
        [Input("subTree")]
        public Input<bool>? SubTree { get; set; }

        public LdapGroupSettingArgs()
        {
        }
        public static new LdapGroupSettingArgs Empty => new LdapGroupSettingArgs();
    }

    public sealed class LdapGroupSettingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An attribute on the group entry which denoting the group description. Used when importing groups.
        /// </summary>
        [Input("descriptionAttribute")]
        public Input<string>? DescriptionAttribute { get; set; }

        /// <summary>
        /// The LDAP filter used to search for group entries. Used for importing groups.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
        /// </summary>
        [Input("groupBaseDn")]
        public Input<string>? GroupBaseDn { get; set; }

        /// <summary>
        /// A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember,member).
        /// </summary>
        [Input("groupMemberAttribute")]
        public Input<string>? GroupMemberAttribute { get; set; }

        /// <summary>
        /// Attribute on the group entry denoting the group name. Used when importing groups.
        /// </summary>
        [Input("groupNameAttribute")]
        public Input<string>? GroupNameAttribute { get; set; }

        /// <summary>
        /// The LDAP setting key you want to use for group retrieval. The value for this field corresponds to 'enabledLdap' field of the ldap group setting XML block of system configuration.
        /// </summary>
        [Input("ldapSettingKey")]
        public Input<string>? LdapSettingKey { get; set; }

        /// <summary>
        /// Ldap group setting name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas:
        /// - STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN.
        /// - DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member.
        /// - HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org indicates that user1 belongs to two groups: uk and developers.
        /// </summary>
        [Input("strategy")]
        public Input<string>? Strategy { get; set; }

        /// <summary>
        /// When set, enables deep search through the sub-tree of the LDAP URL + Search Base. True by default.
        /// </summary>
        [Input("subTree")]
        public Input<bool>? SubTree { get; set; }

        public LdapGroupSettingState()
        {
        }
        public static new LdapGroupSettingState Empty => new LdapGroupSettingState();
    }
}
