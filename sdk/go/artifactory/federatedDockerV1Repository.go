// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a federated Docker repository.
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
//			_, err := artifactory.NewFederatedDockerV1Repository(ctx, "terraform-federated-test-docker-repo", &artifactory.FederatedDockerV1RepositoryArgs{
//				Key: pulumi.String("terraform-federated-test-docker-repo"),
//				Members: artifactory.FederatedDockerV1RepositoryMemberArray{
//					&artifactory.FederatedDockerV1RepositoryMemberArgs{
//						Enabled: pulumi.Bool(true),
//						Url:     pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-docker-repo"),
//					},
//					&artifactory.FederatedDockerV1RepositoryMemberArgs{
//						Enabled: pulumi.Bool(true),
//						Url:     pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-docker-repo-2"),
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
//	$ pulumi import artifactory:index/federatedDockerV1Repository:FederatedDockerV1Repository terraform-federated-test-docker-repo terraform-federated-test-docker-repo
//
// ```
type FederatedDockerV1Repository struct {
	pulumi.CustomResourceState

	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrOutput `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut          pulumi.BoolPtrOutput `pulumi:"blackedOut"`
	BlockPushingSchema1 pulumi.BoolOutput    `pulumi:"blockPushingSchema1"`
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
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringOutput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key           pulumi.StringOutput `pulumi:"key"`
	MaxUniqueTags pulumi.IntOutput    `pulumi:"maxUniqueTags"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members FederatedDockerV1RepositoryMemberArrayOutput `pulumi:"members"`
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
	TagRetention  pulumi.IntOutput       `pulumi:"tagRetention"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrOutput `pulumi:"xrayIndex"`
}

// NewFederatedDockerV1Repository registers a new resource with the given unique name, arguments, and options.
func NewFederatedDockerV1Repository(ctx *pulumi.Context,
	name string, args *FederatedDockerV1RepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedDockerV1Repository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FederatedDockerV1Repository
	err := ctx.RegisterResource("artifactory:index/federatedDockerV1Repository:FederatedDockerV1Repository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedDockerV1Repository gets an existing FederatedDockerV1Repository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedDockerV1Repository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedDockerV1RepositoryState, opts ...pulumi.ResourceOption) (*FederatedDockerV1Repository, error) {
	var resource FederatedDockerV1Repository
	err := ctx.ReadResource("artifactory:index/federatedDockerV1Repository:FederatedDockerV1Repository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedDockerV1Repository resources.
type federatedDockerV1RepositoryState struct {
	ApiVersion *string `pulumi:"apiVersion"`
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut          *bool `pulumi:"blackedOut"`
	BlockPushingSchema1 *bool `pulumi:"blockPushingSchema1"`
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
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key           *string `pulumi:"key"`
	MaxUniqueTags *int    `pulumi:"maxUniqueTags"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members []FederatedDockerV1RepositoryMember `pulumi:"members"`
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
	TagRetention  *int    `pulumi:"tagRetention"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

type FederatedDockerV1RepositoryState struct {
	ApiVersion pulumi.StringPtrInput
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut          pulumi.BoolPtrInput
	BlockPushingSchema1 pulumi.BoolPtrInput
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
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// the identity key of the repo.
	Key           pulumi.StringPtrInput
	MaxUniqueTags pulumi.IntPtrInput
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members FederatedDockerV1RepositoryMemberArrayInput
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
	TagRetention  pulumi.IntPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (FederatedDockerV1RepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedDockerV1RepositoryState)(nil)).Elem()
}

type federatedDockerV1RepositoryArgs struct {
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
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key           string `pulumi:"key"`
	MaxUniqueTags *int   `pulumi:"maxUniqueTags"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members []FederatedDockerV1RepositoryMember `pulumi:"members"`
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

// The set of arguments for constructing a FederatedDockerV1Repository resource.
type FederatedDockerV1RepositoryArgs struct {
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
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// the identity key of the repo.
	Key           pulumi.StringInput
	MaxUniqueTags pulumi.IntPtrInput
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members FederatedDockerV1RepositoryMemberArrayInput
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

func (FederatedDockerV1RepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedDockerV1RepositoryArgs)(nil)).Elem()
}

type FederatedDockerV1RepositoryInput interface {
	pulumi.Input

	ToFederatedDockerV1RepositoryOutput() FederatedDockerV1RepositoryOutput
	ToFederatedDockerV1RepositoryOutputWithContext(ctx context.Context) FederatedDockerV1RepositoryOutput
}

func (*FederatedDockerV1Repository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedDockerV1Repository)(nil)).Elem()
}

func (i *FederatedDockerV1Repository) ToFederatedDockerV1RepositoryOutput() FederatedDockerV1RepositoryOutput {
	return i.ToFederatedDockerV1RepositoryOutputWithContext(context.Background())
}

func (i *FederatedDockerV1Repository) ToFederatedDockerV1RepositoryOutputWithContext(ctx context.Context) FederatedDockerV1RepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedDockerV1RepositoryOutput)
}

// FederatedDockerV1RepositoryArrayInput is an input type that accepts FederatedDockerV1RepositoryArray and FederatedDockerV1RepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedDockerV1RepositoryArrayInput` via:
//
//	FederatedDockerV1RepositoryArray{ FederatedDockerV1RepositoryArgs{...} }
type FederatedDockerV1RepositoryArrayInput interface {
	pulumi.Input

	ToFederatedDockerV1RepositoryArrayOutput() FederatedDockerV1RepositoryArrayOutput
	ToFederatedDockerV1RepositoryArrayOutputWithContext(context.Context) FederatedDockerV1RepositoryArrayOutput
}

type FederatedDockerV1RepositoryArray []FederatedDockerV1RepositoryInput

func (FederatedDockerV1RepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedDockerV1Repository)(nil)).Elem()
}

func (i FederatedDockerV1RepositoryArray) ToFederatedDockerV1RepositoryArrayOutput() FederatedDockerV1RepositoryArrayOutput {
	return i.ToFederatedDockerV1RepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedDockerV1RepositoryArray) ToFederatedDockerV1RepositoryArrayOutputWithContext(ctx context.Context) FederatedDockerV1RepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedDockerV1RepositoryArrayOutput)
}

// FederatedDockerV1RepositoryMapInput is an input type that accepts FederatedDockerV1RepositoryMap and FederatedDockerV1RepositoryMapOutput values.
// You can construct a concrete instance of `FederatedDockerV1RepositoryMapInput` via:
//
//	FederatedDockerV1RepositoryMap{ "key": FederatedDockerV1RepositoryArgs{...} }
type FederatedDockerV1RepositoryMapInput interface {
	pulumi.Input

	ToFederatedDockerV1RepositoryMapOutput() FederatedDockerV1RepositoryMapOutput
	ToFederatedDockerV1RepositoryMapOutputWithContext(context.Context) FederatedDockerV1RepositoryMapOutput
}

type FederatedDockerV1RepositoryMap map[string]FederatedDockerV1RepositoryInput

func (FederatedDockerV1RepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedDockerV1Repository)(nil)).Elem()
}

func (i FederatedDockerV1RepositoryMap) ToFederatedDockerV1RepositoryMapOutput() FederatedDockerV1RepositoryMapOutput {
	return i.ToFederatedDockerV1RepositoryMapOutputWithContext(context.Background())
}

func (i FederatedDockerV1RepositoryMap) ToFederatedDockerV1RepositoryMapOutputWithContext(ctx context.Context) FederatedDockerV1RepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedDockerV1RepositoryMapOutput)
}

type FederatedDockerV1RepositoryOutput struct{ *pulumi.OutputState }

func (FederatedDockerV1RepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedDockerV1Repository)(nil)).Elem()
}

func (o FederatedDockerV1RepositoryOutput) ToFederatedDockerV1RepositoryOutput() FederatedDockerV1RepositoryOutput {
	return o
}

func (o FederatedDockerV1RepositoryOutput) ToFederatedDockerV1RepositoryOutputWithContext(ctx context.Context) FederatedDockerV1RepositoryOutput {
	return o
}

func (o FederatedDockerV1RepositoryOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedDockerV1Repository) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o FederatedDockerV1RepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDockerV1Repository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o FederatedDockerV1RepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDockerV1Repository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o FederatedDockerV1RepositoryOutput) BlockPushingSchema1() pulumi.BoolOutput {
	return o.ApplyT(func(v *FederatedDockerV1Repository) pulumi.BoolOutput { return v.BlockPushingSchema1 }).(pulumi.BoolOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o FederatedDockerV1RepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDockerV1Repository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

// Delete all federated members on `terraform destroy` if set to `true`. Caution: it will delete all the repositories in
// the federation on other Artifactory instances.
func (o FederatedDockerV1RepositoryOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDockerV1Repository) pulumi.BoolPtrOutput { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

// Public description.
func (o FederatedDockerV1RepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedDockerV1Repository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o FederatedDockerV1RepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDockerV1Repository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
// artifacts are excluded.
func (o FederatedDockerV1RepositoryOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedDockerV1Repository) pulumi.StringOutput { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o FederatedDockerV1RepositoryOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedDockerV1Repository) pulumi.StringOutput { return v.IncludesPattern }).(pulumi.StringOutput)
}

// the identity key of the repo.
func (o FederatedDockerV1RepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedDockerV1Repository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

func (o FederatedDockerV1RepositoryOutput) MaxUniqueTags() pulumi.IntOutput {
	return o.ApplyT(func(v *FederatedDockerV1Repository) pulumi.IntOutput { return v.MaxUniqueTags }).(pulumi.IntOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o FederatedDockerV1RepositoryOutput) Members() FederatedDockerV1RepositoryMemberArrayOutput {
	return o.ApplyT(func(v *FederatedDockerV1Repository) FederatedDockerV1RepositoryMemberArrayOutput { return v.Members }).(FederatedDockerV1RepositoryMemberArrayOutput)
}

// Internal description.
func (o FederatedDockerV1RepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedDockerV1Repository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o FederatedDockerV1RepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedDockerV1Repository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o FederatedDockerV1RepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDockerV1Repository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o FederatedDockerV1RepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedDockerV1Repository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o FederatedDockerV1RepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedDockerV1Repository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o FederatedDockerV1RepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedDockerV1Repository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o FederatedDockerV1RepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedDockerV1Repository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o FederatedDockerV1RepositoryOutput) TagRetention() pulumi.IntOutput {
	return o.ApplyT(func(v *FederatedDockerV1Repository) pulumi.IntOutput { return v.TagRetention }).(pulumi.IntOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o FederatedDockerV1RepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDockerV1Repository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type FederatedDockerV1RepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedDockerV1RepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedDockerV1Repository)(nil)).Elem()
}

func (o FederatedDockerV1RepositoryArrayOutput) ToFederatedDockerV1RepositoryArrayOutput() FederatedDockerV1RepositoryArrayOutput {
	return o
}

func (o FederatedDockerV1RepositoryArrayOutput) ToFederatedDockerV1RepositoryArrayOutputWithContext(ctx context.Context) FederatedDockerV1RepositoryArrayOutput {
	return o
}

func (o FederatedDockerV1RepositoryArrayOutput) Index(i pulumi.IntInput) FederatedDockerV1RepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedDockerV1Repository {
		return vs[0].([]*FederatedDockerV1Repository)[vs[1].(int)]
	}).(FederatedDockerV1RepositoryOutput)
}

type FederatedDockerV1RepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedDockerV1RepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedDockerV1Repository)(nil)).Elem()
}

func (o FederatedDockerV1RepositoryMapOutput) ToFederatedDockerV1RepositoryMapOutput() FederatedDockerV1RepositoryMapOutput {
	return o
}

func (o FederatedDockerV1RepositoryMapOutput) ToFederatedDockerV1RepositoryMapOutputWithContext(ctx context.Context) FederatedDockerV1RepositoryMapOutput {
	return o
}

func (o FederatedDockerV1RepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedDockerV1RepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedDockerV1Repository {
		return vs[0].(map[string]*FederatedDockerV1Repository)[vs[1].(string)]
	}).(FederatedDockerV1RepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedDockerV1RepositoryInput)(nil)).Elem(), &FederatedDockerV1Repository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedDockerV1RepositoryArrayInput)(nil)).Elem(), FederatedDockerV1RepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedDockerV1RepositoryMapInput)(nil)).Elem(), FederatedDockerV1RepositoryMap{})
	pulumi.RegisterOutputType(FederatedDockerV1RepositoryOutput{})
	pulumi.RegisterOutputType(FederatedDockerV1RepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedDockerV1RepositoryMapOutput{})
}
