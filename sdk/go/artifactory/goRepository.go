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

// Creates a virtual Go repository.
// Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Go+Registry#GoRegistry-VirtualRepositories).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewGoRepository(ctx, "baz-go", &artifactory.GoRepositoryArgs{
//				Description:                 pulumi.String("A test virtual repo"),
//				ExcludesPattern:             pulumi.String("com/google/**"),
//				ExternalDependenciesEnabled: pulumi.Bool(true),
//				ExternalDependenciesPatterns: pulumi.StringArray{
//					pulumi.String("**/github.com/**"),
//					pulumi.String("**/go.googlesource.com/**"),
//				},
//				IncludesPattern: pulumi.String("com/jfrog/**,cloud/jfrog/**"),
//				Key:             pulumi.String("baz-go"),
//				Notes:           pulumi.String("Internal description"),
//				RepoLayoutRef:   pulumi.String("go-default"),
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
//	$ pulumi import artifactory:index/goRepository:GoRepository baz-go baz-go
//
// ```
type GoRepository struct {
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
	// Shorthand for "Enable 'go-import' Meta Tags" on the UI. This must be set to true in order to use the allow list.
	// When checked (default), Artifactory will automatically follow remote VCS roots in 'go-import' meta tags to download remote modules.
	ExternalDependenciesEnabled pulumi.BoolPtrOutput `pulumi:"externalDependenciesEnabled"`
	// 'go-import' Allow List on the UI.
	ExternalDependenciesPatterns pulumi.StringArrayOutput `pulumi:"externalDependenciesPatterns"`
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

// NewGoRepository registers a new resource with the given unique name, arguments, and options.
func NewGoRepository(ctx *pulumi.Context,
	name string, args *GoRepositoryArgs, opts ...pulumi.ResourceOption) (*GoRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GoRepository
	err := ctx.RegisterResource("artifactory:index/goRepository:GoRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGoRepository gets an existing GoRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGoRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GoRepositoryState, opts ...pulumi.ResourceOption) (*GoRepository, error) {
	var resource GoRepository
	err := ctx.ReadResource("artifactory:index/goRepository:GoRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GoRepository resources.
type goRepositoryState struct {
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
	// Shorthand for "Enable 'go-import' Meta Tags" on the UI. This must be set to true in order to use the allow list.
	// When checked (default), Artifactory will automatically follow remote VCS roots in 'go-import' meta tags to download remote modules.
	ExternalDependenciesEnabled *bool `pulumi:"externalDependenciesEnabled"`
	// 'go-import' Allow List on the UI.
	ExternalDependenciesPatterns []string `pulumi:"externalDependenciesPatterns"`
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

type GoRepositoryState struct {
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
	// Shorthand for "Enable 'go-import' Meta Tags" on the UI. This must be set to true in order to use the allow list.
	// When checked (default), Artifactory will automatically follow remote VCS roots in 'go-import' meta tags to download remote modules.
	ExternalDependenciesEnabled pulumi.BoolPtrInput
	// 'go-import' Allow List on the UI.
	ExternalDependenciesPatterns pulumi.StringArrayInput
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

func (GoRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*goRepositoryState)(nil)).Elem()
}

type goRepositoryArgs struct {
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
	// Shorthand for "Enable 'go-import' Meta Tags" on the UI. This must be set to true in order to use the allow list.
	// When checked (default), Artifactory will automatically follow remote VCS roots in 'go-import' meta tags to download remote modules.
	ExternalDependenciesEnabled *bool `pulumi:"externalDependenciesEnabled"`
	// 'go-import' Allow List on the UI.
	ExternalDependenciesPatterns []string `pulumi:"externalDependenciesPatterns"`
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

// The set of arguments for constructing a GoRepository resource.
type GoRepositoryArgs struct {
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
	// Shorthand for "Enable 'go-import' Meta Tags" on the UI. This must be set to true in order to use the allow list.
	// When checked (default), Artifactory will automatically follow remote VCS roots in 'go-import' meta tags to download remote modules.
	ExternalDependenciesEnabled pulumi.BoolPtrInput
	// 'go-import' Allow List on the UI.
	ExternalDependenciesPatterns pulumi.StringArrayInput
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

func (GoRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*goRepositoryArgs)(nil)).Elem()
}

type GoRepositoryInput interface {
	pulumi.Input

	ToGoRepositoryOutput() GoRepositoryOutput
	ToGoRepositoryOutputWithContext(ctx context.Context) GoRepositoryOutput
}

func (*GoRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**GoRepository)(nil)).Elem()
}

func (i *GoRepository) ToGoRepositoryOutput() GoRepositoryOutput {
	return i.ToGoRepositoryOutputWithContext(context.Background())
}

func (i *GoRepository) ToGoRepositoryOutputWithContext(ctx context.Context) GoRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoRepositoryOutput)
}

func (i *GoRepository) ToOutput(ctx context.Context) pulumix.Output[*GoRepository] {
	return pulumix.Output[*GoRepository]{
		OutputState: i.ToGoRepositoryOutputWithContext(ctx).OutputState,
	}
}

// GoRepositoryArrayInput is an input type that accepts GoRepositoryArray and GoRepositoryArrayOutput values.
// You can construct a concrete instance of `GoRepositoryArrayInput` via:
//
//	GoRepositoryArray{ GoRepositoryArgs{...} }
type GoRepositoryArrayInput interface {
	pulumi.Input

	ToGoRepositoryArrayOutput() GoRepositoryArrayOutput
	ToGoRepositoryArrayOutputWithContext(context.Context) GoRepositoryArrayOutput
}

type GoRepositoryArray []GoRepositoryInput

func (GoRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GoRepository)(nil)).Elem()
}

func (i GoRepositoryArray) ToGoRepositoryArrayOutput() GoRepositoryArrayOutput {
	return i.ToGoRepositoryArrayOutputWithContext(context.Background())
}

func (i GoRepositoryArray) ToGoRepositoryArrayOutputWithContext(ctx context.Context) GoRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoRepositoryArrayOutput)
}

func (i GoRepositoryArray) ToOutput(ctx context.Context) pulumix.Output[[]*GoRepository] {
	return pulumix.Output[[]*GoRepository]{
		OutputState: i.ToGoRepositoryArrayOutputWithContext(ctx).OutputState,
	}
}

// GoRepositoryMapInput is an input type that accepts GoRepositoryMap and GoRepositoryMapOutput values.
// You can construct a concrete instance of `GoRepositoryMapInput` via:
//
//	GoRepositoryMap{ "key": GoRepositoryArgs{...} }
type GoRepositoryMapInput interface {
	pulumi.Input

	ToGoRepositoryMapOutput() GoRepositoryMapOutput
	ToGoRepositoryMapOutputWithContext(context.Context) GoRepositoryMapOutput
}

type GoRepositoryMap map[string]GoRepositoryInput

func (GoRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GoRepository)(nil)).Elem()
}

func (i GoRepositoryMap) ToGoRepositoryMapOutput() GoRepositoryMapOutput {
	return i.ToGoRepositoryMapOutputWithContext(context.Background())
}

func (i GoRepositoryMap) ToGoRepositoryMapOutputWithContext(ctx context.Context) GoRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoRepositoryMapOutput)
}

func (i GoRepositoryMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*GoRepository] {
	return pulumix.Output[map[string]*GoRepository]{
		OutputState: i.ToGoRepositoryMapOutputWithContext(ctx).OutputState,
	}
}

type GoRepositoryOutput struct{ *pulumi.OutputState }

func (GoRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GoRepository)(nil)).Elem()
}

func (o GoRepositoryOutput) ToGoRepositoryOutput() GoRepositoryOutput {
	return o
}

func (o GoRepositoryOutput) ToGoRepositoryOutputWithContext(ctx context.Context) GoRepositoryOutput {
	return o
}

func (o GoRepositoryOutput) ToOutput(ctx context.Context) pulumix.Output[*GoRepository] {
	return pulumix.Output[*GoRepository]{
		OutputState: o.OutputState,
	}
}

// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
// another Artifactory instance.
func (o GoRepositoryOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GoRepository) pulumi.BoolPtrOutput { return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts }).(pulumi.BoolPtrOutput)
}

// Default repository to deploy artifacts.
func (o GoRepositoryOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GoRepository) pulumi.StringPtrOutput { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

// Public description.
func (o GoRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GoRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o GoRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GoRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// Shorthand for "Enable 'go-import' Meta Tags" on the UI. This must be set to true in order to use the allow list.
// When checked (default), Artifactory will automatically follow remote VCS roots in 'go-import' meta tags to download remote modules.
func (o GoRepositoryOutput) ExternalDependenciesEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GoRepository) pulumi.BoolPtrOutput { return v.ExternalDependenciesEnabled }).(pulumi.BoolPtrOutput)
}

// 'go-import' Allow List on the UI.
func (o GoRepositoryOutput) ExternalDependenciesPatterns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GoRepository) pulumi.StringArrayOutput { return v.ExternalDependenciesPatterns }).(pulumi.StringArrayOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o GoRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GoRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
// contain spaces or special characters.
func (o GoRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *GoRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o GoRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GoRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o GoRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *GoRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o GoRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GoRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o GoRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GoRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Repository layout key for the virtual repository
func (o GoRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GoRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The effective list of actual repositories included in this virtual repository.
func (o GoRepositoryOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GoRepository) pulumi.StringArrayOutput { return v.Repositories }).(pulumi.StringArrayOutput)
}

type GoRepositoryArrayOutput struct{ *pulumi.OutputState }

func (GoRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GoRepository)(nil)).Elem()
}

func (o GoRepositoryArrayOutput) ToGoRepositoryArrayOutput() GoRepositoryArrayOutput {
	return o
}

func (o GoRepositoryArrayOutput) ToGoRepositoryArrayOutputWithContext(ctx context.Context) GoRepositoryArrayOutput {
	return o
}

func (o GoRepositoryArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*GoRepository] {
	return pulumix.Output[[]*GoRepository]{
		OutputState: o.OutputState,
	}
}

func (o GoRepositoryArrayOutput) Index(i pulumi.IntInput) GoRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GoRepository {
		return vs[0].([]*GoRepository)[vs[1].(int)]
	}).(GoRepositoryOutput)
}

type GoRepositoryMapOutput struct{ *pulumi.OutputState }

func (GoRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GoRepository)(nil)).Elem()
}

func (o GoRepositoryMapOutput) ToGoRepositoryMapOutput() GoRepositoryMapOutput {
	return o
}

func (o GoRepositoryMapOutput) ToGoRepositoryMapOutputWithContext(ctx context.Context) GoRepositoryMapOutput {
	return o
}

func (o GoRepositoryMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*GoRepository] {
	return pulumix.Output[map[string]*GoRepository]{
		OutputState: o.OutputState,
	}
}

func (o GoRepositoryMapOutput) MapIndex(k pulumi.StringInput) GoRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GoRepository {
		return vs[0].(map[string]*GoRepository)[vs[1].(string)]
	}).(GoRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GoRepositoryInput)(nil)).Elem(), &GoRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*GoRepositoryArrayInput)(nil)).Elem(), GoRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GoRepositoryMapInput)(nil)).Elem(), GoRepositoryMap{})
	pulumi.RegisterOutputType(GoRepositoryOutput{})
	pulumi.RegisterOutputType(GoRepositoryArrayOutput{})
	pulumi.RegisterOutputType(GoRepositoryMapOutput{})
}
