// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a virtual PHP Composer repository.
// Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/PHP+Composer+Repositories#PHPComposerRepositories-VirtualRepositories).
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
//			_, err := artifactory.NewVirtualComposerRepository(ctx, "foo-composer", &artifactory.VirtualComposerRepositoryArgs{
//				Description:     pulumi.String("A test virtual repo"),
//				ExcludesPattern: pulumi.String("com/google/**"),
//				IncludesPattern: pulumi.String("com/jfrog/**,cloud/jfrog/**"),
//				Key:             pulumi.String("foo-composer"),
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
//	$ pulumi import artifactory:index/virtualComposerRepository:VirtualComposerRepository foo-composer foo-composer
//
// ```
type VirtualComposerRepository struct {
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
}

// NewVirtualComposerRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualComposerRepository(ctx *pulumi.Context,
	name string, args *VirtualComposerRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualComposerRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource VirtualComposerRepository
	err := ctx.RegisterResource("artifactory:index/virtualComposerRepository:VirtualComposerRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualComposerRepository gets an existing VirtualComposerRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualComposerRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualComposerRepositoryState, opts ...pulumi.ResourceOption) (*VirtualComposerRepository, error) {
	var resource VirtualComposerRepository
	err := ctx.ReadResource("artifactory:index/virtualComposerRepository:VirtualComposerRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualComposerRepository resources.
type virtualComposerRepositoryState struct {
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
}

type VirtualComposerRepositoryState struct {
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
}

func (VirtualComposerRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualComposerRepositoryState)(nil)).Elem()
}

type virtualComposerRepositoryArgs struct {
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
}

// The set of arguments for constructing a VirtualComposerRepository resource.
type VirtualComposerRepositoryArgs struct {
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
}

func (VirtualComposerRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualComposerRepositoryArgs)(nil)).Elem()
}

type VirtualComposerRepositoryInput interface {
	pulumi.Input

	ToVirtualComposerRepositoryOutput() VirtualComposerRepositoryOutput
	ToVirtualComposerRepositoryOutputWithContext(ctx context.Context) VirtualComposerRepositoryOutput
}

func (*VirtualComposerRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualComposerRepository)(nil)).Elem()
}

func (i *VirtualComposerRepository) ToVirtualComposerRepositoryOutput() VirtualComposerRepositoryOutput {
	return i.ToVirtualComposerRepositoryOutputWithContext(context.Background())
}

func (i *VirtualComposerRepository) ToVirtualComposerRepositoryOutputWithContext(ctx context.Context) VirtualComposerRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualComposerRepositoryOutput)
}

// VirtualComposerRepositoryArrayInput is an input type that accepts VirtualComposerRepositoryArray and VirtualComposerRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualComposerRepositoryArrayInput` via:
//
//	VirtualComposerRepositoryArray{ VirtualComposerRepositoryArgs{...} }
type VirtualComposerRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualComposerRepositoryArrayOutput() VirtualComposerRepositoryArrayOutput
	ToVirtualComposerRepositoryArrayOutputWithContext(context.Context) VirtualComposerRepositoryArrayOutput
}

type VirtualComposerRepositoryArray []VirtualComposerRepositoryInput

func (VirtualComposerRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualComposerRepository)(nil)).Elem()
}

func (i VirtualComposerRepositoryArray) ToVirtualComposerRepositoryArrayOutput() VirtualComposerRepositoryArrayOutput {
	return i.ToVirtualComposerRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualComposerRepositoryArray) ToVirtualComposerRepositoryArrayOutputWithContext(ctx context.Context) VirtualComposerRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualComposerRepositoryArrayOutput)
}

// VirtualComposerRepositoryMapInput is an input type that accepts VirtualComposerRepositoryMap and VirtualComposerRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualComposerRepositoryMapInput` via:
//
//	VirtualComposerRepositoryMap{ "key": VirtualComposerRepositoryArgs{...} }
type VirtualComposerRepositoryMapInput interface {
	pulumi.Input

	ToVirtualComposerRepositoryMapOutput() VirtualComposerRepositoryMapOutput
	ToVirtualComposerRepositoryMapOutputWithContext(context.Context) VirtualComposerRepositoryMapOutput
}

type VirtualComposerRepositoryMap map[string]VirtualComposerRepositoryInput

func (VirtualComposerRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualComposerRepository)(nil)).Elem()
}

func (i VirtualComposerRepositoryMap) ToVirtualComposerRepositoryMapOutput() VirtualComposerRepositoryMapOutput {
	return i.ToVirtualComposerRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualComposerRepositoryMap) ToVirtualComposerRepositoryMapOutputWithContext(ctx context.Context) VirtualComposerRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualComposerRepositoryMapOutput)
}

type VirtualComposerRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualComposerRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualComposerRepository)(nil)).Elem()
}

func (o VirtualComposerRepositoryOutput) ToVirtualComposerRepositoryOutput() VirtualComposerRepositoryOutput {
	return o
}

func (o VirtualComposerRepositoryOutput) ToVirtualComposerRepositoryOutputWithContext(ctx context.Context) VirtualComposerRepositoryOutput {
	return o
}

// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
// another Artifactory instance.
func (o VirtualComposerRepositoryOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualComposerRepository) pulumi.BoolPtrOutput {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

// Default repository to deploy artifacts.
func (o VirtualComposerRepositoryOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualComposerRepository) pulumi.StringPtrOutput { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

// Public description.
func (o VirtualComposerRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualComposerRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o VirtualComposerRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualComposerRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o VirtualComposerRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualComposerRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
// contain spaces or special characters.
func (o VirtualComposerRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualComposerRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o VirtualComposerRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualComposerRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o VirtualComposerRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualComposerRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o VirtualComposerRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualComposerRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o VirtualComposerRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualComposerRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Repository layout key for the local repository
func (o VirtualComposerRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualComposerRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The effective list of actual repositories included in this virtual repository.
func (o VirtualComposerRepositoryOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualComposerRepository) pulumi.StringArrayOutput { return v.Repositories }).(pulumi.StringArrayOutput)
}

type VirtualComposerRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualComposerRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualComposerRepository)(nil)).Elem()
}

func (o VirtualComposerRepositoryArrayOutput) ToVirtualComposerRepositoryArrayOutput() VirtualComposerRepositoryArrayOutput {
	return o
}

func (o VirtualComposerRepositoryArrayOutput) ToVirtualComposerRepositoryArrayOutputWithContext(ctx context.Context) VirtualComposerRepositoryArrayOutput {
	return o
}

func (o VirtualComposerRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualComposerRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualComposerRepository {
		return vs[0].([]*VirtualComposerRepository)[vs[1].(int)]
	}).(VirtualComposerRepositoryOutput)
}

type VirtualComposerRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualComposerRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualComposerRepository)(nil)).Elem()
}

func (o VirtualComposerRepositoryMapOutput) ToVirtualComposerRepositoryMapOutput() VirtualComposerRepositoryMapOutput {
	return o
}

func (o VirtualComposerRepositoryMapOutput) ToVirtualComposerRepositoryMapOutputWithContext(ctx context.Context) VirtualComposerRepositoryMapOutput {
	return o
}

func (o VirtualComposerRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualComposerRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualComposerRepository {
		return vs[0].(map[string]*VirtualComposerRepository)[vs[1].(string)]
	}).(VirtualComposerRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualComposerRepositoryInput)(nil)).Elem(), &VirtualComposerRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualComposerRepositoryArrayInput)(nil)).Elem(), VirtualComposerRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualComposerRepositoryMapInput)(nil)).Elem(), VirtualComposerRepositoryMap{})
	pulumi.RegisterOutputType(VirtualComposerRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualComposerRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualComposerRepositoryMapOutput{})
}
