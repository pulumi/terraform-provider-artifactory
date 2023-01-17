// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.ApiKeyArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.ApiKeyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.ApiKey;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var ci = new ApiKey(&#34;ci&#34;);
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * A user&#39;s API key can be imported using any identifier, e.g.
 * 
 * ```sh
 *  $ pulumi import artifactory:index/apiKey:ApiKey test import
 * ```
 * 
 */
@ResourceType(type="artifactory:index/apiKey:ApiKey")
public class ApiKey extends com.pulumi.resources.CustomResource {
    /**
     * The API key. Deprecated. An upcoming version will support the option to block the usage/creation of API Keys (for admins to set on their platform).
     * In September 2022, the option to block the usage/creation of API Keys will be enabled by default, with the option for admins to change it back to enable API Keys.
     * In January 2023, API Keys will be deprecated all together and the option to use them will no longer be available.
     * It is recommended to use scoped tokens instead - `artifactory.ScopedToken` resource.
     * Please check the [release notes](https://www.jfrog.com/confluence/display/JFROG/Artifactory+Release+Notes#ArtifactoryReleaseNotes-Artifactory7.38.4).
     * 
     */
    @Export(name="apiKey", type=String.class, parameters={})
    private Output<String> apiKey;

    /**
     * @return The API key. Deprecated. An upcoming version will support the option to block the usage/creation of API Keys (for admins to set on their platform).
     * In September 2022, the option to block the usage/creation of API Keys will be enabled by default, with the option for admins to change it back to enable API Keys.
     * In January 2023, API Keys will be deprecated all together and the option to use them will no longer be available.
     * It is recommended to use scoped tokens instead - `artifactory.ScopedToken` resource.
     * Please check the [release notes](https://www.jfrog.com/confluence/display/JFROG/Artifactory+Release+Notes#ArtifactoryReleaseNotes-Artifactory7.38.4).
     * 
     */
    public Output<String> apiKey() {
        return this.apiKey;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ApiKey(String name) {
        this(name, ApiKeyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ApiKey(String name, @Nullable ApiKeyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ApiKey(String name, @Nullable ApiKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/apiKey:ApiKey", name, args == null ? ApiKeyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ApiKey(String name, Output<String> id, @Nullable ApiKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/apiKey:ApiKey", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "apiKey"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static ApiKey get(String name, Output<String> id, @Nullable ApiKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ApiKey(name, id, state, options);
    }
}
