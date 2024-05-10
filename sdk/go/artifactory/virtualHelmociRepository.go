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

// Creates a virtual Helm OCI repository.
//
// Official documentation can be found [here](https://jfrog.com/help/r/jfrog-artifactory-documentation/helm-oci-repositories).
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
//			_, err := artifactory.NewVirtualHelmociRepository(ctx, "my-helmoci-virtual", &artifactory.VirtualHelmociRepositoryArgs{
//				Key:                       pulumi.String("my-helmoci-virtual"),
//				Repositories:              pulumi.StringArray{},
//				Description:               pulumi.String("A test virtual repo"),
//				Notes:                     pulumi.String("Internal description"),
//				IncludesPattern:           pulumi.String("com/jfrog/**,cloud/jfrog/**"),
//				ExcludesPattern:           pulumi.String("com/google/**"),
//				ResolveOciTagsByTimestamp: pulumi.Bool(true),
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
// Virtual Helm OCI repositories can be imported using their name, e.g.
//
// ```sh
// $ pulumi import artifactory:index/virtualHelmociRepository:VirtualHelmociRepository my-helmoci-virtual my-helmoci-virtual
// ```
type VirtualHelmociRepository struct {
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
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.
	Key pulumi.StringOutput `pulumi:"key"`
	// Internal description.
	Notes               pulumi.StringPtrOutput   `pulumi:"notes"`
	PackageType         pulumi.StringOutput      `pulumi:"packageType"`
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrOutput `pulumi:"projectKey"`
	// Repository layout key for the virtual repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayOutput `pulumi:"repositories"`
	// When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
	ResolveOciTagsByTimestamp pulumi.BoolPtrOutput `pulumi:"resolveOciTagsByTimestamp"`
}

// NewVirtualHelmociRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualHelmociRepository(ctx *pulumi.Context,
	name string, args *VirtualHelmociRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualHelmociRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VirtualHelmociRepository
	err := ctx.RegisterResource("artifactory:index/virtualHelmociRepository:VirtualHelmociRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualHelmociRepository gets an existing VirtualHelmociRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualHelmociRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualHelmociRepositoryState, opts ...pulumi.ResourceOption) (*VirtualHelmociRepository, error) {
	var resource VirtualHelmociRepository
	err := ctx.ReadResource("artifactory:index/virtualHelmociRepository:VirtualHelmociRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualHelmociRepository resources.
type virtualHelmociRepositoryState struct {
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
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.
	Key *string `pulumi:"key"`
	// Internal description.
	Notes               *string  `pulumi:"notes"`
	PackageType         *string  `pulumi:"packageType"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// Repository layout key for the virtual repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
	// When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
	ResolveOciTagsByTimestamp *bool `pulumi:"resolveOciTagsByTimestamp"`
}

type VirtualHelmociRepositoryState struct {
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
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.
	Key pulumi.StringPtrInput
	// Internal description.
	Notes               pulumi.StringPtrInput
	PackageType         pulumi.StringPtrInput
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// Repository layout key for the virtual repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
	// When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
	ResolveOciTagsByTimestamp pulumi.BoolPtrInput
}

func (VirtualHelmociRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHelmociRepositoryState)(nil)).Elem()
}

type virtualHelmociRepositoryArgs struct {
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
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.
	Key string `pulumi:"key"`
	// Internal description.
	Notes               *string  `pulumi:"notes"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// Repository layout key for the virtual repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
	// When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
	ResolveOciTagsByTimestamp *bool `pulumi:"resolveOciTagsByTimestamp"`
}

// The set of arguments for constructing a VirtualHelmociRepository resource.
type VirtualHelmociRepositoryArgs struct {
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
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.
	Key pulumi.StringInput
	// Internal description.
	Notes               pulumi.StringPtrInput
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// Repository layout key for the virtual repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
	// When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
	ResolveOciTagsByTimestamp pulumi.BoolPtrInput
}

func (VirtualHelmociRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHelmociRepositoryArgs)(nil)).Elem()
}

type VirtualHelmociRepositoryInput interface {
	pulumi.Input

	ToVirtualHelmociRepositoryOutput() VirtualHelmociRepositoryOutput
	ToVirtualHelmociRepositoryOutputWithContext(ctx context.Context) VirtualHelmociRepositoryOutput
}

func (*VirtualHelmociRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHelmociRepository)(nil)).Elem()
}

func (i *VirtualHelmociRepository) ToVirtualHelmociRepositoryOutput() VirtualHelmociRepositoryOutput {
	return i.ToVirtualHelmociRepositoryOutputWithContext(context.Background())
}

func (i *VirtualHelmociRepository) ToVirtualHelmociRepositoryOutputWithContext(ctx context.Context) VirtualHelmociRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHelmociRepositoryOutput)
}

// VirtualHelmociRepositoryArrayInput is an input type that accepts VirtualHelmociRepositoryArray and VirtualHelmociRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualHelmociRepositoryArrayInput` via:
//
//	VirtualHelmociRepositoryArray{ VirtualHelmociRepositoryArgs{...} }
type VirtualHelmociRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualHelmociRepositoryArrayOutput() VirtualHelmociRepositoryArrayOutput
	ToVirtualHelmociRepositoryArrayOutputWithContext(context.Context) VirtualHelmociRepositoryArrayOutput
}

type VirtualHelmociRepositoryArray []VirtualHelmociRepositoryInput

func (VirtualHelmociRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualHelmociRepository)(nil)).Elem()
}

func (i VirtualHelmociRepositoryArray) ToVirtualHelmociRepositoryArrayOutput() VirtualHelmociRepositoryArrayOutput {
	return i.ToVirtualHelmociRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualHelmociRepositoryArray) ToVirtualHelmociRepositoryArrayOutputWithContext(ctx context.Context) VirtualHelmociRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHelmociRepositoryArrayOutput)
}

// VirtualHelmociRepositoryMapInput is an input type that accepts VirtualHelmociRepositoryMap and VirtualHelmociRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualHelmociRepositoryMapInput` via:
//
//	VirtualHelmociRepositoryMap{ "key": VirtualHelmociRepositoryArgs{...} }
type VirtualHelmociRepositoryMapInput interface {
	pulumi.Input

	ToVirtualHelmociRepositoryMapOutput() VirtualHelmociRepositoryMapOutput
	ToVirtualHelmociRepositoryMapOutputWithContext(context.Context) VirtualHelmociRepositoryMapOutput
}

type VirtualHelmociRepositoryMap map[string]VirtualHelmociRepositoryInput

func (VirtualHelmociRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualHelmociRepository)(nil)).Elem()
}

func (i VirtualHelmociRepositoryMap) ToVirtualHelmociRepositoryMapOutput() VirtualHelmociRepositoryMapOutput {
	return i.ToVirtualHelmociRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualHelmociRepositoryMap) ToVirtualHelmociRepositoryMapOutputWithContext(ctx context.Context) VirtualHelmociRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHelmociRepositoryMapOutput)
}

type VirtualHelmociRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualHelmociRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHelmociRepository)(nil)).Elem()
}

func (o VirtualHelmociRepositoryOutput) ToVirtualHelmociRepositoryOutput() VirtualHelmociRepositoryOutput {
	return o
}

func (o VirtualHelmociRepositoryOutput) ToVirtualHelmociRepositoryOutputWithContext(ctx context.Context) VirtualHelmociRepositoryOutput {
	return o
}

// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
// another Artifactory instance.
func (o VirtualHelmociRepositoryOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualHelmociRepository) pulumi.BoolPtrOutput {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

// Default repository to deploy artifacts.
func (o VirtualHelmociRepositoryOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHelmociRepository) pulumi.StringPtrOutput { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

// Public description.
func (o VirtualHelmociRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHelmociRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o VirtualHelmociRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHelmociRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o VirtualHelmociRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHelmociRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.
func (o VirtualHelmociRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHelmociRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o VirtualHelmociRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHelmociRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o VirtualHelmociRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHelmociRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

func (o VirtualHelmociRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualHelmociRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o VirtualHelmociRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHelmociRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Repository layout key for the virtual repository
func (o VirtualHelmociRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHelmociRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The effective list of actual repositories included in this virtual repository.
func (o VirtualHelmociRepositoryOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualHelmociRepository) pulumi.StringArrayOutput { return v.Repositories }).(pulumi.StringArrayOutput)
}

// When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
func (o VirtualHelmociRepositoryOutput) ResolveOciTagsByTimestamp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualHelmociRepository) pulumi.BoolPtrOutput { return v.ResolveOciTagsByTimestamp }).(pulumi.BoolPtrOutput)
}

type VirtualHelmociRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualHelmociRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualHelmociRepository)(nil)).Elem()
}

func (o VirtualHelmociRepositoryArrayOutput) ToVirtualHelmociRepositoryArrayOutput() VirtualHelmociRepositoryArrayOutput {
	return o
}

func (o VirtualHelmociRepositoryArrayOutput) ToVirtualHelmociRepositoryArrayOutputWithContext(ctx context.Context) VirtualHelmociRepositoryArrayOutput {
	return o
}

func (o VirtualHelmociRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualHelmociRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualHelmociRepository {
		return vs[0].([]*VirtualHelmociRepository)[vs[1].(int)]
	}).(VirtualHelmociRepositoryOutput)
}

type VirtualHelmociRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualHelmociRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualHelmociRepository)(nil)).Elem()
}

func (o VirtualHelmociRepositoryMapOutput) ToVirtualHelmociRepositoryMapOutput() VirtualHelmociRepositoryMapOutput {
	return o
}

func (o VirtualHelmociRepositoryMapOutput) ToVirtualHelmociRepositoryMapOutputWithContext(ctx context.Context) VirtualHelmociRepositoryMapOutput {
	return o
}

func (o VirtualHelmociRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualHelmociRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualHelmociRepository {
		return vs[0].(map[string]*VirtualHelmociRepository)[vs[1].(string)]
	}).(VirtualHelmociRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualHelmociRepositoryInput)(nil)).Elem(), &VirtualHelmociRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualHelmociRepositoryArrayInput)(nil)).Elem(), VirtualHelmociRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualHelmociRepositoryMapInput)(nil)).Elem(), VirtualHelmociRepositoryMap{})
	pulumi.RegisterOutputType(VirtualHelmociRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualHelmociRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualHelmociRepositoryMapOutput{})
}
