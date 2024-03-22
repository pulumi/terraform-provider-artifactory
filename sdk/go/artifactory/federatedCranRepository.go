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

// Creates a federated Cran repository.
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
//			_, err := artifactory.NewFederatedCranRepository(ctx, "terraform-federated-test-cran-repo", &artifactory.FederatedCranRepositoryArgs{
//				Key: pulumi.String("terraform-federated-test-cran-repo"),
//				Members: artifactory.FederatedCranRepositoryMemberArray{
//					&artifactory.FederatedCranRepositoryMemberArgs{
//						Enabled: pulumi.Bool(true),
//						Url:     pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-cran-repo"),
//					},
//					&artifactory.FederatedCranRepositoryMemberArgs{
//						Enabled: pulumi.Bool(true),
//						Url:     pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-cran-repo-2"),
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
// $ pulumi import artifactory:index/federatedCranRepository:FederatedCranRepository terraform-federated-test-cran-repo terraform-federated-test-cran-repo
// ```
type FederatedCranRepository struct {
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
	Members FederatedCranRepositoryMemberArrayOutput `pulumi:"members"`
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

// NewFederatedCranRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedCranRepository(ctx *pulumi.Context,
	name string, args *FederatedCranRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedCranRepository, error) {
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
	var resource FederatedCranRepository
	err := ctx.RegisterResource("artifactory:index/federatedCranRepository:FederatedCranRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedCranRepository gets an existing FederatedCranRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedCranRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedCranRepositoryState, opts ...pulumi.ResourceOption) (*FederatedCranRepository, error) {
	var resource FederatedCranRepository
	err := ctx.ReadResource("artifactory:index/federatedCranRepository:FederatedCranRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedCranRepository resources.
type federatedCranRepositoryState struct {
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
	Members []FederatedCranRepositoryMember `pulumi:"members"`
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

type FederatedCranRepositoryState struct {
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
	Members FederatedCranRepositoryMemberArrayInput
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

func (FederatedCranRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedCranRepositoryState)(nil)).Elem()
}

type federatedCranRepositoryArgs struct {
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
	Members []FederatedCranRepositoryMember `pulumi:"members"`
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

// The set of arguments for constructing a FederatedCranRepository resource.
type FederatedCranRepositoryArgs struct {
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
	Members FederatedCranRepositoryMemberArrayInput
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

func (FederatedCranRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedCranRepositoryArgs)(nil)).Elem()
}

type FederatedCranRepositoryInput interface {
	pulumi.Input

	ToFederatedCranRepositoryOutput() FederatedCranRepositoryOutput
	ToFederatedCranRepositoryOutputWithContext(ctx context.Context) FederatedCranRepositoryOutput
}

func (*FederatedCranRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedCranRepository)(nil)).Elem()
}

func (i *FederatedCranRepository) ToFederatedCranRepositoryOutput() FederatedCranRepositoryOutput {
	return i.ToFederatedCranRepositoryOutputWithContext(context.Background())
}

func (i *FederatedCranRepository) ToFederatedCranRepositoryOutputWithContext(ctx context.Context) FederatedCranRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedCranRepositoryOutput)
}

// FederatedCranRepositoryArrayInput is an input type that accepts FederatedCranRepositoryArray and FederatedCranRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedCranRepositoryArrayInput` via:
//
//	FederatedCranRepositoryArray{ FederatedCranRepositoryArgs{...} }
type FederatedCranRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedCranRepositoryArrayOutput() FederatedCranRepositoryArrayOutput
	ToFederatedCranRepositoryArrayOutputWithContext(context.Context) FederatedCranRepositoryArrayOutput
}

type FederatedCranRepositoryArray []FederatedCranRepositoryInput

func (FederatedCranRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedCranRepository)(nil)).Elem()
}

func (i FederatedCranRepositoryArray) ToFederatedCranRepositoryArrayOutput() FederatedCranRepositoryArrayOutput {
	return i.ToFederatedCranRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedCranRepositoryArray) ToFederatedCranRepositoryArrayOutputWithContext(ctx context.Context) FederatedCranRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedCranRepositoryArrayOutput)
}

// FederatedCranRepositoryMapInput is an input type that accepts FederatedCranRepositoryMap and FederatedCranRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedCranRepositoryMapInput` via:
//
//	FederatedCranRepositoryMap{ "key": FederatedCranRepositoryArgs{...} }
type FederatedCranRepositoryMapInput interface {
	pulumi.Input

	ToFederatedCranRepositoryMapOutput() FederatedCranRepositoryMapOutput
	ToFederatedCranRepositoryMapOutputWithContext(context.Context) FederatedCranRepositoryMapOutput
}

type FederatedCranRepositoryMap map[string]FederatedCranRepositoryInput

func (FederatedCranRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedCranRepository)(nil)).Elem()
}

func (i FederatedCranRepositoryMap) ToFederatedCranRepositoryMapOutput() FederatedCranRepositoryMapOutput {
	return i.ToFederatedCranRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedCranRepositoryMap) ToFederatedCranRepositoryMapOutputWithContext(ctx context.Context) FederatedCranRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedCranRepositoryMapOutput)
}

type FederatedCranRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedCranRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedCranRepository)(nil)).Elem()
}

func (o FederatedCranRepositoryOutput) ToFederatedCranRepositoryOutput() FederatedCranRepositoryOutput {
	return o
}

func (o FederatedCranRepositoryOutput) ToFederatedCranRepositoryOutputWithContext(ctx context.Context) FederatedCranRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o FederatedCranRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCranRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o FederatedCranRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCranRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o FederatedCranRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCranRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

// Delete all federated members on `terraform destroy` if set to `true`. Caution: it will delete all the repositories in
// the federation on other Artifactory instances.
func (o FederatedCranRepositoryOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCranRepository) pulumi.BoolPtrOutput { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

// Public description.
func (o FederatedCranRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedCranRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
func (o FederatedCranRepositoryOutput) DisableProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCranRepository) pulumi.BoolPtrOutput { return v.DisableProxy }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o FederatedCranRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCranRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o FederatedCranRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedCranRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o FederatedCranRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedCranRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// the identity key of the repo.
func (o FederatedCranRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedCranRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o FederatedCranRepositoryOutput) Members() FederatedCranRepositoryMemberArrayOutput {
	return o.ApplyT(func(v *FederatedCranRepository) FederatedCranRepositoryMemberArrayOutput { return v.Members }).(FederatedCranRepositoryMemberArrayOutput)
}

// Internal description.
func (o FederatedCranRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedCranRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o FederatedCranRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedCranRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o FederatedCranRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCranRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o FederatedCranRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedCranRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o FederatedCranRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedCranRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o FederatedCranRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedCranRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Proxy key from Artifactory Proxies settings. Default is empty field. Can't be set if `disableProxy = true`.
func (o FederatedCranRepositoryOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedCranRepository) pulumi.StringPtrOutput { return v.Proxy }).(pulumi.StringPtrOutput)
}

// Repository layout key for the federated repository
func (o FederatedCranRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedCranRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o FederatedCranRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCranRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type FederatedCranRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedCranRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedCranRepository)(nil)).Elem()
}

func (o FederatedCranRepositoryArrayOutput) ToFederatedCranRepositoryArrayOutput() FederatedCranRepositoryArrayOutput {
	return o
}

func (o FederatedCranRepositoryArrayOutput) ToFederatedCranRepositoryArrayOutputWithContext(ctx context.Context) FederatedCranRepositoryArrayOutput {
	return o
}

func (o FederatedCranRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedCranRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedCranRepository {
		return vs[0].([]*FederatedCranRepository)[vs[1].(int)]
	}).(FederatedCranRepositoryOutput)
}

type FederatedCranRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedCranRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedCranRepository)(nil)).Elem()
}

func (o FederatedCranRepositoryMapOutput) ToFederatedCranRepositoryMapOutput() FederatedCranRepositoryMapOutput {
	return o
}

func (o FederatedCranRepositoryMapOutput) ToFederatedCranRepositoryMapOutputWithContext(ctx context.Context) FederatedCranRepositoryMapOutput {
	return o
}

func (o FederatedCranRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedCranRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedCranRepository {
		return vs[0].(map[string]*FederatedCranRepository)[vs[1].(string)]
	}).(FederatedCranRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedCranRepositoryInput)(nil)).Elem(), &FederatedCranRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedCranRepositoryArrayInput)(nil)).Elem(), FederatedCranRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedCranRepositoryMapInput)(nil)).Elem(), FederatedCranRepositoryMap{})
	pulumi.RegisterOutputType(FederatedCranRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedCranRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedCranRepositoryMapOutput{})
}
