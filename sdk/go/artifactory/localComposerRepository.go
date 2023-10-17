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

// Creates a local Composer repository.
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
//			_, err := artifactory.NewLocalComposerRepository(ctx, "terraform-local-test-composer-repo", &artifactory.LocalComposerRepositoryArgs{
//				Key: pulumi.String("terraform-local-test-composer-repo"),
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
//	$ pulumi import artifactory:index/localComposerRepository:LocalComposerRepository terraform-local-test-composer-repo terraform-local-test-composer-repo
//
// ```
type LocalComposerRepository struct {
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

// NewLocalComposerRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalComposerRepository(ctx *pulumi.Context,
	name string, args *LocalComposerRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalComposerRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LocalComposerRepository
	err := ctx.RegisterResource("artifactory:index/localComposerRepository:LocalComposerRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalComposerRepository gets an existing LocalComposerRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalComposerRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalComposerRepositoryState, opts ...pulumi.ResourceOption) (*LocalComposerRepository, error) {
	var resource LocalComposerRepository
	err := ctx.ReadResource("artifactory:index/localComposerRepository:LocalComposerRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalComposerRepository resources.
type localComposerRepositoryState struct {
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

type LocalComposerRepositoryState struct {
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

func (LocalComposerRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localComposerRepositoryState)(nil)).Elem()
}

type localComposerRepositoryArgs struct {
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

// The set of arguments for constructing a LocalComposerRepository resource.
type LocalComposerRepositoryArgs struct {
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

func (LocalComposerRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localComposerRepositoryArgs)(nil)).Elem()
}

type LocalComposerRepositoryInput interface {
	pulumi.Input

	ToLocalComposerRepositoryOutput() LocalComposerRepositoryOutput
	ToLocalComposerRepositoryOutputWithContext(ctx context.Context) LocalComposerRepositoryOutput
}

func (*LocalComposerRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalComposerRepository)(nil)).Elem()
}

func (i *LocalComposerRepository) ToLocalComposerRepositoryOutput() LocalComposerRepositoryOutput {
	return i.ToLocalComposerRepositoryOutputWithContext(context.Background())
}

func (i *LocalComposerRepository) ToLocalComposerRepositoryOutputWithContext(ctx context.Context) LocalComposerRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalComposerRepositoryOutput)
}

func (i *LocalComposerRepository) ToOutput(ctx context.Context) pulumix.Output[*LocalComposerRepository] {
	return pulumix.Output[*LocalComposerRepository]{
		OutputState: i.ToLocalComposerRepositoryOutputWithContext(ctx).OutputState,
	}
}

// LocalComposerRepositoryArrayInput is an input type that accepts LocalComposerRepositoryArray and LocalComposerRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalComposerRepositoryArrayInput` via:
//
//	LocalComposerRepositoryArray{ LocalComposerRepositoryArgs{...} }
type LocalComposerRepositoryArrayInput interface {
	pulumi.Input

	ToLocalComposerRepositoryArrayOutput() LocalComposerRepositoryArrayOutput
	ToLocalComposerRepositoryArrayOutputWithContext(context.Context) LocalComposerRepositoryArrayOutput
}

type LocalComposerRepositoryArray []LocalComposerRepositoryInput

func (LocalComposerRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalComposerRepository)(nil)).Elem()
}

func (i LocalComposerRepositoryArray) ToLocalComposerRepositoryArrayOutput() LocalComposerRepositoryArrayOutput {
	return i.ToLocalComposerRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalComposerRepositoryArray) ToLocalComposerRepositoryArrayOutputWithContext(ctx context.Context) LocalComposerRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalComposerRepositoryArrayOutput)
}

func (i LocalComposerRepositoryArray) ToOutput(ctx context.Context) pulumix.Output[[]*LocalComposerRepository] {
	return pulumix.Output[[]*LocalComposerRepository]{
		OutputState: i.ToLocalComposerRepositoryArrayOutputWithContext(ctx).OutputState,
	}
}

// LocalComposerRepositoryMapInput is an input type that accepts LocalComposerRepositoryMap and LocalComposerRepositoryMapOutput values.
// You can construct a concrete instance of `LocalComposerRepositoryMapInput` via:
//
//	LocalComposerRepositoryMap{ "key": LocalComposerRepositoryArgs{...} }
type LocalComposerRepositoryMapInput interface {
	pulumi.Input

	ToLocalComposerRepositoryMapOutput() LocalComposerRepositoryMapOutput
	ToLocalComposerRepositoryMapOutputWithContext(context.Context) LocalComposerRepositoryMapOutput
}

type LocalComposerRepositoryMap map[string]LocalComposerRepositoryInput

func (LocalComposerRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalComposerRepository)(nil)).Elem()
}

func (i LocalComposerRepositoryMap) ToLocalComposerRepositoryMapOutput() LocalComposerRepositoryMapOutput {
	return i.ToLocalComposerRepositoryMapOutputWithContext(context.Background())
}

func (i LocalComposerRepositoryMap) ToLocalComposerRepositoryMapOutputWithContext(ctx context.Context) LocalComposerRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalComposerRepositoryMapOutput)
}

func (i LocalComposerRepositoryMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*LocalComposerRepository] {
	return pulumix.Output[map[string]*LocalComposerRepository]{
		OutputState: i.ToLocalComposerRepositoryMapOutputWithContext(ctx).OutputState,
	}
}

type LocalComposerRepositoryOutput struct{ *pulumi.OutputState }

func (LocalComposerRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalComposerRepository)(nil)).Elem()
}

func (o LocalComposerRepositoryOutput) ToLocalComposerRepositoryOutput() LocalComposerRepositoryOutput {
	return o
}

func (o LocalComposerRepositoryOutput) ToLocalComposerRepositoryOutputWithContext(ctx context.Context) LocalComposerRepositoryOutput {
	return o
}

func (o LocalComposerRepositoryOutput) ToOutput(ctx context.Context) pulumix.Output[*LocalComposerRepository] {
	return pulumix.Output[*LocalComposerRepository]{
		OutputState: o.OutputState,
	}
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o LocalComposerRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalComposerRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o LocalComposerRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalComposerRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o LocalComposerRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalComposerRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

// Public description.
func (o LocalComposerRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalComposerRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o LocalComposerRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalComposerRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o LocalComposerRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalComposerRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o LocalComposerRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalComposerRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// the identity key of the repo.
func (o LocalComposerRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalComposerRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o LocalComposerRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalComposerRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LocalComposerRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalComposerRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o LocalComposerRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalComposerRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o LocalComposerRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalComposerRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o LocalComposerRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalComposerRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o LocalComposerRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalComposerRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o LocalComposerRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalComposerRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o LocalComposerRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalComposerRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type LocalComposerRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalComposerRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalComposerRepository)(nil)).Elem()
}

func (o LocalComposerRepositoryArrayOutput) ToLocalComposerRepositoryArrayOutput() LocalComposerRepositoryArrayOutput {
	return o
}

func (o LocalComposerRepositoryArrayOutput) ToLocalComposerRepositoryArrayOutputWithContext(ctx context.Context) LocalComposerRepositoryArrayOutput {
	return o
}

func (o LocalComposerRepositoryArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*LocalComposerRepository] {
	return pulumix.Output[[]*LocalComposerRepository]{
		OutputState: o.OutputState,
	}
}

func (o LocalComposerRepositoryArrayOutput) Index(i pulumi.IntInput) LocalComposerRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalComposerRepository {
		return vs[0].([]*LocalComposerRepository)[vs[1].(int)]
	}).(LocalComposerRepositoryOutput)
}

type LocalComposerRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalComposerRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalComposerRepository)(nil)).Elem()
}

func (o LocalComposerRepositoryMapOutput) ToLocalComposerRepositoryMapOutput() LocalComposerRepositoryMapOutput {
	return o
}

func (o LocalComposerRepositoryMapOutput) ToLocalComposerRepositoryMapOutputWithContext(ctx context.Context) LocalComposerRepositoryMapOutput {
	return o
}

func (o LocalComposerRepositoryMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*LocalComposerRepository] {
	return pulumix.Output[map[string]*LocalComposerRepository]{
		OutputState: o.OutputState,
	}
}

func (o LocalComposerRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalComposerRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalComposerRepository {
		return vs[0].(map[string]*LocalComposerRepository)[vs[1].(string)]
	}).(LocalComposerRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalComposerRepositoryInput)(nil)).Elem(), &LocalComposerRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalComposerRepositoryArrayInput)(nil)).Elem(), LocalComposerRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalComposerRepositoryMapInput)(nil)).Elem(), LocalComposerRepositoryMap{})
	pulumi.RegisterOutputType(LocalComposerRepositoryOutput{})
	pulumi.RegisterOutputType(LocalComposerRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalComposerRepositoryMapOutput{})
}
