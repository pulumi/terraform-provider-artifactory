// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LocalCargoRepository struct {
	pulumi.CustomResourceState

	// (Optional) Cargo client does not send credentials when performing download and search for crates. Enable this to allow
	// anonymous access to these resources (only), note that this will override the security anonymous access option. Default
	// value is 'false'.
	AnonymousAccess pulumi.BoolPtrOutput `pulumi:"anonymousAccess"`
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrOutput `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut  pulumi.BoolPtrOutput   `pulumi:"blackedOut"`
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrOutput `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringOutput `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern         pulumi.StringOutput      `pulumi:"includesPattern"`
	IndexCompressionFormats pulumi.StringArrayOutput `pulumi:"indexCompressionFormats"`
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or
	// special characters.
	Key         pulumi.StringOutput    `pulumi:"key"`
	Notes       pulumi.StringPtrOutput `pulumi:"notes"`
	PackageType pulumi.StringOutput    `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrOutput `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey pulumi.StringPtrOutput `pulumi:"projectKey"`
	// List of property set name
	PropertySets pulumi.StringArrayOutput `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrOutput `pulumi:"xrayIndex"`
}

// NewLocalCargoRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalCargoRepository(ctx *pulumi.Context,
	name string, args *LocalCargoRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalCargoRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource LocalCargoRepository
	err := ctx.RegisterResource("artifactory:index/localCargoRepository:LocalCargoRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalCargoRepository gets an existing LocalCargoRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalCargoRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalCargoRepositoryState, opts ...pulumi.ResourceOption) (*LocalCargoRepository, error) {
	var resource LocalCargoRepository
	err := ctx.ReadResource("artifactory:index/localCargoRepository:LocalCargoRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalCargoRepository resources.
type localCargoRepositoryState struct {
	// (Optional) Cargo client does not send credentials when performing download and search for crates. Enable this to allow
	// anonymous access to these resources (only), note that this will override the security anonymous access option. Default
	// value is 'false'.
	AnonymousAccess *bool `pulumi:"anonymousAccess"`
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut  *bool   `pulumi:"blackedOut"`
	Description *string `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect *bool `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern         *string  `pulumi:"includesPattern"`
	IndexCompressionFormats []string `pulumi:"indexCompressionFormats"`
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or
	// special characters.
	Key         *string `pulumi:"key"`
	Notes       *string `pulumi:"notes"`
	PackageType *string `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// List of property set name
	PropertySets []string `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

type LocalCargoRepositoryState struct {
	// (Optional) Cargo client does not send credentials when performing download and search for crates. Enable this to allow
	// anonymous access to these resources (only), note that this will override the security anonymous access option. Default
	// value is 'false'.
	AnonymousAccess pulumi.BoolPtrInput
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut  pulumi.BoolPtrInput
	Description pulumi.StringPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern         pulumi.StringPtrInput
	IndexCompressionFormats pulumi.StringArrayInput
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or
	// special characters.
	Key         pulumi.StringPtrInput
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// List of property set name
	PropertySets pulumi.StringArrayInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (LocalCargoRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localCargoRepositoryState)(nil)).Elem()
}

type localCargoRepositoryArgs struct {
	// (Optional) Cargo client does not send credentials when performing download and search for crates. Enable this to allow
	// anonymous access to these resources (only), note that this will override the security anonymous access option. Default
	// value is 'false'.
	AnonymousAccess *bool `pulumi:"anonymousAccess"`
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut  *bool   `pulumi:"blackedOut"`
	Description *string `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect *bool `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern         *string  `pulumi:"includesPattern"`
	IndexCompressionFormats []string `pulumi:"indexCompressionFormats"`
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or
	// special characters.
	Key   string  `pulumi:"key"`
	Notes *string `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// List of property set name
	PropertySets []string `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a LocalCargoRepository resource.
type LocalCargoRepositoryArgs struct {
	// (Optional) Cargo client does not send credentials when performing download and search for crates. Enable this to allow
	// anonymous access to these resources (only), note that this will override the security anonymous access option. Default
	// value is 'false'.
	AnonymousAccess pulumi.BoolPtrInput
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut  pulumi.BoolPtrInput
	Description pulumi.StringPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern         pulumi.StringPtrInput
	IndexCompressionFormats pulumi.StringArrayInput
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or
	// special characters.
	Key   pulumi.StringInput
	Notes pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// List of property set name
	PropertySets pulumi.StringArrayInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (LocalCargoRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localCargoRepositoryArgs)(nil)).Elem()
}

type LocalCargoRepositoryInput interface {
	pulumi.Input

	ToLocalCargoRepositoryOutput() LocalCargoRepositoryOutput
	ToLocalCargoRepositoryOutputWithContext(ctx context.Context) LocalCargoRepositoryOutput
}

func (*LocalCargoRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalCargoRepository)(nil)).Elem()
}

func (i *LocalCargoRepository) ToLocalCargoRepositoryOutput() LocalCargoRepositoryOutput {
	return i.ToLocalCargoRepositoryOutputWithContext(context.Background())
}

func (i *LocalCargoRepository) ToLocalCargoRepositoryOutputWithContext(ctx context.Context) LocalCargoRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalCargoRepositoryOutput)
}

// LocalCargoRepositoryArrayInput is an input type that accepts LocalCargoRepositoryArray and LocalCargoRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalCargoRepositoryArrayInput` via:
//
//          LocalCargoRepositoryArray{ LocalCargoRepositoryArgs{...} }
type LocalCargoRepositoryArrayInput interface {
	pulumi.Input

	ToLocalCargoRepositoryArrayOutput() LocalCargoRepositoryArrayOutput
	ToLocalCargoRepositoryArrayOutputWithContext(context.Context) LocalCargoRepositoryArrayOutput
}

type LocalCargoRepositoryArray []LocalCargoRepositoryInput

func (LocalCargoRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalCargoRepository)(nil)).Elem()
}

func (i LocalCargoRepositoryArray) ToLocalCargoRepositoryArrayOutput() LocalCargoRepositoryArrayOutput {
	return i.ToLocalCargoRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalCargoRepositoryArray) ToLocalCargoRepositoryArrayOutputWithContext(ctx context.Context) LocalCargoRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalCargoRepositoryArrayOutput)
}

// LocalCargoRepositoryMapInput is an input type that accepts LocalCargoRepositoryMap and LocalCargoRepositoryMapOutput values.
// You can construct a concrete instance of `LocalCargoRepositoryMapInput` via:
//
//          LocalCargoRepositoryMap{ "key": LocalCargoRepositoryArgs{...} }
type LocalCargoRepositoryMapInput interface {
	pulumi.Input

	ToLocalCargoRepositoryMapOutput() LocalCargoRepositoryMapOutput
	ToLocalCargoRepositoryMapOutputWithContext(context.Context) LocalCargoRepositoryMapOutput
}

type LocalCargoRepositoryMap map[string]LocalCargoRepositoryInput

func (LocalCargoRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalCargoRepository)(nil)).Elem()
}

func (i LocalCargoRepositoryMap) ToLocalCargoRepositoryMapOutput() LocalCargoRepositoryMapOutput {
	return i.ToLocalCargoRepositoryMapOutputWithContext(context.Background())
}

func (i LocalCargoRepositoryMap) ToLocalCargoRepositoryMapOutputWithContext(ctx context.Context) LocalCargoRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalCargoRepositoryMapOutput)
}

type LocalCargoRepositoryOutput struct{ *pulumi.OutputState }

func (LocalCargoRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalCargoRepository)(nil)).Elem()
}

func (o LocalCargoRepositoryOutput) ToLocalCargoRepositoryOutput() LocalCargoRepositoryOutput {
	return o
}

func (o LocalCargoRepositoryOutput) ToLocalCargoRepositoryOutputWithContext(ctx context.Context) LocalCargoRepositoryOutput {
	return o
}

type LocalCargoRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalCargoRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalCargoRepository)(nil)).Elem()
}

func (o LocalCargoRepositoryArrayOutput) ToLocalCargoRepositoryArrayOutput() LocalCargoRepositoryArrayOutput {
	return o
}

func (o LocalCargoRepositoryArrayOutput) ToLocalCargoRepositoryArrayOutputWithContext(ctx context.Context) LocalCargoRepositoryArrayOutput {
	return o
}

func (o LocalCargoRepositoryArrayOutput) Index(i pulumi.IntInput) LocalCargoRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalCargoRepository {
		return vs[0].([]*LocalCargoRepository)[vs[1].(int)]
	}).(LocalCargoRepositoryOutput)
}

type LocalCargoRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalCargoRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalCargoRepository)(nil)).Elem()
}

func (o LocalCargoRepositoryMapOutput) ToLocalCargoRepositoryMapOutput() LocalCargoRepositoryMapOutput {
	return o
}

func (o LocalCargoRepositoryMapOutput) ToLocalCargoRepositoryMapOutputWithContext(ctx context.Context) LocalCargoRepositoryMapOutput {
	return o
}

func (o LocalCargoRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalCargoRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalCargoRepository {
		return vs[0].(map[string]*LocalCargoRepository)[vs[1].(string)]
	}).(LocalCargoRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalCargoRepositoryInput)(nil)).Elem(), &LocalCargoRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalCargoRepositoryArrayInput)(nil)).Elem(), LocalCargoRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalCargoRepositoryMapInput)(nil)).Elem(), LocalCargoRepositoryMap{})
	pulumi.RegisterOutputType(LocalCargoRepositoryOutput{})
	pulumi.RegisterOutputType(LocalCargoRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalCargoRepositoryMapOutput{})
}
