// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a federated Puppet repository.
func LookupFederatedPuppetRepository(ctx *pulumi.Context, args *LookupFederatedPuppetRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupFederatedPuppetRepositoryResult, error) {
	var rv LookupFederatedPuppetRepositoryResult
	err := ctx.Invoke("artifactory:index/getFederatedPuppetRepository:getFederatedPuppetRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedPuppetRepository.
type LookupFederatedPuppetRepositoryArgs struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	IncludesPattern        *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             []GetFederatedPuppetRepositoryMember `pulumi:"members"`
	Notes               *string                              `pulumi:"notes"`
	PriorityResolution  *bool                                `pulumi:"priorityResolution"`
	ProjectEnvironments []string                             `pulumi:"projectEnvironments"`
	ProjectKey          *string                              `pulumi:"projectKey"`
	PropertySets        []string                             `pulumi:"propertySets"`
	RepoLayoutRef       *string                              `pulumi:"repoLayoutRef"`
	XrayIndex           *bool                                `pulumi:"xrayIndex"`
}

// A collection of values returned by getFederatedPuppetRepository.
type LookupFederatedPuppetRepositoryResult struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        string  `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id              string `pulumi:"id"`
	IncludesPattern string `pulumi:"includesPattern"`
	Key             string `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             []GetFederatedPuppetRepositoryMember `pulumi:"members"`
	Notes               *string                              `pulumi:"notes"`
	PackageType         string                               `pulumi:"packageType"`
	PriorityResolution  *bool                                `pulumi:"priorityResolution"`
	ProjectEnvironments []string                             `pulumi:"projectEnvironments"`
	ProjectKey          *string                              `pulumi:"projectKey"`
	PropertySets        []string                             `pulumi:"propertySets"`
	RepoLayoutRef       *string                              `pulumi:"repoLayoutRef"`
	XrayIndex           *bool                                `pulumi:"xrayIndex"`
}

func LookupFederatedPuppetRepositoryOutput(ctx *pulumi.Context, args LookupFederatedPuppetRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedPuppetRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedPuppetRepositoryResult, error) {
			args := v.(LookupFederatedPuppetRepositoryArgs)
			r, err := LookupFederatedPuppetRepository(ctx, &args, opts...)
			var s LookupFederatedPuppetRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFederatedPuppetRepositoryResultOutput)
}

// A collection of arguments for invoking getFederatedPuppetRepository.
type LookupFederatedPuppetRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput   `pulumi:"blackedOut"`
	CdnRedirect            pulumi.BoolPtrInput   `pulumi:"cdnRedirect"`
	CleanupOnDelete        pulumi.BoolPtrInput   `pulumi:"cleanupOnDelete"`
	Description            pulumi.StringPtrInput `pulumi:"description"`
	DownloadDirect         pulumi.BoolPtrInput   `pulumi:"downloadDirect"`
	ExcludesPattern        pulumi.StringPtrInput `pulumi:"excludesPattern"`
	IncludesPattern        pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key pulumi.StringInput `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             GetFederatedPuppetRepositoryMemberArrayInput `pulumi:"members"`
	Notes               pulumi.StringPtrInput                        `pulumi:"notes"`
	PriorityResolution  pulumi.BoolPtrInput                          `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayInput                      `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput                        `pulumi:"projectKey"`
	PropertySets        pulumi.StringArrayInput                      `pulumi:"propertySets"`
	RepoLayoutRef       pulumi.StringPtrInput                        `pulumi:"repoLayoutRef"`
	XrayIndex           pulumi.BoolPtrInput                          `pulumi:"xrayIndex"`
}

func (LookupFederatedPuppetRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedPuppetRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedPuppetRepository.
type LookupFederatedPuppetRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedPuppetRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedPuppetRepositoryResult)(nil)).Elem()
}

func (o LookupFederatedPuppetRepositoryResultOutput) ToLookupFederatedPuppetRepositoryResultOutput() LookupFederatedPuppetRepositoryResultOutput {
	return o
}

func (o LookupFederatedPuppetRepositoryResultOutput) ToLookupFederatedPuppetRepositoryResultOutputWithContext(ctx context.Context) LookupFederatedPuppetRepositoryResultOutput {
	return o
}

func (o LookupFederatedPuppetRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedPuppetRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedPuppetRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedPuppetRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedPuppetRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedPuppetRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedPuppetRepositoryResultOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedPuppetRepositoryResult) *bool { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedPuppetRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedPuppetRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedPuppetRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedPuppetRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedPuppetRepositoryResultOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedPuppetRepositoryResult) string { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFederatedPuppetRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedPuppetRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedPuppetRepositoryResultOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedPuppetRepositoryResult) string { return v.IncludesPattern }).(pulumi.StringOutput)
}

func (o LookupFederatedPuppetRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedPuppetRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o LookupFederatedPuppetRepositoryResultOutput) Members() GetFederatedPuppetRepositoryMemberArrayOutput {
	return o.ApplyT(func(v LookupFederatedPuppetRepositoryResult) []GetFederatedPuppetRepositoryMember { return v.Members }).(GetFederatedPuppetRepositoryMemberArrayOutput)
}

func (o LookupFederatedPuppetRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedPuppetRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedPuppetRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedPuppetRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupFederatedPuppetRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedPuppetRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedPuppetRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedPuppetRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedPuppetRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedPuppetRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedPuppetRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedPuppetRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedPuppetRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedPuppetRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedPuppetRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedPuppetRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedPuppetRepositoryResultOutput{})
}
