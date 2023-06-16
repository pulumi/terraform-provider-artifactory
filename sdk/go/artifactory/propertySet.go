// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PropertySet struct {
	pulumi.CustomResourceState

	// Property set name.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of properties that will be part of the property set.
	Properties PropertySetPropertyArrayOutput `pulumi:"properties"`
	// Defines if the list visible and assignable to the repository or artifact.
	Visible pulumi.BoolPtrOutput `pulumi:"visible"`
}

// NewPropertySet registers a new resource with the given unique name, arguments, and options.
func NewPropertySet(ctx *pulumi.Context,
	name string, args *PropertySetArgs, opts ...pulumi.ResourceOption) (*PropertySet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource PropertySet
	err := ctx.RegisterResource("artifactory:index/propertySet:PropertySet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPropertySet gets an existing PropertySet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPropertySet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PropertySetState, opts ...pulumi.ResourceOption) (*PropertySet, error) {
	var resource PropertySet
	err := ctx.ReadResource("artifactory:index/propertySet:PropertySet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PropertySet resources.
type propertySetState struct {
	// Property set name.
	Name *string `pulumi:"name"`
	// A list of properties that will be part of the property set.
	Properties []PropertySetProperty `pulumi:"properties"`
	// Defines if the list visible and assignable to the repository or artifact.
	Visible *bool `pulumi:"visible"`
}

type PropertySetState struct {
	// Property set name.
	Name pulumi.StringPtrInput
	// A list of properties that will be part of the property set.
	Properties PropertySetPropertyArrayInput
	// Defines if the list visible and assignable to the repository or artifact.
	Visible pulumi.BoolPtrInput
}

func (PropertySetState) ElementType() reflect.Type {
	return reflect.TypeOf((*propertySetState)(nil)).Elem()
}

type propertySetArgs struct {
	// Property set name.
	Name *string `pulumi:"name"`
	// A list of properties that will be part of the property set.
	Properties []PropertySetProperty `pulumi:"properties"`
	// Defines if the list visible and assignable to the repository or artifact.
	Visible *bool `pulumi:"visible"`
}

// The set of arguments for constructing a PropertySet resource.
type PropertySetArgs struct {
	// Property set name.
	Name pulumi.StringPtrInput
	// A list of properties that will be part of the property set.
	Properties PropertySetPropertyArrayInput
	// Defines if the list visible and assignable to the repository or artifact.
	Visible pulumi.BoolPtrInput
}

func (PropertySetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*propertySetArgs)(nil)).Elem()
}

type PropertySetInput interface {
	pulumi.Input

	ToPropertySetOutput() PropertySetOutput
	ToPropertySetOutputWithContext(ctx context.Context) PropertySetOutput
}

func (*PropertySet) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertySet)(nil)).Elem()
}

func (i *PropertySet) ToPropertySetOutput() PropertySetOutput {
	return i.ToPropertySetOutputWithContext(context.Background())
}

func (i *PropertySet) ToPropertySetOutputWithContext(ctx context.Context) PropertySetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertySetOutput)
}

// PropertySetArrayInput is an input type that accepts PropertySetArray and PropertySetArrayOutput values.
// You can construct a concrete instance of `PropertySetArrayInput` via:
//
//	PropertySetArray{ PropertySetArgs{...} }
type PropertySetArrayInput interface {
	pulumi.Input

	ToPropertySetArrayOutput() PropertySetArrayOutput
	ToPropertySetArrayOutputWithContext(context.Context) PropertySetArrayOutput
}

type PropertySetArray []PropertySetInput

func (PropertySetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PropertySet)(nil)).Elem()
}

func (i PropertySetArray) ToPropertySetArrayOutput() PropertySetArrayOutput {
	return i.ToPropertySetArrayOutputWithContext(context.Background())
}

func (i PropertySetArray) ToPropertySetArrayOutputWithContext(ctx context.Context) PropertySetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertySetArrayOutput)
}

// PropertySetMapInput is an input type that accepts PropertySetMap and PropertySetMapOutput values.
// You can construct a concrete instance of `PropertySetMapInput` via:
//
//	PropertySetMap{ "key": PropertySetArgs{...} }
type PropertySetMapInput interface {
	pulumi.Input

	ToPropertySetMapOutput() PropertySetMapOutput
	ToPropertySetMapOutputWithContext(context.Context) PropertySetMapOutput
}

type PropertySetMap map[string]PropertySetInput

func (PropertySetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PropertySet)(nil)).Elem()
}

func (i PropertySetMap) ToPropertySetMapOutput() PropertySetMapOutput {
	return i.ToPropertySetMapOutputWithContext(context.Background())
}

func (i PropertySetMap) ToPropertySetMapOutputWithContext(ctx context.Context) PropertySetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertySetMapOutput)
}

type PropertySetOutput struct{ *pulumi.OutputState }

func (PropertySetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertySet)(nil)).Elem()
}

func (o PropertySetOutput) ToPropertySetOutput() PropertySetOutput {
	return o
}

func (o PropertySetOutput) ToPropertySetOutputWithContext(ctx context.Context) PropertySetOutput {
	return o
}

// Property set name.
func (o PropertySetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PropertySet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A list of properties that will be part of the property set.
func (o PropertySetOutput) Properties() PropertySetPropertyArrayOutput {
	return o.ApplyT(func(v *PropertySet) PropertySetPropertyArrayOutput { return v.Properties }).(PropertySetPropertyArrayOutput)
}

// Defines if the list visible and assignable to the repository or artifact.
func (o PropertySetOutput) Visible() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PropertySet) pulumi.BoolPtrOutput { return v.Visible }).(pulumi.BoolPtrOutput)
}

type PropertySetArrayOutput struct{ *pulumi.OutputState }

func (PropertySetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PropertySet)(nil)).Elem()
}

func (o PropertySetArrayOutput) ToPropertySetArrayOutput() PropertySetArrayOutput {
	return o
}

func (o PropertySetArrayOutput) ToPropertySetArrayOutputWithContext(ctx context.Context) PropertySetArrayOutput {
	return o
}

func (o PropertySetArrayOutput) Index(i pulumi.IntInput) PropertySetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PropertySet {
		return vs[0].([]*PropertySet)[vs[1].(int)]
	}).(PropertySetOutput)
}

type PropertySetMapOutput struct{ *pulumi.OutputState }

func (PropertySetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PropertySet)(nil)).Elem()
}

func (o PropertySetMapOutput) ToPropertySetMapOutput() PropertySetMapOutput {
	return o
}

func (o PropertySetMapOutput) ToPropertySetMapOutputWithContext(ctx context.Context) PropertySetMapOutput {
	return o
}

func (o PropertySetMapOutput) MapIndex(k pulumi.StringInput) PropertySetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PropertySet {
		return vs[0].(map[string]*PropertySet)[vs[1].(string)]
	}).(PropertySetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PropertySetInput)(nil)).Elem(), &PropertySet{})
	pulumi.RegisterInputType(reflect.TypeOf((*PropertySetArrayInput)(nil)).Elem(), PropertySetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PropertySetMapInput)(nil)).Elem(), PropertySetMap{})
	pulumi.RegisterOutputType(PropertySetOutput{})
	pulumi.RegisterOutputType(PropertySetArrayOutput{})
	pulumi.RegisterOutputType(PropertySetMapOutput{})
}
