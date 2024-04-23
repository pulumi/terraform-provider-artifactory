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

// Creates a local OCI repository.
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
//			_, err := artifactory.NewLocalOciRepository(ctx, "my-oci-local", &artifactory.LocalOciRepositoryArgs{
//				Key:           pulumi.String("my-oci-local"),
//				TagRetention:  pulumi.Int(3),
//				MaxUniqueTags: pulumi.Int(5),
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
// $ pulumi import artifactory:index/localOciRepository:LocalOciRepository my-oci-local my-oci-local
// ```
type LocalOciRepository struct {
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
	// The maximum number of unique tags of a single OCI image to store in this repository. Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit.
	MaxUniqueTags pulumi.IntPtrOutput `pulumi:"maxUniqueTags"`
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
	// If greater than 1, overwritten tags will be saved by their digest, up to the set up number.
	TagRetention pulumi.IntPtrOutput `pulumi:"tagRetention"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrOutput `pulumi:"xrayIndex"`
}

// NewLocalOciRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalOciRepository(ctx *pulumi.Context,
	name string, args *LocalOciRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalOciRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LocalOciRepository
	err := ctx.RegisterResource("artifactory:index/localOciRepository:LocalOciRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalOciRepository gets an existing LocalOciRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalOciRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalOciRepositoryState, opts ...pulumi.ResourceOption) (*LocalOciRepository, error) {
	var resource LocalOciRepository
	err := ctx.ReadResource("artifactory:index/localOciRepository:LocalOciRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalOciRepository resources.
type localOciRepositoryState struct {
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
	// The maximum number of unique tags of a single OCI image to store in this repository. Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit.
	MaxUniqueTags *int `pulumi:"maxUniqueTags"`
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
	// If greater than 1, overwritten tags will be saved by their digest, up to the set up number.
	TagRetention *int `pulumi:"tagRetention"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

type LocalOciRepositoryState struct {
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
	// The maximum number of unique tags of a single OCI image to store in this repository. Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit.
	MaxUniqueTags pulumi.IntPtrInput
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
	// If greater than 1, overwritten tags will be saved by their digest, up to the set up number.
	TagRetention pulumi.IntPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (LocalOciRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localOciRepositoryState)(nil)).Elem()
}

type localOciRepositoryArgs struct {
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
	// The maximum number of unique tags of a single OCI image to store in this repository. Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit.
	MaxUniqueTags *int `pulumi:"maxUniqueTags"`
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
	// If greater than 1, overwritten tags will be saved by their digest, up to the set up number.
	TagRetention *int `pulumi:"tagRetention"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a LocalOciRepository resource.
type LocalOciRepositoryArgs struct {
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
	// The maximum number of unique tags of a single OCI image to store in this repository. Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit.
	MaxUniqueTags pulumi.IntPtrInput
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
	// If greater than 1, overwritten tags will be saved by their digest, up to the set up number.
	TagRetention pulumi.IntPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (LocalOciRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localOciRepositoryArgs)(nil)).Elem()
}

type LocalOciRepositoryInput interface {
	pulumi.Input

	ToLocalOciRepositoryOutput() LocalOciRepositoryOutput
	ToLocalOciRepositoryOutputWithContext(ctx context.Context) LocalOciRepositoryOutput
}

func (*LocalOciRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalOciRepository)(nil)).Elem()
}

func (i *LocalOciRepository) ToLocalOciRepositoryOutput() LocalOciRepositoryOutput {
	return i.ToLocalOciRepositoryOutputWithContext(context.Background())
}

func (i *LocalOciRepository) ToLocalOciRepositoryOutputWithContext(ctx context.Context) LocalOciRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalOciRepositoryOutput)
}

// LocalOciRepositoryArrayInput is an input type that accepts LocalOciRepositoryArray and LocalOciRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalOciRepositoryArrayInput` via:
//
//	LocalOciRepositoryArray{ LocalOciRepositoryArgs{...} }
type LocalOciRepositoryArrayInput interface {
	pulumi.Input

	ToLocalOciRepositoryArrayOutput() LocalOciRepositoryArrayOutput
	ToLocalOciRepositoryArrayOutputWithContext(context.Context) LocalOciRepositoryArrayOutput
}

type LocalOciRepositoryArray []LocalOciRepositoryInput

func (LocalOciRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalOciRepository)(nil)).Elem()
}

func (i LocalOciRepositoryArray) ToLocalOciRepositoryArrayOutput() LocalOciRepositoryArrayOutput {
	return i.ToLocalOciRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalOciRepositoryArray) ToLocalOciRepositoryArrayOutputWithContext(ctx context.Context) LocalOciRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalOciRepositoryArrayOutput)
}

// LocalOciRepositoryMapInput is an input type that accepts LocalOciRepositoryMap and LocalOciRepositoryMapOutput values.
// You can construct a concrete instance of `LocalOciRepositoryMapInput` via:
//
//	LocalOciRepositoryMap{ "key": LocalOciRepositoryArgs{...} }
type LocalOciRepositoryMapInput interface {
	pulumi.Input

	ToLocalOciRepositoryMapOutput() LocalOciRepositoryMapOutput
	ToLocalOciRepositoryMapOutputWithContext(context.Context) LocalOciRepositoryMapOutput
}

type LocalOciRepositoryMap map[string]LocalOciRepositoryInput

func (LocalOciRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalOciRepository)(nil)).Elem()
}

func (i LocalOciRepositoryMap) ToLocalOciRepositoryMapOutput() LocalOciRepositoryMapOutput {
	return i.ToLocalOciRepositoryMapOutputWithContext(context.Background())
}

func (i LocalOciRepositoryMap) ToLocalOciRepositoryMapOutputWithContext(ctx context.Context) LocalOciRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalOciRepositoryMapOutput)
}

type LocalOciRepositoryOutput struct{ *pulumi.OutputState }

func (LocalOciRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalOciRepository)(nil)).Elem()
}

func (o LocalOciRepositoryOutput) ToLocalOciRepositoryOutput() LocalOciRepositoryOutput {
	return o
}

func (o LocalOciRepositoryOutput) ToLocalOciRepositoryOutputWithContext(ctx context.Context) LocalOciRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o LocalOciRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalOciRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o LocalOciRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalOciRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o LocalOciRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalOciRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

// Public description.
func (o LocalOciRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalOciRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o LocalOciRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalOciRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o LocalOciRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalOciRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o LocalOciRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalOciRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// the identity key of the repo.
func (o LocalOciRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalOciRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The maximum number of unique tags of a single OCI image to store in this repository. Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit.
func (o LocalOciRepositoryOutput) MaxUniqueTags() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LocalOciRepository) pulumi.IntPtrOutput { return v.MaxUniqueTags }).(pulumi.IntPtrOutput)
}

// Internal description.
func (o LocalOciRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalOciRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LocalOciRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalOciRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o LocalOciRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalOciRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o LocalOciRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalOciRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o LocalOciRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalOciRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o LocalOciRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalOciRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o LocalOciRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalOciRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// If greater than 1, overwritten tags will be saved by their digest, up to the set up number.
func (o LocalOciRepositoryOutput) TagRetention() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LocalOciRepository) pulumi.IntPtrOutput { return v.TagRetention }).(pulumi.IntPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o LocalOciRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalOciRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type LocalOciRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalOciRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalOciRepository)(nil)).Elem()
}

func (o LocalOciRepositoryArrayOutput) ToLocalOciRepositoryArrayOutput() LocalOciRepositoryArrayOutput {
	return o
}

func (o LocalOciRepositoryArrayOutput) ToLocalOciRepositoryArrayOutputWithContext(ctx context.Context) LocalOciRepositoryArrayOutput {
	return o
}

func (o LocalOciRepositoryArrayOutput) Index(i pulumi.IntInput) LocalOciRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalOciRepository {
		return vs[0].([]*LocalOciRepository)[vs[1].(int)]
	}).(LocalOciRepositoryOutput)
}

type LocalOciRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalOciRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalOciRepository)(nil)).Elem()
}

func (o LocalOciRepositoryMapOutput) ToLocalOciRepositoryMapOutput() LocalOciRepositoryMapOutput {
	return o
}

func (o LocalOciRepositoryMapOutput) ToLocalOciRepositoryMapOutputWithContext(ctx context.Context) LocalOciRepositoryMapOutput {
	return o
}

func (o LocalOciRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalOciRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalOciRepository {
		return vs[0].(map[string]*LocalOciRepository)[vs[1].(string)]
	}).(LocalOciRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalOciRepositoryInput)(nil)).Elem(), &LocalOciRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalOciRepositoryArrayInput)(nil)).Elem(), LocalOciRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalOciRepositoryMapInput)(nil)).Elem(), LocalOciRepositoryMap{})
	pulumi.RegisterOutputType(LocalOciRepositoryOutput{})
	pulumi.RegisterOutputType(LocalOciRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalOciRepositoryMapOutput{})
}
