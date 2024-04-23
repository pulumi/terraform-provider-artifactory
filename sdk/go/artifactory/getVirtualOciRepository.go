// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v6/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a virtual OCI repository.
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
//			_, err := artifactory.LookupVirtualOciRepository(ctx, &artifactory.LookupVirtualOciRepositoryArgs{
//				Key: "my-oci-virtual",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupVirtualOciRepository(ctx *pulumi.Context, args *LookupVirtualOciRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupVirtualOciRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualOciRepositoryResult
	err := ctx.Invoke("artifactory:index/getVirtualOciRepository:getVirtualOciRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualOciRepository.
type LookupVirtualOciRepositoryArgs struct {
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
	// (Optional) When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
	ResolveOciTagsByTimestamp *bool `pulumi:"resolveOciTagsByTimestamp"`
}

// A collection of values returned by getVirtualOciRepository.
type LookupVirtualOciRepositoryResult struct {
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
	// (Optional) When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
	ResolveOciTagsByTimestamp *bool `pulumi:"resolveOciTagsByTimestamp"`
}

func LookupVirtualOciRepositoryOutput(ctx *pulumi.Context, args LookupVirtualOciRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualOciRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualOciRepositoryResult, error) {
			args := v.(LookupVirtualOciRepositoryArgs)
			r, err := LookupVirtualOciRepository(ctx, &args, opts...)
			var s LookupVirtualOciRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualOciRepositoryResultOutput)
}

// A collection of arguments for invoking getVirtualOciRepository.
type LookupVirtualOciRepositoryOutputArgs struct {
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
	// (Optional) When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
	ResolveOciTagsByTimestamp pulumi.BoolPtrInput `pulumi:"resolveOciTagsByTimestamp"`
}

func (LookupVirtualOciRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualOciRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getVirtualOciRepository.
type LookupVirtualOciRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualOciRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualOciRepositoryResult)(nil)).Elem()
}

func (o LookupVirtualOciRepositoryResultOutput) ToLookupVirtualOciRepositoryResultOutput() LookupVirtualOciRepositoryResultOutput {
	return o
}

func (o LookupVirtualOciRepositoryResultOutput) ToLookupVirtualOciRepositoryResultOutputWithContext(ctx context.Context) LookupVirtualOciRepositoryResultOutput {
	return o
}

func (o LookupVirtualOciRepositoryResultOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualOciRepositoryResult) *bool { return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualOciRepositoryResultOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualOciRepositoryResult) *string { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualOciRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualOciRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualOciRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualOciRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVirtualOciRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualOciRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualOciRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualOciRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualOciRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualOciRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupVirtualOciRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualOciRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualOciRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualOciRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupVirtualOciRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualOciRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupVirtualOciRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualOciRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualOciRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualOciRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualOciRepositoryResultOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualOciRepositoryResult) []string { return v.Repositories }).(pulumi.StringArrayOutput)
}

// (Optional) When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
func (o LookupVirtualOciRepositoryResultOutput) ResolveOciTagsByTimestamp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualOciRepositoryResult) *bool { return v.ResolveOciTagsByTimestamp }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualOciRepositoryResultOutput{})
}
