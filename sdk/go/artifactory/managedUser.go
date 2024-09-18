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

// ## Example Usage
//
// ## Import
//
// ```sh
// $ pulumi import artifactory:index/managedUser:ManagedUser test-user myusername
// ```
type ManagedUser struct {
	pulumi.CustomResourceState

	// (Optional, Default: false) When enabled, this user is an administrator with all the ensuing privileges.
	Admin pulumi.BoolOutput `pulumi:"admin"`
	// (Optional, Default: true) When enabled, this user can only access the system through the REST API. This option cannot be set if the user has Admin privileges.
	DisableUiAccess pulumi.BoolOutput `pulumi:"disableUiAccess"`
	// Email for user.
	Email pulumi.StringOutput `pulumi:"email"`
	// List of groups this user is a part of. **Notes:** If this attribute is not specified then user's group membership is set to empty. User will not be part of default "readers" group automatically.
	Groups pulumi.StringArrayOutput `pulumi:"groups"`
	// (Optional, Default: false) When enabled, disables the fallback mechanism for using an internal password when external authentication (such as LDAP) is enabled.
	InternalPasswordDisabled pulumi.BoolOutput `pulumi:"internalPasswordDisabled"`
	// Username for user. May contain lowercase letters, numbers and symbols: '.-_@' for self-hosted. For SaaS, '+' is also allowed.
	Name pulumi.StringOutput `pulumi:"name"`
	// Password for the user.
	Password pulumi.StringOutput `pulumi:"password"`
	// Password policy to match JFrog Access to provide pre-apply validation. Default values: `uppercase=1`, `lowercase=1`, `special_char=0`, `digit=1`, `length=8`. Also see [Supported Access Configurations](https://jfrog.com/help/r/jfrog-installation-setup-documentation/supported-access-configurations) for more details
	PasswordPolicy ManagedUserPasswordPolicyPtrOutput `pulumi:"passwordPolicy"`
	// (Optional, Default: true) When enabled, this user can update their profile details (except for the password. Only an administrator can update the password). There may be cases in which you want to leave this unset to prevent users from updating their profile. For example, a departmental user with a single password shared between all department members.
	ProfileUpdatable pulumi.BoolOutput `pulumi:"profileUpdatable"`
}

// NewManagedUser registers a new resource with the given unique name, arguments, and options.
func NewManagedUser(ctx *pulumi.Context,
	name string, args *ManagedUserArgs, opts ...pulumi.ResourceOption) (*ManagedUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ManagedUser
	err := ctx.RegisterResource("artifactory:index/managedUser:ManagedUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedUser gets an existing ManagedUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedUserState, opts ...pulumi.ResourceOption) (*ManagedUser, error) {
	var resource ManagedUser
	err := ctx.ReadResource("artifactory:index/managedUser:ManagedUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedUser resources.
type managedUserState struct {
	// (Optional, Default: false) When enabled, this user is an administrator with all the ensuing privileges.
	Admin *bool `pulumi:"admin"`
	// (Optional, Default: true) When enabled, this user can only access the system through the REST API. This option cannot be set if the user has Admin privileges.
	DisableUiAccess *bool `pulumi:"disableUiAccess"`
	// Email for user.
	Email *string `pulumi:"email"`
	// List of groups this user is a part of. **Notes:** If this attribute is not specified then user's group membership is set to empty. User will not be part of default "readers" group automatically.
	Groups []string `pulumi:"groups"`
	// (Optional, Default: false) When enabled, disables the fallback mechanism for using an internal password when external authentication (such as LDAP) is enabled.
	InternalPasswordDisabled *bool `pulumi:"internalPasswordDisabled"`
	// Username for user. May contain lowercase letters, numbers and symbols: '.-_@' for self-hosted. For SaaS, '+' is also allowed.
	Name *string `pulumi:"name"`
	// Password for the user.
	Password *string `pulumi:"password"`
	// Password policy to match JFrog Access to provide pre-apply validation. Default values: `uppercase=1`, `lowercase=1`, `special_char=0`, `digit=1`, `length=8`. Also see [Supported Access Configurations](https://jfrog.com/help/r/jfrog-installation-setup-documentation/supported-access-configurations) for more details
	PasswordPolicy *ManagedUserPasswordPolicy `pulumi:"passwordPolicy"`
	// (Optional, Default: true) When enabled, this user can update their profile details (except for the password. Only an administrator can update the password). There may be cases in which you want to leave this unset to prevent users from updating their profile. For example, a departmental user with a single password shared between all department members.
	ProfileUpdatable *bool `pulumi:"profileUpdatable"`
}

type ManagedUserState struct {
	// (Optional, Default: false) When enabled, this user is an administrator with all the ensuing privileges.
	Admin pulumi.BoolPtrInput
	// (Optional, Default: true) When enabled, this user can only access the system through the REST API. This option cannot be set if the user has Admin privileges.
	DisableUiAccess pulumi.BoolPtrInput
	// Email for user.
	Email pulumi.StringPtrInput
	// List of groups this user is a part of. **Notes:** If this attribute is not specified then user's group membership is set to empty. User will not be part of default "readers" group automatically.
	Groups pulumi.StringArrayInput
	// (Optional, Default: false) When enabled, disables the fallback mechanism for using an internal password when external authentication (such as LDAP) is enabled.
	InternalPasswordDisabled pulumi.BoolPtrInput
	// Username for user. May contain lowercase letters, numbers and symbols: '.-_@' for self-hosted. For SaaS, '+' is also allowed.
	Name pulumi.StringPtrInput
	// Password for the user.
	Password pulumi.StringPtrInput
	// Password policy to match JFrog Access to provide pre-apply validation. Default values: `uppercase=1`, `lowercase=1`, `special_char=0`, `digit=1`, `length=8`. Also see [Supported Access Configurations](https://jfrog.com/help/r/jfrog-installation-setup-documentation/supported-access-configurations) for more details
	PasswordPolicy ManagedUserPasswordPolicyPtrInput
	// (Optional, Default: true) When enabled, this user can update their profile details (except for the password. Only an administrator can update the password). There may be cases in which you want to leave this unset to prevent users from updating their profile. For example, a departmental user with a single password shared between all department members.
	ProfileUpdatable pulumi.BoolPtrInput
}

func (ManagedUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedUserState)(nil)).Elem()
}

type managedUserArgs struct {
	// (Optional, Default: false) When enabled, this user is an administrator with all the ensuing privileges.
	Admin *bool `pulumi:"admin"`
	// (Optional, Default: true) When enabled, this user can only access the system through the REST API. This option cannot be set if the user has Admin privileges.
	DisableUiAccess *bool `pulumi:"disableUiAccess"`
	// Email for user.
	Email string `pulumi:"email"`
	// List of groups this user is a part of. **Notes:** If this attribute is not specified then user's group membership is set to empty. User will not be part of default "readers" group automatically.
	Groups []string `pulumi:"groups"`
	// (Optional, Default: false) When enabled, disables the fallback mechanism for using an internal password when external authentication (such as LDAP) is enabled.
	InternalPasswordDisabled *bool `pulumi:"internalPasswordDisabled"`
	// Username for user. May contain lowercase letters, numbers and symbols: '.-_@' for self-hosted. For SaaS, '+' is also allowed.
	Name *string `pulumi:"name"`
	// Password for the user.
	Password string `pulumi:"password"`
	// Password policy to match JFrog Access to provide pre-apply validation. Default values: `uppercase=1`, `lowercase=1`, `special_char=0`, `digit=1`, `length=8`. Also see [Supported Access Configurations](https://jfrog.com/help/r/jfrog-installation-setup-documentation/supported-access-configurations) for more details
	PasswordPolicy *ManagedUserPasswordPolicy `pulumi:"passwordPolicy"`
	// (Optional, Default: true) When enabled, this user can update their profile details (except for the password. Only an administrator can update the password). There may be cases in which you want to leave this unset to prevent users from updating their profile. For example, a departmental user with a single password shared between all department members.
	ProfileUpdatable *bool `pulumi:"profileUpdatable"`
}

// The set of arguments for constructing a ManagedUser resource.
type ManagedUserArgs struct {
	// (Optional, Default: false) When enabled, this user is an administrator with all the ensuing privileges.
	Admin pulumi.BoolPtrInput
	// (Optional, Default: true) When enabled, this user can only access the system through the REST API. This option cannot be set if the user has Admin privileges.
	DisableUiAccess pulumi.BoolPtrInput
	// Email for user.
	Email pulumi.StringInput
	// List of groups this user is a part of. **Notes:** If this attribute is not specified then user's group membership is set to empty. User will not be part of default "readers" group automatically.
	Groups pulumi.StringArrayInput
	// (Optional, Default: false) When enabled, disables the fallback mechanism for using an internal password when external authentication (such as LDAP) is enabled.
	InternalPasswordDisabled pulumi.BoolPtrInput
	// Username for user. May contain lowercase letters, numbers and symbols: '.-_@' for self-hosted. For SaaS, '+' is also allowed.
	Name pulumi.StringPtrInput
	// Password for the user.
	Password pulumi.StringInput
	// Password policy to match JFrog Access to provide pre-apply validation. Default values: `uppercase=1`, `lowercase=1`, `special_char=0`, `digit=1`, `length=8`. Also see [Supported Access Configurations](https://jfrog.com/help/r/jfrog-installation-setup-documentation/supported-access-configurations) for more details
	PasswordPolicy ManagedUserPasswordPolicyPtrInput
	// (Optional, Default: true) When enabled, this user can update their profile details (except for the password. Only an administrator can update the password). There may be cases in which you want to leave this unset to prevent users from updating their profile. For example, a departmental user with a single password shared between all department members.
	ProfileUpdatable pulumi.BoolPtrInput
}

func (ManagedUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedUserArgs)(nil)).Elem()
}

type ManagedUserInput interface {
	pulumi.Input

	ToManagedUserOutput() ManagedUserOutput
	ToManagedUserOutputWithContext(ctx context.Context) ManagedUserOutput
}

func (*ManagedUser) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedUser)(nil)).Elem()
}

func (i *ManagedUser) ToManagedUserOutput() ManagedUserOutput {
	return i.ToManagedUserOutputWithContext(context.Background())
}

func (i *ManagedUser) ToManagedUserOutputWithContext(ctx context.Context) ManagedUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedUserOutput)
}

// ManagedUserArrayInput is an input type that accepts ManagedUserArray and ManagedUserArrayOutput values.
// You can construct a concrete instance of `ManagedUserArrayInput` via:
//
//	ManagedUserArray{ ManagedUserArgs{...} }
type ManagedUserArrayInput interface {
	pulumi.Input

	ToManagedUserArrayOutput() ManagedUserArrayOutput
	ToManagedUserArrayOutputWithContext(context.Context) ManagedUserArrayOutput
}

type ManagedUserArray []ManagedUserInput

func (ManagedUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ManagedUser)(nil)).Elem()
}

func (i ManagedUserArray) ToManagedUserArrayOutput() ManagedUserArrayOutput {
	return i.ToManagedUserArrayOutputWithContext(context.Background())
}

func (i ManagedUserArray) ToManagedUserArrayOutputWithContext(ctx context.Context) ManagedUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedUserArrayOutput)
}

// ManagedUserMapInput is an input type that accepts ManagedUserMap and ManagedUserMapOutput values.
// You can construct a concrete instance of `ManagedUserMapInput` via:
//
//	ManagedUserMap{ "key": ManagedUserArgs{...} }
type ManagedUserMapInput interface {
	pulumi.Input

	ToManagedUserMapOutput() ManagedUserMapOutput
	ToManagedUserMapOutputWithContext(context.Context) ManagedUserMapOutput
}

type ManagedUserMap map[string]ManagedUserInput

func (ManagedUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ManagedUser)(nil)).Elem()
}

func (i ManagedUserMap) ToManagedUserMapOutput() ManagedUserMapOutput {
	return i.ToManagedUserMapOutputWithContext(context.Background())
}

func (i ManagedUserMap) ToManagedUserMapOutputWithContext(ctx context.Context) ManagedUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedUserMapOutput)
}

type ManagedUserOutput struct{ *pulumi.OutputState }

func (ManagedUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedUser)(nil)).Elem()
}

func (o ManagedUserOutput) ToManagedUserOutput() ManagedUserOutput {
	return o
}

func (o ManagedUserOutput) ToManagedUserOutputWithContext(ctx context.Context) ManagedUserOutput {
	return o
}

// (Optional, Default: false) When enabled, this user is an administrator with all the ensuing privileges.
func (o ManagedUserOutput) Admin() pulumi.BoolOutput {
	return o.ApplyT(func(v *ManagedUser) pulumi.BoolOutput { return v.Admin }).(pulumi.BoolOutput)
}

// (Optional, Default: true) When enabled, this user can only access the system through the REST API. This option cannot be set if the user has Admin privileges.
func (o ManagedUserOutput) DisableUiAccess() pulumi.BoolOutput {
	return o.ApplyT(func(v *ManagedUser) pulumi.BoolOutput { return v.DisableUiAccess }).(pulumi.BoolOutput)
}

// Email for user.
func (o ManagedUserOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedUser) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// List of groups this user is a part of. **Notes:** If this attribute is not specified then user's group membership is set to empty. User will not be part of default "readers" group automatically.
func (o ManagedUserOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedUser) pulumi.StringArrayOutput { return v.Groups }).(pulumi.StringArrayOutput)
}

// (Optional, Default: false) When enabled, disables the fallback mechanism for using an internal password when external authentication (such as LDAP) is enabled.
func (o ManagedUserOutput) InternalPasswordDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ManagedUser) pulumi.BoolOutput { return v.InternalPasswordDisabled }).(pulumi.BoolOutput)
}

// Username for user. May contain lowercase letters, numbers and symbols: '.-_@' for self-hosted. For SaaS, '+' is also allowed.
func (o ManagedUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedUser) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Password for the user.
func (o ManagedUserOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedUser) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// Password policy to match JFrog Access to provide pre-apply validation. Default values: `uppercase=1`, `lowercase=1`, `special_char=0`, `digit=1`, `length=8`. Also see [Supported Access Configurations](https://jfrog.com/help/r/jfrog-installation-setup-documentation/supported-access-configurations) for more details
func (o ManagedUserOutput) PasswordPolicy() ManagedUserPasswordPolicyPtrOutput {
	return o.ApplyT(func(v *ManagedUser) ManagedUserPasswordPolicyPtrOutput { return v.PasswordPolicy }).(ManagedUserPasswordPolicyPtrOutput)
}

// (Optional, Default: true) When enabled, this user can update their profile details (except for the password. Only an administrator can update the password). There may be cases in which you want to leave this unset to prevent users from updating their profile. For example, a departmental user with a single password shared between all department members.
func (o ManagedUserOutput) ProfileUpdatable() pulumi.BoolOutput {
	return o.ApplyT(func(v *ManagedUser) pulumi.BoolOutput { return v.ProfileUpdatable }).(pulumi.BoolOutput)
}

type ManagedUserArrayOutput struct{ *pulumi.OutputState }

func (ManagedUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ManagedUser)(nil)).Elem()
}

func (o ManagedUserArrayOutput) ToManagedUserArrayOutput() ManagedUserArrayOutput {
	return o
}

func (o ManagedUserArrayOutput) ToManagedUserArrayOutputWithContext(ctx context.Context) ManagedUserArrayOutput {
	return o
}

func (o ManagedUserArrayOutput) Index(i pulumi.IntInput) ManagedUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ManagedUser {
		return vs[0].([]*ManagedUser)[vs[1].(int)]
	}).(ManagedUserOutput)
}

type ManagedUserMapOutput struct{ *pulumi.OutputState }

func (ManagedUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ManagedUser)(nil)).Elem()
}

func (o ManagedUserMapOutput) ToManagedUserMapOutput() ManagedUserMapOutput {
	return o
}

func (o ManagedUserMapOutput) ToManagedUserMapOutputWithContext(ctx context.Context) ManagedUserMapOutput {
	return o
}

func (o ManagedUserMapOutput) MapIndex(k pulumi.StringInput) ManagedUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ManagedUser {
		return vs[0].(map[string]*ManagedUser)[vs[1].(string)]
	}).(ManagedUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedUserInput)(nil)).Elem(), &ManagedUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedUserArrayInput)(nil)).Elem(), ManagedUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedUserMapInput)(nil)).Elem(), ManagedUserMap{})
	pulumi.RegisterOutputType(ManagedUserOutput{})
	pulumi.RegisterOutputType(ManagedUserArrayOutput{})
	pulumi.RegisterOutputType(ManagedUserMapOutput{})
}
