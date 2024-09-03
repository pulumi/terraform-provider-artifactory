// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-artifactory/sdk/v7/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource can be used to manage Artifactory's LDAP settings for user authentication.
//
// When specified LDAP setting is active, Artifactory first attempts to authenticate the user against the LDAP server.
// If LDAP authentication fails, it then tries to authenticate via its internal database.
//
// ~>The `LdapSetting` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v7/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Configure Artifactory LDAP setting
//			_, err := artifactory.NewLdapSetting(ctx, "ldap_name", &artifactory.LdapSettingArgs{
//				Key:                      pulumi.String("ldap_name"),
//				Enabled:                  pulumi.Bool(true),
//				LdapUrl:                  pulumi.String("ldap://ldap_server_url"),
//				UserDnPattern:            pulumi.String("uid={0},ou=People"),
//				EmailAttribute:           pulumi.String("mail"),
//				AutoCreateUser:           pulumi.Bool(true),
//				LdapPoisoningProtection:  pulumi.Bool(true),
//				AllowUserToAccessProfile: pulumi.Bool(false),
//				PagingSupportEnabled:     pulumi.Bool(false),
//				SearchFilter:             pulumi.String("(uid={0})"),
//				SearchBase:               pulumi.String("ou=users"),
//				SearchSubTree:            pulumi.Bool(true),
//				ManagerDn:                pulumi.String("mgr_dn"),
//				ManagerPassword:          pulumi.String("mgr_passwd_random"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// Note: `Key` argument has to match to the resource name.Reference Link: [JFrog LDAP](https://www.jfrog.com/confluence/display/JFROG/LDAP)
//
// ## Import
//
// LDAP setting can be imported using the key, e.g.
//
// ```sh
// $ pulumi import artifactory:index/ldapSetting:LdapSetting ldap_name ldap_name
// ```
type LdapSetting struct {
	pulumi.CustomResourceState

	// When set, users created after logging in using LDAP will be able to access their profile page.  Default value is `false`.
	AllowUserToAccessProfile pulumi.BoolPtrOutput `pulumi:"allowUserToAccessProfile"`
	// When set, the system will automatically create new users for those who have logged in using LDAP, and assign them to the default groups.  Default value is `true`.
	AutoCreateUser pulumi.BoolPtrOutput `pulumi:"autoCreateUser"`
	// An attribute that can be used to map a user's email address to a user created automatically in Artifactory. Default value is `mail`.
	// - Note: If blank/empty string input was set for email_attribute, Default value `mail` takes effect. This is to match with Artifactory behavior.
	EmailAttribute pulumi.StringPtrOutput `pulumi:"emailAttribute"`
	// When set, these settings are enabled. Default value is `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The unique ID of the LDAP setting.
	Key pulumi.StringOutput `pulumi:"key"`
	// Protects against LDAP poisoning by filtering out users exposed to vulnerabilities.  Default value is `true`.
	LdapPoisoningProtection pulumi.BoolPtrOutput `pulumi:"ldapPoisoningProtection"`
	// Location of the LDAP server in the following format: ldap://myserver:myport/dc=sampledomain,dc=com. The URL should include the base DN used to search for and/or authenticate users.
	LdapUrl pulumi.StringOutput `pulumi:"ldapUrl"`
	// The full DN of a user with permissions that allow querying the LDAP server. When working with LDAP Groups, the user should have permissions for any extra group attributes such as memberOf.
	ManagerDn pulumi.StringPtrOutput `pulumi:"managerDn"`
	// The password of the user binding to the LDAP server when using "search" authentication.
	ManagerPassword pulumi.StringOutput `pulumi:"managerPassword"`
	// When set, supports paging results for the LDAP server. This feature requires that the LDAP Server supports a PagedResultsControl configuration.  Default value is `true`.
	PagingSupportEnabled pulumi.BoolPtrOutput `pulumi:"pagingSupportEnabled"`
	// The Context name in which to search relative to the base DN in the LDAP URL. Multiple search bases may be specified separated by a pipe ( | ).
	SearchBase pulumi.StringPtrOutput `pulumi:"searchBase"`
	// A filter expression used to search for the user DN that is used in LDAP authentication. This is an LDAP search filter (as defined in 'RFC 2254') with optional arguments. In this case, the username is the only argument, denoted by '{0}'. Possible examples are: uid={0}) - this would search for a username match on the uid attribute. Authentication using LDAP is performed from the DN found if successful. Default value is blank/empty.
	// - Note: LDAP settings should provide a userDnPattern or a searchFilter (or both)
	SearchFilter pulumi.StringPtrOutput `pulumi:"searchFilter"`
	// When set, enables deep search through the sub-tree of the LDAP URL + Search Base.  Default value is `true`.
	SearchSubTree pulumi.BoolPtrOutput `pulumi:"searchSubTree"`
	// A DN pattern used to log users directly in to the LDAP database. This pattern is used to create a DN string for "direct" user authentication, and is relative to the base DN in the LDAP URL. The pattern argument {0} is replaced with the username at runtime. This only works if anonymous binding is allowed and a direct user DN can be used (which is not the default case for Active Directory). For example: uid={0},ou=People. Default value is blank/empty.
	// - Note: LDAP settings should provide a userDnPattern or a searchFilter (or both).
	UserDnPattern pulumi.StringPtrOutput `pulumi:"userDnPattern"`
}

// NewLdapSetting registers a new resource with the given unique name, arguments, and options.
func NewLdapSetting(ctx *pulumi.Context,
	name string, args *LdapSettingArgs, opts ...pulumi.ResourceOption) (*LdapSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.LdapUrl == nil {
		return nil, errors.New("invalid value for required argument 'LdapUrl'")
	}
	if args.ManagerPassword != nil {
		args.ManagerPassword = pulumi.ToSecret(args.ManagerPassword).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"managerPassword",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LdapSetting
	err := ctx.RegisterResource("artifactory:index/ldapSetting:LdapSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLdapSetting gets an existing LdapSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLdapSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LdapSettingState, opts ...pulumi.ResourceOption) (*LdapSetting, error) {
	var resource LdapSetting
	err := ctx.ReadResource("artifactory:index/ldapSetting:LdapSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LdapSetting resources.
type ldapSettingState struct {
	// When set, users created after logging in using LDAP will be able to access their profile page.  Default value is `false`.
	AllowUserToAccessProfile *bool `pulumi:"allowUserToAccessProfile"`
	// When set, the system will automatically create new users for those who have logged in using LDAP, and assign them to the default groups.  Default value is `true`.
	AutoCreateUser *bool `pulumi:"autoCreateUser"`
	// An attribute that can be used to map a user's email address to a user created automatically in Artifactory. Default value is `mail`.
	// - Note: If blank/empty string input was set for email_attribute, Default value `mail` takes effect. This is to match with Artifactory behavior.
	EmailAttribute *string `pulumi:"emailAttribute"`
	// When set, these settings are enabled. Default value is `true`.
	Enabled *bool `pulumi:"enabled"`
	// The unique ID of the LDAP setting.
	Key *string `pulumi:"key"`
	// Protects against LDAP poisoning by filtering out users exposed to vulnerabilities.  Default value is `true`.
	LdapPoisoningProtection *bool `pulumi:"ldapPoisoningProtection"`
	// Location of the LDAP server in the following format: ldap://myserver:myport/dc=sampledomain,dc=com. The URL should include the base DN used to search for and/or authenticate users.
	LdapUrl *string `pulumi:"ldapUrl"`
	// The full DN of a user with permissions that allow querying the LDAP server. When working with LDAP Groups, the user should have permissions for any extra group attributes such as memberOf.
	ManagerDn *string `pulumi:"managerDn"`
	// The password of the user binding to the LDAP server when using "search" authentication.
	ManagerPassword *string `pulumi:"managerPassword"`
	// When set, supports paging results for the LDAP server. This feature requires that the LDAP Server supports a PagedResultsControl configuration.  Default value is `true`.
	PagingSupportEnabled *bool `pulumi:"pagingSupportEnabled"`
	// The Context name in which to search relative to the base DN in the LDAP URL. Multiple search bases may be specified separated by a pipe ( | ).
	SearchBase *string `pulumi:"searchBase"`
	// A filter expression used to search for the user DN that is used in LDAP authentication. This is an LDAP search filter (as defined in 'RFC 2254') with optional arguments. In this case, the username is the only argument, denoted by '{0}'. Possible examples are: uid={0}) - this would search for a username match on the uid attribute. Authentication using LDAP is performed from the DN found if successful. Default value is blank/empty.
	// - Note: LDAP settings should provide a userDnPattern or a searchFilter (or both)
	SearchFilter *string `pulumi:"searchFilter"`
	// When set, enables deep search through the sub-tree of the LDAP URL + Search Base.  Default value is `true`.
	SearchSubTree *bool `pulumi:"searchSubTree"`
	// A DN pattern used to log users directly in to the LDAP database. This pattern is used to create a DN string for "direct" user authentication, and is relative to the base DN in the LDAP URL. The pattern argument {0} is replaced with the username at runtime. This only works if anonymous binding is allowed and a direct user DN can be used (which is not the default case for Active Directory). For example: uid={0},ou=People. Default value is blank/empty.
	// - Note: LDAP settings should provide a userDnPattern or a searchFilter (or both).
	UserDnPattern *string `pulumi:"userDnPattern"`
}

type LdapSettingState struct {
	// When set, users created after logging in using LDAP will be able to access their profile page.  Default value is `false`.
	AllowUserToAccessProfile pulumi.BoolPtrInput
	// When set, the system will automatically create new users for those who have logged in using LDAP, and assign them to the default groups.  Default value is `true`.
	AutoCreateUser pulumi.BoolPtrInput
	// An attribute that can be used to map a user's email address to a user created automatically in Artifactory. Default value is `mail`.
	// - Note: If blank/empty string input was set for email_attribute, Default value `mail` takes effect. This is to match with Artifactory behavior.
	EmailAttribute pulumi.StringPtrInput
	// When set, these settings are enabled. Default value is `true`.
	Enabled pulumi.BoolPtrInput
	// The unique ID of the LDAP setting.
	Key pulumi.StringPtrInput
	// Protects against LDAP poisoning by filtering out users exposed to vulnerabilities.  Default value is `true`.
	LdapPoisoningProtection pulumi.BoolPtrInput
	// Location of the LDAP server in the following format: ldap://myserver:myport/dc=sampledomain,dc=com. The URL should include the base DN used to search for and/or authenticate users.
	LdapUrl pulumi.StringPtrInput
	// The full DN of a user with permissions that allow querying the LDAP server. When working with LDAP Groups, the user should have permissions for any extra group attributes such as memberOf.
	ManagerDn pulumi.StringPtrInput
	// The password of the user binding to the LDAP server when using "search" authentication.
	ManagerPassword pulumi.StringPtrInput
	// When set, supports paging results for the LDAP server. This feature requires that the LDAP Server supports a PagedResultsControl configuration.  Default value is `true`.
	PagingSupportEnabled pulumi.BoolPtrInput
	// The Context name in which to search relative to the base DN in the LDAP URL. Multiple search bases may be specified separated by a pipe ( | ).
	SearchBase pulumi.StringPtrInput
	// A filter expression used to search for the user DN that is used in LDAP authentication. This is an LDAP search filter (as defined in 'RFC 2254') with optional arguments. In this case, the username is the only argument, denoted by '{0}'. Possible examples are: uid={0}) - this would search for a username match on the uid attribute. Authentication using LDAP is performed from the DN found if successful. Default value is blank/empty.
	// - Note: LDAP settings should provide a userDnPattern or a searchFilter (or both)
	SearchFilter pulumi.StringPtrInput
	// When set, enables deep search through the sub-tree of the LDAP URL + Search Base.  Default value is `true`.
	SearchSubTree pulumi.BoolPtrInput
	// A DN pattern used to log users directly in to the LDAP database. This pattern is used to create a DN string for "direct" user authentication, and is relative to the base DN in the LDAP URL. The pattern argument {0} is replaced with the username at runtime. This only works if anonymous binding is allowed and a direct user DN can be used (which is not the default case for Active Directory). For example: uid={0},ou=People. Default value is blank/empty.
	// - Note: LDAP settings should provide a userDnPattern or a searchFilter (or both).
	UserDnPattern pulumi.StringPtrInput
}

func (LdapSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*ldapSettingState)(nil)).Elem()
}

type ldapSettingArgs struct {
	// When set, users created after logging in using LDAP will be able to access their profile page.  Default value is `false`.
	AllowUserToAccessProfile *bool `pulumi:"allowUserToAccessProfile"`
	// When set, the system will automatically create new users for those who have logged in using LDAP, and assign them to the default groups.  Default value is `true`.
	AutoCreateUser *bool `pulumi:"autoCreateUser"`
	// An attribute that can be used to map a user's email address to a user created automatically in Artifactory. Default value is `mail`.
	// - Note: If blank/empty string input was set for email_attribute, Default value `mail` takes effect. This is to match with Artifactory behavior.
	EmailAttribute *string `pulumi:"emailAttribute"`
	// When set, these settings are enabled. Default value is `true`.
	Enabled *bool `pulumi:"enabled"`
	// The unique ID of the LDAP setting.
	Key string `pulumi:"key"`
	// Protects against LDAP poisoning by filtering out users exposed to vulnerabilities.  Default value is `true`.
	LdapPoisoningProtection *bool `pulumi:"ldapPoisoningProtection"`
	// Location of the LDAP server in the following format: ldap://myserver:myport/dc=sampledomain,dc=com. The URL should include the base DN used to search for and/or authenticate users.
	LdapUrl string `pulumi:"ldapUrl"`
	// The full DN of a user with permissions that allow querying the LDAP server. When working with LDAP Groups, the user should have permissions for any extra group attributes such as memberOf.
	ManagerDn *string `pulumi:"managerDn"`
	// The password of the user binding to the LDAP server when using "search" authentication.
	ManagerPassword *string `pulumi:"managerPassword"`
	// When set, supports paging results for the LDAP server. This feature requires that the LDAP Server supports a PagedResultsControl configuration.  Default value is `true`.
	PagingSupportEnabled *bool `pulumi:"pagingSupportEnabled"`
	// The Context name in which to search relative to the base DN in the LDAP URL. Multiple search bases may be specified separated by a pipe ( | ).
	SearchBase *string `pulumi:"searchBase"`
	// A filter expression used to search for the user DN that is used in LDAP authentication. This is an LDAP search filter (as defined in 'RFC 2254') with optional arguments. In this case, the username is the only argument, denoted by '{0}'. Possible examples are: uid={0}) - this would search for a username match on the uid attribute. Authentication using LDAP is performed from the DN found if successful. Default value is blank/empty.
	// - Note: LDAP settings should provide a userDnPattern or a searchFilter (or both)
	SearchFilter *string `pulumi:"searchFilter"`
	// When set, enables deep search through the sub-tree of the LDAP URL + Search Base.  Default value is `true`.
	SearchSubTree *bool `pulumi:"searchSubTree"`
	// A DN pattern used to log users directly in to the LDAP database. This pattern is used to create a DN string for "direct" user authentication, and is relative to the base DN in the LDAP URL. The pattern argument {0} is replaced with the username at runtime. This only works if anonymous binding is allowed and a direct user DN can be used (which is not the default case for Active Directory). For example: uid={0},ou=People. Default value is blank/empty.
	// - Note: LDAP settings should provide a userDnPattern or a searchFilter (or both).
	UserDnPattern *string `pulumi:"userDnPattern"`
}

// The set of arguments for constructing a LdapSetting resource.
type LdapSettingArgs struct {
	// When set, users created after logging in using LDAP will be able to access their profile page.  Default value is `false`.
	AllowUserToAccessProfile pulumi.BoolPtrInput
	// When set, the system will automatically create new users for those who have logged in using LDAP, and assign them to the default groups.  Default value is `true`.
	AutoCreateUser pulumi.BoolPtrInput
	// An attribute that can be used to map a user's email address to a user created automatically in Artifactory. Default value is `mail`.
	// - Note: If blank/empty string input was set for email_attribute, Default value `mail` takes effect. This is to match with Artifactory behavior.
	EmailAttribute pulumi.StringPtrInput
	// When set, these settings are enabled. Default value is `true`.
	Enabled pulumi.BoolPtrInput
	// The unique ID of the LDAP setting.
	Key pulumi.StringInput
	// Protects against LDAP poisoning by filtering out users exposed to vulnerabilities.  Default value is `true`.
	LdapPoisoningProtection pulumi.BoolPtrInput
	// Location of the LDAP server in the following format: ldap://myserver:myport/dc=sampledomain,dc=com. The URL should include the base DN used to search for and/or authenticate users.
	LdapUrl pulumi.StringInput
	// The full DN of a user with permissions that allow querying the LDAP server. When working with LDAP Groups, the user should have permissions for any extra group attributes such as memberOf.
	ManagerDn pulumi.StringPtrInput
	// The password of the user binding to the LDAP server when using "search" authentication.
	ManagerPassword pulumi.StringPtrInput
	// When set, supports paging results for the LDAP server. This feature requires that the LDAP Server supports a PagedResultsControl configuration.  Default value is `true`.
	PagingSupportEnabled pulumi.BoolPtrInput
	// The Context name in which to search relative to the base DN in the LDAP URL. Multiple search bases may be specified separated by a pipe ( | ).
	SearchBase pulumi.StringPtrInput
	// A filter expression used to search for the user DN that is used in LDAP authentication. This is an LDAP search filter (as defined in 'RFC 2254') with optional arguments. In this case, the username is the only argument, denoted by '{0}'. Possible examples are: uid={0}) - this would search for a username match on the uid attribute. Authentication using LDAP is performed from the DN found if successful. Default value is blank/empty.
	// - Note: LDAP settings should provide a userDnPattern or a searchFilter (or both)
	SearchFilter pulumi.StringPtrInput
	// When set, enables deep search through the sub-tree of the LDAP URL + Search Base.  Default value is `true`.
	SearchSubTree pulumi.BoolPtrInput
	// A DN pattern used to log users directly in to the LDAP database. This pattern is used to create a DN string for "direct" user authentication, and is relative to the base DN in the LDAP URL. The pattern argument {0} is replaced with the username at runtime. This only works if anonymous binding is allowed and a direct user DN can be used (which is not the default case for Active Directory). For example: uid={0},ou=People. Default value is blank/empty.
	// - Note: LDAP settings should provide a userDnPattern or a searchFilter (or both).
	UserDnPattern pulumi.StringPtrInput
}

func (LdapSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ldapSettingArgs)(nil)).Elem()
}

type LdapSettingInput interface {
	pulumi.Input

	ToLdapSettingOutput() LdapSettingOutput
	ToLdapSettingOutputWithContext(ctx context.Context) LdapSettingOutput
}

func (*LdapSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**LdapSetting)(nil)).Elem()
}

func (i *LdapSetting) ToLdapSettingOutput() LdapSettingOutput {
	return i.ToLdapSettingOutputWithContext(context.Background())
}

func (i *LdapSetting) ToLdapSettingOutputWithContext(ctx context.Context) LdapSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LdapSettingOutput)
}

// LdapSettingArrayInput is an input type that accepts LdapSettingArray and LdapSettingArrayOutput values.
// You can construct a concrete instance of `LdapSettingArrayInput` via:
//
//	LdapSettingArray{ LdapSettingArgs{...} }
type LdapSettingArrayInput interface {
	pulumi.Input

	ToLdapSettingArrayOutput() LdapSettingArrayOutput
	ToLdapSettingArrayOutputWithContext(context.Context) LdapSettingArrayOutput
}

type LdapSettingArray []LdapSettingInput

func (LdapSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LdapSetting)(nil)).Elem()
}

func (i LdapSettingArray) ToLdapSettingArrayOutput() LdapSettingArrayOutput {
	return i.ToLdapSettingArrayOutputWithContext(context.Background())
}

func (i LdapSettingArray) ToLdapSettingArrayOutputWithContext(ctx context.Context) LdapSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LdapSettingArrayOutput)
}

// LdapSettingMapInput is an input type that accepts LdapSettingMap and LdapSettingMapOutput values.
// You can construct a concrete instance of `LdapSettingMapInput` via:
//
//	LdapSettingMap{ "key": LdapSettingArgs{...} }
type LdapSettingMapInput interface {
	pulumi.Input

	ToLdapSettingMapOutput() LdapSettingMapOutput
	ToLdapSettingMapOutputWithContext(context.Context) LdapSettingMapOutput
}

type LdapSettingMap map[string]LdapSettingInput

func (LdapSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LdapSetting)(nil)).Elem()
}

func (i LdapSettingMap) ToLdapSettingMapOutput() LdapSettingMapOutput {
	return i.ToLdapSettingMapOutputWithContext(context.Background())
}

func (i LdapSettingMap) ToLdapSettingMapOutputWithContext(ctx context.Context) LdapSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LdapSettingMapOutput)
}

type LdapSettingOutput struct{ *pulumi.OutputState }

func (LdapSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LdapSetting)(nil)).Elem()
}

func (o LdapSettingOutput) ToLdapSettingOutput() LdapSettingOutput {
	return o
}

func (o LdapSettingOutput) ToLdapSettingOutputWithContext(ctx context.Context) LdapSettingOutput {
	return o
}

// When set, users created after logging in using LDAP will be able to access their profile page.  Default value is `false`.
func (o LdapSettingOutput) AllowUserToAccessProfile() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LdapSetting) pulumi.BoolPtrOutput { return v.AllowUserToAccessProfile }).(pulumi.BoolPtrOutput)
}

// When set, the system will automatically create new users for those who have logged in using LDAP, and assign them to the default groups.  Default value is `true`.
func (o LdapSettingOutput) AutoCreateUser() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LdapSetting) pulumi.BoolPtrOutput { return v.AutoCreateUser }).(pulumi.BoolPtrOutput)
}

// An attribute that can be used to map a user's email address to a user created automatically in Artifactory. Default value is `mail`.
// - Note: If blank/empty string input was set for email_attribute, Default value `mail` takes effect. This is to match with Artifactory behavior.
func (o LdapSettingOutput) EmailAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LdapSetting) pulumi.StringPtrOutput { return v.EmailAttribute }).(pulumi.StringPtrOutput)
}

// When set, these settings are enabled. Default value is `true`.
func (o LdapSettingOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LdapSetting) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The unique ID of the LDAP setting.
func (o LdapSettingOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *LdapSetting) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Protects against LDAP poisoning by filtering out users exposed to vulnerabilities.  Default value is `true`.
func (o LdapSettingOutput) LdapPoisoningProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LdapSetting) pulumi.BoolPtrOutput { return v.LdapPoisoningProtection }).(pulumi.BoolPtrOutput)
}

// Location of the LDAP server in the following format: ldap://myserver:myport/dc=sampledomain,dc=com. The URL should include the base DN used to search for and/or authenticate users.
func (o LdapSettingOutput) LdapUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *LdapSetting) pulumi.StringOutput { return v.LdapUrl }).(pulumi.StringOutput)
}

// The full DN of a user with permissions that allow querying the LDAP server. When working with LDAP Groups, the user should have permissions for any extra group attributes such as memberOf.
func (o LdapSettingOutput) ManagerDn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LdapSetting) pulumi.StringPtrOutput { return v.ManagerDn }).(pulumi.StringPtrOutput)
}

// The password of the user binding to the LDAP server when using "search" authentication.
func (o LdapSettingOutput) ManagerPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *LdapSetting) pulumi.StringOutput { return v.ManagerPassword }).(pulumi.StringOutput)
}

// When set, supports paging results for the LDAP server. This feature requires that the LDAP Server supports a PagedResultsControl configuration.  Default value is `true`.
func (o LdapSettingOutput) PagingSupportEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LdapSetting) pulumi.BoolPtrOutput { return v.PagingSupportEnabled }).(pulumi.BoolPtrOutput)
}

// The Context name in which to search relative to the base DN in the LDAP URL. Multiple search bases may be specified separated by a pipe ( | ).
func (o LdapSettingOutput) SearchBase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LdapSetting) pulumi.StringPtrOutput { return v.SearchBase }).(pulumi.StringPtrOutput)
}

// A filter expression used to search for the user DN that is used in LDAP authentication. This is an LDAP search filter (as defined in 'RFC 2254') with optional arguments. In this case, the username is the only argument, denoted by '{0}'. Possible examples are: uid={0}) - this would search for a username match on the uid attribute. Authentication using LDAP is performed from the DN found if successful. Default value is blank/empty.
// - Note: LDAP settings should provide a userDnPattern or a searchFilter (or both)
func (o LdapSettingOutput) SearchFilter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LdapSetting) pulumi.StringPtrOutput { return v.SearchFilter }).(pulumi.StringPtrOutput)
}

// When set, enables deep search through the sub-tree of the LDAP URL + Search Base.  Default value is `true`.
func (o LdapSettingOutput) SearchSubTree() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LdapSetting) pulumi.BoolPtrOutput { return v.SearchSubTree }).(pulumi.BoolPtrOutput)
}

// A DN pattern used to log users directly in to the LDAP database. This pattern is used to create a DN string for "direct" user authentication, and is relative to the base DN in the LDAP URL. The pattern argument {0} is replaced with the username at runtime. This only works if anonymous binding is allowed and a direct user DN can be used (which is not the default case for Active Directory). For example: uid={0},ou=People. Default value is blank/empty.
// - Note: LDAP settings should provide a userDnPattern or a searchFilter (or both).
func (o LdapSettingOutput) UserDnPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LdapSetting) pulumi.StringPtrOutput { return v.UserDnPattern }).(pulumi.StringPtrOutput)
}

type LdapSettingArrayOutput struct{ *pulumi.OutputState }

func (LdapSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LdapSetting)(nil)).Elem()
}

func (o LdapSettingArrayOutput) ToLdapSettingArrayOutput() LdapSettingArrayOutput {
	return o
}

func (o LdapSettingArrayOutput) ToLdapSettingArrayOutputWithContext(ctx context.Context) LdapSettingArrayOutput {
	return o
}

func (o LdapSettingArrayOutput) Index(i pulumi.IntInput) LdapSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LdapSetting {
		return vs[0].([]*LdapSetting)[vs[1].(int)]
	}).(LdapSettingOutput)
}

type LdapSettingMapOutput struct{ *pulumi.OutputState }

func (LdapSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LdapSetting)(nil)).Elem()
}

func (o LdapSettingMapOutput) ToLdapSettingMapOutput() LdapSettingMapOutput {
	return o
}

func (o LdapSettingMapOutput) ToLdapSettingMapOutputWithContext(ctx context.Context) LdapSettingMapOutput {
	return o
}

func (o LdapSettingMapOutput) MapIndex(k pulumi.StringInput) LdapSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LdapSetting {
		return vs[0].(map[string]*LdapSetting)[vs[1].(string)]
	}).(LdapSettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LdapSettingInput)(nil)).Elem(), &LdapSetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*LdapSettingArrayInput)(nil)).Elem(), LdapSettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LdapSettingMapInput)(nil)).Elem(), LdapSettingMap{})
	pulumi.RegisterOutputType(LdapSettingOutput{})
	pulumi.RegisterOutputType(LdapSettingArrayOutput{})
	pulumi.RegisterOutputType(LdapSettingMapOutput{})
}
