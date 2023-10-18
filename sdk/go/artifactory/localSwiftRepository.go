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

// Creates a local Swift repository.
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
//			_, err := artifactory.NewLocalSwiftRepository(ctx, "terraform-local-test-swift-repo", &artifactory.LocalSwiftRepositoryArgs{
//				Key: pulumi.String("terraform-local-test-swift-repo"),
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
//	$ pulumi import artifactory:index/localSwiftRepository:LocalSwiftRepository terraform-local-test-swift-repo terraform-local-test-swift-repo
//
// ```
type LocalSwiftRepository struct {
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

// NewLocalSwiftRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalSwiftRepository(ctx *pulumi.Context,
	name string, args *LocalSwiftRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalSwiftRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LocalSwiftRepository
	err := ctx.RegisterResource("artifactory:index/localSwiftRepository:LocalSwiftRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalSwiftRepository gets an existing LocalSwiftRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalSwiftRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalSwiftRepositoryState, opts ...pulumi.ResourceOption) (*LocalSwiftRepository, error) {
	var resource LocalSwiftRepository
	err := ctx.ReadResource("artifactory:index/localSwiftRepository:LocalSwiftRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalSwiftRepository resources.
type localSwiftRepositoryState struct {
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

type LocalSwiftRepositoryState struct {
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

func (LocalSwiftRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localSwiftRepositoryState)(nil)).Elem()
}

type localSwiftRepositoryArgs struct {
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

// The set of arguments for constructing a LocalSwiftRepository resource.
type LocalSwiftRepositoryArgs struct {
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

func (LocalSwiftRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localSwiftRepositoryArgs)(nil)).Elem()
}

type LocalSwiftRepositoryInput interface {
	pulumi.Input

	ToLocalSwiftRepositoryOutput() LocalSwiftRepositoryOutput
	ToLocalSwiftRepositoryOutputWithContext(ctx context.Context) LocalSwiftRepositoryOutput
}

func (*LocalSwiftRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalSwiftRepository)(nil)).Elem()
}

func (i *LocalSwiftRepository) ToLocalSwiftRepositoryOutput() LocalSwiftRepositoryOutput {
	return i.ToLocalSwiftRepositoryOutputWithContext(context.Background())
}

func (i *LocalSwiftRepository) ToLocalSwiftRepositoryOutputWithContext(ctx context.Context) LocalSwiftRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalSwiftRepositoryOutput)
}

func (i *LocalSwiftRepository) ToOutput(ctx context.Context) pulumix.Output[*LocalSwiftRepository] {
	return pulumix.Output[*LocalSwiftRepository]{
		OutputState: i.ToLocalSwiftRepositoryOutputWithContext(ctx).OutputState,
	}
}

// LocalSwiftRepositoryArrayInput is an input type that accepts LocalSwiftRepositoryArray and LocalSwiftRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalSwiftRepositoryArrayInput` via:
//
//	LocalSwiftRepositoryArray{ LocalSwiftRepositoryArgs{...} }
type LocalSwiftRepositoryArrayInput interface {
	pulumi.Input

	ToLocalSwiftRepositoryArrayOutput() LocalSwiftRepositoryArrayOutput
	ToLocalSwiftRepositoryArrayOutputWithContext(context.Context) LocalSwiftRepositoryArrayOutput
}

type LocalSwiftRepositoryArray []LocalSwiftRepositoryInput

func (LocalSwiftRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalSwiftRepository)(nil)).Elem()
}

func (i LocalSwiftRepositoryArray) ToLocalSwiftRepositoryArrayOutput() LocalSwiftRepositoryArrayOutput {
	return i.ToLocalSwiftRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalSwiftRepositoryArray) ToLocalSwiftRepositoryArrayOutputWithContext(ctx context.Context) LocalSwiftRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalSwiftRepositoryArrayOutput)
}

func (i LocalSwiftRepositoryArray) ToOutput(ctx context.Context) pulumix.Output[[]*LocalSwiftRepository] {
	return pulumix.Output[[]*LocalSwiftRepository]{
		OutputState: i.ToLocalSwiftRepositoryArrayOutputWithContext(ctx).OutputState,
	}
}

// LocalSwiftRepositoryMapInput is an input type that accepts LocalSwiftRepositoryMap and LocalSwiftRepositoryMapOutput values.
// You can construct a concrete instance of `LocalSwiftRepositoryMapInput` via:
//
//	LocalSwiftRepositoryMap{ "key": LocalSwiftRepositoryArgs{...} }
type LocalSwiftRepositoryMapInput interface {
	pulumi.Input

	ToLocalSwiftRepositoryMapOutput() LocalSwiftRepositoryMapOutput
	ToLocalSwiftRepositoryMapOutputWithContext(context.Context) LocalSwiftRepositoryMapOutput
}

type LocalSwiftRepositoryMap map[string]LocalSwiftRepositoryInput

func (LocalSwiftRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalSwiftRepository)(nil)).Elem()
}

func (i LocalSwiftRepositoryMap) ToLocalSwiftRepositoryMapOutput() LocalSwiftRepositoryMapOutput {
	return i.ToLocalSwiftRepositoryMapOutputWithContext(context.Background())
}

func (i LocalSwiftRepositoryMap) ToLocalSwiftRepositoryMapOutputWithContext(ctx context.Context) LocalSwiftRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalSwiftRepositoryMapOutput)
}

func (i LocalSwiftRepositoryMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*LocalSwiftRepository] {
	return pulumix.Output[map[string]*LocalSwiftRepository]{
		OutputState: i.ToLocalSwiftRepositoryMapOutputWithContext(ctx).OutputState,
	}
}

type LocalSwiftRepositoryOutput struct{ *pulumi.OutputState }

func (LocalSwiftRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalSwiftRepository)(nil)).Elem()
}

func (o LocalSwiftRepositoryOutput) ToLocalSwiftRepositoryOutput() LocalSwiftRepositoryOutput {
	return o
}

func (o LocalSwiftRepositoryOutput) ToLocalSwiftRepositoryOutputWithContext(ctx context.Context) LocalSwiftRepositoryOutput {
	return o
}

func (o LocalSwiftRepositoryOutput) ToOutput(ctx context.Context) pulumix.Output[*LocalSwiftRepository] {
	return pulumix.Output[*LocalSwiftRepository]{
		OutputState: o.OutputState,
	}
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o LocalSwiftRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalSwiftRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o LocalSwiftRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalSwiftRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o LocalSwiftRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalSwiftRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

// Public description.
func (o LocalSwiftRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalSwiftRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o LocalSwiftRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalSwiftRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o LocalSwiftRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalSwiftRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o LocalSwiftRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalSwiftRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// the identity key of the repo.
func (o LocalSwiftRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalSwiftRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o LocalSwiftRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalSwiftRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LocalSwiftRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalSwiftRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o LocalSwiftRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalSwiftRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o LocalSwiftRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalSwiftRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o LocalSwiftRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalSwiftRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o LocalSwiftRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalSwiftRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o LocalSwiftRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalSwiftRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o LocalSwiftRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalSwiftRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type LocalSwiftRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalSwiftRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalSwiftRepository)(nil)).Elem()
}

func (o LocalSwiftRepositoryArrayOutput) ToLocalSwiftRepositoryArrayOutput() LocalSwiftRepositoryArrayOutput {
	return o
}

func (o LocalSwiftRepositoryArrayOutput) ToLocalSwiftRepositoryArrayOutputWithContext(ctx context.Context) LocalSwiftRepositoryArrayOutput {
	return o
}

func (o LocalSwiftRepositoryArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*LocalSwiftRepository] {
	return pulumix.Output[[]*LocalSwiftRepository]{
		OutputState: o.OutputState,
	}
}

func (o LocalSwiftRepositoryArrayOutput) Index(i pulumi.IntInput) LocalSwiftRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalSwiftRepository {
		return vs[0].([]*LocalSwiftRepository)[vs[1].(int)]
	}).(LocalSwiftRepositoryOutput)
}

type LocalSwiftRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalSwiftRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalSwiftRepository)(nil)).Elem()
}

func (o LocalSwiftRepositoryMapOutput) ToLocalSwiftRepositoryMapOutput() LocalSwiftRepositoryMapOutput {
	return o
}

func (o LocalSwiftRepositoryMapOutput) ToLocalSwiftRepositoryMapOutputWithContext(ctx context.Context) LocalSwiftRepositoryMapOutput {
	return o
}

func (o LocalSwiftRepositoryMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*LocalSwiftRepository] {
	return pulumix.Output[map[string]*LocalSwiftRepository]{
		OutputState: o.OutputState,
	}
}

func (o LocalSwiftRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalSwiftRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalSwiftRepository {
		return vs[0].(map[string]*LocalSwiftRepository)[vs[1].(string)]
	}).(LocalSwiftRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalSwiftRepositoryInput)(nil)).Elem(), &LocalSwiftRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalSwiftRepositoryArrayInput)(nil)).Elem(), LocalSwiftRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalSwiftRepositoryMapInput)(nil)).Elem(), LocalSwiftRepositoryMap{})
	pulumi.RegisterOutputType(LocalSwiftRepositoryOutput{})
	pulumi.RegisterOutputType(LocalSwiftRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalSwiftRepositoryMapOutput{})
}
