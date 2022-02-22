// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Local Docker V2 Repository Resource
//
// Creates a local Docker v2 repository
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-artifactory/sdk/go/artifactory"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := artifactory.NewDockerV2Repository(ctx, "foo", &artifactory.DockerV2RepositoryArgs{
// 			Key:           pulumi.String("foo"),
// 			MaxUniqueTags: pulumi.Int(5),
// 			TagRetention:  pulumi.Int(3),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DockerV2Repository struct {
	pulumi.CustomResourceState

	// The Docker API version to use. This cannot be set
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrOutput `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrOutput `pulumi:"blackedOut"`
	// - When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.
	BlockPushingSchema1 pulumi.BoolOutput      `pulumi:"blockPushingSchema1"`
	Description         pulumi.StringPtrOutput `pulumi:"description"`
	DownloadDirect      pulumi.BoolPtrOutput   `pulumi:"downloadDirect"`
	ExcludesPattern     pulumi.StringOutput    `pulumi:"excludesPattern"`
	IncludesPattern     pulumi.StringOutput    `pulumi:"includesPattern"`
	// - the identity key of the repo
	Key pulumi.StringOutput `pulumi:"key"`
	// - The maximum number of unique tags of a single Docker image to store in this repository.\n" +
	//   Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit.
	//   This only applies to manifest v2
	MaxUniqueTags pulumi.IntPtrOutput    `pulumi:"maxUniqueTags"`
	Notes         pulumi.StringPtrOutput `pulumi:"notes"`
	PackageType   pulumi.StringOutput    `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrOutput `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    pulumi.StringPtrOutput   `pulumi:"projectKey"`
	PropertySets  pulumi.StringArrayOutput `pulumi:"propertySets"`
	RepoLayoutRef pulumi.StringOutput      `pulumi:"repoLayoutRef"`
	// - If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to manifest V2
	TagRetention pulumi.IntPtrOutput `pulumi:"tagRetention"`
	XrayIndex    pulumi.BoolOutput   `pulumi:"xrayIndex"`
}

// NewDockerV2Repository registers a new resource with the given unique name, arguments, and options.
func NewDockerV2Repository(ctx *pulumi.Context,
	name string, args *DockerV2RepositoryArgs, opts ...pulumi.ResourceOption) (*DockerV2Repository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource DockerV2Repository
	err := ctx.RegisterResource("artifactory:index/dockerV2Repository:DockerV2Repository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDockerV2Repository gets an existing DockerV2Repository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDockerV2Repository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DockerV2RepositoryState, opts ...pulumi.ResourceOption) (*DockerV2Repository, error) {
	var resource DockerV2Repository
	err := ctx.ReadResource("artifactory:index/dockerV2Repository:DockerV2Repository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DockerV2Repository resources.
type dockerV2RepositoryState struct {
	// The Docker API version to use. This cannot be set
	ApiVersion *string `pulumi:"apiVersion"`
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool `pulumi:"blackedOut"`
	// - When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.
	BlockPushingSchema1 *bool   `pulumi:"blockPushingSchema1"`
	Description         *string `pulumi:"description"`
	DownloadDirect      *bool   `pulumi:"downloadDirect"`
	ExcludesPattern     *string `pulumi:"excludesPattern"`
	IncludesPattern     *string `pulumi:"includesPattern"`
	// - the identity key of the repo
	Key *string `pulumi:"key"`
	// - The maximum number of unique tags of a single Docker image to store in this repository.\n" +
	//   Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit.
	//   This only applies to manifest v2
	MaxUniqueTags *int    `pulumi:"maxUniqueTags"`
	Notes         *string `pulumi:"notes"`
	PackageType   *string `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    *string  `pulumi:"projectKey"`
	PropertySets  []string `pulumi:"propertySets"`
	RepoLayoutRef *string  `pulumi:"repoLayoutRef"`
	// - If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to manifest V2
	TagRetention *int  `pulumi:"tagRetention"`
	XrayIndex    *bool `pulumi:"xrayIndex"`
}

type DockerV2RepositoryState struct {
	// The Docker API version to use. This cannot be set
	ApiVersion pulumi.StringPtrInput
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	BlackedOut             pulumi.BoolPtrInput
	// - When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.
	BlockPushingSchema1 pulumi.BoolPtrInput
	Description         pulumi.StringPtrInput
	DownloadDirect      pulumi.BoolPtrInput
	ExcludesPattern     pulumi.StringPtrInput
	IncludesPattern     pulumi.StringPtrInput
	// - the identity key of the repo
	Key pulumi.StringPtrInput
	// - The maximum number of unique tags of a single Docker image to store in this repository.\n" +
	//   Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit.
	//   This only applies to manifest v2
	MaxUniqueTags pulumi.IntPtrInput
	Notes         pulumi.StringPtrInput
	PackageType   pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    pulumi.StringPtrInput
	PropertySets  pulumi.StringArrayInput
	RepoLayoutRef pulumi.StringPtrInput
	// - If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to manifest V2
	TagRetention pulumi.IntPtrInput
	XrayIndex    pulumi.BoolPtrInput
}

func (DockerV2RepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*dockerV2RepositoryState)(nil)).Elem()
}

type dockerV2RepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool `pulumi:"blackedOut"`
	// - When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.
	BlockPushingSchema1 *bool   `pulumi:"blockPushingSchema1"`
	Description         *string `pulumi:"description"`
	DownloadDirect      *bool   `pulumi:"downloadDirect"`
	ExcludesPattern     *string `pulumi:"excludesPattern"`
	IncludesPattern     *string `pulumi:"includesPattern"`
	// - the identity key of the repo
	Key string `pulumi:"key"`
	// - The maximum number of unique tags of a single Docker image to store in this repository.\n" +
	//   Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit.
	//   This only applies to manifest v2
	MaxUniqueTags *int    `pulumi:"maxUniqueTags"`
	Notes         *string `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    *string  `pulumi:"projectKey"`
	PropertySets  []string `pulumi:"propertySets"`
	RepoLayoutRef *string  `pulumi:"repoLayoutRef"`
	// - If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to manifest V2
	TagRetention *int  `pulumi:"tagRetention"`
	XrayIndex    *bool `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a DockerV2Repository resource.
type DockerV2RepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	BlackedOut             pulumi.BoolPtrInput
	// - When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.
	BlockPushingSchema1 pulumi.BoolPtrInput
	Description         pulumi.StringPtrInput
	DownloadDirect      pulumi.BoolPtrInput
	ExcludesPattern     pulumi.StringPtrInput
	IncludesPattern     pulumi.StringPtrInput
	// - the identity key of the repo
	Key pulumi.StringInput
	// - The maximum number of unique tags of a single Docker image to store in this repository.\n" +
	//   Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit.
	//   This only applies to manifest v2
	MaxUniqueTags pulumi.IntPtrInput
	Notes         pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    pulumi.StringPtrInput
	PropertySets  pulumi.StringArrayInput
	RepoLayoutRef pulumi.StringPtrInput
	// - If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to manifest V2
	TagRetention pulumi.IntPtrInput
	XrayIndex    pulumi.BoolPtrInput
}

func (DockerV2RepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dockerV2RepositoryArgs)(nil)).Elem()
}

type DockerV2RepositoryInput interface {
	pulumi.Input

	ToDockerV2RepositoryOutput() DockerV2RepositoryOutput
	ToDockerV2RepositoryOutputWithContext(ctx context.Context) DockerV2RepositoryOutput
}

func (*DockerV2Repository) ElementType() reflect.Type {
	return reflect.TypeOf((**DockerV2Repository)(nil)).Elem()
}

func (i *DockerV2Repository) ToDockerV2RepositoryOutput() DockerV2RepositoryOutput {
	return i.ToDockerV2RepositoryOutputWithContext(context.Background())
}

func (i *DockerV2Repository) ToDockerV2RepositoryOutputWithContext(ctx context.Context) DockerV2RepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerV2RepositoryOutput)
}

// DockerV2RepositoryArrayInput is an input type that accepts DockerV2RepositoryArray and DockerV2RepositoryArrayOutput values.
// You can construct a concrete instance of `DockerV2RepositoryArrayInput` via:
//
//          DockerV2RepositoryArray{ DockerV2RepositoryArgs{...} }
type DockerV2RepositoryArrayInput interface {
	pulumi.Input

	ToDockerV2RepositoryArrayOutput() DockerV2RepositoryArrayOutput
	ToDockerV2RepositoryArrayOutputWithContext(context.Context) DockerV2RepositoryArrayOutput
}

type DockerV2RepositoryArray []DockerV2RepositoryInput

func (DockerV2RepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DockerV2Repository)(nil)).Elem()
}

func (i DockerV2RepositoryArray) ToDockerV2RepositoryArrayOutput() DockerV2RepositoryArrayOutput {
	return i.ToDockerV2RepositoryArrayOutputWithContext(context.Background())
}

func (i DockerV2RepositoryArray) ToDockerV2RepositoryArrayOutputWithContext(ctx context.Context) DockerV2RepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerV2RepositoryArrayOutput)
}

// DockerV2RepositoryMapInput is an input type that accepts DockerV2RepositoryMap and DockerV2RepositoryMapOutput values.
// You can construct a concrete instance of `DockerV2RepositoryMapInput` via:
//
//          DockerV2RepositoryMap{ "key": DockerV2RepositoryArgs{...} }
type DockerV2RepositoryMapInput interface {
	pulumi.Input

	ToDockerV2RepositoryMapOutput() DockerV2RepositoryMapOutput
	ToDockerV2RepositoryMapOutputWithContext(context.Context) DockerV2RepositoryMapOutput
}

type DockerV2RepositoryMap map[string]DockerV2RepositoryInput

func (DockerV2RepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DockerV2Repository)(nil)).Elem()
}

func (i DockerV2RepositoryMap) ToDockerV2RepositoryMapOutput() DockerV2RepositoryMapOutput {
	return i.ToDockerV2RepositoryMapOutputWithContext(context.Background())
}

func (i DockerV2RepositoryMap) ToDockerV2RepositoryMapOutputWithContext(ctx context.Context) DockerV2RepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerV2RepositoryMapOutput)
}

type DockerV2RepositoryOutput struct{ *pulumi.OutputState }

func (DockerV2RepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DockerV2Repository)(nil)).Elem()
}

func (o DockerV2RepositoryOutput) ToDockerV2RepositoryOutput() DockerV2RepositoryOutput {
	return o
}

func (o DockerV2RepositoryOutput) ToDockerV2RepositoryOutputWithContext(ctx context.Context) DockerV2RepositoryOutput {
	return o
}

type DockerV2RepositoryArrayOutput struct{ *pulumi.OutputState }

func (DockerV2RepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DockerV2Repository)(nil)).Elem()
}

func (o DockerV2RepositoryArrayOutput) ToDockerV2RepositoryArrayOutput() DockerV2RepositoryArrayOutput {
	return o
}

func (o DockerV2RepositoryArrayOutput) ToDockerV2RepositoryArrayOutputWithContext(ctx context.Context) DockerV2RepositoryArrayOutput {
	return o
}

func (o DockerV2RepositoryArrayOutput) Index(i pulumi.IntInput) DockerV2RepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DockerV2Repository {
		return vs[0].([]*DockerV2Repository)[vs[1].(int)]
	}).(DockerV2RepositoryOutput)
}

type DockerV2RepositoryMapOutput struct{ *pulumi.OutputState }

func (DockerV2RepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DockerV2Repository)(nil)).Elem()
}

func (o DockerV2RepositoryMapOutput) ToDockerV2RepositoryMapOutput() DockerV2RepositoryMapOutput {
	return o
}

func (o DockerV2RepositoryMapOutput) ToDockerV2RepositoryMapOutputWithContext(ctx context.Context) DockerV2RepositoryMapOutput {
	return o
}

func (o DockerV2RepositoryMapOutput) MapIndex(k pulumi.StringInput) DockerV2RepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DockerV2Repository {
		return vs[0].(map[string]*DockerV2Repository)[vs[1].(string)]
	}).(DockerV2RepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DockerV2RepositoryInput)(nil)).Elem(), &DockerV2Repository{})
	pulumi.RegisterInputType(reflect.TypeOf((*DockerV2RepositoryArrayInput)(nil)).Elem(), DockerV2RepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DockerV2RepositoryMapInput)(nil)).Elem(), DockerV2RepositoryMap{})
	pulumi.RegisterOutputType(DockerV2RepositoryOutput{})
	pulumi.RegisterOutputType(DockerV2RepositoryArrayOutput{})
	pulumi.RegisterOutputType(DockerV2RepositoryMapOutput{})
}
