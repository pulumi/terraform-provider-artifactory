// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-artifactory/sdk/v4/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Artifactory Distribution Public Key resource. This can be used to create and manage Artifactory Distribution Public Keys.
//
// See [API description](https://jfrog.com/help/r/jfrog-rest-apis/set-distributionpublic-gpg-key) in the Artifactory documentation for more details. Also the [UI documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/managing-webstart-and-jar-signing) has further details on where to find these keys in Artifactory.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"os"
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v4/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := os.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewDistributionPublicKey(ctx, "my-key", &artifactory.DistributionPublicKeyArgs{
//				Alias:     pulumi.String("my-key"),
//				PublicKey: readFileOrPanic("samples/rsa.pub"),
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
// Distribution Public Key can be imported using the key id, e.g.
//
// ```sh
//
//	$ pulumi import artifactory:index/distributionPublicKey:DistributionPublicKey my-key keyid
//
// ```
type DistributionPublicKey struct {
	pulumi.CustomResourceState

	// Will be used as an identifier when uploading/retrieving the public key via REST API.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// Returns the computed key fingerprint
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// Returns the name and eMail address of issuer
	IssuedBy pulumi.StringOutput `pulumi:"issuedBy"`
	// Returns the date/time when this GPG key was created
	IssuedOn pulumi.StringOutput `pulumi:"issuedOn"`
	// Returns the key id by which this key is referenced in Artifactory
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// The Public key to add as a trusted distribution GPG key.
	//
	// The following additional attributes are exported:
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
	// Returns the date/time when this GPG key expires.
	ValidUntil pulumi.StringOutput `pulumi:"validUntil"`
}

// NewDistributionPublicKey registers a new resource with the given unique name, arguments, and options.
func NewDistributionPublicKey(ctx *pulumi.Context,
	name string, args *DistributionPublicKeyArgs, opts ...pulumi.ResourceOption) (*DistributionPublicKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Alias == nil {
		return nil, errors.New("invalid value for required argument 'Alias'")
	}
	if args.PublicKey == nil {
		return nil, errors.New("invalid value for required argument 'PublicKey'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DistributionPublicKey
	err := ctx.RegisterResource("artifactory:index/distributionPublicKey:DistributionPublicKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDistributionPublicKey gets an existing DistributionPublicKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDistributionPublicKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DistributionPublicKeyState, opts ...pulumi.ResourceOption) (*DistributionPublicKey, error) {
	var resource DistributionPublicKey
	err := ctx.ReadResource("artifactory:index/distributionPublicKey:DistributionPublicKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DistributionPublicKey resources.
type distributionPublicKeyState struct {
	// Will be used as an identifier when uploading/retrieving the public key via REST API.
	Alias *string `pulumi:"alias"`
	// Returns the computed key fingerprint
	Fingerprint *string `pulumi:"fingerprint"`
	// Returns the name and eMail address of issuer
	IssuedBy *string `pulumi:"issuedBy"`
	// Returns the date/time when this GPG key was created
	IssuedOn *string `pulumi:"issuedOn"`
	// Returns the key id by which this key is referenced in Artifactory
	KeyId *string `pulumi:"keyId"`
	// The Public key to add as a trusted distribution GPG key.
	//
	// The following additional attributes are exported:
	PublicKey *string `pulumi:"publicKey"`
	// Returns the date/time when this GPG key expires.
	ValidUntil *string `pulumi:"validUntil"`
}

type DistributionPublicKeyState struct {
	// Will be used as an identifier when uploading/retrieving the public key via REST API.
	Alias pulumi.StringPtrInput
	// Returns the computed key fingerprint
	Fingerprint pulumi.StringPtrInput
	// Returns the name and eMail address of issuer
	IssuedBy pulumi.StringPtrInput
	// Returns the date/time when this GPG key was created
	IssuedOn pulumi.StringPtrInput
	// Returns the key id by which this key is referenced in Artifactory
	KeyId pulumi.StringPtrInput
	// The Public key to add as a trusted distribution GPG key.
	//
	// The following additional attributes are exported:
	PublicKey pulumi.StringPtrInput
	// Returns the date/time when this GPG key expires.
	ValidUntil pulumi.StringPtrInput
}

func (DistributionPublicKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*distributionPublicKeyState)(nil)).Elem()
}

type distributionPublicKeyArgs struct {
	// Will be used as an identifier when uploading/retrieving the public key via REST API.
	Alias string `pulumi:"alias"`
	// The Public key to add as a trusted distribution GPG key.
	//
	// The following additional attributes are exported:
	PublicKey string `pulumi:"publicKey"`
}

// The set of arguments for constructing a DistributionPublicKey resource.
type DistributionPublicKeyArgs struct {
	// Will be used as an identifier when uploading/retrieving the public key via REST API.
	Alias pulumi.StringInput
	// The Public key to add as a trusted distribution GPG key.
	//
	// The following additional attributes are exported:
	PublicKey pulumi.StringInput
}

func (DistributionPublicKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*distributionPublicKeyArgs)(nil)).Elem()
}

type DistributionPublicKeyInput interface {
	pulumi.Input

	ToDistributionPublicKeyOutput() DistributionPublicKeyOutput
	ToDistributionPublicKeyOutputWithContext(ctx context.Context) DistributionPublicKeyOutput
}

func (*DistributionPublicKey) ElementType() reflect.Type {
	return reflect.TypeOf((**DistributionPublicKey)(nil)).Elem()
}

func (i *DistributionPublicKey) ToDistributionPublicKeyOutput() DistributionPublicKeyOutput {
	return i.ToDistributionPublicKeyOutputWithContext(context.Background())
}

func (i *DistributionPublicKey) ToDistributionPublicKeyOutputWithContext(ctx context.Context) DistributionPublicKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DistributionPublicKeyOutput)
}

// DistributionPublicKeyArrayInput is an input type that accepts DistributionPublicKeyArray and DistributionPublicKeyArrayOutput values.
// You can construct a concrete instance of `DistributionPublicKeyArrayInput` via:
//
//	DistributionPublicKeyArray{ DistributionPublicKeyArgs{...} }
type DistributionPublicKeyArrayInput interface {
	pulumi.Input

	ToDistributionPublicKeyArrayOutput() DistributionPublicKeyArrayOutput
	ToDistributionPublicKeyArrayOutputWithContext(context.Context) DistributionPublicKeyArrayOutput
}

type DistributionPublicKeyArray []DistributionPublicKeyInput

func (DistributionPublicKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DistributionPublicKey)(nil)).Elem()
}

func (i DistributionPublicKeyArray) ToDistributionPublicKeyArrayOutput() DistributionPublicKeyArrayOutput {
	return i.ToDistributionPublicKeyArrayOutputWithContext(context.Background())
}

func (i DistributionPublicKeyArray) ToDistributionPublicKeyArrayOutputWithContext(ctx context.Context) DistributionPublicKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DistributionPublicKeyArrayOutput)
}

// DistributionPublicKeyMapInput is an input type that accepts DistributionPublicKeyMap and DistributionPublicKeyMapOutput values.
// You can construct a concrete instance of `DistributionPublicKeyMapInput` via:
//
//	DistributionPublicKeyMap{ "key": DistributionPublicKeyArgs{...} }
type DistributionPublicKeyMapInput interface {
	pulumi.Input

	ToDistributionPublicKeyMapOutput() DistributionPublicKeyMapOutput
	ToDistributionPublicKeyMapOutputWithContext(context.Context) DistributionPublicKeyMapOutput
}

type DistributionPublicKeyMap map[string]DistributionPublicKeyInput

func (DistributionPublicKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DistributionPublicKey)(nil)).Elem()
}

func (i DistributionPublicKeyMap) ToDistributionPublicKeyMapOutput() DistributionPublicKeyMapOutput {
	return i.ToDistributionPublicKeyMapOutputWithContext(context.Background())
}

func (i DistributionPublicKeyMap) ToDistributionPublicKeyMapOutputWithContext(ctx context.Context) DistributionPublicKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DistributionPublicKeyMapOutput)
}

type DistributionPublicKeyOutput struct{ *pulumi.OutputState }

func (DistributionPublicKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DistributionPublicKey)(nil)).Elem()
}

func (o DistributionPublicKeyOutput) ToDistributionPublicKeyOutput() DistributionPublicKeyOutput {
	return o
}

func (o DistributionPublicKeyOutput) ToDistributionPublicKeyOutputWithContext(ctx context.Context) DistributionPublicKeyOutput {
	return o
}

// Will be used as an identifier when uploading/retrieving the public key via REST API.
func (o DistributionPublicKeyOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributionPublicKey) pulumi.StringOutput { return v.Alias }).(pulumi.StringOutput)
}

// Returns the computed key fingerprint
func (o DistributionPublicKeyOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributionPublicKey) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// Returns the name and eMail address of issuer
func (o DistributionPublicKeyOutput) IssuedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributionPublicKey) pulumi.StringOutput { return v.IssuedBy }).(pulumi.StringOutput)
}

// Returns the date/time when this GPG key was created
func (o DistributionPublicKeyOutput) IssuedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributionPublicKey) pulumi.StringOutput { return v.IssuedOn }).(pulumi.StringOutput)
}

// Returns the key id by which this key is referenced in Artifactory
func (o DistributionPublicKeyOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributionPublicKey) pulumi.StringOutput { return v.KeyId }).(pulumi.StringOutput)
}

// The Public key to add as a trusted distribution GPG key.
//
// The following additional attributes are exported:
func (o DistributionPublicKeyOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributionPublicKey) pulumi.StringOutput { return v.PublicKey }).(pulumi.StringOutput)
}

// Returns the date/time when this GPG key expires.
func (o DistributionPublicKeyOutput) ValidUntil() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributionPublicKey) pulumi.StringOutput { return v.ValidUntil }).(pulumi.StringOutput)
}

type DistributionPublicKeyArrayOutput struct{ *pulumi.OutputState }

func (DistributionPublicKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DistributionPublicKey)(nil)).Elem()
}

func (o DistributionPublicKeyArrayOutput) ToDistributionPublicKeyArrayOutput() DistributionPublicKeyArrayOutput {
	return o
}

func (o DistributionPublicKeyArrayOutput) ToDistributionPublicKeyArrayOutputWithContext(ctx context.Context) DistributionPublicKeyArrayOutput {
	return o
}

func (o DistributionPublicKeyArrayOutput) Index(i pulumi.IntInput) DistributionPublicKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DistributionPublicKey {
		return vs[0].([]*DistributionPublicKey)[vs[1].(int)]
	}).(DistributionPublicKeyOutput)
}

type DistributionPublicKeyMapOutput struct{ *pulumi.OutputState }

func (DistributionPublicKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DistributionPublicKey)(nil)).Elem()
}

func (o DistributionPublicKeyMapOutput) ToDistributionPublicKeyMapOutput() DistributionPublicKeyMapOutput {
	return o
}

func (o DistributionPublicKeyMapOutput) ToDistributionPublicKeyMapOutputWithContext(ctx context.Context) DistributionPublicKeyMapOutput {
	return o
}

func (o DistributionPublicKeyMapOutput) MapIndex(k pulumi.StringInput) DistributionPublicKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DistributionPublicKey {
		return vs[0].(map[string]*DistributionPublicKey)[vs[1].(string)]
	}).(DistributionPublicKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DistributionPublicKeyInput)(nil)).Elem(), &DistributionPublicKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*DistributionPublicKeyArrayInput)(nil)).Elem(), DistributionPublicKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DistributionPublicKeyMapInput)(nil)).Elem(), DistributionPublicKeyMap{})
	pulumi.RegisterOutputType(DistributionPublicKeyOutput{})
	pulumi.RegisterOutputType(DistributionPublicKeyArrayOutput{})
	pulumi.RegisterOutputType(DistributionPublicKeyMapOutput{})
}
