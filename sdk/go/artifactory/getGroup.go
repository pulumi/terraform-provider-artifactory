// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v6/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Group Data Source
//
// Provides an Artifactory group datasource. This can be used to read the configuration of groups in artifactory.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v6/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.LookupGroup(ctx, &artifactory.LookupGroupArgs{
//				Name:         "my_group",
//				IncludeUsers: pulumi.StringRef("true"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGroupResult
	err := ctx.Invoke("artifactory:index/getGroup:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroup.
type LookupGroupArgs struct {
	// Any users added to this group will automatically be assigned with admin privileges in the system.
	AdminPrivileges *bool `pulumi:"adminPrivileges"`
	// When this parameter is set, any new users defined in the system are automatically assigned to this group.
	AutoJoin *bool `pulumi:"autoJoin"`
	// A description for the group
	Description *string `pulumi:"description"`
	// New external group ID used to configure the corresponding group in Azure AD.
	ExternalId *string `pulumi:"externalId"`
	// Determines if the group's associated user list will return as an attribute. Default is `false`.
	IncludeUsers *string `pulumi:"includeUsers"`
	// Name of the group.
	Name string `pulumi:"name"`
	// When this override is set, User in the group can set Xray security and compliance policies. Default value is `false`.
	PolicyManager *bool `pulumi:"policyManager"`
	// The realm for the group.
	Realm *string `pulumi:"realm"`
	// The realm attributes for the group.
	RealmAttributes *string `pulumi:"realmAttributes"`
	// When this override is set, User in the group can manage Xray Reports on any resource type. Default value is `false`.
	ReportsManager *bool `pulumi:"reportsManager"`
	// List of users assigned to the group. Set includeUsers to `true` to retrieve this list.
	UsersNames []string `pulumi:"usersNames"`
	// When this override is set, User in the group can manage Xray Watches on any resource type. Default value is `false`.
	WatchManager *bool `pulumi:"watchManager"`
}

// A collection of values returned by getGroup.
type LookupGroupResult struct {
	// Any users added to this group will automatically be assigned with admin privileges in the system.
	AdminPrivileges bool `pulumi:"adminPrivileges"`
	// When this parameter is set, any new users defined in the system are automatically assigned to this group.
	AutoJoin bool `pulumi:"autoJoin"`
	// A description for the group
	Description *string `pulumi:"description"`
	// New external group ID used to configure the corresponding group in Azure AD.
	ExternalId *string `pulumi:"externalId"`
	// The provider-assigned unique ID for this managed resource.
	Id           string  `pulumi:"id"`
	IncludeUsers *string `pulumi:"includeUsers"`
	Name         string  `pulumi:"name"`
	// When this override is set, User in the group can set Xray security and compliance policies. Default value is `false`.
	PolicyManager *bool `pulumi:"policyManager"`
	// The realm for the group.
	Realm string `pulumi:"realm"`
	// The realm attributes for the group.
	RealmAttributes *string `pulumi:"realmAttributes"`
	// When this override is set, User in the group can manage Xray Reports on any resource type. Default value is `false`.
	ReportsManager *bool `pulumi:"reportsManager"`
	// List of users assigned to the group. Set includeUsers to `true` to retrieve this list.
	UsersNames []string `pulumi:"usersNames"`
	// When this override is set, User in the group can manage Xray Watches on any resource type. Default value is `false`.
	WatchManager *bool `pulumi:"watchManager"`
}

func LookupGroupOutput(ctx *pulumi.Context, args LookupGroupOutputArgs, opts ...pulumi.InvokeOption) LookupGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGroupResult, error) {
			args := v.(LookupGroupArgs)
			r, err := LookupGroup(ctx, &args, opts...)
			var s LookupGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGroupResultOutput)
}

// A collection of arguments for invoking getGroup.
type LookupGroupOutputArgs struct {
	// Any users added to this group will automatically be assigned with admin privileges in the system.
	AdminPrivileges pulumi.BoolPtrInput `pulumi:"adminPrivileges"`
	// When this parameter is set, any new users defined in the system are automatically assigned to this group.
	AutoJoin pulumi.BoolPtrInput `pulumi:"autoJoin"`
	// A description for the group
	Description pulumi.StringPtrInput `pulumi:"description"`
	// New external group ID used to configure the corresponding group in Azure AD.
	ExternalId pulumi.StringPtrInput `pulumi:"externalId"`
	// Determines if the group's associated user list will return as an attribute. Default is `false`.
	IncludeUsers pulumi.StringPtrInput `pulumi:"includeUsers"`
	// Name of the group.
	Name pulumi.StringInput `pulumi:"name"`
	// When this override is set, User in the group can set Xray security and compliance policies. Default value is `false`.
	PolicyManager pulumi.BoolPtrInput `pulumi:"policyManager"`
	// The realm for the group.
	Realm pulumi.StringPtrInput `pulumi:"realm"`
	// The realm attributes for the group.
	RealmAttributes pulumi.StringPtrInput `pulumi:"realmAttributes"`
	// When this override is set, User in the group can manage Xray Reports on any resource type. Default value is `false`.
	ReportsManager pulumi.BoolPtrInput `pulumi:"reportsManager"`
	// List of users assigned to the group. Set includeUsers to `true` to retrieve this list.
	UsersNames pulumi.StringArrayInput `pulumi:"usersNames"`
	// When this override is set, User in the group can manage Xray Watches on any resource type. Default value is `false`.
	WatchManager pulumi.BoolPtrInput `pulumi:"watchManager"`
}

func (LookupGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupArgs)(nil)).Elem()
}

// A collection of values returned by getGroup.
type LookupGroupResultOutput struct{ *pulumi.OutputState }

func (LookupGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupResult)(nil)).Elem()
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutput() LookupGroupResultOutput {
	return o
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutputWithContext(ctx context.Context) LookupGroupResultOutput {
	return o
}

// Any users added to this group will automatically be assigned with admin privileges in the system.
func (o LookupGroupResultOutput) AdminPrivileges() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.AdminPrivileges }).(pulumi.BoolOutput)
}

// When this parameter is set, any new users defined in the system are automatically assigned to this group.
func (o LookupGroupResultOutput) AutoJoin() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.AutoJoin }).(pulumi.BoolOutput)
}

// A description for the group
func (o LookupGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// New external group ID used to configure the corresponding group in Azure AD.
func (o LookupGroupResultOutput) ExternalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *string { return v.ExternalId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGroupResultOutput) IncludeUsers() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *string { return v.IncludeUsers }).(pulumi.StringPtrOutput)
}

func (o LookupGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// When this override is set, User in the group can set Xray security and compliance policies. Default value is `false`.
func (o LookupGroupResultOutput) PolicyManager() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *bool { return v.PolicyManager }).(pulumi.BoolPtrOutput)
}

// The realm for the group.
func (o LookupGroupResultOutput) Realm() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Realm }).(pulumi.StringOutput)
}

// The realm attributes for the group.
func (o LookupGroupResultOutput) RealmAttributes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *string { return v.RealmAttributes }).(pulumi.StringPtrOutput)
}

// When this override is set, User in the group can manage Xray Reports on any resource type. Default value is `false`.
func (o LookupGroupResultOutput) ReportsManager() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *bool { return v.ReportsManager }).(pulumi.BoolPtrOutput)
}

// List of users assigned to the group. Set includeUsers to `true` to retrieve this list.
func (o LookupGroupResultOutput) UsersNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []string { return v.UsersNames }).(pulumi.StringArrayOutput)
}

// When this override is set, User in the group can manage Xray Watches on any resource type. Default value is `false`.
func (o LookupGroupResultOutput) WatchManager() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *bool { return v.WatchManager }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupResultOutput{})
}
