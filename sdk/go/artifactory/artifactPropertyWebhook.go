// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory/internal"
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
//	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewLocalGenericRepository(ctx, "my-generic-local", &artifactory.LocalGenericRepositoryArgs{
//				Key: pulumi.String("my-generic-local"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = artifactory.NewArtifactPropertyWebhook(ctx, "artifact-webhook", &artifactory.ArtifactPropertyWebhookArgs{
//				Key: pulumi.String("artifact-property-webhook"),
//				EventTypes: pulumi.StringArray{
//					pulumi.String("added"),
//					pulumi.String("deleted"),
//				},
//				Criteria: &artifactory.ArtifactPropertyWebhookCriteriaArgs{
//					AnyLocal:  pulumi.Bool(true),
//					AnyRemote: pulumi.Bool(false),
//					RepoKeys: pulumi.StringArray{
//						my_generic_local.Key,
//					},
//					IncludePatterns: pulumi.StringArray{
//						pulumi.String("foo/**"),
//					},
//					ExcludePatterns: pulumi.StringArray{
//						pulumi.String("bar/**"),
//					},
//				},
//				Handlers: artifactory.ArtifactPropertyWebhookHandlerArray{
//					&artifactory.ArtifactPropertyWebhookHandlerArgs{
//						Url:    pulumi.String("http://tempurl.org/webhook"),
//						Secret: pulumi.String("some-secret"),
//						Proxy:  pulumi.String("proxy-key"),
//						CustomHttpHeaders: pulumi.StringMap{
//							"header-1": pulumi.String("value-1"),
//							"header-2": pulumi.String("value-2"),
//						},
//					},
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				my_generic_local,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ArtifactPropertyWebhook struct {
	pulumi.CustomResourceState

	// Specifies where the webhook will be applied on which repositories.
	Criteria ArtifactPropertyWebhookCriteriaOutput `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `added`, `deleted`.
	EventTypes pulumi.StringArrayOutput `pulumi:"eventTypes"`
	// At least one is required.
	Handlers ArtifactPropertyWebhookHandlerArrayOutput `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringOutput `pulumi:"key"`
}

// NewArtifactPropertyWebhook registers a new resource with the given unique name, arguments, and options.
func NewArtifactPropertyWebhook(ctx *pulumi.Context,
	name string, args *ArtifactPropertyWebhookArgs, opts ...pulumi.ResourceOption) (*ArtifactPropertyWebhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Criteria == nil {
		return nil, errors.New("invalid value for required argument 'Criteria'")
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
	var resource ArtifactPropertyWebhook
	err := ctx.RegisterResource("artifactory:index/artifactPropertyWebhook:ArtifactPropertyWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetArtifactPropertyWebhook gets an existing ArtifactPropertyWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetArtifactPropertyWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ArtifactPropertyWebhookState, opts ...pulumi.ResourceOption) (*ArtifactPropertyWebhook, error) {
	var resource ArtifactPropertyWebhook
	err := ctx.ReadResource("artifactory:index/artifactPropertyWebhook:ArtifactPropertyWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ArtifactPropertyWebhook resources.
type artifactPropertyWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria *ArtifactPropertyWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `added`, `deleted`.
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []ArtifactPropertyWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key *string `pulumi:"key"`
}

type ArtifactPropertyWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria ArtifactPropertyWebhookCriteriaPtrInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `added`, `deleted`.
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers ArtifactPropertyWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringPtrInput
}

func (ArtifactPropertyWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactPropertyWebhookState)(nil)).Elem()
}

type artifactPropertyWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria ArtifactPropertyWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `added`, `deleted`.
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []ArtifactPropertyWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key string `pulumi:"key"`
}

// The set of arguments for constructing a ArtifactPropertyWebhook resource.
type ArtifactPropertyWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria ArtifactPropertyWebhookCriteriaInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `added`, `deleted`.
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers ArtifactPropertyWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringInput
}

func (ArtifactPropertyWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactPropertyWebhookArgs)(nil)).Elem()
}

type ArtifactPropertyWebhookInput interface {
	pulumi.Input

	ToArtifactPropertyWebhookOutput() ArtifactPropertyWebhookOutput
	ToArtifactPropertyWebhookOutputWithContext(ctx context.Context) ArtifactPropertyWebhookOutput
}

func (*ArtifactPropertyWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**ArtifactPropertyWebhook)(nil)).Elem()
}

func (i *ArtifactPropertyWebhook) ToArtifactPropertyWebhookOutput() ArtifactPropertyWebhookOutput {
	return i.ToArtifactPropertyWebhookOutputWithContext(context.Background())
}

func (i *ArtifactPropertyWebhook) ToArtifactPropertyWebhookOutputWithContext(ctx context.Context) ArtifactPropertyWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactPropertyWebhookOutput)
}

// ArtifactPropertyWebhookArrayInput is an input type that accepts ArtifactPropertyWebhookArray and ArtifactPropertyWebhookArrayOutput values.
// You can construct a concrete instance of `ArtifactPropertyWebhookArrayInput` via:
//
//	ArtifactPropertyWebhookArray{ ArtifactPropertyWebhookArgs{...} }
type ArtifactPropertyWebhookArrayInput interface {
	pulumi.Input

	ToArtifactPropertyWebhookArrayOutput() ArtifactPropertyWebhookArrayOutput
	ToArtifactPropertyWebhookArrayOutputWithContext(context.Context) ArtifactPropertyWebhookArrayOutput
}

type ArtifactPropertyWebhookArray []ArtifactPropertyWebhookInput

func (ArtifactPropertyWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ArtifactPropertyWebhook)(nil)).Elem()
}

func (i ArtifactPropertyWebhookArray) ToArtifactPropertyWebhookArrayOutput() ArtifactPropertyWebhookArrayOutput {
	return i.ToArtifactPropertyWebhookArrayOutputWithContext(context.Background())
}

func (i ArtifactPropertyWebhookArray) ToArtifactPropertyWebhookArrayOutputWithContext(ctx context.Context) ArtifactPropertyWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactPropertyWebhookArrayOutput)
}

// ArtifactPropertyWebhookMapInput is an input type that accepts ArtifactPropertyWebhookMap and ArtifactPropertyWebhookMapOutput values.
// You can construct a concrete instance of `ArtifactPropertyWebhookMapInput` via:
//
//	ArtifactPropertyWebhookMap{ "key": ArtifactPropertyWebhookArgs{...} }
type ArtifactPropertyWebhookMapInput interface {
	pulumi.Input

	ToArtifactPropertyWebhookMapOutput() ArtifactPropertyWebhookMapOutput
	ToArtifactPropertyWebhookMapOutputWithContext(context.Context) ArtifactPropertyWebhookMapOutput
}

type ArtifactPropertyWebhookMap map[string]ArtifactPropertyWebhookInput

func (ArtifactPropertyWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ArtifactPropertyWebhook)(nil)).Elem()
}

func (i ArtifactPropertyWebhookMap) ToArtifactPropertyWebhookMapOutput() ArtifactPropertyWebhookMapOutput {
	return i.ToArtifactPropertyWebhookMapOutputWithContext(context.Background())
}

func (i ArtifactPropertyWebhookMap) ToArtifactPropertyWebhookMapOutputWithContext(ctx context.Context) ArtifactPropertyWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactPropertyWebhookMapOutput)
}

type ArtifactPropertyWebhookOutput struct{ *pulumi.OutputState }

func (ArtifactPropertyWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArtifactPropertyWebhook)(nil)).Elem()
}

func (o ArtifactPropertyWebhookOutput) ToArtifactPropertyWebhookOutput() ArtifactPropertyWebhookOutput {
	return o
}

func (o ArtifactPropertyWebhookOutput) ToArtifactPropertyWebhookOutputWithContext(ctx context.Context) ArtifactPropertyWebhookOutput {
	return o
}

// Specifies where the webhook will be applied on which repositories.
func (o ArtifactPropertyWebhookOutput) Criteria() ArtifactPropertyWebhookCriteriaOutput {
	return o.ApplyT(func(v *ArtifactPropertyWebhook) ArtifactPropertyWebhookCriteriaOutput { return v.Criteria }).(ArtifactPropertyWebhookCriteriaOutput)
}

// Webhook description. Max length 1000 characters.
func (o ArtifactPropertyWebhookOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactPropertyWebhook) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Status of webhook. Default to `true`.
func (o ArtifactPropertyWebhookOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ArtifactPropertyWebhook) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `added`, `deleted`.
func (o ArtifactPropertyWebhookOutput) EventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ArtifactPropertyWebhook) pulumi.StringArrayOutput { return v.EventTypes }).(pulumi.StringArrayOutput)
}

// At least one is required.
func (o ArtifactPropertyWebhookOutput) Handlers() ArtifactPropertyWebhookHandlerArrayOutput {
	return o.ApplyT(func(v *ArtifactPropertyWebhook) ArtifactPropertyWebhookHandlerArrayOutput { return v.Handlers }).(ArtifactPropertyWebhookHandlerArrayOutput)
}

// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
func (o ArtifactPropertyWebhookOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *ArtifactPropertyWebhook) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

type ArtifactPropertyWebhookArrayOutput struct{ *pulumi.OutputState }

func (ArtifactPropertyWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ArtifactPropertyWebhook)(nil)).Elem()
}

func (o ArtifactPropertyWebhookArrayOutput) ToArtifactPropertyWebhookArrayOutput() ArtifactPropertyWebhookArrayOutput {
	return o
}

func (o ArtifactPropertyWebhookArrayOutput) ToArtifactPropertyWebhookArrayOutputWithContext(ctx context.Context) ArtifactPropertyWebhookArrayOutput {
	return o
}

func (o ArtifactPropertyWebhookArrayOutput) Index(i pulumi.IntInput) ArtifactPropertyWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ArtifactPropertyWebhook {
		return vs[0].([]*ArtifactPropertyWebhook)[vs[1].(int)]
	}).(ArtifactPropertyWebhookOutput)
}

type ArtifactPropertyWebhookMapOutput struct{ *pulumi.OutputState }

func (ArtifactPropertyWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ArtifactPropertyWebhook)(nil)).Elem()
}

func (o ArtifactPropertyWebhookMapOutput) ToArtifactPropertyWebhookMapOutput() ArtifactPropertyWebhookMapOutput {
	return o
}

func (o ArtifactPropertyWebhookMapOutput) ToArtifactPropertyWebhookMapOutputWithContext(ctx context.Context) ArtifactPropertyWebhookMapOutput {
	return o
}

func (o ArtifactPropertyWebhookMapOutput) MapIndex(k pulumi.StringInput) ArtifactPropertyWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ArtifactPropertyWebhook {
		return vs[0].(map[string]*ArtifactPropertyWebhook)[vs[1].(string)]
	}).(ArtifactPropertyWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ArtifactPropertyWebhookInput)(nil)).Elem(), &ArtifactPropertyWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*ArtifactPropertyWebhookArrayInput)(nil)).Elem(), ArtifactPropertyWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ArtifactPropertyWebhookMapInput)(nil)).Elem(), ArtifactPropertyWebhookMap{})
	pulumi.RegisterOutputType(ArtifactPropertyWebhookOutput{})
	pulumi.RegisterOutputType(ArtifactPropertyWebhookArrayOutput{})
	pulumi.RegisterOutputType(ArtifactPropertyWebhookMapOutput{})
}
