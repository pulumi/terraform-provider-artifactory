// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v4/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory File Data Source
//
// Provides an Artifactory file datasource. This can be used to download a file from a given Artifactory repository.
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
//			_, err := artifactory.GetFile(ctx, &artifactory.GetFileArgs{
//				OutputPath: "tmp/artifact.zip",
//				Path:       "/path/to/the/artifact.zip",
//				Repository: "repo-key",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetFile(ctx *pulumi.Context, args *GetFileArgs, opts ...pulumi.InvokeOption) (*GetFileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetFileResult
	err := ctx.Invoke("artifactory:index/getFile:getFile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFile.
type GetFileArgs struct {
	// If set to true, an existing file in the outputPath will be overwritten. Default: `false`
	ForceOverwrite *bool `pulumi:"forceOverwrite"`
	// The local path the file should be downloaded to.
	OutputPath string `pulumi:"outputPath"`
	// The path to the file within the repository.
	Path string `pulumi:"path"`
	// If set to `true`, the provider will get the artifact directly from Artifactory without attempting to resolve it or verify it and will delegate this to artifactory
	// if the file exists. More details in the [official documentation](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-RetrieveLatestArtifact)
	PathIsAliased *bool `pulumi:"pathIsAliased"`
	// Name of the repository where the file is stored.
	Repository string `pulumi:"repository"`
}

// A collection of values returned by getFile.
type GetFileResult struct {
	// The time & date when the file was created.
	Created string `pulumi:"created"`
	// The user who created the file.
	CreatedBy string `pulumi:"createdBy"`
	// The URI that can be used to download the file.
	DownloadUri    string `pulumi:"downloadUri"`
	ForceOverwrite *bool  `pulumi:"forceOverwrite"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The time & date when the file was last modified.
	LastModified string `pulumi:"lastModified"`
	// The time & date when the file was last updated.
	LastUpdated string `pulumi:"lastUpdated"`
	// MD5 checksum of the file.
	Md5 string `pulumi:"md5"`
	// The mimetype of the file.
	Mimetype string `pulumi:"mimetype"`
	// The user who last modified the file.
	ModifiedBy    string `pulumi:"modifiedBy"`
	OutputPath    string `pulumi:"outputPath"`
	Path          string `pulumi:"path"`
	PathIsAliased *bool  `pulumi:"pathIsAliased"`
	Repository    string `pulumi:"repository"`
	// SHA1 checksum of the file.
	Sha1 string `pulumi:"sha1"`
	// SHA256 checksum of the file.
	Sha256 string `pulumi:"sha256"`
	// The size of the file.
	Size int `pulumi:"size"`
}

func GetFileOutput(ctx *pulumi.Context, args GetFileOutputArgs, opts ...pulumi.InvokeOption) GetFileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFileResult, error) {
			args := v.(GetFileArgs)
			r, err := GetFile(ctx, &args, opts...)
			var s GetFileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFileResultOutput)
}

// A collection of arguments for invoking getFile.
type GetFileOutputArgs struct {
	// If set to true, an existing file in the outputPath will be overwritten. Default: `false`
	ForceOverwrite pulumi.BoolPtrInput `pulumi:"forceOverwrite"`
	// The local path the file should be downloaded to.
	OutputPath pulumi.StringInput `pulumi:"outputPath"`
	// The path to the file within the repository.
	Path pulumi.StringInput `pulumi:"path"`
	// If set to `true`, the provider will get the artifact directly from Artifactory without attempting to resolve it or verify it and will delegate this to artifactory
	// if the file exists. More details in the [official documentation](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-RetrieveLatestArtifact)
	PathIsAliased pulumi.BoolPtrInput `pulumi:"pathIsAliased"`
	// Name of the repository where the file is stored.
	Repository pulumi.StringInput `pulumi:"repository"`
}

func (GetFileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFileArgs)(nil)).Elem()
}

// A collection of values returned by getFile.
type GetFileResultOutput struct{ *pulumi.OutputState }

func (GetFileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFileResult)(nil)).Elem()
}

func (o GetFileResultOutput) ToGetFileResultOutput() GetFileResultOutput {
	return o
}

func (o GetFileResultOutput) ToGetFileResultOutputWithContext(ctx context.Context) GetFileResultOutput {
	return o
}

// The time & date when the file was created.
func (o GetFileResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileResult) string { return v.Created }).(pulumi.StringOutput)
}

// The user who created the file.
func (o GetFileResultOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileResult) string { return v.CreatedBy }).(pulumi.StringOutput)
}

// The URI that can be used to download the file.
func (o GetFileResultOutput) DownloadUri() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileResult) string { return v.DownloadUri }).(pulumi.StringOutput)
}

func (o GetFileResultOutput) ForceOverwrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetFileResult) *bool { return v.ForceOverwrite }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileResult) string { return v.Id }).(pulumi.StringOutput)
}

// The time & date when the file was last modified.
func (o GetFileResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileResult) string { return v.LastModified }).(pulumi.StringOutput)
}

// The time & date when the file was last updated.
func (o GetFileResultOutput) LastUpdated() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileResult) string { return v.LastUpdated }).(pulumi.StringOutput)
}

// MD5 checksum of the file.
func (o GetFileResultOutput) Md5() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileResult) string { return v.Md5 }).(pulumi.StringOutput)
}

// The mimetype of the file.
func (o GetFileResultOutput) Mimetype() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileResult) string { return v.Mimetype }).(pulumi.StringOutput)
}

// The user who last modified the file.
func (o GetFileResultOutput) ModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileResult) string { return v.ModifiedBy }).(pulumi.StringOutput)
}

func (o GetFileResultOutput) OutputPath() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileResult) string { return v.OutputPath }).(pulumi.StringOutput)
}

func (o GetFileResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileResult) string { return v.Path }).(pulumi.StringOutput)
}

func (o GetFileResultOutput) PathIsAliased() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetFileResult) *bool { return v.PathIsAliased }).(pulumi.BoolPtrOutput)
}

func (o GetFileResultOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileResult) string { return v.Repository }).(pulumi.StringOutput)
}

// SHA1 checksum of the file.
func (o GetFileResultOutput) Sha1() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileResult) string { return v.Sha1 }).(pulumi.StringOutput)
}

// SHA256 checksum of the file.
func (o GetFileResultOutput) Sha256() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileResult) string { return v.Sha256 }).(pulumi.StringOutput)
}

// The size of the file.
func (o GetFileResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v GetFileResult) int { return v.Size }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFileResultOutput{})
}
