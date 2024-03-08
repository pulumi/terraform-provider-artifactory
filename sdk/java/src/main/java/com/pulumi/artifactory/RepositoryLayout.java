// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.RepositoryLayoutArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.RepositoryLayoutState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can be used to manage Artifactory&#39;s Repository Layout settings. See [Repository Layouts](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts) in the Artifactory Wiki documentation for more details.
 * 
 * ~&gt;The `artifactory.RepositoryLayout` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.
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
 * import com.pulumi.artifactory.RepositoryLayout;
 * import com.pulumi.artifactory.RepositoryLayoutArgs;
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
 *         var custom_layout = new RepositoryLayout(&#34;custom-layout&#34;, RepositoryLayoutArgs.builder()        
 *             .artifactPathPattern(&#34;[orgPath]/[module]/[baseRev](-[folderItegRev])/[module]-[baseRev](-[fileItegRev])(-[classifier]).[ext]&#34;)
 *             .descriptorPathPattern(&#34;[orgPath]/[module]/[baseRev](-[folderItegRev])/[module]-[baseRev](-[fileItegRev])(-[classifier]).pom&#34;)
 *             .distinctiveDescriptorPathPattern(true)
 *             .fileIntegrationRevisionRegexp(&#34;Foo|(?:(?:[0-9]{8}.[0-9]{6})-(?:[0-9]+))&#34;)
 *             .folderIntegrationRevisionRegexp(&#34;Foo&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Repository layout can be imported using its name, e.g.
 * 
 * ```sh
 * $ pulumi import artifactory:index/repositoryLayout:RepositoryLayout custom-layout custom-layout
 * ```
 * 
 */
@ResourceType(type="artifactory:index/repositoryLayout:RepositoryLayout")
public class RepositoryLayout extends com.pulumi.resources.CustomResource {
    /**
     * Please refer to: [Path Patterns](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts#RepositoryLayouts-ModulesandPathPatternsusedbyRepositoryLayouts) in the Artifactory Wiki documentation.
     * 
     */
    @Export(name="artifactPathPattern", refs={String.class}, tree="[0]")
    private Output<String> artifactPathPattern;

    /**
     * @return Please refer to: [Path Patterns](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts#RepositoryLayouts-ModulesandPathPatternsusedbyRepositoryLayouts) in the Artifactory Wiki documentation.
     * 
     */
    public Output<String> artifactPathPattern() {
        return this.artifactPathPattern;
    }
    /**
     * Please refer to: [Descriptor Path Patterns](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts#RepositoryLayouts-DescriptorPathPatterns) in the Artifactory Wiki documentation.
     * 
     */
    @Export(name="descriptorPathPattern", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> descriptorPathPattern;

    /**
     * @return Please refer to: [Descriptor Path Patterns](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts#RepositoryLayouts-DescriptorPathPatterns) in the Artifactory Wiki documentation.
     * 
     */
    public Output<Optional<String>> descriptorPathPattern() {
        return Codegen.optional(this.descriptorPathPattern);
    }
    /**
     * When set, `descriptor_path_pattern` will be used. Default to `false`.
     * 
     */
    @Export(name="distinctiveDescriptorPathPattern", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> distinctiveDescriptorPathPattern;

    /**
     * @return When set, `descriptor_path_pattern` will be used. Default to `false`.
     * 
     */
    public Output<Optional<Boolean>> distinctiveDescriptorPathPattern() {
        return Codegen.optional(this.distinctiveDescriptorPathPattern);
    }
    /**
     * A regular expression matching the integration revision string appearing in a file name as part of the artifact&#39;s path. For example, `SNAPSHOT|(?:(?:[0-9]{8}.[0-9]{6})-(?:[0-9]+))`, in Maven. Note! Take care not to introduce any regexp capturing groups within this expression. If not applicable use `.*`
     * 
     */
    @Export(name="fileIntegrationRevisionRegexp", refs={String.class}, tree="[0]")
    private Output<String> fileIntegrationRevisionRegexp;

    /**
     * @return A regular expression matching the integration revision string appearing in a file name as part of the artifact&#39;s path. For example, `SNAPSHOT|(?:(?:[0-9]{8}.[0-9]{6})-(?:[0-9]+))`, in Maven. Note! Take care not to introduce any regexp capturing groups within this expression. If not applicable use `.*`
     * 
     */
    public Output<String> fileIntegrationRevisionRegexp() {
        return this.fileIntegrationRevisionRegexp;
    }
    /**
     * A regular expression matching the integration revision string appearing in a folder name as part of the artifact&#39;s path. For example, `SNAPSHOT`, in Maven. Note! Take care not to introduce any regexp capturing groups within this expression. If not applicable use `.*`
     * 
     */
    @Export(name="folderIntegrationRevisionRegexp", refs={String.class}, tree="[0]")
    private Output<String> folderIntegrationRevisionRegexp;

    /**
     * @return A regular expression matching the integration revision string appearing in a folder name as part of the artifact&#39;s path. For example, `SNAPSHOT`, in Maven. Note! Take care not to introduce any regexp capturing groups within this expression. If not applicable use `.*`
     * 
     */
    public Output<String> folderIntegrationRevisionRegexp() {
        return this.folderIntegrationRevisionRegexp;
    }
    /**
     * Layout name
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Layout name
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RepositoryLayout(String name) {
        this(name, RepositoryLayoutArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RepositoryLayout(String name, RepositoryLayoutArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RepositoryLayout(String name, RepositoryLayoutArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/repositoryLayout:RepositoryLayout", name, args == null ? RepositoryLayoutArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RepositoryLayout(String name, Output<String> id, @Nullable RepositoryLayoutState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/repositoryLayout:RepositoryLayout", name, state, makeResourceOptions(options, id));
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
    public static RepositoryLayout get(String name, Output<String> id, @Nullable RepositoryLayoutState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RepositoryLayout(name, id, state, options);
    }
}
