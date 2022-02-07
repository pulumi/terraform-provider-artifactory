// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Local Go Repository Resource
//
// Creates a local go repository.
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
// 		_, err := artifactory.NewLocalGoRepository(ctx, "terraform_local_test_go_repo", &artifactory.LocalGoRepositoryArgs{
// 			Key: pulumi.String("terraform-local-test-go-repo"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type LocalGoRepository struct {
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

// NewLocalGoRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalGoRepository(ctx *pulumi.Context,
	name string, args *LocalGoRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalGoRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource LocalGoRepository
	err := ctx.RegisterResource("artifactory:index/localGoRepository:LocalGoRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalGoRepository gets an existing LocalGoRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalGoRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalGoRepositoryState, opts ...pulumi.ResourceOption) (*LocalGoRepository, error) {
	var resource LocalGoRepository
	err := ctx.ReadResource("artifactory:index/localGoRepository:LocalGoRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalGoRepository resources.
type localGoRepositoryState struct {
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

type LocalGoRepositoryState struct {
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

func (LocalGoRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localGoRepositoryState)(nil)).Elem()
}

type localGoRepositoryArgs struct {
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

// The set of arguments for constructing a LocalGoRepository resource.
type LocalGoRepositoryArgs struct {
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

func (LocalGoRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localGoRepositoryArgs)(nil)).Elem()
}

type LocalGoRepositoryInput interface {
	pulumi.Input

	ToLocalGoRepositoryOutput() LocalGoRepositoryOutput
	ToLocalGoRepositoryOutputWithContext(ctx context.Context) LocalGoRepositoryOutput
}

func (*LocalGoRepository) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalGoRepository)(nil))
}

func (i *LocalGoRepository) ToLocalGoRepositoryOutput() LocalGoRepositoryOutput {
	return i.ToLocalGoRepositoryOutputWithContext(context.Background())
}

func (i *LocalGoRepository) ToLocalGoRepositoryOutputWithContext(ctx context.Context) LocalGoRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalGoRepositoryOutput)
}

func (i *LocalGoRepository) ToLocalGoRepositoryPtrOutput() LocalGoRepositoryPtrOutput {
	return i.ToLocalGoRepositoryPtrOutputWithContext(context.Background())
}

func (i *LocalGoRepository) ToLocalGoRepositoryPtrOutputWithContext(ctx context.Context) LocalGoRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalGoRepositoryPtrOutput)
}

type LocalGoRepositoryPtrInput interface {
	pulumi.Input

	ToLocalGoRepositoryPtrOutput() LocalGoRepositoryPtrOutput
	ToLocalGoRepositoryPtrOutputWithContext(ctx context.Context) LocalGoRepositoryPtrOutput
}

type localGoRepositoryPtrType LocalGoRepositoryArgs

func (*localGoRepositoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalGoRepository)(nil))
}

func (i *localGoRepositoryPtrType) ToLocalGoRepositoryPtrOutput() LocalGoRepositoryPtrOutput {
	return i.ToLocalGoRepositoryPtrOutputWithContext(context.Background())
}

func (i *localGoRepositoryPtrType) ToLocalGoRepositoryPtrOutputWithContext(ctx context.Context) LocalGoRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalGoRepositoryPtrOutput)
}

// LocalGoRepositoryArrayInput is an input type that accepts LocalGoRepositoryArray and LocalGoRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalGoRepositoryArrayInput` via:
//
//          LocalGoRepositoryArray{ LocalGoRepositoryArgs{...} }
type LocalGoRepositoryArrayInput interface {
	pulumi.Input

	ToLocalGoRepositoryArrayOutput() LocalGoRepositoryArrayOutput
	ToLocalGoRepositoryArrayOutputWithContext(context.Context) LocalGoRepositoryArrayOutput
}

type LocalGoRepositoryArray []LocalGoRepositoryInput

func (LocalGoRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalGoRepository)(nil)).Elem()
}

func (i LocalGoRepositoryArray) ToLocalGoRepositoryArrayOutput() LocalGoRepositoryArrayOutput {
	return i.ToLocalGoRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalGoRepositoryArray) ToLocalGoRepositoryArrayOutputWithContext(ctx context.Context) LocalGoRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalGoRepositoryArrayOutput)
}

// LocalGoRepositoryMapInput is an input type that accepts LocalGoRepositoryMap and LocalGoRepositoryMapOutput values.
// You can construct a concrete instance of `LocalGoRepositoryMapInput` via:
//
//          LocalGoRepositoryMap{ "key": LocalGoRepositoryArgs{...} }
type LocalGoRepositoryMapInput interface {
	pulumi.Input

	ToLocalGoRepositoryMapOutput() LocalGoRepositoryMapOutput
	ToLocalGoRepositoryMapOutputWithContext(context.Context) LocalGoRepositoryMapOutput
}

type LocalGoRepositoryMap map[string]LocalGoRepositoryInput

func (LocalGoRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalGoRepository)(nil)).Elem()
}

func (i LocalGoRepositoryMap) ToLocalGoRepositoryMapOutput() LocalGoRepositoryMapOutput {
	return i.ToLocalGoRepositoryMapOutputWithContext(context.Background())
}

func (i LocalGoRepositoryMap) ToLocalGoRepositoryMapOutputWithContext(ctx context.Context) LocalGoRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalGoRepositoryMapOutput)
}

type LocalGoRepositoryOutput struct{ *pulumi.OutputState }

func (LocalGoRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalGoRepository)(nil))
}

func (o LocalGoRepositoryOutput) ToLocalGoRepositoryOutput() LocalGoRepositoryOutput {
	return o
}

func (o LocalGoRepositoryOutput) ToLocalGoRepositoryOutputWithContext(ctx context.Context) LocalGoRepositoryOutput {
	return o
}

func (o LocalGoRepositoryOutput) ToLocalGoRepositoryPtrOutput() LocalGoRepositoryPtrOutput {
	return o.ToLocalGoRepositoryPtrOutputWithContext(context.Background())
}

func (o LocalGoRepositoryOutput) ToLocalGoRepositoryPtrOutputWithContext(ctx context.Context) LocalGoRepositoryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LocalGoRepository) *LocalGoRepository {
		return &v
	}).(LocalGoRepositoryPtrOutput)
}

type LocalGoRepositoryPtrOutput struct{ *pulumi.OutputState }

func (LocalGoRepositoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalGoRepository)(nil))
}

func (o LocalGoRepositoryPtrOutput) ToLocalGoRepositoryPtrOutput() LocalGoRepositoryPtrOutput {
	return o
}

func (o LocalGoRepositoryPtrOutput) ToLocalGoRepositoryPtrOutputWithContext(ctx context.Context) LocalGoRepositoryPtrOutput {
	return o
}

func (o LocalGoRepositoryPtrOutput) Elem() LocalGoRepositoryOutput {
	return o.ApplyT(func(v *LocalGoRepository) LocalGoRepository {
		if v != nil {
			return *v
		}
		var ret LocalGoRepository
		return ret
	}).(LocalGoRepositoryOutput)
}

type LocalGoRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalGoRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LocalGoRepository)(nil))
}

func (o LocalGoRepositoryArrayOutput) ToLocalGoRepositoryArrayOutput() LocalGoRepositoryArrayOutput {
	return o
}

func (o LocalGoRepositoryArrayOutput) ToLocalGoRepositoryArrayOutputWithContext(ctx context.Context) LocalGoRepositoryArrayOutput {
	return o
}

func (o LocalGoRepositoryArrayOutput) Index(i pulumi.IntInput) LocalGoRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LocalGoRepository {
		return vs[0].([]LocalGoRepository)[vs[1].(int)]
	}).(LocalGoRepositoryOutput)
}

type LocalGoRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalGoRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LocalGoRepository)(nil))
}

func (o LocalGoRepositoryMapOutput) ToLocalGoRepositoryMapOutput() LocalGoRepositoryMapOutput {
	return o
}

func (o LocalGoRepositoryMapOutput) ToLocalGoRepositoryMapOutputWithContext(ctx context.Context) LocalGoRepositoryMapOutput {
	return o
}

func (o LocalGoRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalGoRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LocalGoRepository {
		return vs[0].(map[string]LocalGoRepository)[vs[1].(string)]
	}).(LocalGoRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalGoRepositoryInput)(nil)).Elem(), &LocalGoRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalGoRepositoryPtrInput)(nil)).Elem(), &LocalGoRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalGoRepositoryArrayInput)(nil)).Elem(), LocalGoRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalGoRepositoryMapInput)(nil)).Elem(), LocalGoRepositoryMap{})
	pulumi.RegisterOutputType(LocalGoRepositoryOutput{})
	pulumi.RegisterOutputType(LocalGoRepositoryPtrOutput{})
	pulumi.RegisterOutputType(LocalGoRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalGoRepositoryMapOutput{})
}
