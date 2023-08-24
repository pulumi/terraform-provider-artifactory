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
    /// Provides an Artifactory Scoped Token resource. This can be used to create and manage Artifactory Scoped Tokens.
    /// 
    /// !&gt;Scoped Tokens will be stored in the raw state as plain-text. Read more about sensitive data in
    /// state.
    /// 
    /// ~&gt;Token would not be saved by Artifactory if `expires_in` is less than the persistency threshold value (default to 10800 seconds) set in Access configuration. See [Persistency Threshold](https://jfrog.com/help/r/jfrog-platform-administration-documentation/using-the-revocable-and-persistency-thresholds) for details.
    /// 
    /// ## Example Usage
    /// 
    /// ### Create a new Artifactory scoped token for an existing user
    /// 
    /// resource "artifactory_scoped_token" "scoped_token" {
    ///   username = "existing-user"
    /// }
    /// 
    /// ### **Note:** This assumes that the user `existing-user` has already been created in Artifactory by different means, i.e. manually or in a separate pulumi up.
    /// 
    /// ### Create a new Artifactory user and scoped token
    /// resource "artifactory_user" "new_user" {
    ///   name   = "new_user"
    ///   email  = "new_user@somewhere.com"
    ///   groups = ["readers"]
    /// }
    /// 
    /// resource "artifactory_scoped_token" "scoped_token_user" {
    ///   username = artifactory_user.new_user.name
    /// }
    /// 
    /// ### Creates a new token for groups
    /// resource "artifactory_scoped_token" "scoped_token_group" {
    ///   scopes = ["applied-permissions/groups:readers"]
    /// }
    /// 
    /// ### Create token with expiry
    /// resource "artifactory_scoped_token" "scoped_token_no_expiry" {
    ///   username   = "existing-user"
    ///   expires_in = 7200 // in seconds
    /// }
    /// 
    /// ### Creates a refreshable token
    /// resource "artifactory_scoped_token" "scoped_token_refreshable" {
    ///   username    = "existing-user"
    ///   refreshable = true
    /// }
    /// 
    /// ### Creates an administrator token
    /// resource "artifactory_scoped_token" "admin" {
    ///   username = "admin-user"
    ///   scopes   = ["applied-permissions/admin"]
    /// }
    /// ## References
    /// 
    /// - https://jfrog.com/help/r/jfrog-platform-administration-documentation/access-tokens
    /// - https://jfrog.com/help/r/jfrog-rest-apis/access-tokens
    /// 
    /// ## Import
    /// 
    /// Artifactory **does not** retain scoped tokens, and they cannot be imported into state.
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/scopedToken:ScopedToken")]
    public partial class ScopedToken : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Returns the access token to authenticate to Artifactory.
        /// </summary>
        [Output("accessToken")]
        public Output<string> AccessToken { get; private set; } = null!;

        /// <summary>
        /// A list of the other instances or services that should accept this token identified by their Service-IDs. Limited to total 255 characters. Default to '*@*' if not set. Service ID must begin with valid JFrog service type. Options: jfrt, jfxr, jfpip, jfds, jfmc, jfac, jfevt, jfmd, jfcon, or *. For instructions to retrieve the Artifactory Service ID see this [documentation](https://jfrog.com/help/r/jfrog-rest-apis/get-service-id)
        /// </summary>
        [Output("audiences")]
        public Output<ImmutableArray<string>> Audiences { get; private set; } = null!;

        /// <summary>
        /// Free text token description. Useful for filtering and managing tokens. Limited to 1024 characters.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The amount of time, in seconds, it would take for the token to expire. An admin shall be able to set whether expiry is mandatory, what is the default expiry, and what is the maximum expiry allowed. Must be non-negative. Default value is based on configuration in 'access.config.yaml'. See [API documentation](https://jfrog.com/help/r/jfrog-rest-apis/revoke-token-by-id) for details. Access Token would not be saved by Artifactory if this is less than the persistence threshold value (default to 10800 seconds) set in Access configuration. See [official documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/using-the-revocable-and-persistency-thresholds) for details.
        /// </summary>
        [Output("expiresIn")]
        public Output<int> ExpiresIn { get; private set; } = null!;

        /// <summary>
        /// Returns the token expiry.
        /// </summary>
        [Output("expiry")]
        public Output<int> Expiry { get; private set; } = null!;

        /// <summary>
        /// The grant type used to authenticate the request. In this case, the only value supported is `client_credentials` which is also the default value if this parameter is not specified.
        /// </summary>
        [Output("grantType")]
        public Output<string> GrantType { get; private set; } = null!;

        /// <summary>
        /// Also create a reference token which can be used like an API key.
        /// </summary>
        [Output("includeReferenceToken")]
        public Output<bool> IncludeReferenceToken { get; private set; } = null!;

        /// <summary>
        /// Returns the token issued at date/time.
        /// </summary>
        [Output("issuedAt")]
        public Output<int> IssuedAt { get; private set; } = null!;

        /// <summary>
        /// Returns the token issuer.
        /// </summary>
        [Output("issuer")]
        public Output<string> Issuer { get; private set; } = null!;

        /// <summary>
        /// The project for which this token is created. Enter the project name on which you want to apply this token.
        /// </summary>
        [Output("projectKey")]
        public Output<string?> ProjectKey { get; private set; } = null!;

        /// <summary>
        /// Reference Token (alias to Access Token).
        /// </summary>
        [Output("referenceToken")]
        public Output<string> ReferenceToken { get; private set; } = null!;

        /// <summary>
        /// Refresh token.
        /// </summary>
        [Output("refreshToken")]
        public Output<string> RefreshToken { get; private set; } = null!;

        /// <summary>
        /// Is this token refreshable? Default is `false`.
        /// </summary>
        [Output("refreshable")]
        public Output<bool> Refreshable { get; private set; } = null!;

        /// <summary>
        /// The scope of access that the token provides. Access to the REST API is always provided by default. Administrators can set any scope, while non-admin users can only set the scope to a subset of the groups to which they belong.
        /// The supported scopes include:
        /// </summary>
        [Output("scopes")]
        public Output<ImmutableArray<string>> Scopes { get; private set; } = null!;

        /// <summary>
        /// Returns the token type.
        /// </summary>
        [Output("subject")]
        public Output<string> Subject { get; private set; } = null!;

        /// <summary>
        /// Returns the token type.
        /// </summary>
        [Output("tokenType")]
        public Output<string> TokenType { get; private set; } = null!;

        /// <summary>
        /// The user name for which this token is created. The username is based on the authenticated user - either from the user of the authenticated token or based on the username (if basic auth was used). The username is then used to set the subject of the token: \n\n/users/\n\n. Limited to 255 characters.
        /// </summary>
        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;


        /// <summary>
        /// Create a ScopedToken resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ScopedToken(string name, ScopedTokenArgs? args = null, CustomResourceOptions? options = null)
            : base("artifactory:index/scopedToken:ScopedToken", name, args ?? new ScopedTokenArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ScopedToken(string name, Input<string> id, ScopedTokenState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/scopedToken:ScopedToken", name, state, MakeResourceOptions(options, id))
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
                    "referenceToken",
                    "refreshToken",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ScopedToken resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ScopedToken Get(string name, Input<string> id, ScopedTokenState? state = null, CustomResourceOptions? options = null)
        {
            return new ScopedToken(name, id, state, options);
        }
    }

    public sealed class ScopedTokenArgs : global::Pulumi.ResourceArgs
    {
        [Input("audiences")]
        private InputList<string>? _audiences;

        /// <summary>
        /// A list of the other instances or services that should accept this token identified by their Service-IDs. Limited to total 255 characters. Default to '*@*' if not set. Service ID must begin with valid JFrog service type. Options: jfrt, jfxr, jfpip, jfds, jfmc, jfac, jfevt, jfmd, jfcon, or *. For instructions to retrieve the Artifactory Service ID see this [documentation](https://jfrog.com/help/r/jfrog-rest-apis/get-service-id)
        /// </summary>
        public InputList<string> Audiences
        {
            get => _audiences ?? (_audiences = new InputList<string>());
            set => _audiences = value;
        }

        /// <summary>
        /// Free text token description. Useful for filtering and managing tokens. Limited to 1024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The amount of time, in seconds, it would take for the token to expire. An admin shall be able to set whether expiry is mandatory, what is the default expiry, and what is the maximum expiry allowed. Must be non-negative. Default value is based on configuration in 'access.config.yaml'. See [API documentation](https://jfrog.com/help/r/jfrog-rest-apis/revoke-token-by-id) for details. Access Token would not be saved by Artifactory if this is less than the persistence threshold value (default to 10800 seconds) set in Access configuration. See [official documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/using-the-revocable-and-persistency-thresholds) for details.
        /// </summary>
        [Input("expiresIn")]
        public Input<int>? ExpiresIn { get; set; }

        /// <summary>
        /// The grant type used to authenticate the request. In this case, the only value supported is `client_credentials` which is also the default value if this parameter is not specified.
        /// </summary>
        [Input("grantType")]
        public Input<string>? GrantType { get; set; }

        /// <summary>
        /// Also create a reference token which can be used like an API key.
        /// </summary>
        [Input("includeReferenceToken")]
        public Input<bool>? IncludeReferenceToken { get; set; }

        /// <summary>
        /// The project for which this token is created. Enter the project name on which you want to apply this token.
        /// </summary>
        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        /// <summary>
        /// Is this token refreshable? Default is `false`.
        /// </summary>
        [Input("refreshable")]
        public Input<bool>? Refreshable { get; set; }

        [Input("scopes")]
        private InputList<string>? _scopes;

        /// <summary>
        /// The scope of access that the token provides. Access to the REST API is always provided by default. Administrators can set any scope, while non-admin users can only set the scope to a subset of the groups to which they belong.
        /// The supported scopes include:
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        /// <summary>
        /// The user name for which this token is created. The username is based on the authenticated user - either from the user of the authenticated token or based on the username (if basic auth was used). The username is then used to set the subject of the token: \n\n/users/\n\n. Limited to 255 characters.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ScopedTokenArgs()
        {
        }
        public static new ScopedTokenArgs Empty => new ScopedTokenArgs();
    }

    public sealed class ScopedTokenState : global::Pulumi.ResourceArgs
    {
        [Input("accessToken")]
        private Input<string>? _accessToken;

        /// <summary>
        /// Returns the access token to authenticate to Artifactory.
        /// </summary>
        public Input<string>? AccessToken
        {
            get => _accessToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("audiences")]
        private InputList<string>? _audiences;

        /// <summary>
        /// A list of the other instances or services that should accept this token identified by their Service-IDs. Limited to total 255 characters. Default to '*@*' if not set. Service ID must begin with valid JFrog service type. Options: jfrt, jfxr, jfpip, jfds, jfmc, jfac, jfevt, jfmd, jfcon, or *. For instructions to retrieve the Artifactory Service ID see this [documentation](https://jfrog.com/help/r/jfrog-rest-apis/get-service-id)
        /// </summary>
        public InputList<string> Audiences
        {
            get => _audiences ?? (_audiences = new InputList<string>());
            set => _audiences = value;
        }

        /// <summary>
        /// Free text token description. Useful for filtering and managing tokens. Limited to 1024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The amount of time, in seconds, it would take for the token to expire. An admin shall be able to set whether expiry is mandatory, what is the default expiry, and what is the maximum expiry allowed. Must be non-negative. Default value is based on configuration in 'access.config.yaml'. See [API documentation](https://jfrog.com/help/r/jfrog-rest-apis/revoke-token-by-id) for details. Access Token would not be saved by Artifactory if this is less than the persistence threshold value (default to 10800 seconds) set in Access configuration. See [official documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/using-the-revocable-and-persistency-thresholds) for details.
        /// </summary>
        [Input("expiresIn")]
        public Input<int>? ExpiresIn { get; set; }

        /// <summary>
        /// Returns the token expiry.
        /// </summary>
        [Input("expiry")]
        public Input<int>? Expiry { get; set; }

        /// <summary>
        /// The grant type used to authenticate the request. In this case, the only value supported is `client_credentials` which is also the default value if this parameter is not specified.
        /// </summary>
        [Input("grantType")]
        public Input<string>? GrantType { get; set; }

        /// <summary>
        /// Also create a reference token which can be used like an API key.
        /// </summary>
        [Input("includeReferenceToken")]
        public Input<bool>? IncludeReferenceToken { get; set; }

        /// <summary>
        /// Returns the token issued at date/time.
        /// </summary>
        [Input("issuedAt")]
        public Input<int>? IssuedAt { get; set; }

        /// <summary>
        /// Returns the token issuer.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        /// <summary>
        /// The project for which this token is created. Enter the project name on which you want to apply this token.
        /// </summary>
        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        [Input("referenceToken")]
        private Input<string>? _referenceToken;

        /// <summary>
        /// Reference Token (alias to Access Token).
        /// </summary>
        public Input<string>? ReferenceToken
        {
            get => _referenceToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _referenceToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("refreshToken")]
        private Input<string>? _refreshToken;

        /// <summary>
        /// Refresh token.
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
        /// Is this token refreshable? Default is `false`.
        /// </summary>
        [Input("refreshable")]
        public Input<bool>? Refreshable { get; set; }

        [Input("scopes")]
        private InputList<string>? _scopes;

        /// <summary>
        /// The scope of access that the token provides. Access to the REST API is always provided by default. Administrators can set any scope, while non-admin users can only set the scope to a subset of the groups to which they belong.
        /// The supported scopes include:
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        /// <summary>
        /// Returns the token type.
        /// </summary>
        [Input("subject")]
        public Input<string>? Subject { get; set; }

        /// <summary>
        /// Returns the token type.
        /// </summary>
        [Input("tokenType")]
        public Input<string>? TokenType { get; set; }

        /// <summary>
        /// The user name for which this token is created. The username is based on the authenticated user - either from the user of the authenticated token or based on the username (if basic auth was used). The username is then used to set the subject of the token: \n\n/users/\n\n. Limited to 255 characters.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ScopedTokenState()
        {
        }
        public static new ScopedTokenState Empty => new ScopedTokenState();
    }
}
