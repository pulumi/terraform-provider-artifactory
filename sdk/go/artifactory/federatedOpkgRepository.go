// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Federated Opkg Repository Resource
//
// Creates a federated Opkg repository
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
// 		_, err := artifactory.NewFederatedOpkgRepository(ctx, "terraform-federated-test-opkg-repo", &artifactory.FederatedOpkgRepositoryArgs{
// 			Key: pulumi.String("terraform-federated-test-opkg-repo"),
// 			Members: FederatedOpkgRepositoryMemberArray{
// 				&FederatedOpkgRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-opkg-repo"),
// 				},
// 				&FederatedOpkgRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-opkg-repo-2"),
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
type FederatedOpkgRepository struct {
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
	Members     FederatedOpkgRepositoryMemberArrayOutput `pulumi:"members"`
	Notes       pulumi.StringPtrOutput                   `pulumi:"notes"`
	PackageType pulumi.StringOutput                      `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrOutput     `pulumi:"priorityResolution"`
	PropertySets       pulumi.StringArrayOutput `pulumi:"propertySets"`
	RepoLayoutRef      pulumi.StringOutput      `pulumi:"repoLayoutRef"`
	XrayIndex          pulumi.BoolOutput        `pulumi:"xrayIndex"`
}

// NewFederatedOpkgRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedOpkgRepository(ctx *pulumi.Context,
	name string, args *FederatedOpkgRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedOpkgRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FederatedOpkgRepository
	err := ctx.RegisterResource("artifactory:index/federatedOpkgRepository:FederatedOpkgRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedOpkgRepository gets an existing FederatedOpkgRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedOpkgRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedOpkgRepositoryState, opts ...pulumi.ResourceOption) (*FederatedOpkgRepository, error) {
	var resource FederatedOpkgRepository
	err := ctx.ReadResource("artifactory:index/federatedOpkgRepository:FederatedOpkgRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedOpkgRepository resources.
type federatedOpkgRepositoryState struct {
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
	Members     []FederatedOpkgRepositoryMember `pulumi:"members"`
	Notes       *string                         `pulumi:"notes"`
	PackageType *string                         `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool    `pulumi:"priorityResolution"`
	PropertySets       []string `pulumi:"propertySets"`
	RepoLayoutRef      *string  `pulumi:"repoLayoutRef"`
	XrayIndex          *bool    `pulumi:"xrayIndex"`
}

type FederatedOpkgRepositoryState struct {
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
	Members     FederatedOpkgRepositoryMemberArrayInput
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	PropertySets       pulumi.StringArrayInput
	RepoLayoutRef      pulumi.StringPtrInput
	XrayIndex          pulumi.BoolPtrInput
}

func (FederatedOpkgRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedOpkgRepositoryState)(nil)).Elem()
}

type federatedOpkgRepositoryArgs struct {
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
	Members []FederatedOpkgRepositoryMember `pulumi:"members"`
	Notes   *string                         `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool    `pulumi:"priorityResolution"`
	PropertySets       []string `pulumi:"propertySets"`
	RepoLayoutRef      *string  `pulumi:"repoLayoutRef"`
	XrayIndex          *bool    `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a FederatedOpkgRepository resource.
type FederatedOpkgRepositoryArgs struct {
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
	Members FederatedOpkgRepositoryMemberArrayInput
	Notes   pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	PropertySets       pulumi.StringArrayInput
	RepoLayoutRef      pulumi.StringPtrInput
	XrayIndex          pulumi.BoolPtrInput
}

func (FederatedOpkgRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedOpkgRepositoryArgs)(nil)).Elem()
}

type FederatedOpkgRepositoryInput interface {
	pulumi.Input

	ToFederatedOpkgRepositoryOutput() FederatedOpkgRepositoryOutput
	ToFederatedOpkgRepositoryOutputWithContext(ctx context.Context) FederatedOpkgRepositoryOutput
}

func (*FederatedOpkgRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedOpkgRepository)(nil)).Elem()
}

func (i *FederatedOpkgRepository) ToFederatedOpkgRepositoryOutput() FederatedOpkgRepositoryOutput {
	return i.ToFederatedOpkgRepositoryOutputWithContext(context.Background())
}

func (i *FederatedOpkgRepository) ToFederatedOpkgRepositoryOutputWithContext(ctx context.Context) FederatedOpkgRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedOpkgRepositoryOutput)
}

// FederatedOpkgRepositoryArrayInput is an input type that accepts FederatedOpkgRepositoryArray and FederatedOpkgRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedOpkgRepositoryArrayInput` via:
//
//          FederatedOpkgRepositoryArray{ FederatedOpkgRepositoryArgs{...} }
type FederatedOpkgRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedOpkgRepositoryArrayOutput() FederatedOpkgRepositoryArrayOutput
	ToFederatedOpkgRepositoryArrayOutputWithContext(context.Context) FederatedOpkgRepositoryArrayOutput
}

type FederatedOpkgRepositoryArray []FederatedOpkgRepositoryInput

func (FederatedOpkgRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedOpkgRepository)(nil)).Elem()
}

func (i FederatedOpkgRepositoryArray) ToFederatedOpkgRepositoryArrayOutput() FederatedOpkgRepositoryArrayOutput {
	return i.ToFederatedOpkgRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedOpkgRepositoryArray) ToFederatedOpkgRepositoryArrayOutputWithContext(ctx context.Context) FederatedOpkgRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedOpkgRepositoryArrayOutput)
}

// FederatedOpkgRepositoryMapInput is an input type that accepts FederatedOpkgRepositoryMap and FederatedOpkgRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedOpkgRepositoryMapInput` via:
//
//          FederatedOpkgRepositoryMap{ "key": FederatedOpkgRepositoryArgs{...} }
type FederatedOpkgRepositoryMapInput interface {
	pulumi.Input

	ToFederatedOpkgRepositoryMapOutput() FederatedOpkgRepositoryMapOutput
	ToFederatedOpkgRepositoryMapOutputWithContext(context.Context) FederatedOpkgRepositoryMapOutput
}

type FederatedOpkgRepositoryMap map[string]FederatedOpkgRepositoryInput

func (FederatedOpkgRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedOpkgRepository)(nil)).Elem()
}

func (i FederatedOpkgRepositoryMap) ToFederatedOpkgRepositoryMapOutput() FederatedOpkgRepositoryMapOutput {
	return i.ToFederatedOpkgRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedOpkgRepositoryMap) ToFederatedOpkgRepositoryMapOutputWithContext(ctx context.Context) FederatedOpkgRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedOpkgRepositoryMapOutput)
}

type FederatedOpkgRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedOpkgRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedOpkgRepository)(nil)).Elem()
}

func (o FederatedOpkgRepositoryOutput) ToFederatedOpkgRepositoryOutput() FederatedOpkgRepositoryOutput {
	return o
}

func (o FederatedOpkgRepositoryOutput) ToFederatedOpkgRepositoryOutputWithContext(ctx context.Context) FederatedOpkgRepositoryOutput {
	return o
}

type FederatedOpkgRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedOpkgRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedOpkgRepository)(nil)).Elem()
}

func (o FederatedOpkgRepositoryArrayOutput) ToFederatedOpkgRepositoryArrayOutput() FederatedOpkgRepositoryArrayOutput {
	return o
}

func (o FederatedOpkgRepositoryArrayOutput) ToFederatedOpkgRepositoryArrayOutputWithContext(ctx context.Context) FederatedOpkgRepositoryArrayOutput {
	return o
}

func (o FederatedOpkgRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedOpkgRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedOpkgRepository {
		return vs[0].([]*FederatedOpkgRepository)[vs[1].(int)]
	}).(FederatedOpkgRepositoryOutput)
}

type FederatedOpkgRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedOpkgRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedOpkgRepository)(nil)).Elem()
}

func (o FederatedOpkgRepositoryMapOutput) ToFederatedOpkgRepositoryMapOutput() FederatedOpkgRepositoryMapOutput {
	return o
}

func (o FederatedOpkgRepositoryMapOutput) ToFederatedOpkgRepositoryMapOutputWithContext(ctx context.Context) FederatedOpkgRepositoryMapOutput {
	return o
}

func (o FederatedOpkgRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedOpkgRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedOpkgRepository {
		return vs[0].(map[string]*FederatedOpkgRepository)[vs[1].(string)]
	}).(FederatedOpkgRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedOpkgRepositoryInput)(nil)).Elem(), &FederatedOpkgRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedOpkgRepositoryArrayInput)(nil)).Elem(), FederatedOpkgRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedOpkgRepositoryMapInput)(nil)).Elem(), FederatedOpkgRepositoryMap{})
	pulumi.RegisterOutputType(FederatedOpkgRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedOpkgRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedOpkgRepositoryMapOutput{})
}
