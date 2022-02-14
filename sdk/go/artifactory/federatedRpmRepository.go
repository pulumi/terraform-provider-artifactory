// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Federated Rpm Repository Resource
//
// Creates a federated Rpm repository
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
// 		_, err := artifactory.NewFederatedRpmRepository(ctx, "terraform-federated-test-rpm-repo", &artifactory.FederatedRpmRepositoryArgs{
// 			Key: pulumi.String("terraform-federated-test-rpm-repo"),
// 			Members: FederatedRpmRepositoryMemberArray{
// 				&FederatedRpmRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-rpm-repo"),
// 				},
// 				&FederatedRpmRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-rpm-repo-2"),
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
type FederatedRpmRepository struct {
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
	Members     FederatedRpmRepositoryMemberArrayOutput `pulumi:"members"`
	Notes       pulumi.StringPtrOutput                  `pulumi:"notes"`
	PackageType pulumi.StringOutput                     `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrOutput     `pulumi:"priorityResolution"`
	PropertySets       pulumi.StringArrayOutput `pulumi:"propertySets"`
	RepoLayoutRef      pulumi.StringOutput      `pulumi:"repoLayoutRef"`
	XrayIndex          pulumi.BoolOutput        `pulumi:"xrayIndex"`
}

// NewFederatedRpmRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedRpmRepository(ctx *pulumi.Context,
	name string, args *FederatedRpmRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedRpmRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FederatedRpmRepository
	err := ctx.RegisterResource("artifactory:index/federatedRpmRepository:FederatedRpmRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedRpmRepository gets an existing FederatedRpmRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedRpmRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedRpmRepositoryState, opts ...pulumi.ResourceOption) (*FederatedRpmRepository, error) {
	var resource FederatedRpmRepository
	err := ctx.ReadResource("artifactory:index/federatedRpmRepository:FederatedRpmRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedRpmRepository resources.
type federatedRpmRepositoryState struct {
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
	Members     []FederatedRpmRepositoryMember `pulumi:"members"`
	Notes       *string                        `pulumi:"notes"`
	PackageType *string                        `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool    `pulumi:"priorityResolution"`
	PropertySets       []string `pulumi:"propertySets"`
	RepoLayoutRef      *string  `pulumi:"repoLayoutRef"`
	XrayIndex          *bool    `pulumi:"xrayIndex"`
}

type FederatedRpmRepositoryState struct {
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
	Members     FederatedRpmRepositoryMemberArrayInput
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	PropertySets       pulumi.StringArrayInput
	RepoLayoutRef      pulumi.StringPtrInput
	XrayIndex          pulumi.BoolPtrInput
}

func (FederatedRpmRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedRpmRepositoryState)(nil)).Elem()
}

type federatedRpmRepositoryArgs struct {
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
	Members []FederatedRpmRepositoryMember `pulumi:"members"`
	Notes   *string                        `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool    `pulumi:"priorityResolution"`
	PropertySets       []string `pulumi:"propertySets"`
	RepoLayoutRef      *string  `pulumi:"repoLayoutRef"`
	XrayIndex          *bool    `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a FederatedRpmRepository resource.
type FederatedRpmRepositoryArgs struct {
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
	Members FederatedRpmRepositoryMemberArrayInput
	Notes   pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	PropertySets       pulumi.StringArrayInput
	RepoLayoutRef      pulumi.StringPtrInput
	XrayIndex          pulumi.BoolPtrInput
}

func (FederatedRpmRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedRpmRepositoryArgs)(nil)).Elem()
}

type FederatedRpmRepositoryInput interface {
	pulumi.Input

	ToFederatedRpmRepositoryOutput() FederatedRpmRepositoryOutput
	ToFederatedRpmRepositoryOutputWithContext(ctx context.Context) FederatedRpmRepositoryOutput
}

func (*FederatedRpmRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedRpmRepository)(nil)).Elem()
}

func (i *FederatedRpmRepository) ToFederatedRpmRepositoryOutput() FederatedRpmRepositoryOutput {
	return i.ToFederatedRpmRepositoryOutputWithContext(context.Background())
}

func (i *FederatedRpmRepository) ToFederatedRpmRepositoryOutputWithContext(ctx context.Context) FederatedRpmRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedRpmRepositoryOutput)
}

// FederatedRpmRepositoryArrayInput is an input type that accepts FederatedRpmRepositoryArray and FederatedRpmRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedRpmRepositoryArrayInput` via:
//
//          FederatedRpmRepositoryArray{ FederatedRpmRepositoryArgs{...} }
type FederatedRpmRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedRpmRepositoryArrayOutput() FederatedRpmRepositoryArrayOutput
	ToFederatedRpmRepositoryArrayOutputWithContext(context.Context) FederatedRpmRepositoryArrayOutput
}

type FederatedRpmRepositoryArray []FederatedRpmRepositoryInput

func (FederatedRpmRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedRpmRepository)(nil)).Elem()
}

func (i FederatedRpmRepositoryArray) ToFederatedRpmRepositoryArrayOutput() FederatedRpmRepositoryArrayOutput {
	return i.ToFederatedRpmRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedRpmRepositoryArray) ToFederatedRpmRepositoryArrayOutputWithContext(ctx context.Context) FederatedRpmRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedRpmRepositoryArrayOutput)
}

// FederatedRpmRepositoryMapInput is an input type that accepts FederatedRpmRepositoryMap and FederatedRpmRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedRpmRepositoryMapInput` via:
//
//          FederatedRpmRepositoryMap{ "key": FederatedRpmRepositoryArgs{...} }
type FederatedRpmRepositoryMapInput interface {
	pulumi.Input

	ToFederatedRpmRepositoryMapOutput() FederatedRpmRepositoryMapOutput
	ToFederatedRpmRepositoryMapOutputWithContext(context.Context) FederatedRpmRepositoryMapOutput
}

type FederatedRpmRepositoryMap map[string]FederatedRpmRepositoryInput

func (FederatedRpmRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedRpmRepository)(nil)).Elem()
}

func (i FederatedRpmRepositoryMap) ToFederatedRpmRepositoryMapOutput() FederatedRpmRepositoryMapOutput {
	return i.ToFederatedRpmRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedRpmRepositoryMap) ToFederatedRpmRepositoryMapOutputWithContext(ctx context.Context) FederatedRpmRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedRpmRepositoryMapOutput)
}

type FederatedRpmRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedRpmRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedRpmRepository)(nil)).Elem()
}

func (o FederatedRpmRepositoryOutput) ToFederatedRpmRepositoryOutput() FederatedRpmRepositoryOutput {
	return o
}

func (o FederatedRpmRepositoryOutput) ToFederatedRpmRepositoryOutputWithContext(ctx context.Context) FederatedRpmRepositoryOutput {
	return o
}

type FederatedRpmRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedRpmRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedRpmRepository)(nil)).Elem()
}

func (o FederatedRpmRepositoryArrayOutput) ToFederatedRpmRepositoryArrayOutput() FederatedRpmRepositoryArrayOutput {
	return o
}

func (o FederatedRpmRepositoryArrayOutput) ToFederatedRpmRepositoryArrayOutputWithContext(ctx context.Context) FederatedRpmRepositoryArrayOutput {
	return o
}

func (o FederatedRpmRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedRpmRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedRpmRepository {
		return vs[0].([]*FederatedRpmRepository)[vs[1].(int)]
	}).(FederatedRpmRepositoryOutput)
}

type FederatedRpmRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedRpmRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedRpmRepository)(nil)).Elem()
}

func (o FederatedRpmRepositoryMapOutput) ToFederatedRpmRepositoryMapOutput() FederatedRpmRepositoryMapOutput {
	return o
}

func (o FederatedRpmRepositoryMapOutput) ToFederatedRpmRepositoryMapOutputWithContext(ctx context.Context) FederatedRpmRepositoryMapOutput {
	return o
}

func (o FederatedRpmRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedRpmRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedRpmRepository {
		return vs[0].(map[string]*FederatedRpmRepository)[vs[1].(string)]
	}).(FederatedRpmRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedRpmRepositoryInput)(nil)).Elem(), &FederatedRpmRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedRpmRepositoryArrayInput)(nil)).Elem(), FederatedRpmRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedRpmRepositoryMapInput)(nil)).Elem(), FederatedRpmRepositoryMap{})
	pulumi.RegisterOutputType(FederatedRpmRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedRpmRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedRpmRepositoryMapOutput{})
}
