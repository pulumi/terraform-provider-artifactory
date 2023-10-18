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

// Retrieves a federated Cocoapods repository.
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
//			_, err := artifactory.LookupFederatedCocoapodsRepository(ctx, &artifactory.LookupFederatedCocoapodsRepositoryArgs{
//				Key: "federated-test-cocoapods-repo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFederatedCocoapodsRepository(ctx *pulumi.Context, args *LookupFederatedCocoapodsRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupFederatedCocoapodsRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFederatedCocoapodsRepositoryResult
	err := ctx.Invoke("artifactory:index/getFederatedCocoapodsRepository:getFederatedCocoapodsRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedCocoapodsRepository.
type LookupFederatedCocoapodsRepositoryArgs struct {
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
	Members             []GetFederatedCocoapodsRepositoryMember `pulumi:"members"`
	Notes               *string                                 `pulumi:"notes"`
	PriorityResolution  *bool                                   `pulumi:"priorityResolution"`
	ProjectEnvironments []string                                `pulumi:"projectEnvironments"`
	ProjectKey          *string                                 `pulumi:"projectKey"`
	PropertySets        []string                                `pulumi:"propertySets"`
	RepoLayoutRef       *string                                 `pulumi:"repoLayoutRef"`
	XrayIndex           *bool                                   `pulumi:"xrayIndex"`
}

// A collection of values returned by getFederatedCocoapodsRepository.
type LookupFederatedCocoapodsRepositoryResult struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id              string  `pulumi:"id"`
	IncludesPattern *string `pulumi:"includesPattern"`
	Key             string  `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             []GetFederatedCocoapodsRepositoryMember `pulumi:"members"`
	Notes               *string                                 `pulumi:"notes"`
	PackageType         string                                  `pulumi:"packageType"`
	PriorityResolution  *bool                                   `pulumi:"priorityResolution"`
	ProjectEnvironments []string                                `pulumi:"projectEnvironments"`
	ProjectKey          *string                                 `pulumi:"projectKey"`
	PropertySets        []string                                `pulumi:"propertySets"`
	RepoLayoutRef       *string                                 `pulumi:"repoLayoutRef"`
	XrayIndex           *bool                                   `pulumi:"xrayIndex"`
}

func LookupFederatedCocoapodsRepositoryOutput(ctx *pulumi.Context, args LookupFederatedCocoapodsRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedCocoapodsRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedCocoapodsRepositoryResult, error) {
			args := v.(LookupFederatedCocoapodsRepositoryArgs)
			r, err := LookupFederatedCocoapodsRepository(ctx, &args, opts...)
			var s LookupFederatedCocoapodsRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFederatedCocoapodsRepositoryResultOutput)
}

// A collection of arguments for invoking getFederatedCocoapodsRepository.
type LookupFederatedCocoapodsRepositoryOutputArgs struct {
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
	Members             GetFederatedCocoapodsRepositoryMemberArrayInput `pulumi:"members"`
	Notes               pulumi.StringPtrInput                           `pulumi:"notes"`
	PriorityResolution  pulumi.BoolPtrInput                             `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayInput                         `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput                           `pulumi:"projectKey"`
	PropertySets        pulumi.StringArrayInput                         `pulumi:"propertySets"`
	RepoLayoutRef       pulumi.StringPtrInput                           `pulumi:"repoLayoutRef"`
	XrayIndex           pulumi.BoolPtrInput                             `pulumi:"xrayIndex"`
}

func (LookupFederatedCocoapodsRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedCocoapodsRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedCocoapodsRepository.
type LookupFederatedCocoapodsRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedCocoapodsRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedCocoapodsRepositoryResult)(nil)).Elem()
}

func (o LookupFederatedCocoapodsRepositoryResultOutput) ToLookupFederatedCocoapodsRepositoryResultOutput() LookupFederatedCocoapodsRepositoryResultOutput {
	return o
}

func (o LookupFederatedCocoapodsRepositoryResultOutput) ToLookupFederatedCocoapodsRepositoryResultOutputWithContext(ctx context.Context) LookupFederatedCocoapodsRepositoryResultOutput {
	return o
}

func (o LookupFederatedCocoapodsRepositoryResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupFederatedCocoapodsRepositoryResult] {
	return pulumix.Output[LookupFederatedCocoapodsRepositoryResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupFederatedCocoapodsRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCocoapodsRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCocoapodsRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCocoapodsRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCocoapodsRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCocoapodsRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCocoapodsRepositoryResultOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCocoapodsRepositoryResult) *bool { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCocoapodsRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedCocoapodsRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedCocoapodsRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCocoapodsRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCocoapodsRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedCocoapodsRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFederatedCocoapodsRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedCocoapodsRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedCocoapodsRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedCocoapodsRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedCocoapodsRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedCocoapodsRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o LookupFederatedCocoapodsRepositoryResultOutput) Members() GetFederatedCocoapodsRepositoryMemberArrayOutput {
	return o.ApplyT(func(v LookupFederatedCocoapodsRepositoryResult) []GetFederatedCocoapodsRepositoryMember {
		return v.Members
	}).(GetFederatedCocoapodsRepositoryMemberArrayOutput)
}

func (o LookupFederatedCocoapodsRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedCocoapodsRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedCocoapodsRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedCocoapodsRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupFederatedCocoapodsRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCocoapodsRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCocoapodsRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedCocoapodsRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedCocoapodsRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedCocoapodsRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedCocoapodsRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedCocoapodsRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedCocoapodsRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedCocoapodsRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedCocoapodsRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCocoapodsRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedCocoapodsRepositoryResultOutput{})
}
