// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Federated Helm Repository Resource
//
// Creates a federated Helm repository
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-artifactory/sdk/go/artifactory"
// 	"github.com/pulumi/pulumi-artifactory/sdk/v2/go/artifactory"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := artifactory.NewFederatedHelmRepository(ctx, "terraform-federated-test-helm-repo", &artifactory.FederatedHelmRepositoryArgs{
// 			Key: pulumi.String("terraform-federated-test-helm-repo"),
// 			Members: FederatedHelmRepositoryMemberArray{
// 				&FederatedHelmRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-helm-repo"),
// 				},
// 				&FederatedHelmRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-helm-repo-2"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type FederatedHelmRepository struct {
	pulumi.CustomResourceState

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrOutput   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrOutput   `pulumi:"blackedOut"`
	Description            pulumi.StringPtrOutput `pulumi:"description"`
	DownloadDirect         pulumi.BoolPtrOutput   `pulumi:"downloadDirect"`
	ExcludesPattern        pulumi.StringOutput    `pulumi:"excludesPattern"`
	IncludesPattern        pulumi.StringOutput    `pulumi:"includesPattern"`
	// - the identity key of the repo
	Key pulumi.StringOutput `pulumi:"key"`
	// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
	Members     FederatedHelmRepositoryMemberArrayOutput `pulumi:"members"`
	Notes       pulumi.StringPtrOutput                   `pulumi:"notes"`
	PackageType pulumi.StringOutput                      `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrOutput `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey   pulumi.StringPtrOutput   `pulumi:"projectKey"`
	PropertySets pulumi.StringArrayOutput `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	XrayIndex     pulumi.BoolOutput      `pulumi:"xrayIndex"`
}

// NewFederatedHelmRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedHelmRepository(ctx *pulumi.Context,
	name string, args *FederatedHelmRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedHelmRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FederatedHelmRepository
	err := ctx.RegisterResource("artifactory:index/federatedHelmRepository:FederatedHelmRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedHelmRepository gets an existing FederatedHelmRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedHelmRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedHelmRepositoryState, opts ...pulumi.ResourceOption) (*FederatedHelmRepository, error) {
	var resource FederatedHelmRepository
	err := ctx.ReadResource("artifactory:index/federatedHelmRepository:FederatedHelmRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedHelmRepository resources.
type federatedHelmRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	IncludesPattern        *string `pulumi:"includesPattern"`
	// - the identity key of the repo
	Key *string `pulumi:"key"`
	// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
	Members     []FederatedHelmRepositoryMember `pulumi:"members"`
	Notes       *string                         `pulumi:"notes"`
	PackageType *string                         `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey   *string  `pulumi:"projectKey"`
	PropertySets []string `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	XrayIndex     *bool   `pulumi:"xrayIndex"`
}

type FederatedHelmRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	BlackedOut             pulumi.BoolPtrInput
	Description            pulumi.StringPtrInput
	DownloadDirect         pulumi.BoolPtrInput
	ExcludesPattern        pulumi.StringPtrInput
	IncludesPattern        pulumi.StringPtrInput
	// - the identity key of the repo
	Key pulumi.StringPtrInput
	// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
	Members     FederatedHelmRepositoryMemberArrayInput
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey   pulumi.StringPtrInput
	PropertySets pulumi.StringArrayInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	XrayIndex     pulumi.BoolPtrInput
}

func (FederatedHelmRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedHelmRepositoryState)(nil)).Elem()
}

type federatedHelmRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	IncludesPattern        *string `pulumi:"includesPattern"`
	// - the identity key of the repo
	Key string `pulumi:"key"`
	// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
	Members []FederatedHelmRepositoryMember `pulumi:"members"`
	Notes   *string                         `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey   *string  `pulumi:"projectKey"`
	PropertySets []string `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	XrayIndex     *bool   `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a FederatedHelmRepository resource.
type FederatedHelmRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	BlackedOut             pulumi.BoolPtrInput
	Description            pulumi.StringPtrInput
	DownloadDirect         pulumi.BoolPtrInput
	ExcludesPattern        pulumi.StringPtrInput
	IncludesPattern        pulumi.StringPtrInput
	// - the identity key of the repo
	Key pulumi.StringInput
	// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
	Members FederatedHelmRepositoryMemberArrayInput
	Notes   pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey   pulumi.StringPtrInput
	PropertySets pulumi.StringArrayInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	XrayIndex     pulumi.BoolPtrInput
}

func (FederatedHelmRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedHelmRepositoryArgs)(nil)).Elem()
}

type FederatedHelmRepositoryInput interface {
	pulumi.Input

	ToFederatedHelmRepositoryOutput() FederatedHelmRepositoryOutput
	ToFederatedHelmRepositoryOutputWithContext(ctx context.Context) FederatedHelmRepositoryOutput
}

func (*FederatedHelmRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedHelmRepository)(nil)).Elem()
}

func (i *FederatedHelmRepository) ToFederatedHelmRepositoryOutput() FederatedHelmRepositoryOutput {
	return i.ToFederatedHelmRepositoryOutputWithContext(context.Background())
}

func (i *FederatedHelmRepository) ToFederatedHelmRepositoryOutputWithContext(ctx context.Context) FederatedHelmRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedHelmRepositoryOutput)
}

// FederatedHelmRepositoryArrayInput is an input type that accepts FederatedHelmRepositoryArray and FederatedHelmRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedHelmRepositoryArrayInput` via:
//
//          FederatedHelmRepositoryArray{ FederatedHelmRepositoryArgs{...} }
type FederatedHelmRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedHelmRepositoryArrayOutput() FederatedHelmRepositoryArrayOutput
	ToFederatedHelmRepositoryArrayOutputWithContext(context.Context) FederatedHelmRepositoryArrayOutput
}

type FederatedHelmRepositoryArray []FederatedHelmRepositoryInput

func (FederatedHelmRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedHelmRepository)(nil)).Elem()
}

func (i FederatedHelmRepositoryArray) ToFederatedHelmRepositoryArrayOutput() FederatedHelmRepositoryArrayOutput {
	return i.ToFederatedHelmRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedHelmRepositoryArray) ToFederatedHelmRepositoryArrayOutputWithContext(ctx context.Context) FederatedHelmRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedHelmRepositoryArrayOutput)
}

// FederatedHelmRepositoryMapInput is an input type that accepts FederatedHelmRepositoryMap and FederatedHelmRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedHelmRepositoryMapInput` via:
//
//          FederatedHelmRepositoryMap{ "key": FederatedHelmRepositoryArgs{...} }
type FederatedHelmRepositoryMapInput interface {
	pulumi.Input

	ToFederatedHelmRepositoryMapOutput() FederatedHelmRepositoryMapOutput
	ToFederatedHelmRepositoryMapOutputWithContext(context.Context) FederatedHelmRepositoryMapOutput
}

type FederatedHelmRepositoryMap map[string]FederatedHelmRepositoryInput

func (FederatedHelmRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedHelmRepository)(nil)).Elem()
}

func (i FederatedHelmRepositoryMap) ToFederatedHelmRepositoryMapOutput() FederatedHelmRepositoryMapOutput {
	return i.ToFederatedHelmRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedHelmRepositoryMap) ToFederatedHelmRepositoryMapOutputWithContext(ctx context.Context) FederatedHelmRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedHelmRepositoryMapOutput)
}

type FederatedHelmRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedHelmRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedHelmRepository)(nil)).Elem()
}

func (o FederatedHelmRepositoryOutput) ToFederatedHelmRepositoryOutput() FederatedHelmRepositoryOutput {
	return o
}

func (o FederatedHelmRepositoryOutput) ToFederatedHelmRepositoryOutputWithContext(ctx context.Context) FederatedHelmRepositoryOutput {
	return o
}

type FederatedHelmRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedHelmRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedHelmRepository)(nil)).Elem()
}

func (o FederatedHelmRepositoryArrayOutput) ToFederatedHelmRepositoryArrayOutput() FederatedHelmRepositoryArrayOutput {
	return o
}

func (o FederatedHelmRepositoryArrayOutput) ToFederatedHelmRepositoryArrayOutputWithContext(ctx context.Context) FederatedHelmRepositoryArrayOutput {
	return o
}

func (o FederatedHelmRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedHelmRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedHelmRepository {
		return vs[0].([]*FederatedHelmRepository)[vs[1].(int)]
	}).(FederatedHelmRepositoryOutput)
}

type FederatedHelmRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedHelmRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedHelmRepository)(nil)).Elem()
}

func (o FederatedHelmRepositoryMapOutput) ToFederatedHelmRepositoryMapOutput() FederatedHelmRepositoryMapOutput {
	return o
}

func (o FederatedHelmRepositoryMapOutput) ToFederatedHelmRepositoryMapOutputWithContext(ctx context.Context) FederatedHelmRepositoryMapOutput {
	return o
}

func (o FederatedHelmRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedHelmRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedHelmRepository {
		return vs[0].(map[string]*FederatedHelmRepository)[vs[1].(string)]
	}).(FederatedHelmRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedHelmRepositoryInput)(nil)).Elem(), &FederatedHelmRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedHelmRepositoryArrayInput)(nil)).Elem(), FederatedHelmRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedHelmRepositoryMapInput)(nil)).Elem(), FederatedHelmRepositoryMap{})
	pulumi.RegisterOutputType(FederatedHelmRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedHelmRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedHelmRepositoryMapOutput{})
}
