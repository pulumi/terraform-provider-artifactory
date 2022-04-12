// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Group struct {
	pulumi.CustomResourceState

	AdminPrivileges pulumi.BoolOutput      `pulumi:"adminPrivileges"`
	AutoJoin        pulumi.BoolOutput      `pulumi:"autoJoin"`
	Description     pulumi.StringPtrOutput `pulumi:"description"`
	DetachAllUsers  pulumi.BoolPtrOutput   `pulumi:"detachAllUsers"`
	Name            pulumi.StringOutput    `pulumi:"name"`
	// (Optional) When this override is set, User in the group can set Xray security and compliance policies. Default value is
	// 'false'.
	PolicyManager   pulumi.BoolPtrOutput   `pulumi:"policyManager"`
	Realm           pulumi.StringOutput    `pulumi:"realm"`
	RealmAttributes pulumi.StringPtrOutput `pulumi:"realmAttributes"`
	// (Optional) When this override is set, User in the group can manage Xray Reports. Default value is 'false'.
	ReportsManager pulumi.BoolPtrOutput     `pulumi:"reportsManager"`
	UsersNames     pulumi.StringArrayOutput `pulumi:"usersNames"`
	// (Optional) When this override is set, User in the group can manage Xray Watches on any resource type. Default value is
	// 'false'.
	WatchManager pulumi.BoolPtrOutput `pulumi:"watchManager"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		args = &GroupArgs{}
	}

	var resource Group
	err := ctx.RegisterResource("artifactory:index/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("artifactory:index/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	AdminPrivileges *bool   `pulumi:"adminPrivileges"`
	AutoJoin        *bool   `pulumi:"autoJoin"`
	Description     *string `pulumi:"description"`
	DetachAllUsers  *bool   `pulumi:"detachAllUsers"`
	Name            *string `pulumi:"name"`
	// (Optional) When this override is set, User in the group can set Xray security and compliance policies. Default value is
	// 'false'.
	PolicyManager   *bool   `pulumi:"policyManager"`
	Realm           *string `pulumi:"realm"`
	RealmAttributes *string `pulumi:"realmAttributes"`
	// (Optional) When this override is set, User in the group can manage Xray Reports. Default value is 'false'.
	ReportsManager *bool    `pulumi:"reportsManager"`
	UsersNames     []string `pulumi:"usersNames"`
	// (Optional) When this override is set, User in the group can manage Xray Watches on any resource type. Default value is
	// 'false'.
	WatchManager *bool `pulumi:"watchManager"`
}

type GroupState struct {
	AdminPrivileges pulumi.BoolPtrInput
	AutoJoin        pulumi.BoolPtrInput
	Description     pulumi.StringPtrInput
	DetachAllUsers  pulumi.BoolPtrInput
	Name            pulumi.StringPtrInput
	// (Optional) When this override is set, User in the group can set Xray security and compliance policies. Default value is
	// 'false'.
	PolicyManager   pulumi.BoolPtrInput
	Realm           pulumi.StringPtrInput
	RealmAttributes pulumi.StringPtrInput
	// (Optional) When this override is set, User in the group can manage Xray Reports. Default value is 'false'.
	ReportsManager pulumi.BoolPtrInput
	UsersNames     pulumi.StringArrayInput
	// (Optional) When this override is set, User in the group can manage Xray Watches on any resource type. Default value is
	// 'false'.
	WatchManager pulumi.BoolPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	AdminPrivileges *bool   `pulumi:"adminPrivileges"`
	AutoJoin        *bool   `pulumi:"autoJoin"`
	Description     *string `pulumi:"description"`
	DetachAllUsers  *bool   `pulumi:"detachAllUsers"`
	Name            *string `pulumi:"name"`
	// (Optional) When this override is set, User in the group can set Xray security and compliance policies. Default value is
	// 'false'.
	PolicyManager   *bool   `pulumi:"policyManager"`
	Realm           *string `pulumi:"realm"`
	RealmAttributes *string `pulumi:"realmAttributes"`
	// (Optional) When this override is set, User in the group can manage Xray Reports. Default value is 'false'.
	ReportsManager *bool    `pulumi:"reportsManager"`
	UsersNames     []string `pulumi:"usersNames"`
	// (Optional) When this override is set, User in the group can manage Xray Watches on any resource type. Default value is
	// 'false'.
	WatchManager *bool `pulumi:"watchManager"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	AdminPrivileges pulumi.BoolPtrInput
	AutoJoin        pulumi.BoolPtrInput
	Description     pulumi.StringPtrInput
	DetachAllUsers  pulumi.BoolPtrInput
	Name            pulumi.StringPtrInput
	// (Optional) When this override is set, User in the group can set Xray security and compliance policies. Default value is
	// 'false'.
	PolicyManager   pulumi.BoolPtrInput
	Realm           pulumi.StringPtrInput
	RealmAttributes pulumi.StringPtrInput
	// (Optional) When this override is set, User in the group can manage Xray Reports. Default value is 'false'.
	ReportsManager pulumi.BoolPtrInput
	UsersNames     pulumi.StringArrayInput
	// (Optional) When this override is set, User in the group can manage Xray Watches on any resource type. Default value is
	// 'false'.
	WatchManager pulumi.BoolPtrInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

// GroupArrayInput is an input type that accepts GroupArray and GroupArrayOutput values.
// You can construct a concrete instance of `GroupArrayInput` via:
//
//          GroupArray{ GroupArgs{...} }
type GroupArrayInput interface {
	pulumi.Input

	ToGroupArrayOutput() GroupArrayOutput
	ToGroupArrayOutputWithContext(context.Context) GroupArrayOutput
}

type GroupArray []GroupInput

func (GroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (i GroupArray) ToGroupArrayOutput() GroupArrayOutput {
	return i.ToGroupArrayOutputWithContext(context.Background())
}

func (i GroupArray) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupArrayOutput)
}

// GroupMapInput is an input type that accepts GroupMap and GroupMapOutput values.
// You can construct a concrete instance of `GroupMapInput` via:
//
//          GroupMap{ "key": GroupArgs{...} }
type GroupMapInput interface {
	pulumi.Input

	ToGroupMapOutput() GroupMapOutput
	ToGroupMapOutputWithContext(context.Context) GroupMapOutput
}

type GroupMap map[string]GroupInput

func (GroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (i GroupMap) ToGroupMapOutput() GroupMapOutput {
	return i.ToGroupMapOutputWithContext(context.Background())
}

func (i GroupMap) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMapOutput)
}

type GroupOutput struct{ *pulumi.OutputState }

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

type GroupArrayOutput struct{ *pulumi.OutputState }

func (GroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (o GroupArrayOutput) ToGroupArrayOutput() GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) Index(i pulumi.IntInput) GroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Group {
		return vs[0].([]*Group)[vs[1].(int)]
	}).(GroupOutput)
}

type GroupMapOutput struct{ *pulumi.OutputState }

func (GroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (o GroupMapOutput) ToGroupMapOutput() GroupMapOutput {
	return o
}

func (o GroupMapOutput) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return o
}

func (o GroupMapOutput) MapIndex(k pulumi.StringInput) GroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Group {
		return vs[0].(map[string]*Group)[vs[1].(string)]
	}).(GroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupInput)(nil)).Elem(), &Group{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupArrayInput)(nil)).Elem(), GroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMapInput)(nil)).Elem(), GroupMap{})
	pulumi.RegisterOutputType(GroupOutput{})
	pulumi.RegisterOutputType(GroupArrayOutput{})
	pulumi.RegisterOutputType(GroupMapOutput{})
}
