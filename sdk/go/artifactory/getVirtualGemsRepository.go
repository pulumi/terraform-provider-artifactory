// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a virtual Gems repository.
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
//			_, err := artifactory.LookupVirtualGemsRepository(ctx, &artifactory.LookupVirtualGemsRepositoryArgs{
//				Key: "virtual-gems",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupVirtualGemsRepository(ctx *pulumi.Context, args *LookupVirtualGemsRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupVirtualGemsRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualGemsRepositoryResult
	err := ctx.Invoke("artifactory:index/getVirtualGemsRepository:getVirtualGemsRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualGemsRepository.
type LookupVirtualGemsRepositoryArgs struct {
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

// A collection of values returned by getVirtualGemsRepository.
type LookupVirtualGemsRepositoryResult struct {
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

func LookupVirtualGemsRepositoryOutput(ctx *pulumi.Context, args LookupVirtualGemsRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualGemsRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualGemsRepositoryResult, error) {
			args := v.(LookupVirtualGemsRepositoryArgs)
			r, err := LookupVirtualGemsRepository(ctx, &args, opts...)
			var s LookupVirtualGemsRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualGemsRepositoryResultOutput)
}

// A collection of arguments for invoking getVirtualGemsRepository.
type LookupVirtualGemsRepositoryOutputArgs struct {
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

func (LookupVirtualGemsRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualGemsRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getVirtualGemsRepository.
type LookupVirtualGemsRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualGemsRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualGemsRepositoryResult)(nil)).Elem()
}

func (o LookupVirtualGemsRepositoryResultOutput) ToLookupVirtualGemsRepositoryResultOutput() LookupVirtualGemsRepositoryResultOutput {
	return o
}

func (o LookupVirtualGemsRepositoryResultOutput) ToLookupVirtualGemsRepositoryResultOutputWithContext(ctx context.Context) LookupVirtualGemsRepositoryResultOutput {
	return o
}

func (o LookupVirtualGemsRepositoryResultOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualGemsRepositoryResult) *bool {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualGemsRepositoryResultOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualGemsRepositoryResult) *string { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualGemsRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualGemsRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualGemsRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualGemsRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVirtualGemsRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualGemsRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualGemsRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualGemsRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualGemsRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualGemsRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupVirtualGemsRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualGemsRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualGemsRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualGemsRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupVirtualGemsRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualGemsRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupVirtualGemsRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualGemsRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualGemsRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualGemsRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualGemsRepositoryResultOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualGemsRepositoryResult) []string { return v.Repositories }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualGemsRepositoryResultOutput{})
}
