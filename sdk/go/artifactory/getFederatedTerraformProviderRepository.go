// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v4/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v4/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.LookupFederatedTerraformProviderRepository(ctx, &artifactory.LookupFederatedTerraformProviderRepositoryArgs{
//				Key: "federated-test-terraform-provider-repo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFederatedTerraformProviderRepository(ctx *pulumi.Context, args *LookupFederatedTerraformProviderRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupFederatedTerraformProviderRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFederatedTerraformProviderRepositoryResult
	err := ctx.Invoke("artifactory:index/getFederatedTerraformProviderRepository:getFederatedTerraformProviderRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedTerraformProviderRepository.
type LookupFederatedTerraformProviderRepositoryArgs struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	IncludesPattern        *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             []GetFederatedTerraformProviderRepositoryMember `pulumi:"members"`
	Notes               *string                                         `pulumi:"notes"`
	PriorityResolution  *bool                                           `pulumi:"priorityResolution"`
	ProjectEnvironments []string                                        `pulumi:"projectEnvironments"`
	ProjectKey          *string                                         `pulumi:"projectKey"`
	PropertySets        []string                                        `pulumi:"propertySets"`
	RepoLayoutRef       *string                                         `pulumi:"repoLayoutRef"`
	XrayIndex           *bool                                           `pulumi:"xrayIndex"`
}

// A collection of values returned by getFederatedTerraformProviderRepository.
type LookupFederatedTerraformProviderRepositoryResult struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        string  `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id              string `pulumi:"id"`
	IncludesPattern string `pulumi:"includesPattern"`
	Key             string `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             []GetFederatedTerraformProviderRepositoryMember `pulumi:"members"`
	Notes               *string                                         `pulumi:"notes"`
	PackageType         string                                          `pulumi:"packageType"`
	PriorityResolution  *bool                                           `pulumi:"priorityResolution"`
	ProjectEnvironments []string                                        `pulumi:"projectEnvironments"`
	ProjectKey          *string                                         `pulumi:"projectKey"`
	PropertySets        []string                                        `pulumi:"propertySets"`
	RepoLayoutRef       *string                                         `pulumi:"repoLayoutRef"`
	XrayIndex           *bool                                           `pulumi:"xrayIndex"`
}

func LookupFederatedTerraformProviderRepositoryOutput(ctx *pulumi.Context, args LookupFederatedTerraformProviderRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedTerraformProviderRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedTerraformProviderRepositoryResult, error) {
			args := v.(LookupFederatedTerraformProviderRepositoryArgs)
			r, err := LookupFederatedTerraformProviderRepository(ctx, &args, opts...)
			var s LookupFederatedTerraformProviderRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFederatedTerraformProviderRepositoryResultOutput)
}

// A collection of arguments for invoking getFederatedTerraformProviderRepository.
type LookupFederatedTerraformProviderRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput   `pulumi:"blackedOut"`
	CdnRedirect            pulumi.BoolPtrInput   `pulumi:"cdnRedirect"`
	CleanupOnDelete        pulumi.BoolPtrInput   `pulumi:"cleanupOnDelete"`
	Description            pulumi.StringPtrInput `pulumi:"description"`
	DownloadDirect         pulumi.BoolPtrInput   `pulumi:"downloadDirect"`
	ExcludesPattern        pulumi.StringPtrInput `pulumi:"excludesPattern"`
	IncludesPattern        pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key pulumi.StringInput `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             GetFederatedTerraformProviderRepositoryMemberArrayInput `pulumi:"members"`
	Notes               pulumi.StringPtrInput                                   `pulumi:"notes"`
	PriorityResolution  pulumi.BoolPtrInput                                     `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayInput                                 `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput                                   `pulumi:"projectKey"`
	PropertySets        pulumi.StringArrayInput                                 `pulumi:"propertySets"`
	RepoLayoutRef       pulumi.StringPtrInput                                   `pulumi:"repoLayoutRef"`
	XrayIndex           pulumi.BoolPtrInput                                     `pulumi:"xrayIndex"`
}

func (LookupFederatedTerraformProviderRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedTerraformProviderRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedTerraformProviderRepository.
type LookupFederatedTerraformProviderRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedTerraformProviderRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedTerraformProviderRepositoryResult)(nil)).Elem()
}

func (o LookupFederatedTerraformProviderRepositoryResultOutput) ToLookupFederatedTerraformProviderRepositoryResultOutput() LookupFederatedTerraformProviderRepositoryResultOutput {
	return o
}

func (o LookupFederatedTerraformProviderRepositoryResultOutput) ToLookupFederatedTerraformProviderRepositoryResultOutputWithContext(ctx context.Context) LookupFederatedTerraformProviderRepositoryResultOutput {
	return o
}

func (o LookupFederatedTerraformProviderRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedTerraformProviderRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedTerraformProviderRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedTerraformProviderRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedTerraformProviderRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedTerraformProviderRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedTerraformProviderRepositoryResultOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedTerraformProviderRepositoryResult) *bool { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedTerraformProviderRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedTerraformProviderRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedTerraformProviderRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedTerraformProviderRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedTerraformProviderRepositoryResultOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedTerraformProviderRepositoryResult) string { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFederatedTerraformProviderRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedTerraformProviderRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedTerraformProviderRepositoryResultOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedTerraformProviderRepositoryResult) string { return v.IncludesPattern }).(pulumi.StringOutput)
}

func (o LookupFederatedTerraformProviderRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedTerraformProviderRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o LookupFederatedTerraformProviderRepositoryResultOutput) Members() GetFederatedTerraformProviderRepositoryMemberArrayOutput {
	return o.ApplyT(func(v LookupFederatedTerraformProviderRepositoryResult) []GetFederatedTerraformProviderRepositoryMember {
		return v.Members
	}).(GetFederatedTerraformProviderRepositoryMemberArrayOutput)
}

func (o LookupFederatedTerraformProviderRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedTerraformProviderRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedTerraformProviderRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedTerraformProviderRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupFederatedTerraformProviderRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedTerraformProviderRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedTerraformProviderRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedTerraformProviderRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedTerraformProviderRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedTerraformProviderRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedTerraformProviderRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedTerraformProviderRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedTerraformProviderRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedTerraformProviderRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedTerraformProviderRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedTerraformProviderRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedTerraformProviderRepositoryResultOutput{})
}
