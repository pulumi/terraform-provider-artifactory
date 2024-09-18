// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a federated Cran repository.
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
//			_, err := artifactory.LookupFederatedCranRepository(ctx, &artifactory.LookupFederatedCranRepositoryArgs{
//				Key: "federated-test-cran-repo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFederatedCranRepository(ctx *pulumi.Context, args *LookupFederatedCranRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupFederatedCranRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFederatedCranRepositoryResult
	err := ctx.Invoke("artifactory:index/getFederatedCranRepository:getFederatedCranRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedCranRepository.
type LookupFederatedCranRepositoryArgs struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
	DisableProxy    *bool   `pulumi:"disableProxy"`
	DownloadDirect  *bool   `pulumi:"downloadDirect"`
	ExcludesPattern *string `pulumi:"excludesPattern"`
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             []GetFederatedCranRepositoryMember `pulumi:"members"`
	Notes               *string                            `pulumi:"notes"`
	PriorityResolution  *bool                              `pulumi:"priorityResolution"`
	ProjectEnvironments []string                           `pulumi:"projectEnvironments"`
	ProjectKey          *string                            `pulumi:"projectKey"`
	PropertySets        []string                           `pulumi:"propertySets"`
	// Proxy key from Artifactory Proxies settings.
	Proxy         *string `pulumi:"proxy"`
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	XrayIndex     *bool   `pulumi:"xrayIndex"`
}

// A collection of values returned by getFederatedCranRepository.
type LookupFederatedCranRepositoryResult struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
	DisableProxy    *bool   `pulumi:"disableProxy"`
	DownloadDirect  *bool   `pulumi:"downloadDirect"`
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id              string  `pulumi:"id"`
	IncludesPattern *string `pulumi:"includesPattern"`
	Key             string  `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             []GetFederatedCranRepositoryMember `pulumi:"members"`
	Notes               *string                            `pulumi:"notes"`
	PackageType         string                             `pulumi:"packageType"`
	PriorityResolution  *bool                              `pulumi:"priorityResolution"`
	ProjectEnvironments []string                           `pulumi:"projectEnvironments"`
	ProjectKey          *string                            `pulumi:"projectKey"`
	PropertySets        []string                           `pulumi:"propertySets"`
	// Proxy key from Artifactory Proxies settings.
	Proxy         *string `pulumi:"proxy"`
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	XrayIndex     *bool   `pulumi:"xrayIndex"`
}

func LookupFederatedCranRepositoryOutput(ctx *pulumi.Context, args LookupFederatedCranRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedCranRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedCranRepositoryResult, error) {
			args := v.(LookupFederatedCranRepositoryArgs)
			r, err := LookupFederatedCranRepository(ctx, &args, opts...)
			var s LookupFederatedCranRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFederatedCranRepositoryResultOutput)
}

// A collection of arguments for invoking getFederatedCranRepository.
type LookupFederatedCranRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput   `pulumi:"blackedOut"`
	CdnRedirect            pulumi.BoolPtrInput   `pulumi:"cdnRedirect"`
	CleanupOnDelete        pulumi.BoolPtrInput   `pulumi:"cleanupOnDelete"`
	Description            pulumi.StringPtrInput `pulumi:"description"`
	// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
	DisableProxy    pulumi.BoolPtrInput   `pulumi:"disableProxy"`
	DownloadDirect  pulumi.BoolPtrInput   `pulumi:"downloadDirect"`
	ExcludesPattern pulumi.StringPtrInput `pulumi:"excludesPattern"`
	IncludesPattern pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key pulumi.StringInput `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             GetFederatedCranRepositoryMemberArrayInput `pulumi:"members"`
	Notes               pulumi.StringPtrInput                      `pulumi:"notes"`
	PriorityResolution  pulumi.BoolPtrInput                        `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayInput                    `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput                      `pulumi:"projectKey"`
	PropertySets        pulumi.StringArrayInput                    `pulumi:"propertySets"`
	// Proxy key from Artifactory Proxies settings.
	Proxy         pulumi.StringPtrInput `pulumi:"proxy"`
	RepoLayoutRef pulumi.StringPtrInput `pulumi:"repoLayoutRef"`
	XrayIndex     pulumi.BoolPtrInput   `pulumi:"xrayIndex"`
}

func (LookupFederatedCranRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedCranRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedCranRepository.
type LookupFederatedCranRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedCranRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedCranRepositoryResult)(nil)).Elem()
}

func (o LookupFederatedCranRepositoryResultOutput) ToLookupFederatedCranRepositoryResultOutput() LookupFederatedCranRepositoryResultOutput {
	return o
}

func (o LookupFederatedCranRepositoryResultOutput) ToLookupFederatedCranRepositoryResultOutputWithContext(ctx context.Context) LookupFederatedCranRepositoryResultOutput {
	return o
}

func (o LookupFederatedCranRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCranRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCranRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCranRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCranRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCranRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCranRepositoryResultOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCranRepositoryResult) *bool { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCranRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedCranRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
func (o LookupFederatedCranRepositoryResultOutput) DisableProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCranRepositoryResult) *bool { return v.DisableProxy }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCranRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCranRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCranRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedCranRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFederatedCranRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedCranRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedCranRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedCranRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedCranRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedCranRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o LookupFederatedCranRepositoryResultOutput) Members() GetFederatedCranRepositoryMemberArrayOutput {
	return o.ApplyT(func(v LookupFederatedCranRepositoryResult) []GetFederatedCranRepositoryMember { return v.Members }).(GetFederatedCranRepositoryMemberArrayOutput)
}

func (o LookupFederatedCranRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedCranRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedCranRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedCranRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupFederatedCranRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCranRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCranRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedCranRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedCranRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedCranRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedCranRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedCranRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Proxy key from Artifactory Proxies settings.
func (o LookupFederatedCranRepositoryResultOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedCranRepositoryResult) *string { return v.Proxy }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedCranRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedCranRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedCranRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCranRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedCranRepositoryResultOutput{})
}
