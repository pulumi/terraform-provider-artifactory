// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a local Chef repository.
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
//			_, err := artifactory.LookupLocalChefRepository(ctx, &artifactory.LookupLocalChefRepositoryArgs{
//				Key: "local-test-chef-repo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupLocalChefRepository(ctx *pulumi.Context, args *LookupLocalChefRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupLocalChefRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLocalChefRepositoryResult
	err := ctx.Invoke("artifactory:index/getLocalChefRepository:getLocalChefRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocalChefRepository.
type LookupLocalChefRepositoryArgs struct {
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool `pulumi:"blackedOut"`
	CdnRedirect            *bool `pulumi:"cdnRedirect"`
	// (Optional)
	Description     *string `pulumi:"description"`
	DownloadDirect  *bool   `pulumi:"downloadDirect"`
	ExcludesPattern *string `pulumi:"excludesPattern"`
	IncludesPattern *string `pulumi:"includesPattern"`
	// (Required) the identity key of the repo.
	Key string `pulumi:"key"`
	// (Optional)
	Notes               *string  `pulumi:"notes"`
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	PropertySets        []string `pulumi:"propertySets"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	XrayIndex           *bool    `pulumi:"xrayIndex"`
}

// A collection of values returned by getLocalChefRepository.
type LookupLocalChefRepositoryResult struct {
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool `pulumi:"blackedOut"`
	CdnRedirect            *bool `pulumi:"cdnRedirect"`
	// (Optional)
	Description     *string `pulumi:"description"`
	DownloadDirect  *bool   `pulumi:"downloadDirect"`
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id              string  `pulumi:"id"`
	IncludesPattern *string `pulumi:"includesPattern"`
	// (Required) the identity key of the repo.
	Key string `pulumi:"key"`
	// (Optional)
	Notes               *string  `pulumi:"notes"`
	PackageType         string   `pulumi:"packageType"`
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	PropertySets        []string `pulumi:"propertySets"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	XrayIndex           *bool    `pulumi:"xrayIndex"`
}

func LookupLocalChefRepositoryOutput(ctx *pulumi.Context, args LookupLocalChefRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupLocalChefRepositoryResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupLocalChefRepositoryResultOutput, error) {
			args := v.(LookupLocalChefRepositoryArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("artifactory:index/getLocalChefRepository:getLocalChefRepository", args, LookupLocalChefRepositoryResultOutput{}, options).(LookupLocalChefRepositoryResultOutput), nil
		}).(LookupLocalChefRepositoryResultOutput)
}

// A collection of arguments for invoking getLocalChefRepository.
type LookupLocalChefRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput `pulumi:"blackedOut"`
	CdnRedirect            pulumi.BoolPtrInput `pulumi:"cdnRedirect"`
	// (Optional)
	Description     pulumi.StringPtrInput `pulumi:"description"`
	DownloadDirect  pulumi.BoolPtrInput   `pulumi:"downloadDirect"`
	ExcludesPattern pulumi.StringPtrInput `pulumi:"excludesPattern"`
	IncludesPattern pulumi.StringPtrInput `pulumi:"includesPattern"`
	// (Required) the identity key of the repo.
	Key pulumi.StringInput `pulumi:"key"`
	// (Optional)
	Notes               pulumi.StringPtrInput   `pulumi:"notes"`
	PriorityResolution  pulumi.BoolPtrInput     `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput   `pulumi:"projectKey"`
	PropertySets        pulumi.StringArrayInput `pulumi:"propertySets"`
	RepoLayoutRef       pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	XrayIndex           pulumi.BoolPtrInput     `pulumi:"xrayIndex"`
}

func (LookupLocalChefRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalChefRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getLocalChefRepository.
type LookupLocalChefRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupLocalChefRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalChefRepositoryResult)(nil)).Elem()
}

func (o LookupLocalChefRepositoryResultOutput) ToLookupLocalChefRepositoryResultOutput() LookupLocalChefRepositoryResultOutput {
	return o
}

func (o LookupLocalChefRepositoryResultOutput) ToLookupLocalChefRepositoryResultOutputWithContext(ctx context.Context) LookupLocalChefRepositoryResultOutput {
	return o
}

func (o LookupLocalChefRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalChefRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalChefRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalChefRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalChefRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalChefRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

// (Optional)
func (o LookupLocalChefRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalChefRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupLocalChefRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalChefRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalChefRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalChefRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLocalChefRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalChefRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLocalChefRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalChefRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// (Required) the identity key of the repo.
func (o LookupLocalChefRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalChefRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

// (Optional)
func (o LookupLocalChefRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalChefRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupLocalChefRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalChefRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupLocalChefRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalChefRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalChefRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocalChefRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupLocalChefRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalChefRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupLocalChefRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocalChefRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupLocalChefRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalChefRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupLocalChefRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalChefRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocalChefRepositoryResultOutput{})
}
