// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v6/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a virtual Git LFS repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v6/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.LookupVirtualGitlfsRepository(ctx, &artifactory.LookupVirtualGitlfsRepositoryArgs{
//				Key: "virtual-gitlfs",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupVirtualGitlfsRepository(ctx *pulumi.Context, args *LookupVirtualGitlfsRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupVirtualGitlfsRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualGitlfsRepositoryResult
	err := ctx.Invoke("artifactory:index/getVirtualGitlfsRepository:getVirtualGitlfsRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualGitlfsRepository.
type LookupVirtualGitlfsRepositoryArgs struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool   `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         *string `pulumi:"defaultDeploymentRepo"`
	Description                                   *string `pulumi:"description"`
	ExcludesPattern                               *string `pulumi:"excludesPattern"`
	IncludesPattern                               *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                 string   `pulumi:"key"`
	Notes               *string  `pulumi:"notes"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	Repositories        []string `pulumi:"repositories"`
}

// A collection of values returned by getVirtualGitlfsRepository.
type LookupVirtualGitlfsRepositoryResult struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool   `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         *string `pulumi:"defaultDeploymentRepo"`
	Description                                   *string `pulumi:"description"`
	ExcludesPattern                               *string `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string   `pulumi:"id"`
	IncludesPattern     *string  `pulumi:"includesPattern"`
	Key                 string   `pulumi:"key"`
	Notes               *string  `pulumi:"notes"`
	PackageType         string   `pulumi:"packageType"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	Repositories        []string `pulumi:"repositories"`
}

func LookupVirtualGitlfsRepositoryOutput(ctx *pulumi.Context, args LookupVirtualGitlfsRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualGitlfsRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualGitlfsRepositoryResult, error) {
			args := v.(LookupVirtualGitlfsRepositoryArgs)
			r, err := LookupVirtualGitlfsRepository(ctx, &args, opts...)
			var s LookupVirtualGitlfsRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualGitlfsRepositoryResultOutput)
}

// A collection of arguments for invoking getVirtualGitlfsRepository.
type LookupVirtualGitlfsRepositoryOutputArgs struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput   `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         pulumi.StringPtrInput `pulumi:"defaultDeploymentRepo"`
	Description                                   pulumi.StringPtrInput `pulumi:"description"`
	ExcludesPattern                               pulumi.StringPtrInput `pulumi:"excludesPattern"`
	IncludesPattern                               pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                 pulumi.StringInput      `pulumi:"key"`
	Notes               pulumi.StringPtrInput   `pulumi:"notes"`
	ProjectEnvironments pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput   `pulumi:"projectKey"`
	RepoLayoutRef       pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	Repositories        pulumi.StringArrayInput `pulumi:"repositories"`
}

func (LookupVirtualGitlfsRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualGitlfsRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getVirtualGitlfsRepository.
type LookupVirtualGitlfsRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualGitlfsRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualGitlfsRepositoryResult)(nil)).Elem()
}

func (o LookupVirtualGitlfsRepositoryResultOutput) ToLookupVirtualGitlfsRepositoryResultOutput() LookupVirtualGitlfsRepositoryResultOutput {
	return o
}

func (o LookupVirtualGitlfsRepositoryResultOutput) ToLookupVirtualGitlfsRepositoryResultOutputWithContext(ctx context.Context) LookupVirtualGitlfsRepositoryResultOutput {
	return o
}

func (o LookupVirtualGitlfsRepositoryResultOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualGitlfsRepositoryResult) *bool {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualGitlfsRepositoryResultOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualGitlfsRepositoryResult) *string { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualGitlfsRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualGitlfsRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualGitlfsRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualGitlfsRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVirtualGitlfsRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualGitlfsRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualGitlfsRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualGitlfsRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualGitlfsRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualGitlfsRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupVirtualGitlfsRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualGitlfsRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualGitlfsRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualGitlfsRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupVirtualGitlfsRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualGitlfsRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupVirtualGitlfsRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualGitlfsRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualGitlfsRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualGitlfsRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualGitlfsRepositoryResultOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualGitlfsRepositoryResult) []string { return v.Repositories }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualGitlfsRepositoryResultOutput{})
}
