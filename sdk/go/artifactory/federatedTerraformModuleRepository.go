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

// Creates a federated Generic repository.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := artifactory.NewFederatedTerraformModuleRepository(ctx, "terraform-federated-test-terraformModule-repo", &artifactory.FederatedTerraformModuleRepositoryArgs{
//				Key: pulumi.String("terraform-federated-test-terraform-module-repo"),
//				Members: artifactory.FederatedTerraformModuleRepositoryMemberArray{
//					&artifactory.FederatedTerraformModuleRepositoryMemberArgs{
//						Enabled: pulumi.Bool(true),
//						Url:     pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-terraform_module-repo"),
//					},
//					&artifactory.FederatedTerraformModuleRepositoryMemberArgs{
//						Enabled: pulumi.Bool(true),
//						Url:     pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-terraform_module-repo-2"),
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
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Federated repositories can be imported using their name, e.g.
//
// ```sh
// $ pulumi import artifactory:index/federatedTerraformModuleRepository:FederatedTerraformModuleRepository terraform-federated-test-terraform_module-repo terraform-federated-test-terraform-module-repo
// ```
type FederatedTerraformModuleRepository struct {
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
	// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
	DisableProxy pulumi.BoolPtrOutput `pulumi:"disableProxy"`
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
	Members FederatedTerraformModuleRepositoryMemberArrayOutput `pulumi:"members"`
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
	// Proxy key from Artifactory Proxies settings. Default is empty field. Can't be set if `disableProxy = true`.
	Proxy pulumi.StringPtrOutput `pulumi:"proxy"`
	// Repository layout key for the federated repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrOutput `pulumi:"xrayIndex"`
}

// NewFederatedTerraformModuleRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedTerraformModuleRepository(ctx *pulumi.Context,
	name string, args *FederatedTerraformModuleRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedTerraformModuleRepository, error) {
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
	var resource FederatedTerraformModuleRepository
	err := ctx.RegisterResource("artifactory:index/federatedTerraformModuleRepository:FederatedTerraformModuleRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedTerraformModuleRepository gets an existing FederatedTerraformModuleRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedTerraformModuleRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedTerraformModuleRepositoryState, opts ...pulumi.ResourceOption) (*FederatedTerraformModuleRepository, error) {
	var resource FederatedTerraformModuleRepository
	err := ctx.ReadResource("artifactory:index/federatedTerraformModuleRepository:FederatedTerraformModuleRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedTerraformModuleRepository resources.
type federatedTerraformModuleRepositoryState struct {
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
	// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
	DisableProxy *bool `pulumi:"disableProxy"`
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
	Members []FederatedTerraformModuleRepositoryMember `pulumi:"members"`
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
	// Proxy key from Artifactory Proxies settings. Default is empty field. Can't be set if `disableProxy = true`.
	Proxy *string `pulumi:"proxy"`
	// Repository layout key for the federated repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

type FederatedTerraformModuleRepositoryState struct {
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
	// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
	DisableProxy pulumi.BoolPtrInput
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
	Members FederatedTerraformModuleRepositoryMemberArrayInput
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
	// Proxy key from Artifactory Proxies settings. Default is empty field. Can't be set if `disableProxy = true`.
	Proxy pulumi.StringPtrInput
	// Repository layout key for the federated repository
	RepoLayoutRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (FederatedTerraformModuleRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedTerraformModuleRepositoryState)(nil)).Elem()
}

type federatedTerraformModuleRepositoryArgs struct {
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
	// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
	DisableProxy *bool `pulumi:"disableProxy"`
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
	Members []FederatedTerraformModuleRepositoryMember `pulumi:"members"`
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
	// Proxy key from Artifactory Proxies settings. Default is empty field. Can't be set if `disableProxy = true`.
	Proxy *string `pulumi:"proxy"`
	// Repository layout key for the federated repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a FederatedTerraformModuleRepository resource.
type FederatedTerraformModuleRepositoryArgs struct {
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
	// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
	DisableProxy pulumi.BoolPtrInput
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
	Members FederatedTerraformModuleRepositoryMemberArrayInput
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
	// Proxy key from Artifactory Proxies settings. Default is empty field. Can't be set if `disableProxy = true`.
	Proxy pulumi.StringPtrInput
	// Repository layout key for the federated repository
	RepoLayoutRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (FederatedTerraformModuleRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedTerraformModuleRepositoryArgs)(nil)).Elem()
}

type FederatedTerraformModuleRepositoryInput interface {
	pulumi.Input

	ToFederatedTerraformModuleRepositoryOutput() FederatedTerraformModuleRepositoryOutput
	ToFederatedTerraformModuleRepositoryOutputWithContext(ctx context.Context) FederatedTerraformModuleRepositoryOutput
}

func (*FederatedTerraformModuleRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedTerraformModuleRepository)(nil)).Elem()
}

func (i *FederatedTerraformModuleRepository) ToFederatedTerraformModuleRepositoryOutput() FederatedTerraformModuleRepositoryOutput {
	return i.ToFederatedTerraformModuleRepositoryOutputWithContext(context.Background())
}

func (i *FederatedTerraformModuleRepository) ToFederatedTerraformModuleRepositoryOutputWithContext(ctx context.Context) FederatedTerraformModuleRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedTerraformModuleRepositoryOutput)
}

// FederatedTerraformModuleRepositoryArrayInput is an input type that accepts FederatedTerraformModuleRepositoryArray and FederatedTerraformModuleRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedTerraformModuleRepositoryArrayInput` via:
//
//	FederatedTerraformModuleRepositoryArray{ FederatedTerraformModuleRepositoryArgs{...} }
type FederatedTerraformModuleRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedTerraformModuleRepositoryArrayOutput() FederatedTerraformModuleRepositoryArrayOutput
	ToFederatedTerraformModuleRepositoryArrayOutputWithContext(context.Context) FederatedTerraformModuleRepositoryArrayOutput
}

type FederatedTerraformModuleRepositoryArray []FederatedTerraformModuleRepositoryInput

func (FederatedTerraformModuleRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedTerraformModuleRepository)(nil)).Elem()
}

func (i FederatedTerraformModuleRepositoryArray) ToFederatedTerraformModuleRepositoryArrayOutput() FederatedTerraformModuleRepositoryArrayOutput {
	return i.ToFederatedTerraformModuleRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedTerraformModuleRepositoryArray) ToFederatedTerraformModuleRepositoryArrayOutputWithContext(ctx context.Context) FederatedTerraformModuleRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedTerraformModuleRepositoryArrayOutput)
}

// FederatedTerraformModuleRepositoryMapInput is an input type that accepts FederatedTerraformModuleRepositoryMap and FederatedTerraformModuleRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedTerraformModuleRepositoryMapInput` via:
//
//	FederatedTerraformModuleRepositoryMap{ "key": FederatedTerraformModuleRepositoryArgs{...} }
type FederatedTerraformModuleRepositoryMapInput interface {
	pulumi.Input

	ToFederatedTerraformModuleRepositoryMapOutput() FederatedTerraformModuleRepositoryMapOutput
	ToFederatedTerraformModuleRepositoryMapOutputWithContext(context.Context) FederatedTerraformModuleRepositoryMapOutput
}

type FederatedTerraformModuleRepositoryMap map[string]FederatedTerraformModuleRepositoryInput

func (FederatedTerraformModuleRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedTerraformModuleRepository)(nil)).Elem()
}

func (i FederatedTerraformModuleRepositoryMap) ToFederatedTerraformModuleRepositoryMapOutput() FederatedTerraformModuleRepositoryMapOutput {
	return i.ToFederatedTerraformModuleRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedTerraformModuleRepositoryMap) ToFederatedTerraformModuleRepositoryMapOutputWithContext(ctx context.Context) FederatedTerraformModuleRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedTerraformModuleRepositoryMapOutput)
}

type FederatedTerraformModuleRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedTerraformModuleRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedTerraformModuleRepository)(nil)).Elem()
}

func (o FederatedTerraformModuleRepositoryOutput) ToFederatedTerraformModuleRepositoryOutput() FederatedTerraformModuleRepositoryOutput {
	return o
}

func (o FederatedTerraformModuleRepositoryOutput) ToFederatedTerraformModuleRepositoryOutputWithContext(ctx context.Context) FederatedTerraformModuleRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o FederatedTerraformModuleRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformModuleRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o FederatedTerraformModuleRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformModuleRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o FederatedTerraformModuleRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformModuleRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

// Delete all federated members on `terraform destroy` if set to `true`. Caution: it will delete all the repositories in
// the federation on other Artifactory instances.
func (o FederatedTerraformModuleRepositoryOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformModuleRepository) pulumi.BoolPtrOutput { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

// Public description.
func (o FederatedTerraformModuleRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformModuleRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
func (o FederatedTerraformModuleRepositoryOutput) DisableProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformModuleRepository) pulumi.BoolPtrOutput { return v.DisableProxy }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o FederatedTerraformModuleRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformModuleRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o FederatedTerraformModuleRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformModuleRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o FederatedTerraformModuleRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformModuleRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// the identity key of the repo.
func (o FederatedTerraformModuleRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedTerraformModuleRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o FederatedTerraformModuleRepositoryOutput) Members() FederatedTerraformModuleRepositoryMemberArrayOutput {
	return o.ApplyT(func(v *FederatedTerraformModuleRepository) FederatedTerraformModuleRepositoryMemberArrayOutput {
		return v.Members
	}).(FederatedTerraformModuleRepositoryMemberArrayOutput)
}

// Internal description.
func (o FederatedTerraformModuleRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformModuleRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o FederatedTerraformModuleRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedTerraformModuleRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o FederatedTerraformModuleRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformModuleRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o FederatedTerraformModuleRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedTerraformModuleRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o FederatedTerraformModuleRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformModuleRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o FederatedTerraformModuleRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedTerraformModuleRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Proxy key from Artifactory Proxies settings. Default is empty field. Can't be set if `disableProxy = true`.
func (o FederatedTerraformModuleRepositoryOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformModuleRepository) pulumi.StringPtrOutput { return v.Proxy }).(pulumi.StringPtrOutput)
}

// Repository layout key for the federated repository
func (o FederatedTerraformModuleRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformModuleRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o FederatedTerraformModuleRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformModuleRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type FederatedTerraformModuleRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedTerraformModuleRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedTerraformModuleRepository)(nil)).Elem()
}

func (o FederatedTerraformModuleRepositoryArrayOutput) ToFederatedTerraformModuleRepositoryArrayOutput() FederatedTerraformModuleRepositoryArrayOutput {
	return o
}

func (o FederatedTerraformModuleRepositoryArrayOutput) ToFederatedTerraformModuleRepositoryArrayOutputWithContext(ctx context.Context) FederatedTerraformModuleRepositoryArrayOutput {
	return o
}

func (o FederatedTerraformModuleRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedTerraformModuleRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedTerraformModuleRepository {
		return vs[0].([]*FederatedTerraformModuleRepository)[vs[1].(int)]
	}).(FederatedTerraformModuleRepositoryOutput)
}

type FederatedTerraformModuleRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedTerraformModuleRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedTerraformModuleRepository)(nil)).Elem()
}

func (o FederatedTerraformModuleRepositoryMapOutput) ToFederatedTerraformModuleRepositoryMapOutput() FederatedTerraformModuleRepositoryMapOutput {
	return o
}

func (o FederatedTerraformModuleRepositoryMapOutput) ToFederatedTerraformModuleRepositoryMapOutputWithContext(ctx context.Context) FederatedTerraformModuleRepositoryMapOutput {
	return o
}

func (o FederatedTerraformModuleRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedTerraformModuleRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedTerraformModuleRepository {
		return vs[0].(map[string]*FederatedTerraformModuleRepository)[vs[1].(string)]
	}).(FederatedTerraformModuleRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedTerraformModuleRepositoryInput)(nil)).Elem(), &FederatedTerraformModuleRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedTerraformModuleRepositoryArrayInput)(nil)).Elem(), FederatedTerraformModuleRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedTerraformModuleRepositoryMapInput)(nil)).Elem(), FederatedTerraformModuleRepositoryMap{})
	pulumi.RegisterOutputType(FederatedTerraformModuleRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedTerraformModuleRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedTerraformModuleRepositoryMapOutput{})
}
