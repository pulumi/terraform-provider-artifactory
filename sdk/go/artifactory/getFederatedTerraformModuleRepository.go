// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v4/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
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
//			_, err := artifactory.LookupFederatedTerraformModuleRepository(ctx, &artifactory.LookupFederatedTerraformModuleRepositoryArgs{
//				Key: "federated-test-terraform-module-repo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFederatedTerraformModuleRepository(ctx *pulumi.Context, args *LookupFederatedTerraformModuleRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupFederatedTerraformModuleRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFederatedTerraformModuleRepositoryResult
	err := ctx.Invoke("artifactory:index/getFederatedTerraformModuleRepository:getFederatedTerraformModuleRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedTerraformModuleRepository.
type LookupFederatedTerraformModuleRepositoryArgs struct {
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
	Members             []GetFederatedTerraformModuleRepositoryMember `pulumi:"members"`
	Notes               *string                                       `pulumi:"notes"`
	PriorityResolution  *bool                                         `pulumi:"priorityResolution"`
	ProjectEnvironments []string                                      `pulumi:"projectEnvironments"`
	ProjectKey          *string                                       `pulumi:"projectKey"`
	PropertySets        []string                                      `pulumi:"propertySets"`
	RepoLayoutRef       *string                                       `pulumi:"repoLayoutRef"`
	XrayIndex           *bool                                         `pulumi:"xrayIndex"`
}

// A collection of values returned by getFederatedTerraformModuleRepository.
type LookupFederatedTerraformModuleRepositoryResult struct {
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
	Members             []GetFederatedTerraformModuleRepositoryMember `pulumi:"members"`
	Notes               *string                                       `pulumi:"notes"`
	PackageType         string                                        `pulumi:"packageType"`
	PriorityResolution  *bool                                         `pulumi:"priorityResolution"`
	ProjectEnvironments []string                                      `pulumi:"projectEnvironments"`
	ProjectKey          *string                                       `pulumi:"projectKey"`
	PropertySets        []string                                      `pulumi:"propertySets"`
	RepoLayoutRef       *string                                       `pulumi:"repoLayoutRef"`
	XrayIndex           *bool                                         `pulumi:"xrayIndex"`
}

func LookupFederatedTerraformModuleRepositoryOutput(ctx *pulumi.Context, args LookupFederatedTerraformModuleRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedTerraformModuleRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedTerraformModuleRepositoryResult, error) {
			args := v.(LookupFederatedTerraformModuleRepositoryArgs)
			r, err := LookupFederatedTerraformModuleRepository(ctx, &args, opts...)
			var s LookupFederatedTerraformModuleRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFederatedTerraformModuleRepositoryResultOutput)
}

// A collection of arguments for invoking getFederatedTerraformModuleRepository.
type LookupFederatedTerraformModuleRepositoryOutputArgs struct {
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
	Members             GetFederatedTerraformModuleRepositoryMemberArrayInput `pulumi:"members"`
	Notes               pulumi.StringPtrInput                                 `pulumi:"notes"`
	PriorityResolution  pulumi.BoolPtrInput                                   `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayInput                               `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput                                 `pulumi:"projectKey"`
	PropertySets        pulumi.StringArrayInput                               `pulumi:"propertySets"`
	RepoLayoutRef       pulumi.StringPtrInput                                 `pulumi:"repoLayoutRef"`
	XrayIndex           pulumi.BoolPtrInput                                   `pulumi:"xrayIndex"`
}

func (LookupFederatedTerraformModuleRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedTerraformModuleRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedTerraformModuleRepository.
type LookupFederatedTerraformModuleRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedTerraformModuleRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedTerraformModuleRepositoryResult)(nil)).Elem()
}

func (o LookupFederatedTerraformModuleRepositoryResultOutput) ToLookupFederatedTerraformModuleRepositoryResultOutput() LookupFederatedTerraformModuleRepositoryResultOutput {
	return o
}

func (o LookupFederatedTerraformModuleRepositoryResultOutput) ToLookupFederatedTerraformModuleRepositoryResultOutputWithContext(ctx context.Context) LookupFederatedTerraformModuleRepositoryResultOutput {
	return o
}

func (o LookupFederatedTerraformModuleRepositoryResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupFederatedTerraformModuleRepositoryResult] {
	return pulumix.Output[LookupFederatedTerraformModuleRepositoryResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupFederatedTerraformModuleRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedTerraformModuleRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedTerraformModuleRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedTerraformModuleRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedTerraformModuleRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedTerraformModuleRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedTerraformModuleRepositoryResultOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedTerraformModuleRepositoryResult) *bool { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedTerraformModuleRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedTerraformModuleRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedTerraformModuleRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedTerraformModuleRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedTerraformModuleRepositoryResultOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedTerraformModuleRepositoryResult) string { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFederatedTerraformModuleRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedTerraformModuleRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedTerraformModuleRepositoryResultOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedTerraformModuleRepositoryResult) string { return v.IncludesPattern }).(pulumi.StringOutput)
}

func (o LookupFederatedTerraformModuleRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedTerraformModuleRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o LookupFederatedTerraformModuleRepositoryResultOutput) Members() GetFederatedTerraformModuleRepositoryMemberArrayOutput {
	return o.ApplyT(func(v LookupFederatedTerraformModuleRepositoryResult) []GetFederatedTerraformModuleRepositoryMember {
		return v.Members
	}).(GetFederatedTerraformModuleRepositoryMemberArrayOutput)
}

func (o LookupFederatedTerraformModuleRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedTerraformModuleRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedTerraformModuleRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedTerraformModuleRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupFederatedTerraformModuleRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedTerraformModuleRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedTerraformModuleRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedTerraformModuleRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedTerraformModuleRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedTerraformModuleRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedTerraformModuleRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedTerraformModuleRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedTerraformModuleRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedTerraformModuleRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedTerraformModuleRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedTerraformModuleRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedTerraformModuleRepositoryResultOutput{})
}
