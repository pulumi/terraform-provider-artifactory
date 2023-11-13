// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a remote Generic repository.
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
//			_, err := artifactory.LookupRemoteGenericRepository(ctx, &artifactory.LookupRemoteGenericRepositoryArgs{
//				Key: "remote-generic",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRemoteGenericRepository(ctx *pulumi.Context, args *LookupRemoteGenericRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupRemoteGenericRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRemoteGenericRepositoryResult
	err := ctx.Invoke("artifactory:index/getRemoteGenericRepository:getRemoteGenericRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRemoteGenericRepository.
type LookupRemoteGenericRepositoryArgs struct {
	AllowAnyHostAuth          *bool                                             `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  *int                                              `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool                                             `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool                                             `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool                                             `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool                                             `pulumi:"cdnRedirect"`
	ClientTlsCertificate      *string                                           `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    *GetRemoteGenericRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description               *string                                           `pulumi:"description"`
	DisableProxy              *bool                                             `pulumi:"disableProxy"`
	DisableUrlNormalization   *bool                                             `pulumi:"disableUrlNormalization"`
	DownloadDirect            *bool                                             `pulumi:"downloadDirect"`
	EnableCookieManagement    *bool                                             `pulumi:"enableCookieManagement"`
	ExcludesPattern           *string                                           `pulumi:"excludesPattern"`
	HardFail                  *bool                                             `pulumi:"hardFail"`
	IncludesPattern           *string                                           `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                              string   `pulumi:"key"`
	ListRemoteFolderItems            *bool    `pulumi:"listRemoteFolderItems"`
	LocalAddress                     *string  `pulumi:"localAddress"`
	MetadataRetrievalTimeoutSecs     *int     `pulumi:"metadataRetrievalTimeoutSecs"`
	MismatchingMimeTypesOverrideList *string  `pulumi:"mismatchingMimeTypesOverrideList"`
	MissedCachePeriodSeconds         *int     `pulumi:"missedCachePeriodSeconds"`
	Notes                            *string  `pulumi:"notes"`
	Offline                          *bool    `pulumi:"offline"`
	Password                         *string  `pulumi:"password"`
	PriorityResolution               *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments              []string `pulumi:"projectEnvironments"`
	ProjectKey                       *string  `pulumi:"projectKey"`
	// (Optional, Default: `false`) When set, if query params are included in the request to Artifactory, they will be passed on to the remote repository.
	PropagateQueryParams              *bool    `pulumi:"propagateQueryParams"`
	PropertySets                      []string `pulumi:"propertySets"`
	Proxy                             *string  `pulumi:"proxy"`
	QueryParams                       *string  `pulumi:"queryParams"`
	RemoteRepoLayoutRef               *string  `pulumi:"remoteRepoLayoutRef"`
	RepoLayoutRef                     *string  `pulumi:"repoLayoutRef"`
	RetrievalCachePeriodSeconds       *int     `pulumi:"retrievalCachePeriodSeconds"`
	ShareConfiguration                *bool    `pulumi:"shareConfiguration"`
	SocketTimeoutMillis               *int     `pulumi:"socketTimeoutMillis"`
	StoreArtifactsLocally             *bool    `pulumi:"storeArtifactsLocally"`
	SynchronizeProperties             *bool    `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodHours *int     `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                               *string  `pulumi:"url"`
	Username                          *string  `pulumi:"username"`
	XrayIndex                         *bool    `pulumi:"xrayIndex"`
}

// A collection of values returned by getRemoteGenericRepository.
type LookupRemoteGenericRepositoryResult struct {
	AllowAnyHostAuth          *bool                                            `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  *int                                             `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool                                            `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool                                            `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool                                            `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool                                            `pulumi:"cdnRedirect"`
	ClientTlsCertificate      string                                           `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    GetRemoteGenericRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description               *string                                          `pulumi:"description"`
	DisableProxy              *bool                                            `pulumi:"disableProxy"`
	DisableUrlNormalization   *bool                                            `pulumi:"disableUrlNormalization"`
	DownloadDirect            *bool                                            `pulumi:"downloadDirect"`
	EnableCookieManagement    *bool                                            `pulumi:"enableCookieManagement"`
	ExcludesPattern           *string                                          `pulumi:"excludesPattern"`
	HardFail                  *bool                                            `pulumi:"hardFail"`
	// The provider-assigned unique ID for this managed resource.
	Id                               string   `pulumi:"id"`
	IncludesPattern                  *string  `pulumi:"includesPattern"`
	Key                              string   `pulumi:"key"`
	ListRemoteFolderItems            *bool    `pulumi:"listRemoteFolderItems"`
	LocalAddress                     *string  `pulumi:"localAddress"`
	MetadataRetrievalTimeoutSecs     *int     `pulumi:"metadataRetrievalTimeoutSecs"`
	MismatchingMimeTypesOverrideList *string  `pulumi:"mismatchingMimeTypesOverrideList"`
	MissedCachePeriodSeconds         *int     `pulumi:"missedCachePeriodSeconds"`
	Notes                            *string  `pulumi:"notes"`
	Offline                          *bool    `pulumi:"offline"`
	PackageType                      string   `pulumi:"packageType"`
	Password                         *string  `pulumi:"password"`
	PriorityResolution               *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments              []string `pulumi:"projectEnvironments"`
	ProjectKey                       *string  `pulumi:"projectKey"`
	// (Optional, Default: `false`) When set, if query params are included in the request to Artifactory, they will be passed on to the remote repository.
	PropagateQueryParams              *bool    `pulumi:"propagateQueryParams"`
	PropertySets                      []string `pulumi:"propertySets"`
	Proxy                             *string  `pulumi:"proxy"`
	QueryParams                       *string  `pulumi:"queryParams"`
	RemoteRepoLayoutRef               *string  `pulumi:"remoteRepoLayoutRef"`
	RepoLayoutRef                     *string  `pulumi:"repoLayoutRef"`
	RetrievalCachePeriodSeconds       *int     `pulumi:"retrievalCachePeriodSeconds"`
	ShareConfiguration                bool     `pulumi:"shareConfiguration"`
	SocketTimeoutMillis               *int     `pulumi:"socketTimeoutMillis"`
	StoreArtifactsLocally             *bool    `pulumi:"storeArtifactsLocally"`
	SynchronizeProperties             *bool    `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodHours *int     `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                               *string  `pulumi:"url"`
	Username                          *string  `pulumi:"username"`
	XrayIndex                         *bool    `pulumi:"xrayIndex"`
}

func LookupRemoteGenericRepositoryOutput(ctx *pulumi.Context, args LookupRemoteGenericRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupRemoteGenericRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRemoteGenericRepositoryResult, error) {
			args := v.(LookupRemoteGenericRepositoryArgs)
			r, err := LookupRemoteGenericRepository(ctx, &args, opts...)
			var s LookupRemoteGenericRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRemoteGenericRepositoryResultOutput)
}

// A collection of arguments for invoking getRemoteGenericRepository.
type LookupRemoteGenericRepositoryOutputArgs struct {
	AllowAnyHostAuth          pulumi.BoolPtrInput                                      `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  pulumi.IntPtrInput                                       `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                pulumi.BoolPtrInput                                      `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes pulumi.BoolPtrInput                                      `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        pulumi.BoolPtrInput                                      `pulumi:"bypassHeadRequests"`
	CdnRedirect               pulumi.BoolPtrInput                                      `pulumi:"cdnRedirect"`
	ClientTlsCertificate      pulumi.StringPtrInput                                    `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    GetRemoteGenericRepositoryContentSynchronisationPtrInput `pulumi:"contentSynchronisation"`
	Description               pulumi.StringPtrInput                                    `pulumi:"description"`
	DisableProxy              pulumi.BoolPtrInput                                      `pulumi:"disableProxy"`
	DisableUrlNormalization   pulumi.BoolPtrInput                                      `pulumi:"disableUrlNormalization"`
	DownloadDirect            pulumi.BoolPtrInput                                      `pulumi:"downloadDirect"`
	EnableCookieManagement    pulumi.BoolPtrInput                                      `pulumi:"enableCookieManagement"`
	ExcludesPattern           pulumi.StringPtrInput                                    `pulumi:"excludesPattern"`
	HardFail                  pulumi.BoolPtrInput                                      `pulumi:"hardFail"`
	IncludesPattern           pulumi.StringPtrInput                                    `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                              pulumi.StringInput      `pulumi:"key"`
	ListRemoteFolderItems            pulumi.BoolPtrInput     `pulumi:"listRemoteFolderItems"`
	LocalAddress                     pulumi.StringPtrInput   `pulumi:"localAddress"`
	MetadataRetrievalTimeoutSecs     pulumi.IntPtrInput      `pulumi:"metadataRetrievalTimeoutSecs"`
	MismatchingMimeTypesOverrideList pulumi.StringPtrInput   `pulumi:"mismatchingMimeTypesOverrideList"`
	MissedCachePeriodSeconds         pulumi.IntPtrInput      `pulumi:"missedCachePeriodSeconds"`
	Notes                            pulumi.StringPtrInput   `pulumi:"notes"`
	Offline                          pulumi.BoolPtrInput     `pulumi:"offline"`
	Password                         pulumi.StringPtrInput   `pulumi:"password"`
	PriorityResolution               pulumi.BoolPtrInput     `pulumi:"priorityResolution"`
	ProjectEnvironments              pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey                       pulumi.StringPtrInput   `pulumi:"projectKey"`
	// (Optional, Default: `false`) When set, if query params are included in the request to Artifactory, they will be passed on to the remote repository.
	PropagateQueryParams              pulumi.BoolPtrInput     `pulumi:"propagateQueryParams"`
	PropertySets                      pulumi.StringArrayInput `pulumi:"propertySets"`
	Proxy                             pulumi.StringPtrInput   `pulumi:"proxy"`
	QueryParams                       pulumi.StringPtrInput   `pulumi:"queryParams"`
	RemoteRepoLayoutRef               pulumi.StringPtrInput   `pulumi:"remoteRepoLayoutRef"`
	RepoLayoutRef                     pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	RetrievalCachePeriodSeconds       pulumi.IntPtrInput      `pulumi:"retrievalCachePeriodSeconds"`
	ShareConfiguration                pulumi.BoolPtrInput     `pulumi:"shareConfiguration"`
	SocketTimeoutMillis               pulumi.IntPtrInput      `pulumi:"socketTimeoutMillis"`
	StoreArtifactsLocally             pulumi.BoolPtrInput     `pulumi:"storeArtifactsLocally"`
	SynchronizeProperties             pulumi.BoolPtrInput     `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodHours pulumi.IntPtrInput      `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                               pulumi.StringPtrInput   `pulumi:"url"`
	Username                          pulumi.StringPtrInput   `pulumi:"username"`
	XrayIndex                         pulumi.BoolPtrInput     `pulumi:"xrayIndex"`
}

func (LookupRemoteGenericRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteGenericRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getRemoteGenericRepository.
type LookupRemoteGenericRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupRemoteGenericRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteGenericRepositoryResult)(nil)).Elem()
}

func (o LookupRemoteGenericRepositoryResultOutput) ToLookupRemoteGenericRepositoryResultOutput() LookupRemoteGenericRepositoryResultOutput {
	return o
}

func (o LookupRemoteGenericRepositoryResultOutput) ToLookupRemoteGenericRepositoryResultOutputWithContext(ctx context.Context) LookupRemoteGenericRepositoryResultOutput {
	return o
}

func (o LookupRemoteGenericRepositoryResultOutput) AllowAnyHostAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *bool { return v.AllowAnyHostAuth }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) AssumedOfflinePeriodSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *int { return v.AssumedOfflinePeriodSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) BlockMismatchingMimeTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *bool { return v.BlockMismatchingMimeTypes }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) BypassHeadRequests() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *bool { return v.BypassHeadRequests }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) ClientTlsCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) string { return v.ClientTlsCertificate }).(pulumi.StringOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) ContentSynchronisation() GetRemoteGenericRepositoryContentSynchronisationOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) GetRemoteGenericRepositoryContentSynchronisation {
		return v.ContentSynchronisation
	}).(GetRemoteGenericRepositoryContentSynchronisationOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) DisableProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *bool { return v.DisableProxy }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) DisableUrlNormalization() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *bool { return v.DisableUrlNormalization }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) EnableCookieManagement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *bool { return v.EnableCookieManagement }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) HardFail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *bool { return v.HardFail }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRemoteGenericRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) ListRemoteFolderItems() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *bool { return v.ListRemoteFolderItems }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) LocalAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *string { return v.LocalAddress }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) MetadataRetrievalTimeoutSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *int { return v.MetadataRetrievalTimeoutSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) MismatchingMimeTypesOverrideList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *string { return v.MismatchingMimeTypesOverrideList }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) MissedCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *int { return v.MissedCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) Offline() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *bool { return v.Offline }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// (Optional, Default: `false`) When set, if query params are included in the request to Artifactory, they will be passed on to the remote repository.
func (o LookupRemoteGenericRepositoryResultOutput) PropagateQueryParams() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *bool { return v.PropagateQueryParams }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *string { return v.Proxy }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) QueryParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *string { return v.QueryParams }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) RemoteRepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *string { return v.RemoteRepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) RetrievalCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *int { return v.RetrievalCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) ShareConfiguration() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) bool { return v.ShareConfiguration }).(pulumi.BoolOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) SocketTimeoutMillis() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *int { return v.SocketTimeoutMillis }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) StoreArtifactsLocally() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *bool { return v.StoreArtifactsLocally }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) SynchronizeProperties() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *bool { return v.SynchronizeProperties }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) UnusedArtifactsCleanupPeriodHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *int { return v.UnusedArtifactsCleanupPeriodHours }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteGenericRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGenericRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemoteGenericRepositoryResultOutput{})
}
