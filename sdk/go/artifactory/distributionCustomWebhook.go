// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-artifactory/sdk/v4/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides an Artifactory custom webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v4/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewDistributionCustomWebhook(ctx, "distribution-custom-webhook", &artifactory.DistributionCustomWebhookArgs{
//				Criteria: &artifactory.DistributionCustomWebhookCriteriaArgs{
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
//				Handlers: artifactory.DistributionCustomWebhookHandlerArray{
//					&artifactory.DistributionCustomWebhookHandlerArgs{
//						HttpHeaders: pulumi.StringMap{
//							"headerName1": pulumi.String("value1"),
//							"headerName2": pulumi.String("value2"),
//						},
//						Payload: pulumi.String("{ \"ref\": \"main\" , \"inputs\": { \"artifact_path\": \"test-repo/repo-path\" } }"),
//						Secrets: pulumi.StringMap{
//							"secretName1": pulumi.String("value1"),
//							"secretName2": pulumi.String("value2"),
//						},
//						Url: pulumi.String("https://tempurl.org"),
//					},
//				},
//				Key: pulumi.String("distribution-custom-webhook"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type DistributionCustomWebhook struct {
	pulumi.CustomResourceState

	// Specifies where the webhook will be applied on which repositories.
	Criteria DistributionCustomWebhookCriteriaOutput `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `distributeStarted`, `distributeCompleted`, `distributeAborted`, ` distribute_failed,  `deleteStarted` ,  `deleteCompleted` ,  `deleteFailed`
	EventTypes pulumi.StringArrayOutput `pulumi:"eventTypes"`
	// At least one is required.
	Handlers DistributionCustomWebhookHandlerArrayOutput `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringOutput `pulumi:"key"`
}

// NewDistributionCustomWebhook registers a new resource with the given unique name, arguments, and options.
func NewDistributionCustomWebhook(ctx *pulumi.Context,
	name string, args *DistributionCustomWebhookArgs, opts ...pulumi.ResourceOption) (*DistributionCustomWebhook, error) {
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
	var resource DistributionCustomWebhook
	err := ctx.RegisterResource("artifactory:index/distributionCustomWebhook:DistributionCustomWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDistributionCustomWebhook gets an existing DistributionCustomWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDistributionCustomWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DistributionCustomWebhookState, opts ...pulumi.ResourceOption) (*DistributionCustomWebhook, error) {
	var resource DistributionCustomWebhook
	err := ctx.ReadResource("artifactory:index/distributionCustomWebhook:DistributionCustomWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DistributionCustomWebhook resources.
type distributionCustomWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria *DistributionCustomWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `distributeStarted`, `distributeCompleted`, `distributeAborted`, ` distribute_failed,  `deleteStarted` ,  `deleteCompleted` ,  `deleteFailed`
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []DistributionCustomWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key *string `pulumi:"key"`
}

type DistributionCustomWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria DistributionCustomWebhookCriteriaPtrInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `distributeStarted`, `distributeCompleted`, `distributeAborted`, ` distribute_failed,  `deleteStarted` ,  `deleteCompleted` ,  `deleteFailed`
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers DistributionCustomWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringPtrInput
}

func (DistributionCustomWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*distributionCustomWebhookState)(nil)).Elem()
}

type distributionCustomWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria DistributionCustomWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `distributeStarted`, `distributeCompleted`, `distributeAborted`, ` distribute_failed,  `deleteStarted` ,  `deleteCompleted` ,  `deleteFailed`
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []DistributionCustomWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key string `pulumi:"key"`
}

// The set of arguments for constructing a DistributionCustomWebhook resource.
type DistributionCustomWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria DistributionCustomWebhookCriteriaInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `distributeStarted`, `distributeCompleted`, `distributeAborted`, ` distribute_failed,  `deleteStarted` ,  `deleteCompleted` ,  `deleteFailed`
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers DistributionCustomWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringInput
}

func (DistributionCustomWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*distributionCustomWebhookArgs)(nil)).Elem()
}

type DistributionCustomWebhookInput interface {
	pulumi.Input

	ToDistributionCustomWebhookOutput() DistributionCustomWebhookOutput
	ToDistributionCustomWebhookOutputWithContext(ctx context.Context) DistributionCustomWebhookOutput
}

func (*DistributionCustomWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**DistributionCustomWebhook)(nil)).Elem()
}

func (i *DistributionCustomWebhook) ToDistributionCustomWebhookOutput() DistributionCustomWebhookOutput {
	return i.ToDistributionCustomWebhookOutputWithContext(context.Background())
}

func (i *DistributionCustomWebhook) ToDistributionCustomWebhookOutputWithContext(ctx context.Context) DistributionCustomWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DistributionCustomWebhookOutput)
}

func (i *DistributionCustomWebhook) ToOutput(ctx context.Context) pulumix.Output[*DistributionCustomWebhook] {
	return pulumix.Output[*DistributionCustomWebhook]{
		OutputState: i.ToDistributionCustomWebhookOutputWithContext(ctx).OutputState,
	}
}

// DistributionCustomWebhookArrayInput is an input type that accepts DistributionCustomWebhookArray and DistributionCustomWebhookArrayOutput values.
// You can construct a concrete instance of `DistributionCustomWebhookArrayInput` via:
//
//	DistributionCustomWebhookArray{ DistributionCustomWebhookArgs{...} }
type DistributionCustomWebhookArrayInput interface {
	pulumi.Input

	ToDistributionCustomWebhookArrayOutput() DistributionCustomWebhookArrayOutput
	ToDistributionCustomWebhookArrayOutputWithContext(context.Context) DistributionCustomWebhookArrayOutput
}

type DistributionCustomWebhookArray []DistributionCustomWebhookInput

func (DistributionCustomWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DistributionCustomWebhook)(nil)).Elem()
}

func (i DistributionCustomWebhookArray) ToDistributionCustomWebhookArrayOutput() DistributionCustomWebhookArrayOutput {
	return i.ToDistributionCustomWebhookArrayOutputWithContext(context.Background())
}

func (i DistributionCustomWebhookArray) ToDistributionCustomWebhookArrayOutputWithContext(ctx context.Context) DistributionCustomWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DistributionCustomWebhookArrayOutput)
}

func (i DistributionCustomWebhookArray) ToOutput(ctx context.Context) pulumix.Output[[]*DistributionCustomWebhook] {
	return pulumix.Output[[]*DistributionCustomWebhook]{
		OutputState: i.ToDistributionCustomWebhookArrayOutputWithContext(ctx).OutputState,
	}
}

// DistributionCustomWebhookMapInput is an input type that accepts DistributionCustomWebhookMap and DistributionCustomWebhookMapOutput values.
// You can construct a concrete instance of `DistributionCustomWebhookMapInput` via:
//
//	DistributionCustomWebhookMap{ "key": DistributionCustomWebhookArgs{...} }
type DistributionCustomWebhookMapInput interface {
	pulumi.Input

	ToDistributionCustomWebhookMapOutput() DistributionCustomWebhookMapOutput
	ToDistributionCustomWebhookMapOutputWithContext(context.Context) DistributionCustomWebhookMapOutput
}

type DistributionCustomWebhookMap map[string]DistributionCustomWebhookInput

func (DistributionCustomWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DistributionCustomWebhook)(nil)).Elem()
}

func (i DistributionCustomWebhookMap) ToDistributionCustomWebhookMapOutput() DistributionCustomWebhookMapOutput {
	return i.ToDistributionCustomWebhookMapOutputWithContext(context.Background())
}

func (i DistributionCustomWebhookMap) ToDistributionCustomWebhookMapOutputWithContext(ctx context.Context) DistributionCustomWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DistributionCustomWebhookMapOutput)
}

func (i DistributionCustomWebhookMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*DistributionCustomWebhook] {
	return pulumix.Output[map[string]*DistributionCustomWebhook]{
		OutputState: i.ToDistributionCustomWebhookMapOutputWithContext(ctx).OutputState,
	}
}

type DistributionCustomWebhookOutput struct{ *pulumi.OutputState }

func (DistributionCustomWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DistributionCustomWebhook)(nil)).Elem()
}

func (o DistributionCustomWebhookOutput) ToDistributionCustomWebhookOutput() DistributionCustomWebhookOutput {
	return o
}

func (o DistributionCustomWebhookOutput) ToDistributionCustomWebhookOutputWithContext(ctx context.Context) DistributionCustomWebhookOutput {
	return o
}

func (o DistributionCustomWebhookOutput) ToOutput(ctx context.Context) pulumix.Output[*DistributionCustomWebhook] {
	return pulumix.Output[*DistributionCustomWebhook]{
		OutputState: o.OutputState,
	}
}

// Specifies where the webhook will be applied on which repositories.
func (o DistributionCustomWebhookOutput) Criteria() DistributionCustomWebhookCriteriaOutput {
	return o.ApplyT(func(v *DistributionCustomWebhook) DistributionCustomWebhookCriteriaOutput { return v.Criteria }).(DistributionCustomWebhookCriteriaOutput)
}

// Webhook description. Max length 1000 characters.
func (o DistributionCustomWebhookOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DistributionCustomWebhook) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Status of webhook. Default to `true`.
func (o DistributionCustomWebhookOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DistributionCustomWebhook) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `distributeStarted`, `distributeCompleted`, `distributeAborted`, ` distribute_failed,  `deleteStarted` ,  `deleteCompleted` ,  `deleteFailed`
func (o DistributionCustomWebhookOutput) EventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DistributionCustomWebhook) pulumi.StringArrayOutput { return v.EventTypes }).(pulumi.StringArrayOutput)
}

// At least one is required.
func (o DistributionCustomWebhookOutput) Handlers() DistributionCustomWebhookHandlerArrayOutput {
	return o.ApplyT(func(v *DistributionCustomWebhook) DistributionCustomWebhookHandlerArrayOutput { return v.Handlers }).(DistributionCustomWebhookHandlerArrayOutput)
}

// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
func (o DistributionCustomWebhookOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributionCustomWebhook) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

type DistributionCustomWebhookArrayOutput struct{ *pulumi.OutputState }

func (DistributionCustomWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DistributionCustomWebhook)(nil)).Elem()
}

func (o DistributionCustomWebhookArrayOutput) ToDistributionCustomWebhookArrayOutput() DistributionCustomWebhookArrayOutput {
	return o
}

func (o DistributionCustomWebhookArrayOutput) ToDistributionCustomWebhookArrayOutputWithContext(ctx context.Context) DistributionCustomWebhookArrayOutput {
	return o
}

func (o DistributionCustomWebhookArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*DistributionCustomWebhook] {
	return pulumix.Output[[]*DistributionCustomWebhook]{
		OutputState: o.OutputState,
	}
}

func (o DistributionCustomWebhookArrayOutput) Index(i pulumi.IntInput) DistributionCustomWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DistributionCustomWebhook {
		return vs[0].([]*DistributionCustomWebhook)[vs[1].(int)]
	}).(DistributionCustomWebhookOutput)
}

type DistributionCustomWebhookMapOutput struct{ *pulumi.OutputState }

func (DistributionCustomWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DistributionCustomWebhook)(nil)).Elem()
}

func (o DistributionCustomWebhookMapOutput) ToDistributionCustomWebhookMapOutput() DistributionCustomWebhookMapOutput {
	return o
}

func (o DistributionCustomWebhookMapOutput) ToDistributionCustomWebhookMapOutputWithContext(ctx context.Context) DistributionCustomWebhookMapOutput {
	return o
}

func (o DistributionCustomWebhookMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*DistributionCustomWebhook] {
	return pulumix.Output[map[string]*DistributionCustomWebhook]{
		OutputState: o.OutputState,
	}
}

func (o DistributionCustomWebhookMapOutput) MapIndex(k pulumi.StringInput) DistributionCustomWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DistributionCustomWebhook {
		return vs[0].(map[string]*DistributionCustomWebhook)[vs[1].(string)]
	}).(DistributionCustomWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DistributionCustomWebhookInput)(nil)).Elem(), &DistributionCustomWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*DistributionCustomWebhookArrayInput)(nil)).Elem(), DistributionCustomWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DistributionCustomWebhookMapInput)(nil)).Elem(), DistributionCustomWebhookMap{})
	pulumi.RegisterOutputType(DistributionCustomWebhookOutput{})
	pulumi.RegisterOutputType(DistributionCustomWebhookArrayOutput{})
	pulumi.RegisterOutputType(DistributionCustomWebhookMapOutput{})
}
