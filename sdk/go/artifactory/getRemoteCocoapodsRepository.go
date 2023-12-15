// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v6/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a remote CocoaPods repository.
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
//			_, err := artifactory.LookupRemoteCocoapodsRepository(ctx, &artifactory.LookupRemoteCocoapodsRepositoryArgs{
//				Key: "remote-cocoapods",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRemoteCocoapodsRepository(ctx *pulumi.Context, args *LookupRemoteCocoapodsRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupRemoteCocoapodsRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRemoteCocoapodsRepositoryResult
	err := ctx.Invoke("artifactory:index/getRemoteCocoapodsRepository:getRemoteCocoapodsRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRemoteCocoapodsRepository.
type LookupRemoteCocoapodsRepositoryArgs struct {
	AllowAnyHostAuth          *bool                                               `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  *int                                                `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool                                               `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool                                               `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool                                               `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool                                               `pulumi:"cdnRedirect"`
	ClientTlsCertificate      *string                                             `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    *GetRemoteCocoapodsRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description               *string                                             `pulumi:"description"`
	DisableProxy              *bool                                               `pulumi:"disableProxy"`
	DisableUrlNormalization   *bool                                               `pulumi:"disableUrlNormalization"`
	DownloadDirect            *bool                                               `pulumi:"downloadDirect"`
	EnableCookieManagement    *bool                                               `pulumi:"enableCookieManagement"`
	ExcludesPattern           *string                                             `pulumi:"excludesPattern"`
	HardFail                  *bool                                               `pulumi:"hardFail"`
	IncludesPattern           *string                                             `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                              string  `pulumi:"key"`
	ListRemoteFolderItems            *bool   `pulumi:"listRemoteFolderItems"`
	LocalAddress                     *string `pulumi:"localAddress"`
	MetadataRetrievalTimeoutSecs     *int    `pulumi:"metadataRetrievalTimeoutSecs"`
	MismatchingMimeTypesOverrideList *string `pulumi:"mismatchingMimeTypesOverrideList"`
	MissedCachePeriodSeconds         *int    `pulumi:"missedCachePeriodSeconds"`
	Notes                            *string `pulumi:"notes"`
	Offline                          *bool   `pulumi:"offline"`
	Password                         *string `pulumi:"password"`
	// (Optional) Proxy remote CocoaPods Specs repositories. Default value is `https://github.com/CocoaPods/Specs`.
	PodsSpecsRepoUrl                  *string  `pulumi:"podsSpecsRepoUrl"`
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
	// (Optional) This attribute is used when vcsGitProvider is set to `CUSTOM`. Provided URL will be used as proxy.
	VcsGitDownloadUrl *string `pulumi:"vcsGitDownloadUrl"`
	// (Optional) Artifactory supports proxying the following Git providers out-of-the-box: GitHub or a remote Artifactory instance. Default value is `GITHUB`. Possible values are: `GITHUB`, `BITBUCKET`, `OLDSTASH`, `STASH`, `ARTIFACTORY`, `CUSTOM`.
	VcsGitProvider *string `pulumi:"vcsGitProvider"`
	XrayIndex      *bool   `pulumi:"xrayIndex"`
}

// A collection of values returned by getRemoteCocoapodsRepository.
type LookupRemoteCocoapodsRepositoryResult struct {
	AllowAnyHostAuth          *bool                                              `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  *int                                               `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool                                              `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool                                              `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool                                              `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool                                              `pulumi:"cdnRedirect"`
	ClientTlsCertificate      string                                             `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    GetRemoteCocoapodsRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description               *string                                            `pulumi:"description"`
	DisableProxy              *bool                                              `pulumi:"disableProxy"`
	DisableUrlNormalization   *bool                                              `pulumi:"disableUrlNormalization"`
	DownloadDirect            *bool                                              `pulumi:"downloadDirect"`
	EnableCookieManagement    *bool                                              `pulumi:"enableCookieManagement"`
	ExcludesPattern           *string                                            `pulumi:"excludesPattern"`
	HardFail                  *bool                                              `pulumi:"hardFail"`
	// The provider-assigned unique ID for this managed resource.
	Id                               string  `pulumi:"id"`
	IncludesPattern                  *string `pulumi:"includesPattern"`
	Key                              string  `pulumi:"key"`
	ListRemoteFolderItems            *bool   `pulumi:"listRemoteFolderItems"`
	LocalAddress                     *string `pulumi:"localAddress"`
	MetadataRetrievalTimeoutSecs     *int    `pulumi:"metadataRetrievalTimeoutSecs"`
	MismatchingMimeTypesOverrideList *string `pulumi:"mismatchingMimeTypesOverrideList"`
	MissedCachePeriodSeconds         *int    `pulumi:"missedCachePeriodSeconds"`
	Notes                            *string `pulumi:"notes"`
	Offline                          *bool   `pulumi:"offline"`
	PackageType                      string  `pulumi:"packageType"`
	Password                         *string `pulumi:"password"`
	// (Optional) Proxy remote CocoaPods Specs repositories. Default value is `https://github.com/CocoaPods/Specs`.
	PodsSpecsRepoUrl                  *string  `pulumi:"podsSpecsRepoUrl"`
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
	// (Optional) This attribute is used when vcsGitProvider is set to `CUSTOM`. Provided URL will be used as proxy.
	VcsGitDownloadUrl *string `pulumi:"vcsGitDownloadUrl"`
	// (Optional) Artifactory supports proxying the following Git providers out-of-the-box: GitHub or a remote Artifactory instance. Default value is `GITHUB`. Possible values are: `GITHUB`, `BITBUCKET`, `OLDSTASH`, `STASH`, `ARTIFACTORY`, `CUSTOM`.
	VcsGitProvider *string `pulumi:"vcsGitProvider"`
	XrayIndex      *bool   `pulumi:"xrayIndex"`
}

func LookupRemoteCocoapodsRepositoryOutput(ctx *pulumi.Context, args LookupRemoteCocoapodsRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupRemoteCocoapodsRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRemoteCocoapodsRepositoryResult, error) {
			args := v.(LookupRemoteCocoapodsRepositoryArgs)
			r, err := LookupRemoteCocoapodsRepository(ctx, &args, opts...)
			var s LookupRemoteCocoapodsRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRemoteCocoapodsRepositoryResultOutput)
}

// A collection of arguments for invoking getRemoteCocoapodsRepository.
type LookupRemoteCocoapodsRepositoryOutputArgs struct {
	AllowAnyHostAuth          pulumi.BoolPtrInput                                        `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  pulumi.IntPtrInput                                         `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                pulumi.BoolPtrInput                                        `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes pulumi.BoolPtrInput                                        `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        pulumi.BoolPtrInput                                        `pulumi:"bypassHeadRequests"`
	CdnRedirect               pulumi.BoolPtrInput                                        `pulumi:"cdnRedirect"`
	ClientTlsCertificate      pulumi.StringPtrInput                                      `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    GetRemoteCocoapodsRepositoryContentSynchronisationPtrInput `pulumi:"contentSynchronisation"`
	Description               pulumi.StringPtrInput                                      `pulumi:"description"`
	DisableProxy              pulumi.BoolPtrInput                                        `pulumi:"disableProxy"`
	DisableUrlNormalization   pulumi.BoolPtrInput                                        `pulumi:"disableUrlNormalization"`
	DownloadDirect            pulumi.BoolPtrInput                                        `pulumi:"downloadDirect"`
	EnableCookieManagement    pulumi.BoolPtrInput                                        `pulumi:"enableCookieManagement"`
	ExcludesPattern           pulumi.StringPtrInput                                      `pulumi:"excludesPattern"`
	HardFail                  pulumi.BoolPtrInput                                        `pulumi:"hardFail"`
	IncludesPattern           pulumi.StringPtrInput                                      `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                              pulumi.StringInput    `pulumi:"key"`
	ListRemoteFolderItems            pulumi.BoolPtrInput   `pulumi:"listRemoteFolderItems"`
	LocalAddress                     pulumi.StringPtrInput `pulumi:"localAddress"`
	MetadataRetrievalTimeoutSecs     pulumi.IntPtrInput    `pulumi:"metadataRetrievalTimeoutSecs"`
	MismatchingMimeTypesOverrideList pulumi.StringPtrInput `pulumi:"mismatchingMimeTypesOverrideList"`
	MissedCachePeriodSeconds         pulumi.IntPtrInput    `pulumi:"missedCachePeriodSeconds"`
	Notes                            pulumi.StringPtrInput `pulumi:"notes"`
	Offline                          pulumi.BoolPtrInput   `pulumi:"offline"`
	Password                         pulumi.StringPtrInput `pulumi:"password"`
	// (Optional) Proxy remote CocoaPods Specs repositories. Default value is `https://github.com/CocoaPods/Specs`.
	PodsSpecsRepoUrl                  pulumi.StringPtrInput   `pulumi:"podsSpecsRepoUrl"`
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
	// (Optional) This attribute is used when vcsGitProvider is set to `CUSTOM`. Provided URL will be used as proxy.
	VcsGitDownloadUrl pulumi.StringPtrInput `pulumi:"vcsGitDownloadUrl"`
	// (Optional) Artifactory supports proxying the following Git providers out-of-the-box: GitHub or a remote Artifactory instance. Default value is `GITHUB`. Possible values are: `GITHUB`, `BITBUCKET`, `OLDSTASH`, `STASH`, `ARTIFACTORY`, `CUSTOM`.
	VcsGitProvider pulumi.StringPtrInput `pulumi:"vcsGitProvider"`
	XrayIndex      pulumi.BoolPtrInput   `pulumi:"xrayIndex"`
}

func (LookupRemoteCocoapodsRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteCocoapodsRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getRemoteCocoapodsRepository.
type LookupRemoteCocoapodsRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupRemoteCocoapodsRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteCocoapodsRepositoryResult)(nil)).Elem()
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) ToLookupRemoteCocoapodsRepositoryResultOutput() LookupRemoteCocoapodsRepositoryResultOutput {
	return o
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) ToLookupRemoteCocoapodsRepositoryResultOutputWithContext(ctx context.Context) LookupRemoteCocoapodsRepositoryResultOutput {
	return o
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) AllowAnyHostAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *bool { return v.AllowAnyHostAuth }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) AssumedOfflinePeriodSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *int { return v.AssumedOfflinePeriodSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) BlockMismatchingMimeTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *bool { return v.BlockMismatchingMimeTypes }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) BypassHeadRequests() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *bool { return v.BypassHeadRequests }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) ClientTlsCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) string { return v.ClientTlsCertificate }).(pulumi.StringOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) ContentSynchronisation() GetRemoteCocoapodsRepositoryContentSynchronisationOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) GetRemoteCocoapodsRepositoryContentSynchronisation {
		return v.ContentSynchronisation
	}).(GetRemoteCocoapodsRepositoryContentSynchronisationOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) DisableProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *bool { return v.DisableProxy }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) DisableUrlNormalization() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *bool { return v.DisableUrlNormalization }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) EnableCookieManagement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *bool { return v.EnableCookieManagement }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) HardFail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *bool { return v.HardFail }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRemoteCocoapodsRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) ListRemoteFolderItems() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *bool { return v.ListRemoteFolderItems }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) LocalAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *string { return v.LocalAddress }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) MetadataRetrievalTimeoutSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *int { return v.MetadataRetrievalTimeoutSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) MismatchingMimeTypesOverrideList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *string { return v.MismatchingMimeTypesOverrideList }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) MissedCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *int { return v.MissedCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) Offline() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *bool { return v.Offline }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

// (Optional) Proxy remote CocoaPods Specs repositories. Default value is `https://github.com/CocoaPods/Specs`.
func (o LookupRemoteCocoapodsRepositoryResultOutput) PodsSpecsRepoUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *string { return v.PodsSpecsRepoUrl }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *string { return v.Proxy }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) QueryParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *string { return v.QueryParams }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) RemoteRepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *string { return v.RemoteRepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) RetrievalCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *int { return v.RetrievalCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) ShareConfiguration() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) bool { return v.ShareConfiguration }).(pulumi.BoolOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) SocketTimeoutMillis() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *int { return v.SocketTimeoutMillis }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) StoreArtifactsLocally() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *bool { return v.StoreArtifactsLocally }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) SynchronizeProperties() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *bool { return v.SynchronizeProperties }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) UnusedArtifactsCleanupPeriodHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *int { return v.UnusedArtifactsCleanupPeriodHours }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

// (Optional) This attribute is used when vcsGitProvider is set to `CUSTOM`. Provided URL will be used as proxy.
func (o LookupRemoteCocoapodsRepositoryResultOutput) VcsGitDownloadUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *string { return v.VcsGitDownloadUrl }).(pulumi.StringPtrOutput)
}

// (Optional) Artifactory supports proxying the following Git providers out-of-the-box: GitHub or a remote Artifactory instance. Default value is `GITHUB`. Possible values are: `GITHUB`, `BITBUCKET`, `OLDSTASH`, `STASH`, `ARTIFACTORY`, `CUSTOM`.
func (o LookupRemoteCocoapodsRepositoryResultOutput) VcsGitProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *string { return v.VcsGitProvider }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteCocoapodsRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteCocoapodsRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemoteCocoapodsRepositoryResultOutput{})
}
