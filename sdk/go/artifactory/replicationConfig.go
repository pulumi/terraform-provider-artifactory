// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Note: this resource is deprecated in favor of `PushReplication` resource.
//
// Provides an Artifactory replication config resource. This can be used to create and manage Artifactory replications.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v3/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			providerTestSource, err := artifactory.NewLocalMavenRepository(ctx, "providerTestSource", &artifactory.LocalMavenRepositoryArgs{
//				Key: pulumi.String("provider_test_source"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = artifactory.NewLocalMavenRepository(ctx, "providerTestDest", &artifactory.LocalMavenRepositoryArgs{
//				Key: pulumi.String("provider_test_dest"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = artifactory.NewReplicationConfig(ctx, "foo-rep", &artifactory.ReplicationConfigArgs{
//				CronExp:                pulumi.String("0 0 * * * ?"),
//				EnableEventReplication: pulumi.Bool(true),
//				Replications: artifactory.ReplicationConfigReplicationArray{
//					&artifactory.ReplicationConfigReplicationArgs{
//						Password: pulumi.String("$var.artifactory_password"),
//						Url:      pulumi.String("$var.artifactory_url"),
//						Username: pulumi.String("$var.artifactory_username"),
//					},
//				},
//				RepoKey: providerTestSource.Key,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Replication configs can be imported using their repo key, e.g.
//
// ```sh
//
//	$ pulumi import artifactory:index/replicationConfig:ReplicationConfig foo-rep provider_test_source
//
// ```
type ReplicationConfig struct {
	pulumi.CustomResourceState

	// Cron expression to control the operation frequency.
	CronExp                pulumi.StringOutput                     `pulumi:"cronExp"`
	EnableEventReplication pulumi.BoolOutput                       `pulumi:"enableEventReplication"`
	Replications           ReplicationConfigReplicationArrayOutput `pulumi:"replications"`
	RepoKey                pulumi.StringOutput                     `pulumi:"repoKey"`
}

// NewReplicationConfig registers a new resource with the given unique name, arguments, and options.
func NewReplicationConfig(ctx *pulumi.Context,
	name string, args *ReplicationConfigArgs, opts ...pulumi.ResourceOption) (*ReplicationConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CronExp == nil {
		return nil, errors.New("invalid value for required argument 'CronExp'")
	}
	if args.RepoKey == nil {
		return nil, errors.New("invalid value for required argument 'RepoKey'")
	}
	var resource ReplicationConfig
	err := ctx.RegisterResource("artifactory:index/replicationConfig:ReplicationConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationConfig gets an existing ReplicationConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationConfigState, opts ...pulumi.ResourceOption) (*ReplicationConfig, error) {
	var resource ReplicationConfig
	err := ctx.ReadResource("artifactory:index/replicationConfig:ReplicationConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationConfig resources.
type replicationConfigState struct {
	// Cron expression to control the operation frequency.
	CronExp                *string                        `pulumi:"cronExp"`
	EnableEventReplication *bool                          `pulumi:"enableEventReplication"`
	Replications           []ReplicationConfigReplication `pulumi:"replications"`
	RepoKey                *string                        `pulumi:"repoKey"`
}

type ReplicationConfigState struct {
	// Cron expression to control the operation frequency.
	CronExp                pulumi.StringPtrInput
	EnableEventReplication pulumi.BoolPtrInput
	Replications           ReplicationConfigReplicationArrayInput
	RepoKey                pulumi.StringPtrInput
}

func (ReplicationConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationConfigState)(nil)).Elem()
}

type replicationConfigArgs struct {
	// Cron expression to control the operation frequency.
	CronExp                string                         `pulumi:"cronExp"`
	EnableEventReplication *bool                          `pulumi:"enableEventReplication"`
	Replications           []ReplicationConfigReplication `pulumi:"replications"`
	RepoKey                string                         `pulumi:"repoKey"`
}

// The set of arguments for constructing a ReplicationConfig resource.
type ReplicationConfigArgs struct {
	// Cron expression to control the operation frequency.
	CronExp                pulumi.StringInput
	EnableEventReplication pulumi.BoolPtrInput
	Replications           ReplicationConfigReplicationArrayInput
	RepoKey                pulumi.StringInput
}

func (ReplicationConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationConfigArgs)(nil)).Elem()
}

type ReplicationConfigInput interface {
	pulumi.Input

	ToReplicationConfigOutput() ReplicationConfigOutput
	ToReplicationConfigOutputWithContext(ctx context.Context) ReplicationConfigOutput
}

func (*ReplicationConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationConfig)(nil)).Elem()
}

func (i *ReplicationConfig) ToReplicationConfigOutput() ReplicationConfigOutput {
	return i.ToReplicationConfigOutputWithContext(context.Background())
}

func (i *ReplicationConfig) ToReplicationConfigOutputWithContext(ctx context.Context) ReplicationConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationConfigOutput)
}

// ReplicationConfigArrayInput is an input type that accepts ReplicationConfigArray and ReplicationConfigArrayOutput values.
// You can construct a concrete instance of `ReplicationConfigArrayInput` via:
//
//	ReplicationConfigArray{ ReplicationConfigArgs{...} }
type ReplicationConfigArrayInput interface {
	pulumi.Input

	ToReplicationConfigArrayOutput() ReplicationConfigArrayOutput
	ToReplicationConfigArrayOutputWithContext(context.Context) ReplicationConfigArrayOutput
}

type ReplicationConfigArray []ReplicationConfigInput

func (ReplicationConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicationConfig)(nil)).Elem()
}

func (i ReplicationConfigArray) ToReplicationConfigArrayOutput() ReplicationConfigArrayOutput {
	return i.ToReplicationConfigArrayOutputWithContext(context.Background())
}

func (i ReplicationConfigArray) ToReplicationConfigArrayOutputWithContext(ctx context.Context) ReplicationConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationConfigArrayOutput)
}

// ReplicationConfigMapInput is an input type that accepts ReplicationConfigMap and ReplicationConfigMapOutput values.
// You can construct a concrete instance of `ReplicationConfigMapInput` via:
//
//	ReplicationConfigMap{ "key": ReplicationConfigArgs{...} }
type ReplicationConfigMapInput interface {
	pulumi.Input

	ToReplicationConfigMapOutput() ReplicationConfigMapOutput
	ToReplicationConfigMapOutputWithContext(context.Context) ReplicationConfigMapOutput
}

type ReplicationConfigMap map[string]ReplicationConfigInput

func (ReplicationConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicationConfig)(nil)).Elem()
}

func (i ReplicationConfigMap) ToReplicationConfigMapOutput() ReplicationConfigMapOutput {
	return i.ToReplicationConfigMapOutputWithContext(context.Background())
}

func (i ReplicationConfigMap) ToReplicationConfigMapOutputWithContext(ctx context.Context) ReplicationConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationConfigMapOutput)
}

type ReplicationConfigOutput struct{ *pulumi.OutputState }

func (ReplicationConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationConfig)(nil)).Elem()
}

func (o ReplicationConfigOutput) ToReplicationConfigOutput() ReplicationConfigOutput {
	return o
}

func (o ReplicationConfigOutput) ToReplicationConfigOutputWithContext(ctx context.Context) ReplicationConfigOutput {
	return o
}

// Cron expression to control the operation frequency.
func (o ReplicationConfigOutput) CronExp() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationConfig) pulumi.StringOutput { return v.CronExp }).(pulumi.StringOutput)
}

func (o ReplicationConfigOutput) EnableEventReplication() pulumi.BoolOutput {
	return o.ApplyT(func(v *ReplicationConfig) pulumi.BoolOutput { return v.EnableEventReplication }).(pulumi.BoolOutput)
}

func (o ReplicationConfigOutput) Replications() ReplicationConfigReplicationArrayOutput {
	return o.ApplyT(func(v *ReplicationConfig) ReplicationConfigReplicationArrayOutput { return v.Replications }).(ReplicationConfigReplicationArrayOutput)
}

func (o ReplicationConfigOutput) RepoKey() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationConfig) pulumi.StringOutput { return v.RepoKey }).(pulumi.StringOutput)
}

type ReplicationConfigArrayOutput struct{ *pulumi.OutputState }

func (ReplicationConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicationConfig)(nil)).Elem()
}

func (o ReplicationConfigArrayOutput) ToReplicationConfigArrayOutput() ReplicationConfigArrayOutput {
	return o
}

func (o ReplicationConfigArrayOutput) ToReplicationConfigArrayOutputWithContext(ctx context.Context) ReplicationConfigArrayOutput {
	return o
}

func (o ReplicationConfigArrayOutput) Index(i pulumi.IntInput) ReplicationConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReplicationConfig {
		return vs[0].([]*ReplicationConfig)[vs[1].(int)]
	}).(ReplicationConfigOutput)
}

type ReplicationConfigMapOutput struct{ *pulumi.OutputState }

func (ReplicationConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicationConfig)(nil)).Elem()
}

func (o ReplicationConfigMapOutput) ToReplicationConfigMapOutput() ReplicationConfigMapOutput {
	return o
}

func (o ReplicationConfigMapOutput) ToReplicationConfigMapOutputWithContext(ctx context.Context) ReplicationConfigMapOutput {
	return o
}

func (o ReplicationConfigMapOutput) MapIndex(k pulumi.StringInput) ReplicationConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReplicationConfig {
		return vs[0].(map[string]*ReplicationConfig)[vs[1].(string)]
	}).(ReplicationConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationConfigInput)(nil)).Elem(), &ReplicationConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationConfigArrayInput)(nil)).Elem(), ReplicationConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationConfigMapInput)(nil)).Elem(), ReplicationConfigMap{})
	pulumi.RegisterOutputType(ReplicationConfigOutput{})
	pulumi.RegisterOutputType(ReplicationConfigArrayOutput{})
	pulumi.RegisterOutputType(ReplicationConfigMapOutput{})
}
