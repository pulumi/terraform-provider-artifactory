// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BuildWebhookHandlerArgs extends com.pulumi.resources.ResourceArgs {

    public static final BuildWebhookHandlerArgs Empty = new BuildWebhookHandlerArgs();

    @Import(name="customHttpHeaders")
    private @Nullable Output<Map<String,String>> customHttpHeaders;

    public Optional<Output<Map<String,String>>> customHttpHeaders() {
        return Optional.ofNullable(this.customHttpHeaders);
    }

    @Import(name="proxy")
    private @Nullable Output<String> proxy;

    public Optional<Output<String>> proxy() {
        return Optional.ofNullable(this.proxy);
    }

    @Import(name="secret")
    private @Nullable Output<String> secret;

    public Optional<Output<String>> secret() {
        return Optional.ofNullable(this.secret);
    }

    @Import(name="url", required=true)
    private Output<String> url;

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

        public Builder customHttpHeaders(@Nullable Output<Map<String,String>> customHttpHeaders) {
            $.customHttpHeaders = customHttpHeaders;
            return this;
        }

        public Builder customHttpHeaders(Map<String,String> customHttpHeaders) {
            return customHttpHeaders(Output.of(customHttpHeaders));
        }

        public Builder proxy(@Nullable Output<String> proxy) {
            $.proxy = proxy;
            return this;
        }

        public Builder proxy(String proxy) {
            return proxy(Output.of(proxy));
        }

        public Builder secret(@Nullable Output<String> secret) {
            $.secret = secret;
            return this;
        }

        public Builder secret(String secret) {
            return secret(Output.of(secret));
        }

        public Builder url(Output<String> url) {
            $.url = url;
            return this;
        }

        public Builder url(String url) {
            return url(Output.of(url));
        }

        public BuildWebhookHandlerArgs build() {
            $.url = Objects.requireNonNull($.url, "expected parameter 'url' to be non-null");
            return $;
        }
    }

}
