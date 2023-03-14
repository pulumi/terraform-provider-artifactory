// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a local Debian repository and allows for the creation of a GPG key.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//	"os"
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v3/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := os.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewKeypair(ctx, "some-keypairGPG1", &artifactory.KeypairArgs{
//				PairName:   pulumi.String(fmt.Sprintf("some-keypair%v", random_id.Randid.Id)),
//				PairType:   pulumi.String("GPG"),
//				Alias:      pulumi.String("foo-alias1"),
//				PrivateKey: readFileOrPanic("samples/gpg.priv"),
//				PublicKey:  readFileOrPanic("samples/gpg.pub"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = artifactory.NewKeypair(ctx, "some-keypairGPG2", &artifactory.KeypairArgs{
//				PairName:   pulumi.String(fmt.Sprintf("some-keypair4%v", random_id.Randid.Id)),
//				PairType:   pulumi.String("GPG"),
//				Alias:      pulumi.String("foo-alias2"),
//				PrivateKey: readFileOrPanic("samples/gpg.priv"),
//				PublicKey:  readFileOrPanic("samples/gpg.pub"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = artifactory.NewDebianRepository(ctx, "my-debian-repo", &artifactory.DebianRepositoryArgs{
//				Key:                 pulumi.String("my-debian-repo"),
//				PrimaryKeypairRef:   some_keypairGPG1.PairName,
//				SecondaryKeypairRef: some_keypairGPG2.PairName,
//				IndexCompressionFormats: pulumi.StringArray{
//					pulumi.String("bz2"),
//					pulumi.String("lzma"),
//					pulumi.String("xz"),
//				},
//				TrivialLayout: pulumi.Bool(true),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				some_keypairGPG1,
//				some_keypairGPG2,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Local repositories can be imported using their name, e.g.
//
// ```sh
//
//	$ pulumi import artifactory:index/debianRepository:DebianRepository my-debian-repo my-debian-repo
//
// ```
type DebianRepository struct {
	pulumi.CustomResourceState

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrOutput `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrOutput `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolPtrOutput   `pulumi:"cdnRedirect"`
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrOutput `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringOutput `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringOutput `pulumi:"includesPattern"`
	// The options are Bzip2 (.bz2 extension) (default), LZMA (.lzma extension)
	// and XZ (.xz extension).
	IndexCompressionFormats pulumi.StringArrayOutput `pulumi:"indexCompressionFormats"`
	// the identity key of the repo.
	Key         pulumi.StringOutput    `pulumi:"key"`
	Notes       pulumi.StringPtrOutput `pulumi:"notes"`
	PackageType pulumi.StringOutput    `pulumi:"packageType"`
	// The primary RSA key to be used to sign packages.
	PrimaryKeypairRef pulumi.StringPtrOutput `pulumi:"primaryKeypairRef"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrOutput `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrOutput `pulumi:"projectKey"`
	// List of property set name
	PropertySets pulumi.StringArrayOutput `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	// The secondary RSA key to be used to sign packages.
	SecondaryKeypairRef pulumi.StringPtrOutput `pulumi:"secondaryKeypairRef"`
	// When set, the repository will use the deprecated trivial layout.
	//
	// Deprecated: You shouldn't be using this
	TrivialLayout pulumi.BoolPtrOutput `pulumi:"trivialLayout"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrOutput `pulumi:"xrayIndex"`
}

// NewDebianRepository registers a new resource with the given unique name, arguments, and options.
func NewDebianRepository(ctx *pulumi.Context,
	name string, args *DebianRepositoryArgs, opts ...pulumi.ResourceOption) (*DebianRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource DebianRepository
	err := ctx.RegisterResource("artifactory:index/debianRepository:DebianRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDebianRepository gets an existing DebianRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDebianRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DebianRepositoryState, opts ...pulumi.ResourceOption) (*DebianRepository, error) {
	var resource DebianRepository
	err := ctx.ReadResource("artifactory:index/debianRepository:DebianRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DebianRepository resources.
type debianRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect *bool   `pulumi:"cdnRedirect"`
	Description *string `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect *bool `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// The options are Bzip2 (.bz2 extension) (default), LZMA (.lzma extension)
	// and XZ (.xz extension).
	IndexCompressionFormats []string `pulumi:"indexCompressionFormats"`
	// the identity key of the repo.
	Key         *string `pulumi:"key"`
	Notes       *string `pulumi:"notes"`
	PackageType *string `pulumi:"packageType"`
	// The primary RSA key to be used to sign packages.
	PrimaryKeypairRef *string `pulumi:"primaryKeypairRef"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// List of property set name
	PropertySets []string `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The secondary RSA key to be used to sign packages.
	SecondaryKeypairRef *string `pulumi:"secondaryKeypairRef"`
	// When set, the repository will use the deprecated trivial layout.
	//
	// Deprecated: You shouldn't be using this
	TrivialLayout *bool `pulumi:"trivialLayout"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

type DebianRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolPtrInput
	Description pulumi.StringPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// The options are Bzip2 (.bz2 extension) (default), LZMA (.lzma extension)
	// and XZ (.xz extension).
	IndexCompressionFormats pulumi.StringArrayInput
	// the identity key of the repo.
	Key         pulumi.StringPtrInput
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// The primary RSA key to be used to sign packages.
	PrimaryKeypairRef pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// List of property set name
	PropertySets pulumi.StringArrayInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// The secondary RSA key to be used to sign packages.
	SecondaryKeypairRef pulumi.StringPtrInput
	// When set, the repository will use the deprecated trivial layout.
	//
	// Deprecated: You shouldn't be using this
	TrivialLayout pulumi.BoolPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (DebianRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*debianRepositoryState)(nil)).Elem()
}

type debianRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect *bool   `pulumi:"cdnRedirect"`
	Description *string `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect *bool `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// The options are Bzip2 (.bz2 extension) (default), LZMA (.lzma extension)
	// and XZ (.xz extension).
	IndexCompressionFormats []string `pulumi:"indexCompressionFormats"`
	// the identity key of the repo.
	Key   string  `pulumi:"key"`
	Notes *string `pulumi:"notes"`
	// The primary RSA key to be used to sign packages.
	PrimaryKeypairRef *string `pulumi:"primaryKeypairRef"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// List of property set name
	PropertySets []string `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The secondary RSA key to be used to sign packages.
	SecondaryKeypairRef *string `pulumi:"secondaryKeypairRef"`
	// When set, the repository will use the deprecated trivial layout.
	//
	// Deprecated: You shouldn't be using this
	TrivialLayout *bool `pulumi:"trivialLayout"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a DebianRepository resource.
type DebianRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolPtrInput
	Description pulumi.StringPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// The options are Bzip2 (.bz2 extension) (default), LZMA (.lzma extension)
	// and XZ (.xz extension).
	IndexCompressionFormats pulumi.StringArrayInput
	// the identity key of the repo.
	Key   pulumi.StringInput
	Notes pulumi.StringPtrInput
	// The primary RSA key to be used to sign packages.
	PrimaryKeypairRef pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// List of property set name
	PropertySets pulumi.StringArrayInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// The secondary RSA key to be used to sign packages.
	SecondaryKeypairRef pulumi.StringPtrInput
	// When set, the repository will use the deprecated trivial layout.
	//
	// Deprecated: You shouldn't be using this
	TrivialLayout pulumi.BoolPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (DebianRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*debianRepositoryArgs)(nil)).Elem()
}

type DebianRepositoryInput interface {
	pulumi.Input

	ToDebianRepositoryOutput() DebianRepositoryOutput
	ToDebianRepositoryOutputWithContext(ctx context.Context) DebianRepositoryOutput
}

func (*DebianRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**DebianRepository)(nil)).Elem()
}

func (i *DebianRepository) ToDebianRepositoryOutput() DebianRepositoryOutput {
	return i.ToDebianRepositoryOutputWithContext(context.Background())
}

func (i *DebianRepository) ToDebianRepositoryOutputWithContext(ctx context.Context) DebianRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DebianRepositoryOutput)
}

// DebianRepositoryArrayInput is an input type that accepts DebianRepositoryArray and DebianRepositoryArrayOutput values.
// You can construct a concrete instance of `DebianRepositoryArrayInput` via:
//
//	DebianRepositoryArray{ DebianRepositoryArgs{...} }
type DebianRepositoryArrayInput interface {
	pulumi.Input

	ToDebianRepositoryArrayOutput() DebianRepositoryArrayOutput
	ToDebianRepositoryArrayOutputWithContext(context.Context) DebianRepositoryArrayOutput
}

type DebianRepositoryArray []DebianRepositoryInput

func (DebianRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DebianRepository)(nil)).Elem()
}

func (i DebianRepositoryArray) ToDebianRepositoryArrayOutput() DebianRepositoryArrayOutput {
	return i.ToDebianRepositoryArrayOutputWithContext(context.Background())
}

func (i DebianRepositoryArray) ToDebianRepositoryArrayOutputWithContext(ctx context.Context) DebianRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DebianRepositoryArrayOutput)
}

// DebianRepositoryMapInput is an input type that accepts DebianRepositoryMap and DebianRepositoryMapOutput values.
// You can construct a concrete instance of `DebianRepositoryMapInput` via:
//
//	DebianRepositoryMap{ "key": DebianRepositoryArgs{...} }
type DebianRepositoryMapInput interface {
	pulumi.Input

	ToDebianRepositoryMapOutput() DebianRepositoryMapOutput
	ToDebianRepositoryMapOutputWithContext(context.Context) DebianRepositoryMapOutput
}

type DebianRepositoryMap map[string]DebianRepositoryInput

func (DebianRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DebianRepository)(nil)).Elem()
}

func (i DebianRepositoryMap) ToDebianRepositoryMapOutput() DebianRepositoryMapOutput {
	return i.ToDebianRepositoryMapOutputWithContext(context.Background())
}

func (i DebianRepositoryMap) ToDebianRepositoryMapOutputWithContext(ctx context.Context) DebianRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DebianRepositoryMapOutput)
}

type DebianRepositoryOutput struct{ *pulumi.OutputState }

func (DebianRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DebianRepository)(nil)).Elem()
}

func (o DebianRepositoryOutput) ToDebianRepositoryOutput() DebianRepositoryOutput {
	return o
}

func (o DebianRepositoryOutput) ToDebianRepositoryOutputWithContext(ctx context.Context) DebianRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o DebianRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DebianRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o DebianRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DebianRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o DebianRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DebianRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o DebianRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DebianRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o DebianRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DebianRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
// artifacts are excluded.
func (o DebianRepositoryOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *DebianRepository) pulumi.StringOutput { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o DebianRepositoryOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *DebianRepository) pulumi.StringOutput { return v.IncludesPattern }).(pulumi.StringOutput)
}

// The options are Bzip2 (.bz2 extension) (default), LZMA (.lzma extension)
// and XZ (.xz extension).
func (o DebianRepositoryOutput) IndexCompressionFormats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DebianRepository) pulumi.StringArrayOutput { return v.IndexCompressionFormats }).(pulumi.StringArrayOutput)
}

// the identity key of the repo.
func (o DebianRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *DebianRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

func (o DebianRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DebianRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o DebianRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *DebianRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// The primary RSA key to be used to sign packages.
func (o DebianRepositoryOutput) PrimaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DebianRepository) pulumi.StringPtrOutput { return v.PrimaryKeypairRef }).(pulumi.StringPtrOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o DebianRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DebianRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
// will remain in the Terraform state, which will create state drift during the update.
func (o DebianRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DebianRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o DebianRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DebianRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o DebianRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DebianRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o DebianRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DebianRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The secondary RSA key to be used to sign packages.
func (o DebianRepositoryOutput) SecondaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DebianRepository) pulumi.StringPtrOutput { return v.SecondaryKeypairRef }).(pulumi.StringPtrOutput)
}

// When set, the repository will use the deprecated trivial layout.
//
// Deprecated: You shouldn't be using this
func (o DebianRepositoryOutput) TrivialLayout() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DebianRepository) pulumi.BoolPtrOutput { return v.TrivialLayout }).(pulumi.BoolPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o DebianRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DebianRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type DebianRepositoryArrayOutput struct{ *pulumi.OutputState }

func (DebianRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DebianRepository)(nil)).Elem()
}

func (o DebianRepositoryArrayOutput) ToDebianRepositoryArrayOutput() DebianRepositoryArrayOutput {
	return o
}

func (o DebianRepositoryArrayOutput) ToDebianRepositoryArrayOutputWithContext(ctx context.Context) DebianRepositoryArrayOutput {
	return o
}

func (o DebianRepositoryArrayOutput) Index(i pulumi.IntInput) DebianRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DebianRepository {
		return vs[0].([]*DebianRepository)[vs[1].(int)]
	}).(DebianRepositoryOutput)
}

type DebianRepositoryMapOutput struct{ *pulumi.OutputState }

func (DebianRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DebianRepository)(nil)).Elem()
}

func (o DebianRepositoryMapOutput) ToDebianRepositoryMapOutput() DebianRepositoryMapOutput {
	return o
}

func (o DebianRepositoryMapOutput) ToDebianRepositoryMapOutputWithContext(ctx context.Context) DebianRepositoryMapOutput {
	return o
}

func (o DebianRepositoryMapOutput) MapIndex(k pulumi.StringInput) DebianRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DebianRepository {
		return vs[0].(map[string]*DebianRepository)[vs[1].(string)]
	}).(DebianRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DebianRepositoryInput)(nil)).Elem(), &DebianRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*DebianRepositoryArrayInput)(nil)).Elem(), DebianRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DebianRepositoryMapInput)(nil)).Elem(), DebianRepositoryMap{})
	pulumi.RegisterOutputType(DebianRepositoryOutput{})
	pulumi.RegisterOutputType(DebianRepositoryArrayOutput{})
	pulumi.RegisterOutputType(DebianRepositoryMapOutput{})
}
