// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNugetRepository(ctx *pulumi.Context, args *LookupVirtualNugetRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNugetRepositoryResult, error) {
	var rv LookupVirtualNugetRepositoryResult
	err := ctx.Invoke("artifactory:index/getVirtualNugetRepository:getVirtualNugetRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualNugetRepository.
type LookupVirtualNugetRepositoryArgs struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool    `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         *string  `pulumi:"defaultDeploymentRepo"`
	Description                                   *string  `pulumi:"description"`
	ExcludesPattern                               *string  `pulumi:"excludesPattern"`
	ForceNugetAuthentication                      *bool    `pulumi:"forceNugetAuthentication"`
	IncludesPattern                               *string  `pulumi:"includesPattern"`
	Key                                           string   `pulumi:"key"`
	Notes                                         *string  `pulumi:"notes"`
	ProjectEnvironments                           []string `pulumi:"projectEnvironments"`
	ProjectKey                                    *string  `pulumi:"projectKey"`
	RepoLayoutRef                                 *string  `pulumi:"repoLayoutRef"`
	Repositories                                  []string `pulumi:"repositories"`
}

// A collection of values returned by getVirtualNugetRepository.
type LookupVirtualNugetRepositoryResult struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool   `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         *string `pulumi:"defaultDeploymentRepo"`
	Description                                   *string `pulumi:"description"`
	ExcludesPattern                               *string `pulumi:"excludesPattern"`
	ForceNugetAuthentication                      *bool   `pulumi:"forceNugetAuthentication"`
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

func LookupVirtualNugetRepositoryOutput(ctx *pulumi.Context, args LookupVirtualNugetRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualNugetRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualNugetRepositoryResult, error) {
			args := v.(LookupVirtualNugetRepositoryArgs)
			r, err := LookupVirtualNugetRepository(ctx, &args, opts...)
			var s LookupVirtualNugetRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualNugetRepositoryResultOutput)
}

// A collection of arguments for invoking getVirtualNugetRepository.
type LookupVirtualNugetRepositoryOutputArgs struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput     `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         pulumi.StringPtrInput   `pulumi:"defaultDeploymentRepo"`
	Description                                   pulumi.StringPtrInput   `pulumi:"description"`
	ExcludesPattern                               pulumi.StringPtrInput   `pulumi:"excludesPattern"`
	ForceNugetAuthentication                      pulumi.BoolPtrInput     `pulumi:"forceNugetAuthentication"`
	IncludesPattern                               pulumi.StringPtrInput   `pulumi:"includesPattern"`
	Key                                           pulumi.StringInput      `pulumi:"key"`
	Notes                                         pulumi.StringPtrInput   `pulumi:"notes"`
	ProjectEnvironments                           pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey                                    pulumi.StringPtrInput   `pulumi:"projectKey"`
	RepoLayoutRef                                 pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	Repositories                                  pulumi.StringArrayInput `pulumi:"repositories"`
}

func (LookupVirtualNugetRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNugetRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getVirtualNugetRepository.
type LookupVirtualNugetRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualNugetRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNugetRepositoryResult)(nil)).Elem()
}

func (o LookupVirtualNugetRepositoryResultOutput) ToLookupVirtualNugetRepositoryResultOutput() LookupVirtualNugetRepositoryResultOutput {
	return o
}

func (o LookupVirtualNugetRepositoryResultOutput) ToLookupVirtualNugetRepositoryResultOutputWithContext(ctx context.Context) LookupVirtualNugetRepositoryResultOutput {
	return o
}

func (o LookupVirtualNugetRepositoryResultOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNugetRepositoryResult) *bool {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualNugetRepositoryResultOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNugetRepositoryResult) *string { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNugetRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNugetRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNugetRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNugetRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNugetRepositoryResultOutput) ForceNugetAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNugetRepositoryResult) *bool { return v.ForceNugetAuthentication }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVirtualNugetRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNugetRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualNugetRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNugetRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNugetRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNugetRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupVirtualNugetRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNugetRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNugetRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNugetRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupVirtualNugetRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualNugetRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupVirtualNugetRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNugetRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNugetRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNugetRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNugetRepositoryResultOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualNugetRepositoryResult) []string { return v.Repositories }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualNugetRepositoryResultOutput{})
}
