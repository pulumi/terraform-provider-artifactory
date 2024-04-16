// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Outputs
{

    [OutputType]
    public sealed class ArtifactCustomWebhookHandler
    {
        /// <summary>
        /// HTTP headers you wish to use to invoke the Webhook, comprise key/value pair.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? HttpHeaders;
        /// <summary>
        /// This attribute is used to build the request body. Used in custom webhooks
        /// </summary>
        public readonly string? Payload;
        /// <summary>
        /// Proxy key from Artifactory UI (Administration &gt; Proxies &gt; Configuration).
        /// </summary>
        public readonly string? Proxy;
        /// <summary>
        /// Defines a set of sensitive values (such as, tokens and passwords) that can be injected in the headers and/or payload.Secrets’ values are encrypted. In the header/payload, the value can be invoked using the `{{.secrets.token}}` format, where token is the name provided for the secret value. Comprise key/value pair. **Note:** if multiple handlers are used, same secret name and different secret value for the same url won't work. Example: 
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Secrets;
        /// <summary>
        /// Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private ArtifactCustomWebhookHandler(
            ImmutableDictionary<string, string>? httpHeaders,

            string? payload,

            string? proxy,

            ImmutableDictionary<string, string>? secrets,

            string url)
        {
            HttpHeaders = httpHeaders;
            Payload = payload;
            Proxy = proxy;
            Secrets = secrets;
            Url = url;
        }
    }
}
