// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRemotePypiRepository(ctx *pulumi.Context, args *LookupRemotePypiRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupRemotePypiRepositoryResult, error) {
	var rv LookupRemotePypiRepositoryResult
	err := ctx.Invoke("artifactory:index/getRemotePypiRepository:getRemotePypiRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRemotePypiRepository.
type LookupRemotePypiRepositoryArgs struct {
	AllowAnyHostAuth                  *bool                                          `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs          *int                                           `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                        *bool                                          `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes         *bool                                          `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests                *bool                                          `pulumi:"bypassHeadRequests"`
	CdnRedirect                       *bool                                          `pulumi:"cdnRedirect"`
	ClientTlsCertificate              *string                                        `pulumi:"clientTlsCertificate"`
	ContentSynchronisation            *GetRemotePypiRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description                       *string                                        `pulumi:"description"`
	DownloadDirect                    *bool                                          `pulumi:"downloadDirect"`
	EnableCookieManagement            *bool                                          `pulumi:"enableCookieManagement"`
	ExcludesPattern                   *string                                        `pulumi:"excludesPattern"`
	HardFail                          *bool                                          `pulumi:"hardFail"`
	IncludesPattern                   *string                                        `pulumi:"includesPattern"`
	Key                               string                                         `pulumi:"key"`
	ListRemoteFolderItems             *bool                                          `pulumi:"listRemoteFolderItems"`
	LocalAddress                      *string                                        `pulumi:"localAddress"`
	MetadataRetrievalTimeoutSecs      *int                                           `pulumi:"metadataRetrievalTimeoutSecs"`
	MismatchingMimeTypesOverrideList  *string                                        `pulumi:"mismatchingMimeTypesOverrideList"`
	MissedCachePeriodSeconds          *int                                           `pulumi:"missedCachePeriodSeconds"`
	Notes                             *string                                        `pulumi:"notes"`
	Offline                           *bool                                          `pulumi:"offline"`
	Password                          *string                                        `pulumi:"password"`
	PriorityResolution                *bool                                          `pulumi:"priorityResolution"`
	ProjectEnvironments               []string                                       `pulumi:"projectEnvironments"`
	ProjectKey                        *string                                        `pulumi:"projectKey"`
	PropertySets                      []string                                       `pulumi:"propertySets"`
	Proxy                             *string                                        `pulumi:"proxy"`
	PypiRegistryUrl                   *string                                        `pulumi:"pypiRegistryUrl"`
	PypiRepositorySuffix              *string                                        `pulumi:"pypiRepositorySuffix"`
	QueryParams                       *string                                        `pulumi:"queryParams"`
	RemoteRepoLayoutRef               *string                                        `pulumi:"remoteRepoLayoutRef"`
	RepoLayoutRef                     *string                                        `pulumi:"repoLayoutRef"`
	RetrievalCachePeriodSeconds       *int                                           `pulumi:"retrievalCachePeriodSeconds"`
	ShareConfiguration                *bool                                          `pulumi:"shareConfiguration"`
	SocketTimeoutMillis               *int                                           `pulumi:"socketTimeoutMillis"`
	StoreArtifactsLocally             *bool                                          `pulumi:"storeArtifactsLocally"`
	SynchronizeProperties             *bool                                          `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodHours *int                                           `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                               *string                                        `pulumi:"url"`
	Username                          *string                                        `pulumi:"username"`
	XrayIndex                         *bool                                          `pulumi:"xrayIndex"`
}

// A collection of values returned by getRemotePypiRepository.
type LookupRemotePypiRepositoryResult struct {
	AllowAnyHostAuth          *bool                                         `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  *int                                          `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool                                         `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool                                         `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool                                         `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool                                         `pulumi:"cdnRedirect"`
	ClientTlsCertificate      string                                        `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    GetRemotePypiRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description               *string                                       `pulumi:"description"`
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
	PypiRegistryUrl                   *string  `pulumi:"pypiRegistryUrl"`
	PypiRepositorySuffix              *string  `pulumi:"pypiRepositorySuffix"`
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

func LookupRemotePypiRepositoryOutput(ctx *pulumi.Context, args LookupRemotePypiRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupRemotePypiRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRemotePypiRepositoryResult, error) {
			args := v.(LookupRemotePypiRepositoryArgs)
			r, err := LookupRemotePypiRepository(ctx, &args, opts...)
			var s LookupRemotePypiRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRemotePypiRepositoryResultOutput)
}

// A collection of arguments for invoking getRemotePypiRepository.
type LookupRemotePypiRepositoryOutputArgs struct {
	AllowAnyHostAuth                  pulumi.BoolPtrInput                                   `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs          pulumi.IntPtrInput                                    `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                        pulumi.BoolPtrInput                                   `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes         pulumi.BoolPtrInput                                   `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests                pulumi.BoolPtrInput                                   `pulumi:"bypassHeadRequests"`
	CdnRedirect                       pulumi.BoolPtrInput                                   `pulumi:"cdnRedirect"`
	ClientTlsCertificate              pulumi.StringPtrInput                                 `pulumi:"clientTlsCertificate"`
	ContentSynchronisation            GetRemotePypiRepositoryContentSynchronisationPtrInput `pulumi:"contentSynchronisation"`
	Description                       pulumi.StringPtrInput                                 `pulumi:"description"`
	DownloadDirect                    pulumi.BoolPtrInput                                   `pulumi:"downloadDirect"`
	EnableCookieManagement            pulumi.BoolPtrInput                                   `pulumi:"enableCookieManagement"`
	ExcludesPattern                   pulumi.StringPtrInput                                 `pulumi:"excludesPattern"`
	HardFail                          pulumi.BoolPtrInput                                   `pulumi:"hardFail"`
	IncludesPattern                   pulumi.StringPtrInput                                 `pulumi:"includesPattern"`
	Key                               pulumi.StringInput                                    `pulumi:"key"`
	ListRemoteFolderItems             pulumi.BoolPtrInput                                   `pulumi:"listRemoteFolderItems"`
	LocalAddress                      pulumi.StringPtrInput                                 `pulumi:"localAddress"`
	MetadataRetrievalTimeoutSecs      pulumi.IntPtrInput                                    `pulumi:"metadataRetrievalTimeoutSecs"`
	MismatchingMimeTypesOverrideList  pulumi.StringPtrInput                                 `pulumi:"mismatchingMimeTypesOverrideList"`
	MissedCachePeriodSeconds          pulumi.IntPtrInput                                    `pulumi:"missedCachePeriodSeconds"`
	Notes                             pulumi.StringPtrInput                                 `pulumi:"notes"`
	Offline                           pulumi.BoolPtrInput                                   `pulumi:"offline"`
	Password                          pulumi.StringPtrInput                                 `pulumi:"password"`
	PriorityResolution                pulumi.BoolPtrInput                                   `pulumi:"priorityResolution"`
	ProjectEnvironments               pulumi.StringArrayInput                               `pulumi:"projectEnvironments"`
	ProjectKey                        pulumi.StringPtrInput                                 `pulumi:"projectKey"`
	PropertySets                      pulumi.StringArrayInput                               `pulumi:"propertySets"`
	Proxy                             pulumi.StringPtrInput                                 `pulumi:"proxy"`
	PypiRegistryUrl                   pulumi.StringPtrInput                                 `pulumi:"pypiRegistryUrl"`
	PypiRepositorySuffix              pulumi.StringPtrInput                                 `pulumi:"pypiRepositorySuffix"`
	QueryParams                       pulumi.StringPtrInput                                 `pulumi:"queryParams"`
	RemoteRepoLayoutRef               pulumi.StringPtrInput                                 `pulumi:"remoteRepoLayoutRef"`
	RepoLayoutRef                     pulumi.StringPtrInput                                 `pulumi:"repoLayoutRef"`
	RetrievalCachePeriodSeconds       pulumi.IntPtrInput                                    `pulumi:"retrievalCachePeriodSeconds"`
	ShareConfiguration                pulumi.BoolPtrInput                                   `pulumi:"shareConfiguration"`
	SocketTimeoutMillis               pulumi.IntPtrInput                                    `pulumi:"socketTimeoutMillis"`
	StoreArtifactsLocally             pulumi.BoolPtrInput                                   `pulumi:"storeArtifactsLocally"`
	SynchronizeProperties             pulumi.BoolPtrInput                                   `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodHours pulumi.IntPtrInput                                    `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                               pulumi.StringPtrInput                                 `pulumi:"url"`
	Username                          pulumi.StringPtrInput                                 `pulumi:"username"`
	XrayIndex                         pulumi.BoolPtrInput                                   `pulumi:"xrayIndex"`
}

func (LookupRemotePypiRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemotePypiRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getRemotePypiRepository.
type LookupRemotePypiRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupRemotePypiRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemotePypiRepositoryResult)(nil)).Elem()
}

func (o LookupRemotePypiRepositoryResultOutput) ToLookupRemotePypiRepositoryResultOutput() LookupRemotePypiRepositoryResultOutput {
	return o
}

func (o LookupRemotePypiRepositoryResultOutput) ToLookupRemotePypiRepositoryResultOutputWithContext(ctx context.Context) LookupRemotePypiRepositoryResultOutput {
	return o
}

func (o LookupRemotePypiRepositoryResultOutput) AllowAnyHostAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *bool { return v.AllowAnyHostAuth }).(pulumi.BoolPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) AssumedOfflinePeriodSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *int { return v.AssumedOfflinePeriodSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) BlockMismatchingMimeTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *bool { return v.BlockMismatchingMimeTypes }).(pulumi.BoolPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) BypassHeadRequests() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *bool { return v.BypassHeadRequests }).(pulumi.BoolPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) ClientTlsCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) string { return v.ClientTlsCertificate }).(pulumi.StringOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) ContentSynchronisation() GetRemotePypiRepositoryContentSynchronisationOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) GetRemotePypiRepositoryContentSynchronisation {
		return v.ContentSynchronisation
	}).(GetRemotePypiRepositoryContentSynchronisationOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) EnableCookieManagement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *bool { return v.EnableCookieManagement }).(pulumi.BoolPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) HardFail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *bool { return v.HardFail }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRemotePypiRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) ListRemoteFolderItems() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *bool { return v.ListRemoteFolderItems }).(pulumi.BoolPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) LocalAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *string { return v.LocalAddress }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) MetadataRetrievalTimeoutSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *int { return v.MetadataRetrievalTimeoutSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) MismatchingMimeTypesOverrideList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *string { return v.MismatchingMimeTypesOverrideList }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) MissedCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *int { return v.MissedCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) Offline() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *bool { return v.Offline }).(pulumi.BoolPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *string { return v.Proxy }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) PypiRegistryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *string { return v.PypiRegistryUrl }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) PypiRepositorySuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *string { return v.PypiRepositorySuffix }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) QueryParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *string { return v.QueryParams }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) RemoteRepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *string { return v.RemoteRepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) RetrievalCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *int { return v.RetrievalCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) ShareConfiguration() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) bool { return v.ShareConfiguration }).(pulumi.BoolOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) SocketTimeoutMillis() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *int { return v.SocketTimeoutMillis }).(pulumi.IntPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) StoreArtifactsLocally() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *bool { return v.StoreArtifactsLocally }).(pulumi.BoolPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) SynchronizeProperties() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *bool { return v.SynchronizeProperties }).(pulumi.BoolPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) UnusedArtifactsCleanupPeriodHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *int { return v.UnusedArtifactsCleanupPeriodHours }).(pulumi.IntPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePypiRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePypiRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemotePypiRepositoryResultOutput{})
}
