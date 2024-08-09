// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.ArtifactLifecycleWebhookArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.ArtifactLifecycleWebhookState;
import com.pulumi.artifactory.outputs.ArtifactLifecycleWebhookHandler;
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
 * import com.pulumi.artifactory.ArtifactLifecycleWebhook;
 * import com.pulumi.artifactory.ArtifactLifecycleWebhookArgs;
 * import com.pulumi.artifactory.inputs.ArtifactLifecycleWebhookHandlerArgs;
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
 *         var artifact_lifecycle_webhook = new ArtifactLifecycleWebhook("artifact-lifecycle-webhook", ArtifactLifecycleWebhookArgs.builder()
 *             .key("artifact-lifecycle-webhook")
 *             .eventTypes(            
 *                 "archive",
 *                 "restore")
 *             .handlers(ArtifactLifecycleWebhookHandlerArgs.builder()
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
@ResourceType(type="artifactory:index/artifactLifecycleWebhook:ArtifactLifecycleWebhook")
public class ArtifactLifecycleWebhook extends com.pulumi.resources.CustomResource {
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
     * Status of webhook. Default to `true`
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return Status of webhook. Default to `true`
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * List of event triggers for the Webhook. Allow values: `archive`, `restore`
     * 
     */
    @Export(name="eventTypes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> eventTypes;

    /**
     * @return List of event triggers for the Webhook. Allow values: `archive`, `restore`
     * 
     */
    public Output<List<String>> eventTypes() {
        return this.eventTypes;
    }
    /**
     * At least one is required.
     * 
     */
    @Export(name="handlers", refs={List.class,ArtifactLifecycleWebhookHandler.class}, tree="[0,1]")
    private Output<List<ArtifactLifecycleWebhookHandler>> handlers;

    /**
     * @return At least one is required.
     * 
     */
    public Output<List<ArtifactLifecycleWebhookHandler>> handlers() {
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
    public ArtifactLifecycleWebhook(java.lang.String name) {
        this(name, ArtifactLifecycleWebhookArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ArtifactLifecycleWebhook(java.lang.String name, ArtifactLifecycleWebhookArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ArtifactLifecycleWebhook(java.lang.String name, ArtifactLifecycleWebhookArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/artifactLifecycleWebhook:ArtifactLifecycleWebhook", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ArtifactLifecycleWebhook(java.lang.String name, Output<java.lang.String> id, @Nullable ArtifactLifecycleWebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/artifactLifecycleWebhook:ArtifactLifecycleWebhook", name, state, makeResourceOptions(options, id), false);
    }

    private static ArtifactLifecycleWebhookArgs makeArgs(ArtifactLifecycleWebhookArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ArtifactLifecycleWebhookArgs.Empty : args;
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
    public static ArtifactLifecycleWebhook get(java.lang.String name, Output<java.lang.String> id, @Nullable ArtifactLifecycleWebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ArtifactLifecycleWebhook(name, id, state, options);
    }
}
