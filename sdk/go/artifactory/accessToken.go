// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ## Example Usage
//
// ### S
// ### Create a new Artifactory Access Token for an existing user
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewAccessToken(ctx, "exisingUser", &artifactory.AccessTokenArgs{
//				EndDateRelative: pulumi.String("5m"),
//				Username:        pulumi.String("existing-user"),
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
// Note: This assumes that the user `existing-user` has already been created in Artifactory by different means, i.e. manually or in a separate pulumi up.
// ### Create a new Artifactory User and Access token
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			newUserUser, err := artifactory.NewUser(ctx, "newUserUser", &artifactory.UserArgs{
//				Email: pulumi.String("new_user@somewhere.com"),
//				Groups: pulumi.StringArray{
//					pulumi.String("readers"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = artifactory.NewAccessToken(ctx, "newUserAccessToken", &artifactory.AccessTokenArgs{
//				Username:        newUserUser.Name,
//				EndDateRelative: pulumi.String("5m"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Creates a new token for groups
// This creates a transient user called `temporary-user`.
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewAccessToken(ctx, "temporaryUser", &artifactory.AccessTokenArgs{
//				EndDateRelative: pulumi.String("1h"),
//				Groups: pulumi.StringArray{
//					pulumi.String("readers"),
//				},
//				Username: pulumi.String("temporary-user"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Create token with no expiry
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewAccessToken(ctx, "noExpiry", &artifactory.AccessTokenArgs{
//				EndDateRelative: pulumi.String("0s"),
//				Username:        pulumi.String("existing-user"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Creates a refreshable token
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewAccessToken(ctx, "refreshable", &artifactory.AccessTokenArgs{
//				EndDateRelative: pulumi.String("1m"),
//				Groups: pulumi.StringArray{
//					pulumi.String("readers"),
//				},
//				Refreshable: pulumi.Bool(true),
//				Username:    pulumi.String("refreshable"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Creates an administrator token
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewAccessToken(ctx, "admin", &artifactory.AccessTokenArgs{
//				AdminToken: &artifactory.AccessTokenAdminTokenArgs{
//					InstanceId: pulumi.String("<instance id>"),
//				},
//				EndDateRelative: pulumi.String("1m"),
//				Username:        pulumi.String("admin"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Creates a token with an audience
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewAccessToken(ctx, "audience", &artifactory.AccessTokenArgs{
//				Audience:        pulumi.String("jfrt@*"),
//				EndDateRelative: pulumi.String("1m"),
//				Refreshable:     pulumi.Bool(true),
//				Username:        pulumi.String("audience"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Creates a token with a fixed end date
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewAccessToken(ctx, "fixeddate", &artifactory.AccessTokenArgs{
//				EndDate: pulumi.String("2018-01-01T01:02:03Z"),
//				Groups: pulumi.StringArray{
//					pulumi.String("readers"),
//				},
//				Username: pulumi.String("fixeddate"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Rotate token after it expires
// This example will generate a token that will expire in 1 hour.
//
// If `pulumi up` is run before 1 hour, nothing changes.
// One an hour has passed, `pulumi up` will generate a new token.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory"
//	"github.com/pulumi/pulumi-time/sdk/go/time"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := time.NewRotating(ctx, "nowPlus1Hours", &time.RotatingArgs{
//				RotationHours: pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = artifactory.NewAccessToken(ctx, "rotating", &artifactory.AccessTokenArgs{
//				Username: pulumi.String("rotating"),
//				EndDate:  pulumi.Any(time_rotating.Now_plus_1_hour.Rotation_rfc3339),
//				Groups: pulumi.StringArray{
//					pulumi.String("readers"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## References
//
// - https://www.jfrog.com/confluence/display/ACC1X/Access+Tokens
// - https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-CreateToken
//
// ## Import
//
// Artifactory **does not** retain access tokens and cannot be imported into state.
type AccessToken struct {
	pulumi.CustomResourceState

	// Returns the access token to authenciate to Artifactory
	AccessToken pulumi.StringOutput `pulumi:"accessToken"`
	// (Optional) Specify the `instanceId` in this block to grant this token admin privileges. This can only be created when the authenticated user is an admin. `adminToken` cannot be specified with `groups`.
	AdminToken AccessTokenAdminTokenPtrOutput `pulumi:"adminToken"`
	// (Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set `"jfrt@*"` so the token to be accepted by all Artifactory instances.
	Audience pulumi.StringPtrOutput `pulumi:"audience"`
	// (Optional) The end date which the token is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
	EndDate pulumi.StringOutput `pulumi:"endDate"`
	// (Optional) A relative duration for which the token is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "s", "m", "h".
	EndDateRelative pulumi.StringPtrOutput `pulumi:"endDateRelative"`
	// (Optional) List of groups. The token is granted access based on the permissions of the groups. Specify `["*"]` for all groups that the user belongs to. `groups` cannot be specified with `adminToken`.
	Groups pulumi.StringArrayOutput `pulumi:"groups"`
	// Returns the refresh token when `refreshable` is true, or an empty string when `refreshable` is false.
	RefreshToken pulumi.StringOutput `pulumi:"refreshToken"`
	// (Optional) Is this token refreshable? Defaults to `false`
	Refreshable pulumi.BoolPtrOutput `pulumi:"refreshable"`
	// (Required) The username or subject for the token. A non-admin can only specify their own username. Admins can specify any existing username, or a new name for a temporary token. Temporary tokens require `groups` to be set.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewAccessToken registers a new resource with the given unique name, arguments, and options.
func NewAccessToken(ctx *pulumi.Context,
	name string, args *AccessTokenArgs, opts ...pulumi.ResourceOption) (*AccessToken, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"accessToken",
		"refreshToken",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AccessToken
	err := ctx.RegisterResource("artifactory:index/accessToken:AccessToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessToken gets an existing AccessToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessTokenState, opts ...pulumi.ResourceOption) (*AccessToken, error) {
	var resource AccessToken
	err := ctx.ReadResource("artifactory:index/accessToken:AccessToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessToken resources.
type accessTokenState struct {
	// Returns the access token to authenciate to Artifactory
	AccessToken *string `pulumi:"accessToken"`
	// (Optional) Specify the `instanceId` in this block to grant this token admin privileges. This can only be created when the authenticated user is an admin. `adminToken` cannot be specified with `groups`.
	AdminToken *AccessTokenAdminToken `pulumi:"adminToken"`
	// (Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set `"jfrt@*"` so the token to be accepted by all Artifactory instances.
	Audience *string `pulumi:"audience"`
	// (Optional) The end date which the token is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
	EndDate *string `pulumi:"endDate"`
	// (Optional) A relative duration for which the token is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "s", "m", "h".
	EndDateRelative *string `pulumi:"endDateRelative"`
	// (Optional) List of groups. The token is granted access based on the permissions of the groups. Specify `["*"]` for all groups that the user belongs to. `groups` cannot be specified with `adminToken`.
	Groups []string `pulumi:"groups"`
	// Returns the refresh token when `refreshable` is true, or an empty string when `refreshable` is false.
	RefreshToken *string `pulumi:"refreshToken"`
	// (Optional) Is this token refreshable? Defaults to `false`
	Refreshable *bool `pulumi:"refreshable"`
	// (Required) The username or subject for the token. A non-admin can only specify their own username. Admins can specify any existing username, or a new name for a temporary token. Temporary tokens require `groups` to be set.
	Username *string `pulumi:"username"`
}

type AccessTokenState struct {
	// Returns the access token to authenciate to Artifactory
	AccessToken pulumi.StringPtrInput
	// (Optional) Specify the `instanceId` in this block to grant this token admin privileges. This can only be created when the authenticated user is an admin. `adminToken` cannot be specified with `groups`.
	AdminToken AccessTokenAdminTokenPtrInput
	// (Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set `"jfrt@*"` so the token to be accepted by all Artifactory instances.
	Audience pulumi.StringPtrInput
	// (Optional) The end date which the token is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
	EndDate pulumi.StringPtrInput
	// (Optional) A relative duration for which the token is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "s", "m", "h".
	EndDateRelative pulumi.StringPtrInput
	// (Optional) List of groups. The token is granted access based on the permissions of the groups. Specify `["*"]` for all groups that the user belongs to. `groups` cannot be specified with `adminToken`.
	Groups pulumi.StringArrayInput
	// Returns the refresh token when `refreshable` is true, or an empty string when `refreshable` is false.
	RefreshToken pulumi.StringPtrInput
	// (Optional) Is this token refreshable? Defaults to `false`
	Refreshable pulumi.BoolPtrInput
	// (Required) The username or subject for the token. A non-admin can only specify their own username. Admins can specify any existing username, or a new name for a temporary token. Temporary tokens require `groups` to be set.
	Username pulumi.StringPtrInput
}

func (AccessTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessTokenState)(nil)).Elem()
}

type accessTokenArgs struct {
	// (Optional) Specify the `instanceId` in this block to grant this token admin privileges. This can only be created when the authenticated user is an admin. `adminToken` cannot be specified with `groups`.
	AdminToken *AccessTokenAdminToken `pulumi:"adminToken"`
	// (Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set `"jfrt@*"` so the token to be accepted by all Artifactory instances.
	Audience *string `pulumi:"audience"`
	// (Optional) The end date which the token is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
	EndDate *string `pulumi:"endDate"`
	// (Optional) A relative duration for which the token is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "s", "m", "h".
	EndDateRelative *string `pulumi:"endDateRelative"`
	// (Optional) List of groups. The token is granted access based on the permissions of the groups. Specify `["*"]` for all groups that the user belongs to. `groups` cannot be specified with `adminToken`.
	Groups []string `pulumi:"groups"`
	// (Optional) Is this token refreshable? Defaults to `false`
	Refreshable *bool `pulumi:"refreshable"`
	// (Required) The username or subject for the token. A non-admin can only specify their own username. Admins can specify any existing username, or a new name for a temporary token. Temporary tokens require `groups` to be set.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a AccessToken resource.
type AccessTokenArgs struct {
	// (Optional) Specify the `instanceId` in this block to grant this token admin privileges. This can only be created when the authenticated user is an admin. `adminToken` cannot be specified with `groups`.
	AdminToken AccessTokenAdminTokenPtrInput
	// (Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set `"jfrt@*"` so the token to be accepted by all Artifactory instances.
	Audience pulumi.StringPtrInput
	// (Optional) The end date which the token is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
	EndDate pulumi.StringPtrInput
	// (Optional) A relative duration for which the token is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "s", "m", "h".
	EndDateRelative pulumi.StringPtrInput
	// (Optional) List of groups. The token is granted access based on the permissions of the groups. Specify `["*"]` for all groups that the user belongs to. `groups` cannot be specified with `adminToken`.
	Groups pulumi.StringArrayInput
	// (Optional) Is this token refreshable? Defaults to `false`
	Refreshable pulumi.BoolPtrInput
	// (Required) The username or subject for the token. A non-admin can only specify their own username. Admins can specify any existing username, or a new name for a temporary token. Temporary tokens require `groups` to be set.
	Username pulumi.StringInput
}

func (AccessTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessTokenArgs)(nil)).Elem()
}

type AccessTokenInput interface {
	pulumi.Input

	ToAccessTokenOutput() AccessTokenOutput
	ToAccessTokenOutputWithContext(ctx context.Context) AccessTokenOutput
}

func (*AccessToken) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessToken)(nil)).Elem()
}

func (i *AccessToken) ToAccessTokenOutput() AccessTokenOutput {
	return i.ToAccessTokenOutputWithContext(context.Background())
}

func (i *AccessToken) ToAccessTokenOutputWithContext(ctx context.Context) AccessTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessTokenOutput)
}

func (i *AccessToken) ToOutput(ctx context.Context) pulumix.Output[*AccessToken] {
	return pulumix.Output[*AccessToken]{
		OutputState: i.ToAccessTokenOutputWithContext(ctx).OutputState,
	}
}

// AccessTokenArrayInput is an input type that accepts AccessTokenArray and AccessTokenArrayOutput values.
// You can construct a concrete instance of `AccessTokenArrayInput` via:
//
//	AccessTokenArray{ AccessTokenArgs{...} }
type AccessTokenArrayInput interface {
	pulumi.Input

	ToAccessTokenArrayOutput() AccessTokenArrayOutput
	ToAccessTokenArrayOutputWithContext(context.Context) AccessTokenArrayOutput
}

type AccessTokenArray []AccessTokenInput

func (AccessTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessToken)(nil)).Elem()
}

func (i AccessTokenArray) ToAccessTokenArrayOutput() AccessTokenArrayOutput {
	return i.ToAccessTokenArrayOutputWithContext(context.Background())
}

func (i AccessTokenArray) ToAccessTokenArrayOutputWithContext(ctx context.Context) AccessTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessTokenArrayOutput)
}

func (i AccessTokenArray) ToOutput(ctx context.Context) pulumix.Output[[]*AccessToken] {
	return pulumix.Output[[]*AccessToken]{
		OutputState: i.ToAccessTokenArrayOutputWithContext(ctx).OutputState,
	}
}

// AccessTokenMapInput is an input type that accepts AccessTokenMap and AccessTokenMapOutput values.
// You can construct a concrete instance of `AccessTokenMapInput` via:
//
//	AccessTokenMap{ "key": AccessTokenArgs{...} }
type AccessTokenMapInput interface {
	pulumi.Input

	ToAccessTokenMapOutput() AccessTokenMapOutput
	ToAccessTokenMapOutputWithContext(context.Context) AccessTokenMapOutput
}

type AccessTokenMap map[string]AccessTokenInput

func (AccessTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessToken)(nil)).Elem()
}

func (i AccessTokenMap) ToAccessTokenMapOutput() AccessTokenMapOutput {
	return i.ToAccessTokenMapOutputWithContext(context.Background())
}

func (i AccessTokenMap) ToAccessTokenMapOutputWithContext(ctx context.Context) AccessTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessTokenMapOutput)
}

func (i AccessTokenMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*AccessToken] {
	return pulumix.Output[map[string]*AccessToken]{
		OutputState: i.ToAccessTokenMapOutputWithContext(ctx).OutputState,
	}
}

type AccessTokenOutput struct{ *pulumi.OutputState }

func (AccessTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessToken)(nil)).Elem()
}

func (o AccessTokenOutput) ToAccessTokenOutput() AccessTokenOutput {
	return o
}

func (o AccessTokenOutput) ToAccessTokenOutputWithContext(ctx context.Context) AccessTokenOutput {
	return o
}

func (o AccessTokenOutput) ToOutput(ctx context.Context) pulumix.Output[*AccessToken] {
	return pulumix.Output[*AccessToken]{
		OutputState: o.OutputState,
	}
}

// Returns the access token to authenciate to Artifactory
func (o AccessTokenOutput) AccessToken() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessToken) pulumi.StringOutput { return v.AccessToken }).(pulumi.StringOutput)
}

// (Optional) Specify the `instanceId` in this block to grant this token admin privileges. This can only be created when the authenticated user is an admin. `adminToken` cannot be specified with `groups`.
func (o AccessTokenOutput) AdminToken() AccessTokenAdminTokenPtrOutput {
	return o.ApplyT(func(v *AccessToken) AccessTokenAdminTokenPtrOutput { return v.AdminToken }).(AccessTokenAdminTokenPtrOutput)
}

// (Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set `"jfrt@*"` so the token to be accepted by all Artifactory instances.
func (o AccessTokenOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessToken) pulumi.StringPtrOutput { return v.Audience }).(pulumi.StringPtrOutput)
}

// (Optional) The end date which the token is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
func (o AccessTokenOutput) EndDate() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessToken) pulumi.StringOutput { return v.EndDate }).(pulumi.StringOutput)
}

// (Optional) A relative duration for which the token is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "s", "m", "h".
func (o AccessTokenOutput) EndDateRelative() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessToken) pulumi.StringPtrOutput { return v.EndDateRelative }).(pulumi.StringPtrOutput)
}

// (Optional) List of groups. The token is granted access based on the permissions of the groups. Specify `["*"]` for all groups that the user belongs to. `groups` cannot be specified with `adminToken`.
func (o AccessTokenOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AccessToken) pulumi.StringArrayOutput { return v.Groups }).(pulumi.StringArrayOutput)
}

// Returns the refresh token when `refreshable` is true, or an empty string when `refreshable` is false.
func (o AccessTokenOutput) RefreshToken() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessToken) pulumi.StringOutput { return v.RefreshToken }).(pulumi.StringOutput)
}

// (Optional) Is this token refreshable? Defaults to `false`
func (o AccessTokenOutput) Refreshable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccessToken) pulumi.BoolPtrOutput { return v.Refreshable }).(pulumi.BoolPtrOutput)
}

// (Required) The username or subject for the token. A non-admin can only specify their own username. Admins can specify any existing username, or a new name for a temporary token. Temporary tokens require `groups` to be set.
func (o AccessTokenOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessToken) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type AccessTokenArrayOutput struct{ *pulumi.OutputState }

func (AccessTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessToken)(nil)).Elem()
}

func (o AccessTokenArrayOutput) ToAccessTokenArrayOutput() AccessTokenArrayOutput {
	return o
}

func (o AccessTokenArrayOutput) ToAccessTokenArrayOutputWithContext(ctx context.Context) AccessTokenArrayOutput {
	return o
}

func (o AccessTokenArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*AccessToken] {
	return pulumix.Output[[]*AccessToken]{
		OutputState: o.OutputState,
	}
}

func (o AccessTokenArrayOutput) Index(i pulumi.IntInput) AccessTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccessToken {
		return vs[0].([]*AccessToken)[vs[1].(int)]
	}).(AccessTokenOutput)
}

type AccessTokenMapOutput struct{ *pulumi.OutputState }

func (AccessTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessToken)(nil)).Elem()
}

func (o AccessTokenMapOutput) ToAccessTokenMapOutput() AccessTokenMapOutput {
	return o
}

func (o AccessTokenMapOutput) ToAccessTokenMapOutputWithContext(ctx context.Context) AccessTokenMapOutput {
	return o
}

func (o AccessTokenMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*AccessToken] {
	return pulumix.Output[map[string]*AccessToken]{
		OutputState: o.OutputState,
	}
}

func (o AccessTokenMapOutput) MapIndex(k pulumi.StringInput) AccessTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccessToken {
		return vs[0].(map[string]*AccessToken)[vs[1].(string)]
	}).(AccessTokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessTokenInput)(nil)).Elem(), &AccessToken{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessTokenArrayInput)(nil)).Elem(), AccessTokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessTokenMapInput)(nil)).Elem(), AccessTokenMap{})
	pulumi.RegisterOutputType(AccessTokenOutput{})
	pulumi.RegisterOutputType(AccessTokenArrayOutput{})
	pulumi.RegisterOutputType(AccessTokenMapOutput{})
}
