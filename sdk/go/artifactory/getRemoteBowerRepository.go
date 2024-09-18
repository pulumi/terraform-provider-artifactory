// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a remote Bower repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.LookupRemoteBowerRepository(ctx, &artifactory.LookupRemoteBowerRepositoryArgs{
//				Key: "remote-bower",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRemoteBowerRepository(ctx *pulumi.Context, args *LookupRemoteBowerRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupRemoteBowerRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRemoteBowerRepositoryResult
	err := ctx.Invoke("artifactory:index/getRemoteBowerRepository:getRemoteBowerRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRemoteBowerRepository.
type LookupRemoteBowerRepositoryArgs struct {
	AllowAnyHostAuth          *bool `pulumi:"allowAnyHostAuth"`
	ArchiveBrowsingEnabled    *bool `pulumi:"archiveBrowsingEnabled"`
	AssumedOfflinePeriodSecs  *int  `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool `pulumi:"blockMismatchingMimeTypes"`
	// (Optional) Proxy remote Bower repository. Default value is `https://registry.bower.io`.
	BowerRegistryUrl        *string                                         `pulumi:"bowerRegistryUrl"`
	BypassHeadRequests      *bool                                           `pulumi:"bypassHeadRequests"`
	CdnRedirect             *bool                                           `pulumi:"cdnRedirect"`
	ClientTlsCertificate    *string                                         `pulumi:"clientTlsCertificate"`
	ContentSynchronisation  *GetRemoteBowerRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description             *string                                         `pulumi:"description"`
	DisableProxy            *bool                                           `pulumi:"disableProxy"`
	DisableUrlNormalization *bool                                           `pulumi:"disableUrlNormalization"`
	DownloadDirect          *bool                                           `pulumi:"downloadDirect"`
	EnableCookieManagement  *bool                                           `pulumi:"enableCookieManagement"`
	ExcludesPattern         *string                                         `pulumi:"excludesPattern"`
	HardFail                *bool                                           `pulumi:"hardFail"`
	IncludesPattern         *string                                         `pulumi:"includesPattern"`
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
	VcsGitDownloadUrl                 *string  `pulumi:"vcsGitDownloadUrl"`
	VcsGitProvider                    *string  `pulumi:"vcsGitProvider"`
	XrayIndex                         *bool    `pulumi:"xrayIndex"`
}

// A collection of values returned by getRemoteBowerRepository.
type LookupRemoteBowerRepositoryResult struct {
	AllowAnyHostAuth          *bool `pulumi:"allowAnyHostAuth"`
	ArchiveBrowsingEnabled    *bool `pulumi:"archiveBrowsingEnabled"`
	AssumedOfflinePeriodSecs  *int  `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool `pulumi:"blockMismatchingMimeTypes"`
	// (Optional) Proxy remote Bower repository. Default value is `https://registry.bower.io`.
	BowerRegistryUrl        *string                                        `pulumi:"bowerRegistryUrl"`
	BypassHeadRequests      *bool                                          `pulumi:"bypassHeadRequests"`
	CdnRedirect             *bool                                          `pulumi:"cdnRedirect"`
	ClientTlsCertificate    string                                         `pulumi:"clientTlsCertificate"`
	ContentSynchronisation  GetRemoteBowerRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description             *string                                        `pulumi:"description"`
	DisableProxy            *bool                                          `pulumi:"disableProxy"`
	DisableUrlNormalization *bool                                          `pulumi:"disableUrlNormalization"`
	DownloadDirect          *bool                                          `pulumi:"downloadDirect"`
	EnableCookieManagement  *bool                                          `pulumi:"enableCookieManagement"`
	ExcludesPattern         *string                                        `pulumi:"excludesPattern"`
	HardFail                *bool                                          `pulumi:"hardFail"`
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
	VcsGitDownloadUrl                 *string  `pulumi:"vcsGitDownloadUrl"`
	VcsGitProvider                    *string  `pulumi:"vcsGitProvider"`
	XrayIndex                         *bool    `pulumi:"xrayIndex"`
}

func LookupRemoteBowerRepositoryOutput(ctx *pulumi.Context, args LookupRemoteBowerRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupRemoteBowerRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRemoteBowerRepositoryResult, error) {
			args := v.(LookupRemoteBowerRepositoryArgs)
			r, err := LookupRemoteBowerRepository(ctx, &args, opts...)
			var s LookupRemoteBowerRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRemoteBowerRepositoryResultOutput)
}

// A collection of arguments for invoking getRemoteBowerRepository.
type LookupRemoteBowerRepositoryOutputArgs struct {
	AllowAnyHostAuth          pulumi.BoolPtrInput `pulumi:"allowAnyHostAuth"`
	ArchiveBrowsingEnabled    pulumi.BoolPtrInput `pulumi:"archiveBrowsingEnabled"`
	AssumedOfflinePeriodSecs  pulumi.IntPtrInput  `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                pulumi.BoolPtrInput `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes pulumi.BoolPtrInput `pulumi:"blockMismatchingMimeTypes"`
	// (Optional) Proxy remote Bower repository. Default value is `https://registry.bower.io`.
	BowerRegistryUrl        pulumi.StringPtrInput                                  `pulumi:"bowerRegistryUrl"`
	BypassHeadRequests      pulumi.BoolPtrInput                                    `pulumi:"bypassHeadRequests"`
	CdnRedirect             pulumi.BoolPtrInput                                    `pulumi:"cdnRedirect"`
	ClientTlsCertificate    pulumi.StringPtrInput                                  `pulumi:"clientTlsCertificate"`
	ContentSynchronisation  GetRemoteBowerRepositoryContentSynchronisationPtrInput `pulumi:"contentSynchronisation"`
	Description             pulumi.StringPtrInput                                  `pulumi:"description"`
	DisableProxy            pulumi.BoolPtrInput                                    `pulumi:"disableProxy"`
	DisableUrlNormalization pulumi.BoolPtrInput                                    `pulumi:"disableUrlNormalization"`
	DownloadDirect          pulumi.BoolPtrInput                                    `pulumi:"downloadDirect"`
	EnableCookieManagement  pulumi.BoolPtrInput                                    `pulumi:"enableCookieManagement"`
	ExcludesPattern         pulumi.StringPtrInput                                  `pulumi:"excludesPattern"`
	HardFail                pulumi.BoolPtrInput                                    `pulumi:"hardFail"`
	IncludesPattern         pulumi.StringPtrInput                                  `pulumi:"includesPattern"`
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
	VcsGitDownloadUrl                 pulumi.StringPtrInput   `pulumi:"vcsGitDownloadUrl"`
	VcsGitProvider                    pulumi.StringPtrInput   `pulumi:"vcsGitProvider"`
	XrayIndex                         pulumi.BoolPtrInput     `pulumi:"xrayIndex"`
}

func (LookupRemoteBowerRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteBowerRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getRemoteBowerRepository.
type LookupRemoteBowerRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupRemoteBowerRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteBowerRepositoryResult)(nil)).Elem()
}

func (o LookupRemoteBowerRepositoryResultOutput) ToLookupRemoteBowerRepositoryResultOutput() LookupRemoteBowerRepositoryResultOutput {
	return o
}

func (o LookupRemoteBowerRepositoryResultOutput) ToLookupRemoteBowerRepositoryResultOutputWithContext(ctx context.Context) LookupRemoteBowerRepositoryResultOutput {
	return o
}

func (o LookupRemoteBowerRepositoryResultOutput) AllowAnyHostAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *bool { return v.AllowAnyHostAuth }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) AssumedOfflinePeriodSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *int { return v.AssumedOfflinePeriodSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) BlockMismatchingMimeTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *bool { return v.BlockMismatchingMimeTypes }).(pulumi.BoolPtrOutput)
}

// (Optional) Proxy remote Bower repository. Default value is `https://registry.bower.io`.
func (o LookupRemoteBowerRepositoryResultOutput) BowerRegistryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *string { return v.BowerRegistryUrl }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) BypassHeadRequests() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *bool { return v.BypassHeadRequests }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) ClientTlsCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) string { return v.ClientTlsCertificate }).(pulumi.StringOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) ContentSynchronisation() GetRemoteBowerRepositoryContentSynchronisationOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) GetRemoteBowerRepositoryContentSynchronisation {
		return v.ContentSynchronisation
	}).(GetRemoteBowerRepositoryContentSynchronisationOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) DisableProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *bool { return v.DisableProxy }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) DisableUrlNormalization() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *bool { return v.DisableUrlNormalization }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) EnableCookieManagement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *bool { return v.EnableCookieManagement }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) HardFail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *bool { return v.HardFail }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRemoteBowerRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) ListRemoteFolderItems() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *bool { return v.ListRemoteFolderItems }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) LocalAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *string { return v.LocalAddress }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) MetadataRetrievalTimeoutSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *int { return v.MetadataRetrievalTimeoutSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) MismatchingMimeTypesOverrideList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *string { return v.MismatchingMimeTypesOverrideList }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) MissedCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *int { return v.MissedCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) Offline() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *bool { return v.Offline }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *string { return v.Proxy }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) QueryParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *string { return v.QueryParams }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) RemoteRepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *string { return v.RemoteRepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) RetrievalCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *int { return v.RetrievalCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) ShareConfiguration() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) bool { return v.ShareConfiguration }).(pulumi.BoolOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) SocketTimeoutMillis() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *int { return v.SocketTimeoutMillis }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) StoreArtifactsLocally() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *bool { return v.StoreArtifactsLocally }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) SynchronizeProperties() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *bool { return v.SynchronizeProperties }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) UnusedArtifactsCleanupPeriodHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *int { return v.UnusedArtifactsCleanupPeriodHours }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) VcsGitDownloadUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *string { return v.VcsGitDownloadUrl }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) VcsGitProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *string { return v.VcsGitProvider }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteBowerRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteBowerRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemoteBowerRepositoryResultOutput{})
}
