// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class ArtifactPropertyWebhookHandlerGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("customHttpHeaders")]
        private InputMap<string>? _customHttpHeaders;

        /// <summary>
        /// Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
        /// </summary>
        public InputMap<string> CustomHttpHeaders
        {
            get => _customHttpHeaders ?? (_customHttpHeaders = new InputMap<string>());
            set => _customHttpHeaders = value;
        }

        /// <summary>
        /// Proxy key from Artifactory UI (Administration &gt; Proxies &gt; Configuration).
        /// </summary>
        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        /// <summary>
        /// Secret authentication token that will be sent to the configured URL. The value will be sent as `x-jfrog-event-auth` header.
        /// </summary>
        [Input("secret")]
        public Input<string>? Secret { get; set; }

        /// <summary>
        /// Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public ArtifactPropertyWebhookHandlerGetArgs()
        {
        }
        public static new ArtifactPropertyWebhookHandlerGetArgs Empty => new ArtifactPropertyWebhookHandlerGetArgs();
    }
}
