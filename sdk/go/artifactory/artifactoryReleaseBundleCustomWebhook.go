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

// Provides an Artifactory custom webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.
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
//			_, err := artifactory.NewArtifactoryReleaseBundleCustomWebhook(ctx, "artifactory-release-bundle-custom-webhook", &artifactory.ArtifactoryReleaseBundleCustomWebhookArgs{
//				Key: pulumi.String("artifactory-release-bundle-custom-webhook"),
//				EventTypes: pulumi.StringArray{
//					pulumi.String("received"),
//					pulumi.String("delete_started"),
//					pulumi.String("delete_completed"),
//					pulumi.String("delete_failed"),
//				},
//				Criteria: &artifactory.ArtifactoryReleaseBundleCustomWebhookCriteriaArgs{
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
//				Handlers: artifactory.ArtifactoryReleaseBundleCustomWebhookHandlerArray{
//					&artifactory.ArtifactoryReleaseBundleCustomWebhookHandlerArgs{
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
type ArtifactoryReleaseBundleCustomWebhook struct {
	pulumi.CustomResourceState

	// Specifies where the webhook will be applied on which repositories.
	Criteria ArtifactoryReleaseBundleCustomWebhookCriteriaOutput `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Status of webhook. Default to `true`
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `received`, `deleteStarted`, `deleteCompleted`, `deleteFailed`
	EventTypes pulumi.StringArrayOutput `pulumi:"eventTypes"`
	// At least one is required.
	Handlers ArtifactoryReleaseBundleCustomWebhookHandlerArrayOutput `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringOutput `pulumi:"key"`
}

// NewArtifactoryReleaseBundleCustomWebhook registers a new resource with the given unique name, arguments, and options.
func NewArtifactoryReleaseBundleCustomWebhook(ctx *pulumi.Context,
	name string, args *ArtifactoryReleaseBundleCustomWebhookArgs, opts ...pulumi.ResourceOption) (*ArtifactoryReleaseBundleCustomWebhook, error) {
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
	var resource ArtifactoryReleaseBundleCustomWebhook
	err := ctx.RegisterResource("artifactory:index/artifactoryReleaseBundleCustomWebhook:ArtifactoryReleaseBundleCustomWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetArtifactoryReleaseBundleCustomWebhook gets an existing ArtifactoryReleaseBundleCustomWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetArtifactoryReleaseBundleCustomWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ArtifactoryReleaseBundleCustomWebhookState, opts ...pulumi.ResourceOption) (*ArtifactoryReleaseBundleCustomWebhook, error) {
	var resource ArtifactoryReleaseBundleCustomWebhook
	err := ctx.ReadResource("artifactory:index/artifactoryReleaseBundleCustomWebhook:ArtifactoryReleaseBundleCustomWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ArtifactoryReleaseBundleCustomWebhook resources.
type artifactoryReleaseBundleCustomWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria *ArtifactoryReleaseBundleCustomWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `received`, `deleteStarted`, `deleteCompleted`, `deleteFailed`
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []ArtifactoryReleaseBundleCustomWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key *string `pulumi:"key"`
}

type ArtifactoryReleaseBundleCustomWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria ArtifactoryReleaseBundleCustomWebhookCriteriaPtrInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `received`, `deleteStarted`, `deleteCompleted`, `deleteFailed`
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers ArtifactoryReleaseBundleCustomWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringPtrInput
}

func (ArtifactoryReleaseBundleCustomWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactoryReleaseBundleCustomWebhookState)(nil)).Elem()
}

type artifactoryReleaseBundleCustomWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria ArtifactoryReleaseBundleCustomWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `received`, `deleteStarted`, `deleteCompleted`, `deleteFailed`
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []ArtifactoryReleaseBundleCustomWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key string `pulumi:"key"`
}

// The set of arguments for constructing a ArtifactoryReleaseBundleCustomWebhook resource.
type ArtifactoryReleaseBundleCustomWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria ArtifactoryReleaseBundleCustomWebhookCriteriaInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `received`, `deleteStarted`, `deleteCompleted`, `deleteFailed`
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers ArtifactoryReleaseBundleCustomWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringInput
}

func (ArtifactoryReleaseBundleCustomWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactoryReleaseBundleCustomWebhookArgs)(nil)).Elem()
}

type ArtifactoryReleaseBundleCustomWebhookInput interface {
	pulumi.Input

	ToArtifactoryReleaseBundleCustomWebhookOutput() ArtifactoryReleaseBundleCustomWebhookOutput
	ToArtifactoryReleaseBundleCustomWebhookOutputWithContext(ctx context.Context) ArtifactoryReleaseBundleCustomWebhookOutput
}

func (*ArtifactoryReleaseBundleCustomWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**ArtifactoryReleaseBundleCustomWebhook)(nil)).Elem()
}

func (i *ArtifactoryReleaseBundleCustomWebhook) ToArtifactoryReleaseBundleCustomWebhookOutput() ArtifactoryReleaseBundleCustomWebhookOutput {
	return i.ToArtifactoryReleaseBundleCustomWebhookOutputWithContext(context.Background())
}

func (i *ArtifactoryReleaseBundleCustomWebhook) ToArtifactoryReleaseBundleCustomWebhookOutputWithContext(ctx context.Context) ArtifactoryReleaseBundleCustomWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactoryReleaseBundleCustomWebhookOutput)
}

// ArtifactoryReleaseBundleCustomWebhookArrayInput is an input type that accepts ArtifactoryReleaseBundleCustomWebhookArray and ArtifactoryReleaseBundleCustomWebhookArrayOutput values.
// You can construct a concrete instance of `ArtifactoryReleaseBundleCustomWebhookArrayInput` via:
//
//	ArtifactoryReleaseBundleCustomWebhookArray{ ArtifactoryReleaseBundleCustomWebhookArgs{...} }
type ArtifactoryReleaseBundleCustomWebhookArrayInput interface {
	pulumi.Input

	ToArtifactoryReleaseBundleCustomWebhookArrayOutput() ArtifactoryReleaseBundleCustomWebhookArrayOutput
	ToArtifactoryReleaseBundleCustomWebhookArrayOutputWithContext(context.Context) ArtifactoryReleaseBundleCustomWebhookArrayOutput
}

type ArtifactoryReleaseBundleCustomWebhookArray []ArtifactoryReleaseBundleCustomWebhookInput

func (ArtifactoryReleaseBundleCustomWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ArtifactoryReleaseBundleCustomWebhook)(nil)).Elem()
}

func (i ArtifactoryReleaseBundleCustomWebhookArray) ToArtifactoryReleaseBundleCustomWebhookArrayOutput() ArtifactoryReleaseBundleCustomWebhookArrayOutput {
	return i.ToArtifactoryReleaseBundleCustomWebhookArrayOutputWithContext(context.Background())
}

func (i ArtifactoryReleaseBundleCustomWebhookArray) ToArtifactoryReleaseBundleCustomWebhookArrayOutputWithContext(ctx context.Context) ArtifactoryReleaseBundleCustomWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactoryReleaseBundleCustomWebhookArrayOutput)
}

// ArtifactoryReleaseBundleCustomWebhookMapInput is an input type that accepts ArtifactoryReleaseBundleCustomWebhookMap and ArtifactoryReleaseBundleCustomWebhookMapOutput values.
// You can construct a concrete instance of `ArtifactoryReleaseBundleCustomWebhookMapInput` via:
//
//	ArtifactoryReleaseBundleCustomWebhookMap{ "key": ArtifactoryReleaseBundleCustomWebhookArgs{...} }
type ArtifactoryReleaseBundleCustomWebhookMapInput interface {
	pulumi.Input

	ToArtifactoryReleaseBundleCustomWebhookMapOutput() ArtifactoryReleaseBundleCustomWebhookMapOutput
	ToArtifactoryReleaseBundleCustomWebhookMapOutputWithContext(context.Context) ArtifactoryReleaseBundleCustomWebhookMapOutput
}

type ArtifactoryReleaseBundleCustomWebhookMap map[string]ArtifactoryReleaseBundleCustomWebhookInput

func (ArtifactoryReleaseBundleCustomWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ArtifactoryReleaseBundleCustomWebhook)(nil)).Elem()
}

func (i ArtifactoryReleaseBundleCustomWebhookMap) ToArtifactoryReleaseBundleCustomWebhookMapOutput() ArtifactoryReleaseBundleCustomWebhookMapOutput {
	return i.ToArtifactoryReleaseBundleCustomWebhookMapOutputWithContext(context.Background())
}

func (i ArtifactoryReleaseBundleCustomWebhookMap) ToArtifactoryReleaseBundleCustomWebhookMapOutputWithContext(ctx context.Context) ArtifactoryReleaseBundleCustomWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactoryReleaseBundleCustomWebhookMapOutput)
}

type ArtifactoryReleaseBundleCustomWebhookOutput struct{ *pulumi.OutputState }

func (ArtifactoryReleaseBundleCustomWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArtifactoryReleaseBundleCustomWebhook)(nil)).Elem()
}

func (o ArtifactoryReleaseBundleCustomWebhookOutput) ToArtifactoryReleaseBundleCustomWebhookOutput() ArtifactoryReleaseBundleCustomWebhookOutput {
	return o
}

func (o ArtifactoryReleaseBundleCustomWebhookOutput) ToArtifactoryReleaseBundleCustomWebhookOutputWithContext(ctx context.Context) ArtifactoryReleaseBundleCustomWebhookOutput {
	return o
}

// Specifies where the webhook will be applied on which repositories.
func (o ArtifactoryReleaseBundleCustomWebhookOutput) Criteria() ArtifactoryReleaseBundleCustomWebhookCriteriaOutput {
	return o.ApplyT(func(v *ArtifactoryReleaseBundleCustomWebhook) ArtifactoryReleaseBundleCustomWebhookCriteriaOutput {
		return v.Criteria
	}).(ArtifactoryReleaseBundleCustomWebhookCriteriaOutput)
}

// Webhook description. Max length 1000 characters.
func (o ArtifactoryReleaseBundleCustomWebhookOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactoryReleaseBundleCustomWebhook) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Status of webhook. Default to `true`
func (o ArtifactoryReleaseBundleCustomWebhookOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ArtifactoryReleaseBundleCustomWebhook) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `received`, `deleteStarted`, `deleteCompleted`, `deleteFailed`
func (o ArtifactoryReleaseBundleCustomWebhookOutput) EventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ArtifactoryReleaseBundleCustomWebhook) pulumi.StringArrayOutput { return v.EventTypes }).(pulumi.StringArrayOutput)
}

// At least one is required.
func (o ArtifactoryReleaseBundleCustomWebhookOutput) Handlers() ArtifactoryReleaseBundleCustomWebhookHandlerArrayOutput {
	return o.ApplyT(func(v *ArtifactoryReleaseBundleCustomWebhook) ArtifactoryReleaseBundleCustomWebhookHandlerArrayOutput {
		return v.Handlers
	}).(ArtifactoryReleaseBundleCustomWebhookHandlerArrayOutput)
}

// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
func (o ArtifactoryReleaseBundleCustomWebhookOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *ArtifactoryReleaseBundleCustomWebhook) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

type ArtifactoryReleaseBundleCustomWebhookArrayOutput struct{ *pulumi.OutputState }

func (ArtifactoryReleaseBundleCustomWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ArtifactoryReleaseBundleCustomWebhook)(nil)).Elem()
}

func (o ArtifactoryReleaseBundleCustomWebhookArrayOutput) ToArtifactoryReleaseBundleCustomWebhookArrayOutput() ArtifactoryReleaseBundleCustomWebhookArrayOutput {
	return o
}

func (o ArtifactoryReleaseBundleCustomWebhookArrayOutput) ToArtifactoryReleaseBundleCustomWebhookArrayOutputWithContext(ctx context.Context) ArtifactoryReleaseBundleCustomWebhookArrayOutput {
	return o
}

func (o ArtifactoryReleaseBundleCustomWebhookArrayOutput) Index(i pulumi.IntInput) ArtifactoryReleaseBundleCustomWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ArtifactoryReleaseBundleCustomWebhook {
		return vs[0].([]*ArtifactoryReleaseBundleCustomWebhook)[vs[1].(int)]
	}).(ArtifactoryReleaseBundleCustomWebhookOutput)
}

type ArtifactoryReleaseBundleCustomWebhookMapOutput struct{ *pulumi.OutputState }

func (ArtifactoryReleaseBundleCustomWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ArtifactoryReleaseBundleCustomWebhook)(nil)).Elem()
}

func (o ArtifactoryReleaseBundleCustomWebhookMapOutput) ToArtifactoryReleaseBundleCustomWebhookMapOutput() ArtifactoryReleaseBundleCustomWebhookMapOutput {
	return o
}

func (o ArtifactoryReleaseBundleCustomWebhookMapOutput) ToArtifactoryReleaseBundleCustomWebhookMapOutputWithContext(ctx context.Context) ArtifactoryReleaseBundleCustomWebhookMapOutput {
	return o
}

func (o ArtifactoryReleaseBundleCustomWebhookMapOutput) MapIndex(k pulumi.StringInput) ArtifactoryReleaseBundleCustomWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ArtifactoryReleaseBundleCustomWebhook {
		return vs[0].(map[string]*ArtifactoryReleaseBundleCustomWebhook)[vs[1].(string)]
	}).(ArtifactoryReleaseBundleCustomWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ArtifactoryReleaseBundleCustomWebhookInput)(nil)).Elem(), &ArtifactoryReleaseBundleCustomWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*ArtifactoryReleaseBundleCustomWebhookArrayInput)(nil)).Elem(), ArtifactoryReleaseBundleCustomWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ArtifactoryReleaseBundleCustomWebhookMapInput)(nil)).Elem(), ArtifactoryReleaseBundleCustomWebhookMap{})
	pulumi.RegisterOutputType(ArtifactoryReleaseBundleCustomWebhookOutput{})
	pulumi.RegisterOutputType(ArtifactoryReleaseBundleCustomWebhookArrayOutput{})
	pulumi.RegisterOutputType(ArtifactoryReleaseBundleCustomWebhookMapOutput{})
}
