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
    ///     var test_user = new Artifactory.ManagedUser("test-user", new()
    ///     {
    ///         Name = "terraform",
    ///         Password = "my super secret password",
    ///         Email = "test-user@artifactory-terraform.com",
    ///         Groups = new[]
    ///         {
    ///             "readers",
    ///             "logged-in-users",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import artifactory:index/managedUser:ManagedUser test-user myusername
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/managedUser:ManagedUser")]
    public partial class ManagedUser : global::Pulumi.CustomResource
    {
        /// <summary>
        /// (Optional, Default: false) When enabled, this user is an administrator with all the ensuing privileges.
        /// </summary>
        [Output("admin")]
        public Output<bool> Admin { get; private set; } = null!;

        /// <summary>
        /// (Optional, Default: true) When enabled, this user can only access the system through the REST API. This option cannot be set if the user has Admin privileges.
        /// </summary>
        [Output("disableUiAccess")]
        public Output<bool> DisableUiAccess { get; private set; } = null!;

        /// <summary>
        /// Email for user.
        /// </summary>
        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        /// <summary>
        /// List of groups this user is a part of. **Notes:** If this attribute is not specified then user's group membership is set to empty. User will not be part of default "readers" group automatically.
        /// </summary>
        [Output("groups")]
        public Output<ImmutableArray<string>> Groups { get; private set; } = null!;

        /// <summary>
        /// (Optional, Default: false) When enabled, disables the fallback mechanism for using an internal password when external authentication (such as LDAP) is enabled.
        /// </summary>
        [Output("internalPasswordDisabled")]
        public Output<bool> InternalPasswordDisabled { get; private set; } = null!;

        /// <summary>
        /// Username for user. May contain lowercase letters, numbers and symbols: '.-_@'
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// (Optional, Sensitive) Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// (Optional, Default: true) When enabled, this user can update their profile details (except for the password. Only an administrator can update the password). There may be cases in which you want to leave this unset to prevent users from updating their profile. For example, a departmental user with a single password shared between all department members.
        /// </summary>
        [Output("profileUpdatable")]
        public Output<bool> ProfileUpdatable { get; private set; } = null!;


        /// <summary>
        /// Create a ManagedUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagedUser(string name, ManagedUserArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/managedUser:ManagedUser", name, args ?? new ManagedUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagedUser(string name, Input<string> id, ManagedUserState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/managedUser:ManagedUser", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ManagedUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagedUser Get(string name, Input<string> id, ManagedUserState? state = null, CustomResourceOptions? options = null)
        {
            return new ManagedUser(name, id, state, options);
        }
    }

    public sealed class ManagedUserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Optional, Default: false) When enabled, this user is an administrator with all the ensuing privileges.
        /// </summary>
        [Input("admin")]
        public Input<bool>? Admin { get; set; }

        /// <summary>
        /// (Optional, Default: true) When enabled, this user can only access the system through the REST API. This option cannot be set if the user has Admin privileges.
        /// </summary>
        [Input("disableUiAccess")]
        public Input<bool>? DisableUiAccess { get; set; }

        /// <summary>
        /// Email for user.
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// List of groups this user is a part of. **Notes:** If this attribute is not specified then user's group membership is set to empty. User will not be part of default "readers" group automatically.
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        /// <summary>
        /// (Optional, Default: false) When enabled, disables the fallback mechanism for using an internal password when external authentication (such as LDAP) is enabled.
        /// </summary>
        [Input("internalPasswordDisabled")]
        public Input<bool>? InternalPasswordDisabled { get; set; }

        /// <summary>
        /// Username for user. May contain lowercase letters, numbers and symbols: '.-_@'
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// (Optional, Sensitive) Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// (Optional, Default: true) When enabled, this user can update their profile details (except for the password. Only an administrator can update the password). There may be cases in which you want to leave this unset to prevent users from updating their profile. For example, a departmental user with a single password shared between all department members.
        /// </summary>
        [Input("profileUpdatable")]
        public Input<bool>? ProfileUpdatable { get; set; }

        public ManagedUserArgs()
        {
        }
        public static new ManagedUserArgs Empty => new ManagedUserArgs();
    }

    public sealed class ManagedUserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Optional, Default: false) When enabled, this user is an administrator with all the ensuing privileges.
        /// </summary>
        [Input("admin")]
        public Input<bool>? Admin { get; set; }

        /// <summary>
        /// (Optional, Default: true) When enabled, this user can only access the system through the REST API. This option cannot be set if the user has Admin privileges.
        /// </summary>
        [Input("disableUiAccess")]
        public Input<bool>? DisableUiAccess { get; set; }

        /// <summary>
        /// Email for user.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// List of groups this user is a part of. **Notes:** If this attribute is not specified then user's group membership is set to empty. User will not be part of default "readers" group automatically.
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        /// <summary>
        /// (Optional, Default: false) When enabled, disables the fallback mechanism for using an internal password when external authentication (such as LDAP) is enabled.
        /// </summary>
        [Input("internalPasswordDisabled")]
        public Input<bool>? InternalPasswordDisabled { get; set; }

        /// <summary>
        /// Username for user. May contain lowercase letters, numbers and symbols: '.-_@'
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// (Optional, Sensitive) Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// (Optional, Default: true) When enabled, this user can update their profile details (except for the password. Only an administrator can update the password). There may be cases in which you want to leave this unset to prevent users from updating their profile. For example, a departmental user with a single password shared between all department members.
        /// </summary>
        [Input("profileUpdatable")]
        public Input<bool>? ProfileUpdatable { get; set; }

        public ManagedUserState()
        {
        }
        public static new ManagedUserState Empty => new ManagedUserState();
    }
}
