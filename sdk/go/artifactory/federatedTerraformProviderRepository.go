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
//			_, err := artifactory.NewFederatedTerraformProviderRepository(ctx, "terraform-federated-test-terraform_provider-repo", &artifactory.FederatedTerraformProviderRepositoryArgs{
//				Key: pulumi.String("terraform-federated-test-terraform-provider-repo"),
//				Members: artifactory.FederatedTerraformProviderRepositoryMemberArray{
//					&artifactory.FederatedTerraformProviderRepositoryMemberArgs{
//						Url:     pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-terraform_provider-repo"),
//						Enabled: pulumi.Bool(true),
//					},
//					&artifactory.FederatedTerraformProviderRepositoryMemberArgs{
//						Url:     pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-terraform_provider-repo-2"),
//						Enabled: pulumi.Bool(true),
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
// $ pulumi import artifactory:index/federatedTerraformProviderRepository:FederatedTerraformProviderRepository terraform-federated-test-terraform_provider-repo terraform-federated-test-terraform-provider-repo
// ```
type FederatedTerraformProviderRepository struct {
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
	Members FederatedTerraformProviderRepositoryMemberArrayOutput `pulumi:"members"`
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
	// Proxy key from Artifactory Proxies settings. Default is empty field. Can't be set if `disableProxy = true`.
	Proxy pulumi.StringPtrOutput `pulumi:"proxy"`
	// Repository layout key for the federated repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrOutput `pulumi:"xrayIndex"`
}

// NewFederatedTerraformProviderRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedTerraformProviderRepository(ctx *pulumi.Context,
	name string, args *FederatedTerraformProviderRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedTerraformProviderRepository, error) {
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
	var resource FederatedTerraformProviderRepository
	err := ctx.RegisterResource("artifactory:index/federatedTerraformProviderRepository:FederatedTerraformProviderRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedTerraformProviderRepository gets an existing FederatedTerraformProviderRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedTerraformProviderRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedTerraformProviderRepositoryState, opts ...pulumi.ResourceOption) (*FederatedTerraformProviderRepository, error) {
	var resource FederatedTerraformProviderRepository
	err := ctx.ReadResource("artifactory:index/federatedTerraformProviderRepository:FederatedTerraformProviderRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedTerraformProviderRepository resources.
type federatedTerraformProviderRepositoryState struct {
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
	Members []FederatedTerraformProviderRepositoryMember `pulumi:"members"`
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
	// Proxy key from Artifactory Proxies settings. Default is empty field. Can't be set if `disableProxy = true`.
	Proxy *string `pulumi:"proxy"`
	// Repository layout key for the federated repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

type FederatedTerraformProviderRepositoryState struct {
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
	Members FederatedTerraformProviderRepositoryMemberArrayInput
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
	// Proxy key from Artifactory Proxies settings. Default is empty field. Can't be set if `disableProxy = true`.
	Proxy pulumi.StringPtrInput
	// Repository layout key for the federated repository
	RepoLayoutRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (FederatedTerraformProviderRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedTerraformProviderRepositoryState)(nil)).Elem()
}

type federatedTerraformProviderRepositoryArgs struct {
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
	Members []FederatedTerraformProviderRepositoryMember `pulumi:"members"`
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
	// Proxy key from Artifactory Proxies settings. Default is empty field. Can't be set if `disableProxy = true`.
	Proxy *string `pulumi:"proxy"`
	// Repository layout key for the federated repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a FederatedTerraformProviderRepository resource.
type FederatedTerraformProviderRepositoryArgs struct {
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
	Members FederatedTerraformProviderRepositoryMemberArrayInput
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
	// Proxy key from Artifactory Proxies settings. Default is empty field. Can't be set if `disableProxy = true`.
	Proxy pulumi.StringPtrInput
	// Repository layout key for the federated repository
	RepoLayoutRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (FederatedTerraformProviderRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedTerraformProviderRepositoryArgs)(nil)).Elem()
}

type FederatedTerraformProviderRepositoryInput interface {
	pulumi.Input

	ToFederatedTerraformProviderRepositoryOutput() FederatedTerraformProviderRepositoryOutput
	ToFederatedTerraformProviderRepositoryOutputWithContext(ctx context.Context) FederatedTerraformProviderRepositoryOutput
}

func (*FederatedTerraformProviderRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedTerraformProviderRepository)(nil)).Elem()
}

func (i *FederatedTerraformProviderRepository) ToFederatedTerraformProviderRepositoryOutput() FederatedTerraformProviderRepositoryOutput {
	return i.ToFederatedTerraformProviderRepositoryOutputWithContext(context.Background())
}

func (i *FederatedTerraformProviderRepository) ToFederatedTerraformProviderRepositoryOutputWithContext(ctx context.Context) FederatedTerraformProviderRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedTerraformProviderRepositoryOutput)
}

// FederatedTerraformProviderRepositoryArrayInput is an input type that accepts FederatedTerraformProviderRepositoryArray and FederatedTerraformProviderRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedTerraformProviderRepositoryArrayInput` via:
//
//	FederatedTerraformProviderRepositoryArray{ FederatedTerraformProviderRepositoryArgs{...} }
type FederatedTerraformProviderRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedTerraformProviderRepositoryArrayOutput() FederatedTerraformProviderRepositoryArrayOutput
	ToFederatedTerraformProviderRepositoryArrayOutputWithContext(context.Context) FederatedTerraformProviderRepositoryArrayOutput
}

type FederatedTerraformProviderRepositoryArray []FederatedTerraformProviderRepositoryInput

func (FederatedTerraformProviderRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedTerraformProviderRepository)(nil)).Elem()
}

func (i FederatedTerraformProviderRepositoryArray) ToFederatedTerraformProviderRepositoryArrayOutput() FederatedTerraformProviderRepositoryArrayOutput {
	return i.ToFederatedTerraformProviderRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedTerraformProviderRepositoryArray) ToFederatedTerraformProviderRepositoryArrayOutputWithContext(ctx context.Context) FederatedTerraformProviderRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedTerraformProviderRepositoryArrayOutput)
}

// FederatedTerraformProviderRepositoryMapInput is an input type that accepts FederatedTerraformProviderRepositoryMap and FederatedTerraformProviderRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedTerraformProviderRepositoryMapInput` via:
//
//	FederatedTerraformProviderRepositoryMap{ "key": FederatedTerraformProviderRepositoryArgs{...} }
type FederatedTerraformProviderRepositoryMapInput interface {
	pulumi.Input

	ToFederatedTerraformProviderRepositoryMapOutput() FederatedTerraformProviderRepositoryMapOutput
	ToFederatedTerraformProviderRepositoryMapOutputWithContext(context.Context) FederatedTerraformProviderRepositoryMapOutput
}

type FederatedTerraformProviderRepositoryMap map[string]FederatedTerraformProviderRepositoryInput

func (FederatedTerraformProviderRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedTerraformProviderRepository)(nil)).Elem()
}

func (i FederatedTerraformProviderRepositoryMap) ToFederatedTerraformProviderRepositoryMapOutput() FederatedTerraformProviderRepositoryMapOutput {
	return i.ToFederatedTerraformProviderRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedTerraformProviderRepositoryMap) ToFederatedTerraformProviderRepositoryMapOutputWithContext(ctx context.Context) FederatedTerraformProviderRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedTerraformProviderRepositoryMapOutput)
}

type FederatedTerraformProviderRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedTerraformProviderRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedTerraformProviderRepository)(nil)).Elem()
}

func (o FederatedTerraformProviderRepositoryOutput) ToFederatedTerraformProviderRepositoryOutput() FederatedTerraformProviderRepositoryOutput {
	return o
}

func (o FederatedTerraformProviderRepositoryOutput) ToFederatedTerraformProviderRepositoryOutputWithContext(ctx context.Context) FederatedTerraformProviderRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o FederatedTerraformProviderRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformProviderRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o FederatedTerraformProviderRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformProviderRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o FederatedTerraformProviderRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformProviderRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

// Delete all federated members on `terraform destroy` if set to `true`. Caution: it will delete all the repositories in
// the federation on other Artifactory instances.
func (o FederatedTerraformProviderRepositoryOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformProviderRepository) pulumi.BoolPtrOutput { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

// Public description.
func (o FederatedTerraformProviderRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformProviderRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
func (o FederatedTerraformProviderRepositoryOutput) DisableProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformProviderRepository) pulumi.BoolPtrOutput { return v.DisableProxy }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o FederatedTerraformProviderRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformProviderRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o FederatedTerraformProviderRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformProviderRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o FederatedTerraformProviderRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformProviderRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// the identity key of the repo.
func (o FederatedTerraformProviderRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedTerraformProviderRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o FederatedTerraformProviderRepositoryOutput) Members() FederatedTerraformProviderRepositoryMemberArrayOutput {
	return o.ApplyT(func(v *FederatedTerraformProviderRepository) FederatedTerraformProviderRepositoryMemberArrayOutput {
		return v.Members
	}).(FederatedTerraformProviderRepositoryMemberArrayOutput)
}

// Internal description.
func (o FederatedTerraformProviderRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformProviderRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o FederatedTerraformProviderRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedTerraformProviderRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o FederatedTerraformProviderRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformProviderRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o FederatedTerraformProviderRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedTerraformProviderRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o FederatedTerraformProviderRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformProviderRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o FederatedTerraformProviderRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedTerraformProviderRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Proxy key from Artifactory Proxies settings. Default is empty field. Can't be set if `disableProxy = true`.
func (o FederatedTerraformProviderRepositoryOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformProviderRepository) pulumi.StringPtrOutput { return v.Proxy }).(pulumi.StringPtrOutput)
}

// Repository layout key for the federated repository
func (o FederatedTerraformProviderRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformProviderRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o FederatedTerraformProviderRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedTerraformProviderRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type FederatedTerraformProviderRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedTerraformProviderRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedTerraformProviderRepository)(nil)).Elem()
}

func (o FederatedTerraformProviderRepositoryArrayOutput) ToFederatedTerraformProviderRepositoryArrayOutput() FederatedTerraformProviderRepositoryArrayOutput {
	return o
}

func (o FederatedTerraformProviderRepositoryArrayOutput) ToFederatedTerraformProviderRepositoryArrayOutputWithContext(ctx context.Context) FederatedTerraformProviderRepositoryArrayOutput {
	return o
}

func (o FederatedTerraformProviderRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedTerraformProviderRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedTerraformProviderRepository {
		return vs[0].([]*FederatedTerraformProviderRepository)[vs[1].(int)]
	}).(FederatedTerraformProviderRepositoryOutput)
}

type FederatedTerraformProviderRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedTerraformProviderRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedTerraformProviderRepository)(nil)).Elem()
}

func (o FederatedTerraformProviderRepositoryMapOutput) ToFederatedTerraformProviderRepositoryMapOutput() FederatedTerraformProviderRepositoryMapOutput {
	return o
}

func (o FederatedTerraformProviderRepositoryMapOutput) ToFederatedTerraformProviderRepositoryMapOutputWithContext(ctx context.Context) FederatedTerraformProviderRepositoryMapOutput {
	return o
}

func (o FederatedTerraformProviderRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedTerraformProviderRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedTerraformProviderRepository {
		return vs[0].(map[string]*FederatedTerraformProviderRepository)[vs[1].(string)]
	}).(FederatedTerraformProviderRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedTerraformProviderRepositoryInput)(nil)).Elem(), &FederatedTerraformProviderRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedTerraformProviderRepositoryArrayInput)(nil)).Elem(), FederatedTerraformProviderRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedTerraformProviderRepositoryMapInput)(nil)).Elem(), FederatedTerraformProviderRepositoryMap{})
	pulumi.RegisterOutputType(FederatedTerraformProviderRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedTerraformProviderRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedTerraformProviderRepositoryMapOutput{})
}
