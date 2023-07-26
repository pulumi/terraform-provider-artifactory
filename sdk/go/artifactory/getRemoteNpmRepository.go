// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v4/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a remote Npm repository.
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
//			_, err := artifactory.LookupRemoteNpmRepository(ctx, &artifactory.LookupRemoteNpmRepositoryArgs{
//				Key: "remote-npm",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRemoteNpmRepository(ctx *pulumi.Context, args *LookupRemoteNpmRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupRemoteNpmRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRemoteNpmRepositoryResult
	err := ctx.Invoke("artifactory:index/getRemoteNpmRepository:getRemoteNpmRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRemoteNpmRepository.
type LookupRemoteNpmRepositoryArgs struct {
	AllowAnyHostAuth          *bool                                         `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  *int                                          `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool                                         `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool                                         `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool                                         `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool                                         `pulumi:"cdnRedirect"`
	ClientTlsCertificate      *string                                       `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    *GetRemoteNpmRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
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

// A collection of values returned by getRemoteNpmRepository.
type LookupRemoteNpmRepositoryResult struct {
	AllowAnyHostAuth          *bool                                        `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  *int                                         `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool                                        `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool                                        `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool                                        `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool                                        `pulumi:"cdnRedirect"`
	ClientTlsCertificate      string                                       `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    GetRemoteNpmRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
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

func LookupRemoteNpmRepositoryOutput(ctx *pulumi.Context, args LookupRemoteNpmRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupRemoteNpmRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRemoteNpmRepositoryResult, error) {
			args := v.(LookupRemoteNpmRepositoryArgs)
			r, err := LookupRemoteNpmRepository(ctx, &args, opts...)
			var s LookupRemoteNpmRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRemoteNpmRepositoryResultOutput)
}

// A collection of arguments for invoking getRemoteNpmRepository.
type LookupRemoteNpmRepositoryOutputArgs struct {
	AllowAnyHostAuth          pulumi.BoolPtrInput                                  `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  pulumi.IntPtrInput                                   `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                pulumi.BoolPtrInput                                  `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes pulumi.BoolPtrInput                                  `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        pulumi.BoolPtrInput                                  `pulumi:"bypassHeadRequests"`
	CdnRedirect               pulumi.BoolPtrInput                                  `pulumi:"cdnRedirect"`
	ClientTlsCertificate      pulumi.StringPtrInput                                `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    GetRemoteNpmRepositoryContentSynchronisationPtrInput `pulumi:"contentSynchronisation"`
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

func (LookupRemoteNpmRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteNpmRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getRemoteNpmRepository.
type LookupRemoteNpmRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupRemoteNpmRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteNpmRepositoryResult)(nil)).Elem()
}

func (o LookupRemoteNpmRepositoryResultOutput) ToLookupRemoteNpmRepositoryResultOutput() LookupRemoteNpmRepositoryResultOutput {
	return o
}

func (o LookupRemoteNpmRepositoryResultOutput) ToLookupRemoteNpmRepositoryResultOutputWithContext(ctx context.Context) LookupRemoteNpmRepositoryResultOutput {
	return o
}

func (o LookupRemoteNpmRepositoryResultOutput) AllowAnyHostAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *bool { return v.AllowAnyHostAuth }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) AssumedOfflinePeriodSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *int { return v.AssumedOfflinePeriodSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) BlockMismatchingMimeTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *bool { return v.BlockMismatchingMimeTypes }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) BypassHeadRequests() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *bool { return v.BypassHeadRequests }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) ClientTlsCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) string { return v.ClientTlsCertificate }).(pulumi.StringOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) ContentSynchronisation() GetRemoteNpmRepositoryContentSynchronisationOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) GetRemoteNpmRepositoryContentSynchronisation {
		return v.ContentSynchronisation
	}).(GetRemoteNpmRepositoryContentSynchronisationOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) DisableProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *bool { return v.DisableProxy }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) EnableCookieManagement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *bool { return v.EnableCookieManagement }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) HardFail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *bool { return v.HardFail }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRemoteNpmRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) ListRemoteFolderItems() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *bool { return v.ListRemoteFolderItems }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) LocalAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *string { return v.LocalAddress }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) MetadataRetrievalTimeoutSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *int { return v.MetadataRetrievalTimeoutSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) MismatchingMimeTypesOverrideList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *string { return v.MismatchingMimeTypesOverrideList }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) MissedCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *int { return v.MissedCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) Offline() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *bool { return v.Offline }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *string { return v.Proxy }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) QueryParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *string { return v.QueryParams }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) RemoteRepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *string { return v.RemoteRepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) RetrievalCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *int { return v.RetrievalCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) ShareConfiguration() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) bool { return v.ShareConfiguration }).(pulumi.BoolOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) SocketTimeoutMillis() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *int { return v.SocketTimeoutMillis }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) StoreArtifactsLocally() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *bool { return v.StoreArtifactsLocally }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) SynchronizeProperties() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *bool { return v.SynchronizeProperties }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) UnusedArtifactsCleanupPeriodHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *int { return v.UnusedArtifactsCleanupPeriodHours }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNpmRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNpmRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemoteNpmRepositoryResultOutput{})
}
