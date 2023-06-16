// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a federated Gitlfs repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v3/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.GetFederatedGitlfsRepository(ctx, &artifactory.GetFederatedGitlfsRepositoryArgs{
//				Key: "federated-test-gitlfs-repo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetFederatedGitlfsRepository(ctx *pulumi.Context, args *GetFederatedGitlfsRepositoryArgs, opts ...pulumi.InvokeOption) (*GetFederatedGitlfsRepositoryResult, error) {
	var rv GetFederatedGitlfsRepositoryResult
	err := ctx.Invoke("artifactory:index/getFederatedGitlfsRepository:getFederatedGitlfsRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedGitlfsRepository.
type GetFederatedGitlfsRepositoryArgs struct {
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
	Members             []GetFederatedGitlfsRepositoryMember `pulumi:"members"`
	Notes               *string                              `pulumi:"notes"`
	PriorityResolution  *bool                                `pulumi:"priorityResolution"`
	ProjectEnvironments []string                             `pulumi:"projectEnvironments"`
	ProjectKey          *string                              `pulumi:"projectKey"`
	PropertySets        []string                             `pulumi:"propertySets"`
	RepoLayoutRef       *string                              `pulumi:"repoLayoutRef"`
	XrayIndex           *bool                                `pulumi:"xrayIndex"`
}

// A collection of values returned by getFederatedGitlfsRepository.
type GetFederatedGitlfsRepositoryResult struct {
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
	Members             []GetFederatedGitlfsRepositoryMember `pulumi:"members"`
	Notes               *string                              `pulumi:"notes"`
	PackageType         string                               `pulumi:"packageType"`
	PriorityResolution  *bool                                `pulumi:"priorityResolution"`
	ProjectEnvironments []string                             `pulumi:"projectEnvironments"`
	ProjectKey          *string                              `pulumi:"projectKey"`
	PropertySets        []string                             `pulumi:"propertySets"`
	RepoLayoutRef       *string                              `pulumi:"repoLayoutRef"`
	XrayIndex           *bool                                `pulumi:"xrayIndex"`
}

func GetFederatedGitlfsRepositoryOutput(ctx *pulumi.Context, args GetFederatedGitlfsRepositoryOutputArgs, opts ...pulumi.InvokeOption) GetFederatedGitlfsRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFederatedGitlfsRepositoryResult, error) {
			args := v.(GetFederatedGitlfsRepositoryArgs)
			r, err := GetFederatedGitlfsRepository(ctx, &args, opts...)
			var s GetFederatedGitlfsRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFederatedGitlfsRepositoryResultOutput)
}

// A collection of arguments for invoking getFederatedGitlfsRepository.
type GetFederatedGitlfsRepositoryOutputArgs struct {
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
	Members             GetFederatedGitlfsRepositoryMemberArrayInput `pulumi:"members"`
	Notes               pulumi.StringPtrInput                        `pulumi:"notes"`
	PriorityResolution  pulumi.BoolPtrInput                          `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayInput                      `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput                        `pulumi:"projectKey"`
	PropertySets        pulumi.StringArrayInput                      `pulumi:"propertySets"`
	RepoLayoutRef       pulumi.StringPtrInput                        `pulumi:"repoLayoutRef"`
	XrayIndex           pulumi.BoolPtrInput                          `pulumi:"xrayIndex"`
}

func (GetFederatedGitlfsRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFederatedGitlfsRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedGitlfsRepository.
type GetFederatedGitlfsRepositoryResultOutput struct{ *pulumi.OutputState }

func (GetFederatedGitlfsRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFederatedGitlfsRepositoryResult)(nil)).Elem()
}

func (o GetFederatedGitlfsRepositoryResultOutput) ToGetFederatedGitlfsRepositoryResultOutput() GetFederatedGitlfsRepositoryResultOutput {
	return o
}

func (o GetFederatedGitlfsRepositoryResultOutput) ToGetFederatedGitlfsRepositoryResultOutputWithContext(ctx context.Context) GetFederatedGitlfsRepositoryResultOutput {
	return o
}

func (o GetFederatedGitlfsRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetFederatedGitlfsRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o GetFederatedGitlfsRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetFederatedGitlfsRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o GetFederatedGitlfsRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetFederatedGitlfsRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o GetFederatedGitlfsRepositoryResultOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetFederatedGitlfsRepositoryResult) *bool { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

func (o GetFederatedGitlfsRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFederatedGitlfsRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GetFederatedGitlfsRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetFederatedGitlfsRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o GetFederatedGitlfsRepositoryResultOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v GetFederatedGitlfsRepositoryResult) string { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFederatedGitlfsRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFederatedGitlfsRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetFederatedGitlfsRepositoryResultOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v GetFederatedGitlfsRepositoryResult) string { return v.IncludesPattern }).(pulumi.StringOutput)
}

func (o GetFederatedGitlfsRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GetFederatedGitlfsRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o GetFederatedGitlfsRepositoryResultOutput) Members() GetFederatedGitlfsRepositoryMemberArrayOutput {
	return o.ApplyT(func(v GetFederatedGitlfsRepositoryResult) []GetFederatedGitlfsRepositoryMember { return v.Members }).(GetFederatedGitlfsRepositoryMemberArrayOutput)
}

func (o GetFederatedGitlfsRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFederatedGitlfsRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o GetFederatedGitlfsRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v GetFederatedGitlfsRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o GetFederatedGitlfsRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetFederatedGitlfsRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o GetFederatedGitlfsRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFederatedGitlfsRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o GetFederatedGitlfsRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFederatedGitlfsRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o GetFederatedGitlfsRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFederatedGitlfsRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o GetFederatedGitlfsRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFederatedGitlfsRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o GetFederatedGitlfsRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetFederatedGitlfsRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFederatedGitlfsRepositoryResultOutput{})
}
