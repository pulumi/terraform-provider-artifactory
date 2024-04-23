// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-artifactory/sdk/v6/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a local Go repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v6/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewLocalGoRepository(ctx, "terraform-local-test-go-repo", &artifactory.LocalGoRepositoryArgs{
//				Key: pulumi.String("terraform-local-test-go-repo"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Local repositories can be imported using their name, e.g.
//
// ```sh
// $ pulumi import artifactory:index/localGoRepository:LocalGoRepository terraform-local-test-go-repo terraform-local-test-go-repo
// ```
type LocalGoRepository struct {
	pulumi.CustomResourceState

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrOutput `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrOutput `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolPtrOutput `pulumi:"cdnRedirect"`
	// Public description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrOutput `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrOutput `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrOutput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key pulumi.StringOutput `pulumi:"key"`
	// Internal description.
	Notes       pulumi.StringPtrOutput `pulumi:"notes"`
	PackageType pulumi.StringOutput    `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrOutput `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
	// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
	// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
	// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrOutput `pulumi:"projectKey"`
	// List of property set name
	PropertySets pulumi.StringArrayOutput `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrOutput `pulumi:"xrayIndex"`
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
	opts = internal.PkgResourceDefaultOpts(opts)
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
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect *bool `pulumi:"cdnRedirect"`
	// Public description.
	Description *string `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect *bool `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key *string `pulumi:"key"`
	// Internal description.
	Notes       *string `pulumi:"notes"`
	PackageType *string `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
	// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
	// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
	// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// List of property set name
	PropertySets []string `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

type LocalGoRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolPtrInput
	// Public description.
	Description pulumi.StringPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// the identity key of the repo.
	Key pulumi.StringPtrInput
	// Internal description.
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
	// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
	// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
	// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// List of property set name
	PropertySets pulumi.StringArrayInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (LocalGoRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localGoRepositoryState)(nil)).Elem()
}

type localGoRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect *bool `pulumi:"cdnRedirect"`
	// Public description.
	Description *string `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect *bool `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// Internal description.
	Notes *string `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
	// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
	// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
	// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// List of property set name
	PropertySets []string `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a LocalGoRepository resource.
type LocalGoRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolPtrInput
	// Public description.
	Description pulumi.StringPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// the identity key of the repo.
	Key pulumi.StringInput
	// Internal description.
	Notes pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
	// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
	// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
	// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// List of property set name
	PropertySets pulumi.StringArrayInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
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
	return reflect.TypeOf((**LocalGoRepository)(nil)).Elem()
}

func (i *LocalGoRepository) ToLocalGoRepositoryOutput() LocalGoRepositoryOutput {
	return i.ToLocalGoRepositoryOutputWithContext(context.Background())
}

func (i *LocalGoRepository) ToLocalGoRepositoryOutputWithContext(ctx context.Context) LocalGoRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalGoRepositoryOutput)
}

// LocalGoRepositoryArrayInput is an input type that accepts LocalGoRepositoryArray and LocalGoRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalGoRepositoryArrayInput` via:
//
//	LocalGoRepositoryArray{ LocalGoRepositoryArgs{...} }
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
//	LocalGoRepositoryMap{ "key": LocalGoRepositoryArgs{...} }
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
	return reflect.TypeOf((**LocalGoRepository)(nil)).Elem()
}

func (o LocalGoRepositoryOutput) ToLocalGoRepositoryOutput() LocalGoRepositoryOutput {
	return o
}

func (o LocalGoRepositoryOutput) ToLocalGoRepositoryOutputWithContext(ctx context.Context) LocalGoRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o LocalGoRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalGoRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o LocalGoRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalGoRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o LocalGoRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalGoRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

// Public description.
func (o LocalGoRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalGoRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o LocalGoRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalGoRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o LocalGoRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalGoRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o LocalGoRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalGoRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// the identity key of the repo.
func (o LocalGoRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalGoRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o LocalGoRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalGoRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LocalGoRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalGoRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o LocalGoRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalGoRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o LocalGoRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalGoRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o LocalGoRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalGoRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o LocalGoRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalGoRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o LocalGoRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalGoRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o LocalGoRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalGoRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type LocalGoRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalGoRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalGoRepository)(nil)).Elem()
}

func (o LocalGoRepositoryArrayOutput) ToLocalGoRepositoryArrayOutput() LocalGoRepositoryArrayOutput {
	return o
}

func (o LocalGoRepositoryArrayOutput) ToLocalGoRepositoryArrayOutputWithContext(ctx context.Context) LocalGoRepositoryArrayOutput {
	return o
}

func (o LocalGoRepositoryArrayOutput) Index(i pulumi.IntInput) LocalGoRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalGoRepository {
		return vs[0].([]*LocalGoRepository)[vs[1].(int)]
	}).(LocalGoRepositoryOutput)
}

type LocalGoRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalGoRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalGoRepository)(nil)).Elem()
}

func (o LocalGoRepositoryMapOutput) ToLocalGoRepositoryMapOutput() LocalGoRepositoryMapOutput {
	return o
}

func (o LocalGoRepositoryMapOutput) ToLocalGoRepositoryMapOutputWithContext(ctx context.Context) LocalGoRepositoryMapOutput {
	return o
}

func (o LocalGoRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalGoRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalGoRepository {
		return vs[0].(map[string]*LocalGoRepository)[vs[1].(string)]
	}).(LocalGoRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalGoRepositoryInput)(nil)).Elem(), &LocalGoRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalGoRepositoryArrayInput)(nil)).Elem(), LocalGoRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalGoRepositoryMapInput)(nil)).Elem(), LocalGoRepositoryMap{})
	pulumi.RegisterOutputType(LocalGoRepositoryOutput{})
	pulumi.RegisterOutputType(LocalGoRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalGoRepositoryMapOutput{})
}
