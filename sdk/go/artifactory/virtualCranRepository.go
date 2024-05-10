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

// Creates a virtual Cran repository.
// Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/CRAN+Repositories#CRANRepositories-VirtualRepositories).
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
//			_, err := artifactory.NewVirtualCranRepository(ctx, "foo-cran", &artifactory.VirtualCranRepositoryArgs{
//				Key:             pulumi.String("foo-cran"),
//				Repositories:    pulumi.StringArray{},
//				Description:     pulumi.String("A test virtual repo"),
//				Notes:           pulumi.String("Internal description"),
//				IncludesPattern: pulumi.String("com/jfrog/**,cloud/jfrog/**"),
//				ExcludesPattern: pulumi.String("com/google/**"),
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
// Virtual repositories can be imported using their name, e.g.
//
// ```sh
// $ pulumi import artifactory:index/virtualCranRepository:VirtualCranRepository foo-cran foo-cran
// ```
type VirtualCranRepository struct {
	pulumi.CustomResourceState

	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrOutput `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo pulumi.StringPtrOutput `pulumi:"defaultDeploymentRepo"`
	// Public description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrOutput `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrOutput `pulumi:"includesPattern"`
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
	Key pulumi.StringOutput `pulumi:"key"`
	// Internal description.
	Notes               pulumi.StringPtrOutput   `pulumi:"notes"`
	PackageType         pulumi.StringOutput      `pulumi:"packageType"`
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrOutput `pulumi:"projectKey"`
	// Repository layout key for the virtual repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayOutput `pulumi:"repositories"`
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds pulumi.IntPtrOutput `pulumi:"retrievalCachePeriodSeconds"`
}

// NewVirtualCranRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualCranRepository(ctx *pulumi.Context,
	name string, args *VirtualCranRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualCranRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VirtualCranRepository
	err := ctx.RegisterResource("artifactory:index/virtualCranRepository:VirtualCranRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualCranRepository gets an existing VirtualCranRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualCranRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualCranRepositoryState, opts ...pulumi.ResourceOption) (*VirtualCranRepository, error) {
	var resource VirtualCranRepository
	err := ctx.ReadResource("artifactory:index/virtualCranRepository:VirtualCranRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualCranRepository resources.
type virtualCranRepositoryState struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo *string `pulumi:"defaultDeploymentRepo"`
	// Public description.
	Description *string `pulumi:"description"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
	Key *string `pulumi:"key"`
	// Internal description.
	Notes               *string  `pulumi:"notes"`
	PackageType         *string  `pulumi:"packageType"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// Repository layout key for the virtual repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds *int `pulumi:"retrievalCachePeriodSeconds"`
}

type VirtualCranRepositoryState struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo pulumi.StringPtrInput
	// Public description.
	Description pulumi.StringPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
	Key pulumi.StringPtrInput
	// Internal description.
	Notes               pulumi.StringPtrInput
	PackageType         pulumi.StringPtrInput
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// Repository layout key for the virtual repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds pulumi.IntPtrInput
}

func (VirtualCranRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualCranRepositoryState)(nil)).Elem()
}

type virtualCranRepositoryArgs struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo *string `pulumi:"defaultDeploymentRepo"`
	// Public description.
	Description *string `pulumi:"description"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
	Key string `pulumi:"key"`
	// Internal description.
	Notes               *string  `pulumi:"notes"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// Repository layout key for the virtual repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds *int `pulumi:"retrievalCachePeriodSeconds"`
}

// The set of arguments for constructing a VirtualCranRepository resource.
type VirtualCranRepositoryArgs struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo pulumi.StringPtrInput
	// Public description.
	Description pulumi.StringPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
	Key pulumi.StringInput
	// Internal description.
	Notes               pulumi.StringPtrInput
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// Repository layout key for the virtual repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds pulumi.IntPtrInput
}

func (VirtualCranRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualCranRepositoryArgs)(nil)).Elem()
}

type VirtualCranRepositoryInput interface {
	pulumi.Input

	ToVirtualCranRepositoryOutput() VirtualCranRepositoryOutput
	ToVirtualCranRepositoryOutputWithContext(ctx context.Context) VirtualCranRepositoryOutput
}

func (*VirtualCranRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualCranRepository)(nil)).Elem()
}

func (i *VirtualCranRepository) ToVirtualCranRepositoryOutput() VirtualCranRepositoryOutput {
	return i.ToVirtualCranRepositoryOutputWithContext(context.Background())
}

func (i *VirtualCranRepository) ToVirtualCranRepositoryOutputWithContext(ctx context.Context) VirtualCranRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualCranRepositoryOutput)
}

// VirtualCranRepositoryArrayInput is an input type that accepts VirtualCranRepositoryArray and VirtualCranRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualCranRepositoryArrayInput` via:
//
//	VirtualCranRepositoryArray{ VirtualCranRepositoryArgs{...} }
type VirtualCranRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualCranRepositoryArrayOutput() VirtualCranRepositoryArrayOutput
	ToVirtualCranRepositoryArrayOutputWithContext(context.Context) VirtualCranRepositoryArrayOutput
}

type VirtualCranRepositoryArray []VirtualCranRepositoryInput

func (VirtualCranRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualCranRepository)(nil)).Elem()
}

func (i VirtualCranRepositoryArray) ToVirtualCranRepositoryArrayOutput() VirtualCranRepositoryArrayOutput {
	return i.ToVirtualCranRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualCranRepositoryArray) ToVirtualCranRepositoryArrayOutputWithContext(ctx context.Context) VirtualCranRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualCranRepositoryArrayOutput)
}

// VirtualCranRepositoryMapInput is an input type that accepts VirtualCranRepositoryMap and VirtualCranRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualCranRepositoryMapInput` via:
//
//	VirtualCranRepositoryMap{ "key": VirtualCranRepositoryArgs{...} }
type VirtualCranRepositoryMapInput interface {
	pulumi.Input

	ToVirtualCranRepositoryMapOutput() VirtualCranRepositoryMapOutput
	ToVirtualCranRepositoryMapOutputWithContext(context.Context) VirtualCranRepositoryMapOutput
}

type VirtualCranRepositoryMap map[string]VirtualCranRepositoryInput

func (VirtualCranRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualCranRepository)(nil)).Elem()
}

func (i VirtualCranRepositoryMap) ToVirtualCranRepositoryMapOutput() VirtualCranRepositoryMapOutput {
	return i.ToVirtualCranRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualCranRepositoryMap) ToVirtualCranRepositoryMapOutputWithContext(ctx context.Context) VirtualCranRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualCranRepositoryMapOutput)
}

type VirtualCranRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualCranRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualCranRepository)(nil)).Elem()
}

func (o VirtualCranRepositoryOutput) ToVirtualCranRepositoryOutput() VirtualCranRepositoryOutput {
	return o
}

func (o VirtualCranRepositoryOutput) ToVirtualCranRepositoryOutputWithContext(ctx context.Context) VirtualCranRepositoryOutput {
	return o
}

// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
// another Artifactory instance.
func (o VirtualCranRepositoryOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualCranRepository) pulumi.BoolPtrOutput {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

// Default repository to deploy artifacts.
func (o VirtualCranRepositoryOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualCranRepository) pulumi.StringPtrOutput { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

// Public description.
func (o VirtualCranRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualCranRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o VirtualCranRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualCranRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o VirtualCranRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualCranRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
// contain spaces or special characters.
func (o VirtualCranRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualCranRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o VirtualCranRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualCranRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o VirtualCranRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualCranRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

func (o VirtualCranRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualCranRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o VirtualCranRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualCranRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Repository layout key for the virtual repository
func (o VirtualCranRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualCranRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The effective list of actual repositories included in this virtual repository.
func (o VirtualCranRepositoryOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualCranRepository) pulumi.StringArrayOutput { return v.Repositories }).(pulumi.StringArrayOutput)
}

// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
func (o VirtualCranRepositoryOutput) RetrievalCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualCranRepository) pulumi.IntPtrOutput { return v.RetrievalCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

type VirtualCranRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualCranRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualCranRepository)(nil)).Elem()
}

func (o VirtualCranRepositoryArrayOutput) ToVirtualCranRepositoryArrayOutput() VirtualCranRepositoryArrayOutput {
	return o
}

func (o VirtualCranRepositoryArrayOutput) ToVirtualCranRepositoryArrayOutputWithContext(ctx context.Context) VirtualCranRepositoryArrayOutput {
	return o
}

func (o VirtualCranRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualCranRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualCranRepository {
		return vs[0].([]*VirtualCranRepository)[vs[1].(int)]
	}).(VirtualCranRepositoryOutput)
}

type VirtualCranRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualCranRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualCranRepository)(nil)).Elem()
}

func (o VirtualCranRepositoryMapOutput) ToVirtualCranRepositoryMapOutput() VirtualCranRepositoryMapOutput {
	return o
}

func (o VirtualCranRepositoryMapOutput) ToVirtualCranRepositoryMapOutputWithContext(ctx context.Context) VirtualCranRepositoryMapOutput {
	return o
}

func (o VirtualCranRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualCranRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualCranRepository {
		return vs[0].(map[string]*VirtualCranRepository)[vs[1].(string)]
	}).(VirtualCranRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualCranRepositoryInput)(nil)).Elem(), &VirtualCranRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualCranRepositoryArrayInput)(nil)).Elem(), VirtualCranRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualCranRepositoryMapInput)(nil)).Elem(), VirtualCranRepositoryMap{})
	pulumi.RegisterOutputType(VirtualCranRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualCranRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualCranRepositoryMapOutput{})
}
