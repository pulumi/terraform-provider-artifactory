// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Groups can be imported using their name, e.g.
//
// ```sh
//  $ pulumi import artifactory:index/group:Group terraform-group mygroup
// ```
type Group struct {
	pulumi.CustomResourceState

	// Any users added to this group will automatically be assigned with admin privileges in the system.
	AdminPrivileges pulumi.BoolOutput `pulumi:"adminPrivileges"`
	// When this parameter is set, any new users defined in the system are automatically assigned to this group.
	AutoJoin pulumi.BoolOutput `pulumi:"autoJoin"`
	// A description for the group
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// When this override is set, an empty or missing usernames array will detach all users from the group
	DetachAllUsers pulumi.BoolPtrOutput `pulumi:"detachAllUsers"`
	// Name of the group
	Name pulumi.StringOutput `pulumi:"name"`
	// When this override is set, User in the group can set Xray security and compliance policies. Default value is 'false'.
	PolicyManager pulumi.BoolPtrOutput `pulumi:"policyManager"`
	// The realm for the group.
	Realm pulumi.StringOutput `pulumi:"realm"`
	// The realm attributes for the group.
	RealmAttributes pulumi.StringPtrOutput `pulumi:"realmAttributes"`
	// When this override is set, User in the group can manage Xray Reports on any resource type. Default value is 'false'.
	ReportsManager pulumi.BoolPtrOutput `pulumi:"reportsManager"`
	// List of users assigned to the group. If missing or empty, tf will not manage group membership
	UsersNames pulumi.StringArrayOutput `pulumi:"usersNames"`
	// When this override is set, User in the group can manage Xray Watches on any resource type. Default value is 'false'.
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
	// Any users added to this group will automatically be assigned with admin privileges in the system.
	AdminPrivileges *bool `pulumi:"adminPrivileges"`
	// When this parameter is set, any new users defined in the system are automatically assigned to this group.
	AutoJoin *bool `pulumi:"autoJoin"`
	// A description for the group
	Description *string `pulumi:"description"`
	// When this override is set, an empty or missing usernames array will detach all users from the group
	DetachAllUsers *bool `pulumi:"detachAllUsers"`
	// Name of the group
	Name *string `pulumi:"name"`
	// When this override is set, User in the group can set Xray security and compliance policies. Default value is 'false'.
	PolicyManager *bool `pulumi:"policyManager"`
	// The realm for the group.
	Realm *string `pulumi:"realm"`
	// The realm attributes for the group.
	RealmAttributes *string `pulumi:"realmAttributes"`
	// When this override is set, User in the group can manage Xray Reports on any resource type. Default value is 'false'.
	ReportsManager *bool `pulumi:"reportsManager"`
	// List of users assigned to the group. If missing or empty, tf will not manage group membership
	UsersNames []string `pulumi:"usersNames"`
	// When this override is set, User in the group can manage Xray Watches on any resource type. Default value is 'false'.
	WatchManager *bool `pulumi:"watchManager"`
}

type GroupState struct {
	// Any users added to this group will automatically be assigned with admin privileges in the system.
	AdminPrivileges pulumi.BoolPtrInput
	// When this parameter is set, any new users defined in the system are automatically assigned to this group.
	AutoJoin pulumi.BoolPtrInput
	// A description for the group
	Description pulumi.StringPtrInput
	// When this override is set, an empty or missing usernames array will detach all users from the group
	DetachAllUsers pulumi.BoolPtrInput
	// Name of the group
	Name pulumi.StringPtrInput
	// When this override is set, User in the group can set Xray security and compliance policies. Default value is 'false'.
	PolicyManager pulumi.BoolPtrInput
	// The realm for the group.
	Realm pulumi.StringPtrInput
	// The realm attributes for the group.
	RealmAttributes pulumi.StringPtrInput
	// When this override is set, User in the group can manage Xray Reports on any resource type. Default value is 'false'.
	ReportsManager pulumi.BoolPtrInput
	// List of users assigned to the group. If missing or empty, tf will not manage group membership
	UsersNames pulumi.StringArrayInput
	// When this override is set, User in the group can manage Xray Watches on any resource type. Default value is 'false'.
	WatchManager pulumi.BoolPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// Any users added to this group will automatically be assigned with admin privileges in the system.
	AdminPrivileges *bool `pulumi:"adminPrivileges"`
	// When this parameter is set, any new users defined in the system are automatically assigned to this group.
	AutoJoin *bool `pulumi:"autoJoin"`
	// A description for the group
	Description *string `pulumi:"description"`
	// When this override is set, an empty or missing usernames array will detach all users from the group
	DetachAllUsers *bool `pulumi:"detachAllUsers"`
	// Name of the group
	Name *string `pulumi:"name"`
	// When this override is set, User in the group can set Xray security and compliance policies. Default value is 'false'.
	PolicyManager *bool `pulumi:"policyManager"`
	// The realm for the group.
	Realm *string `pulumi:"realm"`
	// The realm attributes for the group.
	RealmAttributes *string `pulumi:"realmAttributes"`
	// When this override is set, User in the group can manage Xray Reports on any resource type. Default value is 'false'.
	ReportsManager *bool `pulumi:"reportsManager"`
	// List of users assigned to the group. If missing or empty, tf will not manage group membership
	UsersNames []string `pulumi:"usersNames"`
	// When this override is set, User in the group can manage Xray Watches on any resource type. Default value is 'false'.
	WatchManager *bool `pulumi:"watchManager"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// Any users added to this group will automatically be assigned with admin privileges in the system.
	AdminPrivileges pulumi.BoolPtrInput
	// When this parameter is set, any new users defined in the system are automatically assigned to this group.
	AutoJoin pulumi.BoolPtrInput
	// A description for the group
	Description pulumi.StringPtrInput
	// When this override is set, an empty or missing usernames array will detach all users from the group
	DetachAllUsers pulumi.BoolPtrInput
	// Name of the group
	Name pulumi.StringPtrInput
	// When this override is set, User in the group can set Xray security and compliance policies. Default value is 'false'.
	PolicyManager pulumi.BoolPtrInput
	// The realm for the group.
	Realm pulumi.StringPtrInput
	// The realm attributes for the group.
	RealmAttributes pulumi.StringPtrInput
	// When this override is set, User in the group can manage Xray Reports on any resource type. Default value is 'false'.
	ReportsManager pulumi.BoolPtrInput
	// List of users assigned to the group. If missing or empty, tf will not manage group membership
	UsersNames pulumi.StringArrayInput
	// When this override is set, User in the group can manage Xray Watches on any resource type. Default value is 'false'.
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
