// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BuildWebhookHandlerArgs extends com.pulumi.resources.ResourceArgs {

    public static final BuildWebhookHandlerArgs Empty = new BuildWebhookHandlerArgs();

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

    private BuildWebhookHandlerArgs() {}

    private BuildWebhookHandlerArgs(BuildWebhookHandlerArgs $) {
        this.customHttpHeaders = $.customHttpHeaders;
        this.proxy = $.proxy;
        this.secret = $.secret;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BuildWebhookHandlerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BuildWebhookHandlerArgs $;

        public Builder() {
            $ = new BuildWebhookHandlerArgs();
        }

        public Builder(BuildWebhookHandlerArgs defaults) {
            $ = new BuildWebhookHandlerArgs(Objects.requireNonNull(defaults));
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

        public BuildWebhookHandlerArgs build() {
            if ($.url == null) {
                throw new MissingRequiredPropertyException("BuildWebhookHandlerArgs", "url");
            }
            return $;
        }
    }

}
