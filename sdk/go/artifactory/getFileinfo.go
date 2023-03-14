// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory File Info Data Source
//
// Provides an Artifactory fileinfo datasource. This can be used to read metadata of files stored in Artifactory repositories.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v3/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.GetFileinfo(ctx, &artifactory.GetFileinfoArgs{
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
func GetFileinfo(ctx *pulumi.Context, args *GetFileinfoArgs, opts ...pulumi.InvokeOption) (*GetFileinfoResult, error) {
	var rv GetFileinfoResult
	err := ctx.Invoke("artifactory:index/getFileinfo:getFileinfo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFileinfo.
type GetFileinfoArgs struct {
	// The path to the file within the repository.
	Path string `pulumi:"path"`
	// Name of the repository where the file is stored.
	Repository string `pulumi:"repository"`
}

// A collection of values returned by getFileinfo.
type GetFileinfoResult struct {
	// The time & date when the file was created.
	Created string `pulumi:"created"`
	// The user who created the file.
	CreatedBy string `pulumi:"createdBy"`
	// The URI that can be used to download the file.
	DownloadUri string `pulumi:"downloadUri"`
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
	ModifiedBy string `pulumi:"modifiedBy"`
	Path       string `pulumi:"path"`
	Repository string `pulumi:"repository"`
	// SHA1 checksum of the file.
	Sha1 string `pulumi:"sha1"`
	// SHA256 checksum of the file.
	Sha256 string `pulumi:"sha256"`
	// The size of the file.
	Size int `pulumi:"size"`
}

func GetFileinfoOutput(ctx *pulumi.Context, args GetFileinfoOutputArgs, opts ...pulumi.InvokeOption) GetFileinfoResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFileinfoResult, error) {
			args := v.(GetFileinfoArgs)
			r, err := GetFileinfo(ctx, &args, opts...)
			var s GetFileinfoResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFileinfoResultOutput)
}

// A collection of arguments for invoking getFileinfo.
type GetFileinfoOutputArgs struct {
	// The path to the file within the repository.
	Path pulumi.StringInput `pulumi:"path"`
	// Name of the repository where the file is stored.
	Repository pulumi.StringInput `pulumi:"repository"`
}

func (GetFileinfoOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFileinfoArgs)(nil)).Elem()
}

// A collection of values returned by getFileinfo.
type GetFileinfoResultOutput struct{ *pulumi.OutputState }

func (GetFileinfoResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFileinfoResult)(nil)).Elem()
}

func (o GetFileinfoResultOutput) ToGetFileinfoResultOutput() GetFileinfoResultOutput {
	return o
}

func (o GetFileinfoResultOutput) ToGetFileinfoResultOutputWithContext(ctx context.Context) GetFileinfoResultOutput {
	return o
}

// The time & date when the file was created.
func (o GetFileinfoResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileinfoResult) string { return v.Created }).(pulumi.StringOutput)
}

// The user who created the file.
func (o GetFileinfoResultOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileinfoResult) string { return v.CreatedBy }).(pulumi.StringOutput)
}

// The URI that can be used to download the file.
func (o GetFileinfoResultOutput) DownloadUri() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileinfoResult) string { return v.DownloadUri }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFileinfoResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileinfoResult) string { return v.Id }).(pulumi.StringOutput)
}

// The time & date when the file was last modified.
func (o GetFileinfoResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileinfoResult) string { return v.LastModified }).(pulumi.StringOutput)
}

// The time & date when the file was last updated.
func (o GetFileinfoResultOutput) LastUpdated() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileinfoResult) string { return v.LastUpdated }).(pulumi.StringOutput)
}

// MD5 checksum of the file.
func (o GetFileinfoResultOutput) Md5() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileinfoResult) string { return v.Md5 }).(pulumi.StringOutput)
}

// The mimetype of the file.
func (o GetFileinfoResultOutput) Mimetype() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileinfoResult) string { return v.Mimetype }).(pulumi.StringOutput)
}

// The user who last modified the file.
func (o GetFileinfoResultOutput) ModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileinfoResult) string { return v.ModifiedBy }).(pulumi.StringOutput)
}

func (o GetFileinfoResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileinfoResult) string { return v.Path }).(pulumi.StringOutput)
}

func (o GetFileinfoResultOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileinfoResult) string { return v.Repository }).(pulumi.StringOutput)
}

// SHA1 checksum of the file.
func (o GetFileinfoResultOutput) Sha1() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileinfoResult) string { return v.Sha1 }).(pulumi.StringOutput)
}

// SHA256 checksum of the file.
func (o GetFileinfoResultOutput) Sha256() pulumi.StringOutput {
	return o.ApplyT(func(v GetFileinfoResult) string { return v.Sha256 }).(pulumi.StringOutput)
}

// The size of the file.
func (o GetFileinfoResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v GetFileinfoResult) int { return v.Size }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFileinfoResultOutput{})
}
