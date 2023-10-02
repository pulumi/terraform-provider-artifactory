// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

func LookupFederatedDockerRepository(ctx *pulumi.Context, args *LookupFederatedDockerRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupFederatedDockerRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFederatedDockerRepositoryResult
	err := ctx.Invoke("artifactory:index/getFederatedDockerRepository:getFederatedDockerRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedDockerRepository.
type LookupFederatedDockerRepositoryArgs struct {
	ArchiveBrowsingEnabled *bool                                `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool                                `pulumi:"blackedOut"`
	BlockPushingSchema1    *bool                                `pulumi:"blockPushingSchema1"`
	CdnRedirect            *bool                                `pulumi:"cdnRedirect"`
	CleanupOnDelete        *bool                                `pulumi:"cleanupOnDelete"`
	Description            *string                              `pulumi:"description"`
	DownloadDirect         *bool                                `pulumi:"downloadDirect"`
	ExcludesPattern        *string                              `pulumi:"excludesPattern"`
	IncludesPattern        *string                              `pulumi:"includesPattern"`
	Key                    string                               `pulumi:"key"`
	MaxUniqueTags          *int                                 `pulumi:"maxUniqueTags"`
	Members                []GetFederatedDockerRepositoryMember `pulumi:"members"`
	Notes                  *string                              `pulumi:"notes"`
	PriorityResolution     *bool                                `pulumi:"priorityResolution"`
	ProjectEnvironments    []string                             `pulumi:"projectEnvironments"`
	ProjectKey             *string                              `pulumi:"projectKey"`
	PropertySets           []string                             `pulumi:"propertySets"`
	RepoLayoutRef          *string                              `pulumi:"repoLayoutRef"`
	TagRetention           *int                                 `pulumi:"tagRetention"`
	XrayIndex              *bool                                `pulumi:"xrayIndex"`
}

// A collection of values returned by getFederatedDockerRepository.
type LookupFederatedDockerRepositoryResult struct {
	ApiVersion             string  `pulumi:"apiVersion"`
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	BlockPushingSchema1    bool    `pulumi:"blockPushingSchema1"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        string  `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string                               `pulumi:"id"`
	IncludesPattern     string                               `pulumi:"includesPattern"`
	Key                 string                               `pulumi:"key"`
	MaxUniqueTags       *int                                 `pulumi:"maxUniqueTags"`
	Members             []GetFederatedDockerRepositoryMember `pulumi:"members"`
	Notes               *string                              `pulumi:"notes"`
	PackageType         string                               `pulumi:"packageType"`
	PriorityResolution  *bool                                `pulumi:"priorityResolution"`
	ProjectEnvironments []string                             `pulumi:"projectEnvironments"`
	ProjectKey          *string                              `pulumi:"projectKey"`
	PropertySets        []string                             `pulumi:"propertySets"`
	RepoLayoutRef       *string                              `pulumi:"repoLayoutRef"`
	TagRetention        *int                                 `pulumi:"tagRetention"`
	XrayIndex           *bool                                `pulumi:"xrayIndex"`
}

func LookupFederatedDockerRepositoryOutput(ctx *pulumi.Context, args LookupFederatedDockerRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedDockerRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedDockerRepositoryResult, error) {
			args := v.(LookupFederatedDockerRepositoryArgs)
			r, err := LookupFederatedDockerRepository(ctx, &args, opts...)
			var s LookupFederatedDockerRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFederatedDockerRepositoryResultOutput)
}

// A collection of arguments for invoking getFederatedDockerRepository.
type LookupFederatedDockerRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput                          `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput                          `pulumi:"blackedOut"`
	BlockPushingSchema1    pulumi.BoolPtrInput                          `pulumi:"blockPushingSchema1"`
	CdnRedirect            pulumi.BoolPtrInput                          `pulumi:"cdnRedirect"`
	CleanupOnDelete        pulumi.BoolPtrInput                          `pulumi:"cleanupOnDelete"`
	Description            pulumi.StringPtrInput                        `pulumi:"description"`
	DownloadDirect         pulumi.BoolPtrInput                          `pulumi:"downloadDirect"`
	ExcludesPattern        pulumi.StringPtrInput                        `pulumi:"excludesPattern"`
	IncludesPattern        pulumi.StringPtrInput                        `pulumi:"includesPattern"`
	Key                    pulumi.StringInput                           `pulumi:"key"`
	MaxUniqueTags          pulumi.IntPtrInput                           `pulumi:"maxUniqueTags"`
	Members                GetFederatedDockerRepositoryMemberArrayInput `pulumi:"members"`
	Notes                  pulumi.StringPtrInput                        `pulumi:"notes"`
	PriorityResolution     pulumi.BoolPtrInput                          `pulumi:"priorityResolution"`
	ProjectEnvironments    pulumi.StringArrayInput                      `pulumi:"projectEnvironments"`
	ProjectKey             pulumi.StringPtrInput                        `pulumi:"projectKey"`
	PropertySets           pulumi.StringArrayInput                      `pulumi:"propertySets"`
	RepoLayoutRef          pulumi.StringPtrInput                        `pulumi:"repoLayoutRef"`
	TagRetention           pulumi.IntPtrInput                           `pulumi:"tagRetention"`
	XrayIndex              pulumi.BoolPtrInput                          `pulumi:"xrayIndex"`
}

func (LookupFederatedDockerRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedDockerRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedDockerRepository.
type LookupFederatedDockerRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedDockerRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedDockerRepositoryResult)(nil)).Elem()
}

func (o LookupFederatedDockerRepositoryResultOutput) ToLookupFederatedDockerRepositoryResultOutput() LookupFederatedDockerRepositoryResultOutput {
	return o
}

func (o LookupFederatedDockerRepositoryResultOutput) ToLookupFederatedDockerRepositoryResultOutputWithContext(ctx context.Context) LookupFederatedDockerRepositoryResultOutput {
	return o
}

func (o LookupFederatedDockerRepositoryResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupFederatedDockerRepositoryResult] {
	return pulumix.Output[LookupFederatedDockerRepositoryResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupFederatedDockerRepositoryResultOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedDockerRepositoryResult) string { return v.ApiVersion }).(pulumi.StringOutput)
}

func (o LookupFederatedDockerRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedDockerRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedDockerRepositoryResultOutput) BlockPushingSchema1() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupFederatedDockerRepositoryResult) bool { return v.BlockPushingSchema1 }).(pulumi.BoolOutput)
}

func (o LookupFederatedDockerRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedDockerRepositoryResultOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerRepositoryResult) *bool { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedDockerRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedDockerRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedDockerRepositoryResultOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedDockerRepositoryResult) string { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFederatedDockerRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedDockerRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedDockerRepositoryResultOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedDockerRepositoryResult) string { return v.IncludesPattern }).(pulumi.StringOutput)
}

func (o LookupFederatedDockerRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedDockerRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupFederatedDockerRepositoryResultOutput) MaxUniqueTags() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerRepositoryResult) *int { return v.MaxUniqueTags }).(pulumi.IntPtrOutput)
}

func (o LookupFederatedDockerRepositoryResultOutput) Members() GetFederatedDockerRepositoryMemberArrayOutput {
	return o.ApplyT(func(v LookupFederatedDockerRepositoryResult) []GetFederatedDockerRepositoryMember { return v.Members }).(GetFederatedDockerRepositoryMemberArrayOutput)
}

func (o LookupFederatedDockerRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedDockerRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedDockerRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupFederatedDockerRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedDockerRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedDockerRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedDockerRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedDockerRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedDockerRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedDockerRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedDockerRepositoryResultOutput) TagRetention() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerRepositoryResult) *int { return v.TagRetention }).(pulumi.IntPtrOutput)
}

func (o LookupFederatedDockerRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedDockerRepositoryResultOutput{})
}
