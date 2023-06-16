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
    /// Provides an Artifactory certificate resource. This can be used to create and manage Artifactory certificates which can be used as client authentication against remote repositories.
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
    ///     // Create a new Artifactory certificate called my-cert
    ///     var my_cert = new Artifactory.Index.Certificate.Certificate("my-cert", new()
    ///     {
    ///         Alias = "my-cert",
    ///         Content = File.ReadAllText("/path/to/bundle.pem"),
    ///     });
    /// 
    ///     // This can then be used by a remote repository
    ///     var my_remote = new Artifactory.Index.RemoteMavenRepository.RemoteMavenRepository("my-remote", new()
    ///     {
    ///         ClientTlsCertificate = my_cert.Alias,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Certificates can be imported using their alias, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import artifactory:index/certificate:Certificate my-cert my-cert
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/certificate:Certificate")]
    public partial class Certificate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of certificate.
        /// </summary>
        [Output("alias")]
        public Output<string> Alias { get; private set; } = null!;

        /// <summary>
        /// PEM-encoded client certificate and private key.
        /// </summary>
        [Output("content")]
        public Output<string?> Content { get; private set; } = null!;

        [Output("file")]
        public Output<string?> File { get; private set; } = null!;

        /// <summary>
        /// SHA256 fingerprint of the certificate.
        /// </summary>
        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// Name of the certificate authority that issued the certificate.
        /// </summary>
        [Output("issuedBy")]
        public Output<string> IssuedBy { get; private set; } = null!;

        /// <summary>
        /// The time &amp; date when the certificate is valid from.
        /// </summary>
        [Output("issuedOn")]
        public Output<string> IssuedOn { get; private set; } = null!;

        /// <summary>
        /// Name of whom the certificate has been issued to.
        /// </summary>
        [Output("issuedTo")]
        public Output<string> IssuedTo { get; private set; } = null!;

        /// <summary>
        /// The time &amp; date when the certificate expires.
        /// </summary>
        [Output("validUntil")]
        public Output<string> ValidUntil { get; private set; } = null!;


        /// <summary>
        /// Create a Certificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Certificate(string name, CertificateArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/certificate:Certificate", name, args ?? new CertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Certificate(string name, Input<string> id, CertificateState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/certificate:Certificate", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "content",
                    "file",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Certificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Certificate Get(string name, Input<string> id, CertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new Certificate(name, id, state, options);
        }
    }

    public sealed class CertificateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of certificate.
        /// </summary>
        [Input("alias", required: true)]
        public Input<string> Alias { get; set; } = null!;

        [Input("content")]
        private Input<string>? _content;

        /// <summary>
        /// PEM-encoded client certificate and private key.
        /// </summary>
        public Input<string>? Content
        {
            get => _content;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _content = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("file")]
        private Input<string>? _file;
        public Input<string>? File
        {
            get => _file;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _file = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public CertificateArgs()
        {
        }
        public static new CertificateArgs Empty => new CertificateArgs();
    }

    public sealed class CertificateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of certificate.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        [Input("content")]
        private Input<string>? _content;

        /// <summary>
        /// PEM-encoded client certificate and private key.
        /// </summary>
        public Input<string>? Content
        {
            get => _content;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _content = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("file")]
        private Input<string>? _file;
        public Input<string>? File
        {
            get => _file;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _file = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// SHA256 fingerprint of the certificate.
        /// </summary>
        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        /// <summary>
        /// Name of the certificate authority that issued the certificate.
        /// </summary>
        [Input("issuedBy")]
        public Input<string>? IssuedBy { get; set; }

        /// <summary>
        /// The time &amp; date when the certificate is valid from.
        /// </summary>
        [Input("issuedOn")]
        public Input<string>? IssuedOn { get; set; }

        /// <summary>
        /// Name of whom the certificate has been issued to.
        /// </summary>
        [Input("issuedTo")]
        public Input<string>? IssuedTo { get; set; }

        /// <summary>
        /// The time &amp; date when the certificate expires.
        /// </summary>
        [Input("validUntil")]
        public Input<string>? ValidUntil { get; set; }

        public CertificateState()
        {
        }
        public static new CertificateState Empty => new CertificateState();
    }
}
