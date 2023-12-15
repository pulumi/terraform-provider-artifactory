// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-artifactory/sdk/v6/go/artifactory/internal"
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
//	"github.com/pulumi/pulumi-artifactory/sdk/v6/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewDistributionWebhook(ctx, "distribution-webhook", &artifactory.DistributionWebhookArgs{
//				Criteria: &artifactory.DistributionWebhookCriteriaArgs{
//					AnyReleaseBundle: pulumi.Bool(false),
//					ExcludePatterns: pulumi.StringArray{
//						pulumi.String("bar/**"),
//					},
//					IncludePatterns: pulumi.StringArray{
//						pulumi.String("foo/**"),
//					},
//					RegisteredReleaseBundleNames: pulumi.StringArray{
//						pulumi.String("bundle-name"),
//					},
//				},
//				EventTypes: pulumi.StringArray{
//					pulumi.String("distribute_started"),
//					pulumi.String("distribute_completed"),
//					pulumi.String("distribute_aborted"),
//					pulumi.String("distribute_failed"),
//					pulumi.String("delete_started"),
//					pulumi.String("delete_completed"),
//					pulumi.String("delete_failed"),
//				},
//				Handlers: artifactory.DistributionWebhookHandlerArray{
//					&artifactory.DistributionWebhookHandlerArgs{
//						CustomHttpHeaders: pulumi.StringMap{
//							"header-1": pulumi.String("value-1"),
//							"header-2": pulumi.String("value-2"),
//						},
//						Proxy:  pulumi.String("proxy-key"),
//						Secret: pulumi.String("some-secret"),
//						Url:    pulumi.String("http://tempurl.org/webhook"),
//					},
//				},
//				Key: pulumi.String("distribution-webhook"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type DistributionWebhook struct {
	pulumi.CustomResourceState

	// Specifies where the webhook will be applied on which repositories.
	Criteria DistributionWebhookCriteriaOutput `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `distributeStarted`, `distributeCompleted`, `distributeAborted`, ` distribute_failed,  `deleteStarted` ,  `deleteCompleted` ,  `deleteFailed`
	EventTypes pulumi.StringArrayOutput `pulumi:"eventTypes"`
	// At least one is required.
	Handlers DistributionWebhookHandlerArrayOutput `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringOutput `pulumi:"key"`
}

// NewDistributionWebhook registers a new resource with the given unique name, arguments, and options.
func NewDistributionWebhook(ctx *pulumi.Context,
	name string, args *DistributionWebhookArgs, opts ...pulumi.ResourceOption) (*DistributionWebhook, error) {
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
	var resource DistributionWebhook
	err := ctx.RegisterResource("artifactory:index/distributionWebhook:DistributionWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDistributionWebhook gets an existing DistributionWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDistributionWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DistributionWebhookState, opts ...pulumi.ResourceOption) (*DistributionWebhook, error) {
	var resource DistributionWebhook
	err := ctx.ReadResource("artifactory:index/distributionWebhook:DistributionWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DistributionWebhook resources.
type distributionWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria *DistributionWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `distributeStarted`, `distributeCompleted`, `distributeAborted`, ` distribute_failed,  `deleteStarted` ,  `deleteCompleted` ,  `deleteFailed`
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []DistributionWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key *string `pulumi:"key"`
}

type DistributionWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria DistributionWebhookCriteriaPtrInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `distributeStarted`, `distributeCompleted`, `distributeAborted`, ` distribute_failed,  `deleteStarted` ,  `deleteCompleted` ,  `deleteFailed`
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers DistributionWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringPtrInput
}

func (DistributionWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*distributionWebhookState)(nil)).Elem()
}

type distributionWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria DistributionWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `distributeStarted`, `distributeCompleted`, `distributeAborted`, ` distribute_failed,  `deleteStarted` ,  `deleteCompleted` ,  `deleteFailed`
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []DistributionWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key string `pulumi:"key"`
}

// The set of arguments for constructing a DistributionWebhook resource.
type DistributionWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria DistributionWebhookCriteriaInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `distributeStarted`, `distributeCompleted`, `distributeAborted`, ` distribute_failed,  `deleteStarted` ,  `deleteCompleted` ,  `deleteFailed`
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers DistributionWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringInput
}

func (DistributionWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*distributionWebhookArgs)(nil)).Elem()
}

type DistributionWebhookInput interface {
	pulumi.Input

	ToDistributionWebhookOutput() DistributionWebhookOutput
	ToDistributionWebhookOutputWithContext(ctx context.Context) DistributionWebhookOutput
}

func (*DistributionWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**DistributionWebhook)(nil)).Elem()
}

func (i *DistributionWebhook) ToDistributionWebhookOutput() DistributionWebhookOutput {
	return i.ToDistributionWebhookOutputWithContext(context.Background())
}

func (i *DistributionWebhook) ToDistributionWebhookOutputWithContext(ctx context.Context) DistributionWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DistributionWebhookOutput)
}

// DistributionWebhookArrayInput is an input type that accepts DistributionWebhookArray and DistributionWebhookArrayOutput values.
// You can construct a concrete instance of `DistributionWebhookArrayInput` via:
//
//	DistributionWebhookArray{ DistributionWebhookArgs{...} }
type DistributionWebhookArrayInput interface {
	pulumi.Input

	ToDistributionWebhookArrayOutput() DistributionWebhookArrayOutput
	ToDistributionWebhookArrayOutputWithContext(context.Context) DistributionWebhookArrayOutput
}

type DistributionWebhookArray []DistributionWebhookInput

func (DistributionWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DistributionWebhook)(nil)).Elem()
}

func (i DistributionWebhookArray) ToDistributionWebhookArrayOutput() DistributionWebhookArrayOutput {
	return i.ToDistributionWebhookArrayOutputWithContext(context.Background())
}

func (i DistributionWebhookArray) ToDistributionWebhookArrayOutputWithContext(ctx context.Context) DistributionWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DistributionWebhookArrayOutput)
}

// DistributionWebhookMapInput is an input type that accepts DistributionWebhookMap and DistributionWebhookMapOutput values.
// You can construct a concrete instance of `DistributionWebhookMapInput` via:
//
//	DistributionWebhookMap{ "key": DistributionWebhookArgs{...} }
type DistributionWebhookMapInput interface {
	pulumi.Input

	ToDistributionWebhookMapOutput() DistributionWebhookMapOutput
	ToDistributionWebhookMapOutputWithContext(context.Context) DistributionWebhookMapOutput
}

type DistributionWebhookMap map[string]DistributionWebhookInput

func (DistributionWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DistributionWebhook)(nil)).Elem()
}

func (i DistributionWebhookMap) ToDistributionWebhookMapOutput() DistributionWebhookMapOutput {
	return i.ToDistributionWebhookMapOutputWithContext(context.Background())
}

func (i DistributionWebhookMap) ToDistributionWebhookMapOutputWithContext(ctx context.Context) DistributionWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DistributionWebhookMapOutput)
}

type DistributionWebhookOutput struct{ *pulumi.OutputState }

func (DistributionWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DistributionWebhook)(nil)).Elem()
}

func (o DistributionWebhookOutput) ToDistributionWebhookOutput() DistributionWebhookOutput {
	return o
}

func (o DistributionWebhookOutput) ToDistributionWebhookOutputWithContext(ctx context.Context) DistributionWebhookOutput {
	return o
}

// Specifies where the webhook will be applied on which repositories.
func (o DistributionWebhookOutput) Criteria() DistributionWebhookCriteriaOutput {
	return o.ApplyT(func(v *DistributionWebhook) DistributionWebhookCriteriaOutput { return v.Criteria }).(DistributionWebhookCriteriaOutput)
}

// Webhook description. Max length 1000 characters.
func (o DistributionWebhookOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DistributionWebhook) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Status of webhook. Default to `true`.
func (o DistributionWebhookOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DistributionWebhook) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `distributeStarted`, `distributeCompleted`, `distributeAborted`, ` distribute_failed,  `deleteStarted` ,  `deleteCompleted` ,  `deleteFailed`
func (o DistributionWebhookOutput) EventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DistributionWebhook) pulumi.StringArrayOutput { return v.EventTypes }).(pulumi.StringArrayOutput)
}

// At least one is required.
func (o DistributionWebhookOutput) Handlers() DistributionWebhookHandlerArrayOutput {
	return o.ApplyT(func(v *DistributionWebhook) DistributionWebhookHandlerArrayOutput { return v.Handlers }).(DistributionWebhookHandlerArrayOutput)
}

// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
func (o DistributionWebhookOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributionWebhook) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

type DistributionWebhookArrayOutput struct{ *pulumi.OutputState }

func (DistributionWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DistributionWebhook)(nil)).Elem()
}

func (o DistributionWebhookArrayOutput) ToDistributionWebhookArrayOutput() DistributionWebhookArrayOutput {
	return o
}

func (o DistributionWebhookArrayOutput) ToDistributionWebhookArrayOutputWithContext(ctx context.Context) DistributionWebhookArrayOutput {
	return o
}

func (o DistributionWebhookArrayOutput) Index(i pulumi.IntInput) DistributionWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DistributionWebhook {
		return vs[0].([]*DistributionWebhook)[vs[1].(int)]
	}).(DistributionWebhookOutput)
}

type DistributionWebhookMapOutput struct{ *pulumi.OutputState }

func (DistributionWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DistributionWebhook)(nil)).Elem()
}

func (o DistributionWebhookMapOutput) ToDistributionWebhookMapOutput() DistributionWebhookMapOutput {
	return o
}

func (o DistributionWebhookMapOutput) ToDistributionWebhookMapOutputWithContext(ctx context.Context) DistributionWebhookMapOutput {
	return o
}

func (o DistributionWebhookMapOutput) MapIndex(k pulumi.StringInput) DistributionWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DistributionWebhook {
		return vs[0].(map[string]*DistributionWebhook)[vs[1].(string)]
	}).(DistributionWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DistributionWebhookInput)(nil)).Elem(), &DistributionWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*DistributionWebhookArrayInput)(nil)).Elem(), DistributionWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DistributionWebhookMapInput)(nil)).Elem(), DistributionWebhookMap{})
	pulumi.RegisterOutputType(DistributionWebhookOutput{})
	pulumi.RegisterOutputType(DistributionWebhookArrayOutput{})
	pulumi.RegisterOutputType(DistributionWebhookMapOutput{})
}
