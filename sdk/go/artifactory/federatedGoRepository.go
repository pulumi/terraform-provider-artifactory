// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Federated Go Repository Resource
//
// Creates a federated Go repository
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
// 		_, err := artifactory.NewFederatedGoRepository(ctx, "terraform-federated-test-go-repo", &artifactory.FederatedGoRepositoryArgs{
// 			Key: pulumi.String("terraform-federated-test-go-repo"),
// 			Members: FederatedGoRepositoryMemberArray{
// 				&FederatedGoRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-go-repo"),
// 				},
// 				&FederatedGoRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-go-repo-2"),
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
type FederatedGoRepository struct {
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
	Members     FederatedGoRepositoryMemberArrayOutput `pulumi:"members"`
	Notes       pulumi.StringPtrOutput                 `pulumi:"notes"`
	PackageType pulumi.StringOutput                    `pulumi:"packageType"`
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

// NewFederatedGoRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedGoRepository(ctx *pulumi.Context,
	name string, args *FederatedGoRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedGoRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FederatedGoRepository
	err := ctx.RegisterResource("artifactory:index/federatedGoRepository:FederatedGoRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedGoRepository gets an existing FederatedGoRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedGoRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedGoRepositoryState, opts ...pulumi.ResourceOption) (*FederatedGoRepository, error) {
	var resource FederatedGoRepository
	err := ctx.ReadResource("artifactory:index/federatedGoRepository:FederatedGoRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedGoRepository resources.
type federatedGoRepositoryState struct {
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
	Members     []FederatedGoRepositoryMember `pulumi:"members"`
	Notes       *string                       `pulumi:"notes"`
	PackageType *string                       `pulumi:"packageType"`
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

type FederatedGoRepositoryState struct {
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
	Members     FederatedGoRepositoryMemberArrayInput
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

func (FederatedGoRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedGoRepositoryState)(nil)).Elem()
}

type federatedGoRepositoryArgs struct {
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
	Members []FederatedGoRepositoryMember `pulumi:"members"`
	Notes   *string                       `pulumi:"notes"`
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

// The set of arguments for constructing a FederatedGoRepository resource.
type FederatedGoRepositoryArgs struct {
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
	Members FederatedGoRepositoryMemberArrayInput
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

func (FederatedGoRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedGoRepositoryArgs)(nil)).Elem()
}

type FederatedGoRepositoryInput interface {
	pulumi.Input

	ToFederatedGoRepositoryOutput() FederatedGoRepositoryOutput
	ToFederatedGoRepositoryOutputWithContext(ctx context.Context) FederatedGoRepositoryOutput
}

func (*FederatedGoRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedGoRepository)(nil)).Elem()
}

func (i *FederatedGoRepository) ToFederatedGoRepositoryOutput() FederatedGoRepositoryOutput {
	return i.ToFederatedGoRepositoryOutputWithContext(context.Background())
}

func (i *FederatedGoRepository) ToFederatedGoRepositoryOutputWithContext(ctx context.Context) FederatedGoRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedGoRepositoryOutput)
}

// FederatedGoRepositoryArrayInput is an input type that accepts FederatedGoRepositoryArray and FederatedGoRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedGoRepositoryArrayInput` via:
//
//          FederatedGoRepositoryArray{ FederatedGoRepositoryArgs{...} }
type FederatedGoRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedGoRepositoryArrayOutput() FederatedGoRepositoryArrayOutput
	ToFederatedGoRepositoryArrayOutputWithContext(context.Context) FederatedGoRepositoryArrayOutput
}

type FederatedGoRepositoryArray []FederatedGoRepositoryInput

func (FederatedGoRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedGoRepository)(nil)).Elem()
}

func (i FederatedGoRepositoryArray) ToFederatedGoRepositoryArrayOutput() FederatedGoRepositoryArrayOutput {
	return i.ToFederatedGoRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedGoRepositoryArray) ToFederatedGoRepositoryArrayOutputWithContext(ctx context.Context) FederatedGoRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedGoRepositoryArrayOutput)
}

// FederatedGoRepositoryMapInput is an input type that accepts FederatedGoRepositoryMap and FederatedGoRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedGoRepositoryMapInput` via:
//
//          FederatedGoRepositoryMap{ "key": FederatedGoRepositoryArgs{...} }
type FederatedGoRepositoryMapInput interface {
	pulumi.Input

	ToFederatedGoRepositoryMapOutput() FederatedGoRepositoryMapOutput
	ToFederatedGoRepositoryMapOutputWithContext(context.Context) FederatedGoRepositoryMapOutput
}

type FederatedGoRepositoryMap map[string]FederatedGoRepositoryInput

func (FederatedGoRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedGoRepository)(nil)).Elem()
}

func (i FederatedGoRepositoryMap) ToFederatedGoRepositoryMapOutput() FederatedGoRepositoryMapOutput {
	return i.ToFederatedGoRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedGoRepositoryMap) ToFederatedGoRepositoryMapOutputWithContext(ctx context.Context) FederatedGoRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedGoRepositoryMapOutput)
}

type FederatedGoRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedGoRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedGoRepository)(nil)).Elem()
}

func (o FederatedGoRepositoryOutput) ToFederatedGoRepositoryOutput() FederatedGoRepositoryOutput {
	return o
}

func (o FederatedGoRepositoryOutput) ToFederatedGoRepositoryOutputWithContext(ctx context.Context) FederatedGoRepositoryOutput {
	return o
}

type FederatedGoRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedGoRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedGoRepository)(nil)).Elem()
}

func (o FederatedGoRepositoryArrayOutput) ToFederatedGoRepositoryArrayOutput() FederatedGoRepositoryArrayOutput {
	return o
}

func (o FederatedGoRepositoryArrayOutput) ToFederatedGoRepositoryArrayOutputWithContext(ctx context.Context) FederatedGoRepositoryArrayOutput {
	return o
}

func (o FederatedGoRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedGoRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedGoRepository {
		return vs[0].([]*FederatedGoRepository)[vs[1].(int)]
	}).(FederatedGoRepositoryOutput)
}

type FederatedGoRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedGoRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedGoRepository)(nil)).Elem()
}

func (o FederatedGoRepositoryMapOutput) ToFederatedGoRepositoryMapOutput() FederatedGoRepositoryMapOutput {
	return o
}

func (o FederatedGoRepositoryMapOutput) ToFederatedGoRepositoryMapOutputWithContext(ctx context.Context) FederatedGoRepositoryMapOutput {
	return o
}

func (o FederatedGoRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedGoRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedGoRepository {
		return vs[0].(map[string]*FederatedGoRepository)[vs[1].(string)]
	}).(FederatedGoRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedGoRepositoryInput)(nil)).Elem(), &FederatedGoRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedGoRepositoryArrayInput)(nil)).Elem(), FederatedGoRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedGoRepositoryMapInput)(nil)).Elem(), FederatedGoRepositoryMap{})
	pulumi.RegisterOutputType(FederatedGoRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedGoRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedGoRepositoryMapOutput{})
}
