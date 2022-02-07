// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Local Cran Repository Resource
//
// Creates a local cran repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-artifactory/sdk/go/artifactory"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := artifactory.NewLocalCranRepository(ctx, "terraform_local_test_cran_repo", &artifactory.LocalCranRepositoryArgs{
// 			Key: pulumi.String("terraform-local-test-cran-repo"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type LocalCranRepository struct {
	pulumi.CustomResourceState

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrOutput   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrOutput   `pulumi:"blackedOut"`
	Description            pulumi.StringPtrOutput `pulumi:"description"`
	DownloadDirect         pulumi.BoolPtrOutput   `pulumi:"downloadDirect"`
	ExcludesPattern        pulumi.StringOutput    `pulumi:"excludesPattern"`
	IncludesPattern        pulumi.StringOutput    `pulumi:"includesPattern"`
	// - the identity key of the repo
	Key         pulumi.StringOutput    `pulumi:"key"`
	Notes       pulumi.StringPtrOutput `pulumi:"notes"`
	PackageType pulumi.StringOutput    `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrOutput     `pulumi:"priorityResolution"`
	PropertySets       pulumi.StringArrayOutput `pulumi:"propertySets"`
	RepoLayoutRef      pulumi.StringOutput      `pulumi:"repoLayoutRef"`
	XrayIndex          pulumi.BoolOutput        `pulumi:"xrayIndex"`
}

// NewLocalCranRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalCranRepository(ctx *pulumi.Context,
	name string, args *LocalCranRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalCranRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource LocalCranRepository
	err := ctx.RegisterResource("artifactory:index/localCranRepository:LocalCranRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalCranRepository gets an existing LocalCranRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalCranRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalCranRepositoryState, opts ...pulumi.ResourceOption) (*LocalCranRepository, error) {
	var resource LocalCranRepository
	err := ctx.ReadResource("artifactory:index/localCranRepository:LocalCranRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalCranRepository resources.
type localCranRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	IncludesPattern        *string `pulumi:"includesPattern"`
	// - the identity key of the repo
	Key         *string `pulumi:"key"`
	Notes       *string `pulumi:"notes"`
	PackageType *string `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool    `pulumi:"priorityResolution"`
	PropertySets       []string `pulumi:"propertySets"`
	RepoLayoutRef      *string  `pulumi:"repoLayoutRef"`
	XrayIndex          *bool    `pulumi:"xrayIndex"`
}

type LocalCranRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	BlackedOut             pulumi.BoolPtrInput
	Description            pulumi.StringPtrInput
	DownloadDirect         pulumi.BoolPtrInput
	ExcludesPattern        pulumi.StringPtrInput
	IncludesPattern        pulumi.StringPtrInput
	// - the identity key of the repo
	Key         pulumi.StringPtrInput
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	PropertySets       pulumi.StringArrayInput
	RepoLayoutRef      pulumi.StringPtrInput
	XrayIndex          pulumi.BoolPtrInput
}

func (LocalCranRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localCranRepositoryState)(nil)).Elem()
}

type localCranRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	IncludesPattern        *string `pulumi:"includesPattern"`
	// - the identity key of the repo
	Key   string  `pulumi:"key"`
	Notes *string `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool    `pulumi:"priorityResolution"`
	PropertySets       []string `pulumi:"propertySets"`
	RepoLayoutRef      *string  `pulumi:"repoLayoutRef"`
	XrayIndex          *bool    `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a LocalCranRepository resource.
type LocalCranRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	BlackedOut             pulumi.BoolPtrInput
	Description            pulumi.StringPtrInput
	DownloadDirect         pulumi.BoolPtrInput
	ExcludesPattern        pulumi.StringPtrInput
	IncludesPattern        pulumi.StringPtrInput
	// - the identity key of the repo
	Key   pulumi.StringInput
	Notes pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	PropertySets       pulumi.StringArrayInput
	RepoLayoutRef      pulumi.StringPtrInput
	XrayIndex          pulumi.BoolPtrInput
}

func (LocalCranRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localCranRepositoryArgs)(nil)).Elem()
}

type LocalCranRepositoryInput interface {
	pulumi.Input

	ToLocalCranRepositoryOutput() LocalCranRepositoryOutput
	ToLocalCranRepositoryOutputWithContext(ctx context.Context) LocalCranRepositoryOutput
}

func (*LocalCranRepository) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalCranRepository)(nil))
}

func (i *LocalCranRepository) ToLocalCranRepositoryOutput() LocalCranRepositoryOutput {
	return i.ToLocalCranRepositoryOutputWithContext(context.Background())
}

func (i *LocalCranRepository) ToLocalCranRepositoryOutputWithContext(ctx context.Context) LocalCranRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalCranRepositoryOutput)
}

func (i *LocalCranRepository) ToLocalCranRepositoryPtrOutput() LocalCranRepositoryPtrOutput {
	return i.ToLocalCranRepositoryPtrOutputWithContext(context.Background())
}

func (i *LocalCranRepository) ToLocalCranRepositoryPtrOutputWithContext(ctx context.Context) LocalCranRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalCranRepositoryPtrOutput)
}

type LocalCranRepositoryPtrInput interface {
	pulumi.Input

	ToLocalCranRepositoryPtrOutput() LocalCranRepositoryPtrOutput
	ToLocalCranRepositoryPtrOutputWithContext(ctx context.Context) LocalCranRepositoryPtrOutput
}

type localCranRepositoryPtrType LocalCranRepositoryArgs

func (*localCranRepositoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalCranRepository)(nil))
}

func (i *localCranRepositoryPtrType) ToLocalCranRepositoryPtrOutput() LocalCranRepositoryPtrOutput {
	return i.ToLocalCranRepositoryPtrOutputWithContext(context.Background())
}

func (i *localCranRepositoryPtrType) ToLocalCranRepositoryPtrOutputWithContext(ctx context.Context) LocalCranRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalCranRepositoryPtrOutput)
}

// LocalCranRepositoryArrayInput is an input type that accepts LocalCranRepositoryArray and LocalCranRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalCranRepositoryArrayInput` via:
//
//          LocalCranRepositoryArray{ LocalCranRepositoryArgs{...} }
type LocalCranRepositoryArrayInput interface {
	pulumi.Input

	ToLocalCranRepositoryArrayOutput() LocalCranRepositoryArrayOutput
	ToLocalCranRepositoryArrayOutputWithContext(context.Context) LocalCranRepositoryArrayOutput
}

type LocalCranRepositoryArray []LocalCranRepositoryInput

func (LocalCranRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalCranRepository)(nil)).Elem()
}

func (i LocalCranRepositoryArray) ToLocalCranRepositoryArrayOutput() LocalCranRepositoryArrayOutput {
	return i.ToLocalCranRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalCranRepositoryArray) ToLocalCranRepositoryArrayOutputWithContext(ctx context.Context) LocalCranRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalCranRepositoryArrayOutput)
}

// LocalCranRepositoryMapInput is an input type that accepts LocalCranRepositoryMap and LocalCranRepositoryMapOutput values.
// You can construct a concrete instance of `LocalCranRepositoryMapInput` via:
//
//          LocalCranRepositoryMap{ "key": LocalCranRepositoryArgs{...} }
type LocalCranRepositoryMapInput interface {
	pulumi.Input

	ToLocalCranRepositoryMapOutput() LocalCranRepositoryMapOutput
	ToLocalCranRepositoryMapOutputWithContext(context.Context) LocalCranRepositoryMapOutput
}

type LocalCranRepositoryMap map[string]LocalCranRepositoryInput

func (LocalCranRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalCranRepository)(nil)).Elem()
}

func (i LocalCranRepositoryMap) ToLocalCranRepositoryMapOutput() LocalCranRepositoryMapOutput {
	return i.ToLocalCranRepositoryMapOutputWithContext(context.Background())
}

func (i LocalCranRepositoryMap) ToLocalCranRepositoryMapOutputWithContext(ctx context.Context) LocalCranRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalCranRepositoryMapOutput)
}

type LocalCranRepositoryOutput struct{ *pulumi.OutputState }

func (LocalCranRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalCranRepository)(nil))
}

func (o LocalCranRepositoryOutput) ToLocalCranRepositoryOutput() LocalCranRepositoryOutput {
	return o
}

func (o LocalCranRepositoryOutput) ToLocalCranRepositoryOutputWithContext(ctx context.Context) LocalCranRepositoryOutput {
	return o
}

func (o LocalCranRepositoryOutput) ToLocalCranRepositoryPtrOutput() LocalCranRepositoryPtrOutput {
	return o.ToLocalCranRepositoryPtrOutputWithContext(context.Background())
}

func (o LocalCranRepositoryOutput) ToLocalCranRepositoryPtrOutputWithContext(ctx context.Context) LocalCranRepositoryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LocalCranRepository) *LocalCranRepository {
		return &v
	}).(LocalCranRepositoryPtrOutput)
}

type LocalCranRepositoryPtrOutput struct{ *pulumi.OutputState }

func (LocalCranRepositoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalCranRepository)(nil))
}

func (o LocalCranRepositoryPtrOutput) ToLocalCranRepositoryPtrOutput() LocalCranRepositoryPtrOutput {
	return o
}

func (o LocalCranRepositoryPtrOutput) ToLocalCranRepositoryPtrOutputWithContext(ctx context.Context) LocalCranRepositoryPtrOutput {
	return o
}

func (o LocalCranRepositoryPtrOutput) Elem() LocalCranRepositoryOutput {
	return o.ApplyT(func(v *LocalCranRepository) LocalCranRepository {
		if v != nil {
			return *v
		}
		var ret LocalCranRepository
		return ret
	}).(LocalCranRepositoryOutput)
}

type LocalCranRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalCranRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LocalCranRepository)(nil))
}

func (o LocalCranRepositoryArrayOutput) ToLocalCranRepositoryArrayOutput() LocalCranRepositoryArrayOutput {
	return o
}

func (o LocalCranRepositoryArrayOutput) ToLocalCranRepositoryArrayOutputWithContext(ctx context.Context) LocalCranRepositoryArrayOutput {
	return o
}

func (o LocalCranRepositoryArrayOutput) Index(i pulumi.IntInput) LocalCranRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LocalCranRepository {
		return vs[0].([]LocalCranRepository)[vs[1].(int)]
	}).(LocalCranRepositoryOutput)
}

type LocalCranRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalCranRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LocalCranRepository)(nil))
}

func (o LocalCranRepositoryMapOutput) ToLocalCranRepositoryMapOutput() LocalCranRepositoryMapOutput {
	return o
}

func (o LocalCranRepositoryMapOutput) ToLocalCranRepositoryMapOutputWithContext(ctx context.Context) LocalCranRepositoryMapOutput {
	return o
}

func (o LocalCranRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalCranRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LocalCranRepository {
		return vs[0].(map[string]LocalCranRepository)[vs[1].(string)]
	}).(LocalCranRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalCranRepositoryInput)(nil)).Elem(), &LocalCranRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalCranRepositoryPtrInput)(nil)).Elem(), &LocalCranRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalCranRepositoryArrayInput)(nil)).Elem(), LocalCranRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalCranRepositoryMapInput)(nil)).Elem(), LocalCranRepositoryMap{})
	pulumi.RegisterOutputType(LocalCranRepositoryOutput{})
	pulumi.RegisterOutputType(LocalCranRepositoryPtrOutput{})
	pulumi.RegisterOutputType(LocalCranRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalCranRepositoryMapOutput{})
}
