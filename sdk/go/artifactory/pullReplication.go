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

// > This resource is deprecated and replaced by `RemoteRepositoryReplication` for clarity.
//
// Provides an Artifactory pull replication resource. This can be used to create and manage pull replication in Artifactory for a local or remote repo. Pull replication provides a convenient way to proactively populate a remote cache, and is very useful when waiting for new artifacts to arrive on demand (when first requested) is not desirable due to network latency.
//
// See the [Official Documentation](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-PullReplication).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v6/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Create a replication between two artifactory local repositories
//			_, err := artifactory.NewLocalMavenRepository(ctx, "providerTestSource", &artifactory.LocalMavenRepositoryArgs{
//				Key: pulumi.String("provider_test_source"),
//			})
//			if err != nil {
//				return err
//			}
//			providerTestDest, err := artifactory.NewRemoteMavenRepository(ctx, "providerTestDest", &artifactory.RemoteMavenRepositoryArgs{
//				Key:      pulumi.String("provider_test_dest"),
//				Password: pulumi.String("bar"),
//				Url:      pulumi.String(fmt.Sprintf("https://example.com/artifactory/%v", artifactory_local_maven_repository.Artifactory_local_maven_repository.Key)),
//				Username: pulumi.String("foo"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = artifactory.NewPullReplication(ctx, "remote-rep", &artifactory.PullReplicationArgs{
//				CronExp:                pulumi.String("0 0 * * * ?"),
//				EnableEventReplication: pulumi.Bool(true),
//				RepoKey:                providerTestDest.Key,
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
// Pull replication config can be imported using its repo key, e.g.
//
// ```sh
//
//	$ pulumi import artifactory:index/pullReplication:PullReplication foo-rep repository-key
//
// ```
type PullReplication struct {
	pulumi.CustomResourceState

	// When true, enables distributed checksum storage. For more information, see
	// [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
	CheckBinaryExistenceInFilestore pulumi.BoolPtrOutput `pulumi:"checkBinaryExistenceInFilestore"`
	// A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
	CronExp pulumi.StringPtrOutput `pulumi:"cronExp"`
	// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. added, deleted or property change.
	EnableEventReplication pulumi.BoolPtrOutput `pulumi:"enableEventReplication"`
	// When set, this replication will be enabled when saved.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Required for local repository, but not needed for remote repository.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Only artifacts that located in path that matches the subpath within the remote repository will be replicated.
	PathPrefix pulumi.StringPtrOutput `pulumi:"pathPrefix"`
	// Proxy key from Artifactory Proxies setting
	Proxy pulumi.StringPtrOutput `pulumi:"proxy"`
	// Repository name.
	RepoKey             pulumi.StringOutput `pulumi:"repoKey"`
	SocketTimeoutMillis pulumi.IntOutput    `pulumi:"socketTimeoutMillis"`
	// When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata).
	SyncDeletes pulumi.BoolOutput `pulumi:"syncDeletes"`
	// When set, the task also synchronizes the properties of replicated artifacts.
	SyncProperties pulumi.BoolOutput `pulumi:"syncProperties"`
	// When set, artifact download statistics will also be replicated. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery.
	SyncStatistics pulumi.BoolOutput `pulumi:"syncStatistics"`
	// The URL of the target local repository on a remote Artifactory server. For some package types, you need to prefix the repository key in the URL with api/<pkg>.
	// For a list of package types where this is required, see the [note](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-anchorPREFIX).
	// Required for local repository, but not needed for remote repository.
	Url pulumi.StringPtrOutput `pulumi:"url"`
	// Required for local repository, but not needed for remote repository.
	Username pulumi.StringPtrOutput `pulumi:"username"`
}

// NewPullReplication registers a new resource with the given unique name, arguments, and options.
func NewPullReplication(ctx *pulumi.Context,
	name string, args *PullReplicationArgs, opts ...pulumi.ResourceOption) (*PullReplication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RepoKey == nil {
		return nil, errors.New("invalid value for required argument 'RepoKey'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PullReplication
	err := ctx.RegisterResource("artifactory:index/pullReplication:PullReplication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPullReplication gets an existing PullReplication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPullReplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PullReplicationState, opts ...pulumi.ResourceOption) (*PullReplication, error) {
	var resource PullReplication
	err := ctx.ReadResource("artifactory:index/pullReplication:PullReplication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PullReplication resources.
type pullReplicationState struct {
	// When true, enables distributed checksum storage. For more information, see
	// [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
	CheckBinaryExistenceInFilestore *bool `pulumi:"checkBinaryExistenceInFilestore"`
	// A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
	CronExp *string `pulumi:"cronExp"`
	// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. added, deleted or property change.
	EnableEventReplication *bool `pulumi:"enableEventReplication"`
	// When set, this replication will be enabled when saved.
	Enabled *bool `pulumi:"enabled"`
	// Required for local repository, but not needed for remote repository.
	Password *string `pulumi:"password"`
	// Only artifacts that located in path that matches the subpath within the remote repository will be replicated.
	PathPrefix *string `pulumi:"pathPrefix"`
	// Proxy key from Artifactory Proxies setting
	Proxy *string `pulumi:"proxy"`
	// Repository name.
	RepoKey             *string `pulumi:"repoKey"`
	SocketTimeoutMillis *int    `pulumi:"socketTimeoutMillis"`
	// When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata).
	SyncDeletes *bool `pulumi:"syncDeletes"`
	// When set, the task also synchronizes the properties of replicated artifacts.
	SyncProperties *bool `pulumi:"syncProperties"`
	// When set, artifact download statistics will also be replicated. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery.
	SyncStatistics *bool `pulumi:"syncStatistics"`
	// The URL of the target local repository on a remote Artifactory server. For some package types, you need to prefix the repository key in the URL with api/<pkg>.
	// For a list of package types where this is required, see the [note](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-anchorPREFIX).
	// Required for local repository, but not needed for remote repository.
	Url *string `pulumi:"url"`
	// Required for local repository, but not needed for remote repository.
	Username *string `pulumi:"username"`
}

type PullReplicationState struct {
	// When true, enables distributed checksum storage. For more information, see
	// [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
	CheckBinaryExistenceInFilestore pulumi.BoolPtrInput
	// A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
	CronExp pulumi.StringPtrInput
	// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. added, deleted or property change.
	EnableEventReplication pulumi.BoolPtrInput
	// When set, this replication will be enabled when saved.
	Enabled pulumi.BoolPtrInput
	// Required for local repository, but not needed for remote repository.
	Password pulumi.StringPtrInput
	// Only artifacts that located in path that matches the subpath within the remote repository will be replicated.
	PathPrefix pulumi.StringPtrInput
	// Proxy key from Artifactory Proxies setting
	Proxy pulumi.StringPtrInput
	// Repository name.
	RepoKey             pulumi.StringPtrInput
	SocketTimeoutMillis pulumi.IntPtrInput
	// When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata).
	SyncDeletes pulumi.BoolPtrInput
	// When set, the task also synchronizes the properties of replicated artifacts.
	SyncProperties pulumi.BoolPtrInput
	// When set, artifact download statistics will also be replicated. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery.
	SyncStatistics pulumi.BoolPtrInput
	// The URL of the target local repository on a remote Artifactory server. For some package types, you need to prefix the repository key in the URL with api/<pkg>.
	// For a list of package types where this is required, see the [note](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-anchorPREFIX).
	// Required for local repository, but not needed for remote repository.
	Url pulumi.StringPtrInput
	// Required for local repository, but not needed for remote repository.
	Username pulumi.StringPtrInput
}

func (PullReplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*pullReplicationState)(nil)).Elem()
}

type pullReplicationArgs struct {
	// When true, enables distributed checksum storage. For more information, see
	// [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
	CheckBinaryExistenceInFilestore *bool `pulumi:"checkBinaryExistenceInFilestore"`
	// A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
	CronExp *string `pulumi:"cronExp"`
	// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. added, deleted or property change.
	EnableEventReplication *bool `pulumi:"enableEventReplication"`
	// When set, this replication will be enabled when saved.
	Enabled *bool `pulumi:"enabled"`
	// Required for local repository, but not needed for remote repository.
	Password *string `pulumi:"password"`
	// Only artifacts that located in path that matches the subpath within the remote repository will be replicated.
	PathPrefix *string `pulumi:"pathPrefix"`
	// Proxy key from Artifactory Proxies setting
	Proxy *string `pulumi:"proxy"`
	// Repository name.
	RepoKey             string `pulumi:"repoKey"`
	SocketTimeoutMillis *int   `pulumi:"socketTimeoutMillis"`
	// When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata).
	SyncDeletes *bool `pulumi:"syncDeletes"`
	// When set, the task also synchronizes the properties of replicated artifacts.
	SyncProperties *bool `pulumi:"syncProperties"`
	// When set, artifact download statistics will also be replicated. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery.
	SyncStatistics *bool `pulumi:"syncStatistics"`
	// The URL of the target local repository on a remote Artifactory server. For some package types, you need to prefix the repository key in the URL with api/<pkg>.
	// For a list of package types where this is required, see the [note](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-anchorPREFIX).
	// Required for local repository, but not needed for remote repository.
	Url *string `pulumi:"url"`
	// Required for local repository, but not needed for remote repository.
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a PullReplication resource.
type PullReplicationArgs struct {
	// When true, enables distributed checksum storage. For more information, see
	// [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
	CheckBinaryExistenceInFilestore pulumi.BoolPtrInput
	// A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
	CronExp pulumi.StringPtrInput
	// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. added, deleted or property change.
	EnableEventReplication pulumi.BoolPtrInput
	// When set, this replication will be enabled when saved.
	Enabled pulumi.BoolPtrInput
	// Required for local repository, but not needed for remote repository.
	Password pulumi.StringPtrInput
	// Only artifacts that located in path that matches the subpath within the remote repository will be replicated.
	PathPrefix pulumi.StringPtrInput
	// Proxy key from Artifactory Proxies setting
	Proxy pulumi.StringPtrInput
	// Repository name.
	RepoKey             pulumi.StringInput
	SocketTimeoutMillis pulumi.IntPtrInput
	// When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata).
	SyncDeletes pulumi.BoolPtrInput
	// When set, the task also synchronizes the properties of replicated artifacts.
	SyncProperties pulumi.BoolPtrInput
	// When set, artifact download statistics will also be replicated. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery.
	SyncStatistics pulumi.BoolPtrInput
	// The URL of the target local repository on a remote Artifactory server. For some package types, you need to prefix the repository key in the URL with api/<pkg>.
	// For a list of package types where this is required, see the [note](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-anchorPREFIX).
	// Required for local repository, but not needed for remote repository.
	Url pulumi.StringPtrInput
	// Required for local repository, but not needed for remote repository.
	Username pulumi.StringPtrInput
}

func (PullReplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pullReplicationArgs)(nil)).Elem()
}

type PullReplicationInput interface {
	pulumi.Input

	ToPullReplicationOutput() PullReplicationOutput
	ToPullReplicationOutputWithContext(ctx context.Context) PullReplicationOutput
}

func (*PullReplication) ElementType() reflect.Type {
	return reflect.TypeOf((**PullReplication)(nil)).Elem()
}

func (i *PullReplication) ToPullReplicationOutput() PullReplicationOutput {
	return i.ToPullReplicationOutputWithContext(context.Background())
}

func (i *PullReplication) ToPullReplicationOutputWithContext(ctx context.Context) PullReplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PullReplicationOutput)
}

// PullReplicationArrayInput is an input type that accepts PullReplicationArray and PullReplicationArrayOutput values.
// You can construct a concrete instance of `PullReplicationArrayInput` via:
//
//	PullReplicationArray{ PullReplicationArgs{...} }
type PullReplicationArrayInput interface {
	pulumi.Input

	ToPullReplicationArrayOutput() PullReplicationArrayOutput
	ToPullReplicationArrayOutputWithContext(context.Context) PullReplicationArrayOutput
}

type PullReplicationArray []PullReplicationInput

func (PullReplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PullReplication)(nil)).Elem()
}

func (i PullReplicationArray) ToPullReplicationArrayOutput() PullReplicationArrayOutput {
	return i.ToPullReplicationArrayOutputWithContext(context.Background())
}

func (i PullReplicationArray) ToPullReplicationArrayOutputWithContext(ctx context.Context) PullReplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PullReplicationArrayOutput)
}

// PullReplicationMapInput is an input type that accepts PullReplicationMap and PullReplicationMapOutput values.
// You can construct a concrete instance of `PullReplicationMapInput` via:
//
//	PullReplicationMap{ "key": PullReplicationArgs{...} }
type PullReplicationMapInput interface {
	pulumi.Input

	ToPullReplicationMapOutput() PullReplicationMapOutput
	ToPullReplicationMapOutputWithContext(context.Context) PullReplicationMapOutput
}

type PullReplicationMap map[string]PullReplicationInput

func (PullReplicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PullReplication)(nil)).Elem()
}

func (i PullReplicationMap) ToPullReplicationMapOutput() PullReplicationMapOutput {
	return i.ToPullReplicationMapOutputWithContext(context.Background())
}

func (i PullReplicationMap) ToPullReplicationMapOutputWithContext(ctx context.Context) PullReplicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PullReplicationMapOutput)
}

type PullReplicationOutput struct{ *pulumi.OutputState }

func (PullReplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PullReplication)(nil)).Elem()
}

func (o PullReplicationOutput) ToPullReplicationOutput() PullReplicationOutput {
	return o
}

func (o PullReplicationOutput) ToPullReplicationOutputWithContext(ctx context.Context) PullReplicationOutput {
	return o
}

// When true, enables distributed checksum storage. For more information, see
// [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
func (o PullReplicationOutput) CheckBinaryExistenceInFilestore() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PullReplication) pulumi.BoolPtrOutput { return v.CheckBinaryExistenceInFilestore }).(pulumi.BoolPtrOutput)
}

// A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
func (o PullReplicationOutput) CronExp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PullReplication) pulumi.StringPtrOutput { return v.CronExp }).(pulumi.StringPtrOutput)
}

// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. added, deleted or property change.
func (o PullReplicationOutput) EnableEventReplication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PullReplication) pulumi.BoolPtrOutput { return v.EnableEventReplication }).(pulumi.BoolPtrOutput)
}

// When set, this replication will be enabled when saved.
func (o PullReplicationOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *PullReplication) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// Required for local repository, but not needed for remote repository.
func (o PullReplicationOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PullReplication) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// Only artifacts that located in path that matches the subpath within the remote repository will be replicated.
func (o PullReplicationOutput) PathPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PullReplication) pulumi.StringPtrOutput { return v.PathPrefix }).(pulumi.StringPtrOutput)
}

// Proxy key from Artifactory Proxies setting
func (o PullReplicationOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PullReplication) pulumi.StringPtrOutput { return v.Proxy }).(pulumi.StringPtrOutput)
}

// Repository name.
func (o PullReplicationOutput) RepoKey() pulumi.StringOutput {
	return o.ApplyT(func(v *PullReplication) pulumi.StringOutput { return v.RepoKey }).(pulumi.StringOutput)
}

func (o PullReplicationOutput) SocketTimeoutMillis() pulumi.IntOutput {
	return o.ApplyT(func(v *PullReplication) pulumi.IntOutput { return v.SocketTimeoutMillis }).(pulumi.IntOutput)
}

// When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata).
func (o PullReplicationOutput) SyncDeletes() pulumi.BoolOutput {
	return o.ApplyT(func(v *PullReplication) pulumi.BoolOutput { return v.SyncDeletes }).(pulumi.BoolOutput)
}

// When set, the task also synchronizes the properties of replicated artifacts.
func (o PullReplicationOutput) SyncProperties() pulumi.BoolOutput {
	return o.ApplyT(func(v *PullReplication) pulumi.BoolOutput { return v.SyncProperties }).(pulumi.BoolOutput)
}

// When set, artifact download statistics will also be replicated. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery.
func (o PullReplicationOutput) SyncStatistics() pulumi.BoolOutput {
	return o.ApplyT(func(v *PullReplication) pulumi.BoolOutput { return v.SyncStatistics }).(pulumi.BoolOutput)
}

// The URL of the target local repository on a remote Artifactory server. For some package types, you need to prefix the repository key in the URL with api/<pkg>.
// For a list of package types where this is required, see the [note](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-anchorPREFIX).
// Required for local repository, but not needed for remote repository.
func (o PullReplicationOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PullReplication) pulumi.StringPtrOutput { return v.Url }).(pulumi.StringPtrOutput)
}

// Required for local repository, but not needed for remote repository.
func (o PullReplicationOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PullReplication) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

type PullReplicationArrayOutput struct{ *pulumi.OutputState }

func (PullReplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PullReplication)(nil)).Elem()
}

func (o PullReplicationArrayOutput) ToPullReplicationArrayOutput() PullReplicationArrayOutput {
	return o
}

func (o PullReplicationArrayOutput) ToPullReplicationArrayOutputWithContext(ctx context.Context) PullReplicationArrayOutput {
	return o
}

func (o PullReplicationArrayOutput) Index(i pulumi.IntInput) PullReplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PullReplication {
		return vs[0].([]*PullReplication)[vs[1].(int)]
	}).(PullReplicationOutput)
}

type PullReplicationMapOutput struct{ *pulumi.OutputState }

func (PullReplicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PullReplication)(nil)).Elem()
}

func (o PullReplicationMapOutput) ToPullReplicationMapOutput() PullReplicationMapOutput {
	return o
}

func (o PullReplicationMapOutput) ToPullReplicationMapOutputWithContext(ctx context.Context) PullReplicationMapOutput {
	return o
}

func (o PullReplicationMapOutput) MapIndex(k pulumi.StringInput) PullReplicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PullReplication {
		return vs[0].(map[string]*PullReplication)[vs[1].(string)]
	}).(PullReplicationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PullReplicationInput)(nil)).Elem(), &PullReplication{})
	pulumi.RegisterInputType(reflect.TypeOf((*PullReplicationArrayInput)(nil)).Elem(), PullReplicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PullReplicationMapInput)(nil)).Elem(), PullReplicationMap{})
	pulumi.RegisterOutputType(PullReplicationOutput{})
	pulumi.RegisterOutputType(PullReplicationArrayOutput{})
	pulumi.RegisterOutputType(PullReplicationMapOutput{})
}
