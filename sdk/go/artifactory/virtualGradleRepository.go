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

// Creates a virtual Gradle repository.
// Official documentation can be found [here](https://jfrog.com/blog/how-to-set-up-a-private-remote-and-virtual-maven-gradle-registry/).
//
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
//			_, err := artifactory.NewVirtualGradleRepository(ctx, "foo-gradle", &artifactory.VirtualGradleRepositoryArgs{
//				Key:                                  pulumi.String("foo-gradle"),
//				Repositories:                         pulumi.StringArray{},
//				Description:                          pulumi.String("A test virtual repo"),
//				Notes:                                pulumi.String("Internal description"),
//				IncludesPattern:                      pulumi.String("com/jfrog/**,cloud/jfrog/**"),
//				ExcludesPattern:                      pulumi.String("com/google/**"),
//				PomRepositoryReferencesCleanupPolicy: pulumi.String("discard_active_reference"),
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
// $ pulumi import artifactory:index/virtualGradleRepository:VirtualGradleRepository foo-gradle foo-gradle
// ```
type VirtualGradleRepository struct {
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
	// User authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This
	// is also enforced when aggregated repositories support anonymous requests.
	ForceMavenAuthentication pulumi.BoolOutput `pulumi:"forceMavenAuthentication"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrOutput `pulumi:"includesPattern"`
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
	Key pulumi.StringOutput `pulumi:"key"`
	// The keypair used to sign artifacts.
	KeyPair pulumi.StringPtrOutput `pulumi:"keyPair"`
	// Internal description.
	Notes       pulumi.StringPtrOutput `pulumi:"notes"`
	PackageType pulumi.StringOutput    `pulumi:"packageType"`
	// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
	// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
	// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
	PomRepositoryReferencesCleanupPolicy pulumi.StringOutput `pulumi:"pomRepositoryReferencesCleanupPolicy"`
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

// NewVirtualGradleRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualGradleRepository(ctx *pulumi.Context,
	name string, args *VirtualGradleRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualGradleRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VirtualGradleRepository
	err := ctx.RegisterResource("artifactory:index/virtualGradleRepository:VirtualGradleRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualGradleRepository gets an existing VirtualGradleRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualGradleRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualGradleRepositoryState, opts ...pulumi.ResourceOption) (*VirtualGradleRepository, error) {
	var resource VirtualGradleRepository
	err := ctx.ReadResource("artifactory:index/virtualGradleRepository:VirtualGradleRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualGradleRepository resources.
type virtualGradleRepositoryState struct {
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
	// User authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This
	// is also enforced when aggregated repositories support anonymous requests.
	ForceMavenAuthentication *bool `pulumi:"forceMavenAuthentication"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
	Key *string `pulumi:"key"`
	// The keypair used to sign artifacts.
	KeyPair *string `pulumi:"keyPair"`
	// Internal description.
	Notes       *string `pulumi:"notes"`
	PackageType *string `pulumi:"packageType"`
	// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
	// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
	// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
	PomRepositoryReferencesCleanupPolicy *string `pulumi:"pomRepositoryReferencesCleanupPolicy"`
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

type VirtualGradleRepositoryState struct {
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
	// User authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This
	// is also enforced when aggregated repositories support anonymous requests.
	ForceMavenAuthentication pulumi.BoolPtrInput
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
	Key pulumi.StringPtrInput
	// The keypair used to sign artifacts.
	KeyPair pulumi.StringPtrInput
	// Internal description.
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
	// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
	// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
	PomRepositoryReferencesCleanupPolicy pulumi.StringPtrInput
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

func (VirtualGradleRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualGradleRepositoryState)(nil)).Elem()
}

type virtualGradleRepositoryArgs struct {
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
	// User authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This
	// is also enforced when aggregated repositories support anonymous requests.
	ForceMavenAuthentication *bool `pulumi:"forceMavenAuthentication"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
	Key string `pulumi:"key"`
	// The keypair used to sign artifacts.
	KeyPair *string `pulumi:"keyPair"`
	// Internal description.
	Notes *string `pulumi:"notes"`
	// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
	// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
	// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
	PomRepositoryReferencesCleanupPolicy *string `pulumi:"pomRepositoryReferencesCleanupPolicy"`
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

// The set of arguments for constructing a VirtualGradleRepository resource.
type VirtualGradleRepositoryArgs struct {
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
	// User authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This
	// is also enforced when aggregated repositories support anonymous requests.
	ForceMavenAuthentication pulumi.BoolPtrInput
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
	Key pulumi.StringInput
	// The keypair used to sign artifacts.
	KeyPair pulumi.StringPtrInput
	// Internal description.
	Notes pulumi.StringPtrInput
	// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
	// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
	// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
	PomRepositoryReferencesCleanupPolicy pulumi.StringPtrInput
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

func (VirtualGradleRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualGradleRepositoryArgs)(nil)).Elem()
}

type VirtualGradleRepositoryInput interface {
	pulumi.Input

	ToVirtualGradleRepositoryOutput() VirtualGradleRepositoryOutput
	ToVirtualGradleRepositoryOutputWithContext(ctx context.Context) VirtualGradleRepositoryOutput
}

func (*VirtualGradleRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualGradleRepository)(nil)).Elem()
}

func (i *VirtualGradleRepository) ToVirtualGradleRepositoryOutput() VirtualGradleRepositoryOutput {
	return i.ToVirtualGradleRepositoryOutputWithContext(context.Background())
}

func (i *VirtualGradleRepository) ToVirtualGradleRepositoryOutputWithContext(ctx context.Context) VirtualGradleRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualGradleRepositoryOutput)
}

// VirtualGradleRepositoryArrayInput is an input type that accepts VirtualGradleRepositoryArray and VirtualGradleRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualGradleRepositoryArrayInput` via:
//
//	VirtualGradleRepositoryArray{ VirtualGradleRepositoryArgs{...} }
type VirtualGradleRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualGradleRepositoryArrayOutput() VirtualGradleRepositoryArrayOutput
	ToVirtualGradleRepositoryArrayOutputWithContext(context.Context) VirtualGradleRepositoryArrayOutput
}

type VirtualGradleRepositoryArray []VirtualGradleRepositoryInput

func (VirtualGradleRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualGradleRepository)(nil)).Elem()
}

func (i VirtualGradleRepositoryArray) ToVirtualGradleRepositoryArrayOutput() VirtualGradleRepositoryArrayOutput {
	return i.ToVirtualGradleRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualGradleRepositoryArray) ToVirtualGradleRepositoryArrayOutputWithContext(ctx context.Context) VirtualGradleRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualGradleRepositoryArrayOutput)
}

// VirtualGradleRepositoryMapInput is an input type that accepts VirtualGradleRepositoryMap and VirtualGradleRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualGradleRepositoryMapInput` via:
//
//	VirtualGradleRepositoryMap{ "key": VirtualGradleRepositoryArgs{...} }
type VirtualGradleRepositoryMapInput interface {
	pulumi.Input

	ToVirtualGradleRepositoryMapOutput() VirtualGradleRepositoryMapOutput
	ToVirtualGradleRepositoryMapOutputWithContext(context.Context) VirtualGradleRepositoryMapOutput
}

type VirtualGradleRepositoryMap map[string]VirtualGradleRepositoryInput

func (VirtualGradleRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualGradleRepository)(nil)).Elem()
}

func (i VirtualGradleRepositoryMap) ToVirtualGradleRepositoryMapOutput() VirtualGradleRepositoryMapOutput {
	return i.ToVirtualGradleRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualGradleRepositoryMap) ToVirtualGradleRepositoryMapOutputWithContext(ctx context.Context) VirtualGradleRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualGradleRepositoryMapOutput)
}

type VirtualGradleRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualGradleRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualGradleRepository)(nil)).Elem()
}

func (o VirtualGradleRepositoryOutput) ToVirtualGradleRepositoryOutput() VirtualGradleRepositoryOutput {
	return o
}

func (o VirtualGradleRepositoryOutput) ToVirtualGradleRepositoryOutputWithContext(ctx context.Context) VirtualGradleRepositoryOutput {
	return o
}

// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
// another Artifactory instance.
func (o VirtualGradleRepositoryOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualGradleRepository) pulumi.BoolPtrOutput {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

// Default repository to deploy artifacts.
func (o VirtualGradleRepositoryOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualGradleRepository) pulumi.StringPtrOutput { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

// Public description.
func (o VirtualGradleRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualGradleRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o VirtualGradleRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualGradleRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// User authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This
// is also enforced when aggregated repositories support anonymous requests.
func (o VirtualGradleRepositoryOutput) ForceMavenAuthentication() pulumi.BoolOutput {
	return o.ApplyT(func(v *VirtualGradleRepository) pulumi.BoolOutput { return v.ForceMavenAuthentication }).(pulumi.BoolOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o VirtualGradleRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualGradleRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
// contain spaces or special characters.
func (o VirtualGradleRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualGradleRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The keypair used to sign artifacts.
func (o VirtualGradleRepositoryOutput) KeyPair() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualGradleRepository) pulumi.StringPtrOutput { return v.KeyPair }).(pulumi.StringPtrOutput)
}

// Internal description.
func (o VirtualGradleRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualGradleRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o VirtualGradleRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualGradleRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
func (o VirtualGradleRepositoryOutput) PomRepositoryReferencesCleanupPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualGradleRepository) pulumi.StringOutput { return v.PomRepositoryReferencesCleanupPolicy }).(pulumi.StringOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o VirtualGradleRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualGradleRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o VirtualGradleRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualGradleRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Repository layout key for the virtual repository
func (o VirtualGradleRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualGradleRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The effective list of actual repositories included in this virtual repository.
func (o VirtualGradleRepositoryOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualGradleRepository) pulumi.StringArrayOutput { return v.Repositories }).(pulumi.StringArrayOutput)
}

type VirtualGradleRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualGradleRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualGradleRepository)(nil)).Elem()
}

func (o VirtualGradleRepositoryArrayOutput) ToVirtualGradleRepositoryArrayOutput() VirtualGradleRepositoryArrayOutput {
	return o
}

func (o VirtualGradleRepositoryArrayOutput) ToVirtualGradleRepositoryArrayOutputWithContext(ctx context.Context) VirtualGradleRepositoryArrayOutput {
	return o
}

func (o VirtualGradleRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualGradleRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualGradleRepository {
		return vs[0].([]*VirtualGradleRepository)[vs[1].(int)]
	}).(VirtualGradleRepositoryOutput)
}

type VirtualGradleRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualGradleRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualGradleRepository)(nil)).Elem()
}

func (o VirtualGradleRepositoryMapOutput) ToVirtualGradleRepositoryMapOutput() VirtualGradleRepositoryMapOutput {
	return o
}

func (o VirtualGradleRepositoryMapOutput) ToVirtualGradleRepositoryMapOutputWithContext(ctx context.Context) VirtualGradleRepositoryMapOutput {
	return o
}

func (o VirtualGradleRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualGradleRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualGradleRepository {
		return vs[0].(map[string]*VirtualGradleRepository)[vs[1].(string)]
	}).(VirtualGradleRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualGradleRepositoryInput)(nil)).Elem(), &VirtualGradleRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualGradleRepositoryArrayInput)(nil)).Elem(), VirtualGradleRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualGradleRepositoryMapInput)(nil)).Elem(), VirtualGradleRepositoryMap{})
	pulumi.RegisterOutputType(VirtualGradleRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualGradleRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualGradleRepositoryMapOutput{})
}
