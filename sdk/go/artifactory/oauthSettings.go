// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource can be used to manage Artifactory's OAuth SSO settings.
//
// Only a single `OauthSettings` resource is meant to be defined.
//
// ~>The `OauthSettings` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v3/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewOauthSettings(ctx, "oauth", &artifactory.OauthSettingsArgs{
//				AllowUserToAccessProfile: pulumi.Bool(true),
//				Enable:                   pulumi.Bool(true),
//				OauthProviders: artifactory.OauthSettingsOauthProviderArray{
//					&artifactory.OauthSettingsOauthProviderArgs{
//						ApiUrl:       pulumi.String("https://organization.okta.com/oauth2/v1/userinfo"),
//						AuthUrl:      pulumi.String("https://organization.okta.com/oauth2/v1/authorize"),
//						ClientId:     pulumi.String("foo"),
//						ClientSecret: pulumi.String("bar"),
//						Enabled:      pulumi.Bool(false),
//						Name:         pulumi.String("okta"),
//						TokenUrl:     pulumi.String("https://organization.okta.com/oauth2/v1/token"),
//						Type:         pulumi.String("openId"),
//					},
//				},
//				PersistUsers: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Current OAuth SSO settings can be imported using `oauth_settings` as the `ID`. If the resource is being imported, there will be a state drift, because `client_secret` can't be known. There are two options on how to approach this:
//
// 1) Don't set `client_secret` initially, import, then update the config with actual secret; 2) Accept that there is a drift initially and run `pulumi up` twice;
//
// ```sh
//
//	$ pulumi import artifactory:index/oauthSettings:OauthSettings oauth oauth_settings
//
// ```
type OauthSettings struct {
	pulumi.CustomResourceState

	// Allow persisted users to access their profile.  Default value is `false`.
	AllowUserToAccessProfile pulumi.BoolPtrOutput `pulumi:"allowUserToAccessProfile"`
	// Enable OAuth SSO.  Default value is `true`.
	Enable pulumi.BoolPtrOutput `pulumi:"enable"`
	// OAuth provider settings block. Multiple blocks can be defined, at least one is required.
	OauthProviders OauthSettingsOauthProviderArrayOutput `pulumi:"oauthProviders"`
	// Enable the creation of local Artifactory users.  Default value is `false`.
	PersistUsers pulumi.BoolPtrOutput `pulumi:"persistUsers"`
}

// NewOauthSettings registers a new resource with the given unique name, arguments, and options.
func NewOauthSettings(ctx *pulumi.Context,
	name string, args *OauthSettingsArgs, opts ...pulumi.ResourceOption) (*OauthSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OauthProviders == nil {
		return nil, errors.New("invalid value for required argument 'OauthProviders'")
	}
	var resource OauthSettings
	err := ctx.RegisterResource("artifactory:index/oauthSettings:OauthSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOauthSettings gets an existing OauthSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOauthSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OauthSettingsState, opts ...pulumi.ResourceOption) (*OauthSettings, error) {
	var resource OauthSettings
	err := ctx.ReadResource("artifactory:index/oauthSettings:OauthSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OauthSettings resources.
type oauthSettingsState struct {
	// Allow persisted users to access their profile.  Default value is `false`.
	AllowUserToAccessProfile *bool `pulumi:"allowUserToAccessProfile"`
	// Enable OAuth SSO.  Default value is `true`.
	Enable *bool `pulumi:"enable"`
	// OAuth provider settings block. Multiple blocks can be defined, at least one is required.
	OauthProviders []OauthSettingsOauthProvider `pulumi:"oauthProviders"`
	// Enable the creation of local Artifactory users.  Default value is `false`.
	PersistUsers *bool `pulumi:"persistUsers"`
}

type OauthSettingsState struct {
	// Allow persisted users to access their profile.  Default value is `false`.
	AllowUserToAccessProfile pulumi.BoolPtrInput
	// Enable OAuth SSO.  Default value is `true`.
	Enable pulumi.BoolPtrInput
	// OAuth provider settings block. Multiple blocks can be defined, at least one is required.
	OauthProviders OauthSettingsOauthProviderArrayInput
	// Enable the creation of local Artifactory users.  Default value is `false`.
	PersistUsers pulumi.BoolPtrInput
}

func (OauthSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*oauthSettingsState)(nil)).Elem()
}

type oauthSettingsArgs struct {
	// Allow persisted users to access their profile.  Default value is `false`.
	AllowUserToAccessProfile *bool `pulumi:"allowUserToAccessProfile"`
	// Enable OAuth SSO.  Default value is `true`.
	Enable *bool `pulumi:"enable"`
	// OAuth provider settings block. Multiple blocks can be defined, at least one is required.
	OauthProviders []OauthSettingsOauthProvider `pulumi:"oauthProviders"`
	// Enable the creation of local Artifactory users.  Default value is `false`.
	PersistUsers *bool `pulumi:"persistUsers"`
}

// The set of arguments for constructing a OauthSettings resource.
type OauthSettingsArgs struct {
	// Allow persisted users to access their profile.  Default value is `false`.
	AllowUserToAccessProfile pulumi.BoolPtrInput
	// Enable OAuth SSO.  Default value is `true`.
	Enable pulumi.BoolPtrInput
	// OAuth provider settings block. Multiple blocks can be defined, at least one is required.
	OauthProviders OauthSettingsOauthProviderArrayInput
	// Enable the creation of local Artifactory users.  Default value is `false`.
	PersistUsers pulumi.BoolPtrInput
}

func (OauthSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*oauthSettingsArgs)(nil)).Elem()
}

type OauthSettingsInput interface {
	pulumi.Input

	ToOauthSettingsOutput() OauthSettingsOutput
	ToOauthSettingsOutputWithContext(ctx context.Context) OauthSettingsOutput
}

func (*OauthSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**OauthSettings)(nil)).Elem()
}

func (i *OauthSettings) ToOauthSettingsOutput() OauthSettingsOutput {
	return i.ToOauthSettingsOutputWithContext(context.Background())
}

func (i *OauthSettings) ToOauthSettingsOutputWithContext(ctx context.Context) OauthSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OauthSettingsOutput)
}

// OauthSettingsArrayInput is an input type that accepts OauthSettingsArray and OauthSettingsArrayOutput values.
// You can construct a concrete instance of `OauthSettingsArrayInput` via:
//
//	OauthSettingsArray{ OauthSettingsArgs{...} }
type OauthSettingsArrayInput interface {
	pulumi.Input

	ToOauthSettingsArrayOutput() OauthSettingsArrayOutput
	ToOauthSettingsArrayOutputWithContext(context.Context) OauthSettingsArrayOutput
}

type OauthSettingsArray []OauthSettingsInput

func (OauthSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OauthSettings)(nil)).Elem()
}

func (i OauthSettingsArray) ToOauthSettingsArrayOutput() OauthSettingsArrayOutput {
	return i.ToOauthSettingsArrayOutputWithContext(context.Background())
}

func (i OauthSettingsArray) ToOauthSettingsArrayOutputWithContext(ctx context.Context) OauthSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OauthSettingsArrayOutput)
}

// OauthSettingsMapInput is an input type that accepts OauthSettingsMap and OauthSettingsMapOutput values.
// You can construct a concrete instance of `OauthSettingsMapInput` via:
//
//	OauthSettingsMap{ "key": OauthSettingsArgs{...} }
type OauthSettingsMapInput interface {
	pulumi.Input

	ToOauthSettingsMapOutput() OauthSettingsMapOutput
	ToOauthSettingsMapOutputWithContext(context.Context) OauthSettingsMapOutput
}

type OauthSettingsMap map[string]OauthSettingsInput

func (OauthSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OauthSettings)(nil)).Elem()
}

func (i OauthSettingsMap) ToOauthSettingsMapOutput() OauthSettingsMapOutput {
	return i.ToOauthSettingsMapOutputWithContext(context.Background())
}

func (i OauthSettingsMap) ToOauthSettingsMapOutputWithContext(ctx context.Context) OauthSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OauthSettingsMapOutput)
}

type OauthSettingsOutput struct{ *pulumi.OutputState }

func (OauthSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OauthSettings)(nil)).Elem()
}

func (o OauthSettingsOutput) ToOauthSettingsOutput() OauthSettingsOutput {
	return o
}

func (o OauthSettingsOutput) ToOauthSettingsOutputWithContext(ctx context.Context) OauthSettingsOutput {
	return o
}

// Allow persisted users to access their profile.  Default value is `false`.
func (o OauthSettingsOutput) AllowUserToAccessProfile() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OauthSettings) pulumi.BoolPtrOutput { return v.AllowUserToAccessProfile }).(pulumi.BoolPtrOutput)
}

// Enable OAuth SSO.  Default value is `true`.
func (o OauthSettingsOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OauthSettings) pulumi.BoolPtrOutput { return v.Enable }).(pulumi.BoolPtrOutput)
}

// OAuth provider settings block. Multiple blocks can be defined, at least one is required.
func (o OauthSettingsOutput) OauthProviders() OauthSettingsOauthProviderArrayOutput {
	return o.ApplyT(func(v *OauthSettings) OauthSettingsOauthProviderArrayOutput { return v.OauthProviders }).(OauthSettingsOauthProviderArrayOutput)
}

// Enable the creation of local Artifactory users.  Default value is `false`.
func (o OauthSettingsOutput) PersistUsers() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OauthSettings) pulumi.BoolPtrOutput { return v.PersistUsers }).(pulumi.BoolPtrOutput)
}

type OauthSettingsArrayOutput struct{ *pulumi.OutputState }

func (OauthSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OauthSettings)(nil)).Elem()
}

func (o OauthSettingsArrayOutput) ToOauthSettingsArrayOutput() OauthSettingsArrayOutput {
	return o
}

func (o OauthSettingsArrayOutput) ToOauthSettingsArrayOutputWithContext(ctx context.Context) OauthSettingsArrayOutput {
	return o
}

func (o OauthSettingsArrayOutput) Index(i pulumi.IntInput) OauthSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OauthSettings {
		return vs[0].([]*OauthSettings)[vs[1].(int)]
	}).(OauthSettingsOutput)
}

type OauthSettingsMapOutput struct{ *pulumi.OutputState }

func (OauthSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OauthSettings)(nil)).Elem()
}

func (o OauthSettingsMapOutput) ToOauthSettingsMapOutput() OauthSettingsMapOutput {
	return o
}

func (o OauthSettingsMapOutput) ToOauthSettingsMapOutputWithContext(ctx context.Context) OauthSettingsMapOutput {
	return o
}

func (o OauthSettingsMapOutput) MapIndex(k pulumi.StringInput) OauthSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OauthSettings {
		return vs[0].(map[string]*OauthSettings)[vs[1].(string)]
	}).(OauthSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OauthSettingsInput)(nil)).Elem(), &OauthSettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*OauthSettingsArrayInput)(nil)).Elem(), OauthSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OauthSettingsMapInput)(nil)).Elem(), OauthSettingsMap{})
	pulumi.RegisterOutputType(OauthSettingsOutput{})
	pulumi.RegisterOutputType(OauthSettingsArrayOutput{})
	pulumi.RegisterOutputType(OauthSettingsMapOutput{})
}
