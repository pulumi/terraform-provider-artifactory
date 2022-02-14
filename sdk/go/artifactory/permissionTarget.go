// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Permission Target Resource
//
// **Requires Artifactory >= 6.6.0. If using a lower version see here**
//
// Provides an Artifactory permission target resource. This can be used to create and manage Artifactory permission targets.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-artifactory/sdk/go/artifactory"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := artifactory.NewPermissionTarget(ctx, "test-perm", &artifactory.PermissionTargetArgs{
// 			Build: &PermissionTargetBuildArgs{
// 				Actions: &PermissionTargetBuildActionsArgs{
// 					Users: PermissionTargetBuildActionsUserArray{
// 						&PermissionTargetBuildActionsUserArgs{
// 							Name: pulumi.String("anonymous"),
// 							Permissions: pulumi.StringArray{
// 								pulumi.String("read"),
// 								pulumi.String("write"),
// 							},
// 						},
// 					},
// 				},
// 				IncludesPatterns: pulumi.StringArray{
// 					pulumi.String("**"),
// 				},
// 				Repositories: pulumi.StringArray{
// 					pulumi.String("artifactory-build-info"),
// 				},
// 			},
// 			Repo: &PermissionTargetRepoArgs{
// 				Actions: &PermissionTargetRepoActionsArgs{
// 					Groups: PermissionTargetRepoActionsGroupArray{
// 						&PermissionTargetRepoActionsGroupArgs{
// 							Name: pulumi.String("readers"),
// 							Permissions: pulumi.StringArray{
// 								pulumi.String("read"),
// 							},
// 						},
// 					},
// 					Users: PermissionTargetRepoActionsUserArray{
// 						&PermissionTargetRepoActionsUserArgs{
// 							Name: pulumi.String("anonymous"),
// 							Permissions: pulumi.StringArray{
// 								pulumi.String("read"),
// 								pulumi.String("write"),
// 							},
// 						},
// 					},
// 				},
// 				ExcludesPatterns: pulumi.StringArray{
// 					pulumi.String("bar/**"),
// 				},
// 				IncludesPatterns: pulumi.StringArray{
// 					pulumi.String("foo/**"),
// 				},
// 				Repositories: pulumi.StringArray{
// 					pulumi.String("example-repo-local"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Permissions
//
// The provider supports the following `permission` enums:
//
// * `read`
// * `write`
// * `annotate`
// * `delete`
// * `manage`
//
// The values can be mapped to the permissions from the official [documentation](https://www.jfrog.com/confluence/display/JFROG/Permissions):
//
// * `read` - matches `Read` permissions.
// * `write` - matches `  Deploy / Cache / Create ` permissions.
// * `annotate` - matches `Annotate` permissions.
// * `delete` - matches `Delete / Overwrite` permissions.
// * `manage` - matches `Manage` permissions.
//
// ## Import
//
// Permission targets can be imported using their name, e.g.
//
// ```sh
//  $ pulumi import artifactory:index/permissionTarget:PermissionTarget terraform-test-permission mypermission
// ```
type PermissionTarget struct {
	pulumi.CustomResourceState

	// As for repo but for artifactory-build-info permssions.
	Build PermissionTargetBuildPtrOutput `pulumi:"build"`
	// Name of permission
	Name          pulumi.StringOutput                    `pulumi:"name"`
	ReleaseBundle PermissionTargetReleaseBundlePtrOutput `pulumi:"releaseBundle"`
	// Repository permission configuration
	Repo PermissionTargetRepoPtrOutput `pulumi:"repo"`
}

// NewPermissionTarget registers a new resource with the given unique name, arguments, and options.
func NewPermissionTarget(ctx *pulumi.Context,
	name string, args *PermissionTargetArgs, opts ...pulumi.ResourceOption) (*PermissionTarget, error) {
	if args == nil {
		args = &PermissionTargetArgs{}
	}

	var resource PermissionTarget
	err := ctx.RegisterResource("artifactory:index/permissionTarget:PermissionTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPermissionTarget gets an existing PermissionTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPermissionTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PermissionTargetState, opts ...pulumi.ResourceOption) (*PermissionTarget, error) {
	var resource PermissionTarget
	err := ctx.ReadResource("artifactory:index/permissionTarget:PermissionTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PermissionTarget resources.
type permissionTargetState struct {
	// As for repo but for artifactory-build-info permssions.
	Build *PermissionTargetBuild `pulumi:"build"`
	// Name of permission
	Name          *string                        `pulumi:"name"`
	ReleaseBundle *PermissionTargetReleaseBundle `pulumi:"releaseBundle"`
	// Repository permission configuration
	Repo *PermissionTargetRepo `pulumi:"repo"`
}

type PermissionTargetState struct {
	// As for repo but for artifactory-build-info permssions.
	Build PermissionTargetBuildPtrInput
	// Name of permission
	Name          pulumi.StringPtrInput
	ReleaseBundle PermissionTargetReleaseBundlePtrInput
	// Repository permission configuration
	Repo PermissionTargetRepoPtrInput
}

func (PermissionTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionTargetState)(nil)).Elem()
}

type permissionTargetArgs struct {
	// As for repo but for artifactory-build-info permssions.
	Build *PermissionTargetBuild `pulumi:"build"`
	// Name of permission
	Name          *string                        `pulumi:"name"`
	ReleaseBundle *PermissionTargetReleaseBundle `pulumi:"releaseBundle"`
	// Repository permission configuration
	Repo *PermissionTargetRepo `pulumi:"repo"`
}

// The set of arguments for constructing a PermissionTarget resource.
type PermissionTargetArgs struct {
	// As for repo but for artifactory-build-info permssions.
	Build PermissionTargetBuildPtrInput
	// Name of permission
	Name          pulumi.StringPtrInput
	ReleaseBundle PermissionTargetReleaseBundlePtrInput
	// Repository permission configuration
	Repo PermissionTargetRepoPtrInput
}

func (PermissionTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionTargetArgs)(nil)).Elem()
}

type PermissionTargetInput interface {
	pulumi.Input

	ToPermissionTargetOutput() PermissionTargetOutput
	ToPermissionTargetOutputWithContext(ctx context.Context) PermissionTargetOutput
}

func (*PermissionTarget) ElementType() reflect.Type {
	return reflect.TypeOf((**PermissionTarget)(nil)).Elem()
}

func (i *PermissionTarget) ToPermissionTargetOutput() PermissionTargetOutput {
	return i.ToPermissionTargetOutputWithContext(context.Background())
}

func (i *PermissionTarget) ToPermissionTargetOutputWithContext(ctx context.Context) PermissionTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionTargetOutput)
}

// PermissionTargetArrayInput is an input type that accepts PermissionTargetArray and PermissionTargetArrayOutput values.
// You can construct a concrete instance of `PermissionTargetArrayInput` via:
//
//          PermissionTargetArray{ PermissionTargetArgs{...} }
type PermissionTargetArrayInput interface {
	pulumi.Input

	ToPermissionTargetArrayOutput() PermissionTargetArrayOutput
	ToPermissionTargetArrayOutputWithContext(context.Context) PermissionTargetArrayOutput
}

type PermissionTargetArray []PermissionTargetInput

func (PermissionTargetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PermissionTarget)(nil)).Elem()
}

func (i PermissionTargetArray) ToPermissionTargetArrayOutput() PermissionTargetArrayOutput {
	return i.ToPermissionTargetArrayOutputWithContext(context.Background())
}

func (i PermissionTargetArray) ToPermissionTargetArrayOutputWithContext(ctx context.Context) PermissionTargetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionTargetArrayOutput)
}

// PermissionTargetMapInput is an input type that accepts PermissionTargetMap and PermissionTargetMapOutput values.
// You can construct a concrete instance of `PermissionTargetMapInput` via:
//
//          PermissionTargetMap{ "key": PermissionTargetArgs{...} }
type PermissionTargetMapInput interface {
	pulumi.Input

	ToPermissionTargetMapOutput() PermissionTargetMapOutput
	ToPermissionTargetMapOutputWithContext(context.Context) PermissionTargetMapOutput
}

type PermissionTargetMap map[string]PermissionTargetInput

func (PermissionTargetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PermissionTarget)(nil)).Elem()
}

func (i PermissionTargetMap) ToPermissionTargetMapOutput() PermissionTargetMapOutput {
	return i.ToPermissionTargetMapOutputWithContext(context.Background())
}

func (i PermissionTargetMap) ToPermissionTargetMapOutputWithContext(ctx context.Context) PermissionTargetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionTargetMapOutput)
}

type PermissionTargetOutput struct{ *pulumi.OutputState }

func (PermissionTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PermissionTarget)(nil)).Elem()
}

func (o PermissionTargetOutput) ToPermissionTargetOutput() PermissionTargetOutput {
	return o
}

func (o PermissionTargetOutput) ToPermissionTargetOutputWithContext(ctx context.Context) PermissionTargetOutput {
	return o
}

type PermissionTargetArrayOutput struct{ *pulumi.OutputState }

func (PermissionTargetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PermissionTarget)(nil)).Elem()
}

func (o PermissionTargetArrayOutput) ToPermissionTargetArrayOutput() PermissionTargetArrayOutput {
	return o
}

func (o PermissionTargetArrayOutput) ToPermissionTargetArrayOutputWithContext(ctx context.Context) PermissionTargetArrayOutput {
	return o
}

func (o PermissionTargetArrayOutput) Index(i pulumi.IntInput) PermissionTargetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PermissionTarget {
		return vs[0].([]*PermissionTarget)[vs[1].(int)]
	}).(PermissionTargetOutput)
}

type PermissionTargetMapOutput struct{ *pulumi.OutputState }

func (PermissionTargetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PermissionTarget)(nil)).Elem()
}

func (o PermissionTargetMapOutput) ToPermissionTargetMapOutput() PermissionTargetMapOutput {
	return o
}

func (o PermissionTargetMapOutput) ToPermissionTargetMapOutputWithContext(ctx context.Context) PermissionTargetMapOutput {
	return o
}

func (o PermissionTargetMapOutput) MapIndex(k pulumi.StringInput) PermissionTargetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PermissionTarget {
		return vs[0].(map[string]*PermissionTarget)[vs[1].(string)]
	}).(PermissionTargetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionTargetInput)(nil)).Elem(), &PermissionTarget{})
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionTargetArrayInput)(nil)).Elem(), PermissionTargetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionTargetMapInput)(nil)).Elem(), PermissionTargetMap{})
	pulumi.RegisterOutputType(PermissionTargetOutput{})
	pulumi.RegisterOutputType(PermissionTargetArrayOutput{})
	pulumi.RegisterOutputType(PermissionTargetMapOutput{})
}
