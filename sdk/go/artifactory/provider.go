// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v6/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the artifactory package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// This is a access token that can be given to you by your admin under `User Management > Access Tokens`. If not set, the
	// 'api_key' attribute value will be used.
	AccessToken pulumi.StringPtrOutput `pulumi:"accessToken"`
	// API key. If `accessToken` attribute, `JFROG_ACCESS_TOKEN` or `ARTIFACTORY_ACCESS_TOKEN` environment variable is set, the
	// provider will ignore this attribute.
	//
	// Deprecated: An upcoming version will support the option to block the usage/creation of API Keys (for admins to set on their platform).
	// In a future version (scheduled for end of Q3, 2023), the option to disable the usage/creation of API Keys will be available and set to disabled by default. Admins will be able to enable the usage/creation of API Keys.
	// By end of Q4 2024, API Keys will be deprecated all together and the option to use them will no longer be available. See [JFrog API deprecation process](https://jfrog.com/help/r/jfrog-platform-administration-documentation/jfrog-api-key-deprecation-process) for more details.
	ApiKey pulumi.StringPtrOutput `pulumi:"apiKey"`
	// OIDC provider name. See [Configure an OIDC
	// Integration](https://jfrog.com/help/r/jfrog-platform-administration-documentation/configure-an-oidc-integration) for
	// more details.
	OidcProviderName pulumi.StringPtrOutput `pulumi:"oidcProviderName"`
	// Artifactory URL.
	Url pulumi.StringPtrOutput `pulumi:"url"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	if args.CheckLicense == nil {
		args.CheckLicense = pulumi.BoolPtr(false)
	}
	if args.AccessToken != nil {
		args.AccessToken = pulumi.ToSecret(args.AccessToken).(pulumi.StringPtrInput)
	}
	if args.ApiKey != nil {
		args.ApiKey = pulumi.ToSecret(args.ApiKey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"accessToken",
		"apiKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:artifactory", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// This is a access token that can be given to you by your admin under `User Management > Access Tokens`. If not set, the
	// 'api_key' attribute value will be used.
	AccessToken *string `pulumi:"accessToken"`
	// API key. If `accessToken` attribute, `JFROG_ACCESS_TOKEN` or `ARTIFACTORY_ACCESS_TOKEN` environment variable is set, the
	// provider will ignore this attribute.
	//
	// Deprecated: An upcoming version will support the option to block the usage/creation of API Keys (for admins to set on their platform).
	// In a future version (scheduled for end of Q3, 2023), the option to disable the usage/creation of API Keys will be available and set to disabled by default. Admins will be able to enable the usage/creation of API Keys.
	// By end of Q4 2024, API Keys will be deprecated all together and the option to use them will no longer be available. See [JFrog API deprecation process](https://jfrog.com/help/r/jfrog-platform-administration-documentation/jfrog-api-key-deprecation-process) for more details.
	ApiKey *string `pulumi:"apiKey"`
	// Toggle for pre-flight checking of Artifactory Pro and Enterprise license. Default to `true`.
	CheckLicense *bool `pulumi:"checkLicense"`
	// OIDC provider name. See [Configure an OIDC
	// Integration](https://jfrog.com/help/r/jfrog-platform-administration-documentation/configure-an-oidc-integration) for
	// more details.
	OidcProviderName *string `pulumi:"oidcProviderName"`
	// Artifactory URL.
	Url *string `pulumi:"url"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// This is a access token that can be given to you by your admin under `User Management > Access Tokens`. If not set, the
	// 'api_key' attribute value will be used.
	AccessToken pulumi.StringPtrInput
	// API key. If `accessToken` attribute, `JFROG_ACCESS_TOKEN` or `ARTIFACTORY_ACCESS_TOKEN` environment variable is set, the
	// provider will ignore this attribute.
	//
	// Deprecated: An upcoming version will support the option to block the usage/creation of API Keys (for admins to set on their platform).
	// In a future version (scheduled for end of Q3, 2023), the option to disable the usage/creation of API Keys will be available and set to disabled by default. Admins will be able to enable the usage/creation of API Keys.
	// By end of Q4 2024, API Keys will be deprecated all together and the option to use them will no longer be available. See [JFrog API deprecation process](https://jfrog.com/help/r/jfrog-platform-administration-documentation/jfrog-api-key-deprecation-process) for more details.
	ApiKey pulumi.StringPtrInput
	// Toggle for pre-flight checking of Artifactory Pro and Enterprise license. Default to `true`.
	CheckLicense pulumi.BoolPtrInput
	// OIDC provider name. See [Configure an OIDC
	// Integration](https://jfrog.com/help/r/jfrog-platform-administration-documentation/configure-an-oidc-integration) for
	// more details.
	OidcProviderName pulumi.StringPtrInput
	// Artifactory URL.
	Url pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

// This is a access token that can be given to you by your admin under `User Management > Access Tokens`. If not set, the
// 'api_key' attribute value will be used.
func (o ProviderOutput) AccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.AccessToken }).(pulumi.StringPtrOutput)
}

// API key. If `accessToken` attribute, `JFROG_ACCESS_TOKEN` or `ARTIFACTORY_ACCESS_TOKEN` environment variable is set, the
// provider will ignore this attribute.
//
// Deprecated: An upcoming version will support the option to block the usage/creation of API Keys (for admins to set on their platform).
// In a future version (scheduled for end of Q3, 2023), the option to disable the usage/creation of API Keys will be available and set to disabled by default. Admins will be able to enable the usage/creation of API Keys.
// By end of Q4 2024, API Keys will be deprecated all together and the option to use them will no longer be available. See [JFrog API deprecation process](https://jfrog.com/help/r/jfrog-platform-administration-documentation/jfrog-api-key-deprecation-process) for more details.
func (o ProviderOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ApiKey }).(pulumi.StringPtrOutput)
}

// OIDC provider name. See [Configure an OIDC
// Integration](https://jfrog.com/help/r/jfrog-platform-administration-documentation/configure-an-oidc-integration) for
// more details.
func (o ProviderOutput) OidcProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.OidcProviderName }).(pulumi.StringPtrOutput)
}

// Artifactory URL.
func (o ProviderOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Url }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
