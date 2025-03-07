// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a local Cocoapods repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewLocalCocoapodsRepository(ctx, "terraform-local-test-cocoapods-repo", &artifactory.LocalCocoapodsRepositoryArgs{
//				Key: pulumi.String("terraform-local-test-cocoapods-repo"),
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
// $ pulumi import artifactory:index/localCocoapodsRepository:LocalCocoapodsRepository terraform-local-test-cocoapods-repo terraform-local-test-cocoapods-repo
// ```
type LocalCocoapodsRepository struct {
	pulumi.CustomResourceState

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolOutput `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolOutput `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolOutput `pulumi:"cdnRedirect"`
	// Public description.
	Description pulumi.StringOutput `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolOutput `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringOutput `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
	IncludesPattern pulumi.StringOutput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key pulumi.StringOutput `pulumi:"key"`
	// Internal description.
	Notes pulumi.StringOutput `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  pulumi.BoolOutput        `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringOutput `pulumi:"projectKey"`
	// List of property set name
	PropertySets pulumi.StringArrayOutput `pulumi:"propertySets"`
	// Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
	// corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
	RepoLayoutRef pulumi.StringOutput `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolOutput `pulumi:"xrayIndex"`
}

// NewLocalCocoapodsRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalCocoapodsRepository(ctx *pulumi.Context,
	name string, args *LocalCocoapodsRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalCocoapodsRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LocalCocoapodsRepository
	err := ctx.RegisterResource("artifactory:index/localCocoapodsRepository:LocalCocoapodsRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalCocoapodsRepository gets an existing LocalCocoapodsRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalCocoapodsRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalCocoapodsRepositoryState, opts ...pulumi.ResourceOption) (*LocalCocoapodsRepository, error) {
	var resource LocalCocoapodsRepository
	err := ctx.ReadResource("artifactory:index/localCocoapodsRepository:LocalCocoapodsRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalCocoapodsRepository resources.
type localCocoapodsRepositoryState struct {
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
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key *string `pulumi:"key"`
	// Internal description.
	Notes *string `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// List of property set name
	PropertySets []string `pulumi:"propertySets"`
	// Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
	// corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

type LocalCocoapodsRepositoryState struct {
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
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
	IncludesPattern pulumi.StringPtrInput
	// the identity key of the repo.
	Key pulumi.StringPtrInput
	// Internal description.
	Notes pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  pulumi.BoolPtrInput
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// List of property set name
	PropertySets pulumi.StringArrayInput
	// Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
	// corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
	RepoLayoutRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (LocalCocoapodsRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localCocoapodsRepositoryState)(nil)).Elem()
}

type localCocoapodsRepositoryArgs struct {
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
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// Internal description.
	Notes *string `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// List of property set name
	PropertySets []string `pulumi:"propertySets"`
	// Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
	// corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a LocalCocoapodsRepository resource.
type LocalCocoapodsRepositoryArgs struct {
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
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
	IncludesPattern pulumi.StringPtrInput
	// the identity key of the repo.
	Key pulumi.StringInput
	// Internal description.
	Notes pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  pulumi.BoolPtrInput
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// List of property set name
	PropertySets pulumi.StringArrayInput
	// Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
	// corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
	RepoLayoutRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (LocalCocoapodsRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localCocoapodsRepositoryArgs)(nil)).Elem()
}

type LocalCocoapodsRepositoryInput interface {
	pulumi.Input

	ToLocalCocoapodsRepositoryOutput() LocalCocoapodsRepositoryOutput
	ToLocalCocoapodsRepositoryOutputWithContext(ctx context.Context) LocalCocoapodsRepositoryOutput
}

func (*LocalCocoapodsRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalCocoapodsRepository)(nil)).Elem()
}

func (i *LocalCocoapodsRepository) ToLocalCocoapodsRepositoryOutput() LocalCocoapodsRepositoryOutput {
	return i.ToLocalCocoapodsRepositoryOutputWithContext(context.Background())
}

func (i *LocalCocoapodsRepository) ToLocalCocoapodsRepositoryOutputWithContext(ctx context.Context) LocalCocoapodsRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalCocoapodsRepositoryOutput)
}

// LocalCocoapodsRepositoryArrayInput is an input type that accepts LocalCocoapodsRepositoryArray and LocalCocoapodsRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalCocoapodsRepositoryArrayInput` via:
//
//	LocalCocoapodsRepositoryArray{ LocalCocoapodsRepositoryArgs{...} }
type LocalCocoapodsRepositoryArrayInput interface {
	pulumi.Input

	ToLocalCocoapodsRepositoryArrayOutput() LocalCocoapodsRepositoryArrayOutput
	ToLocalCocoapodsRepositoryArrayOutputWithContext(context.Context) LocalCocoapodsRepositoryArrayOutput
}

type LocalCocoapodsRepositoryArray []LocalCocoapodsRepositoryInput

func (LocalCocoapodsRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalCocoapodsRepository)(nil)).Elem()
}

func (i LocalCocoapodsRepositoryArray) ToLocalCocoapodsRepositoryArrayOutput() LocalCocoapodsRepositoryArrayOutput {
	return i.ToLocalCocoapodsRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalCocoapodsRepositoryArray) ToLocalCocoapodsRepositoryArrayOutputWithContext(ctx context.Context) LocalCocoapodsRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalCocoapodsRepositoryArrayOutput)
}

// LocalCocoapodsRepositoryMapInput is an input type that accepts LocalCocoapodsRepositoryMap and LocalCocoapodsRepositoryMapOutput values.
// You can construct a concrete instance of `LocalCocoapodsRepositoryMapInput` via:
//
//	LocalCocoapodsRepositoryMap{ "key": LocalCocoapodsRepositoryArgs{...} }
type LocalCocoapodsRepositoryMapInput interface {
	pulumi.Input

	ToLocalCocoapodsRepositoryMapOutput() LocalCocoapodsRepositoryMapOutput
	ToLocalCocoapodsRepositoryMapOutputWithContext(context.Context) LocalCocoapodsRepositoryMapOutput
}

type LocalCocoapodsRepositoryMap map[string]LocalCocoapodsRepositoryInput

func (LocalCocoapodsRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalCocoapodsRepository)(nil)).Elem()
}

func (i LocalCocoapodsRepositoryMap) ToLocalCocoapodsRepositoryMapOutput() LocalCocoapodsRepositoryMapOutput {
	return i.ToLocalCocoapodsRepositoryMapOutputWithContext(context.Background())
}

func (i LocalCocoapodsRepositoryMap) ToLocalCocoapodsRepositoryMapOutputWithContext(ctx context.Context) LocalCocoapodsRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalCocoapodsRepositoryMapOutput)
}

type LocalCocoapodsRepositoryOutput struct{ *pulumi.OutputState }

func (LocalCocoapodsRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalCocoapodsRepository)(nil)).Elem()
}

func (o LocalCocoapodsRepositoryOutput) ToLocalCocoapodsRepositoryOutput() LocalCocoapodsRepositoryOutput {
	return o
}

func (o LocalCocoapodsRepositoryOutput) ToLocalCocoapodsRepositoryOutputWithContext(ctx context.Context) LocalCocoapodsRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o LocalCocoapodsRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalCocoapodsRepository) pulumi.BoolOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o LocalCocoapodsRepositoryOutput) BlackedOut() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalCocoapodsRepository) pulumi.BoolOutput { return v.BlackedOut }).(pulumi.BoolOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o LocalCocoapodsRepositoryOutput) CdnRedirect() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalCocoapodsRepository) pulumi.BoolOutput { return v.CdnRedirect }).(pulumi.BoolOutput)
}

// Public description.
func (o LocalCocoapodsRepositoryOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalCocoapodsRepository) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o LocalCocoapodsRepositoryOutput) DownloadDirect() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalCocoapodsRepository) pulumi.BoolOutput { return v.DownloadDirect }).(pulumi.BoolOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
// artifacts are excluded.
func (o LocalCocoapodsRepositoryOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalCocoapodsRepository) pulumi.StringOutput { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
func (o LocalCocoapodsRepositoryOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalCocoapodsRepository) pulumi.StringOutput { return v.IncludesPattern }).(pulumi.StringOutput)
}

// the identity key of the repo.
func (o LocalCocoapodsRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalCocoapodsRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o LocalCocoapodsRepositoryOutput) Notes() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalCocoapodsRepository) pulumi.StringOutput { return v.Notes }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o LocalCocoapodsRepositoryOutput) PriorityResolution() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalCocoapodsRepository) pulumi.BoolOutput { return v.PriorityResolution }).(pulumi.BoolOutput)
}

func (o LocalCocoapodsRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalCocoapodsRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o LocalCocoapodsRepositoryOutput) ProjectKey() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalCocoapodsRepository) pulumi.StringOutput { return v.ProjectKey }).(pulumi.StringOutput)
}

// List of property set name
func (o LocalCocoapodsRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalCocoapodsRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
// corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
func (o LocalCocoapodsRepositoryOutput) RepoLayoutRef() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalCocoapodsRepository) pulumi.StringOutput { return v.RepoLayoutRef }).(pulumi.StringOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o LocalCocoapodsRepositoryOutput) XrayIndex() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalCocoapodsRepository) pulumi.BoolOutput { return v.XrayIndex }).(pulumi.BoolOutput)
}

type LocalCocoapodsRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalCocoapodsRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalCocoapodsRepository)(nil)).Elem()
}

func (o LocalCocoapodsRepositoryArrayOutput) ToLocalCocoapodsRepositoryArrayOutput() LocalCocoapodsRepositoryArrayOutput {
	return o
}

func (o LocalCocoapodsRepositoryArrayOutput) ToLocalCocoapodsRepositoryArrayOutputWithContext(ctx context.Context) LocalCocoapodsRepositoryArrayOutput {
	return o
}

func (o LocalCocoapodsRepositoryArrayOutput) Index(i pulumi.IntInput) LocalCocoapodsRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalCocoapodsRepository {
		return vs[0].([]*LocalCocoapodsRepository)[vs[1].(int)]
	}).(LocalCocoapodsRepositoryOutput)
}

type LocalCocoapodsRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalCocoapodsRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalCocoapodsRepository)(nil)).Elem()
}

func (o LocalCocoapodsRepositoryMapOutput) ToLocalCocoapodsRepositoryMapOutput() LocalCocoapodsRepositoryMapOutput {
	return o
}

func (o LocalCocoapodsRepositoryMapOutput) ToLocalCocoapodsRepositoryMapOutputWithContext(ctx context.Context) LocalCocoapodsRepositoryMapOutput {
	return o
}

func (o LocalCocoapodsRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalCocoapodsRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalCocoapodsRepository {
		return vs[0].(map[string]*LocalCocoapodsRepository)[vs[1].(string)]
	}).(LocalCocoapodsRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalCocoapodsRepositoryInput)(nil)).Elem(), &LocalCocoapodsRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalCocoapodsRepositoryArrayInput)(nil)).Elem(), LocalCocoapodsRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalCocoapodsRepositoryMapInput)(nil)).Elem(), LocalCocoapodsRepositoryMap{})
	pulumi.RegisterOutputType(LocalCocoapodsRepositoryOutput{})
	pulumi.RegisterOutputType(LocalCocoapodsRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalCocoapodsRepositoryMapOutput{})
}
