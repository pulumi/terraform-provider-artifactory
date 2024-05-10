// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.KeypairArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.KeypairState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * RSA key pairs are used to sign and verify the Alpine Linux index files in JFrog Artifactory, while GPG key pairs are
 * used to sign and validate packages integrity in JFrog Distribution. The JFrog Platform enables you to manage multiple RSA and GPG signing keys through the Keys Management UI and REST API. The JFrog Platform supports managing multiple pairs of GPG signing keys to sign packages for authentication of several package types such as Debian, Opkg, and RPM through the Keys Management UI and REST API.
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
 * import com.pulumi.artifactory.Keypair;
 * import com.pulumi.artifactory.KeypairArgs;
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
 *         var some_keypair_6543461672124900137 = new Keypair("some-keypair-6543461672124900137", KeypairArgs.builder()        
 *             .pairName("some-keypair-6543461672124900137")
 *             .pairType("RSA")
 *             .alias("some-alias-6543461672124900137")
 *             .privateKey(StdFunctions.file(FileArgs.builder()
 *                 .input("samples/rsa.priv")
 *                 .build()).result())
 *             .publicKey(StdFunctions.file(FileArgs.builder()
 *                 .input("samples/rsa.pub")
 *                 .build()).result())
 *             .passphrase("PASSPHRASE")
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
 * Keypair can be imported using the pair name, e.g.
 * 
 * ```sh
 * $ pulumi import artifactory:index/keypair:Keypair my-keypair my-keypair-name
 * ```
 * 
 */
@ResourceType(type="artifactory:index/keypair:Keypair")
public class Keypair extends com.pulumi.resources.CustomResource {
    /**
     * Will be used as a filename when retrieving the public key via REST API.
     * 
     */
    @Export(name="alias", refs={String.class}, tree="[0]")
    private Output<String> alias;

    /**
     * @return Will be used as a filename when retrieving the public key via REST API.
     * 
     */
    public Output<String> alias() {
        return this.alias;
    }
    /**
     * A unique identifier for the Key Pair record.
     * 
     */
    @Export(name="pairName", refs={String.class}, tree="[0]")
    private Output<String> pairName;

    /**
     * @return A unique identifier for the Key Pair record.
     * 
     */
    public Output<String> pairName() {
        return this.pairName;
    }
    /**
     * Key Pair type. Supported types - GPG and RSA.
     * 
     */
    @Export(name="pairType", refs={String.class}, tree="[0]")
    private Output<String> pairType;

    /**
     * @return Key Pair type. Supported types - GPG and RSA.
     * 
     */
    public Output<String> pairType() {
        return this.pairType;
    }
    /**
     * Passphrase will be used to decrypt the private key. Validated server side.
     * 
     */
    @Export(name="passphrase", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> passphrase;

    /**
     * @return Passphrase will be used to decrypt the private key. Validated server side.
     * 
     */
    public Output<Optional<String>> passphrase() {
        return Codegen.optional(this.passphrase);
    }
    /**
     * Private key. PEM format will be validated. Must not include extranous spaces or tabs.
     * 
     */
    @Export(name="privateKey", refs={String.class}, tree="[0]")
    private Output<String> privateKey;

    /**
     * @return Private key. PEM format will be validated. Must not include extranous spaces or tabs.
     * 
     */
    public Output<String> privateKey() {
        return this.privateKey;
    }
    /**
     * Public key. PEM format will be validated. Must not include extranous spaces or tabs.
     * 
     * Artifactory REST API call &#39;Get Key Pair&#39; doesn&#39;t return attributes `private_key` and `passphrase`, but consumes these keys in the POST call.
     * 
     */
    @Export(name="publicKey", refs={String.class}, tree="[0]")
    private Output<String> publicKey;

    /**
     * @return Public key. PEM format will be validated. Must not include extranous spaces or tabs.
     * 
     * Artifactory REST API call &#39;Get Key Pair&#39; doesn&#39;t return attributes `private_key` and `passphrase`, but consumes these keys in the POST call.
     * 
     */
    public Output<String> publicKey() {
        return this.publicKey;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Keypair(String name) {
        this(name, KeypairArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Keypair(String name, KeypairArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Keypair(String name, KeypairArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/keypair:Keypair", name, args == null ? KeypairArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Keypair(String name, Output<String> id, @Nullable KeypairState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/keypair:Keypair", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "passphrase",
                "privateKey"
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
    public static Keypair get(String name, Output<String> id, @Nullable KeypairState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Keypair(name, id, state, options);
    }
}
