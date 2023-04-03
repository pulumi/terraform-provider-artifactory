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
// ~>This resource has been superseded by the `FederatedDockerV2Repository` resource. This resource will continue to be available in the provider for backward compatibility. For documentation, please refer to the new resource.
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
//			_, err := artifactory.NewFederatedDockerRepository(ctx, "terraform-federated-test-docker-repo", &artifactory.FederatedDockerRepositoryArgs{
//				Key: pulumi.String("terraform-federated-test-docker-repo"),
//				Members: artifactory.FederatedDockerRepositoryMemberArray{
//					&artifactory.FederatedDockerRepositoryMemberArgs{
//						Enabled: pulumi.Bool(true),
//						Url:     pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-docker-repo"),
//					},
//					&artifactory.FederatedDockerRepositoryMemberArgs{
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
//	$ pulumi import artifactory:index/federatedDockerRepository:FederatedDockerRepository terraform-federated-test-docker-repo terraform-federated-test-docker-repo
//
// ```
type FederatedDockerRepository struct {
	pulumi.CustomResourceState

	// The Docker API version to use. This cannot be set
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrOutput `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrOutput `pulumi:"blackedOut"`
	// When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.
	BlockPushingSchema1 pulumi.BoolOutput `pulumi:"blockPushingSchema1"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolPtrOutput `pulumi:"cdnRedirect"`
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
	Key pulumi.StringOutput `pulumi:"key"`
	// The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an
	// image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit. This only
	// applies to manifest v2
	MaxUniqueTags pulumi.IntPtrOutput `pulumi:"maxUniqueTags"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members FederatedDockerRepositoryMemberArrayOutput `pulumi:"members"`
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
	// If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to
	// manifest V2
	TagRetention pulumi.IntPtrOutput `pulumi:"tagRetention"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrOutput `pulumi:"xrayIndex"`
}

// NewFederatedDockerRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedDockerRepository(ctx *pulumi.Context,
	name string, args *FederatedDockerRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedDockerRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FederatedDockerRepository
	err := ctx.RegisterResource("artifactory:index/federatedDockerRepository:FederatedDockerRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedDockerRepository gets an existing FederatedDockerRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedDockerRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedDockerRepositoryState, opts ...pulumi.ResourceOption) (*FederatedDockerRepository, error) {
	var resource FederatedDockerRepository
	err := ctx.ReadResource("artifactory:index/federatedDockerRepository:FederatedDockerRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedDockerRepository resources.
type federatedDockerRepositoryState struct {
	// The Docker API version to use. This cannot be set
	ApiVersion *string `pulumi:"apiVersion"`
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.
	BlockPushingSchema1 *bool `pulumi:"blockPushingSchema1"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect *bool `pulumi:"cdnRedirect"`
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
	Key *string `pulumi:"key"`
	// The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an
	// image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit. This only
	// applies to manifest v2
	MaxUniqueTags *int `pulumi:"maxUniqueTags"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members []FederatedDockerRepositoryMember `pulumi:"members"`
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
	// If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to
	// manifest V2
	TagRetention *int `pulumi:"tagRetention"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

type FederatedDockerRepositoryState struct {
	// The Docker API version to use. This cannot be set
	ApiVersion pulumi.StringPtrInput
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.
	BlockPushingSchema1 pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolPtrInput
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
	Key pulumi.StringPtrInput
	// The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an
	// image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit. This only
	// applies to manifest v2
	MaxUniqueTags pulumi.IntPtrInput
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members FederatedDockerRepositoryMemberArrayInput
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
	// If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to
	// manifest V2
	TagRetention pulumi.IntPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (FederatedDockerRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedDockerRepositoryState)(nil)).Elem()
}

type federatedDockerRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.
	BlockPushingSchema1 *bool `pulumi:"blockPushingSchema1"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect *bool `pulumi:"cdnRedirect"`
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
	Key string `pulumi:"key"`
	// The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an
	// image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit. This only
	// applies to manifest v2
	MaxUniqueTags *int `pulumi:"maxUniqueTags"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members []FederatedDockerRepositoryMember `pulumi:"members"`
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
	// If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to
	// manifest V2
	TagRetention *int `pulumi:"tagRetention"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a FederatedDockerRepository resource.
type FederatedDockerRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.
	BlockPushingSchema1 pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolPtrInput
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
	Key pulumi.StringInput
	// The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an
	// image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit. This only
	// applies to manifest v2
	MaxUniqueTags pulumi.IntPtrInput
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members FederatedDockerRepositoryMemberArrayInput
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
	// If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to
	// manifest V2
	TagRetention pulumi.IntPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (FederatedDockerRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedDockerRepositoryArgs)(nil)).Elem()
}

type FederatedDockerRepositoryInput interface {
	pulumi.Input

	ToFederatedDockerRepositoryOutput() FederatedDockerRepositoryOutput
	ToFederatedDockerRepositoryOutputWithContext(ctx context.Context) FederatedDockerRepositoryOutput
}

func (*FederatedDockerRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedDockerRepository)(nil)).Elem()
}

func (i *FederatedDockerRepository) ToFederatedDockerRepositoryOutput() FederatedDockerRepositoryOutput {
	return i.ToFederatedDockerRepositoryOutputWithContext(context.Background())
}

func (i *FederatedDockerRepository) ToFederatedDockerRepositoryOutputWithContext(ctx context.Context) FederatedDockerRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedDockerRepositoryOutput)
}

// FederatedDockerRepositoryArrayInput is an input type that accepts FederatedDockerRepositoryArray and FederatedDockerRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedDockerRepositoryArrayInput` via:
//
//	FederatedDockerRepositoryArray{ FederatedDockerRepositoryArgs{...} }
type FederatedDockerRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedDockerRepositoryArrayOutput() FederatedDockerRepositoryArrayOutput
	ToFederatedDockerRepositoryArrayOutputWithContext(context.Context) FederatedDockerRepositoryArrayOutput
}

type FederatedDockerRepositoryArray []FederatedDockerRepositoryInput

func (FederatedDockerRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedDockerRepository)(nil)).Elem()
}

func (i FederatedDockerRepositoryArray) ToFederatedDockerRepositoryArrayOutput() FederatedDockerRepositoryArrayOutput {
	return i.ToFederatedDockerRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedDockerRepositoryArray) ToFederatedDockerRepositoryArrayOutputWithContext(ctx context.Context) FederatedDockerRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedDockerRepositoryArrayOutput)
}

// FederatedDockerRepositoryMapInput is an input type that accepts FederatedDockerRepositoryMap and FederatedDockerRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedDockerRepositoryMapInput` via:
//
//	FederatedDockerRepositoryMap{ "key": FederatedDockerRepositoryArgs{...} }
type FederatedDockerRepositoryMapInput interface {
	pulumi.Input

	ToFederatedDockerRepositoryMapOutput() FederatedDockerRepositoryMapOutput
	ToFederatedDockerRepositoryMapOutputWithContext(context.Context) FederatedDockerRepositoryMapOutput
}

type FederatedDockerRepositoryMap map[string]FederatedDockerRepositoryInput

func (FederatedDockerRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedDockerRepository)(nil)).Elem()
}

func (i FederatedDockerRepositoryMap) ToFederatedDockerRepositoryMapOutput() FederatedDockerRepositoryMapOutput {
	return i.ToFederatedDockerRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedDockerRepositoryMap) ToFederatedDockerRepositoryMapOutputWithContext(ctx context.Context) FederatedDockerRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedDockerRepositoryMapOutput)
}

type FederatedDockerRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedDockerRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedDockerRepository)(nil)).Elem()
}

func (o FederatedDockerRepositoryOutput) ToFederatedDockerRepositoryOutput() FederatedDockerRepositoryOutput {
	return o
}

func (o FederatedDockerRepositoryOutput) ToFederatedDockerRepositoryOutputWithContext(ctx context.Context) FederatedDockerRepositoryOutput {
	return o
}

// The Docker API version to use. This cannot be set
func (o FederatedDockerRepositoryOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedDockerRepository) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o FederatedDockerRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDockerRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o FederatedDockerRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDockerRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.
func (o FederatedDockerRepositoryOutput) BlockPushingSchema1() pulumi.BoolOutput {
	return o.ApplyT(func(v *FederatedDockerRepository) pulumi.BoolOutput { return v.BlockPushingSchema1 }).(pulumi.BoolOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o FederatedDockerRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDockerRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

// Public description.
func (o FederatedDockerRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedDockerRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o FederatedDockerRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDockerRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
// artifacts are excluded.
func (o FederatedDockerRepositoryOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedDockerRepository) pulumi.StringOutput { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o FederatedDockerRepositoryOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedDockerRepository) pulumi.StringOutput { return v.IncludesPattern }).(pulumi.StringOutput)
}

// the identity key of the repo.
func (o FederatedDockerRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedDockerRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an
// image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit. This only
// applies to manifest v2
func (o FederatedDockerRepositoryOutput) MaxUniqueTags() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FederatedDockerRepository) pulumi.IntPtrOutput { return v.MaxUniqueTags }).(pulumi.IntPtrOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o FederatedDockerRepositoryOutput) Members() FederatedDockerRepositoryMemberArrayOutput {
	return o.ApplyT(func(v *FederatedDockerRepository) FederatedDockerRepositoryMemberArrayOutput { return v.Members }).(FederatedDockerRepositoryMemberArrayOutput)
}

// Internal description.
func (o FederatedDockerRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedDockerRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o FederatedDockerRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedDockerRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o FederatedDockerRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDockerRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o FederatedDockerRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedDockerRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o FederatedDockerRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedDockerRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o FederatedDockerRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedDockerRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o FederatedDockerRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedDockerRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to
// manifest V2
func (o FederatedDockerRepositoryOutput) TagRetention() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FederatedDockerRepository) pulumi.IntPtrOutput { return v.TagRetention }).(pulumi.IntPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o FederatedDockerRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDockerRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type FederatedDockerRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedDockerRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedDockerRepository)(nil)).Elem()
}

func (o FederatedDockerRepositoryArrayOutput) ToFederatedDockerRepositoryArrayOutput() FederatedDockerRepositoryArrayOutput {
	return o
}

func (o FederatedDockerRepositoryArrayOutput) ToFederatedDockerRepositoryArrayOutputWithContext(ctx context.Context) FederatedDockerRepositoryArrayOutput {
	return o
}

func (o FederatedDockerRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedDockerRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedDockerRepository {
		return vs[0].([]*FederatedDockerRepository)[vs[1].(int)]
	}).(FederatedDockerRepositoryOutput)
}

type FederatedDockerRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedDockerRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedDockerRepository)(nil)).Elem()
}

func (o FederatedDockerRepositoryMapOutput) ToFederatedDockerRepositoryMapOutput() FederatedDockerRepositoryMapOutput {
	return o
}

func (o FederatedDockerRepositoryMapOutput) ToFederatedDockerRepositoryMapOutputWithContext(ctx context.Context) FederatedDockerRepositoryMapOutput {
	return o
}

func (o FederatedDockerRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedDockerRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedDockerRepository {
		return vs[0].(map[string]*FederatedDockerRepository)[vs[1].(string)]
	}).(FederatedDockerRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedDockerRepositoryInput)(nil)).Elem(), &FederatedDockerRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedDockerRepositoryArrayInput)(nil)).Elem(), FederatedDockerRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedDockerRepositoryMapInput)(nil)).Elem(), FederatedDockerRepositoryMap{})
	pulumi.RegisterOutputType(FederatedDockerRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedDockerRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedDockerRepositoryMapOutput{})
}
