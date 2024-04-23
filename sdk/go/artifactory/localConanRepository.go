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

// Creates a local Conan repository.
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
//			_, err := artifactory.NewLocalConanRepository(ctx, "terraform-local-test-conan-repo", &artifactory.LocalConanRepositoryArgs{
//				Key: pulumi.String("terraform-local-test-conan-repo"),
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
// $ pulumi import artifactory:index/localConanRepository:LocalConanRepository terraform-local-test-conan-repo terraform-local-test-conan-repo
// ```
type LocalConanRepository struct {
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
	// Force basic authentication credentials in order to use this repository. Default value is `false`.
	ForceConanAuthentication pulumi.BoolPtrOutput `pulumi:"forceConanAuthentication"`
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

// NewLocalConanRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalConanRepository(ctx *pulumi.Context,
	name string, args *LocalConanRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalConanRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LocalConanRepository
	err := ctx.RegisterResource("artifactory:index/localConanRepository:LocalConanRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalConanRepository gets an existing LocalConanRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalConanRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalConanRepositoryState, opts ...pulumi.ResourceOption) (*LocalConanRepository, error) {
	var resource LocalConanRepository
	err := ctx.ReadResource("artifactory:index/localConanRepository:LocalConanRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalConanRepository resources.
type localConanRepositoryState struct {
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
	// Force basic authentication credentials in order to use this repository. Default value is `false`.
	ForceConanAuthentication *bool `pulumi:"forceConanAuthentication"`
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

type LocalConanRepositoryState struct {
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
	// Force basic authentication credentials in order to use this repository. Default value is `false`.
	ForceConanAuthentication pulumi.BoolPtrInput
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

func (LocalConanRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localConanRepositoryState)(nil)).Elem()
}

type localConanRepositoryArgs struct {
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
	// Force basic authentication credentials in order to use this repository. Default value is `false`.
	ForceConanAuthentication *bool `pulumi:"forceConanAuthentication"`
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

// The set of arguments for constructing a LocalConanRepository resource.
type LocalConanRepositoryArgs struct {
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
	// Force basic authentication credentials in order to use this repository. Default value is `false`.
	ForceConanAuthentication pulumi.BoolPtrInput
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

func (LocalConanRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localConanRepositoryArgs)(nil)).Elem()
}

type LocalConanRepositoryInput interface {
	pulumi.Input

	ToLocalConanRepositoryOutput() LocalConanRepositoryOutput
	ToLocalConanRepositoryOutputWithContext(ctx context.Context) LocalConanRepositoryOutput
}

func (*LocalConanRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalConanRepository)(nil)).Elem()
}

func (i *LocalConanRepository) ToLocalConanRepositoryOutput() LocalConanRepositoryOutput {
	return i.ToLocalConanRepositoryOutputWithContext(context.Background())
}

func (i *LocalConanRepository) ToLocalConanRepositoryOutputWithContext(ctx context.Context) LocalConanRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalConanRepositoryOutput)
}

// LocalConanRepositoryArrayInput is an input type that accepts LocalConanRepositoryArray and LocalConanRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalConanRepositoryArrayInput` via:
//
//	LocalConanRepositoryArray{ LocalConanRepositoryArgs{...} }
type LocalConanRepositoryArrayInput interface {
	pulumi.Input

	ToLocalConanRepositoryArrayOutput() LocalConanRepositoryArrayOutput
	ToLocalConanRepositoryArrayOutputWithContext(context.Context) LocalConanRepositoryArrayOutput
}

type LocalConanRepositoryArray []LocalConanRepositoryInput

func (LocalConanRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalConanRepository)(nil)).Elem()
}

func (i LocalConanRepositoryArray) ToLocalConanRepositoryArrayOutput() LocalConanRepositoryArrayOutput {
	return i.ToLocalConanRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalConanRepositoryArray) ToLocalConanRepositoryArrayOutputWithContext(ctx context.Context) LocalConanRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalConanRepositoryArrayOutput)
}

// LocalConanRepositoryMapInput is an input type that accepts LocalConanRepositoryMap and LocalConanRepositoryMapOutput values.
// You can construct a concrete instance of `LocalConanRepositoryMapInput` via:
//
//	LocalConanRepositoryMap{ "key": LocalConanRepositoryArgs{...} }
type LocalConanRepositoryMapInput interface {
	pulumi.Input

	ToLocalConanRepositoryMapOutput() LocalConanRepositoryMapOutput
	ToLocalConanRepositoryMapOutputWithContext(context.Context) LocalConanRepositoryMapOutput
}

type LocalConanRepositoryMap map[string]LocalConanRepositoryInput

func (LocalConanRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalConanRepository)(nil)).Elem()
}

func (i LocalConanRepositoryMap) ToLocalConanRepositoryMapOutput() LocalConanRepositoryMapOutput {
	return i.ToLocalConanRepositoryMapOutputWithContext(context.Background())
}

func (i LocalConanRepositoryMap) ToLocalConanRepositoryMapOutputWithContext(ctx context.Context) LocalConanRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalConanRepositoryMapOutput)
}

type LocalConanRepositoryOutput struct{ *pulumi.OutputState }

func (LocalConanRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalConanRepository)(nil)).Elem()
}

func (o LocalConanRepositoryOutput) ToLocalConanRepositoryOutput() LocalConanRepositoryOutput {
	return o
}

func (o LocalConanRepositoryOutput) ToLocalConanRepositoryOutputWithContext(ctx context.Context) LocalConanRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o LocalConanRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalConanRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o LocalConanRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalConanRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o LocalConanRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalConanRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

// Public description.
func (o LocalConanRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalConanRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o LocalConanRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalConanRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o LocalConanRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalConanRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// Force basic authentication credentials in order to use this repository. Default value is `false`.
func (o LocalConanRepositoryOutput) ForceConanAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalConanRepository) pulumi.BoolPtrOutput { return v.ForceConanAuthentication }).(pulumi.BoolPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o LocalConanRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalConanRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// the identity key of the repo.
func (o LocalConanRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalConanRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o LocalConanRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalConanRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LocalConanRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalConanRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o LocalConanRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalConanRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o LocalConanRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalConanRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o LocalConanRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalConanRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o LocalConanRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalConanRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o LocalConanRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalConanRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o LocalConanRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalConanRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type LocalConanRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalConanRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalConanRepository)(nil)).Elem()
}

func (o LocalConanRepositoryArrayOutput) ToLocalConanRepositoryArrayOutput() LocalConanRepositoryArrayOutput {
	return o
}

func (o LocalConanRepositoryArrayOutput) ToLocalConanRepositoryArrayOutputWithContext(ctx context.Context) LocalConanRepositoryArrayOutput {
	return o
}

func (o LocalConanRepositoryArrayOutput) Index(i pulumi.IntInput) LocalConanRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalConanRepository {
		return vs[0].([]*LocalConanRepository)[vs[1].(int)]
	}).(LocalConanRepositoryOutput)
}

type LocalConanRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalConanRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalConanRepository)(nil)).Elem()
}

func (o LocalConanRepositoryMapOutput) ToLocalConanRepositoryMapOutput() LocalConanRepositoryMapOutput {
	return o
}

func (o LocalConanRepositoryMapOutput) ToLocalConanRepositoryMapOutputWithContext(ctx context.Context) LocalConanRepositoryMapOutput {
	return o
}

func (o LocalConanRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalConanRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalConanRepository {
		return vs[0].(map[string]*LocalConanRepository)[vs[1].(string)]
	}).(LocalConanRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalConanRepositoryInput)(nil)).Elem(), &LocalConanRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalConanRepositoryArrayInput)(nil)).Elem(), LocalConanRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalConanRepositoryMapInput)(nil)).Elem(), LocalConanRepositoryMap{})
	pulumi.RegisterOutputType(LocalConanRepositoryOutput{})
	pulumi.RegisterOutputType(LocalConanRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalConanRepositoryMapOutput{})
}
