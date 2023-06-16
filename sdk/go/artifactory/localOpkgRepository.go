// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a local Opkg repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v1/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewLocalOpkgRepository(ctx, "terraform-local-test-opkg-repo", &artifactory.LocalOpkgRepositoryArgs{
//				Key: "terraform-local-test-opkg-repo",
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
//
//	$ pulumi import artifactory:index/localOpkgRepository:LocalOpkgRepository terraform-local-test-opkg-repo terraform-local-test-opkg-repo
//
// ```
type LocalOpkgRepository struct {
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
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringOutput `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringOutput `pulumi:"includesPattern"`
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
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
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

// NewLocalOpkgRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalOpkgRepository(ctx *pulumi.Context,
	name string, args *LocalOpkgRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalOpkgRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource LocalOpkgRepository
	err := ctx.RegisterResource("artifactory:index/localOpkgRepository:LocalOpkgRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalOpkgRepository gets an existing LocalOpkgRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalOpkgRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalOpkgRepositoryState, opts ...pulumi.ResourceOption) (*LocalOpkgRepository, error) {
	var resource LocalOpkgRepository
	err := ctx.ReadResource("artifactory:index/localOpkgRepository:LocalOpkgRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalOpkgRepository resources.
type localOpkgRepositoryState struct {
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
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
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
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
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

type LocalOpkgRepositoryState struct {
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
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
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
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
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

func (LocalOpkgRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localOpkgRepositoryState)(nil)).Elem()
}

type localOpkgRepositoryArgs struct {
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
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
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
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
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

// The set of arguments for constructing a LocalOpkgRepository resource.
type LocalOpkgRepositoryArgs struct {
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
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
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
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
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

func (LocalOpkgRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localOpkgRepositoryArgs)(nil)).Elem()
}

type LocalOpkgRepositoryInput interface {
	pulumi.Input

	ToLocalOpkgRepositoryOutput() LocalOpkgRepositoryOutput
	ToLocalOpkgRepositoryOutputWithContext(ctx context.Context) LocalOpkgRepositoryOutput
}

func (*LocalOpkgRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalOpkgRepository)(nil)).Elem()
}

func (i *LocalOpkgRepository) ToLocalOpkgRepositoryOutput() LocalOpkgRepositoryOutput {
	return i.ToLocalOpkgRepositoryOutputWithContext(context.Background())
}

func (i *LocalOpkgRepository) ToLocalOpkgRepositoryOutputWithContext(ctx context.Context) LocalOpkgRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalOpkgRepositoryOutput)
}

// LocalOpkgRepositoryArrayInput is an input type that accepts LocalOpkgRepositoryArray and LocalOpkgRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalOpkgRepositoryArrayInput` via:
//
//	LocalOpkgRepositoryArray{ LocalOpkgRepositoryArgs{...} }
type LocalOpkgRepositoryArrayInput interface {
	pulumi.Input

	ToLocalOpkgRepositoryArrayOutput() LocalOpkgRepositoryArrayOutput
	ToLocalOpkgRepositoryArrayOutputWithContext(context.Context) LocalOpkgRepositoryArrayOutput
}

type LocalOpkgRepositoryArray []LocalOpkgRepositoryInput

func (LocalOpkgRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalOpkgRepository)(nil)).Elem()
}

func (i LocalOpkgRepositoryArray) ToLocalOpkgRepositoryArrayOutput() LocalOpkgRepositoryArrayOutput {
	return i.ToLocalOpkgRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalOpkgRepositoryArray) ToLocalOpkgRepositoryArrayOutputWithContext(ctx context.Context) LocalOpkgRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalOpkgRepositoryArrayOutput)
}

// LocalOpkgRepositoryMapInput is an input type that accepts LocalOpkgRepositoryMap and LocalOpkgRepositoryMapOutput values.
// You can construct a concrete instance of `LocalOpkgRepositoryMapInput` via:
//
//	LocalOpkgRepositoryMap{ "key": LocalOpkgRepositoryArgs{...} }
type LocalOpkgRepositoryMapInput interface {
	pulumi.Input

	ToLocalOpkgRepositoryMapOutput() LocalOpkgRepositoryMapOutput
	ToLocalOpkgRepositoryMapOutputWithContext(context.Context) LocalOpkgRepositoryMapOutput
}

type LocalOpkgRepositoryMap map[string]LocalOpkgRepositoryInput

func (LocalOpkgRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalOpkgRepository)(nil)).Elem()
}

func (i LocalOpkgRepositoryMap) ToLocalOpkgRepositoryMapOutput() LocalOpkgRepositoryMapOutput {
	return i.ToLocalOpkgRepositoryMapOutputWithContext(context.Background())
}

func (i LocalOpkgRepositoryMap) ToLocalOpkgRepositoryMapOutputWithContext(ctx context.Context) LocalOpkgRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalOpkgRepositoryMapOutput)
}

type LocalOpkgRepositoryOutput struct{ *pulumi.OutputState }

func (LocalOpkgRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalOpkgRepository)(nil)).Elem()
}

func (o LocalOpkgRepositoryOutput) ToLocalOpkgRepositoryOutput() LocalOpkgRepositoryOutput {
	return o
}

func (o LocalOpkgRepositoryOutput) ToLocalOpkgRepositoryOutputWithContext(ctx context.Context) LocalOpkgRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o LocalOpkgRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalOpkgRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o LocalOpkgRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalOpkgRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o LocalOpkgRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalOpkgRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

// Public description.
func (o LocalOpkgRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalOpkgRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o LocalOpkgRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalOpkgRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
// artifacts are excluded.
func (o LocalOpkgRepositoryOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalOpkgRepository) pulumi.StringOutput { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o LocalOpkgRepositoryOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalOpkgRepository) pulumi.StringOutput { return v.IncludesPattern }).(pulumi.StringOutput)
}

// the identity key of the repo.
func (o LocalOpkgRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalOpkgRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o LocalOpkgRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalOpkgRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LocalOpkgRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalOpkgRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o LocalOpkgRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalOpkgRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o LocalOpkgRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalOpkgRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o LocalOpkgRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalOpkgRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o LocalOpkgRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalOpkgRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o LocalOpkgRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalOpkgRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o LocalOpkgRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalOpkgRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type LocalOpkgRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalOpkgRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalOpkgRepository)(nil)).Elem()
}

func (o LocalOpkgRepositoryArrayOutput) ToLocalOpkgRepositoryArrayOutput() LocalOpkgRepositoryArrayOutput {
	return o
}

func (o LocalOpkgRepositoryArrayOutput) ToLocalOpkgRepositoryArrayOutputWithContext(ctx context.Context) LocalOpkgRepositoryArrayOutput {
	return o
}

func (o LocalOpkgRepositoryArrayOutput) Index(i pulumi.IntInput) LocalOpkgRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalOpkgRepository {
		return vs[0].([]*LocalOpkgRepository)[vs[1].(int)]
	}).(LocalOpkgRepositoryOutput)
}

type LocalOpkgRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalOpkgRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalOpkgRepository)(nil)).Elem()
}

func (o LocalOpkgRepositoryMapOutput) ToLocalOpkgRepositoryMapOutput() LocalOpkgRepositoryMapOutput {
	return o
}

func (o LocalOpkgRepositoryMapOutput) ToLocalOpkgRepositoryMapOutputWithContext(ctx context.Context) LocalOpkgRepositoryMapOutput {
	return o
}

func (o LocalOpkgRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalOpkgRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalOpkgRepository {
		return vs[0].(map[string]*LocalOpkgRepository)[vs[1].(string)]
	}).(LocalOpkgRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalOpkgRepositoryInput)(nil)).Elem(), &LocalOpkgRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalOpkgRepositoryArrayInput)(nil)).Elem(), LocalOpkgRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalOpkgRepositoryMapInput)(nil)).Elem(), LocalOpkgRepositoryMap{})
	pulumi.RegisterOutputType(LocalOpkgRepositoryOutput{})
	pulumi.RegisterOutputType(LocalOpkgRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalOpkgRepositoryMapOutput{})
}
