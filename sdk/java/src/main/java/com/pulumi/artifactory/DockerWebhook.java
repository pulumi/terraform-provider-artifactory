// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.DockerWebhookArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.DockerWebhookState;
import com.pulumi.artifactory.outputs.DockerWebhookCriteria;
import com.pulumi.artifactory.outputs.DockerWebhookHandler;
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
 * import com.pulumi.artifactory.DockerV2Repository;
 * import com.pulumi.artifactory.DockerV2RepositoryArgs;
 * import com.pulumi.artifactory.DockerWebhook;
 * import com.pulumi.artifactory.DockerWebhookArgs;
 * import com.pulumi.artifactory.inputs.DockerWebhookCriteriaArgs;
 * import com.pulumi.artifactory.inputs.DockerWebhookHandlerArgs;
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
 *         var my_docker_local = new DockerV2Repository(&#34;my-docker-local&#34;, DockerV2RepositoryArgs.builder()        
 *             .key(&#34;my-docker-local&#34;)
 *             .build());
 * 
 *         var docker_webhook = new DockerWebhook(&#34;docker-webhook&#34;, DockerWebhookArgs.builder()        
 *             .key(&#34;docker-webhook&#34;)
 *             .eventTypes(            
 *                 &#34;pushed&#34;,
 *                 &#34;deleted&#34;,
 *                 &#34;promoted&#34;)
 *             .criteria(DockerWebhookCriteriaArgs.builder()
 *                 .anyLocal(true)
 *                 .anyRemote(false)
 *                 .repoKeys(my_docker_local.key())
 *                 .includePatterns(&#34;foo/**&#34;)
 *                 .excludePatterns(&#34;bar/**&#34;)
 *                 .build())
 *             .handlers(DockerWebhookHandlerArgs.builder()
 *                 .url(&#34;http://tempurl.org/webhook&#34;)
 *                 .secret(&#34;some-secret&#34;)
 *                 .proxy(&#34;proxy-key&#34;)
 *                 .customHttpHeaders(Map.ofEntries(
 *                     Map.entry(&#34;header-1&#34;, &#34;value-1&#34;),
 *                     Map.entry(&#34;header-2&#34;, &#34;value-2&#34;)
 *                 ))
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(my_docker_local)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="artifactory:index/dockerWebhook:DockerWebhook")
public class DockerWebhook extends com.pulumi.resources.CustomResource {
    /**
     * Specifies where the webhook will be applied on which repositories.
     * 
     */
    @Export(name="criteria", type=DockerWebhookCriteria.class, parameters={})
    private Output<DockerWebhookCriteria> criteria;

    /**
     * @return Specifies where the webhook will be applied on which repositories.
     * 
     */
    public Output<DockerWebhookCriteria> criteria() {
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
     * List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `pushed`, `deleted`, `promoted`.
     * 
     */
    @Export(name="eventTypes", type=List.class, parameters={String.class})
    private Output<List<String>> eventTypes;

    /**
     * @return List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `pushed`, `deleted`, `promoted`.
     * 
     */
    public Output<List<String>> eventTypes() {
        return this.eventTypes;
    }
    /**
     * At least one is required.
     * 
     */
    @Export(name="handlers", type=List.class, parameters={DockerWebhookHandler.class})
    private Output<List<DockerWebhookHandler>> handlers;

    /**
     * @return At least one is required.
     * 
     */
    public Output<List<DockerWebhookHandler>> handlers() {
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
    public DockerWebhook(String name) {
        this(name, DockerWebhookArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DockerWebhook(String name, DockerWebhookArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DockerWebhook(String name, DockerWebhookArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/dockerWebhook:DockerWebhook", name, args == null ? DockerWebhookArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DockerWebhook(String name, Output<String> id, @Nullable DockerWebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/dockerWebhook:DockerWebhook", name, state, makeResourceOptions(options, id));
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
    public static DockerWebhook get(String name, Output<String> id, @Nullable DockerWebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DockerWebhook(name, id, state, options);
    }
}
