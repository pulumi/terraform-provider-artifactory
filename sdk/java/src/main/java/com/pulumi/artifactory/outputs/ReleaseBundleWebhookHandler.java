// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ReleaseBundleWebhookHandler {
    private @Nullable Map<String,String> customHttpHeaders;
    private @Nullable String proxy;
    private @Nullable String secret;
    private String url;

    private ReleaseBundleWebhookHandler() {}
    public Map<String,String> customHttpHeaders() {
        return this.customHttpHeaders == null ? Map.of() : this.customHttpHeaders;
    }
    public Optional<String> proxy() {
        return Optional.ofNullable(this.proxy);
    }
    public Optional<String> secret() {
        return Optional.ofNullable(this.secret);
    }
    public String url() {
        return this.url;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ReleaseBundleWebhookHandler defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Map<String,String> customHttpHeaders;
        private @Nullable String proxy;
        private @Nullable String secret;
        private String url;
        public Builder() {}
        public Builder(ReleaseBundleWebhookHandler defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.customHttpHeaders = defaults.customHttpHeaders;
    	      this.proxy = defaults.proxy;
    	      this.secret = defaults.secret;
    	      this.url = defaults.url;
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
            this.url = Objects.requireNonNull(url);
            return this;
        }
        public ReleaseBundleWebhookHandler build() {
            final var o = new ReleaseBundleWebhookHandler();
            o.customHttpHeaders = customHttpHeaders;
            o.proxy = proxy;
            o.secret = secret;
            o.url = url;
            return o;
        }
    }
}
