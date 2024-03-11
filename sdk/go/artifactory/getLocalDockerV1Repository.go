// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v6/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a local Docker (v1) repository resource.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v6/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewDockerV1Repository(ctx, "artifactoryLocalTestDockerV1Repository", &artifactory.DockerV1RepositoryArgs{
//				Key: pulumi.String("artifactory_local_test_docker_v1_repository"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetLocalDockerV1Repository(ctx *pulumi.Context, args *GetLocalDockerV1RepositoryArgs, opts ...pulumi.InvokeOption) (*GetLocalDockerV1RepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLocalDockerV1RepositoryResult
	err := ctx.Invoke("artifactory:index/getLocalDockerV1Repository:getLocalDockerV1Repository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocalDockerV1Repository.
type GetLocalDockerV1RepositoryArgs struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	IncludesPattern        *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// The maximum number of unique tags of a single Docker image to store in this repository. Once the
	// number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no
	// limit. This only applies to manifest v2.
	MaxUniqueTags       *int     `pulumi:"maxUniqueTags"`
	Notes               *string  `pulumi:"notes"`
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	PropertySets        []string `pulumi:"propertySets"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	XrayIndex           *bool    `pulumi:"xrayIndex"`
}

// A collection of values returned by getLocalDockerV1Repository.
type GetLocalDockerV1RepositoryResult struct {
	// The Docker API version in use.
	ApiVersion             string `pulumi:"apiVersion"`
	ArchiveBrowsingEnabled *bool  `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool  `pulumi:"blackedOut"`
	// When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to
	// this repository.
	BlockPushingSchema1 bool    `pulumi:"blockPushingSchema1"`
	CdnRedirect         *bool   `pulumi:"cdnRedirect"`
	Description         *string `pulumi:"description"`
	DownloadDirect      *bool   `pulumi:"downloadDirect"`
	ExcludesPattern     *string `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id              string  `pulumi:"id"`
	IncludesPattern *string `pulumi:"includesPattern"`
	Key             string  `pulumi:"key"`
	// The maximum number of unique tags of a single Docker image to store in this repository. Once the
	// number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no
	// limit. This only applies to manifest v2.
	MaxUniqueTags       int      `pulumi:"maxUniqueTags"`
	Notes               *string  `pulumi:"notes"`
	PackageType         string   `pulumi:"packageType"`
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	PropertySets        []string `pulumi:"propertySets"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	// If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This
	// only applies to manifest V2.
	TagRetention int   `pulumi:"tagRetention"`
	XrayIndex    *bool `pulumi:"xrayIndex"`
}

func GetLocalDockerV1RepositoryOutput(ctx *pulumi.Context, args GetLocalDockerV1RepositoryOutputArgs, opts ...pulumi.InvokeOption) GetLocalDockerV1RepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLocalDockerV1RepositoryResult, error) {
			args := v.(GetLocalDockerV1RepositoryArgs)
			r, err := GetLocalDockerV1Repository(ctx, &args, opts...)
			var s GetLocalDockerV1RepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLocalDockerV1RepositoryResultOutput)
}

// A collection of arguments for invoking getLocalDockerV1Repository.
type GetLocalDockerV1RepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput   `pulumi:"blackedOut"`
	CdnRedirect            pulumi.BoolPtrInput   `pulumi:"cdnRedirect"`
	Description            pulumi.StringPtrInput `pulumi:"description"`
	DownloadDirect         pulumi.BoolPtrInput   `pulumi:"downloadDirect"`
	ExcludesPattern        pulumi.StringPtrInput `pulumi:"excludesPattern"`
	IncludesPattern        pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key pulumi.StringInput `pulumi:"key"`
	// The maximum number of unique tags of a single Docker image to store in this repository. Once the
	// number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no
	// limit. This only applies to manifest v2.
	MaxUniqueTags       pulumi.IntPtrInput      `pulumi:"maxUniqueTags"`
	Notes               pulumi.StringPtrInput   `pulumi:"notes"`
	PriorityResolution  pulumi.BoolPtrInput     `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput   `pulumi:"projectKey"`
	PropertySets        pulumi.StringArrayInput `pulumi:"propertySets"`
	RepoLayoutRef       pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	XrayIndex           pulumi.BoolPtrInput     `pulumi:"xrayIndex"`
}

func (GetLocalDockerV1RepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLocalDockerV1RepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getLocalDockerV1Repository.
type GetLocalDockerV1RepositoryResultOutput struct{ *pulumi.OutputState }

func (GetLocalDockerV1RepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLocalDockerV1RepositoryResult)(nil)).Elem()
}

func (o GetLocalDockerV1RepositoryResultOutput) ToGetLocalDockerV1RepositoryResultOutput() GetLocalDockerV1RepositoryResultOutput {
	return o
}

func (o GetLocalDockerV1RepositoryResultOutput) ToGetLocalDockerV1RepositoryResultOutputWithContext(ctx context.Context) GetLocalDockerV1RepositoryResultOutput {
	return o
}

// The Docker API version in use.
func (o GetLocalDockerV1RepositoryResultOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalDockerV1RepositoryResult) string { return v.ApiVersion }).(pulumi.StringOutput)
}

func (o GetLocalDockerV1RepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalDockerV1RepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o GetLocalDockerV1RepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalDockerV1RepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to
// this repository.
func (o GetLocalDockerV1RepositoryResultOutput) BlockPushingSchema1() pulumi.BoolOutput {
	return o.ApplyT(func(v GetLocalDockerV1RepositoryResult) bool { return v.BlockPushingSchema1 }).(pulumi.BoolOutput)
}

func (o GetLocalDockerV1RepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalDockerV1RepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o GetLocalDockerV1RepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalDockerV1RepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GetLocalDockerV1RepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalDockerV1RepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o GetLocalDockerV1RepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalDockerV1RepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLocalDockerV1RepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalDockerV1RepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLocalDockerV1RepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalDockerV1RepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o GetLocalDockerV1RepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalDockerV1RepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

// The maximum number of unique tags of a single Docker image to store in this repository. Once the
// number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no
// limit. This only applies to manifest v2.
func (o GetLocalDockerV1RepositoryResultOutput) MaxUniqueTags() pulumi.IntOutput {
	return o.ApplyT(func(v GetLocalDockerV1RepositoryResult) int { return v.MaxUniqueTags }).(pulumi.IntOutput)
}

func (o GetLocalDockerV1RepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalDockerV1RepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o GetLocalDockerV1RepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalDockerV1RepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o GetLocalDockerV1RepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalDockerV1RepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o GetLocalDockerV1RepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLocalDockerV1RepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o GetLocalDockerV1RepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalDockerV1RepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o GetLocalDockerV1RepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLocalDockerV1RepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o GetLocalDockerV1RepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalDockerV1RepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This
// only applies to manifest V2.
func (o GetLocalDockerV1RepositoryResultOutput) TagRetention() pulumi.IntOutput {
	return o.ApplyT(func(v GetLocalDockerV1RepositoryResult) int { return v.TagRetention }).(pulumi.IntOutput)
}

func (o GetLocalDockerV1RepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalDockerV1RepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLocalDockerV1RepositoryResultOutput{})
}
