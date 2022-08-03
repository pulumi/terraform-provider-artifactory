// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class PermissionTargetsReleaseBundleActionsGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final PermissionTargetsReleaseBundleActionsGroupArgs Empty = new PermissionTargetsReleaseBundleActionsGroupArgs();

    /**
     * Name of permission.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of permission.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    @Import(name="permissions", required=true)
    private Output<List<String>> permissions;

    public Output<List<String>> permissions() {
        return this.permissions;
    }

    private PermissionTargetsReleaseBundleActionsGroupArgs() {}

    private PermissionTargetsReleaseBundleActionsGroupArgs(PermissionTargetsReleaseBundleActionsGroupArgs $) {
        this.name = $.name;
        this.permissions = $.permissions;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PermissionTargetsReleaseBundleActionsGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PermissionTargetsReleaseBundleActionsGroupArgs $;

        public Builder() {
            $ = new PermissionTargetsReleaseBundleActionsGroupArgs();
        }

        public Builder(PermissionTargetsReleaseBundleActionsGroupArgs defaults) {
            $ = new PermissionTargetsReleaseBundleActionsGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Name of permission.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of permission.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder permissions(Output<List<String>> permissions) {
            $.permissions = permissions;
            return this;
        }

        public Builder permissions(List<String> permissions) {
            return permissions(Output.of(permissions));
        }

        public Builder permissions(String... permissions) {
            return permissions(List.of(permissions));
        }

        public PermissionTargetsReleaseBundleActionsGroupArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            $.permissions = Objects.requireNonNull($.permissions, "expected parameter 'permissions' to be non-null");
            return $;
        }
    }

}
