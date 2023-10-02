// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Retrieves a remote Rpm repository.
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
//			_, err := artifactory.LookupRemoteRpmRepository(ctx, &artifactory.LookupRemoteRpmRepositoryArgs{
//				Key: "remote-rpm",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRemoteRpmRepository(ctx *pulumi.Context, args *LookupRemoteRpmRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupRemoteRpmRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRemoteRpmRepositoryResult
	err := ctx.Invoke("artifactory:index/getRemoteRpmRepository:getRemoteRpmRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRemoteRpmRepository.
type LookupRemoteRpmRepositoryArgs struct {
	AllowAnyHostAuth          *bool                                         `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  *int                                          `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool                                         `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool                                         `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool                                         `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool                                         `pulumi:"cdnRedirect"`
	ClientTlsCertificate      *string                                       `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    *GetRemoteRpmRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description               *string                                       `pulumi:"description"`
	DisableProxy              *bool                                         `pulumi:"disableProxy"`
	DownloadDirect            *bool                                         `pulumi:"downloadDirect"`
	EnableCookieManagement    *bool                                         `pulumi:"enableCookieManagement"`
	ExcludesPattern           *string                                       `pulumi:"excludesPattern"`
	HardFail                  *bool                                         `pulumi:"hardFail"`
	IncludesPattern           *string                                       `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                               string   `pulumi:"key"`
	ListRemoteFolderItems             *bool    `pulumi:"listRemoteFolderItems"`
	LocalAddress                      *string  `pulumi:"localAddress"`
	MetadataRetrievalTimeoutSecs      *int     `pulumi:"metadataRetrievalTimeoutSecs"`
	MismatchingMimeTypesOverrideList  *string  `pulumi:"mismatchingMimeTypesOverrideList"`
	MissedCachePeriodSeconds          *int     `pulumi:"missedCachePeriodSeconds"`
	Notes                             *string  `pulumi:"notes"`
	Offline                           *bool    `pulumi:"offline"`
	Password                          *string  `pulumi:"password"`
	PriorityResolution                *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments               []string `pulumi:"projectEnvironments"`
	ProjectKey                        *string  `pulumi:"projectKey"`
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

// A collection of values returned by getRemoteRpmRepository.
type LookupRemoteRpmRepositoryResult struct {
	AllowAnyHostAuth          *bool                                        `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  *int                                         `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool                                        `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool                                        `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool                                        `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool                                        `pulumi:"cdnRedirect"`
	ClientTlsCertificate      string                                       `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    GetRemoteRpmRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description               *string                                      `pulumi:"description"`
	DisableProxy              *bool                                        `pulumi:"disableProxy"`
	DownloadDirect            *bool                                        `pulumi:"downloadDirect"`
	EnableCookieManagement    *bool                                        `pulumi:"enableCookieManagement"`
	ExcludesPattern           *string                                      `pulumi:"excludesPattern"`
	HardFail                  *bool                                        `pulumi:"hardFail"`
	// The provider-assigned unique ID for this managed resource.
	Id                                string   `pulumi:"id"`
	IncludesPattern                   *string  `pulumi:"includesPattern"`
	Key                               string   `pulumi:"key"`
	ListRemoteFolderItems             *bool    `pulumi:"listRemoteFolderItems"`
	LocalAddress                      *string  `pulumi:"localAddress"`
	MetadataRetrievalTimeoutSecs      *int     `pulumi:"metadataRetrievalTimeoutSecs"`
	MismatchingMimeTypesOverrideList  *string  `pulumi:"mismatchingMimeTypesOverrideList"`
	MissedCachePeriodSeconds          *int     `pulumi:"missedCachePeriodSeconds"`
	Notes                             *string  `pulumi:"notes"`
	Offline                           *bool    `pulumi:"offline"`
	PackageType                       string   `pulumi:"packageType"`
	Password                          *string  `pulumi:"password"`
	PriorityResolution                *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments               []string `pulumi:"projectEnvironments"`
	ProjectKey                        *string  `pulumi:"projectKey"`
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

func LookupRemoteRpmRepositoryOutput(ctx *pulumi.Context, args LookupRemoteRpmRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupRemoteRpmRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRemoteRpmRepositoryResult, error) {
			args := v.(LookupRemoteRpmRepositoryArgs)
			r, err := LookupRemoteRpmRepository(ctx, &args, opts...)
			var s LookupRemoteRpmRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRemoteRpmRepositoryResultOutput)
}

// A collection of arguments for invoking getRemoteRpmRepository.
type LookupRemoteRpmRepositoryOutputArgs struct {
	AllowAnyHostAuth          pulumi.BoolPtrInput                                  `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  pulumi.IntPtrInput                                   `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                pulumi.BoolPtrInput                                  `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes pulumi.BoolPtrInput                                  `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        pulumi.BoolPtrInput                                  `pulumi:"bypassHeadRequests"`
	CdnRedirect               pulumi.BoolPtrInput                                  `pulumi:"cdnRedirect"`
	ClientTlsCertificate      pulumi.StringPtrInput                                `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    GetRemoteRpmRepositoryContentSynchronisationPtrInput `pulumi:"contentSynchronisation"`
	Description               pulumi.StringPtrInput                                `pulumi:"description"`
	DisableProxy              pulumi.BoolPtrInput                                  `pulumi:"disableProxy"`
	DownloadDirect            pulumi.BoolPtrInput                                  `pulumi:"downloadDirect"`
	EnableCookieManagement    pulumi.BoolPtrInput                                  `pulumi:"enableCookieManagement"`
	ExcludesPattern           pulumi.StringPtrInput                                `pulumi:"excludesPattern"`
	HardFail                  pulumi.BoolPtrInput                                  `pulumi:"hardFail"`
	IncludesPattern           pulumi.StringPtrInput                                `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                               pulumi.StringInput      `pulumi:"key"`
	ListRemoteFolderItems             pulumi.BoolPtrInput     `pulumi:"listRemoteFolderItems"`
	LocalAddress                      pulumi.StringPtrInput   `pulumi:"localAddress"`
	MetadataRetrievalTimeoutSecs      pulumi.IntPtrInput      `pulumi:"metadataRetrievalTimeoutSecs"`
	MismatchingMimeTypesOverrideList  pulumi.StringPtrInput   `pulumi:"mismatchingMimeTypesOverrideList"`
	MissedCachePeriodSeconds          pulumi.IntPtrInput      `pulumi:"missedCachePeriodSeconds"`
	Notes                             pulumi.StringPtrInput   `pulumi:"notes"`
	Offline                           pulumi.BoolPtrInput     `pulumi:"offline"`
	Password                          pulumi.StringPtrInput   `pulumi:"password"`
	PriorityResolution                pulumi.BoolPtrInput     `pulumi:"priorityResolution"`
	ProjectEnvironments               pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey                        pulumi.StringPtrInput   `pulumi:"projectKey"`
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

func (LookupRemoteRpmRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteRpmRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getRemoteRpmRepository.
type LookupRemoteRpmRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupRemoteRpmRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteRpmRepositoryResult)(nil)).Elem()
}

func (o LookupRemoteRpmRepositoryResultOutput) ToLookupRemoteRpmRepositoryResultOutput() LookupRemoteRpmRepositoryResultOutput {
	return o
}

func (o LookupRemoteRpmRepositoryResultOutput) ToLookupRemoteRpmRepositoryResultOutputWithContext(ctx context.Context) LookupRemoteRpmRepositoryResultOutput {
	return o
}

func (o LookupRemoteRpmRepositoryResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupRemoteRpmRepositoryResult] {
	return pulumix.Output[LookupRemoteRpmRepositoryResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupRemoteRpmRepositoryResultOutput) AllowAnyHostAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *bool { return v.AllowAnyHostAuth }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) AssumedOfflinePeriodSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *int { return v.AssumedOfflinePeriodSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) BlockMismatchingMimeTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *bool { return v.BlockMismatchingMimeTypes }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) BypassHeadRequests() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *bool { return v.BypassHeadRequests }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) ClientTlsCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) string { return v.ClientTlsCertificate }).(pulumi.StringOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) ContentSynchronisation() GetRemoteRpmRepositoryContentSynchronisationOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) GetRemoteRpmRepositoryContentSynchronisation {
		return v.ContentSynchronisation
	}).(GetRemoteRpmRepositoryContentSynchronisationOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) DisableProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *bool { return v.DisableProxy }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) EnableCookieManagement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *bool { return v.EnableCookieManagement }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) HardFail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *bool { return v.HardFail }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRemoteRpmRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) ListRemoteFolderItems() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *bool { return v.ListRemoteFolderItems }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) LocalAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *string { return v.LocalAddress }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) MetadataRetrievalTimeoutSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *int { return v.MetadataRetrievalTimeoutSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) MismatchingMimeTypesOverrideList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *string { return v.MismatchingMimeTypesOverrideList }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) MissedCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *int { return v.MissedCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) Offline() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *bool { return v.Offline }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *string { return v.Proxy }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) QueryParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *string { return v.QueryParams }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) RemoteRepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *string { return v.RemoteRepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) RetrievalCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *int { return v.RetrievalCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) ShareConfiguration() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) bool { return v.ShareConfiguration }).(pulumi.BoolOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) SocketTimeoutMillis() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *int { return v.SocketTimeoutMillis }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) StoreArtifactsLocally() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *bool { return v.StoreArtifactsLocally }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) SynchronizeProperties() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *bool { return v.SynchronizeProperties }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) UnusedArtifactsCleanupPeriodHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *int { return v.UnusedArtifactsCleanupPeriodHours }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteRpmRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteRpmRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemoteRpmRepositoryResultOutput{})
}
