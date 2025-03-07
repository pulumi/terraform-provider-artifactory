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

// Creates a local Docker v1 repository - By choosing a V1 repository, you don't really have many options.
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
//			_, err := artifactory.NewDockerV1Repository(ctx, "foo", &artifactory.DockerV1RepositoryArgs{
//				Key: pulumi.String("foo"),
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
// $ pulumi import artifactory:index/dockerV1Repository:DockerV1Repository foo foo
// ```
type DockerV1Repository struct {
	pulumi.CustomResourceState

	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolOutput `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut          pulumi.BoolOutput `pulumi:"blackedOut"`
	BlockPushingSchema1 pulumi.BoolOutput `pulumi:"blockPushingSchema1"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolOutput `pulumi:"cdnRedirect"`
	// Public description.
	Description pulumi.StringOutput `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolOutput `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringOutput `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
	IncludesPattern pulumi.StringOutput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key           pulumi.StringOutput `pulumi:"key"`
	MaxUniqueTags pulumi.IntOutput    `pulumi:"maxUniqueTags"`
	// Internal description.
	Notes pulumi.StringOutput `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  pulumi.BoolOutput        `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringOutput `pulumi:"projectKey"`
	// List of property set name
	PropertySets pulumi.StringArrayOutput `pulumi:"propertySets"`
	// Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
	// corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
	RepoLayoutRef pulumi.StringOutput `pulumi:"repoLayoutRef"`
	TagRetention  pulumi.IntOutput    `pulumi:"tagRetention"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolOutput `pulumi:"xrayIndex"`
}

// NewDockerV1Repository registers a new resource with the given unique name, arguments, and options.
func NewDockerV1Repository(ctx *pulumi.Context,
	name string, args *DockerV1RepositoryArgs, opts ...pulumi.ResourceOption) (*DockerV1Repository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DockerV1Repository
	err := ctx.RegisterResource("artifactory:index/dockerV1Repository:DockerV1Repository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDockerV1Repository gets an existing DockerV1Repository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDockerV1Repository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DockerV1RepositoryState, opts ...pulumi.ResourceOption) (*DockerV1Repository, error) {
	var resource DockerV1Repository
	err := ctx.ReadResource("artifactory:index/dockerV1Repository:DockerV1Repository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DockerV1Repository resources.
type dockerV1RepositoryState struct {
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
	// Public description.
	Description *string `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect *bool `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key           *string `pulumi:"key"`
	MaxUniqueTags *int    `pulumi:"maxUniqueTags"`
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
	// Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
	// corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	TagRetention  *int    `pulumi:"tagRetention"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

type DockerV1RepositoryState struct {
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
	// Public description.
	Description pulumi.StringPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
	IncludesPattern pulumi.StringPtrInput
	// the identity key of the repo.
	Key           pulumi.StringPtrInput
	MaxUniqueTags pulumi.IntPtrInput
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
	// Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
	// corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
	RepoLayoutRef pulumi.StringPtrInput
	TagRetention  pulumi.IntPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (DockerV1RepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*dockerV1RepositoryState)(nil)).Elem()
}

type dockerV1RepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect *bool `pulumi:"cdnRedirect"`
	// Public description.
	Description *string `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect *bool `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
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
	// Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
	// corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a DockerV1Repository resource.
type DockerV1RepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolPtrInput
	// Public description.
	Description pulumi.StringPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
	IncludesPattern pulumi.StringPtrInput
	// the identity key of the repo.
	Key pulumi.StringInput
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
	// Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
	// corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
	RepoLayoutRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (DockerV1RepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dockerV1RepositoryArgs)(nil)).Elem()
}

type DockerV1RepositoryInput interface {
	pulumi.Input

	ToDockerV1RepositoryOutput() DockerV1RepositoryOutput
	ToDockerV1RepositoryOutputWithContext(ctx context.Context) DockerV1RepositoryOutput
}

func (*DockerV1Repository) ElementType() reflect.Type {
	return reflect.TypeOf((**DockerV1Repository)(nil)).Elem()
}

func (i *DockerV1Repository) ToDockerV1RepositoryOutput() DockerV1RepositoryOutput {
	return i.ToDockerV1RepositoryOutputWithContext(context.Background())
}

func (i *DockerV1Repository) ToDockerV1RepositoryOutputWithContext(ctx context.Context) DockerV1RepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerV1RepositoryOutput)
}

// DockerV1RepositoryArrayInput is an input type that accepts DockerV1RepositoryArray and DockerV1RepositoryArrayOutput values.
// You can construct a concrete instance of `DockerV1RepositoryArrayInput` via:
//
//	DockerV1RepositoryArray{ DockerV1RepositoryArgs{...} }
type DockerV1RepositoryArrayInput interface {
	pulumi.Input

	ToDockerV1RepositoryArrayOutput() DockerV1RepositoryArrayOutput
	ToDockerV1RepositoryArrayOutputWithContext(context.Context) DockerV1RepositoryArrayOutput
}

type DockerV1RepositoryArray []DockerV1RepositoryInput

func (DockerV1RepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DockerV1Repository)(nil)).Elem()
}

func (i DockerV1RepositoryArray) ToDockerV1RepositoryArrayOutput() DockerV1RepositoryArrayOutput {
	return i.ToDockerV1RepositoryArrayOutputWithContext(context.Background())
}

func (i DockerV1RepositoryArray) ToDockerV1RepositoryArrayOutputWithContext(ctx context.Context) DockerV1RepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerV1RepositoryArrayOutput)
}

// DockerV1RepositoryMapInput is an input type that accepts DockerV1RepositoryMap and DockerV1RepositoryMapOutput values.
// You can construct a concrete instance of `DockerV1RepositoryMapInput` via:
//
//	DockerV1RepositoryMap{ "key": DockerV1RepositoryArgs{...} }
type DockerV1RepositoryMapInput interface {
	pulumi.Input

	ToDockerV1RepositoryMapOutput() DockerV1RepositoryMapOutput
	ToDockerV1RepositoryMapOutputWithContext(context.Context) DockerV1RepositoryMapOutput
}

type DockerV1RepositoryMap map[string]DockerV1RepositoryInput

func (DockerV1RepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DockerV1Repository)(nil)).Elem()
}

func (i DockerV1RepositoryMap) ToDockerV1RepositoryMapOutput() DockerV1RepositoryMapOutput {
	return i.ToDockerV1RepositoryMapOutputWithContext(context.Background())
}

func (i DockerV1RepositoryMap) ToDockerV1RepositoryMapOutputWithContext(ctx context.Context) DockerV1RepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerV1RepositoryMapOutput)
}

type DockerV1RepositoryOutput struct{ *pulumi.OutputState }

func (DockerV1RepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DockerV1Repository)(nil)).Elem()
}

func (o DockerV1RepositoryOutput) ToDockerV1RepositoryOutput() DockerV1RepositoryOutput {
	return o
}

func (o DockerV1RepositoryOutput) ToDockerV1RepositoryOutputWithContext(ctx context.Context) DockerV1RepositoryOutput {
	return o
}

func (o DockerV1RepositoryOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *DockerV1Repository) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o DockerV1RepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *DockerV1Repository) pulumi.BoolOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o DockerV1RepositoryOutput) BlackedOut() pulumi.BoolOutput {
	return o.ApplyT(func(v *DockerV1Repository) pulumi.BoolOutput { return v.BlackedOut }).(pulumi.BoolOutput)
}

func (o DockerV1RepositoryOutput) BlockPushingSchema1() pulumi.BoolOutput {
	return o.ApplyT(func(v *DockerV1Repository) pulumi.BoolOutput { return v.BlockPushingSchema1 }).(pulumi.BoolOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o DockerV1RepositoryOutput) CdnRedirect() pulumi.BoolOutput {
	return o.ApplyT(func(v *DockerV1Repository) pulumi.BoolOutput { return v.CdnRedirect }).(pulumi.BoolOutput)
}

// Public description.
func (o DockerV1RepositoryOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *DockerV1Repository) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o DockerV1RepositoryOutput) DownloadDirect() pulumi.BoolOutput {
	return o.ApplyT(func(v *DockerV1Repository) pulumi.BoolOutput { return v.DownloadDirect }).(pulumi.BoolOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
// artifacts are excluded.
func (o DockerV1RepositoryOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *DockerV1Repository) pulumi.StringOutput { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
func (o DockerV1RepositoryOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *DockerV1Repository) pulumi.StringOutput { return v.IncludesPattern }).(pulumi.StringOutput)
}

// the identity key of the repo.
func (o DockerV1RepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *DockerV1Repository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

func (o DockerV1RepositoryOutput) MaxUniqueTags() pulumi.IntOutput {
	return o.ApplyT(func(v *DockerV1Repository) pulumi.IntOutput { return v.MaxUniqueTags }).(pulumi.IntOutput)
}

// Internal description.
func (o DockerV1RepositoryOutput) Notes() pulumi.StringOutput {
	return o.ApplyT(func(v *DockerV1Repository) pulumi.StringOutput { return v.Notes }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o DockerV1RepositoryOutput) PriorityResolution() pulumi.BoolOutput {
	return o.ApplyT(func(v *DockerV1Repository) pulumi.BoolOutput { return v.PriorityResolution }).(pulumi.BoolOutput)
}

func (o DockerV1RepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DockerV1Repository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o DockerV1RepositoryOutput) ProjectKey() pulumi.StringOutput {
	return o.ApplyT(func(v *DockerV1Repository) pulumi.StringOutput { return v.ProjectKey }).(pulumi.StringOutput)
}

// List of property set name
func (o DockerV1RepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DockerV1Repository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
// corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
func (o DockerV1RepositoryOutput) RepoLayoutRef() pulumi.StringOutput {
	return o.ApplyT(func(v *DockerV1Repository) pulumi.StringOutput { return v.RepoLayoutRef }).(pulumi.StringOutput)
}

func (o DockerV1RepositoryOutput) TagRetention() pulumi.IntOutput {
	return o.ApplyT(func(v *DockerV1Repository) pulumi.IntOutput { return v.TagRetention }).(pulumi.IntOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o DockerV1RepositoryOutput) XrayIndex() pulumi.BoolOutput {
	return o.ApplyT(func(v *DockerV1Repository) pulumi.BoolOutput { return v.XrayIndex }).(pulumi.BoolOutput)
}

type DockerV1RepositoryArrayOutput struct{ *pulumi.OutputState }

func (DockerV1RepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DockerV1Repository)(nil)).Elem()
}

func (o DockerV1RepositoryArrayOutput) ToDockerV1RepositoryArrayOutput() DockerV1RepositoryArrayOutput {
	return o
}

func (o DockerV1RepositoryArrayOutput) ToDockerV1RepositoryArrayOutputWithContext(ctx context.Context) DockerV1RepositoryArrayOutput {
	return o
}

func (o DockerV1RepositoryArrayOutput) Index(i pulumi.IntInput) DockerV1RepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DockerV1Repository {
		return vs[0].([]*DockerV1Repository)[vs[1].(int)]
	}).(DockerV1RepositoryOutput)
}

type DockerV1RepositoryMapOutput struct{ *pulumi.OutputState }

func (DockerV1RepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DockerV1Repository)(nil)).Elem()
}

func (o DockerV1RepositoryMapOutput) ToDockerV1RepositoryMapOutput() DockerV1RepositoryMapOutput {
	return o
}

func (o DockerV1RepositoryMapOutput) ToDockerV1RepositoryMapOutputWithContext(ctx context.Context) DockerV1RepositoryMapOutput {
	return o
}

func (o DockerV1RepositoryMapOutput) MapIndex(k pulumi.StringInput) DockerV1RepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DockerV1Repository {
		return vs[0].(map[string]*DockerV1Repository)[vs[1].(string)]
	}).(DockerV1RepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DockerV1RepositoryInput)(nil)).Elem(), &DockerV1Repository{})
	pulumi.RegisterInputType(reflect.TypeOf((*DockerV1RepositoryArrayInput)(nil)).Elem(), DockerV1RepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DockerV1RepositoryMapInput)(nil)).Elem(), DockerV1RepositoryMap{})
	pulumi.RegisterOutputType(DockerV1RepositoryOutput{})
	pulumi.RegisterOutputType(DockerV1RepositoryArrayOutput{})
	pulumi.RegisterOutputType(DockerV1RepositoryMapOutput{})
}
