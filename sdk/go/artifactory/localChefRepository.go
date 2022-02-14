// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Local Chef Repository Resource
//
// Creates a local chef repository.
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
// 		_, err := artifactory.NewLocalChefRepository(ctx, "terraform-local-test-chef-repo", &artifactory.LocalChefRepositoryArgs{
// 			Key: pulumi.String("terraform-local-test-chef-repo"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type LocalChefRepository struct {
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

// NewLocalChefRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalChefRepository(ctx *pulumi.Context,
	name string, args *LocalChefRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalChefRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource LocalChefRepository
	err := ctx.RegisterResource("artifactory:index/localChefRepository:LocalChefRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalChefRepository gets an existing LocalChefRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalChefRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalChefRepositoryState, opts ...pulumi.ResourceOption) (*LocalChefRepository, error) {
	var resource LocalChefRepository
	err := ctx.ReadResource("artifactory:index/localChefRepository:LocalChefRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalChefRepository resources.
type localChefRepositoryState struct {
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

type LocalChefRepositoryState struct {
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

func (LocalChefRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localChefRepositoryState)(nil)).Elem()
}

type localChefRepositoryArgs struct {
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

// The set of arguments for constructing a LocalChefRepository resource.
type LocalChefRepositoryArgs struct {
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

func (LocalChefRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localChefRepositoryArgs)(nil)).Elem()
}

type LocalChefRepositoryInput interface {
	pulumi.Input

	ToLocalChefRepositoryOutput() LocalChefRepositoryOutput
	ToLocalChefRepositoryOutputWithContext(ctx context.Context) LocalChefRepositoryOutput
}

func (*LocalChefRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalChefRepository)(nil)).Elem()
}

func (i *LocalChefRepository) ToLocalChefRepositoryOutput() LocalChefRepositoryOutput {
	return i.ToLocalChefRepositoryOutputWithContext(context.Background())
}

func (i *LocalChefRepository) ToLocalChefRepositoryOutputWithContext(ctx context.Context) LocalChefRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalChefRepositoryOutput)
}

// LocalChefRepositoryArrayInput is an input type that accepts LocalChefRepositoryArray and LocalChefRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalChefRepositoryArrayInput` via:
//
//          LocalChefRepositoryArray{ LocalChefRepositoryArgs{...} }
type LocalChefRepositoryArrayInput interface {
	pulumi.Input

	ToLocalChefRepositoryArrayOutput() LocalChefRepositoryArrayOutput
	ToLocalChefRepositoryArrayOutputWithContext(context.Context) LocalChefRepositoryArrayOutput
}

type LocalChefRepositoryArray []LocalChefRepositoryInput

func (LocalChefRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalChefRepository)(nil)).Elem()
}

func (i LocalChefRepositoryArray) ToLocalChefRepositoryArrayOutput() LocalChefRepositoryArrayOutput {
	return i.ToLocalChefRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalChefRepositoryArray) ToLocalChefRepositoryArrayOutputWithContext(ctx context.Context) LocalChefRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalChefRepositoryArrayOutput)
}

// LocalChefRepositoryMapInput is an input type that accepts LocalChefRepositoryMap and LocalChefRepositoryMapOutput values.
// You can construct a concrete instance of `LocalChefRepositoryMapInput` via:
//
//          LocalChefRepositoryMap{ "key": LocalChefRepositoryArgs{...} }
type LocalChefRepositoryMapInput interface {
	pulumi.Input

	ToLocalChefRepositoryMapOutput() LocalChefRepositoryMapOutput
	ToLocalChefRepositoryMapOutputWithContext(context.Context) LocalChefRepositoryMapOutput
}

type LocalChefRepositoryMap map[string]LocalChefRepositoryInput

func (LocalChefRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalChefRepository)(nil)).Elem()
}

func (i LocalChefRepositoryMap) ToLocalChefRepositoryMapOutput() LocalChefRepositoryMapOutput {
	return i.ToLocalChefRepositoryMapOutputWithContext(context.Background())
}

func (i LocalChefRepositoryMap) ToLocalChefRepositoryMapOutputWithContext(ctx context.Context) LocalChefRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalChefRepositoryMapOutput)
}

type LocalChefRepositoryOutput struct{ *pulumi.OutputState }

func (LocalChefRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalChefRepository)(nil)).Elem()
}

func (o LocalChefRepositoryOutput) ToLocalChefRepositoryOutput() LocalChefRepositoryOutput {
	return o
}

func (o LocalChefRepositoryOutput) ToLocalChefRepositoryOutputWithContext(ctx context.Context) LocalChefRepositoryOutput {
	return o
}

type LocalChefRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalChefRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalChefRepository)(nil)).Elem()
}

func (o LocalChefRepositoryArrayOutput) ToLocalChefRepositoryArrayOutput() LocalChefRepositoryArrayOutput {
	return o
}

func (o LocalChefRepositoryArrayOutput) ToLocalChefRepositoryArrayOutputWithContext(ctx context.Context) LocalChefRepositoryArrayOutput {
	return o
}

func (o LocalChefRepositoryArrayOutput) Index(i pulumi.IntInput) LocalChefRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalChefRepository {
		return vs[0].([]*LocalChefRepository)[vs[1].(int)]
	}).(LocalChefRepositoryOutput)
}

type LocalChefRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalChefRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalChefRepository)(nil)).Elem()
}

func (o LocalChefRepositoryMapOutput) ToLocalChefRepositoryMapOutput() LocalChefRepositoryMapOutput {
	return o
}

func (o LocalChefRepositoryMapOutput) ToLocalChefRepositoryMapOutputWithContext(ctx context.Context) LocalChefRepositoryMapOutput {
	return o
}

func (o LocalChefRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalChefRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalChefRepository {
		return vs[0].(map[string]*LocalChefRepository)[vs[1].(string)]
	}).(LocalChefRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalChefRepositoryInput)(nil)).Elem(), &LocalChefRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalChefRepositoryArrayInput)(nil)).Elem(), LocalChefRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalChefRepositoryMapInput)(nil)).Elem(), LocalChefRepositoryMap{})
	pulumi.RegisterOutputType(LocalChefRepositoryOutput{})
	pulumi.RegisterOutputType(LocalChefRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalChefRepositoryMapOutput{})
}
