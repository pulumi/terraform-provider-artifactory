// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualCondaRepository(ctx *pulumi.Context, args *LookupVirtualCondaRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupVirtualCondaRepositoryResult, error) {
	var rv LookupVirtualCondaRepositoryResult
	err := ctx.Invoke("artifactory:index/getVirtualCondaRepository:getVirtualCondaRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualCondaRepository.
type LookupVirtualCondaRepositoryArgs struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool    `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         *string  `pulumi:"defaultDeploymentRepo"`
	Description                                   *string  `pulumi:"description"`
	ExcludesPattern                               *string  `pulumi:"excludesPattern"`
	IncludesPattern                               *string  `pulumi:"includesPattern"`
	Key                                           string   `pulumi:"key"`
	Notes                                         *string  `pulumi:"notes"`
	ProjectEnvironments                           []string `pulumi:"projectEnvironments"`
	ProjectKey                                    *string  `pulumi:"projectKey"`
	RepoLayoutRef                                 *string  `pulumi:"repoLayoutRef"`
	Repositories                                  []string `pulumi:"repositories"`
	RetrievalCachePeriodSeconds                   *int     `pulumi:"retrievalCachePeriodSeconds"`
}

// A collection of values returned by getVirtualCondaRepository.
type LookupVirtualCondaRepositoryResult struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool   `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         *string `pulumi:"defaultDeploymentRepo"`
	Description                                   *string `pulumi:"description"`
	ExcludesPattern                               *string `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id                          string   `pulumi:"id"`
	IncludesPattern             *string  `pulumi:"includesPattern"`
	Key                         string   `pulumi:"key"`
	Notes                       *string  `pulumi:"notes"`
	PackageType                 string   `pulumi:"packageType"`
	ProjectEnvironments         []string `pulumi:"projectEnvironments"`
	ProjectKey                  *string  `pulumi:"projectKey"`
	RepoLayoutRef               *string  `pulumi:"repoLayoutRef"`
	Repositories                []string `pulumi:"repositories"`
	RetrievalCachePeriodSeconds *int     `pulumi:"retrievalCachePeriodSeconds"`
}

func LookupVirtualCondaRepositoryOutput(ctx *pulumi.Context, args LookupVirtualCondaRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualCondaRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualCondaRepositoryResult, error) {
			args := v.(LookupVirtualCondaRepositoryArgs)
			r, err := LookupVirtualCondaRepository(ctx, &args, opts...)
			var s LookupVirtualCondaRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualCondaRepositoryResultOutput)
}

// A collection of arguments for invoking getVirtualCondaRepository.
type LookupVirtualCondaRepositoryOutputArgs struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput     `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         pulumi.StringPtrInput   `pulumi:"defaultDeploymentRepo"`
	Description                                   pulumi.StringPtrInput   `pulumi:"description"`
	ExcludesPattern                               pulumi.StringPtrInput   `pulumi:"excludesPattern"`
	IncludesPattern                               pulumi.StringPtrInput   `pulumi:"includesPattern"`
	Key                                           pulumi.StringInput      `pulumi:"key"`
	Notes                                         pulumi.StringPtrInput   `pulumi:"notes"`
	ProjectEnvironments                           pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey                                    pulumi.StringPtrInput   `pulumi:"projectKey"`
	RepoLayoutRef                                 pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	Repositories                                  pulumi.StringArrayInput `pulumi:"repositories"`
	RetrievalCachePeriodSeconds                   pulumi.IntPtrInput      `pulumi:"retrievalCachePeriodSeconds"`
}

func (LookupVirtualCondaRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualCondaRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getVirtualCondaRepository.
type LookupVirtualCondaRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualCondaRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualCondaRepositoryResult)(nil)).Elem()
}

func (o LookupVirtualCondaRepositoryResultOutput) ToLookupVirtualCondaRepositoryResultOutput() LookupVirtualCondaRepositoryResultOutput {
	return o
}

func (o LookupVirtualCondaRepositoryResultOutput) ToLookupVirtualCondaRepositoryResultOutputWithContext(ctx context.Context) LookupVirtualCondaRepositoryResultOutput {
	return o
}

func (o LookupVirtualCondaRepositoryResultOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualCondaRepositoryResult) *bool {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualCondaRepositoryResultOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualCondaRepositoryResult) *string { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualCondaRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualCondaRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualCondaRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualCondaRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVirtualCondaRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualCondaRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualCondaRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualCondaRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualCondaRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualCondaRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupVirtualCondaRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualCondaRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualCondaRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualCondaRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupVirtualCondaRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualCondaRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupVirtualCondaRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualCondaRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualCondaRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualCondaRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualCondaRepositoryResultOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualCondaRepositoryResult) []string { return v.Repositories }).(pulumi.StringArrayOutput)
}

func (o LookupVirtualCondaRepositoryResultOutput) RetrievalCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualCondaRepositoryResult) *int { return v.RetrievalCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualCondaRepositoryResultOutput{})
}
