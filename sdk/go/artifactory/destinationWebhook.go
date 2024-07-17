// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-artifactory/sdk/v7/go/artifactory/internal"
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
//	"github.com/pulumi/pulumi-artifactory/sdk/v7/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewDestinationWebhook(ctx, "destination-webhook", &artifactory.DestinationWebhookArgs{
//				Key: pulumi.String("destination-webhook"),
//				EventTypes: pulumi.StringArray{
//					pulumi.String("received"),
//					pulumi.String("delete_started"),
//					pulumi.String("delete_completed"),
//					pulumi.String("delete_failed"),
//				},
//				Criteria: &artifactory.DestinationWebhookCriteriaArgs{
//					AnyReleaseBundle: pulumi.Bool(false),
//					RegisteredReleaseBundleNames: pulumi.StringArray{
//						pulumi.String("bundle-name"),
//					},
//					IncludePatterns: pulumi.StringArray{
//						pulumi.String("foo/**"),
//					},
//					ExcludePatterns: pulumi.StringArray{
//						pulumi.String("bar/**"),
//					},
//				},
//				Handlers: artifactory.DestinationWebhookHandlerArray{
//					&artifactory.DestinationWebhookHandlerArgs{
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
type DestinationWebhook struct {
	pulumi.CustomResourceState

	// Specifies where the webhook will be applied on which repositories.
	Criteria DestinationWebhookCriteriaOutput `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Status of webhook. Default to `true`
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `received`, `deleteStarted`, `deleteCompleted`, `deleteFailed`
	EventTypes pulumi.StringArrayOutput `pulumi:"eventTypes"`
	// At least one is required.
	Handlers DestinationWebhookHandlerArrayOutput `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringOutput `pulumi:"key"`
}

// NewDestinationWebhook registers a new resource with the given unique name, arguments, and options.
func NewDestinationWebhook(ctx *pulumi.Context,
	name string, args *DestinationWebhookArgs, opts ...pulumi.ResourceOption) (*DestinationWebhook, error) {
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
	var resource DestinationWebhook
	err := ctx.RegisterResource("artifactory:index/destinationWebhook:DestinationWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDestinationWebhook gets an existing DestinationWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDestinationWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DestinationWebhookState, opts ...pulumi.ResourceOption) (*DestinationWebhook, error) {
	var resource DestinationWebhook
	err := ctx.ReadResource("artifactory:index/destinationWebhook:DestinationWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DestinationWebhook resources.
type destinationWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria *DestinationWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `received`, `deleteStarted`, `deleteCompleted`, `deleteFailed`
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []DestinationWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key *string `pulumi:"key"`
}

type DestinationWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria DestinationWebhookCriteriaPtrInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `received`, `deleteStarted`, `deleteCompleted`, `deleteFailed`
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers DestinationWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringPtrInput
}

func (DestinationWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationWebhookState)(nil)).Elem()
}

type destinationWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria DestinationWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `received`, `deleteStarted`, `deleteCompleted`, `deleteFailed`
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []DestinationWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key string `pulumi:"key"`
}

// The set of arguments for constructing a DestinationWebhook resource.
type DestinationWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria DestinationWebhookCriteriaInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `received`, `deleteStarted`, `deleteCompleted`, `deleteFailed`
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers DestinationWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringInput
}

func (DestinationWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationWebhookArgs)(nil)).Elem()
}

type DestinationWebhookInput interface {
	pulumi.Input

	ToDestinationWebhookOutput() DestinationWebhookOutput
	ToDestinationWebhookOutputWithContext(ctx context.Context) DestinationWebhookOutput
}

func (*DestinationWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationWebhook)(nil)).Elem()
}

func (i *DestinationWebhook) ToDestinationWebhookOutput() DestinationWebhookOutput {
	return i.ToDestinationWebhookOutputWithContext(context.Background())
}

func (i *DestinationWebhook) ToDestinationWebhookOutputWithContext(ctx context.Context) DestinationWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationWebhookOutput)
}

// DestinationWebhookArrayInput is an input type that accepts DestinationWebhookArray and DestinationWebhookArrayOutput values.
// You can construct a concrete instance of `DestinationWebhookArrayInput` via:
//
//	DestinationWebhookArray{ DestinationWebhookArgs{...} }
type DestinationWebhookArrayInput interface {
	pulumi.Input

	ToDestinationWebhookArrayOutput() DestinationWebhookArrayOutput
	ToDestinationWebhookArrayOutputWithContext(context.Context) DestinationWebhookArrayOutput
}

type DestinationWebhookArray []DestinationWebhookInput

func (DestinationWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DestinationWebhook)(nil)).Elem()
}

func (i DestinationWebhookArray) ToDestinationWebhookArrayOutput() DestinationWebhookArrayOutput {
	return i.ToDestinationWebhookArrayOutputWithContext(context.Background())
}

func (i DestinationWebhookArray) ToDestinationWebhookArrayOutputWithContext(ctx context.Context) DestinationWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationWebhookArrayOutput)
}

// DestinationWebhookMapInput is an input type that accepts DestinationWebhookMap and DestinationWebhookMapOutput values.
// You can construct a concrete instance of `DestinationWebhookMapInput` via:
//
//	DestinationWebhookMap{ "key": DestinationWebhookArgs{...} }
type DestinationWebhookMapInput interface {
	pulumi.Input

	ToDestinationWebhookMapOutput() DestinationWebhookMapOutput
	ToDestinationWebhookMapOutputWithContext(context.Context) DestinationWebhookMapOutput
}

type DestinationWebhookMap map[string]DestinationWebhookInput

func (DestinationWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DestinationWebhook)(nil)).Elem()
}

func (i DestinationWebhookMap) ToDestinationWebhookMapOutput() DestinationWebhookMapOutput {
	return i.ToDestinationWebhookMapOutputWithContext(context.Background())
}

func (i DestinationWebhookMap) ToDestinationWebhookMapOutputWithContext(ctx context.Context) DestinationWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationWebhookMapOutput)
}

type DestinationWebhookOutput struct{ *pulumi.OutputState }

func (DestinationWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationWebhook)(nil)).Elem()
}

func (o DestinationWebhookOutput) ToDestinationWebhookOutput() DestinationWebhookOutput {
	return o
}

func (o DestinationWebhookOutput) ToDestinationWebhookOutputWithContext(ctx context.Context) DestinationWebhookOutput {
	return o
}

// Specifies where the webhook will be applied on which repositories.
func (o DestinationWebhookOutput) Criteria() DestinationWebhookCriteriaOutput {
	return o.ApplyT(func(v *DestinationWebhook) DestinationWebhookCriteriaOutput { return v.Criteria }).(DestinationWebhookCriteriaOutput)
}

// Webhook description. Max length 1000 characters.
func (o DestinationWebhookOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DestinationWebhook) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Status of webhook. Default to `true`
func (o DestinationWebhookOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DestinationWebhook) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `received`, `deleteStarted`, `deleteCompleted`, `deleteFailed`
func (o DestinationWebhookOutput) EventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DestinationWebhook) pulumi.StringArrayOutput { return v.EventTypes }).(pulumi.StringArrayOutput)
}

// At least one is required.
func (o DestinationWebhookOutput) Handlers() DestinationWebhookHandlerArrayOutput {
	return o.ApplyT(func(v *DestinationWebhook) DestinationWebhookHandlerArrayOutput { return v.Handlers }).(DestinationWebhookHandlerArrayOutput)
}

// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
func (o DestinationWebhookOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationWebhook) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

type DestinationWebhookArrayOutput struct{ *pulumi.OutputState }

func (DestinationWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DestinationWebhook)(nil)).Elem()
}

func (o DestinationWebhookArrayOutput) ToDestinationWebhookArrayOutput() DestinationWebhookArrayOutput {
	return o
}

func (o DestinationWebhookArrayOutput) ToDestinationWebhookArrayOutputWithContext(ctx context.Context) DestinationWebhookArrayOutput {
	return o
}

func (o DestinationWebhookArrayOutput) Index(i pulumi.IntInput) DestinationWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DestinationWebhook {
		return vs[0].([]*DestinationWebhook)[vs[1].(int)]
	}).(DestinationWebhookOutput)
}

type DestinationWebhookMapOutput struct{ *pulumi.OutputState }

func (DestinationWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DestinationWebhook)(nil)).Elem()
}

func (o DestinationWebhookMapOutput) ToDestinationWebhookMapOutput() DestinationWebhookMapOutput {
	return o
}

func (o DestinationWebhookMapOutput) ToDestinationWebhookMapOutputWithContext(ctx context.Context) DestinationWebhookMapOutput {
	return o
}

func (o DestinationWebhookMapOutput) MapIndex(k pulumi.StringInput) DestinationWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DestinationWebhook {
		return vs[0].(map[string]*DestinationWebhook)[vs[1].(string)]
	}).(DestinationWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationWebhookInput)(nil)).Elem(), &DestinationWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationWebhookArrayInput)(nil)).Elem(), DestinationWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationWebhookMapInput)(nil)).Elem(), DestinationWebhookMap{})
	pulumi.RegisterOutputType(DestinationWebhookOutput{})
	pulumi.RegisterOutputType(DestinationWebhookArrayOutput{})
	pulumi.RegisterOutputType(DestinationWebhookMapOutput{})
}
