// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.CertificateArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.CertificateState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an Artifactory certificate resource. This can be used to create and manage Artifactory certificates which can be used as client authentication against remote repositories.
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
 * import com.pulumi.artifactory.Certificate;
 * import com.pulumi.artifactory.CertificateArgs;
 * import com.pulumi.artifactory.RemoteMavenRepository;
 * import com.pulumi.artifactory.RemoteMavenRepositoryArgs;
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
 *         // Create a new Artifactory certificate called my-cert
 *         var my_cert = new Certificate("my-cert", CertificateArgs.builder()
 *             .alias("my-cert")
 *             .content(StdFunctions.file(FileArgs.builder()
 *                 .input("/path/to/bundle.pem")
 *                 .build()).result())
 *             .build());
 * 
 *         // This can then be used by a remote repository
 *         var my_remote = new RemoteMavenRepository("my-remote", RemoteMavenRepositoryArgs.builder()
 *             .clientTlsCertificate(my_cert.alias())
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
 * Certificates can be imported using their alias, e.g.
 * 
 * ```sh
 * $ pulumi import artifactory:index/certificate:Certificate my-cert my-cert
 * ```
 * 
 */
@ResourceType(type="artifactory:index/certificate:Certificate")
public class Certificate extends com.pulumi.resources.CustomResource {
    /**
     * Name of certificate.
     * 
     */
    @Export(name="alias", refs={String.class}, tree="[0]")
    private Output<String> alias;

    /**
     * @return Name of certificate.
     * 
     */
    public Output<String> alias() {
        return this.alias;
    }
    /**
     * PEM-encoded client certificate and private key. Cannot be set with `file` attribute simultaneously.
     * 
     */
    @Export(name="content", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> content;

    /**
     * @return PEM-encoded client certificate and private key. Cannot be set with `file` attribute simultaneously.
     * 
     */
    public Output<Optional<String>> content() {
        return Codegen.optional(this.content);
    }
    /**
     * Path to the PEM file. Cannot be set with `content` attribute simultaneously.
     * 
     */
    @Export(name="file", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> file;

    /**
     * @return Path to the PEM file. Cannot be set with `content` attribute simultaneously.
     * 
     */
    public Output<Optional<String>> file() {
        return Codegen.optional(this.file);
    }
    /**
     * SHA256 fingerprint of the certificate.
     * 
     */
    @Export(name="fingerprint", refs={String.class}, tree="[0]")
    private Output<String> fingerprint;

    /**
     * @return SHA256 fingerprint of the certificate.
     * 
     */
    public Output<String> fingerprint() {
        return this.fingerprint;
    }
    /**
     * Name of the certificate authority that issued the certificate.
     * 
     */
    @Export(name="issuedBy", refs={String.class}, tree="[0]")
    private Output<String> issuedBy;

    /**
     * @return Name of the certificate authority that issued the certificate.
     * 
     */
    public Output<String> issuedBy() {
        return this.issuedBy;
    }
    /**
     * The time &amp; date when the certificate is valid from.
     * 
     */
    @Export(name="issuedOn", refs={String.class}, tree="[0]")
    private Output<String> issuedOn;

    /**
     * @return The time &amp; date when the certificate is valid from.
     * 
     */
    public Output<String> issuedOn() {
        return this.issuedOn;
    }
    /**
     * Name of whom the certificate has been issued to.
     * 
     */
    @Export(name="issuedTo", refs={String.class}, tree="[0]")
    private Output<String> issuedTo;

    /**
     * @return Name of whom the certificate has been issued to.
     * 
     */
    public Output<String> issuedTo() {
        return this.issuedTo;
    }
    /**
     * The time &amp; date when the certificate expires.
     * 
     */
    @Export(name="validUntil", refs={String.class}, tree="[0]")
    private Output<String> validUntil;

    /**
     * @return The time &amp; date when the certificate expires.
     * 
     */
    public Output<String> validUntil() {
        return this.validUntil;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Certificate(String name) {
        this(name, CertificateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Certificate(String name, CertificateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Certificate(String name, CertificateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/certificate:Certificate", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private Certificate(String name, Output<String> id, @Nullable CertificateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/certificate:Certificate", name, state, makeResourceOptions(options, id));
    }

    private static CertificateArgs makeArgs(CertificateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? CertificateArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "content",
                "file"
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
    public static Certificate get(String name, Output<String> id, @Nullable CertificateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Certificate(name, id, state, options);
    }
}
