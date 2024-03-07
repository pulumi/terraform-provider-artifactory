// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v6/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a remote Gradle repository.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := artifactory.LookupRemoteGradleRepository(ctx, &artifactory.LookupRemoteGradleRepositoryArgs{
//				Key: "remote-gradle",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupRemoteGradleRepository(ctx *pulumi.Context, args *LookupRemoteGradleRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupRemoteGradleRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRemoteGradleRepositoryResult
	err := ctx.Invoke("artifactory:index/getRemoteGradleRepository:getRemoteGradleRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRemoteGradleRepository.
type LookupRemoteGradleRepositoryArgs struct {
	AllowAnyHostAuth          *bool                                            `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  *int                                             `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool                                            `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool                                            `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool                                            `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool                                            `pulumi:"cdnRedirect"`
	ClientTlsCertificate      *string                                          `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    *GetRemoteGradleRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description               *string                                          `pulumi:"description"`
	DisableProxy              *bool                                            `pulumi:"disableProxy"`
	DisableUrlNormalization   *bool                                            `pulumi:"disableUrlNormalization"`
	DownloadDirect            *bool                                            `pulumi:"downloadDirect"`
	EnableCookieManagement    *bool                                            `pulumi:"enableCookieManagement"`
	ExcludesPattern           *string                                          `pulumi:"excludesPattern"`
	// (Optional, Default: `false`) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.
	FetchJarsEagerly *bool `pulumi:"fetchJarsEagerly"`
	// (Optional, Default: `false`) When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.
	FetchSourcesEagerly *bool `pulumi:"fetchSourcesEagerly"`
	// (Optional, Default: `true`) If set, Artifactory allows you to deploy release artifacts into this repository.
	HandleReleases *bool `pulumi:"handleReleases"`
	// (Optional, Default: `true`) If set, Artifactory allows you to deploy snapshot artifacts into this repository.
	HandleSnapshots *bool   `pulumi:"handleSnapshots"`
	HardFail        *bool   `pulumi:"hardFail"`
	IncludesPattern *string `pulumi:"includesPattern"`
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
	PropertySets                     []string `pulumi:"propertySets"`
	Proxy                            *string  `pulumi:"proxy"`
	QueryParams                      *string  `pulumi:"queryParams"`
	// (Optional, Default: `false`) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".
	RejectInvalidJars *bool `pulumi:"rejectInvalidJars"`
	// (Optional, Default: `generate-if-absent`) Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are `generate-if-absent`, `fail`, `ignore-and-generate`, and `pass-thru`.
	RemoteRepoChecksumPolicyType *string `pulumi:"remoteRepoChecksumPolicyType"`
	RemoteRepoLayoutRef          *string `pulumi:"remoteRepoLayoutRef"`
	RepoLayoutRef                *string `pulumi:"repoLayoutRef"`
	RetrievalCachePeriodSeconds  *int    `pulumi:"retrievalCachePeriodSeconds"`
	ShareConfiguration           *bool   `pulumi:"shareConfiguration"`
	SocketTimeoutMillis          *int    `pulumi:"socketTimeoutMillis"`
	StoreArtifactsLocally        *bool   `pulumi:"storeArtifactsLocally"`
	// (Optional, Default: `true`) - By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to `true`.
	SuppressPomConsistencyChecks      *bool   `pulumi:"suppressPomConsistencyChecks"`
	SynchronizeProperties             *bool   `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodHours *int    `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                               *string `pulumi:"url"`
	Username                          *string `pulumi:"username"`
	XrayIndex                         *bool   `pulumi:"xrayIndex"`
}

// A collection of values returned by getRemoteGradleRepository.
type LookupRemoteGradleRepositoryResult struct {
	AllowAnyHostAuth          *bool                                           `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  *int                                            `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool                                           `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool                                           `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool                                           `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool                                           `pulumi:"cdnRedirect"`
	ClientTlsCertificate      string                                          `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    GetRemoteGradleRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description               *string                                         `pulumi:"description"`
	DisableProxy              *bool                                           `pulumi:"disableProxy"`
	DisableUrlNormalization   *bool                                           `pulumi:"disableUrlNormalization"`
	DownloadDirect            *bool                                           `pulumi:"downloadDirect"`
	EnableCookieManagement    *bool                                           `pulumi:"enableCookieManagement"`
	ExcludesPattern           *string                                         `pulumi:"excludesPattern"`
	// (Optional, Default: `false`) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.
	FetchJarsEagerly *bool `pulumi:"fetchJarsEagerly"`
	// (Optional, Default: `false`) When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.
	FetchSourcesEagerly *bool `pulumi:"fetchSourcesEagerly"`
	// (Optional, Default: `true`) If set, Artifactory allows you to deploy release artifacts into this repository.
	HandleReleases *bool `pulumi:"handleReleases"`
	// (Optional, Default: `true`) If set, Artifactory allows you to deploy snapshot artifacts into this repository.
	HandleSnapshots *bool `pulumi:"handleSnapshots"`
	HardFail        *bool `pulumi:"hardFail"`
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
	PropertySets                     []string `pulumi:"propertySets"`
	Proxy                            *string  `pulumi:"proxy"`
	QueryParams                      *string  `pulumi:"queryParams"`
	// (Optional, Default: `false`) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".
	RejectInvalidJars *bool `pulumi:"rejectInvalidJars"`
	// (Optional, Default: `generate-if-absent`) Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are `generate-if-absent`, `fail`, `ignore-and-generate`, and `pass-thru`.
	RemoteRepoChecksumPolicyType *string `pulumi:"remoteRepoChecksumPolicyType"`
	RemoteRepoLayoutRef          *string `pulumi:"remoteRepoLayoutRef"`
	RepoLayoutRef                *string `pulumi:"repoLayoutRef"`
	RetrievalCachePeriodSeconds  *int    `pulumi:"retrievalCachePeriodSeconds"`
	ShareConfiguration           bool    `pulumi:"shareConfiguration"`
	SocketTimeoutMillis          *int    `pulumi:"socketTimeoutMillis"`
	StoreArtifactsLocally        *bool   `pulumi:"storeArtifactsLocally"`
	// (Optional, Default: `true`) - By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to `true`.
	SuppressPomConsistencyChecks      *bool   `pulumi:"suppressPomConsistencyChecks"`
	SynchronizeProperties             *bool   `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodHours *int    `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                               *string `pulumi:"url"`
	Username                          *string `pulumi:"username"`
	XrayIndex                         *bool   `pulumi:"xrayIndex"`
}

func LookupRemoteGradleRepositoryOutput(ctx *pulumi.Context, args LookupRemoteGradleRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupRemoteGradleRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRemoteGradleRepositoryResult, error) {
			args := v.(LookupRemoteGradleRepositoryArgs)
			r, err := LookupRemoteGradleRepository(ctx, &args, opts...)
			var s LookupRemoteGradleRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRemoteGradleRepositoryResultOutput)
}

// A collection of arguments for invoking getRemoteGradleRepository.
type LookupRemoteGradleRepositoryOutputArgs struct {
	AllowAnyHostAuth          pulumi.BoolPtrInput                                     `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  pulumi.IntPtrInput                                      `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                pulumi.BoolPtrInput                                     `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes pulumi.BoolPtrInput                                     `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        pulumi.BoolPtrInput                                     `pulumi:"bypassHeadRequests"`
	CdnRedirect               pulumi.BoolPtrInput                                     `pulumi:"cdnRedirect"`
	ClientTlsCertificate      pulumi.StringPtrInput                                   `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    GetRemoteGradleRepositoryContentSynchronisationPtrInput `pulumi:"contentSynchronisation"`
	Description               pulumi.StringPtrInput                                   `pulumi:"description"`
	DisableProxy              pulumi.BoolPtrInput                                     `pulumi:"disableProxy"`
	DisableUrlNormalization   pulumi.BoolPtrInput                                     `pulumi:"disableUrlNormalization"`
	DownloadDirect            pulumi.BoolPtrInput                                     `pulumi:"downloadDirect"`
	EnableCookieManagement    pulumi.BoolPtrInput                                     `pulumi:"enableCookieManagement"`
	ExcludesPattern           pulumi.StringPtrInput                                   `pulumi:"excludesPattern"`
	// (Optional, Default: `false`) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.
	FetchJarsEagerly pulumi.BoolPtrInput `pulumi:"fetchJarsEagerly"`
	// (Optional, Default: `false`) When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.
	FetchSourcesEagerly pulumi.BoolPtrInput `pulumi:"fetchSourcesEagerly"`
	// (Optional, Default: `true`) If set, Artifactory allows you to deploy release artifacts into this repository.
	HandleReleases pulumi.BoolPtrInput `pulumi:"handleReleases"`
	// (Optional, Default: `true`) If set, Artifactory allows you to deploy snapshot artifacts into this repository.
	HandleSnapshots pulumi.BoolPtrInput   `pulumi:"handleSnapshots"`
	HardFail        pulumi.BoolPtrInput   `pulumi:"hardFail"`
	IncludesPattern pulumi.StringPtrInput `pulumi:"includesPattern"`
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
	PropertySets                     pulumi.StringArrayInput `pulumi:"propertySets"`
	Proxy                            pulumi.StringPtrInput   `pulumi:"proxy"`
	QueryParams                      pulumi.StringPtrInput   `pulumi:"queryParams"`
	// (Optional, Default: `false`) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".
	RejectInvalidJars pulumi.BoolPtrInput `pulumi:"rejectInvalidJars"`
	// (Optional, Default: `generate-if-absent`) Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are `generate-if-absent`, `fail`, `ignore-and-generate`, and `pass-thru`.
	RemoteRepoChecksumPolicyType pulumi.StringPtrInput `pulumi:"remoteRepoChecksumPolicyType"`
	RemoteRepoLayoutRef          pulumi.StringPtrInput `pulumi:"remoteRepoLayoutRef"`
	RepoLayoutRef                pulumi.StringPtrInput `pulumi:"repoLayoutRef"`
	RetrievalCachePeriodSeconds  pulumi.IntPtrInput    `pulumi:"retrievalCachePeriodSeconds"`
	ShareConfiguration           pulumi.BoolPtrInput   `pulumi:"shareConfiguration"`
	SocketTimeoutMillis          pulumi.IntPtrInput    `pulumi:"socketTimeoutMillis"`
	StoreArtifactsLocally        pulumi.BoolPtrInput   `pulumi:"storeArtifactsLocally"`
	// (Optional, Default: `true`) - By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to `true`.
	SuppressPomConsistencyChecks      pulumi.BoolPtrInput   `pulumi:"suppressPomConsistencyChecks"`
	SynchronizeProperties             pulumi.BoolPtrInput   `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodHours pulumi.IntPtrInput    `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                               pulumi.StringPtrInput `pulumi:"url"`
	Username                          pulumi.StringPtrInput `pulumi:"username"`
	XrayIndex                         pulumi.BoolPtrInput   `pulumi:"xrayIndex"`
}

func (LookupRemoteGradleRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteGradleRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getRemoteGradleRepository.
type LookupRemoteGradleRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupRemoteGradleRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteGradleRepositoryResult)(nil)).Elem()
}

func (o LookupRemoteGradleRepositoryResultOutput) ToLookupRemoteGradleRepositoryResultOutput() LookupRemoteGradleRepositoryResultOutput {
	return o
}

func (o LookupRemoteGradleRepositoryResultOutput) ToLookupRemoteGradleRepositoryResultOutputWithContext(ctx context.Context) LookupRemoteGradleRepositoryResultOutput {
	return o
}

func (o LookupRemoteGradleRepositoryResultOutput) AllowAnyHostAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *bool { return v.AllowAnyHostAuth }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) AssumedOfflinePeriodSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *int { return v.AssumedOfflinePeriodSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) BlockMismatchingMimeTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *bool { return v.BlockMismatchingMimeTypes }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) BypassHeadRequests() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *bool { return v.BypassHeadRequests }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) ClientTlsCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) string { return v.ClientTlsCertificate }).(pulumi.StringOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) ContentSynchronisation() GetRemoteGradleRepositoryContentSynchronisationOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) GetRemoteGradleRepositoryContentSynchronisation {
		return v.ContentSynchronisation
	}).(GetRemoteGradleRepositoryContentSynchronisationOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) DisableProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *bool { return v.DisableProxy }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) DisableUrlNormalization() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *bool { return v.DisableUrlNormalization }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) EnableCookieManagement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *bool { return v.EnableCookieManagement }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// (Optional, Default: `false`) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.
func (o LookupRemoteGradleRepositoryResultOutput) FetchJarsEagerly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *bool { return v.FetchJarsEagerly }).(pulumi.BoolPtrOutput)
}

// (Optional, Default: `false`) When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.
func (o LookupRemoteGradleRepositoryResultOutput) FetchSourcesEagerly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *bool { return v.FetchSourcesEagerly }).(pulumi.BoolPtrOutput)
}

// (Optional, Default: `true`) If set, Artifactory allows you to deploy release artifacts into this repository.
func (o LookupRemoteGradleRepositoryResultOutput) HandleReleases() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *bool { return v.HandleReleases }).(pulumi.BoolPtrOutput)
}

// (Optional, Default: `true`) If set, Artifactory allows you to deploy snapshot artifacts into this repository.
func (o LookupRemoteGradleRepositoryResultOutput) HandleSnapshots() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *bool { return v.HandleSnapshots }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) HardFail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *bool { return v.HardFail }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRemoteGradleRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) ListRemoteFolderItems() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *bool { return v.ListRemoteFolderItems }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) LocalAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *string { return v.LocalAddress }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) MetadataRetrievalTimeoutSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *int { return v.MetadataRetrievalTimeoutSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) MismatchingMimeTypesOverrideList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *string { return v.MismatchingMimeTypesOverrideList }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) MissedCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *int { return v.MissedCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) Offline() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *bool { return v.Offline }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *string { return v.Proxy }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) QueryParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *string { return v.QueryParams }).(pulumi.StringPtrOutput)
}

// (Optional, Default: `false`) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".
func (o LookupRemoteGradleRepositoryResultOutput) RejectInvalidJars() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *bool { return v.RejectInvalidJars }).(pulumi.BoolPtrOutput)
}

// (Optional, Default: `generate-if-absent`) Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are `generate-if-absent`, `fail`, `ignore-and-generate`, and `pass-thru`.
func (o LookupRemoteGradleRepositoryResultOutput) RemoteRepoChecksumPolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *string { return v.RemoteRepoChecksumPolicyType }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) RemoteRepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *string { return v.RemoteRepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) RetrievalCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *int { return v.RetrievalCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) ShareConfiguration() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) bool { return v.ShareConfiguration }).(pulumi.BoolOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) SocketTimeoutMillis() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *int { return v.SocketTimeoutMillis }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) StoreArtifactsLocally() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *bool { return v.StoreArtifactsLocally }).(pulumi.BoolPtrOutput)
}

// (Optional, Default: `true`) - By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to `true`.
func (o LookupRemoteGradleRepositoryResultOutput) SuppressPomConsistencyChecks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *bool { return v.SuppressPomConsistencyChecks }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) SynchronizeProperties() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *bool { return v.SynchronizeProperties }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) UnusedArtifactsCleanupPeriodHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *int { return v.UnusedArtifactsCleanupPeriodHours }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteGradleRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteGradleRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemoteGradleRepositoryResultOutput{})
}
