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

// Provides an Artifactory custom webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.
//
// !>This resource is being deprecated and replaced by `DestinationCustomWebhook` resource.
//
// ## Example Usage
//
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
//			_, err := artifactory.NewReleaseBundleCustomWebhook(ctx, "release-bundle-custom-webhook", &artifactory.ReleaseBundleCustomWebhookArgs{
//				Key: pulumi.String("release-bundle-custom-webhook"),
//				EventTypes: pulumi.StringArray{
//					pulumi.String("created"),
//					pulumi.String("signed"),
//					pulumi.String("deleted"),
//				},
//				Criteria: &artifactory.ReleaseBundleCustomWebhookCriteriaArgs{
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
//				Handlers: artifactory.ReleaseBundleCustomWebhookHandlerArray{
//					&artifactory.ReleaseBundleCustomWebhookHandlerArgs{
//						Url: pulumi.String("https://tempurl.org"),
//						Secrets: pulumi.StringMap{
//							"secretName1": pulumi.String("value1"),
//							"secretName2": pulumi.String("value2"),
//						},
//						HttpHeaders: pulumi.StringMap{
//							"headerName1": pulumi.String("value1"),
//							"headerName2": pulumi.String("value2"),
//						},
//						Payload: pulumi.String("{ \"ref\": \"main\" , \"inputs\": { \"artifact_path\": \"test-repo/repo-path\" } }"),
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
type ReleaseBundleCustomWebhook struct {
	pulumi.CustomResourceState

	// Specifies where the webhook will be applied on which repositories.
	Criteria ReleaseBundleCustomWebhookCriteriaOutput `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `created`, `signed`, `deleted`.
	EventTypes pulumi.StringArrayOutput `pulumi:"eventTypes"`
	// At least one is required.
	Handlers ReleaseBundleCustomWebhookHandlerArrayOutput `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringOutput `pulumi:"key"`
}

// NewReleaseBundleCustomWebhook registers a new resource with the given unique name, arguments, and options.
func NewReleaseBundleCustomWebhook(ctx *pulumi.Context,
	name string, args *ReleaseBundleCustomWebhookArgs, opts ...pulumi.ResourceOption) (*ReleaseBundleCustomWebhook, error) {
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
	var resource ReleaseBundleCustomWebhook
	err := ctx.RegisterResource("artifactory:index/releaseBundleCustomWebhook:ReleaseBundleCustomWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReleaseBundleCustomWebhook gets an existing ReleaseBundleCustomWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReleaseBundleCustomWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReleaseBundleCustomWebhookState, opts ...pulumi.ResourceOption) (*ReleaseBundleCustomWebhook, error) {
	var resource ReleaseBundleCustomWebhook
	err := ctx.ReadResource("artifactory:index/releaseBundleCustomWebhook:ReleaseBundleCustomWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReleaseBundleCustomWebhook resources.
type releaseBundleCustomWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria *ReleaseBundleCustomWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `created`, `signed`, `deleted`.
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []ReleaseBundleCustomWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key *string `pulumi:"key"`
}

type ReleaseBundleCustomWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria ReleaseBundleCustomWebhookCriteriaPtrInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `created`, `signed`, `deleted`.
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers ReleaseBundleCustomWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringPtrInput
}

func (ReleaseBundleCustomWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseBundleCustomWebhookState)(nil)).Elem()
}

type releaseBundleCustomWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria ReleaseBundleCustomWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `created`, `signed`, `deleted`.
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []ReleaseBundleCustomWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key string `pulumi:"key"`
}

// The set of arguments for constructing a ReleaseBundleCustomWebhook resource.
type ReleaseBundleCustomWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria ReleaseBundleCustomWebhookCriteriaInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `created`, `signed`, `deleted`.
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers ReleaseBundleCustomWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringInput
}

func (ReleaseBundleCustomWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseBundleCustomWebhookArgs)(nil)).Elem()
}

type ReleaseBundleCustomWebhookInput interface {
	pulumi.Input

	ToReleaseBundleCustomWebhookOutput() ReleaseBundleCustomWebhookOutput
	ToReleaseBundleCustomWebhookOutputWithContext(ctx context.Context) ReleaseBundleCustomWebhookOutput
}

func (*ReleaseBundleCustomWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**ReleaseBundleCustomWebhook)(nil)).Elem()
}

func (i *ReleaseBundleCustomWebhook) ToReleaseBundleCustomWebhookOutput() ReleaseBundleCustomWebhookOutput {
	return i.ToReleaseBundleCustomWebhookOutputWithContext(context.Background())
}

func (i *ReleaseBundleCustomWebhook) ToReleaseBundleCustomWebhookOutputWithContext(ctx context.Context) ReleaseBundleCustomWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseBundleCustomWebhookOutput)
}

// ReleaseBundleCustomWebhookArrayInput is an input type that accepts ReleaseBundleCustomWebhookArray and ReleaseBundleCustomWebhookArrayOutput values.
// You can construct a concrete instance of `ReleaseBundleCustomWebhookArrayInput` via:
//
//	ReleaseBundleCustomWebhookArray{ ReleaseBundleCustomWebhookArgs{...} }
type ReleaseBundleCustomWebhookArrayInput interface {
	pulumi.Input

	ToReleaseBundleCustomWebhookArrayOutput() ReleaseBundleCustomWebhookArrayOutput
	ToReleaseBundleCustomWebhookArrayOutputWithContext(context.Context) ReleaseBundleCustomWebhookArrayOutput
}

type ReleaseBundleCustomWebhookArray []ReleaseBundleCustomWebhookInput

func (ReleaseBundleCustomWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReleaseBundleCustomWebhook)(nil)).Elem()
}

func (i ReleaseBundleCustomWebhookArray) ToReleaseBundleCustomWebhookArrayOutput() ReleaseBundleCustomWebhookArrayOutput {
	return i.ToReleaseBundleCustomWebhookArrayOutputWithContext(context.Background())
}

func (i ReleaseBundleCustomWebhookArray) ToReleaseBundleCustomWebhookArrayOutputWithContext(ctx context.Context) ReleaseBundleCustomWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseBundleCustomWebhookArrayOutput)
}

// ReleaseBundleCustomWebhookMapInput is an input type that accepts ReleaseBundleCustomWebhookMap and ReleaseBundleCustomWebhookMapOutput values.
// You can construct a concrete instance of `ReleaseBundleCustomWebhookMapInput` via:
//
//	ReleaseBundleCustomWebhookMap{ "key": ReleaseBundleCustomWebhookArgs{...} }
type ReleaseBundleCustomWebhookMapInput interface {
	pulumi.Input

	ToReleaseBundleCustomWebhookMapOutput() ReleaseBundleCustomWebhookMapOutput
	ToReleaseBundleCustomWebhookMapOutputWithContext(context.Context) ReleaseBundleCustomWebhookMapOutput
}

type ReleaseBundleCustomWebhookMap map[string]ReleaseBundleCustomWebhookInput

func (ReleaseBundleCustomWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReleaseBundleCustomWebhook)(nil)).Elem()
}

func (i ReleaseBundleCustomWebhookMap) ToReleaseBundleCustomWebhookMapOutput() ReleaseBundleCustomWebhookMapOutput {
	return i.ToReleaseBundleCustomWebhookMapOutputWithContext(context.Background())
}

func (i ReleaseBundleCustomWebhookMap) ToReleaseBundleCustomWebhookMapOutputWithContext(ctx context.Context) ReleaseBundleCustomWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseBundleCustomWebhookMapOutput)
}

type ReleaseBundleCustomWebhookOutput struct{ *pulumi.OutputState }

func (ReleaseBundleCustomWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReleaseBundleCustomWebhook)(nil)).Elem()
}

func (o ReleaseBundleCustomWebhookOutput) ToReleaseBundleCustomWebhookOutput() ReleaseBundleCustomWebhookOutput {
	return o
}

func (o ReleaseBundleCustomWebhookOutput) ToReleaseBundleCustomWebhookOutputWithContext(ctx context.Context) ReleaseBundleCustomWebhookOutput {
	return o
}

// Specifies where the webhook will be applied on which repositories.
func (o ReleaseBundleCustomWebhookOutput) Criteria() ReleaseBundleCustomWebhookCriteriaOutput {
	return o.ApplyT(func(v *ReleaseBundleCustomWebhook) ReleaseBundleCustomWebhookCriteriaOutput { return v.Criteria }).(ReleaseBundleCustomWebhookCriteriaOutput)
}

// Webhook description. Max length 1000 characters.
func (o ReleaseBundleCustomWebhookOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReleaseBundleCustomWebhook) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Status of webhook. Default to `true`.
func (o ReleaseBundleCustomWebhookOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ReleaseBundleCustomWebhook) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `created`, `signed`, `deleted`.
func (o ReleaseBundleCustomWebhookOutput) EventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReleaseBundleCustomWebhook) pulumi.StringArrayOutput { return v.EventTypes }).(pulumi.StringArrayOutput)
}

// At least one is required.
func (o ReleaseBundleCustomWebhookOutput) Handlers() ReleaseBundleCustomWebhookHandlerArrayOutput {
	return o.ApplyT(func(v *ReleaseBundleCustomWebhook) ReleaseBundleCustomWebhookHandlerArrayOutput { return v.Handlers }).(ReleaseBundleCustomWebhookHandlerArrayOutput)
}

// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
func (o ReleaseBundleCustomWebhookOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *ReleaseBundleCustomWebhook) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

type ReleaseBundleCustomWebhookArrayOutput struct{ *pulumi.OutputState }

func (ReleaseBundleCustomWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReleaseBundleCustomWebhook)(nil)).Elem()
}

func (o ReleaseBundleCustomWebhookArrayOutput) ToReleaseBundleCustomWebhookArrayOutput() ReleaseBundleCustomWebhookArrayOutput {
	return o
}

func (o ReleaseBundleCustomWebhookArrayOutput) ToReleaseBundleCustomWebhookArrayOutputWithContext(ctx context.Context) ReleaseBundleCustomWebhookArrayOutput {
	return o
}

func (o ReleaseBundleCustomWebhookArrayOutput) Index(i pulumi.IntInput) ReleaseBundleCustomWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReleaseBundleCustomWebhook {
		return vs[0].([]*ReleaseBundleCustomWebhook)[vs[1].(int)]
	}).(ReleaseBundleCustomWebhookOutput)
}

type ReleaseBundleCustomWebhookMapOutput struct{ *pulumi.OutputState }

func (ReleaseBundleCustomWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReleaseBundleCustomWebhook)(nil)).Elem()
}

func (o ReleaseBundleCustomWebhookMapOutput) ToReleaseBundleCustomWebhookMapOutput() ReleaseBundleCustomWebhookMapOutput {
	return o
}

func (o ReleaseBundleCustomWebhookMapOutput) ToReleaseBundleCustomWebhookMapOutputWithContext(ctx context.Context) ReleaseBundleCustomWebhookMapOutput {
	return o
}

func (o ReleaseBundleCustomWebhookMapOutput) MapIndex(k pulumi.StringInput) ReleaseBundleCustomWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReleaseBundleCustomWebhook {
		return vs[0].(map[string]*ReleaseBundleCustomWebhook)[vs[1].(string)]
	}).(ReleaseBundleCustomWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseBundleCustomWebhookInput)(nil)).Elem(), &ReleaseBundleCustomWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseBundleCustomWebhookArrayInput)(nil)).Elem(), ReleaseBundleCustomWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseBundleCustomWebhookMapInput)(nil)).Elem(), ReleaseBundleCustomWebhookMap{})
	pulumi.RegisterOutputType(ReleaseBundleCustomWebhookOutput{})
	pulumi.RegisterOutputType(ReleaseBundleCustomWebhookArrayOutput{})
	pulumi.RegisterOutputType(ReleaseBundleCustomWebhookMapOutput{})
}
