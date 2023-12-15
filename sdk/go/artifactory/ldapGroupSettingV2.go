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
// This resource can be used to manage Artifactory's LDAP Group settings for user authentication.
//
// LDAP Groups Add-on allows you to synchronize your LDAP groups with the system and leverage your existing organizational
// structure for managing group-based permissions.
//
// [API documentation](https://jfrog.com/help/r/jfrog-rest-apis/ldap-group-setting), [general documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/ldap).
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
//			_, err := artifactory.NewLdapGroupSettingV2(ctx, "ldapGroupName", &artifactory.LdapGroupSettingV2Args{
//				DescriptionAttribute: pulumi.String("description"),
//				EnabledLdap:          pulumi.String("ldap_name"),
//				Filter:               pulumi.String("(objectClass=groupOfNames)"),
//				ForceAttributeSearch: pulumi.Bool(false),
//				GroupBaseDn:          pulumi.String("CN=Users,DC=MyDomain,DC=com"),
//				GroupMemberAttribute: pulumi.String("uniqueMember"),
//				GroupNameAttribute:   pulumi.String("cn"),
//				Strategy:             pulumi.String("STATIC"),
//				SubTree:              pulumi.Bool(true),
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
// ## Import
//
// ```sh
//
//	$ pulumi import artifactory:index/ldapGroupSettingV2:LdapGroupSettingV2 ldap ldapGroup1
//
// ```
type LdapGroupSettingV2 struct {
	pulumi.CustomResourceState

	// An attribute on the group entry which denoting the group description. Used when importing groups.
	DescriptionAttribute pulumi.StringOutput `pulumi:"descriptionAttribute"`
	// The LDAP setting key you want to use for group retrieval.
	EnabledLdap pulumi.StringOutput `pulumi:"enabledLdap"`
	// The LDAP filter used to search for group entries. Used for importing groups.
	Filter pulumi.StringOutput `pulumi:"filter"`
	// This attribute is used in very specific cases of LDAP group settings. Don't switch it to `false`, unless instructed by the JFrog support team. Default value is `false`.
	ForceAttributeSearch pulumi.BoolOutput `pulumi:"forceAttributeSearch"`
	// A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
	GroupBaseDn pulumi.StringOutput `pulumi:"groupBaseDn"`
	// A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember, member).
	GroupMemberAttribute pulumi.StringOutput `pulumi:"groupMemberAttribute"`
	// Attribute on the group entry denoting the group name. Used when importing groups.
	GroupNameAttribute pulumi.StringOutput `pulumi:"groupNameAttribute"`
	// Ldap group setting name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas: STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN. DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member. HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, `uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org` indicates that `user1` belongs to two groups: `uk` and `developers`. Valid values are: `STATIC`, `DYNAMIC`, `HIERARCHICAL`, case sensitive, all caps.
	Strategy pulumi.StringOutput `pulumi:"strategy"`
	// When set, enables deep search through the sub-tree of the LDAP URL + Search Base. `true` by default. `subTree` can be set to true only with `STATIC` or `DYNAMIC` strategy.
	SubTree pulumi.BoolOutput `pulumi:"subTree"`
}

// NewLdapGroupSettingV2 registers a new resource with the given unique name, arguments, and options.
func NewLdapGroupSettingV2(ctx *pulumi.Context,
	name string, args *LdapGroupSettingV2Args, opts ...pulumi.ResourceOption) (*LdapGroupSettingV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DescriptionAttribute == nil {
		return nil, errors.New("invalid value for required argument 'DescriptionAttribute'")
	}
	if args.Filter == nil {
		return nil, errors.New("invalid value for required argument 'Filter'")
	}
	if args.GroupMemberAttribute == nil {
		return nil, errors.New("invalid value for required argument 'GroupMemberAttribute'")
	}
	if args.GroupNameAttribute == nil {
		return nil, errors.New("invalid value for required argument 'GroupNameAttribute'")
	}
	if args.Strategy == nil {
		return nil, errors.New("invalid value for required argument 'Strategy'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LdapGroupSettingV2
	err := ctx.RegisterResource("artifactory:index/ldapGroupSettingV2:LdapGroupSettingV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLdapGroupSettingV2 gets an existing LdapGroupSettingV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLdapGroupSettingV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LdapGroupSettingV2State, opts ...pulumi.ResourceOption) (*LdapGroupSettingV2, error) {
	var resource LdapGroupSettingV2
	err := ctx.ReadResource("artifactory:index/ldapGroupSettingV2:LdapGroupSettingV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LdapGroupSettingV2 resources.
type ldapGroupSettingV2State struct {
	// An attribute on the group entry which denoting the group description. Used when importing groups.
	DescriptionAttribute *string `pulumi:"descriptionAttribute"`
	// The LDAP setting key you want to use for group retrieval.
	EnabledLdap *string `pulumi:"enabledLdap"`
	// The LDAP filter used to search for group entries. Used for importing groups.
	Filter *string `pulumi:"filter"`
	// This attribute is used in very specific cases of LDAP group settings. Don't switch it to `false`, unless instructed by the JFrog support team. Default value is `false`.
	ForceAttributeSearch *bool `pulumi:"forceAttributeSearch"`
	// A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
	GroupBaseDn *string `pulumi:"groupBaseDn"`
	// A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember, member).
	GroupMemberAttribute *string `pulumi:"groupMemberAttribute"`
	// Attribute on the group entry denoting the group name. Used when importing groups.
	GroupNameAttribute *string `pulumi:"groupNameAttribute"`
	// Ldap group setting name.
	Name *string `pulumi:"name"`
	// The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas: STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN. DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member. HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, `uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org` indicates that `user1` belongs to two groups: `uk` and `developers`. Valid values are: `STATIC`, `DYNAMIC`, `HIERARCHICAL`, case sensitive, all caps.
	Strategy *string `pulumi:"strategy"`
	// When set, enables deep search through the sub-tree of the LDAP URL + Search Base. `true` by default. `subTree` can be set to true only with `STATIC` or `DYNAMIC` strategy.
	SubTree *bool `pulumi:"subTree"`
}

type LdapGroupSettingV2State struct {
	// An attribute on the group entry which denoting the group description. Used when importing groups.
	DescriptionAttribute pulumi.StringPtrInput
	// The LDAP setting key you want to use for group retrieval.
	EnabledLdap pulumi.StringPtrInput
	// The LDAP filter used to search for group entries. Used for importing groups.
	Filter pulumi.StringPtrInput
	// This attribute is used in very specific cases of LDAP group settings. Don't switch it to `false`, unless instructed by the JFrog support team. Default value is `false`.
	ForceAttributeSearch pulumi.BoolPtrInput
	// A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
	GroupBaseDn pulumi.StringPtrInput
	// A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember, member).
	GroupMemberAttribute pulumi.StringPtrInput
	// Attribute on the group entry denoting the group name. Used when importing groups.
	GroupNameAttribute pulumi.StringPtrInput
	// Ldap group setting name.
	Name pulumi.StringPtrInput
	// The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas: STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN. DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member. HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, `uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org` indicates that `user1` belongs to two groups: `uk` and `developers`. Valid values are: `STATIC`, `DYNAMIC`, `HIERARCHICAL`, case sensitive, all caps.
	Strategy pulumi.StringPtrInput
	// When set, enables deep search through the sub-tree of the LDAP URL + Search Base. `true` by default. `subTree` can be set to true only with `STATIC` or `DYNAMIC` strategy.
	SubTree pulumi.BoolPtrInput
}

func (LdapGroupSettingV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*ldapGroupSettingV2State)(nil)).Elem()
}

type ldapGroupSettingV2Args struct {
	// An attribute on the group entry which denoting the group description. Used when importing groups.
	DescriptionAttribute string `pulumi:"descriptionAttribute"`
	// The LDAP setting key you want to use for group retrieval.
	EnabledLdap *string `pulumi:"enabledLdap"`
	// The LDAP filter used to search for group entries. Used for importing groups.
	Filter string `pulumi:"filter"`
	// This attribute is used in very specific cases of LDAP group settings. Don't switch it to `false`, unless instructed by the JFrog support team. Default value is `false`.
	ForceAttributeSearch *bool `pulumi:"forceAttributeSearch"`
	// A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
	GroupBaseDn *string `pulumi:"groupBaseDn"`
	// A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember, member).
	GroupMemberAttribute string `pulumi:"groupMemberAttribute"`
	// Attribute on the group entry denoting the group name. Used when importing groups.
	GroupNameAttribute string `pulumi:"groupNameAttribute"`
	// Ldap group setting name.
	Name *string `pulumi:"name"`
	// The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas: STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN. DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member. HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, `uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org` indicates that `user1` belongs to two groups: `uk` and `developers`. Valid values are: `STATIC`, `DYNAMIC`, `HIERARCHICAL`, case sensitive, all caps.
	Strategy string `pulumi:"strategy"`
	// When set, enables deep search through the sub-tree of the LDAP URL + Search Base. `true` by default. `subTree` can be set to true only with `STATIC` or `DYNAMIC` strategy.
	SubTree *bool `pulumi:"subTree"`
}

// The set of arguments for constructing a LdapGroupSettingV2 resource.
type LdapGroupSettingV2Args struct {
	// An attribute on the group entry which denoting the group description. Used when importing groups.
	DescriptionAttribute pulumi.StringInput
	// The LDAP setting key you want to use for group retrieval.
	EnabledLdap pulumi.StringPtrInput
	// The LDAP filter used to search for group entries. Used for importing groups.
	Filter pulumi.StringInput
	// This attribute is used in very specific cases of LDAP group settings. Don't switch it to `false`, unless instructed by the JFrog support team. Default value is `false`.
	ForceAttributeSearch pulumi.BoolPtrInput
	// A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
	GroupBaseDn pulumi.StringPtrInput
	// A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember, member).
	GroupMemberAttribute pulumi.StringInput
	// Attribute on the group entry denoting the group name. Used when importing groups.
	GroupNameAttribute pulumi.StringInput
	// Ldap group setting name.
	Name pulumi.StringPtrInput
	// The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas: STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN. DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member. HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, `uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org` indicates that `user1` belongs to two groups: `uk` and `developers`. Valid values are: `STATIC`, `DYNAMIC`, `HIERARCHICAL`, case sensitive, all caps.
	Strategy pulumi.StringInput
	// When set, enables deep search through the sub-tree of the LDAP URL + Search Base. `true` by default. `subTree` can be set to true only with `STATIC` or `DYNAMIC` strategy.
	SubTree pulumi.BoolPtrInput
}

func (LdapGroupSettingV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*ldapGroupSettingV2Args)(nil)).Elem()
}

type LdapGroupSettingV2Input interface {
	pulumi.Input

	ToLdapGroupSettingV2Output() LdapGroupSettingV2Output
	ToLdapGroupSettingV2OutputWithContext(ctx context.Context) LdapGroupSettingV2Output
}

func (*LdapGroupSettingV2) ElementType() reflect.Type {
	return reflect.TypeOf((**LdapGroupSettingV2)(nil)).Elem()
}

func (i *LdapGroupSettingV2) ToLdapGroupSettingV2Output() LdapGroupSettingV2Output {
	return i.ToLdapGroupSettingV2OutputWithContext(context.Background())
}

func (i *LdapGroupSettingV2) ToLdapGroupSettingV2OutputWithContext(ctx context.Context) LdapGroupSettingV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(LdapGroupSettingV2Output)
}

// LdapGroupSettingV2ArrayInput is an input type that accepts LdapGroupSettingV2Array and LdapGroupSettingV2ArrayOutput values.
// You can construct a concrete instance of `LdapGroupSettingV2ArrayInput` via:
//
//	LdapGroupSettingV2Array{ LdapGroupSettingV2Args{...} }
type LdapGroupSettingV2ArrayInput interface {
	pulumi.Input

	ToLdapGroupSettingV2ArrayOutput() LdapGroupSettingV2ArrayOutput
	ToLdapGroupSettingV2ArrayOutputWithContext(context.Context) LdapGroupSettingV2ArrayOutput
}

type LdapGroupSettingV2Array []LdapGroupSettingV2Input

func (LdapGroupSettingV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LdapGroupSettingV2)(nil)).Elem()
}

func (i LdapGroupSettingV2Array) ToLdapGroupSettingV2ArrayOutput() LdapGroupSettingV2ArrayOutput {
	return i.ToLdapGroupSettingV2ArrayOutputWithContext(context.Background())
}

func (i LdapGroupSettingV2Array) ToLdapGroupSettingV2ArrayOutputWithContext(ctx context.Context) LdapGroupSettingV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LdapGroupSettingV2ArrayOutput)
}

// LdapGroupSettingV2MapInput is an input type that accepts LdapGroupSettingV2Map and LdapGroupSettingV2MapOutput values.
// You can construct a concrete instance of `LdapGroupSettingV2MapInput` via:
//
//	LdapGroupSettingV2Map{ "key": LdapGroupSettingV2Args{...} }
type LdapGroupSettingV2MapInput interface {
	pulumi.Input

	ToLdapGroupSettingV2MapOutput() LdapGroupSettingV2MapOutput
	ToLdapGroupSettingV2MapOutputWithContext(context.Context) LdapGroupSettingV2MapOutput
}

type LdapGroupSettingV2Map map[string]LdapGroupSettingV2Input

func (LdapGroupSettingV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LdapGroupSettingV2)(nil)).Elem()
}

func (i LdapGroupSettingV2Map) ToLdapGroupSettingV2MapOutput() LdapGroupSettingV2MapOutput {
	return i.ToLdapGroupSettingV2MapOutputWithContext(context.Background())
}

func (i LdapGroupSettingV2Map) ToLdapGroupSettingV2MapOutputWithContext(ctx context.Context) LdapGroupSettingV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LdapGroupSettingV2MapOutput)
}

type LdapGroupSettingV2Output struct{ *pulumi.OutputState }

func (LdapGroupSettingV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**LdapGroupSettingV2)(nil)).Elem()
}

func (o LdapGroupSettingV2Output) ToLdapGroupSettingV2Output() LdapGroupSettingV2Output {
	return o
}

func (o LdapGroupSettingV2Output) ToLdapGroupSettingV2OutputWithContext(ctx context.Context) LdapGroupSettingV2Output {
	return o
}

// An attribute on the group entry which denoting the group description. Used when importing groups.
func (o LdapGroupSettingV2Output) DescriptionAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v *LdapGroupSettingV2) pulumi.StringOutput { return v.DescriptionAttribute }).(pulumi.StringOutput)
}

// The LDAP setting key you want to use for group retrieval.
func (o LdapGroupSettingV2Output) EnabledLdap() pulumi.StringOutput {
	return o.ApplyT(func(v *LdapGroupSettingV2) pulumi.StringOutput { return v.EnabledLdap }).(pulumi.StringOutput)
}

// The LDAP filter used to search for group entries. Used for importing groups.
func (o LdapGroupSettingV2Output) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v *LdapGroupSettingV2) pulumi.StringOutput { return v.Filter }).(pulumi.StringOutput)
}

// This attribute is used in very specific cases of LDAP group settings. Don't switch it to `false`, unless instructed by the JFrog support team. Default value is `false`.
func (o LdapGroupSettingV2Output) ForceAttributeSearch() pulumi.BoolOutput {
	return o.ApplyT(func(v *LdapGroupSettingV2) pulumi.BoolOutput { return v.ForceAttributeSearch }).(pulumi.BoolOutput)
}

// A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
func (o LdapGroupSettingV2Output) GroupBaseDn() pulumi.StringOutput {
	return o.ApplyT(func(v *LdapGroupSettingV2) pulumi.StringOutput { return v.GroupBaseDn }).(pulumi.StringOutput)
}

// A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember, member).
func (o LdapGroupSettingV2Output) GroupMemberAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v *LdapGroupSettingV2) pulumi.StringOutput { return v.GroupMemberAttribute }).(pulumi.StringOutput)
}

// Attribute on the group entry denoting the group name. Used when importing groups.
func (o LdapGroupSettingV2Output) GroupNameAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v *LdapGroupSettingV2) pulumi.StringOutput { return v.GroupNameAttribute }).(pulumi.StringOutput)
}

// Ldap group setting name.
func (o LdapGroupSettingV2Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LdapGroupSettingV2) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas: STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN. DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member. HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, `uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org` indicates that `user1` belongs to two groups: `uk` and `developers`. Valid values are: `STATIC`, `DYNAMIC`, `HIERARCHICAL`, case sensitive, all caps.
func (o LdapGroupSettingV2Output) Strategy() pulumi.StringOutput {
	return o.ApplyT(func(v *LdapGroupSettingV2) pulumi.StringOutput { return v.Strategy }).(pulumi.StringOutput)
}

// When set, enables deep search through the sub-tree of the LDAP URL + Search Base. `true` by default. `subTree` can be set to true only with `STATIC` or `DYNAMIC` strategy.
func (o LdapGroupSettingV2Output) SubTree() pulumi.BoolOutput {
	return o.ApplyT(func(v *LdapGroupSettingV2) pulumi.BoolOutput { return v.SubTree }).(pulumi.BoolOutput)
}

type LdapGroupSettingV2ArrayOutput struct{ *pulumi.OutputState }

func (LdapGroupSettingV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LdapGroupSettingV2)(nil)).Elem()
}

func (o LdapGroupSettingV2ArrayOutput) ToLdapGroupSettingV2ArrayOutput() LdapGroupSettingV2ArrayOutput {
	return o
}

func (o LdapGroupSettingV2ArrayOutput) ToLdapGroupSettingV2ArrayOutputWithContext(ctx context.Context) LdapGroupSettingV2ArrayOutput {
	return o
}

func (o LdapGroupSettingV2ArrayOutput) Index(i pulumi.IntInput) LdapGroupSettingV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LdapGroupSettingV2 {
		return vs[0].([]*LdapGroupSettingV2)[vs[1].(int)]
	}).(LdapGroupSettingV2Output)
}

type LdapGroupSettingV2MapOutput struct{ *pulumi.OutputState }

func (LdapGroupSettingV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LdapGroupSettingV2)(nil)).Elem()
}

func (o LdapGroupSettingV2MapOutput) ToLdapGroupSettingV2MapOutput() LdapGroupSettingV2MapOutput {
	return o
}

func (o LdapGroupSettingV2MapOutput) ToLdapGroupSettingV2MapOutputWithContext(ctx context.Context) LdapGroupSettingV2MapOutput {
	return o
}

func (o LdapGroupSettingV2MapOutput) MapIndex(k pulumi.StringInput) LdapGroupSettingV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LdapGroupSettingV2 {
		return vs[0].(map[string]*LdapGroupSettingV2)[vs[1].(string)]
	}).(LdapGroupSettingV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LdapGroupSettingV2Input)(nil)).Elem(), &LdapGroupSettingV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*LdapGroupSettingV2ArrayInput)(nil)).Elem(), LdapGroupSettingV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*LdapGroupSettingV2MapInput)(nil)).Elem(), LdapGroupSettingV2Map{})
	pulumi.RegisterOutputType(LdapGroupSettingV2Output{})
	pulumi.RegisterOutputType(LdapGroupSettingV2ArrayOutput{})
	pulumi.RegisterOutputType(LdapGroupSettingV2MapOutput{})
}
