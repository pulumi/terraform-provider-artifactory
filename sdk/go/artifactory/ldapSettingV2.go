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

// Provides an Artifactory LDAP Setting resource.
//
// This resource can be used to manage Artifactory's LDAP settings for user authentication.
//
// When specified LDAP setting is active, Artifactory first attempts to authenticate the user against the LDAP server.
// If LDAP authentication fails, it then tries to authenticate via its internal database.
//
// [API documentation](https://jfrog.com/help/r/jfrog-rest-apis/ldap-setting), [general documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/ldap).
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
//			_, err := artifactory.NewLdapSettingV2(ctx, "ldap_name", &artifactory.LdapSettingV2Args{
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
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// ```sh
// $ pulumi import artifactory:index/ldapSettingV2:LdapSettingV2 ldap ldap1
// ```
type LdapSettingV2 struct {
	pulumi.CustomResourceState

	// Auto created users will have access to their profile page and will be able to perform actions such as generating an API key. Default value is `false`.
	AllowUserToAccessProfile pulumi.BoolOutput `pulumi:"allowUserToAccessProfile"`
	// When set, users are automatically created when using LDAP. Otherwise, users are transient and associated with auto-join groups defined in Artifactory. Default value is `true`.
	AutoCreateUser pulumi.BoolOutput `pulumi:"autoCreateUser"`
	// An attribute that can be used to map a user's email address to a user created automatically in Artifactory. Default value is`mail`.
	EmailAttribute pulumi.StringOutput `pulumi:"emailAttribute"`
	// Flag to enable or disable the ldap setting. Default value is `true`.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Ldap setting name.
	Key pulumi.StringOutput `pulumi:"key"`
	// When this is set to `true`, an empty or missing usernames array will detach all users from the group.
	LdapPoisoningProtection pulumi.BoolOutput `pulumi:"ldapPoisoningProtection"`
	// Location of the LDAP server in the following format: `ldap://myldapserver/dc=sampledomain,dc=com`
	LdapUrl pulumi.StringOutput `pulumi:"ldapUrl"`
	// The full DN of the user that binds to the LDAP server to perform user searches. Only used with `search` authentication.
	ManagerDn pulumi.StringOutput `pulumi:"managerDn"`
	// The password of the user that binds to the LDAP server to perform the search. Only used with `search` authentication.
	ManagerPassword pulumi.StringOutput `pulumi:"managerPassword"`
	// When set, supports paging results for the LDAP server. This feature requires that the LDAP server supports a PagedResultsControl configuration. Default value is `true`.
	PagingSupportEnabled pulumi.BoolOutput `pulumi:"pagingSupportEnabled"`
	// A context name to search in relative to the base DN of the LDAP URL. For example, 'ou=users' With the LDAP Group Add-on enabled, it is possible to enter multiple search base entries separated by a pipe ('|') character.
	SearchBase pulumi.StringOutput `pulumi:"searchBase"`
	// A filter expression used to search for the user DN used in LDAP authentication. This is an LDAP search filter (as defined in 'RFC 2254') with optional arguments. In this case, the username is the only argument, and is denoted by '{0}'. Possible examples are: (uid={0}) - This searches for a username match on the attribute. Authentication to LDAP is performed from the DN found if successful.
	SearchFilter pulumi.StringOutput `pulumi:"searchFilter"`
	// When set, enables deep search through the sub tree of the LDAP URL + search base. Default value is `true`.
	SearchSubTree pulumi.BoolOutput `pulumi:"searchSubTree"`
	// A DN pattern that can be used to log users directly in to LDAP. This pattern is used to create a DN string for 'direct' user authentication where the pattern is relative to the base DN in the LDAP URL. The pattern argument {0} is replaced with the username. This only works if anonymous binding is allowed and a direct user DN can be used, which is not the default case for Active Directory (use User DN search filter instead). Example: uid={0},ou=People. Default value is blank/empty.
	UserDnPattern pulumi.StringOutput `pulumi:"userDnPattern"`
}

// NewLdapSettingV2 registers a new resource with the given unique name, arguments, and options.
func NewLdapSettingV2(ctx *pulumi.Context,
	name string, args *LdapSettingV2Args, opts ...pulumi.ResourceOption) (*LdapSettingV2, error) {
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
	var resource LdapSettingV2
	err := ctx.RegisterResource("artifactory:index/ldapSettingV2:LdapSettingV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLdapSettingV2 gets an existing LdapSettingV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLdapSettingV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LdapSettingV2State, opts ...pulumi.ResourceOption) (*LdapSettingV2, error) {
	var resource LdapSettingV2
	err := ctx.ReadResource("artifactory:index/ldapSettingV2:LdapSettingV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LdapSettingV2 resources.
type ldapSettingV2State struct {
	// Auto created users will have access to their profile page and will be able to perform actions such as generating an API key. Default value is `false`.
	AllowUserToAccessProfile *bool `pulumi:"allowUserToAccessProfile"`
	// When set, users are automatically created when using LDAP. Otherwise, users are transient and associated with auto-join groups defined in Artifactory. Default value is `true`.
	AutoCreateUser *bool `pulumi:"autoCreateUser"`
	// An attribute that can be used to map a user's email address to a user created automatically in Artifactory. Default value is`mail`.
	EmailAttribute *string `pulumi:"emailAttribute"`
	// Flag to enable or disable the ldap setting. Default value is `true`.
	Enabled *bool `pulumi:"enabled"`
	// Ldap setting name.
	Key *string `pulumi:"key"`
	// When this is set to `true`, an empty or missing usernames array will detach all users from the group.
	LdapPoisoningProtection *bool `pulumi:"ldapPoisoningProtection"`
	// Location of the LDAP server in the following format: `ldap://myldapserver/dc=sampledomain,dc=com`
	LdapUrl *string `pulumi:"ldapUrl"`
	// The full DN of the user that binds to the LDAP server to perform user searches. Only used with `search` authentication.
	ManagerDn *string `pulumi:"managerDn"`
	// The password of the user that binds to the LDAP server to perform the search. Only used with `search` authentication.
	ManagerPassword *string `pulumi:"managerPassword"`
	// When set, supports paging results for the LDAP server. This feature requires that the LDAP server supports a PagedResultsControl configuration. Default value is `true`.
	PagingSupportEnabled *bool `pulumi:"pagingSupportEnabled"`
	// A context name to search in relative to the base DN of the LDAP URL. For example, 'ou=users' With the LDAP Group Add-on enabled, it is possible to enter multiple search base entries separated by a pipe ('|') character.
	SearchBase *string `pulumi:"searchBase"`
	// A filter expression used to search for the user DN used in LDAP authentication. This is an LDAP search filter (as defined in 'RFC 2254') with optional arguments. In this case, the username is the only argument, and is denoted by '{0}'. Possible examples are: (uid={0}) - This searches for a username match on the attribute. Authentication to LDAP is performed from the DN found if successful.
	SearchFilter *string `pulumi:"searchFilter"`
	// When set, enables deep search through the sub tree of the LDAP URL + search base. Default value is `true`.
	SearchSubTree *bool `pulumi:"searchSubTree"`
	// A DN pattern that can be used to log users directly in to LDAP. This pattern is used to create a DN string for 'direct' user authentication where the pattern is relative to the base DN in the LDAP URL. The pattern argument {0} is replaced with the username. This only works if anonymous binding is allowed and a direct user DN can be used, which is not the default case for Active Directory (use User DN search filter instead). Example: uid={0},ou=People. Default value is blank/empty.
	UserDnPattern *string `pulumi:"userDnPattern"`
}

type LdapSettingV2State struct {
	// Auto created users will have access to their profile page and will be able to perform actions such as generating an API key. Default value is `false`.
	AllowUserToAccessProfile pulumi.BoolPtrInput
	// When set, users are automatically created when using LDAP. Otherwise, users are transient and associated with auto-join groups defined in Artifactory. Default value is `true`.
	AutoCreateUser pulumi.BoolPtrInput
	// An attribute that can be used to map a user's email address to a user created automatically in Artifactory. Default value is`mail`.
	EmailAttribute pulumi.StringPtrInput
	// Flag to enable or disable the ldap setting. Default value is `true`.
	Enabled pulumi.BoolPtrInput
	// Ldap setting name.
	Key pulumi.StringPtrInput
	// When this is set to `true`, an empty or missing usernames array will detach all users from the group.
	LdapPoisoningProtection pulumi.BoolPtrInput
	// Location of the LDAP server in the following format: `ldap://myldapserver/dc=sampledomain,dc=com`
	LdapUrl pulumi.StringPtrInput
	// The full DN of the user that binds to the LDAP server to perform user searches. Only used with `search` authentication.
	ManagerDn pulumi.StringPtrInput
	// The password of the user that binds to the LDAP server to perform the search. Only used with `search` authentication.
	ManagerPassword pulumi.StringPtrInput
	// When set, supports paging results for the LDAP server. This feature requires that the LDAP server supports a PagedResultsControl configuration. Default value is `true`.
	PagingSupportEnabled pulumi.BoolPtrInput
	// A context name to search in relative to the base DN of the LDAP URL. For example, 'ou=users' With the LDAP Group Add-on enabled, it is possible to enter multiple search base entries separated by a pipe ('|') character.
	SearchBase pulumi.StringPtrInput
	// A filter expression used to search for the user DN used in LDAP authentication. This is an LDAP search filter (as defined in 'RFC 2254') with optional arguments. In this case, the username is the only argument, and is denoted by '{0}'. Possible examples are: (uid={0}) - This searches for a username match on the attribute. Authentication to LDAP is performed from the DN found if successful.
	SearchFilter pulumi.StringPtrInput
	// When set, enables deep search through the sub tree of the LDAP URL + search base. Default value is `true`.
	SearchSubTree pulumi.BoolPtrInput
	// A DN pattern that can be used to log users directly in to LDAP. This pattern is used to create a DN string for 'direct' user authentication where the pattern is relative to the base DN in the LDAP URL. The pattern argument {0} is replaced with the username. This only works if anonymous binding is allowed and a direct user DN can be used, which is not the default case for Active Directory (use User DN search filter instead). Example: uid={0},ou=People. Default value is blank/empty.
	UserDnPattern pulumi.StringPtrInput
}

func (LdapSettingV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*ldapSettingV2State)(nil)).Elem()
}

type ldapSettingV2Args struct {
	// Auto created users will have access to their profile page and will be able to perform actions such as generating an API key. Default value is `false`.
	AllowUserToAccessProfile *bool `pulumi:"allowUserToAccessProfile"`
	// When set, users are automatically created when using LDAP. Otherwise, users are transient and associated with auto-join groups defined in Artifactory. Default value is `true`.
	AutoCreateUser *bool `pulumi:"autoCreateUser"`
	// An attribute that can be used to map a user's email address to a user created automatically in Artifactory. Default value is`mail`.
	EmailAttribute *string `pulumi:"emailAttribute"`
	// Flag to enable or disable the ldap setting. Default value is `true`.
	Enabled *bool `pulumi:"enabled"`
	// Ldap setting name.
	Key string `pulumi:"key"`
	// When this is set to `true`, an empty or missing usernames array will detach all users from the group.
	LdapPoisoningProtection *bool `pulumi:"ldapPoisoningProtection"`
	// Location of the LDAP server in the following format: `ldap://myldapserver/dc=sampledomain,dc=com`
	LdapUrl string `pulumi:"ldapUrl"`
	// The full DN of the user that binds to the LDAP server to perform user searches. Only used with `search` authentication.
	ManagerDn *string `pulumi:"managerDn"`
	// The password of the user that binds to the LDAP server to perform the search. Only used with `search` authentication.
	ManagerPassword *string `pulumi:"managerPassword"`
	// When set, supports paging results for the LDAP server. This feature requires that the LDAP server supports a PagedResultsControl configuration. Default value is `true`.
	PagingSupportEnabled *bool `pulumi:"pagingSupportEnabled"`
	// A context name to search in relative to the base DN of the LDAP URL. For example, 'ou=users' With the LDAP Group Add-on enabled, it is possible to enter multiple search base entries separated by a pipe ('|') character.
	SearchBase *string `pulumi:"searchBase"`
	// A filter expression used to search for the user DN used in LDAP authentication. This is an LDAP search filter (as defined in 'RFC 2254') with optional arguments. In this case, the username is the only argument, and is denoted by '{0}'. Possible examples are: (uid={0}) - This searches for a username match on the attribute. Authentication to LDAP is performed from the DN found if successful.
	SearchFilter *string `pulumi:"searchFilter"`
	// When set, enables deep search through the sub tree of the LDAP URL + search base. Default value is `true`.
	SearchSubTree *bool `pulumi:"searchSubTree"`
	// A DN pattern that can be used to log users directly in to LDAP. This pattern is used to create a DN string for 'direct' user authentication where the pattern is relative to the base DN in the LDAP URL. The pattern argument {0} is replaced with the username. This only works if anonymous binding is allowed and a direct user DN can be used, which is not the default case for Active Directory (use User DN search filter instead). Example: uid={0},ou=People. Default value is blank/empty.
	UserDnPattern *string `pulumi:"userDnPattern"`
}

// The set of arguments for constructing a LdapSettingV2 resource.
type LdapSettingV2Args struct {
	// Auto created users will have access to their profile page and will be able to perform actions such as generating an API key. Default value is `false`.
	AllowUserToAccessProfile pulumi.BoolPtrInput
	// When set, users are automatically created when using LDAP. Otherwise, users are transient and associated with auto-join groups defined in Artifactory. Default value is `true`.
	AutoCreateUser pulumi.BoolPtrInput
	// An attribute that can be used to map a user's email address to a user created automatically in Artifactory. Default value is`mail`.
	EmailAttribute pulumi.StringPtrInput
	// Flag to enable or disable the ldap setting. Default value is `true`.
	Enabled pulumi.BoolPtrInput
	// Ldap setting name.
	Key pulumi.StringInput
	// When this is set to `true`, an empty or missing usernames array will detach all users from the group.
	LdapPoisoningProtection pulumi.BoolPtrInput
	// Location of the LDAP server in the following format: `ldap://myldapserver/dc=sampledomain,dc=com`
	LdapUrl pulumi.StringInput
	// The full DN of the user that binds to the LDAP server to perform user searches. Only used with `search` authentication.
	ManagerDn pulumi.StringPtrInput
	// The password of the user that binds to the LDAP server to perform the search. Only used with `search` authentication.
	ManagerPassword pulumi.StringPtrInput
	// When set, supports paging results for the LDAP server. This feature requires that the LDAP server supports a PagedResultsControl configuration. Default value is `true`.
	PagingSupportEnabled pulumi.BoolPtrInput
	// A context name to search in relative to the base DN of the LDAP URL. For example, 'ou=users' With the LDAP Group Add-on enabled, it is possible to enter multiple search base entries separated by a pipe ('|') character.
	SearchBase pulumi.StringPtrInput
	// A filter expression used to search for the user DN used in LDAP authentication. This is an LDAP search filter (as defined in 'RFC 2254') with optional arguments. In this case, the username is the only argument, and is denoted by '{0}'. Possible examples are: (uid={0}) - This searches for a username match on the attribute. Authentication to LDAP is performed from the DN found if successful.
	SearchFilter pulumi.StringPtrInput
	// When set, enables deep search through the sub tree of the LDAP URL + search base. Default value is `true`.
	SearchSubTree pulumi.BoolPtrInput
	// A DN pattern that can be used to log users directly in to LDAP. This pattern is used to create a DN string for 'direct' user authentication where the pattern is relative to the base DN in the LDAP URL. The pattern argument {0} is replaced with the username. This only works if anonymous binding is allowed and a direct user DN can be used, which is not the default case for Active Directory (use User DN search filter instead). Example: uid={0},ou=People. Default value is blank/empty.
	UserDnPattern pulumi.StringPtrInput
}

func (LdapSettingV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*ldapSettingV2Args)(nil)).Elem()
}

type LdapSettingV2Input interface {
	pulumi.Input

	ToLdapSettingV2Output() LdapSettingV2Output
	ToLdapSettingV2OutputWithContext(ctx context.Context) LdapSettingV2Output
}

func (*LdapSettingV2) ElementType() reflect.Type {
	return reflect.TypeOf((**LdapSettingV2)(nil)).Elem()
}

func (i *LdapSettingV2) ToLdapSettingV2Output() LdapSettingV2Output {
	return i.ToLdapSettingV2OutputWithContext(context.Background())
}

func (i *LdapSettingV2) ToLdapSettingV2OutputWithContext(ctx context.Context) LdapSettingV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(LdapSettingV2Output)
}

// LdapSettingV2ArrayInput is an input type that accepts LdapSettingV2Array and LdapSettingV2ArrayOutput values.
// You can construct a concrete instance of `LdapSettingV2ArrayInput` via:
//
//	LdapSettingV2Array{ LdapSettingV2Args{...} }
type LdapSettingV2ArrayInput interface {
	pulumi.Input

	ToLdapSettingV2ArrayOutput() LdapSettingV2ArrayOutput
	ToLdapSettingV2ArrayOutputWithContext(context.Context) LdapSettingV2ArrayOutput
}

type LdapSettingV2Array []LdapSettingV2Input

func (LdapSettingV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LdapSettingV2)(nil)).Elem()
}

func (i LdapSettingV2Array) ToLdapSettingV2ArrayOutput() LdapSettingV2ArrayOutput {
	return i.ToLdapSettingV2ArrayOutputWithContext(context.Background())
}

func (i LdapSettingV2Array) ToLdapSettingV2ArrayOutputWithContext(ctx context.Context) LdapSettingV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LdapSettingV2ArrayOutput)
}

// LdapSettingV2MapInput is an input type that accepts LdapSettingV2Map and LdapSettingV2MapOutput values.
// You can construct a concrete instance of `LdapSettingV2MapInput` via:
//
//	LdapSettingV2Map{ "key": LdapSettingV2Args{...} }
type LdapSettingV2MapInput interface {
	pulumi.Input

	ToLdapSettingV2MapOutput() LdapSettingV2MapOutput
	ToLdapSettingV2MapOutputWithContext(context.Context) LdapSettingV2MapOutput
}

type LdapSettingV2Map map[string]LdapSettingV2Input

func (LdapSettingV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LdapSettingV2)(nil)).Elem()
}

func (i LdapSettingV2Map) ToLdapSettingV2MapOutput() LdapSettingV2MapOutput {
	return i.ToLdapSettingV2MapOutputWithContext(context.Background())
}

func (i LdapSettingV2Map) ToLdapSettingV2MapOutputWithContext(ctx context.Context) LdapSettingV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LdapSettingV2MapOutput)
}

type LdapSettingV2Output struct{ *pulumi.OutputState }

func (LdapSettingV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**LdapSettingV2)(nil)).Elem()
}

func (o LdapSettingV2Output) ToLdapSettingV2Output() LdapSettingV2Output {
	return o
}

func (o LdapSettingV2Output) ToLdapSettingV2OutputWithContext(ctx context.Context) LdapSettingV2Output {
	return o
}

// Auto created users will have access to their profile page and will be able to perform actions such as generating an API key. Default value is `false`.
func (o LdapSettingV2Output) AllowUserToAccessProfile() pulumi.BoolOutput {
	return o.ApplyT(func(v *LdapSettingV2) pulumi.BoolOutput { return v.AllowUserToAccessProfile }).(pulumi.BoolOutput)
}

// When set, users are automatically created when using LDAP. Otherwise, users are transient and associated with auto-join groups defined in Artifactory. Default value is `true`.
func (o LdapSettingV2Output) AutoCreateUser() pulumi.BoolOutput {
	return o.ApplyT(func(v *LdapSettingV2) pulumi.BoolOutput { return v.AutoCreateUser }).(pulumi.BoolOutput)
}

// An attribute that can be used to map a user's email address to a user created automatically in Artifactory. Default value is`mail`.
func (o LdapSettingV2Output) EmailAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v *LdapSettingV2) pulumi.StringOutput { return v.EmailAttribute }).(pulumi.StringOutput)
}

// Flag to enable or disable the ldap setting. Default value is `true`.
func (o LdapSettingV2Output) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *LdapSettingV2) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// Ldap setting name.
func (o LdapSettingV2Output) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *LdapSettingV2) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// When this is set to `true`, an empty or missing usernames array will detach all users from the group.
func (o LdapSettingV2Output) LdapPoisoningProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v *LdapSettingV2) pulumi.BoolOutput { return v.LdapPoisoningProtection }).(pulumi.BoolOutput)
}

// Location of the LDAP server in the following format: `ldap://myldapserver/dc=sampledomain,dc=com`
func (o LdapSettingV2Output) LdapUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *LdapSettingV2) pulumi.StringOutput { return v.LdapUrl }).(pulumi.StringOutput)
}

// The full DN of the user that binds to the LDAP server to perform user searches. Only used with `search` authentication.
func (o LdapSettingV2Output) ManagerDn() pulumi.StringOutput {
	return o.ApplyT(func(v *LdapSettingV2) pulumi.StringOutput { return v.ManagerDn }).(pulumi.StringOutput)
}

// The password of the user that binds to the LDAP server to perform the search. Only used with `search` authentication.
func (o LdapSettingV2Output) ManagerPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *LdapSettingV2) pulumi.StringOutput { return v.ManagerPassword }).(pulumi.StringOutput)
}

// When set, supports paging results for the LDAP server. This feature requires that the LDAP server supports a PagedResultsControl configuration. Default value is `true`.
func (o LdapSettingV2Output) PagingSupportEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *LdapSettingV2) pulumi.BoolOutput { return v.PagingSupportEnabled }).(pulumi.BoolOutput)
}

// A context name to search in relative to the base DN of the LDAP URL. For example, 'ou=users' With the LDAP Group Add-on enabled, it is possible to enter multiple search base entries separated by a pipe ('|') character.
func (o LdapSettingV2Output) SearchBase() pulumi.StringOutput {
	return o.ApplyT(func(v *LdapSettingV2) pulumi.StringOutput { return v.SearchBase }).(pulumi.StringOutput)
}

// A filter expression used to search for the user DN used in LDAP authentication. This is an LDAP search filter (as defined in 'RFC 2254') with optional arguments. In this case, the username is the only argument, and is denoted by '{0}'. Possible examples are: (uid={0}) - This searches for a username match on the attribute. Authentication to LDAP is performed from the DN found if successful.
func (o LdapSettingV2Output) SearchFilter() pulumi.StringOutput {
	return o.ApplyT(func(v *LdapSettingV2) pulumi.StringOutput { return v.SearchFilter }).(pulumi.StringOutput)
}

// When set, enables deep search through the sub tree of the LDAP URL + search base. Default value is `true`.
func (o LdapSettingV2Output) SearchSubTree() pulumi.BoolOutput {
	return o.ApplyT(func(v *LdapSettingV2) pulumi.BoolOutput { return v.SearchSubTree }).(pulumi.BoolOutput)
}

// A DN pattern that can be used to log users directly in to LDAP. This pattern is used to create a DN string for 'direct' user authentication where the pattern is relative to the base DN in the LDAP URL. The pattern argument {0} is replaced with the username. This only works if anonymous binding is allowed and a direct user DN can be used, which is not the default case for Active Directory (use User DN search filter instead). Example: uid={0},ou=People. Default value is blank/empty.
func (o LdapSettingV2Output) UserDnPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LdapSettingV2) pulumi.StringOutput { return v.UserDnPattern }).(pulumi.StringOutput)
}

type LdapSettingV2ArrayOutput struct{ *pulumi.OutputState }

func (LdapSettingV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LdapSettingV2)(nil)).Elem()
}

func (o LdapSettingV2ArrayOutput) ToLdapSettingV2ArrayOutput() LdapSettingV2ArrayOutput {
	return o
}

func (o LdapSettingV2ArrayOutput) ToLdapSettingV2ArrayOutputWithContext(ctx context.Context) LdapSettingV2ArrayOutput {
	return o
}

func (o LdapSettingV2ArrayOutput) Index(i pulumi.IntInput) LdapSettingV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LdapSettingV2 {
		return vs[0].([]*LdapSettingV2)[vs[1].(int)]
	}).(LdapSettingV2Output)
}

type LdapSettingV2MapOutput struct{ *pulumi.OutputState }

func (LdapSettingV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LdapSettingV2)(nil)).Elem()
}

func (o LdapSettingV2MapOutput) ToLdapSettingV2MapOutput() LdapSettingV2MapOutput {
	return o
}

func (o LdapSettingV2MapOutput) ToLdapSettingV2MapOutputWithContext(ctx context.Context) LdapSettingV2MapOutput {
	return o
}

func (o LdapSettingV2MapOutput) MapIndex(k pulumi.StringInput) LdapSettingV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LdapSettingV2 {
		return vs[0].(map[string]*LdapSettingV2)[vs[1].(string)]
	}).(LdapSettingV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LdapSettingV2Input)(nil)).Elem(), &LdapSettingV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*LdapSettingV2ArrayInput)(nil)).Elem(), LdapSettingV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*LdapSettingV2MapInput)(nil)).Elem(), LdapSettingV2Map{})
	pulumi.RegisterOutputType(LdapSettingV2Output{})
	pulumi.RegisterOutputType(LdapSettingV2ArrayOutput{})
	pulumi.RegisterOutputType(LdapSettingV2MapOutput{})
}
