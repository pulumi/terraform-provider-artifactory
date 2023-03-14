// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a federated Cocoapods repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v3/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewFederatedCocoapodsRepository(ctx, "terraform-federated-test-cocoapods-repo", &artifactory.FederatedCocoapodsRepositoryArgs{
//				Key: pulumi.String("terraform-federated-test-cocoapods-repo"),
//				Members: artifactory.FederatedCocoapodsRepositoryMemberArray{
//					&artifactory.FederatedCocoapodsRepositoryMemberArgs{
//						Enabled: pulumi.Bool(true),
//						Url:     pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-cocoapods-repo"),
//					},
//					&artifactory.FederatedCocoapodsRepositoryMemberArgs{
//						Enabled: pulumi.Bool(true),
//						Url:     pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-cocoapods-repo-2"),
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
//	$ pulumi import artifactory:index/federatedCocoapodsRepository:FederatedCocoapodsRepository terraform-federated-test-cocoapods-repo terraform-federated-test-cocoapods-repo
//
// ```
type FederatedCocoapodsRepository struct {
	pulumi.CustomResourceState

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrOutput `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrOutput `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolPtrOutput   `pulumi:"cdnRedirect"`
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrOutput `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringOutput `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringOutput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key pulumi.StringOutput `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members     FederatedCocoapodsRepositoryMemberArrayOutput `pulumi:"members"`
	Notes       pulumi.StringPtrOutput                        `pulumi:"notes"`
	PackageType pulumi.StringOutput                           `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrOutput `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
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

// NewFederatedCocoapodsRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedCocoapodsRepository(ctx *pulumi.Context,
	name string, args *FederatedCocoapodsRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedCocoapodsRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FederatedCocoapodsRepository
	err := ctx.RegisterResource("artifactory:index/federatedCocoapodsRepository:FederatedCocoapodsRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedCocoapodsRepository gets an existing FederatedCocoapodsRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedCocoapodsRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedCocoapodsRepositoryState, opts ...pulumi.ResourceOption) (*FederatedCocoapodsRepository, error) {
	var resource FederatedCocoapodsRepository
	err := ctx.ReadResource("artifactory:index/federatedCocoapodsRepository:FederatedCocoapodsRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedCocoapodsRepository resources.
type federatedCocoapodsRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect *bool   `pulumi:"cdnRedirect"`
	Description *string `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect *bool `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key *string `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members     []FederatedCocoapodsRepositoryMember `pulumi:"members"`
	Notes       *string                              `pulumi:"notes"`
	PackageType *string                              `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
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

type FederatedCocoapodsRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolPtrInput
	Description pulumi.StringPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// the identity key of the repo.
	Key pulumi.StringPtrInput
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members     FederatedCocoapodsRepositoryMemberArrayInput
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
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

func (FederatedCocoapodsRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedCocoapodsRepositoryState)(nil)).Elem()
}

type federatedCocoapodsRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect *bool   `pulumi:"cdnRedirect"`
	Description *string `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect *bool `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members []FederatedCocoapodsRepositoryMember `pulumi:"members"`
	Notes   *string                              `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
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

// The set of arguments for constructing a FederatedCocoapodsRepository resource.
type FederatedCocoapodsRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolPtrInput
	Description pulumi.StringPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// the identity key of the repo.
	Key pulumi.StringInput
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members FederatedCocoapodsRepositoryMemberArrayInput
	Notes   pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
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

func (FederatedCocoapodsRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedCocoapodsRepositoryArgs)(nil)).Elem()
}

type FederatedCocoapodsRepositoryInput interface {
	pulumi.Input

	ToFederatedCocoapodsRepositoryOutput() FederatedCocoapodsRepositoryOutput
	ToFederatedCocoapodsRepositoryOutputWithContext(ctx context.Context) FederatedCocoapodsRepositoryOutput
}

func (*FederatedCocoapodsRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedCocoapodsRepository)(nil)).Elem()
}

func (i *FederatedCocoapodsRepository) ToFederatedCocoapodsRepositoryOutput() FederatedCocoapodsRepositoryOutput {
	return i.ToFederatedCocoapodsRepositoryOutputWithContext(context.Background())
}

func (i *FederatedCocoapodsRepository) ToFederatedCocoapodsRepositoryOutputWithContext(ctx context.Context) FederatedCocoapodsRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedCocoapodsRepositoryOutput)
}

// FederatedCocoapodsRepositoryArrayInput is an input type that accepts FederatedCocoapodsRepositoryArray and FederatedCocoapodsRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedCocoapodsRepositoryArrayInput` via:
//
//	FederatedCocoapodsRepositoryArray{ FederatedCocoapodsRepositoryArgs{...} }
type FederatedCocoapodsRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedCocoapodsRepositoryArrayOutput() FederatedCocoapodsRepositoryArrayOutput
	ToFederatedCocoapodsRepositoryArrayOutputWithContext(context.Context) FederatedCocoapodsRepositoryArrayOutput
}

type FederatedCocoapodsRepositoryArray []FederatedCocoapodsRepositoryInput

func (FederatedCocoapodsRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedCocoapodsRepository)(nil)).Elem()
}

func (i FederatedCocoapodsRepositoryArray) ToFederatedCocoapodsRepositoryArrayOutput() FederatedCocoapodsRepositoryArrayOutput {
	return i.ToFederatedCocoapodsRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedCocoapodsRepositoryArray) ToFederatedCocoapodsRepositoryArrayOutputWithContext(ctx context.Context) FederatedCocoapodsRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedCocoapodsRepositoryArrayOutput)
}

// FederatedCocoapodsRepositoryMapInput is an input type that accepts FederatedCocoapodsRepositoryMap and FederatedCocoapodsRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedCocoapodsRepositoryMapInput` via:
//
//	FederatedCocoapodsRepositoryMap{ "key": FederatedCocoapodsRepositoryArgs{...} }
type FederatedCocoapodsRepositoryMapInput interface {
	pulumi.Input

	ToFederatedCocoapodsRepositoryMapOutput() FederatedCocoapodsRepositoryMapOutput
	ToFederatedCocoapodsRepositoryMapOutputWithContext(context.Context) FederatedCocoapodsRepositoryMapOutput
}

type FederatedCocoapodsRepositoryMap map[string]FederatedCocoapodsRepositoryInput

func (FederatedCocoapodsRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedCocoapodsRepository)(nil)).Elem()
}

func (i FederatedCocoapodsRepositoryMap) ToFederatedCocoapodsRepositoryMapOutput() FederatedCocoapodsRepositoryMapOutput {
	return i.ToFederatedCocoapodsRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedCocoapodsRepositoryMap) ToFederatedCocoapodsRepositoryMapOutputWithContext(ctx context.Context) FederatedCocoapodsRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedCocoapodsRepositoryMapOutput)
}

type FederatedCocoapodsRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedCocoapodsRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedCocoapodsRepository)(nil)).Elem()
}

func (o FederatedCocoapodsRepositoryOutput) ToFederatedCocoapodsRepositoryOutput() FederatedCocoapodsRepositoryOutput {
	return o
}

func (o FederatedCocoapodsRepositoryOutput) ToFederatedCocoapodsRepositoryOutputWithContext(ctx context.Context) FederatedCocoapodsRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o FederatedCocoapodsRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCocoapodsRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o FederatedCocoapodsRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCocoapodsRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o FederatedCocoapodsRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCocoapodsRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o FederatedCocoapodsRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedCocoapodsRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o FederatedCocoapodsRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCocoapodsRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
// artifacts are excluded.
func (o FederatedCocoapodsRepositoryOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedCocoapodsRepository) pulumi.StringOutput { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o FederatedCocoapodsRepositoryOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedCocoapodsRepository) pulumi.StringOutput { return v.IncludesPattern }).(pulumi.StringOutput)
}

// the identity key of the repo.
func (o FederatedCocoapodsRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedCocoapodsRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o FederatedCocoapodsRepositoryOutput) Members() FederatedCocoapodsRepositoryMemberArrayOutput {
	return o.ApplyT(func(v *FederatedCocoapodsRepository) FederatedCocoapodsRepositoryMemberArrayOutput { return v.Members }).(FederatedCocoapodsRepositoryMemberArrayOutput)
}

func (o FederatedCocoapodsRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedCocoapodsRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o FederatedCocoapodsRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedCocoapodsRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o FederatedCocoapodsRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCocoapodsRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
// will remain in the Terraform state, which will create state drift during the update.
func (o FederatedCocoapodsRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedCocoapodsRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o FederatedCocoapodsRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedCocoapodsRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o FederatedCocoapodsRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedCocoapodsRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o FederatedCocoapodsRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedCocoapodsRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o FederatedCocoapodsRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCocoapodsRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type FederatedCocoapodsRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedCocoapodsRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedCocoapodsRepository)(nil)).Elem()
}

func (o FederatedCocoapodsRepositoryArrayOutput) ToFederatedCocoapodsRepositoryArrayOutput() FederatedCocoapodsRepositoryArrayOutput {
	return o
}

func (o FederatedCocoapodsRepositoryArrayOutput) ToFederatedCocoapodsRepositoryArrayOutputWithContext(ctx context.Context) FederatedCocoapodsRepositoryArrayOutput {
	return o
}

func (o FederatedCocoapodsRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedCocoapodsRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedCocoapodsRepository {
		return vs[0].([]*FederatedCocoapodsRepository)[vs[1].(int)]
	}).(FederatedCocoapodsRepositoryOutput)
}

type FederatedCocoapodsRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedCocoapodsRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedCocoapodsRepository)(nil)).Elem()
}

func (o FederatedCocoapodsRepositoryMapOutput) ToFederatedCocoapodsRepositoryMapOutput() FederatedCocoapodsRepositoryMapOutput {
	return o
}

func (o FederatedCocoapodsRepositoryMapOutput) ToFederatedCocoapodsRepositoryMapOutputWithContext(ctx context.Context) FederatedCocoapodsRepositoryMapOutput {
	return o
}

func (o FederatedCocoapodsRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedCocoapodsRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedCocoapodsRepository {
		return vs[0].(map[string]*FederatedCocoapodsRepository)[vs[1].(string)]
	}).(FederatedCocoapodsRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedCocoapodsRepositoryInput)(nil)).Elem(), &FederatedCocoapodsRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedCocoapodsRepositoryArrayInput)(nil)).Elem(), FederatedCocoapodsRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedCocoapodsRepositoryMapInput)(nil)).Elem(), FederatedCocoapodsRepositoryMap{})
	pulumi.RegisterOutputType(FederatedCocoapodsRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedCocoapodsRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedCocoapodsRepositoryMapOutput{})
}
