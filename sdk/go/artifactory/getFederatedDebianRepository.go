// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a federated Debian repository.
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
//			_, err := artifactory.LookupFederatedDebianRepository(ctx, &artifactory.LookupFederatedDebianRepositoryArgs{
//				Key: "federated-test-debian-repo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFederatedDebianRepository(ctx *pulumi.Context, args *LookupFederatedDebianRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupFederatedDebianRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFederatedDebianRepositoryResult
	err := ctx.Invoke("artifactory:index/getFederatedDebianRepository:getFederatedDebianRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedDebianRepository.
type LookupFederatedDebianRepositoryArgs struct {
	ArchiveBrowsingEnabled  *bool    `pulumi:"archiveBrowsingEnabled"`
	BlackedOut              *bool    `pulumi:"blackedOut"`
	CdnRedirect             *bool    `pulumi:"cdnRedirect"`
	CleanupOnDelete         *bool    `pulumi:"cleanupOnDelete"`
	Description             *string  `pulumi:"description"`
	DownloadDirect          *bool    `pulumi:"downloadDirect"`
	ExcludesPattern         *string  `pulumi:"excludesPattern"`
	IncludesPattern         *string  `pulumi:"includesPattern"`
	IndexCompressionFormats []string `pulumi:"indexCompressionFormats"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             []GetFederatedDebianRepositoryMember `pulumi:"members"`
	Notes               *string                              `pulumi:"notes"`
	PrimaryKeypairRef   *string                              `pulumi:"primaryKeypairRef"`
	PriorityResolution  *bool                                `pulumi:"priorityResolution"`
	ProjectEnvironments []string                             `pulumi:"projectEnvironments"`
	ProjectKey          *string                              `pulumi:"projectKey"`
	PropertySets        []string                             `pulumi:"propertySets"`
	RepoLayoutRef       *string                              `pulumi:"repoLayoutRef"`
	SecondaryKeypairRef *string                              `pulumi:"secondaryKeypairRef"`
	// Deprecated: You shouldn't be using this
	TrivialLayout *bool `pulumi:"trivialLayout"`
	XrayIndex     *bool `pulumi:"xrayIndex"`
}

// A collection of values returned by getFederatedDebianRepository.
type LookupFederatedDebianRepositoryResult struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id                      string   `pulumi:"id"`
	IncludesPattern         *string  `pulumi:"includesPattern"`
	IndexCompressionFormats []string `pulumi:"indexCompressionFormats"`
	Key                     string   `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             []GetFederatedDebianRepositoryMember `pulumi:"members"`
	Notes               *string                              `pulumi:"notes"`
	PackageType         string                               `pulumi:"packageType"`
	PrimaryKeypairRef   *string                              `pulumi:"primaryKeypairRef"`
	PriorityResolution  *bool                                `pulumi:"priorityResolution"`
	ProjectEnvironments []string                             `pulumi:"projectEnvironments"`
	ProjectKey          *string                              `pulumi:"projectKey"`
	PropertySets        []string                             `pulumi:"propertySets"`
	RepoLayoutRef       *string                              `pulumi:"repoLayoutRef"`
	SecondaryKeypairRef *string                              `pulumi:"secondaryKeypairRef"`
	// Deprecated: You shouldn't be using this
	TrivialLayout *bool `pulumi:"trivialLayout"`
	XrayIndex     *bool `pulumi:"xrayIndex"`
}

func LookupFederatedDebianRepositoryOutput(ctx *pulumi.Context, args LookupFederatedDebianRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedDebianRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedDebianRepositoryResult, error) {
			args := v.(LookupFederatedDebianRepositoryArgs)
			r, err := LookupFederatedDebianRepository(ctx, &args, opts...)
			var s LookupFederatedDebianRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFederatedDebianRepositoryResultOutput)
}

// A collection of arguments for invoking getFederatedDebianRepository.
type LookupFederatedDebianRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled  pulumi.BoolPtrInput     `pulumi:"archiveBrowsingEnabled"`
	BlackedOut              pulumi.BoolPtrInput     `pulumi:"blackedOut"`
	CdnRedirect             pulumi.BoolPtrInput     `pulumi:"cdnRedirect"`
	CleanupOnDelete         pulumi.BoolPtrInput     `pulumi:"cleanupOnDelete"`
	Description             pulumi.StringPtrInput   `pulumi:"description"`
	DownloadDirect          pulumi.BoolPtrInput     `pulumi:"downloadDirect"`
	ExcludesPattern         pulumi.StringPtrInput   `pulumi:"excludesPattern"`
	IncludesPattern         pulumi.StringPtrInput   `pulumi:"includesPattern"`
	IndexCompressionFormats pulumi.StringArrayInput `pulumi:"indexCompressionFormats"`
	// the identity key of the repo.
	Key pulumi.StringInput `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             GetFederatedDebianRepositoryMemberArrayInput `pulumi:"members"`
	Notes               pulumi.StringPtrInput                        `pulumi:"notes"`
	PrimaryKeypairRef   pulumi.StringPtrInput                        `pulumi:"primaryKeypairRef"`
	PriorityResolution  pulumi.BoolPtrInput                          `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayInput                      `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput                        `pulumi:"projectKey"`
	PropertySets        pulumi.StringArrayInput                      `pulumi:"propertySets"`
	RepoLayoutRef       pulumi.StringPtrInput                        `pulumi:"repoLayoutRef"`
	SecondaryKeypairRef pulumi.StringPtrInput                        `pulumi:"secondaryKeypairRef"`
	// Deprecated: You shouldn't be using this
	TrivialLayout pulumi.BoolPtrInput `pulumi:"trivialLayout"`
	XrayIndex     pulumi.BoolPtrInput `pulumi:"xrayIndex"`
}

func (LookupFederatedDebianRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedDebianRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedDebianRepository.
type LookupFederatedDebianRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedDebianRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedDebianRepositoryResult)(nil)).Elem()
}

func (o LookupFederatedDebianRepositoryResultOutput) ToLookupFederatedDebianRepositoryResultOutput() LookupFederatedDebianRepositoryResultOutput {
	return o
}

func (o LookupFederatedDebianRepositoryResultOutput) ToLookupFederatedDebianRepositoryResultOutputWithContext(ctx context.Context) LookupFederatedDebianRepositoryResultOutput {
	return o
}

func (o LookupFederatedDebianRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDebianRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedDebianRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDebianRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedDebianRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDebianRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedDebianRepositoryResultOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDebianRepositoryResult) *bool { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedDebianRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedDebianRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedDebianRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDebianRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedDebianRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedDebianRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFederatedDebianRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedDebianRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedDebianRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedDebianRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedDebianRepositoryResultOutput) IndexCompressionFormats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedDebianRepositoryResult) []string { return v.IndexCompressionFormats }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedDebianRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedDebianRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o LookupFederatedDebianRepositoryResultOutput) Members() GetFederatedDebianRepositoryMemberArrayOutput {
	return o.ApplyT(func(v LookupFederatedDebianRepositoryResult) []GetFederatedDebianRepositoryMember { return v.Members }).(GetFederatedDebianRepositoryMemberArrayOutput)
}

func (o LookupFederatedDebianRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedDebianRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedDebianRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedDebianRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupFederatedDebianRepositoryResultOutput) PrimaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedDebianRepositoryResult) *string { return v.PrimaryKeypairRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedDebianRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDebianRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedDebianRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedDebianRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedDebianRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedDebianRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedDebianRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedDebianRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedDebianRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedDebianRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedDebianRepositoryResultOutput) SecondaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedDebianRepositoryResult) *string { return v.SecondaryKeypairRef }).(pulumi.StringPtrOutput)
}

// Deprecated: You shouldn't be using this
func (o LookupFederatedDebianRepositoryResultOutput) TrivialLayout() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDebianRepositoryResult) *bool { return v.TrivialLayout }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedDebianRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDebianRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedDebianRepositoryResultOutput{})
}
