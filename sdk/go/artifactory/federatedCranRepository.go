// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Federated Cran Repository Resource
//
// Creates a federated Cran repository
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
// 		_, err := artifactory.NewFederatedCranRepository(ctx, "terraform-federated-test-cran-repo", &artifactory.FederatedCranRepositoryArgs{
// 			Key: pulumi.String("terraform-federated-test-cran-repo"),
// 			Members: FederatedCranRepositoryMemberArray{
// 				&FederatedCranRepositoryMemberArgs{
// 					Enabled: pulumi.Bool(true),
// 					Url:     pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-cran-repo"),
// 				},
// 				&FederatedCranRepositoryMemberArgs{
// 					Enabled: pulumi.Bool(true),
// 					Url:     pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-cran-repo-2"),
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
type FederatedCranRepository struct {
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
	Members     FederatedCranRepositoryMemberArrayOutput `pulumi:"members"`
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

// NewFederatedCranRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedCranRepository(ctx *pulumi.Context,
	name string, args *FederatedCranRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedCranRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FederatedCranRepository
	err := ctx.RegisterResource("artifactory:index/federatedCranRepository:FederatedCranRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedCranRepository gets an existing FederatedCranRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedCranRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedCranRepositoryState, opts ...pulumi.ResourceOption) (*FederatedCranRepository, error) {
	var resource FederatedCranRepository
	err := ctx.ReadResource("artifactory:index/federatedCranRepository:FederatedCranRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedCranRepository resources.
type federatedCranRepositoryState struct {
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
	Members     []FederatedCranRepositoryMember `pulumi:"members"`
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

type FederatedCranRepositoryState struct {
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
	Members     FederatedCranRepositoryMemberArrayInput
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

func (FederatedCranRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedCranRepositoryState)(nil)).Elem()
}

type federatedCranRepositoryArgs struct {
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
	Members []FederatedCranRepositoryMember `pulumi:"members"`
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

// The set of arguments for constructing a FederatedCranRepository resource.
type FederatedCranRepositoryArgs struct {
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
	Members FederatedCranRepositoryMemberArrayInput
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

func (FederatedCranRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedCranRepositoryArgs)(nil)).Elem()
}

type FederatedCranRepositoryInput interface {
	pulumi.Input

	ToFederatedCranRepositoryOutput() FederatedCranRepositoryOutput
	ToFederatedCranRepositoryOutputWithContext(ctx context.Context) FederatedCranRepositoryOutput
}

func (*FederatedCranRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedCranRepository)(nil)).Elem()
}

func (i *FederatedCranRepository) ToFederatedCranRepositoryOutput() FederatedCranRepositoryOutput {
	return i.ToFederatedCranRepositoryOutputWithContext(context.Background())
}

func (i *FederatedCranRepository) ToFederatedCranRepositoryOutputWithContext(ctx context.Context) FederatedCranRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedCranRepositoryOutput)
}

// FederatedCranRepositoryArrayInput is an input type that accepts FederatedCranRepositoryArray and FederatedCranRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedCranRepositoryArrayInput` via:
//
//          FederatedCranRepositoryArray{ FederatedCranRepositoryArgs{...} }
type FederatedCranRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedCranRepositoryArrayOutput() FederatedCranRepositoryArrayOutput
	ToFederatedCranRepositoryArrayOutputWithContext(context.Context) FederatedCranRepositoryArrayOutput
}

type FederatedCranRepositoryArray []FederatedCranRepositoryInput

func (FederatedCranRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedCranRepository)(nil)).Elem()
}

func (i FederatedCranRepositoryArray) ToFederatedCranRepositoryArrayOutput() FederatedCranRepositoryArrayOutput {
	return i.ToFederatedCranRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedCranRepositoryArray) ToFederatedCranRepositoryArrayOutputWithContext(ctx context.Context) FederatedCranRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedCranRepositoryArrayOutput)
}

// FederatedCranRepositoryMapInput is an input type that accepts FederatedCranRepositoryMap and FederatedCranRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedCranRepositoryMapInput` via:
//
//          FederatedCranRepositoryMap{ "key": FederatedCranRepositoryArgs{...} }
type FederatedCranRepositoryMapInput interface {
	pulumi.Input

	ToFederatedCranRepositoryMapOutput() FederatedCranRepositoryMapOutput
	ToFederatedCranRepositoryMapOutputWithContext(context.Context) FederatedCranRepositoryMapOutput
}

type FederatedCranRepositoryMap map[string]FederatedCranRepositoryInput

func (FederatedCranRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedCranRepository)(nil)).Elem()
}

func (i FederatedCranRepositoryMap) ToFederatedCranRepositoryMapOutput() FederatedCranRepositoryMapOutput {
	return i.ToFederatedCranRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedCranRepositoryMap) ToFederatedCranRepositoryMapOutputWithContext(ctx context.Context) FederatedCranRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedCranRepositoryMapOutput)
}

type FederatedCranRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedCranRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedCranRepository)(nil)).Elem()
}

func (o FederatedCranRepositoryOutput) ToFederatedCranRepositoryOutput() FederatedCranRepositoryOutput {
	return o
}

func (o FederatedCranRepositoryOutput) ToFederatedCranRepositoryOutputWithContext(ctx context.Context) FederatedCranRepositoryOutput {
	return o
}

type FederatedCranRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedCranRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedCranRepository)(nil)).Elem()
}

func (o FederatedCranRepositoryArrayOutput) ToFederatedCranRepositoryArrayOutput() FederatedCranRepositoryArrayOutput {
	return o
}

func (o FederatedCranRepositoryArrayOutput) ToFederatedCranRepositoryArrayOutputWithContext(ctx context.Context) FederatedCranRepositoryArrayOutput {
	return o
}

func (o FederatedCranRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedCranRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedCranRepository {
		return vs[0].([]*FederatedCranRepository)[vs[1].(int)]
	}).(FederatedCranRepositoryOutput)
}

type FederatedCranRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedCranRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedCranRepository)(nil)).Elem()
}

func (o FederatedCranRepositoryMapOutput) ToFederatedCranRepositoryMapOutput() FederatedCranRepositoryMapOutput {
	return o
}

func (o FederatedCranRepositoryMapOutput) ToFederatedCranRepositoryMapOutputWithContext(ctx context.Context) FederatedCranRepositoryMapOutput {
	return o
}

func (o FederatedCranRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedCranRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedCranRepository {
		return vs[0].(map[string]*FederatedCranRepository)[vs[1].(string)]
	}).(FederatedCranRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedCranRepositoryInput)(nil)).Elem(), &FederatedCranRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedCranRepositoryArrayInput)(nil)).Elem(), FederatedCranRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedCranRepositoryMapInput)(nil)).Elem(), FederatedCranRepositoryMap{})
	pulumi.RegisterOutputType(FederatedCranRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedCranRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedCranRepositoryMapOutput{})
}
