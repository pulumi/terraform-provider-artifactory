// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRemoteNugetRepository(ctx *pulumi.Context, args *LookupRemoteNugetRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupRemoteNugetRepositoryResult, error) {
	var rv LookupRemoteNugetRepositoryResult
	err := ctx.Invoke("artifactory:index/getRemoteNugetRepository:getRemoteNugetRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRemoteNugetRepository.
type LookupRemoteNugetRepositoryArgs struct {
	AllowAnyHostAuth                  *bool                                           `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs          *int                                            `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                        *bool                                           `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes         *bool                                           `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests                *bool                                           `pulumi:"bypassHeadRequests"`
	CdnRedirect                       *bool                                           `pulumi:"cdnRedirect"`
	ClientTlsCertificate              *string                                         `pulumi:"clientTlsCertificate"`
	ContentSynchronisation            *GetRemoteNugetRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description                       *string                                         `pulumi:"description"`
	DownloadContextPath               *string                                         `pulumi:"downloadContextPath"`
	DownloadDirect                    *bool                                           `pulumi:"downloadDirect"`
	EnableCookieManagement            *bool                                           `pulumi:"enableCookieManagement"`
	ExcludesPattern                   *string                                         `pulumi:"excludesPattern"`
	FeedContextPath                   *string                                         `pulumi:"feedContextPath"`
	ForceNugetAuthentication          *bool                                           `pulumi:"forceNugetAuthentication"`
	HardFail                          *bool                                           `pulumi:"hardFail"`
	IncludesPattern                   *string                                         `pulumi:"includesPattern"`
	Key                               string                                          `pulumi:"key"`
	ListRemoteFolderItems             *bool                                           `pulumi:"listRemoteFolderItems"`
	LocalAddress                      *string                                         `pulumi:"localAddress"`
	MetadataRetrievalTimeoutSecs      *int                                            `pulumi:"metadataRetrievalTimeoutSecs"`
	MismatchingMimeTypesOverrideList  *string                                         `pulumi:"mismatchingMimeTypesOverrideList"`
	MissedCachePeriodSeconds          *int                                            `pulumi:"missedCachePeriodSeconds"`
	Notes                             *string                                         `pulumi:"notes"`
	Offline                           *bool                                           `pulumi:"offline"`
	Password                          *string                                         `pulumi:"password"`
	PriorityResolution                *bool                                           `pulumi:"priorityResolution"`
	ProjectEnvironments               []string                                        `pulumi:"projectEnvironments"`
	ProjectKey                        *string                                         `pulumi:"projectKey"`
	PropertySets                      []string                                        `pulumi:"propertySets"`
	Proxy                             *string                                         `pulumi:"proxy"`
	QueryParams                       *string                                         `pulumi:"queryParams"`
	RemoteRepoLayoutRef               *string                                         `pulumi:"remoteRepoLayoutRef"`
	RepoLayoutRef                     *string                                         `pulumi:"repoLayoutRef"`
	RetrievalCachePeriodSeconds       *int                                            `pulumi:"retrievalCachePeriodSeconds"`
	ShareConfiguration                *bool                                           `pulumi:"shareConfiguration"`
	SocketTimeoutMillis               *int                                            `pulumi:"socketTimeoutMillis"`
	StoreArtifactsLocally             *bool                                           `pulumi:"storeArtifactsLocally"`
	SymbolServerUrl                   *string                                         `pulumi:"symbolServerUrl"`
	SynchronizeProperties             *bool                                           `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodHours *int                                            `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                               *string                                         `pulumi:"url"`
	Username                          *string                                         `pulumi:"username"`
	V3FeedUrl                         *string                                         `pulumi:"v3FeedUrl"`
	XrayIndex                         *bool                                           `pulumi:"xrayIndex"`
}

// A collection of values returned by getRemoteNugetRepository.
type LookupRemoteNugetRepositoryResult struct {
	AllowAnyHostAuth          *bool                                          `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  *int                                           `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool                                          `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool                                          `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool                                          `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool                                          `pulumi:"cdnRedirect"`
	ClientTlsCertificate      string                                         `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    GetRemoteNugetRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description               *string                                        `pulumi:"description"`
	DownloadContextPath       *string                                        `pulumi:"downloadContextPath"`
	DownloadDirect            *bool                                          `pulumi:"downloadDirect"`
	EnableCookieManagement    *bool                                          `pulumi:"enableCookieManagement"`
	ExcludesPattern           *string                                        `pulumi:"excludesPattern"`
	FeedContextPath           *string                                        `pulumi:"feedContextPath"`
	ForceNugetAuthentication  *bool                                          `pulumi:"forceNugetAuthentication"`
	HardFail                  *bool                                          `pulumi:"hardFail"`
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
	SymbolServerUrl                   *string  `pulumi:"symbolServerUrl"`
	SynchronizeProperties             *bool    `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodHours *int     `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                               *string  `pulumi:"url"`
	Username                          *string  `pulumi:"username"`
	V3FeedUrl                         *string  `pulumi:"v3FeedUrl"`
	XrayIndex                         *bool    `pulumi:"xrayIndex"`
}

func LookupRemoteNugetRepositoryOutput(ctx *pulumi.Context, args LookupRemoteNugetRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupRemoteNugetRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRemoteNugetRepositoryResult, error) {
			args := v.(LookupRemoteNugetRepositoryArgs)
			r, err := LookupRemoteNugetRepository(ctx, &args, opts...)
			var s LookupRemoteNugetRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRemoteNugetRepositoryResultOutput)
}

// A collection of arguments for invoking getRemoteNugetRepository.
type LookupRemoteNugetRepositoryOutputArgs struct {
	AllowAnyHostAuth                  pulumi.BoolPtrInput                                    `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs          pulumi.IntPtrInput                                     `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                        pulumi.BoolPtrInput                                    `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes         pulumi.BoolPtrInput                                    `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests                pulumi.BoolPtrInput                                    `pulumi:"bypassHeadRequests"`
	CdnRedirect                       pulumi.BoolPtrInput                                    `pulumi:"cdnRedirect"`
	ClientTlsCertificate              pulumi.StringPtrInput                                  `pulumi:"clientTlsCertificate"`
	ContentSynchronisation            GetRemoteNugetRepositoryContentSynchronisationPtrInput `pulumi:"contentSynchronisation"`
	Description                       pulumi.StringPtrInput                                  `pulumi:"description"`
	DownloadContextPath               pulumi.StringPtrInput                                  `pulumi:"downloadContextPath"`
	DownloadDirect                    pulumi.BoolPtrInput                                    `pulumi:"downloadDirect"`
	EnableCookieManagement            pulumi.BoolPtrInput                                    `pulumi:"enableCookieManagement"`
	ExcludesPattern                   pulumi.StringPtrInput                                  `pulumi:"excludesPattern"`
	FeedContextPath                   pulumi.StringPtrInput                                  `pulumi:"feedContextPath"`
	ForceNugetAuthentication          pulumi.BoolPtrInput                                    `pulumi:"forceNugetAuthentication"`
	HardFail                          pulumi.BoolPtrInput                                    `pulumi:"hardFail"`
	IncludesPattern                   pulumi.StringPtrInput                                  `pulumi:"includesPattern"`
	Key                               pulumi.StringInput                                     `pulumi:"key"`
	ListRemoteFolderItems             pulumi.BoolPtrInput                                    `pulumi:"listRemoteFolderItems"`
	LocalAddress                      pulumi.StringPtrInput                                  `pulumi:"localAddress"`
	MetadataRetrievalTimeoutSecs      pulumi.IntPtrInput                                     `pulumi:"metadataRetrievalTimeoutSecs"`
	MismatchingMimeTypesOverrideList  pulumi.StringPtrInput                                  `pulumi:"mismatchingMimeTypesOverrideList"`
	MissedCachePeriodSeconds          pulumi.IntPtrInput                                     `pulumi:"missedCachePeriodSeconds"`
	Notes                             pulumi.StringPtrInput                                  `pulumi:"notes"`
	Offline                           pulumi.BoolPtrInput                                    `pulumi:"offline"`
	Password                          pulumi.StringPtrInput                                  `pulumi:"password"`
	PriorityResolution                pulumi.BoolPtrInput                                    `pulumi:"priorityResolution"`
	ProjectEnvironments               pulumi.StringArrayInput                                `pulumi:"projectEnvironments"`
	ProjectKey                        pulumi.StringPtrInput                                  `pulumi:"projectKey"`
	PropertySets                      pulumi.StringArrayInput                                `pulumi:"propertySets"`
	Proxy                             pulumi.StringPtrInput                                  `pulumi:"proxy"`
	QueryParams                       pulumi.StringPtrInput                                  `pulumi:"queryParams"`
	RemoteRepoLayoutRef               pulumi.StringPtrInput                                  `pulumi:"remoteRepoLayoutRef"`
	RepoLayoutRef                     pulumi.StringPtrInput                                  `pulumi:"repoLayoutRef"`
	RetrievalCachePeriodSeconds       pulumi.IntPtrInput                                     `pulumi:"retrievalCachePeriodSeconds"`
	ShareConfiguration                pulumi.BoolPtrInput                                    `pulumi:"shareConfiguration"`
	SocketTimeoutMillis               pulumi.IntPtrInput                                     `pulumi:"socketTimeoutMillis"`
	StoreArtifactsLocally             pulumi.BoolPtrInput                                    `pulumi:"storeArtifactsLocally"`
	SymbolServerUrl                   pulumi.StringPtrInput                                  `pulumi:"symbolServerUrl"`
	SynchronizeProperties             pulumi.BoolPtrInput                                    `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodHours pulumi.IntPtrInput                                     `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                               pulumi.StringPtrInput                                  `pulumi:"url"`
	Username                          pulumi.StringPtrInput                                  `pulumi:"username"`
	V3FeedUrl                         pulumi.StringPtrInput                                  `pulumi:"v3FeedUrl"`
	XrayIndex                         pulumi.BoolPtrInput                                    `pulumi:"xrayIndex"`
}

func (LookupRemoteNugetRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteNugetRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getRemoteNugetRepository.
type LookupRemoteNugetRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupRemoteNugetRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteNugetRepositoryResult)(nil)).Elem()
}

func (o LookupRemoteNugetRepositoryResultOutput) ToLookupRemoteNugetRepositoryResultOutput() LookupRemoteNugetRepositoryResultOutput {
	return o
}

func (o LookupRemoteNugetRepositoryResultOutput) ToLookupRemoteNugetRepositoryResultOutputWithContext(ctx context.Context) LookupRemoteNugetRepositoryResultOutput {
	return o
}

func (o LookupRemoteNugetRepositoryResultOutput) AllowAnyHostAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *bool { return v.AllowAnyHostAuth }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) AssumedOfflinePeriodSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *int { return v.AssumedOfflinePeriodSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) BlockMismatchingMimeTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *bool { return v.BlockMismatchingMimeTypes }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) BypassHeadRequests() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *bool { return v.BypassHeadRequests }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) ClientTlsCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) string { return v.ClientTlsCertificate }).(pulumi.StringOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) ContentSynchronisation() GetRemoteNugetRepositoryContentSynchronisationOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) GetRemoteNugetRepositoryContentSynchronisation {
		return v.ContentSynchronisation
	}).(GetRemoteNugetRepositoryContentSynchronisationOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) DownloadContextPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *string { return v.DownloadContextPath }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) EnableCookieManagement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *bool { return v.EnableCookieManagement }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) FeedContextPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *string { return v.FeedContextPath }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) ForceNugetAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *bool { return v.ForceNugetAuthentication }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) HardFail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *bool { return v.HardFail }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRemoteNugetRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) ListRemoteFolderItems() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *bool { return v.ListRemoteFolderItems }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) LocalAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *string { return v.LocalAddress }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) MetadataRetrievalTimeoutSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *int { return v.MetadataRetrievalTimeoutSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) MismatchingMimeTypesOverrideList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *string { return v.MismatchingMimeTypesOverrideList }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) MissedCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *int { return v.MissedCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) Offline() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *bool { return v.Offline }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *string { return v.Proxy }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) QueryParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *string { return v.QueryParams }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) RemoteRepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *string { return v.RemoteRepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) RetrievalCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *int { return v.RetrievalCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) ShareConfiguration() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) bool { return v.ShareConfiguration }).(pulumi.BoolOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) SocketTimeoutMillis() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *int { return v.SocketTimeoutMillis }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) StoreArtifactsLocally() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *bool { return v.StoreArtifactsLocally }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) SymbolServerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *string { return v.SymbolServerUrl }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) SynchronizeProperties() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *bool { return v.SynchronizeProperties }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) UnusedArtifactsCleanupPeriodHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *int { return v.UnusedArtifactsCleanupPeriodHours }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) V3FeedUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *string { return v.V3FeedUrl }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteNugetRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteNugetRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemoteNugetRepositoryResultOutput{})
}
