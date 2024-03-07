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
// <!--Start PulumiCodeChooser -->
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
//			_, err := artifactory.NewBuildWebhook(ctx, "build-webhook", &artifactory.BuildWebhookArgs{
//				Criteria: &artifactory.BuildWebhookCriteriaArgs{
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
//				Handlers: artifactory.BuildWebhookHandlerArray{
//					&artifactory.BuildWebhookHandlerArgs{
//						CustomHttpHeaders: pulumi.StringMap{
//							"header-1": pulumi.String("value-1"),
//							"header-2": pulumi.String("value-2"),
//						},
//						Proxy:  pulumi.String("proxy-key"),
//						Secret: pulumi.String("some-secret"),
//						Url:    pulumi.String("http://tempurl.org/webhook"),
//					},
//				},
//				Key: pulumi.String("build-webhook"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
type BuildWebhook struct {
	pulumi.CustomResourceState

	// Specifies where the webhook will be applied on which repositories.
	Criteria BuildWebhookCriteriaOutput `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `uploaded`, `deleted`, `promoted`.
	EventTypes pulumi.StringArrayOutput `pulumi:"eventTypes"`
	// At least one is required.
	Handlers BuildWebhookHandlerArrayOutput `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringOutput `pulumi:"key"`
}

// NewBuildWebhook registers a new resource with the given unique name, arguments, and options.
func NewBuildWebhook(ctx *pulumi.Context,
	name string, args *BuildWebhookArgs, opts ...pulumi.ResourceOption) (*BuildWebhook, error) {
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
	var resource BuildWebhook
	err := ctx.RegisterResource("artifactory:index/buildWebhook:BuildWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBuildWebhook gets an existing BuildWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBuildWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BuildWebhookState, opts ...pulumi.ResourceOption) (*BuildWebhook, error) {
	var resource BuildWebhook
	err := ctx.ReadResource("artifactory:index/buildWebhook:BuildWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BuildWebhook resources.
type buildWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria *BuildWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `uploaded`, `deleted`, `promoted`.
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []BuildWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key *string `pulumi:"key"`
}

type BuildWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria BuildWebhookCriteriaPtrInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `uploaded`, `deleted`, `promoted`.
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers BuildWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringPtrInput
}

func (BuildWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*buildWebhookState)(nil)).Elem()
}

type buildWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria BuildWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `uploaded`, `deleted`, `promoted`.
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []BuildWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key string `pulumi:"key"`
}

// The set of arguments for constructing a BuildWebhook resource.
type BuildWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria BuildWebhookCriteriaInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `uploaded`, `deleted`, `promoted`.
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers BuildWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringInput
}

func (BuildWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*buildWebhookArgs)(nil)).Elem()
}

type BuildWebhookInput interface {
	pulumi.Input

	ToBuildWebhookOutput() BuildWebhookOutput
	ToBuildWebhookOutputWithContext(ctx context.Context) BuildWebhookOutput
}

func (*BuildWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildWebhook)(nil)).Elem()
}

func (i *BuildWebhook) ToBuildWebhookOutput() BuildWebhookOutput {
	return i.ToBuildWebhookOutputWithContext(context.Background())
}

func (i *BuildWebhook) ToBuildWebhookOutputWithContext(ctx context.Context) BuildWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildWebhookOutput)
}

// BuildWebhookArrayInput is an input type that accepts BuildWebhookArray and BuildWebhookArrayOutput values.
// You can construct a concrete instance of `BuildWebhookArrayInput` via:
//
//	BuildWebhookArray{ BuildWebhookArgs{...} }
type BuildWebhookArrayInput interface {
	pulumi.Input

	ToBuildWebhookArrayOutput() BuildWebhookArrayOutput
	ToBuildWebhookArrayOutputWithContext(context.Context) BuildWebhookArrayOutput
}

type BuildWebhookArray []BuildWebhookInput

func (BuildWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BuildWebhook)(nil)).Elem()
}

func (i BuildWebhookArray) ToBuildWebhookArrayOutput() BuildWebhookArrayOutput {
	return i.ToBuildWebhookArrayOutputWithContext(context.Background())
}

func (i BuildWebhookArray) ToBuildWebhookArrayOutputWithContext(ctx context.Context) BuildWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildWebhookArrayOutput)
}

// BuildWebhookMapInput is an input type that accepts BuildWebhookMap and BuildWebhookMapOutput values.
// You can construct a concrete instance of `BuildWebhookMapInput` via:
//
//	BuildWebhookMap{ "key": BuildWebhookArgs{...} }
type BuildWebhookMapInput interface {
	pulumi.Input

	ToBuildWebhookMapOutput() BuildWebhookMapOutput
	ToBuildWebhookMapOutputWithContext(context.Context) BuildWebhookMapOutput
}

type BuildWebhookMap map[string]BuildWebhookInput

func (BuildWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BuildWebhook)(nil)).Elem()
}

func (i BuildWebhookMap) ToBuildWebhookMapOutput() BuildWebhookMapOutput {
	return i.ToBuildWebhookMapOutputWithContext(context.Background())
}

func (i BuildWebhookMap) ToBuildWebhookMapOutputWithContext(ctx context.Context) BuildWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildWebhookMapOutput)
}

type BuildWebhookOutput struct{ *pulumi.OutputState }

func (BuildWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildWebhook)(nil)).Elem()
}

func (o BuildWebhookOutput) ToBuildWebhookOutput() BuildWebhookOutput {
	return o
}

func (o BuildWebhookOutput) ToBuildWebhookOutputWithContext(ctx context.Context) BuildWebhookOutput {
	return o
}

// Specifies where the webhook will be applied on which repositories.
func (o BuildWebhookOutput) Criteria() BuildWebhookCriteriaOutput {
	return o.ApplyT(func(v *BuildWebhook) BuildWebhookCriteriaOutput { return v.Criteria }).(BuildWebhookCriteriaOutput)
}

// Webhook description. Max length 1000 characters.
func (o BuildWebhookOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BuildWebhook) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Status of webhook. Default to `true`.
func (o BuildWebhookOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BuildWebhook) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `uploaded`, `deleted`, `promoted`.
func (o BuildWebhookOutput) EventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BuildWebhook) pulumi.StringArrayOutput { return v.EventTypes }).(pulumi.StringArrayOutput)
}

// At least one is required.
func (o BuildWebhookOutput) Handlers() BuildWebhookHandlerArrayOutput {
	return o.ApplyT(func(v *BuildWebhook) BuildWebhookHandlerArrayOutput { return v.Handlers }).(BuildWebhookHandlerArrayOutput)
}

// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
func (o BuildWebhookOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildWebhook) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

type BuildWebhookArrayOutput struct{ *pulumi.OutputState }

func (BuildWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BuildWebhook)(nil)).Elem()
}

func (o BuildWebhookArrayOutput) ToBuildWebhookArrayOutput() BuildWebhookArrayOutput {
	return o
}

func (o BuildWebhookArrayOutput) ToBuildWebhookArrayOutputWithContext(ctx context.Context) BuildWebhookArrayOutput {
	return o
}

func (o BuildWebhookArrayOutput) Index(i pulumi.IntInput) BuildWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BuildWebhook {
		return vs[0].([]*BuildWebhook)[vs[1].(int)]
	}).(BuildWebhookOutput)
}

type BuildWebhookMapOutput struct{ *pulumi.OutputState }

func (BuildWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BuildWebhook)(nil)).Elem()
}

func (o BuildWebhookMapOutput) ToBuildWebhookMapOutput() BuildWebhookMapOutput {
	return o
}

func (o BuildWebhookMapOutput) ToBuildWebhookMapOutputWithContext(ctx context.Context) BuildWebhookMapOutput {
	return o
}

func (o BuildWebhookMapOutput) MapIndex(k pulumi.StringInput) BuildWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BuildWebhook {
		return vs[0].(map[string]*BuildWebhook)[vs[1].(string)]
	}).(BuildWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BuildWebhookInput)(nil)).Elem(), &BuildWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*BuildWebhookArrayInput)(nil)).Elem(), BuildWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BuildWebhookMapInput)(nil)).Elem(), BuildWebhookMap{})
	pulumi.RegisterOutputType(BuildWebhookOutput{})
	pulumi.RegisterOutputType(BuildWebhookArrayOutput{})
	pulumi.RegisterOutputType(BuildWebhookMapOutput{})
}
