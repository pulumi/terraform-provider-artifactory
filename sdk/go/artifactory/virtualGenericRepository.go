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

// Creates a virtual Generic repository.
// Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Virtual+Repositories).
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
//			_, err := artifactory.NewVirtualGenericRepository(ctx, "foo-generic", &artifactory.VirtualGenericRepositoryArgs{
//				Key:             pulumi.String("foo-generic"),
//				RepoLayoutRef:   pulumi.String("simple-default"),
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
// $ pulumi import artifactory:index/virtualGenericRepository:VirtualGenericRepository foo-generic foo-generic
// ```
type VirtualGenericRepository struct {
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

// NewVirtualGenericRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualGenericRepository(ctx *pulumi.Context,
	name string, args *VirtualGenericRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualGenericRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VirtualGenericRepository
	err := ctx.RegisterResource("artifactory:index/virtualGenericRepository:VirtualGenericRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualGenericRepository gets an existing VirtualGenericRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualGenericRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualGenericRepositoryState, opts ...pulumi.ResourceOption) (*VirtualGenericRepository, error) {
	var resource VirtualGenericRepository
	err := ctx.ReadResource("artifactory:index/virtualGenericRepository:VirtualGenericRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualGenericRepository resources.
type virtualGenericRepositoryState struct {
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

type VirtualGenericRepositoryState struct {
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

func (VirtualGenericRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualGenericRepositoryState)(nil)).Elem()
}

type virtualGenericRepositoryArgs struct {
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

// The set of arguments for constructing a VirtualGenericRepository resource.
type VirtualGenericRepositoryArgs struct {
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

func (VirtualGenericRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualGenericRepositoryArgs)(nil)).Elem()
}

type VirtualGenericRepositoryInput interface {
	pulumi.Input

	ToVirtualGenericRepositoryOutput() VirtualGenericRepositoryOutput
	ToVirtualGenericRepositoryOutputWithContext(ctx context.Context) VirtualGenericRepositoryOutput
}

func (*VirtualGenericRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualGenericRepository)(nil)).Elem()
}

func (i *VirtualGenericRepository) ToVirtualGenericRepositoryOutput() VirtualGenericRepositoryOutput {
	return i.ToVirtualGenericRepositoryOutputWithContext(context.Background())
}

func (i *VirtualGenericRepository) ToVirtualGenericRepositoryOutputWithContext(ctx context.Context) VirtualGenericRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualGenericRepositoryOutput)
}

// VirtualGenericRepositoryArrayInput is an input type that accepts VirtualGenericRepositoryArray and VirtualGenericRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualGenericRepositoryArrayInput` via:
//
//	VirtualGenericRepositoryArray{ VirtualGenericRepositoryArgs{...} }
type VirtualGenericRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualGenericRepositoryArrayOutput() VirtualGenericRepositoryArrayOutput
	ToVirtualGenericRepositoryArrayOutputWithContext(context.Context) VirtualGenericRepositoryArrayOutput
}

type VirtualGenericRepositoryArray []VirtualGenericRepositoryInput

func (VirtualGenericRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualGenericRepository)(nil)).Elem()
}

func (i VirtualGenericRepositoryArray) ToVirtualGenericRepositoryArrayOutput() VirtualGenericRepositoryArrayOutput {
	return i.ToVirtualGenericRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualGenericRepositoryArray) ToVirtualGenericRepositoryArrayOutputWithContext(ctx context.Context) VirtualGenericRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualGenericRepositoryArrayOutput)
}

// VirtualGenericRepositoryMapInput is an input type that accepts VirtualGenericRepositoryMap and VirtualGenericRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualGenericRepositoryMapInput` via:
//
//	VirtualGenericRepositoryMap{ "key": VirtualGenericRepositoryArgs{...} }
type VirtualGenericRepositoryMapInput interface {
	pulumi.Input

	ToVirtualGenericRepositoryMapOutput() VirtualGenericRepositoryMapOutput
	ToVirtualGenericRepositoryMapOutputWithContext(context.Context) VirtualGenericRepositoryMapOutput
}

type VirtualGenericRepositoryMap map[string]VirtualGenericRepositoryInput

func (VirtualGenericRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualGenericRepository)(nil)).Elem()
}

func (i VirtualGenericRepositoryMap) ToVirtualGenericRepositoryMapOutput() VirtualGenericRepositoryMapOutput {
	return i.ToVirtualGenericRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualGenericRepositoryMap) ToVirtualGenericRepositoryMapOutputWithContext(ctx context.Context) VirtualGenericRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualGenericRepositoryMapOutput)
}

type VirtualGenericRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualGenericRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualGenericRepository)(nil)).Elem()
}

func (o VirtualGenericRepositoryOutput) ToVirtualGenericRepositoryOutput() VirtualGenericRepositoryOutput {
	return o
}

func (o VirtualGenericRepositoryOutput) ToVirtualGenericRepositoryOutputWithContext(ctx context.Context) VirtualGenericRepositoryOutput {
	return o
}

// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
// another Artifactory instance.
func (o VirtualGenericRepositoryOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualGenericRepository) pulumi.BoolPtrOutput {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

// Default repository to deploy artifacts.
func (o VirtualGenericRepositoryOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualGenericRepository) pulumi.StringPtrOutput { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

// Public description.
func (o VirtualGenericRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualGenericRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o VirtualGenericRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualGenericRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o VirtualGenericRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualGenericRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
// contain spaces or special characters.
func (o VirtualGenericRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualGenericRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o VirtualGenericRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualGenericRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o VirtualGenericRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualGenericRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o VirtualGenericRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualGenericRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o VirtualGenericRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualGenericRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Repository layout key for the virtual repository
func (o VirtualGenericRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualGenericRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The effective list of actual repositories included in this virtual repository.
func (o VirtualGenericRepositoryOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualGenericRepository) pulumi.StringArrayOutput { return v.Repositories }).(pulumi.StringArrayOutput)
}

type VirtualGenericRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualGenericRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualGenericRepository)(nil)).Elem()
}

func (o VirtualGenericRepositoryArrayOutput) ToVirtualGenericRepositoryArrayOutput() VirtualGenericRepositoryArrayOutput {
	return o
}

func (o VirtualGenericRepositoryArrayOutput) ToVirtualGenericRepositoryArrayOutputWithContext(ctx context.Context) VirtualGenericRepositoryArrayOutput {
	return o
}

func (o VirtualGenericRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualGenericRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualGenericRepository {
		return vs[0].([]*VirtualGenericRepository)[vs[1].(int)]
	}).(VirtualGenericRepositoryOutput)
}

type VirtualGenericRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualGenericRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualGenericRepository)(nil)).Elem()
}

func (o VirtualGenericRepositoryMapOutput) ToVirtualGenericRepositoryMapOutput() VirtualGenericRepositoryMapOutput {
	return o
}

func (o VirtualGenericRepositoryMapOutput) ToVirtualGenericRepositoryMapOutputWithContext(ctx context.Context) VirtualGenericRepositoryMapOutput {
	return o
}

func (o VirtualGenericRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualGenericRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualGenericRepository {
		return vs[0].(map[string]*VirtualGenericRepository)[vs[1].(string)]
	}).(VirtualGenericRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualGenericRepositoryInput)(nil)).Elem(), &VirtualGenericRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualGenericRepositoryArrayInput)(nil)).Elem(), VirtualGenericRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualGenericRepositoryMapInput)(nil)).Elem(), VirtualGenericRepositoryMap{})
	pulumi.RegisterOutputType(VirtualGenericRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualGenericRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualGenericRepositoryMapOutput{})
}
