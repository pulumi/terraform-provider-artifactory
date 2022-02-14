// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Remote repositories can be imported using their name, e.g.
//
// ```sh
//  $ pulumi import artifactory:index/remoteRepository:RemoteRepository my-remote my-remote
// ```
type RemoteRepository struct {
	pulumi.CustomResourceState

	AllowAnyHostAuth          pulumi.BoolOutput   `pulumi:"allowAnyHostAuth"`
	BlackedOut                pulumi.BoolOutput   `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes pulumi.BoolOutput   `pulumi:"blockMismatchingMimeTypes"`
	BowerRegistryUrl          pulumi.StringOutput `pulumi:"bowerRegistryUrl"`
	BypassHeadRequests        pulumi.BoolOutput   `pulumi:"bypassHeadRequests"`
	ClientTlsCertificate      pulumi.StringOutput `pulumi:"clientTlsCertificate"`
	// Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
	ContentSynchronisation    RemoteRepositoryContentSynchronisationOutput `pulumi:"contentSynchronisation"`
	Description               pulumi.StringPtrOutput                       `pulumi:"description"`
	DownloadContextPath       pulumi.StringPtrOutput                       `pulumi:"downloadContextPath"`
	EnableCookieManagement    pulumi.BoolOutput                            `pulumi:"enableCookieManagement"`
	EnableTokenAuthentication pulumi.BoolOutput                            `pulumi:"enableTokenAuthentication"`
	ExcludesPattern           pulumi.StringOutput                          `pulumi:"excludesPattern"`
	FeedContextPath           pulumi.StringPtrOutput                       `pulumi:"feedContextPath"`
	FetchJarsEagerly          pulumi.BoolOutput                            `pulumi:"fetchJarsEagerly"`
	FetchSourcesEagerly       pulumi.BoolOutput                            `pulumi:"fetchSourcesEagerly"`
	ForceNugetAuthentication  pulumi.BoolOutput                            `pulumi:"forceNugetAuthentication"`
	HandleReleases            pulumi.BoolOutput                            `pulumi:"handleReleases"`
	HandleSnapshots           pulumi.BoolOutput                            `pulumi:"handleSnapshots"`
	HardFail                  pulumi.BoolOutput                            `pulumi:"hardFail"`
	IncludesPattern           pulumi.StringOutput                          `pulumi:"includesPattern"`
	Key                       pulumi.StringOutput                          `pulumi:"key"`
	LocalAddress              pulumi.StringPtrOutput                       `pulumi:"localAddress"`
	MaxUniqueSnapshots        pulumi.IntOutput                             `pulumi:"maxUniqueSnapshots"`
	MissedCachePeriodSeconds  pulumi.IntOutput                             `pulumi:"missedCachePeriodSeconds"`
	Notes                     pulumi.StringPtrOutput                       `pulumi:"notes"`
	Offline                   pulumi.BoolOutput                            `pulumi:"offline"`
	PackageType               pulumi.StringPtrOutput                       `pulumi:"packageType"`
	// Requires password encryption to be turned off `POST /api/system/decrypt`
	Password             pulumi.StringPtrOutput   `pulumi:"password"`
	PropagateQueryParams pulumi.BoolOutput        `pulumi:"propagateQueryParams"`
	PropertySets         pulumi.StringArrayOutput `pulumi:"propertySets"`
	// Proxy key from Artifactory Proxies setting
	Proxy                        pulumi.StringPtrOutput `pulumi:"proxy"`
	PypiRegistryUrl              pulumi.StringOutput    `pulumi:"pypiRegistryUrl"`
	RemoteRepoChecksumPolicyType pulumi.StringOutput    `pulumi:"remoteRepoChecksumPolicyType"`
	RepoLayoutRef                pulumi.StringOutput    `pulumi:"repoLayoutRef"`
	// The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
	RetrievalCachePeriodSeconds       pulumi.IntOutput       `pulumi:"retrievalCachePeriodSeconds"`
	ShareConfiguration                pulumi.BoolOutput      `pulumi:"shareConfiguration"`
	SocketTimeoutMillis               pulumi.IntOutput       `pulumi:"socketTimeoutMillis"`
	StoreArtifactsLocally             pulumi.BoolOutput      `pulumi:"storeArtifactsLocally"`
	SuppressPomConsistencyChecks      pulumi.BoolOutput      `pulumi:"suppressPomConsistencyChecks"`
	SynchronizeProperties             pulumi.BoolOutput      `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodHours pulumi.IntOutput       `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                               pulumi.StringOutput    `pulumi:"url"`
	Username                          pulumi.StringPtrOutput `pulumi:"username"`
	V3FeedUrl                         pulumi.StringPtrOutput `pulumi:"v3FeedUrl"`
	VcsGitDownloadUrl                 pulumi.StringOutput    `pulumi:"vcsGitDownloadUrl"`
	VcsGitProvider                    pulumi.StringOutput    `pulumi:"vcsGitProvider"`
	VcsType                           pulumi.StringOutput    `pulumi:"vcsType"`
	XrayIndex                         pulumi.BoolOutput      `pulumi:"xrayIndex"`
}

// NewRemoteRepository registers a new resource with the given unique name, arguments, and options.
func NewRemoteRepository(ctx *pulumi.Context,
	name string, args *RemoteRepositoryArgs, opts ...pulumi.ResourceOption) (*RemoteRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	var resource RemoteRepository
	err := ctx.RegisterResource("artifactory:index/remoteRepository:RemoteRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRemoteRepository gets an existing RemoteRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRemoteRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemoteRepositoryState, opts ...pulumi.ResourceOption) (*RemoteRepository, error) {
	var resource RemoteRepository
	err := ctx.ReadResource("artifactory:index/remoteRepository:RemoteRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RemoteRepository resources.
type remoteRepositoryState struct {
	AllowAnyHostAuth          *bool   `pulumi:"allowAnyHostAuth"`
	BlackedOut                *bool   `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool   `pulumi:"blockMismatchingMimeTypes"`
	BowerRegistryUrl          *string `pulumi:"bowerRegistryUrl"`
	BypassHeadRequests        *bool   `pulumi:"bypassHeadRequests"`
	ClientTlsCertificate      *string `pulumi:"clientTlsCertificate"`
	// Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
	ContentSynchronisation    *RemoteRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description               *string                                 `pulumi:"description"`
	DownloadContextPath       *string                                 `pulumi:"downloadContextPath"`
	EnableCookieManagement    *bool                                   `pulumi:"enableCookieManagement"`
	EnableTokenAuthentication *bool                                   `pulumi:"enableTokenAuthentication"`
	ExcludesPattern           *string                                 `pulumi:"excludesPattern"`
	FeedContextPath           *string                                 `pulumi:"feedContextPath"`
	FetchJarsEagerly          *bool                                   `pulumi:"fetchJarsEagerly"`
	FetchSourcesEagerly       *bool                                   `pulumi:"fetchSourcesEagerly"`
	ForceNugetAuthentication  *bool                                   `pulumi:"forceNugetAuthentication"`
	HandleReleases            *bool                                   `pulumi:"handleReleases"`
	HandleSnapshots           *bool                                   `pulumi:"handleSnapshots"`
	HardFail                  *bool                                   `pulumi:"hardFail"`
	IncludesPattern           *string                                 `pulumi:"includesPattern"`
	Key                       *string                                 `pulumi:"key"`
	LocalAddress              *string                                 `pulumi:"localAddress"`
	MaxUniqueSnapshots        *int                                    `pulumi:"maxUniqueSnapshots"`
	MissedCachePeriodSeconds  *int                                    `pulumi:"missedCachePeriodSeconds"`
	Notes                     *string                                 `pulumi:"notes"`
	Offline                   *bool                                   `pulumi:"offline"`
	PackageType               *string                                 `pulumi:"packageType"`
	// Requires password encryption to be turned off `POST /api/system/decrypt`
	Password             *string  `pulumi:"password"`
	PropagateQueryParams *bool    `pulumi:"propagateQueryParams"`
	PropertySets         []string `pulumi:"propertySets"`
	// Proxy key from Artifactory Proxies setting
	Proxy                        *string `pulumi:"proxy"`
	PypiRegistryUrl              *string `pulumi:"pypiRegistryUrl"`
	RemoteRepoChecksumPolicyType *string `pulumi:"remoteRepoChecksumPolicyType"`
	RepoLayoutRef                *string `pulumi:"repoLayoutRef"`
	// The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
	RetrievalCachePeriodSeconds       *int    `pulumi:"retrievalCachePeriodSeconds"`
	ShareConfiguration                *bool   `pulumi:"shareConfiguration"`
	SocketTimeoutMillis               *int    `pulumi:"socketTimeoutMillis"`
	StoreArtifactsLocally             *bool   `pulumi:"storeArtifactsLocally"`
	SuppressPomConsistencyChecks      *bool   `pulumi:"suppressPomConsistencyChecks"`
	SynchronizeProperties             *bool   `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodHours *int    `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                               *string `pulumi:"url"`
	Username                          *string `pulumi:"username"`
	V3FeedUrl                         *string `pulumi:"v3FeedUrl"`
	VcsGitDownloadUrl                 *string `pulumi:"vcsGitDownloadUrl"`
	VcsGitProvider                    *string `pulumi:"vcsGitProvider"`
	VcsType                           *string `pulumi:"vcsType"`
	XrayIndex                         *bool   `pulumi:"xrayIndex"`
}

type RemoteRepositoryState struct {
	AllowAnyHostAuth          pulumi.BoolPtrInput
	BlackedOut                pulumi.BoolPtrInput
	BlockMismatchingMimeTypes pulumi.BoolPtrInput
	BowerRegistryUrl          pulumi.StringPtrInput
	BypassHeadRequests        pulumi.BoolPtrInput
	ClientTlsCertificate      pulumi.StringPtrInput
	// Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
	ContentSynchronisation    RemoteRepositoryContentSynchronisationPtrInput
	Description               pulumi.StringPtrInput
	DownloadContextPath       pulumi.StringPtrInput
	EnableCookieManagement    pulumi.BoolPtrInput
	EnableTokenAuthentication pulumi.BoolPtrInput
	ExcludesPattern           pulumi.StringPtrInput
	FeedContextPath           pulumi.StringPtrInput
	FetchJarsEagerly          pulumi.BoolPtrInput
	FetchSourcesEagerly       pulumi.BoolPtrInput
	ForceNugetAuthentication  pulumi.BoolPtrInput
	HandleReleases            pulumi.BoolPtrInput
	HandleSnapshots           pulumi.BoolPtrInput
	HardFail                  pulumi.BoolPtrInput
	IncludesPattern           pulumi.StringPtrInput
	Key                       pulumi.StringPtrInput
	LocalAddress              pulumi.StringPtrInput
	MaxUniqueSnapshots        pulumi.IntPtrInput
	MissedCachePeriodSeconds  pulumi.IntPtrInput
	Notes                     pulumi.StringPtrInput
	Offline                   pulumi.BoolPtrInput
	PackageType               pulumi.StringPtrInput
	// Requires password encryption to be turned off `POST /api/system/decrypt`
	Password             pulumi.StringPtrInput
	PropagateQueryParams pulumi.BoolPtrInput
	PropertySets         pulumi.StringArrayInput
	// Proxy key from Artifactory Proxies setting
	Proxy                        pulumi.StringPtrInput
	PypiRegistryUrl              pulumi.StringPtrInput
	RemoteRepoChecksumPolicyType pulumi.StringPtrInput
	RepoLayoutRef                pulumi.StringPtrInput
	// The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
	RetrievalCachePeriodSeconds       pulumi.IntPtrInput
	ShareConfiguration                pulumi.BoolPtrInput
	SocketTimeoutMillis               pulumi.IntPtrInput
	StoreArtifactsLocally             pulumi.BoolPtrInput
	SuppressPomConsistencyChecks      pulumi.BoolPtrInput
	SynchronizeProperties             pulumi.BoolPtrInput
	UnusedArtifactsCleanupPeriodHours pulumi.IntPtrInput
	Url                               pulumi.StringPtrInput
	Username                          pulumi.StringPtrInput
	V3FeedUrl                         pulumi.StringPtrInput
	VcsGitDownloadUrl                 pulumi.StringPtrInput
	VcsGitProvider                    pulumi.StringPtrInput
	VcsType                           pulumi.StringPtrInput
	XrayIndex                         pulumi.BoolPtrInput
}

func (RemoteRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*remoteRepositoryState)(nil)).Elem()
}

type remoteRepositoryArgs struct {
	AllowAnyHostAuth          *bool   `pulumi:"allowAnyHostAuth"`
	BlackedOut                *bool   `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool   `pulumi:"blockMismatchingMimeTypes"`
	BowerRegistryUrl          *string `pulumi:"bowerRegistryUrl"`
	BypassHeadRequests        *bool   `pulumi:"bypassHeadRequests"`
	ClientTlsCertificate      *string `pulumi:"clientTlsCertificate"`
	// Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
	ContentSynchronisation    *RemoteRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description               *string                                 `pulumi:"description"`
	DownloadContextPath       *string                                 `pulumi:"downloadContextPath"`
	EnableCookieManagement    *bool                                   `pulumi:"enableCookieManagement"`
	EnableTokenAuthentication *bool                                   `pulumi:"enableTokenAuthentication"`
	ExcludesPattern           *string                                 `pulumi:"excludesPattern"`
	FeedContextPath           *string                                 `pulumi:"feedContextPath"`
	FetchJarsEagerly          *bool                                   `pulumi:"fetchJarsEagerly"`
	FetchSourcesEagerly       *bool                                   `pulumi:"fetchSourcesEagerly"`
	ForceNugetAuthentication  *bool                                   `pulumi:"forceNugetAuthentication"`
	HandleReleases            *bool                                   `pulumi:"handleReleases"`
	HandleSnapshots           *bool                                   `pulumi:"handleSnapshots"`
	HardFail                  *bool                                   `pulumi:"hardFail"`
	IncludesPattern           *string                                 `pulumi:"includesPattern"`
	Key                       string                                  `pulumi:"key"`
	LocalAddress              *string                                 `pulumi:"localAddress"`
	MaxUniqueSnapshots        *int                                    `pulumi:"maxUniqueSnapshots"`
	MissedCachePeriodSeconds  *int                                    `pulumi:"missedCachePeriodSeconds"`
	Notes                     *string                                 `pulumi:"notes"`
	Offline                   *bool                                   `pulumi:"offline"`
	PackageType               *string                                 `pulumi:"packageType"`
	// Requires password encryption to be turned off `POST /api/system/decrypt`
	Password             *string  `pulumi:"password"`
	PropagateQueryParams *bool    `pulumi:"propagateQueryParams"`
	PropertySets         []string `pulumi:"propertySets"`
	// Proxy key from Artifactory Proxies setting
	Proxy                        *string `pulumi:"proxy"`
	PypiRegistryUrl              *string `pulumi:"pypiRegistryUrl"`
	RemoteRepoChecksumPolicyType *string `pulumi:"remoteRepoChecksumPolicyType"`
	RepoLayoutRef                *string `pulumi:"repoLayoutRef"`
	// The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
	RetrievalCachePeriodSeconds       *int    `pulumi:"retrievalCachePeriodSeconds"`
	ShareConfiguration                *bool   `pulumi:"shareConfiguration"`
	SocketTimeoutMillis               *int    `pulumi:"socketTimeoutMillis"`
	StoreArtifactsLocally             *bool   `pulumi:"storeArtifactsLocally"`
	SuppressPomConsistencyChecks      *bool   `pulumi:"suppressPomConsistencyChecks"`
	SynchronizeProperties             *bool   `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodHours *int    `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                               string  `pulumi:"url"`
	Username                          *string `pulumi:"username"`
	V3FeedUrl                         *string `pulumi:"v3FeedUrl"`
	VcsGitDownloadUrl                 *string `pulumi:"vcsGitDownloadUrl"`
	VcsGitProvider                    *string `pulumi:"vcsGitProvider"`
	VcsType                           *string `pulumi:"vcsType"`
	XrayIndex                         *bool   `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a RemoteRepository resource.
type RemoteRepositoryArgs struct {
	AllowAnyHostAuth          pulumi.BoolPtrInput
	BlackedOut                pulumi.BoolPtrInput
	BlockMismatchingMimeTypes pulumi.BoolPtrInput
	BowerRegistryUrl          pulumi.StringPtrInput
	BypassHeadRequests        pulumi.BoolPtrInput
	ClientTlsCertificate      pulumi.StringPtrInput
	// Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
	ContentSynchronisation    RemoteRepositoryContentSynchronisationPtrInput
	Description               pulumi.StringPtrInput
	DownloadContextPath       pulumi.StringPtrInput
	EnableCookieManagement    pulumi.BoolPtrInput
	EnableTokenAuthentication pulumi.BoolPtrInput
	ExcludesPattern           pulumi.StringPtrInput
	FeedContextPath           pulumi.StringPtrInput
	FetchJarsEagerly          pulumi.BoolPtrInput
	FetchSourcesEagerly       pulumi.BoolPtrInput
	ForceNugetAuthentication  pulumi.BoolPtrInput
	HandleReleases            pulumi.BoolPtrInput
	HandleSnapshots           pulumi.BoolPtrInput
	HardFail                  pulumi.BoolPtrInput
	IncludesPattern           pulumi.StringPtrInput
	Key                       pulumi.StringInput
	LocalAddress              pulumi.StringPtrInput
	MaxUniqueSnapshots        pulumi.IntPtrInput
	MissedCachePeriodSeconds  pulumi.IntPtrInput
	Notes                     pulumi.StringPtrInput
	Offline                   pulumi.BoolPtrInput
	PackageType               pulumi.StringPtrInput
	// Requires password encryption to be turned off `POST /api/system/decrypt`
	Password             pulumi.StringPtrInput
	PropagateQueryParams pulumi.BoolPtrInput
	PropertySets         pulumi.StringArrayInput
	// Proxy key from Artifactory Proxies setting
	Proxy                        pulumi.StringPtrInput
	PypiRegistryUrl              pulumi.StringPtrInput
	RemoteRepoChecksumPolicyType pulumi.StringPtrInput
	RepoLayoutRef                pulumi.StringPtrInput
	// The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
	RetrievalCachePeriodSeconds       pulumi.IntPtrInput
	ShareConfiguration                pulumi.BoolPtrInput
	SocketTimeoutMillis               pulumi.IntPtrInput
	StoreArtifactsLocally             pulumi.BoolPtrInput
	SuppressPomConsistencyChecks      pulumi.BoolPtrInput
	SynchronizeProperties             pulumi.BoolPtrInput
	UnusedArtifactsCleanupPeriodHours pulumi.IntPtrInput
	Url                               pulumi.StringInput
	Username                          pulumi.StringPtrInput
	V3FeedUrl                         pulumi.StringPtrInput
	VcsGitDownloadUrl                 pulumi.StringPtrInput
	VcsGitProvider                    pulumi.StringPtrInput
	VcsType                           pulumi.StringPtrInput
	XrayIndex                         pulumi.BoolPtrInput
}

func (RemoteRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*remoteRepositoryArgs)(nil)).Elem()
}

type RemoteRepositoryInput interface {
	pulumi.Input

	ToRemoteRepositoryOutput() RemoteRepositoryOutput
	ToRemoteRepositoryOutputWithContext(ctx context.Context) RemoteRepositoryOutput
}

func (*RemoteRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**RemoteRepository)(nil)).Elem()
}

func (i *RemoteRepository) ToRemoteRepositoryOutput() RemoteRepositoryOutput {
	return i.ToRemoteRepositoryOutputWithContext(context.Background())
}

func (i *RemoteRepository) ToRemoteRepositoryOutputWithContext(ctx context.Context) RemoteRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteRepositoryOutput)
}

// RemoteRepositoryArrayInput is an input type that accepts RemoteRepositoryArray and RemoteRepositoryArrayOutput values.
// You can construct a concrete instance of `RemoteRepositoryArrayInput` via:
//
//          RemoteRepositoryArray{ RemoteRepositoryArgs{...} }
type RemoteRepositoryArrayInput interface {
	pulumi.Input

	ToRemoteRepositoryArrayOutput() RemoteRepositoryArrayOutput
	ToRemoteRepositoryArrayOutputWithContext(context.Context) RemoteRepositoryArrayOutput
}

type RemoteRepositoryArray []RemoteRepositoryInput

func (RemoteRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RemoteRepository)(nil)).Elem()
}

func (i RemoteRepositoryArray) ToRemoteRepositoryArrayOutput() RemoteRepositoryArrayOutput {
	return i.ToRemoteRepositoryArrayOutputWithContext(context.Background())
}

func (i RemoteRepositoryArray) ToRemoteRepositoryArrayOutputWithContext(ctx context.Context) RemoteRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteRepositoryArrayOutput)
}

// RemoteRepositoryMapInput is an input type that accepts RemoteRepositoryMap and RemoteRepositoryMapOutput values.
// You can construct a concrete instance of `RemoteRepositoryMapInput` via:
//
//          RemoteRepositoryMap{ "key": RemoteRepositoryArgs{...} }
type RemoteRepositoryMapInput interface {
	pulumi.Input

	ToRemoteRepositoryMapOutput() RemoteRepositoryMapOutput
	ToRemoteRepositoryMapOutputWithContext(context.Context) RemoteRepositoryMapOutput
}

type RemoteRepositoryMap map[string]RemoteRepositoryInput

func (RemoteRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RemoteRepository)(nil)).Elem()
}

func (i RemoteRepositoryMap) ToRemoteRepositoryMapOutput() RemoteRepositoryMapOutput {
	return i.ToRemoteRepositoryMapOutputWithContext(context.Background())
}

func (i RemoteRepositoryMap) ToRemoteRepositoryMapOutputWithContext(ctx context.Context) RemoteRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteRepositoryMapOutput)
}

type RemoteRepositoryOutput struct{ *pulumi.OutputState }

func (RemoteRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemoteRepository)(nil)).Elem()
}

func (o RemoteRepositoryOutput) ToRemoteRepositoryOutput() RemoteRepositoryOutput {
	return o
}

func (o RemoteRepositoryOutput) ToRemoteRepositoryOutputWithContext(ctx context.Context) RemoteRepositoryOutput {
	return o
}

type RemoteRepositoryArrayOutput struct{ *pulumi.OutputState }

func (RemoteRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RemoteRepository)(nil)).Elem()
}

func (o RemoteRepositoryArrayOutput) ToRemoteRepositoryArrayOutput() RemoteRepositoryArrayOutput {
	return o
}

func (o RemoteRepositoryArrayOutput) ToRemoteRepositoryArrayOutputWithContext(ctx context.Context) RemoteRepositoryArrayOutput {
	return o
}

func (o RemoteRepositoryArrayOutput) Index(i pulumi.IntInput) RemoteRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RemoteRepository {
		return vs[0].([]*RemoteRepository)[vs[1].(int)]
	}).(RemoteRepositoryOutput)
}

type RemoteRepositoryMapOutput struct{ *pulumi.OutputState }

func (RemoteRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RemoteRepository)(nil)).Elem()
}

func (o RemoteRepositoryMapOutput) ToRemoteRepositoryMapOutput() RemoteRepositoryMapOutput {
	return o
}

func (o RemoteRepositoryMapOutput) ToRemoteRepositoryMapOutputWithContext(ctx context.Context) RemoteRepositoryMapOutput {
	return o
}

func (o RemoteRepositoryMapOutput) MapIndex(k pulumi.StringInput) RemoteRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RemoteRepository {
		return vs[0].(map[string]*RemoteRepository)[vs[1].(string)]
	}).(RemoteRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteRepositoryInput)(nil)).Elem(), &RemoteRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteRepositoryArrayInput)(nil)).Elem(), RemoteRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteRepositoryMapInput)(nil)).Elem(), RemoteRepositoryMap{})
	pulumi.RegisterOutputType(RemoteRepositoryOutput{})
	pulumi.RegisterOutputType(RemoteRepositoryArrayOutput{})
	pulumi.RegisterOutputType(RemoteRepositoryMapOutput{})
}
