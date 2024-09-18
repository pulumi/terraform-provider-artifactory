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

// Creates a federated Debian repository.
//
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
//			_, err := artifactory.NewFederatedDebianRepository(ctx, "terraform-federated-test-debian-repo", &artifactory.FederatedDebianRepositoryArgs{
//				Key: pulumi.String("terraform-federated-test-debian-repo"),
//				Members: artifactory.FederatedDebianRepositoryMemberArray{
//					&artifactory.FederatedDebianRepositoryMemberArgs{
//						Url:     pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-debian-repo"),
//						Enabled: pulumi.Bool(true),
//					},
//					&artifactory.FederatedDebianRepositoryMemberArgs{
//						Url:     pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-debian-repo-2"),
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
// $ pulumi import artifactory:index/federatedDebianRepository:FederatedDebianRepository terraform-federated-test-debian-repo terraform-federated-test-debian-repo
// ```
type FederatedDebianRepository struct {
	pulumi.CustomResourceState

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrOutput `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrOutput `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect     pulumi.BoolPtrOutput `pulumi:"cdnRedirect"`
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
	IncludesPattern         pulumi.StringPtrOutput   `pulumi:"includesPattern"`
	IndexCompressionFormats pulumi.StringArrayOutput `pulumi:"indexCompressionFormats"`
	// the identity key of the repo.
	Key pulumi.StringOutput `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members FederatedDebianRepositoryMemberArrayOutput `pulumi:"members"`
	// Internal description.
	Notes       pulumi.StringPtrOutput `pulumi:"notes"`
	PackageType pulumi.StringOutput    `pulumi:"packageType"`
	// Primary keypair used to sign artifacts. Default value is empty.
	PrimaryKeypairRef pulumi.StringPtrOutput `pulumi:"primaryKeypairRef"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  pulumi.BoolPtrOutput     `pulumi:"priorityResolution"`
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
	// Secondary keypair used to sign artifacts.
	SecondaryKeypairRef pulumi.StringPtrOutput `pulumi:"secondaryKeypairRef"`
	// When set, the repository will use the deprecated trivial layout.
	//
	// Deprecated: You shouldn't be using this
	TrivialLayout pulumi.BoolPtrOutput `pulumi:"trivialLayout"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrOutput `pulumi:"xrayIndex"`
}

// NewFederatedDebianRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedDebianRepository(ctx *pulumi.Context,
	name string, args *FederatedDebianRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedDebianRepository, error) {
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
	var resource FederatedDebianRepository
	err := ctx.RegisterResource("artifactory:index/federatedDebianRepository:FederatedDebianRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedDebianRepository gets an existing FederatedDebianRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedDebianRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedDebianRepositoryState, opts ...pulumi.ResourceOption) (*FederatedDebianRepository, error) {
	var resource FederatedDebianRepository
	err := ctx.ReadResource("artifactory:index/federatedDebianRepository:FederatedDebianRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedDebianRepository resources.
type federatedDebianRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect     *bool `pulumi:"cdnRedirect"`
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
	IncludesPattern         *string  `pulumi:"includesPattern"`
	IndexCompressionFormats []string `pulumi:"indexCompressionFormats"`
	// the identity key of the repo.
	Key *string `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members []FederatedDebianRepositoryMember `pulumi:"members"`
	// Internal description.
	Notes       *string `pulumi:"notes"`
	PackageType *string `pulumi:"packageType"`
	// Primary keypair used to sign artifacts. Default value is empty.
	PrimaryKeypairRef *string `pulumi:"primaryKeypairRef"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
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
	// Secondary keypair used to sign artifacts.
	SecondaryKeypairRef *string `pulumi:"secondaryKeypairRef"`
	// When set, the repository will use the deprecated trivial layout.
	//
	// Deprecated: You shouldn't be using this
	TrivialLayout *bool `pulumi:"trivialLayout"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

type FederatedDebianRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect     pulumi.BoolPtrInput
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
	IncludesPattern         pulumi.StringPtrInput
	IndexCompressionFormats pulumi.StringArrayInput
	// the identity key of the repo.
	Key pulumi.StringPtrInput
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members FederatedDebianRepositoryMemberArrayInput
	// Internal description.
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// Primary keypair used to sign artifacts. Default value is empty.
	PrimaryKeypairRef pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  pulumi.BoolPtrInput
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
	// Secondary keypair used to sign artifacts.
	SecondaryKeypairRef pulumi.StringPtrInput
	// When set, the repository will use the deprecated trivial layout.
	//
	// Deprecated: You shouldn't be using this
	TrivialLayout pulumi.BoolPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (FederatedDebianRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedDebianRepositoryState)(nil)).Elem()
}

type federatedDebianRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect     *bool `pulumi:"cdnRedirect"`
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
	IncludesPattern         *string  `pulumi:"includesPattern"`
	IndexCompressionFormats []string `pulumi:"indexCompressionFormats"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members []FederatedDebianRepositoryMember `pulumi:"members"`
	// Internal description.
	Notes *string `pulumi:"notes"`
	// Primary keypair used to sign artifacts. Default value is empty.
	PrimaryKeypairRef *string `pulumi:"primaryKeypairRef"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
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
	// Secondary keypair used to sign artifacts.
	SecondaryKeypairRef *string `pulumi:"secondaryKeypairRef"`
	// When set, the repository will use the deprecated trivial layout.
	//
	// Deprecated: You shouldn't be using this
	TrivialLayout *bool `pulumi:"trivialLayout"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a FederatedDebianRepository resource.
type FederatedDebianRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect     pulumi.BoolPtrInput
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
	IncludesPattern         pulumi.StringPtrInput
	IndexCompressionFormats pulumi.StringArrayInput
	// the identity key of the repo.
	Key pulumi.StringInput
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members FederatedDebianRepositoryMemberArrayInput
	// Internal description.
	Notes pulumi.StringPtrInput
	// Primary keypair used to sign artifacts. Default value is empty.
	PrimaryKeypairRef pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  pulumi.BoolPtrInput
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
	// Secondary keypair used to sign artifacts.
	SecondaryKeypairRef pulumi.StringPtrInput
	// When set, the repository will use the deprecated trivial layout.
	//
	// Deprecated: You shouldn't be using this
	TrivialLayout pulumi.BoolPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (FederatedDebianRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedDebianRepositoryArgs)(nil)).Elem()
}

type FederatedDebianRepositoryInput interface {
	pulumi.Input

	ToFederatedDebianRepositoryOutput() FederatedDebianRepositoryOutput
	ToFederatedDebianRepositoryOutputWithContext(ctx context.Context) FederatedDebianRepositoryOutput
}

func (*FederatedDebianRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedDebianRepository)(nil)).Elem()
}

func (i *FederatedDebianRepository) ToFederatedDebianRepositoryOutput() FederatedDebianRepositoryOutput {
	return i.ToFederatedDebianRepositoryOutputWithContext(context.Background())
}

func (i *FederatedDebianRepository) ToFederatedDebianRepositoryOutputWithContext(ctx context.Context) FederatedDebianRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedDebianRepositoryOutput)
}

// FederatedDebianRepositoryArrayInput is an input type that accepts FederatedDebianRepositoryArray and FederatedDebianRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedDebianRepositoryArrayInput` via:
//
//	FederatedDebianRepositoryArray{ FederatedDebianRepositoryArgs{...} }
type FederatedDebianRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedDebianRepositoryArrayOutput() FederatedDebianRepositoryArrayOutput
	ToFederatedDebianRepositoryArrayOutputWithContext(context.Context) FederatedDebianRepositoryArrayOutput
}

type FederatedDebianRepositoryArray []FederatedDebianRepositoryInput

func (FederatedDebianRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedDebianRepository)(nil)).Elem()
}

func (i FederatedDebianRepositoryArray) ToFederatedDebianRepositoryArrayOutput() FederatedDebianRepositoryArrayOutput {
	return i.ToFederatedDebianRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedDebianRepositoryArray) ToFederatedDebianRepositoryArrayOutputWithContext(ctx context.Context) FederatedDebianRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedDebianRepositoryArrayOutput)
}

// FederatedDebianRepositoryMapInput is an input type that accepts FederatedDebianRepositoryMap and FederatedDebianRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedDebianRepositoryMapInput` via:
//
//	FederatedDebianRepositoryMap{ "key": FederatedDebianRepositoryArgs{...} }
type FederatedDebianRepositoryMapInput interface {
	pulumi.Input

	ToFederatedDebianRepositoryMapOutput() FederatedDebianRepositoryMapOutput
	ToFederatedDebianRepositoryMapOutputWithContext(context.Context) FederatedDebianRepositoryMapOutput
}

type FederatedDebianRepositoryMap map[string]FederatedDebianRepositoryInput

func (FederatedDebianRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedDebianRepository)(nil)).Elem()
}

func (i FederatedDebianRepositoryMap) ToFederatedDebianRepositoryMapOutput() FederatedDebianRepositoryMapOutput {
	return i.ToFederatedDebianRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedDebianRepositoryMap) ToFederatedDebianRepositoryMapOutputWithContext(ctx context.Context) FederatedDebianRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedDebianRepositoryMapOutput)
}

type FederatedDebianRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedDebianRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedDebianRepository)(nil)).Elem()
}

func (o FederatedDebianRepositoryOutput) ToFederatedDebianRepositoryOutput() FederatedDebianRepositoryOutput {
	return o
}

func (o FederatedDebianRepositoryOutput) ToFederatedDebianRepositoryOutputWithContext(ctx context.Context) FederatedDebianRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o FederatedDebianRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o FederatedDebianRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o FederatedDebianRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o FederatedDebianRepositoryOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.BoolPtrOutput { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

// Public description.
func (o FederatedDebianRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
func (o FederatedDebianRepositoryOutput) DisableProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.BoolPtrOutput { return v.DisableProxy }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o FederatedDebianRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o FederatedDebianRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o FederatedDebianRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o FederatedDebianRepositoryOutput) IndexCompressionFormats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringArrayOutput { return v.IndexCompressionFormats }).(pulumi.StringArrayOutput)
}

// the identity key of the repo.
func (o FederatedDebianRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o FederatedDebianRepositoryOutput) Members() FederatedDebianRepositoryMemberArrayOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) FederatedDebianRepositoryMemberArrayOutput { return v.Members }).(FederatedDebianRepositoryMemberArrayOutput)
}

// Internal description.
func (o FederatedDebianRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o FederatedDebianRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Primary keypair used to sign artifacts. Default value is empty.
func (o FederatedDebianRepositoryOutput) PrimaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringPtrOutput { return v.PrimaryKeypairRef }).(pulumi.StringPtrOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o FederatedDebianRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o FederatedDebianRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o FederatedDebianRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o FederatedDebianRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Proxy key from Artifactory Proxies settings. Default is empty field. Can't be set if `disableProxy = true`.
func (o FederatedDebianRepositoryOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringPtrOutput { return v.Proxy }).(pulumi.StringPtrOutput)
}

// Repository layout key for the federated repository
func (o FederatedDebianRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Secondary keypair used to sign artifacts.
func (o FederatedDebianRepositoryOutput) SecondaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringPtrOutput { return v.SecondaryKeypairRef }).(pulumi.StringPtrOutput)
}

// When set, the repository will use the deprecated trivial layout.
//
// Deprecated: You shouldn't be using this
func (o FederatedDebianRepositoryOutput) TrivialLayout() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.BoolPtrOutput { return v.TrivialLayout }).(pulumi.BoolPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o FederatedDebianRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type FederatedDebianRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedDebianRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedDebianRepository)(nil)).Elem()
}

func (o FederatedDebianRepositoryArrayOutput) ToFederatedDebianRepositoryArrayOutput() FederatedDebianRepositoryArrayOutput {
	return o
}

func (o FederatedDebianRepositoryArrayOutput) ToFederatedDebianRepositoryArrayOutputWithContext(ctx context.Context) FederatedDebianRepositoryArrayOutput {
	return o
}

func (o FederatedDebianRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedDebianRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedDebianRepository {
		return vs[0].([]*FederatedDebianRepository)[vs[1].(int)]
	}).(FederatedDebianRepositoryOutput)
}

type FederatedDebianRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedDebianRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedDebianRepository)(nil)).Elem()
}

func (o FederatedDebianRepositoryMapOutput) ToFederatedDebianRepositoryMapOutput() FederatedDebianRepositoryMapOutput {
	return o
}

func (o FederatedDebianRepositoryMapOutput) ToFederatedDebianRepositoryMapOutputWithContext(ctx context.Context) FederatedDebianRepositoryMapOutput {
	return o
}

func (o FederatedDebianRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedDebianRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedDebianRepository {
		return vs[0].(map[string]*FederatedDebianRepository)[vs[1].(string)]
	}).(FederatedDebianRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedDebianRepositoryInput)(nil)).Elem(), &FederatedDebianRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedDebianRepositoryArrayInput)(nil)).Elem(), FederatedDebianRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedDebianRepositoryMapInput)(nil)).Elem(), FederatedDebianRepositoryMap{})
	pulumi.RegisterOutputType(FederatedDebianRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedDebianRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedDebianRepositoryMapOutput{})
}
