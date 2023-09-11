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
//			_, err := artifactory.NewDockerV2Repository(ctx, "my-docker-local", &artifactory.DockerV2RepositoryArgs{
//				Key: pulumi.String("my-docker-local"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = artifactory.NewDockerCustomWebhook(ctx, "docker-custom-webhook", &artifactory.DockerCustomWebhookArgs{
//				Key: pulumi.String("docker-custom_webhook"),
//				EventTypes: pulumi.StringArray{
//					pulumi.String("pushed"),
//					pulumi.String("deleted"),
//					pulumi.String("promoted"),
//				},
//				Criteria: &artifactory.DockerCustomWebhookCriteriaArgs{
//					AnyLocal:  pulumi.Bool(true),
//					AnyRemote: pulumi.Bool(false),
//					RepoKeys: pulumi.StringArray{
//						my_docker_local.Key,
//					},
//					IncludePatterns: pulumi.StringArray{
//						pulumi.String("foo/**"),
//					},
//					ExcludePatterns: pulumi.StringArray{
//						pulumi.String("bar/**"),
//					},
//				},
//				Handlers: artifactory.DockerCustomWebhookHandlerArray{
//					&artifactory.DockerCustomWebhookHandlerArgs{
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
//				my_docker_local,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type DockerCustomWebhook struct {
	pulumi.CustomResourceState

	// Specifies where the webhook will be applied on which repositories.
	Criteria DockerCustomWebhookCriteriaOutput `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `pushed`, `deleted`, `promoted`.
	EventTypes pulumi.StringArrayOutput `pulumi:"eventTypes"`
	// At least one is required.
	Handlers DockerCustomWebhookHandlerArrayOutput `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringOutput `pulumi:"key"`
}

// NewDockerCustomWebhook registers a new resource with the given unique name, arguments, and options.
func NewDockerCustomWebhook(ctx *pulumi.Context,
	name string, args *DockerCustomWebhookArgs, opts ...pulumi.ResourceOption) (*DockerCustomWebhook, error) {
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
	var resource DockerCustomWebhook
	err := ctx.RegisterResource("artifactory:index/dockerCustomWebhook:DockerCustomWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDockerCustomWebhook gets an existing DockerCustomWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDockerCustomWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DockerCustomWebhookState, opts ...pulumi.ResourceOption) (*DockerCustomWebhook, error) {
	var resource DockerCustomWebhook
	err := ctx.ReadResource("artifactory:index/dockerCustomWebhook:DockerCustomWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DockerCustomWebhook resources.
type dockerCustomWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria *DockerCustomWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `pushed`, `deleted`, `promoted`.
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []DockerCustomWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key *string `pulumi:"key"`
}

type DockerCustomWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria DockerCustomWebhookCriteriaPtrInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `pushed`, `deleted`, `promoted`.
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers DockerCustomWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringPtrInput
}

func (DockerCustomWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*dockerCustomWebhookState)(nil)).Elem()
}

type dockerCustomWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria DockerCustomWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `pushed`, `deleted`, `promoted`.
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []DockerCustomWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key string `pulumi:"key"`
}

// The set of arguments for constructing a DockerCustomWebhook resource.
type DockerCustomWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria DockerCustomWebhookCriteriaInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `pushed`, `deleted`, `promoted`.
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers DockerCustomWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringInput
}

func (DockerCustomWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dockerCustomWebhookArgs)(nil)).Elem()
}

type DockerCustomWebhookInput interface {
	pulumi.Input

	ToDockerCustomWebhookOutput() DockerCustomWebhookOutput
	ToDockerCustomWebhookOutputWithContext(ctx context.Context) DockerCustomWebhookOutput
}

func (*DockerCustomWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**DockerCustomWebhook)(nil)).Elem()
}

func (i *DockerCustomWebhook) ToDockerCustomWebhookOutput() DockerCustomWebhookOutput {
	return i.ToDockerCustomWebhookOutputWithContext(context.Background())
}

func (i *DockerCustomWebhook) ToDockerCustomWebhookOutputWithContext(ctx context.Context) DockerCustomWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerCustomWebhookOutput)
}

func (i *DockerCustomWebhook) ToOutput(ctx context.Context) pulumix.Output[*DockerCustomWebhook] {
	return pulumix.Output[*DockerCustomWebhook]{
		OutputState: i.ToDockerCustomWebhookOutputWithContext(ctx).OutputState,
	}
}

// DockerCustomWebhookArrayInput is an input type that accepts DockerCustomWebhookArray and DockerCustomWebhookArrayOutput values.
// You can construct a concrete instance of `DockerCustomWebhookArrayInput` via:
//
//	DockerCustomWebhookArray{ DockerCustomWebhookArgs{...} }
type DockerCustomWebhookArrayInput interface {
	pulumi.Input

	ToDockerCustomWebhookArrayOutput() DockerCustomWebhookArrayOutput
	ToDockerCustomWebhookArrayOutputWithContext(context.Context) DockerCustomWebhookArrayOutput
}

type DockerCustomWebhookArray []DockerCustomWebhookInput

func (DockerCustomWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DockerCustomWebhook)(nil)).Elem()
}

func (i DockerCustomWebhookArray) ToDockerCustomWebhookArrayOutput() DockerCustomWebhookArrayOutput {
	return i.ToDockerCustomWebhookArrayOutputWithContext(context.Background())
}

func (i DockerCustomWebhookArray) ToDockerCustomWebhookArrayOutputWithContext(ctx context.Context) DockerCustomWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerCustomWebhookArrayOutput)
}

func (i DockerCustomWebhookArray) ToOutput(ctx context.Context) pulumix.Output[[]*DockerCustomWebhook] {
	return pulumix.Output[[]*DockerCustomWebhook]{
		OutputState: i.ToDockerCustomWebhookArrayOutputWithContext(ctx).OutputState,
	}
}

// DockerCustomWebhookMapInput is an input type that accepts DockerCustomWebhookMap and DockerCustomWebhookMapOutput values.
// You can construct a concrete instance of `DockerCustomWebhookMapInput` via:
//
//	DockerCustomWebhookMap{ "key": DockerCustomWebhookArgs{...} }
type DockerCustomWebhookMapInput interface {
	pulumi.Input

	ToDockerCustomWebhookMapOutput() DockerCustomWebhookMapOutput
	ToDockerCustomWebhookMapOutputWithContext(context.Context) DockerCustomWebhookMapOutput
}

type DockerCustomWebhookMap map[string]DockerCustomWebhookInput

func (DockerCustomWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DockerCustomWebhook)(nil)).Elem()
}

func (i DockerCustomWebhookMap) ToDockerCustomWebhookMapOutput() DockerCustomWebhookMapOutput {
	return i.ToDockerCustomWebhookMapOutputWithContext(context.Background())
}

func (i DockerCustomWebhookMap) ToDockerCustomWebhookMapOutputWithContext(ctx context.Context) DockerCustomWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerCustomWebhookMapOutput)
}

func (i DockerCustomWebhookMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*DockerCustomWebhook] {
	return pulumix.Output[map[string]*DockerCustomWebhook]{
		OutputState: i.ToDockerCustomWebhookMapOutputWithContext(ctx).OutputState,
	}
}

type DockerCustomWebhookOutput struct{ *pulumi.OutputState }

func (DockerCustomWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DockerCustomWebhook)(nil)).Elem()
}

func (o DockerCustomWebhookOutput) ToDockerCustomWebhookOutput() DockerCustomWebhookOutput {
	return o
}

func (o DockerCustomWebhookOutput) ToDockerCustomWebhookOutputWithContext(ctx context.Context) DockerCustomWebhookOutput {
	return o
}

func (o DockerCustomWebhookOutput) ToOutput(ctx context.Context) pulumix.Output[*DockerCustomWebhook] {
	return pulumix.Output[*DockerCustomWebhook]{
		OutputState: o.OutputState,
	}
}

// Specifies where the webhook will be applied on which repositories.
func (o DockerCustomWebhookOutput) Criteria() DockerCustomWebhookCriteriaOutput {
	return o.ApplyT(func(v *DockerCustomWebhook) DockerCustomWebhookCriteriaOutput { return v.Criteria }).(DockerCustomWebhookCriteriaOutput)
}

// Webhook description. Max length 1000 characters.
func (o DockerCustomWebhookOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DockerCustomWebhook) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Status of webhook. Default to `true`.
func (o DockerCustomWebhookOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DockerCustomWebhook) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `pushed`, `deleted`, `promoted`.
func (o DockerCustomWebhookOutput) EventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DockerCustomWebhook) pulumi.StringArrayOutput { return v.EventTypes }).(pulumi.StringArrayOutput)
}

// At least one is required.
func (o DockerCustomWebhookOutput) Handlers() DockerCustomWebhookHandlerArrayOutput {
	return o.ApplyT(func(v *DockerCustomWebhook) DockerCustomWebhookHandlerArrayOutput { return v.Handlers }).(DockerCustomWebhookHandlerArrayOutput)
}

// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
func (o DockerCustomWebhookOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *DockerCustomWebhook) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

type DockerCustomWebhookArrayOutput struct{ *pulumi.OutputState }

func (DockerCustomWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DockerCustomWebhook)(nil)).Elem()
}

func (o DockerCustomWebhookArrayOutput) ToDockerCustomWebhookArrayOutput() DockerCustomWebhookArrayOutput {
	return o
}

func (o DockerCustomWebhookArrayOutput) ToDockerCustomWebhookArrayOutputWithContext(ctx context.Context) DockerCustomWebhookArrayOutput {
	return o
}

func (o DockerCustomWebhookArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*DockerCustomWebhook] {
	return pulumix.Output[[]*DockerCustomWebhook]{
		OutputState: o.OutputState,
	}
}

func (o DockerCustomWebhookArrayOutput) Index(i pulumi.IntInput) DockerCustomWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DockerCustomWebhook {
		return vs[0].([]*DockerCustomWebhook)[vs[1].(int)]
	}).(DockerCustomWebhookOutput)
}

type DockerCustomWebhookMapOutput struct{ *pulumi.OutputState }

func (DockerCustomWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DockerCustomWebhook)(nil)).Elem()
}

func (o DockerCustomWebhookMapOutput) ToDockerCustomWebhookMapOutput() DockerCustomWebhookMapOutput {
	return o
}

func (o DockerCustomWebhookMapOutput) ToDockerCustomWebhookMapOutputWithContext(ctx context.Context) DockerCustomWebhookMapOutput {
	return o
}

func (o DockerCustomWebhookMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*DockerCustomWebhook] {
	return pulumix.Output[map[string]*DockerCustomWebhook]{
		OutputState: o.OutputState,
	}
}

func (o DockerCustomWebhookMapOutput) MapIndex(k pulumi.StringInput) DockerCustomWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DockerCustomWebhook {
		return vs[0].(map[string]*DockerCustomWebhook)[vs[1].(string)]
	}).(DockerCustomWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DockerCustomWebhookInput)(nil)).Elem(), &DockerCustomWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*DockerCustomWebhookArrayInput)(nil)).Elem(), DockerCustomWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DockerCustomWebhookMapInput)(nil)).Elem(), DockerCustomWebhookMap{})
	pulumi.RegisterOutputType(DockerCustomWebhookOutput{})
	pulumi.RegisterOutputType(DockerCustomWebhookArrayOutput{})
	pulumi.RegisterOutputType(DockerCustomWebhookMapOutput{})
}
