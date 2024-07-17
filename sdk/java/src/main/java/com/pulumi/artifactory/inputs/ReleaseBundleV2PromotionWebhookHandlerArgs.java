// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ReleaseBundleV2PromotionWebhookHandlerArgs extends com.pulumi.resources.ResourceArgs {

    public static final ReleaseBundleV2PromotionWebhookHandlerArgs Empty = new ReleaseBundleV2PromotionWebhookHandlerArgs();

    /**
     * Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
     * 
     */
    @Import(name="customHttpHeaders")
    private @Nullable Output<Map<String,String>> customHttpHeaders;

    /**
     * @return Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
     * 
     */
    public Optional<Output<Map<String,String>>> customHttpHeaders() {
        return Optional.ofNullable(this.customHttpHeaders);
    }

    /**
     * Proxy key from Artifactory UI (Administration &gt; Proxies &gt; Configuration).
     * 
     */
    @Import(name="proxy")
    private @Nullable Output<String> proxy;

    /**
     * @return Proxy key from Artifactory UI (Administration &gt; Proxies &gt; Configuration).
     * 
     */
    public Optional<Output<String>> proxy() {
        return Optional.ofNullable(this.proxy);
    }

    /**
     * Secret authentication token that will be sent to the configured URL. The value will be sent as `x-jfrog-event-auth` header.
     * 
     */
    @Import(name="secret")
    private @Nullable Output<String> secret;

    /**
     * @return Secret authentication token that will be sent to the configured URL. The value will be sent as `x-jfrog-event-auth` header.
     * 
     */
    public Optional<Output<String>> secret() {
        return Optional.ofNullable(this.secret);
    }

    /**
     * Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
     * 
     */
    @Import(name="url", required=true)
    private Output<String> url;

    /**
     * @return Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     * When set to `true`, the secret will be used to sign the event payload, allowing the target to validate that the payload content has not been changed and will not be passed as part of the event. If left unset or set to `false`, the secret is passed through the `X-JFrog-Event-Auth` HTTP header.
     * 
     */
    @Import(name="useSecretForSigning")
    private @Nullable Output<Boolean> useSecretForSigning;

    /**
     * @return When set to `true`, the secret will be used to sign the event payload, allowing the target to validate that the payload content has not been changed and will not be passed as part of the event. If left unset or set to `false`, the secret is passed through the `X-JFrog-Event-Auth` HTTP header.
     * 
     */
    public Optional<Output<Boolean>> useSecretForSigning() {
        return Optional.ofNullable(this.useSecretForSigning);
    }

    private ReleaseBundleV2PromotionWebhookHandlerArgs() {}

    private ReleaseBundleV2PromotionWebhookHandlerArgs(ReleaseBundleV2PromotionWebhookHandlerArgs $) {
        this.customHttpHeaders = $.customHttpHeaders;
        this.proxy = $.proxy;
        this.secret = $.secret;
        this.url = $.url;
        this.useSecretForSigning = $.useSecretForSigning;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ReleaseBundleV2PromotionWebhookHandlerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ReleaseBundleV2PromotionWebhookHandlerArgs $;

        public Builder() {
            $ = new ReleaseBundleV2PromotionWebhookHandlerArgs();
        }

        public Builder(ReleaseBundleV2PromotionWebhookHandlerArgs defaults) {
            $ = new ReleaseBundleV2PromotionWebhookHandlerArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param customHttpHeaders Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
         * 
         * @return builder
         * 
         */
        public Builder customHttpHeaders(@Nullable Output<Map<String,String>> customHttpHeaders) {
            $.customHttpHeaders = customHttpHeaders;
            return this;
        }

        /**
         * @param customHttpHeaders Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
         * 
         * @return builder
         * 
         */
        public Builder customHttpHeaders(Map<String,String> customHttpHeaders) {
            return customHttpHeaders(Output.of(customHttpHeaders));
        }

        /**
         * @param proxy Proxy key from Artifactory UI (Administration &gt; Proxies &gt; Configuration).
         * 
         * @return builder
         * 
         */
        public Builder proxy(@Nullable Output<String> proxy) {
            $.proxy = proxy;
            return this;
        }

        /**
         * @param proxy Proxy key from Artifactory UI (Administration &gt; Proxies &gt; Configuration).
         * 
         * @return builder
         * 
         */
        public Builder proxy(String proxy) {
            return proxy(Output.of(proxy));
        }

        /**
         * @param secret Secret authentication token that will be sent to the configured URL. The value will be sent as `x-jfrog-event-auth` header.
         * 
         * @return builder
         * 
         */
        public Builder secret(@Nullable Output<String> secret) {
            $.secret = secret;
            return this;
        }

        /**
         * @param secret Secret authentication token that will be sent to the configured URL. The value will be sent as `x-jfrog-event-auth` header.
         * 
         * @return builder
         * 
         */
        public Builder secret(String secret) {
            return secret(Output.of(secret));
        }

        /**
         * @param url Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
         * 
         * @return builder
         * 
         */
        public Builder url(Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        /**
         * @param useSecretForSigning When set to `true`, the secret will be used to sign the event payload, allowing the target to validate that the payload content has not been changed and will not be passed as part of the event. If left unset or set to `false`, the secret is passed through the `X-JFrog-Event-Auth` HTTP header.
         * 
         * @return builder
         * 
         */
        public Builder useSecretForSigning(@Nullable Output<Boolean> useSecretForSigning) {
            $.useSecretForSigning = useSecretForSigning;
            return this;
        }

        /**
         * @param useSecretForSigning When set to `true`, the secret will be used to sign the event payload, allowing the target to validate that the payload content has not been changed and will not be passed as part of the event. If left unset or set to `false`, the secret is passed through the `X-JFrog-Event-Auth` HTTP header.
         * 
         * @return builder
         * 
         */
        public Builder useSecretForSigning(Boolean useSecretForSigning) {
            return useSecretForSigning(Output.of(useSecretForSigning));
        }

        public ReleaseBundleV2PromotionWebhookHandlerArgs build() {
            if ($.url == null) {
                throw new MissingRequiredPropertyException("ReleaseBundleV2PromotionWebhookHandlerArgs", "url");
            }
            return $;
        }
    }

}
