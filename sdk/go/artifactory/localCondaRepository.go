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

// Creates a local Conda repository.
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
//			_, err := artifactory.NewLocalCondaRepository(ctx, "terraform-local-test-conda-repo", &artifactory.LocalCondaRepositoryArgs{
//				Key: pulumi.String("terraform-local-test-conda-repo"),
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
// $ pulumi import artifactory:index/localCondaRepository:LocalCondaRepository terraform-local-test-conda-repo terraform-local-test-conda-repo
// ```
type LocalCondaRepository struct {
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

// NewLocalCondaRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalCondaRepository(ctx *pulumi.Context,
	name string, args *LocalCondaRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalCondaRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LocalCondaRepository
	err := ctx.RegisterResource("artifactory:index/localCondaRepository:LocalCondaRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalCondaRepository gets an existing LocalCondaRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalCondaRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalCondaRepositoryState, opts ...pulumi.ResourceOption) (*LocalCondaRepository, error) {
	var resource LocalCondaRepository
	err := ctx.ReadResource("artifactory:index/localCondaRepository:LocalCondaRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalCondaRepository resources.
type localCondaRepositoryState struct {
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

type LocalCondaRepositoryState struct {
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

func (LocalCondaRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localCondaRepositoryState)(nil)).Elem()
}

type localCondaRepositoryArgs struct {
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

// The set of arguments for constructing a LocalCondaRepository resource.
type LocalCondaRepositoryArgs struct {
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

func (LocalCondaRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localCondaRepositoryArgs)(nil)).Elem()
}

type LocalCondaRepositoryInput interface {
	pulumi.Input

	ToLocalCondaRepositoryOutput() LocalCondaRepositoryOutput
	ToLocalCondaRepositoryOutputWithContext(ctx context.Context) LocalCondaRepositoryOutput
}

func (*LocalCondaRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalCondaRepository)(nil)).Elem()
}

func (i *LocalCondaRepository) ToLocalCondaRepositoryOutput() LocalCondaRepositoryOutput {
	return i.ToLocalCondaRepositoryOutputWithContext(context.Background())
}

func (i *LocalCondaRepository) ToLocalCondaRepositoryOutputWithContext(ctx context.Context) LocalCondaRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalCondaRepositoryOutput)
}

// LocalCondaRepositoryArrayInput is an input type that accepts LocalCondaRepositoryArray and LocalCondaRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalCondaRepositoryArrayInput` via:
//
//	LocalCondaRepositoryArray{ LocalCondaRepositoryArgs{...} }
type LocalCondaRepositoryArrayInput interface {
	pulumi.Input

	ToLocalCondaRepositoryArrayOutput() LocalCondaRepositoryArrayOutput
	ToLocalCondaRepositoryArrayOutputWithContext(context.Context) LocalCondaRepositoryArrayOutput
}

type LocalCondaRepositoryArray []LocalCondaRepositoryInput

func (LocalCondaRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalCondaRepository)(nil)).Elem()
}

func (i LocalCondaRepositoryArray) ToLocalCondaRepositoryArrayOutput() LocalCondaRepositoryArrayOutput {
	return i.ToLocalCondaRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalCondaRepositoryArray) ToLocalCondaRepositoryArrayOutputWithContext(ctx context.Context) LocalCondaRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalCondaRepositoryArrayOutput)
}

// LocalCondaRepositoryMapInput is an input type that accepts LocalCondaRepositoryMap and LocalCondaRepositoryMapOutput values.
// You can construct a concrete instance of `LocalCondaRepositoryMapInput` via:
//
//	LocalCondaRepositoryMap{ "key": LocalCondaRepositoryArgs{...} }
type LocalCondaRepositoryMapInput interface {
	pulumi.Input

	ToLocalCondaRepositoryMapOutput() LocalCondaRepositoryMapOutput
	ToLocalCondaRepositoryMapOutputWithContext(context.Context) LocalCondaRepositoryMapOutput
}

type LocalCondaRepositoryMap map[string]LocalCondaRepositoryInput

func (LocalCondaRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalCondaRepository)(nil)).Elem()
}

func (i LocalCondaRepositoryMap) ToLocalCondaRepositoryMapOutput() LocalCondaRepositoryMapOutput {
	return i.ToLocalCondaRepositoryMapOutputWithContext(context.Background())
}

func (i LocalCondaRepositoryMap) ToLocalCondaRepositoryMapOutputWithContext(ctx context.Context) LocalCondaRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalCondaRepositoryMapOutput)
}

type LocalCondaRepositoryOutput struct{ *pulumi.OutputState }

func (LocalCondaRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalCondaRepository)(nil)).Elem()
}

func (o LocalCondaRepositoryOutput) ToLocalCondaRepositoryOutput() LocalCondaRepositoryOutput {
	return o
}

func (o LocalCondaRepositoryOutput) ToLocalCondaRepositoryOutputWithContext(ctx context.Context) LocalCondaRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o LocalCondaRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalCondaRepository) pulumi.BoolOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o LocalCondaRepositoryOutput) BlackedOut() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalCondaRepository) pulumi.BoolOutput { return v.BlackedOut }).(pulumi.BoolOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is `false`
func (o LocalCondaRepositoryOutput) CdnRedirect() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalCondaRepository) pulumi.BoolOutput { return v.CdnRedirect }).(pulumi.BoolOutput)
}

// Public description.
func (o LocalCondaRepositoryOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalCondaRepository) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o LocalCondaRepositoryOutput) DownloadDirect() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalCondaRepository) pulumi.BoolOutput { return v.DownloadDirect }).(pulumi.BoolOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
// artifacts are excluded.
func (o LocalCondaRepositoryOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalCondaRepository) pulumi.StringOutput { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
func (o LocalCondaRepositoryOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalCondaRepository) pulumi.StringOutput { return v.IncludesPattern }).(pulumi.StringOutput)
}

// the identity key of the repo.
func (o LocalCondaRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalCondaRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o LocalCondaRepositoryOutput) Notes() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalCondaRepository) pulumi.StringOutput { return v.Notes }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o LocalCondaRepositoryOutput) PriorityResolution() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalCondaRepository) pulumi.BoolOutput { return v.PriorityResolution }).(pulumi.BoolOutput)
}

func (o LocalCondaRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalCondaRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o LocalCondaRepositoryOutput) ProjectKey() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalCondaRepository) pulumi.StringOutput { return v.ProjectKey }).(pulumi.StringOutput)
}

// List of property set name
func (o LocalCondaRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalCondaRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
// corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
func (o LocalCondaRepositoryOutput) RepoLayoutRef() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalCondaRepository) pulumi.StringOutput { return v.RepoLayoutRef }).(pulumi.StringOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o LocalCondaRepositoryOutput) XrayIndex() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalCondaRepository) pulumi.BoolOutput { return v.XrayIndex }).(pulumi.BoolOutput)
}

type LocalCondaRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalCondaRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalCondaRepository)(nil)).Elem()
}

func (o LocalCondaRepositoryArrayOutput) ToLocalCondaRepositoryArrayOutput() LocalCondaRepositoryArrayOutput {
	return o
}

func (o LocalCondaRepositoryArrayOutput) ToLocalCondaRepositoryArrayOutputWithContext(ctx context.Context) LocalCondaRepositoryArrayOutput {
	return o
}

func (o LocalCondaRepositoryArrayOutput) Index(i pulumi.IntInput) LocalCondaRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalCondaRepository {
		return vs[0].([]*LocalCondaRepository)[vs[1].(int)]
	}).(LocalCondaRepositoryOutput)
}

type LocalCondaRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalCondaRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalCondaRepository)(nil)).Elem()
}

func (o LocalCondaRepositoryMapOutput) ToLocalCondaRepositoryMapOutput() LocalCondaRepositoryMapOutput {
	return o
}

func (o LocalCondaRepositoryMapOutput) ToLocalCondaRepositoryMapOutputWithContext(ctx context.Context) LocalCondaRepositoryMapOutput {
	return o
}

func (o LocalCondaRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalCondaRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalCondaRepository {
		return vs[0].(map[string]*LocalCondaRepository)[vs[1].(string)]
	}).(LocalCondaRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalCondaRepositoryInput)(nil)).Elem(), &LocalCondaRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalCondaRepositoryArrayInput)(nil)).Elem(), LocalCondaRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalCondaRepositoryMapInput)(nil)).Elem(), LocalCondaRepositoryMap{})
	pulumi.RegisterOutputType(LocalCondaRepositoryOutput{})
	pulumi.RegisterOutputType(LocalCondaRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalCondaRepositoryMapOutput{})
}
