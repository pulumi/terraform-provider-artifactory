// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Federated Debian Repository Resource
//
// Creates a federated Debian repository
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
// 		_, err := artifactory.NewFederatedDebianRepository(ctx, "terraform-federated-test-debian-repo", &artifactory.FederatedDebianRepositoryArgs{
// 			Key: pulumi.String("terraform-federated-test-debian-repo"),
// 			Members: FederatedDebianRepositoryMemberArray{
// 				&FederatedDebianRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-debian-repo"),
// 				},
// 				&FederatedDebianRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-debian-repo-2"),
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
type FederatedDebianRepository struct {
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
	Members     FederatedDebianRepositoryMemberArrayOutput `pulumi:"members"`
	Notes       pulumi.StringPtrOutput                     `pulumi:"notes"`
	PackageType pulumi.StringOutput                        `pulumi:"packageType"`
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

// NewFederatedDebianRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedDebianRepository(ctx *pulumi.Context,
	name string, args *FederatedDebianRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedDebianRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FederatedDebianRepository
	err := ctx.RegisterResource("artifactory:index/federatedDebianRepository:FederatedDebianRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedDebianRepository gets an existing FederatedDebianRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedDebianRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedDebianRepositoryState, opts ...pulumi.ResourceOption) (*FederatedDebianRepository, error) {
	var resource FederatedDebianRepository
	err := ctx.ReadResource("artifactory:index/federatedDebianRepository:FederatedDebianRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedDebianRepository resources.
type federatedDebianRepositoryState struct {
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
	Members     []FederatedDebianRepositoryMember `pulumi:"members"`
	Notes       *string                           `pulumi:"notes"`
	PackageType *string                           `pulumi:"packageType"`
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

type FederatedDebianRepositoryState struct {
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
	Members     FederatedDebianRepositoryMemberArrayInput
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

func (FederatedDebianRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedDebianRepositoryState)(nil)).Elem()
}

type federatedDebianRepositoryArgs struct {
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
	Members []FederatedDebianRepositoryMember `pulumi:"members"`
	Notes   *string                           `pulumi:"notes"`
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

// The set of arguments for constructing a FederatedDebianRepository resource.
type FederatedDebianRepositoryArgs struct {
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
	Members FederatedDebianRepositoryMemberArrayInput
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

func (FederatedDebianRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedDebianRepositoryArgs)(nil)).Elem()
}

type FederatedDebianRepositoryInput interface {
	pulumi.Input

	ToFederatedDebianRepositoryOutput() FederatedDebianRepositoryOutput
	ToFederatedDebianRepositoryOutputWithContext(ctx context.Context) FederatedDebianRepositoryOutput
}

func (*FederatedDebianRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedDebianRepository)(nil)).Elem()
}

func (i *FederatedDebianRepository) ToFederatedDebianRepositoryOutput() FederatedDebianRepositoryOutput {
	return i.ToFederatedDebianRepositoryOutputWithContext(context.Background())
}

func (i *FederatedDebianRepository) ToFederatedDebianRepositoryOutputWithContext(ctx context.Context) FederatedDebianRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedDebianRepositoryOutput)
}

// FederatedDebianRepositoryArrayInput is an input type that accepts FederatedDebianRepositoryArray and FederatedDebianRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedDebianRepositoryArrayInput` via:
//
//          FederatedDebianRepositoryArray{ FederatedDebianRepositoryArgs{...} }
type FederatedDebianRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedDebianRepositoryArrayOutput() FederatedDebianRepositoryArrayOutput
	ToFederatedDebianRepositoryArrayOutputWithContext(context.Context) FederatedDebianRepositoryArrayOutput
}

type FederatedDebianRepositoryArray []FederatedDebianRepositoryInput

func (FederatedDebianRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedDebianRepository)(nil)).Elem()
}

func (i FederatedDebianRepositoryArray) ToFederatedDebianRepositoryArrayOutput() FederatedDebianRepositoryArrayOutput {
	return i.ToFederatedDebianRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedDebianRepositoryArray) ToFederatedDebianRepositoryArrayOutputWithContext(ctx context.Context) FederatedDebianRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedDebianRepositoryArrayOutput)
}

// FederatedDebianRepositoryMapInput is an input type that accepts FederatedDebianRepositoryMap and FederatedDebianRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedDebianRepositoryMapInput` via:
//
//          FederatedDebianRepositoryMap{ "key": FederatedDebianRepositoryArgs{...} }
type FederatedDebianRepositoryMapInput interface {
	pulumi.Input

	ToFederatedDebianRepositoryMapOutput() FederatedDebianRepositoryMapOutput
	ToFederatedDebianRepositoryMapOutputWithContext(context.Context) FederatedDebianRepositoryMapOutput
}

type FederatedDebianRepositoryMap map[string]FederatedDebianRepositoryInput

func (FederatedDebianRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedDebianRepository)(nil)).Elem()
}

func (i FederatedDebianRepositoryMap) ToFederatedDebianRepositoryMapOutput() FederatedDebianRepositoryMapOutput {
	return i.ToFederatedDebianRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedDebianRepositoryMap) ToFederatedDebianRepositoryMapOutputWithContext(ctx context.Context) FederatedDebianRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedDebianRepositoryMapOutput)
}

type FederatedDebianRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedDebianRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedDebianRepository)(nil)).Elem()
}

func (o FederatedDebianRepositoryOutput) ToFederatedDebianRepositoryOutput() FederatedDebianRepositoryOutput {
	return o
}

func (o FederatedDebianRepositoryOutput) ToFederatedDebianRepositoryOutputWithContext(ctx context.Context) FederatedDebianRepositoryOutput {
	return o
}

type FederatedDebianRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedDebianRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedDebianRepository)(nil)).Elem()
}

func (o FederatedDebianRepositoryArrayOutput) ToFederatedDebianRepositoryArrayOutput() FederatedDebianRepositoryArrayOutput {
	return o
}

func (o FederatedDebianRepositoryArrayOutput) ToFederatedDebianRepositoryArrayOutputWithContext(ctx context.Context) FederatedDebianRepositoryArrayOutput {
	return o
}

func (o FederatedDebianRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedDebianRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedDebianRepository {
		return vs[0].([]*FederatedDebianRepository)[vs[1].(int)]
	}).(FederatedDebianRepositoryOutput)
}

type FederatedDebianRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedDebianRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedDebianRepository)(nil)).Elem()
}

func (o FederatedDebianRepositoryMapOutput) ToFederatedDebianRepositoryMapOutput() FederatedDebianRepositoryMapOutput {
	return o
}

func (o FederatedDebianRepositoryMapOutput) ToFederatedDebianRepositoryMapOutputWithContext(ctx context.Context) FederatedDebianRepositoryMapOutput {
	return o
}

func (o FederatedDebianRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedDebianRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedDebianRepository {
		return vs[0].(map[string]*FederatedDebianRepository)[vs[1].(string)]
	}).(FederatedDebianRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedDebianRepositoryInput)(nil)).Elem(), &FederatedDebianRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedDebianRepositoryArrayInput)(nil)).Elem(), FederatedDebianRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedDebianRepositoryMapInput)(nil)).Elem(), FederatedDebianRepositoryMap{})
	pulumi.RegisterOutputType(FederatedDebianRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedDebianRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedDebianRepositoryMapOutput{})
}
