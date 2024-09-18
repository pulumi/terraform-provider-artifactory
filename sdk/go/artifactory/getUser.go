// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory User Data Source
//
// Provides an Artifactory user data source. This can be used to read the configuration of users in artifactory.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.LookupUser(ctx, &artifactory.LookupUserArgs{
//				Name: "user1",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupUserResult
	err := ctx.Invoke("artifactory:index/getUser:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUser.
type LookupUserArgs struct {
	// When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
	Admin *bool `pulumi:"admin"`
	// When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
	DisableUiAccess *bool `pulumi:"disableUiAccess"`
	// Email for user.
	Email *string `pulumi:"email"`
	// List of groups this user is a part of.
	Groups []string `pulumi:"groups"`
	// When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
	InternalPasswordDisabled *bool `pulumi:"internalPasswordDisabled"`
	// Name of the user.
	Name string `pulumi:"name"`
	// When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
	ProfileUpdatable *bool `pulumi:"profileUpdatable"`
}

// A collection of values returned by getUser.
type LookupUserResult struct {
	// When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
	Admin *bool `pulumi:"admin"`
	// When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
	DisableUiAccess *bool `pulumi:"disableUiAccess"`
	// Email for user.
	Email *string `pulumi:"email"`
	// List of groups this user is a part of.
	Groups []string `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
	InternalPasswordDisabled *bool  `pulumi:"internalPasswordDisabled"`
	Name                     string `pulumi:"name"`
	// When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
	ProfileUpdatable *bool `pulumi:"profileUpdatable"`
}

func LookupUserOutput(ctx *pulumi.Context, args LookupUserOutputArgs, opts ...pulumi.InvokeOption) LookupUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserResult, error) {
			args := v.(LookupUserArgs)
			r, err := LookupUser(ctx, &args, opts...)
			var s LookupUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserResultOutput)
}

// A collection of arguments for invoking getUser.
type LookupUserOutputArgs struct {
	// When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
	Admin pulumi.BoolPtrInput `pulumi:"admin"`
	// When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
	DisableUiAccess pulumi.BoolPtrInput `pulumi:"disableUiAccess"`
	// Email for user.
	Email pulumi.StringPtrInput `pulumi:"email"`
	// List of groups this user is a part of.
	Groups pulumi.StringArrayInput `pulumi:"groups"`
	// When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
	InternalPasswordDisabled pulumi.BoolPtrInput `pulumi:"internalPasswordDisabled"`
	// Name of the user.
	Name pulumi.StringInput `pulumi:"name"`
	// When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
	ProfileUpdatable pulumi.BoolPtrInput `pulumi:"profileUpdatable"`
}

func (LookupUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserArgs)(nil)).Elem()
}

// A collection of values returned by getUser.
type LookupUserResultOutput struct{ *pulumi.OutputState }

func (LookupUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserResult)(nil)).Elem()
}

func (o LookupUserResultOutput) ToLookupUserResultOutput() LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) ToLookupUserResultOutputWithContext(ctx context.Context) LookupUserResultOutput {
	return o
}

// When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
func (o LookupUserResultOutput) Admin() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *bool { return v.Admin }).(pulumi.BoolPtrOutput)
}

// When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
func (o LookupUserResultOutput) DisableUiAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *bool { return v.DisableUiAccess }).(pulumi.BoolPtrOutput)
}

// Email for user.
func (o LookupUserResultOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.Email }).(pulumi.StringPtrOutput)
}

// List of groups this user is a part of.
func (o LookupUserResultOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUserResult) []string { return v.Groups }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
func (o LookupUserResultOutput) InternalPasswordDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *bool { return v.InternalPasswordDisabled }).(pulumi.BoolPtrOutput)
}

func (o LookupUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Name }).(pulumi.StringOutput)
}

// When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
func (o LookupUserResultOutput) ProfileUpdatable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *bool { return v.ProfileUpdatable }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserResultOutput{})
}
