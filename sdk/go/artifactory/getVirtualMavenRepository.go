// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a virtual Maven repository.
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
//			_, err := artifactory.GetVirtualMavenRepository(ctx, &artifactory.GetVirtualMavenRepositoryArgs{
//				Key: "virtual-maven",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetVirtualMavenRepository(ctx *pulumi.Context, args *GetVirtualMavenRepositoryArgs, opts ...pulumi.InvokeOption) (*GetVirtualMavenRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetVirtualMavenRepositoryResult
	err := ctx.Invoke("artifactory:index/getVirtualMavenRepository:getVirtualMavenRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualMavenRepository.
type GetVirtualMavenRepositoryArgs struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool   `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         *string `pulumi:"defaultDeploymentRepo"`
	Description                                   *string `pulumi:"description"`
	ExcludesPattern                               *string `pulumi:"excludesPattern"`
	// (Optional) Forces authentication when fetching from remote repos.
	ForceMavenAuthentication *bool   `pulumi:"forceMavenAuthentication"`
	IncludesPattern          *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key     string  `pulumi:"key"`
	KeyPair *string `pulumi:"keyPair"`
	Notes   *string `pulumi:"notes"`
	// (Optional) One of: `"discardActiveReference", "discardAnyReference", "nothing"`
	PomRepositoryReferencesCleanupPolicy *string  `pulumi:"pomRepositoryReferencesCleanupPolicy"`
	ProjectEnvironments                  []string `pulumi:"projectEnvironments"`
	ProjectKey                           *string  `pulumi:"projectKey"`
	RepoLayoutRef                        *string  `pulumi:"repoLayoutRef"`
	Repositories                         []string `pulumi:"repositories"`
}

// A collection of values returned by getVirtualMavenRepository.
type GetVirtualMavenRepositoryResult struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool   `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         *string `pulumi:"defaultDeploymentRepo"`
	Description                                   *string `pulumi:"description"`
	ExcludesPattern                               *string `pulumi:"excludesPattern"`
	// (Optional) Forces authentication when fetching from remote repos.
	ForceMavenAuthentication bool `pulumi:"forceMavenAuthentication"`
	// The provider-assigned unique ID for this managed resource.
	Id              string  `pulumi:"id"`
	IncludesPattern *string `pulumi:"includesPattern"`
	Key             string  `pulumi:"key"`
	KeyPair         *string `pulumi:"keyPair"`
	Notes           *string `pulumi:"notes"`
	PackageType     string  `pulumi:"packageType"`
	// (Optional) One of: `"discardActiveReference", "discardAnyReference", "nothing"`
	PomRepositoryReferencesCleanupPolicy string   `pulumi:"pomRepositoryReferencesCleanupPolicy"`
	ProjectEnvironments                  []string `pulumi:"projectEnvironments"`
	ProjectKey                           *string  `pulumi:"projectKey"`
	RepoLayoutRef                        *string  `pulumi:"repoLayoutRef"`
	Repositories                         []string `pulumi:"repositories"`
}

func GetVirtualMavenRepositoryOutput(ctx *pulumi.Context, args GetVirtualMavenRepositoryOutputArgs, opts ...pulumi.InvokeOption) GetVirtualMavenRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVirtualMavenRepositoryResult, error) {
			args := v.(GetVirtualMavenRepositoryArgs)
			r, err := GetVirtualMavenRepository(ctx, &args, opts...)
			var s GetVirtualMavenRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVirtualMavenRepositoryResultOutput)
}

// A collection of arguments for invoking getVirtualMavenRepository.
type GetVirtualMavenRepositoryOutputArgs struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput   `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         pulumi.StringPtrInput `pulumi:"defaultDeploymentRepo"`
	Description                                   pulumi.StringPtrInput `pulumi:"description"`
	ExcludesPattern                               pulumi.StringPtrInput `pulumi:"excludesPattern"`
	// (Optional) Forces authentication when fetching from remote repos.
	ForceMavenAuthentication pulumi.BoolPtrInput   `pulumi:"forceMavenAuthentication"`
	IncludesPattern          pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key     pulumi.StringInput    `pulumi:"key"`
	KeyPair pulumi.StringPtrInput `pulumi:"keyPair"`
	Notes   pulumi.StringPtrInput `pulumi:"notes"`
	// (Optional) One of: `"discardActiveReference", "discardAnyReference", "nothing"`
	PomRepositoryReferencesCleanupPolicy pulumi.StringPtrInput   `pulumi:"pomRepositoryReferencesCleanupPolicy"`
	ProjectEnvironments                  pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey                           pulumi.StringPtrInput   `pulumi:"projectKey"`
	RepoLayoutRef                        pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	Repositories                         pulumi.StringArrayInput `pulumi:"repositories"`
}

func (GetVirtualMavenRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualMavenRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getVirtualMavenRepository.
type GetVirtualMavenRepositoryResultOutput struct{ *pulumi.OutputState }

func (GetVirtualMavenRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualMavenRepositoryResult)(nil)).Elem()
}

func (o GetVirtualMavenRepositoryResultOutput) ToGetVirtualMavenRepositoryResultOutput() GetVirtualMavenRepositoryResultOutput {
	return o
}

func (o GetVirtualMavenRepositoryResultOutput) ToGetVirtualMavenRepositoryResultOutputWithContext(ctx context.Context) GetVirtualMavenRepositoryResultOutput {
	return o
}

func (o GetVirtualMavenRepositoryResultOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetVirtualMavenRepositoryResult) *bool { return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts }).(pulumi.BoolPtrOutput)
}

func (o GetVirtualMavenRepositoryResultOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualMavenRepositoryResult) *string { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

func (o GetVirtualMavenRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualMavenRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GetVirtualMavenRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualMavenRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// (Optional) Forces authentication when fetching from remote repos.
func (o GetVirtualMavenRepositoryResultOutput) ForceMavenAuthentication() pulumi.BoolOutput {
	return o.ApplyT(func(v GetVirtualMavenRepositoryResult) bool { return v.ForceMavenAuthentication }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetVirtualMavenRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVirtualMavenRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetVirtualMavenRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualMavenRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o GetVirtualMavenRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GetVirtualMavenRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o GetVirtualMavenRepositoryResultOutput) KeyPair() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualMavenRepositoryResult) *string { return v.KeyPair }).(pulumi.StringPtrOutput)
}

func (o GetVirtualMavenRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualMavenRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o GetVirtualMavenRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v GetVirtualMavenRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

// (Optional) One of: `"discardActiveReference", "discardAnyReference", "nothing"`
func (o GetVirtualMavenRepositoryResultOutput) PomRepositoryReferencesCleanupPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v GetVirtualMavenRepositoryResult) string { return v.PomRepositoryReferencesCleanupPolicy }).(pulumi.StringOutput)
}

func (o GetVirtualMavenRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVirtualMavenRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o GetVirtualMavenRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualMavenRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o GetVirtualMavenRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualMavenRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o GetVirtualMavenRepositoryResultOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVirtualMavenRepositoryResult) []string { return v.Repositories }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVirtualMavenRepositoryResultOutput{})
}
