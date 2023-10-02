// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Retrieves a federated Nuget repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.LookupFederatedNugetRepository(ctx, &artifactory.LookupFederatedNugetRepositoryArgs{
//				Key: "federated-test-nuget-repo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFederatedNugetRepository(ctx *pulumi.Context, args *LookupFederatedNugetRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupFederatedNugetRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFederatedNugetRepositoryResult
	err := ctx.Invoke("artifactory:index/getFederatedNugetRepository:getFederatedNugetRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedNugetRepository.
type LookupFederatedNugetRepositoryArgs struct {
	ArchiveBrowsingEnabled   *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut               *bool   `pulumi:"blackedOut"`
	CdnRedirect              *bool   `pulumi:"cdnRedirect"`
	CleanupOnDelete          *bool   `pulumi:"cleanupOnDelete"`
	Description              *string `pulumi:"description"`
	DownloadDirect           *bool   `pulumi:"downloadDirect"`
	ExcludesPattern          *string `pulumi:"excludesPattern"`
	ForceNugetAuthentication *bool   `pulumi:"forceNugetAuthentication"`
	IncludesPattern          *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                string `pulumi:"key"`
	MaxUniqueSnapshots *int   `pulumi:"maxUniqueSnapshots"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             []GetFederatedNugetRepositoryMember `pulumi:"members"`
	Notes               *string                             `pulumi:"notes"`
	PriorityResolution  *bool                               `pulumi:"priorityResolution"`
	ProjectEnvironments []string                            `pulumi:"projectEnvironments"`
	ProjectKey          *string                             `pulumi:"projectKey"`
	PropertySets        []string                            `pulumi:"propertySets"`
	RepoLayoutRef       *string                             `pulumi:"repoLayoutRef"`
	XrayIndex           *bool                               `pulumi:"xrayIndex"`
}

// A collection of values returned by getFederatedNugetRepository.
type LookupFederatedNugetRepositoryResult struct {
	ArchiveBrowsingEnabled   *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut               *bool   `pulumi:"blackedOut"`
	CdnRedirect              *bool   `pulumi:"cdnRedirect"`
	CleanupOnDelete          *bool   `pulumi:"cleanupOnDelete"`
	Description              *string `pulumi:"description"`
	DownloadDirect           *bool   `pulumi:"downloadDirect"`
	ExcludesPattern          string  `pulumi:"excludesPattern"`
	ForceNugetAuthentication *bool   `pulumi:"forceNugetAuthentication"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string `pulumi:"id"`
	IncludesPattern    string `pulumi:"includesPattern"`
	Key                string `pulumi:"key"`
	MaxUniqueSnapshots *int   `pulumi:"maxUniqueSnapshots"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             []GetFederatedNugetRepositoryMember `pulumi:"members"`
	Notes               *string                             `pulumi:"notes"`
	PackageType         string                              `pulumi:"packageType"`
	PriorityResolution  *bool                               `pulumi:"priorityResolution"`
	ProjectEnvironments []string                            `pulumi:"projectEnvironments"`
	ProjectKey          *string                             `pulumi:"projectKey"`
	PropertySets        []string                            `pulumi:"propertySets"`
	RepoLayoutRef       *string                             `pulumi:"repoLayoutRef"`
	XrayIndex           *bool                               `pulumi:"xrayIndex"`
}

func LookupFederatedNugetRepositoryOutput(ctx *pulumi.Context, args LookupFederatedNugetRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedNugetRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedNugetRepositoryResult, error) {
			args := v.(LookupFederatedNugetRepositoryArgs)
			r, err := LookupFederatedNugetRepository(ctx, &args, opts...)
			var s LookupFederatedNugetRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFederatedNugetRepositoryResultOutput)
}

// A collection of arguments for invoking getFederatedNugetRepository.
type LookupFederatedNugetRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled   pulumi.BoolPtrInput   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut               pulumi.BoolPtrInput   `pulumi:"blackedOut"`
	CdnRedirect              pulumi.BoolPtrInput   `pulumi:"cdnRedirect"`
	CleanupOnDelete          pulumi.BoolPtrInput   `pulumi:"cleanupOnDelete"`
	Description              pulumi.StringPtrInput `pulumi:"description"`
	DownloadDirect           pulumi.BoolPtrInput   `pulumi:"downloadDirect"`
	ExcludesPattern          pulumi.StringPtrInput `pulumi:"excludesPattern"`
	ForceNugetAuthentication pulumi.BoolPtrInput   `pulumi:"forceNugetAuthentication"`
	IncludesPattern          pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                pulumi.StringInput `pulumi:"key"`
	MaxUniqueSnapshots pulumi.IntPtrInput `pulumi:"maxUniqueSnapshots"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             GetFederatedNugetRepositoryMemberArrayInput `pulumi:"members"`
	Notes               pulumi.StringPtrInput                       `pulumi:"notes"`
	PriorityResolution  pulumi.BoolPtrInput                         `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayInput                     `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput                       `pulumi:"projectKey"`
	PropertySets        pulumi.StringArrayInput                     `pulumi:"propertySets"`
	RepoLayoutRef       pulumi.StringPtrInput                       `pulumi:"repoLayoutRef"`
	XrayIndex           pulumi.BoolPtrInput                         `pulumi:"xrayIndex"`
}

func (LookupFederatedNugetRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedNugetRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedNugetRepository.
type LookupFederatedNugetRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedNugetRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedNugetRepositoryResult)(nil)).Elem()
}

func (o LookupFederatedNugetRepositoryResultOutput) ToLookupFederatedNugetRepositoryResultOutput() LookupFederatedNugetRepositoryResultOutput {
	return o
}

func (o LookupFederatedNugetRepositoryResultOutput) ToLookupFederatedNugetRepositoryResultOutputWithContext(ctx context.Context) LookupFederatedNugetRepositoryResultOutput {
	return o
}

func (o LookupFederatedNugetRepositoryResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupFederatedNugetRepositoryResult] {
	return pulumix.Output[LookupFederatedNugetRepositoryResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupFederatedNugetRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedNugetRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedNugetRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedNugetRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedNugetRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedNugetRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedNugetRepositoryResultOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedNugetRepositoryResult) *bool { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedNugetRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedNugetRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedNugetRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedNugetRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedNugetRepositoryResultOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedNugetRepositoryResult) string { return v.ExcludesPattern }).(pulumi.StringOutput)
}

func (o LookupFederatedNugetRepositoryResultOutput) ForceNugetAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedNugetRepositoryResult) *bool { return v.ForceNugetAuthentication }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFederatedNugetRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedNugetRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedNugetRepositoryResultOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedNugetRepositoryResult) string { return v.IncludesPattern }).(pulumi.StringOutput)
}

func (o LookupFederatedNugetRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedNugetRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupFederatedNugetRepositoryResultOutput) MaxUniqueSnapshots() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFederatedNugetRepositoryResult) *int { return v.MaxUniqueSnapshots }).(pulumi.IntPtrOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o LookupFederatedNugetRepositoryResultOutput) Members() GetFederatedNugetRepositoryMemberArrayOutput {
	return o.ApplyT(func(v LookupFederatedNugetRepositoryResult) []GetFederatedNugetRepositoryMember { return v.Members }).(GetFederatedNugetRepositoryMemberArrayOutput)
}

func (o LookupFederatedNugetRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedNugetRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedNugetRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedNugetRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupFederatedNugetRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedNugetRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedNugetRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedNugetRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedNugetRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedNugetRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedNugetRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedNugetRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedNugetRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedNugetRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedNugetRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedNugetRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedNugetRepositoryResultOutput{})
}
