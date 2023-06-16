// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNpmRepository(ctx *pulumi.Context, args *LookupVirtualNpmRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNpmRepositoryResult, error) {
	var rv LookupVirtualNpmRepositoryResult
	err := ctx.Invoke("artifactory:index/getVirtualNpmRepository:getVirtualNpmRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualNpmRepository.
type LookupVirtualNpmRepositoryArgs struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool    `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         *string  `pulumi:"defaultDeploymentRepo"`
	Description                                   *string  `pulumi:"description"`
	ExcludesPattern                               *string  `pulumi:"excludesPattern"`
	ExternalDependenciesEnabled                   *bool    `pulumi:"externalDependenciesEnabled"`
	ExternalDependenciesPatterns                  []string `pulumi:"externalDependenciesPatterns"`
	ExternalDependenciesRemoteRepo                *string  `pulumi:"externalDependenciesRemoteRepo"`
	IncludesPattern                               *string  `pulumi:"includesPattern"`
	Key                                           string   `pulumi:"key"`
	Notes                                         *string  `pulumi:"notes"`
	ProjectEnvironments                           []string `pulumi:"projectEnvironments"`
	ProjectKey                                    *string  `pulumi:"projectKey"`
	RepoLayoutRef                                 *string  `pulumi:"repoLayoutRef"`
	Repositories                                  []string `pulumi:"repositories"`
	RetrievalCachePeriodSeconds                   *int     `pulumi:"retrievalCachePeriodSeconds"`
}

// A collection of values returned by getVirtualNpmRepository.
type LookupVirtualNpmRepositoryResult struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool    `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         *string  `pulumi:"defaultDeploymentRepo"`
	Description                                   *string  `pulumi:"description"`
	ExcludesPattern                               *string  `pulumi:"excludesPattern"`
	ExternalDependenciesEnabled                   *bool    `pulumi:"externalDependenciesEnabled"`
	ExternalDependenciesPatterns                  []string `pulumi:"externalDependenciesPatterns"`
	ExternalDependenciesRemoteRepo                *string  `pulumi:"externalDependenciesRemoteRepo"`
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

func LookupVirtualNpmRepositoryOutput(ctx *pulumi.Context, args LookupVirtualNpmRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualNpmRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualNpmRepositoryResult, error) {
			args := v.(LookupVirtualNpmRepositoryArgs)
			r, err := LookupVirtualNpmRepository(ctx, &args, opts...)
			var s LookupVirtualNpmRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualNpmRepositoryResultOutput)
}

// A collection of arguments for invoking getVirtualNpmRepository.
type LookupVirtualNpmRepositoryOutputArgs struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput     `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         pulumi.StringPtrInput   `pulumi:"defaultDeploymentRepo"`
	Description                                   pulumi.StringPtrInput   `pulumi:"description"`
	ExcludesPattern                               pulumi.StringPtrInput   `pulumi:"excludesPattern"`
	ExternalDependenciesEnabled                   pulumi.BoolPtrInput     `pulumi:"externalDependenciesEnabled"`
	ExternalDependenciesPatterns                  pulumi.StringArrayInput `pulumi:"externalDependenciesPatterns"`
	ExternalDependenciesRemoteRepo                pulumi.StringPtrInput   `pulumi:"externalDependenciesRemoteRepo"`
	IncludesPattern                               pulumi.StringPtrInput   `pulumi:"includesPattern"`
	Key                                           pulumi.StringInput      `pulumi:"key"`
	Notes                                         pulumi.StringPtrInput   `pulumi:"notes"`
	ProjectEnvironments                           pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey                                    pulumi.StringPtrInput   `pulumi:"projectKey"`
	RepoLayoutRef                                 pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	Repositories                                  pulumi.StringArrayInput `pulumi:"repositories"`
	RetrievalCachePeriodSeconds                   pulumi.IntPtrInput      `pulumi:"retrievalCachePeriodSeconds"`
}

func (LookupVirtualNpmRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNpmRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getVirtualNpmRepository.
type LookupVirtualNpmRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualNpmRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNpmRepositoryResult)(nil)).Elem()
}

func (o LookupVirtualNpmRepositoryResultOutput) ToLookupVirtualNpmRepositoryResultOutput() LookupVirtualNpmRepositoryResultOutput {
	return o
}

func (o LookupVirtualNpmRepositoryResultOutput) ToLookupVirtualNpmRepositoryResultOutputWithContext(ctx context.Context) LookupVirtualNpmRepositoryResultOutput {
	return o
}

func (o LookupVirtualNpmRepositoryResultOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNpmRepositoryResult) *bool { return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualNpmRepositoryResultOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNpmRepositoryResult) *string { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNpmRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNpmRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNpmRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNpmRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNpmRepositoryResultOutput) ExternalDependenciesEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNpmRepositoryResult) *bool { return v.ExternalDependenciesEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualNpmRepositoryResultOutput) ExternalDependenciesPatterns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualNpmRepositoryResult) []string { return v.ExternalDependenciesPatterns }).(pulumi.StringArrayOutput)
}

func (o LookupVirtualNpmRepositoryResultOutput) ExternalDependenciesRemoteRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNpmRepositoryResult) *string { return v.ExternalDependenciesRemoteRepo }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVirtualNpmRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNpmRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualNpmRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNpmRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNpmRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNpmRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupVirtualNpmRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNpmRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNpmRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNpmRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupVirtualNpmRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualNpmRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupVirtualNpmRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNpmRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNpmRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNpmRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNpmRepositoryResultOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualNpmRepositoryResult) []string { return v.Repositories }).(pulumi.StringArrayOutput)
}

func (o LookupVirtualNpmRepositoryResultOutput) RetrievalCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualNpmRepositoryResult) *int { return v.RetrievalCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualNpmRepositoryResultOutput{})
}
