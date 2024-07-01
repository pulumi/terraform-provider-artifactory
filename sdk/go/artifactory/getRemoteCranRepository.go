// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v7/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a remote CRAN repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v7/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.LookupRemoteCranRepository(ctx, &artifactory.LookupRemoteCranRepositoryArgs{
//				Key: "remote-cran",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRemoteCranRepository(ctx *pulumi.Context, args *LookupRemoteCranRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupRemoteCranRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRemoteCranRepositoryResult
	err := ctx.Invoke("artifactory:index/getRemoteCranRepository:getRemoteCranRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRemoteCranRepository.
type LookupRemoteCranRepositoryArgs struct {
	AllowAnyHostAuth          *bool                                          `pulumi:"allowAnyHostAuth"`
	ArchiveBrowsingEnabled    *bool                                          `pulumi:"archiveBrowsingEnabled"`
	AssumedOfflinePeriodSecs  *int                                           `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool                                          `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool                                          `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool                                          `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool                                          `pulumi:"cdnRedirect"`
	ClientTlsCertificate      *string                                        `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    *GetRemoteCranRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description               *string                                        `pulumi:"description"`
	DisableProxy              *bool                                          `pulumi:"disableProxy"`
	DisableUrlNormalization   *bool                                          `pulumi:"disableUrlNormalization"`
	DownloadDirect            *bool                                          `pulumi:"downloadDirect"`
	EnableCookieManagement    *bool                                          `pulumi:"enableCookieManagement"`
	ExcludesPattern           *string                                        `pulumi:"excludesPattern"`
	HardFail                  *bool                                          `pulumi:"hardFail"`
	IncludesPattern           *string                                        `pulumi:"includesPattern"`
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

// A collection of values returned by getRemoteCranRepository.
type LookupRemoteCranRepositoryResult struct {
	AllowAnyHostAuth          *bool                                         `pulumi:"allowAnyHostAuth"`
	ArchiveBrowsingEnabled    *bool                                         `pulumi:"archiveBrowsingEnabled"`
	AssumedOfflinePeriodSecs  *int                                          `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool                                         `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool                                         `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool                                         `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool                                         `pulumi:"cdnRedirect"`
	ClientTlsCertificate      string                                        `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    GetRemoteCranRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description               *string                                       `pulumi:"description"`
	DisableProxy              *bool                                         `pulumi:"disableProxy"`
	DisableUrlNormalization   *bool                                         `pulumi:"disableUrlNormalization"`
	DownloadDirect            *bool                                         `pulumi:"downloadDirect"`
	EnableCookieManagement    *bool                                         `pulumi:"enableCookieManagement"`
	ExcludesPattern           *string                                       `pulumi:"excludesPattern"`
	HardFail                  *bool                                         `pulumi:"hardFail"`
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

func LookupRemoteCranRepositoryOutput(ctx *pulumi.Context, args LookupRemoteCranRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupRemoteCranRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRemoteCranRepositoryResult, error) {
			args := v.(LookupRemoteCranRepositoryArgs)
			r, err := LookupRemoteCranRepository(ctx, &args, opts...)
			var s LookupRemoteCranRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRemoteCranRepositoryResultOutput)
}

// A collection of arguments for invoking getRemoteCranRepository.
type LookupRemoteCranRepositoryOutputArgs struct {
	AllowAnyHostAuth          pulumi.BoolPtrInput                                   `pulumi:"allowAnyHostAuth"`
	ArchiveBrowsingEnabled    pulumi.BoolPtrInput                                   `pulumi:"archiveBrowsingEnabled"`
	AssumedOfflinePeriodSecs  pulumi.IntPtrInput                                    `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                pulumi.BoolPtrInput                                   `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes pulumi.BoolPtrInput                                   `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        pulumi.BoolPtrInput                                   `pulumi:"bypassHeadRequests"`
	CdnRedirect               pulumi.BoolPtrInput                                   `pulumi:"cdnRedirect"`
	ClientTlsCertificate      pulumi.StringPtrInput                                 `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    GetRemoteCranRepositoryContentSynchronisationPtrInput `pulumi:"contentSynchronisation"`
	Description               pulumi.StringPtrInput                                 `pulumi:"description"`
	DisableProxy              pulumi.BoolPtrInput                                   `pulumi:"disableProxy"`
	DisableUrlNormalization   pulumi.BoolPtrInput                                   `pulumi:"disableUrlNormalization"`
	DownloadDirect            pulumi.BoolPtrInput                                   `pulumi:"downloadDirect"`
	EnableCookieManagement    pulumi.BoolPtrInput                                   `pulumi:"enableCookieManagement"`
	ExcludesPattern           pulumi.StringPtrInput                                 `pulumi:"excludesPattern"`
	HardFail                  pulumi.BoolPtrInput                                   `pulumi:"hardFail"`
	IncludesPattern           pulumi.StringPtrInput                                 `pulumi:"includesPattern"`
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

func (LookupRemoteCranRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteCranRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getRemoteCranRepository.
type LookupRemoteCranRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupRemoteCranRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteCranRepositoryResult)(nil)).Elem()
}

func (o LookupRemoteCranRepositoryResultOutput) ToLookupRemoteCranRepositoryResultOutput() LookupRemoteCranRepositoryResultOutput {
	return o
}

func (o LookupRemoteCranRepositoryResultOutput) ToLookupRemoteCranRepositoryResultOutputWithContext(ctx context.Context) LookupRemoteCranRepositoryResultOutput {
	return o
}

func (o LookupRemoteCranRepositoryResultOutput) AllowAnyHostAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *bool { return v.AllowAnyHostAuth }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) AssumedOfflinePeriodSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *int { return v.AssumedOfflinePeriodSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) BlockMismatchingMimeTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *bool { return v.BlockMismatchingMimeTypes }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) BypassHeadRequests() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *bool { return v.BypassHeadRequests }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) ClientTlsCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) string { return v.ClientTlsCertificate }).(pulumi.StringOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) ContentSynchronisation() GetRemoteCranRepositoryContentSynchronisationOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) GetRemoteCranRepositoryContentSynchronisation {
		return v.ContentSynchronisation
	}).(GetRemoteCranRepositoryContentSynchronisationOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) DisableProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *bool { return v.DisableProxy }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) DisableUrlNormalization() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *bool { return v.DisableUrlNormalization }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) EnableCookieManagement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *bool { return v.EnableCookieManagement }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) HardFail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *bool { return v.HardFail }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRemoteCranRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) ListRemoteFolderItems() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *bool { return v.ListRemoteFolderItems }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) LocalAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *string { return v.LocalAddress }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) MetadataRetrievalTimeoutSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *int { return v.MetadataRetrievalTimeoutSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) MismatchingMimeTypesOverrideList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *string { return v.MismatchingMimeTypesOverrideList }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) MissedCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *int { return v.MissedCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) Offline() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *bool { return v.Offline }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *string { return v.Proxy }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) QueryParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *string { return v.QueryParams }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) RemoteRepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *string { return v.RemoteRepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) RetrievalCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *int { return v.RetrievalCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) ShareConfiguration() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) bool { return v.ShareConfiguration }).(pulumi.BoolOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) SocketTimeoutMillis() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *int { return v.SocketTimeoutMillis }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) StoreArtifactsLocally() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *bool { return v.StoreArtifactsLocally }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) SynchronizeProperties() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *bool { return v.SynchronizeProperties }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) UnusedArtifactsCleanupPeriodHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *int { return v.UnusedArtifactsCleanupPeriodHours }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCranRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCranRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemoteCranRepositoryResultOutput{})
}
