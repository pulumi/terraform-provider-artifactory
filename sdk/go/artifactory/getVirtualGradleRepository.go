// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v6/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a virtual Gradle repository.
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
//			_, err := artifactory.LookupVirtualGradleRepository(ctx, &artifactory.LookupVirtualGradleRepositoryArgs{
//				Key: "virtual-gradle",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupVirtualGradleRepository(ctx *pulumi.Context, args *LookupVirtualGradleRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupVirtualGradleRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualGradleRepositoryResult
	err := ctx.Invoke("artifactory:index/getVirtualGradleRepository:getVirtualGradleRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualGradleRepository.
type LookupVirtualGradleRepositoryArgs struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool   `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         *string `pulumi:"defaultDeploymentRepo"`
	Description                                   *string `pulumi:"description"`
	ExcludesPattern                               *string `pulumi:"excludesPattern"`
	ForceMavenAuthentication                      *bool   `pulumi:"forceMavenAuthentication"`
	IncludesPattern                               *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// (Optional) The keypair used to sign artifacts.
	KeyPair *string `pulumi:"keyPair"`
	Notes   *string `pulumi:"notes"`
	// (Optional)
	// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
	// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
	// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
	PomRepositoryReferencesCleanupPolicy *string  `pulumi:"pomRepositoryReferencesCleanupPolicy"`
	ProjectEnvironments                  []string `pulumi:"projectEnvironments"`
	ProjectKey                           *string  `pulumi:"projectKey"`
	RepoLayoutRef                        *string  `pulumi:"repoLayoutRef"`
	Repositories                         []string `pulumi:"repositories"`
}

// A collection of values returned by getVirtualGradleRepository.
type LookupVirtualGradleRepositoryResult struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool   `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         *string `pulumi:"defaultDeploymentRepo"`
	Description                                   *string `pulumi:"description"`
	ExcludesPattern                               *string `pulumi:"excludesPattern"`
	ForceMavenAuthentication                      bool    `pulumi:"forceMavenAuthentication"`
	// The provider-assigned unique ID for this managed resource.
	Id              string  `pulumi:"id"`
	IncludesPattern *string `pulumi:"includesPattern"`
	Key             string  `pulumi:"key"`
	// (Optional) The keypair used to sign artifacts.
	KeyPair     *string `pulumi:"keyPair"`
	Notes       *string `pulumi:"notes"`
	PackageType string  `pulumi:"packageType"`
	// (Optional)
	// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
	// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
	// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
	PomRepositoryReferencesCleanupPolicy string   `pulumi:"pomRepositoryReferencesCleanupPolicy"`
	ProjectEnvironments                  []string `pulumi:"projectEnvironments"`
	ProjectKey                           *string  `pulumi:"projectKey"`
	RepoLayoutRef                        *string  `pulumi:"repoLayoutRef"`
	Repositories                         []string `pulumi:"repositories"`
}

func LookupVirtualGradleRepositoryOutput(ctx *pulumi.Context, args LookupVirtualGradleRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualGradleRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualGradleRepositoryResult, error) {
			args := v.(LookupVirtualGradleRepositoryArgs)
			r, err := LookupVirtualGradleRepository(ctx, &args, opts...)
			var s LookupVirtualGradleRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualGradleRepositoryResultOutput)
}

// A collection of arguments for invoking getVirtualGradleRepository.
type LookupVirtualGradleRepositoryOutputArgs struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput   `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         pulumi.StringPtrInput `pulumi:"defaultDeploymentRepo"`
	Description                                   pulumi.StringPtrInput `pulumi:"description"`
	ExcludesPattern                               pulumi.StringPtrInput `pulumi:"excludesPattern"`
	ForceMavenAuthentication                      pulumi.BoolPtrInput   `pulumi:"forceMavenAuthentication"`
	IncludesPattern                               pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key pulumi.StringInput `pulumi:"key"`
	// (Optional) The keypair used to sign artifacts.
	KeyPair pulumi.StringPtrInput `pulumi:"keyPair"`
	Notes   pulumi.StringPtrInput `pulumi:"notes"`
	// (Optional)
	// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
	// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
	// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
	PomRepositoryReferencesCleanupPolicy pulumi.StringPtrInput   `pulumi:"pomRepositoryReferencesCleanupPolicy"`
	ProjectEnvironments                  pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey                           pulumi.StringPtrInput   `pulumi:"projectKey"`
	RepoLayoutRef                        pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	Repositories                         pulumi.StringArrayInput `pulumi:"repositories"`
}

func (LookupVirtualGradleRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualGradleRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getVirtualGradleRepository.
type LookupVirtualGradleRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualGradleRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualGradleRepositoryResult)(nil)).Elem()
}

func (o LookupVirtualGradleRepositoryResultOutput) ToLookupVirtualGradleRepositoryResultOutput() LookupVirtualGradleRepositoryResultOutput {
	return o
}

func (o LookupVirtualGradleRepositoryResultOutput) ToLookupVirtualGradleRepositoryResultOutputWithContext(ctx context.Context) LookupVirtualGradleRepositoryResultOutput {
	return o
}

func (o LookupVirtualGradleRepositoryResultOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualGradleRepositoryResult) *bool {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualGradleRepositoryResultOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualGradleRepositoryResult) *string { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualGradleRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualGradleRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualGradleRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualGradleRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualGradleRepositoryResultOutput) ForceMavenAuthentication() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVirtualGradleRepositoryResult) bool { return v.ForceMavenAuthentication }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVirtualGradleRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualGradleRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualGradleRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualGradleRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualGradleRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualGradleRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

// (Optional) The keypair used to sign artifacts.
func (o LookupVirtualGradleRepositoryResultOutput) KeyPair() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualGradleRepositoryResult) *string { return v.KeyPair }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualGradleRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualGradleRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualGradleRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualGradleRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

// (Optional)
// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
func (o LookupVirtualGradleRepositoryResultOutput) PomRepositoryReferencesCleanupPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualGradleRepositoryResult) string { return v.PomRepositoryReferencesCleanupPolicy }).(pulumi.StringOutput)
}

func (o LookupVirtualGradleRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualGradleRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupVirtualGradleRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualGradleRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualGradleRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualGradleRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualGradleRepositoryResultOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualGradleRepositoryResult) []string { return v.Repositories }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualGradleRepositoryResultOutput{})
}
