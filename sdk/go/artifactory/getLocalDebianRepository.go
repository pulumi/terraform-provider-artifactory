// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a local Debian repository.
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
//			_, err := artifactory.GetLocalDebianRepository(ctx, &artifactory.GetLocalDebianRepositoryArgs{
//				Key: "local-test-debian-repo-basic",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetLocalDebianRepository(ctx *pulumi.Context, args *GetLocalDebianRepositoryArgs, opts ...pulumi.InvokeOption) (*GetLocalDebianRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLocalDebianRepositoryResult
	err := ctx.Invoke("artifactory:index/getLocalDebianRepository:getLocalDebianRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocalDebianRepository.
type GetLocalDebianRepositoryArgs struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	IncludesPattern        *string `pulumi:"includesPattern"`
	// The options are Bzip2 (.bz2 extension) (default), LZMA (.lzma extension)
	// and XZ (.xz extension).
	IndexCompressionFormats []string `pulumi:"indexCompressionFormats"`
	// the identity key of the repo.
	Key   string  `pulumi:"key"`
	Notes *string `pulumi:"notes"`
	// The primary RSA key to be used to sign packages.
	PrimaryKeypairRef   *string  `pulumi:"primaryKeypairRef"`
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	PropertySets        []string `pulumi:"propertySets"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	// The secondary RSA key to be used to sign packages.
	SecondaryKeypairRef *string `pulumi:"secondaryKeypairRef"`
	// When set, the repository will use the deprecated trivial layout.
	//
	// Deprecated: You shouldn't be using this
	TrivialLayout *bool `pulumi:"trivialLayout"`
	XrayIndex     *bool `pulumi:"xrayIndex"`
}

// A collection of values returned by getLocalDebianRepository.
type GetLocalDebianRepositoryResult struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id              string  `pulumi:"id"`
	IncludesPattern *string `pulumi:"includesPattern"`
	// The options are Bzip2 (.bz2 extension) (default), LZMA (.lzma extension)
	// and XZ (.xz extension).
	IndexCompressionFormats []string `pulumi:"indexCompressionFormats"`
	Key                     string   `pulumi:"key"`
	Notes                   *string  `pulumi:"notes"`
	PackageType             string   `pulumi:"packageType"`
	// The primary RSA key to be used to sign packages.
	PrimaryKeypairRef   *string  `pulumi:"primaryKeypairRef"`
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	PropertySets        []string `pulumi:"propertySets"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	// The secondary RSA key to be used to sign packages.
	SecondaryKeypairRef *string `pulumi:"secondaryKeypairRef"`
	// When set, the repository will use the deprecated trivial layout.
	//
	// Deprecated: You shouldn't be using this
	TrivialLayout *bool `pulumi:"trivialLayout"`
	XrayIndex     *bool `pulumi:"xrayIndex"`
}

func GetLocalDebianRepositoryOutput(ctx *pulumi.Context, args GetLocalDebianRepositoryOutputArgs, opts ...pulumi.InvokeOption) GetLocalDebianRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLocalDebianRepositoryResultOutput, error) {
			args := v.(GetLocalDebianRepositoryArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetLocalDebianRepositoryResult
			secret, err := ctx.InvokePackageRaw("artifactory:index/getLocalDebianRepository:getLocalDebianRepository", args, &rv, "", opts...)
			if err != nil {
				return GetLocalDebianRepositoryResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetLocalDebianRepositoryResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetLocalDebianRepositoryResultOutput), nil
			}
			return output, nil
		}).(GetLocalDebianRepositoryResultOutput)
}

// A collection of arguments for invoking getLocalDebianRepository.
type GetLocalDebianRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput   `pulumi:"blackedOut"`
	CdnRedirect            pulumi.BoolPtrInput   `pulumi:"cdnRedirect"`
	Description            pulumi.StringPtrInput `pulumi:"description"`
	DownloadDirect         pulumi.BoolPtrInput   `pulumi:"downloadDirect"`
	ExcludesPattern        pulumi.StringPtrInput `pulumi:"excludesPattern"`
	IncludesPattern        pulumi.StringPtrInput `pulumi:"includesPattern"`
	// The options are Bzip2 (.bz2 extension) (default), LZMA (.lzma extension)
	// and XZ (.xz extension).
	IndexCompressionFormats pulumi.StringArrayInput `pulumi:"indexCompressionFormats"`
	// the identity key of the repo.
	Key   pulumi.StringInput    `pulumi:"key"`
	Notes pulumi.StringPtrInput `pulumi:"notes"`
	// The primary RSA key to be used to sign packages.
	PrimaryKeypairRef   pulumi.StringPtrInput   `pulumi:"primaryKeypairRef"`
	PriorityResolution  pulumi.BoolPtrInput     `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput   `pulumi:"projectKey"`
	PropertySets        pulumi.StringArrayInput `pulumi:"propertySets"`
	RepoLayoutRef       pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	// The secondary RSA key to be used to sign packages.
	SecondaryKeypairRef pulumi.StringPtrInput `pulumi:"secondaryKeypairRef"`
	// When set, the repository will use the deprecated trivial layout.
	//
	// Deprecated: You shouldn't be using this
	TrivialLayout pulumi.BoolPtrInput `pulumi:"trivialLayout"`
	XrayIndex     pulumi.BoolPtrInput `pulumi:"xrayIndex"`
}

func (GetLocalDebianRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLocalDebianRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getLocalDebianRepository.
type GetLocalDebianRepositoryResultOutput struct{ *pulumi.OutputState }

func (GetLocalDebianRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLocalDebianRepositoryResult)(nil)).Elem()
}

func (o GetLocalDebianRepositoryResultOutput) ToGetLocalDebianRepositoryResultOutput() GetLocalDebianRepositoryResultOutput {
	return o
}

func (o GetLocalDebianRepositoryResultOutput) ToGetLocalDebianRepositoryResultOutputWithContext(ctx context.Context) GetLocalDebianRepositoryResultOutput {
	return o
}

func (o GetLocalDebianRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalDebianRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o GetLocalDebianRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalDebianRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o GetLocalDebianRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalDebianRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o GetLocalDebianRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalDebianRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GetLocalDebianRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalDebianRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o GetLocalDebianRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalDebianRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLocalDebianRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalDebianRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLocalDebianRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalDebianRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// The options are Bzip2 (.bz2 extension) (default), LZMA (.lzma extension)
// and XZ (.xz extension).
func (o GetLocalDebianRepositoryResultOutput) IndexCompressionFormats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLocalDebianRepositoryResult) []string { return v.IndexCompressionFormats }).(pulumi.StringArrayOutput)
}

func (o GetLocalDebianRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalDebianRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o GetLocalDebianRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalDebianRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o GetLocalDebianRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalDebianRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

// The primary RSA key to be used to sign packages.
func (o GetLocalDebianRepositoryResultOutput) PrimaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalDebianRepositoryResult) *string { return v.PrimaryKeypairRef }).(pulumi.StringPtrOutput)
}

func (o GetLocalDebianRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalDebianRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o GetLocalDebianRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLocalDebianRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o GetLocalDebianRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalDebianRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o GetLocalDebianRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLocalDebianRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o GetLocalDebianRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalDebianRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The secondary RSA key to be used to sign packages.
func (o GetLocalDebianRepositoryResultOutput) SecondaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalDebianRepositoryResult) *string { return v.SecondaryKeypairRef }).(pulumi.StringPtrOutput)
}

// When set, the repository will use the deprecated trivial layout.
//
// Deprecated: You shouldn't be using this
func (o GetLocalDebianRepositoryResultOutput) TrivialLayout() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalDebianRepositoryResult) *bool { return v.TrivialLayout }).(pulumi.BoolPtrOutput)
}

func (o GetLocalDebianRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalDebianRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLocalDebianRepositoryResultOutput{})
}
