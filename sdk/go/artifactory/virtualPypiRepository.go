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

// Creates a virtual Pypi repository.
// Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/PyPI+Repositories#PyPIRepositories-VirtualRepositories).
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := artifactory.NewVirtualPypiRepository(ctx, "foo-pypi", &artifactory.VirtualPypiRepositoryArgs{
//				Key:             pulumi.String("foo-pypi"),
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
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Virtual repositories can be imported using their name, e.g.
//
// ```sh
// $ pulumi import artifactory:index/virtualPypiRepository:VirtualPypiRepository foo-pypi foo-pypi
// ```
type VirtualPypiRepository struct {
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
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrOutput `pulumi:"projectKey"`
	// Repository layout key for the virtual repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayOutput `pulumi:"repositories"`
}

// NewVirtualPypiRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualPypiRepository(ctx *pulumi.Context,
	name string, args *VirtualPypiRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualPypiRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VirtualPypiRepository
	err := ctx.RegisterResource("artifactory:index/virtualPypiRepository:VirtualPypiRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualPypiRepository gets an existing VirtualPypiRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualPypiRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualPypiRepositoryState, opts ...pulumi.ResourceOption) (*VirtualPypiRepository, error) {
	var resource VirtualPypiRepository
	err := ctx.ReadResource("artifactory:index/virtualPypiRepository:VirtualPypiRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualPypiRepository resources.
type virtualPypiRepositoryState struct {
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
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// Repository layout key for the virtual repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
}

type VirtualPypiRepositoryState struct {
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
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// Repository layout key for the virtual repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
}

func (VirtualPypiRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualPypiRepositoryState)(nil)).Elem()
}

type virtualPypiRepositoryArgs struct {
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
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// Repository layout key for the virtual repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
}

// The set of arguments for constructing a VirtualPypiRepository resource.
type VirtualPypiRepositoryArgs struct {
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
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// Repository layout key for the virtual repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
}

func (VirtualPypiRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualPypiRepositoryArgs)(nil)).Elem()
}

type VirtualPypiRepositoryInput interface {
	pulumi.Input

	ToVirtualPypiRepositoryOutput() VirtualPypiRepositoryOutput
	ToVirtualPypiRepositoryOutputWithContext(ctx context.Context) VirtualPypiRepositoryOutput
}

func (*VirtualPypiRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualPypiRepository)(nil)).Elem()
}

func (i *VirtualPypiRepository) ToVirtualPypiRepositoryOutput() VirtualPypiRepositoryOutput {
	return i.ToVirtualPypiRepositoryOutputWithContext(context.Background())
}

func (i *VirtualPypiRepository) ToVirtualPypiRepositoryOutputWithContext(ctx context.Context) VirtualPypiRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualPypiRepositoryOutput)
}

// VirtualPypiRepositoryArrayInput is an input type that accepts VirtualPypiRepositoryArray and VirtualPypiRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualPypiRepositoryArrayInput` via:
//
//	VirtualPypiRepositoryArray{ VirtualPypiRepositoryArgs{...} }
type VirtualPypiRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualPypiRepositoryArrayOutput() VirtualPypiRepositoryArrayOutput
	ToVirtualPypiRepositoryArrayOutputWithContext(context.Context) VirtualPypiRepositoryArrayOutput
}

type VirtualPypiRepositoryArray []VirtualPypiRepositoryInput

func (VirtualPypiRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualPypiRepository)(nil)).Elem()
}

func (i VirtualPypiRepositoryArray) ToVirtualPypiRepositoryArrayOutput() VirtualPypiRepositoryArrayOutput {
	return i.ToVirtualPypiRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualPypiRepositoryArray) ToVirtualPypiRepositoryArrayOutputWithContext(ctx context.Context) VirtualPypiRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualPypiRepositoryArrayOutput)
}

// VirtualPypiRepositoryMapInput is an input type that accepts VirtualPypiRepositoryMap and VirtualPypiRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualPypiRepositoryMapInput` via:
//
//	VirtualPypiRepositoryMap{ "key": VirtualPypiRepositoryArgs{...} }
type VirtualPypiRepositoryMapInput interface {
	pulumi.Input

	ToVirtualPypiRepositoryMapOutput() VirtualPypiRepositoryMapOutput
	ToVirtualPypiRepositoryMapOutputWithContext(context.Context) VirtualPypiRepositoryMapOutput
}

type VirtualPypiRepositoryMap map[string]VirtualPypiRepositoryInput

func (VirtualPypiRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualPypiRepository)(nil)).Elem()
}

func (i VirtualPypiRepositoryMap) ToVirtualPypiRepositoryMapOutput() VirtualPypiRepositoryMapOutput {
	return i.ToVirtualPypiRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualPypiRepositoryMap) ToVirtualPypiRepositoryMapOutputWithContext(ctx context.Context) VirtualPypiRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualPypiRepositoryMapOutput)
}

type VirtualPypiRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualPypiRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualPypiRepository)(nil)).Elem()
}

func (o VirtualPypiRepositoryOutput) ToVirtualPypiRepositoryOutput() VirtualPypiRepositoryOutput {
	return o
}

func (o VirtualPypiRepositoryOutput) ToVirtualPypiRepositoryOutputWithContext(ctx context.Context) VirtualPypiRepositoryOutput {
	return o
}

// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
// another Artifactory instance.
func (o VirtualPypiRepositoryOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualPypiRepository) pulumi.BoolPtrOutput {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

// Default repository to deploy artifacts.
func (o VirtualPypiRepositoryOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualPypiRepository) pulumi.StringPtrOutput { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

// Public description.
func (o VirtualPypiRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualPypiRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o VirtualPypiRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualPypiRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o VirtualPypiRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualPypiRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
// contain spaces or special characters.
func (o VirtualPypiRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualPypiRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o VirtualPypiRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualPypiRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o VirtualPypiRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualPypiRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o VirtualPypiRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualPypiRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o VirtualPypiRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualPypiRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Repository layout key for the virtual repository
func (o VirtualPypiRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualPypiRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The effective list of actual repositories included in this virtual repository.
func (o VirtualPypiRepositoryOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualPypiRepository) pulumi.StringArrayOutput { return v.Repositories }).(pulumi.StringArrayOutput)
}

type VirtualPypiRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualPypiRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualPypiRepository)(nil)).Elem()
}

func (o VirtualPypiRepositoryArrayOutput) ToVirtualPypiRepositoryArrayOutput() VirtualPypiRepositoryArrayOutput {
	return o
}

func (o VirtualPypiRepositoryArrayOutput) ToVirtualPypiRepositoryArrayOutputWithContext(ctx context.Context) VirtualPypiRepositoryArrayOutput {
	return o
}

func (o VirtualPypiRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualPypiRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualPypiRepository {
		return vs[0].([]*VirtualPypiRepository)[vs[1].(int)]
	}).(VirtualPypiRepositoryOutput)
}

type VirtualPypiRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualPypiRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualPypiRepository)(nil)).Elem()
}

func (o VirtualPypiRepositoryMapOutput) ToVirtualPypiRepositoryMapOutput() VirtualPypiRepositoryMapOutput {
	return o
}

func (o VirtualPypiRepositoryMapOutput) ToVirtualPypiRepositoryMapOutputWithContext(ctx context.Context) VirtualPypiRepositoryMapOutput {
	return o
}

func (o VirtualPypiRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualPypiRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualPypiRepository {
		return vs[0].(map[string]*VirtualPypiRepository)[vs[1].(string)]
	}).(VirtualPypiRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualPypiRepositoryInput)(nil)).Elem(), &VirtualPypiRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualPypiRepositoryArrayInput)(nil)).Elem(), VirtualPypiRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualPypiRepositoryMapInput)(nil)).Elem(), VirtualPypiRepositoryMap{})
	pulumi.RegisterOutputType(VirtualPypiRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualPypiRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualPypiRepositoryMapOutput{})
}
