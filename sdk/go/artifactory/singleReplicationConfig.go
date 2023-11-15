// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Note: this resource is deprecated in favor of `PullReplication` resource.
//
// Provides an Artifactory single replication config resource. This can be used to create and manage a single Artifactory
// replication. Primarily used when pull replication is needed.
//
// **WARNING: This should not be used on a repository with `ReplicationConfig`. Using both together will cause
// unexpected behaviour and will almost certainly cause your replications to break.**
//
// ## Import
//
// Replication configs can be imported using their repo key, e.g.
//
// ```sh
//
//	$ pulumi import artifactory:index/singleReplicationConfig:SingleReplicationConfig foo-rep repository-key
//
// ```
type SingleReplicationConfig struct {
	pulumi.CustomResourceState

	// Cron expression to control the operation frequency.
	CronExp                pulumi.StringOutput `pulumi:"cronExp"`
	EnableEventReplication pulumi.BoolOutput   `pulumi:"enableEventReplication"`
	Enabled                pulumi.BoolOutput   `pulumi:"enabled"`
	// Requires password encryption to be turned off `POST /api/system/decrypt`.
	Password   pulumi.StringOutput    `pulumi:"password"`
	PathPrefix pulumi.StringPtrOutput `pulumi:"pathPrefix"`
	// Proxy key from Artifactory Proxies setting.
	Proxy               pulumi.StringPtrOutput `pulumi:"proxy"`
	RepoKey             pulumi.StringOutput    `pulumi:"repoKey"`
	SocketTimeoutMillis pulumi.IntOutput       `pulumi:"socketTimeoutMillis"`
	SyncDeletes         pulumi.BoolOutput      `pulumi:"syncDeletes"`
	SyncProperties      pulumi.BoolOutput      `pulumi:"syncProperties"`
	SyncStatistics      pulumi.BoolOutput      `pulumi:"syncStatistics"`
	Url                 pulumi.StringPtrOutput `pulumi:"url"`
	Username            pulumi.StringPtrOutput `pulumi:"username"`
}

// NewSingleReplicationConfig registers a new resource with the given unique name, arguments, and options.
func NewSingleReplicationConfig(ctx *pulumi.Context,
	name string, args *SingleReplicationConfigArgs, opts ...pulumi.ResourceOption) (*SingleReplicationConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CronExp == nil {
		return nil, errors.New("invalid value for required argument 'CronExp'")
	}
	if args.RepoKey == nil {
		return nil, errors.New("invalid value for required argument 'RepoKey'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SingleReplicationConfig
	err := ctx.RegisterResource("artifactory:index/singleReplicationConfig:SingleReplicationConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSingleReplicationConfig gets an existing SingleReplicationConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSingleReplicationConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SingleReplicationConfigState, opts ...pulumi.ResourceOption) (*SingleReplicationConfig, error) {
	var resource SingleReplicationConfig
	err := ctx.ReadResource("artifactory:index/singleReplicationConfig:SingleReplicationConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SingleReplicationConfig resources.
type singleReplicationConfigState struct {
	// Cron expression to control the operation frequency.
	CronExp                *string `pulumi:"cronExp"`
	EnableEventReplication *bool   `pulumi:"enableEventReplication"`
	Enabled                *bool   `pulumi:"enabled"`
	// Requires password encryption to be turned off `POST /api/system/decrypt`.
	Password   *string `pulumi:"password"`
	PathPrefix *string `pulumi:"pathPrefix"`
	// Proxy key from Artifactory Proxies setting.
	Proxy               *string `pulumi:"proxy"`
	RepoKey             *string `pulumi:"repoKey"`
	SocketTimeoutMillis *int    `pulumi:"socketTimeoutMillis"`
	SyncDeletes         *bool   `pulumi:"syncDeletes"`
	SyncProperties      *bool   `pulumi:"syncProperties"`
	SyncStatistics      *bool   `pulumi:"syncStatistics"`
	Url                 *string `pulumi:"url"`
	Username            *string `pulumi:"username"`
}

type SingleReplicationConfigState struct {
	// Cron expression to control the operation frequency.
	CronExp                pulumi.StringPtrInput
	EnableEventReplication pulumi.BoolPtrInput
	Enabled                pulumi.BoolPtrInput
	// Requires password encryption to be turned off `POST /api/system/decrypt`.
	Password   pulumi.StringPtrInput
	PathPrefix pulumi.StringPtrInput
	// Proxy key from Artifactory Proxies setting.
	Proxy               pulumi.StringPtrInput
	RepoKey             pulumi.StringPtrInput
	SocketTimeoutMillis pulumi.IntPtrInput
	SyncDeletes         pulumi.BoolPtrInput
	SyncProperties      pulumi.BoolPtrInput
	SyncStatistics      pulumi.BoolPtrInput
	Url                 pulumi.StringPtrInput
	Username            pulumi.StringPtrInput
}

func (SingleReplicationConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*singleReplicationConfigState)(nil)).Elem()
}

type singleReplicationConfigArgs struct {
	// Cron expression to control the operation frequency.
	CronExp                string  `pulumi:"cronExp"`
	EnableEventReplication *bool   `pulumi:"enableEventReplication"`
	Enabled                *bool   `pulumi:"enabled"`
	PathPrefix             *string `pulumi:"pathPrefix"`
	// Proxy key from Artifactory Proxies setting.
	Proxy               *string `pulumi:"proxy"`
	RepoKey             string  `pulumi:"repoKey"`
	SocketTimeoutMillis *int    `pulumi:"socketTimeoutMillis"`
	SyncDeletes         *bool   `pulumi:"syncDeletes"`
	SyncProperties      *bool   `pulumi:"syncProperties"`
	SyncStatistics      *bool   `pulumi:"syncStatistics"`
	Url                 *string `pulumi:"url"`
	Username            *string `pulumi:"username"`
}

// The set of arguments for constructing a SingleReplicationConfig resource.
type SingleReplicationConfigArgs struct {
	// Cron expression to control the operation frequency.
	CronExp                pulumi.StringInput
	EnableEventReplication pulumi.BoolPtrInput
	Enabled                pulumi.BoolPtrInput
	PathPrefix             pulumi.StringPtrInput
	// Proxy key from Artifactory Proxies setting.
	Proxy               pulumi.StringPtrInput
	RepoKey             pulumi.StringInput
	SocketTimeoutMillis pulumi.IntPtrInput
	SyncDeletes         pulumi.BoolPtrInput
	SyncProperties      pulumi.BoolPtrInput
	SyncStatistics      pulumi.BoolPtrInput
	Url                 pulumi.StringPtrInput
	Username            pulumi.StringPtrInput
}

func (SingleReplicationConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*singleReplicationConfigArgs)(nil)).Elem()
}

type SingleReplicationConfigInput interface {
	pulumi.Input

	ToSingleReplicationConfigOutput() SingleReplicationConfigOutput
	ToSingleReplicationConfigOutputWithContext(ctx context.Context) SingleReplicationConfigOutput
}

func (*SingleReplicationConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**SingleReplicationConfig)(nil)).Elem()
}

func (i *SingleReplicationConfig) ToSingleReplicationConfigOutput() SingleReplicationConfigOutput {
	return i.ToSingleReplicationConfigOutputWithContext(context.Background())
}

func (i *SingleReplicationConfig) ToSingleReplicationConfigOutputWithContext(ctx context.Context) SingleReplicationConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SingleReplicationConfigOutput)
}

// SingleReplicationConfigArrayInput is an input type that accepts SingleReplicationConfigArray and SingleReplicationConfigArrayOutput values.
// You can construct a concrete instance of `SingleReplicationConfigArrayInput` via:
//
//	SingleReplicationConfigArray{ SingleReplicationConfigArgs{...} }
type SingleReplicationConfigArrayInput interface {
	pulumi.Input

	ToSingleReplicationConfigArrayOutput() SingleReplicationConfigArrayOutput
	ToSingleReplicationConfigArrayOutputWithContext(context.Context) SingleReplicationConfigArrayOutput
}

type SingleReplicationConfigArray []SingleReplicationConfigInput

func (SingleReplicationConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SingleReplicationConfig)(nil)).Elem()
}

func (i SingleReplicationConfigArray) ToSingleReplicationConfigArrayOutput() SingleReplicationConfigArrayOutput {
	return i.ToSingleReplicationConfigArrayOutputWithContext(context.Background())
}

func (i SingleReplicationConfigArray) ToSingleReplicationConfigArrayOutputWithContext(ctx context.Context) SingleReplicationConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SingleReplicationConfigArrayOutput)
}

// SingleReplicationConfigMapInput is an input type that accepts SingleReplicationConfigMap and SingleReplicationConfigMapOutput values.
// You can construct a concrete instance of `SingleReplicationConfigMapInput` via:
//
//	SingleReplicationConfigMap{ "key": SingleReplicationConfigArgs{...} }
type SingleReplicationConfigMapInput interface {
	pulumi.Input

	ToSingleReplicationConfigMapOutput() SingleReplicationConfigMapOutput
	ToSingleReplicationConfigMapOutputWithContext(context.Context) SingleReplicationConfigMapOutput
}

type SingleReplicationConfigMap map[string]SingleReplicationConfigInput

func (SingleReplicationConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SingleReplicationConfig)(nil)).Elem()
}

func (i SingleReplicationConfigMap) ToSingleReplicationConfigMapOutput() SingleReplicationConfigMapOutput {
	return i.ToSingleReplicationConfigMapOutputWithContext(context.Background())
}

func (i SingleReplicationConfigMap) ToSingleReplicationConfigMapOutputWithContext(ctx context.Context) SingleReplicationConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SingleReplicationConfigMapOutput)
}

type SingleReplicationConfigOutput struct{ *pulumi.OutputState }

func (SingleReplicationConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SingleReplicationConfig)(nil)).Elem()
}

func (o SingleReplicationConfigOutput) ToSingleReplicationConfigOutput() SingleReplicationConfigOutput {
	return o
}

func (o SingleReplicationConfigOutput) ToSingleReplicationConfigOutputWithContext(ctx context.Context) SingleReplicationConfigOutput {
	return o
}

// Cron expression to control the operation frequency.
func (o SingleReplicationConfigOutput) CronExp() pulumi.StringOutput {
	return o.ApplyT(func(v *SingleReplicationConfig) pulumi.StringOutput { return v.CronExp }).(pulumi.StringOutput)
}

func (o SingleReplicationConfigOutput) EnableEventReplication() pulumi.BoolOutput {
	return o.ApplyT(func(v *SingleReplicationConfig) pulumi.BoolOutput { return v.EnableEventReplication }).(pulumi.BoolOutput)
}

func (o SingleReplicationConfigOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *SingleReplicationConfig) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// Requires password encryption to be turned off `POST /api/system/decrypt`.
func (o SingleReplicationConfigOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *SingleReplicationConfig) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

func (o SingleReplicationConfigOutput) PathPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SingleReplicationConfig) pulumi.StringPtrOutput { return v.PathPrefix }).(pulumi.StringPtrOutput)
}

// Proxy key from Artifactory Proxies setting.
func (o SingleReplicationConfigOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SingleReplicationConfig) pulumi.StringPtrOutput { return v.Proxy }).(pulumi.StringPtrOutput)
}

func (o SingleReplicationConfigOutput) RepoKey() pulumi.StringOutput {
	return o.ApplyT(func(v *SingleReplicationConfig) pulumi.StringOutput { return v.RepoKey }).(pulumi.StringOutput)
}

func (o SingleReplicationConfigOutput) SocketTimeoutMillis() pulumi.IntOutput {
	return o.ApplyT(func(v *SingleReplicationConfig) pulumi.IntOutput { return v.SocketTimeoutMillis }).(pulumi.IntOutput)
}

func (o SingleReplicationConfigOutput) SyncDeletes() pulumi.BoolOutput {
	return o.ApplyT(func(v *SingleReplicationConfig) pulumi.BoolOutput { return v.SyncDeletes }).(pulumi.BoolOutput)
}

func (o SingleReplicationConfigOutput) SyncProperties() pulumi.BoolOutput {
	return o.ApplyT(func(v *SingleReplicationConfig) pulumi.BoolOutput { return v.SyncProperties }).(pulumi.BoolOutput)
}

func (o SingleReplicationConfigOutput) SyncStatistics() pulumi.BoolOutput {
	return o.ApplyT(func(v *SingleReplicationConfig) pulumi.BoolOutput { return v.SyncStatistics }).(pulumi.BoolOutput)
}

func (o SingleReplicationConfigOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SingleReplicationConfig) pulumi.StringPtrOutput { return v.Url }).(pulumi.StringPtrOutput)
}

func (o SingleReplicationConfigOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SingleReplicationConfig) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

type SingleReplicationConfigArrayOutput struct{ *pulumi.OutputState }

func (SingleReplicationConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SingleReplicationConfig)(nil)).Elem()
}

func (o SingleReplicationConfigArrayOutput) ToSingleReplicationConfigArrayOutput() SingleReplicationConfigArrayOutput {
	return o
}

func (o SingleReplicationConfigArrayOutput) ToSingleReplicationConfigArrayOutputWithContext(ctx context.Context) SingleReplicationConfigArrayOutput {
	return o
}

func (o SingleReplicationConfigArrayOutput) Index(i pulumi.IntInput) SingleReplicationConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SingleReplicationConfig {
		return vs[0].([]*SingleReplicationConfig)[vs[1].(int)]
	}).(SingleReplicationConfigOutput)
}

type SingleReplicationConfigMapOutput struct{ *pulumi.OutputState }

func (SingleReplicationConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SingleReplicationConfig)(nil)).Elem()
}

func (o SingleReplicationConfigMapOutput) ToSingleReplicationConfigMapOutput() SingleReplicationConfigMapOutput {
	return o
}

func (o SingleReplicationConfigMapOutput) ToSingleReplicationConfigMapOutputWithContext(ctx context.Context) SingleReplicationConfigMapOutput {
	return o
}

func (o SingleReplicationConfigMapOutput) MapIndex(k pulumi.StringInput) SingleReplicationConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SingleReplicationConfig {
		return vs[0].(map[string]*SingleReplicationConfig)[vs[1].(string)]
	}).(SingleReplicationConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SingleReplicationConfigInput)(nil)).Elem(), &SingleReplicationConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*SingleReplicationConfigArrayInput)(nil)).Elem(), SingleReplicationConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SingleReplicationConfigMapInput)(nil)).Elem(), SingleReplicationConfigMap{})
	pulumi.RegisterOutputType(SingleReplicationConfigOutput{})
	pulumi.RegisterOutputType(SingleReplicationConfigArrayOutput{})
	pulumi.RegisterOutputType(SingleReplicationConfigMapOutput{})
}
