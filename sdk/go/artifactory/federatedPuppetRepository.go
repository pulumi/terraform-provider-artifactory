// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Federated Puppet Repository Resource
//
// Creates a federated Puppet repository
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-artifactory/sdk/go/artifactory"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := artifactory.NewFederatedPuppetRepository(ctx, "terraform-federated-test-puppet-repo", &artifactory.FederatedPuppetRepositoryArgs{
// 			Key: pulumi.String("terraform-federated-test-puppet-repo"),
// 			Members: FederatedPuppetRepositoryMemberArray{
// 				&FederatedPuppetRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-puppet-repo"),
// 				},
// 				&FederatedPuppetRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-puppet-repo-2"),
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
type FederatedPuppetRepository struct {
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
	Members     FederatedPuppetRepositoryMemberArrayOutput `pulumi:"members"`
	Notes       pulumi.StringPtrOutput                     `pulumi:"notes"`
	PackageType pulumi.StringOutput                        `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrOutput `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    pulumi.StringPtrOutput   `pulumi:"projectKey"`
	PropertySets  pulumi.StringArrayOutput `pulumi:"propertySets"`
	RepoLayoutRef pulumi.StringOutput      `pulumi:"repoLayoutRef"`
	XrayIndex     pulumi.BoolOutput        `pulumi:"xrayIndex"`
}

// NewFederatedPuppetRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedPuppetRepository(ctx *pulumi.Context,
	name string, args *FederatedPuppetRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedPuppetRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FederatedPuppetRepository
	err := ctx.RegisterResource("artifactory:index/federatedPuppetRepository:FederatedPuppetRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedPuppetRepository gets an existing FederatedPuppetRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedPuppetRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedPuppetRepositoryState, opts ...pulumi.ResourceOption) (*FederatedPuppetRepository, error) {
	var resource FederatedPuppetRepository
	err := ctx.ReadResource("artifactory:index/federatedPuppetRepository:FederatedPuppetRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedPuppetRepository resources.
type federatedPuppetRepositoryState struct {
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
	Members     []FederatedPuppetRepositoryMember `pulumi:"members"`
	Notes       *string                           `pulumi:"notes"`
	PackageType *string                           `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    *string  `pulumi:"projectKey"`
	PropertySets  []string `pulumi:"propertySets"`
	RepoLayoutRef *string  `pulumi:"repoLayoutRef"`
	XrayIndex     *bool    `pulumi:"xrayIndex"`
}

type FederatedPuppetRepositoryState struct {
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
	Members     FederatedPuppetRepositoryMemberArrayInput
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    pulumi.StringPtrInput
	PropertySets  pulumi.StringArrayInput
	RepoLayoutRef pulumi.StringPtrInput
	XrayIndex     pulumi.BoolPtrInput
}

func (FederatedPuppetRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedPuppetRepositoryState)(nil)).Elem()
}

type federatedPuppetRepositoryArgs struct {
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
	Members []FederatedPuppetRepositoryMember `pulumi:"members"`
	Notes   *string                           `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    *string  `pulumi:"projectKey"`
	PropertySets  []string `pulumi:"propertySets"`
	RepoLayoutRef *string  `pulumi:"repoLayoutRef"`
	XrayIndex     *bool    `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a FederatedPuppetRepository resource.
type FederatedPuppetRepositoryArgs struct {
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
	Members FederatedPuppetRepositoryMemberArrayInput
	Notes   pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey    pulumi.StringPtrInput
	PropertySets  pulumi.StringArrayInput
	RepoLayoutRef pulumi.StringPtrInput
	XrayIndex     pulumi.BoolPtrInput
}

func (FederatedPuppetRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedPuppetRepositoryArgs)(nil)).Elem()
}

type FederatedPuppetRepositoryInput interface {
	pulumi.Input

	ToFederatedPuppetRepositoryOutput() FederatedPuppetRepositoryOutput
	ToFederatedPuppetRepositoryOutputWithContext(ctx context.Context) FederatedPuppetRepositoryOutput
}

func (*FederatedPuppetRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedPuppetRepository)(nil)).Elem()
}

func (i *FederatedPuppetRepository) ToFederatedPuppetRepositoryOutput() FederatedPuppetRepositoryOutput {
	return i.ToFederatedPuppetRepositoryOutputWithContext(context.Background())
}

func (i *FederatedPuppetRepository) ToFederatedPuppetRepositoryOutputWithContext(ctx context.Context) FederatedPuppetRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedPuppetRepositoryOutput)
}

// FederatedPuppetRepositoryArrayInput is an input type that accepts FederatedPuppetRepositoryArray and FederatedPuppetRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedPuppetRepositoryArrayInput` via:
//
//          FederatedPuppetRepositoryArray{ FederatedPuppetRepositoryArgs{...} }
type FederatedPuppetRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedPuppetRepositoryArrayOutput() FederatedPuppetRepositoryArrayOutput
	ToFederatedPuppetRepositoryArrayOutputWithContext(context.Context) FederatedPuppetRepositoryArrayOutput
}

type FederatedPuppetRepositoryArray []FederatedPuppetRepositoryInput

func (FederatedPuppetRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedPuppetRepository)(nil)).Elem()
}

func (i FederatedPuppetRepositoryArray) ToFederatedPuppetRepositoryArrayOutput() FederatedPuppetRepositoryArrayOutput {
	return i.ToFederatedPuppetRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedPuppetRepositoryArray) ToFederatedPuppetRepositoryArrayOutputWithContext(ctx context.Context) FederatedPuppetRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedPuppetRepositoryArrayOutput)
}

// FederatedPuppetRepositoryMapInput is an input type that accepts FederatedPuppetRepositoryMap and FederatedPuppetRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedPuppetRepositoryMapInput` via:
//
//          FederatedPuppetRepositoryMap{ "key": FederatedPuppetRepositoryArgs{...} }
type FederatedPuppetRepositoryMapInput interface {
	pulumi.Input

	ToFederatedPuppetRepositoryMapOutput() FederatedPuppetRepositoryMapOutput
	ToFederatedPuppetRepositoryMapOutputWithContext(context.Context) FederatedPuppetRepositoryMapOutput
}

type FederatedPuppetRepositoryMap map[string]FederatedPuppetRepositoryInput

func (FederatedPuppetRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedPuppetRepository)(nil)).Elem()
}

func (i FederatedPuppetRepositoryMap) ToFederatedPuppetRepositoryMapOutput() FederatedPuppetRepositoryMapOutput {
	return i.ToFederatedPuppetRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedPuppetRepositoryMap) ToFederatedPuppetRepositoryMapOutputWithContext(ctx context.Context) FederatedPuppetRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedPuppetRepositoryMapOutput)
}

type FederatedPuppetRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedPuppetRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedPuppetRepository)(nil)).Elem()
}

func (o FederatedPuppetRepositoryOutput) ToFederatedPuppetRepositoryOutput() FederatedPuppetRepositoryOutput {
	return o
}

func (o FederatedPuppetRepositoryOutput) ToFederatedPuppetRepositoryOutputWithContext(ctx context.Context) FederatedPuppetRepositoryOutput {
	return o
}

type FederatedPuppetRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedPuppetRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedPuppetRepository)(nil)).Elem()
}

func (o FederatedPuppetRepositoryArrayOutput) ToFederatedPuppetRepositoryArrayOutput() FederatedPuppetRepositoryArrayOutput {
	return o
}

func (o FederatedPuppetRepositoryArrayOutput) ToFederatedPuppetRepositoryArrayOutputWithContext(ctx context.Context) FederatedPuppetRepositoryArrayOutput {
	return o
}

func (o FederatedPuppetRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedPuppetRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedPuppetRepository {
		return vs[0].([]*FederatedPuppetRepository)[vs[1].(int)]
	}).(FederatedPuppetRepositoryOutput)
}

type FederatedPuppetRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedPuppetRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedPuppetRepository)(nil)).Elem()
}

func (o FederatedPuppetRepositoryMapOutput) ToFederatedPuppetRepositoryMapOutput() FederatedPuppetRepositoryMapOutput {
	return o
}

func (o FederatedPuppetRepositoryMapOutput) ToFederatedPuppetRepositoryMapOutputWithContext(ctx context.Context) FederatedPuppetRepositoryMapOutput {
	return o
}

func (o FederatedPuppetRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedPuppetRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedPuppetRepository {
		return vs[0].(map[string]*FederatedPuppetRepository)[vs[1].(string)]
	}).(FederatedPuppetRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedPuppetRepositoryInput)(nil)).Elem(), &FederatedPuppetRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedPuppetRepositoryArrayInput)(nil)).Elem(), FederatedPuppetRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedPuppetRepositoryMapInput)(nil)).Elem(), FederatedPuppetRepositoryMap{})
	pulumi.RegisterOutputType(FederatedPuppetRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedPuppetRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedPuppetRepositoryMapOutput{})
}
