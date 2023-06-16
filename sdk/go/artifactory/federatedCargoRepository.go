// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a federated Cargo repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v1/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewFederatedCargoRepository(ctx, "terraform-federated-test-cargo-repo", &artifactory.FederatedCargoRepositoryArgs{
//				Key: "terraform-federated-test-cargo-repo",
//				Members: []map[string]interface{}{
//					map[string]interface{}{
//						"enabled": true,
//						"url":     "http://tempurl.org/artifactory/terraform-federated-test-cargo-repo",
//					},
//					map[string]interface{}{
//						"enabled": true,
//						"url":     "http://tempurl2.org/artifactory/terraform-federated-test-cargo-repo-2",
//					},
//				},
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
// Federated repositories can be imported using their name, e.g.
//
// ```sh
//
//	$ pulumi import artifactory:index/federatedCargoRepository:FederatedCargoRepository terraform-federated-test-cargo-repo terraform-federated-test-cargo-repo
//
// ```
type FederatedCargoRepository struct {
	pulumi.CustomResourceState

	// Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous
	// access to these resources (only), note that this will override the security anonymous access option. Default value is
	// 'false'.
	AnonymousAccess pulumi.BoolPtrOutput `pulumi:"anonymousAccess"`
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
	// Enable internal index support based on Cargo sparse index specifications, instead of the default git index. Default
	// value is 'false'.
	EnableSparseIndex pulumi.BoolPtrOutput `pulumi:"enableSparseIndex"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringOutput `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern         pulumi.StringOutput      `pulumi:"includesPattern"`
	IndexCompressionFormats pulumi.StringArrayOutput `pulumi:"indexCompressionFormats"`
	// the identity key of the repo.
	Key pulumi.StringOutput `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members FederatedCargoRepositoryMemberArrayOutput `pulumi:"members"`
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

// NewFederatedCargoRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedCargoRepository(ctx *pulumi.Context,
	name string, args *FederatedCargoRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedCargoRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FederatedCargoRepository
	err := ctx.RegisterResource("artifactory:index/federatedCargoRepository:FederatedCargoRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedCargoRepository gets an existing FederatedCargoRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedCargoRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedCargoRepositoryState, opts ...pulumi.ResourceOption) (*FederatedCargoRepository, error) {
	var resource FederatedCargoRepository
	err := ctx.ReadResource("artifactory:index/federatedCargoRepository:FederatedCargoRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedCargoRepository resources.
type federatedCargoRepositoryState struct {
	// Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous
	// access to these resources (only), note that this will override the security anonymous access option. Default value is
	// 'false'.
	AnonymousAccess *bool `pulumi:"anonymousAccess"`
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
	// Enable internal index support based on Cargo sparse index specifications, instead of the default git index. Default
	// value is 'false'.
	EnableSparseIndex *bool `pulumi:"enableSparseIndex"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern         *string  `pulumi:"includesPattern"`
	IndexCompressionFormats []string `pulumi:"indexCompressionFormats"`
	// the identity key of the repo.
	Key *string `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members []FederatedCargoRepositoryMember `pulumi:"members"`
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

type FederatedCargoRepositoryState struct {
	// Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous
	// access to these resources (only), note that this will override the security anonymous access option. Default value is
	// 'false'.
	AnonymousAccess pulumi.BoolPtrInput
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
	// Enable internal index support based on Cargo sparse index specifications, instead of the default git index. Default
	// value is 'false'.
	EnableSparseIndex pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern         pulumi.StringPtrInput
	IndexCompressionFormats pulumi.StringArrayInput
	// the identity key of the repo.
	Key pulumi.StringPtrInput
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members FederatedCargoRepositoryMemberArrayInput
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

func (FederatedCargoRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedCargoRepositoryState)(nil)).Elem()
}

type federatedCargoRepositoryArgs struct {
	// Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous
	// access to these resources (only), note that this will override the security anonymous access option. Default value is
	// 'false'.
	AnonymousAccess *bool `pulumi:"anonymousAccess"`
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
	// Enable internal index support based on Cargo sparse index specifications, instead of the default git index. Default
	// value is 'false'.
	EnableSparseIndex *bool `pulumi:"enableSparseIndex"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern         *string  `pulumi:"includesPattern"`
	IndexCompressionFormats []string `pulumi:"indexCompressionFormats"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members []FederatedCargoRepositoryMember `pulumi:"members"`
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

// The set of arguments for constructing a FederatedCargoRepository resource.
type FederatedCargoRepositoryArgs struct {
	// Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous
	// access to these resources (only), note that this will override the security anonymous access option. Default value is
	// 'false'.
	AnonymousAccess pulumi.BoolPtrInput
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
	// Enable internal index support based on Cargo sparse index specifications, instead of the default git index. Default
	// value is 'false'.
	EnableSparseIndex pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern         pulumi.StringPtrInput
	IndexCompressionFormats pulumi.StringArrayInput
	// the identity key of the repo.
	Key pulumi.StringInput
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members FederatedCargoRepositoryMemberArrayInput
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

func (FederatedCargoRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedCargoRepositoryArgs)(nil)).Elem()
}

type FederatedCargoRepositoryInput interface {
	pulumi.Input

	ToFederatedCargoRepositoryOutput() FederatedCargoRepositoryOutput
	ToFederatedCargoRepositoryOutputWithContext(ctx context.Context) FederatedCargoRepositoryOutput
}

func (*FederatedCargoRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedCargoRepository)(nil)).Elem()
}

func (i *FederatedCargoRepository) ToFederatedCargoRepositoryOutput() FederatedCargoRepositoryOutput {
	return i.ToFederatedCargoRepositoryOutputWithContext(context.Background())
}

func (i *FederatedCargoRepository) ToFederatedCargoRepositoryOutputWithContext(ctx context.Context) FederatedCargoRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedCargoRepositoryOutput)
}

// FederatedCargoRepositoryArrayInput is an input type that accepts FederatedCargoRepositoryArray and FederatedCargoRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedCargoRepositoryArrayInput` via:
//
//	FederatedCargoRepositoryArray{ FederatedCargoRepositoryArgs{...} }
type FederatedCargoRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedCargoRepositoryArrayOutput() FederatedCargoRepositoryArrayOutput
	ToFederatedCargoRepositoryArrayOutputWithContext(context.Context) FederatedCargoRepositoryArrayOutput
}

type FederatedCargoRepositoryArray []FederatedCargoRepositoryInput

func (FederatedCargoRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedCargoRepository)(nil)).Elem()
}

func (i FederatedCargoRepositoryArray) ToFederatedCargoRepositoryArrayOutput() FederatedCargoRepositoryArrayOutput {
	return i.ToFederatedCargoRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedCargoRepositoryArray) ToFederatedCargoRepositoryArrayOutputWithContext(ctx context.Context) FederatedCargoRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedCargoRepositoryArrayOutput)
}

// FederatedCargoRepositoryMapInput is an input type that accepts FederatedCargoRepositoryMap and FederatedCargoRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedCargoRepositoryMapInput` via:
//
//	FederatedCargoRepositoryMap{ "key": FederatedCargoRepositoryArgs{...} }
type FederatedCargoRepositoryMapInput interface {
	pulumi.Input

	ToFederatedCargoRepositoryMapOutput() FederatedCargoRepositoryMapOutput
	ToFederatedCargoRepositoryMapOutputWithContext(context.Context) FederatedCargoRepositoryMapOutput
}

type FederatedCargoRepositoryMap map[string]FederatedCargoRepositoryInput

func (FederatedCargoRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedCargoRepository)(nil)).Elem()
}

func (i FederatedCargoRepositoryMap) ToFederatedCargoRepositoryMapOutput() FederatedCargoRepositoryMapOutput {
	return i.ToFederatedCargoRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedCargoRepositoryMap) ToFederatedCargoRepositoryMapOutputWithContext(ctx context.Context) FederatedCargoRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedCargoRepositoryMapOutput)
}

type FederatedCargoRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedCargoRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedCargoRepository)(nil)).Elem()
}

func (o FederatedCargoRepositoryOutput) ToFederatedCargoRepositoryOutput() FederatedCargoRepositoryOutput {
	return o
}

func (o FederatedCargoRepositoryOutput) ToFederatedCargoRepositoryOutputWithContext(ctx context.Context) FederatedCargoRepositoryOutput {
	return o
}

// Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous
// access to these resources (only), note that this will override the security anonymous access option. Default value is
// 'false'.
func (o FederatedCargoRepositoryOutput) AnonymousAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCargoRepository) pulumi.BoolPtrOutput { return v.AnonymousAccess }).(pulumi.BoolPtrOutput)
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o FederatedCargoRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCargoRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o FederatedCargoRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCargoRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o FederatedCargoRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCargoRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

// Delete all federated members on `terraform destroy` if set to `true`. Caution: it will delete all the repositories in
// the federation on other Artifactory instances.
func (o FederatedCargoRepositoryOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCargoRepository) pulumi.BoolPtrOutput { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

// Public description.
func (o FederatedCargoRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedCargoRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o FederatedCargoRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCargoRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// Enable internal index support based on Cargo sparse index specifications, instead of the default git index. Default
// value is 'false'.
func (o FederatedCargoRepositoryOutput) EnableSparseIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCargoRepository) pulumi.BoolPtrOutput { return v.EnableSparseIndex }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
// artifacts are excluded.
func (o FederatedCargoRepositoryOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedCargoRepository) pulumi.StringOutput { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o FederatedCargoRepositoryOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedCargoRepository) pulumi.StringOutput { return v.IncludesPattern }).(pulumi.StringOutput)
}

func (o FederatedCargoRepositoryOutput) IndexCompressionFormats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedCargoRepository) pulumi.StringArrayOutput { return v.IndexCompressionFormats }).(pulumi.StringArrayOutput)
}

// the identity key of the repo.
func (o FederatedCargoRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedCargoRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o FederatedCargoRepositoryOutput) Members() FederatedCargoRepositoryMemberArrayOutput {
	return o.ApplyT(func(v *FederatedCargoRepository) FederatedCargoRepositoryMemberArrayOutput { return v.Members }).(FederatedCargoRepositoryMemberArrayOutput)
}

// Internal description.
func (o FederatedCargoRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedCargoRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o FederatedCargoRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedCargoRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o FederatedCargoRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCargoRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o FederatedCargoRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedCargoRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o FederatedCargoRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedCargoRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o FederatedCargoRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedCargoRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o FederatedCargoRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedCargoRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o FederatedCargoRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCargoRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type FederatedCargoRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedCargoRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedCargoRepository)(nil)).Elem()
}

func (o FederatedCargoRepositoryArrayOutput) ToFederatedCargoRepositoryArrayOutput() FederatedCargoRepositoryArrayOutput {
	return o
}

func (o FederatedCargoRepositoryArrayOutput) ToFederatedCargoRepositoryArrayOutputWithContext(ctx context.Context) FederatedCargoRepositoryArrayOutput {
	return o
}

func (o FederatedCargoRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedCargoRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedCargoRepository {
		return vs[0].([]*FederatedCargoRepository)[vs[1].(int)]
	}).(FederatedCargoRepositoryOutput)
}

type FederatedCargoRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedCargoRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedCargoRepository)(nil)).Elem()
}

func (o FederatedCargoRepositoryMapOutput) ToFederatedCargoRepositoryMapOutput() FederatedCargoRepositoryMapOutput {
	return o
}

func (o FederatedCargoRepositoryMapOutput) ToFederatedCargoRepositoryMapOutputWithContext(ctx context.Context) FederatedCargoRepositoryMapOutput {
	return o
}

func (o FederatedCargoRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedCargoRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedCargoRepository {
		return vs[0].(map[string]*FederatedCargoRepository)[vs[1].(string)]
	}).(FederatedCargoRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedCargoRepositoryInput)(nil)).Elem(), &FederatedCargoRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedCargoRepositoryArrayInput)(nil)).Elem(), FederatedCargoRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedCargoRepositoryMapInput)(nil)).Elem(), FederatedCargoRepositoryMap{})
	pulumi.RegisterOutputType(FederatedCargoRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedCargoRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedCargoRepositoryMapOutput{})
}
