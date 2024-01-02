// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class GetPermissionTargetReleaseBundleActionsUserArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetPermissionTargetReleaseBundleActionsUserArgs Empty = new GetPermissionTargetReleaseBundleActionsUserArgs();

    /**
     * Name of the permission target.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of the permission target.
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

    private GetPermissionTargetReleaseBundleActionsUserArgs() {}

    private GetPermissionTargetReleaseBundleActionsUserArgs(GetPermissionTargetReleaseBundleActionsUserArgs $) {
        this.name = $.name;
        this.permissions = $.permissions;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPermissionTargetReleaseBundleActionsUserArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPermissionTargetReleaseBundleActionsUserArgs $;

        public Builder() {
            $ = new GetPermissionTargetReleaseBundleActionsUserArgs();
        }

        public Builder(GetPermissionTargetReleaseBundleActionsUserArgs defaults) {
            $ = new GetPermissionTargetReleaseBundleActionsUserArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Name of the permission target.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the permission target.
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

        public GetPermissionTargetReleaseBundleActionsUserArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetPermissionTargetReleaseBundleActionsUserArgs", "name");
            }
            if ($.permissions == null) {
                throw new MissingRequiredPropertyException("GetPermissionTargetReleaseBundleActionsUserArgs", "permissions");
            }
            return $;
        }
    }

}
