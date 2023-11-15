// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a federated Composer repository.
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
//			_, err := artifactory.NewFederatedComposerRepository(ctx, "terraform-federated-test-composer-repo", &artifactory.FederatedComposerRepositoryArgs{
//				Key: pulumi.String("terraform-federated-test-composer-repo"),
//				Members: artifactory.FederatedComposerRepositoryMemberArray{
//					&artifactory.FederatedComposerRepositoryMemberArgs{
//						Enabled: pulumi.Bool(true),
//						Url:     pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-composer-repo"),
//					},
//					&artifactory.FederatedComposerRepositoryMemberArgs{
//						Enabled: pulumi.Bool(true),
//						Url:     pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-composer-repo-2"),
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
//	$ pulumi import artifactory:index/federatedComposerRepository:FederatedComposerRepository terraform-federated-test-composer-repo terraform-federated-test-composer-repo
//
// ```
type FederatedComposerRepository struct {
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
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrOutput `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrOutput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key pulumi.StringOutput `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members FederatedComposerRepositoryMemberArrayOutput `pulumi:"members"`
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
	// Repository layout key for the federated repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrOutput `pulumi:"xrayIndex"`
}

// NewFederatedComposerRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedComposerRepository(ctx *pulumi.Context,
	name string, args *FederatedComposerRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedComposerRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FederatedComposerRepository
	err := ctx.RegisterResource("artifactory:index/federatedComposerRepository:FederatedComposerRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedComposerRepository gets an existing FederatedComposerRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedComposerRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedComposerRepositoryState, opts ...pulumi.ResourceOption) (*FederatedComposerRepository, error) {
	var resource FederatedComposerRepository
	err := ctx.ReadResource("artifactory:index/federatedComposerRepository:FederatedComposerRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedComposerRepository resources.
type federatedComposerRepositoryState struct {
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
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key *string `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members []FederatedComposerRepositoryMember `pulumi:"members"`
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
	// Repository layout key for the federated repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

type FederatedComposerRepositoryState struct {
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
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// the identity key of the repo.
	Key pulumi.StringPtrInput
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members FederatedComposerRepositoryMemberArrayInput
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
	// Repository layout key for the federated repository
	RepoLayoutRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (FederatedComposerRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedComposerRepositoryState)(nil)).Elem()
}

type federatedComposerRepositoryArgs struct {
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
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members []FederatedComposerRepositoryMember `pulumi:"members"`
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
	// Repository layout key for the federated repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a FederatedComposerRepository resource.
type FederatedComposerRepositoryArgs struct {
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
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// the identity key of the repo.
	Key pulumi.StringInput
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members FederatedComposerRepositoryMemberArrayInput
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
	// Repository layout key for the federated repository
	RepoLayoutRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (FederatedComposerRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedComposerRepositoryArgs)(nil)).Elem()
}

type FederatedComposerRepositoryInput interface {
	pulumi.Input

	ToFederatedComposerRepositoryOutput() FederatedComposerRepositoryOutput
	ToFederatedComposerRepositoryOutputWithContext(ctx context.Context) FederatedComposerRepositoryOutput
}

func (*FederatedComposerRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedComposerRepository)(nil)).Elem()
}

func (i *FederatedComposerRepository) ToFederatedComposerRepositoryOutput() FederatedComposerRepositoryOutput {
	return i.ToFederatedComposerRepositoryOutputWithContext(context.Background())
}

func (i *FederatedComposerRepository) ToFederatedComposerRepositoryOutputWithContext(ctx context.Context) FederatedComposerRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedComposerRepositoryOutput)
}

// FederatedComposerRepositoryArrayInput is an input type that accepts FederatedComposerRepositoryArray and FederatedComposerRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedComposerRepositoryArrayInput` via:
//
//	FederatedComposerRepositoryArray{ FederatedComposerRepositoryArgs{...} }
type FederatedComposerRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedComposerRepositoryArrayOutput() FederatedComposerRepositoryArrayOutput
	ToFederatedComposerRepositoryArrayOutputWithContext(context.Context) FederatedComposerRepositoryArrayOutput
}

type FederatedComposerRepositoryArray []FederatedComposerRepositoryInput

func (FederatedComposerRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedComposerRepository)(nil)).Elem()
}

func (i FederatedComposerRepositoryArray) ToFederatedComposerRepositoryArrayOutput() FederatedComposerRepositoryArrayOutput {
	return i.ToFederatedComposerRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedComposerRepositoryArray) ToFederatedComposerRepositoryArrayOutputWithContext(ctx context.Context) FederatedComposerRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedComposerRepositoryArrayOutput)
}

// FederatedComposerRepositoryMapInput is an input type that accepts FederatedComposerRepositoryMap and FederatedComposerRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedComposerRepositoryMapInput` via:
//
//	FederatedComposerRepositoryMap{ "key": FederatedComposerRepositoryArgs{...} }
type FederatedComposerRepositoryMapInput interface {
	pulumi.Input

	ToFederatedComposerRepositoryMapOutput() FederatedComposerRepositoryMapOutput
	ToFederatedComposerRepositoryMapOutputWithContext(context.Context) FederatedComposerRepositoryMapOutput
}

type FederatedComposerRepositoryMap map[string]FederatedComposerRepositoryInput

func (FederatedComposerRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedComposerRepository)(nil)).Elem()
}

func (i FederatedComposerRepositoryMap) ToFederatedComposerRepositoryMapOutput() FederatedComposerRepositoryMapOutput {
	return i.ToFederatedComposerRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedComposerRepositoryMap) ToFederatedComposerRepositoryMapOutputWithContext(ctx context.Context) FederatedComposerRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedComposerRepositoryMapOutput)
}

type FederatedComposerRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedComposerRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedComposerRepository)(nil)).Elem()
}

func (o FederatedComposerRepositoryOutput) ToFederatedComposerRepositoryOutput() FederatedComposerRepositoryOutput {
	return o
}

func (o FederatedComposerRepositoryOutput) ToFederatedComposerRepositoryOutputWithContext(ctx context.Context) FederatedComposerRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o FederatedComposerRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedComposerRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o FederatedComposerRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedComposerRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o FederatedComposerRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedComposerRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

// Delete all federated members on `terraform destroy` if set to `true`. Caution: it will delete all the repositories in
// the federation on other Artifactory instances.
func (o FederatedComposerRepositoryOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedComposerRepository) pulumi.BoolPtrOutput { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

// Public description.
func (o FederatedComposerRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedComposerRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o FederatedComposerRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedComposerRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o FederatedComposerRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedComposerRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o FederatedComposerRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedComposerRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// the identity key of the repo.
func (o FederatedComposerRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedComposerRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o FederatedComposerRepositoryOutput) Members() FederatedComposerRepositoryMemberArrayOutput {
	return o.ApplyT(func(v *FederatedComposerRepository) FederatedComposerRepositoryMemberArrayOutput { return v.Members }).(FederatedComposerRepositoryMemberArrayOutput)
}

// Internal description.
func (o FederatedComposerRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedComposerRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o FederatedComposerRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedComposerRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o FederatedComposerRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedComposerRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o FederatedComposerRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedComposerRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o FederatedComposerRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedComposerRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o FederatedComposerRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedComposerRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the federated repository
func (o FederatedComposerRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedComposerRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o FederatedComposerRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedComposerRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type FederatedComposerRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedComposerRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedComposerRepository)(nil)).Elem()
}

func (o FederatedComposerRepositoryArrayOutput) ToFederatedComposerRepositoryArrayOutput() FederatedComposerRepositoryArrayOutput {
	return o
}

func (o FederatedComposerRepositoryArrayOutput) ToFederatedComposerRepositoryArrayOutputWithContext(ctx context.Context) FederatedComposerRepositoryArrayOutput {
	return o
}

func (o FederatedComposerRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedComposerRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedComposerRepository {
		return vs[0].([]*FederatedComposerRepository)[vs[1].(int)]
	}).(FederatedComposerRepositoryOutput)
}

type FederatedComposerRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedComposerRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedComposerRepository)(nil)).Elem()
}

func (o FederatedComposerRepositoryMapOutput) ToFederatedComposerRepositoryMapOutput() FederatedComposerRepositoryMapOutput {
	return o
}

func (o FederatedComposerRepositoryMapOutput) ToFederatedComposerRepositoryMapOutputWithContext(ctx context.Context) FederatedComposerRepositoryMapOutput {
	return o
}

func (o FederatedComposerRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedComposerRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedComposerRepository {
		return vs[0].(map[string]*FederatedComposerRepository)[vs[1].(string)]
	}).(FederatedComposerRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedComposerRepositoryInput)(nil)).Elem(), &FederatedComposerRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedComposerRepositoryArrayInput)(nil)).Elem(), FederatedComposerRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedComposerRepositoryMapInput)(nil)).Elem(), FederatedComposerRepositoryMap{})
	pulumi.RegisterOutputType(FederatedComposerRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedComposerRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedComposerRepositoryMapOutput{})
}
