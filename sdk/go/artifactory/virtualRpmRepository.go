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

// Creates a virtual Rpm repository.
// Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/RPM+Repositories).
//
// ## Import
//
// Virtual repositories can be imported using their name, e.g.
//
// ```sh
// $ pulumi import artifactory:index/virtualRpmRepository:VirtualRpmRepository foo-rpm-virtual foo-rpm-virtual
// ```
type VirtualRpmRepository struct {
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
	// The primary GPG key to be used to sign packages.
	PrimaryKeypairRef pulumi.StringPtrOutput `pulumi:"primaryKeypairRef"`
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
	// Secondary keypair used to sign artifacts.
	SecondaryKeypairRef pulumi.StringPtrOutput `pulumi:"secondaryKeypairRef"`
}

// NewVirtualRpmRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualRpmRepository(ctx *pulumi.Context,
	name string, args *VirtualRpmRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualRpmRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VirtualRpmRepository
	err := ctx.RegisterResource("artifactory:index/virtualRpmRepository:VirtualRpmRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualRpmRepository gets an existing VirtualRpmRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualRpmRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualRpmRepositoryState, opts ...pulumi.ResourceOption) (*VirtualRpmRepository, error) {
	var resource VirtualRpmRepository
	err := ctx.ReadResource("artifactory:index/virtualRpmRepository:VirtualRpmRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualRpmRepository resources.
type virtualRpmRepositoryState struct {
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
	// The primary GPG key to be used to sign packages.
	PrimaryKeypairRef *string `pulumi:"primaryKeypairRef"`
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
	// Secondary keypair used to sign artifacts.
	SecondaryKeypairRef *string `pulumi:"secondaryKeypairRef"`
}

type VirtualRpmRepositoryState struct {
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
	// The primary GPG key to be used to sign packages.
	PrimaryKeypairRef pulumi.StringPtrInput
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
	// Secondary keypair used to sign artifacts.
	SecondaryKeypairRef pulumi.StringPtrInput
}

func (VirtualRpmRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualRpmRepositoryState)(nil)).Elem()
}

type virtualRpmRepositoryArgs struct {
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
	// The primary GPG key to be used to sign packages.
	PrimaryKeypairRef *string `pulumi:"primaryKeypairRef"`
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
	// Secondary keypair used to sign artifacts.
	SecondaryKeypairRef *string `pulumi:"secondaryKeypairRef"`
}

// The set of arguments for constructing a VirtualRpmRepository resource.
type VirtualRpmRepositoryArgs struct {
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
	// The primary GPG key to be used to sign packages.
	PrimaryKeypairRef pulumi.StringPtrInput
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
	// Secondary keypair used to sign artifacts.
	SecondaryKeypairRef pulumi.StringPtrInput
}

func (VirtualRpmRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualRpmRepositoryArgs)(nil)).Elem()
}

type VirtualRpmRepositoryInput interface {
	pulumi.Input

	ToVirtualRpmRepositoryOutput() VirtualRpmRepositoryOutput
	ToVirtualRpmRepositoryOutputWithContext(ctx context.Context) VirtualRpmRepositoryOutput
}

func (*VirtualRpmRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualRpmRepository)(nil)).Elem()
}

func (i *VirtualRpmRepository) ToVirtualRpmRepositoryOutput() VirtualRpmRepositoryOutput {
	return i.ToVirtualRpmRepositoryOutputWithContext(context.Background())
}

func (i *VirtualRpmRepository) ToVirtualRpmRepositoryOutputWithContext(ctx context.Context) VirtualRpmRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualRpmRepositoryOutput)
}

// VirtualRpmRepositoryArrayInput is an input type that accepts VirtualRpmRepositoryArray and VirtualRpmRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualRpmRepositoryArrayInput` via:
//
//	VirtualRpmRepositoryArray{ VirtualRpmRepositoryArgs{...} }
type VirtualRpmRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualRpmRepositoryArrayOutput() VirtualRpmRepositoryArrayOutput
	ToVirtualRpmRepositoryArrayOutputWithContext(context.Context) VirtualRpmRepositoryArrayOutput
}

type VirtualRpmRepositoryArray []VirtualRpmRepositoryInput

func (VirtualRpmRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualRpmRepository)(nil)).Elem()
}

func (i VirtualRpmRepositoryArray) ToVirtualRpmRepositoryArrayOutput() VirtualRpmRepositoryArrayOutput {
	return i.ToVirtualRpmRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualRpmRepositoryArray) ToVirtualRpmRepositoryArrayOutputWithContext(ctx context.Context) VirtualRpmRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualRpmRepositoryArrayOutput)
}

// VirtualRpmRepositoryMapInput is an input type that accepts VirtualRpmRepositoryMap and VirtualRpmRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualRpmRepositoryMapInput` via:
//
//	VirtualRpmRepositoryMap{ "key": VirtualRpmRepositoryArgs{...} }
type VirtualRpmRepositoryMapInput interface {
	pulumi.Input

	ToVirtualRpmRepositoryMapOutput() VirtualRpmRepositoryMapOutput
	ToVirtualRpmRepositoryMapOutputWithContext(context.Context) VirtualRpmRepositoryMapOutput
}

type VirtualRpmRepositoryMap map[string]VirtualRpmRepositoryInput

func (VirtualRpmRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualRpmRepository)(nil)).Elem()
}

func (i VirtualRpmRepositoryMap) ToVirtualRpmRepositoryMapOutput() VirtualRpmRepositoryMapOutput {
	return i.ToVirtualRpmRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualRpmRepositoryMap) ToVirtualRpmRepositoryMapOutputWithContext(ctx context.Context) VirtualRpmRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualRpmRepositoryMapOutput)
}

type VirtualRpmRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualRpmRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualRpmRepository)(nil)).Elem()
}

func (o VirtualRpmRepositoryOutput) ToVirtualRpmRepositoryOutput() VirtualRpmRepositoryOutput {
	return o
}

func (o VirtualRpmRepositoryOutput) ToVirtualRpmRepositoryOutputWithContext(ctx context.Context) VirtualRpmRepositoryOutput {
	return o
}

// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
// another Artifactory instance.
func (o VirtualRpmRepositoryOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.BoolPtrOutput {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

// Default repository to deploy artifacts.
func (o VirtualRpmRepositoryOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.StringPtrOutput { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

// Public description.
func (o VirtualRpmRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o VirtualRpmRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o VirtualRpmRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
// contain spaces or special characters.
func (o VirtualRpmRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o VirtualRpmRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o VirtualRpmRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// The primary GPG key to be used to sign packages.
func (o VirtualRpmRepositoryOutput) PrimaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.StringPtrOutput { return v.PrimaryKeypairRef }).(pulumi.StringPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o VirtualRpmRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o VirtualRpmRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Repository layout key for the virtual repository
func (o VirtualRpmRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The effective list of actual repositories included in this virtual repository.
func (o VirtualRpmRepositoryOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.StringArrayOutput { return v.Repositories }).(pulumi.StringArrayOutput)
}

// Secondary keypair used to sign artifacts.
func (o VirtualRpmRepositoryOutput) SecondaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.StringPtrOutput { return v.SecondaryKeypairRef }).(pulumi.StringPtrOutput)
}

type VirtualRpmRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualRpmRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualRpmRepository)(nil)).Elem()
}

func (o VirtualRpmRepositoryArrayOutput) ToVirtualRpmRepositoryArrayOutput() VirtualRpmRepositoryArrayOutput {
	return o
}

func (o VirtualRpmRepositoryArrayOutput) ToVirtualRpmRepositoryArrayOutputWithContext(ctx context.Context) VirtualRpmRepositoryArrayOutput {
	return o
}

func (o VirtualRpmRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualRpmRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualRpmRepository {
		return vs[0].([]*VirtualRpmRepository)[vs[1].(int)]
	}).(VirtualRpmRepositoryOutput)
}

type VirtualRpmRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualRpmRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualRpmRepository)(nil)).Elem()
}

func (o VirtualRpmRepositoryMapOutput) ToVirtualRpmRepositoryMapOutput() VirtualRpmRepositoryMapOutput {
	return o
}

func (o VirtualRpmRepositoryMapOutput) ToVirtualRpmRepositoryMapOutputWithContext(ctx context.Context) VirtualRpmRepositoryMapOutput {
	return o
}

func (o VirtualRpmRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualRpmRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualRpmRepository {
		return vs[0].(map[string]*VirtualRpmRepository)[vs[1].(string)]
	}).(VirtualRpmRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualRpmRepositoryInput)(nil)).Elem(), &VirtualRpmRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualRpmRepositoryArrayInput)(nil)).Elem(), VirtualRpmRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualRpmRepositoryMapInput)(nil)).Elem(), VirtualRpmRepositoryMap{})
	pulumi.RegisterOutputType(VirtualRpmRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualRpmRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualRpmRepositoryMapOutput{})
}
