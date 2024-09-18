// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a federated Rpm repository.
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
//			_, err := artifactory.LookupFederatedRpmRepository(ctx, &artifactory.LookupFederatedRpmRepositoryArgs{
//				Key: "federated-test-rpm-repo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFederatedRpmRepository(ctx *pulumi.Context, args *LookupFederatedRpmRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupFederatedRpmRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFederatedRpmRepositoryResult
	err := ctx.Invoke("artifactory:index/getFederatedRpmRepository:getFederatedRpmRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedRpmRepository.
type LookupFederatedRpmRepositoryArgs struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CalculateYumMetadata   *bool   `pulumi:"calculateYumMetadata"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
	DisableProxy            *bool   `pulumi:"disableProxy"`
	DownloadDirect          *bool   `pulumi:"downloadDirect"`
	EnableFileListsIndexing *bool   `pulumi:"enableFileListsIndexing"`
	ExcludesPattern         *string `pulumi:"excludesPattern"`
	IncludesPattern         *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             []GetFederatedRpmRepositoryMember `pulumi:"members"`
	Notes               *string                           `pulumi:"notes"`
	PrimaryKeypairRef   *string                           `pulumi:"primaryKeypairRef"`
	PriorityResolution  *bool                             `pulumi:"priorityResolution"`
	ProjectEnvironments []string                          `pulumi:"projectEnvironments"`
	ProjectKey          *string                           `pulumi:"projectKey"`
	PropertySets        []string                          `pulumi:"propertySets"`
	// Proxy key from Artifactory Proxies settings.
	Proxy               *string `pulumi:"proxy"`
	RepoLayoutRef       *string `pulumi:"repoLayoutRef"`
	SecondaryKeypairRef *string `pulumi:"secondaryKeypairRef"`
	XrayIndex           *bool   `pulumi:"xrayIndex"`
	YumGroupFileNames   *string `pulumi:"yumGroupFileNames"`
	YumRootDepth        *int    `pulumi:"yumRootDepth"`
}

// A collection of values returned by getFederatedRpmRepository.
type LookupFederatedRpmRepositoryResult struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CalculateYumMetadata   *bool   `pulumi:"calculateYumMetadata"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
	DisableProxy            *bool   `pulumi:"disableProxy"`
	DownloadDirect          *bool   `pulumi:"downloadDirect"`
	EnableFileListsIndexing *bool   `pulumi:"enableFileListsIndexing"`
	ExcludesPattern         *string `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id              string  `pulumi:"id"`
	IncludesPattern *string `pulumi:"includesPattern"`
	Key             string  `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             []GetFederatedRpmRepositoryMember `pulumi:"members"`
	Notes               *string                           `pulumi:"notes"`
	PackageType         string                            `pulumi:"packageType"`
	PrimaryKeypairRef   *string                           `pulumi:"primaryKeypairRef"`
	PriorityResolution  *bool                             `pulumi:"priorityResolution"`
	ProjectEnvironments []string                          `pulumi:"projectEnvironments"`
	ProjectKey          *string                           `pulumi:"projectKey"`
	PropertySets        []string                          `pulumi:"propertySets"`
	// Proxy key from Artifactory Proxies settings.
	Proxy               *string `pulumi:"proxy"`
	RepoLayoutRef       *string `pulumi:"repoLayoutRef"`
	SecondaryKeypairRef *string `pulumi:"secondaryKeypairRef"`
	XrayIndex           *bool   `pulumi:"xrayIndex"`
	YumGroupFileNames   *string `pulumi:"yumGroupFileNames"`
	YumRootDepth        *int    `pulumi:"yumRootDepth"`
}

func LookupFederatedRpmRepositoryOutput(ctx *pulumi.Context, args LookupFederatedRpmRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedRpmRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedRpmRepositoryResult, error) {
			args := v.(LookupFederatedRpmRepositoryArgs)
			r, err := LookupFederatedRpmRepository(ctx, &args, opts...)
			var s LookupFederatedRpmRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFederatedRpmRepositoryResultOutput)
}

// A collection of arguments for invoking getFederatedRpmRepository.
type LookupFederatedRpmRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput   `pulumi:"blackedOut"`
	CalculateYumMetadata   pulumi.BoolPtrInput   `pulumi:"calculateYumMetadata"`
	CdnRedirect            pulumi.BoolPtrInput   `pulumi:"cdnRedirect"`
	CleanupOnDelete        pulumi.BoolPtrInput   `pulumi:"cleanupOnDelete"`
	Description            pulumi.StringPtrInput `pulumi:"description"`
	// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
	DisableProxy            pulumi.BoolPtrInput   `pulumi:"disableProxy"`
	DownloadDirect          pulumi.BoolPtrInput   `pulumi:"downloadDirect"`
	EnableFileListsIndexing pulumi.BoolPtrInput   `pulumi:"enableFileListsIndexing"`
	ExcludesPattern         pulumi.StringPtrInput `pulumi:"excludesPattern"`
	IncludesPattern         pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key pulumi.StringInput `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             GetFederatedRpmRepositoryMemberArrayInput `pulumi:"members"`
	Notes               pulumi.StringPtrInput                     `pulumi:"notes"`
	PrimaryKeypairRef   pulumi.StringPtrInput                     `pulumi:"primaryKeypairRef"`
	PriorityResolution  pulumi.BoolPtrInput                       `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayInput                   `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput                     `pulumi:"projectKey"`
	PropertySets        pulumi.StringArrayInput                   `pulumi:"propertySets"`
	// Proxy key from Artifactory Proxies settings.
	Proxy               pulumi.StringPtrInput `pulumi:"proxy"`
	RepoLayoutRef       pulumi.StringPtrInput `pulumi:"repoLayoutRef"`
	SecondaryKeypairRef pulumi.StringPtrInput `pulumi:"secondaryKeypairRef"`
	XrayIndex           pulumi.BoolPtrInput   `pulumi:"xrayIndex"`
	YumGroupFileNames   pulumi.StringPtrInput `pulumi:"yumGroupFileNames"`
	YumRootDepth        pulumi.IntPtrInput    `pulumi:"yumRootDepth"`
}

func (LookupFederatedRpmRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedRpmRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedRpmRepository.
type LookupFederatedRpmRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedRpmRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedRpmRepositoryResult)(nil)).Elem()
}

func (o LookupFederatedRpmRepositoryResultOutput) ToLookupFederatedRpmRepositoryResultOutput() LookupFederatedRpmRepositoryResultOutput {
	return o
}

func (o LookupFederatedRpmRepositoryResultOutput) ToLookupFederatedRpmRepositoryResultOutputWithContext(ctx context.Context) LookupFederatedRpmRepositoryResultOutput {
	return o
}

func (o LookupFederatedRpmRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedRpmRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedRpmRepositoryResultOutput) CalculateYumMetadata() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) *bool { return v.CalculateYumMetadata }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedRpmRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedRpmRepositoryResultOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) *bool { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedRpmRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
func (o LookupFederatedRpmRepositoryResultOutput) DisableProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) *bool { return v.DisableProxy }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedRpmRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedRpmRepositoryResultOutput) EnableFileListsIndexing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) *bool { return v.EnableFileListsIndexing }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedRpmRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFederatedRpmRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedRpmRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedRpmRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o LookupFederatedRpmRepositoryResultOutput) Members() GetFederatedRpmRepositoryMemberArrayOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) []GetFederatedRpmRepositoryMember { return v.Members }).(GetFederatedRpmRepositoryMemberArrayOutput)
}

func (o LookupFederatedRpmRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedRpmRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupFederatedRpmRepositoryResultOutput) PrimaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) *string { return v.PrimaryKeypairRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedRpmRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedRpmRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedRpmRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedRpmRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Proxy key from Artifactory Proxies settings.
func (o LookupFederatedRpmRepositoryResultOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) *string { return v.Proxy }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedRpmRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedRpmRepositoryResultOutput) SecondaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) *string { return v.SecondaryKeypairRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedRpmRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedRpmRepositoryResultOutput) YumGroupFileNames() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) *string { return v.YumGroupFileNames }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedRpmRepositoryResultOutput) YumRootDepth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFederatedRpmRepositoryResult) *int { return v.YumRootDepth }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedRpmRepositoryResultOutput{})
}
