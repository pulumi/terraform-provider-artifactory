// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a federated SBT repository.
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
//			_, err := artifactory.LookupFederatedSbtRepository(ctx, &artifactory.LookupFederatedSbtRepositoryArgs{
//				Key: "federated-test-sbt-repo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFederatedSbtRepository(ctx *pulumi.Context, args *LookupFederatedSbtRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupFederatedSbtRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFederatedSbtRepositoryResult
	err := ctx.Invoke("artifactory:index/getFederatedSbtRepository:getFederatedSbtRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedSbtRepository.
type LookupFederatedSbtRepositoryArgs struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	ChecksumPolicyType     *string `pulumi:"checksumPolicyType"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
	DisableProxy    *bool   `pulumi:"disableProxy"`
	DownloadDirect  *bool   `pulumi:"downloadDirect"`
	ExcludesPattern *string `pulumi:"excludesPattern"`
	HandleReleases  *bool   `pulumi:"handleReleases"`
	HandleSnapshots *bool   `pulumi:"handleSnapshots"`
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                string `pulumi:"key"`
	MaxUniqueSnapshots *int   `pulumi:"maxUniqueSnapshots"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             []GetFederatedSbtRepositoryMember `pulumi:"members"`
	Notes               *string                           `pulumi:"notes"`
	PriorityResolution  *bool                             `pulumi:"priorityResolution"`
	ProjectEnvironments []string                          `pulumi:"projectEnvironments"`
	ProjectKey          *string                           `pulumi:"projectKey"`
	PropertySets        []string                          `pulumi:"propertySets"`
	// Proxy key from Artifactory Proxies settings.
	Proxy                        *string `pulumi:"proxy"`
	RepoLayoutRef                *string `pulumi:"repoLayoutRef"`
	SnapshotVersionBehavior      *string `pulumi:"snapshotVersionBehavior"`
	SuppressPomConsistencyChecks *bool   `pulumi:"suppressPomConsistencyChecks"`
	XrayIndex                    *bool   `pulumi:"xrayIndex"`
}

// A collection of values returned by getFederatedSbtRepository.
type LookupFederatedSbtRepositoryResult struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	ChecksumPolicyType     *string `pulumi:"checksumPolicyType"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
	DisableProxy    *bool   `pulumi:"disableProxy"`
	DownloadDirect  *bool   `pulumi:"downloadDirect"`
	ExcludesPattern *string `pulumi:"excludesPattern"`
	HandleReleases  *bool   `pulumi:"handleReleases"`
	HandleSnapshots *bool   `pulumi:"handleSnapshots"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string  `pulumi:"id"`
	IncludesPattern    *string `pulumi:"includesPattern"`
	Key                string  `pulumi:"key"`
	MaxUniqueSnapshots *int    `pulumi:"maxUniqueSnapshots"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             []GetFederatedSbtRepositoryMember `pulumi:"members"`
	Notes               *string                           `pulumi:"notes"`
	PackageType         string                            `pulumi:"packageType"`
	PriorityResolution  *bool                             `pulumi:"priorityResolution"`
	ProjectEnvironments []string                          `pulumi:"projectEnvironments"`
	ProjectKey          *string                           `pulumi:"projectKey"`
	PropertySets        []string                          `pulumi:"propertySets"`
	// Proxy key from Artifactory Proxies settings.
	Proxy                        *string `pulumi:"proxy"`
	RepoLayoutRef                *string `pulumi:"repoLayoutRef"`
	SnapshotVersionBehavior      *string `pulumi:"snapshotVersionBehavior"`
	SuppressPomConsistencyChecks *bool   `pulumi:"suppressPomConsistencyChecks"`
	XrayIndex                    *bool   `pulumi:"xrayIndex"`
}

func LookupFederatedSbtRepositoryOutput(ctx *pulumi.Context, args LookupFederatedSbtRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedSbtRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedSbtRepositoryResult, error) {
			args := v.(LookupFederatedSbtRepositoryArgs)
			r, err := LookupFederatedSbtRepository(ctx, &args, opts...)
			var s LookupFederatedSbtRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFederatedSbtRepositoryResultOutput)
}

// A collection of arguments for invoking getFederatedSbtRepository.
type LookupFederatedSbtRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput   `pulumi:"blackedOut"`
	CdnRedirect            pulumi.BoolPtrInput   `pulumi:"cdnRedirect"`
	ChecksumPolicyType     pulumi.StringPtrInput `pulumi:"checksumPolicyType"`
	CleanupOnDelete        pulumi.BoolPtrInput   `pulumi:"cleanupOnDelete"`
	Description            pulumi.StringPtrInput `pulumi:"description"`
	// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
	DisableProxy    pulumi.BoolPtrInput   `pulumi:"disableProxy"`
	DownloadDirect  pulumi.BoolPtrInput   `pulumi:"downloadDirect"`
	ExcludesPattern pulumi.StringPtrInput `pulumi:"excludesPattern"`
	HandleReleases  pulumi.BoolPtrInput   `pulumi:"handleReleases"`
	HandleSnapshots pulumi.BoolPtrInput   `pulumi:"handleSnapshots"`
	IncludesPattern pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                pulumi.StringInput `pulumi:"key"`
	MaxUniqueSnapshots pulumi.IntPtrInput `pulumi:"maxUniqueSnapshots"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             GetFederatedSbtRepositoryMemberArrayInput `pulumi:"members"`
	Notes               pulumi.StringPtrInput                     `pulumi:"notes"`
	PriorityResolution  pulumi.BoolPtrInput                       `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayInput                   `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput                     `pulumi:"projectKey"`
	PropertySets        pulumi.StringArrayInput                   `pulumi:"propertySets"`
	// Proxy key from Artifactory Proxies settings.
	Proxy                        pulumi.StringPtrInput `pulumi:"proxy"`
	RepoLayoutRef                pulumi.StringPtrInput `pulumi:"repoLayoutRef"`
	SnapshotVersionBehavior      pulumi.StringPtrInput `pulumi:"snapshotVersionBehavior"`
	SuppressPomConsistencyChecks pulumi.BoolPtrInput   `pulumi:"suppressPomConsistencyChecks"`
	XrayIndex                    pulumi.BoolPtrInput   `pulumi:"xrayIndex"`
}

func (LookupFederatedSbtRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedSbtRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedSbtRepository.
type LookupFederatedSbtRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedSbtRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedSbtRepositoryResult)(nil)).Elem()
}

func (o LookupFederatedSbtRepositoryResultOutput) ToLookupFederatedSbtRepositoryResultOutput() LookupFederatedSbtRepositoryResultOutput {
	return o
}

func (o LookupFederatedSbtRepositoryResultOutput) ToLookupFederatedSbtRepositoryResultOutputWithContext(ctx context.Context) LookupFederatedSbtRepositoryResultOutput {
	return o
}

func (o LookupFederatedSbtRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedSbtRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedSbtRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedSbtRepositoryResultOutput) ChecksumPolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) *string { return v.ChecksumPolicyType }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedSbtRepositoryResultOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) *bool { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedSbtRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
func (o LookupFederatedSbtRepositoryResultOutput) DisableProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) *bool { return v.DisableProxy }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedSbtRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedSbtRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedSbtRepositoryResultOutput) HandleReleases() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) *bool { return v.HandleReleases }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedSbtRepositoryResultOutput) HandleSnapshots() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) *bool { return v.HandleSnapshots }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFederatedSbtRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedSbtRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedSbtRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupFederatedSbtRepositoryResultOutput) MaxUniqueSnapshots() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) *int { return v.MaxUniqueSnapshots }).(pulumi.IntPtrOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o LookupFederatedSbtRepositoryResultOutput) Members() GetFederatedSbtRepositoryMemberArrayOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) []GetFederatedSbtRepositoryMember { return v.Members }).(GetFederatedSbtRepositoryMemberArrayOutput)
}

func (o LookupFederatedSbtRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedSbtRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupFederatedSbtRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedSbtRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedSbtRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedSbtRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Proxy key from Artifactory Proxies settings.
func (o LookupFederatedSbtRepositoryResultOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) *string { return v.Proxy }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedSbtRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedSbtRepositoryResultOutput) SnapshotVersionBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) *string { return v.SnapshotVersionBehavior }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedSbtRepositoryResultOutput) SuppressPomConsistencyChecks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) *bool { return v.SuppressPomConsistencyChecks }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedSbtRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedSbtRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedSbtRepositoryResultOutput{})
}
