// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.ArtifactWebhookArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.ArtifactWebhookState;
import com.pulumi.artifactory.outputs.ArtifactWebhookCriteria;
import com.pulumi.artifactory.outputs.ArtifactWebhookHandler;
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
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.LocalGenericRepository;
 * import com.pulumi.artifactory.LocalGenericRepositoryArgs;
 * import com.pulumi.artifactory.ArtifactWebhook;
 * import com.pulumi.artifactory.ArtifactWebhookArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var my_generic_local = new LocalGenericRepository(&#34;my-generic-local&#34;, LocalGenericRepositoryArgs.builder()        
 *             .key(&#34;my-generic-local&#34;)
 *             .build());
 * 
 *         var artifact_webhook = new ArtifactWebhook(&#34;artifact-webhook&#34;, ArtifactWebhookArgs.builder()        
 *             .key(&#34;artifact-webhook&#34;)
 *             .eventTypes(            
 *                 &#34;deployed&#34;,
 *                 &#34;deleted&#34;,
 *                 &#34;moved&#34;,
 *                 &#34;copied&#34;)
 *             .criteria(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
 *             .handlers(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(my_generic_local)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="artifactory:index/artifactWebhook:ArtifactWebhook")
public class ArtifactWebhook extends com.pulumi.resources.CustomResource {
    /**
     * Specifies where the webhook will be applied on which repositories.
     * 
     */
    @Export(name="criteria", type=ArtifactWebhookCriteria.class, parameters={})
    private Output<ArtifactWebhookCriteria> criteria;

    /**
     * @return Specifies where the webhook will be applied on which repositories.
     * 
     */
    public Output<ArtifactWebhookCriteria> criteria() {
        return this.criteria;
    }
    /**
     * Webhook description. Max length 1000 characters.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
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
    @Export(name="enabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return Status of webhook. Default to `true`.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `deployed`, `deleted`, `moved`, `copied`, `cached`.
     * 
     */
    @Export(name="eventTypes", type=List.class, parameters={String.class})
    private Output<List<String>> eventTypes;

    /**
     * @return List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `deployed`, `deleted`, `moved`, `copied`, `cached`.
     * 
     */
    public Output<List<String>> eventTypes() {
        return this.eventTypes;
    }
    /**
     * At least one is required.
     * 
     */
    @Export(name="handlers", type=List.class, parameters={ArtifactWebhookHandler.class})
    private Output<List<ArtifactWebhookHandler>> handlers;

    /**
     * @return At least one is required.
     * 
     */
    public Output<List<ArtifactWebhookHandler>> handlers() {
        return this.handlers;
    }
    /**
     * The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     * 
     */
    @Export(name="key", type=String.class, parameters={})
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
    public ArtifactWebhook(String name) {
        this(name, ArtifactWebhookArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ArtifactWebhook(String name, ArtifactWebhookArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ArtifactWebhook(String name, ArtifactWebhookArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/artifactWebhook:ArtifactWebhook", name, args == null ? ArtifactWebhookArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ArtifactWebhook(String name, Output<String> id, @Nullable ArtifactWebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/artifactWebhook:ArtifactWebhook", name, state, makeResourceOptions(options, id));
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
    public static ArtifactWebhook get(String name, Output<String> id, @Nullable ArtifactWebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ArtifactWebhook(name, id, state, options);
    }
}
