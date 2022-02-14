// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PermissionTargets struct {
	pulumi.CustomResourceState

	Build         PermissionTargetsBuildPtrOutput         `pulumi:"build"`
	Name          pulumi.StringOutput                     `pulumi:"name"`
	ReleaseBundle PermissionTargetsReleaseBundlePtrOutput `pulumi:"releaseBundle"`
	Repo          PermissionTargetsRepoPtrOutput          `pulumi:"repo"`
}

// NewPermissionTargets registers a new resource with the given unique name, arguments, and options.
func NewPermissionTargets(ctx *pulumi.Context,
	name string, args *PermissionTargetsArgs, opts ...pulumi.ResourceOption) (*PermissionTargets, error) {
	if args == nil {
		args = &PermissionTargetsArgs{}
	}

	var resource PermissionTargets
	err := ctx.RegisterResource("artifactory:index/permissionTargets:PermissionTargets", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPermissionTargets gets an existing PermissionTargets resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPermissionTargets(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PermissionTargetsState, opts ...pulumi.ResourceOption) (*PermissionTargets, error) {
	var resource PermissionTargets
	err := ctx.ReadResource("artifactory:index/permissionTargets:PermissionTargets", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PermissionTargets resources.
type permissionTargetsState struct {
	Build         *PermissionTargetsBuild         `pulumi:"build"`
	Name          *string                         `pulumi:"name"`
	ReleaseBundle *PermissionTargetsReleaseBundle `pulumi:"releaseBundle"`
	Repo          *PermissionTargetsRepo          `pulumi:"repo"`
}

type PermissionTargetsState struct {
	Build         PermissionTargetsBuildPtrInput
	Name          pulumi.StringPtrInput
	ReleaseBundle PermissionTargetsReleaseBundlePtrInput
	Repo          PermissionTargetsRepoPtrInput
}

func (PermissionTargetsState) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionTargetsState)(nil)).Elem()
}

type permissionTargetsArgs struct {
	Build         *PermissionTargetsBuild         `pulumi:"build"`
	Name          *string                         `pulumi:"name"`
	ReleaseBundle *PermissionTargetsReleaseBundle `pulumi:"releaseBundle"`
	Repo          *PermissionTargetsRepo          `pulumi:"repo"`
}

// The set of arguments for constructing a PermissionTargets resource.
type PermissionTargetsArgs struct {
	Build         PermissionTargetsBuildPtrInput
	Name          pulumi.StringPtrInput
	ReleaseBundle PermissionTargetsReleaseBundlePtrInput
	Repo          PermissionTargetsRepoPtrInput
}

func (PermissionTargetsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionTargetsArgs)(nil)).Elem()
}

type PermissionTargetsInput interface {
	pulumi.Input

	ToPermissionTargetsOutput() PermissionTargetsOutput
	ToPermissionTargetsOutputWithContext(ctx context.Context) PermissionTargetsOutput
}

func (*PermissionTargets) ElementType() reflect.Type {
	return reflect.TypeOf((**PermissionTargets)(nil)).Elem()
}

func (i *PermissionTargets) ToPermissionTargetsOutput() PermissionTargetsOutput {
	return i.ToPermissionTargetsOutputWithContext(context.Background())
}

func (i *PermissionTargets) ToPermissionTargetsOutputWithContext(ctx context.Context) PermissionTargetsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionTargetsOutput)
}

// PermissionTargetsArrayInput is an input type that accepts PermissionTargetsArray and PermissionTargetsArrayOutput values.
// You can construct a concrete instance of `PermissionTargetsArrayInput` via:
//
//          PermissionTargetsArray{ PermissionTargetsArgs{...} }
type PermissionTargetsArrayInput interface {
	pulumi.Input

	ToPermissionTargetsArrayOutput() PermissionTargetsArrayOutput
	ToPermissionTargetsArrayOutputWithContext(context.Context) PermissionTargetsArrayOutput
}

type PermissionTargetsArray []PermissionTargetsInput

func (PermissionTargetsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PermissionTargets)(nil)).Elem()
}

func (i PermissionTargetsArray) ToPermissionTargetsArrayOutput() PermissionTargetsArrayOutput {
	return i.ToPermissionTargetsArrayOutputWithContext(context.Background())
}

func (i PermissionTargetsArray) ToPermissionTargetsArrayOutputWithContext(ctx context.Context) PermissionTargetsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionTargetsArrayOutput)
}

// PermissionTargetsMapInput is an input type that accepts PermissionTargetsMap and PermissionTargetsMapOutput values.
// You can construct a concrete instance of `PermissionTargetsMapInput` via:
//
//          PermissionTargetsMap{ "key": PermissionTargetsArgs{...} }
type PermissionTargetsMapInput interface {
	pulumi.Input

	ToPermissionTargetsMapOutput() PermissionTargetsMapOutput
	ToPermissionTargetsMapOutputWithContext(context.Context) PermissionTargetsMapOutput
}

type PermissionTargetsMap map[string]PermissionTargetsInput

func (PermissionTargetsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PermissionTargets)(nil)).Elem()
}

func (i PermissionTargetsMap) ToPermissionTargetsMapOutput() PermissionTargetsMapOutput {
	return i.ToPermissionTargetsMapOutputWithContext(context.Background())
}

func (i PermissionTargetsMap) ToPermissionTargetsMapOutputWithContext(ctx context.Context) PermissionTargetsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionTargetsMapOutput)
}

type PermissionTargetsOutput struct{ *pulumi.OutputState }

func (PermissionTargetsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PermissionTargets)(nil)).Elem()
}

func (o PermissionTargetsOutput) ToPermissionTargetsOutput() PermissionTargetsOutput {
	return o
}

func (o PermissionTargetsOutput) ToPermissionTargetsOutputWithContext(ctx context.Context) PermissionTargetsOutput {
	return o
}

type PermissionTargetsArrayOutput struct{ *pulumi.OutputState }

func (PermissionTargetsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PermissionTargets)(nil)).Elem()
}

func (o PermissionTargetsArrayOutput) ToPermissionTargetsArrayOutput() PermissionTargetsArrayOutput {
	return o
}

func (o PermissionTargetsArrayOutput) ToPermissionTargetsArrayOutputWithContext(ctx context.Context) PermissionTargetsArrayOutput {
	return o
}

func (o PermissionTargetsArrayOutput) Index(i pulumi.IntInput) PermissionTargetsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PermissionTargets {
		return vs[0].([]*PermissionTargets)[vs[1].(int)]
	}).(PermissionTargetsOutput)
}

type PermissionTargetsMapOutput struct{ *pulumi.OutputState }

func (PermissionTargetsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PermissionTargets)(nil)).Elem()
}

func (o PermissionTargetsMapOutput) ToPermissionTargetsMapOutput() PermissionTargetsMapOutput {
	return o
}

func (o PermissionTargetsMapOutput) ToPermissionTargetsMapOutputWithContext(ctx context.Context) PermissionTargetsMapOutput {
	return o
}

func (o PermissionTargetsMapOutput) MapIndex(k pulumi.StringInput) PermissionTargetsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PermissionTargets {
		return vs[0].(map[string]*PermissionTargets)[vs[1].(string)]
	}).(PermissionTargetsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionTargetsInput)(nil)).Elem(), &PermissionTargets{})
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionTargetsArrayInput)(nil)).Elem(), PermissionTargetsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionTargetsMapInput)(nil)).Elem(), PermissionTargetsMap{})
	pulumi.RegisterOutputType(PermissionTargetsOutput{})
	pulumi.RegisterOutputType(PermissionTargetsArrayOutput{})
	pulumi.RegisterOutputType(PermissionTargetsMapOutput{})
}
