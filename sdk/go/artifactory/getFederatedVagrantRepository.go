// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a federated Vagrant repository.
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
//			_, err := artifactory.LookupFederatedVagrantRepository(ctx, &artifactory.LookupFederatedVagrantRepositoryArgs{
//				Key: "federated-test-vagrant-repo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFederatedVagrantRepository(ctx *pulumi.Context, args *LookupFederatedVagrantRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupFederatedVagrantRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFederatedVagrantRepositoryResult
	err := ctx.Invoke("artifactory:index/getFederatedVagrantRepository:getFederatedVagrantRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedVagrantRepository.
type LookupFederatedVagrantRepositoryArgs struct {
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
	Members             []GetFederatedVagrantRepositoryMember `pulumi:"members"`
	Notes               *string                               `pulumi:"notes"`
	PriorityResolution  *bool                                 `pulumi:"priorityResolution"`
	ProjectEnvironments []string                              `pulumi:"projectEnvironments"`
	ProjectKey          *string                               `pulumi:"projectKey"`
	PropertySets        []string                              `pulumi:"propertySets"`
	RepoLayoutRef       *string                               `pulumi:"repoLayoutRef"`
	XrayIndex           *bool                                 `pulumi:"xrayIndex"`
}

// A collection of values returned by getFederatedVagrantRepository.
type LookupFederatedVagrantRepositoryResult struct {
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
	Members             []GetFederatedVagrantRepositoryMember `pulumi:"members"`
	Notes               *string                               `pulumi:"notes"`
	PackageType         string                                `pulumi:"packageType"`
	PriorityResolution  *bool                                 `pulumi:"priorityResolution"`
	ProjectEnvironments []string                              `pulumi:"projectEnvironments"`
	ProjectKey          *string                               `pulumi:"projectKey"`
	PropertySets        []string                              `pulumi:"propertySets"`
	RepoLayoutRef       *string                               `pulumi:"repoLayoutRef"`
	XrayIndex           *bool                                 `pulumi:"xrayIndex"`
}

func LookupFederatedVagrantRepositoryOutput(ctx *pulumi.Context, args LookupFederatedVagrantRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedVagrantRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedVagrantRepositoryResult, error) {
			args := v.(LookupFederatedVagrantRepositoryArgs)
			r, err := LookupFederatedVagrantRepository(ctx, &args, opts...)
			var s LookupFederatedVagrantRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFederatedVagrantRepositoryResultOutput)
}

// A collection of arguments for invoking getFederatedVagrantRepository.
type LookupFederatedVagrantRepositoryOutputArgs struct {
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
	Members             GetFederatedVagrantRepositoryMemberArrayInput `pulumi:"members"`
	Notes               pulumi.StringPtrInput                         `pulumi:"notes"`
	PriorityResolution  pulumi.BoolPtrInput                           `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayInput                       `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput                         `pulumi:"projectKey"`
	PropertySets        pulumi.StringArrayInput                       `pulumi:"propertySets"`
	RepoLayoutRef       pulumi.StringPtrInput                         `pulumi:"repoLayoutRef"`
	XrayIndex           pulumi.BoolPtrInput                           `pulumi:"xrayIndex"`
}

func (LookupFederatedVagrantRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedVagrantRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedVagrantRepository.
type LookupFederatedVagrantRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedVagrantRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedVagrantRepositoryResult)(nil)).Elem()
}

func (o LookupFederatedVagrantRepositoryResultOutput) ToLookupFederatedVagrantRepositoryResultOutput() LookupFederatedVagrantRepositoryResultOutput {
	return o
}

func (o LookupFederatedVagrantRepositoryResultOutput) ToLookupFederatedVagrantRepositoryResultOutputWithContext(ctx context.Context) LookupFederatedVagrantRepositoryResultOutput {
	return o
}

func (o LookupFederatedVagrantRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedVagrantRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedVagrantRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedVagrantRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedVagrantRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedVagrantRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedVagrantRepositoryResultOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedVagrantRepositoryResult) *bool { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedVagrantRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedVagrantRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedVagrantRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedVagrantRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedVagrantRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedVagrantRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFederatedVagrantRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedVagrantRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedVagrantRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedVagrantRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedVagrantRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedVagrantRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o LookupFederatedVagrantRepositoryResultOutput) Members() GetFederatedVagrantRepositoryMemberArrayOutput {
	return o.ApplyT(func(v LookupFederatedVagrantRepositoryResult) []GetFederatedVagrantRepositoryMember { return v.Members }).(GetFederatedVagrantRepositoryMemberArrayOutput)
}

func (o LookupFederatedVagrantRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedVagrantRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedVagrantRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedVagrantRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupFederatedVagrantRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedVagrantRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedVagrantRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedVagrantRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedVagrantRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedVagrantRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedVagrantRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedVagrantRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedVagrantRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedVagrantRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedVagrantRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedVagrantRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedVagrantRepositoryResultOutput{})
}
