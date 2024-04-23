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

// Creates a virtual Nuget repository.
// Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/NuGet+Repositories#NuGetRepositories-VirtualRepositories).
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
//			_, err := artifactory.NewVirtualNugetRepository(ctx, "foo-nuget", &artifactory.VirtualNugetRepositoryArgs{
//				Key:                      pulumi.String("foo-nuget"),
//				Repositories:             pulumi.StringArray{},
//				Description:              pulumi.String("A test virtual repo"),
//				Notes:                    pulumi.String("Internal description"),
//				IncludesPattern:          pulumi.String("com/jfrog/**,cloud/jfrog/**"),
//				ExcludesPattern:          pulumi.String("com/google/**"),
//				ForceNugetAuthentication: pulumi.Bool(true),
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
// $ pulumi import artifactory:index/virtualNugetRepository:VirtualNugetRepository foo-nuget foo-nuget
// ```
type VirtualNugetRepository struct {
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
	// If set, user authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This is also enforced when aggregated repositories support anonymous requests. Default is `false`.
	ForceNugetAuthentication pulumi.BoolPtrOutput `pulumi:"forceNugetAuthentication"`
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

// NewVirtualNugetRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualNugetRepository(ctx *pulumi.Context,
	name string, args *VirtualNugetRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualNugetRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VirtualNugetRepository
	err := ctx.RegisterResource("artifactory:index/virtualNugetRepository:VirtualNugetRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNugetRepository gets an existing VirtualNugetRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNugetRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNugetRepositoryState, opts ...pulumi.ResourceOption) (*VirtualNugetRepository, error) {
	var resource VirtualNugetRepository
	err := ctx.ReadResource("artifactory:index/virtualNugetRepository:VirtualNugetRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNugetRepository resources.
type virtualNugetRepositoryState struct {
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
	// If set, user authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This is also enforced when aggregated repositories support anonymous requests. Default is `false`.
	ForceNugetAuthentication *bool `pulumi:"forceNugetAuthentication"`
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

type VirtualNugetRepositoryState struct {
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
	// If set, user authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This is also enforced when aggregated repositories support anonymous requests. Default is `false`.
	ForceNugetAuthentication pulumi.BoolPtrInput
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

func (VirtualNugetRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNugetRepositoryState)(nil)).Elem()
}

type virtualNugetRepositoryArgs struct {
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
	// If set, user authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This is also enforced when aggregated repositories support anonymous requests. Default is `false`.
	ForceNugetAuthentication *bool `pulumi:"forceNugetAuthentication"`
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

// The set of arguments for constructing a VirtualNugetRepository resource.
type VirtualNugetRepositoryArgs struct {
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
	// If set, user authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This is also enforced when aggregated repositories support anonymous requests. Default is `false`.
	ForceNugetAuthentication pulumi.BoolPtrInput
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

func (VirtualNugetRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNugetRepositoryArgs)(nil)).Elem()
}

type VirtualNugetRepositoryInput interface {
	pulumi.Input

	ToVirtualNugetRepositoryOutput() VirtualNugetRepositoryOutput
	ToVirtualNugetRepositoryOutputWithContext(ctx context.Context) VirtualNugetRepositoryOutput
}

func (*VirtualNugetRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNugetRepository)(nil)).Elem()
}

func (i *VirtualNugetRepository) ToVirtualNugetRepositoryOutput() VirtualNugetRepositoryOutput {
	return i.ToVirtualNugetRepositoryOutputWithContext(context.Background())
}

func (i *VirtualNugetRepository) ToVirtualNugetRepositoryOutputWithContext(ctx context.Context) VirtualNugetRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNugetRepositoryOutput)
}

// VirtualNugetRepositoryArrayInput is an input type that accepts VirtualNugetRepositoryArray and VirtualNugetRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualNugetRepositoryArrayInput` via:
//
//	VirtualNugetRepositoryArray{ VirtualNugetRepositoryArgs{...} }
type VirtualNugetRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualNugetRepositoryArrayOutput() VirtualNugetRepositoryArrayOutput
	ToVirtualNugetRepositoryArrayOutputWithContext(context.Context) VirtualNugetRepositoryArrayOutput
}

type VirtualNugetRepositoryArray []VirtualNugetRepositoryInput

func (VirtualNugetRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualNugetRepository)(nil)).Elem()
}

func (i VirtualNugetRepositoryArray) ToVirtualNugetRepositoryArrayOutput() VirtualNugetRepositoryArrayOutput {
	return i.ToVirtualNugetRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualNugetRepositoryArray) ToVirtualNugetRepositoryArrayOutputWithContext(ctx context.Context) VirtualNugetRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNugetRepositoryArrayOutput)
}

// VirtualNugetRepositoryMapInput is an input type that accepts VirtualNugetRepositoryMap and VirtualNugetRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualNugetRepositoryMapInput` via:
//
//	VirtualNugetRepositoryMap{ "key": VirtualNugetRepositoryArgs{...} }
type VirtualNugetRepositoryMapInput interface {
	pulumi.Input

	ToVirtualNugetRepositoryMapOutput() VirtualNugetRepositoryMapOutput
	ToVirtualNugetRepositoryMapOutputWithContext(context.Context) VirtualNugetRepositoryMapOutput
}

type VirtualNugetRepositoryMap map[string]VirtualNugetRepositoryInput

func (VirtualNugetRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualNugetRepository)(nil)).Elem()
}

func (i VirtualNugetRepositoryMap) ToVirtualNugetRepositoryMapOutput() VirtualNugetRepositoryMapOutput {
	return i.ToVirtualNugetRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualNugetRepositoryMap) ToVirtualNugetRepositoryMapOutputWithContext(ctx context.Context) VirtualNugetRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNugetRepositoryMapOutput)
}

type VirtualNugetRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualNugetRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNugetRepository)(nil)).Elem()
}

func (o VirtualNugetRepositoryOutput) ToVirtualNugetRepositoryOutput() VirtualNugetRepositoryOutput {
	return o
}

func (o VirtualNugetRepositoryOutput) ToVirtualNugetRepositoryOutputWithContext(ctx context.Context) VirtualNugetRepositoryOutput {
	return o
}

// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
// another Artifactory instance.
func (o VirtualNugetRepositoryOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNugetRepository) pulumi.BoolPtrOutput {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

// Default repository to deploy artifacts.
func (o VirtualNugetRepositoryOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNugetRepository) pulumi.StringPtrOutput { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

// Public description.
func (o VirtualNugetRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNugetRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o VirtualNugetRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNugetRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// If set, user authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This is also enforced when aggregated repositories support anonymous requests. Default is `false`.
func (o VirtualNugetRepositoryOutput) ForceNugetAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNugetRepository) pulumi.BoolPtrOutput { return v.ForceNugetAuthentication }).(pulumi.BoolPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o VirtualNugetRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNugetRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
// contain spaces or special characters.
func (o VirtualNugetRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNugetRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o VirtualNugetRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNugetRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o VirtualNugetRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNugetRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o VirtualNugetRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualNugetRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o VirtualNugetRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNugetRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Repository layout key for the virtual repository
func (o VirtualNugetRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNugetRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The effective list of actual repositories included in this virtual repository.
func (o VirtualNugetRepositoryOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualNugetRepository) pulumi.StringArrayOutput { return v.Repositories }).(pulumi.StringArrayOutput)
}

type VirtualNugetRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualNugetRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualNugetRepository)(nil)).Elem()
}

func (o VirtualNugetRepositoryArrayOutput) ToVirtualNugetRepositoryArrayOutput() VirtualNugetRepositoryArrayOutput {
	return o
}

func (o VirtualNugetRepositoryArrayOutput) ToVirtualNugetRepositoryArrayOutputWithContext(ctx context.Context) VirtualNugetRepositoryArrayOutput {
	return o
}

func (o VirtualNugetRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualNugetRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualNugetRepository {
		return vs[0].([]*VirtualNugetRepository)[vs[1].(int)]
	}).(VirtualNugetRepositoryOutput)
}

type VirtualNugetRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualNugetRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualNugetRepository)(nil)).Elem()
}

func (o VirtualNugetRepositoryMapOutput) ToVirtualNugetRepositoryMapOutput() VirtualNugetRepositoryMapOutput {
	return o
}

func (o VirtualNugetRepositoryMapOutput) ToVirtualNugetRepositoryMapOutputWithContext(ctx context.Context) VirtualNugetRepositoryMapOutput {
	return o
}

func (o VirtualNugetRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualNugetRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualNugetRepository {
		return vs[0].(map[string]*VirtualNugetRepository)[vs[1].(string)]
	}).(VirtualNugetRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualNugetRepositoryInput)(nil)).Elem(), &VirtualNugetRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualNugetRepositoryArrayInput)(nil)).Elem(), VirtualNugetRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualNugetRepositoryMapInput)(nil)).Elem(), VirtualNugetRepositoryMap{})
	pulumi.RegisterOutputType(VirtualNugetRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualNugetRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualNugetRepositoryMapOutput{})
}
