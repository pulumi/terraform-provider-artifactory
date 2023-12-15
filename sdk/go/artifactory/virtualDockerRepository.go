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

// Creates a virtual Docker repository.
// Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Docker+Registry#DockerRegistry-VirtualDockerRepositories).
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
//			_, err := artifactory.NewVirtualDockerRepository(ctx, "foo-docker", &artifactory.VirtualDockerRepositoryArgs{
//				Description:                  pulumi.String("A test virtual repo"),
//				ExcludesPattern:              pulumi.String("com/google/**"),
//				IncludesPattern:              pulumi.String("com/jfrog/**,cloud/jfrog/**"),
//				Key:                          pulumi.String("foo-docker"),
//				Notes:                        pulumi.String("Internal description"),
//				Repositories:                 pulumi.StringArray{},
//				ResolveDockerTagsByTimestamp: pulumi.Bool(true),
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
//	$ pulumi import artifactory:index/virtualDockerRepository:VirtualDockerRepository foo-docker foo-docker
//
// ```
type VirtualDockerRepository struct {
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
	// When enabled, in cases where the same Docker tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
	ResolveDockerTagsByTimestamp pulumi.BoolPtrOutput `pulumi:"resolveDockerTagsByTimestamp"`
}

// NewVirtualDockerRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualDockerRepository(ctx *pulumi.Context,
	name string, args *VirtualDockerRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualDockerRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VirtualDockerRepository
	err := ctx.RegisterResource("artifactory:index/virtualDockerRepository:VirtualDockerRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualDockerRepository gets an existing VirtualDockerRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualDockerRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualDockerRepositoryState, opts ...pulumi.ResourceOption) (*VirtualDockerRepository, error) {
	var resource VirtualDockerRepository
	err := ctx.ReadResource("artifactory:index/virtualDockerRepository:VirtualDockerRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualDockerRepository resources.
type virtualDockerRepositoryState struct {
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
	// When enabled, in cases where the same Docker tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
	ResolveDockerTagsByTimestamp *bool `pulumi:"resolveDockerTagsByTimestamp"`
}

type VirtualDockerRepositoryState struct {
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
	// When enabled, in cases where the same Docker tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
	ResolveDockerTagsByTimestamp pulumi.BoolPtrInput
}

func (VirtualDockerRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualDockerRepositoryState)(nil)).Elem()
}

type virtualDockerRepositoryArgs struct {
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
	// When enabled, in cases where the same Docker tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
	ResolveDockerTagsByTimestamp *bool `pulumi:"resolveDockerTagsByTimestamp"`
}

// The set of arguments for constructing a VirtualDockerRepository resource.
type VirtualDockerRepositoryArgs struct {
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
	// When enabled, in cases where the same Docker tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
	ResolveDockerTagsByTimestamp pulumi.BoolPtrInput
}

func (VirtualDockerRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualDockerRepositoryArgs)(nil)).Elem()
}

type VirtualDockerRepositoryInput interface {
	pulumi.Input

	ToVirtualDockerRepositoryOutput() VirtualDockerRepositoryOutput
	ToVirtualDockerRepositoryOutputWithContext(ctx context.Context) VirtualDockerRepositoryOutput
}

func (*VirtualDockerRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualDockerRepository)(nil)).Elem()
}

func (i *VirtualDockerRepository) ToVirtualDockerRepositoryOutput() VirtualDockerRepositoryOutput {
	return i.ToVirtualDockerRepositoryOutputWithContext(context.Background())
}

func (i *VirtualDockerRepository) ToVirtualDockerRepositoryOutputWithContext(ctx context.Context) VirtualDockerRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDockerRepositoryOutput)
}

// VirtualDockerRepositoryArrayInput is an input type that accepts VirtualDockerRepositoryArray and VirtualDockerRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualDockerRepositoryArrayInput` via:
//
//	VirtualDockerRepositoryArray{ VirtualDockerRepositoryArgs{...} }
type VirtualDockerRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualDockerRepositoryArrayOutput() VirtualDockerRepositoryArrayOutput
	ToVirtualDockerRepositoryArrayOutputWithContext(context.Context) VirtualDockerRepositoryArrayOutput
}

type VirtualDockerRepositoryArray []VirtualDockerRepositoryInput

func (VirtualDockerRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualDockerRepository)(nil)).Elem()
}

func (i VirtualDockerRepositoryArray) ToVirtualDockerRepositoryArrayOutput() VirtualDockerRepositoryArrayOutput {
	return i.ToVirtualDockerRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualDockerRepositoryArray) ToVirtualDockerRepositoryArrayOutputWithContext(ctx context.Context) VirtualDockerRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDockerRepositoryArrayOutput)
}

// VirtualDockerRepositoryMapInput is an input type that accepts VirtualDockerRepositoryMap and VirtualDockerRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualDockerRepositoryMapInput` via:
//
//	VirtualDockerRepositoryMap{ "key": VirtualDockerRepositoryArgs{...} }
type VirtualDockerRepositoryMapInput interface {
	pulumi.Input

	ToVirtualDockerRepositoryMapOutput() VirtualDockerRepositoryMapOutput
	ToVirtualDockerRepositoryMapOutputWithContext(context.Context) VirtualDockerRepositoryMapOutput
}

type VirtualDockerRepositoryMap map[string]VirtualDockerRepositoryInput

func (VirtualDockerRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualDockerRepository)(nil)).Elem()
}

func (i VirtualDockerRepositoryMap) ToVirtualDockerRepositoryMapOutput() VirtualDockerRepositoryMapOutput {
	return i.ToVirtualDockerRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualDockerRepositoryMap) ToVirtualDockerRepositoryMapOutputWithContext(ctx context.Context) VirtualDockerRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDockerRepositoryMapOutput)
}

type VirtualDockerRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualDockerRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualDockerRepository)(nil)).Elem()
}

func (o VirtualDockerRepositoryOutput) ToVirtualDockerRepositoryOutput() VirtualDockerRepositoryOutput {
	return o
}

func (o VirtualDockerRepositoryOutput) ToVirtualDockerRepositoryOutputWithContext(ctx context.Context) VirtualDockerRepositoryOutput {
	return o
}

// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
// another Artifactory instance.
func (o VirtualDockerRepositoryOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualDockerRepository) pulumi.BoolPtrOutput {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

// Default repository to deploy artifacts.
func (o VirtualDockerRepositoryOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualDockerRepository) pulumi.StringPtrOutput { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

// Public description.
func (o VirtualDockerRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualDockerRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o VirtualDockerRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualDockerRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o VirtualDockerRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualDockerRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
// contain spaces or special characters.
func (o VirtualDockerRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualDockerRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o VirtualDockerRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualDockerRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o VirtualDockerRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualDockerRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o VirtualDockerRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualDockerRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o VirtualDockerRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualDockerRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Repository layout key for the virtual repository
func (o VirtualDockerRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualDockerRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The effective list of actual repositories included in this virtual repository.
func (o VirtualDockerRepositoryOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualDockerRepository) pulumi.StringArrayOutput { return v.Repositories }).(pulumi.StringArrayOutput)
}

// When enabled, in cases where the same Docker tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
func (o VirtualDockerRepositoryOutput) ResolveDockerTagsByTimestamp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualDockerRepository) pulumi.BoolPtrOutput { return v.ResolveDockerTagsByTimestamp }).(pulumi.BoolPtrOutput)
}

type VirtualDockerRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualDockerRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualDockerRepository)(nil)).Elem()
}

func (o VirtualDockerRepositoryArrayOutput) ToVirtualDockerRepositoryArrayOutput() VirtualDockerRepositoryArrayOutput {
	return o
}

func (o VirtualDockerRepositoryArrayOutput) ToVirtualDockerRepositoryArrayOutputWithContext(ctx context.Context) VirtualDockerRepositoryArrayOutput {
	return o
}

func (o VirtualDockerRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualDockerRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualDockerRepository {
		return vs[0].([]*VirtualDockerRepository)[vs[1].(int)]
	}).(VirtualDockerRepositoryOutput)
}

type VirtualDockerRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualDockerRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualDockerRepository)(nil)).Elem()
}

func (o VirtualDockerRepositoryMapOutput) ToVirtualDockerRepositoryMapOutput() VirtualDockerRepositoryMapOutput {
	return o
}

func (o VirtualDockerRepositoryMapOutput) ToVirtualDockerRepositoryMapOutputWithContext(ctx context.Context) VirtualDockerRepositoryMapOutput {
	return o
}

func (o VirtualDockerRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualDockerRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualDockerRepository {
		return vs[0].(map[string]*VirtualDockerRepository)[vs[1].(string)]
	}).(VirtualDockerRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualDockerRepositoryInput)(nil)).Elem(), &VirtualDockerRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualDockerRepositoryArrayInput)(nil)).Elem(), VirtualDockerRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualDockerRepositoryMapInput)(nil)).Elem(), VirtualDockerRepositoryMap{})
	pulumi.RegisterOutputType(VirtualDockerRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualDockerRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualDockerRepositoryMapOutput{})
}
