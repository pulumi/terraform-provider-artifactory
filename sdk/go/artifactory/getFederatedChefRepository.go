// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a federated Chef repository.
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
//			_, err := artifactory.LookupFederatedChefRepository(ctx, &artifactory.LookupFederatedChefRepositoryArgs{
//				Key: "federated-test-chef-repo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFederatedChefRepository(ctx *pulumi.Context, args *LookupFederatedChefRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupFederatedChefRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFederatedChefRepositoryResult
	err := ctx.Invoke("artifactory:index/getFederatedChefRepository:getFederatedChefRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedChefRepository.
type LookupFederatedChefRepositoryArgs struct {
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
	Members             []GetFederatedChefRepositoryMember `pulumi:"members"`
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

// A collection of values returned by getFederatedChefRepository.
type LookupFederatedChefRepositoryResult struct {
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
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             []GetFederatedChefRepositoryMember `pulumi:"members"`
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

func LookupFederatedChefRepositoryOutput(ctx *pulumi.Context, args LookupFederatedChefRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedChefRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedChefRepositoryResult, error) {
			args := v.(LookupFederatedChefRepositoryArgs)
			r, err := LookupFederatedChefRepository(ctx, &args, opts...)
			var s LookupFederatedChefRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFederatedChefRepositoryResultOutput)
}

// A collection of arguments for invoking getFederatedChefRepository.
type LookupFederatedChefRepositoryOutputArgs struct {
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
	Members             GetFederatedChefRepositoryMemberArrayInput `pulumi:"members"`
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

func (LookupFederatedChefRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedChefRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedChefRepository.
type LookupFederatedChefRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedChefRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedChefRepositoryResult)(nil)).Elem()
}

func (o LookupFederatedChefRepositoryResultOutput) ToLookupFederatedChefRepositoryResultOutput() LookupFederatedChefRepositoryResultOutput {
	return o
}

func (o LookupFederatedChefRepositoryResultOutput) ToLookupFederatedChefRepositoryResultOutputWithContext(ctx context.Context) LookupFederatedChefRepositoryResultOutput {
	return o
}

func (o LookupFederatedChefRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedChefRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedChefRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedChefRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedChefRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedChefRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedChefRepositoryResultOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedChefRepositoryResult) *bool { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedChefRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedChefRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
func (o LookupFederatedChefRepositoryResultOutput) DisableProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedChefRepositoryResult) *bool { return v.DisableProxy }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedChefRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedChefRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedChefRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedChefRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFederatedChefRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedChefRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedChefRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedChefRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// the identity key of the repo.
func (o LookupFederatedChefRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedChefRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o LookupFederatedChefRepositoryResultOutput) Members() GetFederatedChefRepositoryMemberArrayOutput {
	return o.ApplyT(func(v LookupFederatedChefRepositoryResult) []GetFederatedChefRepositoryMember { return v.Members }).(GetFederatedChefRepositoryMemberArrayOutput)
}

func (o LookupFederatedChefRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedChefRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedChefRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedChefRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupFederatedChefRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedChefRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedChefRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedChefRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedChefRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedChefRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedChefRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedChefRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Proxy key from Artifactory Proxies settings.
func (o LookupFederatedChefRepositoryResultOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedChefRepositoryResult) *string { return v.Proxy }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedChefRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedChefRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedChefRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedChefRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedChefRepositoryResultOutput{})
}
