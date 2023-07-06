// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
//			_, err := artifactory.NewBuildCustomWebhook(ctx, "build-custom-webhook", &artifactory.BuildCustomWebhookArgs{
//				Criteria: &artifactory.BuildCustomWebhookCriteriaArgs{
//					AnyBuild: pulumi.Bool(true),
//					ExcludePatterns: pulumi.StringArray{
//						pulumi.String("bar/**"),
//					},
//					IncludePatterns: pulumi.StringArray{
//						pulumi.String("foo/**"),
//					},
//					SelectedBuilds: pulumi.StringArray{
//						pulumi.String("build-id"),
//					},
//				},
//				EventTypes: pulumi.StringArray{
//					pulumi.String("uploaded"),
//					pulumi.String("deleted"),
//					pulumi.String("promoted"),
//				},
//				Handlers: artifactory.BuildCustomWebhookHandlerArray{
//					&artifactory.BuildCustomWebhookHandlerArgs{
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
//				Key: pulumi.String("build-custom-webhook"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type BuildCustomWebhook struct {
	pulumi.CustomResourceState

	// Specifies where the webhook will be applied on which repositories.
	Criteria BuildCustomWebhookCriteriaOutput `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `uploaded`, `deleted`, `promoted`.
	EventTypes pulumi.StringArrayOutput `pulumi:"eventTypes"`
	// At least one is required.
	Handlers BuildCustomWebhookHandlerArrayOutput `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringOutput `pulumi:"key"`
}

// NewBuildCustomWebhook registers a new resource with the given unique name, arguments, and options.
func NewBuildCustomWebhook(ctx *pulumi.Context,
	name string, args *BuildCustomWebhookArgs, opts ...pulumi.ResourceOption) (*BuildCustomWebhook, error) {
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
	var resource BuildCustomWebhook
	err := ctx.RegisterResource("artifactory:index/buildCustomWebhook:BuildCustomWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBuildCustomWebhook gets an existing BuildCustomWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBuildCustomWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BuildCustomWebhookState, opts ...pulumi.ResourceOption) (*BuildCustomWebhook, error) {
	var resource BuildCustomWebhook
	err := ctx.ReadResource("artifactory:index/buildCustomWebhook:BuildCustomWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BuildCustomWebhook resources.
type buildCustomWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria *BuildCustomWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `uploaded`, `deleted`, `promoted`.
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []BuildCustomWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key *string `pulumi:"key"`
}

type BuildCustomWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria BuildCustomWebhookCriteriaPtrInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `uploaded`, `deleted`, `promoted`.
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers BuildCustomWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringPtrInput
}

func (BuildCustomWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*buildCustomWebhookState)(nil)).Elem()
}

type buildCustomWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria BuildCustomWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `uploaded`, `deleted`, `promoted`.
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []BuildCustomWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key string `pulumi:"key"`
}

// The set of arguments for constructing a BuildCustomWebhook resource.
type BuildCustomWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria BuildCustomWebhookCriteriaInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `uploaded`, `deleted`, `promoted`.
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers BuildCustomWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringInput
}

func (BuildCustomWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*buildCustomWebhookArgs)(nil)).Elem()
}

type BuildCustomWebhookInput interface {
	pulumi.Input

	ToBuildCustomWebhookOutput() BuildCustomWebhookOutput
	ToBuildCustomWebhookOutputWithContext(ctx context.Context) BuildCustomWebhookOutput
}

func (*BuildCustomWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildCustomWebhook)(nil)).Elem()
}

func (i *BuildCustomWebhook) ToBuildCustomWebhookOutput() BuildCustomWebhookOutput {
	return i.ToBuildCustomWebhookOutputWithContext(context.Background())
}

func (i *BuildCustomWebhook) ToBuildCustomWebhookOutputWithContext(ctx context.Context) BuildCustomWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildCustomWebhookOutput)
}

// BuildCustomWebhookArrayInput is an input type that accepts BuildCustomWebhookArray and BuildCustomWebhookArrayOutput values.
// You can construct a concrete instance of `BuildCustomWebhookArrayInput` via:
//
//	BuildCustomWebhookArray{ BuildCustomWebhookArgs{...} }
type BuildCustomWebhookArrayInput interface {
	pulumi.Input

	ToBuildCustomWebhookArrayOutput() BuildCustomWebhookArrayOutput
	ToBuildCustomWebhookArrayOutputWithContext(context.Context) BuildCustomWebhookArrayOutput
}

type BuildCustomWebhookArray []BuildCustomWebhookInput

func (BuildCustomWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BuildCustomWebhook)(nil)).Elem()
}

func (i BuildCustomWebhookArray) ToBuildCustomWebhookArrayOutput() BuildCustomWebhookArrayOutput {
	return i.ToBuildCustomWebhookArrayOutputWithContext(context.Background())
}

func (i BuildCustomWebhookArray) ToBuildCustomWebhookArrayOutputWithContext(ctx context.Context) BuildCustomWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildCustomWebhookArrayOutput)
}

// BuildCustomWebhookMapInput is an input type that accepts BuildCustomWebhookMap and BuildCustomWebhookMapOutput values.
// You can construct a concrete instance of `BuildCustomWebhookMapInput` via:
//
//	BuildCustomWebhookMap{ "key": BuildCustomWebhookArgs{...} }
type BuildCustomWebhookMapInput interface {
	pulumi.Input

	ToBuildCustomWebhookMapOutput() BuildCustomWebhookMapOutput
	ToBuildCustomWebhookMapOutputWithContext(context.Context) BuildCustomWebhookMapOutput
}

type BuildCustomWebhookMap map[string]BuildCustomWebhookInput

func (BuildCustomWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BuildCustomWebhook)(nil)).Elem()
}

func (i BuildCustomWebhookMap) ToBuildCustomWebhookMapOutput() BuildCustomWebhookMapOutput {
	return i.ToBuildCustomWebhookMapOutputWithContext(context.Background())
}

func (i BuildCustomWebhookMap) ToBuildCustomWebhookMapOutputWithContext(ctx context.Context) BuildCustomWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildCustomWebhookMapOutput)
}

type BuildCustomWebhookOutput struct{ *pulumi.OutputState }

func (BuildCustomWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildCustomWebhook)(nil)).Elem()
}

func (o BuildCustomWebhookOutput) ToBuildCustomWebhookOutput() BuildCustomWebhookOutput {
	return o
}

func (o BuildCustomWebhookOutput) ToBuildCustomWebhookOutputWithContext(ctx context.Context) BuildCustomWebhookOutput {
	return o
}

// Specifies where the webhook will be applied on which repositories.
func (o BuildCustomWebhookOutput) Criteria() BuildCustomWebhookCriteriaOutput {
	return o.ApplyT(func(v *BuildCustomWebhook) BuildCustomWebhookCriteriaOutput { return v.Criteria }).(BuildCustomWebhookCriteriaOutput)
}

// Webhook description. Max length 1000 characters.
func (o BuildCustomWebhookOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BuildCustomWebhook) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Status of webhook. Default to `true`.
func (o BuildCustomWebhookOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BuildCustomWebhook) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `uploaded`, `deleted`, `promoted`.
func (o BuildCustomWebhookOutput) EventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BuildCustomWebhook) pulumi.StringArrayOutput { return v.EventTypes }).(pulumi.StringArrayOutput)
}

// At least one is required.
func (o BuildCustomWebhookOutput) Handlers() BuildCustomWebhookHandlerArrayOutput {
	return o.ApplyT(func(v *BuildCustomWebhook) BuildCustomWebhookHandlerArrayOutput { return v.Handlers }).(BuildCustomWebhookHandlerArrayOutput)
}

// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
func (o BuildCustomWebhookOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildCustomWebhook) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

type BuildCustomWebhookArrayOutput struct{ *pulumi.OutputState }

func (BuildCustomWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BuildCustomWebhook)(nil)).Elem()
}

func (o BuildCustomWebhookArrayOutput) ToBuildCustomWebhookArrayOutput() BuildCustomWebhookArrayOutput {
	return o
}

func (o BuildCustomWebhookArrayOutput) ToBuildCustomWebhookArrayOutputWithContext(ctx context.Context) BuildCustomWebhookArrayOutput {
	return o
}

func (o BuildCustomWebhookArrayOutput) Index(i pulumi.IntInput) BuildCustomWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BuildCustomWebhook {
		return vs[0].([]*BuildCustomWebhook)[vs[1].(int)]
	}).(BuildCustomWebhookOutput)
}

type BuildCustomWebhookMapOutput struct{ *pulumi.OutputState }

func (BuildCustomWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BuildCustomWebhook)(nil)).Elem()
}

func (o BuildCustomWebhookMapOutput) ToBuildCustomWebhookMapOutput() BuildCustomWebhookMapOutput {
	return o
}

func (o BuildCustomWebhookMapOutput) ToBuildCustomWebhookMapOutputWithContext(ctx context.Context) BuildCustomWebhookMapOutput {
	return o
}

func (o BuildCustomWebhookMapOutput) MapIndex(k pulumi.StringInput) BuildCustomWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BuildCustomWebhook {
		return vs[0].(map[string]*BuildCustomWebhook)[vs[1].(string)]
	}).(BuildCustomWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BuildCustomWebhookInput)(nil)).Elem(), &BuildCustomWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*BuildCustomWebhookArrayInput)(nil)).Elem(), BuildCustomWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BuildCustomWebhookMapInput)(nil)).Elem(), BuildCustomWebhookMap{})
	pulumi.RegisterOutputType(BuildCustomWebhookOutput{})
	pulumi.RegisterOutputType(BuildCustomWebhookArrayOutput{})
	pulumi.RegisterOutputType(BuildCustomWebhookMapOutput{})
}
