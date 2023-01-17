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
    /// ## Import
    /// 
    /// Artifactory **does not** retain access tokens and cannot be imported into state.
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/accessToken:AccessToken")]
    public partial class AccessToken : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Returns the access token to authenciate to Artifactory
        /// </summary>
        [Output("accessToken")]
        public Output<string> Details { get; private set; } = null!;

        /// <summary>
        /// (Optional) Specify the `instance_id` in this block to grant this token admin privileges. This can only be created when the authenticated user is an admin. `admin_token` cannot be specified with `groups`.
        /// </summary>
        [Output("adminToken")]
        public Output<Outputs.AccessTokenAdminToken?> AdminToken { get; private set; } = null!;

        /// <summary>
        /// (Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set `"jfrt@*"` so the token to be accepted by all Artifactory instances.
        /// </summary>
        [Output("audience")]
        public Output<string?> Audience { get; private set; } = null!;

        /// <summary>
        /// (Optional) The end date which the token is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
        /// </summary>
        [Output("endDate")]
        public Output<string> EndDate { get; private set; } = null!;

        /// <summary>
        /// (Optional) A relative duration for which the token is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "s", "m", "h".
        /// </summary>
        [Output("endDateRelative")]
        public Output<string?> EndDateRelative { get; private set; } = null!;

        /// <summary>
        /// (Optional) List of groups. The token is granted access based on the permissions of the groups. Specify `["*"]` for all groups that the user belongs to. `groups` cannot be specified with `admin_token`.
        /// </summary>
        [Output("groups")]
        public Output<ImmutableArray<string>> Groups { get; private set; } = null!;

        /// <summary>
        /// Returns the refresh token when `refreshable` is true, or an empty string when `refreshable` is false.
        /// </summary>
        [Output("refreshToken")]
        public Output<string> RefreshToken { get; private set; } = null!;

        /// <summary>
        /// (Optional) Is this token refreshable? Defaults to `false`
        /// </summary>
        [Output("refreshable")]
        public Output<bool?> Refreshable { get; private set; } = null!;

        /// <summary>
        /// (Required) The username or subject for the token. A non-admin can only specify their own username. Admins can specify any existing username, or a new name for a temporary token. Temporary tokens require `groups` to be set.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a AccessToken resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccessToken(string name, AccessTokenArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/accessToken:AccessToken", name, args ?? new AccessTokenArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccessToken(string name, Input<string> id, AccessTokenState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/accessToken:AccessToken", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "accessToken",
                    "refreshToken",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AccessToken resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccessToken Get(string name, Input<string> id, AccessTokenState? state = null, CustomResourceOptions? options = null)
        {
            return new AccessToken(name, id, state, options);
        }
    }

    public sealed class AccessTokenArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Optional) Specify the `instance_id` in this block to grant this token admin privileges. This can only be created when the authenticated user is an admin. `admin_token` cannot be specified with `groups`.
        /// </summary>
        [Input("adminToken")]
        public Input<Inputs.AccessTokenAdminTokenArgs>? AdminToken { get; set; }

        /// <summary>
        /// (Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set `"jfrt@*"` so the token to be accepted by all Artifactory instances.
        /// </summary>
        [Input("audience")]
        public Input<string>? Audience { get; set; }

        /// <summary>
        /// (Optional) The end date which the token is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
        /// </summary>
        [Input("endDate")]
        public Input<string>? EndDate { get; set; }

        /// <summary>
        /// (Optional) A relative duration for which the token is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "s", "m", "h".
        /// </summary>
        [Input("endDateRelative")]
        public Input<string>? EndDateRelative { get; set; }

        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// (Optional) List of groups. The token is granted access based on the permissions of the groups. Specify `["*"]` for all groups that the user belongs to. `groups` cannot be specified with `admin_token`.
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        /// <summary>
        /// (Optional) Is this token refreshable? Defaults to `false`
        /// </summary>
        [Input("refreshable")]
        public Input<bool>? Refreshable { get; set; }

        /// <summary>
        /// (Required) The username or subject for the token. A non-admin can only specify their own username. Admins can specify any existing username, or a new name for a temporary token. Temporary tokens require `groups` to be set.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public AccessTokenArgs()
        {
        }
        public static new AccessTokenArgs Empty => new AccessTokenArgs();
    }

    public sealed class AccessTokenState : global::Pulumi.ResourceArgs
    {
        [Input("accessToken")]
        private Input<string>? _accessToken;

        /// <summary>
        /// Returns the access token to authenciate to Artifactory
        /// </summary>
        public Input<string>? Details
        {
            get => _accessToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// (Optional) Specify the `instance_id` in this block to grant this token admin privileges. This can only be created when the authenticated user is an admin. `admin_token` cannot be specified with `groups`.
        /// </summary>
        [Input("adminToken")]
        public Input<Inputs.AccessTokenAdminTokenGetArgs>? AdminToken { get; set; }

        /// <summary>
        /// (Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set `"jfrt@*"` so the token to be accepted by all Artifactory instances.
        /// </summary>
        [Input("audience")]
        public Input<string>? Audience { get; set; }

        /// <summary>
        /// (Optional) The end date which the token is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
        /// </summary>
        [Input("endDate")]
        public Input<string>? EndDate { get; set; }

        /// <summary>
        /// (Optional) A relative duration for which the token is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "s", "m", "h".
        /// </summary>
        [Input("endDateRelative")]
        public Input<string>? EndDateRelative { get; set; }

        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// (Optional) List of groups. The token is granted access based on the permissions of the groups. Specify `["*"]` for all groups that the user belongs to. `groups` cannot be specified with `admin_token`.
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        [Input("refreshToken")]
        private Input<string>? _refreshToken;

        /// <summary>
        /// Returns the refresh token when `refreshable` is true, or an empty string when `refreshable` is false.
        /// </summary>
        public Input<string>? RefreshToken
        {
            get => _refreshToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _refreshToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// (Optional) Is this token refreshable? Defaults to `false`
        /// </summary>
        [Input("refreshable")]
        public Input<bool>? Refreshable { get; set; }

        /// <summary>
        /// (Required) The username or subject for the token. A non-admin can only specify their own username. Admins can specify any existing username, or a new name for a temporary token. Temporary tokens require `groups` to be set.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public AccessTokenState()
        {
        }
        public static new AccessTokenState Empty => new AccessTokenState();
    }
}
