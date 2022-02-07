// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Local Puppet Repository Resource
//
// Creates a local puppet repository.
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
// 		_, err := artifactory.NewLocalPuppetRepository(ctx, "terraform_local_test_puppet_repo", &artifactory.LocalPuppetRepositoryArgs{
// 			Key: pulumi.String("terraform-local-test-puppet-repo"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type LocalPuppetRepository struct {
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

// NewLocalPuppetRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalPuppetRepository(ctx *pulumi.Context,
	name string, args *LocalPuppetRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalPuppetRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource LocalPuppetRepository
	err := ctx.RegisterResource("artifactory:index/localPuppetRepository:LocalPuppetRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalPuppetRepository gets an existing LocalPuppetRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalPuppetRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalPuppetRepositoryState, opts ...pulumi.ResourceOption) (*LocalPuppetRepository, error) {
	var resource LocalPuppetRepository
	err := ctx.ReadResource("artifactory:index/localPuppetRepository:LocalPuppetRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalPuppetRepository resources.
type localPuppetRepositoryState struct {
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

type LocalPuppetRepositoryState struct {
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

func (LocalPuppetRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localPuppetRepositoryState)(nil)).Elem()
}

type localPuppetRepositoryArgs struct {
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

// The set of arguments for constructing a LocalPuppetRepository resource.
type LocalPuppetRepositoryArgs struct {
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

func (LocalPuppetRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localPuppetRepositoryArgs)(nil)).Elem()
}

type LocalPuppetRepositoryInput interface {
	pulumi.Input

	ToLocalPuppetRepositoryOutput() LocalPuppetRepositoryOutput
	ToLocalPuppetRepositoryOutputWithContext(ctx context.Context) LocalPuppetRepositoryOutput
}

func (*LocalPuppetRepository) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalPuppetRepository)(nil))
}

func (i *LocalPuppetRepository) ToLocalPuppetRepositoryOutput() LocalPuppetRepositoryOutput {
	return i.ToLocalPuppetRepositoryOutputWithContext(context.Background())
}

func (i *LocalPuppetRepository) ToLocalPuppetRepositoryOutputWithContext(ctx context.Context) LocalPuppetRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalPuppetRepositoryOutput)
}

func (i *LocalPuppetRepository) ToLocalPuppetRepositoryPtrOutput() LocalPuppetRepositoryPtrOutput {
	return i.ToLocalPuppetRepositoryPtrOutputWithContext(context.Background())
}

func (i *LocalPuppetRepository) ToLocalPuppetRepositoryPtrOutputWithContext(ctx context.Context) LocalPuppetRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalPuppetRepositoryPtrOutput)
}

type LocalPuppetRepositoryPtrInput interface {
	pulumi.Input

	ToLocalPuppetRepositoryPtrOutput() LocalPuppetRepositoryPtrOutput
	ToLocalPuppetRepositoryPtrOutputWithContext(ctx context.Context) LocalPuppetRepositoryPtrOutput
}

type localPuppetRepositoryPtrType LocalPuppetRepositoryArgs

func (*localPuppetRepositoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalPuppetRepository)(nil))
}

func (i *localPuppetRepositoryPtrType) ToLocalPuppetRepositoryPtrOutput() LocalPuppetRepositoryPtrOutput {
	return i.ToLocalPuppetRepositoryPtrOutputWithContext(context.Background())
}

func (i *localPuppetRepositoryPtrType) ToLocalPuppetRepositoryPtrOutputWithContext(ctx context.Context) LocalPuppetRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalPuppetRepositoryPtrOutput)
}

// LocalPuppetRepositoryArrayInput is an input type that accepts LocalPuppetRepositoryArray and LocalPuppetRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalPuppetRepositoryArrayInput` via:
//
//          LocalPuppetRepositoryArray{ LocalPuppetRepositoryArgs{...} }
type LocalPuppetRepositoryArrayInput interface {
	pulumi.Input

	ToLocalPuppetRepositoryArrayOutput() LocalPuppetRepositoryArrayOutput
	ToLocalPuppetRepositoryArrayOutputWithContext(context.Context) LocalPuppetRepositoryArrayOutput
}

type LocalPuppetRepositoryArray []LocalPuppetRepositoryInput

func (LocalPuppetRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalPuppetRepository)(nil)).Elem()
}

func (i LocalPuppetRepositoryArray) ToLocalPuppetRepositoryArrayOutput() LocalPuppetRepositoryArrayOutput {
	return i.ToLocalPuppetRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalPuppetRepositoryArray) ToLocalPuppetRepositoryArrayOutputWithContext(ctx context.Context) LocalPuppetRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalPuppetRepositoryArrayOutput)
}

// LocalPuppetRepositoryMapInput is an input type that accepts LocalPuppetRepositoryMap and LocalPuppetRepositoryMapOutput values.
// You can construct a concrete instance of `LocalPuppetRepositoryMapInput` via:
//
//          LocalPuppetRepositoryMap{ "key": LocalPuppetRepositoryArgs{...} }
type LocalPuppetRepositoryMapInput interface {
	pulumi.Input

	ToLocalPuppetRepositoryMapOutput() LocalPuppetRepositoryMapOutput
	ToLocalPuppetRepositoryMapOutputWithContext(context.Context) LocalPuppetRepositoryMapOutput
}

type LocalPuppetRepositoryMap map[string]LocalPuppetRepositoryInput

func (LocalPuppetRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalPuppetRepository)(nil)).Elem()
}

func (i LocalPuppetRepositoryMap) ToLocalPuppetRepositoryMapOutput() LocalPuppetRepositoryMapOutput {
	return i.ToLocalPuppetRepositoryMapOutputWithContext(context.Background())
}

func (i LocalPuppetRepositoryMap) ToLocalPuppetRepositoryMapOutputWithContext(ctx context.Context) LocalPuppetRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalPuppetRepositoryMapOutput)
}

type LocalPuppetRepositoryOutput struct{ *pulumi.OutputState }

func (LocalPuppetRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalPuppetRepository)(nil))
}

func (o LocalPuppetRepositoryOutput) ToLocalPuppetRepositoryOutput() LocalPuppetRepositoryOutput {
	return o
}

func (o LocalPuppetRepositoryOutput) ToLocalPuppetRepositoryOutputWithContext(ctx context.Context) LocalPuppetRepositoryOutput {
	return o
}

func (o LocalPuppetRepositoryOutput) ToLocalPuppetRepositoryPtrOutput() LocalPuppetRepositoryPtrOutput {
	return o.ToLocalPuppetRepositoryPtrOutputWithContext(context.Background())
}

func (o LocalPuppetRepositoryOutput) ToLocalPuppetRepositoryPtrOutputWithContext(ctx context.Context) LocalPuppetRepositoryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LocalPuppetRepository) *LocalPuppetRepository {
		return &v
	}).(LocalPuppetRepositoryPtrOutput)
}

type LocalPuppetRepositoryPtrOutput struct{ *pulumi.OutputState }

func (LocalPuppetRepositoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalPuppetRepository)(nil))
}

func (o LocalPuppetRepositoryPtrOutput) ToLocalPuppetRepositoryPtrOutput() LocalPuppetRepositoryPtrOutput {
	return o
}

func (o LocalPuppetRepositoryPtrOutput) ToLocalPuppetRepositoryPtrOutputWithContext(ctx context.Context) LocalPuppetRepositoryPtrOutput {
	return o
}

func (o LocalPuppetRepositoryPtrOutput) Elem() LocalPuppetRepositoryOutput {
	return o.ApplyT(func(v *LocalPuppetRepository) LocalPuppetRepository {
		if v != nil {
			return *v
		}
		var ret LocalPuppetRepository
		return ret
	}).(LocalPuppetRepositoryOutput)
}

type LocalPuppetRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalPuppetRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LocalPuppetRepository)(nil))
}

func (o LocalPuppetRepositoryArrayOutput) ToLocalPuppetRepositoryArrayOutput() LocalPuppetRepositoryArrayOutput {
	return o
}

func (o LocalPuppetRepositoryArrayOutput) ToLocalPuppetRepositoryArrayOutputWithContext(ctx context.Context) LocalPuppetRepositoryArrayOutput {
	return o
}

func (o LocalPuppetRepositoryArrayOutput) Index(i pulumi.IntInput) LocalPuppetRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LocalPuppetRepository {
		return vs[0].([]LocalPuppetRepository)[vs[1].(int)]
	}).(LocalPuppetRepositoryOutput)
}

type LocalPuppetRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalPuppetRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LocalPuppetRepository)(nil))
}

func (o LocalPuppetRepositoryMapOutput) ToLocalPuppetRepositoryMapOutput() LocalPuppetRepositoryMapOutput {
	return o
}

func (o LocalPuppetRepositoryMapOutput) ToLocalPuppetRepositoryMapOutputWithContext(ctx context.Context) LocalPuppetRepositoryMapOutput {
	return o
}

func (o LocalPuppetRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalPuppetRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LocalPuppetRepository {
		return vs[0].(map[string]LocalPuppetRepository)[vs[1].(string)]
	}).(LocalPuppetRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalPuppetRepositoryInput)(nil)).Elem(), &LocalPuppetRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalPuppetRepositoryPtrInput)(nil)).Elem(), &LocalPuppetRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalPuppetRepositoryArrayInput)(nil)).Elem(), LocalPuppetRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalPuppetRepositoryMapInput)(nil)).Elem(), LocalPuppetRepositoryMap{})
	pulumi.RegisterOutputType(LocalPuppetRepositoryOutput{})
	pulumi.RegisterOutputType(LocalPuppetRepositoryPtrOutput{})
	pulumi.RegisterOutputType(LocalPuppetRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalPuppetRepositoryMapOutput{})
}
