// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Artifactory webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.
//
// ## Example Usage
//
// .
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewArtifactLifecycleWebhook(ctx, "artifact-lifecycle-webhook", &artifactory.ArtifactLifecycleWebhookArgs{
//				Key: pulumi.String("artifact-lifecycle-webhook"),
//				EventTypes: pulumi.StringArray{
//					pulumi.String("archive"),
//					pulumi.String("restore"),
//				},
//				Handlers: artifactory.ArtifactLifecycleWebhookHandlerArray{
//					&artifactory.ArtifactLifecycleWebhookHandlerArgs{
//						Url:    pulumi.String("http://tempurl.org/webhook"),
//						Secret: pulumi.String("some-secret"),
//						Proxy:  pulumi.String("proxy-key"),
//						CustomHttpHeaders: pulumi.StringMap{
//							"header-1": pulumi.String("value-1"),
//							"header-2": pulumi.String("value-2"),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ArtifactLifecycleWebhook struct {
	pulumi.CustomResourceState

	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Status of webhook. Default to `true`
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// List of event triggers for the Webhook. Allow values: `archive`, `restore`
	EventTypes pulumi.StringArrayOutput `pulumi:"eventTypes"`
	// At least one is required.
	Handlers ArtifactLifecycleWebhookHandlerArrayOutput `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringOutput `pulumi:"key"`
}

// NewArtifactLifecycleWebhook registers a new resource with the given unique name, arguments, and options.
func NewArtifactLifecycleWebhook(ctx *pulumi.Context,
	name string, args *ArtifactLifecycleWebhookArgs, opts ...pulumi.ResourceOption) (*ArtifactLifecycleWebhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventTypes == nil {
		return nil, errors.New("invalid value for required argument 'EventTypes'")
	}
	if args.Handlers == nil {
		return nil, errors.New("invalid value for required argument 'Handlers'")
	}
	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ArtifactLifecycleWebhook
	err := ctx.RegisterResource("artifactory:index/artifactLifecycleWebhook:ArtifactLifecycleWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetArtifactLifecycleWebhook gets an existing ArtifactLifecycleWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetArtifactLifecycleWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ArtifactLifecycleWebhookState, opts ...pulumi.ResourceOption) (*ArtifactLifecycleWebhook, error) {
	var resource ArtifactLifecycleWebhook
	err := ctx.ReadResource("artifactory:index/artifactLifecycleWebhook:ArtifactLifecycleWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ArtifactLifecycleWebhook resources.
type artifactLifecycleWebhookState struct {
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`
	Enabled *bool `pulumi:"enabled"`
	// List of event triggers for the Webhook. Allow values: `archive`, `restore`
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []ArtifactLifecycleWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key *string `pulumi:"key"`
}

type ArtifactLifecycleWebhookState struct {
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`
	Enabled pulumi.BoolPtrInput
	// List of event triggers for the Webhook. Allow values: `archive`, `restore`
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers ArtifactLifecycleWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringPtrInput
}

func (ArtifactLifecycleWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactLifecycleWebhookState)(nil)).Elem()
}

type artifactLifecycleWebhookArgs struct {
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`
	Enabled *bool `pulumi:"enabled"`
	// List of event triggers for the Webhook. Allow values: `archive`, `restore`
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []ArtifactLifecycleWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key string `pulumi:"key"`
}

// The set of arguments for constructing a ArtifactLifecycleWebhook resource.
type ArtifactLifecycleWebhookArgs struct {
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`
	Enabled pulumi.BoolPtrInput
	// List of event triggers for the Webhook. Allow values: `archive`, `restore`
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers ArtifactLifecycleWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringInput
}

func (ArtifactLifecycleWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactLifecycleWebhookArgs)(nil)).Elem()
}

type ArtifactLifecycleWebhookInput interface {
	pulumi.Input

	ToArtifactLifecycleWebhookOutput() ArtifactLifecycleWebhookOutput
	ToArtifactLifecycleWebhookOutputWithContext(ctx context.Context) ArtifactLifecycleWebhookOutput
}

func (*ArtifactLifecycleWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**ArtifactLifecycleWebhook)(nil)).Elem()
}

func (i *ArtifactLifecycleWebhook) ToArtifactLifecycleWebhookOutput() ArtifactLifecycleWebhookOutput {
	return i.ToArtifactLifecycleWebhookOutputWithContext(context.Background())
}

func (i *ArtifactLifecycleWebhook) ToArtifactLifecycleWebhookOutputWithContext(ctx context.Context) ArtifactLifecycleWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactLifecycleWebhookOutput)
}

// ArtifactLifecycleWebhookArrayInput is an input type that accepts ArtifactLifecycleWebhookArray and ArtifactLifecycleWebhookArrayOutput values.
// You can construct a concrete instance of `ArtifactLifecycleWebhookArrayInput` via:
//
//	ArtifactLifecycleWebhookArray{ ArtifactLifecycleWebhookArgs{...} }
type ArtifactLifecycleWebhookArrayInput interface {
	pulumi.Input

	ToArtifactLifecycleWebhookArrayOutput() ArtifactLifecycleWebhookArrayOutput
	ToArtifactLifecycleWebhookArrayOutputWithContext(context.Context) ArtifactLifecycleWebhookArrayOutput
}

type ArtifactLifecycleWebhookArray []ArtifactLifecycleWebhookInput

func (ArtifactLifecycleWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ArtifactLifecycleWebhook)(nil)).Elem()
}

func (i ArtifactLifecycleWebhookArray) ToArtifactLifecycleWebhookArrayOutput() ArtifactLifecycleWebhookArrayOutput {
	return i.ToArtifactLifecycleWebhookArrayOutputWithContext(context.Background())
}

func (i ArtifactLifecycleWebhookArray) ToArtifactLifecycleWebhookArrayOutputWithContext(ctx context.Context) ArtifactLifecycleWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactLifecycleWebhookArrayOutput)
}

// ArtifactLifecycleWebhookMapInput is an input type that accepts ArtifactLifecycleWebhookMap and ArtifactLifecycleWebhookMapOutput values.
// You can construct a concrete instance of `ArtifactLifecycleWebhookMapInput` via:
//
//	ArtifactLifecycleWebhookMap{ "key": ArtifactLifecycleWebhookArgs{...} }
type ArtifactLifecycleWebhookMapInput interface {
	pulumi.Input

	ToArtifactLifecycleWebhookMapOutput() ArtifactLifecycleWebhookMapOutput
	ToArtifactLifecycleWebhookMapOutputWithContext(context.Context) ArtifactLifecycleWebhookMapOutput
}

type ArtifactLifecycleWebhookMap map[string]ArtifactLifecycleWebhookInput

func (ArtifactLifecycleWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ArtifactLifecycleWebhook)(nil)).Elem()
}

func (i ArtifactLifecycleWebhookMap) ToArtifactLifecycleWebhookMapOutput() ArtifactLifecycleWebhookMapOutput {
	return i.ToArtifactLifecycleWebhookMapOutputWithContext(context.Background())
}

func (i ArtifactLifecycleWebhookMap) ToArtifactLifecycleWebhookMapOutputWithContext(ctx context.Context) ArtifactLifecycleWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactLifecycleWebhookMapOutput)
}

type ArtifactLifecycleWebhookOutput struct{ *pulumi.OutputState }

func (ArtifactLifecycleWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArtifactLifecycleWebhook)(nil)).Elem()
}

func (o ArtifactLifecycleWebhookOutput) ToArtifactLifecycleWebhookOutput() ArtifactLifecycleWebhookOutput {
	return o
}

func (o ArtifactLifecycleWebhookOutput) ToArtifactLifecycleWebhookOutputWithContext(ctx context.Context) ArtifactLifecycleWebhookOutput {
	return o
}

// Webhook description. Max length 1000 characters.
func (o ArtifactLifecycleWebhookOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactLifecycleWebhook) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Status of webhook. Default to `true`
func (o ArtifactLifecycleWebhookOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ArtifactLifecycleWebhook) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// List of event triggers for the Webhook. Allow values: `archive`, `restore`
func (o ArtifactLifecycleWebhookOutput) EventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ArtifactLifecycleWebhook) pulumi.StringArrayOutput { return v.EventTypes }).(pulumi.StringArrayOutput)
}

// At least one is required.
func (o ArtifactLifecycleWebhookOutput) Handlers() ArtifactLifecycleWebhookHandlerArrayOutput {
	return o.ApplyT(func(v *ArtifactLifecycleWebhook) ArtifactLifecycleWebhookHandlerArrayOutput { return v.Handlers }).(ArtifactLifecycleWebhookHandlerArrayOutput)
}

// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
func (o ArtifactLifecycleWebhookOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *ArtifactLifecycleWebhook) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

type ArtifactLifecycleWebhookArrayOutput struct{ *pulumi.OutputState }

func (ArtifactLifecycleWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ArtifactLifecycleWebhook)(nil)).Elem()
}

func (o ArtifactLifecycleWebhookArrayOutput) ToArtifactLifecycleWebhookArrayOutput() ArtifactLifecycleWebhookArrayOutput {
	return o
}

func (o ArtifactLifecycleWebhookArrayOutput) ToArtifactLifecycleWebhookArrayOutputWithContext(ctx context.Context) ArtifactLifecycleWebhookArrayOutput {
	return o
}

func (o ArtifactLifecycleWebhookArrayOutput) Index(i pulumi.IntInput) ArtifactLifecycleWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ArtifactLifecycleWebhook {
		return vs[0].([]*ArtifactLifecycleWebhook)[vs[1].(int)]
	}).(ArtifactLifecycleWebhookOutput)
}

type ArtifactLifecycleWebhookMapOutput struct{ *pulumi.OutputState }

func (ArtifactLifecycleWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ArtifactLifecycleWebhook)(nil)).Elem()
}

func (o ArtifactLifecycleWebhookMapOutput) ToArtifactLifecycleWebhookMapOutput() ArtifactLifecycleWebhookMapOutput {
	return o
}

func (o ArtifactLifecycleWebhookMapOutput) ToArtifactLifecycleWebhookMapOutputWithContext(ctx context.Context) ArtifactLifecycleWebhookMapOutput {
	return o
}

func (o ArtifactLifecycleWebhookMapOutput) MapIndex(k pulumi.StringInput) ArtifactLifecycleWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ArtifactLifecycleWebhook {
		return vs[0].(map[string]*ArtifactLifecycleWebhook)[vs[1].(string)]
	}).(ArtifactLifecycleWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ArtifactLifecycleWebhookInput)(nil)).Elem(), &ArtifactLifecycleWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*ArtifactLifecycleWebhookArrayInput)(nil)).Elem(), ArtifactLifecycleWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ArtifactLifecycleWebhookMapInput)(nil)).Elem(), ArtifactLifecycleWebhookMap{})
	pulumi.RegisterOutputType(ArtifactLifecycleWebhookOutput{})
	pulumi.RegisterOutputType(ArtifactLifecycleWebhookArrayOutput{})
	pulumi.RegisterOutputType(ArtifactLifecycleWebhookMapOutput{})
}
