// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FederatedNpmRepository struct {
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
	Key                    pulumi.StringOutput    `pulumi:"key"`
	// The list of Federated members. If a Federated member receives a request that does not include the repository URL, it
	// will automatically be added with the combination of the configured base URL and `key` field value. Note that each of the
	// federated members will need to have a base URL set. Please follow the
	// [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members     FederatedNpmRepositoryMemberArrayOutput `pulumi:"members"`
	Notes       pulumi.StringPtrOutput                  `pulumi:"notes"`
	PackageType pulumi.StringOutput                     `pulumi:"packageType"`
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

// NewFederatedNpmRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedNpmRepository(ctx *pulumi.Context,
	name string, args *FederatedNpmRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedNpmRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FederatedNpmRepository
	err := ctx.RegisterResource("artifactory:index/federatedNpmRepository:FederatedNpmRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedNpmRepository gets an existing FederatedNpmRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedNpmRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedNpmRepositoryState, opts ...pulumi.ResourceOption) (*FederatedNpmRepository, error) {
	var resource FederatedNpmRepository
	err := ctx.ReadResource("artifactory:index/federatedNpmRepository:FederatedNpmRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedNpmRepository resources.
type federatedNpmRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	IncludesPattern        *string `pulumi:"includesPattern"`
	Key                    *string `pulumi:"key"`
	// The list of Federated members. If a Federated member receives a request that does not include the repository URL, it
	// will automatically be added with the combination of the configured base URL and `key` field value. Note that each of the
	// federated members will need to have a base URL set. Please follow the
	// [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members     []FederatedNpmRepositoryMember `pulumi:"members"`
	Notes       *string                        `pulumi:"notes"`
	PackageType *string                        `pulumi:"packageType"`
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

type FederatedNpmRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	BlackedOut             pulumi.BoolPtrInput
	Description            pulumi.StringPtrInput
	DownloadDirect         pulumi.BoolPtrInput
	ExcludesPattern        pulumi.StringPtrInput
	IncludesPattern        pulumi.StringPtrInput
	Key                    pulumi.StringPtrInput
	// The list of Federated members. If a Federated member receives a request that does not include the repository URL, it
	// will automatically be added with the combination of the configured base URL and `key` field value. Note that each of the
	// federated members will need to have a base URL set. Please follow the
	// [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members     FederatedNpmRepositoryMemberArrayInput
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

func (FederatedNpmRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedNpmRepositoryState)(nil)).Elem()
}

type federatedNpmRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	IncludesPattern        *string `pulumi:"includesPattern"`
	Key                    string  `pulumi:"key"`
	// The list of Federated members. If a Federated member receives a request that does not include the repository URL, it
	// will automatically be added with the combination of the configured base URL and `key` field value. Note that each of the
	// federated members will need to have a base URL set. Please follow the
	// [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members []FederatedNpmRepositoryMember `pulumi:"members"`
	Notes   *string                        `pulumi:"notes"`
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

// The set of arguments for constructing a FederatedNpmRepository resource.
type FederatedNpmRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	BlackedOut             pulumi.BoolPtrInput
	Description            pulumi.StringPtrInput
	DownloadDirect         pulumi.BoolPtrInput
	ExcludesPattern        pulumi.StringPtrInput
	IncludesPattern        pulumi.StringPtrInput
	Key                    pulumi.StringInput
	// The list of Federated members. If a Federated member receives a request that does not include the repository URL, it
	// will automatically be added with the combination of the configured base URL and `key` field value. Note that each of the
	// federated members will need to have a base URL set. Please follow the
	// [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members FederatedNpmRepositoryMemberArrayInput
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

func (FederatedNpmRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedNpmRepositoryArgs)(nil)).Elem()
}

type FederatedNpmRepositoryInput interface {
	pulumi.Input

	ToFederatedNpmRepositoryOutput() FederatedNpmRepositoryOutput
	ToFederatedNpmRepositoryOutputWithContext(ctx context.Context) FederatedNpmRepositoryOutput
}

func (*FederatedNpmRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedNpmRepository)(nil)).Elem()
}

func (i *FederatedNpmRepository) ToFederatedNpmRepositoryOutput() FederatedNpmRepositoryOutput {
	return i.ToFederatedNpmRepositoryOutputWithContext(context.Background())
}

func (i *FederatedNpmRepository) ToFederatedNpmRepositoryOutputWithContext(ctx context.Context) FederatedNpmRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedNpmRepositoryOutput)
}

// FederatedNpmRepositoryArrayInput is an input type that accepts FederatedNpmRepositoryArray and FederatedNpmRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedNpmRepositoryArrayInput` via:
//
//          FederatedNpmRepositoryArray{ FederatedNpmRepositoryArgs{...} }
type FederatedNpmRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedNpmRepositoryArrayOutput() FederatedNpmRepositoryArrayOutput
	ToFederatedNpmRepositoryArrayOutputWithContext(context.Context) FederatedNpmRepositoryArrayOutput
}

type FederatedNpmRepositoryArray []FederatedNpmRepositoryInput

func (FederatedNpmRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedNpmRepository)(nil)).Elem()
}

func (i FederatedNpmRepositoryArray) ToFederatedNpmRepositoryArrayOutput() FederatedNpmRepositoryArrayOutput {
	return i.ToFederatedNpmRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedNpmRepositoryArray) ToFederatedNpmRepositoryArrayOutputWithContext(ctx context.Context) FederatedNpmRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedNpmRepositoryArrayOutput)
}

// FederatedNpmRepositoryMapInput is an input type that accepts FederatedNpmRepositoryMap and FederatedNpmRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedNpmRepositoryMapInput` via:
//
//          FederatedNpmRepositoryMap{ "key": FederatedNpmRepositoryArgs{...} }
type FederatedNpmRepositoryMapInput interface {
	pulumi.Input

	ToFederatedNpmRepositoryMapOutput() FederatedNpmRepositoryMapOutput
	ToFederatedNpmRepositoryMapOutputWithContext(context.Context) FederatedNpmRepositoryMapOutput
}

type FederatedNpmRepositoryMap map[string]FederatedNpmRepositoryInput

func (FederatedNpmRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedNpmRepository)(nil)).Elem()
}

func (i FederatedNpmRepositoryMap) ToFederatedNpmRepositoryMapOutput() FederatedNpmRepositoryMapOutput {
	return i.ToFederatedNpmRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedNpmRepositoryMap) ToFederatedNpmRepositoryMapOutputWithContext(ctx context.Context) FederatedNpmRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedNpmRepositoryMapOutput)
}

type FederatedNpmRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedNpmRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedNpmRepository)(nil)).Elem()
}

func (o FederatedNpmRepositoryOutput) ToFederatedNpmRepositoryOutput() FederatedNpmRepositoryOutput {
	return o
}

func (o FederatedNpmRepositoryOutput) ToFederatedNpmRepositoryOutputWithContext(ctx context.Context) FederatedNpmRepositoryOutput {
	return o
}

type FederatedNpmRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedNpmRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedNpmRepository)(nil)).Elem()
}

func (o FederatedNpmRepositoryArrayOutput) ToFederatedNpmRepositoryArrayOutput() FederatedNpmRepositoryArrayOutput {
	return o
}

func (o FederatedNpmRepositoryArrayOutput) ToFederatedNpmRepositoryArrayOutputWithContext(ctx context.Context) FederatedNpmRepositoryArrayOutput {
	return o
}

func (o FederatedNpmRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedNpmRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedNpmRepository {
		return vs[0].([]*FederatedNpmRepository)[vs[1].(int)]
	}).(FederatedNpmRepositoryOutput)
}

type FederatedNpmRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedNpmRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedNpmRepository)(nil)).Elem()
}

func (o FederatedNpmRepositoryMapOutput) ToFederatedNpmRepositoryMapOutput() FederatedNpmRepositoryMapOutput {
	return o
}

func (o FederatedNpmRepositoryMapOutput) ToFederatedNpmRepositoryMapOutputWithContext(ctx context.Context) FederatedNpmRepositoryMapOutput {
	return o
}

func (o FederatedNpmRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedNpmRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedNpmRepository {
		return vs[0].(map[string]*FederatedNpmRepository)[vs[1].(string)]
	}).(FederatedNpmRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedNpmRepositoryInput)(nil)).Elem(), &FederatedNpmRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedNpmRepositoryArrayInput)(nil)).Elem(), FederatedNpmRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedNpmRepositoryMapInput)(nil)).Elem(), FederatedNpmRepositoryMap{})
	pulumi.RegisterOutputType(FederatedNpmRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedNpmRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedNpmRepositoryMapOutput{})
}
