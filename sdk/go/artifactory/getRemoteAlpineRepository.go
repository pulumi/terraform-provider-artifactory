// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v6/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a remote Alpine repository.
//
// ## Example Usage
//
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
//			_, err := artifactory.LookupRemoteAlpineRepository(ctx, &artifactory.LookupRemoteAlpineRepositoryArgs{
//				Key: "remote-alpine",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRemoteAlpineRepository(ctx *pulumi.Context, args *LookupRemoteAlpineRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupRemoteAlpineRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRemoteAlpineRepositoryResult
	err := ctx.Invoke("artifactory:index/getRemoteAlpineRepository:getRemoteAlpineRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRemoteAlpineRepository.
type LookupRemoteAlpineRepositoryArgs struct {
	AllowAnyHostAuth          *bool                                            `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  *int                                             `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool                                            `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool                                            `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool                                            `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool                                            `pulumi:"cdnRedirect"`
	ClientTlsCertificate      *string                                          `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    *GetRemoteAlpineRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description               *string                                          `pulumi:"description"`
	DisableProxy              *bool                                            `pulumi:"disableProxy"`
	DisableUrlNormalization   *bool                                            `pulumi:"disableUrlNormalization"`
	DownloadDirect            *bool                                            `pulumi:"downloadDirect"`
	EnableCookieManagement    *bool                                            `pulumi:"enableCookieManagement"`
	ExcludesPattern           *string                                          `pulumi:"excludesPattern"`
	HardFail                  *bool                                            `pulumi:"hardFail"`
	IncludesPattern           *string                                          `pulumi:"includesPattern"`
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

// A collection of values returned by getRemoteAlpineRepository.
type LookupRemoteAlpineRepositoryResult struct {
	AllowAnyHostAuth          *bool                                           `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  *int                                            `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool                                           `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool                                           `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool                                           `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool                                           `pulumi:"cdnRedirect"`
	ClientTlsCertificate      string                                          `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    GetRemoteAlpineRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description               *string                                         `pulumi:"description"`
	DisableProxy              *bool                                           `pulumi:"disableProxy"`
	DisableUrlNormalization   *bool                                           `pulumi:"disableUrlNormalization"`
	DownloadDirect            *bool                                           `pulumi:"downloadDirect"`
	EnableCookieManagement    *bool                                           `pulumi:"enableCookieManagement"`
	ExcludesPattern           *string                                         `pulumi:"excludesPattern"`
	HardFail                  *bool                                           `pulumi:"hardFail"`
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

func LookupRemoteAlpineRepositoryOutput(ctx *pulumi.Context, args LookupRemoteAlpineRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupRemoteAlpineRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRemoteAlpineRepositoryResult, error) {
			args := v.(LookupRemoteAlpineRepositoryArgs)
			r, err := LookupRemoteAlpineRepository(ctx, &args, opts...)
			var s LookupRemoteAlpineRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRemoteAlpineRepositoryResultOutput)
}

// A collection of arguments for invoking getRemoteAlpineRepository.
type LookupRemoteAlpineRepositoryOutputArgs struct {
	AllowAnyHostAuth          pulumi.BoolPtrInput                                     `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  pulumi.IntPtrInput                                      `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                pulumi.BoolPtrInput                                     `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes pulumi.BoolPtrInput                                     `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        pulumi.BoolPtrInput                                     `pulumi:"bypassHeadRequests"`
	CdnRedirect               pulumi.BoolPtrInput                                     `pulumi:"cdnRedirect"`
	ClientTlsCertificate      pulumi.StringPtrInput                                   `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    GetRemoteAlpineRepositoryContentSynchronisationPtrInput `pulumi:"contentSynchronisation"`
	Description               pulumi.StringPtrInput                                   `pulumi:"description"`
	DisableProxy              pulumi.BoolPtrInput                                     `pulumi:"disableProxy"`
	DisableUrlNormalization   pulumi.BoolPtrInput                                     `pulumi:"disableUrlNormalization"`
	DownloadDirect            pulumi.BoolPtrInput                                     `pulumi:"downloadDirect"`
	EnableCookieManagement    pulumi.BoolPtrInput                                     `pulumi:"enableCookieManagement"`
	ExcludesPattern           pulumi.StringPtrInput                                   `pulumi:"excludesPattern"`
	HardFail                  pulumi.BoolPtrInput                                     `pulumi:"hardFail"`
	IncludesPattern           pulumi.StringPtrInput                                   `pulumi:"includesPattern"`
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

func (LookupRemoteAlpineRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteAlpineRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getRemoteAlpineRepository.
type LookupRemoteAlpineRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupRemoteAlpineRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteAlpineRepositoryResult)(nil)).Elem()
}

func (o LookupRemoteAlpineRepositoryResultOutput) ToLookupRemoteAlpineRepositoryResultOutput() LookupRemoteAlpineRepositoryResultOutput {
	return o
}

func (o LookupRemoteAlpineRepositoryResultOutput) ToLookupRemoteAlpineRepositoryResultOutputWithContext(ctx context.Context) LookupRemoteAlpineRepositoryResultOutput {
	return o
}

func (o LookupRemoteAlpineRepositoryResultOutput) AllowAnyHostAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *bool { return v.AllowAnyHostAuth }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) AssumedOfflinePeriodSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *int { return v.AssumedOfflinePeriodSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) BlockMismatchingMimeTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *bool { return v.BlockMismatchingMimeTypes }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) BypassHeadRequests() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *bool { return v.BypassHeadRequests }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) ClientTlsCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) string { return v.ClientTlsCertificate }).(pulumi.StringOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) ContentSynchronisation() GetRemoteAlpineRepositoryContentSynchronisationOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) GetRemoteAlpineRepositoryContentSynchronisation {
		return v.ContentSynchronisation
	}).(GetRemoteAlpineRepositoryContentSynchronisationOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) DisableProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *bool { return v.DisableProxy }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) DisableUrlNormalization() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *bool { return v.DisableUrlNormalization }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) EnableCookieManagement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *bool { return v.EnableCookieManagement }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) HardFail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *bool { return v.HardFail }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRemoteAlpineRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) ListRemoteFolderItems() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *bool { return v.ListRemoteFolderItems }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) LocalAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *string { return v.LocalAddress }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) MetadataRetrievalTimeoutSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *int { return v.MetadataRetrievalTimeoutSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) MismatchingMimeTypesOverrideList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *string { return v.MismatchingMimeTypesOverrideList }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) MissedCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *int { return v.MissedCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) Offline() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *bool { return v.Offline }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *string { return v.Proxy }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) QueryParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *string { return v.QueryParams }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) RemoteRepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *string { return v.RemoteRepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) RetrievalCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *int { return v.RetrievalCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) ShareConfiguration() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) bool { return v.ShareConfiguration }).(pulumi.BoolOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) SocketTimeoutMillis() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *int { return v.SocketTimeoutMillis }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) StoreArtifactsLocally() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *bool { return v.StoreArtifactsLocally }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) SynchronizeProperties() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *bool { return v.SynchronizeProperties }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) UnusedArtifactsCleanupPeriodHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *int { return v.UnusedArtifactsCleanupPeriodHours }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteAlpineRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteAlpineRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemoteAlpineRepositoryResultOutput{})
}
