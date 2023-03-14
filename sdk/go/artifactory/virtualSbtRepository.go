// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a virtual SBT repository.
// Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/SBT+Repositories#SBTRepositories-VirtualRepositories).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v3/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewVirtualSbtRepository(ctx, "foo-sbt", &artifactory.VirtualSbtRepositoryArgs{
//				Description:                          pulumi.String("A test virtual repo"),
//				ExcludesPattern:                      pulumi.String("com/google/**"),
//				IncludesPattern:                      pulumi.String("com/jfrog/**,cloud/jfrog/**"),
//				Key:                                  pulumi.String("foo-sbt"),
//				Notes:                                pulumi.String("Internal description"),
//				PomRepositoryReferencesCleanupPolicy: pulumi.String("discard_active_reference"),
//				Repositories:                         pulumi.StringArray{},
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
//	$ pulumi import artifactory:index/virtualSbtRepository:VirtualSbtRepository foo-sbt foo-sbt
//
// ```
type VirtualSbtRepository struct {
	pulumi.CustomResourceState

	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrOutput `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo pulumi.StringPtrOutput `pulumi:"defaultDeploymentRepo"`
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
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
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
	PackageType pulumi.StringOutput `pulumi:"packageType"`
	// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
	// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
	// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
	PomRepositoryReferencesCleanupPolicy pulumi.StringOutput `pulumi:"pomRepositoryReferencesCleanupPolicy"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrOutput `pulumi:"projectKey"`
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayOutput `pulumi:"repositories"`
}

// NewVirtualSbtRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualSbtRepository(ctx *pulumi.Context,
	name string, args *VirtualSbtRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualSbtRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource VirtualSbtRepository
	err := ctx.RegisterResource("artifactory:index/virtualSbtRepository:VirtualSbtRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualSbtRepository gets an existing VirtualSbtRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualSbtRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualSbtRepositoryState, opts ...pulumi.ResourceOption) (*VirtualSbtRepository, error) {
	var resource VirtualSbtRepository
	err := ctx.ReadResource("artifactory:index/virtualSbtRepository:VirtualSbtRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualSbtRepository resources.
type virtualSbtRepositoryState struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo *string `pulumi:"defaultDeploymentRepo"`
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
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
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes *string `pulumi:"notes"`
	// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
	PackageType *string `pulumi:"packageType"`
	// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
	// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
	// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
	PomRepositoryReferencesCleanupPolicy *string `pulumi:"pomRepositoryReferencesCleanupPolicy"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
}

type VirtualSbtRepositoryState struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo pulumi.StringPtrInput
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
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
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes pulumi.StringPtrInput
	// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
	PackageType pulumi.StringPtrInput
	// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
	// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
	// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
	PomRepositoryReferencesCleanupPolicy pulumi.StringPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
}

func (VirtualSbtRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualSbtRepositoryState)(nil)).Elem()
}

type virtualSbtRepositoryArgs struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo *string `pulumi:"defaultDeploymentRepo"`
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
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
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes *string `pulumi:"notes"`
	// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
	// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
	// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
	PomRepositoryReferencesCleanupPolicy *string `pulumi:"pomRepositoryReferencesCleanupPolicy"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
}

// The set of arguments for constructing a VirtualSbtRepository resource.
type VirtualSbtRepositoryArgs struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo pulumi.StringPtrInput
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
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
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes pulumi.StringPtrInput
	// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
	// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
	// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
	PomRepositoryReferencesCleanupPolicy pulumi.StringPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
}

func (VirtualSbtRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualSbtRepositoryArgs)(nil)).Elem()
}

type VirtualSbtRepositoryInput interface {
	pulumi.Input

	ToVirtualSbtRepositoryOutput() VirtualSbtRepositoryOutput
	ToVirtualSbtRepositoryOutputWithContext(ctx context.Context) VirtualSbtRepositoryOutput
}

func (*VirtualSbtRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualSbtRepository)(nil)).Elem()
}

func (i *VirtualSbtRepository) ToVirtualSbtRepositoryOutput() VirtualSbtRepositoryOutput {
	return i.ToVirtualSbtRepositoryOutputWithContext(context.Background())
}

func (i *VirtualSbtRepository) ToVirtualSbtRepositoryOutputWithContext(ctx context.Context) VirtualSbtRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualSbtRepositoryOutput)
}

// VirtualSbtRepositoryArrayInput is an input type that accepts VirtualSbtRepositoryArray and VirtualSbtRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualSbtRepositoryArrayInput` via:
//
//	VirtualSbtRepositoryArray{ VirtualSbtRepositoryArgs{...} }
type VirtualSbtRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualSbtRepositoryArrayOutput() VirtualSbtRepositoryArrayOutput
	ToVirtualSbtRepositoryArrayOutputWithContext(context.Context) VirtualSbtRepositoryArrayOutput
}

type VirtualSbtRepositoryArray []VirtualSbtRepositoryInput

func (VirtualSbtRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualSbtRepository)(nil)).Elem()
}

func (i VirtualSbtRepositoryArray) ToVirtualSbtRepositoryArrayOutput() VirtualSbtRepositoryArrayOutput {
	return i.ToVirtualSbtRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualSbtRepositoryArray) ToVirtualSbtRepositoryArrayOutputWithContext(ctx context.Context) VirtualSbtRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualSbtRepositoryArrayOutput)
}

// VirtualSbtRepositoryMapInput is an input type that accepts VirtualSbtRepositoryMap and VirtualSbtRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualSbtRepositoryMapInput` via:
//
//	VirtualSbtRepositoryMap{ "key": VirtualSbtRepositoryArgs{...} }
type VirtualSbtRepositoryMapInput interface {
	pulumi.Input

	ToVirtualSbtRepositoryMapOutput() VirtualSbtRepositoryMapOutput
	ToVirtualSbtRepositoryMapOutputWithContext(context.Context) VirtualSbtRepositoryMapOutput
}

type VirtualSbtRepositoryMap map[string]VirtualSbtRepositoryInput

func (VirtualSbtRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualSbtRepository)(nil)).Elem()
}

func (i VirtualSbtRepositoryMap) ToVirtualSbtRepositoryMapOutput() VirtualSbtRepositoryMapOutput {
	return i.ToVirtualSbtRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualSbtRepositoryMap) ToVirtualSbtRepositoryMapOutputWithContext(ctx context.Context) VirtualSbtRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualSbtRepositoryMapOutput)
}

type VirtualSbtRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualSbtRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualSbtRepository)(nil)).Elem()
}

func (o VirtualSbtRepositoryOutput) ToVirtualSbtRepositoryOutput() VirtualSbtRepositoryOutput {
	return o
}

func (o VirtualSbtRepositoryOutput) ToVirtualSbtRepositoryOutputWithContext(ctx context.Context) VirtualSbtRepositoryOutput {
	return o
}

// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
// another Artifactory instance.
func (o VirtualSbtRepositoryOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualSbtRepository) pulumi.BoolPtrOutput {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

// Default repository to deploy artifacts.
func (o VirtualSbtRepositoryOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualSbtRepository) pulumi.StringPtrOutput { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
func (o VirtualSbtRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualSbtRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o VirtualSbtRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualSbtRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// User authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This
// is also enforced when aggregated repositories support anonymous requests.
func (o VirtualSbtRepositoryOutput) ForceMavenAuthentication() pulumi.BoolOutput {
	return o.ApplyT(func(v *VirtualSbtRepository) pulumi.BoolOutput { return v.ForceMavenAuthentication }).(pulumi.BoolOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o VirtualSbtRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualSbtRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
// contain spaces or special characters.
func (o VirtualSbtRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualSbtRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The keypair used to sign artifacts.
func (o VirtualSbtRepositoryOutput) KeyPair() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualSbtRepository) pulumi.StringPtrOutput { return v.KeyPair }).(pulumi.StringPtrOutput)
}

// A free text field to add additional notes about the repository. These are only visible to the administrator.
func (o VirtualSbtRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualSbtRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
func (o VirtualSbtRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualSbtRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
func (o VirtualSbtRepositoryOutput) PomRepositoryReferencesCleanupPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualSbtRepository) pulumi.StringOutput { return v.PomRepositoryReferencesCleanupPolicy }).(pulumi.StringOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
// will remain in the Terraform state, which will create state drift during the update.
func (o VirtualSbtRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualSbtRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o VirtualSbtRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualSbtRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Repository layout key for the local repository
func (o VirtualSbtRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualSbtRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The effective list of actual repositories included in this virtual repository.
func (o VirtualSbtRepositoryOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualSbtRepository) pulumi.StringArrayOutput { return v.Repositories }).(pulumi.StringArrayOutput)
}

type VirtualSbtRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualSbtRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualSbtRepository)(nil)).Elem()
}

func (o VirtualSbtRepositoryArrayOutput) ToVirtualSbtRepositoryArrayOutput() VirtualSbtRepositoryArrayOutput {
	return o
}

func (o VirtualSbtRepositoryArrayOutput) ToVirtualSbtRepositoryArrayOutputWithContext(ctx context.Context) VirtualSbtRepositoryArrayOutput {
	return o
}

func (o VirtualSbtRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualSbtRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualSbtRepository {
		return vs[0].([]*VirtualSbtRepository)[vs[1].(int)]
	}).(VirtualSbtRepositoryOutput)
}

type VirtualSbtRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualSbtRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualSbtRepository)(nil)).Elem()
}

func (o VirtualSbtRepositoryMapOutput) ToVirtualSbtRepositoryMapOutput() VirtualSbtRepositoryMapOutput {
	return o
}

func (o VirtualSbtRepositoryMapOutput) ToVirtualSbtRepositoryMapOutputWithContext(ctx context.Context) VirtualSbtRepositoryMapOutput {
	return o
}

func (o VirtualSbtRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualSbtRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualSbtRepository {
		return vs[0].(map[string]*VirtualSbtRepository)[vs[1].(string)]
	}).(VirtualSbtRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualSbtRepositoryInput)(nil)).Elem(), &VirtualSbtRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualSbtRepositoryArrayInput)(nil)).Elem(), VirtualSbtRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualSbtRepositoryMapInput)(nil)).Elem(), VirtualSbtRepositoryMap{})
	pulumi.RegisterOutputType(VirtualSbtRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualSbtRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualSbtRepositoryMapOutput{})
}
