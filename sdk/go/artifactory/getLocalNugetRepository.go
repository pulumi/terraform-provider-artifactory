// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a local Nuget repository.
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
//			_, err := artifactory.LookupLocalNugetRepository(ctx, &artifactory.LookupLocalNugetRepositoryArgs{
//				Key: "local-test-nuget-repo-basic",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupLocalNugetRepository(ctx *pulumi.Context, args *LookupLocalNugetRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupLocalNugetRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLocalNugetRepositoryResult
	err := ctx.Invoke("artifactory:index/getLocalNugetRepository:getLocalNugetRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocalNugetRepository.
type LookupLocalNugetRepositoryArgs struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	// Force basic authentication credentials in order to use this repository.
	// Default is `false`.
	ForceNugetAuthentication *bool   `pulumi:"forceNugetAuthentication"`
	IncludesPattern          *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// The maximum number of unique snapshots of a single artifact to store Once the
	// number of snapshots exceeds this setting, older versions are removed A value of 0 (default) indicates there is no
	// limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots  *int     `pulumi:"maxUniqueSnapshots"`
	Notes               *string  `pulumi:"notes"`
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	PropertySets        []string `pulumi:"propertySets"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	XrayIndex           *bool    `pulumi:"xrayIndex"`
}

// A collection of values returned by getLocalNugetRepository.
type LookupLocalNugetRepositoryResult struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	// Force basic authentication credentials in order to use this repository.
	// Default is `false`.
	ForceNugetAuthentication *bool `pulumi:"forceNugetAuthentication"`
	// The provider-assigned unique ID for this managed resource.
	Id              string  `pulumi:"id"`
	IncludesPattern *string `pulumi:"includesPattern"`
	Key             string  `pulumi:"key"`
	// The maximum number of unique snapshots of a single artifact to store Once the
	// number of snapshots exceeds this setting, older versions are removed A value of 0 (default) indicates there is no
	// limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots  *int     `pulumi:"maxUniqueSnapshots"`
	Notes               *string  `pulumi:"notes"`
	PackageType         string   `pulumi:"packageType"`
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	PropertySets        []string `pulumi:"propertySets"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	XrayIndex           *bool    `pulumi:"xrayIndex"`
}

func LookupLocalNugetRepositoryOutput(ctx *pulumi.Context, args LookupLocalNugetRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupLocalNugetRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLocalNugetRepositoryResultOutput, error) {
			args := v.(LookupLocalNugetRepositoryArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupLocalNugetRepositoryResult
			secret, err := ctx.InvokePackageRaw("artifactory:index/getLocalNugetRepository:getLocalNugetRepository", args, &rv, "", opts...)
			if err != nil {
				return LookupLocalNugetRepositoryResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupLocalNugetRepositoryResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupLocalNugetRepositoryResultOutput), nil
			}
			return output, nil
		}).(LookupLocalNugetRepositoryResultOutput)
}

// A collection of arguments for invoking getLocalNugetRepository.
type LookupLocalNugetRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput   `pulumi:"blackedOut"`
	CdnRedirect            pulumi.BoolPtrInput   `pulumi:"cdnRedirect"`
	Description            pulumi.StringPtrInput `pulumi:"description"`
	DownloadDirect         pulumi.BoolPtrInput   `pulumi:"downloadDirect"`
	ExcludesPattern        pulumi.StringPtrInput `pulumi:"excludesPattern"`
	// Force basic authentication credentials in order to use this repository.
	// Default is `false`.
	ForceNugetAuthentication pulumi.BoolPtrInput   `pulumi:"forceNugetAuthentication"`
	IncludesPattern          pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key pulumi.StringInput `pulumi:"key"`
	// The maximum number of unique snapshots of a single artifact to store Once the
	// number of snapshots exceeds this setting, older versions are removed A value of 0 (default) indicates there is no
	// limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots  pulumi.IntPtrInput      `pulumi:"maxUniqueSnapshots"`
	Notes               pulumi.StringPtrInput   `pulumi:"notes"`
	PriorityResolution  pulumi.BoolPtrInput     `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput   `pulumi:"projectKey"`
	PropertySets        pulumi.StringArrayInput `pulumi:"propertySets"`
	RepoLayoutRef       pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	XrayIndex           pulumi.BoolPtrInput     `pulumi:"xrayIndex"`
}

func (LookupLocalNugetRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalNugetRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getLocalNugetRepository.
type LookupLocalNugetRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupLocalNugetRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalNugetRepositoryResult)(nil)).Elem()
}

func (o LookupLocalNugetRepositoryResultOutput) ToLookupLocalNugetRepositoryResultOutput() LookupLocalNugetRepositoryResultOutput {
	return o
}

func (o LookupLocalNugetRepositoryResultOutput) ToLookupLocalNugetRepositoryResultOutputWithContext(ctx context.Context) LookupLocalNugetRepositoryResultOutput {
	return o
}

func (o LookupLocalNugetRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalNugetRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalNugetRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalNugetRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalNugetRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalNugetRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalNugetRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalNugetRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupLocalNugetRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalNugetRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalNugetRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalNugetRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// Force basic authentication credentials in order to use this repository.
// Default is `false`.
func (o LookupLocalNugetRepositoryResultOutput) ForceNugetAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalNugetRepositoryResult) *bool { return v.ForceNugetAuthentication }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLocalNugetRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalNugetRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLocalNugetRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalNugetRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupLocalNugetRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalNugetRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

// The maximum number of unique snapshots of a single artifact to store Once the
// number of snapshots exceeds this setting, older versions are removed A value of 0 (default) indicates there is no
// limit, and unique snapshots are not cleaned up.
func (o LookupLocalNugetRepositoryResultOutput) MaxUniqueSnapshots() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupLocalNugetRepositoryResult) *int { return v.MaxUniqueSnapshots }).(pulumi.IntPtrOutput)
}

func (o LookupLocalNugetRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalNugetRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupLocalNugetRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalNugetRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupLocalNugetRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalNugetRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalNugetRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocalNugetRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupLocalNugetRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalNugetRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupLocalNugetRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocalNugetRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupLocalNugetRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalNugetRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupLocalNugetRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalNugetRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocalNugetRepositoryResultOutput{})
}
