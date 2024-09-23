// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a federated Opkg repository.
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
//			_, err := artifactory.LookupFederatedOpkgRepository(ctx, &artifactory.LookupFederatedOpkgRepositoryArgs{
//				Key: "federated-test-opkg-repo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFederatedOpkgRepository(ctx *pulumi.Context, args *LookupFederatedOpkgRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupFederatedOpkgRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFederatedOpkgRepositoryResult
	err := ctx.Invoke("artifactory:index/getFederatedOpkgRepository:getFederatedOpkgRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedOpkgRepository.
type LookupFederatedOpkgRepositoryArgs struct {
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
	Members             []GetFederatedOpkgRepositoryMember `pulumi:"members"`
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

// A collection of values returned by getFederatedOpkgRepository.
type LookupFederatedOpkgRepositoryResult struct {
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
	Members             []GetFederatedOpkgRepositoryMember `pulumi:"members"`
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

func LookupFederatedOpkgRepositoryOutput(ctx *pulumi.Context, args LookupFederatedOpkgRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedOpkgRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedOpkgRepositoryResultOutput, error) {
			args := v.(LookupFederatedOpkgRepositoryArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupFederatedOpkgRepositoryResult
			secret, err := ctx.InvokePackageRaw("artifactory:index/getFederatedOpkgRepository:getFederatedOpkgRepository", args, &rv, "", opts...)
			if err != nil {
				return LookupFederatedOpkgRepositoryResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupFederatedOpkgRepositoryResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupFederatedOpkgRepositoryResultOutput), nil
			}
			return output, nil
		}).(LookupFederatedOpkgRepositoryResultOutput)
}

// A collection of arguments for invoking getFederatedOpkgRepository.
type LookupFederatedOpkgRepositoryOutputArgs struct {
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
	Members             GetFederatedOpkgRepositoryMemberArrayInput `pulumi:"members"`
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

func (LookupFederatedOpkgRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedOpkgRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedOpkgRepository.
type LookupFederatedOpkgRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedOpkgRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedOpkgRepositoryResult)(nil)).Elem()
}

func (o LookupFederatedOpkgRepositoryResultOutput) ToLookupFederatedOpkgRepositoryResultOutput() LookupFederatedOpkgRepositoryResultOutput {
	return o
}

func (o LookupFederatedOpkgRepositoryResultOutput) ToLookupFederatedOpkgRepositoryResultOutputWithContext(ctx context.Context) LookupFederatedOpkgRepositoryResultOutput {
	return o
}

func (o LookupFederatedOpkgRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedOpkgRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedOpkgRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedOpkgRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedOpkgRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedOpkgRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedOpkgRepositoryResultOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedOpkgRepositoryResult) *bool { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedOpkgRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedOpkgRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
func (o LookupFederatedOpkgRepositoryResultOutput) DisableProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedOpkgRepositoryResult) *bool { return v.DisableProxy }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedOpkgRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedOpkgRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedOpkgRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedOpkgRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFederatedOpkgRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedOpkgRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedOpkgRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedOpkgRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedOpkgRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedOpkgRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o LookupFederatedOpkgRepositoryResultOutput) Members() GetFederatedOpkgRepositoryMemberArrayOutput {
	return o.ApplyT(func(v LookupFederatedOpkgRepositoryResult) []GetFederatedOpkgRepositoryMember { return v.Members }).(GetFederatedOpkgRepositoryMemberArrayOutput)
}

func (o LookupFederatedOpkgRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedOpkgRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedOpkgRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedOpkgRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupFederatedOpkgRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedOpkgRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedOpkgRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedOpkgRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedOpkgRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedOpkgRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedOpkgRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedOpkgRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Proxy key from Artifactory Proxies settings.
func (o LookupFederatedOpkgRepositoryResultOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedOpkgRepositoryResult) *string { return v.Proxy }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedOpkgRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedOpkgRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedOpkgRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedOpkgRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedOpkgRepositoryResultOutput{})
}
