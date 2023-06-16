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

@ResourceType(type="artifactory:index/certificate:Certificate")
public class Certificate extends com.pulumi.resources.CustomResource {
    @Export(name="alias", type=String.class, parameters={})
    private Output<String> alias;

    public Output<String> alias() {
        return this.alias;
    }
    @Export(name="content", type=String.class, parameters={})
    private Output</* @Nullable */ String> content;

    public Output<Optional<String>> content() {
        return Codegen.optional(this.content);
    }
    @Export(name="file", type=String.class, parameters={})
    private Output</* @Nullable */ String> file;

    public Output<Optional<String>> file() {
        return Codegen.optional(this.file);
    }
    @Export(name="fingerprint", type=String.class, parameters={})
    private Output<String> fingerprint;

    public Output<String> fingerprint() {
        return this.fingerprint;
    }
    @Export(name="issuedBy", type=String.class, parameters={})
    private Output<String> issuedBy;

    public Output<String> issuedBy() {
        return this.issuedBy;
    }
    @Export(name="issuedOn", type=String.class, parameters={})
    private Output<String> issuedOn;

    public Output<String> issuedOn() {
        return this.issuedOn;
    }
    @Export(name="issuedTo", type=String.class, parameters={})
    private Output<String> issuedTo;

    public Output<String> issuedTo() {
        return this.issuedTo;
    }
    @Export(name="validUntil", type=String.class, parameters={})
    private Output<String> validUntil;

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
        super("artifactory:index/certificate:Certificate", name, args == null ? CertificateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Certificate(String name, Output<String> id, @Nullable CertificateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/certificate:Certificate", name, state, makeResourceOptions(options, id));
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
