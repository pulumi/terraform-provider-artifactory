// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualChefRepository struct {
	pulumi.CustomResourceState

	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrOutput `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo pulumi.StringPtrOutput `pulumi:"defaultDeploymentRepo"`
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrOutput `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
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
	// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
	// repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrOutput `pulumi:"projectKey"`
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayOutput `pulumi:"repositories"`
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
	// repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds pulumi.IntPtrOutput `pulumi:"retrievalCachePeriodSeconds"`
}

// NewVirtualChefRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualChefRepository(ctx *pulumi.Context,
	name string, args *VirtualChefRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualChefRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource VirtualChefRepository
	err := ctx.RegisterResource("artifactory:index/virtualChefRepository:VirtualChefRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualChefRepository gets an existing VirtualChefRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualChefRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualChefRepositoryState, opts ...pulumi.ResourceOption) (*VirtualChefRepository, error) {
	var resource VirtualChefRepository
	err := ctx.ReadResource("artifactory:index/virtualChefRepository:VirtualChefRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualChefRepository resources.
type virtualChefRepositoryState struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo *string `pulumi:"defaultDeploymentRepo"`
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
	Description *string `pulumi:"description"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
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
	// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
	// repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
	// repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds *int `pulumi:"retrievalCachePeriodSeconds"`
}

type VirtualChefRepositoryState struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo pulumi.StringPtrInput
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
	Description pulumi.StringPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
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
	// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
	// repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
	// repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds pulumi.IntPtrInput
}

func (VirtualChefRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualChefRepositoryState)(nil)).Elem()
}

type virtualChefRepositoryArgs struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo *string `pulumi:"defaultDeploymentRepo"`
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
	Description *string `pulumi:"description"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
	// contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
	// 'libs-release-local').
	Key string `pulumi:"key"`
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes *string `pulumi:"notes"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
	// repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
	// repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds *int `pulumi:"retrievalCachePeriodSeconds"`
}

// The set of arguments for constructing a VirtualChefRepository resource.
type VirtualChefRepositoryArgs struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo pulumi.StringPtrInput
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
	Description pulumi.StringPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
	// contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
	// 'libs-release-local').
	Key pulumi.StringInput
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes pulumi.StringPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
	// repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
	// repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds pulumi.IntPtrInput
}

func (VirtualChefRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualChefRepositoryArgs)(nil)).Elem()
}

type VirtualChefRepositoryInput interface {
	pulumi.Input

	ToVirtualChefRepositoryOutput() VirtualChefRepositoryOutput
	ToVirtualChefRepositoryOutputWithContext(ctx context.Context) VirtualChefRepositoryOutput
}

func (*VirtualChefRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualChefRepository)(nil)).Elem()
}

func (i *VirtualChefRepository) ToVirtualChefRepositoryOutput() VirtualChefRepositoryOutput {
	return i.ToVirtualChefRepositoryOutputWithContext(context.Background())
}

func (i *VirtualChefRepository) ToVirtualChefRepositoryOutputWithContext(ctx context.Context) VirtualChefRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualChefRepositoryOutput)
}

// VirtualChefRepositoryArrayInput is an input type that accepts VirtualChefRepositoryArray and VirtualChefRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualChefRepositoryArrayInput` via:
//
//          VirtualChefRepositoryArray{ VirtualChefRepositoryArgs{...} }
type VirtualChefRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualChefRepositoryArrayOutput() VirtualChefRepositoryArrayOutput
	ToVirtualChefRepositoryArrayOutputWithContext(context.Context) VirtualChefRepositoryArrayOutput
}

type VirtualChefRepositoryArray []VirtualChefRepositoryInput

func (VirtualChefRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualChefRepository)(nil)).Elem()
}

func (i VirtualChefRepositoryArray) ToVirtualChefRepositoryArrayOutput() VirtualChefRepositoryArrayOutput {
	return i.ToVirtualChefRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualChefRepositoryArray) ToVirtualChefRepositoryArrayOutputWithContext(ctx context.Context) VirtualChefRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualChefRepositoryArrayOutput)
}

// VirtualChefRepositoryMapInput is an input type that accepts VirtualChefRepositoryMap and VirtualChefRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualChefRepositoryMapInput` via:
//
//          VirtualChefRepositoryMap{ "key": VirtualChefRepositoryArgs{...} }
type VirtualChefRepositoryMapInput interface {
	pulumi.Input

	ToVirtualChefRepositoryMapOutput() VirtualChefRepositoryMapOutput
	ToVirtualChefRepositoryMapOutputWithContext(context.Context) VirtualChefRepositoryMapOutput
}

type VirtualChefRepositoryMap map[string]VirtualChefRepositoryInput

func (VirtualChefRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualChefRepository)(nil)).Elem()
}

func (i VirtualChefRepositoryMap) ToVirtualChefRepositoryMapOutput() VirtualChefRepositoryMapOutput {
	return i.ToVirtualChefRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualChefRepositoryMap) ToVirtualChefRepositoryMapOutputWithContext(ctx context.Context) VirtualChefRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualChefRepositoryMapOutput)
}

type VirtualChefRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualChefRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualChefRepository)(nil)).Elem()
}

func (o VirtualChefRepositoryOutput) ToVirtualChefRepositoryOutput() VirtualChefRepositoryOutput {
	return o
}

func (o VirtualChefRepositoryOutput) ToVirtualChefRepositoryOutputWithContext(ctx context.Context) VirtualChefRepositoryOutput {
	return o
}

type VirtualChefRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualChefRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualChefRepository)(nil)).Elem()
}

func (o VirtualChefRepositoryArrayOutput) ToVirtualChefRepositoryArrayOutput() VirtualChefRepositoryArrayOutput {
	return o
}

func (o VirtualChefRepositoryArrayOutput) ToVirtualChefRepositoryArrayOutputWithContext(ctx context.Context) VirtualChefRepositoryArrayOutput {
	return o
}

func (o VirtualChefRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualChefRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualChefRepository {
		return vs[0].([]*VirtualChefRepository)[vs[1].(int)]
	}).(VirtualChefRepositoryOutput)
}

type VirtualChefRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualChefRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualChefRepository)(nil)).Elem()
}

func (o VirtualChefRepositoryMapOutput) ToVirtualChefRepositoryMapOutput() VirtualChefRepositoryMapOutput {
	return o
}

func (o VirtualChefRepositoryMapOutput) ToVirtualChefRepositoryMapOutputWithContext(ctx context.Context) VirtualChefRepositoryMapOutput {
	return o
}

func (o VirtualChefRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualChefRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualChefRepository {
		return vs[0].(map[string]*VirtualChefRepository)[vs[1].(string)]
	}).(VirtualChefRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualChefRepositoryInput)(nil)).Elem(), &VirtualChefRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualChefRepositoryArrayInput)(nil)).Elem(), VirtualChefRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualChefRepositoryMapInput)(nil)).Elem(), VirtualChefRepositoryMap{})
	pulumi.RegisterOutputType(VirtualChefRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualChefRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualChefRepositoryMapOutput{})
}
