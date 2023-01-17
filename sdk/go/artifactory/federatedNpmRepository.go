// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a federated Npm repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v2/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewFederatedNpmRepository(ctx, "terraform-federated-test-npm-repo", &artifactory.FederatedNpmRepositoryArgs{
//				Key: pulumi.String("terraform-federated-test-npm-repo"),
//				Members: artifactory.FederatedNpmRepositoryMemberArray{
//					&artifactory.FederatedNpmRepositoryMemberArgs{
//						Enabled: pulumi.Bool(true),
//						Url:     pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-npm-repo"),
//					},
//					&artifactory.FederatedNpmRepositoryMemberArgs{
//						Enabled: pulumi.Bool(true),
//						Url:     pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-npm-repo-2"),
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
//	$ pulumi import artifactory:index/federatedNpmRepository:FederatedNpmRepository terraform-federated-test-npm-repo terraform-federated-test-npm-repo
//
// ```
type FederatedNpmRepository struct {
	pulumi.CustomResourceState

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
	IncludesPattern pulumi.StringOutput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key pulumi.StringOutput `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members     FederatedNpmRepositoryMemberArrayOutput `pulumi:"members"`
	Notes       pulumi.StringPtrOutput                  `pulumi:"notes"`
	PackageType pulumi.StringOutput                     `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrOutput `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
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

// NewFederatedNpmRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedNpmRepository(ctx *pulumi.Context,
	name string, args *FederatedNpmRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedNpmRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FederatedNpmRepository
	err := ctx.RegisterResource("artifactory:index/federatedNpmRepository:FederatedNpmRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedNpmRepository gets an existing FederatedNpmRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedNpmRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedNpmRepositoryState, opts ...pulumi.ResourceOption) (*FederatedNpmRepository, error) {
	var resource FederatedNpmRepository
	err := ctx.ReadResource("artifactory:index/federatedNpmRepository:FederatedNpmRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedNpmRepository resources.
type federatedNpmRepositoryState struct {
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
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key *string `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members     []FederatedNpmRepositoryMember `pulumi:"members"`
	Notes       *string                        `pulumi:"notes"`
	PackageType *string                        `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
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

type FederatedNpmRepositoryState struct {
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
	IncludesPattern pulumi.StringPtrInput
	// the identity key of the repo.
	Key pulumi.StringPtrInput
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members     FederatedNpmRepositoryMemberArrayInput
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
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

func (FederatedNpmRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedNpmRepositoryState)(nil)).Elem()
}

type federatedNpmRepositoryArgs struct {
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
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members []FederatedNpmRepositoryMember `pulumi:"members"`
	Notes   *string                        `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
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

// The set of arguments for constructing a FederatedNpmRepository resource.
type FederatedNpmRepositoryArgs struct {
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
	IncludesPattern pulumi.StringPtrInput
	// the identity key of the repo.
	Key pulumi.StringInput
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members FederatedNpmRepositoryMemberArrayInput
	Notes   pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
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

func (FederatedNpmRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedNpmRepositoryArgs)(nil)).Elem()
}

type FederatedNpmRepositoryInput interface {
	pulumi.Input

	ToFederatedNpmRepositoryOutput() FederatedNpmRepositoryOutput
	ToFederatedNpmRepositoryOutputWithContext(ctx context.Context) FederatedNpmRepositoryOutput
}

func (*FederatedNpmRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedNpmRepository)(nil)).Elem()
}

func (i *FederatedNpmRepository) ToFederatedNpmRepositoryOutput() FederatedNpmRepositoryOutput {
	return i.ToFederatedNpmRepositoryOutputWithContext(context.Background())
}

func (i *FederatedNpmRepository) ToFederatedNpmRepositoryOutputWithContext(ctx context.Context) FederatedNpmRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedNpmRepositoryOutput)
}

// FederatedNpmRepositoryArrayInput is an input type that accepts FederatedNpmRepositoryArray and FederatedNpmRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedNpmRepositoryArrayInput` via:
//
//	FederatedNpmRepositoryArray{ FederatedNpmRepositoryArgs{...} }
type FederatedNpmRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedNpmRepositoryArrayOutput() FederatedNpmRepositoryArrayOutput
	ToFederatedNpmRepositoryArrayOutputWithContext(context.Context) FederatedNpmRepositoryArrayOutput
}

type FederatedNpmRepositoryArray []FederatedNpmRepositoryInput

func (FederatedNpmRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedNpmRepository)(nil)).Elem()
}

func (i FederatedNpmRepositoryArray) ToFederatedNpmRepositoryArrayOutput() FederatedNpmRepositoryArrayOutput {
	return i.ToFederatedNpmRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedNpmRepositoryArray) ToFederatedNpmRepositoryArrayOutputWithContext(ctx context.Context) FederatedNpmRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedNpmRepositoryArrayOutput)
}

// FederatedNpmRepositoryMapInput is an input type that accepts FederatedNpmRepositoryMap and FederatedNpmRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedNpmRepositoryMapInput` via:
//
//	FederatedNpmRepositoryMap{ "key": FederatedNpmRepositoryArgs{...} }
type FederatedNpmRepositoryMapInput interface {
	pulumi.Input

	ToFederatedNpmRepositoryMapOutput() FederatedNpmRepositoryMapOutput
	ToFederatedNpmRepositoryMapOutputWithContext(context.Context) FederatedNpmRepositoryMapOutput
}

type FederatedNpmRepositoryMap map[string]FederatedNpmRepositoryInput

func (FederatedNpmRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedNpmRepository)(nil)).Elem()
}

func (i FederatedNpmRepositoryMap) ToFederatedNpmRepositoryMapOutput() FederatedNpmRepositoryMapOutput {
	return i.ToFederatedNpmRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedNpmRepositoryMap) ToFederatedNpmRepositoryMapOutputWithContext(ctx context.Context) FederatedNpmRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedNpmRepositoryMapOutput)
}

type FederatedNpmRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedNpmRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedNpmRepository)(nil)).Elem()
}

func (o FederatedNpmRepositoryOutput) ToFederatedNpmRepositoryOutput() FederatedNpmRepositoryOutput {
	return o
}

func (o FederatedNpmRepositoryOutput) ToFederatedNpmRepositoryOutputWithContext(ctx context.Context) FederatedNpmRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o FederatedNpmRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedNpmRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o FederatedNpmRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedNpmRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o FederatedNpmRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedNpmRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o FederatedNpmRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedNpmRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
// artifacts are excluded.
func (o FederatedNpmRepositoryOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedNpmRepository) pulumi.StringOutput { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o FederatedNpmRepositoryOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedNpmRepository) pulumi.StringOutput { return v.IncludesPattern }).(pulumi.StringOutput)
}

// the identity key of the repo.
func (o FederatedNpmRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedNpmRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o FederatedNpmRepositoryOutput) Members() FederatedNpmRepositoryMemberArrayOutput {
	return o.ApplyT(func(v *FederatedNpmRepository) FederatedNpmRepositoryMemberArrayOutput { return v.Members }).(FederatedNpmRepositoryMemberArrayOutput)
}

func (o FederatedNpmRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedNpmRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o FederatedNpmRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedNpmRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o FederatedNpmRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedNpmRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
// will remain in the Terraform state, which will create state drift during the update.
func (o FederatedNpmRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedNpmRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o FederatedNpmRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedNpmRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o FederatedNpmRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedNpmRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o FederatedNpmRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedNpmRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o FederatedNpmRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedNpmRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type FederatedNpmRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedNpmRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedNpmRepository)(nil)).Elem()
}

func (o FederatedNpmRepositoryArrayOutput) ToFederatedNpmRepositoryArrayOutput() FederatedNpmRepositoryArrayOutput {
	return o
}

func (o FederatedNpmRepositoryArrayOutput) ToFederatedNpmRepositoryArrayOutputWithContext(ctx context.Context) FederatedNpmRepositoryArrayOutput {
	return o
}

func (o FederatedNpmRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedNpmRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedNpmRepository {
		return vs[0].([]*FederatedNpmRepository)[vs[1].(int)]
	}).(FederatedNpmRepositoryOutput)
}

type FederatedNpmRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedNpmRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedNpmRepository)(nil)).Elem()
}

func (o FederatedNpmRepositoryMapOutput) ToFederatedNpmRepositoryMapOutput() FederatedNpmRepositoryMapOutput {
	return o
}

func (o FederatedNpmRepositoryMapOutput) ToFederatedNpmRepositoryMapOutputWithContext(ctx context.Context) FederatedNpmRepositoryMapOutput {
	return o
}

func (o FederatedNpmRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedNpmRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedNpmRepository {
		return vs[0].(map[string]*FederatedNpmRepository)[vs[1].(string)]
	}).(FederatedNpmRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedNpmRepositoryInput)(nil)).Elem(), &FederatedNpmRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedNpmRepositoryArrayInput)(nil)).Elem(), FederatedNpmRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedNpmRepositoryMapInput)(nil)).Elem(), FederatedNpmRepositoryMap{})
	pulumi.RegisterOutputType(FederatedNpmRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedNpmRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedNpmRepositoryMapOutput{})
}
