// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-artifactory/sdk/v4/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Creates a virtual Chef repository.
// Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Chef+Cookbook+Repositories#ChefCookbookRepositories-VirtualChefSupermarket).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v4/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewVirtualChefRepository(ctx, "foo-chef", &artifactory.VirtualChefRepositoryArgs{
//				Description:     pulumi.String("A test virtual repo"),
//				ExcludesPattern: pulumi.String("com/google/**"),
//				IncludesPattern: pulumi.String("com/jfrog/**,cloud/jfrog/**"),
//				Key:             pulumi.String("foo-chef"),
//				Notes:           pulumi.String("Internal description"),
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
//	$ pulumi import artifactory:index/virtualChefRepository:VirtualChefRepository foo-chef foo-chef
//
// ```
type VirtualChefRepository struct {
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

// NewVirtualChefRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualChefRepository(ctx *pulumi.Context,
	name string, args *VirtualChefRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualChefRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
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

type VirtualChefRepositoryState struct {
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

func (VirtualChefRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualChefRepositoryState)(nil)).Elem()
}

type virtualChefRepositoryArgs struct {
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

// The set of arguments for constructing a VirtualChefRepository resource.
type VirtualChefRepositoryArgs struct {
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

func (i *VirtualChefRepository) ToOutput(ctx context.Context) pulumix.Output[*VirtualChefRepository] {
	return pulumix.Output[*VirtualChefRepository]{
		OutputState: i.ToVirtualChefRepositoryOutputWithContext(ctx).OutputState,
	}
}

// VirtualChefRepositoryArrayInput is an input type that accepts VirtualChefRepositoryArray and VirtualChefRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualChefRepositoryArrayInput` via:
//
//	VirtualChefRepositoryArray{ VirtualChefRepositoryArgs{...} }
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

func (i VirtualChefRepositoryArray) ToOutput(ctx context.Context) pulumix.Output[[]*VirtualChefRepository] {
	return pulumix.Output[[]*VirtualChefRepository]{
		OutputState: i.ToVirtualChefRepositoryArrayOutputWithContext(ctx).OutputState,
	}
}

// VirtualChefRepositoryMapInput is an input type that accepts VirtualChefRepositoryMap and VirtualChefRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualChefRepositoryMapInput` via:
//
//	VirtualChefRepositoryMap{ "key": VirtualChefRepositoryArgs{...} }
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

func (i VirtualChefRepositoryMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*VirtualChefRepository] {
	return pulumix.Output[map[string]*VirtualChefRepository]{
		OutputState: i.ToVirtualChefRepositoryMapOutputWithContext(ctx).OutputState,
	}
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

func (o VirtualChefRepositoryOutput) ToOutput(ctx context.Context) pulumix.Output[*VirtualChefRepository] {
	return pulumix.Output[*VirtualChefRepository]{
		OutputState: o.OutputState,
	}
}

// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
// another Artifactory instance.
func (o VirtualChefRepositoryOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualChefRepository) pulumi.BoolPtrOutput {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

// Default repository to deploy artifacts.
func (o VirtualChefRepositoryOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualChefRepository) pulumi.StringPtrOutput { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

// Public description.
func (o VirtualChefRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualChefRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o VirtualChefRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualChefRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o VirtualChefRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualChefRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
// contain spaces or special characters.
func (o VirtualChefRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualChefRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o VirtualChefRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualChefRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o VirtualChefRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualChefRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o VirtualChefRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualChefRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o VirtualChefRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualChefRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Repository layout key for the local repository
func (o VirtualChefRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualChefRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The effective list of actual repositories included in this virtual repository.
func (o VirtualChefRepositoryOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualChefRepository) pulumi.StringArrayOutput { return v.Repositories }).(pulumi.StringArrayOutput)
}

// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
func (o VirtualChefRepositoryOutput) RetrievalCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualChefRepository) pulumi.IntPtrOutput { return v.RetrievalCachePeriodSeconds }).(pulumi.IntPtrOutput)
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

func (o VirtualChefRepositoryArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*VirtualChefRepository] {
	return pulumix.Output[[]*VirtualChefRepository]{
		OutputState: o.OutputState,
	}
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

func (o VirtualChefRepositoryMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*VirtualChefRepository] {
	return pulumix.Output[map[string]*VirtualChefRepository]{
		OutputState: o.OutputState,
	}
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
