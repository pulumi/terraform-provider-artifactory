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

// Retrieves a federated Npm repository.
//
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
//			_, err := artifactory.LookupFederatedNpmRepository(ctx, &artifactory.LookupFederatedNpmRepositoryArgs{
//				Key: "federated-test-npm-repo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFederatedNpmRepository(ctx *pulumi.Context, args *LookupFederatedNpmRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupFederatedNpmRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFederatedNpmRepositoryResult
	err := ctx.Invoke("artifactory:index/getFederatedNpmRepository:getFederatedNpmRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedNpmRepository.
type LookupFederatedNpmRepositoryArgs struct {
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
	Members             []GetFederatedNpmRepositoryMember `pulumi:"members"`
	Notes               *string                           `pulumi:"notes"`
	PriorityResolution  *bool                             `pulumi:"priorityResolution"`
	ProjectEnvironments []string                          `pulumi:"projectEnvironments"`
	ProjectKey          *string                           `pulumi:"projectKey"`
	PropertySets        []string                          `pulumi:"propertySets"`
	RepoLayoutRef       *string                           `pulumi:"repoLayoutRef"`
	XrayIndex           *bool                             `pulumi:"xrayIndex"`
}

// A collection of values returned by getFederatedNpmRepository.
type LookupFederatedNpmRepositoryResult struct {
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
	Members             []GetFederatedNpmRepositoryMember `pulumi:"members"`
	Notes               *string                           `pulumi:"notes"`
	PackageType         string                            `pulumi:"packageType"`
	PriorityResolution  *bool                             `pulumi:"priorityResolution"`
	ProjectEnvironments []string                          `pulumi:"projectEnvironments"`
	ProjectKey          *string                           `pulumi:"projectKey"`
	PropertySets        []string                          `pulumi:"propertySets"`
	RepoLayoutRef       *string                           `pulumi:"repoLayoutRef"`
	XrayIndex           *bool                             `pulumi:"xrayIndex"`
}

func LookupFederatedNpmRepositoryOutput(ctx *pulumi.Context, args LookupFederatedNpmRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedNpmRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedNpmRepositoryResult, error) {
			args := v.(LookupFederatedNpmRepositoryArgs)
			r, err := LookupFederatedNpmRepository(ctx, &args, opts...)
			var s LookupFederatedNpmRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFederatedNpmRepositoryResultOutput)
}

// A collection of arguments for invoking getFederatedNpmRepository.
type LookupFederatedNpmRepositoryOutputArgs struct {
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
	Members             GetFederatedNpmRepositoryMemberArrayInput `pulumi:"members"`
	Notes               pulumi.StringPtrInput                     `pulumi:"notes"`
	PriorityResolution  pulumi.BoolPtrInput                       `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayInput                   `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput                     `pulumi:"projectKey"`
	PropertySets        pulumi.StringArrayInput                   `pulumi:"propertySets"`
	RepoLayoutRef       pulumi.StringPtrInput                     `pulumi:"repoLayoutRef"`
	XrayIndex           pulumi.BoolPtrInput                       `pulumi:"xrayIndex"`
}

func (LookupFederatedNpmRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedNpmRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedNpmRepository.
type LookupFederatedNpmRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedNpmRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedNpmRepositoryResult)(nil)).Elem()
}

func (o LookupFederatedNpmRepositoryResultOutput) ToLookupFederatedNpmRepositoryResultOutput() LookupFederatedNpmRepositoryResultOutput {
	return o
}

func (o LookupFederatedNpmRepositoryResultOutput) ToLookupFederatedNpmRepositoryResultOutputWithContext(ctx context.Context) LookupFederatedNpmRepositoryResultOutput {
	return o
}

func (o LookupFederatedNpmRepositoryResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupFederatedNpmRepositoryResult] {
	return pulumix.Output[LookupFederatedNpmRepositoryResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupFederatedNpmRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedNpmRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedNpmRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedNpmRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedNpmRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedNpmRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedNpmRepositoryResultOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedNpmRepositoryResult) *bool { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedNpmRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedNpmRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedNpmRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedNpmRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedNpmRepositoryResultOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedNpmRepositoryResult) string { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFederatedNpmRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedNpmRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedNpmRepositoryResultOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedNpmRepositoryResult) string { return v.IncludesPattern }).(pulumi.StringOutput)
}

func (o LookupFederatedNpmRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedNpmRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o LookupFederatedNpmRepositoryResultOutput) Members() GetFederatedNpmRepositoryMemberArrayOutput {
	return o.ApplyT(func(v LookupFederatedNpmRepositoryResult) []GetFederatedNpmRepositoryMember { return v.Members }).(GetFederatedNpmRepositoryMemberArrayOutput)
}

func (o LookupFederatedNpmRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedNpmRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedNpmRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedNpmRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupFederatedNpmRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedNpmRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedNpmRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedNpmRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedNpmRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedNpmRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedNpmRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedNpmRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedNpmRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedNpmRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedNpmRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedNpmRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedNpmRepositoryResultOutput{})
}
