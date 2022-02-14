// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Xray Watch Resource
//
// Provides a Xray watch resource.
//
// ## Import
//
// Watches can be imported using their name, e.g.
//
// ```sh
//  $ pulumi import artifactory:index/xrayWatch:XrayWatch example watch-name
// ```
type XrayWatch struct {
	pulumi.CustomResourceState

	// Whether or not the watch will be active
	Active pulumi.BoolPtrOutput `pulumi:"active"`
	// Nested argument describing policies that will be applied. Defined below.
	AssignedPolicies XrayWatchAssignedPolicyArrayOutput `pulumi:"assignedPolicies"`
	// Description of the watch
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the watch (must be unique)
	Name pulumi.StringOutput `pulumi:"name"`
	// Nested argument describing the resources to be watched. Defined below.
	Resources       XrayWatchResourceArrayOutput `pulumi:"resources"`
	WatchRecipients pulumi.StringArrayOutput     `pulumi:"watchRecipients"`
}

// NewXrayWatch registers a new resource with the given unique name, arguments, and options.
func NewXrayWatch(ctx *pulumi.Context,
	name string, args *XrayWatchArgs, opts ...pulumi.ResourceOption) (*XrayWatch, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AssignedPolicies == nil {
		return nil, errors.New("invalid value for required argument 'AssignedPolicies'")
	}
	if args.Resources == nil {
		return nil, errors.New("invalid value for required argument 'Resources'")
	}
	var resource XrayWatch
	err := ctx.RegisterResource("artifactory:index/xrayWatch:XrayWatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetXrayWatch gets an existing XrayWatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetXrayWatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *XrayWatchState, opts ...pulumi.ResourceOption) (*XrayWatch, error) {
	var resource XrayWatch
	err := ctx.ReadResource("artifactory:index/xrayWatch:XrayWatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering XrayWatch resources.
type xrayWatchState struct {
	// Whether or not the watch will be active
	Active *bool `pulumi:"active"`
	// Nested argument describing policies that will be applied. Defined below.
	AssignedPolicies []XrayWatchAssignedPolicy `pulumi:"assignedPolicies"`
	// Description of the watch
	Description *string `pulumi:"description"`
	// Name of the watch (must be unique)
	Name *string `pulumi:"name"`
	// Nested argument describing the resources to be watched. Defined below.
	Resources       []XrayWatchResource `pulumi:"resources"`
	WatchRecipients []string            `pulumi:"watchRecipients"`
}

type XrayWatchState struct {
	// Whether or not the watch will be active
	Active pulumi.BoolPtrInput
	// Nested argument describing policies that will be applied. Defined below.
	AssignedPolicies XrayWatchAssignedPolicyArrayInput
	// Description of the watch
	Description pulumi.StringPtrInput
	// Name of the watch (must be unique)
	Name pulumi.StringPtrInput
	// Nested argument describing the resources to be watched. Defined below.
	Resources       XrayWatchResourceArrayInput
	WatchRecipients pulumi.StringArrayInput
}

func (XrayWatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*xrayWatchState)(nil)).Elem()
}

type xrayWatchArgs struct {
	// Whether or not the watch will be active
	Active *bool `pulumi:"active"`
	// Nested argument describing policies that will be applied. Defined below.
	AssignedPolicies []XrayWatchAssignedPolicy `pulumi:"assignedPolicies"`
	// Description of the watch
	Description *string `pulumi:"description"`
	// Name of the watch (must be unique)
	Name *string `pulumi:"name"`
	// Nested argument describing the resources to be watched. Defined below.
	Resources       []XrayWatchResource `pulumi:"resources"`
	WatchRecipients []string            `pulumi:"watchRecipients"`
}

// The set of arguments for constructing a XrayWatch resource.
type XrayWatchArgs struct {
	// Whether or not the watch will be active
	Active pulumi.BoolPtrInput
	// Nested argument describing policies that will be applied. Defined below.
	AssignedPolicies XrayWatchAssignedPolicyArrayInput
	// Description of the watch
	Description pulumi.StringPtrInput
	// Name of the watch (must be unique)
	Name pulumi.StringPtrInput
	// Nested argument describing the resources to be watched. Defined below.
	Resources       XrayWatchResourceArrayInput
	WatchRecipients pulumi.StringArrayInput
}

func (XrayWatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*xrayWatchArgs)(nil)).Elem()
}

type XrayWatchInput interface {
	pulumi.Input

	ToXrayWatchOutput() XrayWatchOutput
	ToXrayWatchOutputWithContext(ctx context.Context) XrayWatchOutput
}

func (*XrayWatch) ElementType() reflect.Type {
	return reflect.TypeOf((**XrayWatch)(nil)).Elem()
}

func (i *XrayWatch) ToXrayWatchOutput() XrayWatchOutput {
	return i.ToXrayWatchOutputWithContext(context.Background())
}

func (i *XrayWatch) ToXrayWatchOutputWithContext(ctx context.Context) XrayWatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(XrayWatchOutput)
}

// XrayWatchArrayInput is an input type that accepts XrayWatchArray and XrayWatchArrayOutput values.
// You can construct a concrete instance of `XrayWatchArrayInput` via:
//
//          XrayWatchArray{ XrayWatchArgs{...} }
type XrayWatchArrayInput interface {
	pulumi.Input

	ToXrayWatchArrayOutput() XrayWatchArrayOutput
	ToXrayWatchArrayOutputWithContext(context.Context) XrayWatchArrayOutput
}

type XrayWatchArray []XrayWatchInput

func (XrayWatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*XrayWatch)(nil)).Elem()
}

func (i XrayWatchArray) ToXrayWatchArrayOutput() XrayWatchArrayOutput {
	return i.ToXrayWatchArrayOutputWithContext(context.Background())
}

func (i XrayWatchArray) ToXrayWatchArrayOutputWithContext(ctx context.Context) XrayWatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(XrayWatchArrayOutput)
}

// XrayWatchMapInput is an input type that accepts XrayWatchMap and XrayWatchMapOutput values.
// You can construct a concrete instance of `XrayWatchMapInput` via:
//
//          XrayWatchMap{ "key": XrayWatchArgs{...} }
type XrayWatchMapInput interface {
	pulumi.Input

	ToXrayWatchMapOutput() XrayWatchMapOutput
	ToXrayWatchMapOutputWithContext(context.Context) XrayWatchMapOutput
}

type XrayWatchMap map[string]XrayWatchInput

func (XrayWatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*XrayWatch)(nil)).Elem()
}

func (i XrayWatchMap) ToXrayWatchMapOutput() XrayWatchMapOutput {
	return i.ToXrayWatchMapOutputWithContext(context.Background())
}

func (i XrayWatchMap) ToXrayWatchMapOutputWithContext(ctx context.Context) XrayWatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(XrayWatchMapOutput)
}

type XrayWatchOutput struct{ *pulumi.OutputState }

func (XrayWatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**XrayWatch)(nil)).Elem()
}

func (o XrayWatchOutput) ToXrayWatchOutput() XrayWatchOutput {
	return o
}

func (o XrayWatchOutput) ToXrayWatchOutputWithContext(ctx context.Context) XrayWatchOutput {
	return o
}

type XrayWatchArrayOutput struct{ *pulumi.OutputState }

func (XrayWatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*XrayWatch)(nil)).Elem()
}

func (o XrayWatchArrayOutput) ToXrayWatchArrayOutput() XrayWatchArrayOutput {
	return o
}

func (o XrayWatchArrayOutput) ToXrayWatchArrayOutputWithContext(ctx context.Context) XrayWatchArrayOutput {
	return o
}

func (o XrayWatchArrayOutput) Index(i pulumi.IntInput) XrayWatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *XrayWatch {
		return vs[0].([]*XrayWatch)[vs[1].(int)]
	}).(XrayWatchOutput)
}

type XrayWatchMapOutput struct{ *pulumi.OutputState }

func (XrayWatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*XrayWatch)(nil)).Elem()
}

func (o XrayWatchMapOutput) ToXrayWatchMapOutput() XrayWatchMapOutput {
	return o
}

func (o XrayWatchMapOutput) ToXrayWatchMapOutputWithContext(ctx context.Context) XrayWatchMapOutput {
	return o
}

func (o XrayWatchMapOutput) MapIndex(k pulumi.StringInput) XrayWatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *XrayWatch {
		return vs[0].(map[string]*XrayWatch)[vs[1].(string)]
	}).(XrayWatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*XrayWatchInput)(nil)).Elem(), &XrayWatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*XrayWatchArrayInput)(nil)).Elem(), XrayWatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*XrayWatchMapInput)(nil)).Elem(), XrayWatchMap{})
	pulumi.RegisterOutputType(XrayWatchOutput{})
	pulumi.RegisterOutputType(XrayWatchArrayOutput{})
	pulumi.RegisterOutputType(XrayWatchMapOutput{})
}
