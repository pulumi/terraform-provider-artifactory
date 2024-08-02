// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.MailServerArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.MailServerState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an Artifactory Mail Server resource. This can be used to create and manage Artifactory mail server configuration.
 * 
 * ## Example Usage
 * 
 * ### S
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.MailServer;
 * import com.pulumi.artifactory.MailServerArgs;
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
 *         var mymailserver = new MailServer("mymailserver", MailServerArgs.builder()
 *             .enabled(true)
 *             .artifactoryUrl("http://tempurl.org")
 *             .from("test{@literal @}jfrog.com")
 *             .host("http://tempurl.org")
 *             .username("test-user")
 *             .password("test-password")
 *             .port(25)
 *             .subjectPrefix("[Test]")
 *             .useSsl(true)
 *             .useTls(true)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * ```sh
 * $ pulumi import artifactory:index/mailServer:MailServer my-mail-server mymailserver
 * ```
 * 
 * ~&gt;The `password` attribute is not retrievable from Artifactory thus there will be state drift after importing this resource.
 * 
 */
@ResourceType(type="artifactory:index/mailServer:MailServer")
public class MailServer extends com.pulumi.resources.CustomResource {
    /**
     * The Artifactory URL to to link to in all outgoing messages.
     * 
     */
    @Export(name="artifactoryUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> artifactoryUrl;

    /**
     * @return The Artifactory URL to to link to in all outgoing messages.
     * 
     */
    public Output<Optional<String>> artifactoryUrl() {
        return Codegen.optional(this.artifactoryUrl);
    }
    /**
     * When set, mail notifications are enabled.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enabled;

    /**
     * @return When set, mail notifications are enabled.
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }
    /**
     * The &#39;from&#39; address header to use in all outgoing messages.
     * 
     */
    @Export(name="from", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> from;

    /**
     * @return The &#39;from&#39; address header to use in all outgoing messages.
     * 
     */
    public Output<Optional<String>> from() {
        return Codegen.optional(this.from);
    }
    /**
     * The mail server IP address / DNS.
     * 
     */
    @Export(name="host", refs={String.class}, tree="[0]")
    private Output<String> host;

    /**
     * @return The mail server IP address / DNS.
     * 
     */
    public Output<String> host() {
        return this.host;
    }
    /**
     * The password for authentication with the mail server.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> password;

    /**
     * @return The password for authentication with the mail server.
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * The port number of the mail server.
     * 
     */
    @Export(name="port", refs={Integer.class}, tree="[0]")
    private Output<Integer> port;

    /**
     * @return The port number of the mail server.
     * 
     */
    public Output<Integer> port() {
        return this.port;
    }
    /**
     * A prefix to use for the subject of all outgoing mails.
     * 
     */
    @Export(name="subjectPrefix", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> subjectPrefix;

    /**
     * @return A prefix to use for the subject of all outgoing mails.
     * 
     */
    public Output<Optional<String>> subjectPrefix() {
        return Codegen.optional(this.subjectPrefix);
    }
    /**
     * When set to &#39;true&#39;, uses a secure connection to the mail server.
     * 
     */
    @Export(name="useSsl", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> useSsl;

    /**
     * @return When set to &#39;true&#39;, uses a secure connection to the mail server.
     * 
     */
    public Output<Boolean> useSsl() {
        return this.useSsl;
    }
    /**
     * When set to &#39;true&#39;, uses Transport Layer Security when connecting to the mail server.
     * 
     */
    @Export(name="useTls", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> useTls;

    /**
     * @return When set to &#39;true&#39;, uses Transport Layer Security when connecting to the mail server.
     * 
     */
    public Output<Boolean> useTls() {
        return this.useTls;
    }
    /**
     * The username for authentication with the mail server.
     * 
     */
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> username;

    /**
     * @return The username for authentication with the mail server.
     * 
     */
    public Output<Optional<String>> username() {
        return Codegen.optional(this.username);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MailServer(String name) {
        this(name, MailServerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MailServer(String name, MailServerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MailServer(String name, MailServerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/mailServer:MailServer", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private MailServer(String name, Output<String> id, @Nullable MailServerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/mailServer:MailServer", name, state, makeResourceOptions(options, id));
    }

    private static MailServerArgs makeArgs(MailServerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? MailServerArgs.Empty : args;
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
    public static MailServer get(String name, Output<String> id, @Nullable MailServerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MailServer(name, id, state, options);
    }
}
