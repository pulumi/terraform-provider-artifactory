// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-artifactory/sdk/v6/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Artifactory unmanaged user resource. This can be used to create and manage Artifactory users.
// The password is a required field by the [Artifactory API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-CreateorReplaceUser), but we made it optional in this resource to accommodate the scenario where the password is not needed and will be reset by the actual user later. When the optional attribute `password` is omitted, a random password is generated according to current Artifactory password policy.
//
// > The generated password won't be stored in the TF state and can not be recovered. The user must reset the password to be able to log in. An admin can always generate the access key for the user as well. The password change won't trigger state drift.
//
// > This resource is an alias for `User` resource, it is identical and was added for clarity and compatibility purposes. We don't recommend to use this resource unless there is a specific use case for it. Recommended resource is `ManagedUser`.
//
// ## Example Usage
//
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
//			// Create a new Artifactory user called terraform
//			_, err := artifactory.NewUnmanagedUser(ctx, "test-user", &artifactory.UnmanagedUserArgs{
//				Name:  pulumi.String("terraform"),
//				Email: pulumi.String("test-user@artifactory-terraform.com"),
//				Groups: pulumi.StringArray{
//					pulumi.String("logged-in-users"),
//				},
//				Password: pulumi.String("my super secret password"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Managing groups relationship
//
// See our recommendation on how to manage user-group relationship.
//
// ## Import
//
// Users can be imported using their name, e.g.
//
// ```sh
// $ pulumi import artifactory:index/unmanagedUser:UnmanagedUser test-user myusername
// ```
type UnmanagedUser struct {
	pulumi.CustomResourceState

	// When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
	Admin pulumi.BoolOutput `pulumi:"admin"`
	// When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
	DisableUiAccess pulumi.BoolOutput `pulumi:"disableUiAccess"`
	// Email for user.
	Email pulumi.StringOutput `pulumi:"email"`
	// List of groups this user is a part of. **Notes:** If this attribute is not specified then user's group membership is set to empty. User will not be part of default "readers" group automatically.
	Groups pulumi.StringArrayOutput `pulumi:"groups"`
	// When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
	InternalPasswordDisabled pulumi.BoolOutput `pulumi:"internalPasswordDisabled"`
	// Username for user. May contain lowercase letters, numbers and symbols: '.-_@'
	Name pulumi.StringOutput `pulumi:"name"`
	// Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters.
	Password pulumi.StringOutput `pulumi:"password"`
	// When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
	ProfileUpdatable pulumi.BoolOutput `pulumi:"profileUpdatable"`
}

// NewUnmanagedUser registers a new resource with the given unique name, arguments, and options.
func NewUnmanagedUser(ctx *pulumi.Context,
	name string, args *UnmanagedUserArgs, opts ...pulumi.ResourceOption) (*UnmanagedUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UnmanagedUser
	err := ctx.RegisterResource("artifactory:index/unmanagedUser:UnmanagedUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUnmanagedUser gets an existing UnmanagedUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUnmanagedUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UnmanagedUserState, opts ...pulumi.ResourceOption) (*UnmanagedUser, error) {
	var resource UnmanagedUser
	err := ctx.ReadResource("artifactory:index/unmanagedUser:UnmanagedUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UnmanagedUser resources.
type unmanagedUserState struct {
	// When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
	Admin *bool `pulumi:"admin"`
	// When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
	DisableUiAccess *bool `pulumi:"disableUiAccess"`
	// Email for user.
	Email *string `pulumi:"email"`
	// List of groups this user is a part of. **Notes:** If this attribute is not specified then user's group membership is set to empty. User will not be part of default "readers" group automatically.
	Groups []string `pulumi:"groups"`
	// When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
	InternalPasswordDisabled *bool `pulumi:"internalPasswordDisabled"`
	// Username for user. May contain lowercase letters, numbers and symbols: '.-_@'
	Name *string `pulumi:"name"`
	// Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters.
	Password *string `pulumi:"password"`
	// When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
	ProfileUpdatable *bool `pulumi:"profileUpdatable"`
}

type UnmanagedUserState struct {
	// When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
	Admin pulumi.BoolPtrInput
	// When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
	DisableUiAccess pulumi.BoolPtrInput
	// Email for user.
	Email pulumi.StringPtrInput
	// List of groups this user is a part of. **Notes:** If this attribute is not specified then user's group membership is set to empty. User will not be part of default "readers" group automatically.
	Groups pulumi.StringArrayInput
	// When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
	InternalPasswordDisabled pulumi.BoolPtrInput
	// Username for user. May contain lowercase letters, numbers and symbols: '.-_@'
	Name pulumi.StringPtrInput
	// Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters.
	Password pulumi.StringPtrInput
	// When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
	ProfileUpdatable pulumi.BoolPtrInput
}

func (UnmanagedUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*unmanagedUserState)(nil)).Elem()
}

type unmanagedUserArgs struct {
	// When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
	Admin *bool `pulumi:"admin"`
	// When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
	DisableUiAccess *bool `pulumi:"disableUiAccess"`
	// Email for user.
	Email string `pulumi:"email"`
	// List of groups this user is a part of. **Notes:** If this attribute is not specified then user's group membership is set to empty. User will not be part of default "readers" group automatically.
	Groups []string `pulumi:"groups"`
	// When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
	InternalPasswordDisabled *bool `pulumi:"internalPasswordDisabled"`
	// Username for user. May contain lowercase letters, numbers and symbols: '.-_@'
	Name *string `pulumi:"name"`
	// Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters.
	Password *string `pulumi:"password"`
	// When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
	ProfileUpdatable *bool `pulumi:"profileUpdatable"`
}

// The set of arguments for constructing a UnmanagedUser resource.
type UnmanagedUserArgs struct {
	// When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
	Admin pulumi.BoolPtrInput
	// When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
	DisableUiAccess pulumi.BoolPtrInput
	// Email for user.
	Email pulumi.StringInput
	// List of groups this user is a part of. **Notes:** If this attribute is not specified then user's group membership is set to empty. User will not be part of default "readers" group automatically.
	Groups pulumi.StringArrayInput
	// When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
	InternalPasswordDisabled pulumi.BoolPtrInput
	// Username for user. May contain lowercase letters, numbers and symbols: '.-_@'
	Name pulumi.StringPtrInput
	// Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters.
	Password pulumi.StringPtrInput
	// When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
	ProfileUpdatable pulumi.BoolPtrInput
}

func (UnmanagedUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*unmanagedUserArgs)(nil)).Elem()
}

type UnmanagedUserInput interface {
	pulumi.Input

	ToUnmanagedUserOutput() UnmanagedUserOutput
	ToUnmanagedUserOutputWithContext(ctx context.Context) UnmanagedUserOutput
}

func (*UnmanagedUser) ElementType() reflect.Type {
	return reflect.TypeOf((**UnmanagedUser)(nil)).Elem()
}

func (i *UnmanagedUser) ToUnmanagedUserOutput() UnmanagedUserOutput {
	return i.ToUnmanagedUserOutputWithContext(context.Background())
}

func (i *UnmanagedUser) ToUnmanagedUserOutputWithContext(ctx context.Context) UnmanagedUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UnmanagedUserOutput)
}

// UnmanagedUserArrayInput is an input type that accepts UnmanagedUserArray and UnmanagedUserArrayOutput values.
// You can construct a concrete instance of `UnmanagedUserArrayInput` via:
//
//	UnmanagedUserArray{ UnmanagedUserArgs{...} }
type UnmanagedUserArrayInput interface {
	pulumi.Input

	ToUnmanagedUserArrayOutput() UnmanagedUserArrayOutput
	ToUnmanagedUserArrayOutputWithContext(context.Context) UnmanagedUserArrayOutput
}

type UnmanagedUserArray []UnmanagedUserInput

func (UnmanagedUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UnmanagedUser)(nil)).Elem()
}

func (i UnmanagedUserArray) ToUnmanagedUserArrayOutput() UnmanagedUserArrayOutput {
	return i.ToUnmanagedUserArrayOutputWithContext(context.Background())
}

func (i UnmanagedUserArray) ToUnmanagedUserArrayOutputWithContext(ctx context.Context) UnmanagedUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UnmanagedUserArrayOutput)
}

// UnmanagedUserMapInput is an input type that accepts UnmanagedUserMap and UnmanagedUserMapOutput values.
// You can construct a concrete instance of `UnmanagedUserMapInput` via:
//
//	UnmanagedUserMap{ "key": UnmanagedUserArgs{...} }
type UnmanagedUserMapInput interface {
	pulumi.Input

	ToUnmanagedUserMapOutput() UnmanagedUserMapOutput
	ToUnmanagedUserMapOutputWithContext(context.Context) UnmanagedUserMapOutput
}

type UnmanagedUserMap map[string]UnmanagedUserInput

func (UnmanagedUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UnmanagedUser)(nil)).Elem()
}

func (i UnmanagedUserMap) ToUnmanagedUserMapOutput() UnmanagedUserMapOutput {
	return i.ToUnmanagedUserMapOutputWithContext(context.Background())
}

func (i UnmanagedUserMap) ToUnmanagedUserMapOutputWithContext(ctx context.Context) UnmanagedUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UnmanagedUserMapOutput)
}

type UnmanagedUserOutput struct{ *pulumi.OutputState }

func (UnmanagedUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UnmanagedUser)(nil)).Elem()
}

func (o UnmanagedUserOutput) ToUnmanagedUserOutput() UnmanagedUserOutput {
	return o
}

func (o UnmanagedUserOutput) ToUnmanagedUserOutputWithContext(ctx context.Context) UnmanagedUserOutput {
	return o
}

// When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
func (o UnmanagedUserOutput) Admin() pulumi.BoolOutput {
	return o.ApplyT(func(v *UnmanagedUser) pulumi.BoolOutput { return v.Admin }).(pulumi.BoolOutput)
}

// When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
func (o UnmanagedUserOutput) DisableUiAccess() pulumi.BoolOutput {
	return o.ApplyT(func(v *UnmanagedUser) pulumi.BoolOutput { return v.DisableUiAccess }).(pulumi.BoolOutput)
}

// Email for user.
func (o UnmanagedUserOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *UnmanagedUser) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// List of groups this user is a part of. **Notes:** If this attribute is not specified then user's group membership is set to empty. User will not be part of default "readers" group automatically.
func (o UnmanagedUserOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *UnmanagedUser) pulumi.StringArrayOutput { return v.Groups }).(pulumi.StringArrayOutput)
}

// When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
func (o UnmanagedUserOutput) InternalPasswordDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *UnmanagedUser) pulumi.BoolOutput { return v.InternalPasswordDisabled }).(pulumi.BoolOutput)
}

// Username for user. May contain lowercase letters, numbers and symbols: '.-_@'
func (o UnmanagedUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UnmanagedUser) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters.
func (o UnmanagedUserOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *UnmanagedUser) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
func (o UnmanagedUserOutput) ProfileUpdatable() pulumi.BoolOutput {
	return o.ApplyT(func(v *UnmanagedUser) pulumi.BoolOutput { return v.ProfileUpdatable }).(pulumi.BoolOutput)
}

type UnmanagedUserArrayOutput struct{ *pulumi.OutputState }

func (UnmanagedUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UnmanagedUser)(nil)).Elem()
}

func (o UnmanagedUserArrayOutput) ToUnmanagedUserArrayOutput() UnmanagedUserArrayOutput {
	return o
}

func (o UnmanagedUserArrayOutput) ToUnmanagedUserArrayOutputWithContext(ctx context.Context) UnmanagedUserArrayOutput {
	return o
}

func (o UnmanagedUserArrayOutput) Index(i pulumi.IntInput) UnmanagedUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UnmanagedUser {
		return vs[0].([]*UnmanagedUser)[vs[1].(int)]
	}).(UnmanagedUserOutput)
}

type UnmanagedUserMapOutput struct{ *pulumi.OutputState }

func (UnmanagedUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UnmanagedUser)(nil)).Elem()
}

func (o UnmanagedUserMapOutput) ToUnmanagedUserMapOutput() UnmanagedUserMapOutput {
	return o
}

func (o UnmanagedUserMapOutput) ToUnmanagedUserMapOutputWithContext(ctx context.Context) UnmanagedUserMapOutput {
	return o
}

func (o UnmanagedUserMapOutput) MapIndex(k pulumi.StringInput) UnmanagedUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UnmanagedUser {
		return vs[0].(map[string]*UnmanagedUser)[vs[1].(string)]
	}).(UnmanagedUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UnmanagedUserInput)(nil)).Elem(), &UnmanagedUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*UnmanagedUserArrayInput)(nil)).Elem(), UnmanagedUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UnmanagedUserMapInput)(nil)).Elem(), UnmanagedUserMap{})
	pulumi.RegisterOutputType(UnmanagedUserOutput{})
	pulumi.RegisterOutputType(UnmanagedUserArrayOutput{})
	pulumi.RegisterOutputType(UnmanagedUserMapOutput{})
}
