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
//			_, err := artifactory.NewLocalTerraformModuleRepository(ctx, "terraform-local-test-terraform-module-repo", &artifactory.LocalTerraformModuleRepositoryArgs{
//				Key: pulumi.String("terraform-local-test-terraform-module-repo"),
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
// $ pulumi import artifactory:index/localTerraformModuleRepository:LocalTerraformModuleRepository terraform-local-test-terraform-module-repo terraform-local-test-terraform-module-repo
// ```
type LocalTerraformModuleRepository struct {
	pulumi.CustomResourceState

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolOutput `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolOutput `pulumi:"blackedOut"`
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

// NewLocalTerraformModuleRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalTerraformModuleRepository(ctx *pulumi.Context,
	name string, args *LocalTerraformModuleRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalTerraformModuleRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LocalTerraformModuleRepository
	err := ctx.RegisterResource("artifactory:index/localTerraformModuleRepository:LocalTerraformModuleRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalTerraformModuleRepository gets an existing LocalTerraformModuleRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalTerraformModuleRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalTerraformModuleRepositoryState, opts ...pulumi.ResourceOption) (*LocalTerraformModuleRepository, error) {
	var resource LocalTerraformModuleRepository
	err := ctx.ReadResource("artifactory:index/localTerraformModuleRepository:LocalTerraformModuleRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalTerraformModuleRepository resources.
type localTerraformModuleRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
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

type LocalTerraformModuleRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
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

func (LocalTerraformModuleRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localTerraformModuleRepositoryState)(nil)).Elem()
}

type localTerraformModuleRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
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

// The set of arguments for constructing a LocalTerraformModuleRepository resource.
type LocalTerraformModuleRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
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

func (LocalTerraformModuleRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localTerraformModuleRepositoryArgs)(nil)).Elem()
}

type LocalTerraformModuleRepositoryInput interface {
	pulumi.Input

	ToLocalTerraformModuleRepositoryOutput() LocalTerraformModuleRepositoryOutput
	ToLocalTerraformModuleRepositoryOutputWithContext(ctx context.Context) LocalTerraformModuleRepositoryOutput
}

func (*LocalTerraformModuleRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalTerraformModuleRepository)(nil)).Elem()
}

func (i *LocalTerraformModuleRepository) ToLocalTerraformModuleRepositoryOutput() LocalTerraformModuleRepositoryOutput {
	return i.ToLocalTerraformModuleRepositoryOutputWithContext(context.Background())
}

func (i *LocalTerraformModuleRepository) ToLocalTerraformModuleRepositoryOutputWithContext(ctx context.Context) LocalTerraformModuleRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalTerraformModuleRepositoryOutput)
}

// LocalTerraformModuleRepositoryArrayInput is an input type that accepts LocalTerraformModuleRepositoryArray and LocalTerraformModuleRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalTerraformModuleRepositoryArrayInput` via:
//
//	LocalTerraformModuleRepositoryArray{ LocalTerraformModuleRepositoryArgs{...} }
type LocalTerraformModuleRepositoryArrayInput interface {
	pulumi.Input

	ToLocalTerraformModuleRepositoryArrayOutput() LocalTerraformModuleRepositoryArrayOutput
	ToLocalTerraformModuleRepositoryArrayOutputWithContext(context.Context) LocalTerraformModuleRepositoryArrayOutput
}

type LocalTerraformModuleRepositoryArray []LocalTerraformModuleRepositoryInput

func (LocalTerraformModuleRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalTerraformModuleRepository)(nil)).Elem()
}

func (i LocalTerraformModuleRepositoryArray) ToLocalTerraformModuleRepositoryArrayOutput() LocalTerraformModuleRepositoryArrayOutput {
	return i.ToLocalTerraformModuleRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalTerraformModuleRepositoryArray) ToLocalTerraformModuleRepositoryArrayOutputWithContext(ctx context.Context) LocalTerraformModuleRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalTerraformModuleRepositoryArrayOutput)
}

// LocalTerraformModuleRepositoryMapInput is an input type that accepts LocalTerraformModuleRepositoryMap and LocalTerraformModuleRepositoryMapOutput values.
// You can construct a concrete instance of `LocalTerraformModuleRepositoryMapInput` via:
//
//	LocalTerraformModuleRepositoryMap{ "key": LocalTerraformModuleRepositoryArgs{...} }
type LocalTerraformModuleRepositoryMapInput interface {
	pulumi.Input

	ToLocalTerraformModuleRepositoryMapOutput() LocalTerraformModuleRepositoryMapOutput
	ToLocalTerraformModuleRepositoryMapOutputWithContext(context.Context) LocalTerraformModuleRepositoryMapOutput
}

type LocalTerraformModuleRepositoryMap map[string]LocalTerraformModuleRepositoryInput

func (LocalTerraformModuleRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalTerraformModuleRepository)(nil)).Elem()
}

func (i LocalTerraformModuleRepositoryMap) ToLocalTerraformModuleRepositoryMapOutput() LocalTerraformModuleRepositoryMapOutput {
	return i.ToLocalTerraformModuleRepositoryMapOutputWithContext(context.Background())
}

func (i LocalTerraformModuleRepositoryMap) ToLocalTerraformModuleRepositoryMapOutputWithContext(ctx context.Context) LocalTerraformModuleRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalTerraformModuleRepositoryMapOutput)
}

type LocalTerraformModuleRepositoryOutput struct{ *pulumi.OutputState }

func (LocalTerraformModuleRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalTerraformModuleRepository)(nil)).Elem()
}

func (o LocalTerraformModuleRepositoryOutput) ToLocalTerraformModuleRepositoryOutput() LocalTerraformModuleRepositoryOutput {
	return o
}

func (o LocalTerraformModuleRepositoryOutput) ToLocalTerraformModuleRepositoryOutputWithContext(ctx context.Context) LocalTerraformModuleRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o LocalTerraformModuleRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalTerraformModuleRepository) pulumi.BoolOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o LocalTerraformModuleRepositoryOutput) BlackedOut() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalTerraformModuleRepository) pulumi.BoolOutput { return v.BlackedOut }).(pulumi.BoolOutput)
}

// Public description.
func (o LocalTerraformModuleRepositoryOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalTerraformModuleRepository) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o LocalTerraformModuleRepositoryOutput) DownloadDirect() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalTerraformModuleRepository) pulumi.BoolOutput { return v.DownloadDirect }).(pulumi.BoolOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
// artifacts are excluded.
func (o LocalTerraformModuleRepositoryOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalTerraformModuleRepository) pulumi.StringOutput { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
func (o LocalTerraformModuleRepositoryOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalTerraformModuleRepository) pulumi.StringOutput { return v.IncludesPattern }).(pulumi.StringOutput)
}

// the identity key of the repo.
func (o LocalTerraformModuleRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalTerraformModuleRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o LocalTerraformModuleRepositoryOutput) Notes() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalTerraformModuleRepository) pulumi.StringOutput { return v.Notes }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o LocalTerraformModuleRepositoryOutput) PriorityResolution() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalTerraformModuleRepository) pulumi.BoolOutput { return v.PriorityResolution }).(pulumi.BoolOutput)
}

func (o LocalTerraformModuleRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalTerraformModuleRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o LocalTerraformModuleRepositoryOutput) ProjectKey() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalTerraformModuleRepository) pulumi.StringOutput { return v.ProjectKey }).(pulumi.StringOutput)
}

// List of property set name
func (o LocalTerraformModuleRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalTerraformModuleRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
// corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
func (o LocalTerraformModuleRepositoryOutput) RepoLayoutRef() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalTerraformModuleRepository) pulumi.StringOutput { return v.RepoLayoutRef }).(pulumi.StringOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o LocalTerraformModuleRepositoryOutput) XrayIndex() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalTerraformModuleRepository) pulumi.BoolOutput { return v.XrayIndex }).(pulumi.BoolOutput)
}

type LocalTerraformModuleRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalTerraformModuleRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalTerraformModuleRepository)(nil)).Elem()
}

func (o LocalTerraformModuleRepositoryArrayOutput) ToLocalTerraformModuleRepositoryArrayOutput() LocalTerraformModuleRepositoryArrayOutput {
	return o
}

func (o LocalTerraformModuleRepositoryArrayOutput) ToLocalTerraformModuleRepositoryArrayOutputWithContext(ctx context.Context) LocalTerraformModuleRepositoryArrayOutput {
	return o
}

func (o LocalTerraformModuleRepositoryArrayOutput) Index(i pulumi.IntInput) LocalTerraformModuleRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalTerraformModuleRepository {
		return vs[0].([]*LocalTerraformModuleRepository)[vs[1].(int)]
	}).(LocalTerraformModuleRepositoryOutput)
}

type LocalTerraformModuleRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalTerraformModuleRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalTerraformModuleRepository)(nil)).Elem()
}

func (o LocalTerraformModuleRepositoryMapOutput) ToLocalTerraformModuleRepositoryMapOutput() LocalTerraformModuleRepositoryMapOutput {
	return o
}

func (o LocalTerraformModuleRepositoryMapOutput) ToLocalTerraformModuleRepositoryMapOutputWithContext(ctx context.Context) LocalTerraformModuleRepositoryMapOutput {
	return o
}

func (o LocalTerraformModuleRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalTerraformModuleRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalTerraformModuleRepository {
		return vs[0].(map[string]*LocalTerraformModuleRepository)[vs[1].(string)]
	}).(LocalTerraformModuleRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalTerraformModuleRepositoryInput)(nil)).Elem(), &LocalTerraformModuleRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalTerraformModuleRepositoryArrayInput)(nil)).Elem(), LocalTerraformModuleRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalTerraformModuleRepositoryMapInput)(nil)).Elem(), LocalTerraformModuleRepositoryMap{})
	pulumi.RegisterOutputType(LocalTerraformModuleRepositoryOutput{})
	pulumi.RegisterOutputType(LocalTerraformModuleRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalTerraformModuleRepositoryMapOutput{})
}
