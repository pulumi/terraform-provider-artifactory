// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v4/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Retrieves a virtual Debian repository.
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
//			_, err := artifactory.LookupVirtualDebianRepository(ctx, &artifactory.LookupVirtualDebianRepositoryArgs{
//				Key: "virtual-debian",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupVirtualDebianRepository(ctx *pulumi.Context, args *LookupVirtualDebianRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupVirtualDebianRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualDebianRepositoryResult
	err := ctx.Invoke("artifactory:index/getVirtualDebianRepository:getVirtualDebianRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualDebianRepository.
type LookupVirtualDebianRepositoryArgs struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// (Optional) Specifying  architectures will speed up Artifactory's initial metadata indexing process. The default architecture values are amd64 and i386.
	DebianDefaultArchitectures *string `pulumi:"debianDefaultArchitectures"`
	DefaultDeploymentRepo      *string `pulumi:"defaultDeploymentRepo"`
	Description                *string `pulumi:"description"`
	ExcludesPattern            *string `pulumi:"excludesPattern"`
	IncludesPattern            *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key   string  `pulumi:"key"`
	Notes *string `pulumi:"notes"`
	// (Optional) Index file formats you would like to create in addition to the default Gzip (.gzip extension). Supported values are `bz2`,`lzma` and `xz`. Default value is `bz2`.
	OptionalIndexCompressionFormats []string `pulumi:"optionalIndexCompressionFormats"`
	// (Optional) Primary keypair used to sign artifacts. Default is empty.
	PrimaryKeypairRef   *string  `pulumi:"primaryKeypairRef"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	Repositories        []string `pulumi:"repositories"`
	// (Optional, Default: `7200`) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds *int `pulumi:"retrievalCachePeriodSeconds"`
	// (Optional) Secondary keypair used to sign artifacts. Default is empty.
	SecondaryKeypairRef *string `pulumi:"secondaryKeypairRef"`
}

// A collection of values returned by getVirtualDebianRepository.
type LookupVirtualDebianRepositoryResult struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// (Optional) Specifying  architectures will speed up Artifactory's initial metadata indexing process. The default architecture values are amd64 and i386.
	DebianDefaultArchitectures *string `pulumi:"debianDefaultArchitectures"`
	DefaultDeploymentRepo      *string `pulumi:"defaultDeploymentRepo"`
	Description                *string `pulumi:"description"`
	ExcludesPattern            *string `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id              string  `pulumi:"id"`
	IncludesPattern *string `pulumi:"includesPattern"`
	Key             string  `pulumi:"key"`
	Notes           *string `pulumi:"notes"`
	// (Optional) Index file formats you would like to create in addition to the default Gzip (.gzip extension). Supported values are `bz2`,`lzma` and `xz`. Default value is `bz2`.
	OptionalIndexCompressionFormats []string `pulumi:"optionalIndexCompressionFormats"`
	PackageType                     string   `pulumi:"packageType"`
	// (Optional) Primary keypair used to sign artifacts. Default is empty.
	PrimaryKeypairRef   *string  `pulumi:"primaryKeypairRef"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	Repositories        []string `pulumi:"repositories"`
	// (Optional, Default: `7200`) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds *int `pulumi:"retrievalCachePeriodSeconds"`
	// (Optional) Secondary keypair used to sign artifacts. Default is empty.
	SecondaryKeypairRef *string `pulumi:"secondaryKeypairRef"`
}

func LookupVirtualDebianRepositoryOutput(ctx *pulumi.Context, args LookupVirtualDebianRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualDebianRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualDebianRepositoryResult, error) {
			args := v.(LookupVirtualDebianRepositoryArgs)
			r, err := LookupVirtualDebianRepository(ctx, &args, opts...)
			var s LookupVirtualDebianRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualDebianRepositoryResultOutput)
}

// A collection of arguments for invoking getVirtualDebianRepository.
type LookupVirtualDebianRepositoryOutputArgs struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// (Optional) Specifying  architectures will speed up Artifactory's initial metadata indexing process. The default architecture values are amd64 and i386.
	DebianDefaultArchitectures pulumi.StringPtrInput `pulumi:"debianDefaultArchitectures"`
	DefaultDeploymentRepo      pulumi.StringPtrInput `pulumi:"defaultDeploymentRepo"`
	Description                pulumi.StringPtrInput `pulumi:"description"`
	ExcludesPattern            pulumi.StringPtrInput `pulumi:"excludesPattern"`
	IncludesPattern            pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key   pulumi.StringInput    `pulumi:"key"`
	Notes pulumi.StringPtrInput `pulumi:"notes"`
	// (Optional) Index file formats you would like to create in addition to the default Gzip (.gzip extension). Supported values are `bz2`,`lzma` and `xz`. Default value is `bz2`.
	OptionalIndexCompressionFormats pulumi.StringArrayInput `pulumi:"optionalIndexCompressionFormats"`
	// (Optional) Primary keypair used to sign artifacts. Default is empty.
	PrimaryKeypairRef   pulumi.StringPtrInput   `pulumi:"primaryKeypairRef"`
	ProjectEnvironments pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput   `pulumi:"projectKey"`
	RepoLayoutRef       pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	Repositories        pulumi.StringArrayInput `pulumi:"repositories"`
	// (Optional, Default: `7200`) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds pulumi.IntPtrInput `pulumi:"retrievalCachePeriodSeconds"`
	// (Optional) Secondary keypair used to sign artifacts. Default is empty.
	SecondaryKeypairRef pulumi.StringPtrInput `pulumi:"secondaryKeypairRef"`
}

func (LookupVirtualDebianRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualDebianRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getVirtualDebianRepository.
type LookupVirtualDebianRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualDebianRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualDebianRepositoryResult)(nil)).Elem()
}

func (o LookupVirtualDebianRepositoryResultOutput) ToLookupVirtualDebianRepositoryResultOutput() LookupVirtualDebianRepositoryResultOutput {
	return o
}

func (o LookupVirtualDebianRepositoryResultOutput) ToLookupVirtualDebianRepositoryResultOutputWithContext(ctx context.Context) LookupVirtualDebianRepositoryResultOutput {
	return o
}

func (o LookupVirtualDebianRepositoryResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVirtualDebianRepositoryResult] {
	return pulumix.Output[LookupVirtualDebianRepositoryResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupVirtualDebianRepositoryResultOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualDebianRepositoryResult) *bool {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

// (Optional) Specifying  architectures will speed up Artifactory's initial metadata indexing process. The default architecture values are amd64 and i386.
func (o LookupVirtualDebianRepositoryResultOutput) DebianDefaultArchitectures() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualDebianRepositoryResult) *string { return v.DebianDefaultArchitectures }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualDebianRepositoryResultOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualDebianRepositoryResult) *string { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualDebianRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualDebianRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualDebianRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualDebianRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVirtualDebianRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualDebianRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualDebianRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualDebianRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualDebianRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualDebianRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupVirtualDebianRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualDebianRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

// (Optional) Index file formats you would like to create in addition to the default Gzip (.gzip extension). Supported values are `bz2`,`lzma` and `xz`. Default value is `bz2`.
func (o LookupVirtualDebianRepositoryResultOutput) OptionalIndexCompressionFormats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualDebianRepositoryResult) []string { return v.OptionalIndexCompressionFormats }).(pulumi.StringArrayOutput)
}

func (o LookupVirtualDebianRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualDebianRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

// (Optional) Primary keypair used to sign artifacts. Default is empty.
func (o LookupVirtualDebianRepositoryResultOutput) PrimaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualDebianRepositoryResult) *string { return v.PrimaryKeypairRef }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualDebianRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualDebianRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupVirtualDebianRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualDebianRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualDebianRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualDebianRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualDebianRepositoryResultOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualDebianRepositoryResult) []string { return v.Repositories }).(pulumi.StringArrayOutput)
}

// (Optional, Default: `7200`) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
func (o LookupVirtualDebianRepositoryResultOutput) RetrievalCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualDebianRepositoryResult) *int { return v.RetrievalCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

// (Optional) Secondary keypair used to sign artifacts. Default is empty.
func (o LookupVirtualDebianRepositoryResultOutput) SecondaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualDebianRepositoryResult) *string { return v.SecondaryKeypairRef }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualDebianRepositoryResultOutput{})
}
