// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Local Gems Repository Resource
//
// Creates a local gems repository.
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
// 		_, err := artifactory.NewLocalGemsRepository(ctx, "terraform-local-test-gems-repo", &artifactory.LocalGemsRepositoryArgs{
// 			Key: pulumi.String("terraform-local-test-gems-repo"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type LocalGemsRepository struct {
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
	PriorityResolution pulumi.BoolPtrOutput `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    pulumi.StringPtrOutput   `pulumi:"projectKey"`
	PropertySets  pulumi.StringArrayOutput `pulumi:"propertySets"`
	RepoLayoutRef pulumi.StringOutput      `pulumi:"repoLayoutRef"`
	XrayIndex     pulumi.BoolOutput        `pulumi:"xrayIndex"`
}

// NewLocalGemsRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalGemsRepository(ctx *pulumi.Context,
	name string, args *LocalGemsRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalGemsRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource LocalGemsRepository
	err := ctx.RegisterResource("artifactory:index/localGemsRepository:LocalGemsRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalGemsRepository gets an existing LocalGemsRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalGemsRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalGemsRepositoryState, opts ...pulumi.ResourceOption) (*LocalGemsRepository, error) {
	var resource LocalGemsRepository
	err := ctx.ReadResource("artifactory:index/localGemsRepository:LocalGemsRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalGemsRepository resources.
type localGemsRepositoryState struct {
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
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    *string  `pulumi:"projectKey"`
	PropertySets  []string `pulumi:"propertySets"`
	RepoLayoutRef *string  `pulumi:"repoLayoutRef"`
	XrayIndex     *bool    `pulumi:"xrayIndex"`
}

type LocalGemsRepositoryState struct {
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
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    pulumi.StringPtrInput
	PropertySets  pulumi.StringArrayInput
	RepoLayoutRef pulumi.StringPtrInput
	XrayIndex     pulumi.BoolPtrInput
}

func (LocalGemsRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localGemsRepositoryState)(nil)).Elem()
}

type localGemsRepositoryArgs struct {
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
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    *string  `pulumi:"projectKey"`
	PropertySets  []string `pulumi:"propertySets"`
	RepoLayoutRef *string  `pulumi:"repoLayoutRef"`
	XrayIndex     *bool    `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a LocalGemsRepository resource.
type LocalGemsRepositoryArgs struct {
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
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    pulumi.StringPtrInput
	PropertySets  pulumi.StringArrayInput
	RepoLayoutRef pulumi.StringPtrInput
	XrayIndex     pulumi.BoolPtrInput
}

func (LocalGemsRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localGemsRepositoryArgs)(nil)).Elem()
}

type LocalGemsRepositoryInput interface {
	pulumi.Input

	ToLocalGemsRepositoryOutput() LocalGemsRepositoryOutput
	ToLocalGemsRepositoryOutputWithContext(ctx context.Context) LocalGemsRepositoryOutput
}

func (*LocalGemsRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalGemsRepository)(nil)).Elem()
}

func (i *LocalGemsRepository) ToLocalGemsRepositoryOutput() LocalGemsRepositoryOutput {
	return i.ToLocalGemsRepositoryOutputWithContext(context.Background())
}

func (i *LocalGemsRepository) ToLocalGemsRepositoryOutputWithContext(ctx context.Context) LocalGemsRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalGemsRepositoryOutput)
}

// LocalGemsRepositoryArrayInput is an input type that accepts LocalGemsRepositoryArray and LocalGemsRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalGemsRepositoryArrayInput` via:
//
//          LocalGemsRepositoryArray{ LocalGemsRepositoryArgs{...} }
type LocalGemsRepositoryArrayInput interface {
	pulumi.Input

	ToLocalGemsRepositoryArrayOutput() LocalGemsRepositoryArrayOutput
	ToLocalGemsRepositoryArrayOutputWithContext(context.Context) LocalGemsRepositoryArrayOutput
}

type LocalGemsRepositoryArray []LocalGemsRepositoryInput

func (LocalGemsRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalGemsRepository)(nil)).Elem()
}

func (i LocalGemsRepositoryArray) ToLocalGemsRepositoryArrayOutput() LocalGemsRepositoryArrayOutput {
	return i.ToLocalGemsRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalGemsRepositoryArray) ToLocalGemsRepositoryArrayOutputWithContext(ctx context.Context) LocalGemsRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalGemsRepositoryArrayOutput)
}

// LocalGemsRepositoryMapInput is an input type that accepts LocalGemsRepositoryMap and LocalGemsRepositoryMapOutput values.
// You can construct a concrete instance of `LocalGemsRepositoryMapInput` via:
//
//          LocalGemsRepositoryMap{ "key": LocalGemsRepositoryArgs{...} }
type LocalGemsRepositoryMapInput interface {
	pulumi.Input

	ToLocalGemsRepositoryMapOutput() LocalGemsRepositoryMapOutput
	ToLocalGemsRepositoryMapOutputWithContext(context.Context) LocalGemsRepositoryMapOutput
}

type LocalGemsRepositoryMap map[string]LocalGemsRepositoryInput

func (LocalGemsRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalGemsRepository)(nil)).Elem()
}

func (i LocalGemsRepositoryMap) ToLocalGemsRepositoryMapOutput() LocalGemsRepositoryMapOutput {
	return i.ToLocalGemsRepositoryMapOutputWithContext(context.Background())
}

func (i LocalGemsRepositoryMap) ToLocalGemsRepositoryMapOutputWithContext(ctx context.Context) LocalGemsRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalGemsRepositoryMapOutput)
}

type LocalGemsRepositoryOutput struct{ *pulumi.OutputState }

func (LocalGemsRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalGemsRepository)(nil)).Elem()
}

func (o LocalGemsRepositoryOutput) ToLocalGemsRepositoryOutput() LocalGemsRepositoryOutput {
	return o
}

func (o LocalGemsRepositoryOutput) ToLocalGemsRepositoryOutputWithContext(ctx context.Context) LocalGemsRepositoryOutput {
	return o
}

type LocalGemsRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalGemsRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalGemsRepository)(nil)).Elem()
}

func (o LocalGemsRepositoryArrayOutput) ToLocalGemsRepositoryArrayOutput() LocalGemsRepositoryArrayOutput {
	return o
}

func (o LocalGemsRepositoryArrayOutput) ToLocalGemsRepositoryArrayOutputWithContext(ctx context.Context) LocalGemsRepositoryArrayOutput {
	return o
}

func (o LocalGemsRepositoryArrayOutput) Index(i pulumi.IntInput) LocalGemsRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalGemsRepository {
		return vs[0].([]*LocalGemsRepository)[vs[1].(int)]
	}).(LocalGemsRepositoryOutput)
}

type LocalGemsRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalGemsRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalGemsRepository)(nil)).Elem()
}

func (o LocalGemsRepositoryMapOutput) ToLocalGemsRepositoryMapOutput() LocalGemsRepositoryMapOutput {
	return o
}

func (o LocalGemsRepositoryMapOutput) ToLocalGemsRepositoryMapOutputWithContext(ctx context.Context) LocalGemsRepositoryMapOutput {
	return o
}

func (o LocalGemsRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalGemsRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalGemsRepository {
		return vs[0].(map[string]*LocalGemsRepository)[vs[1].(string)]
	}).(LocalGemsRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalGemsRepositoryInput)(nil)).Elem(), &LocalGemsRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalGemsRepositoryArrayInput)(nil)).Elem(), LocalGemsRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalGemsRepositoryMapInput)(nil)).Elem(), LocalGemsRepositoryMap{})
	pulumi.RegisterOutputType(LocalGemsRepositoryOutput{})
	pulumi.RegisterOutputType(LocalGemsRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalGemsRepositoryMapOutput{})
}
