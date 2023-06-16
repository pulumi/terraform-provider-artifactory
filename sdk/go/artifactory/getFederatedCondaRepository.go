// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFederatedCondaRepository(ctx *pulumi.Context, args *LookupFederatedCondaRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupFederatedCondaRepositoryResult, error) {
	var rv LookupFederatedCondaRepositoryResult
	err := ctx.Invoke("artifactory:index/getFederatedCondaRepository:getFederatedCondaRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedCondaRepository.
type LookupFederatedCondaRepositoryArgs struct {
	ArchiveBrowsingEnabled *bool                               `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool                               `pulumi:"blackedOut"`
	CdnRedirect            *bool                               `pulumi:"cdnRedirect"`
	CleanupOnDelete        *bool                               `pulumi:"cleanupOnDelete"`
	Description            *string                             `pulumi:"description"`
	DownloadDirect         *bool                               `pulumi:"downloadDirect"`
	ExcludesPattern        *string                             `pulumi:"excludesPattern"`
	IncludesPattern        *string                             `pulumi:"includesPattern"`
	Key                    string                              `pulumi:"key"`
	Members                []GetFederatedCondaRepositoryMember `pulumi:"members"`
	Notes                  *string                             `pulumi:"notes"`
	PriorityResolution     *bool                               `pulumi:"priorityResolution"`
	ProjectEnvironments    []string                            `pulumi:"projectEnvironments"`
	ProjectKey             *string                             `pulumi:"projectKey"`
	PropertySets           []string                            `pulumi:"propertySets"`
	RepoLayoutRef          *string                             `pulumi:"repoLayoutRef"`
	XrayIndex              *bool                               `pulumi:"xrayIndex"`
}

// A collection of values returned by getFederatedCondaRepository.
type LookupFederatedCondaRepositoryResult struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        string  `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string                              `pulumi:"id"`
	IncludesPattern     string                              `pulumi:"includesPattern"`
	Key                 string                              `pulumi:"key"`
	Members             []GetFederatedCondaRepositoryMember `pulumi:"members"`
	Notes               *string                             `pulumi:"notes"`
	PackageType         string                              `pulumi:"packageType"`
	PriorityResolution  *bool                               `pulumi:"priorityResolution"`
	ProjectEnvironments []string                            `pulumi:"projectEnvironments"`
	ProjectKey          *string                             `pulumi:"projectKey"`
	PropertySets        []string                            `pulumi:"propertySets"`
	RepoLayoutRef       *string                             `pulumi:"repoLayoutRef"`
	XrayIndex           *bool                               `pulumi:"xrayIndex"`
}

func LookupFederatedCondaRepositoryOutput(ctx *pulumi.Context, args LookupFederatedCondaRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedCondaRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedCondaRepositoryResult, error) {
			args := v.(LookupFederatedCondaRepositoryArgs)
			r, err := LookupFederatedCondaRepository(ctx, &args, opts...)
			var s LookupFederatedCondaRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFederatedCondaRepositoryResultOutput)
}

// A collection of arguments for invoking getFederatedCondaRepository.
type LookupFederatedCondaRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput                         `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput                         `pulumi:"blackedOut"`
	CdnRedirect            pulumi.BoolPtrInput                         `pulumi:"cdnRedirect"`
	CleanupOnDelete        pulumi.BoolPtrInput                         `pulumi:"cleanupOnDelete"`
	Description            pulumi.StringPtrInput                       `pulumi:"description"`
	DownloadDirect         pulumi.BoolPtrInput                         `pulumi:"downloadDirect"`
	ExcludesPattern        pulumi.StringPtrInput                       `pulumi:"excludesPattern"`
	IncludesPattern        pulumi.StringPtrInput                       `pulumi:"includesPattern"`
	Key                    pulumi.StringInput                          `pulumi:"key"`
	Members                GetFederatedCondaRepositoryMemberArrayInput `pulumi:"members"`
	Notes                  pulumi.StringPtrInput                       `pulumi:"notes"`
	PriorityResolution     pulumi.BoolPtrInput                         `pulumi:"priorityResolution"`
	ProjectEnvironments    pulumi.StringArrayInput                     `pulumi:"projectEnvironments"`
	ProjectKey             pulumi.StringPtrInput                       `pulumi:"projectKey"`
	PropertySets           pulumi.StringArrayInput                     `pulumi:"propertySets"`
	RepoLayoutRef          pulumi.StringPtrInput                       `pulumi:"repoLayoutRef"`
	XrayIndex              pulumi.BoolPtrInput                         `pulumi:"xrayIndex"`
}

func (LookupFederatedCondaRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedCondaRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedCondaRepository.
type LookupFederatedCondaRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedCondaRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedCondaRepositoryResult)(nil)).Elem()
}

func (o LookupFederatedCondaRepositoryResultOutput) ToLookupFederatedCondaRepositoryResultOutput() LookupFederatedCondaRepositoryResultOutput {
	return o
}

func (o LookupFederatedCondaRepositoryResultOutput) ToLookupFederatedCondaRepositoryResultOutputWithContext(ctx context.Context) LookupFederatedCondaRepositoryResultOutput {
	return o
}

func (o LookupFederatedCondaRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCondaRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCondaRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCondaRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCondaRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCondaRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCondaRepositoryResultOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCondaRepositoryResult) *bool { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCondaRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedCondaRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedCondaRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCondaRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCondaRepositoryResultOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedCondaRepositoryResult) string { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFederatedCondaRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedCondaRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedCondaRepositoryResultOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedCondaRepositoryResult) string { return v.IncludesPattern }).(pulumi.StringOutput)
}

func (o LookupFederatedCondaRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedCondaRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupFederatedCondaRepositoryResultOutput) Members() GetFederatedCondaRepositoryMemberArrayOutput {
	return o.ApplyT(func(v LookupFederatedCondaRepositoryResult) []GetFederatedCondaRepositoryMember { return v.Members }).(GetFederatedCondaRepositoryMemberArrayOutput)
}

func (o LookupFederatedCondaRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedCondaRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedCondaRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedCondaRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupFederatedCondaRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCondaRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedCondaRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedCondaRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedCondaRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedCondaRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedCondaRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedCondaRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedCondaRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedCondaRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedCondaRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedCondaRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedCondaRepositoryResultOutput{})
}
