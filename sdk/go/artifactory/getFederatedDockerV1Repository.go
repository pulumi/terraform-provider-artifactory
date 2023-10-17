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

// Retrieves a federated Docker repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.LookupFederatedDockerV1Repository(ctx, &artifactory.LookupFederatedDockerV1RepositoryArgs{
//				Key: "federated-test-docker-repo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFederatedDockerV1Repository(ctx *pulumi.Context, args *LookupFederatedDockerV1RepositoryArgs, opts ...pulumi.InvokeOption) (*LookupFederatedDockerV1RepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFederatedDockerV1RepositoryResult
	err := ctx.Invoke("artifactory:index/getFederatedDockerV1Repository:getFederatedDockerV1Repository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedDockerV1Repository.
type LookupFederatedDockerV1RepositoryArgs struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	IncludesPattern        *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key           string `pulumi:"key"`
	MaxUniqueTags *int   `pulumi:"maxUniqueTags"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             []GetFederatedDockerV1RepositoryMember `pulumi:"members"`
	Notes               *string                                `pulumi:"notes"`
	PriorityResolution  *bool                                  `pulumi:"priorityResolution"`
	ProjectEnvironments []string                               `pulumi:"projectEnvironments"`
	ProjectKey          *string                                `pulumi:"projectKey"`
	PropertySets        []string                               `pulumi:"propertySets"`
	RepoLayoutRef       *string                                `pulumi:"repoLayoutRef"`
	XrayIndex           *bool                                  `pulumi:"xrayIndex"`
}

// A collection of values returned by getFederatedDockerV1Repository.
type LookupFederatedDockerV1RepositoryResult struct {
	ApiVersion             string  `pulumi:"apiVersion"`
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	BlockPushingSchema1    bool    `pulumi:"blockPushingSchema1"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id              string  `pulumi:"id"`
	IncludesPattern *string `pulumi:"includesPattern"`
	Key             string  `pulumi:"key"`
	MaxUniqueTags   int     `pulumi:"maxUniqueTags"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             []GetFederatedDockerV1RepositoryMember `pulumi:"members"`
	Notes               *string                                `pulumi:"notes"`
	PackageType         string                                 `pulumi:"packageType"`
	PriorityResolution  *bool                                  `pulumi:"priorityResolution"`
	ProjectEnvironments []string                               `pulumi:"projectEnvironments"`
	ProjectKey          *string                                `pulumi:"projectKey"`
	PropertySets        []string                               `pulumi:"propertySets"`
	RepoLayoutRef       *string                                `pulumi:"repoLayoutRef"`
	TagRetention        int                                    `pulumi:"tagRetention"`
	XrayIndex           *bool                                  `pulumi:"xrayIndex"`
}

func LookupFederatedDockerV1RepositoryOutput(ctx *pulumi.Context, args LookupFederatedDockerV1RepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedDockerV1RepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedDockerV1RepositoryResult, error) {
			args := v.(LookupFederatedDockerV1RepositoryArgs)
			r, err := LookupFederatedDockerV1Repository(ctx, &args, opts...)
			var s LookupFederatedDockerV1RepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFederatedDockerV1RepositoryResultOutput)
}

// A collection of arguments for invoking getFederatedDockerV1Repository.
type LookupFederatedDockerV1RepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput   `pulumi:"blackedOut"`
	CdnRedirect            pulumi.BoolPtrInput   `pulumi:"cdnRedirect"`
	CleanupOnDelete        pulumi.BoolPtrInput   `pulumi:"cleanupOnDelete"`
	Description            pulumi.StringPtrInput `pulumi:"description"`
	DownloadDirect         pulumi.BoolPtrInput   `pulumi:"downloadDirect"`
	ExcludesPattern        pulumi.StringPtrInput `pulumi:"excludesPattern"`
	IncludesPattern        pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key           pulumi.StringInput `pulumi:"key"`
	MaxUniqueTags pulumi.IntPtrInput `pulumi:"maxUniqueTags"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             GetFederatedDockerV1RepositoryMemberArrayInput `pulumi:"members"`
	Notes               pulumi.StringPtrInput                          `pulumi:"notes"`
	PriorityResolution  pulumi.BoolPtrInput                            `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayInput                        `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput                          `pulumi:"projectKey"`
	PropertySets        pulumi.StringArrayInput                        `pulumi:"propertySets"`
	RepoLayoutRef       pulumi.StringPtrInput                          `pulumi:"repoLayoutRef"`
	XrayIndex           pulumi.BoolPtrInput                            `pulumi:"xrayIndex"`
}

func (LookupFederatedDockerV1RepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedDockerV1RepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedDockerV1Repository.
type LookupFederatedDockerV1RepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedDockerV1RepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedDockerV1RepositoryResult)(nil)).Elem()
}

func (o LookupFederatedDockerV1RepositoryResultOutput) ToLookupFederatedDockerV1RepositoryResultOutput() LookupFederatedDockerV1RepositoryResultOutput {
	return o
}

func (o LookupFederatedDockerV1RepositoryResultOutput) ToLookupFederatedDockerV1RepositoryResultOutputWithContext(ctx context.Context) LookupFederatedDockerV1RepositoryResultOutput {
	return o
}

func (o LookupFederatedDockerV1RepositoryResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupFederatedDockerV1RepositoryResult] {
	return pulumix.Output[LookupFederatedDockerV1RepositoryResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupFederatedDockerV1RepositoryResultOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedDockerV1RepositoryResult) string { return v.ApiVersion }).(pulumi.StringOutput)
}

func (o LookupFederatedDockerV1RepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV1RepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedDockerV1RepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV1RepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedDockerV1RepositoryResultOutput) BlockPushingSchema1() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupFederatedDockerV1RepositoryResult) bool { return v.BlockPushingSchema1 }).(pulumi.BoolOutput)
}

func (o LookupFederatedDockerV1RepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV1RepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedDockerV1RepositoryResultOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV1RepositoryResult) *bool { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedDockerV1RepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV1RepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedDockerV1RepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV1RepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedDockerV1RepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV1RepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFederatedDockerV1RepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedDockerV1RepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedDockerV1RepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV1RepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedDockerV1RepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedDockerV1RepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupFederatedDockerV1RepositoryResultOutput) MaxUniqueTags() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFederatedDockerV1RepositoryResult) int { return v.MaxUniqueTags }).(pulumi.IntOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o LookupFederatedDockerV1RepositoryResultOutput) Members() GetFederatedDockerV1RepositoryMemberArrayOutput {
	return o.ApplyT(func(v LookupFederatedDockerV1RepositoryResult) []GetFederatedDockerV1RepositoryMember {
		return v.Members
	}).(GetFederatedDockerV1RepositoryMemberArrayOutput)
}

func (o LookupFederatedDockerV1RepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV1RepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedDockerV1RepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedDockerV1RepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupFederatedDockerV1RepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV1RepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedDockerV1RepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedDockerV1RepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedDockerV1RepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV1RepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedDockerV1RepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedDockerV1RepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedDockerV1RepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV1RepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedDockerV1RepositoryResultOutput) TagRetention() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFederatedDockerV1RepositoryResult) int { return v.TagRetention }).(pulumi.IntOutput)
}

func (o LookupFederatedDockerV1RepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV1RepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedDockerV1RepositoryResultOutput{})
}
