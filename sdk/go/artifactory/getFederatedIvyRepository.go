// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v4/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a federated Ivy repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v4/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.LookupFederatedIvyRepository(ctx, &artifactory.LookupFederatedIvyRepositoryArgs{
//				Key: "federated-test-ivy-repo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFederatedIvyRepository(ctx *pulumi.Context, args *LookupFederatedIvyRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupFederatedIvyRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFederatedIvyRepositoryResult
	err := ctx.Invoke("artifactory:index/getFederatedIvyRepository:getFederatedIvyRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedIvyRepository.
type LookupFederatedIvyRepositoryArgs struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	ChecksumPolicyType     *string `pulumi:"checksumPolicyType"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	HandleReleases         *bool   `pulumi:"handleReleases"`
	HandleSnapshots        *bool   `pulumi:"handleSnapshots"`
	IncludesPattern        *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                string `pulumi:"key"`
	MaxUniqueSnapshots *int   `pulumi:"maxUniqueSnapshots"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members                      []GetFederatedIvyRepositoryMember `pulumi:"members"`
	Notes                        *string                           `pulumi:"notes"`
	PriorityResolution           *bool                             `pulumi:"priorityResolution"`
	ProjectEnvironments          []string                          `pulumi:"projectEnvironments"`
	ProjectKey                   *string                           `pulumi:"projectKey"`
	PropertySets                 []string                          `pulumi:"propertySets"`
	RepoLayoutRef                *string                           `pulumi:"repoLayoutRef"`
	SnapshotVersionBehavior      *string                           `pulumi:"snapshotVersionBehavior"`
	SuppressPomConsistencyChecks *bool                             `pulumi:"suppressPomConsistencyChecks"`
	XrayIndex                    *bool                             `pulumi:"xrayIndex"`
}

// A collection of values returned by getFederatedIvyRepository.
type LookupFederatedIvyRepositoryResult struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	ChecksumPolicyType     *string `pulumi:"checksumPolicyType"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        string  `pulumi:"excludesPattern"`
	HandleReleases         *bool   `pulumi:"handleReleases"`
	HandleSnapshots        *bool   `pulumi:"handleSnapshots"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string `pulumi:"id"`
	IncludesPattern    string `pulumi:"includesPattern"`
	Key                string `pulumi:"key"`
	MaxUniqueSnapshots *int   `pulumi:"maxUniqueSnapshots"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members                      []GetFederatedIvyRepositoryMember `pulumi:"members"`
	Notes                        *string                           `pulumi:"notes"`
	PackageType                  string                            `pulumi:"packageType"`
	PriorityResolution           *bool                             `pulumi:"priorityResolution"`
	ProjectEnvironments          []string                          `pulumi:"projectEnvironments"`
	ProjectKey                   *string                           `pulumi:"projectKey"`
	PropertySets                 []string                          `pulumi:"propertySets"`
	RepoLayoutRef                *string                           `pulumi:"repoLayoutRef"`
	SnapshotVersionBehavior      *string                           `pulumi:"snapshotVersionBehavior"`
	SuppressPomConsistencyChecks *bool                             `pulumi:"suppressPomConsistencyChecks"`
	XrayIndex                    *bool                             `pulumi:"xrayIndex"`
}

func LookupFederatedIvyRepositoryOutput(ctx *pulumi.Context, args LookupFederatedIvyRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedIvyRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedIvyRepositoryResult, error) {
			args := v.(LookupFederatedIvyRepositoryArgs)
			r, err := LookupFederatedIvyRepository(ctx, &args, opts...)
			var s LookupFederatedIvyRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFederatedIvyRepositoryResultOutput)
}

// A collection of arguments for invoking getFederatedIvyRepository.
type LookupFederatedIvyRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput   `pulumi:"blackedOut"`
	CdnRedirect            pulumi.BoolPtrInput   `pulumi:"cdnRedirect"`
	ChecksumPolicyType     pulumi.StringPtrInput `pulumi:"checksumPolicyType"`
	CleanupOnDelete        pulumi.BoolPtrInput   `pulumi:"cleanupOnDelete"`
	Description            pulumi.StringPtrInput `pulumi:"description"`
	DownloadDirect         pulumi.BoolPtrInput   `pulumi:"downloadDirect"`
	ExcludesPattern        pulumi.StringPtrInput `pulumi:"excludesPattern"`
	HandleReleases         pulumi.BoolPtrInput   `pulumi:"handleReleases"`
	HandleSnapshots        pulumi.BoolPtrInput   `pulumi:"handleSnapshots"`
	IncludesPattern        pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                pulumi.StringInput `pulumi:"key"`
	MaxUniqueSnapshots pulumi.IntPtrInput `pulumi:"maxUniqueSnapshots"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members                      GetFederatedIvyRepositoryMemberArrayInput `pulumi:"members"`
	Notes                        pulumi.StringPtrInput                     `pulumi:"notes"`
	PriorityResolution           pulumi.BoolPtrInput                       `pulumi:"priorityResolution"`
	ProjectEnvironments          pulumi.StringArrayInput                   `pulumi:"projectEnvironments"`
	ProjectKey                   pulumi.StringPtrInput                     `pulumi:"projectKey"`
	PropertySets                 pulumi.StringArrayInput                   `pulumi:"propertySets"`
	RepoLayoutRef                pulumi.StringPtrInput                     `pulumi:"repoLayoutRef"`
	SnapshotVersionBehavior      pulumi.StringPtrInput                     `pulumi:"snapshotVersionBehavior"`
	SuppressPomConsistencyChecks pulumi.BoolPtrInput                       `pulumi:"suppressPomConsistencyChecks"`
	XrayIndex                    pulumi.BoolPtrInput                       `pulumi:"xrayIndex"`
}

func (LookupFederatedIvyRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedIvyRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedIvyRepository.
type LookupFederatedIvyRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedIvyRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedIvyRepositoryResult)(nil)).Elem()
}

func (o LookupFederatedIvyRepositoryResultOutput) ToLookupFederatedIvyRepositoryResultOutput() LookupFederatedIvyRepositoryResultOutput {
	return o
}

func (o LookupFederatedIvyRepositoryResultOutput) ToLookupFederatedIvyRepositoryResultOutputWithContext(ctx context.Context) LookupFederatedIvyRepositoryResultOutput {
	return o
}

func (o LookupFederatedIvyRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedIvyRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedIvyRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedIvyRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedIvyRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedIvyRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedIvyRepositoryResultOutput) ChecksumPolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedIvyRepositoryResult) *string { return v.ChecksumPolicyType }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedIvyRepositoryResultOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedIvyRepositoryResult) *bool { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedIvyRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedIvyRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedIvyRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedIvyRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedIvyRepositoryResultOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedIvyRepositoryResult) string { return v.ExcludesPattern }).(pulumi.StringOutput)
}

func (o LookupFederatedIvyRepositoryResultOutput) HandleReleases() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedIvyRepositoryResult) *bool { return v.HandleReleases }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedIvyRepositoryResultOutput) HandleSnapshots() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedIvyRepositoryResult) *bool { return v.HandleSnapshots }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFederatedIvyRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedIvyRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedIvyRepositoryResultOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedIvyRepositoryResult) string { return v.IncludesPattern }).(pulumi.StringOutput)
}

func (o LookupFederatedIvyRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedIvyRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupFederatedIvyRepositoryResultOutput) MaxUniqueSnapshots() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFederatedIvyRepositoryResult) *int { return v.MaxUniqueSnapshots }).(pulumi.IntPtrOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o LookupFederatedIvyRepositoryResultOutput) Members() GetFederatedIvyRepositoryMemberArrayOutput {
	return o.ApplyT(func(v LookupFederatedIvyRepositoryResult) []GetFederatedIvyRepositoryMember { return v.Members }).(GetFederatedIvyRepositoryMemberArrayOutput)
}

func (o LookupFederatedIvyRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedIvyRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedIvyRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedIvyRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupFederatedIvyRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedIvyRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedIvyRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedIvyRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedIvyRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedIvyRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedIvyRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedIvyRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedIvyRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedIvyRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedIvyRepositoryResultOutput) SnapshotVersionBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedIvyRepositoryResult) *string { return v.SnapshotVersionBehavior }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedIvyRepositoryResultOutput) SuppressPomConsistencyChecks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedIvyRepositoryResult) *bool { return v.SuppressPomConsistencyChecks }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedIvyRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedIvyRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedIvyRepositoryResultOutput{})
}
