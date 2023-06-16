// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFederatedGoRepository(ctx *pulumi.Context, args *LookupFederatedGoRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupFederatedGoRepositoryResult, error) {
	var rv LookupFederatedGoRepositoryResult
	err := ctx.Invoke("artifactory:index/getFederatedGoRepository:getFederatedGoRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedGoRepository.
type LookupFederatedGoRepositoryArgs struct {
	ArchiveBrowsingEnabled *bool                            `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool                            `pulumi:"blackedOut"`
	CdnRedirect            *bool                            `pulumi:"cdnRedirect"`
	CleanupOnDelete        *bool                            `pulumi:"cleanupOnDelete"`
	Description            *string                          `pulumi:"description"`
	DownloadDirect         *bool                            `pulumi:"downloadDirect"`
	ExcludesPattern        *string                          `pulumi:"excludesPattern"`
	IncludesPattern        *string                          `pulumi:"includesPattern"`
	Key                    string                           `pulumi:"key"`
	Members                []GetFederatedGoRepositoryMember `pulumi:"members"`
	Notes                  *string                          `pulumi:"notes"`
	PriorityResolution     *bool                            `pulumi:"priorityResolution"`
	ProjectEnvironments    []string                         `pulumi:"projectEnvironments"`
	ProjectKey             *string                          `pulumi:"projectKey"`
	PropertySets           []string                         `pulumi:"propertySets"`
	RepoLayoutRef          *string                          `pulumi:"repoLayoutRef"`
	XrayIndex              *bool                            `pulumi:"xrayIndex"`
}

// A collection of values returned by getFederatedGoRepository.
type LookupFederatedGoRepositoryResult struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        string  `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string                           `pulumi:"id"`
	IncludesPattern     string                           `pulumi:"includesPattern"`
	Key                 string                           `pulumi:"key"`
	Members             []GetFederatedGoRepositoryMember `pulumi:"members"`
	Notes               *string                          `pulumi:"notes"`
	PackageType         string                           `pulumi:"packageType"`
	PriorityResolution  *bool                            `pulumi:"priorityResolution"`
	ProjectEnvironments []string                         `pulumi:"projectEnvironments"`
	ProjectKey          *string                          `pulumi:"projectKey"`
	PropertySets        []string                         `pulumi:"propertySets"`
	RepoLayoutRef       *string                          `pulumi:"repoLayoutRef"`
	XrayIndex           *bool                            `pulumi:"xrayIndex"`
}

func LookupFederatedGoRepositoryOutput(ctx *pulumi.Context, args LookupFederatedGoRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedGoRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedGoRepositoryResult, error) {
			args := v.(LookupFederatedGoRepositoryArgs)
			r, err := LookupFederatedGoRepository(ctx, &args, opts...)
			var s LookupFederatedGoRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFederatedGoRepositoryResultOutput)
}

// A collection of arguments for invoking getFederatedGoRepository.
type LookupFederatedGoRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput                      `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput                      `pulumi:"blackedOut"`
	CdnRedirect            pulumi.BoolPtrInput                      `pulumi:"cdnRedirect"`
	CleanupOnDelete        pulumi.BoolPtrInput                      `pulumi:"cleanupOnDelete"`
	Description            pulumi.StringPtrInput                    `pulumi:"description"`
	DownloadDirect         pulumi.BoolPtrInput                      `pulumi:"downloadDirect"`
	ExcludesPattern        pulumi.StringPtrInput                    `pulumi:"excludesPattern"`
	IncludesPattern        pulumi.StringPtrInput                    `pulumi:"includesPattern"`
	Key                    pulumi.StringInput                       `pulumi:"key"`
	Members                GetFederatedGoRepositoryMemberArrayInput `pulumi:"members"`
	Notes                  pulumi.StringPtrInput                    `pulumi:"notes"`
	PriorityResolution     pulumi.BoolPtrInput                      `pulumi:"priorityResolution"`
	ProjectEnvironments    pulumi.StringArrayInput                  `pulumi:"projectEnvironments"`
	ProjectKey             pulumi.StringPtrInput                    `pulumi:"projectKey"`
	PropertySets           pulumi.StringArrayInput                  `pulumi:"propertySets"`
	RepoLayoutRef          pulumi.StringPtrInput                    `pulumi:"repoLayoutRef"`
	XrayIndex              pulumi.BoolPtrInput                      `pulumi:"xrayIndex"`
}

func (LookupFederatedGoRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedGoRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedGoRepository.
type LookupFederatedGoRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedGoRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedGoRepositoryResult)(nil)).Elem()
}

func (o LookupFederatedGoRepositoryResultOutput) ToLookupFederatedGoRepositoryResultOutput() LookupFederatedGoRepositoryResultOutput {
	return o
}

func (o LookupFederatedGoRepositoryResultOutput) ToLookupFederatedGoRepositoryResultOutputWithContext(ctx context.Context) LookupFederatedGoRepositoryResultOutput {
	return o
}

func (o LookupFederatedGoRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedGoRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedGoRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedGoRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedGoRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedGoRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedGoRepositoryResultOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedGoRepositoryResult) *bool { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedGoRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedGoRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedGoRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedGoRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedGoRepositoryResultOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedGoRepositoryResult) string { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFederatedGoRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedGoRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedGoRepositoryResultOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedGoRepositoryResult) string { return v.IncludesPattern }).(pulumi.StringOutput)
}

func (o LookupFederatedGoRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedGoRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupFederatedGoRepositoryResultOutput) Members() GetFederatedGoRepositoryMemberArrayOutput {
	return o.ApplyT(func(v LookupFederatedGoRepositoryResult) []GetFederatedGoRepositoryMember { return v.Members }).(GetFederatedGoRepositoryMemberArrayOutput)
}

func (o LookupFederatedGoRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedGoRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedGoRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedGoRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupFederatedGoRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedGoRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedGoRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedGoRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedGoRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedGoRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedGoRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedGoRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedGoRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedGoRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedGoRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedGoRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedGoRepositoryResultOutput{})
}
