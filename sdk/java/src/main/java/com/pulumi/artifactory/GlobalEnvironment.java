// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.GlobalEnvironmentArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.GlobalEnvironmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a global environment resource. This can be used to create and manage global environment.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.GlobalEnvironment;
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
 *         var test_env = new GlobalEnvironment(&#34;test-env&#34;);
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * ```sh
 * $ pulumi import artifactory:index/globalEnvironment:GlobalEnvironment dev-env myenv
 * ```
 * 
 */
@ResourceType(type="artifactory:index/globalEnvironment:GlobalEnvironment")
public class GlobalEnvironment extends com.pulumi.resources.CustomResource {
    /**
     * Name must start with a letter and contain letters, digits and `-` character. The maximum length is 32 characters
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name must start with a letter and contain letters, digits and `-` character. The maximum length is 32 characters
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GlobalEnvironment(String name) {
        this(name, GlobalEnvironmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GlobalEnvironment(String name, @Nullable GlobalEnvironmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GlobalEnvironment(String name, @Nullable GlobalEnvironmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/globalEnvironment:GlobalEnvironment", name, args == null ? GlobalEnvironmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GlobalEnvironment(String name, Output<String> id, @Nullable GlobalEnvironmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/globalEnvironment:GlobalEnvironment", name, state, makeResourceOptions(options, id));
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
    public static GlobalEnvironment get(String name, Output<String> id, @Nullable GlobalEnvironmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GlobalEnvironment(name, id, state, options);
    }
}
