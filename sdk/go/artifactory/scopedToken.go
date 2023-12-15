// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v6/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Artifactory Scoped Token resource. This can be used to create and manage Artifactory Scoped Tokens.
//
// !>Scoped Tokens will be stored in the raw state as plain-text. Read more about sensitive data in
// state.
//
// ~>Token would not be saved by Artifactory if `expiresIn` is less than the persistency threshold value (default to 10800 seconds) set in Access configuration. See [Persistency Threshold](https://jfrog.com/help/r/jfrog-platform-administration-documentation/using-the-revocable-and-persistency-thresholds) for details.
//
// ## Example Usage
//
// ### Create a new Artifactory scoped token for an existing user
//
//	resource "artifactory_scoped_token" "scoped_token" {
//	  username = "existing-user"
//	}
//
// ### **Note:** This assumes that the user `existing-user` has already been created in Artifactory by different means, i.e. manually or in a separate pulumi up.
//
// ### Create a new Artifactory user and scoped token
//
//	resource "artifactory_user" "new_user" {
//	  name   = "new_user"
//	  email  = "new_user@somewhere.com"
//	  groups = ["readers"]
//	}
//
//	resource "artifactory_scoped_token" "scoped_token_user" {
//	  username = artifactory_user.new_user.name
//	}
//
// ### Creates a new token for groups
//
//	resource "artifactory_scoped_token" "scoped_token_group" {
//	  scopes = ["applied-permissions/groups:readers"]
//	}
//
// ### Create token with expiry
//
//	resource "artifactory_scoped_token" "scoped_token_no_expiry" {
//	  username   = "existing-user"
//	  expires_in = 7200 // in seconds
//	}
//
// ### Creates a refreshable token
//
//	resource "artifactory_scoped_token" "scoped_token_refreshable" {
//	  username    = "existing-user"
//	  refreshable = true
//	}
//
// ### Creates an administrator token
//
//	resource "artifactory_scoped_token" "admin" {
//	  username = "admin-user"
//	  scopes   = ["applied-permissions/admin"]
//	}
//
// ## References
//
// - https://jfrog.com/help/r/jfrog-platform-administration-documentation/access-tokens
// - https://jfrog.com/help/r/jfrog-rest-apis/access-tokens
//
// ## Import
//
// Artifactory **does not** retain scoped tokens, and they cannot be imported into state.
type ScopedToken struct {
	pulumi.CustomResourceState

	// Returns the access token to authenticate to Artifactory.
	AccessToken pulumi.StringOutput `pulumi:"accessToken"`
	// A list of the other instances or services that should accept this token identified by their Service-IDs. Limited to total 255 characters. Default to '*@*' if not set. Service ID must begin with valid JFrog service type. Options: jfrt, jfxr, jfpip, jfds, jfmc, jfac, jfevt, jfmd, jfcon, or *. For instructions to retrieve the Artifactory Service ID see this [documentation](https://jfrog.com/help/r/jfrog-rest-apis/get-service-id)
	Audiences pulumi.StringArrayOutput `pulumi:"audiences"`
	// Free text token description. Useful for filtering and managing tokens. Limited to 1024 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	// The amount of time, in seconds, it would take for the token to expire. An admin shall be able to set whether expiry is mandatory, what is the default expiry, and what is the maximum expiry allowed. Must be non-negative. Default value is based on configuration in 'access.config.yaml'. See [API documentation](https://jfrog.com/help/r/jfrog-rest-apis/create-token) for details. Access Token would not be saved by Artifactory if this is less than the persistence threshold value (default to 10800 seconds) set in Access configuration. See [official documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/using-the-revocable-and-persistency-thresholds) for details.
	ExpiresIn pulumi.IntOutput `pulumi:"expiresIn"`
	// Returns the token expiry.
	Expiry pulumi.IntOutput `pulumi:"expiry"`
	// The grant type used to authenticate the request. In this case, the only value supported is `clientCredentials` which is also the default value if this parameter is not specified.
	GrantType pulumi.StringOutput `pulumi:"grantType"`
	// Also create a reference token which can be used like an API key.
	IncludeReferenceToken pulumi.BoolOutput `pulumi:"includeReferenceToken"`
	// Returns the token issued at date/time.
	IssuedAt pulumi.IntOutput `pulumi:"issuedAt"`
	// Returns the token issuer.
	Issuer pulumi.StringOutput `pulumi:"issuer"`
	// The project for which this token is created. Enter the project name on which you want to apply this token.
	ProjectKey pulumi.StringPtrOutput `pulumi:"projectKey"`
	// Reference Token (alias to Access Token).
	ReferenceToken pulumi.StringOutput `pulumi:"referenceToken"`
	// Refresh token.
	RefreshToken pulumi.StringOutput `pulumi:"refreshToken"`
	// Is this token refreshable? Default is `false`.
	Refreshable pulumi.BoolOutput `pulumi:"refreshable"`
	// The scope of access that the token provides. Access to the REST API is always provided by default. Administrators can set any scope, while non-admin users can only set the scope to a subset of the groups to which they belong.
	// The supported scopes include:
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
	// Returns the token type.
	Subject pulumi.StringOutput `pulumi:"subject"`
	// Returns the token type.
	TokenType pulumi.StringOutput `pulumi:"tokenType"`
	// The user name for which this token is created. The username is based on the authenticated user - either from the user of the authenticated token or based on the username (if basic auth was used). The username is then used to set the subject of the token: \n\n/users/\n\n. Limited to 255 characters.
	Username pulumi.StringPtrOutput `pulumi:"username"`
}

// NewScopedToken registers a new resource with the given unique name, arguments, and options.
func NewScopedToken(ctx *pulumi.Context,
	name string, args *ScopedTokenArgs, opts ...pulumi.ResourceOption) (*ScopedToken, error) {
	if args == nil {
		args = &ScopedTokenArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"accessToken",
		"referenceToken",
		"refreshToken",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ScopedToken
	err := ctx.RegisterResource("artifactory:index/scopedToken:ScopedToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScopedToken gets an existing ScopedToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScopedToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScopedTokenState, opts ...pulumi.ResourceOption) (*ScopedToken, error) {
	var resource ScopedToken
	err := ctx.ReadResource("artifactory:index/scopedToken:ScopedToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScopedToken resources.
type scopedTokenState struct {
	// Returns the access token to authenticate to Artifactory.
	AccessToken *string `pulumi:"accessToken"`
	// A list of the other instances or services that should accept this token identified by their Service-IDs. Limited to total 255 characters. Default to '*@*' if not set. Service ID must begin with valid JFrog service type. Options: jfrt, jfxr, jfpip, jfds, jfmc, jfac, jfevt, jfmd, jfcon, or *. For instructions to retrieve the Artifactory Service ID see this [documentation](https://jfrog.com/help/r/jfrog-rest-apis/get-service-id)
	Audiences []string `pulumi:"audiences"`
	// Free text token description. Useful for filtering and managing tokens. Limited to 1024 characters.
	Description *string `pulumi:"description"`
	// The amount of time, in seconds, it would take for the token to expire. An admin shall be able to set whether expiry is mandatory, what is the default expiry, and what is the maximum expiry allowed. Must be non-negative. Default value is based on configuration in 'access.config.yaml'. See [API documentation](https://jfrog.com/help/r/jfrog-rest-apis/create-token) for details. Access Token would not be saved by Artifactory if this is less than the persistence threshold value (default to 10800 seconds) set in Access configuration. See [official documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/using-the-revocable-and-persistency-thresholds) for details.
	ExpiresIn *int `pulumi:"expiresIn"`
	// Returns the token expiry.
	Expiry *int `pulumi:"expiry"`
	// The grant type used to authenticate the request. In this case, the only value supported is `clientCredentials` which is also the default value if this parameter is not specified.
	GrantType *string `pulumi:"grantType"`
	// Also create a reference token which can be used like an API key.
	IncludeReferenceToken *bool `pulumi:"includeReferenceToken"`
	// Returns the token issued at date/time.
	IssuedAt *int `pulumi:"issuedAt"`
	// Returns the token issuer.
	Issuer *string `pulumi:"issuer"`
	// The project for which this token is created. Enter the project name on which you want to apply this token.
	ProjectKey *string `pulumi:"projectKey"`
	// Reference Token (alias to Access Token).
	ReferenceToken *string `pulumi:"referenceToken"`
	// Refresh token.
	RefreshToken *string `pulumi:"refreshToken"`
	// Is this token refreshable? Default is `false`.
	Refreshable *bool `pulumi:"refreshable"`
	// The scope of access that the token provides. Access to the REST API is always provided by default. Administrators can set any scope, while non-admin users can only set the scope to a subset of the groups to which they belong.
	// The supported scopes include:
	Scopes []string `pulumi:"scopes"`
	// Returns the token type.
	Subject *string `pulumi:"subject"`
	// Returns the token type.
	TokenType *string `pulumi:"tokenType"`
	// The user name for which this token is created. The username is based on the authenticated user - either from the user of the authenticated token or based on the username (if basic auth was used). The username is then used to set the subject of the token: \n\n/users/\n\n. Limited to 255 characters.
	Username *string `pulumi:"username"`
}

type ScopedTokenState struct {
	// Returns the access token to authenticate to Artifactory.
	AccessToken pulumi.StringPtrInput
	// A list of the other instances or services that should accept this token identified by their Service-IDs. Limited to total 255 characters. Default to '*@*' if not set. Service ID must begin with valid JFrog service type. Options: jfrt, jfxr, jfpip, jfds, jfmc, jfac, jfevt, jfmd, jfcon, or *. For instructions to retrieve the Artifactory Service ID see this [documentation](https://jfrog.com/help/r/jfrog-rest-apis/get-service-id)
	Audiences pulumi.StringArrayInput
	// Free text token description. Useful for filtering and managing tokens. Limited to 1024 characters.
	Description pulumi.StringPtrInput
	// The amount of time, in seconds, it would take for the token to expire. An admin shall be able to set whether expiry is mandatory, what is the default expiry, and what is the maximum expiry allowed. Must be non-negative. Default value is based on configuration in 'access.config.yaml'. See [API documentation](https://jfrog.com/help/r/jfrog-rest-apis/create-token) for details. Access Token would not be saved by Artifactory if this is less than the persistence threshold value (default to 10800 seconds) set in Access configuration. See [official documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/using-the-revocable-and-persistency-thresholds) for details.
	ExpiresIn pulumi.IntPtrInput
	// Returns the token expiry.
	Expiry pulumi.IntPtrInput
	// The grant type used to authenticate the request. In this case, the only value supported is `clientCredentials` which is also the default value if this parameter is not specified.
	GrantType pulumi.StringPtrInput
	// Also create a reference token which can be used like an API key.
	IncludeReferenceToken pulumi.BoolPtrInput
	// Returns the token issued at date/time.
	IssuedAt pulumi.IntPtrInput
	// Returns the token issuer.
	Issuer pulumi.StringPtrInput
	// The project for which this token is created. Enter the project name on which you want to apply this token.
	ProjectKey pulumi.StringPtrInput
	// Reference Token (alias to Access Token).
	ReferenceToken pulumi.StringPtrInput
	// Refresh token.
	RefreshToken pulumi.StringPtrInput
	// Is this token refreshable? Default is `false`.
	Refreshable pulumi.BoolPtrInput
	// The scope of access that the token provides. Access to the REST API is always provided by default. Administrators can set any scope, while non-admin users can only set the scope to a subset of the groups to which they belong.
	// The supported scopes include:
	Scopes pulumi.StringArrayInput
	// Returns the token type.
	Subject pulumi.StringPtrInput
	// Returns the token type.
	TokenType pulumi.StringPtrInput
	// The user name for which this token is created. The username is based on the authenticated user - either from the user of the authenticated token or based on the username (if basic auth was used). The username is then used to set the subject of the token: \n\n/users/\n\n. Limited to 255 characters.
	Username pulumi.StringPtrInput
}

func (ScopedTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*scopedTokenState)(nil)).Elem()
}

type scopedTokenArgs struct {
	// A list of the other instances or services that should accept this token identified by their Service-IDs. Limited to total 255 characters. Default to '*@*' if not set. Service ID must begin with valid JFrog service type. Options: jfrt, jfxr, jfpip, jfds, jfmc, jfac, jfevt, jfmd, jfcon, or *. For instructions to retrieve the Artifactory Service ID see this [documentation](https://jfrog.com/help/r/jfrog-rest-apis/get-service-id)
	Audiences []string `pulumi:"audiences"`
	// Free text token description. Useful for filtering and managing tokens. Limited to 1024 characters.
	Description *string `pulumi:"description"`
	// The amount of time, in seconds, it would take for the token to expire. An admin shall be able to set whether expiry is mandatory, what is the default expiry, and what is the maximum expiry allowed. Must be non-negative. Default value is based on configuration in 'access.config.yaml'. See [API documentation](https://jfrog.com/help/r/jfrog-rest-apis/create-token) for details. Access Token would not be saved by Artifactory if this is less than the persistence threshold value (default to 10800 seconds) set in Access configuration. See [official documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/using-the-revocable-and-persistency-thresholds) for details.
	ExpiresIn *int `pulumi:"expiresIn"`
	// The grant type used to authenticate the request. In this case, the only value supported is `clientCredentials` which is also the default value if this parameter is not specified.
	GrantType *string `pulumi:"grantType"`
	// Also create a reference token which can be used like an API key.
	IncludeReferenceToken *bool `pulumi:"includeReferenceToken"`
	// The project for which this token is created. Enter the project name on which you want to apply this token.
	ProjectKey *string `pulumi:"projectKey"`
	// Is this token refreshable? Default is `false`.
	Refreshable *bool `pulumi:"refreshable"`
	// The scope of access that the token provides. Access to the REST API is always provided by default. Administrators can set any scope, while non-admin users can only set the scope to a subset of the groups to which they belong.
	// The supported scopes include:
	Scopes []string `pulumi:"scopes"`
	// The user name for which this token is created. The username is based on the authenticated user - either from the user of the authenticated token or based on the username (if basic auth was used). The username is then used to set the subject of the token: \n\n/users/\n\n. Limited to 255 characters.
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a ScopedToken resource.
type ScopedTokenArgs struct {
	// A list of the other instances or services that should accept this token identified by their Service-IDs. Limited to total 255 characters. Default to '*@*' if not set. Service ID must begin with valid JFrog service type. Options: jfrt, jfxr, jfpip, jfds, jfmc, jfac, jfevt, jfmd, jfcon, or *. For instructions to retrieve the Artifactory Service ID see this [documentation](https://jfrog.com/help/r/jfrog-rest-apis/get-service-id)
	Audiences pulumi.StringArrayInput
	// Free text token description. Useful for filtering and managing tokens. Limited to 1024 characters.
	Description pulumi.StringPtrInput
	// The amount of time, in seconds, it would take for the token to expire. An admin shall be able to set whether expiry is mandatory, what is the default expiry, and what is the maximum expiry allowed. Must be non-negative. Default value is based on configuration in 'access.config.yaml'. See [API documentation](https://jfrog.com/help/r/jfrog-rest-apis/create-token) for details. Access Token would not be saved by Artifactory if this is less than the persistence threshold value (default to 10800 seconds) set in Access configuration. See [official documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/using-the-revocable-and-persistency-thresholds) for details.
	ExpiresIn pulumi.IntPtrInput
	// The grant type used to authenticate the request. In this case, the only value supported is `clientCredentials` which is also the default value if this parameter is not specified.
	GrantType pulumi.StringPtrInput
	// Also create a reference token which can be used like an API key.
	IncludeReferenceToken pulumi.BoolPtrInput
	// The project for which this token is created. Enter the project name on which you want to apply this token.
	ProjectKey pulumi.StringPtrInput
	// Is this token refreshable? Default is `false`.
	Refreshable pulumi.BoolPtrInput
	// The scope of access that the token provides. Access to the REST API is always provided by default. Administrators can set any scope, while non-admin users can only set the scope to a subset of the groups to which they belong.
	// The supported scopes include:
	Scopes pulumi.StringArrayInput
	// The user name for which this token is created. The username is based on the authenticated user - either from the user of the authenticated token or based on the username (if basic auth was used). The username is then used to set the subject of the token: \n\n/users/\n\n. Limited to 255 characters.
	Username pulumi.StringPtrInput
}

func (ScopedTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scopedTokenArgs)(nil)).Elem()
}

type ScopedTokenInput interface {
	pulumi.Input

	ToScopedTokenOutput() ScopedTokenOutput
	ToScopedTokenOutputWithContext(ctx context.Context) ScopedTokenOutput
}

func (*ScopedToken) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopedToken)(nil)).Elem()
}

func (i *ScopedToken) ToScopedTokenOutput() ScopedTokenOutput {
	return i.ToScopedTokenOutputWithContext(context.Background())
}

func (i *ScopedToken) ToScopedTokenOutputWithContext(ctx context.Context) ScopedTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopedTokenOutput)
}

// ScopedTokenArrayInput is an input type that accepts ScopedTokenArray and ScopedTokenArrayOutput values.
// You can construct a concrete instance of `ScopedTokenArrayInput` via:
//
//	ScopedTokenArray{ ScopedTokenArgs{...} }
type ScopedTokenArrayInput interface {
	pulumi.Input

	ToScopedTokenArrayOutput() ScopedTokenArrayOutput
	ToScopedTokenArrayOutputWithContext(context.Context) ScopedTokenArrayOutput
}

type ScopedTokenArray []ScopedTokenInput

func (ScopedTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScopedToken)(nil)).Elem()
}

func (i ScopedTokenArray) ToScopedTokenArrayOutput() ScopedTokenArrayOutput {
	return i.ToScopedTokenArrayOutputWithContext(context.Background())
}

func (i ScopedTokenArray) ToScopedTokenArrayOutputWithContext(ctx context.Context) ScopedTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopedTokenArrayOutput)
}

// ScopedTokenMapInput is an input type that accepts ScopedTokenMap and ScopedTokenMapOutput values.
// You can construct a concrete instance of `ScopedTokenMapInput` via:
//
//	ScopedTokenMap{ "key": ScopedTokenArgs{...} }
type ScopedTokenMapInput interface {
	pulumi.Input

	ToScopedTokenMapOutput() ScopedTokenMapOutput
	ToScopedTokenMapOutputWithContext(context.Context) ScopedTokenMapOutput
}

type ScopedTokenMap map[string]ScopedTokenInput

func (ScopedTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScopedToken)(nil)).Elem()
}

func (i ScopedTokenMap) ToScopedTokenMapOutput() ScopedTokenMapOutput {
	return i.ToScopedTokenMapOutputWithContext(context.Background())
}

func (i ScopedTokenMap) ToScopedTokenMapOutputWithContext(ctx context.Context) ScopedTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopedTokenMapOutput)
}

type ScopedTokenOutput struct{ *pulumi.OutputState }

func (ScopedTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopedToken)(nil)).Elem()
}

func (o ScopedTokenOutput) ToScopedTokenOutput() ScopedTokenOutput {
	return o
}

func (o ScopedTokenOutput) ToScopedTokenOutputWithContext(ctx context.Context) ScopedTokenOutput {
	return o
}

// Returns the access token to authenticate to Artifactory.
func (o ScopedTokenOutput) AccessToken() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopedToken) pulumi.StringOutput { return v.AccessToken }).(pulumi.StringOutput)
}

// A list of the other instances or services that should accept this token identified by their Service-IDs. Limited to total 255 characters. Default to '*@*' if not set. Service ID must begin with valid JFrog service type. Options: jfrt, jfxr, jfpip, jfds, jfmc, jfac, jfevt, jfmd, jfcon, or *. For instructions to retrieve the Artifactory Service ID see this [documentation](https://jfrog.com/help/r/jfrog-rest-apis/get-service-id)
func (o ScopedTokenOutput) Audiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScopedToken) pulumi.StringArrayOutput { return v.Audiences }).(pulumi.StringArrayOutput)
}

// Free text token description. Useful for filtering and managing tokens. Limited to 1024 characters.
func (o ScopedTokenOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopedToken) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The amount of time, in seconds, it would take for the token to expire. An admin shall be able to set whether expiry is mandatory, what is the default expiry, and what is the maximum expiry allowed. Must be non-negative. Default value is based on configuration in 'access.config.yaml'. See [API documentation](https://jfrog.com/help/r/jfrog-rest-apis/create-token) for details. Access Token would not be saved by Artifactory if this is less than the persistence threshold value (default to 10800 seconds) set in Access configuration. See [official documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/using-the-revocable-and-persistency-thresholds) for details.
func (o ScopedTokenOutput) ExpiresIn() pulumi.IntOutput {
	return o.ApplyT(func(v *ScopedToken) pulumi.IntOutput { return v.ExpiresIn }).(pulumi.IntOutput)
}

// Returns the token expiry.
func (o ScopedTokenOutput) Expiry() pulumi.IntOutput {
	return o.ApplyT(func(v *ScopedToken) pulumi.IntOutput { return v.Expiry }).(pulumi.IntOutput)
}

// The grant type used to authenticate the request. In this case, the only value supported is `clientCredentials` which is also the default value if this parameter is not specified.
func (o ScopedTokenOutput) GrantType() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopedToken) pulumi.StringOutput { return v.GrantType }).(pulumi.StringOutput)
}

// Also create a reference token which can be used like an API key.
func (o ScopedTokenOutput) IncludeReferenceToken() pulumi.BoolOutput {
	return o.ApplyT(func(v *ScopedToken) pulumi.BoolOutput { return v.IncludeReferenceToken }).(pulumi.BoolOutput)
}

// Returns the token issued at date/time.
func (o ScopedTokenOutput) IssuedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *ScopedToken) pulumi.IntOutput { return v.IssuedAt }).(pulumi.IntOutput)
}

// Returns the token issuer.
func (o ScopedTokenOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopedToken) pulumi.StringOutput { return v.Issuer }).(pulumi.StringOutput)
}

// The project for which this token is created. Enter the project name on which you want to apply this token.
func (o ScopedTokenOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopedToken) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Reference Token (alias to Access Token).
func (o ScopedTokenOutput) ReferenceToken() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopedToken) pulumi.StringOutput { return v.ReferenceToken }).(pulumi.StringOutput)
}

// Refresh token.
func (o ScopedTokenOutput) RefreshToken() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopedToken) pulumi.StringOutput { return v.RefreshToken }).(pulumi.StringOutput)
}

// Is this token refreshable? Default is `false`.
func (o ScopedTokenOutput) Refreshable() pulumi.BoolOutput {
	return o.ApplyT(func(v *ScopedToken) pulumi.BoolOutput { return v.Refreshable }).(pulumi.BoolOutput)
}

// The scope of access that the token provides. Access to the REST API is always provided by default. Administrators can set any scope, while non-admin users can only set the scope to a subset of the groups to which they belong.
// The supported scopes include:
func (o ScopedTokenOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScopedToken) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

// Returns the token type.
func (o ScopedTokenOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopedToken) pulumi.StringOutput { return v.Subject }).(pulumi.StringOutput)
}

// Returns the token type.
func (o ScopedTokenOutput) TokenType() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopedToken) pulumi.StringOutput { return v.TokenType }).(pulumi.StringOutput)
}

// The user name for which this token is created. The username is based on the authenticated user - either from the user of the authenticated token or based on the username (if basic auth was used). The username is then used to set the subject of the token: \n\n/users/\n\n. Limited to 255 characters.
func (o ScopedTokenOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopedToken) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

type ScopedTokenArrayOutput struct{ *pulumi.OutputState }

func (ScopedTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScopedToken)(nil)).Elem()
}

func (o ScopedTokenArrayOutput) ToScopedTokenArrayOutput() ScopedTokenArrayOutput {
	return o
}

func (o ScopedTokenArrayOutput) ToScopedTokenArrayOutputWithContext(ctx context.Context) ScopedTokenArrayOutput {
	return o
}

func (o ScopedTokenArrayOutput) Index(i pulumi.IntInput) ScopedTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ScopedToken {
		return vs[0].([]*ScopedToken)[vs[1].(int)]
	}).(ScopedTokenOutput)
}

type ScopedTokenMapOutput struct{ *pulumi.OutputState }

func (ScopedTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScopedToken)(nil)).Elem()
}

func (o ScopedTokenMapOutput) ToScopedTokenMapOutput() ScopedTokenMapOutput {
	return o
}

func (o ScopedTokenMapOutput) ToScopedTokenMapOutputWithContext(ctx context.Context) ScopedTokenMapOutput {
	return o
}

func (o ScopedTokenMapOutput) MapIndex(k pulumi.StringInput) ScopedTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ScopedToken {
		return vs[0].(map[string]*ScopedToken)[vs[1].(string)]
	}).(ScopedTokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScopedTokenInput)(nil)).Elem(), &ScopedToken{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScopedTokenArrayInput)(nil)).Elem(), ScopedTokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScopedTokenMapInput)(nil)).Elem(), ScopedTokenMap{})
	pulumi.RegisterOutputType(ScopedTokenOutput{})
	pulumi.RegisterOutputType(ScopedTokenArrayOutput{})
	pulumi.RegisterOutputType(ScopedTokenMapOutput{})
}
