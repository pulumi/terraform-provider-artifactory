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

// Creates a local Npm repository.
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
//			_, err := artifactory.NewLocalNpmRepository(ctx, "terraform-local-test-npm-repo", &artifactory.LocalNpmRepositoryArgs{
//				Key: pulumi.String("terraform-local-test-npm-repo"),
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
// $ pulumi import artifactory:index/localNpmRepository:LocalNpmRepository terraform-local-test-npm-repo terraform-local-test-npm-repo
// ```
type LocalNpmRepository struct {
	pulumi.CustomResourceState

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolOutput `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolOutput `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is `false`
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

// NewLocalNpmRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalNpmRepository(ctx *pulumi.Context,
	name string, args *LocalNpmRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalNpmRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LocalNpmRepository
	err := ctx.RegisterResource("artifactory:index/localNpmRepository:LocalNpmRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalNpmRepository gets an existing LocalNpmRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalNpmRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalNpmRepositoryState, opts ...pulumi.ResourceOption) (*LocalNpmRepository, error) {
	var resource LocalNpmRepository
	err := ctx.ReadResource("artifactory:index/localNpmRepository:LocalNpmRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalNpmRepository resources.
type localNpmRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is `false`
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

type LocalNpmRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is `false`
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

func (LocalNpmRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localNpmRepositoryState)(nil)).Elem()
}

type localNpmRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is `false`
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

// The set of arguments for constructing a LocalNpmRepository resource.
type LocalNpmRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is `false`
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

func (LocalNpmRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localNpmRepositoryArgs)(nil)).Elem()
}

type LocalNpmRepositoryInput interface {
	pulumi.Input

	ToLocalNpmRepositoryOutput() LocalNpmRepositoryOutput
	ToLocalNpmRepositoryOutputWithContext(ctx context.Context) LocalNpmRepositoryOutput
}

func (*LocalNpmRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalNpmRepository)(nil)).Elem()
}

func (i *LocalNpmRepository) ToLocalNpmRepositoryOutput() LocalNpmRepositoryOutput {
	return i.ToLocalNpmRepositoryOutputWithContext(context.Background())
}

func (i *LocalNpmRepository) ToLocalNpmRepositoryOutputWithContext(ctx context.Context) LocalNpmRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalNpmRepositoryOutput)
}

// LocalNpmRepositoryArrayInput is an input type that accepts LocalNpmRepositoryArray and LocalNpmRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalNpmRepositoryArrayInput` via:
//
//	LocalNpmRepositoryArray{ LocalNpmRepositoryArgs{...} }
type LocalNpmRepositoryArrayInput interface {
	pulumi.Input

	ToLocalNpmRepositoryArrayOutput() LocalNpmRepositoryArrayOutput
	ToLocalNpmRepositoryArrayOutputWithContext(context.Context) LocalNpmRepositoryArrayOutput
}

type LocalNpmRepositoryArray []LocalNpmRepositoryInput

func (LocalNpmRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalNpmRepository)(nil)).Elem()
}

func (i LocalNpmRepositoryArray) ToLocalNpmRepositoryArrayOutput() LocalNpmRepositoryArrayOutput {
	return i.ToLocalNpmRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalNpmRepositoryArray) ToLocalNpmRepositoryArrayOutputWithContext(ctx context.Context) LocalNpmRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalNpmRepositoryArrayOutput)
}

// LocalNpmRepositoryMapInput is an input type that accepts LocalNpmRepositoryMap and LocalNpmRepositoryMapOutput values.
// You can construct a concrete instance of `LocalNpmRepositoryMapInput` via:
//
//	LocalNpmRepositoryMap{ "key": LocalNpmRepositoryArgs{...} }
type LocalNpmRepositoryMapInput interface {
	pulumi.Input

	ToLocalNpmRepositoryMapOutput() LocalNpmRepositoryMapOutput
	ToLocalNpmRepositoryMapOutputWithContext(context.Context) LocalNpmRepositoryMapOutput
}

type LocalNpmRepositoryMap map[string]LocalNpmRepositoryInput

func (LocalNpmRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalNpmRepository)(nil)).Elem()
}

func (i LocalNpmRepositoryMap) ToLocalNpmRepositoryMapOutput() LocalNpmRepositoryMapOutput {
	return i.ToLocalNpmRepositoryMapOutputWithContext(context.Background())
}

func (i LocalNpmRepositoryMap) ToLocalNpmRepositoryMapOutputWithContext(ctx context.Context) LocalNpmRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalNpmRepositoryMapOutput)
}

type LocalNpmRepositoryOutput struct{ *pulumi.OutputState }

func (LocalNpmRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalNpmRepository)(nil)).Elem()
}

func (o LocalNpmRepositoryOutput) ToLocalNpmRepositoryOutput() LocalNpmRepositoryOutput {
	return o
}

func (o LocalNpmRepositoryOutput) ToLocalNpmRepositoryOutputWithContext(ctx context.Context) LocalNpmRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o LocalNpmRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalNpmRepository) pulumi.BoolOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o LocalNpmRepositoryOutput) BlackedOut() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalNpmRepository) pulumi.BoolOutput { return v.BlackedOut }).(pulumi.BoolOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is `false`
func (o LocalNpmRepositoryOutput) CdnRedirect() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalNpmRepository) pulumi.BoolOutput { return v.CdnRedirect }).(pulumi.BoolOutput)
}

// Public description.
func (o LocalNpmRepositoryOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalNpmRepository) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o LocalNpmRepositoryOutput) DownloadDirect() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalNpmRepository) pulumi.BoolOutput { return v.DownloadDirect }).(pulumi.BoolOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
// artifacts are excluded.
func (o LocalNpmRepositoryOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalNpmRepository) pulumi.StringOutput { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
func (o LocalNpmRepositoryOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalNpmRepository) pulumi.StringOutput { return v.IncludesPattern }).(pulumi.StringOutput)
}

// the identity key of the repo.
func (o LocalNpmRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalNpmRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o LocalNpmRepositoryOutput) Notes() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalNpmRepository) pulumi.StringOutput { return v.Notes }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o LocalNpmRepositoryOutput) PriorityResolution() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalNpmRepository) pulumi.BoolOutput { return v.PriorityResolution }).(pulumi.BoolOutput)
}

func (o LocalNpmRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalNpmRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o LocalNpmRepositoryOutput) ProjectKey() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalNpmRepository) pulumi.StringOutput { return v.ProjectKey }).(pulumi.StringOutput)
}

// List of property set name
func (o LocalNpmRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalNpmRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
// corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
func (o LocalNpmRepositoryOutput) RepoLayoutRef() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalNpmRepository) pulumi.StringOutput { return v.RepoLayoutRef }).(pulumi.StringOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o LocalNpmRepositoryOutput) XrayIndex() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalNpmRepository) pulumi.BoolOutput { return v.XrayIndex }).(pulumi.BoolOutput)
}

type LocalNpmRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalNpmRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalNpmRepository)(nil)).Elem()
}

func (o LocalNpmRepositoryArrayOutput) ToLocalNpmRepositoryArrayOutput() LocalNpmRepositoryArrayOutput {
	return o
}

func (o LocalNpmRepositoryArrayOutput) ToLocalNpmRepositoryArrayOutputWithContext(ctx context.Context) LocalNpmRepositoryArrayOutput {
	return o
}

func (o LocalNpmRepositoryArrayOutput) Index(i pulumi.IntInput) LocalNpmRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalNpmRepository {
		return vs[0].([]*LocalNpmRepository)[vs[1].(int)]
	}).(LocalNpmRepositoryOutput)
}

type LocalNpmRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalNpmRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalNpmRepository)(nil)).Elem()
}

func (o LocalNpmRepositoryMapOutput) ToLocalNpmRepositoryMapOutput() LocalNpmRepositoryMapOutput {
	return o
}

func (o LocalNpmRepositoryMapOutput) ToLocalNpmRepositoryMapOutputWithContext(ctx context.Context) LocalNpmRepositoryMapOutput {
	return o
}

func (o LocalNpmRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalNpmRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalNpmRepository {
		return vs[0].(map[string]*LocalNpmRepository)[vs[1].(string)]
	}).(LocalNpmRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalNpmRepositoryInput)(nil)).Elem(), &LocalNpmRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalNpmRepositoryArrayInput)(nil)).Elem(), LocalNpmRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalNpmRepositoryMapInput)(nil)).Elem(), LocalNpmRepositoryMap{})
	pulumi.RegisterOutputType(LocalNpmRepositoryOutput{})
	pulumi.RegisterOutputType(LocalNpmRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalNpmRepositoryMapOutput{})
}
