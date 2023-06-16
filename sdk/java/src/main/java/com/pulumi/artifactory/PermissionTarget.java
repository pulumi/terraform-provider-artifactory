// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.PermissionTargetArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.PermissionTargetState;
import com.pulumi.artifactory.outputs.PermissionTargetBuild;
import com.pulumi.artifactory.outputs.PermissionTargetReleaseBundle;
import com.pulumi.artifactory.outputs.PermissionTargetRepo;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="artifactory:index/permissionTarget:PermissionTarget")
public class PermissionTarget extends com.pulumi.resources.CustomResource {
    /**
     * As for repo but for artifactory-build-info permissions.
     * 
     */
    @Export(name="build", type=PermissionTargetBuild.class, parameters={})
    private Output</* @Nullable */ PermissionTargetBuild> build;

    /**
     * @return As for repo but for artifactory-build-info permissions.
     * 
     */
    public Output<Optional<PermissionTargetBuild>> build() {
        return Codegen.optional(this.build);
    }
    /**
     * Name of permission.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Name of permission.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * As for repo for for release-bundles permissions.
     * 
     */
    @Export(name="releaseBundle", type=PermissionTargetReleaseBundle.class, parameters={})
    private Output</* @Nullable */ PermissionTargetReleaseBundle> releaseBundle;

    /**
     * @return As for repo for for release-bundles permissions.
     * 
     */
    public Output<Optional<PermissionTargetReleaseBundle>> releaseBundle() {
        return Codegen.optional(this.releaseBundle);
    }
    /**
     * Repository permission configuration.
     * 
     */
    @Export(name="repo", type=PermissionTargetRepo.class, parameters={})
    private Output</* @Nullable */ PermissionTargetRepo> repo;

    /**
     * @return Repository permission configuration.
     * 
     */
    public Output<Optional<PermissionTargetRepo>> repo() {
        return Codegen.optional(this.repo);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PermissionTarget(String name) {
        this(name, PermissionTargetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PermissionTarget(String name, @Nullable PermissionTargetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PermissionTarget(String name, @Nullable PermissionTargetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/permissionTarget:PermissionTarget", name, args == null ? PermissionTargetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PermissionTarget(String name, Output<String> id, @Nullable PermissionTargetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/permissionTarget:PermissionTarget", name, state, makeResourceOptions(options, id));
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
    public static PermissionTarget get(String name, Output<String> id, @Nullable PermissionTargetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PermissionTarget(name, id, state, options);
    }
}
