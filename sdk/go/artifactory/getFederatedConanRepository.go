// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a federated Conan repository.
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
//			_, err := artifactory.LookupFederatedConanRepository(ctx, &artifactory.LookupFederatedConanRepositoryArgs{
//				Key: "federated-test-conan-repo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFederatedConanRepository(ctx *pulumi.Context, args *LookupFederatedConanRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupFederatedConanRepositoryResult, error) {
	var rv LookupFederatedConanRepositoryResult
	err := ctx.Invoke("artifactory:index/getFederatedConanRepository:getFederatedConanRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedConanRepository.
type LookupFederatedConanRepositoryArgs struct {
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
	Members             []GetFederatedConanRepositoryMember `pulumi:"members"`
	Notes               *string                             `pulumi:"notes"`
	PriorityResolution  *bool                               `pulumi:"priorityResolution"`
	ProjectEnvironments []string                            `pulumi:"projectEnvironments"`
	ProjectKey          *string                             `pulumi:"projectKey"`
	PropertySets        []string                            `pulumi:"propertySets"`
	RepoLayoutRef       *string                             `pulumi:"repoLayoutRef"`
	XrayIndex           *bool                               `pulumi:"xrayIndex"`
}

// A collection of values returned by getFederatedConanRepository.
type LookupFederatedConanRepositoryResult struct {
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
	Members             []GetFederatedConanRepositoryMember `pulumi:"members"`
	Notes               *string                             `pulumi:"notes"`
	PackageType         string                              `pulumi:"packageType"`
	PriorityResolution  *bool                               `pulumi:"priorityResolution"`
	ProjectEnvironments []string                            `pulumi:"projectEnvironments"`
	ProjectKey          *string                             `pulumi:"projectKey"`
	PropertySets        []string                            `pulumi:"propertySets"`
	RepoLayoutRef       *string                             `pulumi:"repoLayoutRef"`
	XrayIndex           *bool                               `pulumi:"xrayIndex"`
}

func LookupFederatedConanRepositoryOutput(ctx *pulumi.Context, args LookupFederatedConanRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedConanRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedConanRepositoryResult, error) {
			args := v.(LookupFederatedConanRepositoryArgs)
			r, err := LookupFederatedConanRepository(ctx, &args, opts...)
			var s LookupFederatedConanRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFederatedConanRepositoryResultOutput)
}

// A collection of arguments for invoking getFederatedConanRepository.
type LookupFederatedConanRepositoryOutputArgs struct {
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
	Members             GetFederatedConanRepositoryMemberArrayInput `pulumi:"members"`
	Notes               pulumi.StringPtrInput                       `pulumi:"notes"`
	PriorityResolution  pulumi.BoolPtrInput                         `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayInput                     `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput                       `pulumi:"projectKey"`
	PropertySets        pulumi.StringArrayInput                     `pulumi:"propertySets"`
	RepoLayoutRef       pulumi.StringPtrInput                       `pulumi:"repoLayoutRef"`
	XrayIndex           pulumi.BoolPtrInput                         `pulumi:"xrayIndex"`
}

func (LookupFederatedConanRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedConanRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedConanRepository.
type LookupFederatedConanRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedConanRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedConanRepositoryResult)(nil)).Elem()
}

func (o LookupFederatedConanRepositoryResultOutput) ToLookupFederatedConanRepositoryResultOutput() LookupFederatedConanRepositoryResultOutput {
	return o
}

func (o LookupFederatedConanRepositoryResultOutput) ToLookupFederatedConanRepositoryResultOutputWithContext(ctx context.Context) LookupFederatedConanRepositoryResultOutput {
	return o
}

func (o LookupFederatedConanRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedConanRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedConanRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedConanRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedConanRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedConanRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedConanRepositoryResultOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedConanRepositoryResult) *bool { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedConanRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedConanRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedConanRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedConanRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedConanRepositoryResultOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedConanRepositoryResult) string { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFederatedConanRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedConanRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedConanRepositoryResultOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedConanRepositoryResult) string { return v.IncludesPattern }).(pulumi.StringOutput)
}

func (o LookupFederatedConanRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedConanRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o LookupFederatedConanRepositoryResultOutput) Members() GetFederatedConanRepositoryMemberArrayOutput {
	return o.ApplyT(func(v LookupFederatedConanRepositoryResult) []GetFederatedConanRepositoryMember { return v.Members }).(GetFederatedConanRepositoryMemberArrayOutput)
}

func (o LookupFederatedConanRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedConanRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedConanRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedConanRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupFederatedConanRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedConanRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedConanRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedConanRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedConanRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedConanRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedConanRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedConanRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedConanRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedConanRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedConanRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedConanRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedConanRepositoryResultOutput{})
}
