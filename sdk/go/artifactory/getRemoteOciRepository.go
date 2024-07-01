// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v7/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a remote OCI repository.
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
//			_, err := artifactory.LookupRemoteOciRepository(ctx, &artifactory.LookupRemoteOciRepositoryArgs{
//				Key: "my-oci-remote",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRemoteOciRepository(ctx *pulumi.Context, args *LookupRemoteOciRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupRemoteOciRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRemoteOciRepositoryResult
	err := ctx.Invoke("artifactory:index/getRemoteOciRepository:getRemoteOciRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRemoteOciRepository.
type LookupRemoteOciRepositoryArgs struct {
	AllowAnyHostAuth          *bool                                         `pulumi:"allowAnyHostAuth"`
	ArchiveBrowsingEnabled    *bool                                         `pulumi:"archiveBrowsingEnabled"`
	AssumedOfflinePeriodSecs  *int                                          `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool                                         `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool                                         `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool                                         `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool                                         `pulumi:"cdnRedirect"`
	ClientTlsCertificate      *string                                       `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    *GetRemoteOciRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description               *string                                       `pulumi:"description"`
	DisableProxy              *bool                                         `pulumi:"disableProxy"`
	// (Optional) Whether to disable URL normalization.
	DisableUrlNormalization *bool `pulumi:"disableUrlNormalization"`
	DownloadDirect          *bool `pulumi:"downloadDirect"`
	EnableCookieManagement  *bool `pulumi:"enableCookieManagement"`
	// (Optional) Enable token (Bearer) based authentication.
	EnableTokenAuthentication *bool   `pulumi:"enableTokenAuthentication"`
	ExcludesPattern           *string `pulumi:"excludesPattern"`
	// (Optional) Also known as 'Foreign Layers Caching' on the UI.
	ExternalDependenciesEnabled *bool `pulumi:"externalDependenciesEnabled"`
	// (Optional) Optional include patterns to match external URLs. Ant-style path expressions are supported (*, **, ?). For example, specifying `**/github.com/**` will only allow downloading foreign layers from github.com host. Due to SDKv2 limitations, we can't set the default value for the list. This value `[**]` must be assigned to the attribute manually, if user don't specify any other non-default values. We don't want to make this attribute required, but it must be set to avoid the state drift on update. Note: Artifactory assigns `[**]` on update if HCL doesn't have the attribute set or the list is empty.
	ExternalDependenciesPatterns []string `pulumi:"externalDependenciesPatterns"`
	HardFail                     *bool    `pulumi:"hardFail"`
	IncludesPattern              *string  `pulumi:"includesPattern"`
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
	ProjectId                         *string  `pulumi:"projectId"`
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

// A collection of values returned by getRemoteOciRepository.
type LookupRemoteOciRepositoryResult struct {
	AllowAnyHostAuth          *bool                                        `pulumi:"allowAnyHostAuth"`
	ArchiveBrowsingEnabled    *bool                                        `pulumi:"archiveBrowsingEnabled"`
	AssumedOfflinePeriodSecs  *int                                         `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool                                        `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool                                        `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool                                        `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool                                        `pulumi:"cdnRedirect"`
	ClientTlsCertificate      string                                       `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    GetRemoteOciRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description               *string                                      `pulumi:"description"`
	DisableProxy              *bool                                        `pulumi:"disableProxy"`
	// (Optional) Whether to disable URL normalization.
	DisableUrlNormalization *bool `pulumi:"disableUrlNormalization"`
	DownloadDirect          *bool `pulumi:"downloadDirect"`
	EnableCookieManagement  *bool `pulumi:"enableCookieManagement"`
	// (Optional) Enable token (Bearer) based authentication.
	EnableTokenAuthentication bool    `pulumi:"enableTokenAuthentication"`
	ExcludesPattern           *string `pulumi:"excludesPattern"`
	// (Optional) Also known as 'Foreign Layers Caching' on the UI.
	ExternalDependenciesEnabled *bool `pulumi:"externalDependenciesEnabled"`
	// (Optional) Optional include patterns to match external URLs. Ant-style path expressions are supported (*, **, ?). For example, specifying `**/github.com/**` will only allow downloading foreign layers from github.com host. Due to SDKv2 limitations, we can't set the default value for the list. This value `[**]` must be assigned to the attribute manually, if user don't specify any other non-default values. We don't want to make this attribute required, but it must be set to avoid the state drift on update. Note: Artifactory assigns `[**]` on update if HCL doesn't have the attribute set or the list is empty.
	ExternalDependenciesPatterns []string `pulumi:"externalDependenciesPatterns"`
	HardFail                     *bool    `pulumi:"hardFail"`
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
	ProjectId                         *string  `pulumi:"projectId"`
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

func LookupRemoteOciRepositoryOutput(ctx *pulumi.Context, args LookupRemoteOciRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupRemoteOciRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRemoteOciRepositoryResult, error) {
			args := v.(LookupRemoteOciRepositoryArgs)
			r, err := LookupRemoteOciRepository(ctx, &args, opts...)
			var s LookupRemoteOciRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRemoteOciRepositoryResultOutput)
}

// A collection of arguments for invoking getRemoteOciRepository.
type LookupRemoteOciRepositoryOutputArgs struct {
	AllowAnyHostAuth          pulumi.BoolPtrInput                                  `pulumi:"allowAnyHostAuth"`
	ArchiveBrowsingEnabled    pulumi.BoolPtrInput                                  `pulumi:"archiveBrowsingEnabled"`
	AssumedOfflinePeriodSecs  pulumi.IntPtrInput                                   `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                pulumi.BoolPtrInput                                  `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes pulumi.BoolPtrInput                                  `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        pulumi.BoolPtrInput                                  `pulumi:"bypassHeadRequests"`
	CdnRedirect               pulumi.BoolPtrInput                                  `pulumi:"cdnRedirect"`
	ClientTlsCertificate      pulumi.StringPtrInput                                `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    GetRemoteOciRepositoryContentSynchronisationPtrInput `pulumi:"contentSynchronisation"`
	Description               pulumi.StringPtrInput                                `pulumi:"description"`
	DisableProxy              pulumi.BoolPtrInput                                  `pulumi:"disableProxy"`
	// (Optional) Whether to disable URL normalization.
	DisableUrlNormalization pulumi.BoolPtrInput `pulumi:"disableUrlNormalization"`
	DownloadDirect          pulumi.BoolPtrInput `pulumi:"downloadDirect"`
	EnableCookieManagement  pulumi.BoolPtrInput `pulumi:"enableCookieManagement"`
	// (Optional) Enable token (Bearer) based authentication.
	EnableTokenAuthentication pulumi.BoolPtrInput   `pulumi:"enableTokenAuthentication"`
	ExcludesPattern           pulumi.StringPtrInput `pulumi:"excludesPattern"`
	// (Optional) Also known as 'Foreign Layers Caching' on the UI.
	ExternalDependenciesEnabled pulumi.BoolPtrInput `pulumi:"externalDependenciesEnabled"`
	// (Optional) Optional include patterns to match external URLs. Ant-style path expressions are supported (*, **, ?). For example, specifying `**/github.com/**` will only allow downloading foreign layers from github.com host. Due to SDKv2 limitations, we can't set the default value for the list. This value `[**]` must be assigned to the attribute manually, if user don't specify any other non-default values. We don't want to make this attribute required, but it must be set to avoid the state drift on update. Note: Artifactory assigns `[**]` on update if HCL doesn't have the attribute set or the list is empty.
	ExternalDependenciesPatterns pulumi.StringArrayInput `pulumi:"externalDependenciesPatterns"`
	HardFail                     pulumi.BoolPtrInput     `pulumi:"hardFail"`
	IncludesPattern              pulumi.StringPtrInput   `pulumi:"includesPattern"`
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
	ProjectId                         pulumi.StringPtrInput   `pulumi:"projectId"`
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

func (LookupRemoteOciRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteOciRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getRemoteOciRepository.
type LookupRemoteOciRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupRemoteOciRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteOciRepositoryResult)(nil)).Elem()
}

func (o LookupRemoteOciRepositoryResultOutput) ToLookupRemoteOciRepositoryResultOutput() LookupRemoteOciRepositoryResultOutput {
	return o
}

func (o LookupRemoteOciRepositoryResultOutput) ToLookupRemoteOciRepositoryResultOutputWithContext(ctx context.Context) LookupRemoteOciRepositoryResultOutput {
	return o
}

func (o LookupRemoteOciRepositoryResultOutput) AllowAnyHostAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *bool { return v.AllowAnyHostAuth }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) AssumedOfflinePeriodSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *int { return v.AssumedOfflinePeriodSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) BlockMismatchingMimeTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *bool { return v.BlockMismatchingMimeTypes }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) BypassHeadRequests() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *bool { return v.BypassHeadRequests }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) ClientTlsCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) string { return v.ClientTlsCertificate }).(pulumi.StringOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) ContentSynchronisation() GetRemoteOciRepositoryContentSynchronisationOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) GetRemoteOciRepositoryContentSynchronisation {
		return v.ContentSynchronisation
	}).(GetRemoteOciRepositoryContentSynchronisationOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) DisableProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *bool { return v.DisableProxy }).(pulumi.BoolPtrOutput)
}

// (Optional) Whether to disable URL normalization.
func (o LookupRemoteOciRepositoryResultOutput) DisableUrlNormalization() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *bool { return v.DisableUrlNormalization }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) EnableCookieManagement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *bool { return v.EnableCookieManagement }).(pulumi.BoolPtrOutput)
}

// (Optional) Enable token (Bearer) based authentication.
func (o LookupRemoteOciRepositoryResultOutput) EnableTokenAuthentication() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) bool { return v.EnableTokenAuthentication }).(pulumi.BoolOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// (Optional) Also known as 'Foreign Layers Caching' on the UI.
func (o LookupRemoteOciRepositoryResultOutput) ExternalDependenciesEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *bool { return v.ExternalDependenciesEnabled }).(pulumi.BoolPtrOutput)
}

// (Optional) Optional include patterns to match external URLs. Ant-style path expressions are supported (*, **, ?). For example, specifying `**/github.com/**` will only allow downloading foreign layers from github.com host. Due to SDKv2 limitations, we can't set the default value for the list. This value `[**]` must be assigned to the attribute manually, if user don't specify any other non-default values. We don't want to make this attribute required, but it must be set to avoid the state drift on update. Note: Artifactory assigns `[**]` on update if HCL doesn't have the attribute set or the list is empty.
func (o LookupRemoteOciRepositoryResultOutput) ExternalDependenciesPatterns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) []string { return v.ExternalDependenciesPatterns }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) HardFail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *bool { return v.HardFail }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRemoteOciRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) ListRemoteFolderItems() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *bool { return v.ListRemoteFolderItems }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) LocalAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *string { return v.LocalAddress }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) MetadataRetrievalTimeoutSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *int { return v.MetadataRetrievalTimeoutSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) MismatchingMimeTypesOverrideList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *string { return v.MismatchingMimeTypesOverrideList }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) MissedCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *int { return v.MissedCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) Offline() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *bool { return v.Offline }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *string { return v.Proxy }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) QueryParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *string { return v.QueryParams }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) RemoteRepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *string { return v.RemoteRepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) RetrievalCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *int { return v.RetrievalCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) ShareConfiguration() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) bool { return v.ShareConfiguration }).(pulumi.BoolOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) SocketTimeoutMillis() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *int { return v.SocketTimeoutMillis }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) StoreArtifactsLocally() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *bool { return v.StoreArtifactsLocally }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) SynchronizeProperties() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *bool { return v.SynchronizeProperties }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) UnusedArtifactsCleanupPeriodHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *int { return v.UnusedArtifactsCleanupPeriodHours }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteOciRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteOciRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemoteOciRepositoryResultOutput{})
}
