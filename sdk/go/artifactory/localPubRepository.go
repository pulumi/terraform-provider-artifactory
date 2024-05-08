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

// Creates a local Pub repository.
//
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
//			_, err := artifactory.NewLocalPubRepository(ctx, "terraform-local-test-pub-repo", &artifactory.LocalPubRepositoryArgs{
//				Key: pulumi.String("terraform-local-test-pub-repo"),
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
// $ pulumi import artifactory:index/localPubRepository:LocalPubRepository terraform-local-test-pub-repo terraform-local-test-pub-repo
// ```
type LocalPubRepository struct {
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
	PriorityResolution  pulumi.BoolPtrOutput     `pulumi:"priorityResolution"`
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

// NewLocalPubRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalPubRepository(ctx *pulumi.Context,
	name string, args *LocalPubRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalPubRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LocalPubRepository
	err := ctx.RegisterResource("artifactory:index/localPubRepository:LocalPubRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalPubRepository gets an existing LocalPubRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalPubRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalPubRepositoryState, opts ...pulumi.ResourceOption) (*LocalPubRepository, error) {
	var resource LocalPubRepository
	err := ctx.ReadResource("artifactory:index/localPubRepository:LocalPubRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalPubRepository resources.
type localPubRepositoryState struct {
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
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
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

type LocalPubRepositoryState struct {
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
	PriorityResolution  pulumi.BoolPtrInput
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

func (LocalPubRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localPubRepositoryState)(nil)).Elem()
}

type localPubRepositoryArgs struct {
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
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
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

// The set of arguments for constructing a LocalPubRepository resource.
type LocalPubRepositoryArgs struct {
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
	PriorityResolution  pulumi.BoolPtrInput
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

func (LocalPubRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localPubRepositoryArgs)(nil)).Elem()
}

type LocalPubRepositoryInput interface {
	pulumi.Input

	ToLocalPubRepositoryOutput() LocalPubRepositoryOutput
	ToLocalPubRepositoryOutputWithContext(ctx context.Context) LocalPubRepositoryOutput
}

func (*LocalPubRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalPubRepository)(nil)).Elem()
}

func (i *LocalPubRepository) ToLocalPubRepositoryOutput() LocalPubRepositoryOutput {
	return i.ToLocalPubRepositoryOutputWithContext(context.Background())
}

func (i *LocalPubRepository) ToLocalPubRepositoryOutputWithContext(ctx context.Context) LocalPubRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalPubRepositoryOutput)
}

// LocalPubRepositoryArrayInput is an input type that accepts LocalPubRepositoryArray and LocalPubRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalPubRepositoryArrayInput` via:
//
//	LocalPubRepositoryArray{ LocalPubRepositoryArgs{...} }
type LocalPubRepositoryArrayInput interface {
	pulumi.Input

	ToLocalPubRepositoryArrayOutput() LocalPubRepositoryArrayOutput
	ToLocalPubRepositoryArrayOutputWithContext(context.Context) LocalPubRepositoryArrayOutput
}

type LocalPubRepositoryArray []LocalPubRepositoryInput

func (LocalPubRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalPubRepository)(nil)).Elem()
}

func (i LocalPubRepositoryArray) ToLocalPubRepositoryArrayOutput() LocalPubRepositoryArrayOutput {
	return i.ToLocalPubRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalPubRepositoryArray) ToLocalPubRepositoryArrayOutputWithContext(ctx context.Context) LocalPubRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalPubRepositoryArrayOutput)
}

// LocalPubRepositoryMapInput is an input type that accepts LocalPubRepositoryMap and LocalPubRepositoryMapOutput values.
// You can construct a concrete instance of `LocalPubRepositoryMapInput` via:
//
//	LocalPubRepositoryMap{ "key": LocalPubRepositoryArgs{...} }
type LocalPubRepositoryMapInput interface {
	pulumi.Input

	ToLocalPubRepositoryMapOutput() LocalPubRepositoryMapOutput
	ToLocalPubRepositoryMapOutputWithContext(context.Context) LocalPubRepositoryMapOutput
}

type LocalPubRepositoryMap map[string]LocalPubRepositoryInput

func (LocalPubRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalPubRepository)(nil)).Elem()
}

func (i LocalPubRepositoryMap) ToLocalPubRepositoryMapOutput() LocalPubRepositoryMapOutput {
	return i.ToLocalPubRepositoryMapOutputWithContext(context.Background())
}

func (i LocalPubRepositoryMap) ToLocalPubRepositoryMapOutputWithContext(ctx context.Context) LocalPubRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalPubRepositoryMapOutput)
}

type LocalPubRepositoryOutput struct{ *pulumi.OutputState }

func (LocalPubRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalPubRepository)(nil)).Elem()
}

func (o LocalPubRepositoryOutput) ToLocalPubRepositoryOutput() LocalPubRepositoryOutput {
	return o
}

func (o LocalPubRepositoryOutput) ToLocalPubRepositoryOutputWithContext(ctx context.Context) LocalPubRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o LocalPubRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalPubRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o LocalPubRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalPubRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o LocalPubRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalPubRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

// Public description.
func (o LocalPubRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalPubRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o LocalPubRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalPubRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o LocalPubRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalPubRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o LocalPubRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalPubRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// the identity key of the repo.
func (o LocalPubRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalPubRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o LocalPubRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalPubRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LocalPubRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalPubRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o LocalPubRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalPubRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LocalPubRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalPubRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o LocalPubRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalPubRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o LocalPubRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalPubRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o LocalPubRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalPubRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o LocalPubRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalPubRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type LocalPubRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalPubRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalPubRepository)(nil)).Elem()
}

func (o LocalPubRepositoryArrayOutput) ToLocalPubRepositoryArrayOutput() LocalPubRepositoryArrayOutput {
	return o
}

func (o LocalPubRepositoryArrayOutput) ToLocalPubRepositoryArrayOutputWithContext(ctx context.Context) LocalPubRepositoryArrayOutput {
	return o
}

func (o LocalPubRepositoryArrayOutput) Index(i pulumi.IntInput) LocalPubRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalPubRepository {
		return vs[0].([]*LocalPubRepository)[vs[1].(int)]
	}).(LocalPubRepositoryOutput)
}

type LocalPubRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalPubRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalPubRepository)(nil)).Elem()
}

func (o LocalPubRepositoryMapOutput) ToLocalPubRepositoryMapOutput() LocalPubRepositoryMapOutput {
	return o
}

func (o LocalPubRepositoryMapOutput) ToLocalPubRepositoryMapOutputWithContext(ctx context.Context) LocalPubRepositoryMapOutput {
	return o
}

func (o LocalPubRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalPubRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalPubRepository {
		return vs[0].(map[string]*LocalPubRepository)[vs[1].(string)]
	}).(LocalPubRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalPubRepositoryInput)(nil)).Elem(), &LocalPubRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalPubRepositoryArrayInput)(nil)).Elem(), LocalPubRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalPubRepositoryMapInput)(nil)).Elem(), LocalPubRepositoryMap{})
	pulumi.RegisterOutputType(LocalPubRepositoryOutput{})
	pulumi.RegisterOutputType(LocalPubRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalPubRepositoryMapOutput{})
}
