// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a virtual Conan repository.
// Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Conan+Repositories#ConanRepositories-VirtualRepositories).
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
//			_, err := artifactory.NewVirtualConanRepository(ctx, "foo-conan", &artifactory.VirtualConanRepositoryArgs{
//				Description:     pulumi.String("A test virtual repo"),
//				ExcludesPattern: pulumi.String("com/google/**"),
//				IncludesPattern: pulumi.String("com/jfrog/**,cloud/jfrog/**"),
//				Key:             pulumi.String("foo-conan"),
//				Notes:           pulumi.String("Internal description"),
//				RepoLayoutRef:   pulumi.String("conan-default"),
//				Repositories:    pulumi.StringArray{},
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
//
//	$ pulumi import artifactory:index/virtualConanRepository:VirtualConanRepository foo-conan foo-conan
//
// ```
type VirtualConanRepository struct {
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
	Notes       pulumi.StringPtrOutput `pulumi:"notes"`
	PackageType pulumi.StringOutput    `pulumi:"packageType"`
	// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
	// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
	// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
	// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrOutput `pulumi:"projectKey"`
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayOutput `pulumi:"repositories"`
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds pulumi.IntPtrOutput `pulumi:"retrievalCachePeriodSeconds"`
}

// NewVirtualConanRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualConanRepository(ctx *pulumi.Context,
	name string, args *VirtualConanRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualConanRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource VirtualConanRepository
	err := ctx.RegisterResource("artifactory:index/virtualConanRepository:VirtualConanRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualConanRepository gets an existing VirtualConanRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualConanRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualConanRepositoryState, opts ...pulumi.ResourceOption) (*VirtualConanRepository, error) {
	var resource VirtualConanRepository
	err := ctx.ReadResource("artifactory:index/virtualConanRepository:VirtualConanRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualConanRepository resources.
type virtualConanRepositoryState struct {
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
	Notes       *string `pulumi:"notes"`
	PackageType *string `pulumi:"packageType"`
	// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
	// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
	// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
	// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds *int `pulumi:"retrievalCachePeriodSeconds"`
}

type VirtualConanRepositoryState struct {
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
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
	// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
	// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
	// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds pulumi.IntPtrInput
}

func (VirtualConanRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualConanRepositoryState)(nil)).Elem()
}

type virtualConanRepositoryArgs struct {
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
	Notes *string `pulumi:"notes"`
	// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
	// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
	// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
	// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds *int `pulumi:"retrievalCachePeriodSeconds"`
}

// The set of arguments for constructing a VirtualConanRepository resource.
type VirtualConanRepositoryArgs struct {
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
	Notes pulumi.StringPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
	// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
	// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
	// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds pulumi.IntPtrInput
}

func (VirtualConanRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualConanRepositoryArgs)(nil)).Elem()
}

type VirtualConanRepositoryInput interface {
	pulumi.Input

	ToVirtualConanRepositoryOutput() VirtualConanRepositoryOutput
	ToVirtualConanRepositoryOutputWithContext(ctx context.Context) VirtualConanRepositoryOutput
}

func (*VirtualConanRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualConanRepository)(nil)).Elem()
}

func (i *VirtualConanRepository) ToVirtualConanRepositoryOutput() VirtualConanRepositoryOutput {
	return i.ToVirtualConanRepositoryOutputWithContext(context.Background())
}

func (i *VirtualConanRepository) ToVirtualConanRepositoryOutputWithContext(ctx context.Context) VirtualConanRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualConanRepositoryOutput)
}

// VirtualConanRepositoryArrayInput is an input type that accepts VirtualConanRepositoryArray and VirtualConanRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualConanRepositoryArrayInput` via:
//
//	VirtualConanRepositoryArray{ VirtualConanRepositoryArgs{...} }
type VirtualConanRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualConanRepositoryArrayOutput() VirtualConanRepositoryArrayOutput
	ToVirtualConanRepositoryArrayOutputWithContext(context.Context) VirtualConanRepositoryArrayOutput
}

type VirtualConanRepositoryArray []VirtualConanRepositoryInput

func (VirtualConanRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualConanRepository)(nil)).Elem()
}

func (i VirtualConanRepositoryArray) ToVirtualConanRepositoryArrayOutput() VirtualConanRepositoryArrayOutput {
	return i.ToVirtualConanRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualConanRepositoryArray) ToVirtualConanRepositoryArrayOutputWithContext(ctx context.Context) VirtualConanRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualConanRepositoryArrayOutput)
}

// VirtualConanRepositoryMapInput is an input type that accepts VirtualConanRepositoryMap and VirtualConanRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualConanRepositoryMapInput` via:
//
//	VirtualConanRepositoryMap{ "key": VirtualConanRepositoryArgs{...} }
type VirtualConanRepositoryMapInput interface {
	pulumi.Input

	ToVirtualConanRepositoryMapOutput() VirtualConanRepositoryMapOutput
	ToVirtualConanRepositoryMapOutputWithContext(context.Context) VirtualConanRepositoryMapOutput
}

type VirtualConanRepositoryMap map[string]VirtualConanRepositoryInput

func (VirtualConanRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualConanRepository)(nil)).Elem()
}

func (i VirtualConanRepositoryMap) ToVirtualConanRepositoryMapOutput() VirtualConanRepositoryMapOutput {
	return i.ToVirtualConanRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualConanRepositoryMap) ToVirtualConanRepositoryMapOutputWithContext(ctx context.Context) VirtualConanRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualConanRepositoryMapOutput)
}

type VirtualConanRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualConanRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualConanRepository)(nil)).Elem()
}

func (o VirtualConanRepositoryOutput) ToVirtualConanRepositoryOutput() VirtualConanRepositoryOutput {
	return o
}

func (o VirtualConanRepositoryOutput) ToVirtualConanRepositoryOutputWithContext(ctx context.Context) VirtualConanRepositoryOutput {
	return o
}

// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
// another Artifactory instance.
func (o VirtualConanRepositoryOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualConanRepository) pulumi.BoolPtrOutput {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

// Default repository to deploy artifacts.
func (o VirtualConanRepositoryOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualConanRepository) pulumi.StringPtrOutput { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

// Public description.
func (o VirtualConanRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualConanRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o VirtualConanRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualConanRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o VirtualConanRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualConanRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
// contain spaces or special characters.
func (o VirtualConanRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualConanRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o VirtualConanRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualConanRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o VirtualConanRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualConanRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o VirtualConanRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualConanRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o VirtualConanRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualConanRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Repository layout key for the local repository
func (o VirtualConanRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualConanRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The effective list of actual repositories included in this virtual repository.
func (o VirtualConanRepositoryOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualConanRepository) pulumi.StringArrayOutput { return v.Repositories }).(pulumi.StringArrayOutput)
}

// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
func (o VirtualConanRepositoryOutput) RetrievalCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualConanRepository) pulumi.IntPtrOutput { return v.RetrievalCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

type VirtualConanRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualConanRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualConanRepository)(nil)).Elem()
}

func (o VirtualConanRepositoryArrayOutput) ToVirtualConanRepositoryArrayOutput() VirtualConanRepositoryArrayOutput {
	return o
}

func (o VirtualConanRepositoryArrayOutput) ToVirtualConanRepositoryArrayOutputWithContext(ctx context.Context) VirtualConanRepositoryArrayOutput {
	return o
}

func (o VirtualConanRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualConanRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualConanRepository {
		return vs[0].([]*VirtualConanRepository)[vs[1].(int)]
	}).(VirtualConanRepositoryOutput)
}

type VirtualConanRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualConanRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualConanRepository)(nil)).Elem()
}

func (o VirtualConanRepositoryMapOutput) ToVirtualConanRepositoryMapOutput() VirtualConanRepositoryMapOutput {
	return o
}

func (o VirtualConanRepositoryMapOutput) ToVirtualConanRepositoryMapOutputWithContext(ctx context.Context) VirtualConanRepositoryMapOutput {
	return o
}

func (o VirtualConanRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualConanRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualConanRepository {
		return vs[0].(map[string]*VirtualConanRepository)[vs[1].(string)]
	}).(VirtualConanRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualConanRepositoryInput)(nil)).Elem(), &VirtualConanRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualConanRepositoryArrayInput)(nil)).Elem(), VirtualConanRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualConanRepositoryMapInput)(nil)).Elem(), VirtualConanRepositoryMap{})
	pulumi.RegisterOutputType(VirtualConanRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualConanRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualConanRepositoryMapOutput{})
}
