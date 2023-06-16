// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FederatedNugetRepository struct {
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
	// Delete all federated members on `terraform destroy` if set to `true`. Caution: it will delete all the repositories in
	// the federation on other Artifactory instances.
	CleanupOnDelete pulumi.BoolPtrOutput `pulumi:"cleanupOnDelete"`
	// Public description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrOutput `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringOutput `pulumi:"excludesPattern"`
	// Force basic authentication credentials in order to use this repository.
	ForceNugetAuthentication pulumi.BoolPtrOutput `pulumi:"forceNugetAuthentication"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringOutput `pulumi:"includesPattern"`
	// A mandatory identifier for the repository that must be unique. Must be 3 - 10 lowercase alphanumeric and hyphen
	// characters. It cannot begin with a number or contain spaces or special characters.
	Key pulumi.StringOutput `pulumi:"key"`
	// The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting,
	// older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots pulumi.IntPtrOutput `pulumi:"maxUniqueSnapshots"`
	// The list of Federated members. If a Federated member receives a request that does not include the repository URL, it
	// will automatically be added with the combination of the configured base URL and `key` field value. Note that each of the
	// federated members will need to have a base URL set. Please follow the
	// [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members FederatedNugetRepositoryMemberArrayOutput `pulumi:"members"`
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

// NewFederatedNugetRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedNugetRepository(ctx *pulumi.Context,
	name string, args *FederatedNugetRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedNugetRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FederatedNugetRepository
	err := ctx.RegisterResource("artifactory:index/federatedNugetRepository:FederatedNugetRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedNugetRepository gets an existing FederatedNugetRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedNugetRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedNugetRepositoryState, opts ...pulumi.ResourceOption) (*FederatedNugetRepository, error) {
	var resource FederatedNugetRepository
	err := ctx.ReadResource("artifactory:index/federatedNugetRepository:FederatedNugetRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedNugetRepository resources.
type federatedNugetRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect *bool `pulumi:"cdnRedirect"`
	// Delete all federated members on `terraform destroy` if set to `true`. Caution: it will delete all the repositories in
	// the federation on other Artifactory instances.
	CleanupOnDelete *bool `pulumi:"cleanupOnDelete"`
	// Public description.
	Description *string `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect *bool `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// Force basic authentication credentials in order to use this repository.
	ForceNugetAuthentication *bool `pulumi:"forceNugetAuthentication"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// A mandatory identifier for the repository that must be unique. Must be 3 - 10 lowercase alphanumeric and hyphen
	// characters. It cannot begin with a number or contain spaces or special characters.
	Key *string `pulumi:"key"`
	// The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting,
	// older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots *int `pulumi:"maxUniqueSnapshots"`
	// The list of Federated members. If a Federated member receives a request that does not include the repository URL, it
	// will automatically be added with the combination of the configured base URL and `key` field value. Note that each of the
	// federated members will need to have a base URL set. Please follow the
	// [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members []FederatedNugetRepositoryMember `pulumi:"members"`
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

type FederatedNugetRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolPtrInput
	// Delete all federated members on `terraform destroy` if set to `true`. Caution: it will delete all the repositories in
	// the federation on other Artifactory instances.
	CleanupOnDelete pulumi.BoolPtrInput
	// Public description.
	Description pulumi.StringPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// Force basic authentication credentials in order to use this repository.
	ForceNugetAuthentication pulumi.BoolPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// A mandatory identifier for the repository that must be unique. Must be 3 - 10 lowercase alphanumeric and hyphen
	// characters. It cannot begin with a number or contain spaces or special characters.
	Key pulumi.StringPtrInput
	// The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting,
	// older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots pulumi.IntPtrInput
	// The list of Federated members. If a Federated member receives a request that does not include the repository URL, it
	// will automatically be added with the combination of the configured base URL and `key` field value. Note that each of the
	// federated members will need to have a base URL set. Please follow the
	// [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members FederatedNugetRepositoryMemberArrayInput
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

func (FederatedNugetRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedNugetRepositoryState)(nil)).Elem()
}

type federatedNugetRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect *bool `pulumi:"cdnRedirect"`
	// Delete all federated members on `terraform destroy` if set to `true`. Caution: it will delete all the repositories in
	// the federation on other Artifactory instances.
	CleanupOnDelete *bool `pulumi:"cleanupOnDelete"`
	// Public description.
	Description *string `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect *bool `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// Force basic authentication credentials in order to use this repository.
	ForceNugetAuthentication *bool `pulumi:"forceNugetAuthentication"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// A mandatory identifier for the repository that must be unique. Must be 3 - 10 lowercase alphanumeric and hyphen
	// characters. It cannot begin with a number or contain spaces or special characters.
	Key string `pulumi:"key"`
	// The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting,
	// older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots *int `pulumi:"maxUniqueSnapshots"`
	// The list of Federated members. If a Federated member receives a request that does not include the repository URL, it
	// will automatically be added with the combination of the configured base URL and `key` field value. Note that each of the
	// federated members will need to have a base URL set. Please follow the
	// [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members []FederatedNugetRepositoryMember `pulumi:"members"`
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

// The set of arguments for constructing a FederatedNugetRepository resource.
type FederatedNugetRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolPtrInput
	// Delete all federated members on `terraform destroy` if set to `true`. Caution: it will delete all the repositories in
	// the federation on other Artifactory instances.
	CleanupOnDelete pulumi.BoolPtrInput
	// Public description.
	Description pulumi.StringPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// Force basic authentication credentials in order to use this repository.
	ForceNugetAuthentication pulumi.BoolPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// A mandatory identifier for the repository that must be unique. Must be 3 - 10 lowercase alphanumeric and hyphen
	// characters. It cannot begin with a number or contain spaces or special characters.
	Key pulumi.StringInput
	// The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting,
	// older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots pulumi.IntPtrInput
	// The list of Federated members. If a Federated member receives a request that does not include the repository URL, it
	// will automatically be added with the combination of the configured base URL and `key` field value. Note that each of the
	// federated members will need to have a base URL set. Please follow the
	// [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members FederatedNugetRepositoryMemberArrayInput
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

func (FederatedNugetRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedNugetRepositoryArgs)(nil)).Elem()
}

type FederatedNugetRepositoryInput interface {
	pulumi.Input

	ToFederatedNugetRepositoryOutput() FederatedNugetRepositoryOutput
	ToFederatedNugetRepositoryOutputWithContext(ctx context.Context) FederatedNugetRepositoryOutput
}

func (*FederatedNugetRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedNugetRepository)(nil)).Elem()
}

func (i *FederatedNugetRepository) ToFederatedNugetRepositoryOutput() FederatedNugetRepositoryOutput {
	return i.ToFederatedNugetRepositoryOutputWithContext(context.Background())
}

func (i *FederatedNugetRepository) ToFederatedNugetRepositoryOutputWithContext(ctx context.Context) FederatedNugetRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedNugetRepositoryOutput)
}

// FederatedNugetRepositoryArrayInput is an input type that accepts FederatedNugetRepositoryArray and FederatedNugetRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedNugetRepositoryArrayInput` via:
//
//	FederatedNugetRepositoryArray{ FederatedNugetRepositoryArgs{...} }
type FederatedNugetRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedNugetRepositoryArrayOutput() FederatedNugetRepositoryArrayOutput
	ToFederatedNugetRepositoryArrayOutputWithContext(context.Context) FederatedNugetRepositoryArrayOutput
}

type FederatedNugetRepositoryArray []FederatedNugetRepositoryInput

func (FederatedNugetRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedNugetRepository)(nil)).Elem()
}

func (i FederatedNugetRepositoryArray) ToFederatedNugetRepositoryArrayOutput() FederatedNugetRepositoryArrayOutput {
	return i.ToFederatedNugetRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedNugetRepositoryArray) ToFederatedNugetRepositoryArrayOutputWithContext(ctx context.Context) FederatedNugetRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedNugetRepositoryArrayOutput)
}

// FederatedNugetRepositoryMapInput is an input type that accepts FederatedNugetRepositoryMap and FederatedNugetRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedNugetRepositoryMapInput` via:
//
//	FederatedNugetRepositoryMap{ "key": FederatedNugetRepositoryArgs{...} }
type FederatedNugetRepositoryMapInput interface {
	pulumi.Input

	ToFederatedNugetRepositoryMapOutput() FederatedNugetRepositoryMapOutput
	ToFederatedNugetRepositoryMapOutputWithContext(context.Context) FederatedNugetRepositoryMapOutput
}

type FederatedNugetRepositoryMap map[string]FederatedNugetRepositoryInput

func (FederatedNugetRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedNugetRepository)(nil)).Elem()
}

func (i FederatedNugetRepositoryMap) ToFederatedNugetRepositoryMapOutput() FederatedNugetRepositoryMapOutput {
	return i.ToFederatedNugetRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedNugetRepositoryMap) ToFederatedNugetRepositoryMapOutputWithContext(ctx context.Context) FederatedNugetRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedNugetRepositoryMapOutput)
}

type FederatedNugetRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedNugetRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedNugetRepository)(nil)).Elem()
}

func (o FederatedNugetRepositoryOutput) ToFederatedNugetRepositoryOutput() FederatedNugetRepositoryOutput {
	return o
}

func (o FederatedNugetRepositoryOutput) ToFederatedNugetRepositoryOutputWithContext(ctx context.Context) FederatedNugetRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o FederatedNugetRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedNugetRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o FederatedNugetRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedNugetRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o FederatedNugetRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedNugetRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

// Delete all federated members on `terraform destroy` if set to `true`. Caution: it will delete all the repositories in
// the federation on other Artifactory instances.
func (o FederatedNugetRepositoryOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedNugetRepository) pulumi.BoolPtrOutput { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

// Public description.
func (o FederatedNugetRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedNugetRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o FederatedNugetRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedNugetRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
// artifacts are excluded.
func (o FederatedNugetRepositoryOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedNugetRepository) pulumi.StringOutput { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// Force basic authentication credentials in order to use this repository.
func (o FederatedNugetRepositoryOutput) ForceNugetAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedNugetRepository) pulumi.BoolPtrOutput { return v.ForceNugetAuthentication }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o FederatedNugetRepositoryOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedNugetRepository) pulumi.StringOutput { return v.IncludesPattern }).(pulumi.StringOutput)
}

// A mandatory identifier for the repository that must be unique. Must be 3 - 10 lowercase alphanumeric and hyphen
// characters. It cannot begin with a number or contain spaces or special characters.
func (o FederatedNugetRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedNugetRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting,
// older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
func (o FederatedNugetRepositoryOutput) MaxUniqueSnapshots() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FederatedNugetRepository) pulumi.IntPtrOutput { return v.MaxUniqueSnapshots }).(pulumi.IntPtrOutput)
}

// The list of Federated members. If a Federated member receives a request that does not include the repository URL, it
// will automatically be added with the combination of the configured base URL and `key` field value. Note that each of the
// federated members will need to have a base URL set. Please follow the
// [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o FederatedNugetRepositoryOutput) Members() FederatedNugetRepositoryMemberArrayOutput {
	return o.ApplyT(func(v *FederatedNugetRepository) FederatedNugetRepositoryMemberArrayOutput { return v.Members }).(FederatedNugetRepositoryMemberArrayOutput)
}

// Internal description.
func (o FederatedNugetRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedNugetRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o FederatedNugetRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedNugetRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o FederatedNugetRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedNugetRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o FederatedNugetRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedNugetRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o FederatedNugetRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedNugetRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o FederatedNugetRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedNugetRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o FederatedNugetRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedNugetRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o FederatedNugetRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedNugetRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type FederatedNugetRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedNugetRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedNugetRepository)(nil)).Elem()
}

func (o FederatedNugetRepositoryArrayOutput) ToFederatedNugetRepositoryArrayOutput() FederatedNugetRepositoryArrayOutput {
	return o
}

func (o FederatedNugetRepositoryArrayOutput) ToFederatedNugetRepositoryArrayOutputWithContext(ctx context.Context) FederatedNugetRepositoryArrayOutput {
	return o
}

func (o FederatedNugetRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedNugetRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedNugetRepository {
		return vs[0].([]*FederatedNugetRepository)[vs[1].(int)]
	}).(FederatedNugetRepositoryOutput)
}

type FederatedNugetRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedNugetRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedNugetRepository)(nil)).Elem()
}

func (o FederatedNugetRepositoryMapOutput) ToFederatedNugetRepositoryMapOutput() FederatedNugetRepositoryMapOutput {
	return o
}

func (o FederatedNugetRepositoryMapOutput) ToFederatedNugetRepositoryMapOutputWithContext(ctx context.Context) FederatedNugetRepositoryMapOutput {
	return o
}

func (o FederatedNugetRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedNugetRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedNugetRepository {
		return vs[0].(map[string]*FederatedNugetRepository)[vs[1].(string)]
	}).(FederatedNugetRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedNugetRepositoryInput)(nil)).Elem(), &FederatedNugetRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedNugetRepositoryArrayInput)(nil)).Elem(), FederatedNugetRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedNugetRepositoryMapInput)(nil)).Elem(), FederatedNugetRepositoryMap{})
	pulumi.RegisterOutputType(FederatedNugetRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedNugetRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedNugetRepositoryMapOutput{})
}
