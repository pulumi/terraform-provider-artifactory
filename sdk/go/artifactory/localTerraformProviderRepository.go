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
//			_, err := artifactory.NewLocalTerraformProviderRepository(ctx, "terraform-local-test-terraform-provider-repo", &artifactory.LocalTerraformProviderRepositoryArgs{
//				Key: pulumi.String("terraform-local-test-terraform-provider-repo"),
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
//	$ pulumi import artifactory:index/localTerraformProviderRepository:LocalTerraformProviderRepository terraform-local-test-terraform-provider-repo terraform-local-test-terraform-provider-repo
//
// ```
type LocalTerraformProviderRepository struct {
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

// NewLocalTerraformProviderRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalTerraformProviderRepository(ctx *pulumi.Context,
	name string, args *LocalTerraformProviderRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalTerraformProviderRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LocalTerraformProviderRepository
	err := ctx.RegisterResource("artifactory:index/localTerraformProviderRepository:LocalTerraformProviderRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalTerraformProviderRepository gets an existing LocalTerraformProviderRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalTerraformProviderRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalTerraformProviderRepositoryState, opts ...pulumi.ResourceOption) (*LocalTerraformProviderRepository, error) {
	var resource LocalTerraformProviderRepository
	err := ctx.ReadResource("artifactory:index/localTerraformProviderRepository:LocalTerraformProviderRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalTerraformProviderRepository resources.
type localTerraformProviderRepositoryState struct {
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

type LocalTerraformProviderRepositoryState struct {
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

func (LocalTerraformProviderRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localTerraformProviderRepositoryState)(nil)).Elem()
}

type localTerraformProviderRepositoryArgs struct {
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

// The set of arguments for constructing a LocalTerraformProviderRepository resource.
type LocalTerraformProviderRepositoryArgs struct {
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

func (LocalTerraformProviderRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localTerraformProviderRepositoryArgs)(nil)).Elem()
}

type LocalTerraformProviderRepositoryInput interface {
	pulumi.Input

	ToLocalTerraformProviderRepositoryOutput() LocalTerraformProviderRepositoryOutput
	ToLocalTerraformProviderRepositoryOutputWithContext(ctx context.Context) LocalTerraformProviderRepositoryOutput
}

func (*LocalTerraformProviderRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalTerraformProviderRepository)(nil)).Elem()
}

func (i *LocalTerraformProviderRepository) ToLocalTerraformProviderRepositoryOutput() LocalTerraformProviderRepositoryOutput {
	return i.ToLocalTerraformProviderRepositoryOutputWithContext(context.Background())
}

func (i *LocalTerraformProviderRepository) ToLocalTerraformProviderRepositoryOutputWithContext(ctx context.Context) LocalTerraformProviderRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalTerraformProviderRepositoryOutput)
}

// LocalTerraformProviderRepositoryArrayInput is an input type that accepts LocalTerraformProviderRepositoryArray and LocalTerraformProviderRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalTerraformProviderRepositoryArrayInput` via:
//
//	LocalTerraformProviderRepositoryArray{ LocalTerraformProviderRepositoryArgs{...} }
type LocalTerraformProviderRepositoryArrayInput interface {
	pulumi.Input

	ToLocalTerraformProviderRepositoryArrayOutput() LocalTerraformProviderRepositoryArrayOutput
	ToLocalTerraformProviderRepositoryArrayOutputWithContext(context.Context) LocalTerraformProviderRepositoryArrayOutput
}

type LocalTerraformProviderRepositoryArray []LocalTerraformProviderRepositoryInput

func (LocalTerraformProviderRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalTerraformProviderRepository)(nil)).Elem()
}

func (i LocalTerraformProviderRepositoryArray) ToLocalTerraformProviderRepositoryArrayOutput() LocalTerraformProviderRepositoryArrayOutput {
	return i.ToLocalTerraformProviderRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalTerraformProviderRepositoryArray) ToLocalTerraformProviderRepositoryArrayOutputWithContext(ctx context.Context) LocalTerraformProviderRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalTerraformProviderRepositoryArrayOutput)
}

// LocalTerraformProviderRepositoryMapInput is an input type that accepts LocalTerraformProviderRepositoryMap and LocalTerraformProviderRepositoryMapOutput values.
// You can construct a concrete instance of `LocalTerraformProviderRepositoryMapInput` via:
//
//	LocalTerraformProviderRepositoryMap{ "key": LocalTerraformProviderRepositoryArgs{...} }
type LocalTerraformProviderRepositoryMapInput interface {
	pulumi.Input

	ToLocalTerraformProviderRepositoryMapOutput() LocalTerraformProviderRepositoryMapOutput
	ToLocalTerraformProviderRepositoryMapOutputWithContext(context.Context) LocalTerraformProviderRepositoryMapOutput
}

type LocalTerraformProviderRepositoryMap map[string]LocalTerraformProviderRepositoryInput

func (LocalTerraformProviderRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalTerraformProviderRepository)(nil)).Elem()
}

func (i LocalTerraformProviderRepositoryMap) ToLocalTerraformProviderRepositoryMapOutput() LocalTerraformProviderRepositoryMapOutput {
	return i.ToLocalTerraformProviderRepositoryMapOutputWithContext(context.Background())
}

func (i LocalTerraformProviderRepositoryMap) ToLocalTerraformProviderRepositoryMapOutputWithContext(ctx context.Context) LocalTerraformProviderRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalTerraformProviderRepositoryMapOutput)
}

type LocalTerraformProviderRepositoryOutput struct{ *pulumi.OutputState }

func (LocalTerraformProviderRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalTerraformProviderRepository)(nil)).Elem()
}

func (o LocalTerraformProviderRepositoryOutput) ToLocalTerraformProviderRepositoryOutput() LocalTerraformProviderRepositoryOutput {
	return o
}

func (o LocalTerraformProviderRepositoryOutput) ToLocalTerraformProviderRepositoryOutputWithContext(ctx context.Context) LocalTerraformProviderRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o LocalTerraformProviderRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalTerraformProviderRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o LocalTerraformProviderRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalTerraformProviderRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o LocalTerraformProviderRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalTerraformProviderRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

// Public description.
func (o LocalTerraformProviderRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalTerraformProviderRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o LocalTerraformProviderRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalTerraformProviderRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o LocalTerraformProviderRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalTerraformProviderRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o LocalTerraformProviderRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalTerraformProviderRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// the identity key of the repo.
func (o LocalTerraformProviderRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalTerraformProviderRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o LocalTerraformProviderRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalTerraformProviderRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LocalTerraformProviderRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalTerraformProviderRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o LocalTerraformProviderRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalTerraformProviderRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o LocalTerraformProviderRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalTerraformProviderRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o LocalTerraformProviderRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalTerraformProviderRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o LocalTerraformProviderRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalTerraformProviderRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o LocalTerraformProviderRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalTerraformProviderRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o LocalTerraformProviderRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalTerraformProviderRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type LocalTerraformProviderRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalTerraformProviderRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalTerraformProviderRepository)(nil)).Elem()
}

func (o LocalTerraformProviderRepositoryArrayOutput) ToLocalTerraformProviderRepositoryArrayOutput() LocalTerraformProviderRepositoryArrayOutput {
	return o
}

func (o LocalTerraformProviderRepositoryArrayOutput) ToLocalTerraformProviderRepositoryArrayOutputWithContext(ctx context.Context) LocalTerraformProviderRepositoryArrayOutput {
	return o
}

func (o LocalTerraformProviderRepositoryArrayOutput) Index(i pulumi.IntInput) LocalTerraformProviderRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalTerraformProviderRepository {
		return vs[0].([]*LocalTerraformProviderRepository)[vs[1].(int)]
	}).(LocalTerraformProviderRepositoryOutput)
}

type LocalTerraformProviderRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalTerraformProviderRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalTerraformProviderRepository)(nil)).Elem()
}

func (o LocalTerraformProviderRepositoryMapOutput) ToLocalTerraformProviderRepositoryMapOutput() LocalTerraformProviderRepositoryMapOutput {
	return o
}

func (o LocalTerraformProviderRepositoryMapOutput) ToLocalTerraformProviderRepositoryMapOutputWithContext(ctx context.Context) LocalTerraformProviderRepositoryMapOutput {
	return o
}

func (o LocalTerraformProviderRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalTerraformProviderRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalTerraformProviderRepository {
		return vs[0].(map[string]*LocalTerraformProviderRepository)[vs[1].(string)]
	}).(LocalTerraformProviderRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalTerraformProviderRepositoryInput)(nil)).Elem(), &LocalTerraformProviderRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalTerraformProviderRepositoryArrayInput)(nil)).Elem(), LocalTerraformProviderRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalTerraformProviderRepositoryMapInput)(nil)).Elem(), LocalTerraformProviderRepositoryMap{})
	pulumi.RegisterOutputType(LocalTerraformProviderRepositoryOutput{})
	pulumi.RegisterOutputType(LocalTerraformProviderRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalTerraformProviderRepositoryMapOutput{})
}
