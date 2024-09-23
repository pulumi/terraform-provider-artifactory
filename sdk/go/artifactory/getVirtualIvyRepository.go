// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a virtual Ivy repository.
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
//			_, err := artifactory.LookupVirtualIvyRepository(ctx, &artifactory.LookupVirtualIvyRepositoryArgs{
//				Key: "virtual-ivy",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupVirtualIvyRepository(ctx *pulumi.Context, args *LookupVirtualIvyRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupVirtualIvyRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualIvyRepositoryResult
	err := ctx.Invoke("artifactory:index/getVirtualIvyRepository:getVirtualIvyRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualIvyRepository.
type LookupVirtualIvyRepositoryArgs struct {
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

// A collection of values returned by getVirtualIvyRepository.
type LookupVirtualIvyRepositoryResult struct {
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

func LookupVirtualIvyRepositoryOutput(ctx *pulumi.Context, args LookupVirtualIvyRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualIvyRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualIvyRepositoryResultOutput, error) {
			args := v.(LookupVirtualIvyRepositoryArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupVirtualIvyRepositoryResult
			secret, err := ctx.InvokePackageRaw("artifactory:index/getVirtualIvyRepository:getVirtualIvyRepository", args, &rv, "", opts...)
			if err != nil {
				return LookupVirtualIvyRepositoryResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupVirtualIvyRepositoryResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupVirtualIvyRepositoryResultOutput), nil
			}
			return output, nil
		}).(LookupVirtualIvyRepositoryResultOutput)
}

// A collection of arguments for invoking getVirtualIvyRepository.
type LookupVirtualIvyRepositoryOutputArgs struct {
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

func (LookupVirtualIvyRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualIvyRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getVirtualIvyRepository.
type LookupVirtualIvyRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualIvyRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualIvyRepositoryResult)(nil)).Elem()
}

func (o LookupVirtualIvyRepositoryResultOutput) ToLookupVirtualIvyRepositoryResultOutput() LookupVirtualIvyRepositoryResultOutput {
	return o
}

func (o LookupVirtualIvyRepositoryResultOutput) ToLookupVirtualIvyRepositoryResultOutputWithContext(ctx context.Context) LookupVirtualIvyRepositoryResultOutput {
	return o
}

func (o LookupVirtualIvyRepositoryResultOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualIvyRepositoryResult) *bool { return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualIvyRepositoryResultOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualIvyRepositoryResult) *string { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualIvyRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualIvyRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualIvyRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualIvyRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualIvyRepositoryResultOutput) ForceMavenAuthentication() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVirtualIvyRepositoryResult) bool { return v.ForceMavenAuthentication }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVirtualIvyRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualIvyRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualIvyRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualIvyRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualIvyRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualIvyRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

// (Optional) The keypair used to sign artifacts.
func (o LookupVirtualIvyRepositoryResultOutput) KeyPair() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualIvyRepositoryResult) *string { return v.KeyPair }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualIvyRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualIvyRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualIvyRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualIvyRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

// (Optional)
// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
func (o LookupVirtualIvyRepositoryResultOutput) PomRepositoryReferencesCleanupPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualIvyRepositoryResult) string { return v.PomRepositoryReferencesCleanupPolicy }).(pulumi.StringOutput)
}

func (o LookupVirtualIvyRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualIvyRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupVirtualIvyRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualIvyRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualIvyRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualIvyRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualIvyRepositoryResultOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualIvyRepositoryResult) []string { return v.Repositories }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualIvyRepositoryResultOutput{})
}
