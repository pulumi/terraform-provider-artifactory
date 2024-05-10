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

// Creates a local Ivy repository.
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
//			_, err := artifactory.NewLocalIvyRepository(ctx, "terraform-local-test-ivy-repo", &artifactory.LocalIvyRepositoryArgs{
//				Key: pulumi.String("terraform-local-test-ivy-repo"),
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
// Local repositories can be imported using their name, e.g.
//
// ```sh
// $ pulumi import artifactory:index/localIvyRepository:LocalIvyRepository terraform-local-test-ivy-repo terraform-local-test-ivy-repo
// ```
type LocalIvyRepository struct {
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
	// Checksum policy determines how Artifactory behaves when a client checksum for a deployed resource is missing or
	// conflicts with the locally calculated checksum (bad checksum). Options are: "client-checksums", or
	// "server-generated-checksums". Default: "client-checksums"\n For more details, please refer to Checksum Policy -
	// https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy
	ChecksumPolicyType pulumi.StringPtrOutput `pulumi:"checksumPolicyType"`
	// Public description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrOutput `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrOutput `pulumi:"excludesPattern"`
	// If set, Artifactory allows you to deploy release artifacts into this repository.
	HandleReleases pulumi.BoolPtrOutput `pulumi:"handleReleases"`
	// If set, Artifactory allows you to deploy snapshot artifacts into this repository.
	HandleSnapshots pulumi.BoolPtrOutput `pulumi:"handleSnapshots"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrOutput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key pulumi.StringOutput `pulumi:"key"`
	// The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting,
	// older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots pulumi.IntPtrOutput `pulumi:"maxUniqueSnapshots"`
	// Internal description.
	Notes       pulumi.StringPtrOutput `pulumi:"notes"`
	PackageType pulumi.StringOutput    `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  pulumi.BoolPtrOutput     `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrOutput `pulumi:"projectKey"`
	// List of property set name
	PropertySets pulumi.StringArrayOutput `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	// Specifies the naming convention for Maven SNAPSHOT versions. The options are - unique: Version number is based on a
	// time-stamp (default) non-unique: Version number uses a self-overriding naming pattern of
	// artifactId-version-SNAPSHOT.type deployer: Respects the settings in the Maven client that is deploying the artifact.
	SnapshotVersionBehavior pulumi.StringPtrOutput `pulumi:"snapshotVersionBehavior"`
	// By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the
	// groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the
	// deployment with a "409 Conflict" error. You can disable this behavior by setting the Suppress POM Consistency Checks
	// checkbox.
	SuppressPomConsistencyChecks pulumi.BoolPtrOutput `pulumi:"suppressPomConsistencyChecks"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrOutput `pulumi:"xrayIndex"`
}

// NewLocalIvyRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalIvyRepository(ctx *pulumi.Context,
	name string, args *LocalIvyRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalIvyRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LocalIvyRepository
	err := ctx.RegisterResource("artifactory:index/localIvyRepository:LocalIvyRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalIvyRepository gets an existing LocalIvyRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalIvyRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalIvyRepositoryState, opts ...pulumi.ResourceOption) (*LocalIvyRepository, error) {
	var resource LocalIvyRepository
	err := ctx.ReadResource("artifactory:index/localIvyRepository:LocalIvyRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalIvyRepository resources.
type localIvyRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect *bool `pulumi:"cdnRedirect"`
	// Checksum policy determines how Artifactory behaves when a client checksum for a deployed resource is missing or
	// conflicts with the locally calculated checksum (bad checksum). Options are: "client-checksums", or
	// "server-generated-checksums". Default: "client-checksums"\n For more details, please refer to Checksum Policy -
	// https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy
	ChecksumPolicyType *string `pulumi:"checksumPolicyType"`
	// Public description.
	Description *string `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect *bool `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// If set, Artifactory allows you to deploy release artifacts into this repository.
	HandleReleases *bool `pulumi:"handleReleases"`
	// If set, Artifactory allows you to deploy snapshot artifacts into this repository.
	HandleSnapshots *bool `pulumi:"handleSnapshots"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key *string `pulumi:"key"`
	// The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting,
	// older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots *int `pulumi:"maxUniqueSnapshots"`
	// Internal description.
	Notes       *string `pulumi:"notes"`
	PackageType *string `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// List of property set name
	PropertySets []string `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// Specifies the naming convention for Maven SNAPSHOT versions. The options are - unique: Version number is based on a
	// time-stamp (default) non-unique: Version number uses a self-overriding naming pattern of
	// artifactId-version-SNAPSHOT.type deployer: Respects the settings in the Maven client that is deploying the artifact.
	SnapshotVersionBehavior *string `pulumi:"snapshotVersionBehavior"`
	// By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the
	// groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the
	// deployment with a "409 Conflict" error. You can disable this behavior by setting the Suppress POM Consistency Checks
	// checkbox.
	SuppressPomConsistencyChecks *bool `pulumi:"suppressPomConsistencyChecks"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

type LocalIvyRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolPtrInput
	// Checksum policy determines how Artifactory behaves when a client checksum for a deployed resource is missing or
	// conflicts with the locally calculated checksum (bad checksum). Options are: "client-checksums", or
	// "server-generated-checksums". Default: "client-checksums"\n For more details, please refer to Checksum Policy -
	// https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy
	ChecksumPolicyType pulumi.StringPtrInput
	// Public description.
	Description pulumi.StringPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// If set, Artifactory allows you to deploy release artifacts into this repository.
	HandleReleases pulumi.BoolPtrInput
	// If set, Artifactory allows you to deploy snapshot artifacts into this repository.
	HandleSnapshots pulumi.BoolPtrInput
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// the identity key of the repo.
	Key pulumi.StringPtrInput
	// The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting,
	// older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots pulumi.IntPtrInput
	// Internal description.
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  pulumi.BoolPtrInput
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// List of property set name
	PropertySets pulumi.StringArrayInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// Specifies the naming convention for Maven SNAPSHOT versions. The options are - unique: Version number is based on a
	// time-stamp (default) non-unique: Version number uses a self-overriding naming pattern of
	// artifactId-version-SNAPSHOT.type deployer: Respects the settings in the Maven client that is deploying the artifact.
	SnapshotVersionBehavior pulumi.StringPtrInput
	// By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the
	// groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the
	// deployment with a "409 Conflict" error. You can disable this behavior by setting the Suppress POM Consistency Checks
	// checkbox.
	SuppressPomConsistencyChecks pulumi.BoolPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (LocalIvyRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localIvyRepositoryState)(nil)).Elem()
}

type localIvyRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect *bool `pulumi:"cdnRedirect"`
	// Checksum policy determines how Artifactory behaves when a client checksum for a deployed resource is missing or
	// conflicts with the locally calculated checksum (bad checksum). Options are: "client-checksums", or
	// "server-generated-checksums". Default: "client-checksums"\n For more details, please refer to Checksum Policy -
	// https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy
	ChecksumPolicyType *string `pulumi:"checksumPolicyType"`
	// Public description.
	Description *string `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect *bool `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// If set, Artifactory allows you to deploy release artifacts into this repository.
	HandleReleases *bool `pulumi:"handleReleases"`
	// If set, Artifactory allows you to deploy snapshot artifacts into this repository.
	HandleSnapshots *bool `pulumi:"handleSnapshots"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting,
	// older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots *int `pulumi:"maxUniqueSnapshots"`
	// Internal description.
	Notes *string `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// List of property set name
	PropertySets []string `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// Specifies the naming convention for Maven SNAPSHOT versions. The options are - unique: Version number is based on a
	// time-stamp (default) non-unique: Version number uses a self-overriding naming pattern of
	// artifactId-version-SNAPSHOT.type deployer: Respects the settings in the Maven client that is deploying the artifact.
	SnapshotVersionBehavior *string `pulumi:"snapshotVersionBehavior"`
	// By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the
	// groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the
	// deployment with a "409 Conflict" error. You can disable this behavior by setting the Suppress POM Consistency Checks
	// checkbox.
	SuppressPomConsistencyChecks *bool `pulumi:"suppressPomConsistencyChecks"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a LocalIvyRepository resource.
type LocalIvyRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolPtrInput
	// Checksum policy determines how Artifactory behaves when a client checksum for a deployed resource is missing or
	// conflicts with the locally calculated checksum (bad checksum). Options are: "client-checksums", or
	// "server-generated-checksums". Default: "client-checksums"\n For more details, please refer to Checksum Policy -
	// https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy
	ChecksumPolicyType pulumi.StringPtrInput
	// Public description.
	Description pulumi.StringPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// If set, Artifactory allows you to deploy release artifacts into this repository.
	HandleReleases pulumi.BoolPtrInput
	// If set, Artifactory allows you to deploy snapshot artifacts into this repository.
	HandleSnapshots pulumi.BoolPtrInput
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// the identity key of the repo.
	Key pulumi.StringInput
	// The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting,
	// older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots pulumi.IntPtrInput
	// Internal description.
	Notes pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  pulumi.BoolPtrInput
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// List of property set name
	PropertySets pulumi.StringArrayInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// Specifies the naming convention for Maven SNAPSHOT versions. The options are - unique: Version number is based on a
	// time-stamp (default) non-unique: Version number uses a self-overriding naming pattern of
	// artifactId-version-SNAPSHOT.type deployer: Respects the settings in the Maven client that is deploying the artifact.
	SnapshotVersionBehavior pulumi.StringPtrInput
	// By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the
	// groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the
	// deployment with a "409 Conflict" error. You can disable this behavior by setting the Suppress POM Consistency Checks
	// checkbox.
	SuppressPomConsistencyChecks pulumi.BoolPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (LocalIvyRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localIvyRepositoryArgs)(nil)).Elem()
}

type LocalIvyRepositoryInput interface {
	pulumi.Input

	ToLocalIvyRepositoryOutput() LocalIvyRepositoryOutput
	ToLocalIvyRepositoryOutputWithContext(ctx context.Context) LocalIvyRepositoryOutput
}

func (*LocalIvyRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalIvyRepository)(nil)).Elem()
}

func (i *LocalIvyRepository) ToLocalIvyRepositoryOutput() LocalIvyRepositoryOutput {
	return i.ToLocalIvyRepositoryOutputWithContext(context.Background())
}

func (i *LocalIvyRepository) ToLocalIvyRepositoryOutputWithContext(ctx context.Context) LocalIvyRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalIvyRepositoryOutput)
}

// LocalIvyRepositoryArrayInput is an input type that accepts LocalIvyRepositoryArray and LocalIvyRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalIvyRepositoryArrayInput` via:
//
//	LocalIvyRepositoryArray{ LocalIvyRepositoryArgs{...} }
type LocalIvyRepositoryArrayInput interface {
	pulumi.Input

	ToLocalIvyRepositoryArrayOutput() LocalIvyRepositoryArrayOutput
	ToLocalIvyRepositoryArrayOutputWithContext(context.Context) LocalIvyRepositoryArrayOutput
}

type LocalIvyRepositoryArray []LocalIvyRepositoryInput

func (LocalIvyRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalIvyRepository)(nil)).Elem()
}

func (i LocalIvyRepositoryArray) ToLocalIvyRepositoryArrayOutput() LocalIvyRepositoryArrayOutput {
	return i.ToLocalIvyRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalIvyRepositoryArray) ToLocalIvyRepositoryArrayOutputWithContext(ctx context.Context) LocalIvyRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalIvyRepositoryArrayOutput)
}

// LocalIvyRepositoryMapInput is an input type that accepts LocalIvyRepositoryMap and LocalIvyRepositoryMapOutput values.
// You can construct a concrete instance of `LocalIvyRepositoryMapInput` via:
//
//	LocalIvyRepositoryMap{ "key": LocalIvyRepositoryArgs{...} }
type LocalIvyRepositoryMapInput interface {
	pulumi.Input

	ToLocalIvyRepositoryMapOutput() LocalIvyRepositoryMapOutput
	ToLocalIvyRepositoryMapOutputWithContext(context.Context) LocalIvyRepositoryMapOutput
}

type LocalIvyRepositoryMap map[string]LocalIvyRepositoryInput

func (LocalIvyRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalIvyRepository)(nil)).Elem()
}

func (i LocalIvyRepositoryMap) ToLocalIvyRepositoryMapOutput() LocalIvyRepositoryMapOutput {
	return i.ToLocalIvyRepositoryMapOutputWithContext(context.Background())
}

func (i LocalIvyRepositoryMap) ToLocalIvyRepositoryMapOutputWithContext(ctx context.Context) LocalIvyRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalIvyRepositoryMapOutput)
}

type LocalIvyRepositoryOutput struct{ *pulumi.OutputState }

func (LocalIvyRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalIvyRepository)(nil)).Elem()
}

func (o LocalIvyRepositoryOutput) ToLocalIvyRepositoryOutput() LocalIvyRepositoryOutput {
	return o
}

func (o LocalIvyRepositoryOutput) ToLocalIvyRepositoryOutputWithContext(ctx context.Context) LocalIvyRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o LocalIvyRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalIvyRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o LocalIvyRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalIvyRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o LocalIvyRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalIvyRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

// Checksum policy determines how Artifactory behaves when a client checksum for a deployed resource is missing or
// conflicts with the locally calculated checksum (bad checksum). Options are: "client-checksums", or
// "server-generated-checksums". Default: "client-checksums"\n For more details, please refer to Checksum Policy -
// https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy
func (o LocalIvyRepositoryOutput) ChecksumPolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalIvyRepository) pulumi.StringPtrOutput { return v.ChecksumPolicyType }).(pulumi.StringPtrOutput)
}

// Public description.
func (o LocalIvyRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalIvyRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o LocalIvyRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalIvyRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o LocalIvyRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalIvyRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// If set, Artifactory allows you to deploy release artifacts into this repository.
func (o LocalIvyRepositoryOutput) HandleReleases() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalIvyRepository) pulumi.BoolPtrOutput { return v.HandleReleases }).(pulumi.BoolPtrOutput)
}

// If set, Artifactory allows you to deploy snapshot artifacts into this repository.
func (o LocalIvyRepositoryOutput) HandleSnapshots() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalIvyRepository) pulumi.BoolPtrOutput { return v.HandleSnapshots }).(pulumi.BoolPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o LocalIvyRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalIvyRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// the identity key of the repo.
func (o LocalIvyRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalIvyRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting,
// older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
func (o LocalIvyRepositoryOutput) MaxUniqueSnapshots() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LocalIvyRepository) pulumi.IntPtrOutput { return v.MaxUniqueSnapshots }).(pulumi.IntPtrOutput)
}

// Internal description.
func (o LocalIvyRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalIvyRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LocalIvyRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalIvyRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o LocalIvyRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalIvyRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LocalIvyRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalIvyRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o LocalIvyRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalIvyRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o LocalIvyRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalIvyRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o LocalIvyRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalIvyRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Specifies the naming convention for Maven SNAPSHOT versions. The options are - unique: Version number is based on a
// time-stamp (default) non-unique: Version number uses a self-overriding naming pattern of
// artifactId-version-SNAPSHOT.type deployer: Respects the settings in the Maven client that is deploying the artifact.
func (o LocalIvyRepositoryOutput) SnapshotVersionBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalIvyRepository) pulumi.StringPtrOutput { return v.SnapshotVersionBehavior }).(pulumi.StringPtrOutput)
}

// By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the
// groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the
// deployment with a "409 Conflict" error. You can disable this behavior by setting the Suppress POM Consistency Checks
// checkbox.
func (o LocalIvyRepositoryOutput) SuppressPomConsistencyChecks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalIvyRepository) pulumi.BoolPtrOutput { return v.SuppressPomConsistencyChecks }).(pulumi.BoolPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o LocalIvyRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalIvyRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type LocalIvyRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalIvyRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalIvyRepository)(nil)).Elem()
}

func (o LocalIvyRepositoryArrayOutput) ToLocalIvyRepositoryArrayOutput() LocalIvyRepositoryArrayOutput {
	return o
}

func (o LocalIvyRepositoryArrayOutput) ToLocalIvyRepositoryArrayOutputWithContext(ctx context.Context) LocalIvyRepositoryArrayOutput {
	return o
}

func (o LocalIvyRepositoryArrayOutput) Index(i pulumi.IntInput) LocalIvyRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalIvyRepository {
		return vs[0].([]*LocalIvyRepository)[vs[1].(int)]
	}).(LocalIvyRepositoryOutput)
}

type LocalIvyRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalIvyRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalIvyRepository)(nil)).Elem()
}

func (o LocalIvyRepositoryMapOutput) ToLocalIvyRepositoryMapOutput() LocalIvyRepositoryMapOutput {
	return o
}

func (o LocalIvyRepositoryMapOutput) ToLocalIvyRepositoryMapOutputWithContext(ctx context.Context) LocalIvyRepositoryMapOutput {
	return o
}

func (o LocalIvyRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalIvyRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalIvyRepository {
		return vs[0].(map[string]*LocalIvyRepository)[vs[1].(string)]
	}).(LocalIvyRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalIvyRepositoryInput)(nil)).Elem(), &LocalIvyRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalIvyRepositoryArrayInput)(nil)).Elem(), LocalIvyRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalIvyRepositoryMapInput)(nil)).Elem(), LocalIvyRepositoryMap{})
	pulumi.RegisterOutputType(LocalIvyRepositoryOutput{})
	pulumi.RegisterOutputType(LocalIvyRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalIvyRepositoryMapOutput{})
}
