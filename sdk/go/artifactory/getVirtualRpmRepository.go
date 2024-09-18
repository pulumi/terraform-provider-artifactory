// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a virtual Rpm repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.LookupVirtualRpmRepository(ctx, &artifactory.LookupVirtualRpmRepositoryArgs{
//				Key: "virtual-rpm",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupVirtualRpmRepository(ctx *pulumi.Context, args *LookupVirtualRpmRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupVirtualRpmRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualRpmRepositoryResult
	err := ctx.Invoke("artifactory:index/getVirtualRpmRepository:getVirtualRpmRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualRpmRepository.
type LookupVirtualRpmRepositoryArgs struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool   `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         *string `pulumi:"defaultDeploymentRepo"`
	Description                                   *string `pulumi:"description"`
	ExcludesPattern                               *string `pulumi:"excludesPattern"`
	IncludesPattern                               *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key   string  `pulumi:"key"`
	Notes *string `pulumi:"notes"`
	// (Optional) The primary GPG key to be used to sign packages.
	PrimaryKeypairRef   *string  `pulumi:"primaryKeypairRef"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	Repositories        []string `pulumi:"repositories"`
	// (Optional) The secondary GPG key to be used to sign packages.
	SecondaryKeypairRef *string `pulumi:"secondaryKeypairRef"`
}

// A collection of values returned by getVirtualRpmRepository.
type LookupVirtualRpmRepositoryResult struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool   `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         *string `pulumi:"defaultDeploymentRepo"`
	Description                                   *string `pulumi:"description"`
	ExcludesPattern                               *string `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id              string  `pulumi:"id"`
	IncludesPattern *string `pulumi:"includesPattern"`
	Key             string  `pulumi:"key"`
	Notes           *string `pulumi:"notes"`
	PackageType     string  `pulumi:"packageType"`
	// (Optional) The primary GPG key to be used to sign packages.
	PrimaryKeypairRef   *string  `pulumi:"primaryKeypairRef"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	Repositories        []string `pulumi:"repositories"`
	// (Optional) The secondary GPG key to be used to sign packages.
	SecondaryKeypairRef *string `pulumi:"secondaryKeypairRef"`
}

func LookupVirtualRpmRepositoryOutput(ctx *pulumi.Context, args LookupVirtualRpmRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualRpmRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualRpmRepositoryResult, error) {
			args := v.(LookupVirtualRpmRepositoryArgs)
			r, err := LookupVirtualRpmRepository(ctx, &args, opts...)
			var s LookupVirtualRpmRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualRpmRepositoryResultOutput)
}

// A collection of arguments for invoking getVirtualRpmRepository.
type LookupVirtualRpmRepositoryOutputArgs struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput   `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         pulumi.StringPtrInput `pulumi:"defaultDeploymentRepo"`
	Description                                   pulumi.StringPtrInput `pulumi:"description"`
	ExcludesPattern                               pulumi.StringPtrInput `pulumi:"excludesPattern"`
	IncludesPattern                               pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key   pulumi.StringInput    `pulumi:"key"`
	Notes pulumi.StringPtrInput `pulumi:"notes"`
	// (Optional) The primary GPG key to be used to sign packages.
	PrimaryKeypairRef   pulumi.StringPtrInput   `pulumi:"primaryKeypairRef"`
	ProjectEnvironments pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput   `pulumi:"projectKey"`
	RepoLayoutRef       pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	Repositories        pulumi.StringArrayInput `pulumi:"repositories"`
	// (Optional) The secondary GPG key to be used to sign packages.
	SecondaryKeypairRef pulumi.StringPtrInput `pulumi:"secondaryKeypairRef"`
}

func (LookupVirtualRpmRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualRpmRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getVirtualRpmRepository.
type LookupVirtualRpmRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualRpmRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualRpmRepositoryResult)(nil)).Elem()
}

func (o LookupVirtualRpmRepositoryResultOutput) ToLookupVirtualRpmRepositoryResultOutput() LookupVirtualRpmRepositoryResultOutput {
	return o
}

func (o LookupVirtualRpmRepositoryResultOutput) ToLookupVirtualRpmRepositoryResultOutputWithContext(ctx context.Context) LookupVirtualRpmRepositoryResultOutput {
	return o
}

func (o LookupVirtualRpmRepositoryResultOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualRpmRepositoryResult) *bool { return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualRpmRepositoryResultOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualRpmRepositoryResult) *string { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualRpmRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualRpmRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualRpmRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualRpmRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVirtualRpmRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualRpmRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualRpmRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualRpmRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualRpmRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualRpmRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupVirtualRpmRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualRpmRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualRpmRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualRpmRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

// (Optional) The primary GPG key to be used to sign packages.
func (o LookupVirtualRpmRepositoryResultOutput) PrimaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualRpmRepositoryResult) *string { return v.PrimaryKeypairRef }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualRpmRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualRpmRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupVirtualRpmRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualRpmRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualRpmRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualRpmRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualRpmRepositoryResultOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualRpmRepositoryResult) []string { return v.Repositories }).(pulumi.StringArrayOutput)
}

// (Optional) The secondary GPG key to be used to sign packages.
func (o LookupVirtualRpmRepositoryResultOutput) SecondaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualRpmRepositoryResult) *string { return v.SecondaryKeypairRef }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualRpmRepositoryResultOutput{})
}
