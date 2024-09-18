// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a remote Maven repository.
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
//			_, err := artifactory.LookupRemoteMavenRepository(ctx, &artifactory.LookupRemoteMavenRepositoryArgs{
//				Key: "remote-maven",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRemoteMavenRepository(ctx *pulumi.Context, args *LookupRemoteMavenRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupRemoteMavenRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRemoteMavenRepositoryResult
	err := ctx.Invoke("artifactory:index/getRemoteMavenRepository:getRemoteMavenRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRemoteMavenRepository.
type LookupRemoteMavenRepositoryArgs struct {
	AllowAnyHostAuth          *bool                                           `pulumi:"allowAnyHostAuth"`
	ArchiveBrowsingEnabled    *bool                                           `pulumi:"archiveBrowsingEnabled"`
	AssumedOfflinePeriodSecs  *int                                            `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool                                           `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool                                           `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool                                           `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool                                           `pulumi:"cdnRedirect"`
	ClientTlsCertificate      *string                                         `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    *GetRemoteMavenRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Curated                   *bool                                           `pulumi:"curated"`
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
	HandleSnapshots *bool   `pulumi:"handleSnapshots"`
	HardFail        *bool   `pulumi:"hardFail"`
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                   string  `pulumi:"key"`
	ListRemoteFolderItems *bool   `pulumi:"listRemoteFolderItems"`
	LocalAddress          *string `pulumi:"localAddress"`
	MaxUniqueSnapshots    *int    `pulumi:"maxUniqueSnapshots"`
	// (Optional, Default: 60) This value refers to the number of seconds to wait for retrieval from the remote before serving locally cached artifact or fail the request. Can not set to 0 or negative.
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
	// (Optional, Default: `true`) By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to `true`.
	SuppressPomConsistencyChecks      *bool   `pulumi:"suppressPomConsistencyChecks"`
	SynchronizeProperties             *bool   `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodHours *int    `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                               *string `pulumi:"url"`
	Username                          *string `pulumi:"username"`
	XrayIndex                         *bool   `pulumi:"xrayIndex"`
}

// A collection of values returned by getRemoteMavenRepository.
type LookupRemoteMavenRepositoryResult struct {
	AllowAnyHostAuth          *bool                                          `pulumi:"allowAnyHostAuth"`
	ArchiveBrowsingEnabled    *bool                                          `pulumi:"archiveBrowsingEnabled"`
	AssumedOfflinePeriodSecs  *int                                           `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool                                          `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool                                          `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool                                          `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool                                          `pulumi:"cdnRedirect"`
	ClientTlsCertificate      string                                         `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    GetRemoteMavenRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Curated                   *bool                                          `pulumi:"curated"`
	Description               *string                                        `pulumi:"description"`
	DisableProxy              *bool                                          `pulumi:"disableProxy"`
	DisableUrlNormalization   *bool                                          `pulumi:"disableUrlNormalization"`
	DownloadDirect            *bool                                          `pulumi:"downloadDirect"`
	EnableCookieManagement    *bool                                          `pulumi:"enableCookieManagement"`
	ExcludesPattern           *string                                        `pulumi:"excludesPattern"`
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
	Id                    string  `pulumi:"id"`
	IncludesPattern       *string `pulumi:"includesPattern"`
	Key                   string  `pulumi:"key"`
	ListRemoteFolderItems *bool   `pulumi:"listRemoteFolderItems"`
	LocalAddress          *string `pulumi:"localAddress"`
	MaxUniqueSnapshots    *int    `pulumi:"maxUniqueSnapshots"`
	// (Optional, Default: 60) This value refers to the number of seconds to wait for retrieval from the remote before serving locally cached artifact or fail the request. Can not set to 0 or negative.
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
	// (Optional, Default: `true`) By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to `true`.
	SuppressPomConsistencyChecks      *bool   `pulumi:"suppressPomConsistencyChecks"`
	SynchronizeProperties             *bool   `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodHours *int    `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                               *string `pulumi:"url"`
	Username                          *string `pulumi:"username"`
	XrayIndex                         *bool   `pulumi:"xrayIndex"`
}

func LookupRemoteMavenRepositoryOutput(ctx *pulumi.Context, args LookupRemoteMavenRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupRemoteMavenRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRemoteMavenRepositoryResult, error) {
			args := v.(LookupRemoteMavenRepositoryArgs)
			r, err := LookupRemoteMavenRepository(ctx, &args, opts...)
			var s LookupRemoteMavenRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRemoteMavenRepositoryResultOutput)
}

// A collection of arguments for invoking getRemoteMavenRepository.
type LookupRemoteMavenRepositoryOutputArgs struct {
	AllowAnyHostAuth          pulumi.BoolPtrInput                                    `pulumi:"allowAnyHostAuth"`
	ArchiveBrowsingEnabled    pulumi.BoolPtrInput                                    `pulumi:"archiveBrowsingEnabled"`
	AssumedOfflinePeriodSecs  pulumi.IntPtrInput                                     `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                pulumi.BoolPtrInput                                    `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes pulumi.BoolPtrInput                                    `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        pulumi.BoolPtrInput                                    `pulumi:"bypassHeadRequests"`
	CdnRedirect               pulumi.BoolPtrInput                                    `pulumi:"cdnRedirect"`
	ClientTlsCertificate      pulumi.StringPtrInput                                  `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    GetRemoteMavenRepositoryContentSynchronisationPtrInput `pulumi:"contentSynchronisation"`
	Curated                   pulumi.BoolPtrInput                                    `pulumi:"curated"`
	Description               pulumi.StringPtrInput                                  `pulumi:"description"`
	DisableProxy              pulumi.BoolPtrInput                                    `pulumi:"disableProxy"`
	DisableUrlNormalization   pulumi.BoolPtrInput                                    `pulumi:"disableUrlNormalization"`
	DownloadDirect            pulumi.BoolPtrInput                                    `pulumi:"downloadDirect"`
	EnableCookieManagement    pulumi.BoolPtrInput                                    `pulumi:"enableCookieManagement"`
	ExcludesPattern           pulumi.StringPtrInput                                  `pulumi:"excludesPattern"`
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
	Key                   pulumi.StringInput    `pulumi:"key"`
	ListRemoteFolderItems pulumi.BoolPtrInput   `pulumi:"listRemoteFolderItems"`
	LocalAddress          pulumi.StringPtrInput `pulumi:"localAddress"`
	MaxUniqueSnapshots    pulumi.IntPtrInput    `pulumi:"maxUniqueSnapshots"`
	// (Optional, Default: 60) This value refers to the number of seconds to wait for retrieval from the remote before serving locally cached artifact or fail the request. Can not set to 0 or negative.
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
	// (Optional, Default: `true`) By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to `true`.
	SuppressPomConsistencyChecks      pulumi.BoolPtrInput   `pulumi:"suppressPomConsistencyChecks"`
	SynchronizeProperties             pulumi.BoolPtrInput   `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodHours pulumi.IntPtrInput    `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                               pulumi.StringPtrInput `pulumi:"url"`
	Username                          pulumi.StringPtrInput `pulumi:"username"`
	XrayIndex                         pulumi.BoolPtrInput   `pulumi:"xrayIndex"`
}

func (LookupRemoteMavenRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteMavenRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getRemoteMavenRepository.
type LookupRemoteMavenRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupRemoteMavenRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteMavenRepositoryResult)(nil)).Elem()
}

func (o LookupRemoteMavenRepositoryResultOutput) ToLookupRemoteMavenRepositoryResultOutput() LookupRemoteMavenRepositoryResultOutput {
	return o
}

func (o LookupRemoteMavenRepositoryResultOutput) ToLookupRemoteMavenRepositoryResultOutputWithContext(ctx context.Context) LookupRemoteMavenRepositoryResultOutput {
	return o
}

func (o LookupRemoteMavenRepositoryResultOutput) AllowAnyHostAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *bool { return v.AllowAnyHostAuth }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) AssumedOfflinePeriodSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *int { return v.AssumedOfflinePeriodSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) BlockMismatchingMimeTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *bool { return v.BlockMismatchingMimeTypes }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) BypassHeadRequests() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *bool { return v.BypassHeadRequests }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) ClientTlsCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) string { return v.ClientTlsCertificate }).(pulumi.StringOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) ContentSynchronisation() GetRemoteMavenRepositoryContentSynchronisationOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) GetRemoteMavenRepositoryContentSynchronisation {
		return v.ContentSynchronisation
	}).(GetRemoteMavenRepositoryContentSynchronisationOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) Curated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *bool { return v.Curated }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) DisableProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *bool { return v.DisableProxy }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) DisableUrlNormalization() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *bool { return v.DisableUrlNormalization }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) EnableCookieManagement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *bool { return v.EnableCookieManagement }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// (Optional, Default: `false`) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.
func (o LookupRemoteMavenRepositoryResultOutput) FetchJarsEagerly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *bool { return v.FetchJarsEagerly }).(pulumi.BoolPtrOutput)
}

// (Optional, Default: `false`) When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.
func (o LookupRemoteMavenRepositoryResultOutput) FetchSourcesEagerly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *bool { return v.FetchSourcesEagerly }).(pulumi.BoolPtrOutput)
}

// (Optional, Default: `true`) If set, Artifactory allows you to deploy release artifacts into this repository.
func (o LookupRemoteMavenRepositoryResultOutput) HandleReleases() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *bool { return v.HandleReleases }).(pulumi.BoolPtrOutput)
}

// (Optional, Default: `true`) If set, Artifactory allows you to deploy snapshot artifacts into this repository.
func (o LookupRemoteMavenRepositoryResultOutput) HandleSnapshots() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *bool { return v.HandleSnapshots }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) HardFail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *bool { return v.HardFail }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRemoteMavenRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) ListRemoteFolderItems() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *bool { return v.ListRemoteFolderItems }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) LocalAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *string { return v.LocalAddress }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) MaxUniqueSnapshots() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *int { return v.MaxUniqueSnapshots }).(pulumi.IntPtrOutput)
}

// (Optional, Default: 60) This value refers to the number of seconds to wait for retrieval from the remote before serving locally cached artifact or fail the request. Can not set to 0 or negative.
func (o LookupRemoteMavenRepositoryResultOutput) MetadataRetrievalTimeoutSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *int { return v.MetadataRetrievalTimeoutSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) MismatchingMimeTypesOverrideList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *string { return v.MismatchingMimeTypesOverrideList }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) MissedCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *int { return v.MissedCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) Offline() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *bool { return v.Offline }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *string { return v.Proxy }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) QueryParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *string { return v.QueryParams }).(pulumi.StringPtrOutput)
}

// (Optional, Default: `false`) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".
func (o LookupRemoteMavenRepositoryResultOutput) RejectInvalidJars() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *bool { return v.RejectInvalidJars }).(pulumi.BoolPtrOutput)
}

// (Optional, Default: `generate-if-absent`) Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are `generate-if-absent`, `fail`, `ignore-and-generate`, and `pass-thru`.
func (o LookupRemoteMavenRepositoryResultOutput) RemoteRepoChecksumPolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *string { return v.RemoteRepoChecksumPolicyType }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) RemoteRepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *string { return v.RemoteRepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) RetrievalCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *int { return v.RetrievalCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) ShareConfiguration() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) bool { return v.ShareConfiguration }).(pulumi.BoolOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) SocketTimeoutMillis() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *int { return v.SocketTimeoutMillis }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) StoreArtifactsLocally() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *bool { return v.StoreArtifactsLocally }).(pulumi.BoolPtrOutput)
}

// (Optional, Default: `true`) By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to `true`.
func (o LookupRemoteMavenRepositoryResultOutput) SuppressPomConsistencyChecks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *bool { return v.SuppressPomConsistencyChecks }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) SynchronizeProperties() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *bool { return v.SynchronizeProperties }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) UnusedArtifactsCleanupPeriodHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *int { return v.UnusedArtifactsCleanupPeriodHours }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteMavenRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteMavenRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemoteMavenRepositoryResultOutput{})
}
