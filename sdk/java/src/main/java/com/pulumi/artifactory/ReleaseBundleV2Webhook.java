// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.ReleaseBundleV2WebhookArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.ReleaseBundleV2WebhookState;
import com.pulumi.artifactory.outputs.ReleaseBundleV2WebhookCriteria;
import com.pulumi.artifactory.outputs.ReleaseBundleV2WebhookHandler;
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
 * Provides an Artifactory webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.
 * 
 * ## Example Usage
 * 
 * .
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.ReleaseBundleV2Webhook;
 * import com.pulumi.artifactory.ReleaseBundleV2WebhookArgs;
 * import com.pulumi.artifactory.inputs.ReleaseBundleV2WebhookCriteriaArgs;
 * import com.pulumi.artifactory.inputs.ReleaseBundleV2WebhookHandlerArgs;
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
 *         var release_bundle_v2_webhook = new ReleaseBundleV2Webhook("release-bundle-v2-webhook", ReleaseBundleV2WebhookArgs.builder()
 *             .key("release-bundle-v2-webhook")
 *             .eventTypes(            
 *                 "release_bundle_v2_started",
 *                 "release_bundle_v2_failed",
 *                 "release_bundle_v2_completed")
 *             .criteria(ReleaseBundleV2WebhookCriteriaArgs.builder()
 *                 .anyReleaseBundle(false)
 *                 .selectedReleaseBundles("bundle-name")
 *                 .includePatterns("foo/**")
 *                 .excludePatterns("bar/**")
 *                 .build())
 *             .handlers(ReleaseBundleV2WebhookHandlerArgs.builder()
 *                 .url("http://tempurl.org/webhook")
 *                 .secret("some-secret")
 *                 .proxy("proxy-key")
 *                 .customHttpHeaders(Map.ofEntries(
 *                     Map.entry("header-1", "value-1"),
 *                     Map.entry("header-2", "value-2")
 *                 ))
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="artifactory:index/releaseBundleV2Webhook:ReleaseBundleV2Webhook")
public class ReleaseBundleV2Webhook extends com.pulumi.resources.CustomResource {
    /**
     * Specifies where the webhook will be applied on which repositories.
     * 
     */
    @Export(name="criteria", refs={ReleaseBundleV2WebhookCriteria.class}, tree="[0]")
    private Output<ReleaseBundleV2WebhookCriteria> criteria;

    /**
     * @return Specifies where the webhook will be applied on which repositories.
     * 
     */
    public Output<ReleaseBundleV2WebhookCriteria> criteria() {
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
     * List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `release_bundle_v2_started`, `release_bundle_v2_failed`, `release_bundle_v2_completed`.
     * 
     */
    @Export(name="eventTypes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> eventTypes;

    /**
     * @return List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `release_bundle_v2_started`, `release_bundle_v2_failed`, `release_bundle_v2_completed`.
     * 
     */
    public Output<List<String>> eventTypes() {
        return this.eventTypes;
    }
    /**
     * At least one is required.
     * 
     */
    @Export(name="handlers", refs={List.class,ReleaseBundleV2WebhookHandler.class}, tree="[0,1]")
    private Output<List<ReleaseBundleV2WebhookHandler>> handlers;

    /**
     * @return At least one is required.
     * 
     */
    public Output<List<ReleaseBundleV2WebhookHandler>> handlers() {
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
    public ReleaseBundleV2Webhook(java.lang.String name) {
        this(name, ReleaseBundleV2WebhookArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ReleaseBundleV2Webhook(java.lang.String name, ReleaseBundleV2WebhookArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ReleaseBundleV2Webhook(java.lang.String name, ReleaseBundleV2WebhookArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/releaseBundleV2Webhook:ReleaseBundleV2Webhook", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ReleaseBundleV2Webhook(java.lang.String name, Output<java.lang.String> id, @Nullable ReleaseBundleV2WebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/releaseBundleV2Webhook:ReleaseBundleV2Webhook", name, state, makeResourceOptions(options, id), false);
    }

    private static ReleaseBundleV2WebhookArgs makeArgs(ReleaseBundleV2WebhookArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ReleaseBundleV2WebhookArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static ReleaseBundleV2Webhook get(java.lang.String name, Output<java.lang.String> id, @Nullable ReleaseBundleV2WebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ReleaseBundleV2Webhook(name, id, state, options);
    }
}
