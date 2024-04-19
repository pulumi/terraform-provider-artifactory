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

// Provides an Artifactory Mail Server resource. This can be used to create and manage Artifactory mail server configuration.
//
// ## Example Usage
//
// ### S
//
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
//			_, err := artifactory.NewMailServer(ctx, "mymailserver", &artifactory.MailServerArgs{
//				Enabled:        pulumi.Bool(true),
//				ArtifactoryUrl: pulumi.String("http://tempurl.org"),
//				From:           pulumi.String("test@jfrog.com"),
//				Host:           pulumi.String("http://tempurl.org"),
//				Username:       pulumi.String("test-user"),
//				Password:       pulumi.String("test-password"),
//				Port:           pulumi.Int(25),
//				SubjectPrefix:  pulumi.String("[Test]"),
//				UseSsl:         pulumi.Bool(true),
//				UseTls:         pulumi.Bool(true),
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
//
// ## Import
//
// ```sh
// $ pulumi import artifactory:index/mailServer:MailServer my-mail-server mymailserver
// ```
//
// ~>The `password` attribute is not retrievable from Artifactory thus there will be state drift after importing this resource.
type MailServer struct {
	pulumi.CustomResourceState

	// The Artifactory URL to to link to in all outgoing messages.
	ArtifactoryUrl pulumi.StringPtrOutput `pulumi:"artifactoryUrl"`
	// When set, mail notifications are enabled.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The 'from' address header to use in all outgoing messages.
	From pulumi.StringPtrOutput `pulumi:"from"`
	// The mail server IP address / DNS.
	Host pulumi.StringOutput `pulumi:"host"`
	// The password for authentication with the mail server.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The port number of the mail server.
	Port pulumi.IntOutput `pulumi:"port"`
	// A prefix to use for the subject of all outgoing mails.
	SubjectPrefix pulumi.StringPtrOutput `pulumi:"subjectPrefix"`
	// When set to 'true', uses a secure connection to the mail server.
	UseSsl pulumi.BoolOutput `pulumi:"useSsl"`
	// When set to 'true', uses Transport Layer Security when connecting to the mail server.
	UseTls pulumi.BoolOutput `pulumi:"useTls"`
	// The username for authentication with the mail server.
	Username pulumi.StringPtrOutput `pulumi:"username"`
}

// NewMailServer registers a new resource with the given unique name, arguments, and options.
func NewMailServer(ctx *pulumi.Context,
	name string, args *MailServerArgs, opts ...pulumi.ResourceOption) (*MailServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.Host == nil {
		return nil, errors.New("invalid value for required argument 'Host'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MailServer
	err := ctx.RegisterResource("artifactory:index/mailServer:MailServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMailServer gets an existing MailServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMailServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MailServerState, opts ...pulumi.ResourceOption) (*MailServer, error) {
	var resource MailServer
	err := ctx.ReadResource("artifactory:index/mailServer:MailServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MailServer resources.
type mailServerState struct {
	// The Artifactory URL to to link to in all outgoing messages.
	ArtifactoryUrl *string `pulumi:"artifactoryUrl"`
	// When set, mail notifications are enabled.
	Enabled *bool `pulumi:"enabled"`
	// The 'from' address header to use in all outgoing messages.
	From *string `pulumi:"from"`
	// The mail server IP address / DNS.
	Host *string `pulumi:"host"`
	// The password for authentication with the mail server.
	Password *string `pulumi:"password"`
	// The port number of the mail server.
	Port *int `pulumi:"port"`
	// A prefix to use for the subject of all outgoing mails.
	SubjectPrefix *string `pulumi:"subjectPrefix"`
	// When set to 'true', uses a secure connection to the mail server.
	UseSsl *bool `pulumi:"useSsl"`
	// When set to 'true', uses Transport Layer Security when connecting to the mail server.
	UseTls *bool `pulumi:"useTls"`
	// The username for authentication with the mail server.
	Username *string `pulumi:"username"`
}

type MailServerState struct {
	// The Artifactory URL to to link to in all outgoing messages.
	ArtifactoryUrl pulumi.StringPtrInput
	// When set, mail notifications are enabled.
	Enabled pulumi.BoolPtrInput
	// The 'from' address header to use in all outgoing messages.
	From pulumi.StringPtrInput
	// The mail server IP address / DNS.
	Host pulumi.StringPtrInput
	// The password for authentication with the mail server.
	Password pulumi.StringPtrInput
	// The port number of the mail server.
	Port pulumi.IntPtrInput
	// A prefix to use for the subject of all outgoing mails.
	SubjectPrefix pulumi.StringPtrInput
	// When set to 'true', uses a secure connection to the mail server.
	UseSsl pulumi.BoolPtrInput
	// When set to 'true', uses Transport Layer Security when connecting to the mail server.
	UseTls pulumi.BoolPtrInput
	// The username for authentication with the mail server.
	Username pulumi.StringPtrInput
}

func (MailServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*mailServerState)(nil)).Elem()
}

type mailServerArgs struct {
	// The Artifactory URL to to link to in all outgoing messages.
	ArtifactoryUrl *string `pulumi:"artifactoryUrl"`
	// When set, mail notifications are enabled.
	Enabled bool `pulumi:"enabled"`
	// The 'from' address header to use in all outgoing messages.
	From *string `pulumi:"from"`
	// The mail server IP address / DNS.
	Host string `pulumi:"host"`
	// The password for authentication with the mail server.
	Password *string `pulumi:"password"`
	// The port number of the mail server.
	Port int `pulumi:"port"`
	// A prefix to use for the subject of all outgoing mails.
	SubjectPrefix *string `pulumi:"subjectPrefix"`
	// When set to 'true', uses a secure connection to the mail server.
	UseSsl *bool `pulumi:"useSsl"`
	// When set to 'true', uses Transport Layer Security when connecting to the mail server.
	UseTls *bool `pulumi:"useTls"`
	// The username for authentication with the mail server.
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a MailServer resource.
type MailServerArgs struct {
	// The Artifactory URL to to link to in all outgoing messages.
	ArtifactoryUrl pulumi.StringPtrInput
	// When set, mail notifications are enabled.
	Enabled pulumi.BoolInput
	// The 'from' address header to use in all outgoing messages.
	From pulumi.StringPtrInput
	// The mail server IP address / DNS.
	Host pulumi.StringInput
	// The password for authentication with the mail server.
	Password pulumi.StringPtrInput
	// The port number of the mail server.
	Port pulumi.IntInput
	// A prefix to use for the subject of all outgoing mails.
	SubjectPrefix pulumi.StringPtrInput
	// When set to 'true', uses a secure connection to the mail server.
	UseSsl pulumi.BoolPtrInput
	// When set to 'true', uses Transport Layer Security when connecting to the mail server.
	UseTls pulumi.BoolPtrInput
	// The username for authentication with the mail server.
	Username pulumi.StringPtrInput
}

func (MailServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mailServerArgs)(nil)).Elem()
}

type MailServerInput interface {
	pulumi.Input

	ToMailServerOutput() MailServerOutput
	ToMailServerOutputWithContext(ctx context.Context) MailServerOutput
}

func (*MailServer) ElementType() reflect.Type {
	return reflect.TypeOf((**MailServer)(nil)).Elem()
}

func (i *MailServer) ToMailServerOutput() MailServerOutput {
	return i.ToMailServerOutputWithContext(context.Background())
}

func (i *MailServer) ToMailServerOutputWithContext(ctx context.Context) MailServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MailServerOutput)
}

// MailServerArrayInput is an input type that accepts MailServerArray and MailServerArrayOutput values.
// You can construct a concrete instance of `MailServerArrayInput` via:
//
//	MailServerArray{ MailServerArgs{...} }
type MailServerArrayInput interface {
	pulumi.Input

	ToMailServerArrayOutput() MailServerArrayOutput
	ToMailServerArrayOutputWithContext(context.Context) MailServerArrayOutput
}

type MailServerArray []MailServerInput

func (MailServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MailServer)(nil)).Elem()
}

func (i MailServerArray) ToMailServerArrayOutput() MailServerArrayOutput {
	return i.ToMailServerArrayOutputWithContext(context.Background())
}

func (i MailServerArray) ToMailServerArrayOutputWithContext(ctx context.Context) MailServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MailServerArrayOutput)
}

// MailServerMapInput is an input type that accepts MailServerMap and MailServerMapOutput values.
// You can construct a concrete instance of `MailServerMapInput` via:
//
//	MailServerMap{ "key": MailServerArgs{...} }
type MailServerMapInput interface {
	pulumi.Input

	ToMailServerMapOutput() MailServerMapOutput
	ToMailServerMapOutputWithContext(context.Context) MailServerMapOutput
}

type MailServerMap map[string]MailServerInput

func (MailServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MailServer)(nil)).Elem()
}

func (i MailServerMap) ToMailServerMapOutput() MailServerMapOutput {
	return i.ToMailServerMapOutputWithContext(context.Background())
}

func (i MailServerMap) ToMailServerMapOutputWithContext(ctx context.Context) MailServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MailServerMapOutput)
}

type MailServerOutput struct{ *pulumi.OutputState }

func (MailServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MailServer)(nil)).Elem()
}

func (o MailServerOutput) ToMailServerOutput() MailServerOutput {
	return o
}

func (o MailServerOutput) ToMailServerOutputWithContext(ctx context.Context) MailServerOutput {
	return o
}

// The Artifactory URL to to link to in all outgoing messages.
func (o MailServerOutput) ArtifactoryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MailServer) pulumi.StringPtrOutput { return v.ArtifactoryUrl }).(pulumi.StringPtrOutput)
}

// When set, mail notifications are enabled.
func (o MailServerOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *MailServer) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// The 'from' address header to use in all outgoing messages.
func (o MailServerOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MailServer) pulumi.StringPtrOutput { return v.From }).(pulumi.StringPtrOutput)
}

// The mail server IP address / DNS.
func (o MailServerOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *MailServer) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

// The password for authentication with the mail server.
func (o MailServerOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MailServer) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The port number of the mail server.
func (o MailServerOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *MailServer) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// A prefix to use for the subject of all outgoing mails.
func (o MailServerOutput) SubjectPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MailServer) pulumi.StringPtrOutput { return v.SubjectPrefix }).(pulumi.StringPtrOutput)
}

// When set to 'true', uses a secure connection to the mail server.
func (o MailServerOutput) UseSsl() pulumi.BoolOutput {
	return o.ApplyT(func(v *MailServer) pulumi.BoolOutput { return v.UseSsl }).(pulumi.BoolOutput)
}

// When set to 'true', uses Transport Layer Security when connecting to the mail server.
func (o MailServerOutput) UseTls() pulumi.BoolOutput {
	return o.ApplyT(func(v *MailServer) pulumi.BoolOutput { return v.UseTls }).(pulumi.BoolOutput)
}

// The username for authentication with the mail server.
func (o MailServerOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MailServer) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

type MailServerArrayOutput struct{ *pulumi.OutputState }

func (MailServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MailServer)(nil)).Elem()
}

func (o MailServerArrayOutput) ToMailServerArrayOutput() MailServerArrayOutput {
	return o
}

func (o MailServerArrayOutput) ToMailServerArrayOutputWithContext(ctx context.Context) MailServerArrayOutput {
	return o
}

func (o MailServerArrayOutput) Index(i pulumi.IntInput) MailServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MailServer {
		return vs[0].([]*MailServer)[vs[1].(int)]
	}).(MailServerOutput)
}

type MailServerMapOutput struct{ *pulumi.OutputState }

func (MailServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MailServer)(nil)).Elem()
}

func (o MailServerMapOutput) ToMailServerMapOutput() MailServerMapOutput {
	return o
}

func (o MailServerMapOutput) ToMailServerMapOutputWithContext(ctx context.Context) MailServerMapOutput {
	return o
}

func (o MailServerMapOutput) MapIndex(k pulumi.StringInput) MailServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MailServer {
		return vs[0].(map[string]*MailServer)[vs[1].(string)]
	}).(MailServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MailServerInput)(nil)).Elem(), &MailServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*MailServerArrayInput)(nil)).Elem(), MailServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MailServerMapInput)(nil)).Elem(), MailServerMap{})
	pulumi.RegisterOutputType(MailServerOutput{})
	pulumi.RegisterOutputType(MailServerArrayOutput{})
	pulumi.RegisterOutputType(MailServerMapOutput{})
}
