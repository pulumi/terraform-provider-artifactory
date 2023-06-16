// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    /// <summary>
    /// RSA key pairs are used to sign and verify the Alpine Linux index files in JFrog Artifactory, while GPG key pairs are
    /// used to sign and validate packages integrity in JFrog Distribution. The JFrog Platform enables you to manage multiple
    /// RSA and GPG signing keys through the Keys Management UI and REST API. The JFrog Platform supports managing multiple
    /// pairs of GPG signing keys to sign packages for authentication of several package types such as Debian, Opkg, and RPM
    /// through the Keys Management UI and REST API.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.IO;
    /// using System.Linq;
    /// using Pulumi;
    /// using Artifactory = Pulumi.Artifactory;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var some_keypair6543461672124900137 = new Artifactory.Index.Keypair.Keypair("some-keypair6543461672124900137", new()
    ///     {
    ///         PairName = "some-keypair6543461672124900137",
    ///         PairType = "RSA",
    ///         Alias = "foo-alias6543461672124900137",
    ///         PrivateKey = File.ReadAllText("samples/rsa.priv"),
    ///         PublicKey = File.ReadAllText("samples/rsa.pub"),
    ///         Passphrase = "PASSPHRASE",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Keypair can be imported using their name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import artifactory:index/keypair:Keypair my-keypair my-keypair
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/keypair:Keypair")]
    public partial class Keypair : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Will be used as a filename when retrieving the public key via REST API.
        /// </summary>
        [Output("alias")]
        public Output<string> Alias { get; private set; } = null!;

        /// <summary>
        /// A unique identifier for the Key Pair record.
        /// </summary>
        [Output("pairName")]
        public Output<string> PairName { get; private set; } = null!;

        /// <summary>
        /// Key Pair type. Supported types - GPG and RSA.
        /// </summary>
        [Output("pairType")]
        public Output<string> PairType { get; private set; } = null!;

        /// <summary>
        /// Passphrase will be used to decrypt the private key. Validated server side.
        /// </summary>
        [Output("passphrase")]
        public Output<string?> Passphrase { get; private set; } = null!;

        /// <summary>
        /// Private key. PEM format will be validated.
        /// </summary>
        [Output("privateKey")]
        public Output<string> PrivateKey { get; private set; } = null!;

        /// <summary>
        /// Public key. PEM format will be validated.
        /// </summary>
        [Output("publicKey")]
        public Output<string> PublicKey { get; private set; } = null!;

        /// <summary>
        /// Unknown usage. Returned in the json payload and cannot be set.
        /// 
        /// Artifactory REST API call Get Key Pair doesn't return keys `private_key` and `passphrase`, but consumes these keys in the POST call.
        /// </summary>
        [Output("unavailable")]
        public Output<bool> Unavailable { get; private set; } = null!;


        /// <summary>
        /// Create a Keypair resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Keypair(string name, KeypairArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/keypair:Keypair", name, args ?? new KeypairArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Keypair(string name, Input<string> id, KeypairState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/keypair:Keypair", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "passphrase",
                    "privateKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Keypair resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Keypair Get(string name, Input<string> id, KeypairState? state = null, CustomResourceOptions? options = null)
        {
            return new Keypair(name, id, state, options);
        }
    }

    public sealed class KeypairArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Will be used as a filename when retrieving the public key via REST API.
        /// </summary>
        [Input("alias", required: true)]
        public Input<string> Alias { get; set; } = null!;

        /// <summary>
        /// A unique identifier for the Key Pair record.
        /// </summary>
        [Input("pairName", required: true)]
        public Input<string> PairName { get; set; } = null!;

        /// <summary>
        /// Key Pair type. Supported types - GPG and RSA.
        /// </summary>
        [Input("pairType", required: true)]
        public Input<string> PairType { get; set; } = null!;

        [Input("passphrase")]
        private Input<string>? _passphrase;

        /// <summary>
        /// Passphrase will be used to decrypt the private key. Validated server side.
        /// </summary>
        public Input<string>? Passphrase
        {
            get => _passphrase;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _passphrase = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("privateKey", required: true)]
        private Input<string>? _privateKey;

        /// <summary>
        /// Private key. PEM format will be validated.
        /// </summary>
        public Input<string>? PrivateKey
        {
            get => _privateKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privateKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Public key. PEM format will be validated.
        /// </summary>
        [Input("publicKey", required: true)]
        public Input<string> PublicKey { get; set; } = null!;

        public KeypairArgs()
        {
        }
        public static new KeypairArgs Empty => new KeypairArgs();
    }

    public sealed class KeypairState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Will be used as a filename when retrieving the public key via REST API.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// A unique identifier for the Key Pair record.
        /// </summary>
        [Input("pairName")]
        public Input<string>? PairName { get; set; }

        /// <summary>
        /// Key Pair type. Supported types - GPG and RSA.
        /// </summary>
        [Input("pairType")]
        public Input<string>? PairType { get; set; }

        [Input("passphrase")]
        private Input<string>? _passphrase;

        /// <summary>
        /// Passphrase will be used to decrypt the private key. Validated server side.
        /// </summary>
        public Input<string>? Passphrase
        {
            get => _passphrase;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _passphrase = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("privateKey")]
        private Input<string>? _privateKey;

        /// <summary>
        /// Private key. PEM format will be validated.
        /// </summary>
        public Input<string>? PrivateKey
        {
            get => _privateKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privateKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Public key. PEM format will be validated.
        /// </summary>
        [Input("publicKey")]
        public Input<string>? PublicKey { get; set; }

        /// <summary>
        /// Unknown usage. Returned in the json payload and cannot be set.
        /// 
        /// Artifactory REST API call Get Key Pair doesn't return keys `private_key` and `passphrase`, but consumes these keys in the POST call.
        /// </summary>
        [Input("unavailable")]
        public Input<bool>? Unavailable { get; set; }

        public KeypairState()
        {
        }
        public static new KeypairState Empty => new KeypairState();
    }
}
