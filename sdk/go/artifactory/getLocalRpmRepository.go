// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v4/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLocalRpmRepository(ctx *pulumi.Context, args *LookupLocalRpmRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupLocalRpmRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLocalRpmRepositoryResult
	err := ctx.Invoke("artifactory:index/getLocalRpmRepository:getLocalRpmRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocalRpmRepository.
type LookupLocalRpmRepositoryArgs struct {
	ArchiveBrowsingEnabled  *bool    `pulumi:"archiveBrowsingEnabled"`
	BlackedOut              *bool    `pulumi:"blackedOut"`
	CalculateYumMetadata    *bool    `pulumi:"calculateYumMetadata"`
	CdnRedirect             *bool    `pulumi:"cdnRedirect"`
	Description             *string  `pulumi:"description"`
	DownloadDirect          *bool    `pulumi:"downloadDirect"`
	EnableFileListsIndexing *bool    `pulumi:"enableFileListsIndexing"`
	ExcludesPattern         *string  `pulumi:"excludesPattern"`
	IncludesPattern         *string  `pulumi:"includesPattern"`
	Key                     string   `pulumi:"key"`
	Notes                   *string  `pulumi:"notes"`
	PrimaryKeypairRef       *string  `pulumi:"primaryKeypairRef"`
	PriorityResolution      *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments     []string `pulumi:"projectEnvironments"`
	ProjectKey              *string  `pulumi:"projectKey"`
	PropertySets            []string `pulumi:"propertySets"`
	RepoLayoutRef           *string  `pulumi:"repoLayoutRef"`
	SecondaryKeypairRef     *string  `pulumi:"secondaryKeypairRef"`
	XrayIndex               *bool    `pulumi:"xrayIndex"`
	YumGroupFileNames       *string  `pulumi:"yumGroupFileNames"`
	YumRootDepth            *int     `pulumi:"yumRootDepth"`
}

// A collection of values returned by getLocalRpmRepository.
type LookupLocalRpmRepositoryResult struct {
	ArchiveBrowsingEnabled  *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut              *bool   `pulumi:"blackedOut"`
	CalculateYumMetadata    *bool   `pulumi:"calculateYumMetadata"`
	CdnRedirect             *bool   `pulumi:"cdnRedirect"`
	Description             *string `pulumi:"description"`
	DownloadDirect          *bool   `pulumi:"downloadDirect"`
	EnableFileListsIndexing *bool   `pulumi:"enableFileListsIndexing"`
	ExcludesPattern         string  `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string   `pulumi:"id"`
	IncludesPattern     string   `pulumi:"includesPattern"`
	Key                 string   `pulumi:"key"`
	Notes               *string  `pulumi:"notes"`
	PackageType         string   `pulumi:"packageType"`
	PrimaryKeypairRef   *string  `pulumi:"primaryKeypairRef"`
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	PropertySets        []string `pulumi:"propertySets"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	SecondaryKeypairRef *string  `pulumi:"secondaryKeypairRef"`
	XrayIndex           *bool    `pulumi:"xrayIndex"`
	YumGroupFileNames   *string  `pulumi:"yumGroupFileNames"`
	YumRootDepth        *int     `pulumi:"yumRootDepth"`
}

func LookupLocalRpmRepositoryOutput(ctx *pulumi.Context, args LookupLocalRpmRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupLocalRpmRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLocalRpmRepositoryResult, error) {
			args := v.(LookupLocalRpmRepositoryArgs)
			r, err := LookupLocalRpmRepository(ctx, &args, opts...)
			var s LookupLocalRpmRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLocalRpmRepositoryResultOutput)
}

// A collection of arguments for invoking getLocalRpmRepository.
type LookupLocalRpmRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled  pulumi.BoolPtrInput     `pulumi:"archiveBrowsingEnabled"`
	BlackedOut              pulumi.BoolPtrInput     `pulumi:"blackedOut"`
	CalculateYumMetadata    pulumi.BoolPtrInput     `pulumi:"calculateYumMetadata"`
	CdnRedirect             pulumi.BoolPtrInput     `pulumi:"cdnRedirect"`
	Description             pulumi.StringPtrInput   `pulumi:"description"`
	DownloadDirect          pulumi.BoolPtrInput     `pulumi:"downloadDirect"`
	EnableFileListsIndexing pulumi.BoolPtrInput     `pulumi:"enableFileListsIndexing"`
	ExcludesPattern         pulumi.StringPtrInput   `pulumi:"excludesPattern"`
	IncludesPattern         pulumi.StringPtrInput   `pulumi:"includesPattern"`
	Key                     pulumi.StringInput      `pulumi:"key"`
	Notes                   pulumi.StringPtrInput   `pulumi:"notes"`
	PrimaryKeypairRef       pulumi.StringPtrInput   `pulumi:"primaryKeypairRef"`
	PriorityResolution      pulumi.BoolPtrInput     `pulumi:"priorityResolution"`
	ProjectEnvironments     pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey              pulumi.StringPtrInput   `pulumi:"projectKey"`
	PropertySets            pulumi.StringArrayInput `pulumi:"propertySets"`
	RepoLayoutRef           pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	SecondaryKeypairRef     pulumi.StringPtrInput   `pulumi:"secondaryKeypairRef"`
	XrayIndex               pulumi.BoolPtrInput     `pulumi:"xrayIndex"`
	YumGroupFileNames       pulumi.StringPtrInput   `pulumi:"yumGroupFileNames"`
	YumRootDepth            pulumi.IntPtrInput      `pulumi:"yumRootDepth"`
}

func (LookupLocalRpmRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalRpmRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getLocalRpmRepository.
type LookupLocalRpmRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupLocalRpmRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalRpmRepositoryResult)(nil)).Elem()
}

func (o LookupLocalRpmRepositoryResultOutput) ToLookupLocalRpmRepositoryResultOutput() LookupLocalRpmRepositoryResultOutput {
	return o
}

func (o LookupLocalRpmRepositoryResultOutput) ToLookupLocalRpmRepositoryResultOutputWithContext(ctx context.Context) LookupLocalRpmRepositoryResultOutput {
	return o
}

func (o LookupLocalRpmRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalRpmRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalRpmRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalRpmRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalRpmRepositoryResultOutput) CalculateYumMetadata() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalRpmRepositoryResult) *bool { return v.CalculateYumMetadata }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalRpmRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalRpmRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalRpmRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalRpmRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupLocalRpmRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalRpmRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalRpmRepositoryResultOutput) EnableFileListsIndexing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalRpmRepositoryResult) *bool { return v.EnableFileListsIndexing }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalRpmRepositoryResultOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalRpmRepositoryResult) string { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLocalRpmRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalRpmRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLocalRpmRepositoryResultOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalRpmRepositoryResult) string { return v.IncludesPattern }).(pulumi.StringOutput)
}

func (o LookupLocalRpmRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalRpmRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupLocalRpmRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalRpmRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupLocalRpmRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalRpmRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupLocalRpmRepositoryResultOutput) PrimaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalRpmRepositoryResult) *string { return v.PrimaryKeypairRef }).(pulumi.StringPtrOutput)
}

func (o LookupLocalRpmRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalRpmRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalRpmRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocalRpmRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupLocalRpmRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalRpmRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupLocalRpmRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocalRpmRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupLocalRpmRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalRpmRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupLocalRpmRepositoryResultOutput) SecondaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalRpmRepositoryResult) *string { return v.SecondaryKeypairRef }).(pulumi.StringPtrOutput)
}

func (o LookupLocalRpmRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalRpmRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalRpmRepositoryResultOutput) YumGroupFileNames() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalRpmRepositoryResult) *string { return v.YumGroupFileNames }).(pulumi.StringPtrOutput)
}

func (o LookupLocalRpmRepositoryResultOutput) YumRootDepth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupLocalRpmRepositoryResult) *int { return v.YumRootDepth }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocalRpmRepositoryResultOutput{})
}
