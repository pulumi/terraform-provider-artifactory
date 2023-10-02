// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Creates a virtual Swift repository.
// Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Swift+Repositories#SwiftRepositories-VirtualRepositories).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-swift/sdk/v1/go/swift"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := index.NewSwift(ctx, "foo-conan", &index.SwiftArgs{
//				Key:             "foo-conan",
//				RepoLayoutRef:   "swift-default",
//				Repositories:    []interface{}{},
//				Description:     "A test virtual repo",
//				Notes:           "Internal description",
//				IncludesPattern: "com/jfrog/**,cloud/jfrog/**",
//				ExcludesPattern: "com/google/**",
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
//	$ pulumi import artifactory:index/virtualSwiftRepository:VirtualSwiftRepository foo-swift foo-swift
//
// ```
type VirtualSwiftRepository struct {
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
	// Repository layout key for the virtual repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayOutput `pulumi:"repositories"`
}

// NewVirtualSwiftRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualSwiftRepository(ctx *pulumi.Context,
	name string, args *VirtualSwiftRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualSwiftRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VirtualSwiftRepository
	err := ctx.RegisterResource("artifactory:index/virtualSwiftRepository:VirtualSwiftRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualSwiftRepository gets an existing VirtualSwiftRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualSwiftRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualSwiftRepositoryState, opts ...pulumi.ResourceOption) (*VirtualSwiftRepository, error) {
	var resource VirtualSwiftRepository
	err := ctx.ReadResource("artifactory:index/virtualSwiftRepository:VirtualSwiftRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualSwiftRepository resources.
type virtualSwiftRepositoryState struct {
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
	// Repository layout key for the virtual repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
}

type VirtualSwiftRepositoryState struct {
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
	// Repository layout key for the virtual repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
}

func (VirtualSwiftRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualSwiftRepositoryState)(nil)).Elem()
}

type virtualSwiftRepositoryArgs struct {
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
	// Repository layout key for the virtual repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
}

// The set of arguments for constructing a VirtualSwiftRepository resource.
type VirtualSwiftRepositoryArgs struct {
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
	// Repository layout key for the virtual repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
}

func (VirtualSwiftRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualSwiftRepositoryArgs)(nil)).Elem()
}

type VirtualSwiftRepositoryInput interface {
	pulumi.Input

	ToVirtualSwiftRepositoryOutput() VirtualSwiftRepositoryOutput
	ToVirtualSwiftRepositoryOutputWithContext(ctx context.Context) VirtualSwiftRepositoryOutput
}

func (*VirtualSwiftRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualSwiftRepository)(nil)).Elem()
}

func (i *VirtualSwiftRepository) ToVirtualSwiftRepositoryOutput() VirtualSwiftRepositoryOutput {
	return i.ToVirtualSwiftRepositoryOutputWithContext(context.Background())
}

func (i *VirtualSwiftRepository) ToVirtualSwiftRepositoryOutputWithContext(ctx context.Context) VirtualSwiftRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualSwiftRepositoryOutput)
}

func (i *VirtualSwiftRepository) ToOutput(ctx context.Context) pulumix.Output[*VirtualSwiftRepository] {
	return pulumix.Output[*VirtualSwiftRepository]{
		OutputState: i.ToVirtualSwiftRepositoryOutputWithContext(ctx).OutputState,
	}
}

// VirtualSwiftRepositoryArrayInput is an input type that accepts VirtualSwiftRepositoryArray and VirtualSwiftRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualSwiftRepositoryArrayInput` via:
//
//	VirtualSwiftRepositoryArray{ VirtualSwiftRepositoryArgs{...} }
type VirtualSwiftRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualSwiftRepositoryArrayOutput() VirtualSwiftRepositoryArrayOutput
	ToVirtualSwiftRepositoryArrayOutputWithContext(context.Context) VirtualSwiftRepositoryArrayOutput
}

type VirtualSwiftRepositoryArray []VirtualSwiftRepositoryInput

func (VirtualSwiftRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualSwiftRepository)(nil)).Elem()
}

func (i VirtualSwiftRepositoryArray) ToVirtualSwiftRepositoryArrayOutput() VirtualSwiftRepositoryArrayOutput {
	return i.ToVirtualSwiftRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualSwiftRepositoryArray) ToVirtualSwiftRepositoryArrayOutputWithContext(ctx context.Context) VirtualSwiftRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualSwiftRepositoryArrayOutput)
}

func (i VirtualSwiftRepositoryArray) ToOutput(ctx context.Context) pulumix.Output[[]*VirtualSwiftRepository] {
	return pulumix.Output[[]*VirtualSwiftRepository]{
		OutputState: i.ToVirtualSwiftRepositoryArrayOutputWithContext(ctx).OutputState,
	}
}

// VirtualSwiftRepositoryMapInput is an input type that accepts VirtualSwiftRepositoryMap and VirtualSwiftRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualSwiftRepositoryMapInput` via:
//
//	VirtualSwiftRepositoryMap{ "key": VirtualSwiftRepositoryArgs{...} }
type VirtualSwiftRepositoryMapInput interface {
	pulumi.Input

	ToVirtualSwiftRepositoryMapOutput() VirtualSwiftRepositoryMapOutput
	ToVirtualSwiftRepositoryMapOutputWithContext(context.Context) VirtualSwiftRepositoryMapOutput
}

type VirtualSwiftRepositoryMap map[string]VirtualSwiftRepositoryInput

func (VirtualSwiftRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualSwiftRepository)(nil)).Elem()
}

func (i VirtualSwiftRepositoryMap) ToVirtualSwiftRepositoryMapOutput() VirtualSwiftRepositoryMapOutput {
	return i.ToVirtualSwiftRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualSwiftRepositoryMap) ToVirtualSwiftRepositoryMapOutputWithContext(ctx context.Context) VirtualSwiftRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualSwiftRepositoryMapOutput)
}

func (i VirtualSwiftRepositoryMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*VirtualSwiftRepository] {
	return pulumix.Output[map[string]*VirtualSwiftRepository]{
		OutputState: i.ToVirtualSwiftRepositoryMapOutputWithContext(ctx).OutputState,
	}
}

type VirtualSwiftRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualSwiftRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualSwiftRepository)(nil)).Elem()
}

func (o VirtualSwiftRepositoryOutput) ToVirtualSwiftRepositoryOutput() VirtualSwiftRepositoryOutput {
	return o
}

func (o VirtualSwiftRepositoryOutput) ToVirtualSwiftRepositoryOutputWithContext(ctx context.Context) VirtualSwiftRepositoryOutput {
	return o
}

func (o VirtualSwiftRepositoryOutput) ToOutput(ctx context.Context) pulumix.Output[*VirtualSwiftRepository] {
	return pulumix.Output[*VirtualSwiftRepository]{
		OutputState: o.OutputState,
	}
}

// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
// another Artifactory instance.
func (o VirtualSwiftRepositoryOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualSwiftRepository) pulumi.BoolPtrOutput {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

// Default repository to deploy artifacts.
func (o VirtualSwiftRepositoryOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualSwiftRepository) pulumi.StringPtrOutput { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

// Public description.
func (o VirtualSwiftRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualSwiftRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o VirtualSwiftRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualSwiftRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o VirtualSwiftRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualSwiftRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
// contain spaces or special characters.
func (o VirtualSwiftRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualSwiftRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o VirtualSwiftRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualSwiftRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o VirtualSwiftRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualSwiftRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o VirtualSwiftRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualSwiftRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o VirtualSwiftRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualSwiftRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Repository layout key for the virtual repository
func (o VirtualSwiftRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualSwiftRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The effective list of actual repositories included in this virtual repository.
func (o VirtualSwiftRepositoryOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualSwiftRepository) pulumi.StringArrayOutput { return v.Repositories }).(pulumi.StringArrayOutput)
}

type VirtualSwiftRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualSwiftRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualSwiftRepository)(nil)).Elem()
}

func (o VirtualSwiftRepositoryArrayOutput) ToVirtualSwiftRepositoryArrayOutput() VirtualSwiftRepositoryArrayOutput {
	return o
}

func (o VirtualSwiftRepositoryArrayOutput) ToVirtualSwiftRepositoryArrayOutputWithContext(ctx context.Context) VirtualSwiftRepositoryArrayOutput {
	return o
}

func (o VirtualSwiftRepositoryArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*VirtualSwiftRepository] {
	return pulumix.Output[[]*VirtualSwiftRepository]{
		OutputState: o.OutputState,
	}
}

func (o VirtualSwiftRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualSwiftRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualSwiftRepository {
		return vs[0].([]*VirtualSwiftRepository)[vs[1].(int)]
	}).(VirtualSwiftRepositoryOutput)
}

type VirtualSwiftRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualSwiftRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualSwiftRepository)(nil)).Elem()
}

func (o VirtualSwiftRepositoryMapOutput) ToVirtualSwiftRepositoryMapOutput() VirtualSwiftRepositoryMapOutput {
	return o
}

func (o VirtualSwiftRepositoryMapOutput) ToVirtualSwiftRepositoryMapOutputWithContext(ctx context.Context) VirtualSwiftRepositoryMapOutput {
	return o
}

func (o VirtualSwiftRepositoryMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*VirtualSwiftRepository] {
	return pulumix.Output[map[string]*VirtualSwiftRepository]{
		OutputState: o.OutputState,
	}
}

func (o VirtualSwiftRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualSwiftRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualSwiftRepository {
		return vs[0].(map[string]*VirtualSwiftRepository)[vs[1].(string)]
	}).(VirtualSwiftRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualSwiftRepositoryInput)(nil)).Elem(), &VirtualSwiftRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualSwiftRepositoryArrayInput)(nil)).Elem(), VirtualSwiftRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualSwiftRepositoryMapInput)(nil)).Elem(), VirtualSwiftRepositoryMap{})
	pulumi.RegisterOutputType(VirtualSwiftRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualSwiftRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualSwiftRepositoryMapOutput{})
}
