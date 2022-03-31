// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Virtual Gradle Repository Resource
//
// Provides an Artifactory virtual repository resource, but with specific gradle features. This should be preferred over the original
// one-size-fits-all `VirtualRepository`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-artifactory/sdk/v2/go/artifactory"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := artifactory.NewVirtualGradleRepository(ctx, "foo-gradle", &artifactory.VirtualGradleRepositoryArgs{
// 			Description:                          pulumi.String("A test virtual repo"),
// 			ExcludesPattern:                      pulumi.String("com/google/**"),
// 			IncludesPattern:                      pulumi.String("com/jfrog/**,cloud/jfrog/**"),
// 			Key:                                  pulumi.String("foo-gradle"),
// 			Notes:                                pulumi.String("Internal description"),
// 			PomRepositoryReferencesCleanupPolicy: pulumi.String("discard_active_reference"),
// 			Repositories:                         pulumi.StringArray{},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Virtual repositories can be imported using their name, e.g.
//
// ```sh
//  $ pulumi import artifactory:index/virtualGradleRepository:VirtualGradleRepository foo foo
// ```
type VirtualGradleRepository struct {
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
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrOutput `pulumi:"includesPattern"`
	// The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
	// contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
	// 'libs-release-local').
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
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
	// repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrOutput `pulumi:"projectKey"`
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayOutput `pulumi:"repositories"`
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
	// repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds pulumi.IntPtrOutput `pulumi:"retrievalCachePeriodSeconds"`
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
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
	Description *string `pulumi:"description"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// User authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This
	// is also enforced when aggregated repositories support anonymous requests.
	ForceMavenAuthentication *bool `pulumi:"forceMavenAuthentication"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
	// contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
	// 'libs-release-local').
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
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
	// repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
	// repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds *int `pulumi:"retrievalCachePeriodSeconds"`
}

type VirtualGradleRepositoryState struct {
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
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
	// contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
	// 'libs-release-local').
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
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
	// repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
	// repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds pulumi.IntPtrInput
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
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
	Description *string `pulumi:"description"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// User authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This
	// is also enforced when aggregated repositories support anonymous requests.
	ForceMavenAuthentication *bool `pulumi:"forceMavenAuthentication"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
	// contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
	// 'libs-release-local').
	Key string `pulumi:"key"`
	// The keypair used to sign artifacts.
	KeyPair *string `pulumi:"keyPair"`
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes *string `pulumi:"notes"`
	// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
	// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
	// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
	PomRepositoryReferencesCleanupPolicy *string `pulumi:"pomRepositoryReferencesCleanupPolicy"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
	// repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
	// repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds *int `pulumi:"retrievalCachePeriodSeconds"`
}

// The set of arguments for constructing a VirtualGradleRepository resource.
type VirtualGradleRepositoryArgs struct {
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
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
	// contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
	// 'libs-release-local').
	Key pulumi.StringInput
	// The keypair used to sign artifacts.
	KeyPair pulumi.StringPtrInput
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes pulumi.StringPtrInput
	// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
	// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
	// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
	PomRepositoryReferencesCleanupPolicy pulumi.StringPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
	// repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
	// repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds pulumi.IntPtrInput
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
//          VirtualGradleRepositoryArray{ VirtualGradleRepositoryArgs{...} }
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
//          VirtualGradleRepositoryMap{ "key": VirtualGradleRepositoryArgs{...} }
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
