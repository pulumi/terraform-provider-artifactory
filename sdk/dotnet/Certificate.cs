// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    [ArtifactoryResourceType("artifactory:index/certificate:Certificate")]
    public partial class Certificate : global::Pulumi.CustomResource
    {
        [Output("alias")]
        public Output<string> Alias { get; private set; } = null!;

        [Output("content")]
        public Output<string?> Content { get; private set; } = null!;

        [Output("file")]
        public Output<string?> File { get; private set; } = null!;

        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        [Output("issuedBy")]
        public Output<string> IssuedBy { get; private set; } = null!;

        [Output("issuedOn")]
        public Output<string> IssuedOn { get; private set; } = null!;

        [Output("issuedTo")]
        public Output<string> IssuedTo { get; private set; } = null!;

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
        [Input("alias", required: true)]
        public Input<string> Alias { get; set; } = null!;

        [Input("content")]
        private Input<string>? _content;
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
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        [Input("content")]
        private Input<string>? _content;
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

        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        [Input("issuedBy")]
        public Input<string>? IssuedBy { get; set; }

        [Input("issuedOn")]
        public Input<string>? IssuedOn { get; set; }

        [Input("issuedTo")]
        public Input<string>? IssuedTo { get; set; }

        [Input("validUntil")]
        public Input<string>? ValidUntil { get; set; }

        public CertificateState()
        {
        }
        public static new CertificateState Empty => new CertificateState();
    }
}
