// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFederatedGenericRepository(ctx *pulumi.Context, args *LookupFederatedGenericRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupFederatedGenericRepositoryResult, error) {
	var rv LookupFederatedGenericRepositoryResult
	err := ctx.Invoke("artifactory:index/getFederatedGenericRepository:getFederatedGenericRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedGenericRepository.
type LookupFederatedGenericRepositoryArgs struct {
	ArchiveBrowsingEnabled *bool                                 `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool                                 `pulumi:"blackedOut"`
	CdnRedirect            *bool                                 `pulumi:"cdnRedirect"`
	CleanupOnDelete        *bool                                 `pulumi:"cleanupOnDelete"`
	Description            *string                               `pulumi:"description"`
	DownloadDirect         *bool                                 `pulumi:"downloadDirect"`
	ExcludesPattern        *string                               `pulumi:"excludesPattern"`
	IncludesPattern        *string                               `pulumi:"includesPattern"`
	Key                    string                                `pulumi:"key"`
	Members                []GetFederatedGenericRepositoryMember `pulumi:"members"`
	Notes                  *string                               `pulumi:"notes"`
	PriorityResolution     *bool                                 `pulumi:"priorityResolution"`
	ProjectEnvironments    []string                              `pulumi:"projectEnvironments"`
	ProjectKey             *string                               `pulumi:"projectKey"`
	PropertySets           []string                              `pulumi:"propertySets"`
	RepoLayoutRef          *string                               `pulumi:"repoLayoutRef"`
	XrayIndex              *bool                                 `pulumi:"xrayIndex"`
}

// A collection of values returned by getFederatedGenericRepository.
type LookupFederatedGenericRepositoryResult struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        string  `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string                                `pulumi:"id"`
	IncludesPattern     string                                `pulumi:"includesPattern"`
	Key                 string                                `pulumi:"key"`
	Members             []GetFederatedGenericRepositoryMember `pulumi:"members"`
	Notes               *string                               `pulumi:"notes"`
	PackageType         string                                `pulumi:"packageType"`
	PriorityResolution  *bool                                 `pulumi:"priorityResolution"`
	ProjectEnvironments []string                              `pulumi:"projectEnvironments"`
	ProjectKey          *string                               `pulumi:"projectKey"`
	PropertySets        []string                              `pulumi:"propertySets"`
	RepoLayoutRef       *string                               `pulumi:"repoLayoutRef"`
	XrayIndex           *bool                                 `pulumi:"xrayIndex"`
}

func LookupFederatedGenericRepositoryOutput(ctx *pulumi.Context, args LookupFederatedGenericRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedGenericRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedGenericRepositoryResult, error) {
			args := v.(LookupFederatedGenericRepositoryArgs)
			r, err := LookupFederatedGenericRepository(ctx, &args, opts...)
			var s LookupFederatedGenericRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFederatedGenericRepositoryResultOutput)
}

// A collection of arguments for invoking getFederatedGenericRepository.
type LookupFederatedGenericRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput                           `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput                           `pulumi:"blackedOut"`
	CdnRedirect            pulumi.BoolPtrInput                           `pulumi:"cdnRedirect"`
	CleanupOnDelete        pulumi.BoolPtrInput                           `pulumi:"cleanupOnDelete"`
	Description            pulumi.StringPtrInput                         `pulumi:"description"`
	DownloadDirect         pulumi.BoolPtrInput                           `pulumi:"downloadDirect"`
	ExcludesPattern        pulumi.StringPtrInput                         `pulumi:"excludesPattern"`
	IncludesPattern        pulumi.StringPtrInput                         `pulumi:"includesPattern"`
	Key                    pulumi.StringInput                            `pulumi:"key"`
	Members                GetFederatedGenericRepositoryMemberArrayInput `pulumi:"members"`
	Notes                  pulumi.StringPtrInput                         `pulumi:"notes"`
	PriorityResolution     pulumi.BoolPtrInput                           `pulumi:"priorityResolution"`
	ProjectEnvironments    pulumi.StringArrayInput                       `pulumi:"projectEnvironments"`
	ProjectKey             pulumi.StringPtrInput                         `pulumi:"projectKey"`
	PropertySets           pulumi.StringArrayInput                       `pulumi:"propertySets"`
	RepoLayoutRef          pulumi.StringPtrInput                         `pulumi:"repoLayoutRef"`
	XrayIndex              pulumi.BoolPtrInput                           `pulumi:"xrayIndex"`
}

func (LookupFederatedGenericRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedGenericRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedGenericRepository.
type LookupFederatedGenericRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedGenericRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedGenericRepositoryResult)(nil)).Elem()
}

func (o LookupFederatedGenericRepositoryResultOutput) ToLookupFederatedGenericRepositoryResultOutput() LookupFederatedGenericRepositoryResultOutput {
	return o
}

func (o LookupFederatedGenericRepositoryResultOutput) ToLookupFederatedGenericRepositoryResultOutputWithContext(ctx context.Context) LookupFederatedGenericRepositoryResultOutput {
	return o
}

func (o LookupFederatedGenericRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedGenericRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedGenericRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedGenericRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedGenericRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedGenericRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedGenericRepositoryResultOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedGenericRepositoryResult) *bool { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedGenericRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedGenericRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedGenericRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedGenericRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedGenericRepositoryResultOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedGenericRepositoryResult) string { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFederatedGenericRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedGenericRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedGenericRepositoryResultOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedGenericRepositoryResult) string { return v.IncludesPattern }).(pulumi.StringOutput)
}

func (o LookupFederatedGenericRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedGenericRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupFederatedGenericRepositoryResultOutput) Members() GetFederatedGenericRepositoryMemberArrayOutput {
	return o.ApplyT(func(v LookupFederatedGenericRepositoryResult) []GetFederatedGenericRepositoryMember { return v.Members }).(GetFederatedGenericRepositoryMemberArrayOutput)
}

func (o LookupFederatedGenericRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedGenericRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedGenericRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedGenericRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupFederatedGenericRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedGenericRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedGenericRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedGenericRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedGenericRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedGenericRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedGenericRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedGenericRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedGenericRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedGenericRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedGenericRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedGenericRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedGenericRepositoryResultOutput{})
}
