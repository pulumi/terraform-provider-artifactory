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

// RSA key pairs are used to sign and verify the Alpine Linux index files in JFrog Artifactory, while GPG key pairs are
// used to sign and validate packages integrity in JFrog Distribution. The JFrog Platform enables you to manage multiple RSA and GPG signing keys through the Keys Management UI and REST API. The JFrog Platform supports managing multiple pairs of GPG signing keys to sign packages for authentication of several package types such as Debian, Opkg, and RPM through the Keys Management UI and REST API.
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
//	"github.com/pulumi/pulumi-artifactory/sdk/v5/go/artifactory"
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
//			_, err := artifactory.NewKeypair(ctx, "some-keypair-6543461672124900137", &artifactory.KeypairArgs{
//				PairName:   pulumi.String("some-keypair-6543461672124900137"),
//				PairType:   pulumi.String("RSA"),
//				Alias:      pulumi.String("some-alias-6543461672124900137"),
//				PrivateKey: readFileOrPanic("samples/rsa.priv"),
//				PublicKey:  readFileOrPanic("samples/rsa.pub"),
//				Passphrase: pulumi.String("PASSPHRASE"),
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
// Keypair can be imported using the pair name, e.g.
//
// ```sh
//
//	$ pulumi import artifactory:index/keypair:Keypair my-keypair my-keypair-name
//
// ```
type Keypair struct {
	pulumi.CustomResourceState

	// Will be used as a filename when retrieving the public key via REST API.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// A unique identifier for the Key Pair record.
	PairName pulumi.StringOutput `pulumi:"pairName"`
	// Key Pair type. Supported types - GPG and RSA.
	PairType pulumi.StringOutput `pulumi:"pairType"`
	// Passphrase will be used to decrypt the private key. Validated server side.
	Passphrase pulumi.StringPtrOutput `pulumi:"passphrase"`
	// Private key. PEM format will be validated. Must not include extranous spaces or tabs.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// Public key. PEM format will be validated. Must not include extranous spaces or tabs.
	//
	// Artifactory REST API call 'Get Key Pair' doesn't return attributes `privateKey` and `passphrase`, but consumes these keys in the POST call.
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
}

// NewKeypair registers a new resource with the given unique name, arguments, and options.
func NewKeypair(ctx *pulumi.Context,
	name string, args *KeypairArgs, opts ...pulumi.ResourceOption) (*Keypair, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Alias == nil {
		return nil, errors.New("invalid value for required argument 'Alias'")
	}
	if args.PairName == nil {
		return nil, errors.New("invalid value for required argument 'PairName'")
	}
	if args.PairType == nil {
		return nil, errors.New("invalid value for required argument 'PairType'")
	}
	if args.PrivateKey == nil {
		return nil, errors.New("invalid value for required argument 'PrivateKey'")
	}
	if args.PublicKey == nil {
		return nil, errors.New("invalid value for required argument 'PublicKey'")
	}
	if args.Passphrase != nil {
		args.Passphrase = pulumi.ToSecret(args.Passphrase).(pulumi.StringPtrInput)
	}
	if args.PrivateKey != nil {
		args.PrivateKey = pulumi.ToSecret(args.PrivateKey).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"passphrase",
		"privateKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Keypair
	err := ctx.RegisterResource("artifactory:index/keypair:Keypair", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeypair gets an existing Keypair resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeypair(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeypairState, opts ...pulumi.ResourceOption) (*Keypair, error) {
	var resource Keypair
	err := ctx.ReadResource("artifactory:index/keypair:Keypair", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Keypair resources.
type keypairState struct {
	// Will be used as a filename when retrieving the public key via REST API.
	Alias *string `pulumi:"alias"`
	// A unique identifier for the Key Pair record.
	PairName *string `pulumi:"pairName"`
	// Key Pair type. Supported types - GPG and RSA.
	PairType *string `pulumi:"pairType"`
	// Passphrase will be used to decrypt the private key. Validated server side.
	Passphrase *string `pulumi:"passphrase"`
	// Private key. PEM format will be validated. Must not include extranous spaces or tabs.
	PrivateKey *string `pulumi:"privateKey"`
	// Public key. PEM format will be validated. Must not include extranous spaces or tabs.
	//
	// Artifactory REST API call 'Get Key Pair' doesn't return attributes `privateKey` and `passphrase`, but consumes these keys in the POST call.
	PublicKey *string `pulumi:"publicKey"`
}

type KeypairState struct {
	// Will be used as a filename when retrieving the public key via REST API.
	Alias pulumi.StringPtrInput
	// A unique identifier for the Key Pair record.
	PairName pulumi.StringPtrInput
	// Key Pair type. Supported types - GPG and RSA.
	PairType pulumi.StringPtrInput
	// Passphrase will be used to decrypt the private key. Validated server side.
	Passphrase pulumi.StringPtrInput
	// Private key. PEM format will be validated. Must not include extranous spaces or tabs.
	PrivateKey pulumi.StringPtrInput
	// Public key. PEM format will be validated. Must not include extranous spaces or tabs.
	//
	// Artifactory REST API call 'Get Key Pair' doesn't return attributes `privateKey` and `passphrase`, but consumes these keys in the POST call.
	PublicKey pulumi.StringPtrInput
}

func (KeypairState) ElementType() reflect.Type {
	return reflect.TypeOf((*keypairState)(nil)).Elem()
}

type keypairArgs struct {
	// Will be used as a filename when retrieving the public key via REST API.
	Alias string `pulumi:"alias"`
	// A unique identifier for the Key Pair record.
	PairName string `pulumi:"pairName"`
	// Key Pair type. Supported types - GPG and RSA.
	PairType string `pulumi:"pairType"`
	// Passphrase will be used to decrypt the private key. Validated server side.
	Passphrase *string `pulumi:"passphrase"`
	// Private key. PEM format will be validated. Must not include extranous spaces or tabs.
	PrivateKey string `pulumi:"privateKey"`
	// Public key. PEM format will be validated. Must not include extranous spaces or tabs.
	//
	// Artifactory REST API call 'Get Key Pair' doesn't return attributes `privateKey` and `passphrase`, but consumes these keys in the POST call.
	PublicKey string `pulumi:"publicKey"`
}

// The set of arguments for constructing a Keypair resource.
type KeypairArgs struct {
	// Will be used as a filename when retrieving the public key via REST API.
	Alias pulumi.StringInput
	// A unique identifier for the Key Pair record.
	PairName pulumi.StringInput
	// Key Pair type. Supported types - GPG and RSA.
	PairType pulumi.StringInput
	// Passphrase will be used to decrypt the private key. Validated server side.
	Passphrase pulumi.StringPtrInput
	// Private key. PEM format will be validated. Must not include extranous spaces or tabs.
	PrivateKey pulumi.StringInput
	// Public key. PEM format will be validated. Must not include extranous spaces or tabs.
	//
	// Artifactory REST API call 'Get Key Pair' doesn't return attributes `privateKey` and `passphrase`, but consumes these keys in the POST call.
	PublicKey pulumi.StringInput
}

func (KeypairArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keypairArgs)(nil)).Elem()
}

type KeypairInput interface {
	pulumi.Input

	ToKeypairOutput() KeypairOutput
	ToKeypairOutputWithContext(ctx context.Context) KeypairOutput
}

func (*Keypair) ElementType() reflect.Type {
	return reflect.TypeOf((**Keypair)(nil)).Elem()
}

func (i *Keypair) ToKeypairOutput() KeypairOutput {
	return i.ToKeypairOutputWithContext(context.Background())
}

func (i *Keypair) ToKeypairOutputWithContext(ctx context.Context) KeypairOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeypairOutput)
}

func (i *Keypair) ToOutput(ctx context.Context) pulumix.Output[*Keypair] {
	return pulumix.Output[*Keypair]{
		OutputState: i.ToKeypairOutputWithContext(ctx).OutputState,
	}
}

// KeypairArrayInput is an input type that accepts KeypairArray and KeypairArrayOutput values.
// You can construct a concrete instance of `KeypairArrayInput` via:
//
//	KeypairArray{ KeypairArgs{...} }
type KeypairArrayInput interface {
	pulumi.Input

	ToKeypairArrayOutput() KeypairArrayOutput
	ToKeypairArrayOutputWithContext(context.Context) KeypairArrayOutput
}

type KeypairArray []KeypairInput

func (KeypairArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Keypair)(nil)).Elem()
}

func (i KeypairArray) ToKeypairArrayOutput() KeypairArrayOutput {
	return i.ToKeypairArrayOutputWithContext(context.Background())
}

func (i KeypairArray) ToKeypairArrayOutputWithContext(ctx context.Context) KeypairArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeypairArrayOutput)
}

func (i KeypairArray) ToOutput(ctx context.Context) pulumix.Output[[]*Keypair] {
	return pulumix.Output[[]*Keypair]{
		OutputState: i.ToKeypairArrayOutputWithContext(ctx).OutputState,
	}
}

// KeypairMapInput is an input type that accepts KeypairMap and KeypairMapOutput values.
// You can construct a concrete instance of `KeypairMapInput` via:
//
//	KeypairMap{ "key": KeypairArgs{...} }
type KeypairMapInput interface {
	pulumi.Input

	ToKeypairMapOutput() KeypairMapOutput
	ToKeypairMapOutputWithContext(context.Context) KeypairMapOutput
}

type KeypairMap map[string]KeypairInput

func (KeypairMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Keypair)(nil)).Elem()
}

func (i KeypairMap) ToKeypairMapOutput() KeypairMapOutput {
	return i.ToKeypairMapOutputWithContext(context.Background())
}

func (i KeypairMap) ToKeypairMapOutputWithContext(ctx context.Context) KeypairMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeypairMapOutput)
}

func (i KeypairMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Keypair] {
	return pulumix.Output[map[string]*Keypair]{
		OutputState: i.ToKeypairMapOutputWithContext(ctx).OutputState,
	}
}

type KeypairOutput struct{ *pulumi.OutputState }

func (KeypairOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Keypair)(nil)).Elem()
}

func (o KeypairOutput) ToKeypairOutput() KeypairOutput {
	return o
}

func (o KeypairOutput) ToKeypairOutputWithContext(ctx context.Context) KeypairOutput {
	return o
}

func (o KeypairOutput) ToOutput(ctx context.Context) pulumix.Output[*Keypair] {
	return pulumix.Output[*Keypair]{
		OutputState: o.OutputState,
	}
}

// Will be used as a filename when retrieving the public key via REST API.
func (o KeypairOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v *Keypair) pulumi.StringOutput { return v.Alias }).(pulumi.StringOutput)
}

// A unique identifier for the Key Pair record.
func (o KeypairOutput) PairName() pulumi.StringOutput {
	return o.ApplyT(func(v *Keypair) pulumi.StringOutput { return v.PairName }).(pulumi.StringOutput)
}

// Key Pair type. Supported types - GPG and RSA.
func (o KeypairOutput) PairType() pulumi.StringOutput {
	return o.ApplyT(func(v *Keypair) pulumi.StringOutput { return v.PairType }).(pulumi.StringOutput)
}

// Passphrase will be used to decrypt the private key. Validated server side.
func (o KeypairOutput) Passphrase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Keypair) pulumi.StringPtrOutput { return v.Passphrase }).(pulumi.StringPtrOutput)
}

// Private key. PEM format will be validated. Must not include extranous spaces or tabs.
func (o KeypairOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Keypair) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

// Public key. PEM format will be validated. Must not include extranous spaces or tabs.
//
// Artifactory REST API call 'Get Key Pair' doesn't return attributes `privateKey` and `passphrase`, but consumes these keys in the POST call.
func (o KeypairOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Keypair) pulumi.StringOutput { return v.PublicKey }).(pulumi.StringOutput)
}

type KeypairArrayOutput struct{ *pulumi.OutputState }

func (KeypairArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Keypair)(nil)).Elem()
}

func (o KeypairArrayOutput) ToKeypairArrayOutput() KeypairArrayOutput {
	return o
}

func (o KeypairArrayOutput) ToKeypairArrayOutputWithContext(ctx context.Context) KeypairArrayOutput {
	return o
}

func (o KeypairArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Keypair] {
	return pulumix.Output[[]*Keypair]{
		OutputState: o.OutputState,
	}
}

func (o KeypairArrayOutput) Index(i pulumi.IntInput) KeypairOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Keypair {
		return vs[0].([]*Keypair)[vs[1].(int)]
	}).(KeypairOutput)
}

type KeypairMapOutput struct{ *pulumi.OutputState }

func (KeypairMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Keypair)(nil)).Elem()
}

func (o KeypairMapOutput) ToKeypairMapOutput() KeypairMapOutput {
	return o
}

func (o KeypairMapOutput) ToKeypairMapOutputWithContext(ctx context.Context) KeypairMapOutput {
	return o
}

func (o KeypairMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Keypair] {
	return pulumix.Output[map[string]*Keypair]{
		OutputState: o.OutputState,
	}
}

func (o KeypairMapOutput) MapIndex(k pulumi.StringInput) KeypairOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Keypair {
		return vs[0].(map[string]*Keypair)[vs[1].(string)]
	}).(KeypairOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeypairInput)(nil)).Elem(), &Keypair{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeypairArrayInput)(nil)).Elem(), KeypairArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeypairMapInput)(nil)).Elem(), KeypairMap{})
	pulumi.RegisterOutputType(KeypairOutput{})
	pulumi.RegisterOutputType(KeypairArrayOutput{})
	pulumi.RegisterOutputType(KeypairMapOutput{})
}
