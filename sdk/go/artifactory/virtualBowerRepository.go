// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a virtual Bower repository.
// Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Bower+Repositories#BowerRepositories-VirtualRepositories).
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v1/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewVirtualBowerRepository(ctx, "foo-bower", &artifactory.VirtualBowerRepositoryArgs{
//				Description:                 "A test virtual repo",
//				ExcludesPattern:             "com/google/**",
//				ExternalDependenciesEnabled: false,
//				IncludesPattern:             "com/jfrog/**,cloud/jfrog/**",
//				Key:                         "foo-bower",
//				Notes:                       "Internal description",
//				Repositories:                []interface{}{},
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
//	$ pulumi import artifactory:index/virtualBowerRepository:VirtualBowerRepository foo-bower foo-bower
//
// ```
type VirtualBowerRepository struct {
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
	// When set, external dependencies are rewritten. Default value is false.
	ExternalDependenciesEnabled pulumi.BoolPtrOutput `pulumi:"externalDependenciesEnabled"`
	// An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
	ExternalDependenciesPatterns pulumi.StringArrayOutput `pulumi:"externalDependenciesPatterns"`
	// The remote repository aggregated by this virtual repository in which the external dependency will be cached.
	ExternalDependenciesRemoteRepo pulumi.StringPtrOutput `pulumi:"externalDependenciesRemoteRepo"`
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
}

// NewVirtualBowerRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualBowerRepository(ctx *pulumi.Context,
	name string, args *VirtualBowerRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualBowerRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource VirtualBowerRepository
	err := ctx.RegisterResource("artifactory:index/virtualBowerRepository:VirtualBowerRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualBowerRepository gets an existing VirtualBowerRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualBowerRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualBowerRepositoryState, opts ...pulumi.ResourceOption) (*VirtualBowerRepository, error) {
	var resource VirtualBowerRepository
	err := ctx.ReadResource("artifactory:index/virtualBowerRepository:VirtualBowerRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualBowerRepository resources.
type virtualBowerRepositoryState struct {
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
	// When set, external dependencies are rewritten. Default value is false.
	ExternalDependenciesEnabled *bool `pulumi:"externalDependenciesEnabled"`
	// An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
	ExternalDependenciesPatterns []string `pulumi:"externalDependenciesPatterns"`
	// The remote repository aggregated by this virtual repository in which the external dependency will be cached.
	ExternalDependenciesRemoteRepo *string `pulumi:"externalDependenciesRemoteRepo"`
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
}

type VirtualBowerRepositoryState struct {
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
	// When set, external dependencies are rewritten. Default value is false.
	ExternalDependenciesEnabled pulumi.BoolPtrInput
	// An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
	ExternalDependenciesPatterns pulumi.StringArrayInput
	// The remote repository aggregated by this virtual repository in which the external dependency will be cached.
	ExternalDependenciesRemoteRepo pulumi.StringPtrInput
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
}

func (VirtualBowerRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualBowerRepositoryState)(nil)).Elem()
}

type virtualBowerRepositoryArgs struct {
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
	// When set, external dependencies are rewritten. Default value is false.
	ExternalDependenciesEnabled *bool `pulumi:"externalDependenciesEnabled"`
	// An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
	ExternalDependenciesPatterns []string `pulumi:"externalDependenciesPatterns"`
	// The remote repository aggregated by this virtual repository in which the external dependency will be cached.
	ExternalDependenciesRemoteRepo *string `pulumi:"externalDependenciesRemoteRepo"`
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
}

// The set of arguments for constructing a VirtualBowerRepository resource.
type VirtualBowerRepositoryArgs struct {
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
	// When set, external dependencies are rewritten. Default value is false.
	ExternalDependenciesEnabled pulumi.BoolPtrInput
	// An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
	ExternalDependenciesPatterns pulumi.StringArrayInput
	// The remote repository aggregated by this virtual repository in which the external dependency will be cached.
	ExternalDependenciesRemoteRepo pulumi.StringPtrInput
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
}

func (VirtualBowerRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualBowerRepositoryArgs)(nil)).Elem()
}

type VirtualBowerRepositoryInput interface {
	pulumi.Input

	ToVirtualBowerRepositoryOutput() VirtualBowerRepositoryOutput
	ToVirtualBowerRepositoryOutputWithContext(ctx context.Context) VirtualBowerRepositoryOutput
}

func (*VirtualBowerRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualBowerRepository)(nil)).Elem()
}

func (i *VirtualBowerRepository) ToVirtualBowerRepositoryOutput() VirtualBowerRepositoryOutput {
	return i.ToVirtualBowerRepositoryOutputWithContext(context.Background())
}

func (i *VirtualBowerRepository) ToVirtualBowerRepositoryOutputWithContext(ctx context.Context) VirtualBowerRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualBowerRepositoryOutput)
}

// VirtualBowerRepositoryArrayInput is an input type that accepts VirtualBowerRepositoryArray and VirtualBowerRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualBowerRepositoryArrayInput` via:
//
//	VirtualBowerRepositoryArray{ VirtualBowerRepositoryArgs{...} }
type VirtualBowerRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualBowerRepositoryArrayOutput() VirtualBowerRepositoryArrayOutput
	ToVirtualBowerRepositoryArrayOutputWithContext(context.Context) VirtualBowerRepositoryArrayOutput
}

type VirtualBowerRepositoryArray []VirtualBowerRepositoryInput

func (VirtualBowerRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualBowerRepository)(nil)).Elem()
}

func (i VirtualBowerRepositoryArray) ToVirtualBowerRepositoryArrayOutput() VirtualBowerRepositoryArrayOutput {
	return i.ToVirtualBowerRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualBowerRepositoryArray) ToVirtualBowerRepositoryArrayOutputWithContext(ctx context.Context) VirtualBowerRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualBowerRepositoryArrayOutput)
}

// VirtualBowerRepositoryMapInput is an input type that accepts VirtualBowerRepositoryMap and VirtualBowerRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualBowerRepositoryMapInput` via:
//
//	VirtualBowerRepositoryMap{ "key": VirtualBowerRepositoryArgs{...} }
type VirtualBowerRepositoryMapInput interface {
	pulumi.Input

	ToVirtualBowerRepositoryMapOutput() VirtualBowerRepositoryMapOutput
	ToVirtualBowerRepositoryMapOutputWithContext(context.Context) VirtualBowerRepositoryMapOutput
}

type VirtualBowerRepositoryMap map[string]VirtualBowerRepositoryInput

func (VirtualBowerRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualBowerRepository)(nil)).Elem()
}

func (i VirtualBowerRepositoryMap) ToVirtualBowerRepositoryMapOutput() VirtualBowerRepositoryMapOutput {
	return i.ToVirtualBowerRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualBowerRepositoryMap) ToVirtualBowerRepositoryMapOutputWithContext(ctx context.Context) VirtualBowerRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualBowerRepositoryMapOutput)
}

type VirtualBowerRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualBowerRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualBowerRepository)(nil)).Elem()
}

func (o VirtualBowerRepositoryOutput) ToVirtualBowerRepositoryOutput() VirtualBowerRepositoryOutput {
	return o
}

func (o VirtualBowerRepositoryOutput) ToVirtualBowerRepositoryOutputWithContext(ctx context.Context) VirtualBowerRepositoryOutput {
	return o
}

// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
// another Artifactory instance.
func (o VirtualBowerRepositoryOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualBowerRepository) pulumi.BoolPtrOutput {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

// Default repository to deploy artifacts.
func (o VirtualBowerRepositoryOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualBowerRepository) pulumi.StringPtrOutput { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

// Public description.
func (o VirtualBowerRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualBowerRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o VirtualBowerRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualBowerRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// When set, external dependencies are rewritten. Default value is false.
func (o VirtualBowerRepositoryOutput) ExternalDependenciesEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualBowerRepository) pulumi.BoolPtrOutput { return v.ExternalDependenciesEnabled }).(pulumi.BoolPtrOutput)
}

// An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
func (o VirtualBowerRepositoryOutput) ExternalDependenciesPatterns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualBowerRepository) pulumi.StringArrayOutput { return v.ExternalDependenciesPatterns }).(pulumi.StringArrayOutput)
}

// The remote repository aggregated by this virtual repository in which the external dependency will be cached.
func (o VirtualBowerRepositoryOutput) ExternalDependenciesRemoteRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualBowerRepository) pulumi.StringPtrOutput { return v.ExternalDependenciesRemoteRepo }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o VirtualBowerRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualBowerRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
// contain spaces or special characters.
func (o VirtualBowerRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualBowerRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o VirtualBowerRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualBowerRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o VirtualBowerRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualBowerRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o VirtualBowerRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualBowerRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o VirtualBowerRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualBowerRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Repository layout key for the local repository
func (o VirtualBowerRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualBowerRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The effective list of actual repositories included in this virtual repository.
func (o VirtualBowerRepositoryOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualBowerRepository) pulumi.StringArrayOutput { return v.Repositories }).(pulumi.StringArrayOutput)
}

type VirtualBowerRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualBowerRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualBowerRepository)(nil)).Elem()
}

func (o VirtualBowerRepositoryArrayOutput) ToVirtualBowerRepositoryArrayOutput() VirtualBowerRepositoryArrayOutput {
	return o
}

func (o VirtualBowerRepositoryArrayOutput) ToVirtualBowerRepositoryArrayOutputWithContext(ctx context.Context) VirtualBowerRepositoryArrayOutput {
	return o
}

func (o VirtualBowerRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualBowerRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualBowerRepository {
		return vs[0].([]*VirtualBowerRepository)[vs[1].(int)]
	}).(VirtualBowerRepositoryOutput)
}

type VirtualBowerRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualBowerRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualBowerRepository)(nil)).Elem()
}

func (o VirtualBowerRepositoryMapOutput) ToVirtualBowerRepositoryMapOutput() VirtualBowerRepositoryMapOutput {
	return o
}

func (o VirtualBowerRepositoryMapOutput) ToVirtualBowerRepositoryMapOutputWithContext(ctx context.Context) VirtualBowerRepositoryMapOutput {
	return o
}

func (o VirtualBowerRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualBowerRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualBowerRepository {
		return vs[0].(map[string]*VirtualBowerRepository)[vs[1].(string)]
	}).(VirtualBowerRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualBowerRepositoryInput)(nil)).Elem(), &VirtualBowerRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualBowerRepositoryArrayInput)(nil)).Elem(), VirtualBowerRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualBowerRepositoryMapInput)(nil)).Elem(), VirtualBowerRepositoryMap{})
	pulumi.RegisterOutputType(VirtualBowerRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualBowerRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualBowerRepositoryMapOutput{})
}
