// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
//			_, err := artifactory.GetLocalTerraformbackendRepository(ctx, &artifactory.GetLocalTerraformbackendRepositoryArgs{
//				Key: "local-test-terraformbackend-repo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetLocalTerraformbackendRepository(ctx *pulumi.Context, args *GetLocalTerraformbackendRepositoryArgs, opts ...pulumi.InvokeOption) (*GetLocalTerraformbackendRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLocalTerraformbackendRepositoryResult
	err := ctx.Invoke("artifactory:index/getLocalTerraformbackendRepository:getLocalTerraformbackendRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocalTerraformbackendRepository.
type GetLocalTerraformbackendRepositoryArgs struct {
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

// A collection of values returned by getLocalTerraformbackendRepository.
type GetLocalTerraformbackendRepositoryResult struct {
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

func GetLocalTerraformbackendRepositoryOutput(ctx *pulumi.Context, args GetLocalTerraformbackendRepositoryOutputArgs, opts ...pulumi.InvokeOption) GetLocalTerraformbackendRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLocalTerraformbackendRepositoryResult, error) {
			args := v.(GetLocalTerraformbackendRepositoryArgs)
			r, err := GetLocalTerraformbackendRepository(ctx, &args, opts...)
			var s GetLocalTerraformbackendRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLocalTerraformbackendRepositoryResultOutput)
}

// A collection of arguments for invoking getLocalTerraformbackendRepository.
type GetLocalTerraformbackendRepositoryOutputArgs struct {
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

func (GetLocalTerraformbackendRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLocalTerraformbackendRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getLocalTerraformbackendRepository.
type GetLocalTerraformbackendRepositoryResultOutput struct{ *pulumi.OutputState }

func (GetLocalTerraformbackendRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLocalTerraformbackendRepositoryResult)(nil)).Elem()
}

func (o GetLocalTerraformbackendRepositoryResultOutput) ToGetLocalTerraformbackendRepositoryResultOutput() GetLocalTerraformbackendRepositoryResultOutput {
	return o
}

func (o GetLocalTerraformbackendRepositoryResultOutput) ToGetLocalTerraformbackendRepositoryResultOutputWithContext(ctx context.Context) GetLocalTerraformbackendRepositoryResultOutput {
	return o
}

func (o GetLocalTerraformbackendRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalTerraformbackendRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o GetLocalTerraformbackendRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalTerraformbackendRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o GetLocalTerraformbackendRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalTerraformbackendRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o GetLocalTerraformbackendRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalTerraformbackendRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GetLocalTerraformbackendRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalTerraformbackendRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o GetLocalTerraformbackendRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalTerraformbackendRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLocalTerraformbackendRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalTerraformbackendRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLocalTerraformbackendRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalTerraformbackendRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// the identity key of the repo.
func (o GetLocalTerraformbackendRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalTerraformbackendRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o GetLocalTerraformbackendRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalTerraformbackendRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o GetLocalTerraformbackendRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalTerraformbackendRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o GetLocalTerraformbackendRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalTerraformbackendRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o GetLocalTerraformbackendRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLocalTerraformbackendRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o GetLocalTerraformbackendRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalTerraformbackendRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o GetLocalTerraformbackendRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLocalTerraformbackendRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o GetLocalTerraformbackendRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalTerraformbackendRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o GetLocalTerraformbackendRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalTerraformbackendRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLocalTerraformbackendRepositoryResultOutput{})
}
