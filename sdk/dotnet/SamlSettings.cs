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
    /// This resource can be used to manage Artifactory's SAML SSO settings.
    /// 
    /// Only a single `artifactory.SamlSettings` resource is meant to be defined.
    /// 
    /// ~&gt;The `artifactory.SamlSettings` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.
    /// 
    /// !&gt;This resource is deprecated in favor of platform_saml_settings resource in the Platform provider.
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
    ///     // Configure Artifactory SAML SSO settings
    ///     var saml = new Artifactory.SamlSettings("saml", new()
    ///     {
    ///         Enable = true,
    ///         ServiceProviderName = "okta",
    ///         LoginUrl = "test-login-url",
    ///         LogoutUrl = "test-logout-url",
    ///         Certificate = "test-certificate",
    ///         EmailAttribute = "email",
    ///         GroupAttribute = "groups",
    ///         NoAutoUserCreation = false,
    ///         AllowUserToAccessProfile = true,
    ///         AutoRedirect = true,
    ///         SyncGroups = true,
    ///         VerifyAudienceRestriction = true,
    ///         UseEncryptedAssertion = false,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Current SAML SSO settings can be imported using `saml_settings` as the `ID`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import artifactory:index/samlSettings:SamlSettings saml saml_settings
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/samlSettings:SamlSettings")]
    public partial class SamlSettings : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Allow persisted users to access their profile.  Default value is `true`.
        /// </summary>
        [Output("allowUserToAccessProfile")]
        public Output<bool?> AllowUserToAccessProfile { get; private set; } = null!;

        /// <summary>
        /// Auto redirect to login through the IdP when clicking on Artifactory's login link.  Default value is `false`.
        /// </summary>
        [Output("autoRedirect")]
        public Output<bool?> AutoRedirect { get; private set; } = null!;

        /// <summary>
        /// SAML certificate that contains the public key for the IdP service provider.  Used by Artifactory to verify sign-in requests. Default value is ``.
        /// </summary>
        [Output("certificate")]
        public Output<string?> Certificate { get; private set; } = null!;

        /// <summary>
        /// Name of the attribute in the SAML response from the IdP that contains the user's email. Default value is ``.
        /// </summary>
        [Output("emailAttribute")]
        public Output<string?> EmailAttribute { get; private set; } = null!;

        /// <summary>
        /// Enable SAML SSO.  Default value is `true`.
        /// </summary>
        [Output("enable")]
        public Output<bool?> Enable { get; private set; } = null!;

        /// <summary>
        /// Name of the attribute in the SAML response from the IdP that contains the user's group memberships. Default value is ``.
        /// </summary>
        [Output("groupAttribute")]
        public Output<string?> GroupAttribute { get; private set; } = null!;

        /// <summary>
        /// Service provider login url configured on the IdP.
        /// </summary>
        [Output("loginUrl")]
        public Output<string> LoginUrl { get; private set; } = null!;

        /// <summary>
        /// Service provider logout url, or where to redirect after user logs out.
        /// </summary>
        [Output("logoutUrl")]
        public Output<string> LogoutUrl { get; private set; } = null!;

        /// <summary>
        /// When automatic user creation is off, authenticated users are not automatically created inside Artifactory. Instead, for every request from an SSO user, the user is temporarily associated with default groups (if such groups are defined), and the permissions for these groups apply. Without auto-user creation, you must manually create the user inside Artifactory to manage user permissions not attached to their default groups. Default value is `false`.
        /// </summary>
        [Output("noAutoUserCreation")]
        public Output<bool?> NoAutoUserCreation { get; private set; } = null!;

        /// <summary>
        /// The SAML service provider name. This should be a URI that is also known as the entityID, providerID, or entity identity.
        /// </summary>
        [Output("serviceProviderName")]
        public Output<string> ServiceProviderName { get; private set; } = null!;

        /// <summary>
        /// Associate user with Artifactory groups based on the `group_attribute` provided in the SAML response from the identity provider.  Default value is `false`.
        /// </summary>
        [Output("syncGroups")]
        public Output<bool?> SyncGroups { get; private set; } = null!;

        /// <summary>
        /// When set, an X.509 public certificate will be created by Artifactory. Download this certificate and upload it to your IDP and choose your own encryption algorithm. This process will let you encrypt the assertion section in your SAML response. Default value is `false`.
        /// </summary>
        [Output("useEncryptedAssertion")]
        public Output<bool?> UseEncryptedAssertion { get; private set; } = null!;

        /// <summary>
        /// Enable "audience", or who the SAML assertion is intended for.  Ensures that the correct service provider intended for Artifactory is used on the IdP.  Default value is `true`.
        /// </summary>
        [Output("verifyAudienceRestriction")]
        public Output<bool?> VerifyAudienceRestriction { get; private set; } = null!;


        /// <summary>
        /// Create a SamlSettings resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SamlSettings(string name, SamlSettingsArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/samlSettings:SamlSettings", name, args ?? new SamlSettingsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SamlSettings(string name, Input<string> id, SamlSettingsState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/samlSettings:SamlSettings", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SamlSettings resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SamlSettings Get(string name, Input<string> id, SamlSettingsState? state = null, CustomResourceOptions? options = null)
        {
            return new SamlSettings(name, id, state, options);
        }
    }

    public sealed class SamlSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allow persisted users to access their profile.  Default value is `true`.
        /// </summary>
        [Input("allowUserToAccessProfile")]
        public Input<bool>? AllowUserToAccessProfile { get; set; }

        /// <summary>
        /// Auto redirect to login through the IdP when clicking on Artifactory's login link.  Default value is `false`.
        /// </summary>
        [Input("autoRedirect")]
        public Input<bool>? AutoRedirect { get; set; }

        /// <summary>
        /// SAML certificate that contains the public key for the IdP service provider.  Used by Artifactory to verify sign-in requests. Default value is ``.
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        /// <summary>
        /// Name of the attribute in the SAML response from the IdP that contains the user's email. Default value is ``.
        /// </summary>
        [Input("emailAttribute")]
        public Input<string>? EmailAttribute { get; set; }

        /// <summary>
        /// Enable SAML SSO.  Default value is `true`.
        /// </summary>
        [Input("enable")]
        public Input<bool>? Enable { get; set; }

        /// <summary>
        /// Name of the attribute in the SAML response from the IdP that contains the user's group memberships. Default value is ``.
        /// </summary>
        [Input("groupAttribute")]
        public Input<string>? GroupAttribute { get; set; }

        /// <summary>
        /// Service provider login url configured on the IdP.
        /// </summary>
        [Input("loginUrl", required: true)]
        public Input<string> LoginUrl { get; set; } = null!;

        /// <summary>
        /// Service provider logout url, or where to redirect after user logs out.
        /// </summary>
        [Input("logoutUrl", required: true)]
        public Input<string> LogoutUrl { get; set; } = null!;

        /// <summary>
        /// When automatic user creation is off, authenticated users are not automatically created inside Artifactory. Instead, for every request from an SSO user, the user is temporarily associated with default groups (if such groups are defined), and the permissions for these groups apply. Without auto-user creation, you must manually create the user inside Artifactory to manage user permissions not attached to their default groups. Default value is `false`.
        /// </summary>
        [Input("noAutoUserCreation")]
        public Input<bool>? NoAutoUserCreation { get; set; }

        /// <summary>
        /// The SAML service provider name. This should be a URI that is also known as the entityID, providerID, or entity identity.
        /// </summary>
        [Input("serviceProviderName", required: true)]
        public Input<string> ServiceProviderName { get; set; } = null!;

        /// <summary>
        /// Associate user with Artifactory groups based on the `group_attribute` provided in the SAML response from the identity provider.  Default value is `false`.
        /// </summary>
        [Input("syncGroups")]
        public Input<bool>? SyncGroups { get; set; }

        /// <summary>
        /// When set, an X.509 public certificate will be created by Artifactory. Download this certificate and upload it to your IDP and choose your own encryption algorithm. This process will let you encrypt the assertion section in your SAML response. Default value is `false`.
        /// </summary>
        [Input("useEncryptedAssertion")]
        public Input<bool>? UseEncryptedAssertion { get; set; }

        /// <summary>
        /// Enable "audience", or who the SAML assertion is intended for.  Ensures that the correct service provider intended for Artifactory is used on the IdP.  Default value is `true`.
        /// </summary>
        [Input("verifyAudienceRestriction")]
        public Input<bool>? VerifyAudienceRestriction { get; set; }

        public SamlSettingsArgs()
        {
        }
        public static new SamlSettingsArgs Empty => new SamlSettingsArgs();
    }

    public sealed class SamlSettingsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allow persisted users to access their profile.  Default value is `true`.
        /// </summary>
        [Input("allowUserToAccessProfile")]
        public Input<bool>? AllowUserToAccessProfile { get; set; }

        /// <summary>
        /// Auto redirect to login through the IdP when clicking on Artifactory's login link.  Default value is `false`.
        /// </summary>
        [Input("autoRedirect")]
        public Input<bool>? AutoRedirect { get; set; }

        /// <summary>
        /// SAML certificate that contains the public key for the IdP service provider.  Used by Artifactory to verify sign-in requests. Default value is ``.
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        /// <summary>
        /// Name of the attribute in the SAML response from the IdP that contains the user's email. Default value is ``.
        /// </summary>
        [Input("emailAttribute")]
        public Input<string>? EmailAttribute { get; set; }

        /// <summary>
        /// Enable SAML SSO.  Default value is `true`.
        /// </summary>
        [Input("enable")]
        public Input<bool>? Enable { get; set; }

        /// <summary>
        /// Name of the attribute in the SAML response from the IdP that contains the user's group memberships. Default value is ``.
        /// </summary>
        [Input("groupAttribute")]
        public Input<string>? GroupAttribute { get; set; }

        /// <summary>
        /// Service provider login url configured on the IdP.
        /// </summary>
        [Input("loginUrl")]
        public Input<string>? LoginUrl { get; set; }

        /// <summary>
        /// Service provider logout url, or where to redirect after user logs out.
        /// </summary>
        [Input("logoutUrl")]
        public Input<string>? LogoutUrl { get; set; }

        /// <summary>
        /// When automatic user creation is off, authenticated users are not automatically created inside Artifactory. Instead, for every request from an SSO user, the user is temporarily associated with default groups (if such groups are defined), and the permissions for these groups apply. Without auto-user creation, you must manually create the user inside Artifactory to manage user permissions not attached to their default groups. Default value is `false`.
        /// </summary>
        [Input("noAutoUserCreation")]
        public Input<bool>? NoAutoUserCreation { get; set; }

        /// <summary>
        /// The SAML service provider name. This should be a URI that is also known as the entityID, providerID, or entity identity.
        /// </summary>
        [Input("serviceProviderName")]
        public Input<string>? ServiceProviderName { get; set; }

        /// <summary>
        /// Associate user with Artifactory groups based on the `group_attribute` provided in the SAML response from the identity provider.  Default value is `false`.
        /// </summary>
        [Input("syncGroups")]
        public Input<bool>? SyncGroups { get; set; }

        /// <summary>
        /// When set, an X.509 public certificate will be created by Artifactory. Download this certificate and upload it to your IDP and choose your own encryption algorithm. This process will let you encrypt the assertion section in your SAML response. Default value is `false`.
        /// </summary>
        [Input("useEncryptedAssertion")]
        public Input<bool>? UseEncryptedAssertion { get; set; }

        /// <summary>
        /// Enable "audience", or who the SAML assertion is intended for.  Ensures that the correct service provider intended for Artifactory is used on the IdP.  Default value is `true`.
        /// </summary>
        [Input("verifyAudienceRestriction")]
        public Input<bool>? VerifyAudienceRestriction { get; set; }

        public SamlSettingsState()
        {
        }
        public static new SamlSettingsState Empty => new SamlSettingsState();
    }
}
