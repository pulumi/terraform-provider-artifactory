// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Local NPM Repository Resource
//
// Creates a local npm repository.
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
// 		_, err := artifactory.NewLocalNpmRepository(ctx, "terraform_local_test_npm_repo", &artifactory.LocalNpmRepositoryArgs{
// 			Key: pulumi.String("terraform-local-test-npm-repo"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type LocalNpmRepository struct {
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

// NewLocalNpmRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalNpmRepository(ctx *pulumi.Context,
	name string, args *LocalNpmRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalNpmRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource LocalNpmRepository
	err := ctx.RegisterResource("artifactory:index/localNpmRepository:LocalNpmRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalNpmRepository gets an existing LocalNpmRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalNpmRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalNpmRepositoryState, opts ...pulumi.ResourceOption) (*LocalNpmRepository, error) {
	var resource LocalNpmRepository
	err := ctx.ReadResource("artifactory:index/localNpmRepository:LocalNpmRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalNpmRepository resources.
type localNpmRepositoryState struct {
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

type LocalNpmRepositoryState struct {
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

func (LocalNpmRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localNpmRepositoryState)(nil)).Elem()
}

type localNpmRepositoryArgs struct {
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

// The set of arguments for constructing a LocalNpmRepository resource.
type LocalNpmRepositoryArgs struct {
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

func (LocalNpmRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localNpmRepositoryArgs)(nil)).Elem()
}

type LocalNpmRepositoryInput interface {
	pulumi.Input

	ToLocalNpmRepositoryOutput() LocalNpmRepositoryOutput
	ToLocalNpmRepositoryOutputWithContext(ctx context.Context) LocalNpmRepositoryOutput
}

func (*LocalNpmRepository) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalNpmRepository)(nil))
}

func (i *LocalNpmRepository) ToLocalNpmRepositoryOutput() LocalNpmRepositoryOutput {
	return i.ToLocalNpmRepositoryOutputWithContext(context.Background())
}

func (i *LocalNpmRepository) ToLocalNpmRepositoryOutputWithContext(ctx context.Context) LocalNpmRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalNpmRepositoryOutput)
}

func (i *LocalNpmRepository) ToLocalNpmRepositoryPtrOutput() LocalNpmRepositoryPtrOutput {
	return i.ToLocalNpmRepositoryPtrOutputWithContext(context.Background())
}

func (i *LocalNpmRepository) ToLocalNpmRepositoryPtrOutputWithContext(ctx context.Context) LocalNpmRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalNpmRepositoryPtrOutput)
}

type LocalNpmRepositoryPtrInput interface {
	pulumi.Input

	ToLocalNpmRepositoryPtrOutput() LocalNpmRepositoryPtrOutput
	ToLocalNpmRepositoryPtrOutputWithContext(ctx context.Context) LocalNpmRepositoryPtrOutput
}

type localNpmRepositoryPtrType LocalNpmRepositoryArgs

func (*localNpmRepositoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalNpmRepository)(nil))
}

func (i *localNpmRepositoryPtrType) ToLocalNpmRepositoryPtrOutput() LocalNpmRepositoryPtrOutput {
	return i.ToLocalNpmRepositoryPtrOutputWithContext(context.Background())
}

func (i *localNpmRepositoryPtrType) ToLocalNpmRepositoryPtrOutputWithContext(ctx context.Context) LocalNpmRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalNpmRepositoryPtrOutput)
}

// LocalNpmRepositoryArrayInput is an input type that accepts LocalNpmRepositoryArray and LocalNpmRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalNpmRepositoryArrayInput` via:
//
//          LocalNpmRepositoryArray{ LocalNpmRepositoryArgs{...} }
type LocalNpmRepositoryArrayInput interface {
	pulumi.Input

	ToLocalNpmRepositoryArrayOutput() LocalNpmRepositoryArrayOutput
	ToLocalNpmRepositoryArrayOutputWithContext(context.Context) LocalNpmRepositoryArrayOutput
}

type LocalNpmRepositoryArray []LocalNpmRepositoryInput

func (LocalNpmRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalNpmRepository)(nil)).Elem()
}

func (i LocalNpmRepositoryArray) ToLocalNpmRepositoryArrayOutput() LocalNpmRepositoryArrayOutput {
	return i.ToLocalNpmRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalNpmRepositoryArray) ToLocalNpmRepositoryArrayOutputWithContext(ctx context.Context) LocalNpmRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalNpmRepositoryArrayOutput)
}

// LocalNpmRepositoryMapInput is an input type that accepts LocalNpmRepositoryMap and LocalNpmRepositoryMapOutput values.
// You can construct a concrete instance of `LocalNpmRepositoryMapInput` via:
//
//          LocalNpmRepositoryMap{ "key": LocalNpmRepositoryArgs{...} }
type LocalNpmRepositoryMapInput interface {
	pulumi.Input

	ToLocalNpmRepositoryMapOutput() LocalNpmRepositoryMapOutput
	ToLocalNpmRepositoryMapOutputWithContext(context.Context) LocalNpmRepositoryMapOutput
}

type LocalNpmRepositoryMap map[string]LocalNpmRepositoryInput

func (LocalNpmRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalNpmRepository)(nil)).Elem()
}

func (i LocalNpmRepositoryMap) ToLocalNpmRepositoryMapOutput() LocalNpmRepositoryMapOutput {
	return i.ToLocalNpmRepositoryMapOutputWithContext(context.Background())
}

func (i LocalNpmRepositoryMap) ToLocalNpmRepositoryMapOutputWithContext(ctx context.Context) LocalNpmRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalNpmRepositoryMapOutput)
}

type LocalNpmRepositoryOutput struct{ *pulumi.OutputState }

func (LocalNpmRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalNpmRepository)(nil))
}

func (o LocalNpmRepositoryOutput) ToLocalNpmRepositoryOutput() LocalNpmRepositoryOutput {
	return o
}

func (o LocalNpmRepositoryOutput) ToLocalNpmRepositoryOutputWithContext(ctx context.Context) LocalNpmRepositoryOutput {
	return o
}

func (o LocalNpmRepositoryOutput) ToLocalNpmRepositoryPtrOutput() LocalNpmRepositoryPtrOutput {
	return o.ToLocalNpmRepositoryPtrOutputWithContext(context.Background())
}

func (o LocalNpmRepositoryOutput) ToLocalNpmRepositoryPtrOutputWithContext(ctx context.Context) LocalNpmRepositoryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LocalNpmRepository) *LocalNpmRepository {
		return &v
	}).(LocalNpmRepositoryPtrOutput)
}

type LocalNpmRepositoryPtrOutput struct{ *pulumi.OutputState }

func (LocalNpmRepositoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalNpmRepository)(nil))
}

func (o LocalNpmRepositoryPtrOutput) ToLocalNpmRepositoryPtrOutput() LocalNpmRepositoryPtrOutput {
	return o
}

func (o LocalNpmRepositoryPtrOutput) ToLocalNpmRepositoryPtrOutputWithContext(ctx context.Context) LocalNpmRepositoryPtrOutput {
	return o
}

func (o LocalNpmRepositoryPtrOutput) Elem() LocalNpmRepositoryOutput {
	return o.ApplyT(func(v *LocalNpmRepository) LocalNpmRepository {
		if v != nil {
			return *v
		}
		var ret LocalNpmRepository
		return ret
	}).(LocalNpmRepositoryOutput)
}

type LocalNpmRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalNpmRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LocalNpmRepository)(nil))
}

func (o LocalNpmRepositoryArrayOutput) ToLocalNpmRepositoryArrayOutput() LocalNpmRepositoryArrayOutput {
	return o
}

func (o LocalNpmRepositoryArrayOutput) ToLocalNpmRepositoryArrayOutputWithContext(ctx context.Context) LocalNpmRepositoryArrayOutput {
	return o
}

func (o LocalNpmRepositoryArrayOutput) Index(i pulumi.IntInput) LocalNpmRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LocalNpmRepository {
		return vs[0].([]LocalNpmRepository)[vs[1].(int)]
	}).(LocalNpmRepositoryOutput)
}

type LocalNpmRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalNpmRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LocalNpmRepository)(nil))
}

func (o LocalNpmRepositoryMapOutput) ToLocalNpmRepositoryMapOutput() LocalNpmRepositoryMapOutput {
	return o
}

func (o LocalNpmRepositoryMapOutput) ToLocalNpmRepositoryMapOutputWithContext(ctx context.Context) LocalNpmRepositoryMapOutput {
	return o
}

func (o LocalNpmRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalNpmRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LocalNpmRepository {
		return vs[0].(map[string]LocalNpmRepository)[vs[1].(string)]
	}).(LocalNpmRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalNpmRepositoryInput)(nil)).Elem(), &LocalNpmRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalNpmRepositoryPtrInput)(nil)).Elem(), &LocalNpmRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalNpmRepositoryArrayInput)(nil)).Elem(), LocalNpmRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalNpmRepositoryMapInput)(nil)).Elem(), LocalNpmRepositoryMap{})
	pulumi.RegisterOutputType(LocalNpmRepositoryOutput{})
	pulumi.RegisterOutputType(LocalNpmRepositoryPtrOutput{})
	pulumi.RegisterOutputType(LocalNpmRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalNpmRepositoryMapOutput{})
}
