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
    public sealed class UserWebhookHandler
    {
        /// <summary>
        /// Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? CustomHttpHeaders;
        /// <summary>
        /// Proxy key from Artifactory UI (Administration &gt; Proxies &gt; Configuration).
        /// </summary>
        public readonly string? Proxy;
        /// <summary>
        /// Secret authentication token that will be sent to the configured URL. The value will be sent as `x-jfrog-event-auth` header.
        /// </summary>
        public readonly string? Secret;
        /// <summary>
        /// Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
        /// </summary>
        public readonly string Url;
        /// <summary>
        /// When set to `true`, the secret will be used to sign the event payload, allowing the target to validate that the payload content has not been changed and will not be passed as part of the event. If left unset or set to `false`, the secret is passed through the `X-JFrog-Event-Auth` HTTP header.
        /// </summary>
        public readonly bool? UseSecretForSigning;

        [OutputConstructor]
        private UserWebhookHandler(
            ImmutableDictionary<string, string>? customHttpHeaders,

            string? proxy,

            string? secret,

            string url,

            bool? useSecretForSigning)
        {
            CustomHttpHeaders = customHttpHeaders;
            Proxy = proxy;
            Secret = secret;
            Url = url;
            UseSecretForSigning = useSecretForSigning;
        }
    }
}
