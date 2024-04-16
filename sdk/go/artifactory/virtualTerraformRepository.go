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
//			_, err := artifactory.NewVirtualTerraformRepository(ctx, "terraform-virtual", &artifactory.VirtualTerraformRepositoryArgs{
//				Key:             pulumi.String("terraform-remote"),
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
// $ pulumi import artifactory:index/virtualTerraformRepository:VirtualTerraformRepository terraform-virtual terraform-remote
// ```
type VirtualTerraformRepository struct {
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

// NewVirtualTerraformRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualTerraformRepository(ctx *pulumi.Context,
	name string, args *VirtualTerraformRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualTerraformRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VirtualTerraformRepository
	err := ctx.RegisterResource("artifactory:index/virtualTerraformRepository:VirtualTerraformRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualTerraformRepository gets an existing VirtualTerraformRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualTerraformRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualTerraformRepositoryState, opts ...pulumi.ResourceOption) (*VirtualTerraformRepository, error) {
	var resource VirtualTerraformRepository
	err := ctx.ReadResource("artifactory:index/virtualTerraformRepository:VirtualTerraformRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualTerraformRepository resources.
type virtualTerraformRepositoryState struct {
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

type VirtualTerraformRepositoryState struct {
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

func (VirtualTerraformRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualTerraformRepositoryState)(nil)).Elem()
}

type virtualTerraformRepositoryArgs struct {
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

// The set of arguments for constructing a VirtualTerraformRepository resource.
type VirtualTerraformRepositoryArgs struct {
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

func (VirtualTerraformRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualTerraformRepositoryArgs)(nil)).Elem()
}

type VirtualTerraformRepositoryInput interface {
	pulumi.Input

	ToVirtualTerraformRepositoryOutput() VirtualTerraformRepositoryOutput
	ToVirtualTerraformRepositoryOutputWithContext(ctx context.Context) VirtualTerraformRepositoryOutput
}

func (*VirtualTerraformRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualTerraformRepository)(nil)).Elem()
}

func (i *VirtualTerraformRepository) ToVirtualTerraformRepositoryOutput() VirtualTerraformRepositoryOutput {
	return i.ToVirtualTerraformRepositoryOutputWithContext(context.Background())
}

func (i *VirtualTerraformRepository) ToVirtualTerraformRepositoryOutputWithContext(ctx context.Context) VirtualTerraformRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualTerraformRepositoryOutput)
}

// VirtualTerraformRepositoryArrayInput is an input type that accepts VirtualTerraformRepositoryArray and VirtualTerraformRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualTerraformRepositoryArrayInput` via:
//
//	VirtualTerraformRepositoryArray{ VirtualTerraformRepositoryArgs{...} }
type VirtualTerraformRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualTerraformRepositoryArrayOutput() VirtualTerraformRepositoryArrayOutput
	ToVirtualTerraformRepositoryArrayOutputWithContext(context.Context) VirtualTerraformRepositoryArrayOutput
}

type VirtualTerraformRepositoryArray []VirtualTerraformRepositoryInput

func (VirtualTerraformRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualTerraformRepository)(nil)).Elem()
}

func (i VirtualTerraformRepositoryArray) ToVirtualTerraformRepositoryArrayOutput() VirtualTerraformRepositoryArrayOutput {
	return i.ToVirtualTerraformRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualTerraformRepositoryArray) ToVirtualTerraformRepositoryArrayOutputWithContext(ctx context.Context) VirtualTerraformRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualTerraformRepositoryArrayOutput)
}

// VirtualTerraformRepositoryMapInput is an input type that accepts VirtualTerraformRepositoryMap and VirtualTerraformRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualTerraformRepositoryMapInput` via:
//
//	VirtualTerraformRepositoryMap{ "key": VirtualTerraformRepositoryArgs{...} }
type VirtualTerraformRepositoryMapInput interface {
	pulumi.Input

	ToVirtualTerraformRepositoryMapOutput() VirtualTerraformRepositoryMapOutput
	ToVirtualTerraformRepositoryMapOutputWithContext(context.Context) VirtualTerraformRepositoryMapOutput
}

type VirtualTerraformRepositoryMap map[string]VirtualTerraformRepositoryInput

func (VirtualTerraformRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualTerraformRepository)(nil)).Elem()
}

func (i VirtualTerraformRepositoryMap) ToVirtualTerraformRepositoryMapOutput() VirtualTerraformRepositoryMapOutput {
	return i.ToVirtualTerraformRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualTerraformRepositoryMap) ToVirtualTerraformRepositoryMapOutputWithContext(ctx context.Context) VirtualTerraformRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualTerraformRepositoryMapOutput)
}

type VirtualTerraformRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualTerraformRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualTerraformRepository)(nil)).Elem()
}

func (o VirtualTerraformRepositoryOutput) ToVirtualTerraformRepositoryOutput() VirtualTerraformRepositoryOutput {
	return o
}

func (o VirtualTerraformRepositoryOutput) ToVirtualTerraformRepositoryOutputWithContext(ctx context.Context) VirtualTerraformRepositoryOutput {
	return o
}

// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
// another Artifactory instance.
func (o VirtualTerraformRepositoryOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualTerraformRepository) pulumi.BoolPtrOutput {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

// Default repository to deploy artifacts.
func (o VirtualTerraformRepositoryOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualTerraformRepository) pulumi.StringPtrOutput { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

// Public description.
func (o VirtualTerraformRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualTerraformRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o VirtualTerraformRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualTerraformRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o VirtualTerraformRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualTerraformRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
// contain spaces or special characters.
func (o VirtualTerraformRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualTerraformRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o VirtualTerraformRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualTerraformRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o VirtualTerraformRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualTerraformRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o VirtualTerraformRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualTerraformRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o VirtualTerraformRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualTerraformRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Repository layout key for the virtual repository
func (o VirtualTerraformRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualTerraformRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The effective list of actual repositories included in this virtual repository.
func (o VirtualTerraformRepositoryOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualTerraformRepository) pulumi.StringArrayOutput { return v.Repositories }).(pulumi.StringArrayOutput)
}

type VirtualTerraformRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualTerraformRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualTerraformRepository)(nil)).Elem()
}

func (o VirtualTerraformRepositoryArrayOutput) ToVirtualTerraformRepositoryArrayOutput() VirtualTerraformRepositoryArrayOutput {
	return o
}

func (o VirtualTerraformRepositoryArrayOutput) ToVirtualTerraformRepositoryArrayOutputWithContext(ctx context.Context) VirtualTerraformRepositoryArrayOutput {
	return o
}

func (o VirtualTerraformRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualTerraformRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualTerraformRepository {
		return vs[0].([]*VirtualTerraformRepository)[vs[1].(int)]
	}).(VirtualTerraformRepositoryOutput)
}

type VirtualTerraformRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualTerraformRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualTerraformRepository)(nil)).Elem()
}

func (o VirtualTerraformRepositoryMapOutput) ToVirtualTerraformRepositoryMapOutput() VirtualTerraformRepositoryMapOutput {
	return o
}

func (o VirtualTerraformRepositoryMapOutput) ToVirtualTerraformRepositoryMapOutputWithContext(ctx context.Context) VirtualTerraformRepositoryMapOutput {
	return o
}

func (o VirtualTerraformRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualTerraformRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualTerraformRepository {
		return vs[0].(map[string]*VirtualTerraformRepository)[vs[1].(string)]
	}).(VirtualTerraformRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualTerraformRepositoryInput)(nil)).Elem(), &VirtualTerraformRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualTerraformRepositoryArrayInput)(nil)).Elem(), VirtualTerraformRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualTerraformRepositoryMapInput)(nil)).Elem(), VirtualTerraformRepositoryMap{})
	pulumi.RegisterOutputType(VirtualTerraformRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualTerraformRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualTerraformRepositoryMapOutput{})
}
