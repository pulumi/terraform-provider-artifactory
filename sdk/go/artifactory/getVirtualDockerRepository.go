// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualDockerRepository(ctx *pulumi.Context, args *LookupVirtualDockerRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupVirtualDockerRepositoryResult, error) {
	var rv LookupVirtualDockerRepositoryResult
	err := ctx.Invoke("artifactory:index/getVirtualDockerRepository:getVirtualDockerRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualDockerRepository.
type LookupVirtualDockerRepositoryArgs struct {
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
	ResolveDockerTagsByTimestamp                  *bool    `pulumi:"resolveDockerTagsByTimestamp"`
}

// A collection of values returned by getVirtualDockerRepository.
type LookupVirtualDockerRepositoryResult struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool   `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         *string `pulumi:"defaultDeploymentRepo"`
	Description                                   *string `pulumi:"description"`
	ExcludesPattern                               *string `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id                           string   `pulumi:"id"`
	IncludesPattern              *string  `pulumi:"includesPattern"`
	Key                          string   `pulumi:"key"`
	Notes                        *string  `pulumi:"notes"`
	PackageType                  string   `pulumi:"packageType"`
	ProjectEnvironments          []string `pulumi:"projectEnvironments"`
	ProjectKey                   *string  `pulumi:"projectKey"`
	RepoLayoutRef                *string  `pulumi:"repoLayoutRef"`
	Repositories                 []string `pulumi:"repositories"`
	ResolveDockerTagsByTimestamp *bool    `pulumi:"resolveDockerTagsByTimestamp"`
}

func LookupVirtualDockerRepositoryOutput(ctx *pulumi.Context, args LookupVirtualDockerRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualDockerRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualDockerRepositoryResult, error) {
			args := v.(LookupVirtualDockerRepositoryArgs)
			r, err := LookupVirtualDockerRepository(ctx, &args, opts...)
			var s LookupVirtualDockerRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualDockerRepositoryResultOutput)
}

// A collection of arguments for invoking getVirtualDockerRepository.
type LookupVirtualDockerRepositoryOutputArgs struct {
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
	ResolveDockerTagsByTimestamp                  pulumi.BoolPtrInput     `pulumi:"resolveDockerTagsByTimestamp"`
}

func (LookupVirtualDockerRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualDockerRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getVirtualDockerRepository.
type LookupVirtualDockerRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualDockerRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualDockerRepositoryResult)(nil)).Elem()
}

func (o LookupVirtualDockerRepositoryResultOutput) ToLookupVirtualDockerRepositoryResultOutput() LookupVirtualDockerRepositoryResultOutput {
	return o
}

func (o LookupVirtualDockerRepositoryResultOutput) ToLookupVirtualDockerRepositoryResultOutputWithContext(ctx context.Context) LookupVirtualDockerRepositoryResultOutput {
	return o
}

func (o LookupVirtualDockerRepositoryResultOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualDockerRepositoryResult) *bool {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualDockerRepositoryResultOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualDockerRepositoryResult) *string { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualDockerRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualDockerRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualDockerRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualDockerRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVirtualDockerRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualDockerRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualDockerRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualDockerRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualDockerRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualDockerRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupVirtualDockerRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualDockerRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualDockerRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualDockerRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupVirtualDockerRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualDockerRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupVirtualDockerRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualDockerRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualDockerRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualDockerRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualDockerRepositoryResultOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualDockerRepositoryResult) []string { return v.Repositories }).(pulumi.StringArrayOutput)
}

func (o LookupVirtualDockerRepositoryResultOutput) ResolveDockerTagsByTimestamp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualDockerRepositoryResult) *bool { return v.ResolveDockerTagsByTimestamp }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualDockerRepositoryResultOutput{})
}
