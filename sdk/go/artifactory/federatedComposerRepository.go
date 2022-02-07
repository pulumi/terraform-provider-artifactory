// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Federated Composer Repository Resource
//
// Creates a federated Composer repository
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
// 		_, err := artifactory.NewFederatedComposerRepository(ctx, "terraform_federated_test_composer_repo", &artifactory.FederatedComposerRepositoryArgs{
// 			Key: pulumi.String("terraform-federated-test-composer-repo"),
// 			Members: FederatedComposerRepositoryMemberArray{
// 				&FederatedComposerRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-composer-repo"),
// 				},
// 				&FederatedComposerRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-composer-repo-2"),
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
type FederatedComposerRepository struct {
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
	Members     FederatedComposerRepositoryMemberArrayOutput `pulumi:"members"`
	Notes       pulumi.StringPtrOutput                       `pulumi:"notes"`
	PackageType pulumi.StringOutput                          `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrOutput     `pulumi:"priorityResolution"`
	PropertySets       pulumi.StringArrayOutput `pulumi:"propertySets"`
	RepoLayoutRef      pulumi.StringOutput      `pulumi:"repoLayoutRef"`
	XrayIndex          pulumi.BoolOutput        `pulumi:"xrayIndex"`
}

// NewFederatedComposerRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedComposerRepository(ctx *pulumi.Context,
	name string, args *FederatedComposerRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedComposerRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FederatedComposerRepository
	err := ctx.RegisterResource("artifactory:index/federatedComposerRepository:FederatedComposerRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedComposerRepository gets an existing FederatedComposerRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedComposerRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedComposerRepositoryState, opts ...pulumi.ResourceOption) (*FederatedComposerRepository, error) {
	var resource FederatedComposerRepository
	err := ctx.ReadResource("artifactory:index/federatedComposerRepository:FederatedComposerRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedComposerRepository resources.
type federatedComposerRepositoryState struct {
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
	Members     []FederatedComposerRepositoryMember `pulumi:"members"`
	Notes       *string                             `pulumi:"notes"`
	PackageType *string                             `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool    `pulumi:"priorityResolution"`
	PropertySets       []string `pulumi:"propertySets"`
	RepoLayoutRef      *string  `pulumi:"repoLayoutRef"`
	XrayIndex          *bool    `pulumi:"xrayIndex"`
}

type FederatedComposerRepositoryState struct {
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
	Members     FederatedComposerRepositoryMemberArrayInput
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	PropertySets       pulumi.StringArrayInput
	RepoLayoutRef      pulumi.StringPtrInput
	XrayIndex          pulumi.BoolPtrInput
}

func (FederatedComposerRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedComposerRepositoryState)(nil)).Elem()
}

type federatedComposerRepositoryArgs struct {
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
	Members []FederatedComposerRepositoryMember `pulumi:"members"`
	Notes   *string                             `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool    `pulumi:"priorityResolution"`
	PropertySets       []string `pulumi:"propertySets"`
	RepoLayoutRef      *string  `pulumi:"repoLayoutRef"`
	XrayIndex          *bool    `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a FederatedComposerRepository resource.
type FederatedComposerRepositoryArgs struct {
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
	Members FederatedComposerRepositoryMemberArrayInput
	Notes   pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	PropertySets       pulumi.StringArrayInput
	RepoLayoutRef      pulumi.StringPtrInput
	XrayIndex          pulumi.BoolPtrInput
}

func (FederatedComposerRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedComposerRepositoryArgs)(nil)).Elem()
}

type FederatedComposerRepositoryInput interface {
	pulumi.Input

	ToFederatedComposerRepositoryOutput() FederatedComposerRepositoryOutput
	ToFederatedComposerRepositoryOutputWithContext(ctx context.Context) FederatedComposerRepositoryOutput
}

func (*FederatedComposerRepository) ElementType() reflect.Type {
	return reflect.TypeOf((*FederatedComposerRepository)(nil))
}

func (i *FederatedComposerRepository) ToFederatedComposerRepositoryOutput() FederatedComposerRepositoryOutput {
	return i.ToFederatedComposerRepositoryOutputWithContext(context.Background())
}

func (i *FederatedComposerRepository) ToFederatedComposerRepositoryOutputWithContext(ctx context.Context) FederatedComposerRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedComposerRepositoryOutput)
}

func (i *FederatedComposerRepository) ToFederatedComposerRepositoryPtrOutput() FederatedComposerRepositoryPtrOutput {
	return i.ToFederatedComposerRepositoryPtrOutputWithContext(context.Background())
}

func (i *FederatedComposerRepository) ToFederatedComposerRepositoryPtrOutputWithContext(ctx context.Context) FederatedComposerRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedComposerRepositoryPtrOutput)
}

type FederatedComposerRepositoryPtrInput interface {
	pulumi.Input

	ToFederatedComposerRepositoryPtrOutput() FederatedComposerRepositoryPtrOutput
	ToFederatedComposerRepositoryPtrOutputWithContext(ctx context.Context) FederatedComposerRepositoryPtrOutput
}

type federatedComposerRepositoryPtrType FederatedComposerRepositoryArgs

func (*federatedComposerRepositoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedComposerRepository)(nil))
}

func (i *federatedComposerRepositoryPtrType) ToFederatedComposerRepositoryPtrOutput() FederatedComposerRepositoryPtrOutput {
	return i.ToFederatedComposerRepositoryPtrOutputWithContext(context.Background())
}

func (i *federatedComposerRepositoryPtrType) ToFederatedComposerRepositoryPtrOutputWithContext(ctx context.Context) FederatedComposerRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedComposerRepositoryPtrOutput)
}

// FederatedComposerRepositoryArrayInput is an input type that accepts FederatedComposerRepositoryArray and FederatedComposerRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedComposerRepositoryArrayInput` via:
//
//          FederatedComposerRepositoryArray{ FederatedComposerRepositoryArgs{...} }
type FederatedComposerRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedComposerRepositoryArrayOutput() FederatedComposerRepositoryArrayOutput
	ToFederatedComposerRepositoryArrayOutputWithContext(context.Context) FederatedComposerRepositoryArrayOutput
}

type FederatedComposerRepositoryArray []FederatedComposerRepositoryInput

func (FederatedComposerRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedComposerRepository)(nil)).Elem()
}

func (i FederatedComposerRepositoryArray) ToFederatedComposerRepositoryArrayOutput() FederatedComposerRepositoryArrayOutput {
	return i.ToFederatedComposerRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedComposerRepositoryArray) ToFederatedComposerRepositoryArrayOutputWithContext(ctx context.Context) FederatedComposerRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedComposerRepositoryArrayOutput)
}

// FederatedComposerRepositoryMapInput is an input type that accepts FederatedComposerRepositoryMap and FederatedComposerRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedComposerRepositoryMapInput` via:
//
//          FederatedComposerRepositoryMap{ "key": FederatedComposerRepositoryArgs{...} }
type FederatedComposerRepositoryMapInput interface {
	pulumi.Input

	ToFederatedComposerRepositoryMapOutput() FederatedComposerRepositoryMapOutput
	ToFederatedComposerRepositoryMapOutputWithContext(context.Context) FederatedComposerRepositoryMapOutput
}

type FederatedComposerRepositoryMap map[string]FederatedComposerRepositoryInput

func (FederatedComposerRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedComposerRepository)(nil)).Elem()
}

func (i FederatedComposerRepositoryMap) ToFederatedComposerRepositoryMapOutput() FederatedComposerRepositoryMapOutput {
	return i.ToFederatedComposerRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedComposerRepositoryMap) ToFederatedComposerRepositoryMapOutputWithContext(ctx context.Context) FederatedComposerRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedComposerRepositoryMapOutput)
}

type FederatedComposerRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedComposerRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FederatedComposerRepository)(nil))
}

func (o FederatedComposerRepositoryOutput) ToFederatedComposerRepositoryOutput() FederatedComposerRepositoryOutput {
	return o
}

func (o FederatedComposerRepositoryOutput) ToFederatedComposerRepositoryOutputWithContext(ctx context.Context) FederatedComposerRepositoryOutput {
	return o
}

func (o FederatedComposerRepositoryOutput) ToFederatedComposerRepositoryPtrOutput() FederatedComposerRepositoryPtrOutput {
	return o.ToFederatedComposerRepositoryPtrOutputWithContext(context.Background())
}

func (o FederatedComposerRepositoryOutput) ToFederatedComposerRepositoryPtrOutputWithContext(ctx context.Context) FederatedComposerRepositoryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FederatedComposerRepository) *FederatedComposerRepository {
		return &v
	}).(FederatedComposerRepositoryPtrOutput)
}

type FederatedComposerRepositoryPtrOutput struct{ *pulumi.OutputState }

func (FederatedComposerRepositoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedComposerRepository)(nil))
}

func (o FederatedComposerRepositoryPtrOutput) ToFederatedComposerRepositoryPtrOutput() FederatedComposerRepositoryPtrOutput {
	return o
}

func (o FederatedComposerRepositoryPtrOutput) ToFederatedComposerRepositoryPtrOutputWithContext(ctx context.Context) FederatedComposerRepositoryPtrOutput {
	return o
}

func (o FederatedComposerRepositoryPtrOutput) Elem() FederatedComposerRepositoryOutput {
	return o.ApplyT(func(v *FederatedComposerRepository) FederatedComposerRepository {
		if v != nil {
			return *v
		}
		var ret FederatedComposerRepository
		return ret
	}).(FederatedComposerRepositoryOutput)
}

type FederatedComposerRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedComposerRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FederatedComposerRepository)(nil))
}

func (o FederatedComposerRepositoryArrayOutput) ToFederatedComposerRepositoryArrayOutput() FederatedComposerRepositoryArrayOutput {
	return o
}

func (o FederatedComposerRepositoryArrayOutput) ToFederatedComposerRepositoryArrayOutputWithContext(ctx context.Context) FederatedComposerRepositoryArrayOutput {
	return o
}

func (o FederatedComposerRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedComposerRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FederatedComposerRepository {
		return vs[0].([]FederatedComposerRepository)[vs[1].(int)]
	}).(FederatedComposerRepositoryOutput)
}

type FederatedComposerRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedComposerRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FederatedComposerRepository)(nil))
}

func (o FederatedComposerRepositoryMapOutput) ToFederatedComposerRepositoryMapOutput() FederatedComposerRepositoryMapOutput {
	return o
}

func (o FederatedComposerRepositoryMapOutput) ToFederatedComposerRepositoryMapOutputWithContext(ctx context.Context) FederatedComposerRepositoryMapOutput {
	return o
}

func (o FederatedComposerRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedComposerRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FederatedComposerRepository {
		return vs[0].(map[string]FederatedComposerRepository)[vs[1].(string)]
	}).(FederatedComposerRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedComposerRepositoryInput)(nil)).Elem(), &FederatedComposerRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedComposerRepositoryPtrInput)(nil)).Elem(), &FederatedComposerRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedComposerRepositoryArrayInput)(nil)).Elem(), FederatedComposerRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedComposerRepositoryMapInput)(nil)).Elem(), FederatedComposerRepositoryMap{})
	pulumi.RegisterOutputType(FederatedComposerRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedComposerRepositoryPtrOutput{})
	pulumi.RegisterOutputType(FederatedComposerRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedComposerRepositoryMapOutput{})
}
