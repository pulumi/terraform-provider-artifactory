// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v4/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

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
	ArchiveBrowsingEnabled       *bool    `pulumi:"archiveBrowsingEnabled"`
	BlackedOut                   *bool    `pulumi:"blackedOut"`
	CdnRedirect                  *bool    `pulumi:"cdnRedirect"`
	ChecksumPolicyType           *string  `pulumi:"checksumPolicyType"`
	Description                  *string  `pulumi:"description"`
	DownloadDirect               *bool    `pulumi:"downloadDirect"`
	ExcludesPattern              *string  `pulumi:"excludesPattern"`
	HandleReleases               *bool    `pulumi:"handleReleases"`
	HandleSnapshots              *bool    `pulumi:"handleSnapshots"`
	IncludesPattern              *string  `pulumi:"includesPattern"`
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

// A collection of values returned by getLocalIvyRepository.
type LookupLocalIvyRepositoryResult struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	ChecksumPolicyType     *string `pulumi:"checksumPolicyType"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        string  `pulumi:"excludesPattern"`
	HandleReleases         *bool   `pulumi:"handleReleases"`
	HandleSnapshots        *bool   `pulumi:"handleSnapshots"`
	// The provider-assigned unique ID for this managed resource.
	Id                           string   `pulumi:"id"`
	IncludesPattern              string   `pulumi:"includesPattern"`
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
	ArchiveBrowsingEnabled       pulumi.BoolPtrInput     `pulumi:"archiveBrowsingEnabled"`
	BlackedOut                   pulumi.BoolPtrInput     `pulumi:"blackedOut"`
	CdnRedirect                  pulumi.BoolPtrInput     `pulumi:"cdnRedirect"`
	ChecksumPolicyType           pulumi.StringPtrInput   `pulumi:"checksumPolicyType"`
	Description                  pulumi.StringPtrInput   `pulumi:"description"`
	DownloadDirect               pulumi.BoolPtrInput     `pulumi:"downloadDirect"`
	ExcludesPattern              pulumi.StringPtrInput   `pulumi:"excludesPattern"`
	HandleReleases               pulumi.BoolPtrInput     `pulumi:"handleReleases"`
	HandleSnapshots              pulumi.BoolPtrInput     `pulumi:"handleSnapshots"`
	IncludesPattern              pulumi.StringPtrInput   `pulumi:"includesPattern"`
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

func (o LookupLocalIvyRepositoryResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupLocalIvyRepositoryResult] {
	return pulumix.Output[LookupLocalIvyRepositoryResult]{
		OutputState: o.OutputState,
	}
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

func (o LookupLocalIvyRepositoryResultOutput) ChecksumPolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) *string { return v.ChecksumPolicyType }).(pulumi.StringPtrOutput)
}

func (o LookupLocalIvyRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupLocalIvyRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalIvyRepositoryResultOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) string { return v.ExcludesPattern }).(pulumi.StringOutput)
}

func (o LookupLocalIvyRepositoryResultOutput) HandleReleases() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) *bool { return v.HandleReleases }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalIvyRepositoryResultOutput) HandleSnapshots() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) *bool { return v.HandleSnapshots }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLocalIvyRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLocalIvyRepositoryResultOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) string { return v.IncludesPattern }).(pulumi.StringOutput)
}

func (o LookupLocalIvyRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

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

func (o LookupLocalIvyRepositoryResultOutput) SnapshotVersionBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) *string { return v.SnapshotVersionBehavior }).(pulumi.StringPtrOutput)
}

func (o LookupLocalIvyRepositoryResultOutput) SuppressPomConsistencyChecks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) *bool { return v.SuppressPomConsistencyChecks }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalIvyRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalIvyRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocalIvyRepositoryResultOutput{})
}
