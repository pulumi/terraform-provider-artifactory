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

// Creates a local Vagrant repository.
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
//			_, err := artifactory.NewLocalVagrantRepository(ctx, "terraform-local-test-vagrant-repo", &artifactory.LocalVagrantRepositoryArgs{
//				Key: pulumi.String("terraform-local-test-vagrant-repo"),
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
//	$ pulumi import artifactory:index/localVagrantRepository:LocalVagrantRepository terraform-local-test-vagrant-repo terraform-local-test-vagrant-repo
//
// ```
type LocalVagrantRepository struct {
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

// NewLocalVagrantRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalVagrantRepository(ctx *pulumi.Context,
	name string, args *LocalVagrantRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalVagrantRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LocalVagrantRepository
	err := ctx.RegisterResource("artifactory:index/localVagrantRepository:LocalVagrantRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalVagrantRepository gets an existing LocalVagrantRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalVagrantRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalVagrantRepositoryState, opts ...pulumi.ResourceOption) (*LocalVagrantRepository, error) {
	var resource LocalVagrantRepository
	err := ctx.ReadResource("artifactory:index/localVagrantRepository:LocalVagrantRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalVagrantRepository resources.
type localVagrantRepositoryState struct {
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

type LocalVagrantRepositoryState struct {
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

func (LocalVagrantRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localVagrantRepositoryState)(nil)).Elem()
}

type localVagrantRepositoryArgs struct {
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

// The set of arguments for constructing a LocalVagrantRepository resource.
type LocalVagrantRepositoryArgs struct {
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

func (LocalVagrantRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localVagrantRepositoryArgs)(nil)).Elem()
}

type LocalVagrantRepositoryInput interface {
	pulumi.Input

	ToLocalVagrantRepositoryOutput() LocalVagrantRepositoryOutput
	ToLocalVagrantRepositoryOutputWithContext(ctx context.Context) LocalVagrantRepositoryOutput
}

func (*LocalVagrantRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalVagrantRepository)(nil)).Elem()
}

func (i *LocalVagrantRepository) ToLocalVagrantRepositoryOutput() LocalVagrantRepositoryOutput {
	return i.ToLocalVagrantRepositoryOutputWithContext(context.Background())
}

func (i *LocalVagrantRepository) ToLocalVagrantRepositoryOutputWithContext(ctx context.Context) LocalVagrantRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalVagrantRepositoryOutput)
}

func (i *LocalVagrantRepository) ToOutput(ctx context.Context) pulumix.Output[*LocalVagrantRepository] {
	return pulumix.Output[*LocalVagrantRepository]{
		OutputState: i.ToLocalVagrantRepositoryOutputWithContext(ctx).OutputState,
	}
}

// LocalVagrantRepositoryArrayInput is an input type that accepts LocalVagrantRepositoryArray and LocalVagrantRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalVagrantRepositoryArrayInput` via:
//
//	LocalVagrantRepositoryArray{ LocalVagrantRepositoryArgs{...} }
type LocalVagrantRepositoryArrayInput interface {
	pulumi.Input

	ToLocalVagrantRepositoryArrayOutput() LocalVagrantRepositoryArrayOutput
	ToLocalVagrantRepositoryArrayOutputWithContext(context.Context) LocalVagrantRepositoryArrayOutput
}

type LocalVagrantRepositoryArray []LocalVagrantRepositoryInput

func (LocalVagrantRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalVagrantRepository)(nil)).Elem()
}

func (i LocalVagrantRepositoryArray) ToLocalVagrantRepositoryArrayOutput() LocalVagrantRepositoryArrayOutput {
	return i.ToLocalVagrantRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalVagrantRepositoryArray) ToLocalVagrantRepositoryArrayOutputWithContext(ctx context.Context) LocalVagrantRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalVagrantRepositoryArrayOutput)
}

func (i LocalVagrantRepositoryArray) ToOutput(ctx context.Context) pulumix.Output[[]*LocalVagrantRepository] {
	return pulumix.Output[[]*LocalVagrantRepository]{
		OutputState: i.ToLocalVagrantRepositoryArrayOutputWithContext(ctx).OutputState,
	}
}

// LocalVagrantRepositoryMapInput is an input type that accepts LocalVagrantRepositoryMap and LocalVagrantRepositoryMapOutput values.
// You can construct a concrete instance of `LocalVagrantRepositoryMapInput` via:
//
//	LocalVagrantRepositoryMap{ "key": LocalVagrantRepositoryArgs{...} }
type LocalVagrantRepositoryMapInput interface {
	pulumi.Input

	ToLocalVagrantRepositoryMapOutput() LocalVagrantRepositoryMapOutput
	ToLocalVagrantRepositoryMapOutputWithContext(context.Context) LocalVagrantRepositoryMapOutput
}

type LocalVagrantRepositoryMap map[string]LocalVagrantRepositoryInput

func (LocalVagrantRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalVagrantRepository)(nil)).Elem()
}

func (i LocalVagrantRepositoryMap) ToLocalVagrantRepositoryMapOutput() LocalVagrantRepositoryMapOutput {
	return i.ToLocalVagrantRepositoryMapOutputWithContext(context.Background())
}

func (i LocalVagrantRepositoryMap) ToLocalVagrantRepositoryMapOutputWithContext(ctx context.Context) LocalVagrantRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalVagrantRepositoryMapOutput)
}

func (i LocalVagrantRepositoryMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*LocalVagrantRepository] {
	return pulumix.Output[map[string]*LocalVagrantRepository]{
		OutputState: i.ToLocalVagrantRepositoryMapOutputWithContext(ctx).OutputState,
	}
}

type LocalVagrantRepositoryOutput struct{ *pulumi.OutputState }

func (LocalVagrantRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalVagrantRepository)(nil)).Elem()
}

func (o LocalVagrantRepositoryOutput) ToLocalVagrantRepositoryOutput() LocalVagrantRepositoryOutput {
	return o
}

func (o LocalVagrantRepositoryOutput) ToLocalVagrantRepositoryOutputWithContext(ctx context.Context) LocalVagrantRepositoryOutput {
	return o
}

func (o LocalVagrantRepositoryOutput) ToOutput(ctx context.Context) pulumix.Output[*LocalVagrantRepository] {
	return pulumix.Output[*LocalVagrantRepository]{
		OutputState: o.OutputState,
	}
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o LocalVagrantRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalVagrantRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o LocalVagrantRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalVagrantRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o LocalVagrantRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalVagrantRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

// Public description.
func (o LocalVagrantRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalVagrantRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o LocalVagrantRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalVagrantRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o LocalVagrantRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalVagrantRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o LocalVagrantRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalVagrantRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// the identity key of the repo.
func (o LocalVagrantRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalVagrantRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o LocalVagrantRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalVagrantRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LocalVagrantRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalVagrantRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o LocalVagrantRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalVagrantRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o LocalVagrantRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalVagrantRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o LocalVagrantRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalVagrantRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o LocalVagrantRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalVagrantRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o LocalVagrantRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalVagrantRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o LocalVagrantRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalVagrantRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type LocalVagrantRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalVagrantRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalVagrantRepository)(nil)).Elem()
}

func (o LocalVagrantRepositoryArrayOutput) ToLocalVagrantRepositoryArrayOutput() LocalVagrantRepositoryArrayOutput {
	return o
}

func (o LocalVagrantRepositoryArrayOutput) ToLocalVagrantRepositoryArrayOutputWithContext(ctx context.Context) LocalVagrantRepositoryArrayOutput {
	return o
}

func (o LocalVagrantRepositoryArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*LocalVagrantRepository] {
	return pulumix.Output[[]*LocalVagrantRepository]{
		OutputState: o.OutputState,
	}
}

func (o LocalVagrantRepositoryArrayOutput) Index(i pulumi.IntInput) LocalVagrantRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalVagrantRepository {
		return vs[0].([]*LocalVagrantRepository)[vs[1].(int)]
	}).(LocalVagrantRepositoryOutput)
}

type LocalVagrantRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalVagrantRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalVagrantRepository)(nil)).Elem()
}

func (o LocalVagrantRepositoryMapOutput) ToLocalVagrantRepositoryMapOutput() LocalVagrantRepositoryMapOutput {
	return o
}

func (o LocalVagrantRepositoryMapOutput) ToLocalVagrantRepositoryMapOutputWithContext(ctx context.Context) LocalVagrantRepositoryMapOutput {
	return o
}

func (o LocalVagrantRepositoryMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*LocalVagrantRepository] {
	return pulumix.Output[map[string]*LocalVagrantRepository]{
		OutputState: o.OutputState,
	}
}

func (o LocalVagrantRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalVagrantRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalVagrantRepository {
		return vs[0].(map[string]*LocalVagrantRepository)[vs[1].(string)]
	}).(LocalVagrantRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalVagrantRepositoryInput)(nil)).Elem(), &LocalVagrantRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalVagrantRepositoryArrayInput)(nil)).Elem(), LocalVagrantRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalVagrantRepositoryMapInput)(nil)).Elem(), LocalVagrantRepositoryMap{})
	pulumi.RegisterOutputType(LocalVagrantRepositoryOutput{})
	pulumi.RegisterOutputType(LocalVagrantRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalVagrantRepositoryMapOutput{})
}
