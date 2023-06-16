// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a local repository replication resource, also referred to as Artifactory push replication. This can be used to create and manage Artifactory local repository replications using [Push Replication API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-SetRepositoryReplicationConfiguration).
// Push replication is used to synchronize Local Repositories, and is implemented by the Artifactory server on the near end invoking a synchronization of artifacts to the far end.
// See the [Official Documentation](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-PushReplication).
// This resource can create the replication of local repository to single repository on the remote server.
//
// ## Import
//
// Push replication configs can be imported using their repo key, e.g.
//
// ```sh
//
//	$ pulumi import artifactory:index/localRepositorySingleReplication:LocalRepositorySingleReplication foo-rep provider_test_source
//
// ```
type LocalRepositorySingleReplication struct {
	pulumi.CustomResourceState

	// Enabling the `checkBinaryExistenceInFilestore` flag requires an Enterprise Plus license. When true, enables distributed checksum storage. For more information, see [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
	CheckBinaryExistenceInFilestore pulumi.BoolPtrOutput `pulumi:"checkBinaryExistenceInFilestore"`
	// A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
	CronExp pulumi.StringPtrOutput `pulumi:"cronExp"`
	// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is `false`.
	EnableEventReplication pulumi.BoolPtrOutput `pulumi:"enableEventReplication"`
	// When set, enables replication of this repository to the target specified in `url` attribute. Default value is `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`. By default, no artifacts are excluded.
	ExcludePathPrefixPattern pulumi.StringPtrOutput `pulumi:"excludePathPrefixPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included `(**/*)`.
	IncludePathPrefixPattern pulumi.StringPtrOutput `pulumi:"includePathPrefixPattern"`
	// Use either the HTTP authentication password or [identity token](https://www.jfrog.com/confluence/display/JFROG/User+Profile#UserProfile-IdentityTokenidentitytoken).
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Proxy key from Artifactory Proxies settings. The proxy configuration will be used when communicating with the remote instance.
	Proxy pulumi.StringPtrOutput `pulumi:"proxy"`
	// Replication ID, the value is unknown until the resource is created. Can't be set or updated.
	ReplicationKey pulumi.StringOutput `pulumi:"replicationKey"`
	// Repository name.
	RepoKey pulumi.StringOutput `pulumi:"repoKey"`
	// The network timeout in milliseconds to use for remote operations. Default value is `15000`.
	SocketTimeoutMillis pulumi.IntPtrOutput `pulumi:"socketTimeoutMillis"`
	// When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value is `false`.
	SyncDeletes pulumi.BoolPtrOutput `pulumi:"syncDeletes"`
	// When set, the task also synchronizes the properties of replicated artifacts. Default value is `true`.
	SyncProperties pulumi.BoolPtrOutput `pulumi:"syncProperties"`
	// When set, the task also synchronizes artifact download statistics. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery. Default value is `false`
	SyncStatistics pulumi.BoolPtrOutput `pulumi:"syncStatistics"`
	// The URL of the target local repository on a remote Artifactory server. Use the format `https://<artifactory_url>/artifactory/<repository_name>`.
	Url pulumi.StringOutput `pulumi:"url"`
	// Username on the remote Artifactory instance.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewLocalRepositorySingleReplication registers a new resource with the given unique name, arguments, and options.
func NewLocalRepositorySingleReplication(ctx *pulumi.Context,
	name string, args *LocalRepositorySingleReplicationArgs, opts ...pulumi.ResourceOption) (*LocalRepositorySingleReplication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RepoKey == nil {
		return nil, errors.New("invalid value for required argument 'RepoKey'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	var resource LocalRepositorySingleReplication
	err := ctx.RegisterResource("artifactory:index/localRepositorySingleReplication:LocalRepositorySingleReplication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalRepositorySingleReplication gets an existing LocalRepositorySingleReplication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalRepositorySingleReplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalRepositorySingleReplicationState, opts ...pulumi.ResourceOption) (*LocalRepositorySingleReplication, error) {
	var resource LocalRepositorySingleReplication
	err := ctx.ReadResource("artifactory:index/localRepositorySingleReplication:LocalRepositorySingleReplication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalRepositorySingleReplication resources.
type localRepositorySingleReplicationState struct {
	// Enabling the `checkBinaryExistenceInFilestore` flag requires an Enterprise Plus license. When true, enables distributed checksum storage. For more information, see [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
	CheckBinaryExistenceInFilestore *bool `pulumi:"checkBinaryExistenceInFilestore"`
	// A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
	CronExp *string `pulumi:"cronExp"`
	// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is `false`.
	EnableEventReplication *bool `pulumi:"enableEventReplication"`
	// When set, enables replication of this repository to the target specified in `url` attribute. Default value is `true`.
	Enabled *bool `pulumi:"enabled"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`. By default, no artifacts are excluded.
	ExcludePathPrefixPattern *string `pulumi:"excludePathPrefixPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included `(**/*)`.
	IncludePathPrefixPattern *string `pulumi:"includePathPrefixPattern"`
	// Use either the HTTP authentication password or [identity token](https://www.jfrog.com/confluence/display/JFROG/User+Profile#UserProfile-IdentityTokenidentitytoken).
	Password *string `pulumi:"password"`
	// Proxy key from Artifactory Proxies settings. The proxy configuration will be used when communicating with the remote instance.
	Proxy *string `pulumi:"proxy"`
	// Replication ID, the value is unknown until the resource is created. Can't be set or updated.
	ReplicationKey *string `pulumi:"replicationKey"`
	// Repository name.
	RepoKey *string `pulumi:"repoKey"`
	// The network timeout in milliseconds to use for remote operations. Default value is `15000`.
	SocketTimeoutMillis *int `pulumi:"socketTimeoutMillis"`
	// When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value is `false`.
	SyncDeletes *bool `pulumi:"syncDeletes"`
	// When set, the task also synchronizes the properties of replicated artifacts. Default value is `true`.
	SyncProperties *bool `pulumi:"syncProperties"`
	// When set, the task also synchronizes artifact download statistics. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery. Default value is `false`
	SyncStatistics *bool `pulumi:"syncStatistics"`
	// The URL of the target local repository on a remote Artifactory server. Use the format `https://<artifactory_url>/artifactory/<repository_name>`.
	Url *string `pulumi:"url"`
	// Username on the remote Artifactory instance.
	Username *string `pulumi:"username"`
}

type LocalRepositorySingleReplicationState struct {
	// Enabling the `checkBinaryExistenceInFilestore` flag requires an Enterprise Plus license. When true, enables distributed checksum storage. For more information, see [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
	CheckBinaryExistenceInFilestore pulumi.BoolPtrInput
	// A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
	CronExp pulumi.StringPtrInput
	// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is `false`.
	EnableEventReplication pulumi.BoolPtrInput
	// When set, enables replication of this repository to the target specified in `url` attribute. Default value is `true`.
	Enabled pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`. By default, no artifacts are excluded.
	ExcludePathPrefixPattern pulumi.StringPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included `(**/*)`.
	IncludePathPrefixPattern pulumi.StringPtrInput
	// Use either the HTTP authentication password or [identity token](https://www.jfrog.com/confluence/display/JFROG/User+Profile#UserProfile-IdentityTokenidentitytoken).
	Password pulumi.StringPtrInput
	// Proxy key from Artifactory Proxies settings. The proxy configuration will be used when communicating with the remote instance.
	Proxy pulumi.StringPtrInput
	// Replication ID, the value is unknown until the resource is created. Can't be set or updated.
	ReplicationKey pulumi.StringPtrInput
	// Repository name.
	RepoKey pulumi.StringPtrInput
	// The network timeout in milliseconds to use for remote operations. Default value is `15000`.
	SocketTimeoutMillis pulumi.IntPtrInput
	// When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value is `false`.
	SyncDeletes pulumi.BoolPtrInput
	// When set, the task also synchronizes the properties of replicated artifacts. Default value is `true`.
	SyncProperties pulumi.BoolPtrInput
	// When set, the task also synchronizes artifact download statistics. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery. Default value is `false`
	SyncStatistics pulumi.BoolPtrInput
	// The URL of the target local repository on a remote Artifactory server. Use the format `https://<artifactory_url>/artifactory/<repository_name>`.
	Url pulumi.StringPtrInput
	// Username on the remote Artifactory instance.
	Username pulumi.StringPtrInput
}

func (LocalRepositorySingleReplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*localRepositorySingleReplicationState)(nil)).Elem()
}

type localRepositorySingleReplicationArgs struct {
	// Enabling the `checkBinaryExistenceInFilestore` flag requires an Enterprise Plus license. When true, enables distributed checksum storage. For more information, see [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
	CheckBinaryExistenceInFilestore *bool `pulumi:"checkBinaryExistenceInFilestore"`
	// A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
	CronExp *string `pulumi:"cronExp"`
	// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is `false`.
	EnableEventReplication *bool `pulumi:"enableEventReplication"`
	// When set, enables replication of this repository to the target specified in `url` attribute. Default value is `true`.
	Enabled *bool `pulumi:"enabled"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`. By default, no artifacts are excluded.
	ExcludePathPrefixPattern *string `pulumi:"excludePathPrefixPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included `(**/*)`.
	IncludePathPrefixPattern *string `pulumi:"includePathPrefixPattern"`
	// Use either the HTTP authentication password or [identity token](https://www.jfrog.com/confluence/display/JFROG/User+Profile#UserProfile-IdentityTokenidentitytoken).
	Password *string `pulumi:"password"`
	// Proxy key from Artifactory Proxies settings. The proxy configuration will be used when communicating with the remote instance.
	Proxy *string `pulumi:"proxy"`
	// Replication ID, the value is unknown until the resource is created. Can't be set or updated.
	ReplicationKey *string `pulumi:"replicationKey"`
	// Repository name.
	RepoKey string `pulumi:"repoKey"`
	// The network timeout in milliseconds to use for remote operations. Default value is `15000`.
	SocketTimeoutMillis *int `pulumi:"socketTimeoutMillis"`
	// When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value is `false`.
	SyncDeletes *bool `pulumi:"syncDeletes"`
	// When set, the task also synchronizes the properties of replicated artifacts. Default value is `true`.
	SyncProperties *bool `pulumi:"syncProperties"`
	// When set, the task also synchronizes artifact download statistics. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery. Default value is `false`
	SyncStatistics *bool `pulumi:"syncStatistics"`
	// The URL of the target local repository on a remote Artifactory server. Use the format `https://<artifactory_url>/artifactory/<repository_name>`.
	Url string `pulumi:"url"`
	// Username on the remote Artifactory instance.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a LocalRepositorySingleReplication resource.
type LocalRepositorySingleReplicationArgs struct {
	// Enabling the `checkBinaryExistenceInFilestore` flag requires an Enterprise Plus license. When true, enables distributed checksum storage. For more information, see [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
	CheckBinaryExistenceInFilestore pulumi.BoolPtrInput
	// A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
	CronExp pulumi.StringPtrInput
	// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is `false`.
	EnableEventReplication pulumi.BoolPtrInput
	// When set, enables replication of this repository to the target specified in `url` attribute. Default value is `true`.
	Enabled pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`. By default, no artifacts are excluded.
	ExcludePathPrefixPattern pulumi.StringPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included `(**/*)`.
	IncludePathPrefixPattern pulumi.StringPtrInput
	// Use either the HTTP authentication password or [identity token](https://www.jfrog.com/confluence/display/JFROG/User+Profile#UserProfile-IdentityTokenidentitytoken).
	Password pulumi.StringPtrInput
	// Proxy key from Artifactory Proxies settings. The proxy configuration will be used when communicating with the remote instance.
	Proxy pulumi.StringPtrInput
	// Replication ID, the value is unknown until the resource is created. Can't be set or updated.
	ReplicationKey pulumi.StringPtrInput
	// Repository name.
	RepoKey pulumi.StringInput
	// The network timeout in milliseconds to use for remote operations. Default value is `15000`.
	SocketTimeoutMillis pulumi.IntPtrInput
	// When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value is `false`.
	SyncDeletes pulumi.BoolPtrInput
	// When set, the task also synchronizes the properties of replicated artifacts. Default value is `true`.
	SyncProperties pulumi.BoolPtrInput
	// When set, the task also synchronizes artifact download statistics. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery. Default value is `false`
	SyncStatistics pulumi.BoolPtrInput
	// The URL of the target local repository on a remote Artifactory server. Use the format `https://<artifactory_url>/artifactory/<repository_name>`.
	Url pulumi.StringInput
	// Username on the remote Artifactory instance.
	Username pulumi.StringInput
}

func (LocalRepositorySingleReplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localRepositorySingleReplicationArgs)(nil)).Elem()
}

type LocalRepositorySingleReplicationInput interface {
	pulumi.Input

	ToLocalRepositorySingleReplicationOutput() LocalRepositorySingleReplicationOutput
	ToLocalRepositorySingleReplicationOutputWithContext(ctx context.Context) LocalRepositorySingleReplicationOutput
}

func (*LocalRepositorySingleReplication) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalRepositorySingleReplication)(nil)).Elem()
}

func (i *LocalRepositorySingleReplication) ToLocalRepositorySingleReplicationOutput() LocalRepositorySingleReplicationOutput {
	return i.ToLocalRepositorySingleReplicationOutputWithContext(context.Background())
}

func (i *LocalRepositorySingleReplication) ToLocalRepositorySingleReplicationOutputWithContext(ctx context.Context) LocalRepositorySingleReplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalRepositorySingleReplicationOutput)
}

// LocalRepositorySingleReplicationArrayInput is an input type that accepts LocalRepositorySingleReplicationArray and LocalRepositorySingleReplicationArrayOutput values.
// You can construct a concrete instance of `LocalRepositorySingleReplicationArrayInput` via:
//
//	LocalRepositorySingleReplicationArray{ LocalRepositorySingleReplicationArgs{...} }
type LocalRepositorySingleReplicationArrayInput interface {
	pulumi.Input

	ToLocalRepositorySingleReplicationArrayOutput() LocalRepositorySingleReplicationArrayOutput
	ToLocalRepositorySingleReplicationArrayOutputWithContext(context.Context) LocalRepositorySingleReplicationArrayOutput
}

type LocalRepositorySingleReplicationArray []LocalRepositorySingleReplicationInput

func (LocalRepositorySingleReplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalRepositorySingleReplication)(nil)).Elem()
}

func (i LocalRepositorySingleReplicationArray) ToLocalRepositorySingleReplicationArrayOutput() LocalRepositorySingleReplicationArrayOutput {
	return i.ToLocalRepositorySingleReplicationArrayOutputWithContext(context.Background())
}

func (i LocalRepositorySingleReplicationArray) ToLocalRepositorySingleReplicationArrayOutputWithContext(ctx context.Context) LocalRepositorySingleReplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalRepositorySingleReplicationArrayOutput)
}

// LocalRepositorySingleReplicationMapInput is an input type that accepts LocalRepositorySingleReplicationMap and LocalRepositorySingleReplicationMapOutput values.
// You can construct a concrete instance of `LocalRepositorySingleReplicationMapInput` via:
//
//	LocalRepositorySingleReplicationMap{ "key": LocalRepositorySingleReplicationArgs{...} }
type LocalRepositorySingleReplicationMapInput interface {
	pulumi.Input

	ToLocalRepositorySingleReplicationMapOutput() LocalRepositorySingleReplicationMapOutput
	ToLocalRepositorySingleReplicationMapOutputWithContext(context.Context) LocalRepositorySingleReplicationMapOutput
}

type LocalRepositorySingleReplicationMap map[string]LocalRepositorySingleReplicationInput

func (LocalRepositorySingleReplicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalRepositorySingleReplication)(nil)).Elem()
}

func (i LocalRepositorySingleReplicationMap) ToLocalRepositorySingleReplicationMapOutput() LocalRepositorySingleReplicationMapOutput {
	return i.ToLocalRepositorySingleReplicationMapOutputWithContext(context.Background())
}

func (i LocalRepositorySingleReplicationMap) ToLocalRepositorySingleReplicationMapOutputWithContext(ctx context.Context) LocalRepositorySingleReplicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalRepositorySingleReplicationMapOutput)
}

type LocalRepositorySingleReplicationOutput struct{ *pulumi.OutputState }

func (LocalRepositorySingleReplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalRepositorySingleReplication)(nil)).Elem()
}

func (o LocalRepositorySingleReplicationOutput) ToLocalRepositorySingleReplicationOutput() LocalRepositorySingleReplicationOutput {
	return o
}

func (o LocalRepositorySingleReplicationOutput) ToLocalRepositorySingleReplicationOutputWithContext(ctx context.Context) LocalRepositorySingleReplicationOutput {
	return o
}

// Enabling the `checkBinaryExistenceInFilestore` flag requires an Enterprise Plus license. When true, enables distributed checksum storage. For more information, see [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
func (o LocalRepositorySingleReplicationOutput) CheckBinaryExistenceInFilestore() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalRepositorySingleReplication) pulumi.BoolPtrOutput {
		return v.CheckBinaryExistenceInFilestore
	}).(pulumi.BoolPtrOutput)
}

// A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
func (o LocalRepositorySingleReplicationOutput) CronExp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalRepositorySingleReplication) pulumi.StringPtrOutput { return v.CronExp }).(pulumi.StringPtrOutput)
}

// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is `false`.
func (o LocalRepositorySingleReplicationOutput) EnableEventReplication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalRepositorySingleReplication) pulumi.BoolPtrOutput { return v.EnableEventReplication }).(pulumi.BoolPtrOutput)
}

// When set, enables replication of this repository to the target specified in `url` attribute. Default value is `true`.
func (o LocalRepositorySingleReplicationOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalRepositorySingleReplication) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`. By default, no artifacts are excluded.
func (o LocalRepositorySingleReplicationOutput) ExcludePathPrefixPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalRepositorySingleReplication) pulumi.StringPtrOutput { return v.ExcludePathPrefixPattern }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included `(**/*)`.
func (o LocalRepositorySingleReplicationOutput) IncludePathPrefixPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalRepositorySingleReplication) pulumi.StringPtrOutput { return v.IncludePathPrefixPattern }).(pulumi.StringPtrOutput)
}

// Use either the HTTP authentication password or [identity token](https://www.jfrog.com/confluence/display/JFROG/User+Profile#UserProfile-IdentityTokenidentitytoken).
func (o LocalRepositorySingleReplicationOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalRepositorySingleReplication) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// Proxy key from Artifactory Proxies settings. The proxy configuration will be used when communicating with the remote instance.
func (o LocalRepositorySingleReplicationOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalRepositorySingleReplication) pulumi.StringPtrOutput { return v.Proxy }).(pulumi.StringPtrOutput)
}

// Replication ID, the value is unknown until the resource is created. Can't be set or updated.
func (o LocalRepositorySingleReplicationOutput) ReplicationKey() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalRepositorySingleReplication) pulumi.StringOutput { return v.ReplicationKey }).(pulumi.StringOutput)
}

// Repository name.
func (o LocalRepositorySingleReplicationOutput) RepoKey() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalRepositorySingleReplication) pulumi.StringOutput { return v.RepoKey }).(pulumi.StringOutput)
}

// The network timeout in milliseconds to use for remote operations. Default value is `15000`.
func (o LocalRepositorySingleReplicationOutput) SocketTimeoutMillis() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LocalRepositorySingleReplication) pulumi.IntPtrOutput { return v.SocketTimeoutMillis }).(pulumi.IntPtrOutput)
}

// When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value is `false`.
func (o LocalRepositorySingleReplicationOutput) SyncDeletes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalRepositorySingleReplication) pulumi.BoolPtrOutput { return v.SyncDeletes }).(pulumi.BoolPtrOutput)
}

// When set, the task also synchronizes the properties of replicated artifacts. Default value is `true`.
func (o LocalRepositorySingleReplicationOutput) SyncProperties() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalRepositorySingleReplication) pulumi.BoolPtrOutput { return v.SyncProperties }).(pulumi.BoolPtrOutput)
}

// When set, the task also synchronizes artifact download statistics. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery. Default value is `false`
func (o LocalRepositorySingleReplicationOutput) SyncStatistics() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalRepositorySingleReplication) pulumi.BoolPtrOutput { return v.SyncStatistics }).(pulumi.BoolPtrOutput)
}

// The URL of the target local repository on a remote Artifactory server. Use the format `https://<artifactory_url>/artifactory/<repository_name>`.
func (o LocalRepositorySingleReplicationOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalRepositorySingleReplication) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// Username on the remote Artifactory instance.
func (o LocalRepositorySingleReplicationOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalRepositorySingleReplication) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type LocalRepositorySingleReplicationArrayOutput struct{ *pulumi.OutputState }

func (LocalRepositorySingleReplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalRepositorySingleReplication)(nil)).Elem()
}

func (o LocalRepositorySingleReplicationArrayOutput) ToLocalRepositorySingleReplicationArrayOutput() LocalRepositorySingleReplicationArrayOutput {
	return o
}

func (o LocalRepositorySingleReplicationArrayOutput) ToLocalRepositorySingleReplicationArrayOutputWithContext(ctx context.Context) LocalRepositorySingleReplicationArrayOutput {
	return o
}

func (o LocalRepositorySingleReplicationArrayOutput) Index(i pulumi.IntInput) LocalRepositorySingleReplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalRepositorySingleReplication {
		return vs[0].([]*LocalRepositorySingleReplication)[vs[1].(int)]
	}).(LocalRepositorySingleReplicationOutput)
}

type LocalRepositorySingleReplicationMapOutput struct{ *pulumi.OutputState }

func (LocalRepositorySingleReplicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalRepositorySingleReplication)(nil)).Elem()
}

func (o LocalRepositorySingleReplicationMapOutput) ToLocalRepositorySingleReplicationMapOutput() LocalRepositorySingleReplicationMapOutput {
	return o
}

func (o LocalRepositorySingleReplicationMapOutput) ToLocalRepositorySingleReplicationMapOutputWithContext(ctx context.Context) LocalRepositorySingleReplicationMapOutput {
	return o
}

func (o LocalRepositorySingleReplicationMapOutput) MapIndex(k pulumi.StringInput) LocalRepositorySingleReplicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalRepositorySingleReplication {
		return vs[0].(map[string]*LocalRepositorySingleReplication)[vs[1].(string)]
	}).(LocalRepositorySingleReplicationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalRepositorySingleReplicationInput)(nil)).Elem(), &LocalRepositorySingleReplication{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalRepositorySingleReplicationArrayInput)(nil)).Elem(), LocalRepositorySingleReplicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalRepositorySingleReplicationMapInput)(nil)).Elem(), LocalRepositorySingleReplicationMap{})
	pulumi.RegisterOutputType(LocalRepositorySingleReplicationOutput{})
	pulumi.RegisterOutputType(LocalRepositorySingleReplicationArrayOutput{})
	pulumi.RegisterOutputType(LocalRepositorySingleReplicationMapOutput{})
}
