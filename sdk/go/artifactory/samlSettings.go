// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory SAML SSO Settings Resource
//
// This resource can be used to manage Artifactory's SAML SSO settings.
//
// Only a single `SamlSettings` resource is meant to be defined.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-artifactory/sdk/go/artifactory"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := artifactory.NewSamlSettings(ctx, "saml", &artifactory.SamlSettingsArgs{
// 			AllowUserToAccessProfile:  pulumi.Bool(true),
// 			AutoRedirect:              pulumi.Bool(true),
// 			Certificate:               pulumi.String("test-certificate"),
// 			EmailAttribute:            pulumi.String("email"),
// 			Enable:                    pulumi.Bool(true),
// 			GroupAttribute:            pulumi.String("groups"),
// 			LoginUrl:                  pulumi.String("test-login-url"),
// 			LogoutUrl:                 pulumi.String("test-logout-url"),
// 			NoAutoUserCreation:        pulumi.Bool(false),
// 			ServiceProviderName:       pulumi.String("okta"),
// 			SyncGroups:                pulumi.Bool(true),
// 			VerifyAudienceRestriction: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Current SAML SSO settings can be imported using `saml_settings` as the `ID`, e.g.
//
// ```sh
//  $ pulumi import artifactory:index/samlSettings:SamlSettings saml saml_settings
// ```
type SamlSettings struct {
	pulumi.CustomResourceState

	// Allow persisted users to access their profile.  Default value is `true`.
	AllowUserToAccessProfile pulumi.BoolPtrOutput `pulumi:"allowUserToAccessProfile"`
	// Auto redirect to login through the IdP when clicking on Artifactory's login link.  Default value is `false`.
	AutoRedirect pulumi.BoolPtrOutput `pulumi:"autoRedirect"`
	// SAML certificate that contains the public key for the IdP service provider.  Used by Artifactory to verify sign-in requests.
	Certificate pulumi.StringPtrOutput `pulumi:"certificate"`
	// Name of the attribute in the SAML response from the IdP that contains the user's email.
	EmailAttribute pulumi.StringPtrOutput `pulumi:"emailAttribute"`
	// Enable SAML SSO.  Default value is `true`.
	Enable pulumi.BoolPtrOutput `pulumi:"enable"`
	// Name of the attribute in the SAML response from the IdP that contains the user's group memberships.
	GroupAttribute pulumi.StringPtrOutput `pulumi:"groupAttribute"`
	// Service provider login url configured on the IdP.
	LoginUrl pulumi.StringOutput `pulumi:"loginUrl"`
	// Service provider logout url, or where to redirect after user logs out.
	LogoutUrl pulumi.StringOutput `pulumi:"logoutUrl"`
	// Enable the creation of local Artifactory users.  Default value is `false`.
	NoAutoUserCreation pulumi.BoolPtrOutput `pulumi:"noAutoUserCreation"`
	// Name of the service provider configured on the .
	ServiceProviderName pulumi.StringOutput `pulumi:"serviceProviderName"`
	// Associate user with Artifactory groups based on the `groupAttribute` provided in the SAML response from the identity provider.  Default value is `false`.
	SyncGroups pulumi.BoolPtrOutput `pulumi:"syncGroups"`
	// Enable "audience", or who the SAML assertion is intended for.  Ensures that the correct service provider intended for Artifactory is used on the IdP.  Default value is `true`.
	VerifyAudienceRestriction pulumi.BoolPtrOutput `pulumi:"verifyAudienceRestriction"`
}

// NewSamlSettings registers a new resource with the given unique name, arguments, and options.
func NewSamlSettings(ctx *pulumi.Context,
	name string, args *SamlSettingsArgs, opts ...pulumi.ResourceOption) (*SamlSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LoginUrl == nil {
		return nil, errors.New("invalid value for required argument 'LoginUrl'")
	}
	if args.LogoutUrl == nil {
		return nil, errors.New("invalid value for required argument 'LogoutUrl'")
	}
	if args.ServiceProviderName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceProviderName'")
	}
	var resource SamlSettings
	err := ctx.RegisterResource("artifactory:index/samlSettings:SamlSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSamlSettings gets an existing SamlSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSamlSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SamlSettingsState, opts ...pulumi.ResourceOption) (*SamlSettings, error) {
	var resource SamlSettings
	err := ctx.ReadResource("artifactory:index/samlSettings:SamlSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SamlSettings resources.
type samlSettingsState struct {
	// Allow persisted users to access their profile.  Default value is `true`.
	AllowUserToAccessProfile *bool `pulumi:"allowUserToAccessProfile"`
	// Auto redirect to login through the IdP when clicking on Artifactory's login link.  Default value is `false`.
	AutoRedirect *bool `pulumi:"autoRedirect"`
	// SAML certificate that contains the public key for the IdP service provider.  Used by Artifactory to verify sign-in requests.
	Certificate *string `pulumi:"certificate"`
	// Name of the attribute in the SAML response from the IdP that contains the user's email.
	EmailAttribute *string `pulumi:"emailAttribute"`
	// Enable SAML SSO.  Default value is `true`.
	Enable *bool `pulumi:"enable"`
	// Name of the attribute in the SAML response from the IdP that contains the user's group memberships.
	GroupAttribute *string `pulumi:"groupAttribute"`
	// Service provider login url configured on the IdP.
	LoginUrl *string `pulumi:"loginUrl"`
	// Service provider logout url, or where to redirect after user logs out.
	LogoutUrl *string `pulumi:"logoutUrl"`
	// Enable the creation of local Artifactory users.  Default value is `false`.
	NoAutoUserCreation *bool `pulumi:"noAutoUserCreation"`
	// Name of the service provider configured on the .
	ServiceProviderName *string `pulumi:"serviceProviderName"`
	// Associate user with Artifactory groups based on the `groupAttribute` provided in the SAML response from the identity provider.  Default value is `false`.
	SyncGroups *bool `pulumi:"syncGroups"`
	// Enable "audience", or who the SAML assertion is intended for.  Ensures that the correct service provider intended for Artifactory is used on the IdP.  Default value is `true`.
	VerifyAudienceRestriction *bool `pulumi:"verifyAudienceRestriction"`
}

type SamlSettingsState struct {
	// Allow persisted users to access their profile.  Default value is `true`.
	AllowUserToAccessProfile pulumi.BoolPtrInput
	// Auto redirect to login through the IdP when clicking on Artifactory's login link.  Default value is `false`.
	AutoRedirect pulumi.BoolPtrInput
	// SAML certificate that contains the public key for the IdP service provider.  Used by Artifactory to verify sign-in requests.
	Certificate pulumi.StringPtrInput
	// Name of the attribute in the SAML response from the IdP that contains the user's email.
	EmailAttribute pulumi.StringPtrInput
	// Enable SAML SSO.  Default value is `true`.
	Enable pulumi.BoolPtrInput
	// Name of the attribute in the SAML response from the IdP that contains the user's group memberships.
	GroupAttribute pulumi.StringPtrInput
	// Service provider login url configured on the IdP.
	LoginUrl pulumi.StringPtrInput
	// Service provider logout url, or where to redirect after user logs out.
	LogoutUrl pulumi.StringPtrInput
	// Enable the creation of local Artifactory users.  Default value is `false`.
	NoAutoUserCreation pulumi.BoolPtrInput
	// Name of the service provider configured on the .
	ServiceProviderName pulumi.StringPtrInput
	// Associate user with Artifactory groups based on the `groupAttribute` provided in the SAML response from the identity provider.  Default value is `false`.
	SyncGroups pulumi.BoolPtrInput
	// Enable "audience", or who the SAML assertion is intended for.  Ensures that the correct service provider intended for Artifactory is used on the IdP.  Default value is `true`.
	VerifyAudienceRestriction pulumi.BoolPtrInput
}

func (SamlSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*samlSettingsState)(nil)).Elem()
}

type samlSettingsArgs struct {
	// Allow persisted users to access their profile.  Default value is `true`.
	AllowUserToAccessProfile *bool `pulumi:"allowUserToAccessProfile"`
	// Auto redirect to login through the IdP when clicking on Artifactory's login link.  Default value is `false`.
	AutoRedirect *bool `pulumi:"autoRedirect"`
	// SAML certificate that contains the public key for the IdP service provider.  Used by Artifactory to verify sign-in requests.
	Certificate *string `pulumi:"certificate"`
	// Name of the attribute in the SAML response from the IdP that contains the user's email.
	EmailAttribute *string `pulumi:"emailAttribute"`
	// Enable SAML SSO.  Default value is `true`.
	Enable *bool `pulumi:"enable"`
	// Name of the attribute in the SAML response from the IdP that contains the user's group memberships.
	GroupAttribute *string `pulumi:"groupAttribute"`
	// Service provider login url configured on the IdP.
	LoginUrl string `pulumi:"loginUrl"`
	// Service provider logout url, or where to redirect after user logs out.
	LogoutUrl string `pulumi:"logoutUrl"`
	// Enable the creation of local Artifactory users.  Default value is `false`.
	NoAutoUserCreation *bool `pulumi:"noAutoUserCreation"`
	// Name of the service provider configured on the .
	ServiceProviderName string `pulumi:"serviceProviderName"`
	// Associate user with Artifactory groups based on the `groupAttribute` provided in the SAML response from the identity provider.  Default value is `false`.
	SyncGroups *bool `pulumi:"syncGroups"`
	// Enable "audience", or who the SAML assertion is intended for.  Ensures that the correct service provider intended for Artifactory is used on the IdP.  Default value is `true`.
	VerifyAudienceRestriction *bool `pulumi:"verifyAudienceRestriction"`
}

// The set of arguments for constructing a SamlSettings resource.
type SamlSettingsArgs struct {
	// Allow persisted users to access their profile.  Default value is `true`.
	AllowUserToAccessProfile pulumi.BoolPtrInput
	// Auto redirect to login through the IdP when clicking on Artifactory's login link.  Default value is `false`.
	AutoRedirect pulumi.BoolPtrInput
	// SAML certificate that contains the public key for the IdP service provider.  Used by Artifactory to verify sign-in requests.
	Certificate pulumi.StringPtrInput
	// Name of the attribute in the SAML response from the IdP that contains the user's email.
	EmailAttribute pulumi.StringPtrInput
	// Enable SAML SSO.  Default value is `true`.
	Enable pulumi.BoolPtrInput
	// Name of the attribute in the SAML response from the IdP that contains the user's group memberships.
	GroupAttribute pulumi.StringPtrInput
	// Service provider login url configured on the IdP.
	LoginUrl pulumi.StringInput
	// Service provider logout url, or where to redirect after user logs out.
	LogoutUrl pulumi.StringInput
	// Enable the creation of local Artifactory users.  Default value is `false`.
	NoAutoUserCreation pulumi.BoolPtrInput
	// Name of the service provider configured on the .
	ServiceProviderName pulumi.StringInput
	// Associate user with Artifactory groups based on the `groupAttribute` provided in the SAML response from the identity provider.  Default value is `false`.
	SyncGroups pulumi.BoolPtrInput
	// Enable "audience", or who the SAML assertion is intended for.  Ensures that the correct service provider intended for Artifactory is used on the IdP.  Default value is `true`.
	VerifyAudienceRestriction pulumi.BoolPtrInput
}

func (SamlSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*samlSettingsArgs)(nil)).Elem()
}

type SamlSettingsInput interface {
	pulumi.Input

	ToSamlSettingsOutput() SamlSettingsOutput
	ToSamlSettingsOutputWithContext(ctx context.Context) SamlSettingsOutput
}

func (*SamlSettings) ElementType() reflect.Type {
	return reflect.TypeOf((*SamlSettings)(nil))
}

func (i *SamlSettings) ToSamlSettingsOutput() SamlSettingsOutput {
	return i.ToSamlSettingsOutputWithContext(context.Background())
}

func (i *SamlSettings) ToSamlSettingsOutputWithContext(ctx context.Context) SamlSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamlSettingsOutput)
}

func (i *SamlSettings) ToSamlSettingsPtrOutput() SamlSettingsPtrOutput {
	return i.ToSamlSettingsPtrOutputWithContext(context.Background())
}

func (i *SamlSettings) ToSamlSettingsPtrOutputWithContext(ctx context.Context) SamlSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamlSettingsPtrOutput)
}

type SamlSettingsPtrInput interface {
	pulumi.Input

	ToSamlSettingsPtrOutput() SamlSettingsPtrOutput
	ToSamlSettingsPtrOutputWithContext(ctx context.Context) SamlSettingsPtrOutput
}

type samlSettingsPtrType SamlSettingsArgs

func (*samlSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SamlSettings)(nil))
}

func (i *samlSettingsPtrType) ToSamlSettingsPtrOutput() SamlSettingsPtrOutput {
	return i.ToSamlSettingsPtrOutputWithContext(context.Background())
}

func (i *samlSettingsPtrType) ToSamlSettingsPtrOutputWithContext(ctx context.Context) SamlSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamlSettingsPtrOutput)
}

// SamlSettingsArrayInput is an input type that accepts SamlSettingsArray and SamlSettingsArrayOutput values.
// You can construct a concrete instance of `SamlSettingsArrayInput` via:
//
//          SamlSettingsArray{ SamlSettingsArgs{...} }
type SamlSettingsArrayInput interface {
	pulumi.Input

	ToSamlSettingsArrayOutput() SamlSettingsArrayOutput
	ToSamlSettingsArrayOutputWithContext(context.Context) SamlSettingsArrayOutput
}

type SamlSettingsArray []SamlSettingsInput

func (SamlSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SamlSettings)(nil)).Elem()
}

func (i SamlSettingsArray) ToSamlSettingsArrayOutput() SamlSettingsArrayOutput {
	return i.ToSamlSettingsArrayOutputWithContext(context.Background())
}

func (i SamlSettingsArray) ToSamlSettingsArrayOutputWithContext(ctx context.Context) SamlSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamlSettingsArrayOutput)
}

// SamlSettingsMapInput is an input type that accepts SamlSettingsMap and SamlSettingsMapOutput values.
// You can construct a concrete instance of `SamlSettingsMapInput` via:
//
//          SamlSettingsMap{ "key": SamlSettingsArgs{...} }
type SamlSettingsMapInput interface {
	pulumi.Input

	ToSamlSettingsMapOutput() SamlSettingsMapOutput
	ToSamlSettingsMapOutputWithContext(context.Context) SamlSettingsMapOutput
}

type SamlSettingsMap map[string]SamlSettingsInput

func (SamlSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SamlSettings)(nil)).Elem()
}

func (i SamlSettingsMap) ToSamlSettingsMapOutput() SamlSettingsMapOutput {
	return i.ToSamlSettingsMapOutputWithContext(context.Background())
}

func (i SamlSettingsMap) ToSamlSettingsMapOutputWithContext(ctx context.Context) SamlSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamlSettingsMapOutput)
}

type SamlSettingsOutput struct{ *pulumi.OutputState }

func (SamlSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SamlSettings)(nil))
}

func (o SamlSettingsOutput) ToSamlSettingsOutput() SamlSettingsOutput {
	return o
}

func (o SamlSettingsOutput) ToSamlSettingsOutputWithContext(ctx context.Context) SamlSettingsOutput {
	return o
}

func (o SamlSettingsOutput) ToSamlSettingsPtrOutput() SamlSettingsPtrOutput {
	return o.ToSamlSettingsPtrOutputWithContext(context.Background())
}

func (o SamlSettingsOutput) ToSamlSettingsPtrOutputWithContext(ctx context.Context) SamlSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SamlSettings) *SamlSettings {
		return &v
	}).(SamlSettingsPtrOutput)
}

type SamlSettingsPtrOutput struct{ *pulumi.OutputState }

func (SamlSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SamlSettings)(nil))
}

func (o SamlSettingsPtrOutput) ToSamlSettingsPtrOutput() SamlSettingsPtrOutput {
	return o
}

func (o SamlSettingsPtrOutput) ToSamlSettingsPtrOutputWithContext(ctx context.Context) SamlSettingsPtrOutput {
	return o
}

func (o SamlSettingsPtrOutput) Elem() SamlSettingsOutput {
	return o.ApplyT(func(v *SamlSettings) SamlSettings {
		if v != nil {
			return *v
		}
		var ret SamlSettings
		return ret
	}).(SamlSettingsOutput)
}

type SamlSettingsArrayOutput struct{ *pulumi.OutputState }

func (SamlSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SamlSettings)(nil))
}

func (o SamlSettingsArrayOutput) ToSamlSettingsArrayOutput() SamlSettingsArrayOutput {
	return o
}

func (o SamlSettingsArrayOutput) ToSamlSettingsArrayOutputWithContext(ctx context.Context) SamlSettingsArrayOutput {
	return o
}

func (o SamlSettingsArrayOutput) Index(i pulumi.IntInput) SamlSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SamlSettings {
		return vs[0].([]SamlSettings)[vs[1].(int)]
	}).(SamlSettingsOutput)
}

type SamlSettingsMapOutput struct{ *pulumi.OutputState }

func (SamlSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SamlSettings)(nil))
}

func (o SamlSettingsMapOutput) ToSamlSettingsMapOutput() SamlSettingsMapOutput {
	return o
}

func (o SamlSettingsMapOutput) ToSamlSettingsMapOutputWithContext(ctx context.Context) SamlSettingsMapOutput {
	return o
}

func (o SamlSettingsMapOutput) MapIndex(k pulumi.StringInput) SamlSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SamlSettings {
		return vs[0].(map[string]SamlSettings)[vs[1].(string)]
	}).(SamlSettingsOutput)
}

func init() {
	pulumi.RegisterOutputType(SamlSettingsOutput{})
	pulumi.RegisterOutputType(SamlSettingsPtrOutput{})
	pulumi.RegisterOutputType(SamlSettingsArrayOutput{})
	pulumi.RegisterOutputType(SamlSettingsMapOutput{})
}
