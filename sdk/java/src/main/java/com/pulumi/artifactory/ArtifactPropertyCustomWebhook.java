// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.ArtifactPropertyCustomWebhookArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.ArtifactPropertyCustomWebhookState;
import com.pulumi.artifactory.outputs.ArtifactPropertyCustomWebhookCriteria;
import com.pulumi.artifactory.outputs.ArtifactPropertyCustomWebhookHandler;
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
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.LocalGenericRepository;
 * import com.pulumi.artifactory.LocalGenericRepositoryArgs;
 * import com.pulumi.artifactory.ArtifactPropertyCustomWebhook;
 * import com.pulumi.artifactory.ArtifactPropertyCustomWebhookArgs;
 * import com.pulumi.artifactory.inputs.ArtifactPropertyCustomWebhookCriteriaArgs;
 * import com.pulumi.artifactory.inputs.ArtifactPropertyCustomWebhookHandlerArgs;
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
 *         var my_generic_local = new LocalGenericRepository("my-generic-local", LocalGenericRepositoryArgs.builder()
 *             .key("my-generic-local")
 *             .build());
 * 
 *         var artifact_custom_webhook = new ArtifactPropertyCustomWebhook("artifact-custom-webhook", ArtifactPropertyCustomWebhookArgs.builder()
 *             .key("artifact-property-custom-webhook")
 *             .eventTypes(            
 *                 "added",
 *                 "deleted")
 *             .criteria(ArtifactPropertyCustomWebhookCriteriaArgs.builder()
 *                 .anyLocal(true)
 *                 .anyRemote(false)
 *                 .repoKeys(my_generic_local.key())
 *                 .includePatterns("foo/**")
 *                 .excludePatterns("bar/**")
 *                 .build())
 *             .handlers(ArtifactPropertyCustomWebhookHandlerArgs.builder()
 *                 .url("https://tempurl.org")
 *                 .secrets(Map.ofEntries(
 *                     Map.entry("secretName1", "value1"),
 *                     Map.entry("secretName2", "value2")
 *                 ))
 *                 .httpHeaders(Map.ofEntries(
 *                     Map.entry("headerName1", "value1"),
 *                     Map.entry("headerName2", "value2")
 *                 ))
 *                 .payload("{ \"ref\": \"main\" , \"inputs\": { \"artifact_path\": \"test-repo/repo-path\" } }")
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(my_generic_local)
 *                 .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="artifactory:index/artifactPropertyCustomWebhook:ArtifactPropertyCustomWebhook")
public class ArtifactPropertyCustomWebhook extends com.pulumi.resources.CustomResource {
    /**
     * Specifies where the webhook will be applied on which repositories.
     * 
     */
    @Export(name="criteria", refs={ArtifactPropertyCustomWebhookCriteria.class}, tree="[0]")
    private Output<ArtifactPropertyCustomWebhookCriteria> criteria;

    /**
     * @return Specifies where the webhook will be applied on which repositories.
     * 
     */
    public Output<ArtifactPropertyCustomWebhookCriteria> criteria() {
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
     * List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `added`, `deleted`.
     * 
     */
    @Export(name="eventTypes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> eventTypes;

    /**
     * @return List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `added`, `deleted`.
     * 
     */
    public Output<List<String>> eventTypes() {
        return this.eventTypes;
    }
    /**
     * At least one is required.
     * 
     */
    @Export(name="handlers", refs={List.class,ArtifactPropertyCustomWebhookHandler.class}, tree="[0,1]")
    private Output<List<ArtifactPropertyCustomWebhookHandler>> handlers;

    /**
     * @return At least one is required.
     * 
     */
    public Output<List<ArtifactPropertyCustomWebhookHandler>> handlers() {
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
    public ArtifactPropertyCustomWebhook(String name) {
        this(name, ArtifactPropertyCustomWebhookArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ArtifactPropertyCustomWebhook(String name, ArtifactPropertyCustomWebhookArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ArtifactPropertyCustomWebhook(String name, ArtifactPropertyCustomWebhookArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/artifactPropertyCustomWebhook:ArtifactPropertyCustomWebhook", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private ArtifactPropertyCustomWebhook(String name, Output<String> id, @Nullable ArtifactPropertyCustomWebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/artifactPropertyCustomWebhook:ArtifactPropertyCustomWebhook", name, state, makeResourceOptions(options, id));
    }

    private static ArtifactPropertyCustomWebhookArgs makeArgs(ArtifactPropertyCustomWebhookArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ArtifactPropertyCustomWebhookArgs.Empty : args;
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
    public static ArtifactPropertyCustomWebhook get(String name, Output<String> id, @Nullable ArtifactPropertyCustomWebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ArtifactPropertyCustomWebhook(name, id, state, options);
    }
}
