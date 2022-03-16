// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Virtual Generic Repository Resource
//
// Provides an Artifactory virtual repository resource with generic package type. This should be preferred over the original
// one-size-fits-all `VirtualRepository`.
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
// 		_, err := artifactory.NewVirtualGenericRepository(ctx, "foo-generic", &artifactory.VirtualGenericRepositoryArgs{
// 			Description:     pulumi.String("A test virtual repo"),
// 			ExcludesPattern: pulumi.String("com/google/**"),
// 			IncludesPattern: pulumi.String("com/jfrog/**,cloud/jfrog/**"),
// 			Key:             pulumi.String("foo-generic"),
// 			Notes:           pulumi.String("Internal description"),
// 			RepoLayoutRef:   pulumi.String("simple-default"),
// 			Repositories:    pulumi.StringArray{},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Virtual repositories can be imported using their name, e.g.
//
// ```sh
//  $ pulumi import artifactory:index/virtualGenericRepository:VirtualGenericRepository foo foo
// ```
type VirtualGenericRepository struct {
	pulumi.CustomResourceState

	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrOutput `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo pulumi.StringPtrOutput `pulumi:"defaultDeploymentRepo"`
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no artifacts are excluded.
	ExcludesPattern pulumi.StringPtrOutput `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrOutput `pulumi:"includesPattern"`
	// The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
	// contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
	// 'libs-release-local').
	Key pulumi.StringOutput `pulumi:"key"`
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
	PackageType pulumi.StringOutput `pulumi:"packageType"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrOutput `pulumi:"projectKey"`
	// Repository layout key for the virtual repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayOutput `pulumi:"repositories"`
	// - This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching. Default: 7200 seconds.
	RetrievalCachePeriodSeconds pulumi.IntPtrOutput `pulumi:"retrievalCachePeriodSeconds"`
}

// NewVirtualGenericRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualGenericRepository(ctx *pulumi.Context,
	name string, args *VirtualGenericRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualGenericRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource VirtualGenericRepository
	err := ctx.RegisterResource("artifactory:index/virtualGenericRepository:VirtualGenericRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualGenericRepository gets an existing VirtualGenericRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualGenericRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualGenericRepositoryState, opts ...pulumi.ResourceOption) (*VirtualGenericRepository, error) {
	var resource VirtualGenericRepository
	err := ctx.ReadResource("artifactory:index/virtualGenericRepository:VirtualGenericRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualGenericRepository resources.
type virtualGenericRepositoryState struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo *string `pulumi:"defaultDeploymentRepo"`
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
	Description *string `pulumi:"description"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
	// contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
	// 'libs-release-local').
	Key *string `pulumi:"key"`
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes *string `pulumi:"notes"`
	// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
	PackageType *string `pulumi:"packageType"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// Repository layout key for the virtual repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
	// - This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching. Default: 7200 seconds.
	RetrievalCachePeriodSeconds *int `pulumi:"retrievalCachePeriodSeconds"`
}

type VirtualGenericRepositoryState struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo pulumi.StringPtrInput
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
	Description pulumi.StringPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
	// contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
	// 'libs-release-local').
	Key pulumi.StringPtrInput
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes pulumi.StringPtrInput
	// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
	PackageType pulumi.StringPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// Repository layout key for the virtual repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
	// - This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching. Default: 7200 seconds.
	RetrievalCachePeriodSeconds pulumi.IntPtrInput
}

func (VirtualGenericRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualGenericRepositoryState)(nil)).Elem()
}

type virtualGenericRepositoryArgs struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo *string `pulumi:"defaultDeploymentRepo"`
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
	Description *string `pulumi:"description"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
	// contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
	// 'libs-release-local').
	Key string `pulumi:"key"`
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes *string `pulumi:"notes"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// Repository layout key for the virtual repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
	// - This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching. Default: 7200 seconds.
	RetrievalCachePeriodSeconds *int `pulumi:"retrievalCachePeriodSeconds"`
}

// The set of arguments for constructing a VirtualGenericRepository resource.
type VirtualGenericRepositoryArgs struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo pulumi.StringPtrInput
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
	Description pulumi.StringPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
	// contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
	// 'libs-release-local').
	Key pulumi.StringInput
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes pulumi.StringPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// Repository layout key for the virtual repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
	// - This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching. Default: 7200 seconds.
	RetrievalCachePeriodSeconds pulumi.IntPtrInput
}

func (VirtualGenericRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualGenericRepositoryArgs)(nil)).Elem()
}

type VirtualGenericRepositoryInput interface {
	pulumi.Input

	ToVirtualGenericRepositoryOutput() VirtualGenericRepositoryOutput
	ToVirtualGenericRepositoryOutputWithContext(ctx context.Context) VirtualGenericRepositoryOutput
}

func (*VirtualGenericRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualGenericRepository)(nil)).Elem()
}

func (i *VirtualGenericRepository) ToVirtualGenericRepositoryOutput() VirtualGenericRepositoryOutput {
	return i.ToVirtualGenericRepositoryOutputWithContext(context.Background())
}

func (i *VirtualGenericRepository) ToVirtualGenericRepositoryOutputWithContext(ctx context.Context) VirtualGenericRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualGenericRepositoryOutput)
}

// VirtualGenericRepositoryArrayInput is an input type that accepts VirtualGenericRepositoryArray and VirtualGenericRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualGenericRepositoryArrayInput` via:
//
//          VirtualGenericRepositoryArray{ VirtualGenericRepositoryArgs{...} }
type VirtualGenericRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualGenericRepositoryArrayOutput() VirtualGenericRepositoryArrayOutput
	ToVirtualGenericRepositoryArrayOutputWithContext(context.Context) VirtualGenericRepositoryArrayOutput
}

type VirtualGenericRepositoryArray []VirtualGenericRepositoryInput

func (VirtualGenericRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualGenericRepository)(nil)).Elem()
}

func (i VirtualGenericRepositoryArray) ToVirtualGenericRepositoryArrayOutput() VirtualGenericRepositoryArrayOutput {
	return i.ToVirtualGenericRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualGenericRepositoryArray) ToVirtualGenericRepositoryArrayOutputWithContext(ctx context.Context) VirtualGenericRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualGenericRepositoryArrayOutput)
}

// VirtualGenericRepositoryMapInput is an input type that accepts VirtualGenericRepositoryMap and VirtualGenericRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualGenericRepositoryMapInput` via:
//
//          VirtualGenericRepositoryMap{ "key": VirtualGenericRepositoryArgs{...} }
type VirtualGenericRepositoryMapInput interface {
	pulumi.Input

	ToVirtualGenericRepositoryMapOutput() VirtualGenericRepositoryMapOutput
	ToVirtualGenericRepositoryMapOutputWithContext(context.Context) VirtualGenericRepositoryMapOutput
}

type VirtualGenericRepositoryMap map[string]VirtualGenericRepositoryInput

func (VirtualGenericRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualGenericRepository)(nil)).Elem()
}

func (i VirtualGenericRepositoryMap) ToVirtualGenericRepositoryMapOutput() VirtualGenericRepositoryMapOutput {
	return i.ToVirtualGenericRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualGenericRepositoryMap) ToVirtualGenericRepositoryMapOutputWithContext(ctx context.Context) VirtualGenericRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualGenericRepositoryMapOutput)
}

type VirtualGenericRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualGenericRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualGenericRepository)(nil)).Elem()
}

func (o VirtualGenericRepositoryOutput) ToVirtualGenericRepositoryOutput() VirtualGenericRepositoryOutput {
	return o
}

func (o VirtualGenericRepositoryOutput) ToVirtualGenericRepositoryOutputWithContext(ctx context.Context) VirtualGenericRepositoryOutput {
	return o
}

type VirtualGenericRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualGenericRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualGenericRepository)(nil)).Elem()
}

func (o VirtualGenericRepositoryArrayOutput) ToVirtualGenericRepositoryArrayOutput() VirtualGenericRepositoryArrayOutput {
	return o
}

func (o VirtualGenericRepositoryArrayOutput) ToVirtualGenericRepositoryArrayOutputWithContext(ctx context.Context) VirtualGenericRepositoryArrayOutput {
	return o
}

func (o VirtualGenericRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualGenericRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualGenericRepository {
		return vs[0].([]*VirtualGenericRepository)[vs[1].(int)]
	}).(VirtualGenericRepositoryOutput)
}

type VirtualGenericRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualGenericRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualGenericRepository)(nil)).Elem()
}

func (o VirtualGenericRepositoryMapOutput) ToVirtualGenericRepositoryMapOutput() VirtualGenericRepositoryMapOutput {
	return o
}

func (o VirtualGenericRepositoryMapOutput) ToVirtualGenericRepositoryMapOutputWithContext(ctx context.Context) VirtualGenericRepositoryMapOutput {
	return o
}

func (o VirtualGenericRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualGenericRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualGenericRepository {
		return vs[0].(map[string]*VirtualGenericRepository)[vs[1].(string)]
	}).(VirtualGenericRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualGenericRepositoryInput)(nil)).Elem(), &VirtualGenericRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualGenericRepositoryArrayInput)(nil)).Elem(), VirtualGenericRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualGenericRepositoryMapInput)(nil)).Elem(), VirtualGenericRepositoryMap{})
	pulumi.RegisterOutputType(VirtualGenericRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualGenericRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualGenericRepositoryMapOutput{})
}
