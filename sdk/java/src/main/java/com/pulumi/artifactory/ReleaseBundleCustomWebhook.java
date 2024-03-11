// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.ReleaseBundleCustomWebhookArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.ReleaseBundleCustomWebhookState;
import com.pulumi.artifactory.outputs.ReleaseBundleCustomWebhookCriteria;
import com.pulumi.artifactory.outputs.ReleaseBundleCustomWebhookHandler;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an Artifactory custom webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.ReleaseBundleCustomWebhook;
 * import com.pulumi.artifactory.ReleaseBundleCustomWebhookArgs;
 * import com.pulumi.artifactory.inputs.ReleaseBundleCustomWebhookCriteriaArgs;
 * import com.pulumi.artifactory.inputs.ReleaseBundleCustomWebhookHandlerArgs;
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
 *         var release_bundle_custom_webhook = new ReleaseBundleCustomWebhook(&#34;release-bundle-custom-webhook&#34;, ReleaseBundleCustomWebhookArgs.builder()        
 *             .criteria(ReleaseBundleCustomWebhookCriteriaArgs.builder()
 *                 .anyReleaseBundle(false)
 *                 .excludePatterns(&#34;bar/**&#34;)
 *                 .includePatterns(&#34;foo/**&#34;)
 *                 .registeredReleaseBundleNames(&#34;bundle-name&#34;)
 *                 .build())
 *             .eventTypes(            
 *                 &#34;created&#34;,
 *                 &#34;signed&#34;,
 *                 &#34;deleted&#34;)
 *             .handlers(ReleaseBundleCustomWebhookHandlerArgs.builder()
 *                 .httpHeaders(Map.ofEntries(
 *                     Map.entry(&#34;headerName1&#34;, &#34;value1&#34;),
 *                     Map.entry(&#34;headerName2&#34;, &#34;value2&#34;)
 *                 ))
 *                 .payload(&#34;{ \&#34;ref\&#34;: \&#34;main\&#34; , \&#34;inputs\&#34;: { \&#34;artifact_path\&#34;: \&#34;test-repo/repo-path\&#34; } }&#34;)
 *                 .secrets(Map.ofEntries(
 *                     Map.entry(&#34;secretName1&#34;, &#34;value1&#34;),
 *                     Map.entry(&#34;secretName2&#34;, &#34;value2&#34;)
 *                 ))
 *                 .url(&#34;https://tempurl.org&#34;)
 *                 .build())
 *             .key(&#34;release-bundle-custom-webhook&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="artifactory:index/releaseBundleCustomWebhook:ReleaseBundleCustomWebhook")
public class ReleaseBundleCustomWebhook extends com.pulumi.resources.CustomResource {
    /**
     * Specifies where the webhook will be applied on which repositories.
     * 
     */
    @Export(name="criteria", refs={ReleaseBundleCustomWebhookCriteria.class}, tree="[0]")
    private Output<ReleaseBundleCustomWebhookCriteria> criteria;

    /**
     * @return Specifies where the webhook will be applied on which repositories.
     * 
     */
    public Output<ReleaseBundleCustomWebhookCriteria> criteria() {
        return this.criteria;
    }
    /**
     * Webhook description. Max length 1000 characters.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Webhook description. Max length 1000 characters.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Status of webhook. Default to `true`.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return Status of webhook. Default to `true`.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `created`, `signed`, `deleted`.
     * 
     */
    @Export(name="eventTypes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> eventTypes;

    /**
     * @return List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `created`, `signed`, `deleted`.
     * 
     */
    public Output<List<String>> eventTypes() {
        return this.eventTypes;
    }
    /**
     * At least one is required.
     * 
     */
    @Export(name="handlers", refs={List.class,ReleaseBundleCustomWebhookHandler.class}, tree="[0,1]")
    private Output<List<ReleaseBundleCustomWebhookHandler>> handlers;

    /**
     * @return At least one is required.
     * 
     */
    public Output<List<ReleaseBundleCustomWebhookHandler>> handlers() {
        return this.handlers;
    }
    /**
     * The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output<String> key;

    /**
     * @return The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     * 
     */
    public Output<String> key() {
        return this.key;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ReleaseBundleCustomWebhook(String name) {
        this(name, ReleaseBundleCustomWebhookArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ReleaseBundleCustomWebhook(String name, ReleaseBundleCustomWebhookArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ReleaseBundleCustomWebhook(String name, ReleaseBundleCustomWebhookArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/releaseBundleCustomWebhook:ReleaseBundleCustomWebhook", name, args == null ? ReleaseBundleCustomWebhookArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ReleaseBundleCustomWebhook(String name, Output<String> id, @Nullable ReleaseBundleCustomWebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/releaseBundleCustomWebhook:ReleaseBundleCustomWebhook", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
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
    public static ReleaseBundleCustomWebhook get(String name, Output<String> id, @Nullable ReleaseBundleCustomWebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ReleaseBundleCustomWebhook(name, id, state, options);
    }
}
