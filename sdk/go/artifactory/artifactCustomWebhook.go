// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory/internal"
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
//			_, err = artifactory.NewArtifactCustomWebhook(ctx, "artifact-custom-webhook", &artifactory.ArtifactCustomWebhookArgs{
//				Key: pulumi.String("artifact-custom-webhook"),
//				EventTypes: pulumi.StringArray{
//					pulumi.String("deployed"),
//					pulumi.String("deleted"),
//					pulumi.String("moved"),
//					pulumi.String("copied"),
//				},
//				Criteria: &artifactory.ArtifactCustomWebhookCriteriaArgs{
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
//				Handlers: artifactory.ArtifactCustomWebhookHandlerArray{
//					&artifactory.ArtifactCustomWebhookHandlerArgs{
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
type ArtifactCustomWebhook struct {
	pulumi.CustomResourceState

	// Specifies where the webhook will be applied on which repositories.
	Criteria ArtifactCustomWebhookCriteriaOutput `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `deployed`, `deleted`, `moved`, `copied`, `cached`.
	EventTypes pulumi.StringArrayOutput `pulumi:"eventTypes"`
	// At least one is required.
	Handlers ArtifactCustomWebhookHandlerArrayOutput `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringOutput `pulumi:"key"`
}

// NewArtifactCustomWebhook registers a new resource with the given unique name, arguments, and options.
func NewArtifactCustomWebhook(ctx *pulumi.Context,
	name string, args *ArtifactCustomWebhookArgs, opts ...pulumi.ResourceOption) (*ArtifactCustomWebhook, error) {
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
	var resource ArtifactCustomWebhook
	err := ctx.RegisterResource("artifactory:index/artifactCustomWebhook:ArtifactCustomWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetArtifactCustomWebhook gets an existing ArtifactCustomWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetArtifactCustomWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ArtifactCustomWebhookState, opts ...pulumi.ResourceOption) (*ArtifactCustomWebhook, error) {
	var resource ArtifactCustomWebhook
	err := ctx.ReadResource("artifactory:index/artifactCustomWebhook:ArtifactCustomWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ArtifactCustomWebhook resources.
type artifactCustomWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria *ArtifactCustomWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `deployed`, `deleted`, `moved`, `copied`, `cached`.
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []ArtifactCustomWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key *string `pulumi:"key"`
}

type ArtifactCustomWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria ArtifactCustomWebhookCriteriaPtrInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `deployed`, `deleted`, `moved`, `copied`, `cached`.
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers ArtifactCustomWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringPtrInput
}

func (ArtifactCustomWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactCustomWebhookState)(nil)).Elem()
}

type artifactCustomWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria ArtifactCustomWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `deployed`, `deleted`, `moved`, `copied`, `cached`.
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []ArtifactCustomWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key string `pulumi:"key"`
}

// The set of arguments for constructing a ArtifactCustomWebhook resource.
type ArtifactCustomWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria ArtifactCustomWebhookCriteriaInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `deployed`, `deleted`, `moved`, `copied`, `cached`.
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers ArtifactCustomWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringInput
}

func (ArtifactCustomWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactCustomWebhookArgs)(nil)).Elem()
}

type ArtifactCustomWebhookInput interface {
	pulumi.Input

	ToArtifactCustomWebhookOutput() ArtifactCustomWebhookOutput
	ToArtifactCustomWebhookOutputWithContext(ctx context.Context) ArtifactCustomWebhookOutput
}

func (*ArtifactCustomWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**ArtifactCustomWebhook)(nil)).Elem()
}

func (i *ArtifactCustomWebhook) ToArtifactCustomWebhookOutput() ArtifactCustomWebhookOutput {
	return i.ToArtifactCustomWebhookOutputWithContext(context.Background())
}

func (i *ArtifactCustomWebhook) ToArtifactCustomWebhookOutputWithContext(ctx context.Context) ArtifactCustomWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactCustomWebhookOutput)
}

func (i *ArtifactCustomWebhook) ToOutput(ctx context.Context) pulumix.Output[*ArtifactCustomWebhook] {
	return pulumix.Output[*ArtifactCustomWebhook]{
		OutputState: i.ToArtifactCustomWebhookOutputWithContext(ctx).OutputState,
	}
}

// ArtifactCustomWebhookArrayInput is an input type that accepts ArtifactCustomWebhookArray and ArtifactCustomWebhookArrayOutput values.
// You can construct a concrete instance of `ArtifactCustomWebhookArrayInput` via:
//
//	ArtifactCustomWebhookArray{ ArtifactCustomWebhookArgs{...} }
type ArtifactCustomWebhookArrayInput interface {
	pulumi.Input

	ToArtifactCustomWebhookArrayOutput() ArtifactCustomWebhookArrayOutput
	ToArtifactCustomWebhookArrayOutputWithContext(context.Context) ArtifactCustomWebhookArrayOutput
}

type ArtifactCustomWebhookArray []ArtifactCustomWebhookInput

func (ArtifactCustomWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ArtifactCustomWebhook)(nil)).Elem()
}

func (i ArtifactCustomWebhookArray) ToArtifactCustomWebhookArrayOutput() ArtifactCustomWebhookArrayOutput {
	return i.ToArtifactCustomWebhookArrayOutputWithContext(context.Background())
}

func (i ArtifactCustomWebhookArray) ToArtifactCustomWebhookArrayOutputWithContext(ctx context.Context) ArtifactCustomWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactCustomWebhookArrayOutput)
}

func (i ArtifactCustomWebhookArray) ToOutput(ctx context.Context) pulumix.Output[[]*ArtifactCustomWebhook] {
	return pulumix.Output[[]*ArtifactCustomWebhook]{
		OutputState: i.ToArtifactCustomWebhookArrayOutputWithContext(ctx).OutputState,
	}
}

// ArtifactCustomWebhookMapInput is an input type that accepts ArtifactCustomWebhookMap and ArtifactCustomWebhookMapOutput values.
// You can construct a concrete instance of `ArtifactCustomWebhookMapInput` via:
//
//	ArtifactCustomWebhookMap{ "key": ArtifactCustomWebhookArgs{...} }
type ArtifactCustomWebhookMapInput interface {
	pulumi.Input

	ToArtifactCustomWebhookMapOutput() ArtifactCustomWebhookMapOutput
	ToArtifactCustomWebhookMapOutputWithContext(context.Context) ArtifactCustomWebhookMapOutput
}

type ArtifactCustomWebhookMap map[string]ArtifactCustomWebhookInput

func (ArtifactCustomWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ArtifactCustomWebhook)(nil)).Elem()
}

func (i ArtifactCustomWebhookMap) ToArtifactCustomWebhookMapOutput() ArtifactCustomWebhookMapOutput {
	return i.ToArtifactCustomWebhookMapOutputWithContext(context.Background())
}

func (i ArtifactCustomWebhookMap) ToArtifactCustomWebhookMapOutputWithContext(ctx context.Context) ArtifactCustomWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactCustomWebhookMapOutput)
}

func (i ArtifactCustomWebhookMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ArtifactCustomWebhook] {
	return pulumix.Output[map[string]*ArtifactCustomWebhook]{
		OutputState: i.ToArtifactCustomWebhookMapOutputWithContext(ctx).OutputState,
	}
}

type ArtifactCustomWebhookOutput struct{ *pulumi.OutputState }

func (ArtifactCustomWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArtifactCustomWebhook)(nil)).Elem()
}

func (o ArtifactCustomWebhookOutput) ToArtifactCustomWebhookOutput() ArtifactCustomWebhookOutput {
	return o
}

func (o ArtifactCustomWebhookOutput) ToArtifactCustomWebhookOutputWithContext(ctx context.Context) ArtifactCustomWebhookOutput {
	return o
}

func (o ArtifactCustomWebhookOutput) ToOutput(ctx context.Context) pulumix.Output[*ArtifactCustomWebhook] {
	return pulumix.Output[*ArtifactCustomWebhook]{
		OutputState: o.OutputState,
	}
}

// Specifies where the webhook will be applied on which repositories.
func (o ArtifactCustomWebhookOutput) Criteria() ArtifactCustomWebhookCriteriaOutput {
	return o.ApplyT(func(v *ArtifactCustomWebhook) ArtifactCustomWebhookCriteriaOutput { return v.Criteria }).(ArtifactCustomWebhookCriteriaOutput)
}

// Webhook description. Max length 1000 characters.
func (o ArtifactCustomWebhookOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactCustomWebhook) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Status of webhook. Default to `true`.
func (o ArtifactCustomWebhookOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ArtifactCustomWebhook) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `deployed`, `deleted`, `moved`, `copied`, `cached`.
func (o ArtifactCustomWebhookOutput) EventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ArtifactCustomWebhook) pulumi.StringArrayOutput { return v.EventTypes }).(pulumi.StringArrayOutput)
}

// At least one is required.
func (o ArtifactCustomWebhookOutput) Handlers() ArtifactCustomWebhookHandlerArrayOutput {
	return o.ApplyT(func(v *ArtifactCustomWebhook) ArtifactCustomWebhookHandlerArrayOutput { return v.Handlers }).(ArtifactCustomWebhookHandlerArrayOutput)
}

// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
func (o ArtifactCustomWebhookOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *ArtifactCustomWebhook) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

type ArtifactCustomWebhookArrayOutput struct{ *pulumi.OutputState }

func (ArtifactCustomWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ArtifactCustomWebhook)(nil)).Elem()
}

func (o ArtifactCustomWebhookArrayOutput) ToArtifactCustomWebhookArrayOutput() ArtifactCustomWebhookArrayOutput {
	return o
}

func (o ArtifactCustomWebhookArrayOutput) ToArtifactCustomWebhookArrayOutputWithContext(ctx context.Context) ArtifactCustomWebhookArrayOutput {
	return o
}

func (o ArtifactCustomWebhookArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ArtifactCustomWebhook] {
	return pulumix.Output[[]*ArtifactCustomWebhook]{
		OutputState: o.OutputState,
	}
}

func (o ArtifactCustomWebhookArrayOutput) Index(i pulumi.IntInput) ArtifactCustomWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ArtifactCustomWebhook {
		return vs[0].([]*ArtifactCustomWebhook)[vs[1].(int)]
	}).(ArtifactCustomWebhookOutput)
}

type ArtifactCustomWebhookMapOutput struct{ *pulumi.OutputState }

func (ArtifactCustomWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ArtifactCustomWebhook)(nil)).Elem()
}

func (o ArtifactCustomWebhookMapOutput) ToArtifactCustomWebhookMapOutput() ArtifactCustomWebhookMapOutput {
	return o
}

func (o ArtifactCustomWebhookMapOutput) ToArtifactCustomWebhookMapOutputWithContext(ctx context.Context) ArtifactCustomWebhookMapOutput {
	return o
}

func (o ArtifactCustomWebhookMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ArtifactCustomWebhook] {
	return pulumix.Output[map[string]*ArtifactCustomWebhook]{
		OutputState: o.OutputState,
	}
}

func (o ArtifactCustomWebhookMapOutput) MapIndex(k pulumi.StringInput) ArtifactCustomWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ArtifactCustomWebhook {
		return vs[0].(map[string]*ArtifactCustomWebhook)[vs[1].(string)]
	}).(ArtifactCustomWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ArtifactCustomWebhookInput)(nil)).Elem(), &ArtifactCustomWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*ArtifactCustomWebhookArrayInput)(nil)).Elem(), ArtifactCustomWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ArtifactCustomWebhookMapInput)(nil)).Elem(), ArtifactCustomWebhookMap{})
	pulumi.RegisterOutputType(ArtifactCustomWebhookOutput{})
	pulumi.RegisterOutputType(ArtifactCustomWebhookArrayOutput{})
	pulumi.RegisterOutputType(ArtifactCustomWebhookMapOutput{})
}
