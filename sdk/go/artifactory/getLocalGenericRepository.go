// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a local generic repository.
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
//			_, err := artifactory.LookupLocalGenericRepository(ctx, &artifactory.LookupLocalGenericRepositoryArgs{
//				Key: "local-test-generic-repo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupLocalGenericRepository(ctx *pulumi.Context, args *LookupLocalGenericRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupLocalGenericRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLocalGenericRepositoryResult
	err := ctx.Invoke("artifactory:index/getLocalGenericRepository:getLocalGenericRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocalGenericRepository.
type LookupLocalGenericRepositoryArgs struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	IncludesPattern        *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                 string   `pulumi:"key"`
	Notes               *string  `pulumi:"notes"`
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	PropertySets        []string `pulumi:"propertySets"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	XrayIndex           *bool    `pulumi:"xrayIndex"`
}

// A collection of values returned by getLocalGenericRepository.
type LookupLocalGenericRepositoryResult struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id              string  `pulumi:"id"`
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                 string   `pulumi:"key"`
	Notes               *string  `pulumi:"notes"`
	PackageType         string   `pulumi:"packageType"`
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	PropertySets        []string `pulumi:"propertySets"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	XrayIndex           *bool    `pulumi:"xrayIndex"`
}

func LookupLocalGenericRepositoryOutput(ctx *pulumi.Context, args LookupLocalGenericRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupLocalGenericRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLocalGenericRepositoryResult, error) {
			args := v.(LookupLocalGenericRepositoryArgs)
			r, err := LookupLocalGenericRepository(ctx, &args, opts...)
			var s LookupLocalGenericRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLocalGenericRepositoryResultOutput)
}

// A collection of arguments for invoking getLocalGenericRepository.
type LookupLocalGenericRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput   `pulumi:"blackedOut"`
	CdnRedirect            pulumi.BoolPtrInput   `pulumi:"cdnRedirect"`
	Description            pulumi.StringPtrInput `pulumi:"description"`
	DownloadDirect         pulumi.BoolPtrInput   `pulumi:"downloadDirect"`
	ExcludesPattern        pulumi.StringPtrInput `pulumi:"excludesPattern"`
	IncludesPattern        pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                 pulumi.StringInput      `pulumi:"key"`
	Notes               pulumi.StringPtrInput   `pulumi:"notes"`
	PriorityResolution  pulumi.BoolPtrInput     `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput   `pulumi:"projectKey"`
	PropertySets        pulumi.StringArrayInput `pulumi:"propertySets"`
	RepoLayoutRef       pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	XrayIndex           pulumi.BoolPtrInput     `pulumi:"xrayIndex"`
}

func (LookupLocalGenericRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalGenericRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getLocalGenericRepository.
type LookupLocalGenericRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupLocalGenericRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalGenericRepositoryResult)(nil)).Elem()
}

func (o LookupLocalGenericRepositoryResultOutput) ToLookupLocalGenericRepositoryResultOutput() LookupLocalGenericRepositoryResultOutput {
	return o
}

func (o LookupLocalGenericRepositoryResultOutput) ToLookupLocalGenericRepositoryResultOutputWithContext(ctx context.Context) LookupLocalGenericRepositoryResultOutput {
	return o
}

func (o LookupLocalGenericRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalGenericRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalGenericRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalGenericRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalGenericRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalGenericRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalGenericRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalGenericRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupLocalGenericRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalGenericRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalGenericRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalGenericRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLocalGenericRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalGenericRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLocalGenericRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalGenericRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// the identity key of the repo.
func (o LookupLocalGenericRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalGenericRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupLocalGenericRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalGenericRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupLocalGenericRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalGenericRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupLocalGenericRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalGenericRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalGenericRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocalGenericRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupLocalGenericRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalGenericRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupLocalGenericRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocalGenericRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupLocalGenericRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalGenericRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupLocalGenericRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalGenericRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocalGenericRepositoryResultOutput{})
}
