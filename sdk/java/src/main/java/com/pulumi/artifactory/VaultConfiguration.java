// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.VaultConfigurationArgs;
import com.pulumi.artifactory.inputs.VaultConfigurationState;
import com.pulumi.artifactory.outputs.VaultConfigurationConfig;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * This resource enables you to configure an external vault connector to use as a centralized secret management tool for the keys used to sign packages. For more information, see [JFrog documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/vault).
 * This feature is supported with Enterprise X and Enterprise+ licenses.
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
 * import com.pulumi.artifactory.VaultConfiguration;
 * import com.pulumi.artifactory.VaultConfigurationArgs;
 * import com.pulumi.artifactory.inputs.VaultConfigurationConfigArgs;
 * import com.pulumi.artifactory.inputs.VaultConfigurationConfigAuthArgs;
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
 *         var my_vault_config_app_role = new VaultConfiguration("my-vault-config-app-role", VaultConfigurationArgs.builder()
 *             .name("my-vault-config-app-role")
 *             .config(VaultConfigurationConfigArgs.builder()
 *                 .url("http://127.0.0.1:8200")
 *                 .auth(VaultConfigurationConfigAuthArgs.builder()
 *                     .type("AppRole")
 *                     .roleId("1b62ff05...")
 *                     .secretId("acbd6657...")
 *                     .build())
 *                 .mounts(VaultConfigurationConfigMountArgs.builder()
 *                     .path("secret")
 *                     .type("KV2")
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var my_vault_config_cert = new VaultConfiguration("my-vault-config-cert", VaultConfigurationArgs.builder()
 *             .name("my-vault-config-cert")
 *             .config(VaultConfigurationConfigArgs.builder()
 *                 .url("http://127.0.0.1:8200")
 *                 .auth(VaultConfigurationConfigAuthArgs.builder()
 *                     .type("Certificate")
 *                     .certificate(StdFunctions.file(FileArgs.builder()
 *                         .input("samples/public.pem")
 *                         .build()).result())
 *                     .certificateKey(StdFunctions.file(FileArgs.builder()
 *                         .input("samples/private.pem")
 *                         .build()).result())
 *                     .build())
 *                 .mounts(VaultConfigurationConfigMountArgs.builder()
 *                     .path("secret")
 *                     .type("KV2")
 *                     .build())
 *                 .build())
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
 * $ pulumi import artifactory:index/vaultConfiguration:VaultConfiguration my-vault-config my-vault-config
 * ```
 * 
 */
@ResourceType(type="artifactory:index/vaultConfiguration:VaultConfiguration")
public class VaultConfiguration extends com.pulumi.resources.CustomResource {
    @Export(name="config", refs={VaultConfigurationConfig.class}, tree="[0]")
    private Output<VaultConfigurationConfig> config;

    public Output<VaultConfigurationConfig> config() {
        return this.config;
    }
    /**
     * Name of the Vault configuration
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the Vault configuration
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VaultConfiguration(String name) {
        this(name, VaultConfigurationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VaultConfiguration(String name, VaultConfigurationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VaultConfiguration(String name, VaultConfigurationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/vaultConfiguration:VaultConfiguration", name, args == null ? VaultConfigurationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VaultConfiguration(String name, Output<String> id, @Nullable VaultConfigurationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/vaultConfiguration:VaultConfiguration", name, state, makeResourceOptions(options, id));
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
    public static VaultConfiguration get(String name, Output<String> id, @Nullable VaultConfigurationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VaultConfiguration(name, id, state, options);
    }
}
