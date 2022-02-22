// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Artifact Webhook Resource
//
// Provides an Artifactory webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.
//
// ## Example Usage
//
// .
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-artifactory/sdk/go/artifactory"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := artifactory.NewLocalGenericRepository(ctx, "my-generic-local", &artifactory.LocalGenericRepositoryArgs{
// 			Key: pulumi.String("my-generic-local"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = artifactory.NewArtifactWebhook(ctx, "artifact-webhook", &artifactory.ArtifactWebhookArgs{
// 			Key: pulumi.String("artifact-webhook"),
// 			EventTypes: pulumi.StringArray{
// 				pulumi.String("deployed"),
// 				pulumi.String("deleted"),
// 				pulumi.String("moved"),
// 				pulumi.String("copied"),
// 			},
// 			Criteria: &ArtifactWebhookCriteriaArgs{
// 				AnyLocal:  pulumi.Bool(true),
// 				AnyRemote: pulumi.Bool(false),
// 				RepoKeys: pulumi.StringArray{
// 					my_generic_local.Key,
// 				},
// 				IncludePatterns: pulumi.StringArray{
// 					pulumi.String("foo/**"),
// 				},
// 				ExcludePatterns: pulumi.StringArray{
// 					pulumi.String("bar/**"),
// 				},
// 			},
// 			Url:    pulumi.String("http://tempurl.org/webhook"),
// 			Secret: pulumi.String("some-secret"),
// 			Proxy:  pulumi.String("proxy-key"),
// 			CustomHttpHeaders: pulumi.StringMap{
// 				"header-1": pulumi.String("value-1"),
// 				"header-2": pulumi.String("value-2"),
// 			},
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			my_generic_local,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ArtifactWebhook struct {
	pulumi.CustomResourceState

	// Specifies where the webhook will be applied on which repositories.
	Criteria ArtifactWebhookCriteriaOutput `pulumi:"criteria"`
	// Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
	CustomHttpHeaders pulumi.StringMapOutput `pulumi:"customHttpHeaders"`
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Status of webhook. Default to 'true'
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: "deployed", "deleted", "moved", "copied"
	EventTypes pulumi.StringArrayOutput `pulumi:"eventTypes"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringOutput `pulumi:"key"`
	// Proxy key from Artifactory Proxies setting
	Proxy pulumi.StringPtrOutput `pulumi:"proxy"`
	// Secret authentication token that will be sent to the configured URL
	Secret pulumi.StringPtrOutput `pulumi:"secret"`
	// Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewArtifactWebhook registers a new resource with the given unique name, arguments, and options.
func NewArtifactWebhook(ctx *pulumi.Context,
	name string, args *ArtifactWebhookArgs, opts ...pulumi.ResourceOption) (*ArtifactWebhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Criteria == nil {
		return nil, errors.New("invalid value for required argument 'Criteria'")
	}
	if args.EventTypes == nil {
		return nil, errors.New("invalid value for required argument 'EventTypes'")
	}
	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	var resource ArtifactWebhook
	err := ctx.RegisterResource("artifactory:index/artifactWebhook:ArtifactWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetArtifactWebhook gets an existing ArtifactWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetArtifactWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ArtifactWebhookState, opts ...pulumi.ResourceOption) (*ArtifactWebhook, error) {
	var resource ArtifactWebhook
	err := ctx.ReadResource("artifactory:index/artifactWebhook:ArtifactWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ArtifactWebhook resources.
type artifactWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria *ArtifactWebhookCriteria `pulumi:"criteria"`
	// Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
	CustomHttpHeaders map[string]string `pulumi:"customHttpHeaders"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to 'true'
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: "deployed", "deleted", "moved", "copied"
	EventTypes []string `pulumi:"eventTypes"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key *string `pulumi:"key"`
	// Proxy key from Artifactory Proxies setting
	Proxy *string `pulumi:"proxy"`
	// Secret authentication token that will be sent to the configured URL
	Secret *string `pulumi:"secret"`
	// Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
	Url *string `pulumi:"url"`
}

type ArtifactWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria ArtifactWebhookCriteriaPtrInput
	// Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
	CustomHttpHeaders pulumi.StringMapInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to 'true'
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: "deployed", "deleted", "moved", "copied"
	EventTypes pulumi.StringArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringPtrInput
	// Proxy key from Artifactory Proxies setting
	Proxy pulumi.StringPtrInput
	// Secret authentication token that will be sent to the configured URL
	Secret pulumi.StringPtrInput
	// Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
	Url pulumi.StringPtrInput
}

func (ArtifactWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactWebhookState)(nil)).Elem()
}

type artifactWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria ArtifactWebhookCriteria `pulumi:"criteria"`
	// Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
	CustomHttpHeaders map[string]string `pulumi:"customHttpHeaders"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to 'true'
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: "deployed", "deleted", "moved", "copied"
	EventTypes []string `pulumi:"eventTypes"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key string `pulumi:"key"`
	// Proxy key from Artifactory Proxies setting
	Proxy *string `pulumi:"proxy"`
	// Secret authentication token that will be sent to the configured URL
	Secret *string `pulumi:"secret"`
	// Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a ArtifactWebhook resource.
type ArtifactWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria ArtifactWebhookCriteriaInput
	// Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
	CustomHttpHeaders pulumi.StringMapInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to 'true'
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: "deployed", "deleted", "moved", "copied"
	EventTypes pulumi.StringArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringInput
	// Proxy key from Artifactory Proxies setting
	Proxy pulumi.StringPtrInput
	// Secret authentication token that will be sent to the configured URL
	Secret pulumi.StringPtrInput
	// Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
	Url pulumi.StringInput
}

func (ArtifactWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactWebhookArgs)(nil)).Elem()
}

type ArtifactWebhookInput interface {
	pulumi.Input

	ToArtifactWebhookOutput() ArtifactWebhookOutput
	ToArtifactWebhookOutputWithContext(ctx context.Context) ArtifactWebhookOutput
}

func (*ArtifactWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**ArtifactWebhook)(nil)).Elem()
}

func (i *ArtifactWebhook) ToArtifactWebhookOutput() ArtifactWebhookOutput {
	return i.ToArtifactWebhookOutputWithContext(context.Background())
}

func (i *ArtifactWebhook) ToArtifactWebhookOutputWithContext(ctx context.Context) ArtifactWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactWebhookOutput)
}

// ArtifactWebhookArrayInput is an input type that accepts ArtifactWebhookArray and ArtifactWebhookArrayOutput values.
// You can construct a concrete instance of `ArtifactWebhookArrayInput` via:
//
//          ArtifactWebhookArray{ ArtifactWebhookArgs{...} }
type ArtifactWebhookArrayInput interface {
	pulumi.Input

	ToArtifactWebhookArrayOutput() ArtifactWebhookArrayOutput
	ToArtifactWebhookArrayOutputWithContext(context.Context) ArtifactWebhookArrayOutput
}

type ArtifactWebhookArray []ArtifactWebhookInput

func (ArtifactWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ArtifactWebhook)(nil)).Elem()
}

func (i ArtifactWebhookArray) ToArtifactWebhookArrayOutput() ArtifactWebhookArrayOutput {
	return i.ToArtifactWebhookArrayOutputWithContext(context.Background())
}

func (i ArtifactWebhookArray) ToArtifactWebhookArrayOutputWithContext(ctx context.Context) ArtifactWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactWebhookArrayOutput)
}

// ArtifactWebhookMapInput is an input type that accepts ArtifactWebhookMap and ArtifactWebhookMapOutput values.
// You can construct a concrete instance of `ArtifactWebhookMapInput` via:
//
//          ArtifactWebhookMap{ "key": ArtifactWebhookArgs{...} }
type ArtifactWebhookMapInput interface {
	pulumi.Input

	ToArtifactWebhookMapOutput() ArtifactWebhookMapOutput
	ToArtifactWebhookMapOutputWithContext(context.Context) ArtifactWebhookMapOutput
}

type ArtifactWebhookMap map[string]ArtifactWebhookInput

func (ArtifactWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ArtifactWebhook)(nil)).Elem()
}

func (i ArtifactWebhookMap) ToArtifactWebhookMapOutput() ArtifactWebhookMapOutput {
	return i.ToArtifactWebhookMapOutputWithContext(context.Background())
}

func (i ArtifactWebhookMap) ToArtifactWebhookMapOutputWithContext(ctx context.Context) ArtifactWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactWebhookMapOutput)
}

type ArtifactWebhookOutput struct{ *pulumi.OutputState }

func (ArtifactWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArtifactWebhook)(nil)).Elem()
}

func (o ArtifactWebhookOutput) ToArtifactWebhookOutput() ArtifactWebhookOutput {
	return o
}

func (o ArtifactWebhookOutput) ToArtifactWebhookOutputWithContext(ctx context.Context) ArtifactWebhookOutput {
	return o
}

type ArtifactWebhookArrayOutput struct{ *pulumi.OutputState }

func (ArtifactWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ArtifactWebhook)(nil)).Elem()
}

func (o ArtifactWebhookArrayOutput) ToArtifactWebhookArrayOutput() ArtifactWebhookArrayOutput {
	return o
}

func (o ArtifactWebhookArrayOutput) ToArtifactWebhookArrayOutputWithContext(ctx context.Context) ArtifactWebhookArrayOutput {
	return o
}

func (o ArtifactWebhookArrayOutput) Index(i pulumi.IntInput) ArtifactWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ArtifactWebhook {
		return vs[0].([]*ArtifactWebhook)[vs[1].(int)]
	}).(ArtifactWebhookOutput)
}

type ArtifactWebhookMapOutput struct{ *pulumi.OutputState }

func (ArtifactWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ArtifactWebhook)(nil)).Elem()
}

func (o ArtifactWebhookMapOutput) ToArtifactWebhookMapOutput() ArtifactWebhookMapOutput {
	return o
}

func (o ArtifactWebhookMapOutput) ToArtifactWebhookMapOutputWithContext(ctx context.Context) ArtifactWebhookMapOutput {
	return o
}

func (o ArtifactWebhookMapOutput) MapIndex(k pulumi.StringInput) ArtifactWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ArtifactWebhook {
		return vs[0].(map[string]*ArtifactWebhook)[vs[1].(string)]
	}).(ArtifactWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ArtifactWebhookInput)(nil)).Elem(), &ArtifactWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*ArtifactWebhookArrayInput)(nil)).Elem(), ArtifactWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ArtifactWebhookMapInput)(nil)).Elem(), ArtifactWebhookMap{})
	pulumi.RegisterOutputType(ArtifactWebhookOutput{})
	pulumi.RegisterOutputType(ArtifactWebhookArrayOutput{})
	pulumi.RegisterOutputType(ArtifactWebhookMapOutput{})
}
