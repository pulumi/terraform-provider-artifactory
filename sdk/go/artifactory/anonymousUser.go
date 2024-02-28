// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v6/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
//			// Define a new Artifactory 'anonymous' user for import
//			_, err := artifactory.NewAnonymousUser(ctx, "anonymous", nil)
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
//	$ pulumi import artifactory:index/anonymousUser:AnonymousUser anonymous-user anonymous
//
// ```
type AnonymousUser struct {
	pulumi.CustomResourceState

	// Username for anonymous user. This is only for ensuring resource schema is valid for Terraform. This is not meant to be
	// set or updated in the HCL.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewAnonymousUser registers a new resource with the given unique name, arguments, and options.
func NewAnonymousUser(ctx *pulumi.Context,
	name string, args *AnonymousUserArgs, opts ...pulumi.ResourceOption) (*AnonymousUser, error) {
	if args == nil {
		args = &AnonymousUserArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AnonymousUser
	err := ctx.RegisterResource("artifactory:index/anonymousUser:AnonymousUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnonymousUser gets an existing AnonymousUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnonymousUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnonymousUserState, opts ...pulumi.ResourceOption) (*AnonymousUser, error) {
	var resource AnonymousUser
	err := ctx.ReadResource("artifactory:index/anonymousUser:AnonymousUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AnonymousUser resources.
type anonymousUserState struct {
	// Username for anonymous user. This is only for ensuring resource schema is valid for Terraform. This is not meant to be
	// set or updated in the HCL.
	Name *string `pulumi:"name"`
}

type AnonymousUserState struct {
	// Username for anonymous user. This is only for ensuring resource schema is valid for Terraform. This is not meant to be
	// set or updated in the HCL.
	Name pulumi.StringPtrInput
}

func (AnonymousUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*anonymousUserState)(nil)).Elem()
}

type anonymousUserArgs struct {
	// Username for anonymous user. This is only for ensuring resource schema is valid for Terraform. This is not meant to be
	// set or updated in the HCL.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a AnonymousUser resource.
type AnonymousUserArgs struct {
	// Username for anonymous user. This is only for ensuring resource schema is valid for Terraform. This is not meant to be
	// set or updated in the HCL.
	Name pulumi.StringPtrInput
}

func (AnonymousUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*anonymousUserArgs)(nil)).Elem()
}

type AnonymousUserInput interface {
	pulumi.Input

	ToAnonymousUserOutput() AnonymousUserOutput
	ToAnonymousUserOutputWithContext(ctx context.Context) AnonymousUserOutput
}

func (*AnonymousUser) ElementType() reflect.Type {
	return reflect.TypeOf((**AnonymousUser)(nil)).Elem()
}

func (i *AnonymousUser) ToAnonymousUserOutput() AnonymousUserOutput {
	return i.ToAnonymousUserOutputWithContext(context.Background())
}

func (i *AnonymousUser) ToAnonymousUserOutputWithContext(ctx context.Context) AnonymousUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnonymousUserOutput)
}

// AnonymousUserArrayInput is an input type that accepts AnonymousUserArray and AnonymousUserArrayOutput values.
// You can construct a concrete instance of `AnonymousUserArrayInput` via:
//
//	AnonymousUserArray{ AnonymousUserArgs{...} }
type AnonymousUserArrayInput interface {
	pulumi.Input

	ToAnonymousUserArrayOutput() AnonymousUserArrayOutput
	ToAnonymousUserArrayOutputWithContext(context.Context) AnonymousUserArrayOutput
}

type AnonymousUserArray []AnonymousUserInput

func (AnonymousUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AnonymousUser)(nil)).Elem()
}

func (i AnonymousUserArray) ToAnonymousUserArrayOutput() AnonymousUserArrayOutput {
	return i.ToAnonymousUserArrayOutputWithContext(context.Background())
}

func (i AnonymousUserArray) ToAnonymousUserArrayOutputWithContext(ctx context.Context) AnonymousUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnonymousUserArrayOutput)
}

// AnonymousUserMapInput is an input type that accepts AnonymousUserMap and AnonymousUserMapOutput values.
// You can construct a concrete instance of `AnonymousUserMapInput` via:
//
//	AnonymousUserMap{ "key": AnonymousUserArgs{...} }
type AnonymousUserMapInput interface {
	pulumi.Input

	ToAnonymousUserMapOutput() AnonymousUserMapOutput
	ToAnonymousUserMapOutputWithContext(context.Context) AnonymousUserMapOutput
}

type AnonymousUserMap map[string]AnonymousUserInput

func (AnonymousUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AnonymousUser)(nil)).Elem()
}

func (i AnonymousUserMap) ToAnonymousUserMapOutput() AnonymousUserMapOutput {
	return i.ToAnonymousUserMapOutputWithContext(context.Background())
}

func (i AnonymousUserMap) ToAnonymousUserMapOutputWithContext(ctx context.Context) AnonymousUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnonymousUserMapOutput)
}

type AnonymousUserOutput struct{ *pulumi.OutputState }

func (AnonymousUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AnonymousUser)(nil)).Elem()
}

func (o AnonymousUserOutput) ToAnonymousUserOutput() AnonymousUserOutput {
	return o
}

func (o AnonymousUserOutput) ToAnonymousUserOutputWithContext(ctx context.Context) AnonymousUserOutput {
	return o
}

// Username for anonymous user. This is only for ensuring resource schema is valid for Terraform. This is not meant to be
// set or updated in the HCL.
func (o AnonymousUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AnonymousUser) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type AnonymousUserArrayOutput struct{ *pulumi.OutputState }

func (AnonymousUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AnonymousUser)(nil)).Elem()
}

func (o AnonymousUserArrayOutput) ToAnonymousUserArrayOutput() AnonymousUserArrayOutput {
	return o
}

func (o AnonymousUserArrayOutput) ToAnonymousUserArrayOutputWithContext(ctx context.Context) AnonymousUserArrayOutput {
	return o
}

func (o AnonymousUserArrayOutput) Index(i pulumi.IntInput) AnonymousUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AnonymousUser {
		return vs[0].([]*AnonymousUser)[vs[1].(int)]
	}).(AnonymousUserOutput)
}

type AnonymousUserMapOutput struct{ *pulumi.OutputState }

func (AnonymousUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AnonymousUser)(nil)).Elem()
}

func (o AnonymousUserMapOutput) ToAnonymousUserMapOutput() AnonymousUserMapOutput {
	return o
}

func (o AnonymousUserMapOutput) ToAnonymousUserMapOutputWithContext(ctx context.Context) AnonymousUserMapOutput {
	return o
}

func (o AnonymousUserMapOutput) MapIndex(k pulumi.StringInput) AnonymousUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AnonymousUser {
		return vs[0].(map[string]*AnonymousUser)[vs[1].(string)]
	}).(AnonymousUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AnonymousUserInput)(nil)).Elem(), &AnonymousUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*AnonymousUserArrayInput)(nil)).Elem(), AnonymousUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AnonymousUserMapInput)(nil)).Elem(), AnonymousUserMap{})
	pulumi.RegisterOutputType(AnonymousUserOutput{})
	pulumi.RegisterOutputType(AnonymousUserArrayOutput{})
	pulumi.RegisterOutputType(AnonymousUserMapOutput{})
}
