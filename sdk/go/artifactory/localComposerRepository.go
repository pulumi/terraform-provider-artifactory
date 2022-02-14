// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Local Php-Composer Repository Resource
//
// Creates a local composer repository.
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
// 		_, err := artifactory.NewLocalComposerRepository(ctx, "terraform-local-test-composer-repo", &artifactory.LocalComposerRepositoryArgs{
// 			Key: pulumi.String("terraform-local-test-composer-repo"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type LocalComposerRepository struct {
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

// NewLocalComposerRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalComposerRepository(ctx *pulumi.Context,
	name string, args *LocalComposerRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalComposerRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource LocalComposerRepository
	err := ctx.RegisterResource("artifactory:index/localComposerRepository:LocalComposerRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalComposerRepository gets an existing LocalComposerRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalComposerRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalComposerRepositoryState, opts ...pulumi.ResourceOption) (*LocalComposerRepository, error) {
	var resource LocalComposerRepository
	err := ctx.ReadResource("artifactory:index/localComposerRepository:LocalComposerRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalComposerRepository resources.
type localComposerRepositoryState struct {
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

type LocalComposerRepositoryState struct {
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

func (LocalComposerRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localComposerRepositoryState)(nil)).Elem()
}

type localComposerRepositoryArgs struct {
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

// The set of arguments for constructing a LocalComposerRepository resource.
type LocalComposerRepositoryArgs struct {
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

func (LocalComposerRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localComposerRepositoryArgs)(nil)).Elem()
}

type LocalComposerRepositoryInput interface {
	pulumi.Input

	ToLocalComposerRepositoryOutput() LocalComposerRepositoryOutput
	ToLocalComposerRepositoryOutputWithContext(ctx context.Context) LocalComposerRepositoryOutput
}

func (*LocalComposerRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalComposerRepository)(nil)).Elem()
}

func (i *LocalComposerRepository) ToLocalComposerRepositoryOutput() LocalComposerRepositoryOutput {
	return i.ToLocalComposerRepositoryOutputWithContext(context.Background())
}

func (i *LocalComposerRepository) ToLocalComposerRepositoryOutputWithContext(ctx context.Context) LocalComposerRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalComposerRepositoryOutput)
}

// LocalComposerRepositoryArrayInput is an input type that accepts LocalComposerRepositoryArray and LocalComposerRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalComposerRepositoryArrayInput` via:
//
//          LocalComposerRepositoryArray{ LocalComposerRepositoryArgs{...} }
type LocalComposerRepositoryArrayInput interface {
	pulumi.Input

	ToLocalComposerRepositoryArrayOutput() LocalComposerRepositoryArrayOutput
	ToLocalComposerRepositoryArrayOutputWithContext(context.Context) LocalComposerRepositoryArrayOutput
}

type LocalComposerRepositoryArray []LocalComposerRepositoryInput

func (LocalComposerRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalComposerRepository)(nil)).Elem()
}

func (i LocalComposerRepositoryArray) ToLocalComposerRepositoryArrayOutput() LocalComposerRepositoryArrayOutput {
	return i.ToLocalComposerRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalComposerRepositoryArray) ToLocalComposerRepositoryArrayOutputWithContext(ctx context.Context) LocalComposerRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalComposerRepositoryArrayOutput)
}

// LocalComposerRepositoryMapInput is an input type that accepts LocalComposerRepositoryMap and LocalComposerRepositoryMapOutput values.
// You can construct a concrete instance of `LocalComposerRepositoryMapInput` via:
//
//          LocalComposerRepositoryMap{ "key": LocalComposerRepositoryArgs{...} }
type LocalComposerRepositoryMapInput interface {
	pulumi.Input

	ToLocalComposerRepositoryMapOutput() LocalComposerRepositoryMapOutput
	ToLocalComposerRepositoryMapOutputWithContext(context.Context) LocalComposerRepositoryMapOutput
}

type LocalComposerRepositoryMap map[string]LocalComposerRepositoryInput

func (LocalComposerRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalComposerRepository)(nil)).Elem()
}

func (i LocalComposerRepositoryMap) ToLocalComposerRepositoryMapOutput() LocalComposerRepositoryMapOutput {
	return i.ToLocalComposerRepositoryMapOutputWithContext(context.Background())
}

func (i LocalComposerRepositoryMap) ToLocalComposerRepositoryMapOutputWithContext(ctx context.Context) LocalComposerRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalComposerRepositoryMapOutput)
}

type LocalComposerRepositoryOutput struct{ *pulumi.OutputState }

func (LocalComposerRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalComposerRepository)(nil)).Elem()
}

func (o LocalComposerRepositoryOutput) ToLocalComposerRepositoryOutput() LocalComposerRepositoryOutput {
	return o
}

func (o LocalComposerRepositoryOutput) ToLocalComposerRepositoryOutputWithContext(ctx context.Context) LocalComposerRepositoryOutput {
	return o
}

type LocalComposerRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalComposerRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalComposerRepository)(nil)).Elem()
}

func (o LocalComposerRepositoryArrayOutput) ToLocalComposerRepositoryArrayOutput() LocalComposerRepositoryArrayOutput {
	return o
}

func (o LocalComposerRepositoryArrayOutput) ToLocalComposerRepositoryArrayOutputWithContext(ctx context.Context) LocalComposerRepositoryArrayOutput {
	return o
}

func (o LocalComposerRepositoryArrayOutput) Index(i pulumi.IntInput) LocalComposerRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalComposerRepository {
		return vs[0].([]*LocalComposerRepository)[vs[1].(int)]
	}).(LocalComposerRepositoryOutput)
}

type LocalComposerRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalComposerRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalComposerRepository)(nil)).Elem()
}

func (o LocalComposerRepositoryMapOutput) ToLocalComposerRepositoryMapOutput() LocalComposerRepositoryMapOutput {
	return o
}

func (o LocalComposerRepositoryMapOutput) ToLocalComposerRepositoryMapOutputWithContext(ctx context.Context) LocalComposerRepositoryMapOutput {
	return o
}

func (o LocalComposerRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalComposerRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalComposerRepository {
		return vs[0].(map[string]*LocalComposerRepository)[vs[1].(string)]
	}).(LocalComposerRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalComposerRepositoryInput)(nil)).Elem(), &LocalComposerRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalComposerRepositoryArrayInput)(nil)).Elem(), LocalComposerRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalComposerRepositoryMapInput)(nil)).Elem(), LocalComposerRepositoryMap{})
	pulumi.RegisterOutputType(LocalComposerRepositoryOutput{})
	pulumi.RegisterOutputType(LocalComposerRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalComposerRepositoryMapOutput{})
}
