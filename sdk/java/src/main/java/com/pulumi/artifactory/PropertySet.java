// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.PropertySetArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.PropertySetState;
import com.pulumi.artifactory.outputs.PropertySetProperty;
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
 * Provides an Artifactory Property Set resource.
 * This resource configuration corresponds to &#39;propertySets&#39; config block in system configuration XML
 * (REST endpoint: artifactory/api/system/configuration).
 * 
 * ~&gt;The `artifactory.PropertySet` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.
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
 * import com.pulumi.artifactory.PropertySet;
 * import com.pulumi.artifactory.PropertySetArgs;
 * import com.pulumi.artifactory.inputs.PropertySetPropertyArgs;
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
 *         var foo = new PropertySet("foo", PropertySetArgs.builder()
 *             .name("property-set1")
 *             .visible(true)
 *             .properties(            
 *                 PropertySetPropertyArgs.builder()
 *                     .name("set1property1")
 *                     .predefinedValues(                    
 *                         PropertySetPropertyPredefinedValueArgs.builder()
 *                             .name("passed-QA")
 *                             .defaultValue(true)
 *                             .build(),
 *                         PropertySetPropertyPredefinedValueArgs.builder()
 *                             .name("failed-QA")
 *                             .defaultValue(false)
 *                             .build())
 *                     .closedPredefinedValues(true)
 *                     .multipleChoice(true)
 *                     .build(),
 *                 PropertySetPropertyArgs.builder()
 *                     .name("set1property2")
 *                     .predefinedValues(                    
 *                         PropertySetPropertyPredefinedValueArgs.builder()
 *                             .name("passed-QA")
 *                             .defaultValue(true)
 *                             .build(),
 *                         PropertySetPropertyPredefinedValueArgs.builder()
 *                             .name("failed-QA")
 *                             .defaultValue(false)
 *                             .build())
 *                     .closedPredefinedValues(false)
 *                     .multipleChoice(false)
 *                     .build())
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
 * Current Property Set can be imported using `property-set1` as the `ID`, e.g.
 * 
 * ```sh
 * $ pulumi import artifactory:index/propertySet:PropertySet foo property-set1
 * ```
 * 
 */
@ResourceType(type="artifactory:index/propertySet:PropertySet")
public class PropertySet extends com.pulumi.resources.CustomResource {
    /**
     * Property set name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Property set name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * A list of properties that will be part of the property set.
     * 
     */
    @Export(name="properties", refs={List.class,PropertySetProperty.class}, tree="[0,1]")
    private Output</* @Nullable */ List<PropertySetProperty>> properties;

    /**
     * @return A list of properties that will be part of the property set.
     * 
     */
    public Output<Optional<List<PropertySetProperty>>> properties() {
        return Codegen.optional(this.properties);
    }
    /**
     * Defines if the list visible and assignable to the repository or artifact. Default value is `true`.
     * 
     */
    @Export(name="visible", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> visible;

    /**
     * @return Defines if the list visible and assignable to the repository or artifact. Default value is `true`.
     * 
     */
    public Output<Boolean> visible() {
        return this.visible;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PropertySet(String name) {
        this(name, PropertySetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PropertySet(String name, @Nullable PropertySetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PropertySet(String name, @Nullable PropertySetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/propertySet:PropertySet", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private PropertySet(String name, Output<String> id, @Nullable PropertySetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/propertySet:PropertySet", name, state, makeResourceOptions(options, id));
    }

    private static PropertySetArgs makeArgs(@Nullable PropertySetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? PropertySetArgs.Empty : args;
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
    public static PropertySet get(String name, Output<String> id, @Nullable PropertySetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PropertySet(name, id, state, options);
    }
}
