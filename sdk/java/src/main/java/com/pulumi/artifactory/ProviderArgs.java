// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProviderArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderArgs Empty = new ProviderArgs();

    /**
     * This is a access token that can be given to you by your admin under `User Management &gt; Access Tokens`. If not set, the
     * &#39;api_key&#39; attribute value will be used.
     * 
     */
    @Import(name="accessToken")
    private @Nullable Output<String> accessToken;

    /**
     * @return This is a access token that can be given to you by your admin under `User Management &gt; Access Tokens`. If not set, the
     * &#39;api_key&#39; attribute value will be used.
     * 
     */
    public Optional<Output<String>> accessToken() {
        return Optional.ofNullable(this.accessToken);
    }

    /**
     * API key. If `access_token` attribute, `JFROG_ACCESS_TOKEN` or `ARTIFACTORY_ACCESS_TOKEN` environment variable is set,
     * the provider will ignore this attribute.
     * 
     * @deprecated
     * An upcoming version will support the option to block the usage/creation of API Keys (for admins to set on their platform).
     * In a future version (scheduled for end of Q3, 2023), the option to disable the usage/creation of API Keys will be available and set to disabled by default. Admins will be able to enable the usage/creation of API Keys.
     * By end of Q4 2024, API Keys will be deprecated all together and the option to use them will no longer be available. See [JFrog API deprecation process](https://jfrog.com/help/r/jfrog-platform-administration-documentation/jfrog-api-key-deprecation-process) for more details.
     * 
     */
    @Deprecated /* An upcoming version will support the option to block the usage/creation of API Keys (for admins to set on their platform).
In a future version (scheduled for end of Q3, 2023), the option to disable the usage/creation of API Keys will be available and set to disabled by default. Admins will be able to enable the usage/creation of API Keys.
By end of Q4 2024, API Keys will be deprecated all together and the option to use them will no longer be available. See [JFrog API deprecation process](https://jfrog.com/help/r/jfrog-platform-administration-documentation/jfrog-api-key-deprecation-process) for more details. */
    @Import(name="apiKey")
    private @Nullable Output<String> apiKey;

    /**
     * @return API key. If `access_token` attribute, `JFROG_ACCESS_TOKEN` or `ARTIFACTORY_ACCESS_TOKEN` environment variable is set,
     * the provider will ignore this attribute.
     * 
     * @deprecated
     * An upcoming version will support the option to block the usage/creation of API Keys (for admins to set on their platform).
     * In a future version (scheduled for end of Q3, 2023), the option to disable the usage/creation of API Keys will be available and set to disabled by default. Admins will be able to enable the usage/creation of API Keys.
     * By end of Q4 2024, API Keys will be deprecated all together and the option to use them will no longer be available. See [JFrog API deprecation process](https://jfrog.com/help/r/jfrog-platform-administration-documentation/jfrog-api-key-deprecation-process) for more details.
     * 
     */
    @Deprecated /* An upcoming version will support the option to block the usage/creation of API Keys (for admins to set on their platform).
In a future version (scheduled for end of Q3, 2023), the option to disable the usage/creation of API Keys will be available and set to disabled by default. Admins will be able to enable the usage/creation of API Keys.
By end of Q4 2024, API Keys will be deprecated all together and the option to use them will no longer be available. See [JFrog API deprecation process](https://jfrog.com/help/r/jfrog-platform-administration-documentation/jfrog-api-key-deprecation-process) for more details. */
    public Optional<Output<String>> apiKey() {
        return Optional.ofNullable(this.apiKey);
    }

    /**
     * OIDC provider name. See [Configure an OIDC
     * Integration](https://jfrog.com/help/r/jfrog-platform-administration-documentation/configure-an-oidc-integration) for
     * more details.
     * 
     */
    @Import(name="oidcProviderName")
    private @Nullable Output<String> oidcProviderName;

    /**
     * @return OIDC provider name. See [Configure an OIDC
     * Integration](https://jfrog.com/help/r/jfrog-platform-administration-documentation/configure-an-oidc-integration) for
     * more details.
     * 
     */
    public Optional<Output<String>> oidcProviderName() {
        return Optional.ofNullable(this.oidcProviderName);
    }

    /**
     * Artifactory URL.
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return Artifactory URL.
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    private ProviderArgs() {}

    private ProviderArgs(ProviderArgs $) {
        this.accessToken = $.accessToken;
        this.apiKey = $.apiKey;
        this.oidcProviderName = $.oidcProviderName;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderArgs $;

        public Builder() {
            $ = new ProviderArgs();
        }

        public Builder(ProviderArgs defaults) {
            $ = new ProviderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessToken This is a access token that can be given to you by your admin under `User Management &gt; Access Tokens`. If not set, the
         * &#39;api_key&#39; attribute value will be used.
         * 
         * @return builder
         * 
         */
        public Builder accessToken(@Nullable Output<String> accessToken) {
            $.accessToken = accessToken;
            return this;
        }

        /**
         * @param accessToken This is a access token that can be given to you by your admin under `User Management &gt; Access Tokens`. If not set, the
         * &#39;api_key&#39; attribute value will be used.
         * 
         * @return builder
         * 
         */
        public Builder accessToken(String accessToken) {
            return accessToken(Output.of(accessToken));
        }

        /**
         * @param apiKey API key. If `access_token` attribute, `JFROG_ACCESS_TOKEN` or `ARTIFACTORY_ACCESS_TOKEN` environment variable is set,
         * the provider will ignore this attribute.
         * 
         * @return builder
         * 
         * @deprecated
         * An upcoming version will support the option to block the usage/creation of API Keys (for admins to set on their platform).
         * In a future version (scheduled for end of Q3, 2023), the option to disable the usage/creation of API Keys will be available and set to disabled by default. Admins will be able to enable the usage/creation of API Keys.
         * By end of Q4 2024, API Keys will be deprecated all together and the option to use them will no longer be available. See [JFrog API deprecation process](https://jfrog.com/help/r/jfrog-platform-administration-documentation/jfrog-api-key-deprecation-process) for more details.
         * 
         */
        @Deprecated /* An upcoming version will support the option to block the usage/creation of API Keys (for admins to set on their platform).
In a future version (scheduled for end of Q3, 2023), the option to disable the usage/creation of API Keys will be available and set to disabled by default. Admins will be able to enable the usage/creation of API Keys.
By end of Q4 2024, API Keys will be deprecated all together and the option to use them will no longer be available. See [JFrog API deprecation process](https://jfrog.com/help/r/jfrog-platform-administration-documentation/jfrog-api-key-deprecation-process) for more details. */
        public Builder apiKey(@Nullable Output<String> apiKey) {
            $.apiKey = apiKey;
            return this;
        }

        /**
         * @param apiKey API key. If `access_token` attribute, `JFROG_ACCESS_TOKEN` or `ARTIFACTORY_ACCESS_TOKEN` environment variable is set,
         * the provider will ignore this attribute.
         * 
         * @return builder
         * 
         * @deprecated
         * An upcoming version will support the option to block the usage/creation of API Keys (for admins to set on their platform).
         * In a future version (scheduled for end of Q3, 2023), the option to disable the usage/creation of API Keys will be available and set to disabled by default. Admins will be able to enable the usage/creation of API Keys.
         * By end of Q4 2024, API Keys will be deprecated all together and the option to use them will no longer be available. See [JFrog API deprecation process](https://jfrog.com/help/r/jfrog-platform-administration-documentation/jfrog-api-key-deprecation-process) for more details.
         * 
         */
        @Deprecated /* An upcoming version will support the option to block the usage/creation of API Keys (for admins to set on their platform).
In a future version (scheduled for end of Q3, 2023), the option to disable the usage/creation of API Keys will be available and set to disabled by default. Admins will be able to enable the usage/creation of API Keys.
By end of Q4 2024, API Keys will be deprecated all together and the option to use them will no longer be available. See [JFrog API deprecation process](https://jfrog.com/help/r/jfrog-platform-administration-documentation/jfrog-api-key-deprecation-process) for more details. */
        public Builder apiKey(String apiKey) {
            return apiKey(Output.of(apiKey));
        }

        /**
         * @param oidcProviderName OIDC provider name. See [Configure an OIDC
         * Integration](https://jfrog.com/help/r/jfrog-platform-administration-documentation/configure-an-oidc-integration) for
         * more details.
         * 
         * @return builder
         * 
         */
        public Builder oidcProviderName(@Nullable Output<String> oidcProviderName) {
            $.oidcProviderName = oidcProviderName;
            return this;
        }

        /**
         * @param oidcProviderName OIDC provider name. See [Configure an OIDC
         * Integration](https://jfrog.com/help/r/jfrog-platform-administration-documentation/configure-an-oidc-integration) for
         * more details.
         * 
         * @return builder
         * 
         */
        public Builder oidcProviderName(String oidcProviderName) {
            return oidcProviderName(Output.of(oidcProviderName));
        }

        /**
         * @param url Artifactory URL.
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url Artifactory URL.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        public ProviderArgs build() {
            return $;
        }
    }

}
