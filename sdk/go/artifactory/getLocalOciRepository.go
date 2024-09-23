// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a local OCI repository resource
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
//			_, err := artifactory.LookupLocalOciRepository(ctx, &artifactory.LookupLocalOciRepositoryArgs{
//				Key: "my-oci-local",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupLocalOciRepository(ctx *pulumi.Context, args *LookupLocalOciRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupLocalOciRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLocalOciRepositoryResult
	err := ctx.Invoke("artifactory:index/getLocalOciRepository:getLocalOciRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocalOciRepository.
type LookupLocalOciRepositoryArgs struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	IncludesPattern        *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit.
	MaxUniqueTags       *int     `pulumi:"maxUniqueTags"`
	Notes               *string  `pulumi:"notes"`
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	PropertySets        []string `pulumi:"propertySets"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	// If greater than 1, overwritten tags will be saved by their digest, up to the set up number.
	TagRetention *int  `pulumi:"tagRetention"`
	XrayIndex    *bool `pulumi:"xrayIndex"`
}

// A collection of values returned by getLocalOciRepository.
type LookupLocalOciRepositoryResult struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id              string  `pulumi:"id"`
	IncludesPattern *string `pulumi:"includesPattern"`
	Key             string  `pulumi:"key"`
	// The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit.
	MaxUniqueTags       *int     `pulumi:"maxUniqueTags"`
	Notes               *string  `pulumi:"notes"`
	PackageType         string   `pulumi:"packageType"`
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	PropertySets        []string `pulumi:"propertySets"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	// If greater than 1, overwritten tags will be saved by their digest, up to the set up number.
	TagRetention *int  `pulumi:"tagRetention"`
	XrayIndex    *bool `pulumi:"xrayIndex"`
}

func LookupLocalOciRepositoryOutput(ctx *pulumi.Context, args LookupLocalOciRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupLocalOciRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLocalOciRepositoryResultOutput, error) {
			args := v.(LookupLocalOciRepositoryArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupLocalOciRepositoryResult
			secret, err := ctx.InvokePackageRaw("artifactory:index/getLocalOciRepository:getLocalOciRepository", args, &rv, "", opts...)
			if err != nil {
				return LookupLocalOciRepositoryResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupLocalOciRepositoryResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupLocalOciRepositoryResultOutput), nil
			}
			return output, nil
		}).(LookupLocalOciRepositoryResultOutput)
}

// A collection of arguments for invoking getLocalOciRepository.
type LookupLocalOciRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput   `pulumi:"blackedOut"`
	CdnRedirect            pulumi.BoolPtrInput   `pulumi:"cdnRedirect"`
	Description            pulumi.StringPtrInput `pulumi:"description"`
	DownloadDirect         pulumi.BoolPtrInput   `pulumi:"downloadDirect"`
	ExcludesPattern        pulumi.StringPtrInput `pulumi:"excludesPattern"`
	IncludesPattern        pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key pulumi.StringInput `pulumi:"key"`
	// The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit.
	MaxUniqueTags       pulumi.IntPtrInput      `pulumi:"maxUniqueTags"`
	Notes               pulumi.StringPtrInput   `pulumi:"notes"`
	PriorityResolution  pulumi.BoolPtrInput     `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput   `pulumi:"projectKey"`
	PropertySets        pulumi.StringArrayInput `pulumi:"propertySets"`
	RepoLayoutRef       pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	// If greater than 1, overwritten tags will be saved by their digest, up to the set up number.
	TagRetention pulumi.IntPtrInput  `pulumi:"tagRetention"`
	XrayIndex    pulumi.BoolPtrInput `pulumi:"xrayIndex"`
}

func (LookupLocalOciRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalOciRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getLocalOciRepository.
type LookupLocalOciRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupLocalOciRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalOciRepositoryResult)(nil)).Elem()
}

func (o LookupLocalOciRepositoryResultOutput) ToLookupLocalOciRepositoryResultOutput() LookupLocalOciRepositoryResultOutput {
	return o
}

func (o LookupLocalOciRepositoryResultOutput) ToLookupLocalOciRepositoryResultOutputWithContext(ctx context.Context) LookupLocalOciRepositoryResultOutput {
	return o
}

func (o LookupLocalOciRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalOciRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalOciRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalOciRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalOciRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalOciRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalOciRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalOciRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupLocalOciRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalOciRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalOciRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalOciRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLocalOciRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalOciRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLocalOciRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalOciRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupLocalOciRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalOciRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

// The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit.
func (o LookupLocalOciRepositoryResultOutput) MaxUniqueTags() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupLocalOciRepositoryResult) *int { return v.MaxUniqueTags }).(pulumi.IntPtrOutput)
}

func (o LookupLocalOciRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalOciRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupLocalOciRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalOciRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupLocalOciRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalOciRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalOciRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocalOciRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupLocalOciRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalOciRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupLocalOciRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocalOciRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupLocalOciRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalOciRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// If greater than 1, overwritten tags will be saved by their digest, up to the set up number.
func (o LookupLocalOciRepositoryResultOutput) TagRetention() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupLocalOciRepositoryResult) *int { return v.TagRetention }).(pulumi.IntPtrOutput)
}

func (o LookupLocalOciRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalOciRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocalOciRepositoryResultOutput{})
}
