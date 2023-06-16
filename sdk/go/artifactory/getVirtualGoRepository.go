// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetVirtualGoRepository(ctx *pulumi.Context, args *GetVirtualGoRepositoryArgs, opts ...pulumi.InvokeOption) (*GetVirtualGoRepositoryResult, error) {
	var rv GetVirtualGoRepositoryResult
	err := ctx.Invoke("artifactory:index/getVirtualGoRepository:getVirtualGoRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualGoRepository.
type GetVirtualGoRepositoryArgs struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool    `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         *string  `pulumi:"defaultDeploymentRepo"`
	Description                                   *string  `pulumi:"description"`
	ExcludesPattern                               *string  `pulumi:"excludesPattern"`
	ExternalDependenciesEnabled                   *bool    `pulumi:"externalDependenciesEnabled"`
	ExternalDependenciesPatterns                  []string `pulumi:"externalDependenciesPatterns"`
	IncludesPattern                               *string  `pulumi:"includesPattern"`
	Key                                           string   `pulumi:"key"`
	Notes                                         *string  `pulumi:"notes"`
	ProjectEnvironments                           []string `pulumi:"projectEnvironments"`
	ProjectKey                                    *string  `pulumi:"projectKey"`
	RepoLayoutRef                                 *string  `pulumi:"repoLayoutRef"`
	Repositories                                  []string `pulumi:"repositories"`
}

// A collection of values returned by getVirtualGoRepository.
type GetVirtualGoRepositoryResult struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool    `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         *string  `pulumi:"defaultDeploymentRepo"`
	Description                                   *string  `pulumi:"description"`
	ExcludesPattern                               *string  `pulumi:"excludesPattern"`
	ExternalDependenciesEnabled                   *bool    `pulumi:"externalDependenciesEnabled"`
	ExternalDependenciesPatterns                  []string `pulumi:"externalDependenciesPatterns"`
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

func GetVirtualGoRepositoryOutput(ctx *pulumi.Context, args GetVirtualGoRepositoryOutputArgs, opts ...pulumi.InvokeOption) GetVirtualGoRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVirtualGoRepositoryResult, error) {
			args := v.(GetVirtualGoRepositoryArgs)
			r, err := GetVirtualGoRepository(ctx, &args, opts...)
			var s GetVirtualGoRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVirtualGoRepositoryResultOutput)
}

// A collection of arguments for invoking getVirtualGoRepository.
type GetVirtualGoRepositoryOutputArgs struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput     `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         pulumi.StringPtrInput   `pulumi:"defaultDeploymentRepo"`
	Description                                   pulumi.StringPtrInput   `pulumi:"description"`
	ExcludesPattern                               pulumi.StringPtrInput   `pulumi:"excludesPattern"`
	ExternalDependenciesEnabled                   pulumi.BoolPtrInput     `pulumi:"externalDependenciesEnabled"`
	ExternalDependenciesPatterns                  pulumi.StringArrayInput `pulumi:"externalDependenciesPatterns"`
	IncludesPattern                               pulumi.StringPtrInput   `pulumi:"includesPattern"`
	Key                                           pulumi.StringInput      `pulumi:"key"`
	Notes                                         pulumi.StringPtrInput   `pulumi:"notes"`
	ProjectEnvironments                           pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey                                    pulumi.StringPtrInput   `pulumi:"projectKey"`
	RepoLayoutRef                                 pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	Repositories                                  pulumi.StringArrayInput `pulumi:"repositories"`
}

func (GetVirtualGoRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualGoRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getVirtualGoRepository.
type GetVirtualGoRepositoryResultOutput struct{ *pulumi.OutputState }

func (GetVirtualGoRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualGoRepositoryResult)(nil)).Elem()
}

func (o GetVirtualGoRepositoryResultOutput) ToGetVirtualGoRepositoryResultOutput() GetVirtualGoRepositoryResultOutput {
	return o
}

func (o GetVirtualGoRepositoryResultOutput) ToGetVirtualGoRepositoryResultOutputWithContext(ctx context.Context) GetVirtualGoRepositoryResultOutput {
	return o
}

func (o GetVirtualGoRepositoryResultOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) *bool { return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts }).(pulumi.BoolPtrOutput)
}

func (o GetVirtualGoRepositoryResultOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) *string { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

func (o GetVirtualGoRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GetVirtualGoRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

func (o GetVirtualGoRepositoryResultOutput) ExternalDependenciesEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) *bool { return v.ExternalDependenciesEnabled }).(pulumi.BoolPtrOutput)
}

func (o GetVirtualGoRepositoryResultOutput) ExternalDependenciesPatterns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) []string { return v.ExternalDependenciesPatterns }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetVirtualGoRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetVirtualGoRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o GetVirtualGoRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o GetVirtualGoRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o GetVirtualGoRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o GetVirtualGoRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o GetVirtualGoRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o GetVirtualGoRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o GetVirtualGoRepositoryResultOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) []string { return v.Repositories }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVirtualGoRepositoryResultOutput{})
}
