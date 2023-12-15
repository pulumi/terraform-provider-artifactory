// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v6/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a local Ivy repository.
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
//			_, err := artifactory.LookupLocalIvyRepository(ctx, &artifactory.LookupLocalIvyRepositoryArgs{
//				Key: "local-test-ivy-repo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupLocalIvyRepository(ctx *pulumi.Context, args *LookupLocalIvyRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupLocalIvyRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLocalIvyRepositoryResult
	err := ctx.Invoke("artifactory:index/getLocalIvyRepository:getLocalIvyRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocalIvyRepository.
type LookupLocalIvyRepositoryArgs struct {
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool `pulumi:"blackedOut"`
	CdnRedirect            *bool `pulumi:"cdnRedirect"`
	// Checksum policy determines how Artifactory behaves when a client checksum for a deployed
	// resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
	// `client-checksums` and `generated-checksums`. For more details, please refer
	// to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
	// .
	ChecksumPolicyType *string `pulumi:"checksumPolicyType"`
	Description        *string `pulumi:"description"`
	DownloadDirect     *bool   `pulumi:"downloadDirect"`
	ExcludesPattern    *string `pulumi:"excludesPattern"`
	// If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`
	// .
	HandleReleases *bool `pulumi:"handleReleases"`
	// If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default
	// is `true`.
	HandleSnapshots *bool   `pulumi:"handleSnapshots"`
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// The maximum number of unique snapshots of a single artifact to store. Once the number of
	// snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and
	// unique snapshots are not cleaned up.
	MaxUniqueSnapshots  *int     `pulumi:"maxUniqueSnapshots"`
	Notes               *string  `pulumi:"notes"`
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	PropertySets        []string `pulumi:"propertySets"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	// Specifies the naming convention for Maven SNAPSHOT versions. The options are
	// ---
	SnapshotVersionBehavior *string `pulumi:"snapshotVersionBehavior"`
	// By default, Artifactory keeps your repositories healthy by refusing POMs with
	// incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed
	// path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting the
	// Suppress POM Consistency Checks checkbox. True by default for Gradle repository.
	SuppressPomConsistencyChecks *bool `pulumi:"suppressPomConsistencyChecks"`
	XrayIndex                    *bool `pulumi:"xrayIndex"`
}

// A collection of values returned by getLocalIvyRepository.
type LookupLocalIvyRepositoryResult struct {
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool `pulumi:"blackedOut"`
	CdnRedirect            *bool `pulumi:"cdnRedirect"`
	// Checksum policy determines how Artifactory behaves when a client checksum for a deployed
	// resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
	// `client-checksums` and `generated-checksums`. For more details, please refer
	// to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
	// .
	ChecksumPolicyType *string `pulumi:"checksumPolicyType"`
	Description        *string `pulumi:"description"`
	DownloadDirect     *bool   `pulumi:"downloadDirect"`
	ExcludesPattern    *string `pulumi:"excludesPattern"`
	// If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`
	// .
	HandleReleases *bool `pulumi:"handleReleases"`
	// If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default
	// is `true`.
	HandleSnapshots *bool `pulumi:"handleSnapshots"`
	// The provider-assigned unique ID for this managed resource.
	Id              string  `pulumi:"id"`
	IncludesPattern *string `pulumi:"includesPattern"`
	Key             string  `pulumi:"key"`
	// The maximum number of unique snapshots of a single artifact to store. Once the number of
	// snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and
	// unique snapshots are not cleaned up.
	MaxUniqueSnapshots  *int     `pulumi:"maxUniqueSnapshots"`
	Notes               *string  `pulumi:"notes"`
	PackageType         string   `pulumi:"packageType"`
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	PropertySets        []string `pulumi:"propertySets"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	// Specifies the naming convention for Maven SNAPSHOT versions. The options are
	// ---
	SnapshotVersionBehavior *string `pulumi:"snapshotVersionBehavior"`
	// By default, Artifactory keeps your repositories healthy by refusing POMs with
	// incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed
	// path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting the
	// Suppress POM Consistency Checks checkbox. True by default for Gradle repository.
	SuppressPomConsistencyChecks *bool `pulumi:"suppressPomConsistencyChecks"`
	XrayIndex                    *bool `pulumi:"xrayIndex"`
}

func LookupLocalIvyRepositoryOutput(ctx *pulumi.Context, args LookupLocalIvyRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupLocalIvyRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLocalIvyRepositoryResult, error) {
			args := v.(LookupLocalIvyRepositoryArgs)
			r, err := LookupLocalIvyRepository(ctx, &args, opts...)
			var s LookupLocalIvyRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLocalIvyRepositoryResultOutput)
}

// A collection of arguments for invoking getLocalIvyRepository.
type LookupLocalIvyRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput `pulumi:"blackedOut"`
	CdnRedirect            pulumi.BoolPtrInput `pulumi:"cdnRedirect"`
	// Checksum policy determines how Artifactory behaves when a client checksum for a deployed
	// resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
	// `client-checksums` and `generated-checksums`. For more details, please refer
	// to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
	// .
	ChecksumPolicyType pulumi.StringPtrInput `pulumi:"checksumPolicyType"`
	Description        pulumi.StringPtrInput `pulumi:"description"`
	DownloadDirect     pulumi.BoolPtrInput   `pulumi:"downloadDirect"`
	ExcludesPattern    pulumi.StringPtrInput `pulumi:"excludesPattern"`
	// If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`
	// .
	HandleReleases pulumi.BoolPtrInput `pulumi:"handleReleases"`
	// If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default
	// is `true`.
	HandleSnapshots pulumi.BoolPtrInput   `pulumi:"handleSnapshots"`
	IncludesPattern pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key pulumi.StringInput `pulumi:"key"`
	// The maximum number of unique snapshots of a single artifact to store. Once the number of
	// snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and
	// unique snapshots are not cleaned up.
	MaxUniqueSnapshots  pulumi.IntPtrInput      `pulumi:"maxUniqueSnapshots"`
	Notes               pulumi.StringPtrInput   `pulumi:"notes"`
	PriorityResolution  pulumi.BoolPtrInput     `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput   `pulumi:"projectKey"`
	PropertySets        pulumi.StringArrayInput `pulumi:"propertySets"`
	RepoLayoutRef       pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	// Specifies the naming convention for Maven SNAPSHOT versions. The options are
	// ---
	SnapshotVersionBehavior pulumi.StringPtrInput `pulumi:"snapshotVersionBehavior"`
	// By default, Artifactory keeps your repositories healthy by refusing POMs with
	// incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed
	// path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting the
	// Suppress POM Consistency Checks checkbox. True by default for Gradle repository.
	SuppressPomConsistencyChecks pulumi.BoolPtrInput `pulumi:"suppressPomConsistencyChecks"`
	XrayIndex                    pulumi.BoolPtrInput `pulumi:"xrayIndex"`
}

func (LookupLocalIvyRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalIvyRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getLocalIvyRepository.
type LookupLocalIvyRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupLocalIvyRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalIvyRepositoryResult)(nil)).Elem()
}

func (o LookupLocalIvyRepositoryResultOutput) ToLookupLocalIvyRepositoryResultOutput() LookupLocalIvyRepositoryResultOutput {
	return o
}

func (o LookupLocalIvyRepositoryResultOutput) ToLookupLocalIvyRepositoryResultOutputWithContext(ctx context.Context) LookupLocalIvyRepositoryResultOutput {
	return o
}

func (o LookupLocalIvyRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalIvyRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalIvyRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

// Checksum policy determines how Artifactory behaves when a client checksum for a deployed
// resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
// `client-checksums` and `generated-checksums`. For more details, please refer
// to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
// .
func (o LookupLocalIvyRepositoryResultOutput) ChecksumPolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) *string { return v.ChecksumPolicyType }).(pulumi.StringPtrOutput)
}

func (o LookupLocalIvyRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupLocalIvyRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalIvyRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`
// .
func (o LookupLocalIvyRepositoryResultOutput) HandleReleases() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) *bool { return v.HandleReleases }).(pulumi.BoolPtrOutput)
}

// If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default
// is `true`.
func (o LookupLocalIvyRepositoryResultOutput) HandleSnapshots() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) *bool { return v.HandleSnapshots }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLocalIvyRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLocalIvyRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupLocalIvyRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

// The maximum number of unique snapshots of a single artifact to store. Once the number of
// snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and
// unique snapshots are not cleaned up.
func (o LookupLocalIvyRepositoryResultOutput) MaxUniqueSnapshots() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) *int { return v.MaxUniqueSnapshots }).(pulumi.IntPtrOutput)
}

func (o LookupLocalIvyRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupLocalIvyRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupLocalIvyRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalIvyRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupLocalIvyRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupLocalIvyRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupLocalIvyRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Specifies the naming convention for Maven SNAPSHOT versions. The options are
// ---
func (o LookupLocalIvyRepositoryResultOutput) SnapshotVersionBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) *string { return v.SnapshotVersionBehavior }).(pulumi.StringPtrOutput)
}

// By default, Artifactory keeps your repositories healthy by refusing POMs with
// incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed
// path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting the
// Suppress POM Consistency Checks checkbox. True by default for Gradle repository.
func (o LookupLocalIvyRepositoryResultOutput) SuppressPomConsistencyChecks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) *bool { return v.SuppressPomConsistencyChecks }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalIvyRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocalIvyRepositoryResultOutput{})
}
