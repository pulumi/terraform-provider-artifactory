// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class DestinationCustomWebhookHandlerGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("httpHeaders")]
        private InputMap<string>? _httpHeaders;

        /// <summary>
        /// HTTP headers you wish to use to invoke the Webhook, comprise key/value pair.
        /// </summary>
        public InputMap<string> HttpHeaders
        {
            get => _httpHeaders ?? (_httpHeaders = new InputMap<string>());
            set => _httpHeaders = value;
        }

        /// <summary>
        /// Specifies the HTTP method for the URL that the Webhook invokes. Allowed values are: `GET`, `POST`, `PUT`, `PATCH`, `DELETE`.
        /// </summary>
        [Input("method")]
        public Input<string>? Method { get; set; }

        /// <summary>
        /// This attribute is used to build the request body. Used in custom webhooks
        /// </summary>
        [Input("payload")]
        public Input<string>? Payload { get; set; }

        /// <summary>
        /// Proxy key from Artifactory UI (Administration &gt; Proxies &gt; Configuration).
        /// </summary>
        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        [Input("secrets")]
        private InputMap<string>? _secrets;

        /// <summary>
        /// Defines a set of sensitive values (such as, tokens and passwords) that can be injected in the headers and/or payload.Secrets’ values are encrypted. In the header/payload, the value can be invoked using the `{{.secrets.token}}` format, where token is the name provided for the secret value. Comprise key/value pair. **Note:** if multiple handlers are used, same secret name and different secret value for the same url won't work. Example:
        /// </summary>
        public InputMap<string> Secrets
        {
            get => _secrets ?? (_secrets = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _secrets = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public DestinationCustomWebhookHandlerGetArgs()
        {
        }
        public static new DestinationCustomWebhookHandlerGetArgs Empty => new DestinationCustomWebhookHandlerGetArgs();
    }
}
