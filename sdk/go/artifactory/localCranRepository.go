// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Creates a local Cran repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewLocalCranRepository(ctx, "terraform-local-test-cran-repo", &artifactory.LocalCranRepositoryArgs{
//				Key: pulumi.String("terraform-local-test-cran-repo"),
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
//	$ pulumi import artifactory:index/localCranRepository:LocalCranRepository terraform-local-test-cran-repo terraform-local-test-cran-repo
//
// ```
type LocalCranRepository struct {
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

// NewLocalCranRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalCranRepository(ctx *pulumi.Context,
	name string, args *LocalCranRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalCranRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
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

type LocalCranRepositoryState struct {
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

func (LocalCranRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localCranRepositoryState)(nil)).Elem()
}

type localCranRepositoryArgs struct {
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

// The set of arguments for constructing a LocalCranRepository resource.
type LocalCranRepositoryArgs struct {
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

func (LocalCranRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localCranRepositoryArgs)(nil)).Elem()
}

type LocalCranRepositoryInput interface {
	pulumi.Input

	ToLocalCranRepositoryOutput() LocalCranRepositoryOutput
	ToLocalCranRepositoryOutputWithContext(ctx context.Context) LocalCranRepositoryOutput
}

func (*LocalCranRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalCranRepository)(nil)).Elem()
}

func (i *LocalCranRepository) ToLocalCranRepositoryOutput() LocalCranRepositoryOutput {
	return i.ToLocalCranRepositoryOutputWithContext(context.Background())
}

func (i *LocalCranRepository) ToLocalCranRepositoryOutputWithContext(ctx context.Context) LocalCranRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalCranRepositoryOutput)
}

func (i *LocalCranRepository) ToOutput(ctx context.Context) pulumix.Output[*LocalCranRepository] {
	return pulumix.Output[*LocalCranRepository]{
		OutputState: i.ToLocalCranRepositoryOutputWithContext(ctx).OutputState,
	}
}

// LocalCranRepositoryArrayInput is an input type that accepts LocalCranRepositoryArray and LocalCranRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalCranRepositoryArrayInput` via:
//
//	LocalCranRepositoryArray{ LocalCranRepositoryArgs{...} }
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

func (i LocalCranRepositoryArray) ToOutput(ctx context.Context) pulumix.Output[[]*LocalCranRepository] {
	return pulumix.Output[[]*LocalCranRepository]{
		OutputState: i.ToLocalCranRepositoryArrayOutputWithContext(ctx).OutputState,
	}
}

// LocalCranRepositoryMapInput is an input type that accepts LocalCranRepositoryMap and LocalCranRepositoryMapOutput values.
// You can construct a concrete instance of `LocalCranRepositoryMapInput` via:
//
//	LocalCranRepositoryMap{ "key": LocalCranRepositoryArgs{...} }
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

func (i LocalCranRepositoryMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*LocalCranRepository] {
	return pulumix.Output[map[string]*LocalCranRepository]{
		OutputState: i.ToLocalCranRepositoryMapOutputWithContext(ctx).OutputState,
	}
}

type LocalCranRepositoryOutput struct{ *pulumi.OutputState }

func (LocalCranRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalCranRepository)(nil)).Elem()
}

func (o LocalCranRepositoryOutput) ToLocalCranRepositoryOutput() LocalCranRepositoryOutput {
	return o
}

func (o LocalCranRepositoryOutput) ToLocalCranRepositoryOutputWithContext(ctx context.Context) LocalCranRepositoryOutput {
	return o
}

func (o LocalCranRepositoryOutput) ToOutput(ctx context.Context) pulumix.Output[*LocalCranRepository] {
	return pulumix.Output[*LocalCranRepository]{
		OutputState: o.OutputState,
	}
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o LocalCranRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalCranRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o LocalCranRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalCranRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o LocalCranRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalCranRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

// Public description.
func (o LocalCranRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalCranRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o LocalCranRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalCranRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
// artifacts are excluded.
func (o LocalCranRepositoryOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalCranRepository) pulumi.StringOutput { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o LocalCranRepositoryOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalCranRepository) pulumi.StringOutput { return v.IncludesPattern }).(pulumi.StringOutput)
}

// the identity key of the repo.
func (o LocalCranRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalCranRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o LocalCranRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalCranRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LocalCranRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalCranRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o LocalCranRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalCranRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o LocalCranRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalCranRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o LocalCranRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalCranRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o LocalCranRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalCranRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o LocalCranRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalCranRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o LocalCranRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalCranRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type LocalCranRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalCranRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalCranRepository)(nil)).Elem()
}

func (o LocalCranRepositoryArrayOutput) ToLocalCranRepositoryArrayOutput() LocalCranRepositoryArrayOutput {
	return o
}

func (o LocalCranRepositoryArrayOutput) ToLocalCranRepositoryArrayOutputWithContext(ctx context.Context) LocalCranRepositoryArrayOutput {
	return o
}

func (o LocalCranRepositoryArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*LocalCranRepository] {
	return pulumix.Output[[]*LocalCranRepository]{
		OutputState: o.OutputState,
	}
}

func (o LocalCranRepositoryArrayOutput) Index(i pulumi.IntInput) LocalCranRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalCranRepository {
		return vs[0].([]*LocalCranRepository)[vs[1].(int)]
	}).(LocalCranRepositoryOutput)
}

type LocalCranRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalCranRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalCranRepository)(nil)).Elem()
}

func (o LocalCranRepositoryMapOutput) ToLocalCranRepositoryMapOutput() LocalCranRepositoryMapOutput {
	return o
}

func (o LocalCranRepositoryMapOutput) ToLocalCranRepositoryMapOutputWithContext(ctx context.Context) LocalCranRepositoryMapOutput {
	return o
}

func (o LocalCranRepositoryMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*LocalCranRepository] {
	return pulumix.Output[map[string]*LocalCranRepository]{
		OutputState: o.OutputState,
	}
}

func (o LocalCranRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalCranRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalCranRepository {
		return vs[0].(map[string]*LocalCranRepository)[vs[1].(string)]
	}).(LocalCranRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalCranRepositoryInput)(nil)).Elem(), &LocalCranRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalCranRepositoryArrayInput)(nil)).Elem(), LocalCranRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalCranRepositoryMapInput)(nil)).Elem(), LocalCranRepositoryMap{})
	pulumi.RegisterOutputType(LocalCranRepositoryOutput{})
	pulumi.RegisterOutputType(LocalCranRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalCranRepositoryMapOutput{})
}
