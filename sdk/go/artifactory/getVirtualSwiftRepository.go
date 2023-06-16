// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a virtual Swift repository.
func LookupVirtualSwiftRepository(ctx *pulumi.Context, args *LookupVirtualSwiftRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupVirtualSwiftRepositoryResult, error) {
	var rv LookupVirtualSwiftRepositoryResult
	err := ctx.Invoke("artifactory:index/getVirtualSwiftRepository:getVirtualSwiftRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualSwiftRepository.
type LookupVirtualSwiftRepositoryArgs struct {
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

// A collection of values returned by getVirtualSwiftRepository.
type LookupVirtualSwiftRepositoryResult struct {
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

func LookupVirtualSwiftRepositoryOutput(ctx *pulumi.Context, args LookupVirtualSwiftRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualSwiftRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualSwiftRepositoryResult, error) {
			args := v.(LookupVirtualSwiftRepositoryArgs)
			r, err := LookupVirtualSwiftRepository(ctx, &args, opts...)
			var s LookupVirtualSwiftRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualSwiftRepositoryResultOutput)
}

// A collection of arguments for invoking getVirtualSwiftRepository.
type LookupVirtualSwiftRepositoryOutputArgs struct {
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

func (LookupVirtualSwiftRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualSwiftRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getVirtualSwiftRepository.
type LookupVirtualSwiftRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualSwiftRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualSwiftRepositoryResult)(nil)).Elem()
}

func (o LookupVirtualSwiftRepositoryResultOutput) ToLookupVirtualSwiftRepositoryResultOutput() LookupVirtualSwiftRepositoryResultOutput {
	return o
}

func (o LookupVirtualSwiftRepositoryResultOutput) ToLookupVirtualSwiftRepositoryResultOutputWithContext(ctx context.Context) LookupVirtualSwiftRepositoryResultOutput {
	return o
}

func (o LookupVirtualSwiftRepositoryResultOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualSwiftRepositoryResult) *bool {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualSwiftRepositoryResultOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualSwiftRepositoryResult) *string { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualSwiftRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualSwiftRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualSwiftRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualSwiftRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVirtualSwiftRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualSwiftRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualSwiftRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualSwiftRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualSwiftRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualSwiftRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupVirtualSwiftRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualSwiftRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualSwiftRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualSwiftRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupVirtualSwiftRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualSwiftRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupVirtualSwiftRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualSwiftRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualSwiftRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualSwiftRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualSwiftRepositoryResultOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualSwiftRepositoryResult) []string { return v.Repositories }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualSwiftRepositoryResultOutput{})
}
