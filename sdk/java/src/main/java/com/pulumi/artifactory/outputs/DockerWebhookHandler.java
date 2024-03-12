// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DockerWebhookHandler {
    /**
     * @return Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
     * 
     */
    private @Nullable Map<String,String> customHttpHeaders;
    /**
     * @return Proxy key from Artifactory UI (Administration &gt; Proxies &gt; Configuration).
     * 
     */
    private @Nullable String proxy;
    /**
     * @return Secret authentication token that will be sent to the configured URL. The value will be sent as `x-jfrog-event-auth` header.
     * 
     */
    private @Nullable String secret;
    /**
     * @return Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
     * 
     */
    private String url;
    /**
     * @return When set to `true`, the secret will be used to sign the event payload, allowing the target to validate that the payload content has not been changed and will not be passed as part of the event. If left unset or set to `false`, the secret is passed through the `X-JFrog-Event-Auth` HTTP header.
     * 
     */
    private @Nullable Boolean useSecretForSigning;

    private DockerWebhookHandler() {}
    /**
     * @return Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
     * 
     */
    public Map<String,String> customHttpHeaders() {
        return this.customHttpHeaders == null ? Map.of() : this.customHttpHeaders;
    }
    /**
     * @return Proxy key from Artifactory UI (Administration &gt; Proxies &gt; Configuration).
     * 
     */
    public Optional<String> proxy() {
        return Optional.ofNullable(this.proxy);
    }
    /**
     * @return Secret authentication token that will be sent to the configured URL. The value will be sent as `x-jfrog-event-auth` header.
     * 
     */
    public Optional<String> secret() {
        return Optional.ofNullable(this.secret);
    }
    /**
     * @return Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
     * 
     */
    public String url() {
        return this.url;
    }
    /**
     * @return When set to `true`, the secret will be used to sign the event payload, allowing the target to validate that the payload content has not been changed and will not be passed as part of the event. If left unset or set to `false`, the secret is passed through the `X-JFrog-Event-Auth` HTTP header.
     * 
     */
    public Optional<Boolean> useSecretForSigning() {
        return Optional.ofNullable(this.useSecretForSigning);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DockerWebhookHandler defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Map<String,String> customHttpHeaders;
        private @Nullable String proxy;
        private @Nullable String secret;
        private String url;
        private @Nullable Boolean useSecretForSigning;
        public Builder() {}
        public Builder(DockerWebhookHandler defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.customHttpHeaders = defaults.customHttpHeaders;
    	      this.proxy = defaults.proxy;
    	      this.secret = defaults.secret;
    	      this.url = defaults.url;
    	      this.useSecretForSigning = defaults.useSecretForSigning;
        }

        @CustomType.Setter
        public Builder customHttpHeaders(@Nullable Map<String,String> customHttpHeaders) {

            this.customHttpHeaders = customHttpHeaders;
            return this;
        }
        @CustomType.Setter
        public Builder proxy(@Nullable String proxy) {

            this.proxy = proxy;
            return this;
        }
        @CustomType.Setter
        public Builder secret(@Nullable String secret) {

            this.secret = secret;
            return this;
        }
        @CustomType.Setter
        public Builder url(String url) {
            if (url == null) {
              throw new MissingRequiredPropertyException("DockerWebhookHandler", "url");
            }
            this.url = url;
            return this;
        }
        @CustomType.Setter
        public Builder useSecretForSigning(@Nullable Boolean useSecretForSigning) {

            this.useSecretForSigning = useSecretForSigning;
            return this;
        }
        public DockerWebhookHandler build() {
            final var _resultValue = new DockerWebhookHandler();
            _resultValue.customHttpHeaders = customHttpHeaders;
            _resultValue.proxy = proxy;
            _resultValue.secret = secret;
            _resultValue.url = url;
            _resultValue.useSecretForSigning = useSecretForSigning;
            return _resultValue;
        }
    }
}
