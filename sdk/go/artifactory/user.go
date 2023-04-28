// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Artifactory user resource. This can be used to create and manage Artifactory users.
// The password is a required field by the [Artifactory API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-CreateorReplaceUser), but we made it optional in this resource to accommodate the scenario where the password is not needed and will be reset by the actual user later.\
// When the optional attribute `password` is omitted, a random password is generated according to current Artifactory password policy.
//
// > The generated password won't be stored in the TF state and can not be recovered. The user must reset the password to be able to log in. An admin can always generate the access key for the user as well. The password change won't trigger state drift. We don't recommend to use this resource unless there is a specific use case for it. Recommended resource is `ManagedUser`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v3/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewUser(ctx, "test-user", &artifactory.UserArgs{
//				Email: pulumi.String("test-user@artifactory-terraform.com"),
//				Groups: pulumi.StringArray{
//					pulumi.String("logged-in-users"),
//					pulumi.String("readers"),
//				},
//				Name:     pulumi.String("terraform"),
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
// ## Managing groups relationship
//
// See our recommendation on how to manage user-group relationship.
//
// ## Import
//
// Users can be imported using their name, e.g.
//
// ```sh
//
//	$ pulumi import artifactory:index/user:User test-user myusername
//
// ```
type User struct {
	pulumi.CustomResourceState

	// When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
	Admin pulumi.BoolPtrOutput `pulumi:"admin"`
	// When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
	DisableUiAccess pulumi.BoolPtrOutput `pulumi:"disableUiAccess"`
	// Email for user.
	Email pulumi.StringOutput `pulumi:"email"`
	// List of groups this user is a part of. **Notes:** If this attribute is not specified then user's group membership set to empty. User will not be part of default "readers" group automatically.
	Groups pulumi.StringArrayOutput `pulumi:"groups"`
	// When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
	InternalPasswordDisabled pulumi.BoolPtrOutput `pulumi:"internalPasswordDisabled"`
	// Username for user.
	Name pulumi.StringOutput `pulumi:"name"`
	// Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
	ProfileUpdatable pulumi.BoolPtrOutput `pulumi:"profileUpdatable"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	var resource User
	err := ctx.RegisterResource("artifactory:index/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("artifactory:index/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
	Admin *bool `pulumi:"admin"`
	// When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
	DisableUiAccess *bool `pulumi:"disableUiAccess"`
	// Email for user.
	Email *string `pulumi:"email"`
	// List of groups this user is a part of. **Notes:** If this attribute is not specified then user's group membership set to empty. User will not be part of default "readers" group automatically.
	Groups []string `pulumi:"groups"`
	// When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
	InternalPasswordDisabled *bool `pulumi:"internalPasswordDisabled"`
	// Username for user.
	Name *string `pulumi:"name"`
	// Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters.
	Password *string `pulumi:"password"`
	// When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
	ProfileUpdatable *bool `pulumi:"profileUpdatable"`
}

type UserState struct {
	// When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
	Admin pulumi.BoolPtrInput
	// When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
	DisableUiAccess pulumi.BoolPtrInput
	// Email for user.
	Email pulumi.StringPtrInput
	// List of groups this user is a part of. **Notes:** If this attribute is not specified then user's group membership set to empty. User will not be part of default "readers" group automatically.
	Groups pulumi.StringArrayInput
	// When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
	InternalPasswordDisabled pulumi.BoolPtrInput
	// Username for user.
	Name pulumi.StringPtrInput
	// Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters.
	Password pulumi.StringPtrInput
	// When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
	ProfileUpdatable pulumi.BoolPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
	Admin *bool `pulumi:"admin"`
	// When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
	DisableUiAccess *bool `pulumi:"disableUiAccess"`
	// Email for user.
	Email string `pulumi:"email"`
	// List of groups this user is a part of. **Notes:** If this attribute is not specified then user's group membership set to empty. User will not be part of default "readers" group automatically.
	Groups []string `pulumi:"groups"`
	// When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
	InternalPasswordDisabled *bool `pulumi:"internalPasswordDisabled"`
	// Username for user.
	Name string `pulumi:"name"`
	// Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters.
	Password *string `pulumi:"password"`
	// When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
	ProfileUpdatable *bool `pulumi:"profileUpdatable"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
	Admin pulumi.BoolPtrInput
	// When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
	DisableUiAccess pulumi.BoolPtrInput
	// Email for user.
	Email pulumi.StringInput
	// List of groups this user is a part of. **Notes:** If this attribute is not specified then user's group membership set to empty. User will not be part of default "readers" group automatically.
	Groups pulumi.StringArrayInput
	// When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
	InternalPasswordDisabled pulumi.BoolPtrInput
	// Username for user.
	Name pulumi.StringInput
	// Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters.
	Password pulumi.StringPtrInput
	// When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
	ProfileUpdatable pulumi.BoolPtrInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

type UserInput interface {
	pulumi.Input

	ToUserOutput() UserOutput
	ToUserOutputWithContext(ctx context.Context) UserOutput
}

func (*User) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

// UserArrayInput is an input type that accepts UserArray and UserArrayOutput values.
// You can construct a concrete instance of `UserArrayInput` via:
//
//	UserArray{ UserArgs{...} }
type UserArrayInput interface {
	pulumi.Input

	ToUserArrayOutput() UserArrayOutput
	ToUserArrayOutputWithContext(context.Context) UserArrayOutput
}

type UserArray []UserInput

func (UserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (i UserArray) ToUserArrayOutput() UserArrayOutput {
	return i.ToUserArrayOutputWithContext(context.Background())
}

func (i UserArray) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserArrayOutput)
}

// UserMapInput is an input type that accepts UserMap and UserMapOutput values.
// You can construct a concrete instance of `UserMapInput` via:
//
//	UserMap{ "key": UserArgs{...} }
type UserMapInput interface {
	pulumi.Input

	ToUserMapOutput() UserMapOutput
	ToUserMapOutputWithContext(context.Context) UserMapOutput
}

type UserMap map[string]UserInput

func (UserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (i UserMap) ToUserMapOutput() UserMapOutput {
	return i.ToUserMapOutputWithContext(context.Background())
}

func (i UserMap) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMapOutput)
}

type UserOutput struct{ *pulumi.OutputState }

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

// When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
func (o UserOutput) Admin() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.Admin }).(pulumi.BoolPtrOutput)
}

// When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
func (o UserOutput) DisableUiAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.DisableUiAccess }).(pulumi.BoolPtrOutput)
}

// Email for user.
func (o UserOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// List of groups this user is a part of. **Notes:** If this attribute is not specified then user's group membership set to empty. User will not be part of default "readers" group automatically.
func (o UserOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *User) pulumi.StringArrayOutput { return v.Groups }).(pulumi.StringArrayOutput)
}

// When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
func (o UserOutput) InternalPasswordDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.InternalPasswordDisabled }).(pulumi.BoolPtrOutput)
}

// Username for user.
func (o UserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters.
func (o UserOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
func (o UserOutput) ProfileUpdatable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.ProfileUpdatable }).(pulumi.BoolPtrOutput)
}

type UserArrayOutput struct{ *pulumi.OutputState }

func (UserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (o UserArrayOutput) ToUserArrayOutput() UserArrayOutput {
	return o
}

func (o UserArrayOutput) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return o
}

func (o UserArrayOutput) Index(i pulumi.IntInput) UserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *User {
		return vs[0].([]*User)[vs[1].(int)]
	}).(UserOutput)
}

type UserMapOutput struct{ *pulumi.OutputState }

func (UserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (o UserMapOutput) ToUserMapOutput() UserMapOutput {
	return o
}

func (o UserMapOutput) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return o
}

func (o UserMapOutput) MapIndex(k pulumi.StringInput) UserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *User {
		return vs[0].(map[string]*User)[vs[1].(string)]
	}).(UserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserInput)(nil)).Elem(), &User{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserArrayInput)(nil)).Elem(), UserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserMapInput)(nil)).Elem(), UserMap{})
	pulumi.RegisterOutputType(UserOutput{})
	pulumi.RegisterOutputType(UserArrayOutput{})
	pulumi.RegisterOutputType(UserMapOutput{})
}
