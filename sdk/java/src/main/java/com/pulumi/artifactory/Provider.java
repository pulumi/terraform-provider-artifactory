// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.ProviderArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The provider type for the artifactory package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 * 
 */
@ResourceType(type="pulumi:providers:artifactory")
public class Provider extends com.pulumi.resources.ProviderResource {
    /**
     * This is a access token that can be given to you by your admin under `User Management &gt; Access Tokens`. If not set, the
     * &#39;api_key&#39; attribute value will be used.
     * 
     */
    @Export(name="accessToken", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> accessToken;

    /**
     * @return This is a access token that can be given to you by your admin under `User Management &gt; Access Tokens`. If not set, the
     * &#39;api_key&#39; attribute value will be used.
     * 
     */
    public Output<Optional<String>> accessToken() {
        return Codegen.optional(this.accessToken);
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
    @Export(name="apiKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> apiKey;

    /**
     * @return API key. If `access_token` attribute, `JFROG_ACCESS_TOKEN` or `ARTIFACTORY_ACCESS_TOKEN` environment variable is set,
     * the provider will ignore this attribute.
     * 
     */
    public Output<Optional<String>> apiKey() {
        return Codegen.optional(this.apiKey);
    }
    /**
     * OIDC provider name. See [Configure an OIDC
     * Integration](https://jfrog.com/help/r/jfrog-platform-administration-documentation/configure-an-oidc-integration) for
     * more details.
     * 
     */
    @Export(name="oidcProviderName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> oidcProviderName;

    /**
     * @return OIDC provider name. See [Configure an OIDC
     * Integration](https://jfrog.com/help/r/jfrog-platform-administration-documentation/configure-an-oidc-integration) for
     * more details.
     * 
     */
    public Output<Optional<String>> oidcProviderName() {
        return Codegen.optional(this.oidcProviderName);
    }
    @Export(name="tfcCredentialTagName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> tfcCredentialTagName;

    public Output<Optional<String>> tfcCredentialTagName() {
        return Codegen.optional(this.tfcCredentialTagName);
    }
    /**
     * Artifactory URL.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> url;

    /**
     * @return Artifactory URL.
     * 
     */
    public Output<Optional<String>> url() {
        return Codegen.optional(this.url);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Provider(java.lang.String name) {
        this(name, ProviderArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Provider(java.lang.String name, @Nullable ProviderArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Provider(java.lang.String name, @Nullable ProviderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private static ProviderArgs makeArgs(@Nullable ProviderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ProviderArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "accessToken",
                "apiKey"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

}
