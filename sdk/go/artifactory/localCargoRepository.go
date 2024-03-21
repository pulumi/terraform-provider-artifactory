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

// Creates a local Cargo repository.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := artifactory.NewLocalCargoRepository(ctx, "terraform-local-test-cargo-repo-basic", &artifactory.LocalCargoRepositoryArgs{
//				AnonymousAccess:   pulumi.Bool(false),
//				EnableSparseIndex: pulumi.Bool(true),
//				Key:               pulumi.String("terraform-local-test-cargo-repo-basic"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Local repositories can be imported using their name, e.g.
//
// ```sh
// $ pulumi import artifactory:index/localCargoRepository:LocalCargoRepository terraform-local-test-cargo-repo-basic terraform-local-test-cargo-repo-basic
// ```
type LocalCargoRepository struct {
	pulumi.CustomResourceState

	// Cargo client does not send credentials when performing download and search for crates.
	// Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option. Default value is `false`.
	AnonymousAccess pulumi.BoolPtrOutput `pulumi:"anonymousAccess"`
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
	// Enable internal index support based on Cargo sparse index specifications, instead of the default git index. Default value is `false`.
	EnableSparseIndex pulumi.BoolPtrOutput `pulumi:"enableSparseIndex"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrOutput `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern         pulumi.StringPtrOutput   `pulumi:"includesPattern"`
	IndexCompressionFormats pulumi.StringArrayOutput `pulumi:"indexCompressionFormats"`
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

// NewLocalCargoRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalCargoRepository(ctx *pulumi.Context,
	name string, args *LocalCargoRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalCargoRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LocalCargoRepository
	err := ctx.RegisterResource("artifactory:index/localCargoRepository:LocalCargoRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalCargoRepository gets an existing LocalCargoRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalCargoRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalCargoRepositoryState, opts ...pulumi.ResourceOption) (*LocalCargoRepository, error) {
	var resource LocalCargoRepository
	err := ctx.ReadResource("artifactory:index/localCargoRepository:LocalCargoRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalCargoRepository resources.
type localCargoRepositoryState struct {
	// Cargo client does not send credentials when performing download and search for crates.
	// Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option. Default value is `false`.
	AnonymousAccess *bool `pulumi:"anonymousAccess"`
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
	// Enable internal index support based on Cargo sparse index specifications, instead of the default git index. Default value is `false`.
	EnableSparseIndex *bool `pulumi:"enableSparseIndex"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern         *string  `pulumi:"includesPattern"`
	IndexCompressionFormats []string `pulumi:"indexCompressionFormats"`
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

type LocalCargoRepositoryState struct {
	// Cargo client does not send credentials when performing download and search for crates.
	// Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option. Default value is `false`.
	AnonymousAccess pulumi.BoolPtrInput
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
	// Enable internal index support based on Cargo sparse index specifications, instead of the default git index. Default value is `false`.
	EnableSparseIndex pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern         pulumi.StringPtrInput
	IndexCompressionFormats pulumi.StringArrayInput
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

func (LocalCargoRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localCargoRepositoryState)(nil)).Elem()
}

type localCargoRepositoryArgs struct {
	// Cargo client does not send credentials when performing download and search for crates.
	// Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option. Default value is `false`.
	AnonymousAccess *bool `pulumi:"anonymousAccess"`
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
	// Enable internal index support based on Cargo sparse index specifications, instead of the default git index. Default value is `false`.
	EnableSparseIndex *bool `pulumi:"enableSparseIndex"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern         *string  `pulumi:"includesPattern"`
	IndexCompressionFormats []string `pulumi:"indexCompressionFormats"`
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

// The set of arguments for constructing a LocalCargoRepository resource.
type LocalCargoRepositoryArgs struct {
	// Cargo client does not send credentials when performing download and search for crates.
	// Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option. Default value is `false`.
	AnonymousAccess pulumi.BoolPtrInput
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
	// Enable internal index support based on Cargo sparse index specifications, instead of the default git index. Default value is `false`.
	EnableSparseIndex pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern         pulumi.StringPtrInput
	IndexCompressionFormats pulumi.StringArrayInput
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

func (LocalCargoRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localCargoRepositoryArgs)(nil)).Elem()
}

type LocalCargoRepositoryInput interface {
	pulumi.Input

	ToLocalCargoRepositoryOutput() LocalCargoRepositoryOutput
	ToLocalCargoRepositoryOutputWithContext(ctx context.Context) LocalCargoRepositoryOutput
}

func (*LocalCargoRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalCargoRepository)(nil)).Elem()
}

func (i *LocalCargoRepository) ToLocalCargoRepositoryOutput() LocalCargoRepositoryOutput {
	return i.ToLocalCargoRepositoryOutputWithContext(context.Background())
}

func (i *LocalCargoRepository) ToLocalCargoRepositoryOutputWithContext(ctx context.Context) LocalCargoRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalCargoRepositoryOutput)
}

// LocalCargoRepositoryArrayInput is an input type that accepts LocalCargoRepositoryArray and LocalCargoRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalCargoRepositoryArrayInput` via:
//
//	LocalCargoRepositoryArray{ LocalCargoRepositoryArgs{...} }
type LocalCargoRepositoryArrayInput interface {
	pulumi.Input

	ToLocalCargoRepositoryArrayOutput() LocalCargoRepositoryArrayOutput
	ToLocalCargoRepositoryArrayOutputWithContext(context.Context) LocalCargoRepositoryArrayOutput
}

type LocalCargoRepositoryArray []LocalCargoRepositoryInput

func (LocalCargoRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalCargoRepository)(nil)).Elem()
}

func (i LocalCargoRepositoryArray) ToLocalCargoRepositoryArrayOutput() LocalCargoRepositoryArrayOutput {
	return i.ToLocalCargoRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalCargoRepositoryArray) ToLocalCargoRepositoryArrayOutputWithContext(ctx context.Context) LocalCargoRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalCargoRepositoryArrayOutput)
}

// LocalCargoRepositoryMapInput is an input type that accepts LocalCargoRepositoryMap and LocalCargoRepositoryMapOutput values.
// You can construct a concrete instance of `LocalCargoRepositoryMapInput` via:
//
//	LocalCargoRepositoryMap{ "key": LocalCargoRepositoryArgs{...} }
type LocalCargoRepositoryMapInput interface {
	pulumi.Input

	ToLocalCargoRepositoryMapOutput() LocalCargoRepositoryMapOutput
	ToLocalCargoRepositoryMapOutputWithContext(context.Context) LocalCargoRepositoryMapOutput
}

type LocalCargoRepositoryMap map[string]LocalCargoRepositoryInput

func (LocalCargoRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalCargoRepository)(nil)).Elem()
}

func (i LocalCargoRepositoryMap) ToLocalCargoRepositoryMapOutput() LocalCargoRepositoryMapOutput {
	return i.ToLocalCargoRepositoryMapOutputWithContext(context.Background())
}

func (i LocalCargoRepositoryMap) ToLocalCargoRepositoryMapOutputWithContext(ctx context.Context) LocalCargoRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalCargoRepositoryMapOutput)
}

type LocalCargoRepositoryOutput struct{ *pulumi.OutputState }

func (LocalCargoRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalCargoRepository)(nil)).Elem()
}

func (o LocalCargoRepositoryOutput) ToLocalCargoRepositoryOutput() LocalCargoRepositoryOutput {
	return o
}

func (o LocalCargoRepositoryOutput) ToLocalCargoRepositoryOutputWithContext(ctx context.Context) LocalCargoRepositoryOutput {
	return o
}

// Cargo client does not send credentials when performing download and search for crates.
// Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option. Default value is `false`.
func (o LocalCargoRepositoryOutput) AnonymousAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalCargoRepository) pulumi.BoolPtrOutput { return v.AnonymousAccess }).(pulumi.BoolPtrOutput)
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o LocalCargoRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalCargoRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o LocalCargoRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalCargoRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o LocalCargoRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalCargoRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

// Public description.
func (o LocalCargoRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalCargoRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o LocalCargoRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalCargoRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// Enable internal index support based on Cargo sparse index specifications, instead of the default git index. Default value is `false`.
func (o LocalCargoRepositoryOutput) EnableSparseIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalCargoRepository) pulumi.BoolPtrOutput { return v.EnableSparseIndex }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o LocalCargoRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalCargoRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o LocalCargoRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalCargoRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LocalCargoRepositoryOutput) IndexCompressionFormats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalCargoRepository) pulumi.StringArrayOutput { return v.IndexCompressionFormats }).(pulumi.StringArrayOutput)
}

// the identity key of the repo.
func (o LocalCargoRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalCargoRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o LocalCargoRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalCargoRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LocalCargoRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalCargoRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o LocalCargoRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalCargoRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o LocalCargoRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalCargoRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o LocalCargoRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalCargoRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o LocalCargoRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalCargoRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o LocalCargoRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalCargoRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o LocalCargoRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalCargoRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type LocalCargoRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalCargoRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalCargoRepository)(nil)).Elem()
}

func (o LocalCargoRepositoryArrayOutput) ToLocalCargoRepositoryArrayOutput() LocalCargoRepositoryArrayOutput {
	return o
}

func (o LocalCargoRepositoryArrayOutput) ToLocalCargoRepositoryArrayOutputWithContext(ctx context.Context) LocalCargoRepositoryArrayOutput {
	return o
}

func (o LocalCargoRepositoryArrayOutput) Index(i pulumi.IntInput) LocalCargoRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalCargoRepository {
		return vs[0].([]*LocalCargoRepository)[vs[1].(int)]
	}).(LocalCargoRepositoryOutput)
}

type LocalCargoRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalCargoRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalCargoRepository)(nil)).Elem()
}

func (o LocalCargoRepositoryMapOutput) ToLocalCargoRepositoryMapOutput() LocalCargoRepositoryMapOutput {
	return o
}

func (o LocalCargoRepositoryMapOutput) ToLocalCargoRepositoryMapOutputWithContext(ctx context.Context) LocalCargoRepositoryMapOutput {
	return o
}

func (o LocalCargoRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalCargoRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalCargoRepository {
		return vs[0].(map[string]*LocalCargoRepository)[vs[1].(string)]
	}).(LocalCargoRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalCargoRepositoryInput)(nil)).Elem(), &LocalCargoRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalCargoRepositoryArrayInput)(nil)).Elem(), LocalCargoRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalCargoRepositoryMapInput)(nil)).Elem(), LocalCargoRepositoryMap{})
	pulumi.RegisterOutputType(LocalCargoRepositoryOutput{})
	pulumi.RegisterOutputType(LocalCargoRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalCargoRepositoryMapOutput{})
}
