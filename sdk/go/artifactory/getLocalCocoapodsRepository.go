// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v4/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLocalCocoapodsRepository(ctx *pulumi.Context, args *LookupLocalCocoapodsRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupLocalCocoapodsRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLocalCocoapodsRepositoryResult
	err := ctx.Invoke("artifactory:index/getLocalCocoapodsRepository:getLocalCocoapodsRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocalCocoapodsRepository.
type LookupLocalCocoapodsRepositoryArgs struct {
	ArchiveBrowsingEnabled *bool    `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool    `pulumi:"blackedOut"`
	CdnRedirect            *bool    `pulumi:"cdnRedirect"`
	Description            *string  `pulumi:"description"`
	DownloadDirect         *bool    `pulumi:"downloadDirect"`
	ExcludesPattern        *string  `pulumi:"excludesPattern"`
	IncludesPattern        *string  `pulumi:"includesPattern"`
	Key                    string   `pulumi:"key"`
	Notes                  *string  `pulumi:"notes"`
	PriorityResolution     *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments    []string `pulumi:"projectEnvironments"`
	ProjectKey             *string  `pulumi:"projectKey"`
	PropertySets           []string `pulumi:"propertySets"`
	RepoLayoutRef          *string  `pulumi:"repoLayoutRef"`
	XrayIndex              *bool    `pulumi:"xrayIndex"`
}

// A collection of values returned by getLocalCocoapodsRepository.
type LookupLocalCocoapodsRepositoryResult struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        string  `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string   `pulumi:"id"`
	IncludesPattern     string   `pulumi:"includesPattern"`
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

func LookupLocalCocoapodsRepositoryOutput(ctx *pulumi.Context, args LookupLocalCocoapodsRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupLocalCocoapodsRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLocalCocoapodsRepositoryResult, error) {
			args := v.(LookupLocalCocoapodsRepositoryArgs)
			r, err := LookupLocalCocoapodsRepository(ctx, &args, opts...)
			var s LookupLocalCocoapodsRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLocalCocoapodsRepositoryResultOutput)
}

// A collection of arguments for invoking getLocalCocoapodsRepository.
type LookupLocalCocoapodsRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput     `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput     `pulumi:"blackedOut"`
	CdnRedirect            pulumi.BoolPtrInput     `pulumi:"cdnRedirect"`
	Description            pulumi.StringPtrInput   `pulumi:"description"`
	DownloadDirect         pulumi.BoolPtrInput     `pulumi:"downloadDirect"`
	ExcludesPattern        pulumi.StringPtrInput   `pulumi:"excludesPattern"`
	IncludesPattern        pulumi.StringPtrInput   `pulumi:"includesPattern"`
	Key                    pulumi.StringInput      `pulumi:"key"`
	Notes                  pulumi.StringPtrInput   `pulumi:"notes"`
	PriorityResolution     pulumi.BoolPtrInput     `pulumi:"priorityResolution"`
	ProjectEnvironments    pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey             pulumi.StringPtrInput   `pulumi:"projectKey"`
	PropertySets           pulumi.StringArrayInput `pulumi:"propertySets"`
	RepoLayoutRef          pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	XrayIndex              pulumi.BoolPtrInput     `pulumi:"xrayIndex"`
}

func (LookupLocalCocoapodsRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalCocoapodsRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getLocalCocoapodsRepository.
type LookupLocalCocoapodsRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupLocalCocoapodsRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalCocoapodsRepositoryResult)(nil)).Elem()
}

func (o LookupLocalCocoapodsRepositoryResultOutput) ToLookupLocalCocoapodsRepositoryResultOutput() LookupLocalCocoapodsRepositoryResultOutput {
	return o
}

func (o LookupLocalCocoapodsRepositoryResultOutput) ToLookupLocalCocoapodsRepositoryResultOutputWithContext(ctx context.Context) LookupLocalCocoapodsRepositoryResultOutput {
	return o
}

func (o LookupLocalCocoapodsRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalCocoapodsRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalCocoapodsRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalCocoapodsRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalCocoapodsRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalCocoapodsRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalCocoapodsRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalCocoapodsRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupLocalCocoapodsRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalCocoapodsRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalCocoapodsRepositoryResultOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalCocoapodsRepositoryResult) string { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLocalCocoapodsRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalCocoapodsRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLocalCocoapodsRepositoryResultOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalCocoapodsRepositoryResult) string { return v.IncludesPattern }).(pulumi.StringOutput)
}

func (o LookupLocalCocoapodsRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalCocoapodsRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupLocalCocoapodsRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalCocoapodsRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupLocalCocoapodsRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalCocoapodsRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupLocalCocoapodsRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalCocoapodsRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalCocoapodsRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocalCocoapodsRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupLocalCocoapodsRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalCocoapodsRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupLocalCocoapodsRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocalCocoapodsRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupLocalCocoapodsRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalCocoapodsRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupLocalCocoapodsRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalCocoapodsRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocalCocoapodsRepositoryResultOutput{})
}
