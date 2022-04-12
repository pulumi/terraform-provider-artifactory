// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class SamlSettings extends pulumi.CustomResource {
    /**
     * Get an existing SamlSettings resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SamlSettingsState, opts?: pulumi.CustomResourceOptions): SamlSettings {
        return new SamlSettings(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/samlSettings:SamlSettings';

    /**
     * Returns true if the given object is an instance of SamlSettings.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SamlSettings {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SamlSettings.__pulumiType;
    }

    /**
     * (Optional) Allow persisted users to access their profile. Default value is "true".
     */
    public readonly allowUserToAccessProfile!: pulumi.Output<boolean | undefined>;
    /**
     * (Optional) Auto redirect to login through the IdP when clicking on Artifactory's login link. Default value is "false".
     */
    public readonly autoRedirect!: pulumi.Output<boolean | undefined>;
    /**
     * (Optional) SAML certificate that contains the public key for the IdP service provider. Used by Artifactory to verify
     * sign-in requests. Default value is "".
     */
    public readonly certificate!: pulumi.Output<string | undefined>;
    /**
     * (Optional) Name of the attribute in the SAML response from the IdP that contains the user's email. Default value is "".
     */
    public readonly emailAttribute!: pulumi.Output<string | undefined>;
    /**
     * (Optional) Enable SAML SSO. Default value is "true".
     */
    public readonly enable!: pulumi.Output<boolean | undefined>;
    /**
     * (Optional) Name of the attribute in the SAML response from the IdP that contains the user's group memberships. Default
     * value is "".
     */
    public readonly groupAttribute!: pulumi.Output<string | undefined>;
    /**
     * (Required) Service provider login url configured on the IdP.
     */
    public readonly loginUrl!: pulumi.Output<string>;
    /**
     * (Required) Service provider logout url, or where to redirect after user logs out.
     */
    public readonly logoutUrl!: pulumi.Output<string>;
    /**
     * (Optional) When automatic user creation is off, authenticated users are not automatically created inside Artifactory.
     * Instead, for every request from an SSO user, the user is temporarily associated with default groups (if such groups are
     * defined), and the permissions for these groups apply. Without auto-user creation, you must manually create the user
     * inside Artifactory to manage user permissions not attached to their default groups. Default value is "false".
     */
    public readonly noAutoUserCreation!: pulumi.Output<boolean | undefined>;
    /**
     * (Required) The SAML service provider name. This should be a URI that is also known as the entityID, providerID, or
     * entity identity.
     */
    public readonly serviceProviderName!: pulumi.Output<string>;
    /**
     * (Optional) Associate user with Artifactory groups based on the "group_attribute" provided in the SAML response from the
     * identity provider. Default value is "false".
     */
    public readonly syncGroups!: pulumi.Output<boolean | undefined>;
    /**
     * (Optional) When set, an X.509 public certificate will be created by Artifactory. Download this certificate and upload it
     * to your IDP and choose your own encryption algorithm. This process will let you encrypt the assertion section in your
     * SAML response. Default value is "false".
     */
    public readonly useEncryptedAssertion!: pulumi.Output<boolean | undefined>;
    /**
     * (Optional) Enable "audience", or who the SAML assertion is intended for. Ensures that the correct service provider
     * intended for Artifactory is used on the IdP. Default value is "true".
     */
    public readonly verifyAudienceRestriction!: pulumi.Output<boolean | undefined>;

    /**
     * Create a SamlSettings resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SamlSettingsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SamlSettingsArgs | SamlSettingsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SamlSettingsState | undefined;
            resourceInputs["allowUserToAccessProfile"] = state ? state.allowUserToAccessProfile : undefined;
            resourceInputs["autoRedirect"] = state ? state.autoRedirect : undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["emailAttribute"] = state ? state.emailAttribute : undefined;
            resourceInputs["enable"] = state ? state.enable : undefined;
            resourceInputs["groupAttribute"] = state ? state.groupAttribute : undefined;
            resourceInputs["loginUrl"] = state ? state.loginUrl : undefined;
            resourceInputs["logoutUrl"] = state ? state.logoutUrl : undefined;
            resourceInputs["noAutoUserCreation"] = state ? state.noAutoUserCreation : undefined;
            resourceInputs["serviceProviderName"] = state ? state.serviceProviderName : undefined;
            resourceInputs["syncGroups"] = state ? state.syncGroups : undefined;
            resourceInputs["useEncryptedAssertion"] = state ? state.useEncryptedAssertion : undefined;
            resourceInputs["verifyAudienceRestriction"] = state ? state.verifyAudienceRestriction : undefined;
        } else {
            const args = argsOrState as SamlSettingsArgs | undefined;
            if ((!args || args.loginUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'loginUrl'");
            }
            if ((!args || args.logoutUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'logoutUrl'");
            }
            if ((!args || args.serviceProviderName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceProviderName'");
            }
            resourceInputs["allowUserToAccessProfile"] = args ? args.allowUserToAccessProfile : undefined;
            resourceInputs["autoRedirect"] = args ? args.autoRedirect : undefined;
            resourceInputs["certificate"] = args ? args.certificate : undefined;
            resourceInputs["emailAttribute"] = args ? args.emailAttribute : undefined;
            resourceInputs["enable"] = args ? args.enable : undefined;
            resourceInputs["groupAttribute"] = args ? args.groupAttribute : undefined;
            resourceInputs["loginUrl"] = args ? args.loginUrl : undefined;
            resourceInputs["logoutUrl"] = args ? args.logoutUrl : undefined;
            resourceInputs["noAutoUserCreation"] = args ? args.noAutoUserCreation : undefined;
            resourceInputs["serviceProviderName"] = args ? args.serviceProviderName : undefined;
            resourceInputs["syncGroups"] = args ? args.syncGroups : undefined;
            resourceInputs["useEncryptedAssertion"] = args ? args.useEncryptedAssertion : undefined;
            resourceInputs["verifyAudienceRestriction"] = args ? args.verifyAudienceRestriction : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SamlSettings.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SamlSettings resources.
 */
export interface SamlSettingsState {
    /**
     * (Optional) Allow persisted users to access their profile. Default value is "true".
     */
    allowUserToAccessProfile?: pulumi.Input<boolean>;
    /**
     * (Optional) Auto redirect to login through the IdP when clicking on Artifactory's login link. Default value is "false".
     */
    autoRedirect?: pulumi.Input<boolean>;
    /**
     * (Optional) SAML certificate that contains the public key for the IdP service provider. Used by Artifactory to verify
     * sign-in requests. Default value is "".
     */
    certificate?: pulumi.Input<string>;
    /**
     * (Optional) Name of the attribute in the SAML response from the IdP that contains the user's email. Default value is "".
     */
    emailAttribute?: pulumi.Input<string>;
    /**
     * (Optional) Enable SAML SSO. Default value is "true".
     */
    enable?: pulumi.Input<boolean>;
    /**
     * (Optional) Name of the attribute in the SAML response from the IdP that contains the user's group memberships. Default
     * value is "".
     */
    groupAttribute?: pulumi.Input<string>;
    /**
     * (Required) Service provider login url configured on the IdP.
     */
    loginUrl?: pulumi.Input<string>;
    /**
     * (Required) Service provider logout url, or where to redirect after user logs out.
     */
    logoutUrl?: pulumi.Input<string>;
    /**
     * (Optional) When automatic user creation is off, authenticated users are not automatically created inside Artifactory.
     * Instead, for every request from an SSO user, the user is temporarily associated with default groups (if such groups are
     * defined), and the permissions for these groups apply. Without auto-user creation, you must manually create the user
     * inside Artifactory to manage user permissions not attached to their default groups. Default value is "false".
     */
    noAutoUserCreation?: pulumi.Input<boolean>;
    /**
     * (Required) The SAML service provider name. This should be a URI that is also known as the entityID, providerID, or
     * entity identity.
     */
    serviceProviderName?: pulumi.Input<string>;
    /**
     * (Optional) Associate user with Artifactory groups based on the "group_attribute" provided in the SAML response from the
     * identity provider. Default value is "false".
     */
    syncGroups?: pulumi.Input<boolean>;
    /**
     * (Optional) When set, an X.509 public certificate will be created by Artifactory. Download this certificate and upload it
     * to your IDP and choose your own encryption algorithm. This process will let you encrypt the assertion section in your
     * SAML response. Default value is "false".
     */
    useEncryptedAssertion?: pulumi.Input<boolean>;
    /**
     * (Optional) Enable "audience", or who the SAML assertion is intended for. Ensures that the correct service provider
     * intended for Artifactory is used on the IdP. Default value is "true".
     */
    verifyAudienceRestriction?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a SamlSettings resource.
 */
export interface SamlSettingsArgs {
    /**
     * (Optional) Allow persisted users to access their profile. Default value is "true".
     */
    allowUserToAccessProfile?: pulumi.Input<boolean>;
    /**
     * (Optional) Auto redirect to login through the IdP when clicking on Artifactory's login link. Default value is "false".
     */
    autoRedirect?: pulumi.Input<boolean>;
    /**
     * (Optional) SAML certificate that contains the public key for the IdP service provider. Used by Artifactory to verify
     * sign-in requests. Default value is "".
     */
    certificate?: pulumi.Input<string>;
    /**
     * (Optional) Name of the attribute in the SAML response from the IdP that contains the user's email. Default value is "".
     */
    emailAttribute?: pulumi.Input<string>;
    /**
     * (Optional) Enable SAML SSO. Default value is "true".
     */
    enable?: pulumi.Input<boolean>;
    /**
     * (Optional) Name of the attribute in the SAML response from the IdP that contains the user's group memberships. Default
     * value is "".
     */
    groupAttribute?: pulumi.Input<string>;
    /**
     * (Required) Service provider login url configured on the IdP.
     */
    loginUrl: pulumi.Input<string>;
    /**
     * (Required) Service provider logout url, or where to redirect after user logs out.
     */
    logoutUrl: pulumi.Input<string>;
    /**
     * (Optional) When automatic user creation is off, authenticated users are not automatically created inside Artifactory.
     * Instead, for every request from an SSO user, the user is temporarily associated with default groups (if such groups are
     * defined), and the permissions for these groups apply. Without auto-user creation, you must manually create the user
     * inside Artifactory to manage user permissions not attached to their default groups. Default value is "false".
     */
    noAutoUserCreation?: pulumi.Input<boolean>;
    /**
     * (Required) The SAML service provider name. This should be a URI that is also known as the entityID, providerID, or
     * entity identity.
     */
    serviceProviderName: pulumi.Input<string>;
    /**
     * (Optional) Associate user with Artifactory groups based on the "group_attribute" provided in the SAML response from the
     * identity provider. Default value is "false".
     */
    syncGroups?: pulumi.Input<boolean>;
    /**
     * (Optional) When set, an X.509 public certificate will be created by Artifactory. Download this certificate and upload it
     * to your IDP and choose your own encryption algorithm. This process will let you encrypt the assertion section in your
     * SAML response. Default value is "false".
     */
    useEncryptedAssertion?: pulumi.Input<boolean>;
    /**
     * (Optional) Enable "audience", or who the SAML assertion is intended for. Ensures that the correct service provider
     * intended for Artifactory is used on the IdP. Default value is "true".
     */
    verifyAudienceRestriction?: pulumi.Input<boolean>;
}
