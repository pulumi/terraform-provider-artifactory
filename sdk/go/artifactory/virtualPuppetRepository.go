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

// Creates a virtual Puppet repository.
// Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Puppet+Repositories#PuppetRepositories-VirtualPuppetRepository).
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
//			_, err := artifactory.NewVirtualPuppetRepository(ctx, "foo-puppet", &artifactory.VirtualPuppetRepositoryArgs{
//				Description:     pulumi.String("A test virtual repo"),
//				ExcludesPattern: pulumi.String("com/google/**"),
//				IncludesPattern: pulumi.String("com/jfrog/**,cloud/jfrog/**"),
//				Key:             pulumi.String("foo-puppet"),
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
//	$ pulumi import artifactory:index/virtualPuppetRepository:VirtualPuppetRepository foo-puppet foo-puppet
//
// ```
type VirtualPuppetRepository struct {
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

// NewVirtualPuppetRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualPuppetRepository(ctx *pulumi.Context,
	name string, args *VirtualPuppetRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualPuppetRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VirtualPuppetRepository
	err := ctx.RegisterResource("artifactory:index/virtualPuppetRepository:VirtualPuppetRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualPuppetRepository gets an existing VirtualPuppetRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualPuppetRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualPuppetRepositoryState, opts ...pulumi.ResourceOption) (*VirtualPuppetRepository, error) {
	var resource VirtualPuppetRepository
	err := ctx.ReadResource("artifactory:index/virtualPuppetRepository:VirtualPuppetRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualPuppetRepository resources.
type virtualPuppetRepositoryState struct {
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

type VirtualPuppetRepositoryState struct {
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

func (VirtualPuppetRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualPuppetRepositoryState)(nil)).Elem()
}

type virtualPuppetRepositoryArgs struct {
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

// The set of arguments for constructing a VirtualPuppetRepository resource.
type VirtualPuppetRepositoryArgs struct {
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

func (VirtualPuppetRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualPuppetRepositoryArgs)(nil)).Elem()
}

type VirtualPuppetRepositoryInput interface {
	pulumi.Input

	ToVirtualPuppetRepositoryOutput() VirtualPuppetRepositoryOutput
	ToVirtualPuppetRepositoryOutputWithContext(ctx context.Context) VirtualPuppetRepositoryOutput
}

func (*VirtualPuppetRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualPuppetRepository)(nil)).Elem()
}

func (i *VirtualPuppetRepository) ToVirtualPuppetRepositoryOutput() VirtualPuppetRepositoryOutput {
	return i.ToVirtualPuppetRepositoryOutputWithContext(context.Background())
}

func (i *VirtualPuppetRepository) ToVirtualPuppetRepositoryOutputWithContext(ctx context.Context) VirtualPuppetRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualPuppetRepositoryOutput)
}

func (i *VirtualPuppetRepository) ToOutput(ctx context.Context) pulumix.Output[*VirtualPuppetRepository] {
	return pulumix.Output[*VirtualPuppetRepository]{
		OutputState: i.ToVirtualPuppetRepositoryOutputWithContext(ctx).OutputState,
	}
}

// VirtualPuppetRepositoryArrayInput is an input type that accepts VirtualPuppetRepositoryArray and VirtualPuppetRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualPuppetRepositoryArrayInput` via:
//
//	VirtualPuppetRepositoryArray{ VirtualPuppetRepositoryArgs{...} }
type VirtualPuppetRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualPuppetRepositoryArrayOutput() VirtualPuppetRepositoryArrayOutput
	ToVirtualPuppetRepositoryArrayOutputWithContext(context.Context) VirtualPuppetRepositoryArrayOutput
}

type VirtualPuppetRepositoryArray []VirtualPuppetRepositoryInput

func (VirtualPuppetRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualPuppetRepository)(nil)).Elem()
}

func (i VirtualPuppetRepositoryArray) ToVirtualPuppetRepositoryArrayOutput() VirtualPuppetRepositoryArrayOutput {
	return i.ToVirtualPuppetRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualPuppetRepositoryArray) ToVirtualPuppetRepositoryArrayOutputWithContext(ctx context.Context) VirtualPuppetRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualPuppetRepositoryArrayOutput)
}

func (i VirtualPuppetRepositoryArray) ToOutput(ctx context.Context) pulumix.Output[[]*VirtualPuppetRepository] {
	return pulumix.Output[[]*VirtualPuppetRepository]{
		OutputState: i.ToVirtualPuppetRepositoryArrayOutputWithContext(ctx).OutputState,
	}
}

// VirtualPuppetRepositoryMapInput is an input type that accepts VirtualPuppetRepositoryMap and VirtualPuppetRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualPuppetRepositoryMapInput` via:
//
//	VirtualPuppetRepositoryMap{ "key": VirtualPuppetRepositoryArgs{...} }
type VirtualPuppetRepositoryMapInput interface {
	pulumi.Input

	ToVirtualPuppetRepositoryMapOutput() VirtualPuppetRepositoryMapOutput
	ToVirtualPuppetRepositoryMapOutputWithContext(context.Context) VirtualPuppetRepositoryMapOutput
}

type VirtualPuppetRepositoryMap map[string]VirtualPuppetRepositoryInput

func (VirtualPuppetRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualPuppetRepository)(nil)).Elem()
}

func (i VirtualPuppetRepositoryMap) ToVirtualPuppetRepositoryMapOutput() VirtualPuppetRepositoryMapOutput {
	return i.ToVirtualPuppetRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualPuppetRepositoryMap) ToVirtualPuppetRepositoryMapOutputWithContext(ctx context.Context) VirtualPuppetRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualPuppetRepositoryMapOutput)
}

func (i VirtualPuppetRepositoryMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*VirtualPuppetRepository] {
	return pulumix.Output[map[string]*VirtualPuppetRepository]{
		OutputState: i.ToVirtualPuppetRepositoryMapOutputWithContext(ctx).OutputState,
	}
}

type VirtualPuppetRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualPuppetRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualPuppetRepository)(nil)).Elem()
}

func (o VirtualPuppetRepositoryOutput) ToVirtualPuppetRepositoryOutput() VirtualPuppetRepositoryOutput {
	return o
}

func (o VirtualPuppetRepositoryOutput) ToVirtualPuppetRepositoryOutputWithContext(ctx context.Context) VirtualPuppetRepositoryOutput {
	return o
}

func (o VirtualPuppetRepositoryOutput) ToOutput(ctx context.Context) pulumix.Output[*VirtualPuppetRepository] {
	return pulumix.Output[*VirtualPuppetRepository]{
		OutputState: o.OutputState,
	}
}

// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
// another Artifactory instance.
func (o VirtualPuppetRepositoryOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualPuppetRepository) pulumi.BoolPtrOutput {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

// Default repository to deploy artifacts.
func (o VirtualPuppetRepositoryOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualPuppetRepository) pulumi.StringPtrOutput { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

// Public description.
func (o VirtualPuppetRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualPuppetRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o VirtualPuppetRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualPuppetRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o VirtualPuppetRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualPuppetRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
// contain spaces or special characters.
func (o VirtualPuppetRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualPuppetRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o VirtualPuppetRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualPuppetRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o VirtualPuppetRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualPuppetRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o VirtualPuppetRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualPuppetRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o VirtualPuppetRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualPuppetRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Repository layout key for the local repository
func (o VirtualPuppetRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualPuppetRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The effective list of actual repositories included in this virtual repository.
func (o VirtualPuppetRepositoryOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualPuppetRepository) pulumi.StringArrayOutput { return v.Repositories }).(pulumi.StringArrayOutput)
}

type VirtualPuppetRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualPuppetRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualPuppetRepository)(nil)).Elem()
}

func (o VirtualPuppetRepositoryArrayOutput) ToVirtualPuppetRepositoryArrayOutput() VirtualPuppetRepositoryArrayOutput {
	return o
}

func (o VirtualPuppetRepositoryArrayOutput) ToVirtualPuppetRepositoryArrayOutputWithContext(ctx context.Context) VirtualPuppetRepositoryArrayOutput {
	return o
}

func (o VirtualPuppetRepositoryArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*VirtualPuppetRepository] {
	return pulumix.Output[[]*VirtualPuppetRepository]{
		OutputState: o.OutputState,
	}
}

func (o VirtualPuppetRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualPuppetRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualPuppetRepository {
		return vs[0].([]*VirtualPuppetRepository)[vs[1].(int)]
	}).(VirtualPuppetRepositoryOutput)
}

type VirtualPuppetRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualPuppetRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualPuppetRepository)(nil)).Elem()
}

func (o VirtualPuppetRepositoryMapOutput) ToVirtualPuppetRepositoryMapOutput() VirtualPuppetRepositoryMapOutput {
	return o
}

func (o VirtualPuppetRepositoryMapOutput) ToVirtualPuppetRepositoryMapOutputWithContext(ctx context.Context) VirtualPuppetRepositoryMapOutput {
	return o
}

func (o VirtualPuppetRepositoryMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*VirtualPuppetRepository] {
	return pulumix.Output[map[string]*VirtualPuppetRepository]{
		OutputState: o.OutputState,
	}
}

func (o VirtualPuppetRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualPuppetRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualPuppetRepository {
		return vs[0].(map[string]*VirtualPuppetRepository)[vs[1].(string)]
	}).(VirtualPuppetRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualPuppetRepositoryInput)(nil)).Elem(), &VirtualPuppetRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualPuppetRepositoryArrayInput)(nil)).Elem(), VirtualPuppetRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualPuppetRepositoryMapInput)(nil)).Elem(), VirtualPuppetRepositoryMap{})
	pulumi.RegisterOutputType(VirtualPuppetRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualPuppetRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualPuppetRepositoryMapOutput{})
}
