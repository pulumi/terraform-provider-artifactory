// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the artifactory package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// This is a bearer token that can be given to you by your admin under `Identity and Access`
	AccessToken pulumi.StringPtrOutput `pulumi:"accessToken"`
	// Deprecated: Xray and projects functionality will not work with any auth method other than access tokens (Bearer)
	ApiKey pulumi.StringPtrOutput `pulumi:"apiKey"`
	// Insider note: You may actually use an api_key as the password. This will get your around xray limitations instead of a
	// bearer token
	//
	// Deprecated: Xray and projects functionality will not work with any auth method other than access tokens (Bearer)
	Password pulumi.StringPtrOutput `pulumi:"password"`
	Url      pulumi.StringPtrOutput `pulumi:"url"`
	// Deprecated: Xray and projects functionality will not work with any auth method other than access tokens (Bearer)
	Username pulumi.StringPtrOutput `pulumi:"username"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:artifactory", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// This is a bearer token that can be given to you by your admin under `Identity and Access`
	AccessToken *string `pulumi:"accessToken"`
	// Deprecated: Xray and projects functionality will not work with any auth method other than access tokens (Bearer)
	ApiKey *string `pulumi:"apiKey"`
	// Insider note: You may actually use an api_key as the password. This will get your around xray limitations instead of a
	// bearer token
	//
	// Deprecated: Xray and projects functionality will not work with any auth method other than access tokens (Bearer)
	Password *string `pulumi:"password"`
	Url      *string `pulumi:"url"`
	// Deprecated: Xray and projects functionality will not work with any auth method other than access tokens (Bearer)
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// This is a bearer token that can be given to you by your admin under `Identity and Access`
	AccessToken pulumi.StringPtrInput
	// Deprecated: Xray and projects functionality will not work with any auth method other than access tokens (Bearer)
	ApiKey pulumi.StringPtrInput
	// Insider note: You may actually use an api_key as the password. This will get your around xray limitations instead of a
	// bearer token
	//
	// Deprecated: Xray and projects functionality will not work with any auth method other than access tokens (Bearer)
	Password pulumi.StringPtrInput
	Url      pulumi.StringPtrInput
	// Deprecated: Xray and projects functionality will not work with any auth method other than access tokens (Bearer)
	Username pulumi.StringPtrInput
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
	return reflect.TypeOf((*Provider)(nil))
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

func (i *Provider) ToProviderPtrOutput() ProviderPtrOutput {
	return i.ToProviderPtrOutputWithContext(context.Background())
}

func (i *Provider) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderPtrOutput)
}

type ProviderPtrInput interface {
	pulumi.Input

	ToProviderPtrOutput() ProviderPtrOutput
	ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput
}

type providerPtrType ProviderArgs

func (*providerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil))
}

func (i *providerPtrType) ToProviderPtrOutput() ProviderPtrOutput {
	return i.ToProviderPtrOutputWithContext(context.Background())
}

func (i *providerPtrType) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderPtrOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Provider)(nil))
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderPtrOutput() ProviderPtrOutput {
	return o.ToProviderPtrOutputWithContext(context.Background())
}

func (o ProviderOutput) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Provider) *Provider {
		return &v
	}).(ProviderPtrOutput)
}

type ProviderPtrOutput struct{ *pulumi.OutputState }

func (ProviderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil))
}

func (o ProviderPtrOutput) ToProviderPtrOutput() ProviderPtrOutput {
	return o
}

func (o ProviderPtrOutput) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return o
}

func (o ProviderPtrOutput) Elem() ProviderOutput {
	return o.ApplyT(func(v *Provider) Provider {
		if v != nil {
			return *v
		}
		var ret Provider
		return ret
	}).(ProviderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderPtrInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
	pulumi.RegisterOutputType(ProviderPtrOutput{})
}
