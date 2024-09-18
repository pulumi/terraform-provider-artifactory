// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource enables you to creates a new Release Bundle v2, uniquely identified by a combination of repository key, name, and version. For more information, see [Understanding Release Bundles v2](https://jfrog.com/help/r/jfrog-artifactory-documentation/understanding-release-bundles-v2) and [REST API](https://jfrog.com/help/r/jfrog-rest-apis/create-release-bundle-v2-version).
type ReleaseBundleV2 struct {
	pulumi.CustomResourceState

	// Timestamp when the new version was created (ISO 8601 standard).
	Created pulumi.StringOutput `pulumi:"created"`
	// The user who created the Release Bundle.
	CreatedBy pulumi.StringOutput `pulumi:"createdBy"`
	// Key-pair name to use for signature creation
	KeypairName pulumi.StringOutput `pulumi:"keypairName"`
	// Name of Release Bundle
	Name pulumi.StringOutput `pulumi:"name"`
	// Project key the Release Bundle belongs to
	ProjectKey pulumi.StringPtrOutput `pulumi:"projectKey"`
	// The unique identifier of the Artifactory instance where the Release Bundle was created.
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
	// Determines whether to skip the resolution of the Docker manifest, which adds the image layers to the Release Bundle. The default value is `false` (the manifest is resolved and image layers are included).
	SkipDockerManifestResolution pulumi.BoolOutput `pulumi:"skipDockerManifestResolution"`
	// Defines specific repositories to include in the promotion. If this property is left undefined, all repositories (except those specifically excluded) are included in the promotion. Important: If one or more repositories are specifically included, all other repositories are excluded (regardless of what is defined in `excludedRepositoryKeys`).
	Source ReleaseBundleV2SourceOutput `pulumi:"source"`
	// Source type. Valid values: `aql`, `artifacts`, `builds`, `releaseBundles`
	SourceType pulumi.StringOutput `pulumi:"sourceType"`
	// Version to promote
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewReleaseBundleV2 registers a new resource with the given unique name, arguments, and options.
func NewReleaseBundleV2(ctx *pulumi.Context,
	name string, args *ReleaseBundleV2Args, opts ...pulumi.ResourceOption) (*ReleaseBundleV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeypairName == nil {
		return nil, errors.New("invalid value for required argument 'KeypairName'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	if args.SourceType == nil {
		return nil, errors.New("invalid value for required argument 'SourceType'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ReleaseBundleV2
	err := ctx.RegisterResource("artifactory:index/releaseBundleV2:ReleaseBundleV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReleaseBundleV2 gets an existing ReleaseBundleV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReleaseBundleV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReleaseBundleV2State, opts ...pulumi.ResourceOption) (*ReleaseBundleV2, error) {
	var resource ReleaseBundleV2
	err := ctx.ReadResource("artifactory:index/releaseBundleV2:ReleaseBundleV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReleaseBundleV2 resources.
type releaseBundleV2State struct {
	// Timestamp when the new version was created (ISO 8601 standard).
	Created *string `pulumi:"created"`
	// The user who created the Release Bundle.
	CreatedBy *string `pulumi:"createdBy"`
	// Key-pair name to use for signature creation
	KeypairName *string `pulumi:"keypairName"`
	// Name of Release Bundle
	Name *string `pulumi:"name"`
	// Project key the Release Bundle belongs to
	ProjectKey *string `pulumi:"projectKey"`
	// The unique identifier of the Artifactory instance where the Release Bundle was created.
	ServiceId *string `pulumi:"serviceId"`
	// Determines whether to skip the resolution of the Docker manifest, which adds the image layers to the Release Bundle. The default value is `false` (the manifest is resolved and image layers are included).
	SkipDockerManifestResolution *bool `pulumi:"skipDockerManifestResolution"`
	// Defines specific repositories to include in the promotion. If this property is left undefined, all repositories (except those specifically excluded) are included in the promotion. Important: If one or more repositories are specifically included, all other repositories are excluded (regardless of what is defined in `excludedRepositoryKeys`).
	Source *ReleaseBundleV2Source `pulumi:"source"`
	// Source type. Valid values: `aql`, `artifacts`, `builds`, `releaseBundles`
	SourceType *string `pulumi:"sourceType"`
	// Version to promote
	Version *string `pulumi:"version"`
}

type ReleaseBundleV2State struct {
	// Timestamp when the new version was created (ISO 8601 standard).
	Created pulumi.StringPtrInput
	// The user who created the Release Bundle.
	CreatedBy pulumi.StringPtrInput
	// Key-pair name to use for signature creation
	KeypairName pulumi.StringPtrInput
	// Name of Release Bundle
	Name pulumi.StringPtrInput
	// Project key the Release Bundle belongs to
	ProjectKey pulumi.StringPtrInput
	// The unique identifier of the Artifactory instance where the Release Bundle was created.
	ServiceId pulumi.StringPtrInput
	// Determines whether to skip the resolution of the Docker manifest, which adds the image layers to the Release Bundle. The default value is `false` (the manifest is resolved and image layers are included).
	SkipDockerManifestResolution pulumi.BoolPtrInput
	// Defines specific repositories to include in the promotion. If this property is left undefined, all repositories (except those specifically excluded) are included in the promotion. Important: If one or more repositories are specifically included, all other repositories are excluded (regardless of what is defined in `excludedRepositoryKeys`).
	Source ReleaseBundleV2SourcePtrInput
	// Source type. Valid values: `aql`, `artifacts`, `builds`, `releaseBundles`
	SourceType pulumi.StringPtrInput
	// Version to promote
	Version pulumi.StringPtrInput
}

func (ReleaseBundleV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseBundleV2State)(nil)).Elem()
}

type releaseBundleV2Args struct {
	// Key-pair name to use for signature creation
	KeypairName string `pulumi:"keypairName"`
	// Name of Release Bundle
	Name *string `pulumi:"name"`
	// Project key the Release Bundle belongs to
	ProjectKey *string `pulumi:"projectKey"`
	// Determines whether to skip the resolution of the Docker manifest, which adds the image layers to the Release Bundle. The default value is `false` (the manifest is resolved and image layers are included).
	SkipDockerManifestResolution *bool `pulumi:"skipDockerManifestResolution"`
	// Defines specific repositories to include in the promotion. If this property is left undefined, all repositories (except those specifically excluded) are included in the promotion. Important: If one or more repositories are specifically included, all other repositories are excluded (regardless of what is defined in `excludedRepositoryKeys`).
	Source ReleaseBundleV2Source `pulumi:"source"`
	// Source type. Valid values: `aql`, `artifacts`, `builds`, `releaseBundles`
	SourceType string `pulumi:"sourceType"`
	// Version to promote
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a ReleaseBundleV2 resource.
type ReleaseBundleV2Args struct {
	// Key-pair name to use for signature creation
	KeypairName pulumi.StringInput
	// Name of Release Bundle
	Name pulumi.StringPtrInput
	// Project key the Release Bundle belongs to
	ProjectKey pulumi.StringPtrInput
	// Determines whether to skip the resolution of the Docker manifest, which adds the image layers to the Release Bundle. The default value is `false` (the manifest is resolved and image layers are included).
	SkipDockerManifestResolution pulumi.BoolPtrInput
	// Defines specific repositories to include in the promotion. If this property is left undefined, all repositories (except those specifically excluded) are included in the promotion. Important: If one or more repositories are specifically included, all other repositories are excluded (regardless of what is defined in `excludedRepositoryKeys`).
	Source ReleaseBundleV2SourceInput
	// Source type. Valid values: `aql`, `artifacts`, `builds`, `releaseBundles`
	SourceType pulumi.StringInput
	// Version to promote
	Version pulumi.StringInput
}

func (ReleaseBundleV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseBundleV2Args)(nil)).Elem()
}

type ReleaseBundleV2Input interface {
	pulumi.Input

	ToReleaseBundleV2Output() ReleaseBundleV2Output
	ToReleaseBundleV2OutputWithContext(ctx context.Context) ReleaseBundleV2Output
}

func (*ReleaseBundleV2) ElementType() reflect.Type {
	return reflect.TypeOf((**ReleaseBundleV2)(nil)).Elem()
}

func (i *ReleaseBundleV2) ToReleaseBundleV2Output() ReleaseBundleV2Output {
	return i.ToReleaseBundleV2OutputWithContext(context.Background())
}

func (i *ReleaseBundleV2) ToReleaseBundleV2OutputWithContext(ctx context.Context) ReleaseBundleV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseBundleV2Output)
}

// ReleaseBundleV2ArrayInput is an input type that accepts ReleaseBundleV2Array and ReleaseBundleV2ArrayOutput values.
// You can construct a concrete instance of `ReleaseBundleV2ArrayInput` via:
//
//	ReleaseBundleV2Array{ ReleaseBundleV2Args{...} }
type ReleaseBundleV2ArrayInput interface {
	pulumi.Input

	ToReleaseBundleV2ArrayOutput() ReleaseBundleV2ArrayOutput
	ToReleaseBundleV2ArrayOutputWithContext(context.Context) ReleaseBundleV2ArrayOutput
}

type ReleaseBundleV2Array []ReleaseBundleV2Input

func (ReleaseBundleV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReleaseBundleV2)(nil)).Elem()
}

func (i ReleaseBundleV2Array) ToReleaseBundleV2ArrayOutput() ReleaseBundleV2ArrayOutput {
	return i.ToReleaseBundleV2ArrayOutputWithContext(context.Background())
}

func (i ReleaseBundleV2Array) ToReleaseBundleV2ArrayOutputWithContext(ctx context.Context) ReleaseBundleV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseBundleV2ArrayOutput)
}

// ReleaseBundleV2MapInput is an input type that accepts ReleaseBundleV2Map and ReleaseBundleV2MapOutput values.
// You can construct a concrete instance of `ReleaseBundleV2MapInput` via:
//
//	ReleaseBundleV2Map{ "key": ReleaseBundleV2Args{...} }
type ReleaseBundleV2MapInput interface {
	pulumi.Input

	ToReleaseBundleV2MapOutput() ReleaseBundleV2MapOutput
	ToReleaseBundleV2MapOutputWithContext(context.Context) ReleaseBundleV2MapOutput
}

type ReleaseBundleV2Map map[string]ReleaseBundleV2Input

func (ReleaseBundleV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReleaseBundleV2)(nil)).Elem()
}

func (i ReleaseBundleV2Map) ToReleaseBundleV2MapOutput() ReleaseBundleV2MapOutput {
	return i.ToReleaseBundleV2MapOutputWithContext(context.Background())
}

func (i ReleaseBundleV2Map) ToReleaseBundleV2MapOutputWithContext(ctx context.Context) ReleaseBundleV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseBundleV2MapOutput)
}

type ReleaseBundleV2Output struct{ *pulumi.OutputState }

func (ReleaseBundleV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**ReleaseBundleV2)(nil)).Elem()
}

func (o ReleaseBundleV2Output) ToReleaseBundleV2Output() ReleaseBundleV2Output {
	return o
}

func (o ReleaseBundleV2Output) ToReleaseBundleV2OutputWithContext(ctx context.Context) ReleaseBundleV2Output {
	return o
}

// Timestamp when the new version was created (ISO 8601 standard).
func (o ReleaseBundleV2Output) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *ReleaseBundleV2) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

// The user who created the Release Bundle.
func (o ReleaseBundleV2Output) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *ReleaseBundleV2) pulumi.StringOutput { return v.CreatedBy }).(pulumi.StringOutput)
}

// Key-pair name to use for signature creation
func (o ReleaseBundleV2Output) KeypairName() pulumi.StringOutput {
	return o.ApplyT(func(v *ReleaseBundleV2) pulumi.StringOutput { return v.KeypairName }).(pulumi.StringOutput)
}

// Name of Release Bundle
func (o ReleaseBundleV2Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReleaseBundleV2) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Project key the Release Bundle belongs to
func (o ReleaseBundleV2Output) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReleaseBundleV2) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// The unique identifier of the Artifactory instance where the Release Bundle was created.
func (o ReleaseBundleV2Output) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReleaseBundleV2) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

// Determines whether to skip the resolution of the Docker manifest, which adds the image layers to the Release Bundle. The default value is `false` (the manifest is resolved and image layers are included).
func (o ReleaseBundleV2Output) SkipDockerManifestResolution() pulumi.BoolOutput {
	return o.ApplyT(func(v *ReleaseBundleV2) pulumi.BoolOutput { return v.SkipDockerManifestResolution }).(pulumi.BoolOutput)
}

// Defines specific repositories to include in the promotion. If this property is left undefined, all repositories (except those specifically excluded) are included in the promotion. Important: If one or more repositories are specifically included, all other repositories are excluded (regardless of what is defined in `excludedRepositoryKeys`).
func (o ReleaseBundleV2Output) Source() ReleaseBundleV2SourceOutput {
	return o.ApplyT(func(v *ReleaseBundleV2) ReleaseBundleV2SourceOutput { return v.Source }).(ReleaseBundleV2SourceOutput)
}

// Source type. Valid values: `aql`, `artifacts`, `builds`, `releaseBundles`
func (o ReleaseBundleV2Output) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *ReleaseBundleV2) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

// Version to promote
func (o ReleaseBundleV2Output) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *ReleaseBundleV2) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type ReleaseBundleV2ArrayOutput struct{ *pulumi.OutputState }

func (ReleaseBundleV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReleaseBundleV2)(nil)).Elem()
}

func (o ReleaseBundleV2ArrayOutput) ToReleaseBundleV2ArrayOutput() ReleaseBundleV2ArrayOutput {
	return o
}

func (o ReleaseBundleV2ArrayOutput) ToReleaseBundleV2ArrayOutputWithContext(ctx context.Context) ReleaseBundleV2ArrayOutput {
	return o
}

func (o ReleaseBundleV2ArrayOutput) Index(i pulumi.IntInput) ReleaseBundleV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReleaseBundleV2 {
		return vs[0].([]*ReleaseBundleV2)[vs[1].(int)]
	}).(ReleaseBundleV2Output)
}

type ReleaseBundleV2MapOutput struct{ *pulumi.OutputState }

func (ReleaseBundleV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReleaseBundleV2)(nil)).Elem()
}

func (o ReleaseBundleV2MapOutput) ToReleaseBundleV2MapOutput() ReleaseBundleV2MapOutput {
	return o
}

func (o ReleaseBundleV2MapOutput) ToReleaseBundleV2MapOutputWithContext(ctx context.Context) ReleaseBundleV2MapOutput {
	return o
}

func (o ReleaseBundleV2MapOutput) MapIndex(k pulumi.StringInput) ReleaseBundleV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReleaseBundleV2 {
		return vs[0].(map[string]*ReleaseBundleV2)[vs[1].(string)]
	}).(ReleaseBundleV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseBundleV2Input)(nil)).Elem(), &ReleaseBundleV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseBundleV2ArrayInput)(nil)).Elem(), ReleaseBundleV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseBundleV2MapInput)(nil)).Elem(), ReleaseBundleV2Map{})
	pulumi.RegisterOutputType(ReleaseBundleV2Output{})
	pulumi.RegisterOutputType(ReleaseBundleV2ArrayOutput{})
	pulumi.RegisterOutputType(ReleaseBundleV2MapOutput{})
}
