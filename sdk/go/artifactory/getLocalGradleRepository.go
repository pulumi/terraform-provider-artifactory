// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v7/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a local Gradle repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v7/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.LookupLocalGradleRepository(ctx, &artifactory.LookupLocalGradleRepositoryArgs{
//				Key: "local-test-gradle-repo-basic",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// * ## `snapshotVersionBehavior` - Specifies the naming convention for Maven SNAPSHOT versions. The options are
//   - `unique`: Version number is based on a time-stamp (default)
//   - `non-unique`: Version number uses a self-overriding naming pattern of artifactId-version-SNAPSHOT.type
//   - `deployer`: Respects the settings in the Maven client that is deploying the artifact.
//   - `maxUniqueSnapshots` - The maximum number of unique snapshots of a single artifact to store. Once the
//     number of snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no
//     limit, and unique snapshots are not cleaned up.
//   - `handleReleases` - If set, Artifactory allows you to deploy release artifacts into this repository.
//     Default is `true`.
//   - `handleSnapshots` - If set, Artifactory allows you to deploy snapshot artifacts into this repository.
//     Default is `true`.
//   - `suppressPomConsistencyChecks` - By default, Artifactory keeps your repositories healthy by refusing
//     POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match
//     the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by
//     setting the Suppress POM Consistency Checks checkbox. True by default for Gradle repository.
func LookupLocalGradleRepository(ctx *pulumi.Context, args *LookupLocalGradleRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupLocalGradleRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLocalGradleRepositoryResult
	err := ctx.Invoke("artifactory:index/getLocalGradleRepository:getLocalGradleRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocalGradleRepository.
type LookupLocalGradleRepositoryArgs struct {
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool `pulumi:"blackedOut"`
	CdnRedirect            *bool `pulumi:"cdnRedirect"`
	// Checksum policy determines how Artifactory behaves when a client checksum for a
	// deployed resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
	// `client-checksums` and `generated-checksums`. For more details, please refer
	// to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
	// .
	ChecksumPolicyType *string `pulumi:"checksumPolicyType"`
	Description        *string `pulumi:"description"`
	DownloadDirect     *bool   `pulumi:"downloadDirect"`
	ExcludesPattern    *string `pulumi:"excludesPattern"`
	HandleReleases     *bool   `pulumi:"handleReleases"`
	HandleSnapshots    *bool   `pulumi:"handleSnapshots"`
	IncludesPattern    *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                          string   `pulumi:"key"`
	MaxUniqueSnapshots           *int     `pulumi:"maxUniqueSnapshots"`
	Notes                        *string  `pulumi:"notes"`
	PriorityResolution           *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments          []string `pulumi:"projectEnvironments"`
	ProjectKey                   *string  `pulumi:"projectKey"`
	PropertySets                 []string `pulumi:"propertySets"`
	RepoLayoutRef                *string  `pulumi:"repoLayoutRef"`
	SnapshotVersionBehavior      *string  `pulumi:"snapshotVersionBehavior"`
	SuppressPomConsistencyChecks *bool    `pulumi:"suppressPomConsistencyChecks"`
	XrayIndex                    *bool    `pulumi:"xrayIndex"`
}

// A collection of values returned by getLocalGradleRepository.
type LookupLocalGradleRepositoryResult struct {
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool `pulumi:"blackedOut"`
	CdnRedirect            *bool `pulumi:"cdnRedirect"`
	// Checksum policy determines how Artifactory behaves when a client checksum for a
	// deployed resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
	// `client-checksums` and `generated-checksums`. For more details, please refer
	// to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
	// .
	ChecksumPolicyType *string `pulumi:"checksumPolicyType"`
	Description        *string `pulumi:"description"`
	DownloadDirect     *bool   `pulumi:"downloadDirect"`
	ExcludesPattern    *string `pulumi:"excludesPattern"`
	HandleReleases     *bool   `pulumi:"handleReleases"`
	HandleSnapshots    *bool   `pulumi:"handleSnapshots"`
	// The provider-assigned unique ID for this managed resource.
	Id                           string   `pulumi:"id"`
	IncludesPattern              *string  `pulumi:"includesPattern"`
	Key                          string   `pulumi:"key"`
	MaxUniqueSnapshots           *int     `pulumi:"maxUniqueSnapshots"`
	Notes                        *string  `pulumi:"notes"`
	PackageType                  string   `pulumi:"packageType"`
	PriorityResolution           *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments          []string `pulumi:"projectEnvironments"`
	ProjectKey                   *string  `pulumi:"projectKey"`
	PropertySets                 []string `pulumi:"propertySets"`
	RepoLayoutRef                *string  `pulumi:"repoLayoutRef"`
	SnapshotVersionBehavior      *string  `pulumi:"snapshotVersionBehavior"`
	SuppressPomConsistencyChecks *bool    `pulumi:"suppressPomConsistencyChecks"`
	XrayIndex                    *bool    `pulumi:"xrayIndex"`
}

func LookupLocalGradleRepositoryOutput(ctx *pulumi.Context, args LookupLocalGradleRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupLocalGradleRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLocalGradleRepositoryResult, error) {
			args := v.(LookupLocalGradleRepositoryArgs)
			r, err := LookupLocalGradleRepository(ctx, &args, opts...)
			var s LookupLocalGradleRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLocalGradleRepositoryResultOutput)
}

// A collection of arguments for invoking getLocalGradleRepository.
type LookupLocalGradleRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput `pulumi:"blackedOut"`
	CdnRedirect            pulumi.BoolPtrInput `pulumi:"cdnRedirect"`
	// Checksum policy determines how Artifactory behaves when a client checksum for a
	// deployed resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
	// `client-checksums` and `generated-checksums`. For more details, please refer
	// to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
	// .
	ChecksumPolicyType pulumi.StringPtrInput `pulumi:"checksumPolicyType"`
	Description        pulumi.StringPtrInput `pulumi:"description"`
	DownloadDirect     pulumi.BoolPtrInput   `pulumi:"downloadDirect"`
	ExcludesPattern    pulumi.StringPtrInput `pulumi:"excludesPattern"`
	HandleReleases     pulumi.BoolPtrInput   `pulumi:"handleReleases"`
	HandleSnapshots    pulumi.BoolPtrInput   `pulumi:"handleSnapshots"`
	IncludesPattern    pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                          pulumi.StringInput      `pulumi:"key"`
	MaxUniqueSnapshots           pulumi.IntPtrInput      `pulumi:"maxUniqueSnapshots"`
	Notes                        pulumi.StringPtrInput   `pulumi:"notes"`
	PriorityResolution           pulumi.BoolPtrInput     `pulumi:"priorityResolution"`
	ProjectEnvironments          pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey                   pulumi.StringPtrInput   `pulumi:"projectKey"`
	PropertySets                 pulumi.StringArrayInput `pulumi:"propertySets"`
	RepoLayoutRef                pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	SnapshotVersionBehavior      pulumi.StringPtrInput   `pulumi:"snapshotVersionBehavior"`
	SuppressPomConsistencyChecks pulumi.BoolPtrInput     `pulumi:"suppressPomConsistencyChecks"`
	XrayIndex                    pulumi.BoolPtrInput     `pulumi:"xrayIndex"`
}

func (LookupLocalGradleRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalGradleRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getLocalGradleRepository.
type LookupLocalGradleRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupLocalGradleRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalGradleRepositoryResult)(nil)).Elem()
}

func (o LookupLocalGradleRepositoryResultOutput) ToLookupLocalGradleRepositoryResultOutput() LookupLocalGradleRepositoryResultOutput {
	return o
}

func (o LookupLocalGradleRepositoryResultOutput) ToLookupLocalGradleRepositoryResultOutputWithContext(ctx context.Context) LookupLocalGradleRepositoryResultOutput {
	return o
}

func (o LookupLocalGradleRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalGradleRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalGradleRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalGradleRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalGradleRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalGradleRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

// Checksum policy determines how Artifactory behaves when a client checksum for a
// deployed resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
// `client-checksums` and `generated-checksums`. For more details, please refer
// to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
// .
func (o LookupLocalGradleRepositoryResultOutput) ChecksumPolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalGradleRepositoryResult) *string { return v.ChecksumPolicyType }).(pulumi.StringPtrOutput)
}

func (o LookupLocalGradleRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalGradleRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupLocalGradleRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalGradleRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalGradleRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalGradleRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupLocalGradleRepositoryResultOutput) HandleReleases() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalGradleRepositoryResult) *bool { return v.HandleReleases }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalGradleRepositoryResultOutput) HandleSnapshots() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalGradleRepositoryResult) *bool { return v.HandleSnapshots }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLocalGradleRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalGradleRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLocalGradleRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalGradleRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupLocalGradleRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalGradleRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupLocalGradleRepositoryResultOutput) MaxUniqueSnapshots() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupLocalGradleRepositoryResult) *int { return v.MaxUniqueSnapshots }).(pulumi.IntPtrOutput)
}

func (o LookupLocalGradleRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalGradleRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupLocalGradleRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalGradleRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupLocalGradleRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalGradleRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalGradleRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocalGradleRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupLocalGradleRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalGradleRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupLocalGradleRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocalGradleRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupLocalGradleRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalGradleRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupLocalGradleRepositoryResultOutput) SnapshotVersionBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalGradleRepositoryResult) *string { return v.SnapshotVersionBehavior }).(pulumi.StringPtrOutput)
}

func (o LookupLocalGradleRepositoryResultOutput) SuppressPomConsistencyChecks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalGradleRepositoryResult) *bool { return v.SuppressPomConsistencyChecks }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalGradleRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalGradleRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocalGradleRepositoryResultOutput{})
}
