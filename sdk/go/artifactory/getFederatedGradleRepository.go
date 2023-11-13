// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a federated Gradle repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.LookupFederatedGradleRepository(ctx, &artifactory.LookupFederatedGradleRepositoryArgs{
//				Key: "federated-test-gradle-repo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFederatedGradleRepository(ctx *pulumi.Context, args *LookupFederatedGradleRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupFederatedGradleRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFederatedGradleRepositoryResult
	err := ctx.Invoke("artifactory:index/getFederatedGradleRepository:getFederatedGradleRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedGradleRepository.
type LookupFederatedGradleRepositoryArgs struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	ChecksumPolicyType     *string `pulumi:"checksumPolicyType"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	HandleReleases         *bool   `pulumi:"handleReleases"`
	HandleSnapshots        *bool   `pulumi:"handleSnapshots"`
	IncludesPattern        *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                string `pulumi:"key"`
	MaxUniqueSnapshots *int   `pulumi:"maxUniqueSnapshots"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members                      []GetFederatedGradleRepositoryMember `pulumi:"members"`
	Notes                        *string                              `pulumi:"notes"`
	PriorityResolution           *bool                                `pulumi:"priorityResolution"`
	ProjectEnvironments          []string                             `pulumi:"projectEnvironments"`
	ProjectKey                   *string                              `pulumi:"projectKey"`
	PropertySets                 []string                             `pulumi:"propertySets"`
	RepoLayoutRef                *string                              `pulumi:"repoLayoutRef"`
	SnapshotVersionBehavior      *string                              `pulumi:"snapshotVersionBehavior"`
	SuppressPomConsistencyChecks *bool                                `pulumi:"suppressPomConsistencyChecks"`
	XrayIndex                    *bool                                `pulumi:"xrayIndex"`
}

// A collection of values returned by getFederatedGradleRepository.
type LookupFederatedGradleRepositoryResult struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	ChecksumPolicyType     *string `pulumi:"checksumPolicyType"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	HandleReleases         *bool   `pulumi:"handleReleases"`
	HandleSnapshots        *bool   `pulumi:"handleSnapshots"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string  `pulumi:"id"`
	IncludesPattern    *string `pulumi:"includesPattern"`
	Key                string  `pulumi:"key"`
	MaxUniqueSnapshots *int    `pulumi:"maxUniqueSnapshots"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members                      []GetFederatedGradleRepositoryMember `pulumi:"members"`
	Notes                        *string                              `pulumi:"notes"`
	PackageType                  string                               `pulumi:"packageType"`
	PriorityResolution           *bool                                `pulumi:"priorityResolution"`
	ProjectEnvironments          []string                             `pulumi:"projectEnvironments"`
	ProjectKey                   *string                              `pulumi:"projectKey"`
	PropertySets                 []string                             `pulumi:"propertySets"`
	RepoLayoutRef                *string                              `pulumi:"repoLayoutRef"`
	SnapshotVersionBehavior      *string                              `pulumi:"snapshotVersionBehavior"`
	SuppressPomConsistencyChecks *bool                                `pulumi:"suppressPomConsistencyChecks"`
	XrayIndex                    *bool                                `pulumi:"xrayIndex"`
}

func LookupFederatedGradleRepositoryOutput(ctx *pulumi.Context, args LookupFederatedGradleRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedGradleRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedGradleRepositoryResult, error) {
			args := v.(LookupFederatedGradleRepositoryArgs)
			r, err := LookupFederatedGradleRepository(ctx, &args, opts...)
			var s LookupFederatedGradleRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFederatedGradleRepositoryResultOutput)
}

// A collection of arguments for invoking getFederatedGradleRepository.
type LookupFederatedGradleRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput   `pulumi:"blackedOut"`
	CdnRedirect            pulumi.BoolPtrInput   `pulumi:"cdnRedirect"`
	ChecksumPolicyType     pulumi.StringPtrInput `pulumi:"checksumPolicyType"`
	CleanupOnDelete        pulumi.BoolPtrInput   `pulumi:"cleanupOnDelete"`
	Description            pulumi.StringPtrInput `pulumi:"description"`
	DownloadDirect         pulumi.BoolPtrInput   `pulumi:"downloadDirect"`
	ExcludesPattern        pulumi.StringPtrInput `pulumi:"excludesPattern"`
	HandleReleases         pulumi.BoolPtrInput   `pulumi:"handleReleases"`
	HandleSnapshots        pulumi.BoolPtrInput   `pulumi:"handleSnapshots"`
	IncludesPattern        pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                pulumi.StringInput `pulumi:"key"`
	MaxUniqueSnapshots pulumi.IntPtrInput `pulumi:"maxUniqueSnapshots"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members                      GetFederatedGradleRepositoryMemberArrayInput `pulumi:"members"`
	Notes                        pulumi.StringPtrInput                        `pulumi:"notes"`
	PriorityResolution           pulumi.BoolPtrInput                          `pulumi:"priorityResolution"`
	ProjectEnvironments          pulumi.StringArrayInput                      `pulumi:"projectEnvironments"`
	ProjectKey                   pulumi.StringPtrInput                        `pulumi:"projectKey"`
	PropertySets                 pulumi.StringArrayInput                      `pulumi:"propertySets"`
	RepoLayoutRef                pulumi.StringPtrInput                        `pulumi:"repoLayoutRef"`
	SnapshotVersionBehavior      pulumi.StringPtrInput                        `pulumi:"snapshotVersionBehavior"`
	SuppressPomConsistencyChecks pulumi.BoolPtrInput                          `pulumi:"suppressPomConsistencyChecks"`
	XrayIndex                    pulumi.BoolPtrInput                          `pulumi:"xrayIndex"`
}

func (LookupFederatedGradleRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedGradleRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedGradleRepository.
type LookupFederatedGradleRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedGradleRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedGradleRepositoryResult)(nil)).Elem()
}

func (o LookupFederatedGradleRepositoryResultOutput) ToLookupFederatedGradleRepositoryResultOutput() LookupFederatedGradleRepositoryResultOutput {
	return o
}

func (o LookupFederatedGradleRepositoryResultOutput) ToLookupFederatedGradleRepositoryResultOutputWithContext(ctx context.Context) LookupFederatedGradleRepositoryResultOutput {
	return o
}

func (o LookupFederatedGradleRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedGradleRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedGradleRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedGradleRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedGradleRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedGradleRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedGradleRepositoryResultOutput) ChecksumPolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedGradleRepositoryResult) *string { return v.ChecksumPolicyType }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedGradleRepositoryResultOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedGradleRepositoryResult) *bool { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedGradleRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedGradleRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedGradleRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedGradleRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedGradleRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedGradleRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedGradleRepositoryResultOutput) HandleReleases() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedGradleRepositoryResult) *bool { return v.HandleReleases }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedGradleRepositoryResultOutput) HandleSnapshots() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedGradleRepositoryResult) *bool { return v.HandleSnapshots }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFederatedGradleRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedGradleRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedGradleRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedGradleRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedGradleRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedGradleRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupFederatedGradleRepositoryResultOutput) MaxUniqueSnapshots() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFederatedGradleRepositoryResult) *int { return v.MaxUniqueSnapshots }).(pulumi.IntPtrOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o LookupFederatedGradleRepositoryResultOutput) Members() GetFederatedGradleRepositoryMemberArrayOutput {
	return o.ApplyT(func(v LookupFederatedGradleRepositoryResult) []GetFederatedGradleRepositoryMember { return v.Members }).(GetFederatedGradleRepositoryMemberArrayOutput)
}

func (o LookupFederatedGradleRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedGradleRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedGradleRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedGradleRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupFederatedGradleRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedGradleRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedGradleRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedGradleRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedGradleRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedGradleRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedGradleRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedGradleRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedGradleRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedGradleRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedGradleRepositoryResultOutput) SnapshotVersionBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedGradleRepositoryResult) *string { return v.SnapshotVersionBehavior }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedGradleRepositoryResultOutput) SuppressPomConsistencyChecks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedGradleRepositoryResult) *bool { return v.SuppressPomConsistencyChecks }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedGradleRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedGradleRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedGradleRepositoryResultOutput{})
}
