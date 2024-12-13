// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a federated Cargo repository.
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
//			_, err := artifactory.LookupFederatedCargoRepository(ctx, &artifactory.LookupFederatedCargoRepositoryArgs{
//				Key: "federated-test-cargo-repo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFederatedCargoRepository(ctx *pulumi.Context, args *LookupFederatedCargoRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupFederatedCargoRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFederatedCargoRepositoryResult
	err := ctx.Invoke("artifactory:index/getFederatedCargoRepository:getFederatedCargoRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedCargoRepository.
type LookupFederatedCargoRepositoryArgs struct {
	AnonymousAccess        *bool   `pulumi:"anonymousAccess"`
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
	DisableProxy            *bool    `pulumi:"disableProxy"`
	DownloadDirect          *bool    `pulumi:"downloadDirect"`
	EnableSparseIndex       *bool    `pulumi:"enableSparseIndex"`
	ExcludesPattern         *string  `pulumi:"excludesPattern"`
	IncludesPattern         *string  `pulumi:"includesPattern"`
	IndexCompressionFormats []string `pulumi:"indexCompressionFormats"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             []GetFederatedCargoRepositoryMember `pulumi:"members"`
	Notes               *string                             `pulumi:"notes"`
	PriorityResolution  *bool                               `pulumi:"priorityResolution"`
	ProjectEnvironments []string                            `pulumi:"projectEnvironments"`
	ProjectKey          *string                             `pulumi:"projectKey"`
	PropertySets        []string                            `pulumi:"propertySets"`
	// Proxy key from Artifactory Proxies settings.
	Proxy         *string `pulumi:"proxy"`
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	XrayIndex     *bool   `pulumi:"xrayIndex"`
}

// A collection of values returned by getFederatedCargoRepository.
type LookupFederatedCargoRepositoryResult struct {
	AnonymousAccess        *bool   `pulumi:"anonymousAccess"`
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
	DisableProxy      *bool   `pulumi:"disableProxy"`
	DownloadDirect    *bool   `pulumi:"downloadDirect"`
	EnableSparseIndex *bool   `pulumi:"enableSparseIndex"`
	ExcludesPattern   *string `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id                      string   `pulumi:"id"`
	IncludesPattern         *string  `pulumi:"includesPattern"`
	IndexCompressionFormats []string `pulumi:"indexCompressionFormats"`
	Key                     string   `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             []GetFederatedCargoRepositoryMember `pulumi:"members"`
	Notes               *string                             `pulumi:"notes"`
	PackageType         string                              `pulumi:"packageType"`
	PriorityResolution  *bool                               `pulumi:"priorityResolution"`
	ProjectEnvironments []string                            `pulumi:"projectEnvironments"`
	ProjectKey          *string                             `pulumi:"projectKey"`
	PropertySets        []string                            `pulumi:"propertySets"`
	// Proxy key from Artifactory Proxies settings.
	Proxy         *string `pulumi:"proxy"`
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	XrayIndex     *bool   `pulumi:"xrayIndex"`
}

func LookupFederatedCargoRepositoryOutput(ctx *pulumi.Context, args LookupFederatedCargoRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedCargoRepositoryResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupFederatedCargoRepositoryResultOutput, error) {
			args := v.(LookupFederatedCargoRepositoryArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("artifactory:index/getFederatedCargoRepository:getFederatedCargoRepository", args, LookupFederatedCargoRepositoryResultOutput{}, options).(LookupFederatedCargoRepositoryResultOutput), nil
		}).(LookupFederatedCargoRepositoryResultOutput)
}

// A collection of arguments for invoking getFederatedCargoRepository.
type LookupFederatedCargoRepositoryOutputArgs struct {
	AnonymousAccess        pulumi.BoolPtrInput   `pulumi:"anonymousAccess"`
	ArchiveBrowsingEnabled pulumi.BoolPtrInput   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput   `pulumi:"blackedOut"`
	CdnRedirect            pulumi.BoolPtrInput   `pulumi:"cdnRedirect"`
	CleanupOnDelete        pulumi.BoolPtrInput   `pulumi:"cleanupOnDelete"`
	Description            pulumi.StringPtrInput `pulumi:"description"`
	// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
	DisableProxy            pulumi.BoolPtrInput     `pulumi:"disableProxy"`
	DownloadDirect          pulumi.BoolPtrInput     `pulumi:"downloadDirect"`
	EnableSparseIndex       pulumi.BoolPtrInput     `pulumi:"enableSparseIndex"`
	ExcludesPattern         pulumi.StringPtrInput   `pulumi:"excludesPattern"`
	IncludesPattern         pulumi.StringPtrInput   `pulumi:"includesPattern"`
	IndexCompressionFormats pulumi.StringArrayInput `pulumi:"indexCompressionFormats"`
	// the identity key of the repo.
	Key pulumi.StringInput `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             GetFederatedCargoRepositoryMemberArrayInput `pulumi:"members"`
	Notes               pulumi.StringPtrInput                       `pulumi:"notes"`
	PriorityResolution  pulumi.BoolPtrInput                         `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayInput                     `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput                       `pulumi:"projectKey"`
	PropertySets        pulumi.StringArrayInput                     `pulumi:"propertySets"`
	// Proxy key from Artifactory Proxies settings.
	Proxy         pulumi.StringPtrInput `pulumi:"proxy"`
	RepoLayoutRef pulumi.StringPtrInput `pulumi:"repoLayoutRef"`
	XrayIndex     pulumi.BoolPtrInput   `pulumi:"xrayIndex"`
}

func (LookupFederatedCargoRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedCargoRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedCargoRepository.
type LookupFederatedCargoRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedCargoRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedCargoRepositoryResult)(nil)).Elem()
}

func (o LookupFederatedCargoRepositoryResultOutput) ToLookupFederatedCargoRepositoryResultOutput() LookupFederatedCargoRepositoryResultOutput {
	return o
}

func (o LookupFederatedCargoRepositoryResultOutput) ToLookupFederatedCargoRepositoryResultOutputWithContext(ctx context.Context) LookupFederatedCargoRepositoryResultOutput {
	return o
}

func (o LookupFederatedCargoRepositoryResultOutput) AnonymousAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCargoRepositoryResult) *bool { return v.AnonymousAccess }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCargoRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCargoRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCargoRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCargoRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCargoRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCargoRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCargoRepositoryResultOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCargoRepositoryResult) *bool { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCargoRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedCargoRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
func (o LookupFederatedCargoRepositoryResultOutput) DisableProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCargoRepositoryResult) *bool { return v.DisableProxy }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCargoRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCargoRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCargoRepositoryResultOutput) EnableSparseIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCargoRepositoryResult) *bool { return v.EnableSparseIndex }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCargoRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedCargoRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFederatedCargoRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedCargoRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedCargoRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedCargoRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedCargoRepositoryResultOutput) IndexCompressionFormats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedCargoRepositoryResult) []string { return v.IndexCompressionFormats }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedCargoRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedCargoRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o LookupFederatedCargoRepositoryResultOutput) Members() GetFederatedCargoRepositoryMemberArrayOutput {
	return o.ApplyT(func(v LookupFederatedCargoRepositoryResult) []GetFederatedCargoRepositoryMember { return v.Members }).(GetFederatedCargoRepositoryMemberArrayOutput)
}

func (o LookupFederatedCargoRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedCargoRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedCargoRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedCargoRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupFederatedCargoRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCargoRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCargoRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedCargoRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedCargoRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedCargoRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedCargoRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedCargoRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Proxy key from Artifactory Proxies settings.
func (o LookupFederatedCargoRepositoryResultOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedCargoRepositoryResult) *string { return v.Proxy }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedCargoRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedCargoRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedCargoRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCargoRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedCargoRepositoryResultOutput{})
}
